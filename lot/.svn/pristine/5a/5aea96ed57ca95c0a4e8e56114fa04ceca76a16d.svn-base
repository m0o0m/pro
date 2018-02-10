<?php

namespace our_backend\controllers\gameapi;

use Yii;
use our_backend\controllers\Controller;
use common\helpers\Helper;
use common\models\OpentimeModel as otime;

class OpentimeController extends Controller {

    /**
     * **********************************************************
     *  开盘时间列表展示         @author  Rom  *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $page = isset($get['page']) ? intval($get['page']) : 1;
        $pagenum = isset($get['pageNum']) ? intval($get['pageNum']) : 100;
        $fc_type = isset($get['fc_type']) ? $get['fc_type'] : null;
        $qishu = isset($get['qishu']) ? $get['qishu'] : null;
        //判断分页是否正确
        if (!is_int($page) || $page <= 0) {
            $page = 1;
        }
        if (!is_int($pagenum) || $pagenum == 0) {
            $pagenum = 100;
        } else {
            $pagenum = abs($pagenum);
        }

        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $key => $val) {
            $tmp[$val['type']] = $val;
        }
        $games = $tmp;

        $tabname = Helper::GetOpenTimeTableNameByType($fc_type); //获取表名

        $where = array();
        if ($qishu && is_numeric($qishu) && $qishu > 0) {
            $where['qishu'] = $qishu;
        }
        if (!in_array($fc_type, ['liuhecai', 'jnd_bs', 'jnd_28'])) {
            $where['fc_type'] = $fc_type;
        }
        $totalnum = otime::getDataTotalNum($tabname, $where); //获取当前彩种的总条数
        $total_page = ceil($totalnum / $pagenum); //总页码数
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        $condition['offset'] = ($page - 1) * $pagenum;  //开始条数
        $condition['limit'] = $pagenum;
        $condition['tabname'] = $tabname;
        $condition['field'] = array('id', 'qishu', 'kaipan', 'fengpan', 'kaijiang', 'status');
        $data = otime::index($condition, $where);

        foreach ($data as $k => $v) {
            $data[$k]['fc_type'] = $games[$fc_type]['name'];
        }
        $pagedata['page'] = $page;
        $pagedata['totalpage'] = $total_page;
        $pagedata['totalnum'] = $totalnum;

        $render = ['data' => $data, 'games' => $games, 'pagedata' => $pagedata, 'fc_type' => $fc_type];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    /**
     * 编辑页面
     */
    public function actionOpen_edit_info() {
        $post = Yii::$app->request->post();
        $error = false;
        $id = isset($post['id']) && !empty($post['id']) ? $post['id'] : 0;
        if (empty($post['fc_type']) || !isset($post['fc_type'])) {
            $error = true;
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '请选择彩种';
        }
        if ($error) {
            echo json_encode(['code' => $code]);
            die;
        }
        $where['id'] = $id;
        $tabname = Helper::GetOpenTimeTableNameByType($post['fc_type']); //获取表名
        $data = otime::getOneSql($where, $tabname);
        if ($data) {
            $code['ErrorCode'] = 1;
            $code['ErrorMsg'] = '获取编辑数据成功';
            echo json_encode(['code' => $code, 'data' => $data]);
        } else {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '获取编辑数据失败';
            echo json_encode(['code' => $code]);
        }
    }

    /**
     *  修改 提交
     */
    public function actionOpen_update() {
        $post = Yii::$app->request->post();
        $error = false;
        if (empty($post['id']) || !isset($post['id'])) {
            $error = true;
            $msg = 'ID不能为空!';
        } else if (empty($post['fc_type']) || !isset($post['fc_type'])) {
            $error = true;
            $msg = '彩种类型不能为空!';
        } else if (empty($post['qishu']) || !isset($post['qishu'])) {
            $error = true;
            $msg = '期数不能为空!';
        } else if (empty($post['kaipan']) || !isset($post['kaipan'])) {
            $error = true;
            $msg = '开盘时间不能为空!';
        } else if (empty($post['fengpan']) || !isset($post['fengpan'])) {
            $error = true;
            $msg = '封盘时间不能为空!';
        } else if (empty($post['kaijiang']) || !isset($post['kaijiang'])) {
            $error = true;
            $msg = '开奖时间不能为空!';
        }

        if ($error) {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = $msg;
            echo json_encode(['code' => $code]);
            die;
        }
        $tabname = Helper::GetOpenTimeTableNameByType($post['fc_type']); //获取表名
        $where = ['id' => $post['id']];
        //开盘时间为年月日时分秒的彩种
        $open_ymd = array('jnd_28', 'jnd_bs', 'liuhecai');
        $tabname = Helper::GetOpenTimeTableNameByType($post['fc_type']); //获取表名
        if (in_array($post['fc_type'], $open_ymd)) {
            $data = [
                'qishu' => $post['qishu'],
                'kaipan' => $post['kaipan'],
                'fengpan' => $post['fengpan'],
                'kaijiang' => $post['kaijiang'],
                'status' => $post['status']
            ];
        } else {
            $data = [
                'qishu' => $post['qishu'],
                'kaipan' => date('H:i:s', strtotime($post['kaipan'])),
                'fengpan' => date('H:i:s', strtotime($post['fengpan'])),
                'kaijiang' => date('H:i:s', strtotime($post['kaijiang'])),
                'status' => $post['status']
            ];
        }
        $result = otime::upEditOpenData($data, $where, $tabname);
        if ($result) {
            $redis = Yii::$app->redis;
            $redis_key = 'fengpan_time_' . $post['fc_type'];
            $redis->del($redis_key);
            $code['ErrorCode'] = 1;
            $code['ErrorMsg'] = '更新成功';
            echo json_encode(['code' => $code]);
            die;
        } else {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '更新失败';
            echo json_encode(['code' => $code]);
            die;
        }
    }

    /**
     * 新增
     */
    public static function actionOpen_add() {
        $post = Yii::$app->request->post();
        $error = false;
        if (empty($post['fc_type']) || !isset($post['fc_type'])) {
            $error = true;
            $msg = '彩种类型不能为空!';
        } else if (empty($post['qishu']) || !isset($post['qishu'])) {
            $error = true;
            $msg = '期数不能为空!';
        } else if (empty($post['kaipan']) || !isset($post['kaipan'])) {
            $error = true;
            $msg = '开盘时间不能为空!';
        } else if (empty($post['fengpan']) || !isset($post['fengpan'])) {
            $error = true;
            $msg = '封盘时间不能为空!';
        } else if (empty($post['kaijiang']) || !isset($post['kaijiang'])) {
            $error = true;
            $msg = '开奖时间不能为空!';
        } else if (empty($post['status'])) {
            $post['status'] = 1;
        }
        if ($error) {
            $ErrorCode = 2;
            $ErrorMsg = $msg;
            echo json_encode(['ErrorCode' => $ErrorCode, 'ErrorMsg' => $ErrorMsg]);
            die;
        }
        //开盘时间为年月日时分秒的彩种
        $open_ymd = array('jnd_28', 'jnd_bs', 'liuhecai');
        $tabname = Helper::GetOpenTimeTableNameByType($post['fc_type']); //获取表名
        if (in_array($post['fc_type'], $open_ymd)) {
            $data = [
                'qishu' => $post['qishu'],
                'fc_type' => $post['fc_type'],
                'kaipan' => $post['kaipan'],
                'fengpan' => $post['fengpan'],
                'kaijiang' => $post['kaijiang'],
                'status' => $post['status']
            ];
        } else {
            $data = [
                'qishu' => $post['qishu'],
                'fc_type' => $post['fc_type'],
                'kaipan' => date('H:i:s', strtotime($post['kaipan'])),
                'fengpan' => date('H:i:s', strtotime($post['fengpan'])),
                'kaijiang' => date('H:i:s', strtotime($post['kaijiang'])),
                'status' => $post['status']
            ];
        }
        // var_dump($post);
        // var_dump($data);exit;
        $result = otime::addOpenData($tabname, $data);
        if ($result) {
            $redis = Yii::$app->redis;
            $redis_key = 'fengpan_time_' . $post['fc_type'];
            $redis->del($redis_key);
            $ErrorCode = 1;
            $ErrorMsg = '保存成功';
            echo json_encode(['ErrorCode' => $ErrorCode, 'ErrorMsg' => $ErrorMsg]);
            die;
        } else {
            $ErrorCode = 1;
            $ErrorMsg = '保存失败';
            echo json_encode(['ErrorCode' => $ErrorCode, 'ErrorMsg' => $ErrorMsg]);
            die;
        }
    }

    public static function actionOpen_delete() {
        $post = Yii::$app->request->post();
        $error = false;
        if (empty($post['id']) || !isset($post['id'])) {
            $error = true;
            $msg = 'ID不能为空!';
        } else if (empty($post['fc_type']) || !isset($post['fc_type'])) {
            $error = true;
            $msg = '彩种类型不能为空!';
        }
        if ($error) {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = $msg;
            echo json_encode(['code' => $code]);
            die;
        }
        $tabname = Helper::GetOpenTimeTableNameByType($post['fc_type']); //获取表名
        $where = ['id' => $post['id']];
        $result = otime::delOpenData($where, $tabname);
        if ($result) {
            $code['ErrorCode'] = 1;
            $code['ErrorMsg'] = '删除成功';
            echo json_encode(['code' => $code]);
            die;
        } else {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '删除失败';
            echo json_encode(['code' => $code]);
            die;
        }
    }

}
