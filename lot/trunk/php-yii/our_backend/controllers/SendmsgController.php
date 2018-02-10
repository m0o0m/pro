<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\Models\SendmsgModel;

class SendmsgController extends Controller {

    /**
     * **********************************************************
     *  消息推送列表             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $starttime = isset($get['starttime']) ? strtotime($get['starttime']) : null;
        $endtime = isset($get['endtime']) ? strtotime($get['endtime']) : null;
        $sendtype = isset($get['sendtype']) ? $get['sendtype'] : null;
        $sendfor = isset($get['sendfor']) ? $get['sendfor'] : null;
        $page = isset($get['page']) ? intval($get['page']) : 1;
        $pagenum = isset($get['pagenum']) ? abs(intval($get['pagenum'])) : 100;


        //拼接查询条件
        $where = array();
        if ((!empty($sendfor)) && (in_array($sendfor, array(1, 2)))) {
            $where['sendfor'] = $sendfor; //推送目标1前台2后台
        }
        if ((!empty($sendtype)) && (in_array($sendtype, array(1, 2)))) {
            $where['sendtype'] = $sendtype; //1手动推送2自动推送
        }

        //指定时间查询
        if ($starttime && $endtime) {
            $where = ['and', $where, array('between', 'addtime', $starttime, $endtime)];
        } elseif (!empty($endtime)) {
            $where = ['and', $where, array('<=', 'addtime', $endtime)];
        } elseif (!empty($starttime)) {
            $where = ['and', $where, 'addtime>=' . $starttime];
        } else {
            
        }

        //查询处理数据
        $data = array();
        $count = SendmsgModel::getCount($where);
        if ($count > 0) {
            $total_page = ceil($count / $pagenum); //总页码数
            if ($page > $total_page && $total_page != 0)
                $page = $total_page;
            $offset = ($page - 1) * $pagenum;  //开始条数
            $limit = $pagenum;       //每页显示条数

            $data = SendmsgModel::getList($where, $offset, $limit);

            if (!empty($data)) {
                foreach ($data as $key => $val) {
                    $data[$key]['addtime'] = date('Y-m-d H:i', $val['addtime']);
                    if (!empty($val['sendtime'])) {
                        $data[$key]['sendtime'] = date('Y-m-d H:i', $val['sendtime']);
                    }
                }
            }
        } else {
            $data = array();
            $total_page = 1;
        }

        $render = ['data' => $data, 'pagecount' => $total_page, 'page' => $page];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  添加推送消息                @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionInsert() {
        $post = Yii::$app->request->post();
        $sendtype = isset($post['sendtype']) ? $post['sendtype'] : 1;
        $sendfor = isset($post['sendfor']) ? $post['sendfor'] : 1;
        $sendtime = isset($post['sendtime']) ? strtotime($post['sendtime']) : '';
        $content = isset($post['content']) ? $post['content'] : '';

        $result = array();
        $data = array();

        if (empty($content)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '推送消息不能为空!';
            return json_encode($result);
        }

        if ($sendtype == 2) {
            if (empty($sendtime)) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '自动推送必须选择推送时间!';
                return json_encode($result);
            }

            if ($sendtime <= time()) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '推送时间不能小于当前时间!';
                return json_encode($result);
            }
        }

        $data['sendtype'] = $sendtype;
        $data['sendfor'] = $sendfor;
        $data['content'] = $content;
        if (!empty($sendtime)) {
            $data['sendtime'] = $sendtime;
        }
        $data['is_send'] = 1;
        $data['addtime'] = time();
        $res = SendmsgModel::insert($data);
        if ($res) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '添加推送消息成功!';
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '添加推送消息失败!';
        }
        return json_encode($result);
    }

}
