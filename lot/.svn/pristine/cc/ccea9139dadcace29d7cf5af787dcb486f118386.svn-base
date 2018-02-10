<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\LineModel as line;
use common\models\LineModel;
use common\models\AgentModel;
use common\helpers\Curl;

class LineController extends Controller {

    /**
     * **********************************************************
     *  线路管理列表             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $line_list = $this->getLines();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'page' => 1,
                'status' => '',
                'type' => '',
                'line' => $line_list
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $linename = isset($get['line_name']) ? $get['line_name'] : null;
        $page = isset($get['page']) ? intval($get['page']) : 1;
        $pagenum = isset($get['pagenum']) ? abs(intval($get['pagenum'])) : 100;
        $line_id = isset($get['line_id']) ? $get['line_id'] : null;
        $status = isset($get['status']) ? intval($get['status']) : null;
        $type = isset($get['type']) ? intval($get['type']) : null;
        //判断页码合法性
        if (!is_int($page) || $page <= 0) {
            $page = 1;
        }
        if (!is_int($pagenum) || $pagenum == 0) {
            $pagenum = 100;
        } else {
            $pagenum = abs($pagenum);
        }
        //拼接查询条件
        $where = array();
        if (!empty($line_id)) {
            $where['line_id'] = $line_id;
        }
        if (!empty($status) && (intval($status) != 0)) {
            $where['status'] = $status;
        }
        if (!empty($type) && (intval($type) != 0)) {
            $where['type'] = $type;
        }
        if (!empty($linename)) {
            $where = ['and', $where, ['like', 'line_name', $linename]];
        }
        $line_count = line::countLin($where);
        if ($line_count > 0) {
            $total_page = ceil($line_count / $pagenum); //总页码数
            if ($page > $total_page && $total_page != 0)
                $page = $total_page;
            $offset = ($page - 1) * $pagenum;  //开始条数
            $limit = $pagenum;       //每页显示条数

            $arr = line::queryLin($where, $offset, $limit);
        }else {
            $arr = array();
        }

        if ($arr) {
            //处理数组
            $data = array();
            foreach ($arr as $key => $val) {
                $data[$key] = $arr[$key];
                $data[$key]['addtime'] = date('Y-m-d H:i:s', $val['addtime']);
                $data[$key]['updatetime'] = date('Y-m-d H:i:s', $val['updatetime']);
            }
            $status = array(); //线路状态
            $type = array(); //交易模式
            $status = array('', '启用', '停用', '维护');
            $type = array('', '钱包模式', '额度转换');
            $experience = array('', '否', '是');
        } else {
            $data = array();
            $page = 1;
            $total_page = 1;
            $experience = array('', '否', '是');
        }

        $render = [
            'data' => $data,
            'pagecount' => $total_page,
            'page' => $page,
            'status' => $status,
            'type' => $type,
            'experience' => $experience,
            'line' => $line_list
        ];


        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  线路状态更改             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUpdatestatus() {
        $post = Yii::$app->request->post();
        //接收数据
        $id = isset($post['id']) ? intval($post['id']) : null;
        $status = isset($post['status']) ? intval($post['status']) : null;
        //验证数据完整性
        if (empty($id) || empty($status)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '找不到主键';
            echo json_encode($result);
            exit;
        }
        if (($id == 0) || ($status == 0)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '数据类型不正确';
            echo json_encode($result);
            exit;
        }

        $where = array();
        $where['id'] = $id;

        $res = line::updateone('status', $status, $where);
        if ($res) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '更新成功！';
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '更新失败！';
        }
        echo json_encode($result);
        exit;
    }

    /**
     * **********************************************************
     *  额度分配                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionUpdatemoney() {
        $session = Yii::$app->session;
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : null;
        $pattern = isset($post['pattern']) ? abs($post['pattern']) : null;
        $new_money = $cash_num = isset($post['money']) ? abs($post['money']) : null;
        $line_id = isset($post['line_id']) ? $post['line_id'] : null;
        $res = array();
        if (empty($line_id) || empty($id) || empty($pattern) || empty($new_money) || (!is_numeric($id)) || (!is_numeric($pattern)) || (!is_numeric($new_money))) {
            $res['ErrorCode'] = 2;
            $res['ErrorMsg'] = '参数不正确！';
            return json_encode($res);
        }

        //查询旧金额
        $old_money = line::oldmoney($id);
        if ((!$old_money) && ($old_money != 0) && ($old_money != '0.00')) {
            $res['ErrorCode'] = 2;
            $res['ErrorMsg'] = '获取信息失败';
            return json_encode($res);
        }
        //确定新金额
        if ($pattern == 1) { //存入
            $str = '存入';
            $money_field = 'money+' . $new_money;
            $new_money += $old_money;
        } elseif ($pattern == 2) {//取出
            $str = '取出';
            $money_field = 'money-' . $new_money;
            $new_money = $old_money - $new_money;
            if ($new_money < 0) {
                $res['ErrorCode'] = 2;
                $res['ErrorMsg'] = '金额不足！';
                return json_encode($res);
            }
        } else {
            $res['ErrorCode'] = 2;
            $res['ErrorMsg'] = '参数不正确！';
            return json_encode($res);
        }
        $where = array();
        $where['line_id'] = $line_id;
        $where['id'] = $id;
        $money_field = new \yii\db\Expression($money_field);
        //更新金额
        $result = line::updateone('money', $money_field, $where);
        if ($result) {
            $uid = $session['uid'];
            $remark = '站点' . $str . '额度' . $cash_num . '元，操作人:' . $session['login_user'];
            $time = time();
            //写入现金记录
            $data = array();
            $data['line_id'] = $line_id;
            $data['cash_num'] = $cash_num;
            $data['hander_id'] = $uid;
            $data['cash_balance'] = $new_money;
            $data['remark'] = $remark;
            $data['cash_type'] = $pattern;
            $data['addtime'] = time();
            $data['addday'] = date('Ymd', time());

            $result = line::add('line_cash_record', $data);

            if ($result) {
                //写入mongo日志
                $remark = '站点' . $line_id . $str . '额度' . $cash_num . '元,' . '原有额度' . $old_money . '元,' . '现有额度' . $new_money . '元,操作人:' . $session['login_user'];
                parent::insertOperateLog('', '', $remark);

                $res['ErrorCode'] = 1;
                $res['ErrorMsg'] = '额度分配成功';
                return json_encode($res);
            } else {
                $res['ErrorCode'] = 2;
                $res['ErrorMsg'] = '额度分配失败';
                return json_encode($res);
            }
        }
        $res['ErrorCode'] = 2;
        $res['ErrorMsg'] = '额度分配失败!';
        return json_encode($res);
    }

    /**
     * **********************************************************
     *  线路基本信息修改页面(表单) @author ruizuo qisuichen    *
     * **********************************************************
     */
    public function actionLineform() {
        $get = Yii::$app->request->get();
        //接收数据
        $id = isset($get['id']) ? intval($get['id']) : null;
        $line_id = isset($get['line_id']) ? $get['line_id'] : null;
        //验证数据
        $result = array();
        if (empty($id) || empty($line_id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '找不到主键';
            echo json_encode($result);
            exit;
        }
        if ($id == 0) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '找不到主键';
            echo json_encode($result);
            exit;
        }
        //拼接查询条件
        $where = array();
        $where['id'] = $id;
        $where['line_id'] = $line_id;


        $data = line::lineForm($where);
        if ($data) {
            $data['oldaddtime'] = $data['addtime'];
            $data['addtime'] = date('Y-m-d H:i:s', $data['addtime']);
            $data['updatetime'] = date('Y-m-d H:i:s', $data['updatetime']);
            $render = [
                'data' => $data
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('lineform.html', $render);
            } else {
                return $this->render('lineform.html', $render);
            }
        } else {
            $this->redirect('/line/index');
        }
    }

    /**
     * **********************************************************
     * 线路基本信息 修改( 提交表单) @author ruizuo qiyongsheng  *
     * **********************************************************
     */
    public function actionBaseupdate() {
        $post = Yii::$app->request->post();
        //接收数据
        $id = isset($post['id']) ? intval($post['id']) : null;
        $oldaddtime = isset($post['oldaddtime']) ? $post['oldaddtime'] : null;
        $linename = isset($post['line_name']) ? $post['line_name'] : null;
        $line_id = isset($post['line_id']) ? $post['line_id'] : null;
        $status = isset($post['status']) ? intval($post['status']) : null;
        $type = isset($post['type']) ? intval($post['type']) : null;
        $url = isset($post['url']) ? $post['url'] : null;
        $api_key = isset($post['api_key']) ? $post['api_key'] : null;
        $deskey = isset($post['deskey']) ? $post['deskey'] : null;
        $experience = isset($post['experience']) ? $post['experience'] : null;
        // $money = isset($post['money']) ? $post['money'] : null;
        $result = array(); //返回数据储存数组
        //验证数据完整性
        if (empty($id) || $id == 0 || empty($oldaddtime)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '找不到主键';
            echo json_encode($result);
            exit;
        }

        if (empty($linename) ||
                empty($line_id) ||
                empty($status) ||
                empty($type) ||
                empty($url) ||
                empty($api_key) ||
                empty($deskey)
        ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '数据不完整';
            echo json_encode($result);
            exit;
        }

        if (mb_strlen($api_key) != 17) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'md5Key长度只能为17位';
            echo json_encode($result);
            exit;
        }

        if (mb_strlen($deskey) != 8) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'desKey长度只能为8位';
            echo json_encode($result);
            exit;
        }

        //处理数组
        $data = array();

        $data['line_name'] = $linename;
        // $data['line_id'] = $line_id;
        $data['status'] = $status;
        $data['type'] = $type;
        $data['url'] = $url;
        $data['md5key'] = $api_key;
        $data['deskey'] = $deskey;
        $data['is_shiwan'] = $experience;
        $data['updatetime'] = time();
        //条件拼接
        $where = array();
        $where['id'] = $id;
        $where['addtime'] = $oldaddtime;
        $old = $mongo = array();
        $old = line::lineForm(array('id' => $id));
        $res = line::baseUpdate($where, $data);
        if ($res) {
            //写入mongo
            parent::insertOperateLog(json_encode($old), json_encode($data), '修改线路' . $line_id);
            $r = $this->refreshGoApi();
            $result['ErrorCode'] = 1;
            $result['r'] = $r;
            $result['ErrorMsg'] = '更新成功！';
            //清除缓存
            Yii::$app->redis->del('line_list');
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '更新失败！';
        }
        echo json_encode($result);
        exit;
    }

    //刷新golang接口线路缓存
    public function refreshGoApi() {
        $api_url = Yii::$app->params['golangApi'] . '/refresh/key';
        $api_res = Curl::run($api_url);

        $spider_url = Yii::$app->params['golangSpider'] . '/refresh/key';
        $spider_res = Curl::run($spider_url);

        return ['api' => $api_res, 'spider' => $spider_res];
    }

    /**
     * **********************************************************
     *  线路基本信息 新增页面    @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionInsert() {
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('add.html');
        } else {
            return $this->render('add.html');
        }
    }

    /**
     * **********************************************************
     *  线路基本信息 新增        @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionAdd() {
        $post = Yii::$app->request->post();
        // 接收数据
        $linename = isset($post['line_name']) ? $post['line_name'] : null;
        // $money = isset($post['money']) ? intval($post['money']) : 0;
        $line_id = isset($post['line_id']) ? $post['line_id'] : null;
        $status = isset($post['status']) ? intval($post['status']) : null;
        $type = isset($post['type']) ? intval($post['type']) : null;
        $url = isset($post['url']) ? $post['url'] : null;
        $api_key = isset($post['api_key']) ? $post['api_key'] : null;
        $deskey = isset($post['deskey']) ? $post['deskey'] : null;
        $experience = isset($post['experience']) ? $post['experience'] : null;
        //验证数据完整性
        if (empty($linename) ||
                empty($line_id) ||
                empty($status) ||
                empty($type) ||
                empty($url) ||
                empty($api_key) ||
                empty($deskey)
        ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '数据不完整';
            echo json_encode($result);
            exit;
        }
        if (!preg_match("/^[a-zA-z\#]*$/", $line_id)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '线路id只能是纯字母';
            echo json_encode($result);
            exit;
        }

        if (mb_strlen($api_key) != 17) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'md5Key长度只能为17位';
            echo json_encode($result);
            exit;
        }

        if (mb_strlen($deskey) != 8) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'desKey长度只能为8位';
            echo json_encode($result);
            exit;
        }
        //查看数据是否存在
        $is_exist = line::countLin(['line_id' => $line_id]);
        if ($is_exist) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '该线路已经存在,请勿重复添加';
            echo json_encode($result);
            exit;
        }
        $time = time();
        //处理数组
        $data = array();
        // $data['id'] = 'next value for MYCATSEQ_GLOBAL';
        $data['line_name'] = $linename;
        $data['money'] = 0;
        $data['line_id'] = $line_id;
        $data['status'] = $status;
        $data['type'] = $type;
        $data['url'] = $url;
        $data['md5key'] = $api_key;
        $data['deskey'] = $deskey;
        $data['is_shiwan'] = $experience;
        $data['addtime'] = $data['updatetime'] = time();

        $res = line::add('sys_line_list', $data);
        if ($res) {
            //写入mongo日志
            parent::insertOperateLog('', json_encode($data), '增加线路' . $line_id);
            //清除缓存
            Yii::$app->redis->del('line_list');
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '新增成功';
            echo json_encode($result);
            exit;
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '新增失败';
            echo json_encode($result);
            exit;
        }
    }

}
