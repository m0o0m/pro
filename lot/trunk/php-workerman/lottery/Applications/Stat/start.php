<?php

use \Workerman\Worker;
use \helper\MysqlPdo as Db;
use \Applications\Stat\Lib\Common;
use \Workerman\Lib\Timer;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("http://0.0.0.0:2346");

$worker->name = 'Stat';

$worker->count = 1;

$worker->onWorkerStart = function($worker) {
    Timer::add(30, function() {
        $num = 10000; //达到此数值时从内存中读取并处理操作数据库
        Common::getDataFromRedis($num);
    });
};

$worker->onMessage = function($connection, $header) {
    Common::reStat($connection, $header);
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
