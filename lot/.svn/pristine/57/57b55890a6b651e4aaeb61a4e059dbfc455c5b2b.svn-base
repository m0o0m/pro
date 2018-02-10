<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use backend\models\AdminModel;

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
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('tree.html');
        } else {
            return $this->render('tree.html');
        }
    }

    public function actionTree2() {
        $allMenu = AdminModel::getAllRoute();
        $oneLevel = [];
        $twoLevel = [];
        $threeLevel = [];
        if (!empty($allMenu)) {
            foreach ($allMenu as $k => $v) {
                if ($v['level'] == 1) {
                    $oneLevel[] = $v;
                } else if ($v['level'] == 2) {
                    $twoLevel[] = $v;
                } else if ($v['level'] == 3) {
                    $threeLevel[] = $v;
                }
            }
        }

        $render = [
            'oneLevel' => $oneLevel,
            'twoLevel' => $twoLevel,
            'threeLevel' => $threeLevel
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('tree2.html', $render);
        } else {
            return $this->render('tree2.html', $render);
        }
    }


/**
      ***********************************************************
      *  公共部分内容展示           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionGetdata(){
         //顶部余额展示
        $session = Yii::$app->session;
        $uid = $session->get('uid'); //用户id
        $login_user = $session->get('login_user'); //用户id
        $pid = $session->get('pid'); //用户父级id
        $line_id = $session->get('line_id'); //线路id
        $agent_id = $session->get('agent_id'); //代理id
        $user_type = $session->get('user_type'); //帐号类型
        $tab;
        $where;
        $online = 0;
        $redis = Yii::$app->redis;
        /* '账号类型(1.管理员 2.管理员子帐号 3.代理 4.代理子帐号)' */
        switch ($user_type) {
            case 1:
            case 2:
                $tab = 'sys_line_list';
                $where = ['line_id' => $line_id];
                $online = $redis->hlen($line_id . '_uidToken');//在线会员
                break;
            case 3:
            case 4:
                $tab = 'user_agent';
                $where = ['id' => $agent_id];
                $online = $redis->hlen($agent_id . '_uidToken');//在线会员
                break;
        }
        $money = AdminModel::queryMoney($tab, $where);

        $return['money'] = $money;//余额
        $return['online'] = $online;//在线会员人数

        //公告跑马灯
        $where  =  array();
        $where[] = 'and';
        $where[] = ['or', ['=', 'line_id',$line_id], ['=', 'line_id', '']];
        $where[] = ['=', 'status', 1];
        $notice = AdminModel::getNoticeInfo($where);
        if($notice){
            $return['notice'] = $notice;
        }else{
            $return['notice'] = array();
        }

        echo json_encode($return);die;
    }
    
}
