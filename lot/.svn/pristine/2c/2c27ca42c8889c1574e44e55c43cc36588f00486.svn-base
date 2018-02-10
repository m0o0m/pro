<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;
use \helper\SpiderCommon;
use \helper\RedisConPool;

class Jnd_28 {

    static $qishu = "";
    static $wait = [];
    static $code = "cakeno";
    static $name = 'jndklb';
    static $check_arr = array(//检测开奖结果合法性
        'type' => 'jnd_28',
        'maxexpect' => 9000000, //期数最大值
        'minexpect' => 2000000, //期数最小值
        'maxball' => 80, //单球最大值
        'minball' => 1, //单球最小值
        'ballcount' => 20    //球的个数
    );

    public static function getData() {
        $redis = RedisConPool::getInstace();
        $redis_key = 'auto_for28_jnd_bs';
        $data = $redis->get($redis_key);
        if(!$data) return array();
        $data = json_decode($data,true);

        if (empty($data) || !isset($data[0])) return array();
        $data[0]['opencode'] = SpiderCommon::regroup_auto_interval($data[0]['opencode'], 3);

        return $data;
    }

    public static function getContinueData($time) {
        $url = str_replace('[continue]', curl_config::$continue, curl_config::$continue_url);
        $url = str_replace('[token]', curl_config::$token, $url);
        $url = str_replace('[code]', static::$code, $url);
        $url .= $time;

        $temp_time = strtotime($time);
        $time2 = date('Ymd', $temp_time); //转换成接口需要的时间格式
        //彩票控采集链接
        $cpk_url = str_replace('[caipiaokong]', curl_config::$caipiaokong, curl_config::$caipiaokong_url);
        $cpk_url = str_replace('[caipiaokong_token]', curl_config::$caipiaokong_token, $cpk_url);
        $cpk_url = str_replace('[code]', static::$name, $cpk_url);
        $cpk_url = str_replace('[caipiaokong_uid]', curl_config::$caipiaokong_uid, $cpk_url);
        $cpk_url = str_replace(['[caipiaokong_num]'], 400, $cpk_url);
        $cpk_url .= $time2;

        $data = SpiderCommon::spiderData(static::$check_arr, $url, $cpk_url, 2);
        if (empty($data)) {
            return $data;
        }
        $new_arr = array();
        foreach ($data as $key => $val) {
            foreach ($val as $k => $v) {
                if ($k == 'opencode') {
                    $new_arr[$key]['opencode'] = SpiderCommon::regroup_auto_interval('jnd_28',$v, 3);
                } else {
                    $new_arr[$key][$k] = $v;
                }
            }
        }
        return $new_arr;
    }

}
