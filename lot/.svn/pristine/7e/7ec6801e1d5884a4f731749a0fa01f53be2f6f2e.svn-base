<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AdminModel;
use common\models\LogModel;
use common\helpers\mongoTables;

class AdminController extends Controller {

    public function actionIndex() {
        $roles = $this->getAllRole();
        $get = Yii::$app->request->get();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'roles' => $roles,
                'pagecount' => 1,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $where = ['and'];
        if (isset($get['login_user']) && !empty($get['login_user'])) {
            $where[] = 'login_user="' . trim($get['login_user']) . '"';
        }
        if (isset($get['login_name']) && !empty($get['login_name'])) {
            $where[] = 'login_name="' . trim($get['login_name']) . '"';
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where[] = 'is_delete=' . trim($get['status']);
        }
        if (isset($get['role']) && !empty($get['role'])) {
            $where[] = 'role_id=' . trim($get['role']);
        }
        //添加时间
        if (isset($get['startTime']) && !empty($get['startTime'])) {
            $where[] = ['>=', 'addtime', strtotime(trim($get['startTime']))];
        }
        if (isset($get['endTime']) && !empty($get['endTime'])) {
            $where[] = ['<=', 'addtime', strtotime(trim($get['startTime']))];
        }
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $data = AdminModel::getAdminList($where, $offset, $pagenum);
        $data = $this->trans($data, $roles);
        $count = AdminModel::getAdminCount($where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $render = [
            'data' => $data,
            'roles' => $roles,
            'pagecount' => $pagecount,
            'page' => $page
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    //翻译
    public function trans($data, $roles) {
        //角色
        $role_arr = array();
        if($roles){
            foreach($roles as $val){
                $role_arr[$val['id']] = $val['role_name'];
            }
        }

        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                //状态
                if ($v['is_delete'] == 1) {
                    $data[$k]['deleteTxt'] = '有效';
                } else if ($v['is_delete'] == 2) {
                    $data[$k]['deleteTxt'] = '无效';
                }
                //时间
                if (!empty($v['addtime'])) {
                    $data[$k]['addDate'] = date('Y-m-d H:i:s', $v['addtime']);
                }
                if (!empty($v['updatetime'])) {
                    $data[$k]['updateDate'] = date('Y-m-d H:i:s', $v['updatetime']);
                }
                //是否显示踢线
                $redis = Yii::$app->redis;
                $data[$k]['is_withdrawals'] = false;
                $is_exists = $redis->exists('ourbackend_UserOnline_' . $v['id']);
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


                //角色
                $data[$k]['role_name'] = isset($role_arr[$v['role_id']]) ? $role_arr[$v['role_id']] : '无';
            }
        }
        return $data;
    }

    //编辑
    public function actionEdit() {
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : '';
        $data = [];
        $access = [];
        if (!empty($id)) {
            $data = AdminModel::getOneAdmin(['id' => $id]);
            $data['addTimeTxt'] = !empty($data['addtime']) ? date('Y-m-d H:i:s', $data['addtime']) : '';
            $data['updateTimeTxt'] = !empty($data['updatetime']) ? date('Y-m-d H:i:s', $data['updatetime']) : '';
        }

        //所有角色
        $roles = AdminModel::getAllRole();

        $render = [
            'roles' => $roles,
            'type' => $type,
            'id' => $id,
            'data' => $data
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form.html', $render);
        } else {
            return $this->render('form.html', $render);
        }
    }

    public function actionSave() {
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $login_ip = isset($post['login_ip']) ? $post['login_ip'] : '';
        $status = isset($post['status']) ? $post['status'] : '';
        $roleId = isset($post['roleId']) ? $post['roleId'] : '';

        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if (strlen($login_user) < 4 || strlen($login_user) > 20) {
            echo json_encode(['code' => 400, 'msg' => '账号长度为4-20位!']);
            die;
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($roleId)) {
            echo json_encode(['code' => 400, 'msg' => '角色不能为空!']);
            die;
        }

        // 唯一
        $has = AdminModel::getAdminCount(['and', ['=', 'login_user', $login_user], ['<>', 'id', $id]]);
        if ($has > 0) {
            echo json_encode(['code' => 400, 'msg' => '该账号已存在!']);
            die;
        }

        $arr = [
            'login_user' => $login_user,
            'login_name' => $login_name,
            'login_ip' => $login_ip,
            'is_delete' => $status,
            'role_id' => $roleId
        ];

        if (!empty($id)) {
            $arr['updatetime'] = time();
            $res = AdminModel::updateAdmin($id, $arr);
        } else {
            $arr['addtime'] = time();
            $arr['login_pwd'] = md5(md5(123456));
            $res = AdminModel::addAdmin($arr);
        }

        if ($res) {
            echo json_encode(['code' => 200, 'msg' => '更新成功']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '更新失败']);
            die;
        }
    }

    //个人中心
    public function actionCenter() {
        $session = Yii::$app->session;
        $uid = $session['uid'];
        $logTable = mongoTables::getTable('historyLogin'); //历史登录记录
        $loginLog = LogModel::getLoginLogs($logTable, $uid, 2, array('ptype', 'ip', 'adddate'));
        $data = array();
        if ($loginLog && isset($loginLog[1])) {
            $loginLog = $loginLog[1];
            $data['ip'] = $loginLog['ip'];
            $data['adddate'] = $loginLog['adddate'];
        } else {
            $data['adddate'] = $data['ip'] = '本月是第一次登录';
        }

        $data['login_user'] = $session['login_user']; //登录账号
        $data['login_name'] = $session['login_name']; //登录账号昵称
        $data['role_name'] = $session['role_name']; //角色

        $render = ['data' => $data];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('center.html', $render);
        } else {
            return $this->render('center.html', $render);
        }
    }

    //修改密码
    public function actionUpdatepwd() {
        $post = Yii::$app->request->post();
        $oldPWd = isset($post['oldPwd']) ? $post['oldPwd'] : '';
        $newPWd = isset($post['newPwd']) ? $post['newPwd'] : '';
        $confirmPWd = isset($post['confirmPwd']) ? $post['confirmPwd'] : '';
        if (empty($oldPWd)) {
            $errorState = true;
            $errorMsg = '原始密码不能为空';
        } else if (strlen($oldPWd) > 20 || strlen($oldPWd) < 6) {
            $errorState = true;
            $errorMsg = '原始密码长度只能为6-20位!';
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $oldPWd)) {
            $errorState = true;
            $errorMsg = '原始密码只能为数字字母下划线!';
        } else if (empty($newPWd)) {
            $errorState = true;
            $errorMsg = '新密码不能为空';
        } else if (strlen($newPWd) > 20 || strlen($newPWd) < 6) {
            $errorState = true;
            $errorMsg = '新密码长度只能为6-20位!';
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $newPWd)) {
            $errorState = true;
            $errorMsg = '新密码只能为数字字母下划线!';
        } else if (empty($confirmPWd)) {
            $errorState = true;
            $errorMsg = '确认密码不能为空';
        } else if (strlen($confirmPWd) > 20 || strlen($confirmPWd) < 6) {
            $errorState = true;
            $errorMsg = '确认密码长度只能为6-20位!';
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $confirmPWd)) {
            $errorState = true;
            $errorMsg = '确认密码只能为数字字母下划线!';
        } else if ($newPWd != $confirmPWd) {
            $errorState = true;
            $errorMsg = '新密码与确认密码不一致!';
        } else {
            $errorState = false;
            $errorMsg = '';
        }

        if ($errorState) {
            echo json_encode(['code' => 400, 'msg' => $errorMsg]);
            die;
        }

        $uid = Yii::$app->session->get('uid');
        $where = [
            'id' => $uid,
            'login_pwd' => md5(md5($oldPWd))
        ];
        //验证原始密码是否正确
        $check = AdminModel::checkPwd($where);
        if (!$check) {
            echo json_encode(['code' => 400, 'msg' => '原始密码不对!']);
            die;
        }

        $updateArr = [
            'login_pwd' => md5(md5($newPWd)),
            'updatetime' => time()
        ];
        $res = AdminModel::updatePwd($uid, $updateArr);

        if ($res) {
            echo json_encode(['code' => 200, 'msg' => '修改成功']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '修改失败,请稍后再试!']);
            die;
        }
    }

/**
      ***********************************************************
      *  获取所有角色           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function getAllRole(){
        $redis = \Yii::$app->redis;
        $redis_key = 'our_backend_all_roles';
        $roles = $redis->get($redis_key);
        if($roles){
            $roles = json_decode($roles, true);
        }else{
            $roles = AdminModel::getAllRole();
            $redis->setex($redis_key, 30, json_encode($roles));
        }

        return $roles;
    }
    
/**
      ***********************************************************
      *  强制管理员或代理后台超管下线   @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionWithdrawals(){
        $post = Yii::$app->request->post();
        $uid = isset($post['uid']) ? $post['uid'] : '';
        $type = isset($post['type']) ? $post['type'] : '';

        $result = array();
        $result['ErrorCode'] = 2; 
        if( empty($uid) || !is_numeric($uid) || empty($type) ){
            $result['ErrorMsg'] = '参数不正确！';
            echo json_encode($result);die;
        }

        $redis = Yii::$app->redis;
        
        // agentbackend_UserOnline_' . $user['uid']
        // ourbackend_UserOnline_' . $user['uid']
        $key = $type . '_' . $uid;
        if(!$redis->exists($key)){
            $result['ErrorMsg'] = '该管理员已经离线';
            echo json_encode($result);die;
        }

        if(!$redis->del($key)){
            $result['ErrorMsg'] = '踢线失败';
            echo json_encode($result);die;
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '操作成功!';
        echo json_encode($result);die;
    }
    
    
}
