<?php

namespace our_backend\controllers\gameapi;

use Yii;
use our_backend\controllers\Controller;
use common\models\AutoModel as selfModel;
use common\models\LogModel;
use common\helpers\Curl;
use common\helpers\lotteryOrm;

class AutoController extends Controller {

    public function actionIndex() {
        $type_arr = $this->getAllFcTypes();
        $data = array();
        foreach ($type_arr as $key => $value) {
            $data[$key]['fc_type'] = $value['type'];
            $data[$key]['fc_name'] = $value['name'];
        }
        $data['games'] = $data;

        $data['trans'] = $this->transData();
        $result = $this->search();

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

    public function actionInsert() {
        $result = $this->insert();

        echo json_encode($result);
        exit;
    }

    public function actionUpdate() {
        $result = $this->update();

        echo json_encode($result);
        exit;
    }

    /*
     * ******************************************************************************
     */

    public function search() {
        $query = $request = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            return;
        }

        $fc_type = !empty($request['fc_type']) ? trim($request['fc_type']) : '';
        $qishu = isset($request['qishu']) ? floatval($request['qishu']) : 0;
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $status = isset($request['status']) ? trim($request['status']) : 0;
        $page = isset($request['page']) ? intval($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? intval($request['pagenum']) : 10;

        if (!$fc_type) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '请选择彩种';
            return $result;
        }

        $where = ['and'];
        // $where[] = ['status'=>1];
        if ($qishu) {
            $where[] = ['like', 'qishu', $qishu];
        }
        // 时间 处理
        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;
        // 时间筛选
        if ($starttime && $endtime) {
            $where[] = ['between', 'addtime', $starttime, $endtime];
        } elseif (!empty($starttime)) {
            $where[] = ['>', 'addtime', $starttime];
        } elseif (!empty($endtime)) {
            $where[] = ['<', 'addtime', $endtime];
        }
        if ($status) {
            $where[] = ['=', 'status', $status];
        }

        // 分页
        $total_count = selfModel::getCount($fc_type, $where);
        $total_page = ceil($total_count / $pagenum);
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        if ($page <= 0)
            $page = 1;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $rows = selfModel::getList($fc_type, $where, $offset, $limit);

        $rows = $this->trans($rows);

        $lotteryOrm = new lotteryOrm();
        $ball_num = $lotteryOrm->getBallNum($fc_type);
        //取出开奖结果
        $tmp_arr = array();
        for ($i = 1; $i <= $ball_num; $i++) {
            $tmp_arr[] = 'ball_' . $i;
        }
        foreach ($rows as $key => $val) {
            foreach ($val as $k => &$v) {
                $ball = array();
                if (in_array($k, $tmp_arr)) {
                    $ball = explode(',', $v);
                    $rows[$key][$k] = $ball[0];
                }
            }
        }

        $oauto = LogModel::getLogs('auto', ['fc_type' => $fc_type], 0, 999, []);
        $diff = [];
        foreach ($rows as &$item) {
            foreach ($oauto as $key => $auto) {
                if ($item['qishu'] == $auto['expect']) {
                    $opencode = explode(',', $auto['opencode']);
                    $item['oauto'] = $opencode;
                    for ($i = 0; $i < $ball_num; $i++) {
                        if ($item['ball_' . ($i + 1)] != $opencode[$i]) {
                            $item['diff'][] = $i;
                        }
                    }
                }
            }
        }unset($item);

        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $key => $val) {
            $tmp[$val['type']] = $val;
        }
        $games = $tmp;
        $result['ErrorCode'] = 1;
        $result['fc_type'] = $games[$fc_type]['name'];
        $result['ball_num'] = $ball_num;
        $result['data'] = $rows;
        $result['diff'] = $diff;
        $result['page'] = $page;
        $result['pagenum'] = $pagenum;
        $result['totalpage'] = $total_page;
        return $result;
    }

    public function insert() {
        $redis = Yii::$app->redis;
        $request = Yii::$app->request->get();
        $fc_type = isset($request['fc_type']) ? trim($request['fc_type']) : '';
        $qishu = isset($request['qishu']) && is_numeric($request['qishu']) ? trim($request['qishu']) : '';
        $datetime = isset($request['datetime']) ? trim($request['datetime']) : '';
        $ball_num = isset($request['ball_num']) ? intval($request['ball_num']) : 0;

        $result['ErrorCode'] = 2;
        if ( empty($fc_type) ) {
            $result['ErrorMsg'] = '请选择彩种';
            return $result;
        }
        if ( empty($qishu) ) {
            $result['ErrorMsg'] = '请输入期数';
            return $result;
        }
        if ( empty($datetime) ) {
            $result['ErrorMsg'] = '请选择开奖时间';
            return $result;
        }
        for ($i = 1; $i <= $ball_num; $i++) {
            $ball = isset($request['ball_' . $i]) ? trim($request['ball_' . $i]) : false;
            if ($ball === false) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '请将开奖结果填写完整';
                return $result;
            }
            if ($ball === '') {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '开奖结果不能为空';
                return $result;
            }
            $balls['ball_' . $i] = $ball;
        }

        // 检测重复
        $res = selfModel::getCount($fc_type, ['qishu' => $qishu]);
        if ($res > 0) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '期数已存在';
            return $result;
        }

        $datetime = strtotime($datetime) ? strtotime($datetime) : $datetime;

        $ball_arr = array();
        foreach ($balls as $key => $ball) {
            $ball_arr[] = $ball;
        }
        //验证开奖结果合法性
        $check = $this->checkBall( $fc_type, $ball_arr, $ball_num);
        if($check['ErrorCode'] == 2){
            return $check;
        }
        //获取算法结果
        if (empty($host = Yii::$app->params['app_spider_host_http'])) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未配置 : app_spider_host_http';
            return $result;
        }
        $res = Curl::run($host, 'post', array(
                    'todo' => 'get_res', // 获取算法结果数组
                    'fc_type' => $fc_type, //彩种名称
                    'balls' => $ball_arr //开奖结果
        ));
        // var_dump($res);return;
        if (!$res) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '获取开奖结果算法失败，请确认进程开启';
            return $result;
        }
        $values = json_decode($res, true);
        $values['qishu'] = $qishu;
        $values['datetime'] = $datetime;
        $values['addtime'] = time();
        $values['status'] = 1;
        $last_qishu = selfModel::getLastQishu($fc_type);
        $sqlQuery = selfModel::insert($fc_type, $values);

        // if (!$sqlQuery) {
        //     $result['ErrorCode'] = 2;
        //     $result['ErrorMsg'] = '添加失败';
        //     return $result;
        // }
        $new = $values;
        $new['fc_type'] = $fc_type;
        $username = Yii::$app->session->get('uname');

        //触发结算//推送开奖结果
        if ($sqlQuery) {
            $remark = $username . " 添加了开奖结果 ";
            $this->insertOperateLog('', json_encode($new), $remark); //日志
            //更改redis中最后一期数据
            $auto_key = 'all_list_auto';
            if ($last_qishu && ($qishu > $last_qishu)) {
                $values['type'] = $fc_type;
                $redis->hset($auto_key, $fc_type, json_encode($values));
                //推送
                if(in_array($fc_type, ['bj_10', 'jsfc'])){
                    $bj_color_arr = ['','bj-yellow','bj-blue','bj-black','bj-orange','bj-azure','bj-deepblue','bj-silver','bj-red','bj-brown','bj-green'];
                    $tmp_auto = array();
                    foreach($ball_arr as $key=>$val){
                        $tmp_auto[$key]['ball'] = $val;
                        $tmp_auto[$key]['color'] = isset($bj_color_arr[intval($val)]) ? $bj_color_arr[intval($val)] : $val;
                    }
                    $ball_arr = $tmp_auto;
                }
                $return = array();
                if(in_array($fc_type, ['liuhecai', 'jsliuhecai'])){
                   $return = parent::getLastAutoByType('liuhecai');
                   $return['cmd'] = 'lottery';
                }else{
                    $return['cmd'] = 'lottery';
                    $return['fc_type'] = $fc_type;
                    $return['qishu'] = $qishu;
                    $return['datetime'] = $datetime;
                    $return['ball'] = $ball_arr;
                }
                if (empty($host = Yii::$app->params['app_push_host_http'])) {
                    $result['ErrorCode'] = 2;
                    $result['ErrorMsg'] = '未配置 : app_push_host_http';
                    return $result;
                }
                $send_ip = $host;
                $res = Curl::run($send_ip, 'get', $return);
            }

            //结算
            $res = self::betApi('allbalance', $fc_type, $qishu);
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '添加开奖结果成功，请稍后查看该期注单是否结算。<br/>如未结算，请检查进程后，到注单列表未结算注单再次尝试结算';
            return $result;
        }
    }

    public function update() {
        $request = Yii::$app->request->get();
        $redis = Yii::$app->redis;
        $fc_type = isset($request['fc_type']) ? trim($request['fc_type']) : '';
        $id = isset($request['id']) ? intval($request['id']) : 0;
        $qishu = isset($request['edit_qishu']) && is_numeric($request['edit_qishu']) ? trim($request['edit_qishu']) : '';
        $datetime = isset($request['datetime']) ? trim($request['datetime']) : '';
        $ball_num = isset($request['ball_num']) ? intval($request['ball_num']) : 0;

        $info = isset($request['info']) ? trim($request['info']) : '';
        $result['ErrorCode'] = 2;
        if ( empty($fc_type) ) {
            $result['ErrorMsg'] = '请选择彩种';
            return $result;
        }
        if ( empty($id) ) {
            $result['ErrorMsg'] = 'id参数丢失';
            return $result;
        }
        if ( empty($datetime) ) {
            $result['ErrorMsg'] = '请选择开奖时间';
            return $result;
        }
        for ($i = 1; $i <= $ball_num; $i++) {
            $ball = isset($request['ball_' . $i]) ? trim($request['ball_' . $i]) : false;
            if ($ball === false) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '请将开奖结果填写完整';
                return $result;
            }
            if ($ball === '') {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '开奖结果不能为空';
                return $result;
            }
            $balls[] = $ball;
        }
        //获取算法结果
        if (empty($host = Yii::$app->params['app_spider_host_http'])) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未配置 : app_spider_host_http';
            return $result;
        }
        $res = Curl::run($host, 'post', array(
                    'todo' => 'get_res', // 获取算法结果数组
                    'fc_type' => $fc_type, //彩种名称
                    'balls' => $balls //开奖结果
        ));
        // var_dump($res);return;
        if (!$res) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '获取开奖结果算法失败，请确认进程开启';
            return $result;
        }
        
        $set = json_decode($res, true);
        // // 检测重复
        // $res = selfModel::getCount($fc_type, ['and', ['<>', 'id', $id], ['=', 'qishu', $qishu]]);
        // if ($res > 0) {
        //     $result['ErrorCode'] = 2;
        //     $result['ErrorMsg'] = '期数已存在';
        //     return $result;
        // }

        $where['id'] = $id;
        $datetime = strtotime($datetime) ? strtotime($datetime) : $datetime;
        $set['datetime'] = $datetime;

        $last_qishu = self::getLastQishu($fc_type);
        $sqlQuery = selfModel::update($fc_type, $set, $where);

        if (!$sqlQuery) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未修改';
            return $result;
        }

        $set['fc_type'] = $fc_type;

        $new = $set;
        parent::insertOperateLog($info, json_encode($new),'修改开奖结果');
        //更改redis中最后一期数据
        $auto_key = 'all_list_auto';
        $set['qishu'] = $qishu;
        if ($last_qishu && $qishu >= $last_qishu) {
            $redis->hset($auto_key, $fc_type, json_encode($set));
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '修改成功';
        $set['datedate'] = date('Y-m-d H:i:s', $set['datetime']);
        $set['datedatetime'] = date('Y-m-d H:i:s', $set['datetime']);
        $set['datetime'] = date('Y-m-d H:i:s', $set['datetime']);
        unset($set['fc_type']);
        $result['set'] = $set;
        return $result;
    }

    public function transData() {

        $result['status'] = [1 => '未结算', 2 => '已结算', 3 => '正在结算'];

        return $result;
    }

    public function trans($data) {
        foreach ($data as $k => &$v) {
            $v['datedate'] = date('Y-m-d H:i:s', $v['datetime']);
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            $v['datedatetime'] = date('Y-m-d H:i:s', $v['datetime']);
        }unset($v);
        return $data;
    }

    /**
     * **********************************************************
     *  检测某期已结算注单条数   @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionCheckbet() {
        $post = Yii::$app->request->post();
        $fc_type = isset($post['fc_type']) ? trim($post['fc_type']) : '';
        $qishu = isset($post['qishu']) ? floatval($post['qishu']) : 0;
        if ( empty($fc_type) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '请选择彩种';
            return json_encode($result);
        }
        if (empty($qishu) || (!is_numeric($qishu))) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '期数不正确';
            return json_encode($result);
        }
        $betTable = selfModel::getBetTableName($fc_type);
        $bet_count = selfModel::getOKBetCount($betTable, $fc_type, $qishu);

        if ($bet_count) {
            $result['ErrorCode'] = 3;
            $result['ErrorMsg'] = $bet_count; //返回已经结算注单的条数
        } else {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '该期没有结算过的注单';
        }
        return json_encode($result);
    }

    /**
     * **********************************************************
     *     根据时间和日期采集开奖结果  @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionGetauto() {
        $post = Yii::$app->request->post();
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : null;
        $time = isset($post['time']) ? strtotime($post['time']) : null;

        $result = array();
        $result['ErrorCode'] = 2;
        if ( empty($fc_type) ) {
            $result['ErrorMsg'] = '请选择彩种';
            return json_encode($result);
        }
        if ( empty($time) ) {
            $result['ErrorMsg'] = '请选择补采时间';
            return json_encode($result);
        }

        //如果是下列类型彩种无接口不能复采
        $not_arr = array(
                'fc_3d', //福彩3d
                'pl_3', //排列3
                'dm_28', //丹麦28
                'dm_klc', //丹麦快乐彩
                'liuhecai',//六合彩
                'jsliuhecai',//极速六合彩
                'jsfc',//极速飞车
                'ffc_o',//分分彩
                'lfc_o',//两分彩
                'dj_o',//东京1.5分
                'mg_o' //美国45秒
            );
        if (in_array($fc_type, $not_arr)) {
            $result['ErrorMsg'] = '该彩种不支持复采开奖结果';
            return json_encode($result);
        }
        //访问commom.php进程
        if (empty($host = Yii::$app->params['app_spider_host_http'])) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未配置 : app_spider_host_http';
            return $result;
        }
        $res = Curl::run($host, 'post', array(
                    'todo' => 'getauto',
                    'fc_type' => $fc_type, //彩种名称
                    'time' => $time//补彩的时间
        ));
        if ($res) {
            $res = json_decode($res, true);
            if(isset($res['ErrorCode']) && $res['ErrorCode'] == 2){
                return json_encode($res);
            }
            //插入数据库
            $balls_arr = $res['balls'];
            $qishu_list = $res['qishu_list'];
            $all_count = count($balls_arr);
            $tmp_count = 0;
            foreach ($balls_arr as $val) {
                $tmp_res = selfModel::insert($fc_type, $val);
                if ($tmp_res)
                    $tmp_count++;
            }

            if ($tmp_count) {
                foreach ($qishu_list as $qishu) {
                    self::betApi('allbalance', $fc_type, $qishu); //解发结算
                }
                $result['ErrorCode'] = 1;
                if ($tmp_count != $all_count) {
                    $result['ErrorMsg'] = '采集成功，其中' . ($all_count - $tmp_count) . '条数据未插入成功';
                } else {
                    $result['ErrorMsg'] = '插入成功，默认触发结算';
                }
                return json_encode($result);
            } else {
                $result['ErrorMsg'] = '插入数据失败';
                $result['ErrorCode'] = 2;
                return json_encode($result);
            }
            return $res; //json格式
        } else {
            $result['ErrorMsg'] = '无法连接到相关进程';
            $result['ErrorCode'] = 2;
            return json_encode($result);
        }
    }



    /**
     * **********************************************************
     *  获取最后一期期数           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getLastQishu($fc_type) {
        $qishu = selfModel::getLastQishu($fc_type);
        return $qishu;
    }


   /**
     * **********************************************************
     *  与other结算模块对接      @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function betApi($todo, $type, $periods, $id = array()) {
        if (empty($host = Yii::$app->params['app_lottery_host_http'])) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '未配置 : app_lottery_host_http';
            return $result;
        }
        $res = Curl::run($host, 'post', array(
                    'todo' => $todo, //单条结算balance批量结算allbalance 批量回滚rollback
                    'fc_type' => $type, //彩种名称
                    'id' => $id, //注单id
                    'periods' => $periods//期数
        ));
        $result = array(); //储存运行结果

        if ($res) {
            return $res; //json格式
        } else {
            $result['ErrorMsg'] = '无法连接到相关进程';
            $result['ErrorCode'] = 2;
            return json_encode($result);
        }
    }

/**
      ***********************************************************
      *  验证开奖结果合法性           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function checkBall($fc_type, $ball_arr, $num){
        $result = array();
        //检测开奖结果是否重复    
        $result['ErrorCode'] = 2;
        $tmp_ball_arr = array_unique($ball_arr);
        $tmp_arr = ['pc_28', 'jnd_28', 'dm_28', 'bj_28', 'ffc_o', 'lfc_o', 'els_o', 'dj_o', 'mg_o', 'mnl_o'];
        if(count($tmp_ball_arr) != $num && !in_array($fc_type, $tmp_arr)){
            $result['ErrorMsg'] = '开奖结果不能有重复的值';
            return $result;
        }

        //验证开奖球最大值和最小值
        $max = $min = 0;
        switch ($fc_type) {
            case 'liuhecai':
            case 'jsliuhecai':
                $max = 49;
                $min = 1;
                break;
            case 'bj_10':
            case 'jsfc':
            case 'xdl_10':
                $max = 10;
                $min = 1;
                break;
            case 'dm_klc':
            case 'bj_kl8':
            case 'jnd_bs':
                $max = 80;
                $min = 1;
            case 'gd_ten':
            case 'cq_ten':
                $max = 20;
                $min = 1;
                break;
            case 'gx_k3':
            case 'ah_k3':
            case 'js_k3':
                $max = 6;
                $min = 1;
                break;
            case 'gd_11':
            case 'jx_11':
            case 'sd_11':
                $max = 11;
                $min = 1;
                break;
            default:
                $min = 0;
                $max = 9;
                break;
        }

        foreach($ball_arr as $key=>$val){
            if(!is_numeric($val)){
                unset($ball_arr[$key]);
            }

            if($val > $max || $val < $min){
                $result['ErrorMsg'] = '开奖结果必须在' . $min . '~' . $max . '之间,您输入的是：' . $val;
                return $result;
            }
        }

        if(count($ball_arr) != $num){
            $result['ErrorMsg'] = '开奖结果只能是数字';
            return $result;
        }

        $result['ErrorCode'] = 1;
        return $result;
    }  
    
    
}
