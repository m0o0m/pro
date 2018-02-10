<?php

use \Workerman\Worker;
use \helper\MysqlPdo as pdo;
use \helper\Common_helper;
use \helper\MongoDBManager;
use \config\config;
use \helper\IdWork;
use \helper\Encrypt;
use helper\TcpConPoll;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

Worker::$stdoutFile = __DIR__ . '/../../logs/worker.log';

$http_worker = new Worker("text://0.0.0.0:10003");

$http_worker->name = 'bet_rollback';
$http_worker->count = 8;

$http_worker->onMessage = function($connection, $data) {
    $mongo = new MongoDBManager(); //操作mongo数据库
    $lottery = pdo::instance('lottery');
    $user_tab = "my_user";
    $user_record_tab = "my_user_cash_record";
    $data = json_decode($data, true);
    $type = $data['type'];
    $bets = $data['info'];
    $bet_table = $data['bet_table'];
    $periods = $data['periods'];
    $uid = $data['uid'];

    $insert_cash_record = [];///现金记录信息
    $mycat_insert = [];     //mycat不支持批量插入而且字符类型要匹配
    $update = [];           ///对注单的修改信息
    $money_info = []; //需要扣除金额集合
    $order_num_arr = []; //订单号集合
    //现金流水表里cash_do_type字段：新增6 为彩票回滚时扣除用户已胜的金额
    $money = 0; //将金额初始为零
    $new_time = Common_helper::timechange(time());

    $userInfo = array();
    $user_sql = '';
    $user_sql = 'select * from ' . $user_tab . ' where uid=' . $uid;
    $userInfo = $lottery->row($user_sql);
    $user_money = $userInfo['money'];


     //查看是否钱包模式
    foreach ($data['info'] as $did => $bet) {
        $is_wallet = $bet['is_wallet'];
        break;
    }
       
    if($is_wallet){ // 钱包模式下请求会员余额
        // $send_info = Encrypt::GetBalance( $userInfo['uname'], "CNY");
        // $socket = TcpConPoll::getInstace();
        // $send_res = $socket::send($send_info);
        // if( !isset($socket_res['code']) || $socket_res['code'] != 1000 || !isset($socket_res['member']['balance']) ){
        //     Common_helper::write_error_log('get wallet user balance wrong' , 2, $type . '.txt');
        //     $connection->send("FALSE");
        //     return false;
        // }
        // $wallet_money = $socket_res['member']['balance'];
        $wallet_money = 2000;
    }

    
    foreach ($bets as $key => $v) {
            $line_id = $v['line_id'];
            $at_id = $v['at_id'];
            $agent_name = $v['at_username'];
            $ptype = $v['ptype'];
            $uname = $v['uname'];
            $order_num_arr[] = $v['order_num'];

            //1为未结算 2为赢 3为输 4为和局 5为无效
            switch ($v['status']) {
                case '2':
                case '4':
                case '6':
                case '7':
                    $back = $v['back']; //应扣金额
                    $money -= $back;
                    $cash_money = $money + $user_money;
                    if($is_wallet) $cash_money = $money + $wallet_money;

                    $update[] = 'update ' . "`$bet_table` set win=" . $v['win'] . ",odds=" . $v['odds'] . ",valid_bet=" . $v['bet'] . ",js=1,status=1 where periods=$periods and id=" . $v['id'] . " and at_id=" . $v['at_id'] . " and js=2"; //还原注单表

                    $mongo_update[] = array(
                        'where' => array('id'=>strval($v['id']),'at_id' => strval($v['at_id']), 'js'=>'2' ),
                        'set' => array('win' => strval($v['win']), 'odds' => strval($v['odds']), 'js' => '1', 'status' => '1', 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                    );
                    //现金记录
                    $insert_cash_record[] = getCashSql($userInfo, $v, $back, $cash_money, $type, $periods);

                    $mycat_insert[] = getMycatCashSql($userInfo, $v, $back, $cash_money, $type, $periods);

                    break;
                case '3':
                    $back = $v['back']; //应扣金额
                    $money -= $back;
                    $cash_money = $money + $user_money;
                    if($is_wallet) $cash_money = $money + $wallet_money;

                    $update[] = 'update ' . "`$bet_table` set win=" . $v['win'] . ",odds=" . $v['odds'] . ",js=1,status=1 where periods=$periods and id=". $v['id'] ." and at_id=" . $v['at_id'] . " and js=2"; //还原注单表

                    if ($v['return_water'] > 0) {
                        //现金记录
                        $insert_cash_record[] = getCashSql($userInfo, $v, $back, $cash_money, $type, $periods);

                        $mycat_insert[] = getMycatCashSql($userInfo, $v, $back, $cash_money, $type, $periods);

                    } 

                    $mongo_update[] = array(
                        'where' => array('id'=>strval($v['id']),'at_id' => strval($v['at_id']), 'js'=>'2'),
                        'set' => array('win' => $v['win'], 'odds' => strval($v['odds']), 'js' => '1', 'status' => '1', 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                    );
                    break;
                default:
                    // echo 'invalid bet!' . PHP_EOL;
                    Common_helper::write_error_log('invalid bet for rollback type=' . $type . ' qishu=' . $periods . ' bet_id=' . $v['id'], 2);
                    break;
            }
        $money_info[] = $back;
    }

    $lottery->beginTrans();

    //还原注单
    if (!empty($update)) {
        foreach ($update as $sql) {
            if (!$lottery->query($sql)) {
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }

    //同步mongodb
    if (!empty($mongo_update)) {
        foreach ($mongo_update as $val) {
            $res = $mongo->updateData(
                    'bets', $val['where'], array('$set' => $val['set'])
            );
            if(!$res){
                echo 'updata mongodb wrong';
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }

    $mycat = isset(config::$is_mycat) ? config::$is_mycat : false;

    if($mycat){
        foreach($mycat_insert as $sql){
            if (!$lottery->query($sql)) {               //写入现金记录
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }else{
        //现金流水记录
        if (!empty($insert_cash_record)) {
           //正常模式
            $cash_field = "insert into $user_record_tab ( `uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`,`uname`,`fc_type`,`periods`, `is_shiwan`) values ";
        
            if (count($insert_cash_record) != 1){
                $cash_record = $cash_field . implode(',', $insert_cash_record);
            }else{
                $cash_record = $cash_field . $insert_cash_record[0];
            }

            if (!$lottery->query($cash_record)) {               //写入现金记录
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }


    $money = array_sum($money_info);//应扣除的金额

     //钱包模式下，变更代理余额并写入现金纪录
    if($is_wallet && $money > 0){
        //查询代理当前余额
        $agent_money_sql = 'select money from my_user_agent where id=' . $at_id;
        $agent_money = $lottery->single($agent_money_sql);
        if (!$agent_money) {
            Common_helper::write_error_log('wallet get agent money wrong on rollback' . PHP_EOL . $agent_money_sql, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }

        $at_sql = "update my_user_agent set money=money+" . $money . " where id=$at_id";
        if (!$lottery->query($at_sql)) {
            Common_helper::write_error_log('wallet update agent money wrong on rollback' . PHP_EOL . $at_sql, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }
        $at_cash_num = $money;
        $at_cash_blance = $agent_money + $money;
        $at_remark = '会员:' . $userInfo['uname'] . "注单初始化返还金额,彩种:$type, 期数:$periods";
        if($mycat){
            $at_cash_sql = "insert into `lottery`.`my_agent_cash_record` (`id`, `agent_id`, `agent_user`, `line_id`, `hander_id`, `cash_num`, `cash_balance`, `remark`, `addtime`, `addday`, `cash_type`) values ( next value for MYCATSEQ_AGENTCASHRECORD, $at_id, '$agent_name', '$line_id', $uid, $at_cash_num, $at_cash_blance, '$at_remark', " . time() . ',' . date('Ymd') . ',1)';
        }else{
             $at_cash_sql = "insert into `lottery`.`my_agent_cash_record` (`agent_id`, `agent_user`, `line_id`, `hander_id`, `cash_num`, `cash_balance`, `remark`, `addtime`, `addday`, `cash_type`) values ($at_id, '$agent_name', '$line_id', $uid, $at_cash_num, $at_cash_blance, '$at_remark', " . time() . ',' . date('Ymd') . ',1)';
        }

        if(!$lottery->query($at_cash_sql)){
            Common_helper::write_error_log('wallet insert agent cash_record wrong' . PHP_EOL . $at_sql, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }

    }

    $work = new IdWork(1023);
     //修改会员余额
     if ($money > 0) {
        if($is_wallet){ //钱包模式请求派彩 派彩失败写入纪录
            $count = count($order_num_arr);
            if ($count > 1) {
                $remark_did = $order_num_arr[0] . '~' . $order_num_arr[$count - 1];
            } else {
                $remark_did = $order_num_arr[0];
            }
            $type_arr = Common_helper::getAllType('new');
            $fc_type_name = $type_arr[$type];
            $requestId = $work->nextId();
            $requestId = strval($requestId);
            $remark = '彩票注单(初始化注单):' . $remark_did . ',类型:' . $fc_type_name  . ',共计:' . $count . '单'; //Lottery 
            $send_info = Encrypt::Transfer($uname, "CNY",  0 - $money ,  $requestId ,$remark);
            // $socket = TcpConPoll::getInstace();
            // $send_res = $socket::send($send_info);
            // if(!isset($socket_res['code']) || $socket_res['code'] != 1000 ){
            //     Common_helper::write_error_log('wallet update user Money wrong' . PHP_EOL, 2, $type . '.txt');

            //     $wrong_sql = 'insert into `manage`.`my_wallet_error` (`uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `remark`, `ptype`, `addtime`, `addday`, `uname`, `fc_type`, `periods`) values' . "($uid, '$line_id', $at_id, 1, 2, '$tmp_dids', $money, '请求派彩失败', $ptype," .time(). ', ' . date('Ymd') .  ", '$uname', '$type', $qishu)";

            //     $manage->query($wrong_sql);
            //     $lottery->rollBackTrans();
            //     $connection->send("FALSE");
            //     return false;
            // }
        }else{
            $update_user = "update $user_tab set money=money-" . $money . " where uid=$uid";
            if (!$lottery->query($update_user)) {
                    Common_helper::write_error_log('update user Money wrong' . PHP_EOL . $update_user, 2, $type . '.txt');
                    $lottery->rollBackTrans();
                    $connection->send("FALSE");
                    return false;
            }
        }

    }


    $lottery->commitTrans();
    // echo 'balance end' . PHP_EOL;
    $connection->send("TRUE");
};

//获取mycat现金纪录sql
function getMycatCashSql( $userInfo, $bet, $win, $cash_balance, $type, $qishu, $is_wallet = false )
{
    $user_record_tab = "my_user_cash_record";
    $uid = $userInfo['uid'];
    $uname = $userInfo['uname'];
    $is_shiwan = $userInfo['is_shiwan'];
    $line_id = $bet['line_id'];
    $at_id = $bet['at_id'];
    $order_num = $bet['order_num'];
    $ptype = $bet['ptype'];
    $remark = "GOBACK Lottery note#:$order_num,#typesof#:#$type";
    $new_time = Common_helper::timechange(time());
    $addtime = time();

    $mycat_field = "insert into $user_record_tab ( `id` ,`uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`,`uname`,`fc_type`,`periods`,`is_shiwan`) values";

    $sql = $mycat_field . "( 'next value for MYCATSEQ_USERCASHRECORD', $uid, '$line_id', $at_id, 2, 6, $order_num, $win, $cash_balance, '$remark', $ptype,  $addtime, $new_time, '$uname', '$type', $qishu, $is_shiwan)";
    return $sql;
}

//获取现金纪录sql
function getCashSql( $userInfo, $bet, $win, $cash_balance, $type, $qishu){
    $uid = $userInfo['uid'];
    $uname = $userInfo['uname'];
    $is_shiwan = $userInfo['is_shiwan'];
    $line_id = $bet['line_id'];
    $at_id = $bet['at_id'];
    $order_num = $bet['order_num'];
    $ptype = $bet['ptype'];
    $remark = "GOBACK Lottery note#:$order_num,#typesof#:#$type";
    $new_time = Common_helper::timechange(time());
    $addtime = time();
   
   $sql = "($uid, '$line_id', $at_id, 2, 4, $order_num, $win, $cash_balance, '$remark', $ptype,  $addtime, $new_time, '$uname', '$type', $qishu, $is_shiwan)";
    return $sql;
}

// 运行worker
if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
?>
