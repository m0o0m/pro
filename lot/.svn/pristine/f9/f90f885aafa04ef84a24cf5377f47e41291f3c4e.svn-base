<?php

namespace libraries\spider;

use \helper\GetQishuOpentime;
use \config\curl_config;
use \helper\SpiderCommon;

class Els_o {

    static $qishu = "";
    static $wait = [];
    static $name = 'ru15s';
    static $check_arr = array(//检测开奖结果合法性
        'type' => 'els_o',
        'maxexpect' => 20501212001, //期数最大值
        'minexpect' => 20180101001, //期数最小值
        'maxball' => 9, //单球最大值
        'minball' => 0, //单球最小值
        'ballcount' => 5    //球的个数
    );

    public static function getData() {
        //开奖网
        $url = ''; //未购买
        //彩票控采集链接
        $cpk_url = str_replace('[caipiaokong]', curl_config::$caipiaokong, curl_config::$caipiaokong_url);
        $cpk_url = str_replace('[caipiaokong_token]', curl_config::$caipiaokong_token, $cpk_url);
        $cpk_url = str_replace('[code]', static::$name, $cpk_url);
        $cpk_url = str_replace('[caipiaokong_uid]', curl_config::$caipiaokong_uid, $cpk_url);
        $cpk_url = str_replace(['[caipiaokong_num]'], curl_config::$caipiaokong_num, $cpk_url);

        $data = SpiderCommon::spiderData(static::$check_arr, $url, $cpk_url);

        return $data;
    }

    public static function getContinueData($time) {
        $url = '';

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
            $new_arr[$key] = $val;
            $new_arr[$key]['opencode'] = SpiderCommon::regroup_auto_continuous($val['opencode'], 5);
        }
        return $new_arr;
    }
}
