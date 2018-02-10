<?php

ini_set("memory_limit", "-1");

use \Workerman\Worker;
use \Workerman\Connection\AsyncTcpConnection;
use \Workerman\Lib\Timer;
use \helper\RedisConPool as Redis;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("http://0.0.0.0:6666");

$worker->name = 'user_cash_http';

$worker->count = 1;

$worker->onWorkerStart = function($worker) {

    $Redis = Redis::getInstace();
    for ($i = 1; $i <= 16; $i++) {
        $data = $Redis->lpop("user_cash_sql");

        $$i = new AsyncTcpConnection('text://127.0.0.1:7777');

        $$i->onConnect = function($Async) use ($data) {
            $Async->send($data);
        };

        $$i->onMessage = function($Async, $response) {
            $Redis = Redis::getInstace();
            $data = $Redis->lpop("user_cash_sql");

            $Async->send($data);
        };

        $$i->connect();
    }
};

$worker->onMessage = function($connection, $header) {
    
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
