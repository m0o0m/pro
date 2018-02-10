<?php

use \Workerman\Worker;
use \Workerman\WebServer;
use \GatewayWorker\Gateway;
use \GatewayWorker\BusinessWorker;
use \Workerman\Autoloader;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new BusinessWorker();

$worker->name = 'PushBusinessWorker';

$worker->count = 1;

$worker->registerAddress = '0.0.0.0:1258';

if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
