<?php

use \Workerman\Worker;
use \helper\MysqlPdo as pdo;
use \helper\Common_helper;
use \Workerman\Lib\Timer;
use \helper\RedisConPool;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

Worker::$stdoutFile = __DIR__ . '/../../logs/worker.log';

$worker = new Worker('http://0.0.0.0:3341');
$worker->name = 'jnd_bs_opentime';
$worker->onWorkerStart = function($worker) {
    //冬令时，官网时间跟北京时间相差的是16个小时
    //夏令时，官网时间跟北京时间相差的是15个小时
    $hour = 16; //根据情况手工更改成16

    $class = 'libraries\spider\\jnd_bs_opentime';
    $res = $class::getData('my_jnd_bs_opentime', $hour);
    date_default_timezone_set("Asia/Shanghai");

    Timer::add(120, function()use($class, $hour) {
        //判断当前时间，如果在19:07至第二天12:00之间运行采集
        $today_start = strtotime(date('Y-m-d', time()) . '19:07:00');
        $today_end = strtotime(date('Y-m-d', time()) . '12:00:00');
        //当前时间小于今天的12点 或者 大于今天的19点。
        if ((time() < $today_end ) || (time() > $today_start)) {
            $class::getData('my_jnd_bs_opentime', $hour);
        }
    });
};



// 运行worker
if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
?>
