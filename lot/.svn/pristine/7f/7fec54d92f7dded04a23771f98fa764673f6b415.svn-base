<?php

namespace backend\controllers;

use Yii;
use backend\models\AdminModel;
use common\helpers\mongoTables;
use common\models\LogModel;
use common\helpers\Helper;
class LoginController extends \yii\web\Controller {

    public $layout = 'login.html';

    public function actions() {
        return [
            'captcha' => [
                'class' => 'yii\captcha\CaptchaAction',
                'maxLength' => 4, //生成的验证码最大长度
                'minLength' => 4  //生成的验证码最短长度
            ]
        ];
    }

    //登陆页面
    public function actionLogin() {
        if (Yii::$app->session->get('uid')) {
            //登陆状态就跳到个人主页
            $this->redirect('/admin/center');
        }
        $render = ['data' => 11111];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('login.html', $render);
        } else {
            return $this->render('login.html', $render);
        }
    }

    //登陆提交
    public function actionLogindo() {
        $post = Yii::$app->request->post();
        $uname = trim($post['uname']);
        $pwd = trim($post['pwd']);
        $vcode = trim($post['vcode']);
        $rem = $post['rem'];
        $logintype = $post['logintype'];
        $code = 400;
        $errorState = false;

        if (empty($uname)) {
            $errorState = true;
            $errorMsg = '用户名不能为空!';
        } else if (strlen($uname) > 20 || strlen($uname) < 4) {
            $errorState = true;
            $errorMsg = '用户名长度只能为4-20位!';
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $uname)) {
            $errorState = true;
            $errorMsg = '用户名只能为数字字母下划线!';
        } else if (empty($pwd)) {
            $errorState = true;
            $errorMsg = '密码不能为空!';
        } else if (strlen($pwd) > 20 || strlen($pwd) < 6) {
            $errorState = true;
            $errorMsg = '密码长度只能为6-20位!';
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $pwd)) {
            $errorState = true;
            $errorMsg = '密码只能为数字字母下划线!';
        } else if (empty($vcode)) {
            $errorState = true;
            $errorMsg = '验证码不能为空!';
        } else if (strlen($vcode) != 4) {
            $errorState = true;
            $errorMsg = '请输入4位验证码!';
        } else if (!$this->createAction('captcha')->validate($vcode, false)) {
            $code = 300;
            $errorState = true;
            $errorMsg = '验证码错误!';
        } else {
            $errorState = false;
            $errorMsg = '';
        }

        if ($errorState) {
            echo json_encode(['code' => $code, 'msg' => $errorMsg]);
            die;
        }

        //根据登入方式获取表名
        $table = 'agent_admin';
        $conn = \Yii::$app->manage_db;


        //数据库验证
        $params = [
            'username' => $uname,
            'password' => $pwd,
            'table' => $table,
            'conn' => $conn
        ];
        $user = AdminModel::login($params);
        $ip = Helper::getIpAddress();
        if (!$user) {
            //登录失败
            $log = [
                'uid' => 0,
                'login_user' => $uname,
                'line_id' => '',
                'state' => 2
            ];
        } else {
            if (!empty($user['login_ip'])) {
                $allow_ips = explode(",", trim($user['login_ip']));
                if (!in_array($ip, $allow_ips)) {
                    //登录失败
                    $log = [
                        'uid' => 0,
                        'login_user' => $uname,
                        'state' => 2,
                        'line_id' => ''
                    ];
                    //登陆日志
                    $this->writeLog($log);
                    echo json_encode(['code' => 300, 'msg' => 'IP配置失败!']);
                    die;
                }
            }

            //登录成功
            $log = [
                'uid' => $user['uid'],
                'login_user' => $user['login_user'],
                'line_id' => $user['line_id'],
                'state' => 1
            ];
        }

        //登陆日志
        $this->writeLog($log);

        if (!$user) {
            echo json_encode(['code' => 300, 'msg' => '用户名或密码错误!']);
            die;
        } else {
            //写session信息
            $this->writeSession($user);
            //写入登录redis 哈希key(更新在线人数及登录时间)
            $this->updateOnlineRedisKey($user);
            echo json_encode(['code' => 200, 'msg' => '登录成功!']);
            die;
        }
    }

    public function writeLog($log) {
        $ip = Helper::getIpAddress();
        $logTable = mongoTables::getTable('historyLogin');
        $logInsert = [
            'uid' => $log['uid'],
            'uname' => $log['login_user'],
            'ptype' => 2,
            'line_id' => $log['line_id'], //客户端后台
            'ip' => $ip,
            'addtime' => time(),
            'adddate' => date("Y-m-d H:i:s"),
            'state' => $log['state']
        ];
        LogModel::insertLog($logTable, $logInsert);
    }

    public function writeSession($user) {
        // 菜单按级别分组
        $oneLevel = [];
        $twoLevel = [];
        $threeLevel = [];
        $allRoutes = [];
        if (is_array($user['menu']) && !empty($user['menu'])) {
            foreach ($user['menu'] as $k => $v) {
                if ($v['level'] == 1) {
                    $oneArr['menuName'] = $v['menu_name'];
                    $oneArr['menuUrl'] = $v['menu'];
                    $oneArr['id'] = $v['id'];
                    $oneArr['iconClass'] = $v['icon_class'];
                    $oneLevel[] = $oneArr;
                } else if ($v['level'] == 2) {
                    $twoArr['menuName'] = $v['menu_name'];
                    $twoArr['menuUrl'] = $v['menu'];
                    $twoArr['id'] = $v['id'];
                    $twoArr['pid'] = $v['pid'];
                    $twoLevel[] = $twoArr;
                } else if ($v['level'] == 3) {
                    $threeArr['menuName'] = $v['menu_name'];
                    $threeArr['menuUrl'] = $v['menu'];
                    $threeArr['id'] = $v['id'];
                    $threeArr['pid'] = $v['pid'];
                    $threeLevel[] = $threeArr;
                }
                if (trim($v['menu']) != '#') {
                    $allRoutes[] = trim($v['menu']);
                }
            }
        }

        //三级菜单装入二级
        if (!empty($twoLevel)) {
            foreach ($twoLevel as $tk => $tv) {
                foreach ($threeLevel as $thk => $thv) {
                    if ($tv['id'] == $thv['pid']) {
                        $twoLevel[$tk]['threeLevel'][] = $thv['menuUrl'];
                    }
                }
            }
        }


        //二级菜单装入一级
        if (!empty($oneLevel)) {
            foreach ($oneLevel as $kk => $vv) {
                foreach ($twoLevel as $key => $val) {
                    if ($vv['id'] == $val['pid']) {
                        $oneLevel[$kk]['twoLevelMenu'][] = $val;
                    }
                }
            }
        }

        $session = Yii::$app->session;

        //判断是否开启
        if (!$session->isActive) {
            $session->open();
        }


        $session->set('line_id', $user['line_id']);
        $session->set('uid', $user['uid']);
        $session->set('pid', $user['pid']);
        $session->set('uname', $user['login_user']);
        $session->set('login_user', $user['login_user']);
        $session->set('login_name', $user['login_name']);
        $session->set('user_type', $user['user_type']);
        $session->set('role', $user['role']);
        $session->set('role_id', $user['role_id']);
        $session->set('role_name', $user['role_name']);
        $session->set('role_access', $user['role_access']);
        $session->set('menu', $oneLevel);
        $session->set('allRoutes', $allRoutes);
        if (isset($user['son_role'])) {
            $session->set('son_role', $user['son_role']);
        }
        if (isset($user['agent_id'])) {
            $session->set('agent_id', $user['agent_id']);
        }
        if (isset($user['agent_name'])) {
            $session->set('agent_name', $user['agent_name']);
        }
        $session->close();
    }

    //更新在线人数
    public function updateOnlineRedisKey($user) {
        $redisKey = 'agentbackend_UserOnline_' . $user['uid'];
        $session = Yii::$app->session;
        $redisConn = Yii::$app->redis;
        $data = [
            'ssid' => $session->id,
            'uid' => $user['uid'],
            'uname' => $user['login_user'],
            'time' => time(),
        ];
        $redisConn->setex($redisKey, 3600, json_encode($data));
    }

    //退出登陆
    public function actionLoginout() {
        $session = Yii::$app->session;
        $uid = $session['uid'];
        $redis = Yii::$app->redis;
        $redis->del('agentbackend_UserOnline_' . $uid);

        if (!$session->isActive) {
            $session->open();
        }
        $session->destroy();
        $session->close();


        return $this->redirect('/login/login');
    }

/**
      ***********************************************************
      *  维护页面           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionMaintain(){
        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        $wh = $this->game_is_maintain($line_id);
        if(isset($wh['return']) && $wh['return'] == 1){
             return $this->redirect('/admin/center');
             return false;
        }
        $return = array();
        $return['starttime'] = isset($wh['starttime']) ? date('Y-m-d H:i:s', $wh['starttime']) : '未知';
        $return['endtime'] = isset($wh['endtime']) ? date('Y-m-d H:i:s', $wh['endtime']) : '未知';
        $return['remark'] = isset($wh['remark']) ? $wh['remark'] : '例行维护';

        $render = ['data' => $return];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('/other/maintain.html', $render);
        } else {
            return $this->render('/other/maintain.html', $render);
        }
    }

    //判断线路是否维护
    public function game_is_maintain($line_id) {
        $result['return'] = 1;
        $result['remark'] = '';
        $result['starttime'] = '';
        $result['endtime'] = '';

        $data = Yii::$app->redis->get("maintain_agent_all_line_ids");
        $data = json_decode($data, true);
        if ($data) {
            $return = in_array($line_id, $data['line_id']) ? 2 : 1;
            $result['return'] = $return;
            //如果设置了开始时间和结束时间 验证是否过期
            if(isset($data['endtime']) && time() >= $data['endtime']){
                 $result['return'] = 1;
            }
            $result['remark'] = $data['remark'];
            $result['starttime'] = $data['starttime'];
            $result['endtime'] = $data['endtime'];
        }
        return $result;
    }
}
