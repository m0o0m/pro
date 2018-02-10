<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use common\models\StatModel as selfModel;
use common\models\AgentModel;
use common\models\UserModel;
use yii\helpers\ArrayHelper;

class StatController extends Controller {

    public function actionIndex() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        $data['show'] = $this->showData();

        $request = Yii::$app->request->get();
        $request['line_id'] = Yii::$app->session->get('line_id'); // 所在线路
        if($user_type == 3 || $user_type == 4){
            $request['at_id'] = Yii::$app->session->get('agent_id'); // 代理
        }
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

    public function actionLine() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if($user_type != 1 && $user_type != 2){
            $this->wrong_msg();
        }
        $data['show'] = $this->showData();

        $request = Yii::$app->request->get();
        $request['tab'] = 'line';
        $request['line_id'] = Yii::$app->session->get('line_id'); // 所在线路
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


    public function actionTask() {
        if ($this->task_is_running()) {
            $result['ErrorMsg'] = '任务正在运行中，请稍后再试';
            $result['ErrorCode'] = 2;
        } else {
            $result['ErrorMsg'] = 'OK';
            $result['ErrorCode'] = 1;
        }

        return json_encode($result);
    }

    public function actionExport() {
        $request = Yii::$app->request->get();
        $result = $this->search($request, $this->showData());
        $data = $result['data'];

        if (empty($data)) {
            echo '数据为空';
            return;
        }

        $tab = isset($request['tab']) ? trim($request['tab']) : '';
        switch($tab){
            case 'line':
                $filename = '线路注单统计报表' . '-' . date('Y_m_d-H_i_s');
                break;
            default:
                $filename = '会员注单统计报表' . '-' . date('Y_m_d-H_i_s');
                break;
        }

        $fields = [
            'id' => 'ID',
            'fc_typeTxt' => '彩种',
            'line_id' => '线路',
            // 'sh_id'=>'股东ID',
            // 'ua_id'=>'总代ID',
            'at_id' => '代理ID',
            'agentname' => '代理账号',
            'uid' => '会员ID',
            'username' => '会员账号',
            'bet' => '总注单',
            'bet_count' => '总笔数',
            'valid_bet' => '有效注单',
            'valid_bet_count' => '有效笔数',
            'win' => '赢金额',
            'win_count' => '赢笔数',
            'addday' => '注单日期',
            'updatedate' => '最近更新时间'
        ];

        $HTML = $this->getExportContent($data, $fields);

        header("Content-type:application/vnd.ms-excel;charset=UTF-8");
        header("Content-Disposition:attachment;filename=" . $filename . ".xls");
        echo $HTML;
        return;
    }

    public function actionChart() {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        $data['show'] = $this->showData();
        $request = Yii::$app->request->get();
        $request['line_id'] = Yii::$app->session->get('line_id'); // 所在线路
        if($user_type == 3 || $user_type == 4){
            $request['at_id'] = Yii::$app->session->get('agent_id'); // 代理
        }
        $stat = $this->getStat($request, $data['show']);

        $chart = [];
        if (!empty($stat)) {
            $chart['xaxis_categories'] = array_keys($stat);
            $chart['series_type'] = 'column'; // line/spline/column 直线图/曲线图/柱状图
            $chart['title'] = '注单走势图';
            $chart['subtitle'] = date('Y-m-d H:i');
            $chart['yaxis_title'] = '注单量 (元)';
            $chart['tooltip_valuesuffix'] = '元';
            $chart['tooltip_valuesuffix2'] = '笔';

            $chart_type = isset($request['chart_type']) ? trim($request['chart_type']) : '';
            switch ($chart_type) {
                case 'bar':
                    $chart['height'] = count($chart['xaxis_categories']) * 120;
                    break;
                case 'column':
                    $chart['width'] = count($chart['xaxis_categories']) * 120;
                    break;
                default:
                    $chart['width'] = count($chart['xaxis_categories']) * 120;
                    break;
            }

            $col = [
                'bet' => '总注单',
                'bet_count' => '总笔数',
                'valid_bet' => '有效注单',
                'valid_bet_count' => '有效笔数',
                'win' => '赢金额',
                'win_count' => '赢笔数'
            ];
            foreach ($col as $key => $val) {
                $chart['series'][$key]['name'] = $val;
                $chart['series'][$key]['data'] = array_column($stat, $key, 'index');
            }
        }

        $render = [
            'data' => $data,
            'chart' => $chart
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('chart.html', $render);
        } else {
            return $this->render('chart.html', $render);
        }
    }

    /*
     * *************************************************************************
     */

    public function search($request, $showData) {
        $query = $request;
        unset($query['_pjax']);
        unset($query['tab']);
        unset($query['line_id']);
        unset($query['at_id']);
        if (empty($query)) {
            return;
        }

        $tab = isset($request['tab']) ? trim($request['tab']) : '';
        $line_id = isset($request['line_id']) ? trim($request['line_id']) : '';
        $agentname = isset($request['agentname']) ? trim($request['agentname']) : '';
        $at_id = isset($request['at_id']) ? trim($request['at_id']) : '';
        $username = isset($request['username']) ? trim($request['username']) : '';
        $fc_type = isset($request['fc_type']) ? trim($request['fc_type']) : '';
        $order_sort = isset($request['order_sort']) ? trim($request['order_sort']) : '';
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $page = isset($request['page']) ? intval($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? intval($request['pagenum']) : 100;

        // 时间 处理
        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;

        // 查询条件
        $where[] = 'and';
        if ($line_id) {
            $where[] = ['=', 'line_id', $line_id];
        }
        if($at_id){
            $where[] = ['=', 'at_id', $at_id];
        }
        if ($agentname) {
            $where[] = ['=', 'at_username', $agentname];
        }
        if ($username) {
            $where[] = ['=', 'uname', $username];
        }
        if ($fc_type) {
            $where[] = ['=', 'fc_type', $fc_type];
        }
        $orderby = ''; // 排序
        if ($order_sort) {
            list($order, $sort) = explode('-', $order_sort);
            $orderby = "$order $sort";
        } else {
            $orderby = "addday desc";
        }
        // 时间筛选
        if ($starttime) {
            $startdate = date('Ymd', $starttime);
            $where[] = ['>=', 'addday', $startdate];
        }
        if ($endtime) {
            $enddate = date('Ymd', $endtime);
            $where[] = ['<=', 'addday', $enddate];
        }

        // 分页
        if ($tab == 'line') {
            $recordcount = selfModel::get_count_line($where);
        } else {
            $recordcount = selfModel::get_count($where);
        }
        $pagecount = ceil($recordcount / $pagenum);
        if ($page > $pagecount && $pagecount != 0)
            $page = $pagecount;
        if ($page <= 0)
            $page = 1;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        if ($tab == 'line') {
            $data = selfModel::get_list_line($where, $offset, $limit, $orderby);
        } else {
            $data = selfModel::get_list($where, $offset, $limit, $orderby);
        }
        $data = $this->trans($data, $showData, $tab); // 翻译
        $total = $this->getTotalData($where,$data, $tab); //统计
        $result['data'] = $data;
        $result['page'] = $page;
        $result['pagenum'] = $pagenum;
        $result['pagecount'] = $pagecount;
        $result['export_query'] = $request;
        $result['total'] = $total;
        return $result;
    }

    public function getExportContent($data, $fields) {
        list($first) = $data;
        foreach ($fields as $fieldkey => $fieldname) {
            if (!key_exists($fieldkey, $first))
                unset($fields[$fieldkey]);
        }
        $tab = 'style="width:100%;font-size:14px;"';
        $tab_tr = 'style="height:36px;border-bottom:1px solid;"';
        $tab_tr_th = 'style="text-align:center;"';
        $tab_tr_td = 'style="text-align:center;"';
        $HTML = '<meta http-equiv=Content-Type content="text/html; charset=utf-8">' . PHP_EOL;
        $HTML .= '<table ' . $tab . '>' . PHP_EOL;
        $HTML .= '<tr ' . $tab_tr . '>';
        foreach ($fields as $fieldkey => $fieldname) {
            $HTML .= '<th ' . $tab_tr_th . '>' . $fieldname . '</th>';
        }
        $HTML .= '</tr>' . PHP_EOL;

        foreach ($data as $v) {
            $HTML .= '<tr ' . $tab_tr . '>';
            foreach ($fields as $fieldkey => $fieldname) {
                $tab_tr_td_win = ($fieldkey == 'win') ? ($v[$fieldkey] > 0) ? 'style="color:red;"' : 'style="color:green;"' : '';
                $HTML .= '<td ' . $tab_tr_td . $tab_tr_td_win . '>' . $v[$fieldkey] . '</td>';
            }
            $HTML .= '</tr>' . PHP_EOL;
        }
        $HTML .= '</table>';

        return $HTML;
    }

    public function getStat($request, $showData) {

        $display = isset($request['display']) ? trim($request['display']) : '';
        if (in_array($display, [1,2]) && in_array(Yii::$app->session->get('user_type'), [1,2]))
            $request['tab'] = 'line'; // 显示类型为线路/彩种 则查线路表
        $result = $this->search($request, $showData);

        $stat = [];
        if ($result['data']) {
            foreach ($result['data'] as $key => $val) {
                switch ($display) {
                    case '1':
                        $ckey = $val['line_id'];
                        break;
                    case '2':
                        $ckey = key_exists($val['fc_type'], $showData['games']) ? $showData['games'][$val['fc_type']]['name'] : $val['fc_type'];
                        break;
                    case '3':
                        $ckey = key_exists($val['at_id'], $showData['agents']) ? $showData['agents'][$val['at_id']]['login_name'] : $val['at_id'];
                        break;
                    case '4':
                        $ckey = key_exists($val['uid'], $showData['users']) ? $showData['users'][$val['uid']]['uname'] : $val['uid'];
                        break;
                    default:
                        $ckey = $val['line_id'];
                        break;
                }
                @$stat[$ckey]['index'] = $ckey;
                @$stat[$ckey]['bet'] += $val['bet'];
                @$stat[$ckey]['bet_count'] += $val['bet_count'];
                @$stat[$ckey]['valid_bet'] += $val['valid_bet'];
                @$stat[$ckey]['valid_bet_count'] += $val['valid_bet_count'];
                @$stat[$ckey]['win'] += $val['win'];
                @$stat[$ckey]['win_count'] += $val['win_count'];
            }
        }
        return $stat;
    }

    public function showData($type = '') {

        switch ($type) {
            case 'lines':
                $lines = parent::getLines();
                $data = ArrayHelper::index($lines, 'line_id');
                break;
            case 'games':
                $games = parent::getAllFcTypes();
                $data = ArrayHelper::index($games, 'type');
                break;
            case 'agents':
                $agents = AgentModel::getAgent('at');
                $data = ArrayHelper::index($agents, 'id');
                break;
            case 'users':
                $users = UserModel::get_items([]);
                $data = ArrayHelper::index($users, 'uid');
                break;
            default:
                $lines = parent::getLines();
                $games = parent::getAllFcTypes();
                $agents = AgentModel::getAgent('at');
                $users = UserModel::get_items([]);

                $lines = ArrayHelper::index($lines, 'line_id');
                $games = ArrayHelper::index($games, 'type');
                $agents = ArrayHelper::index($agents, 'id');
                $users = ArrayHelper::index($users, 'uid');

                $data['lines'] = $lines;
                $data['games'] = $games;
                $data['agents'] = $agents;
                $data['users'] = $users;
                break;
        }

        return $data;
    }

    public function trans($data, $showData, $tab) {
        // $showData = $this->showData();

        foreach ($data as $k => &$v) {
            if (isset($v['fc_type']))
                $v['fc_typeTxt'] = key_exists($v['fc_type'], $showData['games']) ? $showData['games'][$v['fc_type']]['name'] : $v['fc_type'];
            if (isset($v['uid']))
                $v['username'] = key_exists($v['uid'], $showData['users']) ? $showData['users'][$v['uid']]['uname'] : $v['uid'];
            if (isset($v['at_id']))
                $v['agentname'] = key_exists($v['at_id'], $showData['agents']) ? $showData['agents'][$v['at_id']]['login_name'] : $v['at_id'];
            if (isset($v['addtime']))
                $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            if (isset($v['updatetime']))
                $v['updatedate'] = date('Y-m-d H:i:s', $v['updatetime']);
            if($tab == 'line'){
                $data[$k]['money'] = number_format($v['valid_bet'] - $v['win'], 2, '.', ',');
            }else{
                $data[$k]['money'] = number_format($v['win'] - $v['valid_bet'], 2, '.', ',');
            }
        }unset($v);

        return $data;
    }
/**
      ***********************************************************
      *  统计当前页与所有符合条件的数据 @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function getTotalData($where, $data, $tab = ''){
       $bet = $bet_count = $valid_bet = $valid_bet_count = $win = $win_count = 0;
       $all_bet = $all_bet_count = $all_valid = $all_valid_count = $all_win = $all_win_count = 0;
       if(!empty($data)){
           foreach($data as $key=>$val){
                $bet += $val['bet'];
                $bet_count += $val['bet_count'];
                $valid_bet += $val['valid_bet'];
                $valid_bet_count += $val['valid_bet_count'];
                $win += $val['win'];
                $win_count += $val['win_count'];
           }
       }

       $all_bet = selfModel::getSum('SUM(bet) ', $where, $tab);
       $all_bet_count = selfModel::getSum('SUM(bet_count)', $where, $tab);
       $all_valid = selfModel::getSum('SUM(valid_bet)', $where, $tab);
       $all_valid_count = selfModel::getSum('SUM(valid_bet_count)', $where, $tab);
       $all_win = selfModel::getSum('SUM(win)', $where, $tab);
       $all_win_count = selfModel::getSum('SUM(win_count)', $where, $tab);

       $return = array();
       $return['bet'] = number_format($bet, '2', '.', ',');
       $return['bet_count'] = number_format($bet_count, '0', '.', ',');
       $return['valid_bet'] = number_format($valid_bet, '2', '.', ',');
       $return['valid_bet_count'] = number_format($valid_bet_count, '0', '.', ',');
       $return['win'] = number_format($win, '2', '.', ',');
       $return['win_count'] = number_format($win_count, '0', '.', ',');
       $return['all_bet'] = number_format($all_bet, '2', '.', ',');
       $return['all_bet_count'] = number_format($all_bet_count, '0', '.', ',');
       $return['all_valid'] = number_format($all_valid, '2', '.', ',');
       $return['all_valid_count'] = number_format($all_valid_count, '0', '.', ',');
       $return['all_win'] = number_format($all_win, '2', '.', ',');
       $return['all_win_count'] = number_format($all_win_count, '0', '.', ',');
       if($tab == 'line'){
           $return['money'] = number_format($valid_bet - $win, '0', '.', ',');
           $return['all_money'] = number_format($all_valid - $all_win, '0', '.', ',');
       }else{
           $return['money'] = number_format($win - $valid_bet, '0', '.', ',');
           $return['all_money'] = number_format($all_win - $all_valid, '0', '.', ',');
       }
       return $return;
    }

}
