<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use common\models\NoticefrontModel as selfModel;
use common\models\AgentModel;
use yii\helpers\ArrayHelper;

class NoticefrontController extends Controller {

    public function actionIndex() {
        $data['show'] = $this->showData();
        $data['trans'] = $this->transData();

        $request = Yii::$app->request->get();
        $result = $this->search($request, $data['show']);

        $render = [
            'data' => $data,
            'result' => $result
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    public function actionAdd() {
        if (!Yii::$app->request->isPost) {
            $render = [
                'show' => $this->showData()
            ];

            $request = Yii::$app->request->get();
            $id = isset($request['id']) ? trim($request['id']) : null;
            if (is_numeric($id)) {
                $render['item'] = selfModel::getOne(['id' => $id]);
                // if ($render['item']['agent_id']) {
                //     $aids = explode(',', $render['item']['agent_id']);
                //     foreach ($aids as $aid) {
                //         $aid = explode('_', $aid);
                //         if (!empty($aid[1]))
                //             $render['item']['aids'][] = $aid[1];
                //     }
                // }
                // if ($render['item']['agent_type']) {
                //     $render['item']['agents'] = AgentModel::getAgent($render['item']['agent_type'], $render['item']['line_id']);
                // }
            }

            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('_form.html', $render);
            } else {
                return $this->render('_form.html', $render);
            }
        } else {
            $request = Yii::$app->request->post();
            $result = $this->addSubmit($request);

            return json_encode($result);
        }
    }

    // public function actionAgent() {
    //     $request = Yii::$app->request->post();
    //     $result = $this->agent($request); // 切换代理类型

    //     return json_encode($result);
    // }

    public function actionEnable() {
        $request = Yii::$app->request->post();
        $id = isset($request['id']) ? trim($request['id']) : null;
        $enable = isset($request['enable']) ? trim($request['enable']) : null;

        $result = $this->enable($id, $enable);

        return json_encode($result);
    }

    /*
     * ******************************************************************************
     */

    public function search($request, $showData) {
        $query = $request;
        unset($query['_pjax']);
        if (empty($query)) {
            return;
        }

        $keyword = isset($request['keyword']) ? trim($request['keyword']) : '';
        $type = isset($request['type']) ? trim($request['type']) : '';
        $line_id = isset($request['line_id']) ? trim($request['line_id']) : '';
        $add_endtime = isset($request['add_endtime']) ? trim($request['add_endtime']) : '';
        $add_starttime = isset($request['add_starttime']) ? trim($request['add_starttime']) : '';
        $send_endtime = isset($request['send_endtime']) ? trim($request['send_endtime']) : '';
        $send_starttime = isset($request['send_starttime']) ? trim($request['send_starttime']) : '';
        $status = isset($request['status']) ? intval($request['status']) : 0;
        $page = isset($request['page']) ? intval($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? intval($request['pagenum']) : 100;

        // 查询条件
        $where[] = 'and';
        if ($keyword) {
            $where[] = [
                'OR',
                ['LIKE', 'title', $keyword],
                ['LIKE', 'content', $keyword]
            ];
        }
        if ($type || $type === '0') {
            $where[] = ['=', 'type', $type];
        }
        if ($line_id) {
            $where[] = ['=', 'line_id', $line_id];
        }
        if ($status) {
            $where[] = ['=', 'status', $status];
        }

        // 时间 处理
        $add_endtime = strtotime($add_endtime) ? strtotime($add_endtime) : $add_endtime;
        $add_starttime = strtotime($add_starttime) ? strtotime($add_starttime) : $add_starttime;
        $send_endtime = strtotime($send_endtime) ? strtotime($send_endtime) : $send_endtime;
        $send_starttime = strtotime($send_starttime) ? strtotime($send_starttime) : $send_starttime;
        // 时间筛选
        if ($add_starttime && $add_endtime) {
            $where[] = ['between', 'addtime', $add_starttime, $add_endtime];
        } elseif (!empty($add_starttime)) {
            $where[] = ['>', 'addtime', $add_starttime];
        } elseif (!empty($add_endtime)) {
            $where[] = ['<', 'addtime', $add_endtime];
        }
        if ($send_starttime && $send_endtime) {
            $where[] = ['between', 'sendtime', $send_starttime, $send_endtime];
        } elseif (!empty($send_starttime)) {
            $where[] = ['>', 'sendtime', $send_starttime];
        } elseif (!empty($send_endtime)) {
            $where[] = ['<', 'sendtime', $send_endtime];
        }

        // 分页
        $total_count = selfModel::getCount($where);
        $total_page = ceil($total_count / $pagenum);
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        if ($page <= 0)
            $page = 1;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $rows = selfModel::getList($where, $offset, $limit, ['id'=>SORT_DESC]);

        // 翻译 与 时间格式
        $rows = $this->trans($rows, $showData);

        $result['ErrorCode'] = 1;
        $result['data'] = $rows;
        $result['page'] = $page;
        $result['pagenum'] = $pagenum;
        $result['totalpage'] = $total_page;
        return $result;
    }

    public function showData($type = '') {

        switch($type){
            case 'lines':
                $lines = parent::getLines();
                $data = ArrayHelper::index($lines, 'line_id');
                break;
            case 'agents':
                $agents = AgentModel::getAgent('at');
                $data = ArrayHelper::index($agents, 'id');
                break;
            default:
                $lines = parent::getLines();
                $agents = AgentModel::getAgent('at');

                $lines = ArrayHelper::index($lines, 'line_id');
                $agents = ArrayHelper::index($agents, 'id');

                $data['lines'] = $lines;
                $data['agents'] = $agents;
                $data['types'] = ['所有平台', 'pc', 'wap', 'app'];
                $data['atypes'] = ['sh' => '股东', 'ua' => '总代', 'at' => '代理'];
                break;
        }

        return $data;
    }

    public function transData() {

        $result['status'] = [1 => '有效', 2 => '无效'];

        return $result;
    }

    public function trans($data, $showData) {
        // $showData = $this->showData();
        $transData = $this->transData();

        foreach ($data as $k => &$v) {
            $v['typeTxt'] = key_exists($v['type'], $showData['types']) ? $showData['types'][$v['type']] : $v['type'];
            $v['agentTxt'] = key_exists($v['agent_id'], $showData['agents']) ? $showData['agents'][$v['agent_id']]['login_name'] : $v['agent_id'];
            $v['statusTxt'] = key_exists($v['status'], $transData['status']) ? $transData['status'][$v['status']] : $v['status'];
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
        }unset($v);

        return $data;
    }

    // public static function agent($request) {
    //     $line_id = isset($request['line_id']) ? trim($request['line_id']) : null;
    //     $agent_type = isset($request['agent_type']) ? trim($request['agent_type']) : null;

    //     if (/*empty($line_id) || */empty($agent_type)) {
    //         return [];
    //     }
    //     return AgentModel::getAgent($agent_type, $line_id);
    // }

    public function addSubmit($request) {
        $line_id = isset($request['line_id']) ? trim($request['line_id']) : '';
        $agent_type = isset($request['agent_type']) ? trim($request['agent_type']) : '';
        $agent_id = isset($request['agent_id']) ? $request['agent_id'] : '';
        $type = isset($request['type']) ? intval($request['type']) : 0;
        $title = isset($request['title']) ? trim($request['title']) : '';
        $content = isset($request['content']) ? trim($request['content']) : '';
        $remark = isset($request['remark']) ? trim($request['remark']) : '';

        $info = isset($request['info']) ? trim($request['info']) : '';

        if (empty($title)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '标题不能为空';
            return $result;
        }
        if (empty($content)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '内容不能为空';
            return $result;
        }

        if($agent_id){
            foreach($agent_id as $v){
                if(isset($aid))
                    $aid .= ',' . $agent_type . '_' . $v;
                else
                    $aid = $agent_type . '_' . $v;
            }
        }else{
            $aid = '';
        }
        $cols = [
            'line_id' => $line_id,
            'agent_type' => $agent_type,
            'agent_id' => $aid,
            'type' => $type,
            'title' => $title,
            'content' => $content,
            'remark' => $remark
        ];
        $request = Yii::$app->request->post();
        $id = isset($request['id']) ? trim($request['id']) : '';
        $remark = '';
        if($id){
            $where['id'] = $id;
            $sqlQuery = selfModel::update($cols, $where);
            if (!$sqlQuery) {
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '未修改';
                return $result;
            }
            $result['ErrorMsg'] = '修改成功';
            $remark = '修改前台公告';
        }else{
            $cols['addtime'] = time();
            $sqlQuery = selfModel::insert($cols);
            $result['ErrorMsg'] = '添加成功';
            $remark = '添加前台公告';
        }

        $new = $cols;
        $this->insertOperateLog($info, json_encode($new), $remark);

        $this->batchDelRedis('after', '_site_notice' );// 前台线路公告信息的key

        $result['ErrorCode'] = 0;
        return $result;
    }

    public function enable($id, $enable) {
        $request = Yii::$app->request->post();

        $info = isset($request['info']) ? trim($request['info']) : '';

        if (empty($id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'id不能为空';
            return $result;
        }

        if($enable == 1){
            $enable = 2;
            $msg = '关闭成功';
        }else{
            $enable = 1;
            $msg = '启用成功';
        }

        $where['id'] = $id;
        $set = [
            'status' => $enable
        ];
        $sqlQuery = selfModel::update($set, $where);

        if (!$sqlQuery) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '删除失败';
            return $result;
        }

        $new = $set;

        $info = json_decode($info, true);
        $this->batchDelRedis('after', '_site_notice' );// 前台线路公告信息的key

        $result['ErrorCode'] = 0;
        $result['ErrorMsg'] = $msg;
        return $result;
    }
}
