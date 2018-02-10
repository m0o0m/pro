<?php

namespace backend\controllers;

use Yii;
use common\models\LogModel;
use common\helpers\mongoTables;
use backend\controllers\Controller;

class LogController extends Controller {

    //翻译
    public function trans($data) {
        foreach ($data as $k => $v) {
            if (isset($v['old']) && is_array($v['old']))
                $data[$k]['old'] = json_encode($v['old']);
            if (isset($v['new']) && is_array($v['new']))
                $data[$k]['new'] = json_encode($v['new']);
            if (isset($v['ptype'])) {
                switch ($v['ptype']) {
                    case 1:
                        $data[$k]['ptypeTxt'] = '总后台';
                        break;
                    case 2:
                        $data[$k]['ptypeTxt'] = '后台';
                        break;
                    case 3:
                        $data[$k]['ptypeTxt'] = '前台PC';
                        break;
                    case 4:
                        $data[$k]['ptypeTxt'] = '前台WAP';
                        break;
                    default :
                        $data[$k]['ptypeTxt'] = '未知';
                        break;
                }
            }
            if (isset($v['addtime'])) {
                $data[$k]['addtime'] = date('Y-m-d H:i:s', $v['addtime']);
            }
            //登录状态
            if (isset($v['state'])) {
                if ($v['state'] == 1)
                    $data[$k]['newState'] = '登录成功';
                if ($v['state'] == 2)
                    $data[$k]['newState'] = '登录失败';
            }
            //中文名字
            if (isset($v['menu_name'])) {
                $data[$k]['menu_name'] = (string) $v['menu_name'];
            }
        }
        return $data;
    }

    /**
     * **********************************************************
     *  获取条件               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function getWhere($get, $where = array()) {
        $session = Yii::$app->session;
        $line_id = $session['line_id'];
        $user_type = $session['user_type'];
        $where['line_id'] = $line_id;
        // 根据登录者身份进行展示
        if (in_array($user_type, [3,4])) { //代理
            $where['uid'] = $session['agent_id'];
        }

        //平台
        if (isset($get['ptype']) && !empty($get['ptype'])) {
            $where['ptype'] = (int) trim($get['ptype']);
        }
        //日期
        $starttime = isset($get['starttime']) ? trim($get['starttime']) . ' 00:00:00' : '';
        $endtime = isset($get['endtime']) ? trim($get['endtime']) . '23:59:59' : '';
        $starttime = strtotime($starttime) ? strtotime($starttime) : '';
        $endtime = strtotime($endtime) ? strtotime($endtime) : '';
        if ($starttime && $endtime) {
            $where['addtime'] = ['$gt' => $starttime, '$lt' => $endtime];
        } elseif ($starttime) {
            $where['addtime'] = ['$gt' => $starttime];
        } elseif ($endtime) {
            $where['addtime'] = ['$lt' => $endtime];
        }
        return $where;
    }

    /**
     * **********************************************************
     *  获取页码                  @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function getPage($get) {
        $result = array();
        $result['pageNum'] = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $result['page'] = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $result['offset'] = ( $result['page'] - 1 ) * $result['pageNum'];
        return $result;
    }

    /**
     * **********************************************************
     *  历史登录日志             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionHistory() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('history.html');
            } else {
                return $this->render('history.html');
            }
        }

        // 判断权限
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if (!in_array($user_type, [1,2])) {
            $this->wrong_msg();
        }

        $page_data = $this->getPage($get);
        $where = [];
        //会员帐号
        if (isset($get['uname']) && !empty($get['uname'])) {
            $where['uname'] = trim($get['uname']);
        }
        $where = $this->getWhere($get, $where);
        //获取要查询的表名
        $mongoTable = mongoTables::getTable('historyLogin');
        $count = LogModel::getLogsCount($mongoTable, $where);
        if ($count <= 0) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('history.html');
            } else {
                return $this->render('history.html');
            }
        }

        $orderby = ['addtime' => SORT_DESC];
        $data = LogModel::getLogs($mongoTable, $where, $page_data['offset'], $page_data['pageNum'], $orderby);
        $data = $this->trans($data);
        $pagecount = ceil($count / $page_data['pageNum']);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page_data['page']
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('history.html', $render);
        } else {
            return $this->render('history.html', $render);
        }
    }

    //其它操作日志
    public function actionOperate() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('operate.html');
            } else {
                return $this->render('operate.html');
            }
        }

        // 判断权限
        $session = Yii::$app->session;
        $user_type = $session['user_type'];
        if (!in_array($user_type, [1,2])) {
            $this->wrong_msg();
        }

        $page_data = $this->getPage($get);
        $where = [];
        //操作人
        if (isset($get['uname']) && !empty($get['uname'])) {
            $where['uname'] = trim($get['uname']);
        }
        $where = $this->getWhere($get, $where);
        //获取要查询的表名
        $mongoTable = mongoTables::getTable('operate');
        $count = LogModel::getLogsCount($mongoTable, $where);
        if ($count <= 0) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('operate.html');
            } else {
                return $this->render('operate.html');
            }
        }

        $orderby = ['addtime' => SORT_DESC];
        $data = LogModel::getLogs($mongoTable, $where, $page_data['offset'], $page_data['pageNum'], $orderby);
        $data = $this->trans($data);
        $pagecount = ceil($count / $page_data['pageNum']);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page_data['page']
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('operate.html', $render);
        } else {
            return $this->render('operate.html', $render);
        }
    }

}
