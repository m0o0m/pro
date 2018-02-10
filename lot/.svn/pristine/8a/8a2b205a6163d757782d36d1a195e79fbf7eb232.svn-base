<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use common\models\UserModel;
use common\helpers\mongoTables;
use common\models\LogModel;
use common\helpers\Helper;

class UserController extends Controller {

    public function actionIndex() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            $render = [
                'data' => [],
                'line' => $this->getLines()
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
        $line_list = $this->getLines();
        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page,
            'line' => $line_list
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    public function where($get) {
        $where = ['and'];
        if (isset($get['lineId']) && !empty($get['lineId'])) {
            $where[] = 'line_id="' . trim($get['lineId']) . '"';
        }
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
        return $where;
    }

    //翻译
    public function trans($data) {
        if (is_array($data) && !empty($data)) {
            $redis = Yii::$app->redis;
            $uid_arr = array();
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
     *  会员祥情                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    // public function actionUserdetail() {
    //     $post = Yii::$app->request->post();
    //     $uid = isset($post['uid']) ? $post['uid'] : null;
    //     $addtime = isset($post['time']) ? $post['time'] : null;
    //     if (empty($uid) || (!is_numeric($uid)) || empty($addtime) || (!is_numeric($addtime))) {
    //         $result['ErrorCode'] = 2;
    //         $result['ErrorMsg'] = '参数不正确！';
    //         return json_encode($result);
    //     }

    //     $where = array();
    //     $where['uid'] = $uid;
    //     $where['addtime'] = $addtime;
    //     //获取用户信息
    //     $user_info = UserModel::getOneUser($where);
    //     if (empty($user_info)) {
    //         $result['ErrorCode'] = 2;
    //         $result['ErrorMsg'] = '用户信息不存在';
    //         return json_encode($result);
    //     }

    //     $translate = array('', 'PC', 'WAP', 'APP', '导入');
    //     $result = array();
    //     $result['ErrorCode'] = 1;
    //     $result['res']['username'] = $user_info['uname'];
    //     $result['res']['ip'] = $user_info['create_ip'];
    //     $result['res']['time'] = date('Y-m-d H:i:s', $user_info['addtime']);
    //     $result['res']['device'] = $translate[$user_info['ptype']];

    //     //查询登录日志
    //     $logTable = mongoTables::getTable('historyLogin');
    //     // $loginLog = LogModel::getLogs($logTable,array('uid'=>$uid) , 0 , 5 , array('ptype','ip','addtime'));
    //     $loginLog = LogModel::getLoginLogs($logTable, $uid, 5, array('ptype', 'ip', 'adddate'));
    //     foreach ($loginLog as $key => $val) {
    //         $result['res']['log'][$key] = $val;
    //         $result['res']['log'][$key]['addtime'] = $val['adddate'];
    //         $result['res']['log'][$key]['device'] = $translate[$val['ptype']];
    //     }

    //     return json_encode($result);
    // }

    /**
     * **********************************************************
     *  封停 启用                @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUpdatestatus() {
        $post = Yii::$app->request->post();
        $uid = isset($post['uid']) ? $post['uid'] : null;
        $addtime = isset($post['time']) ? $post['time'] : null;
        $updatetime = isset($post['updatetime']) ? $post['updatetime'] : null;
        $status = isset($post['status']) ? $post['status'] : null;
        $stop = isset($post['stop']) ? $post['stop'] : null;
        $remark = isset($post['remark']) ? $post['remark'] : null;
        $uname = isset($post['uname']) ? $post['uname'] : null;
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
            $old_status = 1;
            $status = 2;
        } elseif ($status == 2) {
            $old_status = 2;
            $status = 1;
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数不正确！';
            return json_encode($result);
        }
        //修改条件
        $where = array();
        $where['uid'] = $uid;
        $where['addtime'] = $addtime;

        //数据
        $data = array();
        $data['status'] = $status;
        $data['updatetime'] = time();
        $old_info = array(); //旧数据
        if ($status == 2) {
            $data['remark'] = $stop;
            $old_info['remark'] = null;
        } else {
            $data['remark'] = null;
            $old_info['remark'] = $remark;
        }

        //旧信息
        $old_info['uname'] = $uname;
        $old_info['updatetime'] = $updatetime;
        $old_info['status'] = $old_status;
        $old_info['uid'] = $uid;
        //新信息
        $new_info = $data;
        $new_info['uid'] = $uid;
        $new_info['uname'] = $uname;
        $res = UserModel::updateUser($data, $where);
        if ($res) {
            //LogModel::addOperateLog('user', json_encode($old_info), json_encode($new_info), $uname);
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '修改成功！';
            $mongo['type'] = "user";
            //before更新之前
            $mongo['b_uname'] = $uname;
            $mongo['b_uid'] = $uid;
            $mongo['b_remark'] = $old_info['remark'];
            $mongo['b_status'] = $old_info['status'];
            $mongo['b_updatetime'] = $old_info['updatetime'];
            //after更新之后
            $mongo['a_uname'] = $uname;
            $mongo['a_uid'] = $uid;
            $mongo['a_remark'] = $data['remark'];
            $mongo['a_status'] = $data['status'];
            $mongo['a_updatetime'] = time();
            $mongo['addday'] = date('Ymd');
            $collection = Yii::$app->mongodb->getCollection('action_info');
            $ress = $collection->insert($mongo);
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
    

}
