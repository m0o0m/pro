<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\PushModel as selfModel;
use common\models\LineModel;
use yii\helpers\ArrayHelper;

class PushController extends Controller {

    public function init() {
        date_default_timezone_set('PRC');
    }

    public function actionResetdata() {// 重制测试数据
        \Yii::$app->manage_db
                ->createCommand('TRUNCATE `my_push_list`;TRUNCATE `my_push_queue`;')
                ->execute();
        $timestamp = time();
        $starttime = $timestamp - 10 * 60;
        $interval = 5 * 60; // 每条记录间隔
        $pc = 3 * 60; // 偏差
        for ($i = 1; $i <= 666; $i++) {
            $values = [
                'title' => '测试标题' . $i,
                'content' => '测试内容' . $i,
                'remark' => '备注' . $i,
                'sendtime' => $starttime + $interval * $i + ( rand(1, $pc * 2) - $pc ),
                'addtime' => $timestamp,
                'timezone' => date_default_timezone_get(),
                'enable' => '1'
            ];
            selfModel::insert($values);
        }
        echo 'Reset success';
    }

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
        $data['show'] = $this->showData();

        $render = [
            'data' => $data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('_form.html', $render);
        } else {
            return $this->render('_form.html', $render);
        }
    }

    public function actionEdit() {
        $data['show'] = $this->showData();

        $request = Yii::$app->request->get();
        $id = isset($request['id']) ? intval($request['id']) : 0;
        $item = selfModel::getOne(['id' => $id]);

        $item['clients'] = explode(',', $item['clients']);
        $item['lines'] = explode(',', $item['lines']);
        $item['games'] = explode(',', $item['games']);
        $item['groups'] = explode(',', $item['groups']);
        $item['users'] = explode(',', $item['users']);
        $item['sendtime'] = date('Y-m-d H:i:s', $item['sendtime']);

        $render = [
            'data' => $data,
            'item' => $item
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('_form.html', $render);
        } else {
            return $this->render('_form.html', $render);
        }
    }

    public function actionQuery() {
        $result = $this->query();

        echo json_encode($result);
        exit;
    }

    public function actionEnable() {
        $result = $this->enable();

        echo json_encode($result);
        exit;
    }

    public function actionUnread() {
        $result = $this->unread();

        $request = Yii::$app->request->get();
        $jsonpCallback = isset($request['callback']) ? trim($request['callback']) : false;
        if ($jsonpCallback !== false) {
            echo $jsonpCallback . '(' . json_encode($result) . ')'; // 跨域提交
            exit;
        }
        echo json_encode($result);
        exit;
    }

    public function actionBind() {
        $result = $this->bind();

        echo json_encode($result);
        exit;
    }

    public function actionRead() {
        $result = $this->read();

        $request = Yii::$app->request->get();
        $jsonpCallback = isset($request['callback']) ? trim($request['callback']) : false;
        if ($jsonpCallback !== false) {
            echo $jsonpCallback . '(' . json_encode($result) . ')'; // 跨域提交
            exit;
        }
        echo json_encode($result);
        exit;
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
        $add_endtime = isset($request['add_endtime']) ? trim($request['add_endtime']) : '';
        $add_starttime = isset($request['add_starttime']) ? trim($request['add_starttime']) : '';
        $send_endtime = isset($request['send_endtime']) ? trim($request['send_endtime']) : '';
        $send_starttime = isset($request['send_starttime']) ? trim($request['send_starttime']) : '';
        $status = isset($request['status']) ? intval($request['status']) : 0;
        $enable = isset($request['enable']) ? trim($request['enable']) : '';
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
        if ($status) {
            $where[] = ['=', 'status', $status];
        }
        if ($enable != '') {
            $where[] = ['=', 'enable', $enable];
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

        $rows = selfModel::getList($where, $offset, $limit);

        // 翻译 与 时间格式
        $rows = $this->trans($rows, $showData);

        $result['ErrorCode'] = 1;
        $result['data'] = $rows;
        $result['page'] = $page;
        $result['pagenum'] = $pagenum;
        $result['totalpage'] = $total_page;
        return $result;
    }

    public function query() {
        $request = Yii::$app->request->post();
        $fields = ['client_limit', 'line_limit', 'game_limit', 'group_limit', 'user_limit', 'title', 'content', 'remark', 'sendtime'];
        foreach ($fields as $field) {
            $$field = isset($request[$field]) ? trim($request[$field]) : false;
            if ($$field === false) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '参数错误:' . $field;
                return $result;
            }
            if ($$field == '') {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '参数不能为空:' . $field;
                return $result;
            }
        }
        $fields = ['client_limit' => 'clients', 'line_limit' => 'lines', 'game_limit' => 'games', 'group_limit' => 'groups', 'user_limit' => 'users'];
        foreach ($fields as $k => $field) {
            $$field = isset($request[$field]) ? ( is_array($request[$field]) ? implode(',', $request[$field]) : trim($request[$field]) ) : false;
            if ($$k == 1) {
                if ($$field === false) {
                    $result['ErrorCode'] = 2;
                    $result['ErrorMsg'] = '参数错误:' . $field;
                    return $result;
                }
                if ($$field == '') {
                    $result['ErrorCode'] = 2;
                    $result['ErrorMsg'] = '参数不能为空:' . $field;
                    return $result;
                }
            }
        }

        $sendtime = strtotime($sendtime) ? strtotime($sendtime) : $sendtime;

        $cols = [
            'client_limit' => $client_limit,
            'clients' => $clients,
            'line_limit' => $line_limit,
            'lines' => $lines,
            'game_limit' => $game_limit,
            'games' => $games,
            'group_limit' => $group_limit,
            'groups' => $groups,
            'user_limit' => $user_limit,
            'users' => $users,
            'title' => $title,
            'content' => $content,
            'remark' => $remark,
            'sendtime' => $sendtime
        ];
        $act = isset($request['act']) ? trim($request['act']) : false;
        if ($act == 'insert') {
            $info = '';
            $values['addtime'] = time();
            $sqlQuery = selfModel::insert($cols);

            $result['ErrorMsg'] = '添加成功';
        } else if ($act == 'update') {
            $id = isset($request['id']) ? trim($request['id']) : false;
            $info = isset($request['info']) ? trim($request['info']) : false;
            $where['id'] = $id;
            $sqlQuery = selfModel::update($cols, $where);

            if (!$sqlQuery) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '未修改';
                return $result;
            }

            $result['ErrorMsg'] = '修改成功';
        }
        $new = $cols;
        $username = Yii::$app->session->get('uname');
        $remark = $username . " 修改了 推送消息";
        $this->insertOperateLog($info, json_encode($new), $remark);

        $result['ErrorCode'] = 1;
        return $result;
    }

    public function enable() {
        $request = Yii::$app->request->post();
        $id = isset($request['id']) ? intval($request['id']) : 0;
        $enable = isset($request['enable']) ? intval($request['enable']) : 0;

        if (empty($id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数错误';
            return $result;
        }

        if ($enable == 1) {
            $enable = 0;
            $msg = '关闭成功';
        } else {
            $enable = 1;
            $msg = '启用成功';
        }

        $where['id'] = $id;
        $set = [
            'enable' => $enable
        ];
        $sqlQuery = selfModel::update($set, $where);

        if (!$sqlQuery) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '修改失败';
            return $result;
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = $msg;
        return $result;
    }

    public function unread() {
        $request = Yii::$app->request->get();
        $group_list = isset($request['group_list']) ? trim($request['group_list']) : false;
        $uid = isset($request['uid']) ? trim($request['uid']) : false;

        if (empty($group_list)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数错误';
            return $result;
        }

        $group_list = json_decode($group_list, true);

        if (!isset($group_list['line_id']))
            $group_list['line_id'] = ' ';
        if (!isset($group_list['fc_type']))
            $group_list['fc_type'] = ' ';
        if (!isset($group_list['agent_id']))
            $group_list['agent_id'] = ' ';

        $timestamp = time();
        $where[] = 'AND';
        $where[] = ['=', 'enable', 1];
        $where[] = ['>', 'status', 0];
        $where[] = ['between', 'addtime', ( $timestamp - 7 * 24 * 3600 ), $timestamp];
        $where[] = [
            'OR',
            ['=', 'client_limit', 0],
            [
                'AND',
                ['=', 'client_limit', 1],
                [
                    'OR',
                    ['LIKE', 'clients', $group_list['client'] . ','],
                    ['LIKE', 'clients', ',' . $group_list['client'] . ','],
                    ['LIKE', 'clients', ',' . $group_list['client']]
                ],
                [
                    'OR',
                    ['=', 'user_limit', 0],
                    [
                        'AND',
                        ['=', 'user_limit', 1],
                        ['LIKE', 'users', $uid]
                    ]
                ]
            ]
        ];
        $where[] = [
            'OR',
            ['=', 'line_limit', 0],
            [
                'AND',
                ['=', 'line_limit', 1],
                [
                    'OR',
                    ['LIKE', 'lines', $group_list['line_id'] . ','],
                    ['LIKE', 'lines', ',' . $group_list['line_id'] . ','],
                    ['LIKE', 'lines', ',' . $group_list['line_id']]
                ]
            ]
        ];
        $where[] = [
            'OR',
            ['=', 'game_limit', 0],
            [
                'AND',
                ['=', 'game_limit', 1],
                [
                    'OR',
                    ['LIKE', 'games', $group_list['fc_type'] . ','],
                    ['LIKE', 'games', ',' . $group_list['fc_type'] . ','],
                    ['LIKE', 'games', ',' . $group_list['fc_type']]
                ]
            ]
        ];
        $where[] = [
            'OR',
            ['=', 'agent_limit', 0],
            [
                'AND',
                ['=', 'agent_limit', 1],
                [
                    'OR',
                    ['LIKE', 'agents', $group_list['agent_id'] . ','],
                    ['LIKE', 'agents', ',' . $group_list['agent_id'] . ','],
                    ['LIKE', 'agents', ',' . $group_list['agent_id']]
                ]
            ]
        ];
        $items = selfModel::getItems($where); // 获取 符合条件的 已发送的 消息
        $result = [];
        foreach ($items as $item) {
            $qcount = selfModel::getCount(['pid' => $item['id'], 'client' => $group_list['client'], 'uid' => $uid], true); // 查询是否已有队列记录
            if ($qcount <= 0) {
                $result['data'][] = $item;
            }
        }
        $result['ErrorCode'] = 1;
        return $result;
    }

    public function bind() {
        $request = Yii::$app->request->get();
        $client_id = isset($request['client_id']) ? trim($request['client_id']) : '';

        if (empty($client_id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数错误';
            return $result;
        }

        $_SESSION['client_id'] = $client_id;
        $result['ErrorCode'] = 1;
        return $result;
    }

    // 用户读取 两种方式:1.读一条插入一条,2.一次性插入所有队列,读取时修改状态，这里用的是第一种
    public function read() {
        $request = Yii::$app->request->get();
        $id = isset($request['id']) ? intval($request['id']) : '';
        $client = isset($request['client']) ? trim($request['client']) : '';
        $uid = isset($request['uid']) ? trim($request['uid']) : '';

        if (empty($id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '参数错误';
            return $result;
        }

        $values = [
            'pid' => $id,
            'client' => $client,
            'uid' => $uid,
            'addtime' => time(),
            'sendtime' => time(),
            'status' => 1
        ];
        $sqlQuery = selfModel::insert($values, true);

        if (!$sqlQuery) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '修改失败';
            return $result;
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '修改成功';
        return $result;
    }

    public function showData($type = '') {

        switch($type){
            case 'lines':
                $lines = parent::getLines();
                $data = ArrayHelper::index($lines, 'line_id');
                break;
            case 'games':
                $games = parent::getAllFcTypes();
                $data = ArrayHelper::index($games, 'type');
                break;
            default:
                $lines = parent::getLines();
                $games = parent::getAllFcTypes();

                $lines = ArrayHelper::index($lines, 'line_id');
                $games = ArrayHelper::index($games, 'type');

                $data['lines'] = $lines;
                $data['games'] = $games;
                $data['clients'] = ['admin' => '管理后台', 'agent' => '代理后台', 'front' => '前台'];
                break;
        }

        return $data;
    }

    public function transData() {

        $data['status'] = [0 => '未推送', 1 => '自动推送', 2 => '手动推送'];
        $data['enable'] = [0 => '已关闭', 1 => '已启用'];

        return $data;
    }

    public function trans($data, $showData) {
        // $showData = $this->showData();
        $transData = $this->transData();

        foreach ($data as $k => &$v) {
            $v['statusTxt'] = key_exists($v['status'], $transData['status']) ? $transData['status'][$v['status']] : $v['status'];
            $v['enableTxt'] = key_exists($v['enable'], $transData['enable']) ? $transData['enable'][$v['enable']] : $v['enable'];
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            $v['senddate'] = date('Y-m-d H:i:s', $v['sendtime']);
            $v['senddatetime'] = date('Y-m-d H:i:s', $v['sendtime']);
        }unset($v);

        return $data;
    }

}
