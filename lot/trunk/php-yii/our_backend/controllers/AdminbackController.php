<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AdminbackModel as selfModel;

class AdminbackController extends Controller {

/**
	  ***********************************************************
	  *  超管列表页           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public function actionIndex(){
        $lines = $this->getLines(); //获取线路
        $get = Yii::$app->request->get();
        $line_id = isset($get['line_id']) ? $get['line_id'] : '';
        $login_user = isset($get['login_user']) ? $get['login_user'] : '';
        $login_name = isset($get['login_name']) ? $get['login_name'] : '';
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 50;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;

        $where = array();
        $where['user_type'] = 1;
        if($line_id) $where['line_id'] = $line_id;
        if($login_user) $where['login_user'] = $login_user;
        if($login_name) $where['login_name'] = $login_name;
       	
       	$data = array();
       	$data = selfModel::getData($where, $offset, $pagenum);
        $count = count($data);
            // $count = selfModel::getCount($where);
        if(!$count){
             $render = [
                'data' => array(),
                'lines' => $lines,
                'pagecount' => 1,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }

        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
       	foreach($data as $key=>$val){
       		$data[$key]['addtime'] = date('Y-m-d H:i:s', $val['addtime']);
       		$data[$key]['updatetime'] = date('Y-m-d H:i:s', $val['updatetime']);
       		$state = ['', '有效', '无效'];
       		if(isset($state[$val['is_delete']])){
       			$data[$key]['deleteTxt'] = $state[$val['is_delete']];
       		}else{
       			$data[$key]['deleteTxt'] = '未知';
       		}

             //是否显示踢线
            $redis = Yii::$app->redis;
            $data[$key]['is_withdrawals'] = false;
            $is_exists = $redis->exists('agentbackend_UserOnline_' . $val['id']);
            if($is_exists) $data[$key]['is_withdrawals'] = true;
            $get = Yii::$app->request->get();
            $online = isset($get['online']) ? $get['online'] : null;
            if($online){
                if($online == 1 && !$is_exists){
                    unset($data[$key]);
                    continue;
                }
                if($online == 2 && $is_exists){
                    unset($data[$key]);
                    continue;
                }
            }
       	}
        $render = [
            'data' => $data,
            'lines' => $lines,
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
	  *  增加和修改超级管理员         @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
 	public function actionSave() {
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : '';
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['confPwd']) ? $post['confPwd'] : '';
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
        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'login_ip'   => $login_ip,
            'is_delete' => $is_delete,
            'updatetime' => time(),
            'user_type' => 1
        ];

 
        $manage_db = \Yii::$app->manage_db;
        $manage_tab = \Yii::$app->manage_db->tablePrefix . 'agent_admin';
        if (!empty($id)) {
            $oldInfo = selfModel::getOneData(array('id'=>$id));
            unset($data['line_id']);
            unset($data['login_user']); //不能修改帐号
            if(!empty($pwd)){
            	if($pwd != $conf_pwd){
            		echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            		die;
            	}
            	$data['login_pwd'] = md5(md5($pwd));
            }
            //修改
            $res = selfModel::update($data, ['id' => $id]);
        } else {
            //新增
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            //帐号唯一性
            $is_exist = selfModel::getCount(['login_user'=>$login_user]);
            if($is_exist){
                echo json_encode(['code' => 400, 'msg' => '帐号已经存在!']);
                die;
            }
            //每条线路只有一个超管
            $line_exist = selfModel::getCount(['line_id'=>$line_id, 'user_type'=> 1]);
            if($line_exist){
            	echo json_encode(['code' => 400, 'msg' => '该线路已经拥有超级管理员!']);
                die;
            }

            //插入
          	$res = selfModel::insert($data);
         	
        }

        if ($res) {
            //mongo日志
            if(!empty($id)){
                $remark = '修改超级管理员：' . $login_user;
                parent::insertOperateLog(json_encode($oldInfo),json_encode($data),$remark);
            }else{//新增
                $remark = '为线路：' . $line_id .'添加超级管理员：' . $login_user;
                parent::insertOperateLog('',json_encode($data),$remark);
            }

            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }
	
}