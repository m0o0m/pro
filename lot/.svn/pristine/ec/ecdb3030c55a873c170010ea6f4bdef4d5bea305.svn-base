<?php

namespace libraries\spider;

use \helper\Curl;
use \helper\MysqlPdo as pdo;
use \config\curl_config;
use \helper\RedisConPool;

class Jnd_bs_opentime {

    static $qishu = "";
    static $wait = [];
    static $code = "cakeno";

    // static $url = "http://lotto.bclc.com/services2/keno/draw/latest/today?_=1481041604376";
    // static $continue_url = "http://c.apiplus.net/daily.do?token=6d8caf8c25bc44f5&code=cakeno&format=json&date=";

    public static function getData($table, $hour) {
        $redis = RedisConPool::getInstace();
        $manage = pdo::instance('manage');
        $redis_key = $table . 'last_qishu';
        $url = curl_config::$jnd_bs_opentime_url;
        $content = Curl::run($url, 'get');
        if (empty($content))
            return;
        $data = json_decode($content, TRUE);
        if (empty($data))
            return;
        //如果redis里不存在 从数据库查询最后一期的期数
        $last_sql = 'select qishu from ' . $table . ' order by qishu desc limit 1';
        $last_qishu = $redis->get($redis_key);
        if (!$last_qishu) {
            $last_qishu = $manage->row($last_sql);
            $last_qishu = $last_qishu['qishu'];
            $redis->setex($redis_key, 600, $last_qishu);
        }

        $last_expect = $data[0]; //最后一期
        $expect = $last_expect['drawNbr'];
        if($expect <= $last_qishu){//如果采集到的期数小于数据库中最后一期，不使用
            return;
        }
        $jnd_date = $last_expect['drawDate'];
        $jnd_time = $last_expect['drawTime'];
        $str_time = $jnd_date . ' ' . $jnd_time;
        $time = strtotime($str_time); //本期开奖时间(按采集数据最后一期为本期)
        $next_expect = $expect - 1; //本期(循环中会先执行增加1)
        $next_kaipan = $time + $hour * 3600 - 7 * 60; //下期开盘时间(顺延15小时)
        $next_fengpan = $next_kaipan + 3 * 60; //下期封盘时间(顺延15小时3分钟)
        $next_kaijiang = $next_fengpan + 30; //下期开奖时间(顺延15小时3分钟30秒)
        // $next_expect = $expect + 1; //下期期数
        // $next_kaipan = $time + $hour * 3600; //下期开盘时间(顺延15小时)
        // $next_fengpan = $time + $hour * 3600 + 3 * 60; //下期封盘时间(顺延15小时3分钟)
        // $next_kaijiang = $time + $hour * 3600 + 3 * 60 + 30; //下期开奖时间(顺延15小时3分钟30秒)

        $fix_time = date('His', strtotime('19:00:00'));
        //如果下期开盘时间 在 19：00：00 以前 开盘时间为19:07:00 以后顺延一天
        $tmp_time = ' 19:07:00';
        if($hour == 16){
            $tmp_time = ' 20:07:00';
        }
        if (date('His', $next_kaipan) <= $fix_time) {
            $time1 = date('Y-m-d', $next_kaipan) . $tmp_time;
        } else {
            $time1 = date('Y-m-d', ($next_kaipan + 3600 * 24)) . $tmp_time;
        }

        $time1 = strtotime($time1);
        $time2 = $next_kaipan;
        $addtime = 210; //顺延3分钟30秒
        $num = 0;

        $arr = array();
        while (($num < 400) && ($time2 < $time1)) {
            $next_expect += 1;
            $num += 1;
            $next_kaipan += $addtime;
            $next_fengpan += $addtime;
            $next_kaijiang += $addtime;
            $time2 = $next_kaijiang;
            $qishu = '';
            $kaipan = '';
            $fengpan = '';
            $kaijiang = '';
            $qishu = $next_expect;
            $kaipan = date('Y-m-d H:i:s', $next_kaipan);
            $fengpan = date('Y-m-d H:i:s', $next_fengpan);
            $kaijiang = date('Y-m-d H:i:s', $next_kaijiang);
            $arr[$qishu]['kaipan'] = $kaipan;
            $arr[$qishu]['fengpan'] = $fengpan;
            $arr[$qishu]['kaijiang'] = $kaijiang;
        }
        if (empty($arr))
            return;

        $sql = array();
        $tmp_qishu = '';
        foreach ($arr as $qishu => $val) {
            if (!empty($last_qishu)) {
                if ($qishu > $last_qishu) {
                    $sql[] = "('{$qishu}','{$val['kaipan']}','{$val['fengpan']}','{$val['kaijiang']}')";
                    $tmp_qishu = $qishu;
                }
            } else {
                $sql[] = "('{$qishu}','{$val['kaipan']}','{$val['fengpan']}','{$val['kaijiang']}')";
                $tmp_qishu = $qishu;
            }
        }
        if (empty($sql))
            return true;
        if (count($sql) != 1) {
            $insert = "insert into `" . $table . "` (qishu,kaipan,fengpan,kaijiang) values " . implode(',', $sql);
        } else {
            $insert = "insert into `" . $table . "` (qishu,kaipan,fengpan,kaijiang) values " . $sql[0];
        }
        $res = $manage->query($insert);
        if ($res) {
            if ($tmp_qishu) {
                $redis->setex($redis_key, 600, $tmp_qishu); //将最后一期存入redis
            }
        }
        return $res;
    }

}
