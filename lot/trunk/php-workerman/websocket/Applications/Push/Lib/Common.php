<?php

namespace Applications\Push\Lib;

use \Applications\Common\Config\Config;
use \Applications\Common\Helper\Db;
use \Applications\Common\Helper\Redis;
use \Applications\Common\Helper\GetQishuOpentime;
use \workerman\Lib\Timer;
use \GatewayWorker\Lib\Gateway;

class Common {

    // 彩种 有变动时需重启
    public static $fc_type = [
        'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'mnl_o', 'dj_o', 'mg_o', 'xdl_10',
        'fc_3d', 'pl_3', 'liuhecai', 'cq_ssc', 'tj_ssc', 'xj_ssc', 'bj_10',
        'ah_k3', 'jl_k3', 'gx_k3', 'js_k3', 'gd_ten', 'cq_ten', 'gd_11', 'sd_11',
        'jx_11', 'bj_kl8', 'jnd_bs', 'dm_klc', 'bj_28', 'pc_28', 'jnd_28', 'dm_28'
    ];

    // 初始化
    public static function init($worker) {
        date_default_timezone_set('PRC');
    }

    // 用户分组
    public static function join($client_id, $message_data) {
        $group_list = $message_data['group_list'];

        if ($group_list[0] == 'online') {
            //后台展示前台在线会员
            Gateway::bindUid($client_id, $message_data['uid']);
            return;
        }

        $fc_type = isset($message_data['fc_type']) ? $message_data['fc_type'] : '';
        if (!in_array($fc_type, self::$fc_type)) {
            Gateway::closeClient($client_id);
        }
        foreach ($group_list as $type => $group) {
            Gateway::joinGroup($client_id, $group . $fc_type);
        }
    }

    // 封盘时间推送
    public static function lefttime($worker) {
        self::lefttimeCode();
        $timer_id = Timer::add(30, function() use (&$timer_id) {
                    self::lefttimeCode();
                });
    }

    public static function lefttimeCode() {
        $msg['cmd'] = 'lefttime';
        $timestamp = time();
        $games = self::getAllFcTypes();
        foreach ($games as $k => $v) {
            $fc_type = $v['type'];
            $periods = GetQishuOpentime::get_qishu($fc_type);

            $game_time = GetQishuOpentime::getGametime($fc_type, $periods);

            if ( strtotime($game_time['fengpan']) <= $timestamp && $timestamp <= strtotime($game_time['kaijiang']) ) {
                $periods = GetQishuOpentime::new_qishu($fc_type, $periods);
                $game_time = GetQishuOpentime::getGametime($fc_type, $periods);
            } elseif ((strtotime($game_time['fengpan']) < $timestamp) && in_array($fc_type, ['bj_kl8', 'bj_10', 'bj_28', 'pc_28', 'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'mnl_o', 'dj_o', 'mg_o', 'xdl_10'])) {
                $periods = GetQishuOpentime::new_qishu($fc_type, $periods);
                $game_time = GetQishuOpentime::getGametime($fc_type, $periods);
            }

            $data['fc_type'] = $fc_type;
            $data['qishu'] = $periods;
            $data['now_time'] = $timestamp;
            $data['open_time'] = strtotime($game_time['kaipan']);
            $data['close_time'] = strtotime($game_time['fengpan']);

            $msg['data'] = $data;
            $cond1 = $data['close_time'] - $data['now_time'] > 60;
            $cond2 = $fc_type != 'ffc_o';
            if ($cond1 && $cond2) {
                Gateway::sendToGroup('lefttime' . $fc_type, json_encode($msg));
            }
        }
    }

    public static function getAllFcTypes() {
        $redis = Redis::instance();

        $redis_key = 'c_lot_all_game_site';
        $data = $redis->get($redis_key);
        $data = json_decode($data, true);

        if (empty($data)) {
            $data = Db::instance('manage')->select('*')->from('my_fc_games')->where('state IN (1,3)')->query();
            $redis->set($redis_key, json_encode($data));
        }
        return $data;
    }

    // 开奖结果推送
    public static function auto($client_id, $message_data) {
        $msg['cmd'] = 'auto';
        $msg['data'] = [];
        $fc_type = isset($message_data['fc_type']) ? $message_data['fc_type'] : '';
        if ($fc_type) {
            $data['fc_type'] = $fc_type;
            $data['qishu'] = isset($message_data['qishu']) ? $message_data['qishu'] : '';
            $data['datetime'] = isset($message_data['datetime']) ? $message_data['datetime'] : '';
            $data['ball'] = isset($message_data['ball']) ? $message_data['ball'] : '';

            $msg['data'] = $data;
            Gateway::sendToGroup('auto' . $fc_type, json_encode($msg));
        }
        Gateway::sendToCurrentClient(json_encode($msg));
    }

    //在线会员推送
    public static function online($client_id, $message_data) {
        $data = isset($message_data['data']) ? $message_data['data'] : '';
        if (empty($data)) {
            return;
        }
        $data = json_decode($data,true);

        foreach ($data as $k => $v) {
            //给每一个用户发送
            $send = [
                'cmd' => 'online',
                'count' => $v
            ];
            Gateway::sendToUid($k, json_encode($send));
        }

        //发送给当前客户端
        Gateway::sendToCurrentClient(json_encode($data));
    }

}
