<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\TaocanModel as selfModel;
//use yii\db\Query;
use yii\mongodb\Query;

class TaocanController extends Controller {

    /**
     * **********************************************************
     *  套餐管理列表             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $t_name = isset($get['t_name']) ? $get['t_name'] : '';
        $t_money = isset($get['t_money']) ? $get['t_money'] : '';

        $linelist = $this->getLines();
        $tmp_list = selfModel::get_set_all([]); //关联数据

        $where = [];
        if (!empty($t_name))
            $where['tname'] = $t_name;
        if (!empty($t_money))
            $where['pktc'] = $t_money;
        $tmp_tlist = selfModel::get_all($where); //套餐数据
        $arr = array();
        $list = array();
        $tlist = array();
        $line_count = array();
        //套餐列表 tid作键
        if (!empty($tmp_tlist)) {
            foreach ($tmp_tlist as $key => $val) {
                $tlist[$val['id']] = $val;
            }
        }
        //关联列表 线路id作键
        if (!empty($tmp_list)) {
            foreach ($tmp_list as $key => $val) {
                $list[$val['line_id']] = $val;
                $line_count[$val['tid']][] = $val['line_id'];
            }
        }

        foreach ($linelist as $key => $val) {
            $arr[$key]['line_id'] = $val['line_id']; //线路id
            if (isset($list[$val['id']])) {
                $arr[$key]['id'] = $val['id']; //线路数字id
                $arr[$key]['tid'] = $list[$val['id']]['tid']; //套餐id
                $arr[$key]['color'] = $tlist[$list[$val['id']]['tid']]['color']; //套餐颜色
                $arr[$key]['name'] = $tlist[$list[$val['id']]['tid']] ['tname'];
            } else {
                $arr[$key]['id'] = $val['id'];
                $arr[$key]['tid'] = 0;
                $arr[$key]['color'] = '#4f990e';
            }
        }

        foreach ($tmp_tlist as $key => $tlist1) {
            $data['tid'] = $tlist1['id'];
            $tmp_tlist[$key]['num'] = isset($line_count[$tlist1['id']]) ? count($line_count[$tlist1['id']]) : 0;
        }

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', array('tlist' => $tmp_tlist, 'list' => $arr));
        } else {
            return $this->render('index.html', array('tlist' => $tmp_tlist, 'list' => $arr));
        }
    }

    //新增套餐
    public function actionInsert() {
        $post = Yii::$app->request->post();
        $result = array();
        $result['ErrorCode'] = 2;
        $data['tname'] = isset($post['tname']) ? $post['tname'] : null;
        $data['pktc'] = isset($post['pktc']) ? $post['pktc'] : null;
        $data['add_time'] = date('Y-m-d H:i:s', time());
        $data['color'] = isset($post['color']) ? $post['color'] : null;
        ;
        if (empty($data['tname'])) {
            $result['ErrorMsg'] = '请输入套餐名称';
            echo json_encode($result);
            die;
        }
        if (empty($data['pktc'])) {
            $result['ErrorMsg'] = '请输入套餐比例';
            echo json_encode($result);
            die;
        }
        $count = selfModel::get_count(['tname' => $data['tname']]);
        if ($count) {
            $result['ErrorMsg'] = '该套餐已经存在';
            echo json_encode($result);
            die;
        }
        $res = selfModel::_insert($data);
        if (!$res) {
            $result['ErrorMsg'] = '添加失败';
            echo json_encode($result);
            die;
        }
        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '添加成功';
        return json_encode($result);
    }

    //获取套餐信息
    public function actionGetone() {
        $post = Yii::$app->request->post();
        $where['id'] = isset($post['id']) ? $post['id'] : 1;
        $result = selfModel::get_one($where);
        return json_encode($result);
    }

    //套餐修改
    public function actionEdit() {
        $post = Yii::$app->request->post();
        $result['ErrorCode'] = 2;
        $data['tname'] = isset($post['tname']) ? $post['tname'] : '';
        $data['pktc'] = isset($post['pktc']) ? $post['pktc'] : '';
        $data['color'] = isset($post['color']) ? $post['color'] : '';
        $data['update_time'] = date('Y-m-d H:i:s');
        if (empty($data['tname'])) {
            $result['ErrorMsg'] = '请输入套餐名称';
            echo json_encode($result);
            die;
        }
        if (empty($data['pktc'])) {
            $result['ErrorMsg'] = '请输入套餐比例';
            echo json_encode($result);
            die;
        }
        $where['id'] = isset($post['id']) ? $post['id'] : '';
        $res = selfModel::_update($data, $where);
        if (!$res) {
            $result['ErrorMsg'] = '修改失败';
            echo json_encode($result);
            die;
        }
        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '修改成功';
        return json_encode($result);
    }

    //站点套餐修改
    public function actionChange() {
        $post = Yii::$app->request->post();
        $site = isset($post['site']) ? $post['site'] : '';
        if (empty($post['tid']))
            return 0;
        $where = array();
        if (empty($site)) {//取消该套餐所有站点
            $where['tid'] = $post['tid'];
            selfModel::set_del($where);
            return true;
        }
        //删除原有该套餐信息，将新的信息添加进去
        $database = \Yii::$app->manage_db;
        $transaction = $database->beginTransaction();
        $where = array();
        $where['line_id'] = $site;
        $count = selfModel::get_set_count($where);
        if ($count > 0) {
            $res = selfModel::set_del($where);
            if (!$res) {
                $transaction->rollBack();
                return 0;
            }
        }
        selfModel::set_del(['tid' => $post['tid']]);
        foreach ($site as $val) {
            $data = array();
            $data['tid'] = $post['tid'];
            $data['line_id'] = $val;
            $res = selfModel::set_insert($data);
            if (!$res) {
                $transaction->rollBack();
                return 0;
            }
        }
        $transaction->commit();
        return true;
    }

}
