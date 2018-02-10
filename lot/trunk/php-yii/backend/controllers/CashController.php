<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use common\models\CashModel;
use common\models\LineModel;
use common\models\AgentModel;

class CashController extends Controller {

    /**
     * **********************************************************
     *  翻译                     @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function trans($arr, $trans_arr) {
          $cash_type = array('', '存入', '取出', '');
        $cash_do_type = array('', '彩票下注', '彩票派彩', '彩票和局返还本金', 'api转入', 'api转出', '回滚', '');
        $ptype = array('', 'PC', 'WAP', 'APP', '');

        foreach ($arr as $key => $val) {
            if (in_array('cash_type', $trans_arr))
                $arr[$key]['cash_type_txt'] = $cash_type[$val['cash_type']];
            if (in_array('cash_do_type', $trans_arr))
                $arr[$key]['cash_do'] = $cash_do_type[$val['cash_do_type']];
            if (in_array('ptype', $trans_arr))
                $arr[$key]['ptype_txt'] = $ptype[$val['ptype']];
            if (in_array('time', $trans_arr))
                $arr[$key]['time'] = date('Y-m-d H:i:s', $val['addtime']);
            //下面的数组是备注翻译
            if (in_array('remark', $trans_arr) && isset($val['remark'])) {
                $remark_trans = array(
                    'Lottery note' => '订单号',
                    'Lottery note#' => '订单号',
                    'GOBACK Lottery note#' => '回滚订单号',
                    '#typesof#' => '彩种',
                    '#typesof#:#' => '彩种:',
                    'type' => '彩种',
                    "'" => ' ',
                    '.' => ' '
                );
                $tmp_remark = strtr($val['remark'], $remark_trans);

                $games = $this->getAllFcTypes();
                $new_games = array();
                foreach ($games as $k => $v) {
                    $new_games[$v['type']] = $v['name'];
                }
                $tmp_html = $tmp_str = '';
                $tmp_remark = strtr($tmp_remark, $new_games);
                $tmp_remark = explode(',', $tmp_remark);
                foreach($tmp_remark as $v){
                    $tmp_str .=  '<tr><td>' . $v . '</td></tr>';
                }
                $tmp_html = explode(',', $val['dids']);
                $tmp_str .= '<tr><td>祥细订单号（下方）：</td></tr>';
                foreach($tmp_html as $k => $v){
                    $tmp_str .=  '<tr><td>' . '第' . ($k+1) . '单: ' . $v . '</td></tr>';
                }
                $arr[$key]['remark_txt'] = $tmp_str;
            }
        }
        return $arr;
    }

    /**
     * **********************************************************
     *  会员现金记录             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUserrecord() {
        $session = Yii::$app->session;
        $get = Yii::$app->request->get();
        $line_id = $session['line_id'];
        $user_type = $session['user_type'];

        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'user_type' => $user_type,
                'pagecount' => 1,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('userRecord.html', $render);
            } else {
                return $this->render('userRecord.html', $render);
            }
        }

        $ptype = isset($get['ptype']) ? $get['ptype'] : null;
        $cash_type = isset($get['cash_type']) ? $get['cash_type'] : null;
        $cash_do = isset($get['cash_do_type']) ? $get['cash_do_type'] : null;
        $uname = isset($get['uname']) ? $get['uname'] : null;
        $agentName = isset($get['agentName']) ? $get['agentName'] : null;
        $page = isset($get['page']) ? $get['page'] : 1;
        $pagenum = isset($get['pagenum']) ? $get['pagenum'] : 100;
        $offset = ($page - 1) * $pagenum;
        $starttime = isset($get['starttime']) ? $get['starttime'] : null;
        $endtime = isset($get['endtime']) ? $get['endtime'] : null;
        if (!empty($starttime)) {
            $starttime = strtotime($starttime) ? strtotime($starttime . ' 00:00:00') : null;
        }
        if (!empty($endtime)) {
            $endtime = strtotime($endtime) ? strtotime($endtime . ' 23:59:59') : null;
        }

        $table = 'user_cash_record';
        $field = '*';
        //根据登录者身份进行展示
        $where = $this->loginWhere();
        if ($ptype) {
            if (in_array($ptype, array(1, 2, 3)))
                $where['ptype'] = $ptype;
        }

        if ($cash_type) {
            if (in_array($cash_type, array(1, 2)))
                $where['cash_type'] = $cash_type;
        }

        if ($cash_do) {
            if (in_array($cash_do, array(1, 2, 3, 4, 5, 6)))
                $where['cash_do_type'] = $cash_do;
        }

        if ($uname)
            $where['uname'] = $uname;
       
        if ($starttime && $endtime) {
            $where = ['and', $where, array('between', 'addtime', $starttime, $endtime)];
        } elseif (!empty($starttime)) {
            $where = ['and', $where, 'addtime>=' . $starttime];
        } elseif (!empty($endtime)) {
            $where = ['and', $where, array('<=', 'addtime', $endtime)];
        }

        $where = ['and', $where, ['line_id' => $line_id]];
        $count = CashModel::getCount($table, $where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        if ($count) {
            $data = CashModel::getData($table, $field, $where, $offset, $pagenum);
        }else {
            $data = array();
        }
        $data = $this->trans($data, array('cash_type', 'cash_do_type', 'time', 'ptype', 'remark'));
        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page,
            'user_type' => $user_type
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('userRecord.html', $render);
        } else {
            return $this->render('userRecord.html', $render);
        }
    }

    /**
     * **********************************************************
     *  代理现金记录             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionAgentrecord() {
        $get = Yii::$app->request->get();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('agentRecord.html', $render);
            } else {
                return $this->render('agentRecord.html', $render);
            }
        }
        $cash_type = isset($get['cash_type']) ? $get['cash_type'] : null;
        $page = isset($get['page']) ? $get['page'] : 1;
        $pagenum = isset($get['pagenum']) ? $get['pagenum'] : 100;
        $offset = ($page - 1) * $pagenum;
        $starttime = isset($get['starttime']) ? $get['starttime'] : null;
        $endtime = isset($get['endtime']) ? $get['endtime'] : null;
        if (!empty($starttime)) {
            $starttime = strtotime($starttime) ? strtotime($starttime . ' 00:00:00') : null;
        }
        if (!empty($endtime)) {
            $endtime = strtotime($endtime) ? strtotime($endtime . ' 23:59:59') : null;
        }
        $table = 'agent_cash_record';
        $field = '*';
        //根据登录者身份进行展示
        $where = $this->loginWhere();

        if ($cash_type) {
            if (in_array($cash_type, array(1, 2)))
                $where['cash_type'] = $cash_type;
        }

       
        if ($starttime && $endtime) {
            $where = ['and', $where, array('between', 'addtime', $starttime, $endtime)];
        } elseif (!empty($starttime)) {
            $where = ['and', $where, 'addtime>=' . $starttime];
        } elseif (!empty($endtime)) {
            $where = ['and', $where, array('<=', 'addtime', $endtime)];
        }

        $count = CashModel::getCount($table, $where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        if ($count) {
            $data = CashModel::getData($table, $field, $where, $offset, $pagenum);
        } else {
            $data = array();
        }
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $data = $this->trans($data, array('cash_type', 'time'));
        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('agentRecord.html', $render);
        } else {
            return $this->render('agentRecord.html', $render);
        }
    }

    public function actionShrecord() {
        $agents = AgentModel::getMoreAgentByCondition(Yii::$app->manage_db, 'my_agent_admin', []);
        $atype = 'sh';
        $render = self::record($atype);
        $render['atype'] = $atype;
        $render['agents'] = $agents;
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('record.html', $render);
        } else {
            return $this->render('record.html', $render);
        }
    }

    public function actionUarecord() {
        $agents = AgentModel::getMoreAgentByCondition(Yii::$app->manage_db, 'my_agent_admin', []);

        $atype = 'ua';
        $render = self::record($atype);
        $render['atype'] = $atype;
        $render['agents'] = $agents;
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('record.html', $render);
        } else {
            return $this->render('record.html', $render);
        }
    }

    public function record($atype) {
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if($user_type != 1 && $user_type != 2){
            $this->wrong_msg();
        }
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            return;
        }
        $agent_name = isset($get['agent_name']) ? $get['agent_name'] : null;
        $cash_type = isset($get['cash_type']) ? $get['cash_type'] : null;
        $starttime = isset($get['starttime']) ? $get['starttime'] : null;
        $endtime = isset($get['endtime']) ? $get['endtime'] : null;
        $page = isset($get['page']) ? $get['page'] : 1;
        $pagenum = isset($get['pagenum']) ? $get['pagenum'] : 100;

       if (!empty($starttime)) {
            $starttime = strtotime($starttime) ? strtotime($starttime . ' 00:00:00') : null;
        }
        if (!empty($endtime)) {
            $endtime = strtotime($endtime) ? strtotime($endtime . ' 23:59:59') : null;
        }

        $table = $atype . '_cash_record';
        $where = [];
        $where = $this->loginWhere();
        if ($agent_name) {
            $where[$atype . '_user'] = $agent_name;
        }
        if ($cash_type) {
            if (in_array($cash_type, [1, 2]))
                $where['cash_type'] = $cash_type;
        }

        if ($starttime && $endtime) {
            $where = ['and', $where, ['between', 'addtime', $starttime, $endtime]];
        } elseif (!empty($starttime)) {
            $where = ['and', $where, 'addtime>=' . $starttime];
        } elseif (!empty($endtime)) {
            $where = ['and', $where, ['<=', 'addtime', $endtime]];
        }

        // 分页
        $total_count = CashModel::getCount($table, $where);
        $total_page = ceil($total_count / $pagenum);
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        if ($page <= 0)
            $page = 1;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        if ($total_count) {
            $data = CashModel::getData($table, '*', $where, $offset, $limit);
        } else {
            $data = [];
        }
        $data = $this->trans($data, ['cash_type', 'time']);

        $render = [
            'data' => $data,
            'page' => $page,
            'total_page' => $total_page
        ];
        return $render;
    }

}
