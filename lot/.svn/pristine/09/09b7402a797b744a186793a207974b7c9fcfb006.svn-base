<?php

ini_set("memory_limit", "-1");

use \Workerman\Worker;
use \Workerman\Connection\AsyncTcpConnection;
use \Workerman\Lib\Timer;
use Applications\Tools\Lib\Common;
use \helper\RedisConPool as Redis;
use Applications\Tools\Lib\Games;
use \helper\MysqlPdo as pdo;
use \helper\Encrypt;
use helper\TcpConPoll;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

$worker = new Worker("http://0.0.0.0:8888");

$worker->name = 'test';

$worker->count = 1;

$worker->onWorkerStart = function($worker) {

    //$redis = Redis::getInstace();
    //$data = $redis->rpop('test_sql');
    //var_dump($data);
    //return;

    /* $redis=Redis::getInstace();
      $s= time();

      $data=$redis->hgetall('test_balance');

      $e=time();

      echo $e-$s;
      return;
     */
    /* $lottery = pdo::instance('lottery');
      $sql = 'select * from my_bet_record limit 1';
      $data = $lottery->query($sql);
      var_dump($data[0]); */

    //$data="INSERT INTO `my_user_cash_record`(`uid` ,`line_id` ,`agent_id` ,`cash_type` ,`cash_do_type` ,`dids` ,`cash_num` ,`cash_balance` ,`remark` ,`ptype` ,`addtime` ,`addday` ,`uname` ,`fc_type` ,`periods` ,`is_shiwan`)VALUES('152','aab','7','2','1','17112582325370614','1.00','5464.60','Lottery note17112582325370614,type:pl_3,共计:1单' ,'1','1511582325','20171125','test','pl_3','17322','1')";



    $data = 1;
    for ($i = 1; $i <= 100; $i++) {
        //$redis = Redis::getInstace();
        //$data = $redis->rpop('test_sql');
        $$i = new AsyncTcpConnection('text://127.0.0.1:9999');
        $$i->onConnect = function($Async)use($data) {
            $res = $Async->send($data);
        };
        $$i->onMessage = function($Async, $responed)use($data) {
            if ($data) {
                $Async->send(json_encode($data));
            } else {
                $Async->close();
            }
        };
        $$i->connect();
    }


    /*
      $Redis = Redis::getInstace();
      for ($i = 1; $i <= 16; $i++) {
      $data = $Redis->lpop("user_cash_sql");

      $$i = new AsyncTcpConnection('text://127.0.0.1:10101');

      $$i->onConnect = function($Async) use ($data) {
      $Async->send($data);
      };

      $$i->onMessage = function($Async, $response) {
      $Redis = Redis::getInstace();
      $data = $Redis->lpop("user_cash_sql");

      $Async->send($data);
      };

      $$i->connect();
      } */






    return;

    //tcp钱包测试
    $str = Encrypt::GetBalance('aaa_jiutong2', "CNY");
    //$strr = Encrypt::Transfer("2017092006228763~201709200622879","aaa_jiutong2","CNY",-20,"重庆时时彩");
    //$st = Encrypt::checkTransfer('1');
    $socket = TcpConPoll::getInstace();
    var_dump($socket::send($str));
    /*
      $games = new Games();
      //$games->login();
      $data = $games->spider();
      $data = json_decode($data, true)['data'];
      if (!empty($data)) {
      //回传 删key
      $keys = [];
      foreach ($data as $k => $v) {
      $keys[] = $v['Score'];
      }

      if (!empty($keys)) {
      $keys_str = implode('|', $keys);
      $res = $games->delkeys($keys_str);
      print_r($res);
      }
      } */
};

$worker->onMessage = function($connection, $header) {
    
};

if (!defined('GLOBAL_START')) {
    Worker::runAll();
}
