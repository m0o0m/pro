<?php

namespace Applications\Tools\Lib;

use \config\config;
use \helper\MysqlPdo as pdo;
use \helper\RedisConPool as Redis;
use \workerman\Lib\Timer;
use \helper\IdWork;
use \helper\Curl;

class Common {

    /**
     * **********************************************************
     *  获取需要补采的注单数据        @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function actionData($data) {
        $data = json_decode($data, true);
        $result = array();
        $result['ErrorCode'] = 2;

        $line_id = isset($data['line_id']) ? $data['line_id'] : '';
        $addday = isset($data['addday']) ? $data['addday'] : '';
        $fc_type = isset($data['fc_type']) ? $data['fc_type'] : '';
        $addtime = isset($data['addtime']) ? $data['addtime'] : '';

        if ($line_id == '' || $addday == '' || $addtime == '') {
            $result['ErrorMsg'] = '请求参数不正确';
            return json_encode($result);
        }
        $redis = Redis::getInstace();
        $addday = date('Ymd', strtotime($addday));

        //处理sql
        $old_sql = ' from my_bet_record where line_id=' . "'{$line_id}'" . ' and addday=' . $addday;
        if ($fc_type)
            $old_sql .= ' and fc_type=' . "'{$fc_type}'";

        $time_arr = self::dayTime($addday);

        //按时间段拼装sql
        $sql_arr = [];
        foreach ($time_arr as $key=>$val) {
            $sql_arr[$key]['count_sql'] ='select count(id)' .  $old_sql . ' and addtime>=' . $val['start'] . ' and addtime<' . $val['end'];
            $sql_arr[$key]['sql'] = 'select *' . $old_sql . ' and addtime>=' . $val['start'] . ' and addtime<' . $val['end'];
        }
        $lottery = pdo::instance('lottery');
        $manage = pdo::instance('manage');
        $work = new IdWork(1023); //获取集合唯一id
        $total_num = 0; //总查询出的数量
        $fail_num = 0; //存入采集redis失败的数量

        $pullredis = Redis::getInstace('pull');
        foreach ($sql_arr as $sql) {
            $count = 0;
            $count = $lottery->single($sql['count_sql']);
            if(!$count) continue;
            $data = array();
            $data_sql = $sql['sql'] . ' limit ' . $count;
            $data = $lottery->query($data_sql);
            if (!empty($data)) {
                foreach ($data as $key => $val) {
                    $total_num += 1;
                    $score = $work->nextId();
                    $score = substr($score, -16);//score最大范围16位，超出自动转科学计数
                    $val['score'] = strval($score);
                    $res = $pullredis->zadd('spider_' . $line_id . '_data', $score, json_encode($val));
                    if (!$res) {
                        $fail_num += 1;
                    }
                }
            }
        }

        $remark = 'total_num:' . $total_num . ', fail_num: ' . $fail_num;
        $status_sql = 'update my_spider_record set status=2,remark=' . "'{$remark}'" . ' where addtime=' . $addtime . ' and line_id=' . "'{$line_id}'" . ' and addday=' . $addday . ' and fc_type=' . "'{$fc_type}'";

        $action_key = 'waitSpiderAction'; //判定数据是否正在处理的键
        if ($manage->query($status_sql)) {
            $redis->del($action_key);
        }
    }

    /**
     * **********************************************************
     *  按分钟切分一天         @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function dayTime($addday, $minutes = 10) {
        $seconds = $minutes * 60;
        $start_day = strtotime(date('Ymd', strtotime($addday))); //每天开始
        $end_day = $start_day + 86400; //每天结束
        $arr = array();
        while ($start_day < $end_day) {
            $tmp = array();
            $tmp['start'] = $start_day;
            $start_day += $seconds;
            if ($start_day >= $end_day) {
                //如果开始和结束间隔小于10分钟，不生成这条数据
                if (($start_day - $seconds) < ($end_day - 600)) {
                    $tmp['end'] = $end_day;
                    $arr[] = $tmp;
                    break;
                } else {
                    $tmp['start'] = $start_day - $seconds;
                    $tmp['end'] = $end_day;
                    $arr[] = $tmp;
                    break;
                }
            }
            $tmp['end'] = $start_day;
            $arr[] = $tmp;
        }

        // foreach($arr as $key =>$val){
        // 	echo $key . '->开始：' . date('Y-m-d H:i:s',$val['start']) . PHP_EOL . $key . '->结束：' . date('Y-m-d H:i:s',$val['end']) . PHP_EOL;
        // }

        return $arr;
    }

}
