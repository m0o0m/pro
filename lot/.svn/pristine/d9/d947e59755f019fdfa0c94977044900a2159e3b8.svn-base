<?php

ini_set("memory_limit", "-1");

use \Workerman\Worker;
use \Workerman\Connection\AsyncTcpConnection;
use \helper\MysqlPdo as pdo;
use \helper\RedisConPool;
use \helper\Curl;
use \config\config;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

Worker::$stdoutFile = __DIR__ . '/../../logs/worker.log';

$worker = new Worker('http://0.0.0.0:10004');
$worker->name = 'betApiForAuto';

$worker->onWorkerStart = function($worker) {
    // echo 'common start now' . PHP_EOL;
};

$worker->onMessage = function($connection, $data) {
    $post = $data['post'];
    //获取所有彩种玩法结果
    if ($post['todo'] == 'get_res') {
        $balls = isset($post['balls']) ? $post['balls'] : false;
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : false;
        if (!$balls || !$fc_type) {
            $result['ErrorMsg'] = '进程：参数错误';
            $connection->send(json_encode($result));
            return;
        }
        $class = 'libraries\auto\\' . ucfirst($fc_type) . 'Auto';
        $result['ErrorCode'] = 2;
        $auto_res = array();
        $auto_res = $class::get_auto($balls);
        if (empty($auto_res)) {
            $result['ErrorMsg'] = '获取玩法数组失败';
            $connection->send(json_encode($result));
            return;
        }
        $connection->send(json_encode($auto_res));
        return;
    }

    //补彩开奖结果
    if ($post['todo'] == 'getauto') {
        $manage = pdo::instance('manage');
        $result['ErrorCode'] = 2;
        //获取所有数据
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : null;
        $time = isset($post['time']) ? $post['time'] : null;
        if (empty($time) || empty($fc_type)) {
            $result['ErrorMsg'] = '进程：参数错误';
            $connection->send(json_encode($result));
            return false;
        } else {
            $time = date('Ymd', $time); //转换成接口需要的时间格式
        }
        $spideclass = 'libraries\spider\\' . $fc_type;
        $data = $spideclass::getContinueData($time);
        if (empty($data)) {
            $result['ErrorMsg'] = '暂未取得数据，请稍候重试';
            $connection->send(json_encode($result));
            return false;
        }
        //查询是否在数据库中存在
        $count = count($data) - 1;
        if ((!isset($data[$count]['expect'])) || (!isset($data[0]['expect']))) {
            $result['ErrorMsg'] = '暂未取得有效数据，请稍候重试';
            $connection->send(json_encode($result));
            return false;
        }
        $small_qishu = $data[$count]['expect'];
        $big_qishu = $data[0]['expect'];
        $auto_table = config::$tablePrefix . 'auto_' . $fc_type;
        $qishu_arr = array();
        $sql = 'select qishu from ' . $auto_table . ' where qishu between ' . $small_qishu . ' and ' . $big_qishu;
        $qishu_arr = $manage->query($sql);
        if (!empty($qishu_arr)) {
            foreach ($qishu_arr as $key => $val) {
                $qishu_arr[$key] = $val['qishu'];
            }
        }

        $qishu_list = array();

        $return = array();
        foreach ($data as $key => $val) {
            if ($qishu_arr) {//检测数据库中该期结果是否存在
                if (in_array($val['expect'], $qishu_arr)) {
                    continue;
                }
            }
            $qishu_list[] = $val['expect']; //用于处理结算
            $open_time = strtotime($val["opentime"]);
            $open_res = explode(',', $val["opencode"]);
            $class = 'libraries\auto\\' . ucfirst($fc_type) . 'Auto';
            $auto_res = array();
            $auto_res = $class::get_auto($open_res);
            if (empty($auto_res)) {
                $result['ErrorMsg'] = '获取玩法数组失败';
                $connection->send(json_encode($result));
                return;
            }
            $auto_res['qishu'] = $val['expect'];
            $auto_res['datetime'] = $open_time;
            $auto_res['addtime'] = time();
            $return[] = $auto_res;
        }
        //返回数据
        if (!empty($return)) {
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = '补采数据成功';
            $result['qishu_list'] = $qishu_list;
            $result['balls'] = $return;
            $connection->send(json_encode($result));
            return true;
        } else {
            $result['ErrorMsg'] = '该彩种当天没有遗漏数据';
            $connection->send(json_encode($result));
            return false;
        }
        return false;
    }


    $result['ErrorMsg'] = '非法指令，无法处理';
    $result['ErrorCode'] = 2;
    $connection->send(json_encode($result));
    return false;
};



// 运行worker
if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
?>