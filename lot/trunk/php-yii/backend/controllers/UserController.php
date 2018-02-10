<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use common\models\UserModel;
use common\models\LineModel;
use common\helpers\mongoTables;
use common\models\LogModel;
use common\helpers\Helper;
use common\helpers\IdWork;

class UserController extends Controller {

    public function actionIndex() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            $render = [
                'data' => []
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $data = UserModel::getUserList($where, $offset, $pagenum);
        $data = $this->trans($data);
        // $count = UserModel::getUserCount($where);
        $count = count($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;



        foreach ($data as $key => &$val) {
            $val['old_info'] = json_encode($val);
        }

        $is_wallet = $this->check_is_wallet();
        $render = [
            'is_wallet' => $is_wallet,
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page
        ];


        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    public function where($get) {
        $where = ['and'];
        if (isset($get['userName']) && !empty($get['userName'])) {
            $where[] = 'uname="' . trim($get['userName']) . '"';
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where[] = 'status=' . trim($get['status']);
        }
        if (isset($get['shiwan']) && !empty($get['shiwan'])) {
            $where[] = 'is_shiwan=' . trim($get['shiwan']);
        }
        //添加时间
        if (isset($get['startTime']) && !empty($get['startTime'])) {
            $where[] = ['>=', 'addtime', strtotime(trim($get['startTime']))];
        }
        if (isset($get['endTime']) && !empty($get['endTime'])) {
            $where[] = ['<=', 'addtime', strtotime(trim($get['endTime']))];
        }
         //根据登录者身份进行展示
        $loginWhere = $this->loginWhere();
        if(isset($loginWhere['line_id'])){
            $where[] = "line_id='{$loginWhere['line_id']}'";
        }
        if (isset($loginWhere['agent_id'])) { //代理
            $where[] = "agent_id='{$loginWhere['agent_id']}'";
        }

        return $where;
    }

    //翻译
    public function trans($data) {
        if (is_array($data) && !empty($data)) {
            $redis = Yii::$app->redis;
            foreach ($data as $k => $v) {
                 //代理名称
                $data[$k]['at_name'] = '';
                $at_arr[] = $v['agent_id'];
                //状态
                if ($v['status'] == 1) {
                    $data[$k]['statusTxt'] = '正常';
                } else if ($v['status'] == 2) {
                    $data[$k]['statusTxt'] = '停用';
                }
                //注册设备
                $translate = array('', 'PC', 'WAP', 'APP', '导入');
                $data[$k]['dev'] = isset($translate[$v['ptype']]) ? $translate[$v['ptype']] : '';
                //是否试玩
                if($v['is_shiwan'] == 1){
                    $data[$k]['shiwan'] = '正式';
                }else{
                    $data[$k]['shiwan'] = '试玩';
                }
                //是否显示踢线
                $line_id = $v['line_id'];
                $data[$k]['is_withdrawals'] = false;
                $token= $redis->hget($line_id . '_uidToken',$v['uid']);
                $is_exists = $redis-> hexists('userOnLine_front',$token);
                if($is_exists) $data[$k]['is_withdrawals'] = true;

                $get = Yii::$app->request->get();
                $online = isset($get['online']) ? $get['online'] : null;
                if($online){
                    if($online == 1 && !$is_exists){
                        unset($data[$k]);
                        continue;
                    }
                    if($online == 2 && $is_exists){
                        unset($data[$k]);
                        continue;
                    }
                }
                //时间
                if (!empty($v['addtime'])) {
                    $data[$k]['addDate'] = date('Y-m-d H:i:s', $v['addtime']);
                }
                if (!empty($v['updatetime'])) {
                    $data[$k]['updateDate'] = date('Y-m-d H:i:s', $v['updatetime']);
                }
            }

             //查询代理信息
            $agent_data = UserModel::getAgent(['id'=>$at_arr], 0, count($at_arr));
            $at_name_arr = array();
            if($agent_data){
                foreach($agent_data as $key=>$val){
                    $at_name_arr[$val['id']] = $val['login_user'];
                }
            }
            if(!empty($agent_data)){
                foreach($data as $k=>$v){
                     $data[$k]['at_name'] = isset($at_name_arr[$v['agent_id']]) ? $at_name_arr[$v['agent_id']] : '';
                }
            }
            
        }
        return $data;
    }

     /**
     * **********************************************************
     *  修改用户                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
     public function actionUpdatepwd(){
        $post = Yii::$app->request->post();
        $uid = isset($post['uid']) ? $post['uid'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $uname = isset($post['uname']) ? $post['uname'] : '';
        $old_time = isset($post['updatetime']) ? $post['updatetime'] : '';
        $result['ErrorCode'] = 2;
        if(empty($uid) || empty($pwd)){
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }
        $data = array();
        $data['pword'] = md5(md5(123456));
        $data['updatetime'] = time();
        $res = UserModel::updateUser($data, ['uid'=>$uid]);
        if ($res) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '密码初始化成功';
            $old_info = array();
            $old_info['uid'] = $uid;
            $old_info['uname'] = $uname;
            $old_info['pword'] = $pwd;
            $old_info['updatetime'] = $old_time;

            $new_info = array();
            $new_info['uid'] = $uid;
            $new_info['uname'] = $uname;
            $new_info['pword'] = md5(md5(123456)); //新密码
            $new_info['updatetime'] = $data['updatetime']; //新更新日期

            $session = Yii::$app->session;
            $operate_uid = $session['uid'];
            $operate_name = $session['uname'];
            $operate_type = $session['user_type'];
            $arr = ['1'=> '管理员', '2'=> '管理员子帐号', '3'=> '代理', '4'=> '代理子帐号'];
            $operate_type = isset($arr[$operate_type]) ? $arr[$operate_type] : '未知管理员';
            $remark = '编号为:' . $operate_uid  . ',帐号为:' . $operate_name . '的' . $operate_type . '初始化了会员“' . $uname . '”的密码'; 
            //插入修改mongo日志
            parent::insertOperateLog(json_encode($old_info),json_encode($new_info), $remark);
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '修改失败';
        }

        return json_encode($result);
     }
    /**
     * **********************************************************
     *  额度分配                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUpdatemoney() {
        $res = array();
        $res['ErrorCode'] = 2;
        $is_wallet = $this->check_is_wallet();
        if($is_wallet){
            $res['ErrorMsg'] = '钱包模式不支持额度分配';
            return json_encode($res);
        }
        
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        $line_id = $session['line_id'];
        if($user_type != 1) parent::wrong_msg(); //只有超管能操控会员金额
        $post = Yii::$app->request->post();
        $uid = isset($post['uid']) ? $post['uid'] : null;
        $pattern = isset($post['pattern']) ? abs($post['pattern']) : null;
        $new_money = $cash_num = isset($post['money']) ? abs($post['money']) : null;
        

        if (empty($uid) || empty($pattern) || empty($new_money) || (!is_numeric($uid)) || (!is_numeric($pattern)) || (!is_numeric($new_money))) {
            $res['ErrorMsg'] = '非法金额';
            return json_encode($res);
        }

        //查询旧信息
        $old_info = UserModel::getOneUser(['uid'=>$uid]);
        if(!$old_info){
            $res['ErrorMsg'] = '获取会员信息失败';
            return json_encode($res);
        }
        //查询线路金额
        $line_money = UserModel::getLineMoney(['line_id'=>$line_id]);
        $old_line_money = $line_money;
        if(!$line_money){
            $res['ErrorMsg'] = '获取信息失败';
            return json_encode($res);
        }

        $old_money = $old_info['money'];
        
        //确定新金额
        if ($pattern == 1) { //存入
            $str = '存入';
            $cash_type = 8; 
            $user_field = 'money+' . $new_money;
            $line_field = 'money-' . $new_money;
            $user_money =  $old_money +  $new_money;
            $line_money =  $line_money - $new_money;
            if($line_money < 0){
                $res['ErrorMsg'] = '站点金额不足！';
                return json_encode($res);
            }
        } elseif ($pattern == 2) {//取出
            $str = '取出';
            $cash_type = 9; 
            $user_field = 'money-' . $new_money;
            $line_field = 'money+' . $new_money;
            $user_money =  $old_money -  $new_money;
            $line_money =  $line_money + $new_money;
            if ($user_money < 0) {
                $res['ErrorMsg'] = '会员金额不足！';
                return json_encode($res);
            }
        } else {
            $res['ErrorMsg'] = '参数不正确！';
            return json_encode($res);
        }

        $user_field = new \yii\db\Expression($user_field);
        $line_field = new \yii\db\Expression($line_field);

        
        //更新金额
        $database = \Yii::$app->db;
        $transaction = $database->beginTransaction();//开启事务
        $line_res = UserModel::updateLineMoney($line_field, ['line_id'=>$line_id]);
        if(!$line_res){
            $transaction->rollBack();
            $res['ErrorMsg'] = '更新站点金额失败';
            return json_encode($res);
        }
        $user_res = UserModel::updateUser(['money'=>$user_field, 'updatetime'=>time()], ['uid'=>$uid]);
        if(!$user_res){
            $transaction->rollBack();
            $res['ErrorMsg'] = '更新会员金额失败';
            return json_encode($res);
        }

        //写入会员现金记录
        $remark = '会员' . $old_info['uname'] . $str . '额度' . $new_money . '元,' . '原有额度' . $old_money . '元,' . '现有额度' . $user_money . '元。站点原有额度' . $old_line_money .'元,现有额度' . $line_money . '元,操作人:' . $session['login_user'];
        $work = new IdWork(1023);
        $did = $work->nextId();
        $cash = array();
        $cash['uid'] = $uid;
        $cash['line_id'] = $old_info['line_id'];
        $cash['agent_id'] = $old_info['agent_id'];
        $cash['cash_type'] = $pattern;
        $cash['cash_do_type'] = $cash_type;
        $cash['dids'] = $did;
        $cash['cash_num'] = $new_money;
        $cash['cash_balance'] = $old_money;
        $cash['remark'] = $remark;
        $cash['ptype'] = 1;
        $cash['addtime'] = time();
        $cash['addday'] = date('Ymd');
        $cash['uname'] = $old_info['uname'];
        $cash['fc_type'] = '';
        $cash['periods'] = '';
        $is_mycat = isset(Yii::$app->params['is_mycat']) ? Yii::$app->params['is_mycat'] : false;
        if($is_mycat){
            $cash_sql = "insert into `lottery`.`my_user_cash_record` (`id`, `uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`, `uname`, `fc_type`, `periods`, `is_shiwan`) values ( next value for MYCATSEQ_USERCASHRECORD, " . $uid . ", '" . $old_info['line_id'] . "', " . $old_info['agent_id'] . ", $pattern, $cash_type,$did,$new_money,$old_money," . "'" . $remark . "', 1," . time() . ", " . date('Ymd') . ", '" . $old_info['uname'] . "', '' ,0,1)";
            $cash_res = UserModel::insert($cash_sql);
        }else{
            $cash_res = UserModel::insertCashRecord($cash);
        }
        if(!$cash_res){
            $transaction->rollBack();
            $res['ErrorMsg'] = '写入会员现金记录失败';
            return json_encode($res);
        }
        $transaction->commit();

        //写入mongo日志
        parent::insertOperateLog('', '', $remark);

        $res['ErrorCode'] = 1;
        $res['ErrorMsg'] = '额度分配成功';
        return json_encode($res);
    
    }

    /**
     * **********************************************************
     *  会员祥情                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUserdetail() {
        $post = Yii::$app->request->post();
        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        $uid = isset($post['uid']) ? $post['uid'] : null;
        $addtime = isset($post['time']) ? $post['time'] : null;
        if (empty($uid) || (!is_numeric($uid)) || empty($addtime) || (!is_numeric($addtime))) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }

        $where = array();
        $where['uid'] = $uid;
        $where['addtime'] = $addtime;
        $where['line_id'] = $line_id;
        //获取用户信息
        $user_info = UserModel::getOneUser($where);
        if (empty($user_info)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '用户信息不存在';
            return json_encode($result);
        }

        $translate = array('', 'PC', 'WAP', 'APP', '导入');
        $result = array();
        $result['ErrorCode'] = 1;
        $result['res']['username'] = $user_info['uname'];
        $result['res']['ip'] = $user_info['create_ip'];
        $result['res']['time'] = date('Y-m-d H:i:s', $user_info['addtime']);
        $result['res']['device'] = $translate[$user_info['ptype']];

        //查询登录日志
        $logTable = mongoTables::getTable('historyLogin');
        $loginLog = LogModel::getLoginLogs($logTable, $uid, 5, array('ptype', 'ip', 'adddate'));
        foreach ($loginLog as $key => $val) {
            $result['res']['log'][$key] = $val;
            $result['res']['log'][$key]['addtime'] = $val['adddate'];
            $result['res']['log'][$key]['device'] = $translate[$val['ptype']];
        }

        return json_encode($result);
    }

    /**
     * **********************************************************
     *  封停 启用                @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUpdatestatus() {
        $post = Yii::$app->request->post();
        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        $uid = isset($post['uid']) ? $post['uid'] : null;
        $addtime = isset($post['time']) ? $post['time'] : null;
        $status = $old_status = isset($post['status']) ? $post['status'] : null;
        $stop = isset($post['stop']) ? $post['stop'] : null;
        $oldInfo = isset($post['oldInfo']) ? $post['oldInfo'] : null;
        $tmp_info = json_decode($oldInfo, true);
        if (empty($uid) || (!is_numeric($uid)) || empty($addtime) || (!is_numeric($addtime)) || empty($status) || (!is_numeric($status))) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }

        if ($status == 1) {
            if (empty($stop)) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '请输入封停原因';
                return json_encode($result);
            }
            $status = 2;
        } elseif ($status == 2) {
            $status = 1;
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }

        $where = array();
        $where['uid'] = $uid;
        $where['addtime'] = $addtime;
        $where['line_id'] = $line_id;


        if ($status == $tmp_info['status']) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '抱歉，您没做什么修改！';
            return json_encode($result);
        }
        $data = array();
        $result = array();
        $data['status'] = $status;
        $data['updatetime'] = time();
        if ($status == 2) {
            $data['remark'] = $stop;
            $tmp_info['remark'] = $stop;
        } else {
            $tmp_info['remark'] = null;
            $data['remark'] = null;
        }
        $tmp_info['status'] = $status;
        $res = UserModel::updateUser($data, $where);
        if ($res) {
            parent::insertOperateLog($oldInfo, json_encode($tmp_info), $tmp_info['uname']);
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '修改成功！';
            return json_encode($result);
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '修改失败';
            return json_encode($result);
        }
    }

    //更改状态
    function actionUpdatestate() {
        $post = Yii::$app->request->post();
        $id = isset($post['Id']) ? $post['Id'] : 0;
        $line_id = isset($post['Line_id']) ? $post['Line_id'] : '';
        $table = \Yii::$app->db->tablePrefix . 'user_problem';
        if ($id && $line_id) {
            $map['id'] = $id;
            $map['line_id'] = $line_id;
            $data = array(
                'state' => 1
            );
            $rs = UserModel::upState($table, $data, $map);
            if ($rs) {
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '修改成功！';
            } else {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '修改失败！';
            }
            return json_encode($result);
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }
    }

/**
      ***********************************************************
      *  踢线（强制会员下线）         @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    function actionWithdrawals(){
        $post = Yii::$app->request->post();
        $uid = isset($post['uid']) ? $post['uid'] : '';
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        $at_id = isset($post['at_id']) ? $post['at_id'] : '';
        $ids = isset($post['ids']) ? $post['ids'] : '';
        $is_batch = isset($post['is_batch']) ? $post['is_batch'] : '';
        $result = array();
        $result['ErrorCode'] = 2; 
       
        //批量操作
        if($ids && $is_batch){
            $id_arr = array();
            $tmp_arr = explode('#', $ids);
            foreach($tmp_arr as $key=>$val){
                $tmp = explode(',', $val);
                if(count($tmp) != 3) continue;
                $id_arr[$key]['line_id'] = $tmp[0];
                $id_arr[$key]['agent_id'] = $tmp[1];
                $id_arr[$key]['uid'] = $tmp[2];
            }
            if(empty($id_arr)){
                 $result['ErrorMsg'] = '参数不正确！';
                 echo json_encode($result);die;
            }

            foreach($id_arr as $val){
               $this->getout($val['line_id'], $val['agent_id'], $val['uid']);
            }
        }else{
            //单个操作
             if(empty($uid) || !is_numeric($uid) || empty($line_id) || empty($at_id)){
                $result['ErrorMsg'] = '参数不正确！';
                echo json_encode($result);die;
            }
            $this->getout($line_id, $at_id, $uid);
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '操作成功!';
        echo json_encode($result);die;

    }
    //清除redis 不返回结果
    function getout($line_id, $at_id, $uid){
        $line_key = $line_id .  '_uidToken';
        $at_key = $at_id .  '_uidToken';

        $redis = Yii::$app->redis;
        $token = $redis->hget($line_key, $uid);
        if(!$token){
            $token = $redis->hget($at_key, $uid);
        }
        // if(!$token){
        //     $result['ErrorMsg'] = '获取token失败';
        //     echo json_encode($result);die;
        // }

        $res = $redis->hdel('userOnLine_front',$token);
        if($res){
            $redis->hdel($line_key, $uid);
            $redis->hdel($at_key, $uid);
        }
        // else{
        //     $result['ErrorMsg'] = '操作失败,原因：该会员已经退出登录';
        //     echo json_encode($result);die;
        // }

    }

/**
      ***********************************************************
      *  查看线路是否钱包模式           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function check_is_wallet(){
        $lines = $this->getLines();
        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        foreach($lines as $val){
            if($val['line_id'] == $line_id && $val['type'] == 1){
                return true;
            }
        }

        return false;
    }
    

}
