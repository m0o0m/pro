<?php

namespace backend\controllers;

use Yii;
use yii\web\UnauthorizedHttpException;
use backend\controllers\Controller;
use backend\models\AdminModel;
use yii\mongodb\Query;
use common\models\LogModel;
use common\helpers\mongoTables;

class AdminController extends Controller {

    //管理员子帐号列表
    public function actionIndex() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if($user_type != 1){
            $this->wrong_msg();
        }
        $roles = $this->getAllRole();
        $get = Yii::$app->request->get();
        if(isset($get['_pjax']))  unset($get['_pjax']);
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
        $where = $this->where($get);
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

    public function where($get) {
        $session = Yii::$app->session;

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
        $where[] = 'line_id=' . "'{$session['line_id']}'";
        $where[] = 'user_type=2';
        $where[] = 'pid=' . $session['uid'];  
        return $where;
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
                //角色
                $data[$k]['role_name'] = isset($role_arr[$v['role_id']]) ? $role_arr[$v['role_id']] : '无';
                //账号类型(1:管理员,2:管理员子账号,3:代理账号,4:代理子账号)
                if ($v['user_type'] == 1) {
                    $data[$k]['user_type_txt'] = '管理员';
                } elseif ($v['user_type'] == 2) {
                    $data[$k]['user_type_txt'] = '管理员子账号';
                } elseif ($v['user_type'] == 3) {
                    $data[$k]['user_type_txt'] = '代理';
                } elseif ($v['user_type'] == 4) {
                    $data[$k]['user_type_txt'] = '代理子账号';
                } 
            }
        }
        return $data;
    }

   /**
      ***********************************************************
      *  管理员子帐号的新增和修改        @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionEdit(){
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $session = Yii::$app->session;
        if( $session['user_type'] != 1 ){
            $this->wrong_msg();
        }
        //所有路由
        $route_where = [];
        $route_where['is_delete'] = 1;
        $route_where['is_agent'] = 1;
        $routes = AdminModel::getAllRoute($route_where);
        $data = array();
        $access = array();
        if (!empty($id)) {
            $data = AdminModel::getOneAdmin(['id' => $id, 'line_id'=>$session['line_id']]);
            $data['addTimeTxt'] = !empty($data['addtime']) ? date('Y-m-d H:i:s', $data['addtime']) : '';
            $data['updateTimeTxt'] = !empty($data['updatetime']) ? date('Y-m-d H:i:s', $data['updatetime']) : '';
        }

        $roles = AdminModel::getAllRole(\Yii::$app->manage_db->tablePrefix . 'admin_role', [ 'is_delete'=>1, 'line_id'=>$session['line_id'] ]);
        $oneLevel = $twoLevel = $threeLevel = array();
        foreach ($routes as $k => $v) {
                if ($v['level'] == 1) {
                    $oneLevel[] = $v;
                } else if ($v['level'] == 2) {
                    $twoLevel[] = $v;
                } else if ($v['level'] == 3) {
                    $threeLevel[] = $v;
                }
        }
        $render = array();
        $render = [
            'id' => $id,
            'data' => $data,//子帐号信息
            'roles' => $roles, //角色
            'oneLevel' => $oneLevel,
            'twoLevel' => $twoLevel,
            'threeLevel' => $threeLevel
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form.html', $render);
        } else {
            return $this->render('form.html', $render);
        }
      

    }
    public function actionSave(){
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : '';
        $roleId =  isset($post['roleId']) ? $post['roleId'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $login_ip = isset($post['login_ip']) ? $post['login_ip'] : '';
        $status = isset($post['status']) ? $post['status'] : '';
        $access = isset($post['access']) ? $post['access'] : '';
        $session = Yii::$app->session;
        //只有管理员才能添加自己的子帐号
        if($session['user_type'] != 1){
            echo json_encode(['code' => 400, 'msg' => '您没有权限操作此功能!']);
            die;
        }
        //判断是否开启
        if (!$session->isActive) {
            $session->open();
        }
        $line_id = $session->get('line_id');
        $pid = $session->get('uid');
        $uname = $session->get('uname');
        $user_type = 2;//管理员子帐号
        $log_str = '子帐号';
        $session->close();

        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($roleId)) {
            echo json_encode(['code' => 400, 'msg' => '请为子帐号分配角色!']);
            die;
        }

        //账号唯一性
        $oldInfo = array();
        if (empty($id)) {
            $ifExist = AdminModel::getOneAdmin(['login_user' => $login_user]);
            if ($ifExist) {
                echo json_encode(['code' => 400, 'msg' => '该账号已存在!']);
                die;
            }
        }else{
            $oldInfo = AdminModel::getOneAdmin(['id' => $id]);
        }
        $arr = [
            'login_user' => $login_user,
            'login_name' => $login_name,
            'login_ip' => $login_ip,
            'is_delete' => $status,
            'role_id' => $roleId
        ];

        if (!empty($id)) {
            //修改
            $arr['updatetime'] = time();
            $res = AdminModel::updateAdmin($id, $arr);
        } else {
            //新增
            $arr['line_id'] = $line_id;
            $arr['user_type'] = $user_type;
            $arr['pid'] = $pid;
            $arr['addtime'] = time();
            $arr['updatetime'] = time();
            $arr['login_pwd'] = md5(md5(123456));
            $res = AdminModel::addAdmin($arr);
        }

        if ($res) {
            //mongo日志
            $old = array();
            $new = $arr;
            if(!empty($id)){
                $remark = '超级管理员:' . $uname . ' 修改了' . $log_str . ':' . $oldInfo['login_name'];
                parent::insertOperateLog(json_encode($oldInfo),json_encode($new),$remark);
            }else{//新增
                $remark = '超级管理员:' . $uname . ' 添加了' . $log_str . ':' . $arr['login_name'];
                parent::insertOperateLog('',json_encode($new),$remark);
            }

            echo json_encode(['code' => 200, 'msg' => '操作成功']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '操作失败']);
            die;
        }

    }

    //个人中心
    public function actionCenter() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        $user_type_txt = '未知';
        switch ($user_type) {
            case 1:
                $user_type_txt = '管理员';
                break;
            case 2:
                $user_type_txt = '管理员子账号';
                break;
            case 3:
                $user_type_txt = '代理';
                break;
            case 4:
                $user_type_txt = '代理子账号';
                break;
        }

        switch ($user_type) {
            case 1:
            case 2:
                $line_id = $session['line_id']; //线路id
                $tab = 'sys_line_list';
                $where = ['line_id' => $line_id];
                break;
            case 3:
            case 4:
                $agent_id = $session['agent_id']; //代理id
                $tab = 'user_agent';
                $where = ['id' => $agent_id];
                break;
        }
        $money = AdminModel::queryMoney($tab, $where);

        $data = [
            'login_user' => $session->get('login_user'),
            'login_name' => $session->get('login_name'),
            'role_name' => $session->get('role_name'),
            'user_type' => $session->get('user_type'),
            'user_type_txt' => $user_type_txt,
            'money' => $money
        ];

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
        $session = Yii::$app->session;
        $redis_key = $session['line_id']  . '_backend_all_roles';
        $roles = $redis->get($redis_key);
        if($roles){
            $roles = json_decode($roles, true);
        }else{
            $roles = AdminModel::getAllRole(\Yii::$app->manage_db->tablePrefix . 'admin_role', [ 'is_delete'=>1, 'line_id'=>$session['line_id'] ]);
            $redis->setex($redis_key, 30, json_encode($roles));
        }
        foreach($roles as $key=>$val){
            if($val['role'] == 'agent'){
                unset($roles[$key]);
            }
        }
        return $roles;
    }
}
