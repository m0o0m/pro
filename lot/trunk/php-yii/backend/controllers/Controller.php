<?php

namespace backend\controllers;

use Yii;
use backend\models\AdminModel;

class Controller extends \common\controllers\Controller {

    //默认布局文件为main.html
    public $layout = 'main.html';

    public function beforeAction($action) {
        //当不是以module建的三层文件夹 module->id 就是项目根目录文件夹的名字
        $projectName = Yii::$app->id; //配置文件main.php中的配置
        $module = Yii::$app->controller->module->id;
        $controller = Yii::$app->controller->id;
        $action = Yii::$app->controller->action->id;
        //拼接当前路由
        if ($projectName == $module) {
            $route = $controller . '/' . $action;
        } else {
            $route = $module . '/' . $controller . '/' . $action;
        }

        //统计人数使用
        $session = Yii::$app->session;

        //判断是否开启session
        if (!$session->isActive) {
            $session->open();
        }
        $ssid = $session->id;
        $user_type = $session['user_type'];
        $line_id = $session['line_id'];
        if (in_array($user_type, [1, 2])) {
            $session->set('unikey', md5($line_id));
        } elseif (in_array($user_type, [3, 4])) {
            $agent_id = $session['agent_id'];
            $session->set('unikey', md5($line_id . '_' . $agent_id));
        }

        //当前路由传入到视图(布局文件main.html=> 左侧菜单栏选中)
        Yii::$app->view->params['route'] = $route;
        Yii::$app->view->params['time'] = date('Y/m/d H:i:s', time());


        //判断是否登录
        if (!$session->get('uid') || ($session->get('uid') && $ssid != $this->checkOnlineRedisKey())) {
            $this->redirect(['/login/loginout']);
            return false;
        }

        //判断是否维护
        $wh = $this->game_is_maintain(1, 'agent', $line_id);
        if ($wh['return'] == 2) {
            return $this->redirect(['/login/maintain']);
            return false;
        }

        //判断是否有权限
        if (Yii::$app->session->get('user_type') != 1) {
            //权限验证(超管不需要验证)
            $allRoutes = Yii::$app->session->get('allRoutes');
            if (!in_array($route, $allRoutes) && !in_array($route, $this->filter())) {
                $get = Yii::$app->request->get();
                if (isset($get['_pjax'])) {
                    $this->redirect(['/other/access']);
                    return false;
                }
                if (Yii::$app->request->isAjax) {
                    $result = array();
                    $result['ErrorCode'] = 2;
                    $result['ErrorMsg'] = '您没有操作权限';
                    echo json_encode($result);
                    die;
                } else {
                    $this->redirect(['/other/access']);
                    return false;
                }
            }
        }


        return parent::beforeAction($action);
    }

    //不需要任何权限就可以访问的页面
    public function filter() {
        return [
            'admin/center',
            'other/access',
            'login/maintain'
        ];
    }

    //在线会员判断
    public function checkOnlineRedisKey() {
        $session = Yii::$app->session;
        $uid = $session->get('uid');
        if (empty($uid)) {
            return false;
        }
        $redis = Yii::$app->redis;
        $redisKey = 'agentbackend_UserOnline_' . $uid;
        $json = $redis->get($redisKey);
        if (!empty($json)) {
            $data = json_decode($json, true);
            $data['time'] = time();
            $redis->setex($redisKey, 3600, json_encode($data));
            return $data['ssid'];
        }
        return false;
    }

    //无权限时自动跳转
    public function wrong_msg() {
        $get = Yii::$app->request->get();
        if (isset($get['_pjax'])) {
            $this->redirect(['/other/access']);
            return false;
        }
        if (Yii::$app->request->isAjax) {
            $result = array();
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '您没有操作权限';
            echo json_encode($result);
            die;
        } else {
            $this->redirect(['/other/access']);
            return false;
        }
    }

    /**
     * **********************************************************
     *  列表页根据登录身份不同必加判断条件 @author ruizuo qiyongsheng *
     * **********************************************************
     */
    public function loginWhere($at = 'agent_id') {
        $session = Yii::$app->session;
        $utype = $session['user_type'];
        //目前只有管理员和代理及子帐号登录
        $where = array();
        $where['line_id'] = $session['line_id'];
        if ($utype == 3 || $utype == 4) {
            $where[$at] = $session['agent_id'];
        }
        return $where;
    }

}
