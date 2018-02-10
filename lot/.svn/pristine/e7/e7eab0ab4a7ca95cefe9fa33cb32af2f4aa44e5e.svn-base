<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;
use \helper\SpiderCommon;
use \helper\RedisConPool;

class Dm_28 {

    static $qishu = "";
    static $wait = [];
    static $type1 = 'get_lotto.php';
    static $type2 = 'tidligere-danske28-resultater/';

    public static function getData() {
        $redis = RedisConPool::getInstace();
        $redis_key = 'auto_for28_dm_klc';
        $data = $redis->get($redis_key);
        if(!$data) return array();
        $data = json_decode($data,true);

        if(empty($data) || !isset($data[0])) return array();
        $auto = $data[0];
        if(!isset($auto['opencode'])) return array();
        
        $auto['opencode'] = SpiderCommon::regroup_auto_continuous($auto['opencode'], 3);
        $arr = array();
        $arr[] = $auto;
        return $arr;
    }

    public static function getContinueData($time) {
        self::getData();
    }

}
