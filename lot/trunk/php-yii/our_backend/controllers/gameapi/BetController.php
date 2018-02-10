<?php

namespace our_backend\controllers\gameapi;

use Yii;
use our_backend\controllers\Controller;
use common\models\BetModel;
use common\helpers\Helper;
use yii\helpers\ArrayHelper;
use common\helpers\Curl;

class BetController extends Controller {

    /**
     * **********************************************************
     *  注单管理列表             @author Rom    *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $page = !empty($get['page']) ? abs(intval($get['page'])) : 1;
        $pagenum = !empty($get['pagenum']) ? abs(intval($get['pagenum'])) : 200;

        if (!is_int($page) || $page <= 0) {
            $page = 1;
        }
        if (!is_int($pagenum) || $pagenum == 0) {
            $pagenum = 200;
        } else {
            $pagenum = abs($pagenum);
        }
        $starttime = isset($get['starttime']) ? $get['starttime'] : null;
        $endtime = isset($get['endtime']) ? $get['endtime'] : null;
        if(empty($starttime)){
            $starttime = strtotime(date('Y-m-d'));
        }else{
            $starttime = strtotime($starttime);
        }
        if(empty($endtime)){
            $endtime = time();
        }else{
            $endtime = strtotime($endtime);
        }
        $addday = date('Ymd', $starttime);
        $endday = date('Ymd', $endtime);
        $between_day = $endday - $addday;
        if($between_day < 0 ){
            $addday = $endday;
        }elseif($between_day > 4){
            $addday = $endday - 4;
        }

        $line_id = isset($post['line_id']) ? $post['line_id'] : null; //站点id
        $agent_list = array();
        if ((!empty($line_id))){
            $res = $this->actionGetagent($line_id);
            $agent_list = json_decode($res, true);
            if($agent_list['ErrorCode'] == 1){
                $agent_list = $agent_list['ErrorMsg'];
            }
        }
        $where = $this->getWhere();
        $games = $this->getAllFcTypes();
        $lines = $this->getLines(); //获取线路
        $html = array();
        //无检索条件，直接跳出查询
        if(isset($get['_pjax'])) unset($get['_pjax']);

        if (empty($get)) {
            $data = array();
            $render = [
                'data' => $data,
                'games' => $games,
                'lines' => $lines,
                'totalbet' => $html,
                'starttime' =>date('Y/m/d',strtotime($addday)),
                'endtime' => date('Y/m/d',strtotime($endday)),
                'pagedata' => ['totalpage' => 1]
            ];

            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }

        $totalnum = BetModel::getDataTotalNum($where); //获取当前彩种的总条数
        $condition ['field'] = array('id', 'addtime', 'periods', 'bet_info', 'order_num', 'bet', 'odds', 'ptype', 'at_username' ,'line_id', 'status', 'uname', 'win', 'fc_type', 'js'); //查询字段
        $totalpage = ceil($totalnum / $pagenum); //总页码数
        if ($page > $totalpage && $totalnum != 0)
            $page = $totalpage;
        $condition['offset'] = ($page - 1) * $pagenum;  //开始条数
        $condition['limit'] = $pagenum;
        $data = betModel::index($condition, $where);
        $total_data = array(); //列表页统计数据
        if (empty($data)) {
            $data = array();
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '获取数据失败';
            $pagedata['page'] = 1;
            $pagedata['totalpage'] = 1;
        } else {
            $bet_count = $count_a = $count_b = $money = $page_money_a = $page_money_b = $page_ok_bet = $page_ok_money = $win = 0; //用于统计当前页数据
            $type_arr = $this->getAllFcTypes();
            $new_type_arr = array();
            foreach($type_arr as $key=>$val){
                $new_type_arr[$val['type']] = $val['name'];
            }
            foreach ($data as $k => $v) {

                $data[$k]['addtime'] = date('Y-m-d H:i:s', $v['addtime']);
                $data[$k]['oldstatus'] = $v['status'];
                $data[$k]['oldInfo'] = json_encode($v);

                switch ($data[$k]['status']) {
                    case 1:
                        $data[$k]['status'] = '未结算';
                        break;
                    case 2:
                        $data[$k]['status'] = '赢';
                        break;
                    case 6:
                        $data[$k]['status'] = '赢(1)';
                        break;
                    case 7:
                        $data[$k]['status'] = '赢(2)';
                        break;
                    case 3:
                        $data[$k]['status'] = '输';
                        break;
                    case 4:
                        $data[$k]['status'] = '和局';
                        break;
                    case 5:
                        $data[$k]['status'] = '无效';
                        break;
                }
                switch ($data[$k]['ptype']) {
                    case 1:
                        $data[$k]['ptype'] = 'PC端';
                        break;
                    case 2:
                        $data[$k]['ptype'] = 'WAP端';
                        break;
                    case 3:
                        $data[$k]['ptype'] = 'APP';
                        break;
                }
                if(isset($new_type_arr[$v['fc_type']])){
                    $data[$k]['type_name'] = $new_type_arr[$v['fc_type']];
                }else{
                     $data[$k]['type_name'] = $v['fc_type'];
                }
                $tmp_bet_info = $this->new_bet_info($v['fc_type'], $v['bet_info']);
                $data[$k]['bet_info'] = $tmp_bet_info['gameplayTxt'] . ' ' . $tmp_bet_info['input_nameTxt'];

                //统计当前页
                $money += $v['bet'];
                $bet_count += 1;
                if ($v['js'] == 1 && $v['status'] == 1) {//未结算注单
                    $count_a += 1; //总笔数
                    $page_money_a += $v['bet']; //总打码
                } else {//已结算注单
                    $count_b += 1; //当前页总笔数
                    $page_money_b += $v['bet']; //当前页总打码
                    if ($v['status'] != 4) {
                        $page_ok_bet += 1; //当前页有效笔数
                        $page_ok_money += $v['bet']; //当前页有效打码
                        $win += $v['win'];
                    }
                }
            }

            //拼接字符串
            $html['str_page_all'] = '总笔数:' . $bet_count . '笔 ' . ' 总打码:' . $money . '元';

            $html['str_page_a'] = '总笔数:' . $count_a . '笔 ' . ' 总打码:' . $page_money_a . '元';

            $html['str_page_b'] = '总笔数:' . $count_b . '笔 ' . ' 总打码:' . $page_money_b . '元' . ' 有效笔数: ' . $page_ok_bet . '笔 ' . ' 有效打码:' . $page_ok_money . '元' . ' 盈利:' . number_format($page_ok_money - $win, '2', '.', ',') . '元';
           

            $code['ErrorCode'] = 1;
            $pagedata['page'] = $page;
            $pagedata['totalpage'] = $totalpage;
            $pagedata['totalnum'] = $totalnum;
        }
        $render = [
            'data' => $data,
            'games' => $games,
            'lines' => $lines,
            'agents' => $agent_list,
            'starttime' =>date('Y/m/d',strtotime($addday)),
            'endtime' => date('Y/m/d',strtotime($endday)),
            'code' => $code,
            'pagedata' => $pagedata,
            'totalbet' => $html
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('_list.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }
/**
      ***********************************************************
      *  获取列表页条件           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function getWhere(){
        $post = Yii::$app->request->get();
        $did = isset($post['did']) ? trim($post['did']) : null; //  订单号
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : null; //彩种类型
        $js = isset($post['js']) ? abs(intval($post['js'])) : null;  //结算状态
        $compare = isset($post['compare']) ? abs(intval($post['compare'])) : 1; //比较期数
        $periods = isset($post['periods']) ? abs(intval($post['periods'])) : null; //彩票期数
        $ptype = isset($post['ptype']) ? abs(intval($post['ptype'])) : null; //注单来源
        $line_id = isset($post['line_id']) ? $post['line_id'] : null; //站点id
        $agent_id = isset($post['agent_id']) ? $post['agent_id'] : null; //代理id
        $uname = isset($post['uname']) ? trim($post['uname']) : null; //投注账号
        $starttime = isset($post['starttime']) ? $post['starttime'] : null;
        $endtime = isset($post['endtime']) ? $post['endtime'] : null;
        if(empty($starttime)){
            $starttime = strtotime(date('Y-m-d'));
        }else{
            $starttime = strtotime($starttime);
        }
        if(empty($endtime)){
            $endtime = time();
        }else{
            $endtime = strtotime($endtime);
        }
        $addday = date('Ymd', $starttime);
        $endday = date('Ymd', $endtime);
        $between_day = $endday - $addday;
        if($between_day < 0 ){
            $addday = $endday;
        }elseif($between_day > 4){
            $addday = $endday - 4;
        }
       
        switch ($compare) {
            case 1:
                $compare = '=';
                break;
            case 2:
                $compare = '<';
                break;
            case 3:
                $compare = '<=';
                break;
            case 4:
                $compare = '>';
                break;
            case 5:
                $compare = '>=';
                break;
        }
        //查询条件
        $agent_list = array();
        $where = array();
        if (!empty($fc_type))
            $where['fc_type'] = $fc_type;
        if ((!empty($did)) && ($did != 0))
            $where['order_num'] = $did; //订单号
        if ((!empty($js)) && ($js != 0))
            $where['js'] = $js; //结算状态
        if ((!empty($ptype)) && ($ptype != 0))
            $where['ptype'] = $ptype; //'1为pc端 2为wap端  3为app'
        if ((!empty($line_id))){
            $where['line_id'] = $line_id; //线路
        }
        if(!empty($agent_id)){
            $where['at_id'] = $agent_id;
        }

        if ((!empty($uname)))
            $where['uname'] = $uname; //账号
        if ((!empty($periods)) && ($periods != 0) && $compare != '')
            $where = ['and', $where, array($compare, 'periods', $periods)]; //彩票期数
            
        //时间条件 
        if ($endtime > time()) {
            $endtime = time();
        }

        $where = ['and', $where, 'addday>=' . $addday];
        $where = ['and', $where, 'addday<=' . $endday];

        return $where;
    }
    
/**
      ***********************************************************
      *  指定条件统计数据响应ajax请求  @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionTotal(){
            $where = $this->getWhere();
            $get = Yii::$app->request->get();
            if(empty($where) || empty($get)) return 0;
            $type = isset($get['type']) ? $get['type'] : 1;
            if(isset($where['js'])) unset($where['js']);
            $result = array();
            //总计，该彩种 总数据
            $str = '';
            if($type == 1){
                $data = betModel::totalcount('sum(bet) as bet,count(id) as count', $where); //所有
                if(empty($data)) return '0';
                if(empty($data['bet'])) $data['bet'] = 0;
                $str = ' 总笔数:' . $data['count'] . '笔 ' . ' 总打码:' . number_format($data['bet'], '2', '.', ','). '元';
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = $str;
                return json_encode($result);
            }elseif($type == 2){
                $where = ['and', $where, ['status'=> 1]];
                $where = ['and', $where, ['js'=> 1]];
                $data = betModel::totalcount('sum(bet) as bet,count(id) as count', $where); //未结算  
                if(empty($data)) return '0';
                if(empty($data['bet'])) $data['bet'] = 0;
                $str = ' 总笔数:' . $data['count'] . '笔 ' . ' 总打码:' . number_format($data['bet'], '2', '.', ',') . '元';
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = $str;
                return json_encode($result);
            }elseif($type == 3){
                $where = ['and', $where, ['js'=> 2]];
                $data = betModel::totalcount('sum(bet) as bet,count(id) as count', $where); //已经结算
                $where = ['and', $where, ['status'=>[2,3,5,6,7]]];
                // $where['status'] = [2,3,5,6,7];
                $data_a = betModel::totalcount('sum(bet) as bet,count(id) as count,sum(win) as win', $where); //已经结算有效打码 
                if(empty($data) || empty($data_a)) return '0';
                if(empty($data['bet'])) $data['bet'] = 0;
                if(empty($data_a['bet'])) $data_a['bet'] = 0;
                if(empty($data_a['win'])) $data_a['win'] = 0;
                
                $str = ' 总笔数:' . $data['count'] . '笔 ' . ' 总打码:' . number_format($data['bet'], '2', '.', ',') . ' 有效笔数:' . $data_a['count'] . '笔 ' . ' 有效打码:' . number_format($data_a['bet'], '2', '.', ',') . ' 盈利:' . ($data_a['bet'] - $data_a['win']. '元');
                $result['ErrorCode'] = 3;
                $result['ErrorMsg'] = $str;
                return json_encode($result);
            }else{
                $result['ErrorCode'] = 4;
                $result['ErrorMsg'] = '处理类型不明确';
                return json_encode($result);
            }
    }
    
    public function actionGetagent($line_id = ''){
        $post = Yii::$app->request->post();
        $line_id = isset($post['line_id']) ? $post['line_id'] : $line_id;
        $result = array();
        $result['ErrorCode'] = 2;
        if($line_id == ''){
            $result['ErrorMsg'] = '请选择线路';
            echo json_encode($result);die;
        }
        $redis = \Yii::$app->redis;
        $redis_key = 'agent_list';
        $agent_arr = $redis->get($redis_key);
        if($agent_arr){
            $agent_arr = json_decode($agent_arr, true);
        }else{
            $agent_arr = BetModel::getAllAgent();
            if(!$agent_arr){
                $result['ErrorMsg'] = '获取代理数据失败';
                echo json_encode($result);die;
            }
            $redis->setex($redis_key, 600, json_encode($agent_arr));
        }
        $arr = array();
        foreach($agent_arr as $key=>$val){
            if($val['line_id'] == $line_id){
                $arr[] = $val;
            }
        }
        if(empty($arr)){
            $result['ErrorMsg'] = '获取代理数据失败';
            echo json_encode($result);die;
        }

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = $arr;
        return json_encode($result);
    }

    public function actionDetail() {
        $get = Yii::$app->request->get();
        $error = false;
        $id = isset($get['id']) && !empty($get['id']) ? $get['id'] : 0;
        $fc_type = isset($get['fc_type']) && !empty($get['fc_type']) ? $get['fc_type'] : null;
        if (empty($id) || empty($fc_type)) {
            $error = true;
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = 'id和彩种类型不能为空';
        }
        if ($error) {
            echo json_encode(['code' => $code]);
            die;
        }
        $where['id'] = $id;
        $data = BetModel::getJoinOneSql($where); //注单基本信息
        $tmp_bet_info =  $this->new_bet_info($data['fc_type'], $data['bet_info']); 
        $data['bet_info'] = $tmp_bet_info['gameplayTxt'] . ' ' . $tmp_bet_info['input_nameTxt'];//翻译投注内容
        if (!empty($data)) {
            $userInfo = BetModel::getUserInfo($data['uid']);
            if(!empty($userInfo)){
                $userInfo['regtime'] = date('Y-m-d H:i:s', $userInfo['addtime']);
                $userInfo['user_status'] = $userInfo['status'];
                unset($userInfo['addtime']);
                unset($userInfo['status']);
                $data = array_merge($userInfo, $data);
            }
            $data['addtime'] = date('Y-m-d H:i:s', $data['addtime']);
            $code['ErrorCode'] = 1;
            $code['ErrorMsg'] = '获取编辑数据成功';
        } else {
            $code['ErrorCode'] = 2;
            $code['ErrorMsg'] = '获取编辑数据失败';
            $data = array();
        }
        $render = [
            'data' => $data,
            'code' => $code
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('detail.html', $render);
        } else {
            return $this->render('detail.html', $render);
        }
    }

    // public function actionEdit() {
    //     $get = Yii::$app->request->get();
    //     $error = false;
    //     $id = isset($get['id']) && !empty($get['id']) ? $get['id'] : 0;
    //     $fc_type = isset($get['fc_type']) && !empty($get['fc_type']) ? $get['fc_type'] : ' ';
    //     if (empty($get['id']) || empty($get['fc_type'])) {
    //         $error = true;
    //         $code['ErrorCode'] = 2;
    //         $code['ErrorMsg'] = '参数错误';
    //     }
    //     if ($error) {
    //         echo json_encode(['code' => $code]);
    //         die;
    //     }
    //     $where['id'] = $id;
    //     $tabname = Helper::GetBetTableNameByType($fc_type); //获取表名
    //     $data = BetModel::getOneSql($where, $tabname);
    //     if ($data) {
    //         $data['addtime'] = date('Y-m-d H:i:s', $data['addtime']);
    //         $code['ErrorCode'] = 1;
    //         $code['ErrorMsg'] = '获取编辑数据成功';
    //     } else {
    //         $code['ErrorCode'] = 2;
    //         $code['ErrorMsg'] = '获取编辑数据失败';
    //         $data = array();
    //     }
    //     $render = ['data' => $data];
    //     if (Yii::$app->request->isAjax) {
    //         return $this->renderAjax('edit.html', $render);
    //     } else {
    //         return $this->render('edit.html', $render);
    //     }
    // }

    // public function actionUpdate() {
    //     $post = Yii::$app->request->post();
    //     $error = false;
    //     if (empty($post['id']) || !isset($post['id'])) {
    //         $error = true;
    //         $msg = 'ID不能为空！';
    //     } else if (empty($post['fc_type']) || !isset($post['fc_type'])) {
    //         $error = true;
    //         $msg = '彩种类型不能为空！';
    //     } else if (empty($post['line_id']) || !isset($post['line_id'])) {
    //         $error = true;
    //         $msg = '线路ID不能为空！';
    //     } else if (empty($post['uname']) || !isset($post['uname'])) {
    //         $error = true;
    //         $msg = '用户账号不能为空！';
    //     } else if (empty($post['periods']) || !isset($post['periods'])) {
    //         $error = true;
    //         $msg = '彩票期数不能为空！';
    //     } else if (empty($post['did']) || !isset($post['did'])) {
    //         $error = true;
    //         $msg = '订单号不能为空！';
    //     } else if (empty($post['money']) || !isset($post['money'])) {
    //         $error = true;
    //         $msg = '下注金额不能为空！';
    //     } else if (empty($post['bet_info']) || !isset($post['bet_info'])) {
    //         $error = true;
    //         $msg = '下注内容不能为空！';
    //     } else if (empty($post['odds']) || !isset($post['odds'])) {
    //         $error = true;
    //         $msg = '赔率不能为空！';
    //     } else if (empty($post['win']) || !isset($post['win'])) {
    //         $error = true;
    //         $msg = '派彩金额不能为空！';
    //     } else if (empty($post['addtime']) || !isset($post['addtime'])) {
    //         $error = true;
    //         $msg = '下注时间不能为空！';
    //     } else if (empty($post['ptype']) || !isset($post['ptype'])) {
    //         $error = true;
    //         $msg = '彩种类型不能为空！';
    //     } else if (empty($post['status']) || !isset($post['status'])) {
    //         $error = true;
    //         $msg = '结算状态不能为空！';
    //     }

    //     if ($error) {
    //         $code['ErrorCode'] = 2;
    //         $code['ErrorMsg'] = $msg;
    //         echo json_encode(['code' => $code]);
    //         die;
    //     }

    //     $where = ['id' => $post['id']];
    //     $data = [
    //         'line_id' => $post['line_id'],
    //         'uname' => $post['uname'],
    //         'periods' => $post['periods'],
    //         'did' => $post['did'],
    //         'money' => $post['money'],
    //         'bet_info' => $post['bet_info'],
    //         'odds' => $post['odds'],
    //         'win' => $post['win'],
    //         'addtime' => strtotime($post['addtime']),
    //         'ptype' => $post['ptype'],
    //         'status' => $post['status']
    //     ];
    //     $result = BetModel::upEditOpenData($data, $where);
    //     if ($result) {
    //         $code['ErrorCode'] = 1;
    //         $code['ErrorMsg'] = '保存成功';
    //         echo json_encode(['code' => $code]);
    //         die;
    //     } else {
    //         $code['ErrorCode'] = 2;
    //         $code['ErrorMsg'] = '保存失败';
    //         echo json_encode(['code' => $code]);
    //         die;
    //     }
    // }

    public function actionBatchbalance() {
        $result = self::batchBalance();

        echo ($result);
        die;
    }

    /**
     * **********************************************************
     *  批量结算回滚             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    //处理前端数据格式 前端数据 I注单id,期数,彩种
    public static function batchData($data) {
        //处理数据
        $new_arr = array();
        $new_arr = explode('I', ltrim($data, 'I'));
        $new_data = array();
        $all_count = count($new_arr);
        foreach ($new_arr as $key => $val) {
            $tmp_data = array();
            $tmp_data = explode(',', $val);
            if ($tmp_data) {
                if (count($tmp_data) == 3) {
                    $new_data[$tmp_data[2]][$tmp_data[1]][] = $tmp_data[0];
                }
            }
        }
        $new_data['all_count'] = $all_count;
        return $new_data;
    }

    //批量结算
    public function actionBatch_balance() {
        $request = Yii::$app->request->post();
        $ids = ArrayHelper::getValue($request, 'ids', '');

        $result['ErrorCode'] = 1;
        if (empty($ids)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'id参数缺失';
            return json_encode($result);
        }
        $new_data = self::batchData($ids);

        //循环结算
        $count = 0; //计算结算成功的次数
        $periods_count = $new_data['all_count'];
        unset($new_data['all_count']);
        foreach ($new_data as $type => $val) {
            foreach ($val as $periods => $ids) {
                if (is_numeric($periods) && is_array($ids)) {
                    $res = self::betApi('balance', $type, $periods, $ids);
                    $res = json_decode($res, true);
                    if ($res['ErrorCode'] == 1 && isset($res['ok'])) {
                        $count += $res['ok'];
                    } elseif (is_array($res)) {
                        return json_encode($res);
                    }
                }
            }
        }

        $result = array();
        if ($count == $periods_count) { //全部结算完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '批量结算成功';
        } elseif ($count > 0) {//部分结算完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '结算成功:&nbsp;' . $count . '&nbsp;条' . '<br/>' . '结算失败:&nbsp;' . ($periods_count - $count) . '&nbsp;条<br/>可能失败原因:<br/>1.该注单已经结算过<br/>2.没有开奖结果<br/>3.非法注单<br/>4.关键进程未打开';
        } else {//结算失败
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '批量结算失败,所选注单中没有可结算的注单';
        }

        return json_encode($result);
    }

    //批量回滚
    public function actionBatch_rollback() {
        $request = Yii::$app->request->post();
        $ids = ArrayHelper::getValue($request, 'ids', '');

        $result['ErrorCode'] = 1;
        if (empty($ids)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'id参数缺失';
            return json_encode($result);
        }
        $new_data = self::batchData($ids);

        //循环结算
        $periods_count = $new_data['all_count'];
        unset($new_data['all_count']);
        $count = 0;
        foreach ($new_data as $type => $val) {
            foreach ($val as $periods => $ids) {
                if (is_numeric($periods) && is_array($ids)) {
                    $res = self::betApi('rollback', $type, $periods, $ids);
                    $res = json_decode($res, true);
                    if ($res['ErrorCode'] == 1 && isset($res['ok'])) {
                        $count += $res['ok']; //回滚成功的条数
                    }
                }
            }
        }

        $result = array();
        if ($count == $periods_count) { //全部回滚完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '批量回滚成功';
        } elseif ($count > 0) {//部分回滚完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '回滚成功:&nbsp;' . $count . '&nbsp;条' . '<br/>' . '回滚失败:&nbsp;' . ($periods_count - $count) . '&nbsp;条<br/>可能失败原因:<br/>1.该注单没有结算或者已经初始化<br/>2.非法注单<br/>3.关键进程未打开';
        } else {//回滚失败
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '批量回滚失败，所选注单中没有可以回滚的注单或进程未开启';
        }

        return json_encode($result);
    }

    //按期数进行回滚和结算
    public function actionBy_periods() {
        $post = Yii::$app->request->post();
        $todo = isset($post['todo']) ? $post['todo'] : null;
        $type = isset($post['type']) ? $post['type'] : null;
        $periods = isset($post['periods']) ? $post['periods'] : null;

        if ( empty($periods) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '期数不能为空';
            return json_encode($result);
        }
        if ( empty($todo) ||  !is_numeric($todo) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '处理方式不明确';
            return json_encode($result);
        }
        if ( empty($type) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '彩种类型不能为空！';
            return json_encode($result);
        }


        if ($todo == 1) { // 整期结算
            $todo = 'allbalance';
        } elseif ($todo == 2) {//整期回滚
            $todo = 'rollback';
        } else {
            $todo = null;
        }

        if (!empty($todo)) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '程序正在运行中,请稍候查看结果';
            echo json_encode($result);
            $res = self::betApi($todo, $type, $periods);
            // if($res){
            //     return $res;
            // }
        } else {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '处理方式不明确';
            return json_encode($result);
        }
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
     * **********************************************************
     *  未结算注单查询           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionNobalance() {
        $get = Yii::$app->request->get();
        $fc_type = isset($get['fc_type']) ? $get['fc_type'] : 'liuhecai';
        $starttime = isset($get['s_addtime']) ? $get['s_addtime'] : null;
        $endtime = isset($get['e_addtime']) ? $get['e_addtime'] : null;
        // 如果没有传时间将以今天0点开始截止时间取现在
        if(empty($starttime)){
            $starttime = strtotime(date('Y-m-d'));
        }else{
            $starttime = strtotime($starttime) ? strtotime($starttime) : time();
        }
        if(empty($endtime)){
            $endtime = time();
        }else{
            $endtime = strtotime($endtime) ? strtotime($endtime) : time();
        }
        $addday = date('Ymd', $starttime);
        $endday = date('Ymd', $endtime);
        $between_day = $endday - $addday;
        if($between_day < 0 ){
            $addday = $endday;
        }elseif($between_day > 4){
            $addday = $endday - 4;
        }

        $type_arr = $this->getAllFcTypes(); 
        $new_type_arr = array();
        foreach($type_arr as $key=>$val){
            $new_type_arr[$val['type']] = $val['name'];
        }  
        
        // 查询所有的未结算注单
        //拼接查询条件
        $data = array();
        $sql = 'select fc_type,periods,count(id) as cid from my_bet_record where js=1 and status=1 and addday>=' . "'{$addday}'"  . ' and addday<=' . "'{$endday}'" . ' group by fc_type,periods order by periods desc';
        $arr = BetModel::getNoBalanceBet($sql);
        if(empty($arr)){
            $render = [
                'data' => array(),
                'starttime' =>date('Y/m/d',strtotime($addday)),
                'endtime' => date('Y/m/d',strtotime($endday)),
                'fc_type' => $type_arr
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('notbalance.html', $render);
            } else {
                return $this->render('notbalance.html', $render);
            }
        }
        $data_arr = array();
        foreach($arr as $key=>$v){
                $data_arr[$v['fc_type']][$v['periods']] = $v['cid'];
        }
        // data格式:array[彩种][期数][id]
        //查询对应期数是否有开奖结果 有开奖结果而没结算的原因
        $last_data = array();
       
        foreach ($data_arr as $type => $val) {
            foreach ($val as $periods => $bet_count) {
                //从数据库查询，查看是否拥有开奖结果 有的话查询出开奖结果的插入时间及结算状态
                $auto_data = array();
                $auto_data = BetModel::getAutoData($type, $periods);
                if ($auto_data) {
                    $last_data[$type][$periods]['kaijiang'] = '已开奖'; //有开奖结果
                    if ($auto_data['status'] == 3) {
                        $last_data[$type][$periods]['msg'] = '正在结算中'; //有开奖结果正常结算中
                    } elseif ($auto_data['status'] == 2) {
                        $last_data[$type][$periods]['msg'] = '注单未正常结算'; //正常结算完成，可能是回滚注单异常注单
                    } elseif ($auto_data['status'] == 1) {
                        $last_data[$type][$periods]['msg'] = '结算异常，未触发结算'; //未开启结算
                    } else {
                        $last_data[$type][$periods]['msg'] = '未知原因，联系管理员'; //其它原因
                    }
                    //判断开奖结果是否超出3分钟
                    if ((time() - $auto_data['addtime']) > 180) {
                        $last_data[$type][$periods]['other'] = 1; //超过 异常
                    } else {
                        $last_data[$type][$periods]['other'] = 2; //未超过
                        $last_data[$type][$periods]['msg'] = '正在请求结算进程'; //未开启结算
                    }
                } else {
                    $last_data[$type][$periods]['msg'] = null; //正常结算完成，可能是回滚注单异常注单
                    $last_data[$type][$periods]['kaijiang'] = '未开奖'; //无开奖结果
                    $last_data[$type][$periods]['other'] = 2; //未开奖
                }
                $last_data[$type][$periods]['count'] = $bet_count;
                //createCommand()->getRawSql();;
            }
        }
        
        $render = [
            'data' => $last_data,
            'starttime' =>date('Y/m/d',strtotime($addday)),
            'endtime' => date('Y/m/d',strtotime($endday)),
            'fc_type' => $type_arr,
            'new_fc_type' => $new_type_arr
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('notbalance.html', $render);
        } else {
            return $this->render('notbalance.html', $render);
        }
    }

    /**
     * **********************************************************
     *  ajax根据期数彩种请求结算 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionSecondblance() {
        $post = Yii::$app->request->post();
        $fc_type = isset($post['type']) ? $post['type'] : null; //彩种
        $periods = isset($post['periods']) ? $post['periods'] : null; //期数
        $result = array();
        if ((!is_numeric($periods)) || empty($fc_type) || empty($periods)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '数据不合法';
            return json_encode($result);
        }
        //调用结算接口 尝试第二次结算
        $auto_data = BetModel::getAutoData($fc_type, $periods);
        $auto_status = $auto_data['status'];
        if ($auto_status == 3) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '自动结算进程正在进行中，请稍候重试。如果十分钟后仍然不可结算，请检查开奖结果状态值';
            return json_encode($result);
        }
        // return $res;
        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '程序正在运行中,请稍候查看结果';
        echo json_encode($result);

        $res = self::betApi('allbalance', $fc_type, $periods);
    }



    /**
     * **********************************************************
     * 设置注单无效              @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionInvalid() {
        $post = Yii::$app->request->post();
        $act = isset($post['act']) ? $post['act'] : null;
        $ids = isset($post['ids']) ? $post['ids'] : null;

        $result = array();
        if ( empty($act) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '处理方式不明确';
            return json_encode($result);
        }
        if ( empty($ids) ) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = 'id参数缺失';
            return json_encode($result);
        }

        //处理数据
        $new_data = self::batchData($ids);
        //循环结算
        unset($new_data['all_count']);
        $id_arr = array();
        foreach ($new_data as $type => $val) {
            $type = $type;
            foreach ($val as $periods => $ids) {
                foreach ($ids as $key => $id) {
                    $id_arr[] = $id;
                }
            }
        }
        $id_count = count($id_arr); //所有选中的注单
        $where = array();
        $set = array();
        if ($act == 'invalid') {//批量设置无效
            $where['js'] = 1;
            $where['status'] = 1;
            $where = ['and', ['in', 'id', $id_arr], $where];
            $return_str = '设置无效';
            $set['status'] = 5;
        } elseif ($act == 'ok') {//批量恢复
            $where['js'] = 1;
            $where['status'] = 5;
            $where = ['and', ['in', 'id', $id_arr], $where];
            $return_str = '恢复无效注单';
            $set['status'] = 1;
        }

        $res = BetModel::invalid($type, $set, $where);
        $result = array();
        $log_id = implode(',&nbsp;', $id_arr);
        if ($res == $id_count) { //全部完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '批量' . $return_str . '成功';
            parent::insertOperateLog('', '', '操作彩种:' . $type . $result['ErrorMsg'] . ',注单id:' . $log_id);
            ;
        } elseif ($res > 0) {//部分完成
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = $return_str . '成功:&nbsp;' . $res . '&nbsp;条' . '<br/>' . $return_str . '失败:&nbsp;' . ($id_count - $res) . '&nbsp;条<br/>可能失败原因:<br/>1.该注单已经结算过<br/>2.该注单状态未变动';
            parent::insertOperateLog('', '', '操作彩种:' . $type . $result['ErrorMsg'] . ',注单id:' . $log_id);
            ;
        } else {//操作失败
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '选中' . $id_count . '条,批量' . $return_str . '全部失败';
            parent::insertOperateLog('', '', '操作彩种:' . $type . $result['ErrorMsg'] . ',注单id:' . $log_id);
            ;
        }

        return json_encode($result);
    }

}
