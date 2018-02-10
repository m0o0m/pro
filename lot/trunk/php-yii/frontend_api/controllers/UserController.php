<?php

namespace frontend_api\controllers;

use Yii;
use frontend_api\controllers\Controller;
use frontend_api\models\UserModel;
use common\models\CashModel;
use common\models\BetModel;
use common\models\StatModel;
use common\helpers\Encrypt;
use common\helpers\TcpConPoll;
class UserController extends Controller {

    //我喜欢的,我的收藏
    public function actionLoves() {
        //所有彩种
        $alltype_data = $this->getAllFcTypes();
        //我喜欢的彩种
        $like_data = $this->getMyLovesFc();

        //数据整合
        $like = $recom = [];
        if (empty($like_data)) {
            //返回推荐彩种
            foreach ($alltype_data as $k => $v) {
                if ($v['is_recom'] == 2) {
                    $recom[] = $v;
                }
            }
        } else {
            //返回我喜欢的彩种
            $like_arr = explode(',', $like_data['fc_types']);
            foreach ($alltype_data as $k => $v) {
                if (in_array($v['type'], $like_arr)) {
                    $like[] = $v;
                }
            }
        }

        $fina = [
            'ErrorCode' => 1,
            'ErrorMsg' => '成功',
            'Data' => [
                'like' => $like,
                'recom' => $recom
            ]
        ];

        echo json_encode($fina);
        die;
    }

    //会员余额
    public function actionUserbalance() {
        $map = array();
        $line_id = $this->user->line_id;
        $map['uid'] = $this->user->uid;
        $is_wallet = $this->user->is_wallet;

        $user_data = UserModel::getUserInfo($map); //取会员信息
        if($is_wallet){
            $send_info = Encrypt::GetBalance( $user_data['uname'], "CNY");
            $socket = TcpConPoll::getInstace();
            $send_res = $socket::send($send_info);
            if( !isset($socket_res['code']) || $socket_res['code'] != 1000 || !isset($socket_res['member']['balance']) ){
                $like_data['Data'] = array();
                $like_data['ErrorCode'] = 2;
                $like_data['ErrorMsg'] = '获取余额失败';
                echo json_encode($like_data);
                die;
            }
            $user_data['money'] = $socket_res['member']['balance'];
        }else{
            if(!$user_data || !isset($user_data['money'])){
                $like_data['Data'] = array();
                $like_data['ErrorCode'] = 2;
                $like_data['ErrorMsg'] = '获取余额失败';
                echo json_encode($like_data);
                die;
            }
        }
        $new_arr = array();
        // $new_arr['uid'] = $user_data['uid'];
        $new_arr['uname'] = $user_data['uname'];
        $new_arr['money'] = $user_data['money'];
        $new_arr['currency'] = $user_data['currency'];
        $like_data['Data'] = $new_arr;
        $like_data['ErrorCode'] = 1;
        $like_data['ErrorMsg'] = '成功';
        // $redis->setex($redis_key, 5 ,'1');
        echo json_encode($like_data);
        die;
    }

    //更新我的收藏
    public function actionUpdateloves() {
        $post = Yii::$app->request->post();
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : null;
        if (empty($fc_type)) {
            $data['Data'] = [];
            $data['ErrorCode'] = 0;
            $data['ErrorMsg'] = '参数非法!';
            echo json_encode($data);
            die;
        }

        $line_id = $this->user->line_id;
        $agent_id = $this->user->agent_id;
        $uid = $this->user->uid;
        $where = [
            'line_id' => $line_id,
            'agent_id' => $agent_id,
            'uid' => $uid
        ];
        //先查是否有收藏
        $if_data = $this->getMyLovesFc();
        if (!empty($if_data)) {
            //更新
            $if_arr = explode(',', $if_data['fc_types']);
            if (in_array($fc_type, $if_arr)) {
                $index = array_search($fc_type, $if_arr);
                unset($if_arr[$index]);
                $msg = '取消收藏成功!';
            } else {
                $if_arr[] = $fc_type;
                $msg = '收藏成功!';
            }
            $update = [
                'fc_types' => implode(',', $if_arr),
                'datetime' => time()
            ];
            $res = UserModel::updateMyLoves($update, $where);
        } else {
            //新增
            $insert = [
                'line_id' => $line_id,
                'agent_id' => $agent_id,
                'uid' => $uid,
                'fc_types' => $fc_type,
                'datetime' => time()
            ];
            $res = UserModel::addMyLoves($insert);
            $msg = '收藏成功!';
        }

        if ($res) {
            $redis_key = 'loveFcType_' . $agent_id . '_' . $uid;
            Yii::$app->redis->del($redis_key);
            $data['Data'] = [];
            $data['ErrorCode'] = 1;
            $data['ErrorMsg'] = $msg;
            echo json_encode($data);
            die;
        } else {
            $data['Data'] = [];
            $data['ErrorCode'] = 0;
            $data['ErrorMsg'] = $msg;
            echo json_encode($data);
            die;
        }
    }

    //我的注单(投注记录)
    public function actionBets() {
        $request = Yii::$app->request->post();
        $bet_status = isset($request['wind']) ? trim($request['wind']) : '';
        $periods = isset($request['periods']) ? trim($request['periods']) : '';
        $fc_type = isset($request['fc_type']) ? trim($request['fc_type']) : '';
        $order_num = isset($request['order_num']) ? trim($request['order_num']) : '';
        $day = isset($request['day']) ? trim($request['day']) : '';
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $page = isset($request['page']) ? trim($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? trim($request['pagenum']) : 10;

        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        $where[] = 'and';
        $where[] = ['=', 'uid', $this->user->uid];
        $where[] = ['=', 'at_id', $this->user->agent_id];
        if ($bet_status) {
            $where[] = ['=', 'js', $bet_status];
        }
        if ($periods) {
            $where[] = ['=', 'periods', $periods];
        }
        if ($fc_type) {
            $where[] = ['=', 'fc_type', $fc_type];
        }
        if ($order_num) {
            $where[] = ['=', 'order_num', $order_num];
        }
        // 时间筛选
        if ($day) {
            // $starttime = strtotime($day);
            // $endtime = strtotime($day . ' 23:59:59');
            $day = date('Ymd',strtotime($day));
        }else{
            $day = date('Ymd');
        }
        $where[] = ['=', 'addday', $day];
        if ($starttime) {
            $where[] = ['>', 'addtime', $starttime];
        }
        if ($endtime) {
            $where[] = ['<', 'addtime', $endtime];
        }

        //前端所有彩种下拉菜单
        $type_arr = $this->getAllFcTypes();
        $menu_arr = array();
        foreach($type_arr as $key=>$val){
            $menu_arr[$key]['label'] = $val['name'];
            $menu_arr[$key]['value'] = $val['type'];
        }

        // 分页
        $recordcount = BetModel::getBetRecordCount($where);
        $pagecount = ceil($recordcount / $pagenum);
        if ($page > $pagecount && $pagecount > 0)
            $page = $pagecount;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $data = BetModel::getBetRecord($where, $offset, $limit, 'id desc');

        $data = $this->transBets($data);

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = 'OK';
        $result['Page'] = $page;
        $result['Pagenum'] = $pagenum;
        $result['Pagecount'] = $pagecount;
        $result['Recordcount'] = $recordcount;
        $result['Data'] = $data;
        $result['type_list'] = $menu_arr;
        return json_encode($result);
    }

    public function transBets($data) {
        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $v) {
            $tmp[$v['type']] = $v;
        }
        $games = $tmp;

        foreach ($data as &$v) {
            $tmp_bet_info = $this->new_bet_info($v['fc_type'], $v['bet_info']);
            $v['gameplayTxt'] = $tmp_bet_info['gameplayTxt'];
            $v['input_nameTxt'] = $tmp_bet_info['input_nameTxt'];

            if (isset($v['fc_type']))
                $v['fc_typeTxt'] = key_exists($v['fc_type'], $games) ? $games[$v['fc_type']]['name'] : $v['fc_type'];
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);

            if ($v['status'] == 1) {
                $v['win'] = '未结算';
            } elseif ($v['status'] == 5) {
                $v['win'] = '无效';
            }
            $v['can_win'] = sprintf("%.2f",$v['bet'] * $v['odds']);
        }unset($v);

        return $data;
    }

    //现金流水,交易明细
    public function actionCashs() {
        $request = Yii::$app->request->post();
        $cash_type = isset($request['cash_type']) ? trim($request['cash_type']) : '';
        $day = isset($request['day']) ? trim($request['day']) : '';
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $page = isset($request['page']) ? trim($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? trim($request['pagenum']) : 10;

        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        $where[] = 'and';
        $where[] = ['=', 'uid', $this->user->uid];
        $where[] = ['=', 'agent_id', $this->user->agent_id];
        if ($cash_type) {
            // 1=>彩票下注, 2=>彩票派彩,3=>彩票和局,4=>额度转入, 5=>额度转出, 6=>注单取消, 7=>注单无效, 
            if(in_array($cash_type, [2,3,4,7,8,9])){//现金存入
                $where[] = ['=', 'cash_type', 1];
                $where[] = ['=', 'cash_do_type', $cash_type];
            }elseif(in_array($cash_type, [1,5,6])){//现金取出
                $where[] = ['=', 'cash_type', 2];
                $where[] = ['=', 'cash_do_type', $cash_type];
            }
        }
        // 时间筛选
        if ($day) {
            $starttime = strtotime($day);
            $endtime = strtotime($day . ' 23:59:59');
        }
        if ($starttime) {
            $where[] = ['>', 'addtime', $starttime];
        }
        if ($endtime) {
            $where[] = ['<', 'addtime', $endtime];
        }

        // 分页
        $recordcount = CashModel::getCashRecordCount($where);
        $pagecount = ceil($recordcount / $pagenum);
        if ($page > $pagecount && $pagecount > 0)
            $page = $pagecount;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $data = CashModel::getCashRecord($where, $offset, $limit, 'id desc');

        $data = $this->transCashs($data);

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = 'OK';
        $result['Page'] = $page;
        $result['Pagenum'] = $pagenum;
        $result['Pagecount'] = $pagecount;
        $result['Recordcount'] = $recordcount;
        $result['Data'] = $data;
        return json_encode($result);
    }

    public function transCashs($data) {
        $cash_type = array(1=>'存入', 2=>'取出');
        $cash_do_type = array(1=>'彩票下注', 2=>'彩票派彩', 3=>'彩票和局', 4=>'额度转入', 5=>'额度转出', 6=>'注单取消', 7=>'注单无效');
        $remark_trans = array(
                    'Lottery note' => '订单号',
                    'Lottery note#' => '订单号',
                    'GOBACK Lottery note#' => '回滚订单号',
                    '#typesof#' => '彩种',
                    '#typesof#:#' => '彩种:',
                    'type' => '彩种',
                    "'" => ' ',
                    '.' => ' '
                );
        $games = $this->getAllFcTypes();
        $new_games = array();
        foreach ($games as $k => $v) {
            $new_games[$v['type']] = $v['name'];
        }
        foreach ($data as &$v) {
            if (isset($v['cash_type']))
                $v['cash_typeTxt'] = key_exists($v['cash_type'], $cash_type) ? $cash_type[$v['cash_type']] : $v['cash_type'];
            if (isset($v['cash_do_type']))
                $v['cash_do_typeTxt'] = key_exists($v['cash_do_type'], $cash_do_type) ? $cash_do_type[$v['cash_do_type']] : $v['cash_do_type'];
            if (isset($v['addtime']))
                $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            if(isset($v['remark'])){
                $tmp_remark = strtr($v['remark'], $remark_trans);
                $v['remark'] = strtr($tmp_remark, $new_games);
            } 
        }unset($v);

        return $data;
    }

    //报表统计
    public function actionReport() {
        $request = Yii::$app->request->post();
        $week = isset($request['week']) ? trim($request['week']) : 'last';
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';

        if (empty($week)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '查询时间不能为空';
            return json_encode($result);
        }

        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        $where[] = 'and';
        $where[] = ['=', 'uid', $this->user->uid];
        // 时间筛选
        if ($week) {
            $d = (date('N') - 1) * 24 * 60 * 60;
            $h = date('H') * 60 * 60;
            $i = date('i') * 60;
            $s = date('s');
            $t = $d + $h + $i + $s;
            $thisweek_starttime = time() - $t;
            if ($week == 'this') {
                $starttime = $thisweek_starttime;
            } elseif ($week == 'last') {
                $starttime = $thisweek_starttime - 7 * 24 * 60 * 60;
                $endtime = $thisweek_starttime - 1;
            }
        }
        if ($starttime) {
            $startdate = date('Ymd', $starttime);
            $where[] = ['>=', 'addday', $startdate];
        }
        if ($endtime) {
            $enddate = date('Ymd', $endtime);
            $where[] = ['<=', 'addday', $enddate];
        }

        $data = StatModel::get_list($where, '', '', 'id desc');
        $total = [];
        $tmp = [];
        foreach($data as $k => $v){
            $v['win'] = sprintf("%.2f",$v['win'] - $v['valid_bet']); // win = win - valid_bet

            @$total['bet'] += $v['bet'];
            @$total['bet_count'] += $v['bet_count'];
            @$total['valid_bet'] += $v['valid_bet'];
            @$total['valid_bet_count'] += $v['valid_bet_count'];
            @$total['win'] += $v['win'];
            @$total['win_count'] += $v['win_count'];

            @$tmp[$v['addday']]['addday'] = date('Y-m-d', strtotime($v['addday']));
            @$tmp[$v['addday']]['bet'] += $v['bet'];
            @$tmp[$v['addday']]['bet_count'] += $v['bet_count'];
            @$tmp[$v['addday']]['valid_bet'] += $v['valid_bet'];
            @$tmp[$v['addday']]['valid_bet_count'] += $v['valid_bet_count'];
            @$tmp[$v['addday']]['win'] += $v['win'];
            @$tmp[$v['addday']]['win_count'] += $v['win_count'];

            $tmp[$v['addday']]['list'][] = $v;
        }
        rsort($tmp);
        $data = $tmp;

        $data = $this->transReport($data);

        foreach($data as &$v){
            $v['bet'] = sprintf("%.2f",$v['bet']);
            $v['valid_bet'] = sprintf("%.2f",$v['valid_bet']);
            $v['win'] = sprintf("%.2f",$v['win']);
        }unset($v);
        $total['bet'] = isset($total['bet']) ? sprintf("%.2f",$total['bet']) : 0;
        $total['valid_bet'] = isset($total['valid_bet']) ? sprintf("%.2f",$total['valid_bet']) : 0;
        $total['win'] = isset($total['win']) ? sprintf("%.2f",$total['win']) : 0;

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = 'OK';
        $result['Total'] = $total;
        $result['Data'] = $data;
        return json_encode($result);
    }

    public function transReport($data) {
        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $v) {
            $tmp[$v['type']] = $v;
        }
        $games = $tmp;

        foreach ($data as &$day) {
            foreach ($day['list'] as &$v) {
                if (isset($v['fc_type']))
                    $v['fc_typeTxt'] = key_exists($v['fc_type'], $games) ? $games[$v['fc_type']]['name'] : $v['fc_type'];
                if (isset($v['addtime']))
                    $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            }unset($v);
        }

        return $data;
    }

    //公告消息
    public function actionNotice() {
        $request = Yii::$app->request->post();
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $page = isset($request['page']) ? trim($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? trim($request['pagenum']) : 10;

        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        $where[] = 'and';
        $where[] = ['or', ['=', 'line_id',$this->user->line_id], ['=', 'line_id', '']];
        $where[] = ['or', ['=', 'type',$this->user->client], ['=', 'type', '0']];
        $where[] = ['=', 'status', 1];
        // 时间筛选
        if ($starttime) {
            $where[] = ['>=', 'addtime', $starttime];
        }
        if ($endtime) {
            $where[] = ['<=', 'addtime', $endtime];
        }

        // 分页
        $recordcount = \common\models\NoticefrontModel::get_count($where);
        $pagecount = ceil($recordcount / $pagenum);
        if ($page > $pagecount && $pagecount > 0)
            $page = $pagecount;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $data = \common\models\NoticefrontModel::get_list($where, $offset, $limit, 'id desc');

        $data = $this->transNotice($data);

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = 'OK';
        $result['Page'] = $page;
        $result['Pagenum'] = $pagenum;
        $result['Pagecount'] = $pagecount;
        $result['Recordcount'] = $recordcount;
        $result['Data'] = $data;
        return json_encode($result);
    }

    public function transNotice($data) {
        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $v) {
            $tmp[$v['type']] = $v;
        }
        $games = $tmp;

        foreach ($data as &$v) {
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
        }unset($v);

        return $data;
    }

    //个人消息
    public function actionMessage() {
        $request = Yii::$app->request->post();
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $page = isset($request['page']) ? trim($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? trim($request['pagenum']) : 10;

        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        $where[] = 'and';
        $where[] = ['=', 'uid', $this->user->uid];
        // 时间筛选
        if ($starttime) {
            $where[] = ['>=', 'addtime', $starttime];
        }
        if ($endtime) {
            $where[] = ['<=', 'addtime', $endtime];
        }

        // 分页
        $recordcount = 0;
        $pagecount = ceil($recordcount / $pagenum);
        if ($page > $pagecount && $pagecount > 0)
            $page = $pagecount;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $data = [];

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = 'OK';
        $result['Page'] = $page;
        $result['Pagenum'] = $pagenum;
        $result['Pagecount'] = $pagecount;
        $result['Recordcount'] = $recordcount;
        $result['Data'] = $data;
        return json_encode($result);
    }

}
