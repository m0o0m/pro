<?php

use \Workerman\Worker;
use \Workerman\Lib\Timer;
use Applications\Tools\Lib\Common;
use \helper\RedisConPool as Redis;
use Applications\Tools\Lib\Games;
use \helper\MysqlPdo as pdo;
use \helper\Encrypt;
use helper\TcpConPoll;
use \helper\IdWork;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("text://0.0.0.0:9999");

$worker->name = 'test';

$worker->count = 100;

$worker->onWorkerStart = function($worker) {
    
};

$worker->onMessage = function($connection, $data) {
   
    if (empty($data)) {
        sleep(1);
    }
    
    //$redis= Redis::getInstace();
    
    
    //$redis->lpush('test_sql',$data);
    
    
    $sql="insert into `manage`.`my_user_cash_record` ( `uid`, `line_id`, `agent_id`, `cash_type`, `cash_do_type`, `dids`, `cash_num`, `cash_balance`, `remark`, `ptype`, `addtime`, `addday`, `uname`, `fc_type`, `periods`, `is_shiwan`) values ( '152', 'aab', '7', '2', '1', '17112582325370614', '1.00', '5464.60', 'Lottery note17112582325370614,type:pl_3,共计:1单', '1', '1511582325', '20171125', 'test', 'pl_3', '17322', '1')";
    $db= pdo::instance('manage');
    
    $db->query($sql);



    $connection->send('success');
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
