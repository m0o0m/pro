<?php

use \Workerman\Worker;
use \Workerman\WebServer;
use \GatewayWorker\Gateway;
use \GatewayWorker\BusinessWorker;
use \Workerman\Autoloader;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$gateway = new Gateway("http://0.0.0.0:9528");

$gateway->name = 'PushHttp';

$gateway->count = 4;

$gateway->registerAddress = '127.0.0.1:1258';

if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
