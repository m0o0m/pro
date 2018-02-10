<?php

use \Workerman\Worker;
use \helper\Curl;
use \Workerman\Lib\Timer;
use \helper\MysqlPdo as pdo;
use \helper\RedisConPool;
use \helper\Common_helper;
use \config\config;

require_once __DIR__ . '/../../Workerman/Autoloader.php';

Worker::$stdoutFile = __DIR__ . '/../../logs/worker.log';

$worker = new Worker('http://0.0.0.0:3338');
$worker->name = 'klcspider';
$worker->onWorkerStart = function($worker) {
    ///初始化载入所有彩种最新期数
    $typelist = [
        'bj_kl8',
        'dm_klc',
        'jnd_bs'
    ];

    foreach ($typelist as $type) {
        $class = 'libraries\spider\\' . $type;
        $class::$qishu = Common_helper::GetNewNumber($type); //根据type 获取最后一期期数
    }

    Timer::add(10, function()use($typelist) {
        foreach ($typelist as $type) {
            $class = 'libraries\spider\\' . $type;
            $data = $class::getData();                   ////十秒采集一波数据
            if (empty($data) || (!isset($data[0]['expect']))) { 
                return;//如果未采集到数据
            }

            //获取对应的结算进程地址
            $port_arr = Common_helper::getPort($type);
            $port = $port_arr[1];//8692
            $ip = 'http://' . config::$balance . ':' . $port;

            if ($class::$qishu < $data[0]['expect']) { //如果期数小于采集的期数
                $auto_class = 'libraries\auto\\' . ucfirst($type) . 'Auto';
                $new_auto = $auto_class::get_auto(explode(',', $data[0]['opencode']));
                $new_arr = array();
                $new_arr[0] = $data[0];
                $new_arr[1] = $new_auto;
                $res = Common_helper::addNewNumber($new_arr, $type, 20);

                $qishu = $data[0]['expect']; //获得最新的期数
                if ($res) {
                    //将开奖结果存入redis供幸运彩使用
                    $redis = RedisConPool::getInstace();
                    $redis_key = 'auto_for28_' . $type;
                    if($type == 'bj_kl8'){
                        $redis->set($redis_key . '_2', json_encode($data));
                    }
                    $redis->set($redis_key, json_encode($data));
                    //触发结算
                    $res = Curl::run($ip, 'post', array(
                                'todo' => 'balance',
                                'type' => $type,
                                'qishu' => $qishu
                    ));


                    if ($res != 'start') {
                        if (strlen($res) == 25) {
                            $class::$qishu = $qishu;
                            // echo 'success';
                        }
                        // echo 'The data maybe exist,or the database insert wrong for expect='. $qishu .' type=' . $type . PHP_EOL;
                    } else {
                        $class::$qishu = $qishu;
                        // echo 'success';
                    }
                }
            }

            ////复核待采队列
            foreach ($class::$wait as $qishu => $time) {
                $data = $class::getContinueData($time);
                foreach ($data as $val) {
                    if ($val['expect'] == $qishu) {
                        $res = Common_helper::addNewNumber($val, $type, 20);
                        if (!$res) {
                            $class::$wait[$qishu] = $time;                  ////采集失败进入复采队列
                        } else {
                            $res = Curl::run($ip, 'post', array(
                                        'todo' => 'balance',
                                        'type' => $type,
                                        'qishu' => $qishu
                            ));

                            if ($res != 'start') {
                                $class::$wait[$qishu] = $time;              ////启动结算失败进入复采队列
                            } else {
                                unset($class::$wait[$qishu]);
                            }
                        }
                        return;
                    }
                }
            }
        }
    });
};

$worker->onMessage = function($connection, $data) {
    ////补采集措施
    $type = $_POST['type'];
    $time = $_POST['time'];
    $qishu = $_POST['qishu'];
    $class = 'libraries\spider\\' . $type;
    $class::$wait[$qishu] = $time;                  ////进入复采队列
    $connection->send('get');
    return;
};


// 运行worker
if(!defined('GLOBAL_START')) {
    Worker::runAll();
}
?>


