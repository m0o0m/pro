<?php

namespace libraries\spider;

use \helper\Curl;
use \config\curl_config;

class Pl_3 {

    static $qishu = "";
    static $wait = [];
    static $type = "pl3";

    // static $url = "http://caipiao.163.com/order/pl3/";
    // static $continue_url = "http://www.lecai.com/lottery/pl3/";

    public static function getData() {
        $url = str_replace('[caipiao163]', curl_config::$caipiao163, curl_config::$caipiao163_url);
        $url = str_replace('[caipiao163_type]', static::$type, $url);
        $continue_url = str_replace('[lecai]', curl_config::$lecai, curl_config::$lecai_url);
        $continue_url = str_replace('[lecai_type]', static::$type, $continue_url);


        // 采集一
        $str = Curl::run($url, 'get');
        //如果采集不到数据
        if (empty($str)) {
            // echo 'Can not collection the url:' . $url . PHP_EOL;
            return array();
        }
        $preg = '/kjgg\"> <b([\d\D]*)<\/span>/';
        preg_match_all($preg, $str, $data);

        if(!isset($data[1][0])){
            return array();
        }
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
        $new_arr['expect'] = intval(trim($expect[1]));
        $new_arr['opencode'] = trim($opencode[1][0]) . ',' . trim($opencode[1][1]) . ',' . trim($opencode[1][2]);
        $new_arr['opentime'] = $timec[1];
        $new_arr['opentimestamp'] = strtotime($timec[1]);
        //采集二
        $html = Curl::run($continue_url, 'get');
        //如果采集不到数据
        if (empty($html)) {
            // echo 'Can not collection the url:' . $continue_url . PHP_EOL;
            return array();
        }
        $last_expect = array();
        if ($html) {
            $expect = array();
            $opencode = array();
            $html = str_replace(' ', '', $html);
            $expect_preg = '/<span>排列3<fontclass="numcolor">(.*?)<\/font>开奖<\/span>/';
            preg_match_all($expect_preg, $html, $expect);
            if (isset($expect[1][0])) {
                $expect = intval($expect[1][0]);
            }
            $opencode_preg = '/<ulclass="redball">(.*?)<\/ul><divclass="clear"><\/div><\/div>/';
            preg_match_all($opencode_preg, $html, $opencode);
            if (isset($opencode[1][0])) {
                $opencode = $opencode[1][0];
                $ball_preg = '/<li>(.*?)<\/li>/';
                preg_match_all($ball_preg, $opencode, $ball);
                if (isset($ball[1]) && isset($ball[0]) && isset($ball[2]) && !empty($expect)) {
                    $ball = $ball[1];
                    $opencode = $ball[0] . ',' . $ball[1] . ',' . $ball[2];
                    $last_expect = array();
                    $last_expect['expect'] = $expect;
                    $last_expect['opencode'] = $opencode;
                } else {
                    $last_expect = array();
                }
            }
        } else {
            $last_expect = array();
            //echo 'URL http://www.lecai.com/lottery/pl3/ did not collect data'; 
        }
        $data = array();
        $data[0] = $new_arr;
        //如果采集二采集到数据，就和采集一对比，如果采集二未采集到，直接返回采集一
        if ($last_expect) {
            if (($last_expect['expect'] == $new_arr['expect']) && ($last_expect['opencode'] == $new_arr['opencode'])) {
                return $data;
            } elseif ($last_expect['expect'] < $new_arr['expect']) {//采集二的数据暂时停止更新
                return $data;
            } else {
                return array();
            }
        } else {
            return $data;
        }
    }

    public static function getContinueData($time) {
        self::getData();
    }

}
