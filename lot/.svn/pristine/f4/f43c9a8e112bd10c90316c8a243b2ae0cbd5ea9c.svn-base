<?php

use \Workerman\Worker;
use \Workerman\Lib\Timer;
use \helper\RedisConPool as Redis;
use \helper\MysqlPdo as pdo;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("text://0.0.0.0:7777");

$worker->name = 'user_cash_text';

$worker->count = 16;

$worker->onWorkerStart = function($worker) {
    
};

$worker->onMessage = function($connection, $data) {

    if (empty($data)) {
        sleep(1);
    }

    $db = pdo::instance('lottery');

    try {
        $db->query($data);
    } catch (\Exception $e) {
        $redis = Redis::getInstace();
        $redis->lpush('user_cash_sql', $data);
    }

    $connection->send('success');
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
