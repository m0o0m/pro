<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AdminModel;
use our_backend\models\OtherModel;

class OtherController extends Controller {

    //权限静态页面
    public function actionAccess() {
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('access.html');
        } else {
            return $this->render('access.html');
        }
    }

    public function actionTree() {
        return $this->render('tree.html');
    }

    /**
     * **********************************************************
     *  在线人数           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionOnline() {
        //在线会员
        $redis = Yii::$app->redis;
        $online = Yii::$app->view->params['online'] = $redis->hlen('userOnLine_front');
        if (!$online)
            $online = 0;
        echo $online;
        die;
    }

    /**
     * **********************************************************
     *  IP白名单列表           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionWhitelist() {
        $get = Yii::$app->request->get();
        $line_list = $this->getLines();
        $line_id = isset($get['line_id']) ? $get['line_id'] : '';
        $type = isset($get['type']) ? $get['type'] : '';
        $state = isset($get['state']) ? $get['state'] : '';

        $where = array();
        if ($line_id)
            $where['line_id'] = $line_id;
        if ($type)
            $where['type'] = $type;
        if ($state)
            $where['state'] = $state;
        $count = OtherModel::getCount($where);
        if (!$count) {
            $render = [
                'data' => array(),
                'lines' => $line_list,
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('whitelist.html', $render);
            } else {
                return $this->render('whitelist.html', $render);
            }
        }
        $data = array();
        $data = OtherModel::getList($where, 0, $count);
        $type_arr = array('', '前台api', 'golang接入api', 'golang采集api', 'golang钱包模式');
        $state_arr = array('', '启用', '停用');
        foreach ($data as $key => $val) {
            if($val['type'] == 4) $data[$key]['line_id'] = '钱包';
            $data[$key]['typeTxt'] = isset($type_arr[$val['type']]) ? $type_arr[$val['type']] : '未知';
            $data[$key]['addtime'] = date('Y-m-d H:i:s', $val['addtime']);
            $data[$key]['updatetime'] = date('Y-m-d H:i:s', $val['updatetime']);
            $data[$key]['stateTxt'] = isset($state_arr[$val['state']]) ? $state_arr[$val['state']] : '未知';
            //每行展示三个ip
            $tmp = explode(',', $val['ip']);
            $ip_str = '';
            foreach ($tmp as $k => $v) {
                $ip_str .= $v . ' &nbsp; ';
                if (($k + 1) % 3 == 0) {
                    $ip_str .= '<br/>';
                }
            }
            $data[$key]['ipTxt'] = $ip_str;
        }
        $render = [
            'data' => $data,
            'lines' => $line_list,
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('whitelist.html', $render);
        } else {
            return $this->render('whitelist.html', $render);
        }
    }

    /**
     * **********************************************************
     *  添加更新IP白名单           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionSavewhite() {
        $redis = Yii::$app->redis;
        $post = Yii::$app->request->post();
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        $ip = isset($post['ip']) ? $post['ip'] : '';
        $id = isset($post['id']) ? $post['id'] : '';
        $type = isset($post['type']) ? $post['type'] : '';
        $status = isset($post['status']) ? $post['status'] : 1;
        $remark = isset($post['remark']) ? $post['remark'] : 1;

        $result = array();
        $result['ErrorCode'] = 2;
        if (empty($line_id)) {
            $result['ErrorMsg'] = '请选择线路';
        } elseif (empty($ip)) {
            $result['ErrorMsg'] = '请输入IP';
        } elseif (empty($type)) {
            $result['ErrorMsg'] = '请选择类型';
        }

        $is_wallet = false;
        if($line_id == 'wallet' && $type != 4){
            $result['ErrorMsg'] = '钱包模式下类型必须是钱包模式';
        }elseif($type == 4 && $line_id != 'wallet'){
            $result['ErrorMsg'] = '钱包模式时线路必须选择钱包模式';
        }elseif($line_id == 'wallet' && $type == 4){
            $is_wallet = true;
        }

        if (isset($result['ErrorMsg'])) {
            echo json_encode($result);
            die;
        }

        $preg = '/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){3}$/';
        $ip_arr = explode(',', $ip);
        foreach ($ip_arr as $val) {
            if (!preg_match($preg, $val)) {
                $result['ErrorMsg'] = 'ip格式不正确';
                echo json_encode($result);
                die;
            }
        }
        if (mb_strlen($remark) > 150) {
            $result['ErrorMsg'] = '备注内容长度超出范围';
            echo json_encode($result);
            die;
        }
        //生成redis
        $redis_key = 'whitelist_' . $line_id . '_';
        if ($type == 1) {
            $redis_key .= 'front';
        } elseif ($type == 2) {
            $redis_key .= 'join';
        } elseif ($type == 3) {
            $redis_key .= 'spider';
        } elseif($type == 4){
            $redis_key = 'whitelist_wallet';
        } else {
            $result['ErrorMsg'] = 'IP类型未指定';
            echo json_encode($result);
            die;
        }

        $session = Yii::$app->session;
        $admin_name = $session['login_user'];
        $data = array();
        $data['ip'] = $ip;
        $data['state'] = $status;
        $data['remark'] = $remark;
        $data['updatetime'] = time();
        if ($id) {
            $oldInfo = OtherModel::getOneData(['id' => $id]);
            $res = OtherModel::update($data, ['id' => $id]);
            if (!$res) {
                $result['ErrorMsg'] = '更新失败！';
                echo json_encode($result);
                die;
            }
            if ($status == 1) {
                $redis->set($redis_key, json_encode(explode(',', $ip)), false);
            }
            if ($status == 2) {
                $redis->del($redis_key);
            }
            $msg = '更新成功！';

            //写入日志
            $remark = '管理员: ' . $admin_name . ' 修改了线路 ' . $line_id . '  IP白名单';
            if($is_wallet){
                $remark = '管理员: ' . $admin_name . ' 修改了钱包模式的IP白名单';
            }
            parent::insertOperateLog(json_encode($oldInfo), json_encode($data), $remark);
        } else {
            //检测纪录是否存在
            $count = OtherModel::getCount(['line_id' => $line_id, 'type' => $type]);
            if ($count) {
                $result['ErrorMsg'] = '该线路下已经有相同类型的纪录';
                if($is_wallet){
                    $result['ErrorMsg'] = '钱包模式纪录已经存在,只能更新不能再次添加';
                }
                echo json_encode($result);
                die;
            }
            $data['type'] = $type;
            $data['addtime'] = $data['updatetime'];
            $data['line_id'] = $line_id;
            $res = OtherModel::insert($data);
            if (!$res) {
                $result['ErrorMsg'] = '添加IP白名单失败';
                echo json_encode($result);
                die;
            }
            if ($status == 1) {
                $redis->set($redis_key, json_encode(explode(',', $ip)), false);
            }
            $msg = '添加成功';

            //写入日志
            $remark = '管理员: ' . $admin_name . ' 为线路 ' . $line_id . ' 添加了IP白名单';
            if($is_wallet){
                $remark = '管理员: ' . $admin_name . ' 添加了钱包模式的IP白名单';
            }
            parent::insertOperateLog('', json_encode($data), $remark);
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = $msg;
        echo json_encode($result);
        die;
    }

}
