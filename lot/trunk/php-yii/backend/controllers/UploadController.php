<?php

namespace backend\controllers;

use Yii;
use backend\controllers\Controller;
use common\models\UploadModel as selfModel;
use common\models\AgentModel;
use yii\helpers\ArrayHelper;

class UploadController extends Controller {

    public function actionIndex() {
        $data['show'] = $this->showData();
        $data['trans'] = $this->transData();

        $request = Yii::$app->request->get();
        $request['line_id'] = Yii::$app->session->get('line_id'); // 所在线路
        $result = $this->search($request, $data['show']);
        if (isset($result['data'])) {
            foreach ($result['data'] as &$item) {
                $item['showpath'] = Yii::$app->params['upurl'] . $item['filepath'];
            }unset($item);
        }

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

        $item['addtime'] = date('Y-m-d H:i:s', $item['addtime']);

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

    public function actionEnable() {
        $result = $this->enable();

        echo json_encode($result);
        exit;
    }

    public function actionUp() {
        $result = $this->up();

        echo json_encode($result);
        exit;
    }

    /*
     * ******************************************************************************
     */

    public function search($request, $showData) {
        $query = $request;
        unset($query['_pjax']);
        unset($query['line_id']);
        if (empty($query)) {
            return;
        }

        $line_id = isset($request['line_id']) ? trim($request['line_id']) : '';
        $keyword = isset($request['keyword']) ? trim($request['keyword']) : '';
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $enable = isset($request['enable']) ? trim($request['enable']) : '';
        $page = isset($request['page']) ? intval($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? intval($request['pagenum']) : 100;

        // 查询条件
        $where[] = 'and';
        if ($keyword) {
            $where[] = [
                'OR',
                ['LIKE', 'filename', $keyword],
                ['LIKE', 'original_filename', $keyword]
            ];
        }
        if ($line_id) {
            $where[] = ['=', 'line_id', $line_id];
        }
        if ($enable != '') {
            $where[] = ['=', 'enable', $enable];
        }

        // 时间 处理
        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;
        // 时间筛选
        if ($starttime && $endtime) {
            $where[] = ['between', 'addtime', $starttime, $endtime];
        } elseif ($starttime) {
            $where[] = ['>', 'addtime', $starttime];
        } elseif ($endtime) {
            $where[] = ['<', 'addtime', $endtime];
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

    public function enable() {
        $request = Yii::$app->request->post();
        $id = isset($request['id']) ? intval($request['id']) : 0;
        $enable = isset($request['enable']) ? intval($request['enable']) : 0;

        if (empty($id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'ID缺失';
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

    public function up() {
        $line_id = Yii::$app->session['line_id'];
        $agent_id = Yii::$app->session['uid'];

        if (!$_FILES) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '请选择文件';
            return $result;
        }

        $signParams['line_id'] = $data['line_id'] = $line_id;
        $signParams['agent_id'] = $data['agent_id'] = $agent_id;
        $signParams['nonce_str'] = $data['nonce_str'] = time();
        $data['sign'] = $this->getSign($signParams);

        //缩略图宽高
        $data['width'] = isset($_REQUEST['width']) ? $_REQUEST['width'] : null;
        $data['height'] = isset($_REQUEST['height']) ? $_REQUEST['height'] : null;
        $data['is_thumb'] = isset($_REQUEST['is_thumb']) ? $_REQUEST['is_thumb'] : 2;

        foreach ($_FILES as $filename => $file) {
            // $data[$filename] = new \CURLFile($file['tmp_name'], $file['type'], $file['name']);
            $data[$filename] = curl_file_create($file['tmp_name'], $file['type'], $file['name']);
        }

        if (empty($host = Yii::$app->params['app_upload_host_http'])) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未配置 : app_upload_host_http';
            return $result;
        }
        $res = $this->https_request($host, true, $data);
        $res = json_decode($res, true);

        if (!$res) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '请求服务失败';
            return $result;
        }

        if (isset($res['errcode']) && $res['errcode'] != 0) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '上传出错' . ' : errcode ' . $res['errcode'] . ', errmsg ' . $res['errmsg'];
            return $result;
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '上传成功';
        $result['returnCode'] = $res;
        return $result;
    }

    // 校验值 生成规则
    public function getSign($params) {
        asort($params, 2);
        foreach ($params as $k => $v) {
            if (!empty($v)) {
                if (!isset($signTemp))
                    $signTemp = $k . '=' . $v;
                else
                    $signTemp .= '&' . $k . '=' . $v;
            }
        }
        $signTemp = $signTemp . '&key=' . $this->getKey();
        $sign = strtoupper(MD5($signTemp));
        return $sign;
    }

    // 校验密钥 可更改
    public function getKey() {
        return 'vantone';
    }

    // 获取网页内容函数
    public function https_request($url, $post = TRUE, $data = []) {
        $ch = curl_init($url);

        curl_setopt($ch, CURLOPT_POST, $post); // 是否POST方式
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data); // 数据
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE); // 为TRUE时 return内容，为FALSE时 直接在页面echo 且返回bool结果
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);

        $result = curl_exec($ch);

        curl_close($ch);
        return $result;
    }

    public function showData($type = '') {

        switch($type){
            case 'agents':
                $agents = AgentModel::getAgent('at');
                $data = ArrayHelper::index($agents, 'id');
                break;
            default:
                $agents = AgentModel::getAgent('at');

                $agents = ArrayHelper::index($agents, 'id');

                $data['agents'] = $agents;
                break;
        }

        return $data;
    }

    public function transData() {

        $result['enable'] = [0 => '已关闭', 1 => '已开启'];

        return $result;
    }

    public function trans($data, $showData) {
        // $showData = $this->showData();
        $transData = $this->transData();

        foreach ($data as $k => &$v) {
            $v['agentTxt'] = array_key_exists($v['agent_id'], $showData['agents']) ? $showData['agents'][$v['agent_id']]['login_name'] : $v['agent_id'];
            $v['enableTxt'] = array_key_exists($v['enable'], $transData['enable']) ? $transData['enable'][$v['enable']] : $v['enable'];
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
        }unset($v);

        return $data;
    }

}
