<?php

namespace libraries\spider;

use \helper\Curl;
use \helper\GetQishuOpentime;
use \helper\SpiderCommon;

class Mg_o {

    static $qishu = "";
    static $wait = [];

    public static function getData() {

        $content = Curl::run('http://powerball-lottery.net/45/Result', 'get');
        $content = json_decode($content, true);

        if (empty($content['data'])) {
            return;
        }
        foreach ($content['data'] as $key => $val) {
            $data[$key]['opentime'] = date('Y-m-d H:i:s', strtotime($val['time']));
            $data[$key]['expect'] = $val['issue'];
            $data[$key]['opencode'] = $val['num'];
            $data[$key]['opentimestamp'] = strtotime($val['time']);

            $data[$key]['oldcode'] = $data[$key]['opencode'];
            $data[$key]['opencode'] = SpiderCommon::regroup_auto_continuous($data[$key]['opencode'], 5);
        }

        return $data;
    }

/*
    public static function getData() {
        $content = Curl::run('http://powerball-lottery.net/', 'get');

        preg_match('/<div[^>]*class="jspPane"[^>]*>(.*?) <\/div>/si', $content, $match); // 目标容器
        preg_match_all('/<li[^>]*>(.*?)<\/li>/si', $match[0], $match2); // 所有行

        foreach ($match2[0] as $key => $val) {
            // 目标内容
            preg_match('/<h1[^>]*class="time"[^>]*>(.*?)<\/h1>/si', $val, $match4);
            preg_match('/<h2[^>]*class="issue"[^>]*>(.*?)<\/h2>/si', $val, $match5);
            preg_match_all('/<span[^>]*>(.*?)<\/span>/si', $val, $match6);

            $data[$key]['opentime'] = date('Y-m-d H:i:s', strtotime($match4[1]));
            $data[$key]['expect'] = str_replace('Draw No.', '', $match5[1]);
            $data[$key]['opencode'] = implode(',', $match6[1]);
            $data[$key]['opentimestamp'] = strtotime($match4[1]);

            $data[$key]['oldcode'] = $data[$key]['opencode'];
            $data[$key]['opencode'] = SpiderCommon::regroup_auto_continuous($data[$key]['opencode'], 5);
        }

        return $data;
    }
*/

    public static function getContinueData() {
        self::getData();
    }

}
