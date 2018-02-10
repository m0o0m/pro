<?php

namespace task_sql;

use \Workerman\Worker;
use \helper\MysqlPdo as pdo;
use \helper\Common_helper;
use \helper\MongoDBManager;
use \helper\RedisConPool;
use \config\config;
use \helper\IdWork;
use \helper\Encrypt;
use helper\TcpConPoll;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

Worker::$stdoutFile = __DIR__ . '/../../logs/worker.log';

$http_worker = new Worker("text://0.0.0.0:10001");

$http_worker->name = 'Mysql_Fast_Task';
$http_worker->count = 16;

$http_worker->onMessage = function($connection, $data) {
    $lottery = pdo::instance('lottery');
    $manage = pdo::instance('manage');
    $user_tab = "my_user";
    $user_record_tab = "my_user_cash_record";
    $data = json_decode($data, true);
    $type = $data['type'];
    $qishu = $data['qishu'];
    $c_bet_tab = Common_helper::GetBetTableNameByType($type);

    $uid = $data['uid'];
    $sql = "select * from $user_tab where uid=$uid";
    $userInfo = $lottery->row($sql);

    $money = 0;             ///用户金额变动
    $win = [];              ///赢的注单id
    $fail = [];             ///输的注单id
    $insert = [];           ///现金记录信息
    $mycat_insert = [];     //mycat不支持批量插入而且字符类型要匹配
    $update = [];           ///对注单的修改信息
    $mongo_update = [];     //同步更新mongo数据表Yii::$app->mongodb->getCollection('autoLog_201708')->update(array(),array())
    $redis_data = [];       //外部接口访问时从redis拉取数据
    $mongo_msg = [];  //中奖信息插入mongo数据库
    $new_time = Common_helper::timechange(time());
    $order_num_arr = []; //订单号集合

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

    foreach ($data['info'] as $did => $bet) {
        //钱包模式下写入现金纪录的数据
        $line_id = $bet['line_id'];
        $at_id = $bet['at_id'];
        $ptype = $bet['ptype'];
        $uname = $bet['uname'];
        $agent_name = $bet['at_username'];
        $order_num_arr[] = $bet['order_num'];
        //方便外部采集接口使用 同样放入队列供统计数据使用(所有类型转成字符串)
        foreach($bet as $k=>$v){
            $tmp_redis_data[$k] = is_numeric($v) ? strval($v) : $v;
        }
        $tmp_redis_data['updatetime'] = strval(time());

        $mongo_where = array(); //mongo表更新条件
        $mongo_where = array('order_num' => strval($bet['order_num']), 'fc_type' => $type, 'at_id'=>strval($bet['at_id']), 'js'=>'1');

        switch ($bet['status']) {
            case 2:     ///赢
                $win = SlicIdByTime($bet['order_num'], $bet['addtime'], $win, $bet['at_id'], $bet['win'], $bet['odds']); //胜的注单id  $win[起始时间/结束时间/ 代理id
                $tmp = $bet['win'] + $bet['return_water'];
                $money += $tmp;
                $cash_balance = $money + $userInfo['money']; //用户现金增加
                if($is_wallet) $cash_balance = $money + $wallet_money;
                $mycat_insert[] = getMycatCashSql($userInfo, $bet, $tmp, $cash_balance, 2, $type, $qishu);

                $insert[] = getCashSql($userInfo, $bet, $tmp, $cash_balance, 2, $type, $qishu);

                $mongo_update[] = array(
                    'where' => $mongo_where,
                    'set' => array('win' => strval($bet['win']), 'js' => '2', 'status' => '2', 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                );

                $tmp_redis_data['win'] = strval($tmp);
                $tmp_redis_data['status'] = '2';
                $tmp_redis_data['js'] = '2';
                $redis_data[] = $tmp_redis_data;
                break;

            case 3:     ///输
                $tmp_win = 0 - $bet['bet'] + $bet['return_water'];
                $fail = SlicIdByTime($bet['order_num'], $bet['addtime'], $fail, $bet['at_id'], $bet['win']);
                $mongo_update[] = array(
                    'where' =>  $mongo_where,
                    'set' => array('win' => '0', 'js' => '2', 'status' => '3', 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                );
                if ($bet['return_water'] > 0) {
                    $money += $bet['return_water'];
                    $cash_balance = $money + $userInfo['money'];
                    if($is_wallet) $cash_balance = $money + $wallet_money;

                    $mycat_insert[] = getMycatCashSql($userInfo, $bet, $bet['return_water'], $cash_balance, 2, $type, $qishu);

                    $insert[] = getCashSql($userInfo, $bet, $bet['return_water'], $cash_balance, 2, $type, $qishu);
                }

                $tmp_redis_data['win'] = '0';
                $tmp_redis_data['status'] = '3';
                $tmp_redis_data['js'] = '2';
                $redis_data[] = $tmp_redis_data;
                break;

            case 4:     ///和
                $update[] = "update $c_bet_tab set win='" . $bet['win'] . "', js=2 , status=4 , valid_bet=0 , updatetime='" . time() . "',updateday='" . $new_time . "' where addday=" . date('Ymd',$bet['addtime']) . " and order_num='" . $bet['order_num'] . "' and at_id='" . $bet['at_id'] . "' and js=1"; //主要更改status为和
                $mongo_update[] = array(
                    'where' => $mongo_where,
                    'set' => array('win' => strval($bet['win']), 'js' => '2', 'status' => '4', 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                );

                $money += $bet['win'];
                $cash_balance = $money + $userInfo['money'];
                if($is_wallet) $cash_balance = $money + $wallet_money;

                $mycat_insert[] = getMycatCashSql($userInfo, $bet, $bet['bet'], $cash_balance, 3, $type, $qishu);

                $insert[] = getCashSql($userInfo, $bet, $bet['bet'], $cash_balance, 3, $type, $qishu);

                $tmp_redis_data['win'] = '0';
                $tmp_redis_data['status'] = '4';
                $tmp_redis_data['js'] = '2';
                $tmp_redis_data['valid_bet'] = '0';
                $redis_data[] = $tmp_redis_data;
                break;

            case 5:     ///多玩法赢
            case 6:
                if($bet['status'] == 5) $status = 6;
                if($bet['status'] == 6) $status = 7;
                $update[] = "update $c_bet_tab set win='" . $bet['win'] . "', odds='" . $bet['odds'] . "', js=2 , status=" . $status . ", updatetime='" . time() . "',updateday='" . $new_time . "' where addday=" . date('Ymd',$bet['addtime']) . "' and order_num='" . $bet['order_num'] . "' and at_id='" . $bet['at_id'] . "' and js=1"; //主要更改倍率

                $mongo_update[] = array(
                    'where' =>  $mongo_where,
                    'set' => array('win' => strval($bet['win']), 'odds' => strval($bet['odds']), 'js' => '2', 'status' => strval($status), 'updatetime' => strval(time()), 'updateday' => strval($new_time))
                );

                $tmp = $bet['win'] + $bet['return_water'];
                $money += $tmp;
                $cash_balance = $money + $userInfo['money'];
                if($is_wallet) $cash_balance = $money + $wallet_money;
              
                $mycat_insert[] = getMycatCashSql($userInfo, $bet, $tmp, $cash_balance, 2, $type, $qishu);

                $insert[] = getCashSql($userInfo, $bet, $tmp, $cash_balance, 2, $type, $qishu);

                $tmp_redis_data['win'] = strval($tmp);
                $tmp_redis_data['status'] = strval($status);
                $tmp_redis_data['odds'] = strval($bet['odds']);
                $tmp_redis_data['js'] = '2';
                $redis_data[] = $tmp_redis_data;
                break;
        }

    }

    //开启事务，执行sql
    $lottery->beginTrans();
   //更新胜的注单
    if (!empty($win)) {
        foreach ($win as $time_info => $id_list) {
            $time = explode('/', $time_info);
            if (count($id_list) > 1)
                $win_sql = "update $c_bet_tab set js=2,status=2,win=" . $time[2] . ', odds=' . $time[3] . ", updatetime='" . time() . "',updateday='" . $new_time . "' where addday = " . $time[0] . " and at_id='" . $time[1] . "' and js=1 and order_num in (" . implode(',', $id_list) . ")";
            else
               $win_sql = "update $c_bet_tab set js=2,status=2,win=" . $time[2] . ', odds=' . $time[3] . ", updatetime='" . time() . "',updateday='" . $new_time . "' where addday = " . $time[0] . " and at_id='" . $time[1] . "' and js=1 and order_num= " . $id_list[0];
            
            if (!$lottery->query($win_sql)) {
                    Common_helper::write_error_log('update win bets wrong' . PHP_EOL . $win, 2, $type . '.txt');
                    $lottery->rollBackTrans();
                    $connection->send("FALSE");
                    return false;
            }
        }
    }

    //更新输的注单
    if (!empty($fail)) {
        foreach ($fail as $time_info => $id_list) {
            $time = explode('/', $time_info);
            if (count($id_list) > 1)
                $fail_sql = "update $c_bet_tab set win='0',js=2, status=3, updatetime='" . time() . "',updateday='" . $new_time . "' where addday = " . $time[0] . " and at_id='" . $time[1] . "' and js=1 and order_num in (" . implode(',', $id_list) . ")";
            else
                  $fail_sql = "update $c_bet_tab set win='0',js=2, status=3, updatetime='" . time() . "',updateday='" . $new_time . "' where addday = " . $time[0] . " and at_id='" . $time[1] . "' and js=1 and order_num = " . $id_list[0];

            if (!$lottery->query($fail_sql)) {
                    Common_helper::write_error_log('update fail bets wrong' . PHP_EOL . $fail, 2, $type . '.txt');
                    $lottery->rollBackTrans();
                    $connection->send("FALSE");
                    return false;
            }
        }
    }

    //更新和或者特殊玩法的注单
    if (!empty($update)) {
        foreach($update as $other_sql){
            if (!$lottery->query($other_sql)) {
                Common_helper::write_error_log('update other_play_bets wrong' . PHP_EOL . $one, 2, $type . '.txt');
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }

    //现金流水记录
    $mycat = isset(config::$is_mycat) ? config::$is_mycat : false;
    if($mycat){
        foreach($mycat_insert as $sql){
            if (!$lottery->query($sql)) {
                Common_helper::write_error_log('insert cash_record wrong' . PHP_EOL . $sql, 2, $type . '.txt');
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }else{
        if (!empty($insert)) {

            $cash_field = "insert into $user_record_tab ( `uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`,`uname`,`fc_type`,`periods`, `is_shiwan`) values ";

            if (count($insert) != 1){
                $cash_record = $cash_field . implode(',', $insert);
            }else{
                $cash_record = $cash_field . $insert[0];
            }

            if (!$lottery->query($cash_record)) {
                Common_helper::write_error_log('insert cash_record wrong' . PHP_EOL . $cash_record, 2, $type . '.txt');
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                return false;
            }
        }
    }

    //钱包模式下，变更代理余额并写入现金纪录
    if($is_wallet && $money > 0){
        //查询代理当前余额
        $agent_money_sql = 'select money from my_user_agent where id=' . $at_id;
        $agent_money = $lottery->single($agent_money_sql);
        if (!$agent_money) {
            Common_helper::write_error_log('wallet get agent money wrong' . PHP_EOL . $agent_money_sql, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }
        //代理余额不足
        if( ($agent_money - $money) < 0 ){
            Common_helper::write_error_log('There is not enough amount of money for agent ' . $agent_name . PHP_EOL, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }

        $at_sql = "update my_user_agent set money=money-" . $money . " where id=$at_id";
        if (!$lottery->query($at_sql)) {
            Common_helper::write_error_log('wallet update agent money wrong' . PHP_EOL . $at_sql, 2, $type . '.txt');
            $lottery->rollBackTrans();
            $connection->send("FALSE");
            return false;
        }
        $at_cash_num = $money;
        $at_cash_blance = $agent_money - $money;
        $at_remark = '会员:' . $userInfo['uname'] . "派彩,彩种:$type, 期数:$qishu";
        if($mycat){
            $at_cash_sql = "insert into `lottery`.`my_agent_cash_record` (`id`, `agent_id`, `agent_user`, `line_id`, `hander_id`, `cash_num`, `cash_balance`, `remark`, `addtime`, `addday`, `cash_type`) values ( next value for MYCATSEQ_AGENTCASHRECORD, $at_id, '$agent_name', '$line_id', $uid, $at_cash_num, $at_cash_blance, '$at_remark', " . time() . ',' . date('Ymd') . ',2)';
        }else{
             $at_cash_sql = "insert into `lottery`.`my_agent_cash_record` (`agent_id`, `agent_user`, `line_id`, `hander_id`, `cash_num`, `cash_balance`, `remark`, `addtime`, `addday`, `cash_type`) values ($at_id, '$agent_name', '$line_id', $uid, $at_cash_num, $at_cash_blance, '$at_remark', " . time() . ',' . date('Ymd') . ',2)';
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
            $remark = '彩票注单(下注):' . $remark_did . ',类型:' . $fc_type_name  . ',共计:' . $count . '单'; //Lottery 
            $send_info = Encrypt::Transfer($uname, "CNY",  $money ,  $requestId ,$remark);
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
            $update_user = "update $user_tab set money=money+" . $money . " where uid=$uid";
            if (!$lottery->query($update_user)) {
                    Common_helper::write_error_log('update user Money wrong' . PHP_EOL . $update_user, 2, $type . '.txt');
                    $lottery->rollBackTrans();
                    $connection->send("FALSE");
                    return false;
            }
        }

    }


    //注单采集缓存
    if (!empty($redis_data)) {
        $pullredis = RedisConPool::getInstace('pull');
        foreach ($redis_data as $key=>$val) {
            $score =  $work->nextId();
            $score = substr($score, -16);//score最大范围16位，超出自动转科学计数
            $val['score'] = $score;
            $res = $pullredis->zadd('spider_' . $userInfo['line_id'] . '_data', $score, json_encode($val));
            if ( ($res !== 1) && ($res !== 0) ) {
                $lottery->rollBackTrans();
                $connection->send("FALSE");
                Common_helper::write_error_log('lost bet for betspider:' . PHP_EOL . json_encode($val), 3, $type . '.txt');
                return false;
            }
        }
    }
   
  
 
    if(!empty($redis_data)){
        $redis = RedisConPool::getInstace();
        foreach($redis_data as $val){
            $tmp_data = array();
            $tmp_data['fc_type'] = $val['fc_type'];
            $tmp_data['line_id'] = $val['line_id'];
            $tmp_data['sh_id'] = $val['sh_id'];
            $tmp_data['ua_id'] = $val['ua_id'];
            $tmp_data['at_id'] = $val['at_id'];
            $tmp_data['bet'] = $val['bet'];
            $tmp_data['valid_bet'] = $val['valid_bet'];
            $tmp_data['uid'] = $val['uid'];
            $tmp_data['win'] = $val['win'];
            $tmp_data['at_username'] = $val['at_username'];
            $tmp_data['uname'] = $val['uname'];
            $tmp_data['js'] = $val['js'];
            $tmp_data['status'] = $val['status'];
            $tmp_data['addday'] = $val['addday'];
            $tmp_data['updatetime'] = time();
            $redis->rpush('AccountFromRedis' , json_encode($tmp_data));
        }
    }


    
    $lottery->commitTrans();

    

    //同步mongodb  （数据量达到一百万时，超级慢）
    // if (!empty($mongo_update)) {
    //     $mongo = new MongoDBManager(); //操作mongo数据库
    //     foreach ($mongo_update as $val) {
    //         $mongo->updateData(
    //                 'bets', $val['where'], array('$set' => $val['set'])
    //         );
    //     }
    // }

    $connection->send("TRUE");
   
};

//$bet['order_num']订单号   $bet['addtime']注单时间    $win赢的注单id   $bet['win']金额//获取胜的注单id集合
function SlicIdByTime($id, $time, $id_list, $at_id, $win, $odds = '') {
    $addday = date('Ymd',$time); //索引
    $key = $addday. '/' . $at_id . '/' . $win . '/' . $odds; //下注日期 代理id 胜的金额$bet['win'] 赔率
    if (!isset($id_list[$key]))
        $id_list[$key] = [];
    $id_list[$key][] = $id; //根据每天的日期不同区分开
    return $id_list; //返回胜的注单id集合
}

//获取mycat现金纪录sql
function getMycatCashSql( $userInfo, $bet, $win, $cash_balance, $cash_do_type, $type, $qishu )
{
    $user_record_tab = "my_user_cash_record";
    $uid = $userInfo['uid'];
    $uname = $userInfo['uname'];
    $is_shiwan = $userInfo['is_shiwan'];
    $line_id = $bet['line_id'];
    $at_id = $bet['at_id'];
    $order_num = $bet['order_num'];
    $ptype = $bet['ptype'];
    $remark = "Lottery note#:$order_num,#typesof#:#$type";
    $new_time = Common_helper::timechange(time());
    $addtime = time();
    $mycat_field = "insert into $user_record_tab ( `id` ,`uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`,`uname`,`fc_type`,`periods`,`is_shiwan`) values";

    $sql = $mycat_field . "( 'next value for MYCATSEQ_USERCASHRECORD', $uid, '$line_id', $at_id, 1, $cash_do_type, $order_num, $win, $cash_balance, '$remark', $ptype,  $addtime, $new_time, '$uname', '$type', $qishu, $is_shiwan)";

    return $sql;
}

//获取现金纪录sql
function getCashSql( $userInfo, $bet, $win, $cash_balance, $cash_do_type, $type, $qishu ){
    $uid = $userInfo['uid'];
    $uname = $userInfo['uname'];
    $is_shiwan = $userInfo['is_shiwan'];
    $line_id = $bet['line_id'];
    $at_id = $bet['at_id'];
    $order_num = $bet['order_num'];
    $ptype = $bet['ptype'];
    $remark = "Lottery note#:$order_num,#typesof#:#$type";
    $new_time = Common_helper::timechange(time());
    $addtime = time();

    $sql = "($uid, '$line_id', $at_id, 1, $cash_do_type, $order_num, $win, $cash_balance, '$remark', $ptype,  $addtime, $new_time, '$uname', '$type', $qishu, $is_shiwan)";

    return $sql;
}


// 运行worker
if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
?>
