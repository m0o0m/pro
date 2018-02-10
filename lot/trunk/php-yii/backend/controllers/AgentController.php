<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use backend\models\AdminModel;
use common\models\AgentModel;

class AgentController extends Controller {

/**
      ***********************************************************
      *  股东列表                             *
      ***********************************************************
*/
    public function actionShindex(){
        $session = $this->get_user_type('admin');//验证权限
        $line_id = $session['line_id'];
        $utype = $session['utype'];
        $get = Yii::$app->request->get();
        if(isset($get['_pjax']))  unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'user_type' => $utype,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('sh_index.html', $render);
            } else {
                return $this->render('sh_index.html', $render);
            }
        }
        $tab_arr = self::get_tab('user_sh');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database,$tab,$where);
        $data = AgentModel::getAgentList($database,$tab,$where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'user_type' => $session['utype'],
            'data' => $data,
            'pagecount' => $pagecount,
            'user_type' => $utype,
            'page' => $page
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('sh_index.html', $render);
        } else {
            return $this->render('sh_index.html', $render);
        }
    }
/**
      ***********************************************************
      *  股东编辑页面                                       *
      ***********************************************************
*/
    public function actionSh_edit() {
        $session = $this->get_user_type('admin');//验证权限
        $get = Yii::$app->request->get();
        $line_id = $session['line_id'];
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        $tab_arr = self::get_tab('user_sh');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = [];
        $sh_data = [];
        if ($type == 'update') {
            $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id,'line_id'=>$line_id]);
            if (empty($data)) {
                return $this->redirect('/agent/shindex');
            }
        }
        // $ua_data = AgentModel::getMoreAgentByCondition(['is_delete' => 1, 'user_type' => 3, 'agent_type' => 2]);
        $ua_data = [];
        $render = [
            'ua_data' => $ua_data,
            'data' => $data,
            'sh_data' => $sh_data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('sh_form.html', $render);
        } else {
            return $this->render('sh_form.html', $render);
        }
    }
/**
      ***********************************************************
      *  股东保存新增和修改                          *
      ***********************************************************
*/

    public function actionSh_save() {
        $session = $this->get_user_type('admin');//验证权限
        $tab_arr = self::get_tab('user_sh');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = $session['line_id'];
        $uname = $session['uname'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        //1.管理员 2.管理员子帐号 3.代理 4.代理子帐号
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : 0;
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if(strlen($login_user) < 4 || strlen($login_user) > 20){
            echo json_encode(['code' => 400, 'msg' => '账号长度为4-20位!']);
            die;
        }  else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }

        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        if (!empty($id)) {
            unset($data['line_id']);
            $old = AgentModel::getOneAgentByCondition($database,$tab,array('id'=>$id));
            //修改
            $data['updatetime'] = time();
            $res = AgentModel::updateAgent($database,$tab,$data, ['id' => $id,'line_id'=>$line_id]);
        } else {
            //新增
            $is_exist = AgentModel::getAgentCount($database,$tab,array('login_user'=>$login_user));
            if($is_exist){
                echo json_encode(['code' => 400, 'msg' => '帐号已经存在!']);
                die;
            }
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $res = AgentModel::insertAgent($database,$tab,$data);
        }

        if ($res) {
             //mongo日志
            $arr = $data;
            if(!empty($id)){
                $oldInfo = $old;
                $remark = '超级管理员:' . $uname . ' 修改了股东:' . $oldInfo['login_user'];
                parent::insertOperateLog(json_encode($oldInfo),json_encode($arr),$remark);
            }else{//新增
                $remark = '超级管理员:' . $uname . ' 添加了股东:' . $arr['login_user'];
                parent::insertOperateLog('',json_encode($arr),$remark);
            }
            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }

/**
      ***********************************************************
      *  总代列表                             *
      ***********************************************************
*/
    public function actionUaindex(){
        $session =  $this->get_user_type('sh');//验证权限
        $uid = $session['uid'];
        $get = Yii::$app->request->get();
        $utype = $session['utype'];
        if(isset($get['_pjax']))  unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'user_type' => $utype,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('ua_index.html', $render);
            } else {
                return $this->render('ua_index.html', $render);
            }
        }
        $tab_arr = self::get_tab('user_ua');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $where = $this->where($get);
        if($utype != 1 && $utype != 2){
            $where['pid'] = $uid;//如果不是管理员，是股东，只能查看自己属下的总代
        }
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database,$tab,$where);
        $data = AgentModel::getAgentList($database,$tab,$where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'user_type' => $utype,
            'page' => $page
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('ua_index.html', $render);
        } else {
            return $this->render('ua_index.html', $render);
        }
    }
/**
      ***********************************************************
      *  总代编辑页面                                       *
      ***********************************************************
*/
    public function actionUa_edit() {
        $session = $this->get_user_type('sh');//验证权限
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $utype = $session['utype'];

        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';



        $tab_arr = self::get_tab('user_ua');
        $sh_table_arr = self::get_tab('user_sh');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];

        $data = [];
        $sh_data = [];
        $is_sh = false;
        $add = false;
        //判断是增加还是修改，如果是管理员，增加时显示选择股东，如果是股东只能添加到自己旗下
        if ($type == 'update') {
            $this->get_parent('ua',$id);

            $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);
            if (empty($data)) {
                return $this->redirect('/agent/uaindex');
            }
        }else{//增加
            if($utype == 1){//如果是管理员
                $sh_data = AgentModel::getMoreAgentByCondition($sh_table_arr['database'],$sh_table_arr['tab'],['line_id'=>$line_id]);
                $sh_arr = array();
                foreach($sh_data as $key=>$val){
                    $sh_arr[$key]['id'] = $val['id'];
                    $sh_arr[$key]['login_name'] = $val['login_name'];
                    $sh_arr[$key]['login_user'] = $val['login_user'];
                }
            }else{
                $add = true; //股东
            }
        }

        $render = [
            'data' => $data,
            'sh_data' => $sh_data,
            'add' => $add
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('ua_form.html', $render);
        } else {
            return $this->render('ua_form.html', $render);
        }
    }
/**
      ***********************************************************
      *  总代保存新增和修改                          *
      ***********************************************************
*/

    public function actionUa_save() {
        $session = $this->get_user_type('sh');//验证权限
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $utype = $session['utype'];
        $uname = $session['uname'];
        $tab_arr = self::get_tab('user_ua');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : 0;
        if($utype != 1)$pid = $uid;//股东操作
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if(strlen($login_user) < 4 || strlen($login_user) > 20){
            echo json_encode(['code' => 400, 'msg' => '账号长度为4-20位!']);
            die;
        }  else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }

        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        if (!empty($id)) {
            unset($data['line_id']);
            unset($data['pid']); //不能修改上级
            $oldInfo = AgentModel::getOneAgentByCondition($database,$tab,array('id'=>$id));
            //修改
            $data['updatetime'] = time();
            $res = AgentModel::updateAgent($database,$tab,$data, ['id' => $id]);
        } else {
            //新增
            $is_exist = AgentModel::getAgentCount($database,$tab,array('login_user'=>$login_user));
            if($is_exist > 0){
                echo json_encode(['code' => 400, 'msg' => '帐号已经存在!']);
                die;
            }
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $res = AgentModel::insertAgent($database,$tab,$data);
        }

        if ($res) {
            if($utype == 1) $uname_str = '超级管理员';
            if($utype == 6) $uname_str = '股东';
             //mongo日志
            $arr = $data;
            if(!empty($id)){
                $remark = $uname_str. ':' . $uname . ' 修改了总代:' . $oldInfo['login_user'];
                parent::insertOperateLog(json_encode($oldInfo),json_encode($arr),$remark);
            }else{//新增
                $remark = $uname_str. ':' . $uname . ' 添加了总代:' . $arr['login_user'];
                parent::insertOperateLog('',json_encode($arr),$remark);
            }
            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }


/**
      ***********************************************************
      *  代理列表                             *
      ***********************************************************
*/
    public function actionIndex(){
        $session = $this->get_user_type('ua');//验证权限
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $utype = $session['utype'];

        $get = Yii::$app->request->get();
        if(isset($get['_pjax']))  unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'user_type' => $utype,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $tab_arr = self::get_tab('user_agent');
        $ua_table_arr = self::get_tab('user_ua');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $where = $this->where($get);
        if($utype == 6){ // 如果是股东，查看旗下所有代理
            $ua_arr = AgentModel::getMoreAgentByCondition($ua_table_arr['database'],$ua_table_arr['tab'],['pid'=>$uid]);
            if(empty($ua_arr)){
                 $render = [
                        'data' => $data,
                        'pagecount' => $pagecount,
                        'user_type' => $utype,
                        'page' => $page
                    ];
                return $this->render('index.html', $render);

            }
            $ua_id_arr = array();
            foreach($ua_arr as $val){
                $ua_id_arr[] = $val['id'];
            }
            $where['pid'] = $ua_id_arr;
        }

        if($utype == 7)$where['pid'] = $uid;//如果是总代，只能查看自己旗下代理
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database,$tab,$where);
        $data = AgentModel::getAgentList($database,$tab,$where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'user_type' => $utype,
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
/**
      ***********************************************************
      *  代理编辑页面                                       *
      ***********************************************************
*/
    public function actionEdit() {
        $session = $this->get_user_type('ua');//验证权限
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $utype = $session['utype'];
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        $tab_arr = self::get_tab('user_agent');
        $sh_table_arr = self::get_tab('user_sh');
        $ua_table_arr = self::get_tab('user_ua');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = [];
        $sh_data = [];
        $sh_arr = array();
        $ua_arr = array();
        $add = false;
        if ($type == 'update') {
            $this->get_parent('agent',$id);
            $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);
            if (empty($data)) {
                return $this->redirect('/agent/index');
            }
        }else{
            if($utype == 1){//如果是管理员
                $sh_data = AgentModel::getMoreAgentByCondition($sh_table_arr['database'],$sh_table_arr['tab'],['line_id'=>$line_id]);
                foreach($sh_data as $key=>$val){
                    $sh_arr[$key]['id'] = $val['id'];
                    $sh_arr[$key]['login_name'] = $val['login_name'];
                    $sh_arr[$key]['login_user'] = $val['login_user'];
                }
            }elseif($utype == 6){//如果是股东
                $ua_data = AgentModel::getMoreAgentByCondition($ua_table_arr['database'],$ua_table_arr['tab'],['pid'=>$uid]);
                $ua_arr = array();
                foreach($ua_data as $key=>$val){
                    $ua_arr[$key]['id'] = $val['id'];
                    $ua_arr[$key]['login_name'] = $val['login_name'];
                    $ua_arr[$key]['login_user'] = $val['login_user'];
                }
            }elseif($utype == 7){
                $add = true;//总代
            }
        }

        $render = [
            'data' => $data,
            'add' => $add,
            'sh_data' => $sh_arr,
            'ua_data' => $ua_arr
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form.html', $render);
        } else {
            return $this->render('form.html', $render);
        }
    }
/**
      ***********************************************************
      *  代理新增和修改                          *
      ***********************************************************
*/

    public function actionSave() {
        $session = $this->get_user_type('ua');//验证权限
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $utype = $session['utype'];
        $uname = $session['uname'];
        if($utype != 1){//目前调整为只有超管能操作代理
            echo json_encode(['code' => 400, 'msg' => '您没有权限操作该功能!']);
            die;
        }
        $tab_arr = self::get_tab('user_agent');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
		$login_ip = isset($post['login_ip']) ? $post['login_ip'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if(strlen($login_user) < 4 || strlen($login_user) > 20){
            echo json_encode(['code' => 400, 'msg' => '账号长度为4-20位!']);
            die;
        }  else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }
        if(empty($pid))$pid = $uid;//总代添加的代理
        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'login_ip'   => $login_ip,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        $data['updatetime'] = time();
        //开启事务 同步my_agent_admin
        $manage_db = \Yii::$app->manage_db;
        $manage_tab = \Yii::$app->manage_db->tablePrefix . 'agent_admin';
        $transaction = $database->beginTransaction();
        if (!empty($id)) {
            $oldInfo = AgentModel::getOneAgentByCondition($database,$tab,array('id'=>$id));
            unset($data['line_id']);
            unset($data['pid']); //不能修改上级
            unset($data['login_user']); //不能修改帐号
            //修改
            $res = AgentModel::updateAgent($database,$tab,$data, ['id' => $id]);
            if(!$res) $transaction->rollBack();
            $res = AgentModel::updateAgent($manage_db,$manage_tab,$data,['login_user'=>$login_user]);
            if(!$res){
                $transaction->rollBack();
            }else{
                $transaction->commit();
            }
        } else {
            //新增
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $role_tab = \Yii::$app->manage_db->tablePrefix . 'admin_role';
            $role_id = AdminModel::getRoleId($role_tab, ['line_id'=>$line_id, 'role'=>'agent', 'is_agent'=>2]);
            if(!$role_id){
                echo json_encode(['code' => 400, 'msg' => '请先到权限管理菜单分配代理角色!']);
                die;
            }
            //帐号唯一性
            $is_exist = AgentModel::getAgentCount($database,$tab,['login_user'=>$login_user]);
            $admin_exist = AgentModel::getAgentCount($manage_db,$manage_tab,['login_user'=>$login_user]);
            if($is_exist || $admin_exist){
                echo json_encode(['code' => 400, 'msg' => '帐号已经存在!']);
                die;
            }
            $data['role_id'] = $role_id;
            //插入
            $res = AgentModel::insertAgent($database,$tab,$data);
            if(!$res) $transaction->rollBack();
            $data['user_type'] = 3;
            $res = AgentModel::insertAgent($manage_db,$manage_tab,$data);
            if(!$res){
                $transaction->rollBack();
            }else{
                $transaction->commit();
            }
        }

        if ($res) {
            //mongo日志
            $arr = $data;
            if($utype == 1) $uname_str = '超级管理员';
            if($utype == 6) $uname_str = '股东';
            if($utype == 7) $uname_str = '总代';
            if(!empty($id)){
                $remark = $uname_str .  ':' . $uname . ' 修改了代理:' . $oldInfo['login_user'];
                parent::insertOperateLog(json_encode($oldInfo),json_encode($arr),$remark);
            }else{//新增
                $remark = $uname_str .  ':'  . $uname . ' 添加了代理:' . $arr['login_user'];
                parent::insertOperateLog('',json_encode($arr),$remark);
            }

            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }
/**
      ***********************************************************
      *  代理子帐号的新增和修改        @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionEditson(){
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $other = isset($get['other']) ? $get['other'] : '';
        $session = Yii::$app->session;
        if( $session['user_type'] != 3 && $other != 'detail'){
            $this->wrong_msg();
        }
        $role_tab = \Yii::$app->manage_db->tablePrefix . 'admin_role';
        //代理角色信息
        $role_data = AdminModel::getOneRole($role_tab, ['line_id'=>$session['line_id'], 'role'=>'agent', 'is_agent'=>2]);
        if(!$role_data){
            $this->wrong_msg();
        }
        //所有路由
        $route_where = [];
        $route_where['is_delete'] = 1;
        $route_where['is_agent'] = 2;
        $routes = AdminModel::getAllRoute($route_where);
        $data = array();
        $access = array();
        if (!empty($id)) {
            $data = AdminModel::getOneAdmin(['id' => $id, 'line_id'=>$session['line_id']]);
            $data['addTimeTxt'] = !empty($data['addtime']) ? date('Y-m-d H:i:s', $data['addtime']) : '';
            $data['updateTimeTxt'] = !empty($data['updatetime']) ? date('Y-m-d H:i:s', $data['updatetime']) : '';
            $access = explode(',', $data['son_role']);
        }
        //找出代理的路由
        $agent_role = explode(',', $role_data['access_id']);
        foreach($routes as $key=>$val){
            if(!in_array($val['id'], $agent_role)){
                unset($routes[$key]);
            }
        }
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
            'access' => $access, //已经分配的权限
            'oneLevel' => $oneLevel,
            'twoLevel' => $twoLevel,
            'threeLevel' => $threeLevel
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form_son.html', $render);
        } else {
            return $this->render('form_son.html', $render);
        }
      

    }
    public function actionSaveson(){
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $login_ip = isset($post['login_ip']) ? $post['login_ip'] : '';
        $status = isset($post['status']) ? $post['status'] : '';
        $access = isset($post['access']) ? $post['access'] : '';
        $session = Yii::$app->session;
        //只有代理才能添加自己的子帐号
        if($session['user_type'] != 3){
            echo json_encode(['code' => 400, 'msg' => '您没有权限操作此功能!']);
            die;
        }
        //判断是否开启
        if (!$session->isActive) {
            $session->open();
        }
        $line_id = $session->get('line_id');
        $pid = $session->get('uid');//父id全用agent_admin表的
        $uname = $session->get('uname');
        $user_type =4;//代理子帐号
        $log_str = '代理子帐号';
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
        } elseif (empty($access)) {
            echo json_encode(['code' => 400, 'msg' => '请为子帐号分配权限!']);
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
            'son_role' => implode(',', $access)
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
                $remark = '代理:' . $uname . ' 修改了' . $log_str . ':' . $oldInfo['login_name'];
                parent::insertOperateLog(json_encode($oldInfo),json_encode($new),$remark);
            }else{//新增
                $remark = '代理:' . $uname . ' 添加了' . $log_str . ':' . $arr['login_name'];
                parent::insertOperateLog('',json_encode($new),$remark);
            }

            echo json_encode(['code' => 200, 'msg' => '操作成功']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '操作失败']);
            die;
        }

    }

    //查询条件
    public function where($get) {
        $where = [];
        if (isset($get['login_user']) && !empty($get['login_user'])) {
            $where['login_user'] = trim($get['login_user']);
        }
        if (isset($get['login_name']) && !empty($get['login_name'])) {
            $where['login_name'] =  trim($get['login_name']);
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where['is_delete'] =  trim($get['status']);
        }

        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        $where['line_id'] =  $line_id ;
        
        return $where;
    }

    //翻译
    public function trans($data) {
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
            }
        }
        return $data;
    }



    //详情页面
    public function actionDetail() {
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        if(empty($type) || !in_array($type, ['user_sh','user_ua','user_agent']) || empty($id)){
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }

        $tab_arr = self::get_tab($type);
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);

        if (empty($data)) {
            echo '<script>alert("获取信息失败"); history.back();</script>';
            exit;
        }

        if ($type == 'user_sh') {
            $data['agent_type_txt'] = '股东';
        } elseif ($type == 'user_ua') {
            $data['agent_type_txt'] = '总代';
            $this->get_parent('ua',$id);
            $ptab = \Yii::$app->db->tablePrefix . 'user_sh'; //股东表
            $parent = AgentModel::getOneAgentByCondition($database,$ptab,['id' => $data['pid']]);
        } elseif ($type == 'user_agent') {
            $this->get_parent('agent',$id);
            $ptab = \Yii::$app->db->tablePrefix . 'user_ua';//总代表
            $parent = AgentModel::getOneAgentByCondition($database,$ptab,['id' => $data['pid']]);
            $data['agent_type_txt'] = '代理';
        }

        $data['parent'] = isset($parent['login_user']) ? $parent['login_user'] : "";
        $data['parent_name'] = isset($parent['login_name']) ? $parent['login_name'] : "";

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('detail.html', ['data' => $data]);
        } else {
            return $this->render('detail.html', ['data' => $data]);
        }
    }

    //获取线路下的股东
    public function actionGetlines() {
        $post = Yii::$app->request->post();
        $tab_arr = self::get_tab('user_sh');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        }
        $where = [
            'is_delete' => 1,
            'line_id' => $line_id
        ];
        $data = AgentModel::getMoreAgentByCondition($database,$tab,$where);
        echo json_encode(['data' => $data, 'code' => 200]);
        die;
    }

    //获取股东下的总代
    public function actionGetagents() {
        $post = Yii::$app->request->post();
        $tab_arr = self::get_tab('user_ua');
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $sh_id = isset($post['sh_id']) ? $post['sh_id'] : '';
        if (empty($sh_id)) {
            echo json_encode(['code' => 400, 'msg' => '股东id不能为空!']);
            die;
        }
        $where = [
            'is_delete' => 1,
            'pid' => $sh_id,
        ];
        $data = AgentModel::getMoreAgentByCondition($database,$tab,$where);
        echo json_encode(['data' => $data, 'code' => 200]);
        die;
    }


     //分配代理角色
    public function actionRole() {
        $post = Yii::$app->request->post();
        $session = Yii::$app->session;
        $agent_id = isset($post['id']) ? $post['id'] : '';
        if (empty($agent_id)) {
            echo json_encode(['code' => 400, 'msg' => '代理不能为空']);
            die;
        }
        //获取role_id
        $role_tab = \Yii::$app->manage_db->tablePrefix . 'admin_role';
        $role_id = AdminModel::getRoleId($role_tab, ['line_id'=>$session['line_id'], 'role'=>'agent', 'is_agent'=>2]);
        if(!$role_id){
            echo json_encode(['code' => 400, 'msg' => '请先到权限管理菜单分配代理角色']);
            die;
        }
        $tab_arr = self::get_tab('user_agent');
        //入库
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $res = AgentModel::updateAgent($database,$tab,['role_id' => $role_id], ['id' => $agent_id]);

        echo json_encode(['code' => 200, 'msg' => '分配成功!']);
        die;
    }

    //代理子帐号列表
    public function actionAgent_son() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if($user_type == 4) $this->wrong_msg();//子帐号没权限查看
        $get = Yii::$app->request->get();
        if(isset($get['_pjax']))  unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'user_type' => $user_type,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('agent_son.html', $render);
            } else {
                return $this->render('agent_son.html', $render);
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
        //添加时间
        if (isset($get['startTime']) && !empty($get['startTime'])) {
            $where[] = ['>=', 'addtime', strtotime(trim($get['startTime']))];
        }
        if (isset($get['endTime']) && !empty($get['endTime'])) {
            $where[] = ['<=', 'addtime', strtotime(trim($get['startTime']))];
        }
        //权限
        $tmp_where = $this->loginWhere();
        if(isset($tmp_where['line_id'])){
            $where[] = 'line_id=' . "'{$tmp_where['line_id']}'";
        }
        if(isset($tmp_where['agent_id'])){
            $where[] = 'pid=' . $session['uid']; //关联的是agent_amdin表 
        }

        $where[] = 'user_type=4';//代理子帐号
        //分页
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $data = AdminModel::getAdminList($where, $offset, $pagenum);
        $data = $this->trans($data);
        $count = AdminModel::getAdminCount($where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'user_type' => $user_type,
            'page' => $page
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('agent_son.html', $render);
        } else {
            return $this->render('agent_son.html', $render);
        }
    }


    //密码表单
    public function actionPassword() {
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : null;
        if(empty($type) || !in_array($type, ['user_sh','user_ua','user_agent']) || empty($id)){
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }
        if($type == 'user_ua'){
            $this->get_parent('ua',$id);
        }elseif($type == 'user_agent'){
            $this->get_parent('agent',$id);
        }
        $tab_arr = self::get_tab($type);
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);

        if (empty($data)) {
            return $this->redirect('/agent/index');
        }

        $render = [
            'data' => $data
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('password.html', $render);
        } else {
            return $this->render('password.html', $render);
        }
    }
    //密码修改(入库)
    public function actionSavepwd() {
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $type = isset($post['type']) ? $post['type'] : null;
        $login_user = isset($post['login_user']) ? $post['login_user'] : null;
        if(empty($type) || !in_array($type, ['user_sh','user_ua','user_agent'])){
            echo json_encode(['code' => 400, 'msg' => '缺失重要参数!']);
            die;
        }
        if (empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '代理不能为空!']);
            die;
        } else if (empty($pwd)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }
        $tab_arr = self::get_tab($type);
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $manage_db = \Yii::$app->manage_db;
        $manage_tab = \Yii::$app->manage_db->tablePrefix . 'agent_admin';
        if($type == 'user_agent'){
            $transaction = $database->beginTransaction();
            $res = AgentModel::updateAgent($database,$tab,['login_pwd' => md5(md5($pwd))], ['id' => $id]);
            if(!$res) $transaction->rollBack();
            $res = AgentModel::updateAgent($manage_db,$manage_tab,['login_pwd' => md5(md5($pwd))], ['login_user' => $login_user]);
            if(!$res){
                $transaction->rollBack();
            }else{
                $transaction->commit();
            }
        }else{
            $res = AgentModel::updateAgent($database,$tab,['login_pwd' => md5(md5($pwd))], ['id' => $id]);
        }

        echo json_encode(['code' => 200, 'msg' => '保存成功!']);
        die;
    }

 /**
       ***********************************************************
       *  额度分配表单与入库           @author ruizuo qiyongsheng    *
       ***********************************************************
 */
    public function actionMoney() {
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : null;
        if(empty($type) || !in_array($type, ['user_sh','user_ua','user_agent']) || empty($id)){
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }
        $session = $this->get_user_type('ua');
        $line_id = $session['line_id'];
        $utype = $session['utype'];
        $uid = $session['uid'];

        if($type == 'user_sh'){
            $this->get_parent('admin',$id);//只有管理员能给股东分配额度
        }elseif($type == 'user_ua'){
            $this->get_parent('ua',$id);//验证是否是股东下的总代
        }elseif($type == 'user_agent'){
            $this->get_parent('agent',$id);//验证是否是股东或总代下的代理
        }

        $tab_arr = self::get_tab($type);
        $database  = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);
        if (empty($data)) {
           echo '<script>alert("获取信息失败"); history.back();</script>';
            exit;
        }
        //查看管理员是否和操作的下属是同一条线路
        if($line_id != $data['line_id']) $this->wrong_msg();

        $render = [
            'data' => $data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('money.html', $render);
        } else {
            return $this->render('money.html', $render);
        }
    }

    //额度分配 入库
    public function actionSetmoney() {
        $session = $this->get_user_type('ua');
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $agent_type = isset($post['type']) ? $post['agent_type'] : 0;
        $type = isset($post['agent_type']) ? $post['type'] : 0;
        $money = isset($post['money']) ? $post['money'] : 0;
        if (empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '操作对象不能为空!']);
            die;
        } else if (!in_array($type, [1, 2])) {
            echo json_encode(['code' => 400, 'msg' => '交易模式错误!']);
            die;
        } else if (empty($money)) {
            echo json_encode(['code' => 400, 'msg' => '交易金额不能为空!']);
            die;
        } else if (!preg_match('/^[+]{0,1}(\d+)$|^[+]{0,1}(\d+\.\d+)$/', $money)) {
            echo json_encode(['code' => 400, 'msg' => '交易金额只能为正数!']);
        }elseif(empty($agent_type) || !in_array($agent_type, ['user_sh','user_ua','user_agent'])){
            echo json_encode(['code' => 400, 'msg' => '操作对象错误!']);
            die;
        }

        $user_type = $session['utype'];
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $tab_arr = self::get_tab($agent_type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];

        //查询操作对象的信息
        $data = AgentModel::getOneAgentByCondition($database,$tab,['id' => $id]);
        if(empty($data)){
            $transaction->rollBack();
            return json_encode(['code' => 400, 'msg' => '获取操作对象信息失败!']);
        }
        if($data['line_id'] != $line_id) $this->wrong_msg();


         //验证权限
        if($agent_type == 'user_sh'){//操作对象为股东
           if($user_type != 1) $this->wrong_msg();
        }elseif($agent_type == 'user_ua'){//登录身份为股东，操作对象为总代
            if($user_type == 6) $this->get_parent('ua',$id);
        }elseif($agent_type == 'user_agent'){//操作对象为代理
            if($user_type == 6 || $user_type == 7) $this->get_parent('agent',$id);
        }
        //查询操作者的信息
        if($user_type == 1){
             $info = AgentModel::getlineinfo($line_id);
             $log_user = '线路:' . $line_id;
        }elseif($user_type == 6){
             $tabInfo = self::get_tab('user_sh');
             $info = AgentModel::getOneAgentByCondition($tabInfo['database'],$tabInfo['tab'],['id'=>$uid]);
             $log_user = '股东:' . $info['login_name'];
        }elseif($user_type == 7){
             $tabInfo = self::get_tab('user_ua');
             $info = AgentModel::getOneAgentByCondition($tabInfo['database'],$tabInfo['tab'],['id'=>$uid]);
             $log_user = '代理:' . $info['login_name'];
        }else{
            $this->wrong_msg();
        }

        if(empty($info)){
            return ['code' => 400, 'msg' => '获取操作者信息失败!'];
        }
        if($type == 1){//存入
            if($info['money'] < $money) return json_encode(['code' => 400, 'msg' => '操作者余额不足!']);
        }elseif($type == 2){//取出
            if($data['money'] < $money) return json_encode(['code' => 400, 'msg' => '操作者余额不足!']);
        }

        //开启事务
        $transaction = $database->beginTransaction();
        try {
            if($type == 1){//存入
                $new_money_field_a = 'money-' . $money;
                $new_money_field = 'money+' . $money;
                $new_money_a = $info['money'] - $money;//操作者
                $new_money = $data['money'] + $money;//操作对象
            }
            if($type == 2){//取出
                $new_money_field_a = 'money+' . $money;
                $new_money_field = 'money-' . $money;
                $new_money_a = $info['money'] + $money;
                $new_money = $data['money'] - $money;
            }

            $new_money_field_a = new \yii\db\Expression($new_money_field_a);
            $new_money_field = new \yii\db\Expression($new_money_field);
            //操作数据库
            if($user_type == 1){//管理员操作，变更线路金额
                $res_a = AgentModel::updatelinemoney('money',$new_money_field_a,['line_id'=>$line_id]);
                if(!$res_a){
                    $transaction->rollBack();
                    return json_encode(['code' => 400, 'msg' => '更新上级金额失败!']);
                }

            }else{//变更股东或总代金额
                $res_a = AgentModel::updateAgent($tabInfo['database'],$tabInfo['tab'],['money'=>$new_money_field_a],['id'=>$uid]);
                if(!$res_a){
                    $transaction->rollBack();
                    return json_encode(['code' => 400, 'msg' => '更新上级金额失败!']);
                }
            }
            $res = AgentModel::updateAgent($database,$tab,['money'=>$new_money_field],['id'=>$id]);
            if(!$res){
                $transaction->rollBack();
                return json_encode(['code' => 400, 'msg' => '更新金额失败!']);
            }
            //写入现金记录
             //股东数据和总代数据
            $user_name = $session['uname'];
            if($type == 1){
                $cash_str_a = '取出';
                $cash_str_b = '存入';
            }elseif($type == 2){
                $cash_str_a = '存入';
                $cash_str_b = '取出';
            }
            $other_insert = [
                'line_id' => $line_id,
                'cash_num' => $money,
                'addtime' => time(),
                'addday' => date('Ymd'),
                'cash_type' => $type
            ];

            $cash_database = \Yii::$app->db;
                //操作对象
            if($agent_type == 'user_sh'){//操作对象为股东
                $cash_tab_sub = Yii::$app->db->tablePrefix . 'sh_cash_record';
                $insert = $other_insert;
                $insert['sh_id'] = $data['id'];
                $insert['sh_user'] = $data['login_user'];
                $log_name = '股东';
            }elseif($agent_type == 'user_ua'){//操作对象为总代
                $cash_tab_sub = Yii::$app->db->tablePrefix . 'ua_cash_record';
                $insert = $other_insert;
                $insert['ua_id'] = $data['id'];
                $insert['ua_user'] = $data['login_user'];
                $log_name = '总代';
            }elseif($agent_type == 'user_agent'){//操作对象为代理
                $cash_tab_sub = Yii::$app->db->tablePrefix . 'agent_cash_record';
                $insert = $other_insert;
                $insert['agent_id'] = $data['id'];
                $insert['agent_user'] = $data['login_user'];
                if($type == 1){
                    $insert['cash_balance'] = $data['money'] + $money;
                }else{
                    $insert['cash_balance'] = $data['money'] - $money;
                }
                $log_name = '代理';
            }
            $insert['remark'] = $cash_str_b . '额度' . $money . '元,操作人:' . $user_name;

            //操作者
            if($user_type == 1){
                $cash_tab = Yii::$app->db->tablePrefix . 'line_cash_record';
                 //线路数据
                $insert_a = [
                    'line_id' => $line_id,
                    'cash_num' => $money,
                    'cash_balance' => $new_money_a,
                    'remark' =>  $cash_str_a . '额度' . $money . '元,操作人:' . $user_name,
                    'addtime' => time(),
                    'addday' => date('Ymd'),
                    'cash_type' => $type
                ];
            }elseif($user_type == 6){
                $cash_tab = Yii::$app->db->tablePrefix . 'sh_cash_record';
                $insert_a = $other_insert;
                $insert_a['sh_id'] = $uid;
                $insert_a['sh_user'] = $user_name;

            }elseif($user_type == 7){
                $cash_tab = Yii::$app->db->tablePrefix . 'ua_cash_record';
                $insert_a = $other_insert;
                $insert_a['ua_id'] = $uid;
                $insert_a['ua_user'] = $user_name;
            }
            $insert_a['remark'] = $cash_str_a . '额度' . $money . '元,操作人:' . $user_name;
            if($type == 1){
                $insert_a['cash_type'] = 2;
            }else{
                $insert_a['cash_type'] = 1;
            }

            //插入数据库
            $res = AgentModel::insertAgent($cash_database,$cash_tab,$insert_a);
            if(!$res){
                $transaction->rollBack();
                return json_encode(['code' => 400, 'msg' => '记录操作者现金记录失败!']);
            }
            $res2 = AgentModel::insertAgent($cash_database,$cash_tab_sub,$insert);
             if(!$res2){
                $transaction->rollBack();
                return json_encode(['code' => 400, 'msg' => '记录操作对象现金记录失败!']);
            }
            //mongo日志
            if($type == 1){
                $remark = $log_user . '取出金额' . $money .'分配给' . $log_name . $data['login_name'] . ',' . $log_user . '原有金额是' . $info['money'] . ',当前金额是' . $new_money_a . ',' .$log_name . '原有金额是' .$data['money']. ',当前金额是' . $new_money . ',操作日期' . date('Y-m-d',time()) . ',操作人:' . $user_name;
            }else{
                $remark = $log_user . '从'  . $log_name . $data['login_name'] .  '中取出金额' . $money . ',' . $log_user  .'原有金额是'.$info['money'] .',当前金额是,' . $new_money_a . ',' .$log_name . '原有金额是' .$data['money']. ',当前金额是'. $new_money . ',操作日期' . date('Y-m-d',time()) . ',操作人:' . $user_name;
            }
            parent::insertOperateLog(json_encode($info),json_encode($data),$remark);

            $transaction->commit();
            return json_encode(['code' => 200, 'msg' => '额度分配成功!']);

        } catch (\Exception $e) {
            $transaction->rollBack();
            return json_encode(['code' => 400, 'msg' => $e->getMessage()]);
        }

    }



/**
      ***********************************************************
      *  获取数据库及表名           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function get_tab($type = 'user_sh'){
        $data = array();
        if($type == 'user_sh'){ //股东
             $database  = \Yii::$app->db;
             $tab = \Yii::$app->db->tablePrefix . 'user_sh';
        }elseif($type == 'user_ua'){//总代
            $database  = \Yii::$app->db;
            $tab = \Yii::$app->db->tablePrefix . 'user_ua';
        }elseif($type == 'user_agent'){//代理
            $database  = \Yii::$app->db;
            $tab = \Yii::$app->db->tablePrefix . 'user_agent';
        }
        $data['database'] = $database;
        $data['tab'] = $tab;
        return $data;
    }


    //从session中获取user_type判断登录者的身份
    public function get_user_type($type = ''){
        $session = Yii::$app->session;
        $user_type =  $session['user_type'];
        if($type == 'admin'){//管理员及管理员子帐号
            $arr = [1,2];
        }elseif($type == 'sh'){ //管理员及管理员子帐号 股东
            $arr = [1,2,6];
        }elseif($type == 'ua'){//管理员及管理员子帐号 股东 总代
            $arr = [1,2,6,7];
        }else{
            $arr = [];
        }

        if(!in_array($user_type, $arr)){
            if (Yii::$app->request->isAjax){
                 $this->wrong_msg();
            }else{
                 $this->wrong_msg();
            }
        }else{
            $new_arr = array();
            $new_arr['uid'] = $session['uid'];
            $new_arr['line_id'] = $session['line_id'];
            $new_arr['utype'] = $user_type;
            $new_arr['uname'] = $session['uname'];
            return $new_arr;
        }

    }
/**
      ***********************************************************
      *  查询代理或股东的上级，验证操作是否合法,防止篡改地址栏            *
      ***********************************************************
*/
    public function get_parent($type,$id){
        $session = Yii::$app->session;
        $user_type =  $session['user_type'];
        $line_id = $session['line_id'];
        $uid = $session['uid'];
        $ua_tab_arr = self::get_tab('user_ua');
        $agent_tab_arr = self::get_tab('user_agent');

        if($user_type == 1){
            return;
        }elseif($user_type == 6){//登录身份为股东
           if($type == 'ua'){//验证是否属于自己旗下的总代
                $data = AgentModel::getMoreAgentByCondition($ua_tab_arr['database'],$ua_tab_arr['tab'],['pid'=>$uid,'line_id'=>$line_id]);
                $res = false;
                if(!empty($data)){
                    foreach($data as $val){
                        if($id == $val['id']){
                            $res = true;
                            break;
                        }
                    }
                }
                if(!$res){
                     $this->wrong_msg();
                }
            }elseif($type == 'agent'){//验证是否属于自己旗下的代理
                //查询父id
                $one = AgentModel::getOneAgentByCondition($agent_tab_arr['database'],$agent_tab_arr['tab'],['id' => $id,'line_id'=>$line_id]);
                if(!$one) $this->wrong_msg();
                $pid = $one['pid'];
                $data = AgentModel::getMoreAgentByCondition($ua_tab_arr['database'],$ua_tab_arr['tab'],['pid'=>$uid,'line_id'=>$line_id]);//总代信息
                $res = false;
                foreach($data as $val){
                    if($pid == $val['id']) $res = true;
                }
                if(!$res) $this->wrong_msg();

            }
        }elseif($user_type == 7){//登录身份为总代
            if($type == 'agent'){//验证是否属于自己旗下的代理
                 $data = AgentModel::getMoreAgentByCondition($agent_tab_arr['database'],$agent_tab_arr['tab'],['pid'=>$uid,'line_id'=>$line_id]);
                $res = false;
                if(!empty($data)){
                    foreach($data as $val){
                        if($id == $val['id']){
                            $res = true;
                            break;
                        }
                    }
                }
                if(!$res){
                    $this->wrong_msg();
                }
            }
        }else{
             $this->wrong_msg();
        }
    }

}
