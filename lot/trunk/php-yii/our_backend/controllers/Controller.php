<?php

namespace our_backend\controllers;

use Yii;

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
        //当前路由传入到视图(布局文件main.html=> 左侧菜单栏选中)
        Yii::$app->view->params['route'] = $route;
        Yii::$app->view->params['time'] = date('Y/m/d H:i:s', time()); //当前时间


        $session = Yii::$app->session;
        $session->set('unikey', md5('pkadmin'));

        //判断是否开启session
        if (!$session->isActive) {
            $session->open();
        }
        $ssid = $session->id;

        //推送uid
        $session->set('unikey', md5('pkadmin'));


        if (!$session->get('uid') || ($session->get('uid') && $ssid != $this->checkOnlineRedisKey())) {
            $this->redirect(['/login/loginout']);
            return false;
        } elseif ($session->get('role') != 'super') {
            //权限验证(超管不需要验证)
            $allRoutes = $session->get('allRoutes');
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
            'other/access'
        ];
    }

    //在线会员
    public function checkOnlineRedisKey() {
        $session = Yii::$app->session;
        $uid = $session->get('uid');
        if (empty($uid)) {
            return false;
        }
        $redis = Yii::$app->redis;
        $redisKey = 'ourbackend_UserOnline_' . $uid;
        $json = $redis->get($redisKey);
        if (!empty($json)) {
            $data = json_decode($json, true);
            $data['time'] = time();
            $redis->setex($redisKey, 3600, json_encode($data));
            return $data['ssid'];
        }
        return false;
    }

}
