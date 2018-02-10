<?php

use \Workerman\Worker;
use \Workerman\Lib\Timer;
use \helper\RedisConPool as Redis;
use \helper\Curl;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("http://0.0.0.0:2349");

$worker->name = 'forvedOffLogin';

$worker->count = 1;

$worker->onWorkerStart = function($worker) {

    Timer::add(60, function() {
        $redis = Redis::getInstace();

        $info = $redis->hgetall('userOnLine_front');
        if (empty($info)) {
            return;
        }
        $all_token = array_keys($info); //所有token
        $at_arr = $line_arr = array();
        $total = 0;
        foreach ($info as $token => $val) {
            $tmp = json_decode($val, true);
            $tmp_time = isset($tmp['time']) ? $tmp['time'] : null;
            $uid = isset($tmp['uid']) ? $tmp['uid'] : null;
            $line_id = isset($tmp['line_id']) ? $tmp['line_id'] : null;
            $agent_id = isset($tmp['agent_id']) ? $tmp['agent_id'] : null;

            //指定时间后自动踢线
            if (empty($tmp_time) || (time() - $tmp_time) > 1200) {
                if ($uid) {
                    $redis->hdel($line_id . '_uidToken', $uid);
                    $redis->hdel($agent_id . '_uidToken', $uid);
                }
                $redis->hdel('userOnLine_front', $token);
            }
            //统计所有在线会员
            $total += 1;
            //统计线路在线会员
            if ($line_id) {
                $line_arr[$line_id]['line_id'] = $line_id;
                if (isset($line_arr[$line_id]['sum'])) {
                    $line_arr[$line_id]['sum'] += 1;
                } else {
                    $line_arr[$line_id]['sum'] = 1;
                }
            }
            //统计代理在线会员
            if ($agent_id) {
                $tmp_key = $line_id . '_' . $agent_id;
                $at_arr[$tmp_key]['at_id'] = $agent_id;
                if (isset($at_arr[$tmp_key]['sum'])) {
                    $at_arr[$tmp_key]['sum'] += 1;
                } else {
                    $at_arr[$tmp_key]['sum'] = 1;
                }
            }
        }

        $curl_data = array();
        if (!empty($line_arr)) {
            foreach ($line_arr as $line_id => $val) {
                $line_key = md5($line_id);
                $curl_data[$line_key] = $val['sum'];
            }
        }

        if (!empty($at_arr)) {
            foreach ($at_arr as $unikey => $val) {
                $at_key = md5($unikey);
                $curl_data[$at_key] = $val['sum'];
            }
        }

        $curl_data[md5('pkadmin')] = $total;

        Curl::run('http://127.0.0.1:9528', 'post', array(
            'cmd' => 'online',
            'data' => json_encode($curl_data)
        ));
    });
};

$worker->onMessage = function($connection, $header) {
    
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
