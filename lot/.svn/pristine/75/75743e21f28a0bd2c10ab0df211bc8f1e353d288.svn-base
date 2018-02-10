<?php

namespace frontend_api\controllers;

use Yii;
use frontend_api\controllers\Controller;
use \common\helpers\Helper;
use frontend_api\models\BetModel;
use frontend_api\models\UserModel;
use frontend_api\models\CommonModel;
use common\helpers\IdWork;
use common\helpers\Encrypt;
use common\helpers\TcpConPoll;
/**
 * Bet controller
 */
class BetController extends Controller {

    //下注页面彩种
    public function actionAlltype() {
        //所有彩种
        $alltypes = $this->getAllFcTypes();
        //我喜欢的
        $liketypes = $this->getMyLovesFc();
        $like_arr = explode(',', $liketypes['fc_types']);

        $fina = $data = [];
        //判断是否维护
        $line_id = $this->user->line_id;
        $ptype = $this->user->client;
         //所有彩种维护信息
        $maintain_data = $this->getAllTypeMaintain(2, $ptype, $line_id);

        foreach ($alltypes as $k => $v) {
            $data = [
                "type" => $v['type'],
                "name" => $v['name'],
                "is_hot" => $v['is_hot'],
                "order_by" => $v['order_by'],
                "img_path" => $v['img_path'],
                "ltype" => $v['ltype']
            ];
            $is_maintain = $maintain_data[$v['type']];
            $data['is_wh'] = $is_maintain['return'];
            $data['wh_content'] = $is_maintain['remark'];
            if (in_array($v['type'], $like_arr)) {
                $data['is_like'] = 2;
            } else {
                $data['is_like'] = 1;
            }
            $fina[] = $data;
        }

        $echo['Data'] = $fina;
        $echo['ErrorCode'] = 1;
        $echo['ErrorMsg'] = '获取数据成功';

        echo json_encode($echo);
        die;
    }

    //根据彩种type类型获取赔率 开关盘时间 开奖结果 期数等信息
    public function actionGetgameindex() {
        $params = Yii::$app->request->post();
        // $get = Yii::$app->request->get();
        // $params = !empty($post) ? $post : $get;
        $fc_type = isset($params['fc_type']) ? $params['fc_type'] : 'fc_3d';
        $gameplay = isset($params['gameplay']) ? intval($params['gameplay']) : null;
        $pankou = isset($params['pankou']) ? $params['pankou'] : 'A';
        $echo['Data'] = array();
        $echo['ErrorCode'] = 2;

        //判断是否维护
        $line_id = $this->user->line_id;
        $ptype = $this->user->client;
        $is_maintain = $this->game_is_maintain(2,$ptype,$line_id,$fc_type);
        $echo['is_wh'] = $is_maintain['return'];
        $echo['wh_content'] =   $is_maintain['remark'];

        if (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai')) {
            $echo['ErrorMsg'] = '获取该玩法相关数据失败';
            //六合彩只返回一个玩法数据，检测玩法是否存在
            if (empty($gameplay)) {
                echo json_encode($echo);
                die;
            }
            $where = ['and', ['id'=>$gameplay], ['or', ['fc_type' => 'liuhecai'], ['fc_type' => 'jsliuhecai']] ];
            $gameplay = BetModel::getGameplayById($where);
            if (!$gameplay) {
                echo json_encode($echo);
                die;
            }
        }

        //赔率
        $odds = $this->getOddsByFcType($fc_type, $pankou, $gameplay);

        if (!$odds) {
            $echo['ErrorMsg'] = '获取该列表相关数据失败';
            echo json_encode($echo);
            die;
        }
        foreach ($odds as $key => $val) {
            $odds[$key]['money'] = ''; //前端要求加入些字段
        }
        if (in_array($fc_type, ['bj_kl8', 'dm_klc', 'jnd_bs'])) {
            $odds = self::getNewOdds($odds);
        }
        if (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'JointMark') {
            $odds = self::getLianMaOdds($odds); //连码二中特特殊赔率处理
        }
        if(($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && in_array($gameplay, ['EndNum','EndNum_Even'])){//尾数 尾数连处理input_name
            $odds = self::getEndNum($odds);
        }
        //期数
        $qishu = $this->get_fc_qishu($fc_type);
        //封盘时间
        $time_arr = $this->getFengpanByQishu($fc_type, $qishu);
        if (!$time_arr) {
            $echo['ErrorMsg'] = '获取开封盘时间错误';
            echo json_encode($echo);
            die;
        } else {
            $closetime = $time_arr;
        }
        //如果当前时间在当前期数的封盘时间和开奖时间之间，所投注单为下一期
        if ( time() >= strtotime($time_arr['fengpan']) && time() <= strtotime($time_arr['kaijiang']) ) {
            $qishu = self::getNextQishu($fc_type, $qishu);
            $closetime = $this->getFengpanByQishu($fc_type, $qishu);
            if (!$closetime) {
                $echo['ErrorMsg'] = '获取开盘时间错误';
                echo json_encode($echo);
                die;
            }
        }elseif( (time() > strtotime($closetime['fengpan'])) && in_array($fc_type,['bj_kl8', 'bj_10', 'bj_28', 'pc_28', 'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'dj_o', 'xdl_10', 'mg_o', 'mnl_o']) ){
            $qishu = self::getNextQishu($fc_type, $qishu);
            $closetime = $this->getFengpanByQishu($fc_type, $qishu);
        }

        //彩种及图片 当前期数
        $games_arr = $this->getAllFcTypes();
        if (!$games_arr) {
            $echo['ErrorMsg'] = '获取彩种信息失败';
            echo json_encode($echo);
            die;
        }
        foreach ($games_arr as $val) {
            $new_games_arr[$val['type']] = $val;
        }
        $tmp_arr = array();
        $tmp_arr = $new_games_arr[$fc_type];
        $c_data = array();
        $c_data['fc_name'] = $tmp_arr['name'];
        $c_data['qishu'] = $qishu;
        $c_data['img_path'] = $tmp_arr['img_path'];

        $new_closetime = array();
        $new_closetime['kaipan'] = strtotime($closetime['kaipan']);
        $new_closetime['fengpan'] = strtotime($closetime['fengpan']);
        $new_closetime['kaijiang'] = strtotime($closetime['kaijiang']);
        $new_closetime['now_time'] = time();

        $lastAuto = $this->getLastAutoByType($fc_type); //最近一期开奖结果
        if(in_array($fc_type, ['bj_10', 'xdl_10', 'jsfc'])){
            if(isset($lastAuto['ball'])){
                foreach($lastAuto['ball'] as $key=>$val){
                    $lastAuto['ball'][$key]['ball'] = $val['number'];
                    unset($lastAuto['ball'][$key]['number']);
                }
            }
        }
        $echo['Data'] = ['odds' => $odds, 'c_data' => $c_data, 'closetime' => $new_closetime, 'auto' => $lastAuto];
        if (in_array($fc_type, ['liuhecai', 'jsliuhecai'])) {
            if ($gameplay == 'EndNum_Even' || $gameplay == 'EndNum') {
                $echo['Data']['endnum'] = self::endNum();
            }
            if (in_array($gameplay, ['Animal', 'Te_Animal', 'SumAnimal', 'AnimalEven', 'Just_Animal', 'All_Animal'])) {
                $echo['shengxiao'] = self::get_animal(true);
            }
            if (in_array($gameplay, ['Tema', 'JustCode', 'JustCode_Te', 'JointMark'])) {
                $num_arr = array();
                for ($i = 1; $i <= 49; $i++) {
                    $num_arr[] = $i;
                }
                $wave_arr = self::wave();
                foreach ($odds as $key => $val) {
                    if (is_numeric($val['input_name']) && in_array($val['input_name'], $num_arr)) {
                        foreach ($wave_arr as $color => $v) {
                            if (in_array($val['input_name'], $v)) {
                                $odds[$key]['color'] = $color;
                            }
                        }
                    }
                }
                $echo['shengxiao'] = self::get_animal();
                $echo['Data']['odds'] = $odds;
            }
        }else if(in_array($fc_type, ['bj_10', 'jsfc', 'xdl_10'])){//返回颜色
            $bj_play_arr = ['first','second','third','fourth','fifth','sixth','seventh','eighth','ninth','tenth'];
            $bj_color_arr = ['','bj-yellow','bj-blue','bj-black','bj-orange','bj-azure','bj-deepblue','bj-silver','bj-red','bj-brown','bj-green'];
            foreach($odds as $key=>$val){
                if(in_array($val['gameplay'], $bj_play_arr) && is_numeric($val['input_name'])){
                    $odds[$key]['color'] = isset($bj_color_arr[$val['input_name']]) ? $bj_color_arr[$val['input_name']] : '';
                }
            }
            $echo['Data']['odds'] = $odds;
        }else if($fc_type == 'pc_28'){//返回颜色
            $wave_arr = array();
            $wave_arr['pc-red'] = [3 , 6 , 9 , 12 , 15 , 18 , 21 , 24];
            $wave_arr['pc-green'] = [1 , 4 , 7 , 10 , 16 , 19 , 22 , 25];
            $wave_arr['pc-blue'] = [2 , 5 , 8 , 11 , 17 , 20 , 23 , 26];
            foreach($odds as $key=>$val){
                if(is_numeric($val['input_name'])){
                    $odds[$key]['color'] = 'pc-white';
                    foreach($wave_arr as $k => $v){
                        if(in_array($val['input_name'], $v)){
                            $odds[$key]['color'] = $k;
                        }
                    }
                }
            }
            $echo['Data']['odds'] = $odds;
        }
        $echo['ErrorCode'] = 1;
        $echo['ErrorMsg'] = '获取数据成功';
        echo json_encode($echo);
        die;
    }

    /**
     * **********************************************************
     *  下注           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionAddbet() {
        $post = Yii::$app->request->post();
        $fc_type = isset($post['fc_type']) ? trim($post['fc_type'], '"') : null;
        $qishu = isset($post['qishu']) ? trim($post['qishu'], '"') : null;
        $bet_data = isset($post['data']) ? $post['data'] : null;
        $result = array();
        $result['ErrorCode'] = 2;

        if (empty($fc_type) || empty($qishu) || empty($bet_data)) {
            $result['ErrorMsg'] = '参数缺失';
            echo json_encode($result);
            die;
        }
        $games_arr = $this->getAllFcTypes();
        $new_games_arr = $type_arr = array();
        foreach($games_arr as $val){
            $new_games_arr[$val['type']] = $val['name'];
            $type_arr[] = $val['type'];
        }
        if(!in_array($fc_type, $type_arr)){
            $result['ErrorMsg'] = '未知彩种';
            echo json_encode($result);
            die;
        }

        $bet_data = json_decode($bet_data, true);

        // 获取用户uid
        $uid = $this->user->uid;
        $line_id = $this->user->line_id;
        $ptype = $this->user->client;
        $is_shiwan = $this->user->is_shiwan;//是否试玩
        $is_wallet =  $this->user->is_wallet; //是否钱包模式
        if (!$uid) {
            $result['ErrorMsg'] = '请登录后操作';
            echo json_encode($result);
            die;
        }
        //查看彩种是否维护
        $is_maintain = $this->game_is_maintain(2,$ptype,$line_id,$fc_type);
        if($is_maintain){
            $maintain = isset($is_maintain['return']) ? $is_maintain['return'] : 0;
            $maintain_txt = isset($is_maintain['remark']) ? $is_maintain['remark'] : '一般维护';
            if($maintain == 2){
                $result['ErrorMsg'] = '该彩种正在维护中';
                $result['is_wh'] = 2;
                $result['wh_content'] =   $maintain_txt;
                echo json_encode($result); die;
            }
        }
        //阻止注单重复提交(将数据保存到redis 5秒，如果两次提交的数据完全相同，视为重复)
        $redis = \Yii::$app->redis;
        $redis_key = 'tmp_bet_' . $uid;
        $is_exist = $redis->get($redis_key);
        if ($is_exist == json_encode($bet_data)) {
            $result['ErrorMsg'] = '请勿重新下注';
            echo json_encode($result);
            die;
        } else {
            $redis->setex($redis_key, 5, json_encode($bet_data));
        }

        //获取用户信息
        $user_info = UserModel::getUserInfo(['uid' => $uid]);
        if (!$user_info) {
            $result['ErrorMsg'] = '用户信息异常';
            echo json_encode($result);
            die;
        }
        if($user_info['status'] != 1){
            $die_str = $user_info['remark'] ? $user_info['remark'] : '未知';
            $result['ErrorMsg'] = '抱歉,您的帐号被封停,封停原因:' . $die_str;
            echo json_encode($result);
            die;
        }

        //查询代理名称
        $agent_info = BetModel::getAgentInfo($user_info['agent_id']);
        $agent_name = $agent_info['login_user'];
        $agent_money = $agent_info['money'];
        
        if (!$agent_name) {
            $result['ErrorMsg'] = '代理信息异常';
            echo json_encode($result);
            die;
        }
        //验证注单合法性
        $bet_sum_money = 0; //总下注金额
        if(! $is_wallet){
            $user_money = $user_info['money'];
        }

        $did_str = '';
        foreach ($bet_data as $key => $val) {
            $tmp = false;
            if (empty($val['money'])) {
                $result['ErrorMsg'] = '下注金额低于最小金额1元';
                echo json_encode($result);
                die;
            }
            $money = intval($val['money']);
            if ($money < 1) {
                $result['ErrorMsg'] = '下注金额低于最小金额1元';
                echo json_encode($result);
                die;
            }
            $bet_sum_money += $money;
            $tmp = self::check_bet($fc_type, $bet_sum_money, $val, $user_info['agent_id']);
            if (!$tmp) {
                $result['ErrorMsg'] = '注单无效';
                echo json_encode($result);
                die;
            }
            $bet_data[$key]['play_id'] = $tmp['play_id'];
            if($is_wallet){//钱包模式不请求余额
                $bet_data[$key]['assets'] = 0;
                $bet_data[$key]['balance'] = 0;
            }else{
                $bet_data[$key]['assets'] = $user_money; //下注前金额
                $user_money -= $money;
                if ($user_money < 0) {
                    $result['ErrorMsg'] = '您的余额不足，请充值';
                    echo json_encode($result);
                    die;
                }
                $bet_data[$key]['balance'] = $user_money; //下注后金额
            }
            $bet_data[$key]['order_num'] = $tmp['order_num'];
            if ($val['gameplay'] == 'Tema' && $val['pankou'] == 'B') {
                $bet_data[$key]['return_water'] = $money * '0.01';
            } else {
                $bet_data[$key]['return_water'] = 0;
            }

            $did_str .= $tmp['order_num'] . ',';
        }

        //获取当前的期数与开封盘时间
        $tmp_qishu = $this->get_fc_qishu($fc_type); //当前期数
        $game_time = $this->getFengpanByQishu($fc_type, $tmp_qishu);
        //检测开封盘时间是否正确
        if (!$game_time) {
            $result['ErrorMsg'] = '获取开封盘时间错误';
            echo json_encode($result);
            die;
        }
        //如果当前时间在当前期数的封盘时间和开奖时间之间，所投注单为下一期
        if (time() >= strtotime($game_time['fengpan']) && time() < strtotime($game_time['kaijiang'])) {//进入下一期
            $tmp_qishu = self::getNextQishu($fc_type, $tmp_qishu);
        }elseif( (time() > strtotime($game_time['fengpan'])) && in_array($fc_type,['bj_kl8', 'bj_10', 'bj_28', 'pc_28', 'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'dj_o', 'xdl_10', 'mg_o', 'mnl_o']) ){
            $tmp_qishu = self::getNextQishu($fc_type, $tmp_qishu);
            $game_time = $this->getFengpanByQishu($fc_type, $tmp_qishu);
        }
        //有些彩种如果是当天最后一期，获取的期数和开封盘时间是明天的，所以取消小于开盘时间条件
        if (time() > strtotime($game_time['kaijiang'])) {
            $result['ErrorMsg'] = '获取开封盘时间错误!!!';
            echo json_encode($result);
            die;
        }

        if ($tmp_qishu != $qishu) {
            $result['ErrorMsg'] = '期数错误,请刷新页面重试';
            echo json_encode($result);
            die;
        }

        $ip = Helper::getIpAddress();
        $ip = sprintf('%u',ip2long($ip));//转成整形
        //拼接注单sql
        $insert = array();
        $field = [
            'line_id', 'at_id', 'ua_id', 'sh_id', 'at_username', 'uid',
            'uname', 'order_num', 'bet', 'valid_bet', 'balance', 'assets',
            'fc_type', 'odds', 'periods', 'win', 'handicap',
            'addtime', 'addday', 'updatetime', 'updateday', 'bet_info',
            'ptype', 'js', 'status', 'return_water', 'bet_ip', 'play_id', 'is_shiwan'
        ];

        $insert_sql = ''; //原生sql拼接
        foreach ($bet_data as $key => $val) {
            $gameplay = $val['gameplay'];
            $pankou = isset($val['pankou']) ? $val['pankou'] : 'A';
            if ($pankou == 'A') {
                $pankou = 1;
            } else {
                $pankou = 2;
            }
            //处理特殊玩法的赔率
            if (
                    (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') &&
                    $val['gameplay'] == 'JointMark' &&
                    stripos($val['mingxi'], '/')) ||
                    ( in_array($fc_type, ['bj_kl8', 'jnd_bs', 'dm_klc']) &&
                    stripos($val['mingxi'], '/') )
            ) {
                $odds = explode(',', $val['mingxi']);
                $odds = explode(':', $odds[0]);
                $odds = $odds[1];
            } elseif (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'AnimalEven') {
            //生肖连的赔率：如果选中的有本命年的生肖，则使用本命年的生肖的赔率，因为本命年生肖多一个数字       
                $odds_arr = $this->getOddsByFcType($fc_type, 'A', $gameplay);
                $odds = self::getAnimalEvenOdds($odds_arr, $val['input_name'], $val['mingxi']);
            } elseif(($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'PassTest'){
                //过关的赔率
                $odds_arr = $this->getOddsByFcType($fc_type, 'A', $gameplay);
                $odds = self::getPassTestOdds($odds_arr, $val['input_name'], $val['mingxi']);
            } else {
                //正常赔率
                $odds = $val['odd'];
            }

            $bet_info = $val['gameplay'] . '#' . $val['input_name'] . '#' . $val['mingxi'];
            $insert[] = [
                $user_info['line_id'], $user_info['agent_id'], $user_info['ua_id'],
                $user_info['sh_id'], $agent_name, $user_info['uid'], $user_info['uname'],
                $val['order_num'], $val['money'], $val['money'],
                $val['balance'], $val['assets'], $fc_type, $odds, $qishu,
                $val['money'] * $odds, $pankou, time(), date('Ymd'),
                time(), date('Ymd'), $bet_info, $ptype, 1, 1,
                $val['return_water'], $ip, $val['play_id'], $is_shiwan
            ];

            //原生sql用于mycat
            $sql_field = 'insert into `lottery`.`my_bet_record` ( `id`, `line_id`, `at_id`, `ua_id`, `sh_id`, `at_username`, `uid`, `uname`, `order_num`, `bet`, `valid_bet`, `balance`, `assets`, `fc_type`, `odds`, `periods`, `win`, `handicap`, `addtime`, `addday`, `updatetime`, `updateday`, `bet_info`, `ptype`, `js`, `status`, `return_water`, `bet_ip`, `play_id`, `is_shiwan`) values ';
            $insert_sql .=  $sql_field .  '(' . "next value for MYCATSEQ_BETRECORD,'" . $user_info['line_id'] . "', " .  $user_info['agent_id'] . "," . $user_info['ua_id'] . "," . $user_info['sh_id'] . ", '" . $agent_name . "', " . $user_info['uid'] . ", '" . $user_info['uname'] . "', " . $val['order_num'] . ",'" . $val['money'] . "', '" . $val['money'] . "', '" .  $val['balance'] . "', '" . $val['assets'] . "', '" . $fc_type  . "', " . $odds . "," . $qishu  . ", '" . $val['money'] * $odds .  "', " . $pankou . ", " . time() . ',' . date('Ymd') . ', ' . time() . ', ' .  date('Ymd') . ", '" . $bet_info . "', " . $ptype . ", 1, 1, '" . $val['return_water'] . "', " . $ip . "," . $val['play_id'] . ', ' . $is_shiwan . ')###';
        }


        $transaction = Yii::$app->db->beginTransaction();

        try{
            //备注内容（现金纪录及请求接口额度转换）
            $work = new IdWork(1023);
            $count = count($bet_data);
            if ($count > 1) {
                $remark_did = $bet_data[0]['order_num'] . '~' . $bet_data[$count - 1]['order_num'];
            } else {
                $remark_did = $bet_data[0]['order_num'];
            }
            //扣除会员金额
            if(!$is_wallet){
                $field_money =  new \yii\db\Expression('money-' . $bet_sum_money);
                //使用注释掉的sql
                // $changeMoney = BetModel::updateUserMoney(['money' => $field_money], [ 'and', ['and',['=', 'uid', $uid], ['>=', 'money', $bet_sum_money]], ['=','agent_id', intval($user_info['agent_id'])] ]);
                $changeMoney = BetModel::updateUserMoney(['money' => $field_money], [ 'and',['=', 'uid', $uid], ['>=', 'money', $bet_sum_money] ]);
                if (!$changeMoney) {
                    $transaction->rollBack();
                    $result['ErrorMsg'] = '扣除会员金额失败';
                    echo json_encode($result);
                    die;
                }
            }else{
                //钱包模式，请求扣除会员金额
                $fc_type_name = $new_games_arr[$fc_type];
                $remark = '彩票注单(下注):' . $remark_did . ',类型:' . $fc_type_name  . ',共计:' . $count . '单'; //Lottery 
                $requestId = $work->nextId();
                $requestId = strval($requestId);
                // $send_info = Encrypt::Transfer( $user_info['uname'], "CNY", 0 - $bet_sum_money , $requestId ,$remark);
                // $socket = TcpConPoll::getInstace();
                // $socket_res = $socket::send($send_info);
                // $socket_res = json_decode($socket_res, true);
                // if(!isset($socket_res['code']) || $socket_res['code'] != 1000 || !isset($socket_res['member']['balance'])){
                //     $transaction->rollBack();
                //     $result['ErrorMsg'] = '操作会员金额失败';
                //     echo json_encode($result);
                //     die;
                // }
                // $wallet_money = $socket_res['member']['balance'] + $bet_sum_money;
                $wallet_money = 1000;
            }


          
            $cash = array(); //现金记录数据
            $remark = 'Lottery note：' . $remark_did . ' , type:' . $fc_type . ',共计:' . $count . '单'; //Lottery note：17100618015013618717~17100618015045352665 , type:bj_kl8 .',共计:'.4.'单'
            if($is_wallet) $remark .= "(requestId:$requestId)";
            $cash['uid'] = $uid;
            $cash['line_id'] = $user_info['line_id'];
            $cash['agent_id'] = $user_info['agent_id'];
            $cash['uname'] = $user_info['uname'];
            $cash['fc_type'] = $fc_type;
            $cash['periods'] = $qishu;
            $cash['cash_type'] = 2;
            $cash['cash_do_type'] = 1;
            $cash['dids'] = rtrim($did_str, ',');
            $cash['cash_num'] = $bet_sum_money;
            $cash['remark'] = $remark;
            $cash['ptype'] = $ptype;
            $cash['addtime'] = time();
            $cash['addday'] = date('Ymd');


            //查看是不是mycat分库（插入注单表数据）
            $is_mycat = isset(Yii::$app->params['is_mycat']) ? Yii::$app->params['is_mycat'] : false;
            if($is_mycat){
                $insertBet = 0;
                $sql_arr = explode('###', rtrim($insert_sql, '###'));
                foreach($sql_arr as $sql){
                    $res = BetModel::insert($sql);
                    if($res) $insertBet += 1;
                }
                if (count($sql_arr) != $insertBet) {
                    $transaction->rollBack();
                    if($is_wallet) self::errorCashRecord($cash);
                    $result['ErrorMsg'] = '写入注单失败';
                    echo json_encode($result);
                    die;
                }
            }else{
                $insertBet = BetModel::insertBet($field, $insert);
                if (!$insertBet) {
                    $transaction->rollBack();
                    if($is_wallet) self::errorCashRecord($cash);
                    $result['ErrorMsg'] = '写入注单失败';
                    echo json_encode($result);
                    die;
                }
            }

            //写入现金纪录
            if($is_wallet){
                $cash_balance = $wallet_money - $bet_sum_money;
            }else{
                $cash_balance = $user_info['money'] - $bet_sum_money;
            }
            if($is_mycat){
                $cash_sql = "insert into `lottery`.`my_user_cash_record` (`id`, `uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`, `uname`, `fc_type`, `periods`, `is_shiwan`) values ( next value for MYCATSEQ_USERCASHRECORD, " . $uid . ", '" . $user_info['line_id'] . "', " . $user_info['agent_id'] . ", 2, 1, '" . rtrim($did_str, ',') . "', '" . $bet_sum_money . "', '" . $cash_balance . "', '" . $remark . "', " . $ptype . ", " . time() . ", " . date('Ymd') . ", '" . $user_info['uname'] . "', '" . $fc_type . "', " . $qishu . ', ' . $is_shiwan .  ")";
                $insertCash = BetModel::insert($cash_sql);
            }else{
                $user_cash = $cash;
                $user_cash['cash_balance'] = $cash_balance;
                $user_cash['is_shiwan'] = $is_shiwan;
                $insertCash = BetModel::insertCashRecord('user_cash_record',$user_cash);
            }
            if (!$insertCash) {
                $transaction->rollBack();
                if($is_wallet) self::errorCashRecord($cash);
                $result['ErrorMsg'] = '纪录失败';
                echo json_encode($result);
                die;
            }

            //钱包模式增加代理现金纪录  
            if($is_wallet){
                $line_id = $user_info['line_id'];
                $agent_id = $user_info['agent_id'];
                $at_cash_num = $bet_sum_money;
                $at_cash_blance = $agent_money + $bet_sum_money;
                $at_remark = '会员:' . $user_info['uname'] . "下注,彩种:$fc_type, 期数:$qishu";
                //钱包模式更改代理金额
                $at_field_money =  new \yii\db\Expression('money+' . $bet_sum_money);
                $changeAtMoney = BetModel::updateAgentMoney(['money' => $at_field_money], ['=', 'id', $agent_id]);
                if (!$changeAtMoney) {
                    $transaction->rollBack();
                    self::errorCashRecord($cash);
                    $result['ErrorMsg'] = '代理金额变更失败';
                    echo json_encode($result);
                    die;
                }

                if($is_mycat){
                    //代理现金纪录
                    $at_cash_sql = "insert into `lottery`.`my_agent_cash_record` (`id`, `agent_id`, `agent_user`, `line_id`, `hander_id`, `cash_num`, `cash_balance`, `remark`, `addtime`, `addday`, `cash_type`) values ( next value for MYCATSEQ_AGENTCASHRECORD, $agent_id, '$agent_name', '$line_id', $uid, $at_cash_num, $at_cash_blance, '$at_remark', " . time() . ',' . date('Ymd') . ',1)';
                    $insertAtCash = BetModel::insert($cash_sql);
                }else{
                    $at_cash = array();
                    $at_cash['agent_id'] = $agent_id;
                    $at_cash['agent_user'] = $agent_name;
                    $at_cash['line_id'] = $line_id;
                    $at_cash['hander_id'] = $uid;
                    $at_cash['cash_num'] = $at_cash_num;
                    $at_cash['cash_balance'] = $at_cash_blance;
                    $at_cash['remark'] = $at_remark;
                    $at_cash['addtime'] = time();
                    $at_cash['addday'] = date('Ymd');
                    $at_cash['cash_type'] = 1;

                    $insertAtCash = BetModel::insertCashRecord('agent_cash_record', $at_cash);
                }
                if (!$insertCash) {
                    $transaction->rollBack();
                    if($is_wallet) self::errorCashRecord($cash);
                    $result['ErrorMsg'] = '代理纪录失败';
                    echo json_encode($result);
                    die;
                }

            }


            //组装下注注单数据
            $old_bets = array();
            foreach($insert as $val){
                $old_bets[] = array_combine($field, $val);
            }

            $pullredis = \Yii::$app->pullredis;
            $mongo_data = array();
            //存储用于外部采集注单的数据
            foreach ($old_bets as $key => $val) {
                $score =  $work->nextId();
                $score = substr($score, -16);//score最大范围16位，超出自动转科学计数
                foreach($val as $k=>$v){
                    if(is_numeric($v)) $val[$k] = strval($v);
                }
                $val['score'] = strval($score);
                $res = $pullredis->zadd('spider_' . $user_info['line_id'] . '_data', $score, json_encode($val));
                if (!$res) {
                    $transaction->rollBack();
                    if($is_wallet) self::errorCashRecord($cash);
                    $result['ErrorMsg'] = '缓存数据失败';
                    echo json_encode($result);
                    die;
                }
                $mongo_data[] = $val;
            }
            //插入Mongo bets表
            if(empty($mongo_data)){
                $transaction->rollBack();
                if($is_wallet) self::errorCashRecord($cash);
                $result['ErrorMsg'] = '获取备份数据失败';
                echo json_encode($result);
                die;
            }
            $res = BetModel::insertMongoBets($mongo_data);
            if (!$res) {
                $transaction->rollBack();
                if($is_wallet) self::errorCashRecord($cash);
                $result['ErrorMsg'] = '备份数据失败';
                echo json_encode($result);
                die;
            }
            //存储用于结算的缓存数据
            $redis_key = 'for_balance-' .  $fc_type . '-' . $qishu;
            $user_key = 'user_'. $uid;
            $user_old_bet = $redis->hget($redis_key, $user_key);
            //如果用户有在当前彩种当前期数下过注，取出与现在的注单合并
            if($user_old_bet){
                $user_old_bet = json_decode($user_old_bet, true);
                $old_bets = array_merge($user_old_bet, $old_bets);
            }
            $redis->hset($redis_key, $user_key, json_encode($old_bets));

        }catch (Exception $e) {
            $transaction->rollBack();
            if($is_wallet) self::errorCashRecord($cash);
            $result['ErrorMsg'] = '下注失败';
            echo json_encode($result);
            die;
        }

        $transaction->commit();

        $result['ErrorMsg'] = '下注成功';
        $result['ErrorCode'] = 1;
        echo json_encode($result);
        die;
    }

    /**
     * **********************************************************
     *  验证注单合法性           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private function check_bet($fc_type, $sum_money, $data, $at_id) {
        $gameplay = $data['gameplay'];
        $pankou = isset($data['pankou']) ? $data['pankou'] : 'A';
        $mingxi = $data['mingxi'];
        $remark = $data['remark'];
        $odd = $data['odd'];
        $money = $data['money'];
        $input_name = $data['input_name'];
        $result = $return = array();
        $result['ErrorCode'] = 2;
        if (empty($gameplay) ||  empty($remark) || empty($odd) || empty($money) || $gameplay == '' || $odd == '' || $money == '') {
            $result['ErrorMsg'] = '关键参数缺失';
            echo json_encode($result);
            die;
        }

        //验证下注内容是否存在
        if($input_name === ''){
            $result['ErrorMsg'] = '请选择下注内容';
            echo json_encode($result);
            die;
        }
        if( mb_strlen($input_name, 'utf-8') >= 250 ){
            $result['ErrorMsg'] = '下注内容过多';
            echo json_encode($result);
            die;
        }
        if($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai'){
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou, $gameplay);
        }else{
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou);
        }
        self::checkAllInput($fc_type, $odds_arr, $input_name, $gameplay);   
        //检测金额
        if (!is_numeric($money)) {
            $result['ErrorMsg'] = '非法金额';
            echo json_encode($result);
            die;
        }
        //验证盘口
        if (!in_array($pankou, ['A', 'B'])) {
            $result['ErrorMsg'] = '不能识别的盘口';
            echo json_encode($result);
            die;
        }
        //验证玩法的合法性
        $play_id = self::getNewPlayId($fc_type, $gameplay);
        if (!$play_id) {
            $result['ErrorMsg'] = '玩法不存在';
            echo json_encode($result);
            die;
        } else {
            $return['play_id'] = $play_id;
        }

        //快乐十分连码 任选
        if (in_array($fc_type, ['gd_ten', 'cq_ten']) && in_array($gameplay, ['random_choose_two', 'random_choose_two_group', 'random_choose_three', 'random_choose_four', 'random_choose_five'])) {

            self::check_ten_inputname($gameplay, $input_name);
        }
        //十一选五任选直选组选
        if (in_array($fc_type, ['sd_11', 'gd_11', 'jx_11']) && in_array($gameplay, ['random_choose', 'group_choose', 'vertical_choose'])) {
            self::check_11_inputname($gameplay, $input_name, $mingxi);
        }
        //PC28特码包三
        if($fc_type == 'pc_28' && $input_name == 'Tema_in_Three'){
            self::check_dd_inputname($mingxi);
        }
        //六合彩尾数连 全不中
        if (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'EndNum_Even') {
            self::check_weilian_inputname($input_name, $mingxi);
        }
        if (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'AllMiss') {
            self::check_allmiss_inputname($input_name, $mingxi);
        }
        //六合彩尾数
        if(($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'EndNum'){
           self::check_endnum_inputname($input_name);
        }

        //验证赔率是否正确
        if (in_array($fc_type, ['bj_kl8', 'dm_klc', 'jnd_bs']) && in_array($gameplay, ['choose_one', 'choose_two', 'choose_three', 'choose_four', 'choose_five'])) {
            //检测用户下注数据是否正确
            self::check_klc_inputname($gameplay, $input_name);
            //快乐彩特殊玩法赔率
            $odds_arr = array();
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou, $gameplay);

            $tmp_arr = array();
            $odds_arr = self::getNewOdds($odds_arr);
            foreach ($odds_arr as $key => $val) {
                $tmp_arr[$val['gameplay']] = $val['odd'];
            }
            if (!isset($tmp_arr[$gameplay])) {
                $result['ErrorMsg'] = '赔率异常-';
                echo json_encode($result);
                die;
            }
            $odds = $tmp_arr[$gameplay];
        } elseif (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'JointMark') {
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou, $gameplay);
            $tmp_arr = array();
            $odds_arr = self::getLianMaOdds($odds_arr);
            foreach ($odds_arr as $key => $val) {
                $tmp_arr[$val['mingxi']] = $val['odd'];
            }
            if (!isset($tmp_arr[$mingxi])) {
                $result['ErrorMsg'] = '赔率异常--';
                echo json_encode($result);
                die;
            }
            $odds = $tmp_arr[$mingxi];
        } elseif (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'AnimalEven') {
            self::check_shengxiaolian_inputname($input_name, $mingxi);
            //生肖连赔率
            $odds_arr = $this->getOddsByFcType($fc_type, 'A', $gameplay);
            $odds = self::getAnimalEvenOdds($odds_arr, $input_name, $mingxi);
        } elseif (($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') &&  $gameplay == 'SumAnimal') {
            //合肖赔率
            self::check_hexiao_inputname($input_name, $mingxi);
            $odds_arr = $this->getOddsByFcType($fc_type, 'A', $gameplay);
            foreach ($odds_arr as $val) {
                $tmp_arr[$val['mingxi']] = $val['odd'];
            }
            $odds = $tmp_arr[$mingxi];
        } elseif(($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'EndNum_Even'){
            //尾数连赔率
            $odds_arr = $this->getOddsByFcType($fc_type, 'A', $gameplay);
            $odds = self::getEndNumOdds($odds_arr , $input_name, $mingxi);
        } elseif(in_array($fc_type, ['gd_ten', 'cq_ten']) && in_array($gameplay, ['random_choose_two', 'random_choose_two_group', 'random_choose_three', 'random_choose_four', 'random_choose_five'])){
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou);
            foreach ($odds_arr as $key => $val) {
                $tmp_arr[$val['gameplay']] = $val['odd'];
            }
            if (!isset($tmp_arr[$gameplay])) {
                $result['ErrorMsg'] = '赔率异常---';
                echo json_encode($result);
                die;
            }
            $odds = $tmp_arr[$gameplay];
        } elseif(($fc_type == 'liuhecai' || $fc_type == 'jsliuhecai') && $gameplay == 'PassTest'){
            $odds = $odd;
        } else {
            //正常赔率
            $odds_arr = $this->getOddsByFcType($fc_type, $pankou);
            foreach ($odds_arr as $key => $val) {
                $tmp_arr[$val['remark']] = $val['odd'];
            }
            if (!isset($tmp_arr[$remark])) {
                $result['ErrorMsg'] = '赔率异常---';
                echo json_encode($result);
                die;
            }
            $odds = $tmp_arr[$remark];
        }
        if (!$odds || (intval($odd) !== intval($odds)) || $odds == 0 || $odds == '0.00') {
            $result['ErrorMsg'] = '赔率异常----';
            echo json_encode($result);
            die;
        }
        //验证有没超出限额
        $limit = $this->getLimitByFcType($fc_type, $pankou);
        $limit_data = isset($limit[$gameplay]) ? $limit[$gameplay] : null;
        if (empty($limit_data)) {
            $result['ErrorMsg'] = '获取限额信息失败';
            echo json_encode($result);
            die;
        }
        if ($money < $limit_data['limit_min'] || $money > $limit_data['single_note_max']) {
            $result['ErrorMsg'] = '单笔下注金额超出范围';
            echo json_encode($result);
            die;
        }
        if ($sum_money > $limit_data['single_field_max']) {
            $result['ErrorMsg'] = '本期下注金额超出范围';
            echo json_encode($result);
            die;
        }
        $work = new IdWork(1023);
        $return['order_num'] =  $work->nextId();

        return $return;
    }

    /**
     * **********************************************************
     *  获取play_id           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function getNewPlayId($fc_type, $gameplay) {
        //获取初始限额并查看play_id是否存在
          //开奖结果存储方式：同一大分类彩种相同玩法的只存其中一个的play_id
        if ($fc_type == 'pl_3')
            $fc_type = 'fc_3d';
        if ($fc_type == 'cq_ten')
            $fc_type = 'gd_ten';
        if ($fc_type == 'jnd_28')
            $fc_type = $gameplay = 'dm_28';
        if (in_array($fc_type, ['dm_klc', 'jnd_bs']))
            $fc_type = 'bj_kl8';
        if (in_array($fc_type, ['xj_ssc', 'tj_ssc']))
            $fc_type = 'cq_ssc';
        if (in_array($fc_type, ['gx_k3', 'js_k3']))
            $fc_type = 'ah_k3';
        if (in_array($fc_type, ['sd_11', 'jx_11']))
            $fc_type = 'gd_11';
        if(in_array($fc_type, ['jsfc','xdl_10']))
            $fc_type = 'bj_10';
        if(in_array($fc_type, ['lfc_o','els_o','dj_o','mnl_o', 'mg_o']))
            $fc_type = 'ffc_o';
        if($fc_type == 'jsliuhecai') $fc_type = 'liuhecai';
        
        $redis = \Yii::$app->redis;
        $init_key = 'init_fc_limit';
        $init_limit = $redis->hget($init_key, $fc_type);
        if (empty($init_limit)) {
            $init_limit = CommonModel::getInitLimitByFcType($fc_type);
            $redis->hset($init_key, $fc_type, json_encode($init_limit, true));
        } else {
            $init_limit = json_decode($init_limit, true);
            $is_exist_id = false;
            foreach($init_limit as $key=>$val){
                if(!isset($val['id'])){
                    $is_exist_id = true;
                    break;
                }
            }
            if($is_exist_id){
                $init_limit = CommonModel::getInitLimitByFcType($fc_type);
                $redis->hset($init_key, $fc_type, json_encode($init_limit, true));
            }
        }


        foreach($init_limit as $key=>$val){
            if($gameplay == $val['gameplay']){
                return $val['id'];
            }
        }
        $result['ErrorMsg'] = '获取玩法id失败';
        $result['ErrorCode'] = '2';
        echo json_encode($result);
        die;
    }

    /**
     * **********************************************************
     *  获取下一期期数           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function getNextQishu($fc_type, $qishu) {
        $type_array = array('tj_ssc', 'cq_ssc', 'xj_ssc', 'gd_ten', 'cq_ten', 'js_k3', 'gx_k3', 'ah_k3', 'gd_11', 'sd_11', 'jx_11', 'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'dj_o', 'mnl_o', 'xdl_10', 'mg_o');
        if (in_array($fc_type, $type_array)) {
            switch ($fc_type) {
                case 'tj_ssc':
                    $end_qishu = 84;
                    $zero = 2;
                    break;
                case 'cq_ssc':
                    $end_qishu = 120;
                    $zero = 2;
                    break;
                case 'xj_ssc':
                    $end_qishu = 96;
                    $zero = 2;
                    break;
                case 'gd_ten':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'cq_ten':
                    $end_qishu = 97;
                    $zero = 2;
                    break;
                case 'js_k3':
                    $end_qishu = 82;
                    $zero = 2;
                    break;
                case 'gx_k3':
                    $end_qishu = 78;
                    $zero = 2;
                    break;
                case 'ah_k3':
                    $end_qishu = 80;
                    $zero = 2;
                    break;
                case 'gd_11':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'sd_11':
                    $end_qishu = 87;
                    $zero = 1;
                    break;
                case 'jx_11':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'ffc_o':
                case 'jsfc' :
                    $end_qishu = 1440;
                    $zero = 1;
                    break;
                case 'lfc_o':
                case 'jsliuhecai':
                    $end_qishu = 720;
                    $zero = 1;
                break;
                case 'els_o' :
                case 'xdl_10':
                    $end_qishu = 960;
                    $zero = 1;
                break;
                case 'dj_o' :
                    $end_qishu = 920;
                    $zero = 1;
                break;
                case 'mnl_o':
                case 'mg_o' :
                    $end_qishu = 1920;
                    $zero = 1;
                break;
            }
            return self::jiajian($fc_type, $qishu, $end_qishu, $zero);
        } else {
            return $qishu + 1;
        }
    }

    /**
     * **********************************************************
     *  当期数增加或减少到尽头时返回正确的期数                 *
     * **********************************************************
     */
    private static function jiajian($fc_type, $qishu, $end_qishu, $zero = 1) {
        $six_arr = ['cq_ten', 'ffc_o', 'jsfc', 'lfc_o', 'mnl_o'];
        $start_qishu = 1;
        $old_qishu = $qishu;
        $str_len = strlen($qishu);
        if (in_array($fc_type, $six_arr)) {
            $date_len = 6;
        } else {
            $date_len = 8;
        }
        $date = substr($qishu, 0, $date_len);
        $qishu = substr($qishu, $date_len, $str_len - $date_len);
        $qishu = intval($qishu);

        if ($qishu == $end_qishu) { //期数自增
            $tomorrow = date('Ymd', strtotime('+ 1days', strtotime($date)));
            if(in_array($fc_type, $six_arr)){
                $tomorrow = date('ymd', strtotime('+ 1days', strtotime($date)));
            }
            if ($zero == 1) {
                $qishu = $tomorrow . str_pad( $now_minute, 2 ,"0",STR_PAD_LEFT );
            } else {
                $qishu = $tomorrow . str_pad( $now_minute, 3 ,"0",STR_PAD_LEFT );
            }

            if( in_array($fc_type, ['ffc_o', 'jsfc', 'mg_o', 'mnl_o']) ){
                $qishu = $tomorrow . '0001';
            }elseif( in_array($fc_type, ['lfc_o', 'els_o', 'xdl_10', 'dj_o', 'jsliuhecai']) ){
                $qishu = $tomorrow . '001';
            }

        } else {
            $qishu = $old_qishu + 1;
        }
        return $qishu;
    }

   

    /**
     * **********************************************************
     *  获取快乐彩特殊玩法的赔率     @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function getNewOdds($odds) {
        $two = '';
        $three_3 = '';
        $three_2 = '';
        $four_4 = '';
        $four_3 = '';
        $four_2 = '';
        $five_5 = '';
        $five_4 = '';
        $five_3 = '';
        foreach ($odds as $key => $val) {
            if (!isset($val['mingxi']))
                continue;
            switch ($val['mingxi']) {
                case 'two_in_two':
                    $two = '2/2:' . floor($val['odd']);
                    break;
                case 'three_in_three':
                    $three_3 = '3/3:' . floor($val['odd']);
                    break;
                case 'three_in_two':
                    $three_2 = '3/2:' . floor($val['odd']);
                    break;
                case 'four_in_four':
                    $four_4 = '4/4:' . floor($val['odd']);
                    break;
                case 'four_in_three':
                    $four_3 = '4/3:' . floor($val['odd']);
                    break;
                case 'four_in_two':
                    $four_2 = '4/2:' . floor($val['odd']);
                    break;
                case 'five_in_five':
                    $five_5 = '5/5:' . floor($val['odd']);
                    break;
                case 'five_in_four':
                    $five_4 = '5/4:' . floor($val['odd']);
                    break;
                case 'five_in_three':
                    $five_3 = '5/3:' . floor($val['odd']);
                    break;
            }
        }
        foreach ($odds as $key => $val) {
            switch ($val['gameplay']) {
                case 'choose_one':
                    $odds[$key]['remark'] = '选一';
                    break;
                case 'choose_two':
                    $odds[$key]['remark'] = '选二';
                    $odds[$key]['odd'] = $two;
                    $odds[$key]['mingxi'] = $two;
                    break;
                case 'choose_three':
                    $odds[$key]['remark'] = '选三';
                    $odds[$key]['odd'] = $three_3 . ',' . $three_2;
                    $odds[$key]['mingxi'] = $three_3 . ',' . $three_2;
                    break;
                case 'choose_four':
                    $odds[$key]['remark'] = '选四';
                    $odds[$key]['odd'] = $four_4 . ',' . $four_3 . ',' . $four_2;
                    $odds[$key]['mingxi'] = $four_4 . ',' . $four_3 . ',' . $four_2;
                    break;
                case 'choose_five':
                    $odds[$key]['remark'] = '选五';
                    $odds[$key]['odd'] = $five_5 . ',' . $five_4 . ',' . $five_3;
                    $odds[$key]['mingxi'] = $five_5 . ',' . $five_4 . ',' . $five_3;
                    break;
                default:

                    break;
            }
        }
        $new_arr = array();
        foreach ($odds as $key => $val) {
            $new_arr[$val['remark']] = $val;
        }
        $odds = array_values($new_arr);
        return $odds;
    }

    /**
     * **********************************************************
     *  获取六合彩连码二中特特殊赔率     @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function getLianMaOdds($odds) {
        $in_te = '';
        $in_two = '';
        $three_in_three = '';
        $three_in_two = '';
        $tmp = array();
        foreach ($odds as $key => $val) {
            if ($val['mingxi'] == 'in_te') {
                $in_te = $val['odd'];
            }
            if ($val['mingxi'] == 'in_two') {
                $in_two = $val['odd'];
            }
            if ($val['mingxi'] == 'in_three') {
                $three_in_three = $val['odd'];
            }
            if ($val['mingxi'] == 'third_in_second') {
                $three_in_two = $val['odd'];
            }
        }
        foreach ($odds as $key => $val) {
            if ($val['mingxi'] == 'in_te' || $val['mingxi'] == 'in_two') {
                $odds[$key]['odd'] = floor($in_te) . '/' . floor($in_two);
                $odds[$key]['remark'] = '连码#二中特#';
                $odds[$key]['mingxi'] = '2/2:' . floor($in_te) . ',' . '2/1:' . floor($in_two);
            }

            if ($val['mingxi'] == 'in_three' || $val['mingxi'] == 'third_in_second') {
                $odds[$key]['odd'] = floor($three_in_three) . '/' . floor($three_in_two);
                $odds[$key]['remark'] = '连码#三中二#';
                $odds[$key]['mingxi'] = '3/3:' . floor($three_in_three) . ',' . '3/2:' . floor($three_in_two);
            }
        }
        $new_arr = array();
        foreach ($odds as $key => $val) {
            $new_arr[$val['remark']] = $val;
        }
        $odds = array_values($new_arr);
        return $odds;
    }

    /**
     * **********************************************************
     *  获取六合彩生肖连的赔率       @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function getAnimalEvenOdds($odds_arr, $input_name, $mingxi) {
        $result['ErrorCode'] = 2;
        //生肖数组
        $animal_arr = array('mouse', 'cattle', 'tiger', 'rabbit', 'dragon', 'snake', 'horse', 'sheep', 'monkey', 'chicken', 'dog', 'pig');
        $year = date('Y');
        $year_animal = $animal_arr[(($year - 4) % 12)]; //获取本命年生肖
        //检测是否在生肖数组内
        $input_arr = explode(',', $input_name);
        if (count($input_arr) < 2) {
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }

        foreach ($input_arr as $val) {
            if (!in_array($val, $input_arr)) {
                $result['ErrorMsg'] = '不能识别的下注内容';
                echo json_encode($result);
                die;
            }
        }

        $other_animal = '';
        $is_year = false;
        if (in_array($year_animal, $input_arr)) {
            $is_year = true; //包含本命年生肖
        } else {
            $other_animal = $input_arr[0]; //除了本命年其它赔率一样
        }
        $tmp_arr = array();
        foreach ($odds_arr as $val) {
            $tmp_arr[$val['mingxi'] . '_' . $val['input_name']] = $val['odd'];
        }
        if ($is_year) {
            //本命年赔率
            $odd = isset($tmp_arr[$mingxi . '_' . $year_animal]) ? $tmp_arr[$mingxi . '_' . $year_animal] : '';
        } else {
            $odd = isset($tmp_arr[$mingxi . '_' . $other_animal]) ? $tmp_arr[$mingxi . '_' . $other_animal] : '';
        }
        if (empty($odd)) {
            $result['ErrorMsg'] = '获取生肖连赔率异常';
            echo json_encode($result);
            die;
        }

        return $odd;
    }

    /**
      ***********************************************************
      *  获取六合彩过关的赔率         @author ruizuo qiyongsheng    *
      ***********************************************************
    */
    private static function getPassTestOdds($odds_arr, $input_name, $mingxi){
        //检测下注合法性
        $result = array();
        $result['ErrorCode'] = 2;
        $tmp = true;
        if(empty($input_name) || empty($mingxi)) $tmp = false;
        $input_arr = explode(',', $input_name);
        $ball_arr = explode(',', $mingxi);
        if(count($input_arr) < 2) $tmp = false;
        if(count($input_arr) != count($ball_arr)) $tmp = false;
        if($tmp != true){
            $result['ErrorMsg'] = '下注内容错误---';
            echo json_encode($result);
            die;
        }
        if(count($input_arr) > 8){
            $result['ErrorMsg'] = '抱歉您最多只能选择8注';
            echo json_encode($result);
            die;
        }
        if(!$odds_arr){
            $result['ErrorMsg'] = '赔率异常~';
            echo json_encode($result);
            die;
        }
        $new_odds_arr = array();
        foreach($odds_arr as $val){
            if(!isset($val['mingxi']) || !isset($val['input_name'])){
                $result['ErrorMsg'] = '赔率异常~~';
                echo json_encode($result);
                die;
            }
            $new_odds_arr[$val['mingxi'] . $val['input_name']] = $val['odd'];
        }

        $ball = ['JustCode_one', 'JustCode_two', 'JustCode_three', 'JustCode_four', 'JustCode_five', 'JustCode_six'];
        $play = ['big', 'small', 'single', 'double', 'blue_wave', 'red_wave', 'green_wave'];
        $game_arr = array();
        foreach($input_arr as $key=>$val){
            $game_arr[] = $ball_arr[$key] . ':' . $val;
        }
        $new_odds = array();
        foreach($game_arr as $key=>$val){
            $tmp = array();
            $tmp = explode(':', $val);
            if(count($tmp) != 2){
                $result['ErrorMsg'] = '下注内容错误--';
                echo json_encode($result);
                die;
            }
            if(!in_array($tmp[0], $ball) || !in_array($tmp[1], $play)){
                $result['ErrorMsg'] = '下注内容错误-';
                echo json_encode($result);
                die;
            }
            if(!isset($new_odds_arr[$tmp[0] . $tmp[1]])){
                $result['ErrorMsg'] = '赔率异常~~~';
                echo json_encode($result);
                die;
            }
            $new_odds[] = $new_odds_arr[$tmp[0] . $tmp[1]];
        }
        //获取赔率（为所有赔率的总乘积）
        $odds =  array_product($new_odds);
        if($odds <= 0){
             if(!isset($new_odds_arr[$tmp[0] . $tmp[1]])){
                $result['ErrorMsg'] = '赔率异常~~~~';
                echo json_encode($result);
                die;
            }
        }
        return $odds;
    }
/**
      ***********************************************************
      *  获取六合彩尾数连赔率         @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    private static function getEndNumOdds($odds_arr, $input_name, $mingxi){
        $input_arr = explode(',', $input_name);
        $num = count($input_arr);
        $tmp_num = 0;
        if(in_array($mingxi, ['two_end_in', 'two_end_not_in'])) $tmp_num = 2;
        if(in_array($mingxi, ['three_end_in', 'three_end_not_in'])) $tmp_num = 3;
        if(in_array($mingxi, ['four_end_in', 'four_end_not_in'])) $tmp_num = 4;
        if($num != $tmp_num){
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
        $new_odds = array();
        foreach($odds_arr as $key=>$val){
            if(isset($val['input_name']) && isset($val['mingxi'])){
                $new_odds[$val['input_name'] . '_' .$val['mingxi']] = $val['odd'];
            }
        }
        //选中0尾后，如果是连中，选赔率高的，如果连不中，选赔率低的
        $is_exist = stripos($input_name, '0');
        if($is_exist || $is_exist === 0){
            $odds = isset($new_odds['zero_end_' . $mingxi]) ? $new_odds['zero_end_' . $mingxi] : '';
        }else{
            $odds = isset($new_odds['one_end_' . $mingxi]) ? $new_odds['one_end_' . $mingxi] : '';
        } 
        
        if($odds == ''){
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '赔率异常。。。';
            echo json_encode($result);
            die;
        }
        return $odds;
    }
    
    
/**
      ***********************************************************
      *  处理尾数 尾数连赔率数据       @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    private static function getEndNum($odds){
        foreach($odds as $key=>$val){
            switch($val['input_name']){
                case 'zero_end':
                    $odds[$key]['input_name'] = 0;
                break;
                 case 'one_end':
                    $odds[$key]['input_name'] = 1;
                break;
                 case 'two_end':
                    $odds[$key]['input_name'] = 2;
                break;
                 case 'three_end':
                    $odds[$key]['input_name'] = 3;
                break;
                 case 'four_end':
                    $odds[$key]['input_name'] = 4;
                break;
                 case 'five_end':
                    $odds[$key]['input_name'] = 5;
                break;
                 case 'six_end':
                    $odds[$key]['input_name'] = 6;
                break;
                 case 'seven_end':
                    $odds[$key]['input_name'] = 7;
                break;
                 case 'eight_end':
                    $odds[$key]['input_name'] = 8;
                break;
                case 'nine_end':
                    $odds[$key]['input_name'] = 9;
                break;
            }
        }
        return $odds;
    }

/**
      ***********************************************************
      *  获取彩种所有合法的input_name数组 @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    private static function checkAllInput( $type, $odds, $input_name, $gameplay = ''){
        //特殊玩法不在此处 单独验证
       if (in_array($type, ['bj_kl8', 'dm_klc', 'jnd_bs']) && in_array($gameplay, ['choose_one', 'choose_two', 'choose_three', 'choose_four', 'choose_five'])) return;
       if(in_array($type, ['gd_ten', 'cq_ten']) && in_array($gameplay, ['random_choose_two', 'random_choose_two_group', 'random_choose_three', 'random_choose_four', 'random_choose_five'])) return;
       if (in_array($type, ['sd_11', 'gd_11', 'jx_11']) && in_array($gameplay, ['random_choose', 'group_choose', 'vertical_choose'])) return;
       if(($type == 'liuhecai' || $type == 'jsliuhecai') && in_array($gameplay, ['PassTest', 'JointMark', 'SumAnimal', 'AnimalEven', 'EndNum_Even', 'AllMiss', 'EndNum'])) return;
       if($type == 'pc_28' && $input_name == 'Tema_in_Three') return;


        //正常玩法 
        $input_arr = array();
        foreach($odds as $key=>$val){
            $input_arr[] = $val['input_name'];
        }
        $input_arr = array_unique($input_arr);
        if(!in_array($input_name, $input_arr)){
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }

        return;
    }
    
    /**
     * **********************************************************
     *  检测快乐彩特殊玩法下注内容是否合法@author ruizuo qiyongsheng  *
     * **********************************************************
     */
    private static function check_klc_inputname($gameplay, $input_name) {
        $num = 0;
        switch ($gameplay) {
            case 'choose_one':
                $num = 1;
                break;
            case 'choose_two':
                $num = 2;
                break;
            case 'choose_three':
                $num = 3;
                break;
            case 'choose_four':
                $num = 4;
                break;
            case 'choose_five':
                $num = 5;
                break;
        }
        $tmp = explode(',', $input_name);
        foreach($tmp as $val){
            if(!is_numeric($val)){
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '下注内容错误';
                echo json_encode($result);
                die;
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }

    /**
     * **********************************************************
     *  检测快乐十分下注内容是否合法   @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_ten_inputname($gameplay, $input_name) {
        $num = 0;
        switch ($gameplay) {
            case 'random_choose_two':
                $num = 2;
                break;
            case 'random_choose_two_group':
                $num = 2;
                break;
            case 'random_choose_three':
                $num = 3;
                break;
            case 'random_choose_four':
                $num = 4;
                break;
            case 'random_choose_five':
                $num = 5;
                break;
        }
        $tmp = explode(',', $input_name);
        foreach($tmp as $val){
            if(!is_numeric($val)){
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '下注内容错误';
                echo json_encode($result);
                die;
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }

    /**
     * **********************************************************
     *  检测十一选五下注内容是否合法   @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_11_inputname($gameplay, $input_name, $mingxi) {
        if (empty($gameplay) || empty($input_name) || empty($mingxi)) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
        $num = 0;
        $res = true;
        $tmp = count(explode(',', $input_name));
        if ($gameplay == 'random_choose') {
            switch ($mingxi) {
                case 'one_in_one':
                    $num = 1;
                    break;
                case 'two_in_two':
                    $num = 2;
                    break;
                case 'three_in_three':
                    $num = 3;
                    break;
                case 'four_in_four':
                    $num = 4;
                    break;
                case 'five_in_five':
                    $num = 5;
                    break;
                case 'six_in_five':
                    $num = 6;
                    break;
                case 'seven_in_five':
                    $num = 7;
                    break;
                default:
                    $num = 'wrong';
                    break;
            }
            if ($tmp != $num)
                $res = false;
        }elseif ($gameplay == 'group_choose' || $gameplay == 'vertical_choose') {
            switch ($mingxi) {
                case 'before_three':
                    $num = 3;
                    break;
                case 'before_two':
                    $num = 2;
                    break;
                default:
                    $num = 'wrong';
                    break;
            }
            if ($tmp != $num)
                $res = false;
        }else {
            $tmp = false;
        }

        if ($tmp == false) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }
    /**
          ***********************************************************
          *  检测pc_28特码包三下注内容 @author ruizuo qiyongsheng    *
          ***********************************************************
    */
        private static function check_dd_inputname($mingxi){
            $tmp = explode(',', $mingxi);
            $tmp = array_unique($tmp);
             if (count($tmp) != 3) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '下注内容错误';
                echo json_encode($result);
                die;
            }
            foreach($tmp as $v){
                if($v < 0 || $v > 27){
                    $result['ErrorCode'] = 2;
                    $result['ErrorMsg'] = '下注内容错误';
                    echo json_encode($result);
                    die;
                }
            }
        }
        
    /**
     * **********************************************************
     *  检测六合彩合肖下注内容        @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_hexiao_inputname($input_name, $mingxi) {
        $num = 0;
        switch ($mingxi) {
            case 'two_Animal':
                $num = 2;
                break;
            case 'three_Animal':
                $num = 3;
                break;
            case 'four_Animal':
                $num = 4;
                break;
            case 'five_Animal':
                $num = 5;
                break;
            case 'six_Animal':
                $num = 6;
                break;
            case 'seven_Animal':
                $num = 7;
                break;
            case 'eight_Animal':
                $num = 8;
                break;
            case 'nine_Animal':
                $num = 9;
                break;
            case 'ten_Animal':
                $num = 10;
                break;
            case 'elven_Animal':
                $num = 11;
                break;
            default:
                $num = 'wrong';
                break;
        }
        $tmp = explode(',', $input_name);
        foreach ($tmp as $key => $val) {
            if (!in_array($val, self::animal_arr())) {
                unset($tmp[$key]);
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }

    /**
     * **********************************************************
     *  检测六合彩生肖连下注内容      @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_shengxiaolian_inputname($input_name, $mingxi) {
        $num = 0;
        switch ($mingxi) {
            case 'two_Animal_in':
            case 'two_Animal_not_in':
                $num = 2;
                break;
            case 'three_Animal_in':
            case 'three_Animal_not_in':
                $num = 3;
                break;
            case 'four_Animal_in':
            case 'four_Animal_not_in':
                $num = 4;
                break;
            case 'five_Animal_in':
                $num = 5;
                break;
            default:
                $num = 'wrong';
                break;
        }
        $tmp = explode(',', $input_name);
        foreach ($tmp as $key => $val) {
            if (!in_array($val, self::animal_arr())) {
                unset($tmp[$key]);
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }
    /**
      ***********************************************************
      *  检测六合彩尾数下注内容           @author ruizuo qiyongsheng    *
      ***********************************************************
    */
    private static function check_endnum_inputname($input_name){
        if(!is_numeric($input_name) || $input_name < 0 || $input_name > 9){
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }
        
        
    /**
     * **********************************************************
     *  检测六合彩尾数连下注内容      @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_weilian_inputname($input_name, $mingxi) {
        $num = 0;
        switch ($mingxi) {
            case 'two_end_in':
            case 'two_end_not_in':
                $num = 2;
                break;
            case 'three_end_in':
            case 'three_end_not_in':
                $num = 3;
                break;
            case 'four_end_in':
            case 'four_end_not_in':
                $num = 4;
                break;
            default:
                $num = 'wrong';
                break;
        }
        $tmp = explode(',', $input_name);
        foreach ($tmp as $key => $val) {
            if ($val < 0 || $val > 9) {
                unset($tmp[$key]);
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }

    /**
     * **********************************************************
     *  检测六合彩全不中下注内容      @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function check_allmiss_inputname($input_name, $mingxi) {
        $num = 0;
        switch ($mingxi) {
            case 'five_not_in':
                $num = 5;
                break;
            case 'six_not_in':
                $num = 6;
                break;
            case 'seven_not_in':
                $num = 7;
                break;
            case 'eight_not_in':
                $num = 8;
                break;
            case 'nine_not_in':
                $num = 9;
                break;
            case 'ten_not_in':
                $num = 10;
                break;
            case 'elven_not_in':
                $num = 11;
                break;
            case 'twelve_not_in':
                $num = 12;
                break;
            default:
                $num = 'wrong';
                break;
        }
        $tmp = explode(',', $input_name);
        foreach ($tmp as $key => $val) {
            if ($val < 1 || $val > 49) {
                unset($tmp[$key]);
            }
        }
        if (count($tmp) != $num) {
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '下注内容错误';
            echo json_encode($result);
            die;
        }
    }

    /**
     * **********************************************************
     *  十二生肖数组              @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function animal_arr() {
        return array('mouse', 'cattle', 'tiger', 'rabbit', 'dragon', 'snake', 'horse', 'sheep', 'monkey', 'chicken', 'dog', 'pig');
    }

    /**
     * **********************************************************
     *  获取六合彩生肖数组           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function get_animal($type = '') {
        $animal_arr = self::animal_arr();
        $year = date('Y');
        $year = $animal_arr[(($year - 4) % 12)]; //获取本命年生肖
        $shenxiao_arr = array('pig', 'dog', 'chicken', 'monkey', 'sheep', 'horse', 'snake', 'dragon', 'rabbit', 'tiger', 'cattle', 'mouse');

        $num_arr = array(
            array(1, 13, 25, 37, 49),
            array(2, 14, 26, 38),
            array(3, 15, 27, 39),
            array(4, 16, 28, 40),
            array(5, 17, 29, 41),
            array(6, 18, 30, 42),
            array(7, 19, 31, 43),
            array(8, 20, 32, 44),
            array(9, 21, 33, 45),
            array(10, 22, 34, 46),
            array(11, 23, 35, 47),
            array(12, 24, 36, 48)
        );


        $key = array_search($year, $shenxiao_arr);

        $begin_arr = array_splice($shenxiao_arr, $key);

        $end_arr = array_splice($shenxiao_arr, $key - 12);

        $new_shenxiao_arr = array_merge($begin_arr, $end_arr);

        $shenxiao_num_arr = array();
        foreach ($new_shenxiao_arr as $k => $v) {
            $shenxiao_num_arr[$v] = $num_arr[$k];
        }
        if ($type) { //加上色波
            $return = array();
            foreach ($shenxiao_num_arr as $key => $val) {
                foreach ($val as $k => $v) {
                    $return[$key][$k]['num'] = $v;
                    $return[$key][$k]['color'] = self::wave($v);
                }
            }
            return $return;
        }
        return $shenxiao_num_arr; //返回新的生肖数组
    }

    /**
     * **********************************************************
     *  六合彩波色                 @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function wave($num = '') {
        $ball = [];
        $ball['red'] = [1, 2, 7, 8, 12, 13, 18, 19, 23, 24, 29, 30, 34, 35, 40, 45, 46];
        $ball['blue'] = [3, 4, 9, 10, 14, 15, 20, 25, 26, 31, 36, 37, 41, 42, 47, 48];
        $ball['green'] = [5, 6, 11, 16, 17, 21, 22, 27, 28, 32, 33, 38, 39, 43, 44, 49];
        if (empty($num))
            return $ball;
        foreach ($ball as $wave => $val) {
            if (in_array($num, $val)) {
                return $wave;
            }
        }
    }

    /**
     * **********************************************************
     *  六合彩尾数连           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    private static function endNum() {
        $arr = $data = $return = array();
        for ($i = 1; $i <= 49; $i++) {
            $arr[] = $i;
        }
        foreach ($arr as $val) {
            $end_num = $val % 10;
            $tmp_arr = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];
            $data[$tmp_arr[$end_num] . '_end'][] = $val;
        }

        foreach ($data as $key => $val) {
            foreach ($val as $k => $v) {
                $return[$key][$k]['num'] = $v;
                $return[$key][$k]['color'] = self::wave($v);
            }
        }
        return $return;
    }

    /**
          ***********************************************************
          *  钱包模式事务未完成写入失败现金纪录@author ruizuo qiyongsheng    *
          ***********************************************************
    */
        
    private static function errorCashRecord($data){
        BetModel::insertErrorCashRecord($data);
    }

}
