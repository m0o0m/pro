<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\MaintainModel as selfModel;
use yii\helpers\ArrayHelper;

class MaintainController extends Controller {

    public function actionIndex() {
        $list = $this->getMaintain();
        $showData = $this->showData();

        $render = [
            'list' => $list
        ];
        $render = array_merge($render, $showData);
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    public function actionCreate() {
        $request = Yii::$app->request->post();
        $return = $this->setMaintain($request);

        $result = [];
        if ($return == 1) {
            $this->resetRedis($request);
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '修改成功';
        } elseif ($return == 0) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '修改失败';
        } elseif ($return == 2) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '事务失败';
        } elseif ($return == 3) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '彩种为空';
        } 
        return json_encode($result);
    }

    public function actionClose() {
        $request = Yii::$app->request->post();
        $return = $this->setMaintain($request);

        $result = [];
        if ($return == 1) {
            $this->resetRedis($request);
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '删除成功';
        } elseif ($return == 0) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '删除失败';
        } elseif ($return == 2) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '事务失败';
        } elseif ($return == 3) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '彩种为空';
        } 
        return json_encode($result);
    }

    public function resetRedis($request) {
        if ($request['mtype'] == 2 && in_array($request['ptype'], ['pc', 'wap'])) {
            $lines = self::showData('lines');
            foreach ($lines as $line) {
                Yii::$app->redis->HDEL("maintain_{$request['ptype']}_one_line_ids", $line['line_id']);
            }
        } else {
            Yii::$app->redis->DEL("maintain_{$request['ptype']}_all_line_ids");
        }
        $this->setRedis($request);
    }

    public function setRedis($request) {
        $data = $this->getMaintain();
        foreach ($data as $_mtype => $v) {
            foreach ($v as $_ptype => $v2) {
                if ($_mtype == 2 && in_array($_ptype, ['pc', 'wap'])) {
                    foreach ($v2['line_list'] as $_line_id => $v3) {
                        Yii::$app->redis->hset("maintain_{$_ptype}_one_line_ids", $_line_id, json_encode($v3));
                    }
                } else {

                    if (in_array($_ptype, ['api', 'spider'])) {
                        //不压缩
                        Yii::$app->redis->set("maintain_{$_ptype}_all_line_ids", json_encode($v2), false);
                    } else {
                        Yii::$app->redis->set("maintain_{$_ptype}_all_line_ids", json_encode($v2));
                    }
                }
            }
        }
    }

    public function setMaintain($request) {
        $result['ErrorCode'] = 2;
        if (empty($request['mtype'])) {
             $result['ErrorMsg'] = '维护类型不能为空';
             echo json_encode($result);die;
        }

        if ( empty($request['ptype']) ) {
            $result['ErrorMsg'] = '平台类型不能为空';
             echo json_encode($result);die;
        }
        if ( empty($request['starttime']) || empty($request['endtime']) ) {
            $result['ErrorMsg'] = '维护时间不能为空';
             echo json_encode($result);die;
        }
        if (  empty($request['line_id']) ) {
             $result['ErrorMsg'] = '线路id不能为空';
             echo json_encode($result);die;
        }

        $where['mtype'] = $request['mtype'];
        $where['ptype'] = $request['ptype'];

        $cols['starttime'] = isset($request['starttime']) ? strtotime($request['starttime']) : 0;
        $cols['endtime'] = isset($request['endtime']) ? strtotime($request['endtime']) : 0;
        $cols['remark'] = isset($request['remark']) ? $request['remark'] : '';
        if (isset($request['act']) && $request['act'] == 'close') {
            $cols['starttime'] = 0; // 清除开始时间
            $cols['endtime'] = 0; // 清除结束时间
            $cols['remark'] = ''; // 清除备注
        }

        if ($request['mtype'] == 2 && in_array($request['ptype'], ['pc', 'wap'])) {
            if (empty($request['fc_type'])) {
                return 3;
            }

            $transaction = \Yii::$app->manage_db->beginTransaction();

            $cols['fc_type'] = '';
            selfModel::_update($cols, $where); // 把该终端下所有线路的fc_type清空

            foreach ($request['line_id'] as $line_id) {

                $where['line_id'] = $line_id;

                $cols['fc_type'] = isset($request['fc_type'][$line_id]) ? implode(',', $request['fc_type'][$line_id]) : '';
                if (isset($request['act']) && $request['act'] == 'close') {
                    $cols['fc_type'] = ''; // 清除彩种
                }

                if (!self::editMaintain($cols, $where)) {
                    $transaction->rollBack();
                    return 2;
                }
            }
            $transaction->commit();
            return true;
        } else {

            $cols['line_id'] = implode(',', $request['line_id']);
            if (isset($request['act']) && $request['act'] == 'close') {
                $cols['line_id'] = ''; // 清除线路
            }

            $cols['fc_type'] = '';

            return self::editMaintain($cols, $where);
        }
    }

    public function editMaintain($cols, $where) {
        $has = selfModel::get_count($where);
        $timestamp = time();
        if ($has == 0) {
            $cols['addtime'] = $timestamp;
            $cols['updatetime'] = $timestamp;
            $cols = array_merge($cols, $where);
            return selfModel::_insert($cols);
        } else {
            $cols['updatetime'] = $timestamp;
            return selfModel::_update($cols, $where);
        }
    }

    public function showData($type = '') {

        switch ($type) {
            case 'ptypes':
                $data = ['pc' => 'PC', 'wap' => 'WAP', 'agent' => '客户后台', 'spider' => '采集', 'api' => 'API'];
                break;
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
                $data['ptypes'] = ['pc' => 'PC', 'wap' => 'WAP', 'agent' => '客户后台', 'spider' => '采集', 'api' => 'API'];
                break;
        }

        return $data;
    }

    /**
     * 获取维护信息
     * @return 全网维护: data[1][$ptype]['line_id']
     *         一般维护: data[2][pc|wap]['fc_type'][$line_id]
     * @author Frank
     */
    public function getMaintain() {
        $maintain_list = selfModel::get_all();
        $data = [];
        foreach ($maintain_list as $m) {
            if (!empty($m['fc_type'])) {
                $m['fc_type'] = explode(',', $m['fc_type']);
            }
            $m['startdate'] = !empty($m['starttime']) ? date('Y-m-d H:i:s', $m['starttime']) : '';
            $m['enddate'] = !empty($m['endtime']) ? date('Y-m-d H:i:s', $m['endtime']) : '';
            if ($m['mtype'] == 2 && in_array($m['ptype'], ['pc', 'wap'])) {
                if (!empty($m['fc_type'])) {
                    $data[$m['mtype']][$m['ptype']]['line_list'][$m['line_id']] = $m;
                    $data[$m['mtype']][$m['ptype']]['line_id'][] = $m['line_id'];
                    $data[$m['mtype']][$m['ptype']]['fc_type'][$m['line_id']] = $m['fc_type'];
                    $data[$m['mtype']][$m['ptype']] = array_merge($m, $data[$m['mtype']][$m['ptype']]);
                    unset($data[$m['mtype']][$m['ptype']]['id']);
                }
            } else {
                if (!empty($m['line_id'])) {
                    $m['line_id'] = explode(',', $m['line_id']);
                    $data[$m['mtype']][$m['ptype']] = $m;
                }
            }
        }
        return $data;
    }

}
