<?php

use \Workerman\Worker;
use \Workerman\Lib\Timer;
use Applications\Tools\Lib\Common;
use \helper\RedisConPool as Redis;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("http://0.0.0.0:2350");

$worker->name = 'spider_data';

$worker->count = 1;

$worker->onWorkerStart = function($worker) {
    $redis = Redis::getInstace();
    $redis_key = 'spiderBetForApi_list';

    Timer::add(10, function()use($redis, $redis_key) {
        $action_key = 'waitSpiderAction'; //正在处理中
        if ($redis->exists($action_key))
            return;
        $data = $redis->rpop($redis_key);
        if ($data) {
            Common::actionData($data);
        }
    });
};

$worker->onMessage = function($connection, $header) {
    if (!isset($_POST['todo']) || $_POST['todo'] != 'next')
        return;
    $redis = Redis::getInstace();
    $action_key = 'waitSpiderAction'; //正在处理中
    $data = $redis->rpop($redis_key);
    if ($data) {
        $redis->set($action_key, 1);
        Common::actionData($data);
    }
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
