<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;
use \helper\SpiderCommon;
use \helper\RedisConPool;

class Bj_28 {

    static $qishu = "";
    static $wait = [];
    static $code = "bjkl8";
    static $name = 'bjklb';
    static $check_arr = array(//检测开奖结果合法性
        'type' => 'bj_28',
        'maxexpect' => 7890600, //期数最大值
        'minexpect' => 695394, //期数最小值
        'maxball' => 80, //单球最大值
        'minball' => 1, //单球最小值
        'ballcount' => 20    //球的个数
    );

    public static function getData() {
        // 第一区 = 取开奖号码的第2/5/8/11/14/17位置的6个数，相加的总和，总和的尾数为一区数字
        // 第二区 = 取开奖号码的第3/6/9/12/15/18位置的6个数，相加的总和，总和的尾数为二区数字
        // 第三区 = 取开奖号码的第4/7/10/13/16/19位置的6个数，相加的总和，总和的尾数为三区数字
        //采集接口一
        $redis = RedisConPool::getInstace();
        $redis_key = 'auto_for28_bj_kl8';
        $data = $redis->get($redis_key);
        if(!$data) return array();
        $data = json_decode($data,true);

        if(empty($data) || !isset($data[0])) return array();
       
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
                    $new_arr[$key]['opencode'] = SpiderCommon::regroup_auto_interval($v, 3);
                } else {
                    $new_arr[$key][$k] = $v;
                }
            }
        }
        return $new_arr;
    }

}
