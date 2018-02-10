<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;

class Fc_3d {

    static $qishu = "";
    static $wait = [];
    static $type = "3d";

    // static $url = "http://kj.zhcw.com/kjData/2012/zhcw_3d_index_last30.js";
    // static $continue_url = "http://caipiao.163.com/order/3d/";

    public static function getData() {
        $url = curl_config::$zhcw;
        $continue_url = str_replace('[caipiao163]', curl_config::$caipiao163, curl_config::$caipiao163_url);
        $continue_url = str_replace('[caipiao163_type]', static::$type, $continue_url);
        //采集一
        $str = Curl::run($url, 'get');
        //如果采集不到数据
        if (empty($str)) {
            // echo 'Can not collection the url:' . $url . PHP_EOL;
            return array();
        }
        $data_preg = '/eval\(\'\(\'\s[+]\s\'(.*)\'\s[+]\s\'\\)\'\);/';
        preg_match_all($data_preg, $str, $data);
        $data = $data[1];
        if(!isset($data[0]) || empty($data[0])) return array();
        $arr = json_decode(trim($data[0]), true);
        if(empty($arr)) return array();
        $new_arr = array();
        foreach ($arr as $key => $val) {
            $temp_ball = '';
            $temp_ball = explode(' ', $val['kjZNum']);
            $ball = $temp_ball[0] . ',' . $temp_ball[1] . ',' . $temp_ball[2];
            $new_arr[$key]['expect'] = $key;
            $new_arr[$key]['opencode'] = $ball;
            $new_arr[$key]['opentime'] = $val['kjDate'] . ' 21:15';
            $new_arr[$key]['opentimestamp'] = strtotime($val['kjDate'] . ' 21:15');
        }
        $last_expect_key = array_keys($new_arr, max($new_arr));
        $last_expect_key = $last_expect_key[0];
        $last_expect = $new_arr[$last_expect_key];

        //采集二
        $str = '';
        $str = Curl::run($continue_url, 'get');
        //如果采集不到数据
        if (empty($str)) {
            // echo 'Can not collection the url:' . $continue_url . PHP_EOL;
            return array();
        }
        $data = array();
        $arr = array();
        $new_arr = array();
        $preg = '/kjgg\"><b([\d\D]*)<\/span><\/div>/';
        preg_match_all($preg, $str, $data);
        if(!isset($data[1][0])) return array();
        $data = $data[1][0];
        $expect_preg = '/第(.*?)期/';
        $opencode_preg = '/red_ball">(.*?)<\/span/';
        $time_preg = '/时间：(.*?)<br/';
        $data .= '</span>';
        // var_dump($data);
        $new_arr = array();
        preg_match($expect_preg, $data, $expect);
        preg_match_all($opencode_preg, $data, $opencode);
        preg_match($time_preg, $data, $timec);
        if(!isset($expect[1])) return array();
        $new_arr['expect'] = intval(trim($expect[1]));
        if(!isset($opencode[1][1])) return array();
        $new_arr['opencode'] = trim($opencode[1][0]) . ',' . trim($opencode[1][1]) . ',' . trim($opencode[1][2]);
        $new_arr['opentime'] = $timec[1];
        $new_arr['opentimestamp'] = strtotime($timec[1]);

        $data = array();
        if (($last_expect['expect'] == $new_arr['expect']) && ($last_expect['opencode'] == $new_arr['opencode'])) {
            $data[0] = $new_arr;
            return $data;
        } else {
            return array();
        }
    }

    public static function getContinueData($time) {
        self::getData();
    }

}
