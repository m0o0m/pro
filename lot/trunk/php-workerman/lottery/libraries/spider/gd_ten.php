<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;
use \helper\SpiderCommon as common;

class Gd_ten {

    static $qishu = "";
    static $wait = [];
    static $code = "gdklsf";
    static $name = 'gdklsf';
    static $check_arr = array(//检测开奖结果合法性
        'type' => 'gd_ten',
        'maxexpect' => 2050010100, //期数最大值
        'minexpect' => 2016010101, //期数最小值
        'maxball' => 20, //单球最大值
        'minball' => 1, //单球最小值
        'ballcount' => 8    //球的个数
    );

    public static function getData() {

        //采集接口一
        $url = str_replace('[domin]', curl_config::$domin, curl_config::$url);
        $url = str_replace('[token]', curl_config::$token, $url);
        $url = str_replace('[code]', static::$code, $url);

        //彩票控采集链接
        $cpk_url = str_replace('[caipiaokong]', curl_config::$caipiaokong, curl_config::$caipiaokong_url);
        $cpk_url = str_replace('[caipiaokong_token]', curl_config::$caipiaokong_token, $cpk_url);
        $cpk_url = str_replace('[code]', static::$name, $cpk_url);
        $cpk_url = str_replace('[caipiaokong_uid]', curl_config::$caipiaokong_uid, $cpk_url);
        $cpk_url = str_replace(['[caipiaokong_num]'], curl_config::$caipiaokong_num, $cpk_url);

        $data = common::spiderData(static::$check_arr, $url, $cpk_url);
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

        $data = common::spiderData(static::$check_arr, $url, $cpk_url, 2);
        return $data;
    }

}
