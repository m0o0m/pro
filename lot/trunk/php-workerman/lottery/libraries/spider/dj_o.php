<?php

namespace libraries\spider;

use \helper\Curl;
use \helper\GetQishuOpentime;
use \helper\SpiderCommon;

class Dj_o {

    static $qishu = "";
    static $wait = [];

    public static function getData() {

        $content = Curl::run('http://tokyokeno.jp/game/history', 'get');
        if(!$content) return array();
        preg_match('/<tbody[^>]*id="content_items"[^>]*>(.*?) <\/tbody>/si', $content, $match); // 目标容器
        preg_match_all('/<tr[^>]*>(.*?)<\/tr>/si', $match[0], $match2); // 所有行
        if(!isset($match2[0])) return array();
        foreach ($match2[0] as $key => $val) {
            preg_match_all('/<td[^>]*>(.*?)<\/td>/si', $val, $match3); // 目标内容

            $data[$key]['opentime'] = $match3[1][0];
            $data[$key]['expect'] = $match3[1][1];
            $data[$key]['opencode'] = strip_tags($match3[1][2]);
            $data[$key]['opentimestamp'] = strtotime($match3[1][0]);

            $data[$key]['expect'] = date('Ymd', $data[$key]['opentimestamp']) . $data[$key]['expect']; // 期数拼上日期
            $data[$key]['oldcode'] = $data[$key]['opencode'];
            $data[$key]['opencode'] = SpiderCommon::regroup_auto_continuous($data[$key]['opencode'], 5);
        }

        return $data;
    }

    public static function getContinueData() {
        self::getData();
    }

}
