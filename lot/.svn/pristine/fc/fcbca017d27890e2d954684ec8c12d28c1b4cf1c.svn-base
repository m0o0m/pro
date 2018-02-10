<?php

namespace Applications\Stat\Lib;

use \config\config;
use \helper\MysqlPdo as Db;
use \helper\RedisConPool as Redis;
use \workerman\Lib\Timer;

class Common {

    //获取内存中的注单数量并执行结算
    public static function getDataFromRedis($num) {
        $redis = Redis::getInstace();
        $redis_key = 'AccountFromRedis';

        //从redis取出指定条数数据
        $data = array();
        $data = $redis->lrange($redis_key, 0, $num - 1);
        if(empty($data)) return;
        //删除已经取出的数据
        $redis->ltrim($redis_key, $num, -1);
        $new_arr = array();
        foreach ($data as $val) {
            $new_arr[] = json_decode($val, true);
        }
        self::addStatData($new_arr);
    }

    //统计数据来自redis
    public static function addStatData($data) {
        if (!is_array($data) || empty($data)) {
            return false;
        }
        /* $test = [
          ['unique_key' => 'a', 'bet' => 100, 'num' => 20],
          ['unique_key' => 'b', 'bet' => 100, 'num' => 20],
          ['unique_key' => 'c', 'bet' => 100, 'num' => 20],
          ['unique_key' => 'a', 'bet' => 50, 'num' => 10]
          ];

          $new = [];
          foreach ($test as $k => $v) {
          if (isset($new[$v['unique_key']])) {
          $new[$v['unique_key']]['bet'] += $v['bet'];
          $new[$v['unique_key']]['num'] += $v['num'];
          } else {
          $new[$v['unique_key']] = $v;
          }
          } */
        //var_dump($new);
        //$data = self::getData();
        $day_data = [];
        $line_data = [];
        //1:按天分组
        foreach ($data as $k => $v) {
            //fc_type 和 uid组合成唯一键
            $v['unique_key'] = $v['uid'] . $v['fc_type'];
            //替换不合法数据(测试模拟的数据)
            $v['addday'] = str_replace('-', '', $v['addday']);
            $day_data[$v['addday']][] = $v;
            //line_id 和 fc_type组合成唯一键
            $v['unique_key'] = $v['line_id'] . $v['fc_type'];
            $line_data[$v['addday']][] = $v;
        }
        /* count(1) as bet_count,
          sum(bet) as bet,
          sum(case when status in (2, 3,6,7) then 1 else 0 end) as valid_bet_count,
          sum(case when status in (2, 3,6,7) then bet else 0 end) as valid_bet,
          sum(case when status in (2,6,7) then 1 else 0 end) as win_count,
          sum(case when status in (2,6,7) then bet else 0 end) as win */
        //2:按会员分组求和(总打码,有效打码,总派彩,总注单数,有效注单数,赢注单数)
        $account_data = [];
        $fina_day_data = [];
        foreach ($day_data as $day => $user) {
            foreach ($user as $key => $val) {
                if (isset($account_data[$val['unique_key']])) {
                    //总注单数
                    $account_data[$val['unique_key']]['bet_count'] += 1;
                    //总打码
                    $account_data[$val['unique_key']]['bet'] += $val['bet'];
                    //有效注单数 有效打码
                    if (in_array($val['status'], ['2', '3', '6', '7'])) {
                        $account_data[$val['unique_key']]['valid_bet_count'] += 1;
                        $account_data[$val['unique_key']]['valid_bet'] += $val['valid_bet'];
                    }
                    //赢注单数 总派彩
                    if (in_array($val['status'], ['2', '6', '7'])) {
                        $account_data[$val['unique_key']]['win_count'] += 1;
                        $account_data[$val['unique_key']]['win'] += $val['win'];
                    }
                } else {
                    $account_data[$val['unique_key']] = $val;
                    //总注单数
                    $account_data[$val['unique_key']]['bet_count'] = 1;
                    //有效注单数 有效打码
                    if (in_array($val['status'], ['2', '3', '6', '7'])) {
                        $account_data[$val['unique_key']]['valid_bet_count'] = 1;
                        $account_data[$val['unique_key']]['valid_bet'] = $val['valid_bet'];
                    } else {
                        $account_data[$val['unique_key']]['valid_bet_count'] = 0;
                        $account_data[$val['unique_key']]['valid_bet'] = 0;
                    }
                    //赢注单数 总派彩
                    if (in_array($val['status'], ['2', '6', '7'])) {
                        $account_data[$val['unique_key']]['win_count'] = 1;
                        $account_data[$val['unique_key']]['win'] = $val['win'];
                    } else {
                        $account_data[$val['unique_key']]['win_count'] = 0;
                        $account_data[$val['unique_key']]['win'] = 0;
                    }
                }
            }
            $fina_day_data[$day] = $account_data;
        }
        //3:入库
        foreach ($fina_day_data as $addday => $value) {
            self::accountByDay($addday, $value);
        }


        //线路统计
        $account = [];
        $fina_line_data = [];
        foreach($line_data as $day=>$line){
            foreach($line as $key => $val){
                if(isset($account[$val['unique_key']])){
                     //总注单数
                    $account[$val['unique_key']]['bet_count'] += 1;
                    //总打码
                    $account[$val['unique_key']]['bet'] += $val['bet'];
                    //有效注单数 有效打码
                    if (in_array($val['status'], ['2', '3', '6', '7'])) {
                        $account[$val['unique_key']]['valid_bet_count'] += 1;
                        $account[$val['unique_key']]['valid_bet'] += $val['valid_bet'];
                    }
                    //赢注单数 总派彩
                    if (in_array($val['status'], ['2', '6', '7'])) {
                        $account[$val['unique_key']]['win_count'] += 1;
                        $account[$val['unique_key']]['win'] += $val['win'];
                    }
                }else{
                    $account[$val['unique_key']] = $val;
                    //总注单数
                    $account[$val['unique_key']]['bet_count'] = 1;
                    //有效注单数 有效打码
                    if (in_array($val['status'], ['2', '3', '6', '7'])) {
                        $account[$val['unique_key']]['valid_bet_count'] = 1;
                        $account[$val['unique_key']]['valid_bet'] = $val['valid_bet'];
                    } else {
                        $account[$val['unique_key']]['valid_bet_count'] = 0;
                        $account[$val['unique_key']]['valid_bet'] = 0;
                    }
                    //赢注单数 总派彩
                    if (in_array($val['status'], ['2', '6', '7'])) {
                        $account[$val['unique_key']]['win_count'] = 1;
                        $account[$val['unique_key']]['win'] = $val['win'];
                    } else {
                        $account[$val['unique_key']]['win_count'] = 0;
                        $account[$val['unique_key']]['win'] = 0;
                    }
                }
            }
            $fina_line_data[$day] = $account;
        }
        //入库
        foreach($fina_line_data as $day => $val){
            self::accountLineByDay($day, $val);
        }
    }

    //会员统计入库(按天)
    public static function accountByDay($addday, $data) {
        //这个地方的$data就等于 select ... from my_bet_record where ... group by fc_type,uid
        //获取统计那一天记录表里面所有uid
        $uids = self::getUids($addday, '', '');

        //变一维数组
        if (!empty($uids)) {
            foreach ($uids as $k => $v) {
                $uids_arr[] = $v['unique_key'];
            }
        } else {
            $uids_arr = array();
        }
        $insert_sql = 'insert into my_stat_data_bet (`fc_type`,`line_id` ,`sh_id`,`ua_id`,  `at_id`,`uid`, `at_username`, `uname`, `bet_count`,`bet` ,`valid_bet_count` , `valid_bet` ,`win_count` , `win` , `addday`,`addtime`,`updatetime`) values ';
        $if_insert = false;
        $if_update = false;
        $update_arr = [];
        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                if (!in_array($v['uid'] . $v['fc_type'], $uids_arr)) {
                    $if_insert = true;
                    //拼接插入sql
                    $insert_sql .= '("' . $v['fc_type'] . '","' . $v['line_id'] . '",' . $v['sh_id'] . ',' . $v['ua_id']
                            . ',' . $v['at_id'] . ',' . $v['uid'] . ',"' . $v['at_username'] . '","' . $v['uname'] . '",'
                            . $v['bet_count'] . ',' . $v['bet'] . ',' . $v['valid_bet_count'] .
                            ',' . $v['valid_bet'] . ',' . $v['win_count'] . ',' . $v['win'] . ','
                            . $addday . ',' . time() . ',' . time() . ')' . ',';
                } else {
                    $if_update = true;
                    $v['updatetime'] = time();
                    $update_arr[] = $v;
                }
            }
        }
        //定义要更新的字段 方便做循环
        $update_cloumns = array('bet_count', 'bet', 'valid_bet_count', 'valid_bet', 'win_count', 'win', 'updatetime');
        //分组update 一次update3000条数据
        $update_arr_group = array_chunk($update_arr, 3000);
        $updatesql_arr = array();
        foreach ($update_arr_group as $key => $val) {
            $update = 'update my_stat_data_bet set ';
            $set = '';
            foreach ($update_cloumns as $cloumn) {
                $where_in = '';
                $case = '';
                //累加update
                if (in_array($cloumn, ['bet_count', 'bet', 'valid_bet' ,'valid_bet_count', 'win_count', 'win'])) {
                    $set .= $cloumn . '=' . $cloumn . '+(case concat(uid,fc_type) ';
                    foreach ($val as $k => $v) {
                        $case .= 'when "' . $v['uid'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                        $where_in .= '"' . $v['uid'] . $v['fc_type'] . '",';
                    }
                    $case .= ' end) ,';
                } else {
                    $set .= $cloumn . '= case concat(uid,fc_type) ';
                    foreach ($val as $k => $v) {
                        $case .= 'when "' . $v['uid'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                        $where_in .= '"' . $v['uid'] . $v['fc_type'] . '",';
                    }
                    $case .= ' end ,';
                }

                $set .= $case;
            }
            //去拼接的最后的一个逗号
            $where_in = substr($where_in, 0, -1);
            $set = substr($set, 0, -1);
            //更新
            //$update_sql .= $update . $set . ' where day_type=' . $datetype . ' and uid in (' . $where_in . ')' . ';';
            $update_sql = $update . $set . ' where addday=' . $addday . ' and concat(uid,fc_type) in (' . $where_in . ')';
            $updatesql_arr[] = $update_sql;
        }

        //插入
        $insert_sql = substr($insert_sql, 0, -1);
        //执行sql
        if ($if_insert) {
            Db::instance('manage')->query($insert_sql);
        }
        if ($if_update) {
            foreach ($updatesql_arr as $k => $v) {
                Db::instance('manage')->query($v);
            }
        }
    }

    //线路统计入库（按天）
    public static function accountLineByDay($addday, $data){
        //获取统计那一天记录表里面所有line_id
        $line_day = self::getLineIds($addday, '', '');
         //变一维数组
        $line_arr = [];
        if (!empty($line_day)) {
            foreach ($line_day as $k => $v) {
                $line_arr[] = $v['unique_key'];
            }
        } else {
            $line_arr = array();
        }

        $insert_sql = 'insert into my_stat_data_line_bet(`fc_type`,`line_id`, `bet` , `bet_count`, `valid_bet`, `valid_bet_count`, `win` , `win_count` , `addday`, `addtime`, `updatetime`) values ';

        $if_insert = false;
        $if_update = false;
        $update_arr = [];
        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                if (!in_array($v['line_id'] . $v['fc_type'], $line_arr)) {
                    $if_insert = true;
                    //拼接插入sql
                    $insert_sql .= '("' . $v['fc_type'] . '","' . $v['line_id'] . '",' . $v['bet'] . ',' . $v['bet_count'] . ',' . $v['valid_bet'] . ',' . $v['valid_bet_count'] . ',' . $v['win'] . ',' . $v['win_count'] . ','. $addday . ',' . time() . ',' . time() . ')' . ',';
                } else {
                    $if_update = true;
                    $v['updatetime'] = time();
                    $update_arr[] = $v;
                }
            }
        }

        //定义要更新的字段 方便做循环
        $update_cloumns = array('bet_count', 'bet', 'valid_bet_count', 'valid_bet', 'win_count', 'win', 'updatetime');
        //分组update 一次update3000条数据
        $update_arr_group = array_chunk($update_arr, 3000);
        $updatesql_arr = array();
        foreach ($update_arr_group as $key => $val) {
            $update = 'update my_stat_data_line_bet set ';
            $set = '';
            foreach ($update_cloumns as $cloumn) {
                $where_in = '';
                $case = '';
                //累加update
                if (in_array($cloumn, ['bet_count', 'bet', 'valid_bet', 'valid_bet_count', 'win_count', 'win'])) {
                    $set .= $cloumn . '=' . $cloumn . '+(case concat(line_id,fc_type) ';
                    foreach ($val as $k => $v) {
                        $case .= 'when "' . $v['line_id'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                        $where_in .= '"' . $v['line_id'] . $v['fc_type'] . '",';
                    }
                    $case .= ' end) ,';
                } else {
                    $set .= $cloumn . '= case concat(line_id,fc_type) ';
                    foreach ($val as $k => $v) {
                        $case .= 'when "' . $v['line_id'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                        $where_in .= '"' . $v['line_id'] . $v['fc_type'] . '",';
                    }
                    $case .= ' end ,';
                }

                $set .= $case;
            }
            //去拼接的最后的一个逗号
            $where_in = substr($where_in, 0, -1);
            $set = substr($set, 0, -1);
            //更新
            //$update_sql .= $update . $set . ' where day_type=' . $datetype . ' and uid in (' . $where_in . ')' . ';';
            $update_sql = $update . $set . ' where addday=' . $addday . ' and concat(line_id,fc_type) in (' . $where_in . ')';
            $updatesql_arr[] = $update_sql;
        }

        //插入
        $insert_sql = substr($insert_sql, 0, -1);
        //执行sql
        if ($if_insert) {
            Db::instance('manage')->query($insert_sql);
        }
        if ($if_update) {
            foreach ($updatesql_arr as $k => $v) {
                Db::instance('manage')->query($v);
            }
        }
    }

/**
      ***********************************************************
      *                      下面是重新统计                        *
      ***********************************************************
*/
    
    public static function reStat($connection, $header) {
        $redis = Redis::getInstace();
        $redis->set('bet_restat_task_is_running', 1);
        $result['connect'] = 1;

        $post = $header['post'];
        $datetype = isset($post['datetype']) ? $post['datetype'] : 'day';
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        $fc_type = isset($post['fc_type']) ? $post['fc_type'] : '';
        $day = isset($post['day']) ? $post['day'] : '';
        $starttime = isset($post['starttime']) ? trim($post['starttime']) : '';
        $endtime = isset($post['endtime']) ? trim($post['endtime']) : '';
        $timezone = isset($post['timezone']) ? trim($post['timezone']) : '';

        if ($datetype == 'day') {
            if ($day) {
                self::reStatByDay($day, $line_id, $fc_type);
            }
        } elseif ($datetype == 'interval') {
            if ($starttime && $endtime) {
                if ($timezone) {
                    date_default_timezone_set($timezone);
                }
                $many = date('z', $endtime) - date('z', $starttime);
                for ($i = 0; $i <= $many; $i++) {
                    $day = date('Ymd', strtotime(date('Y-m-d', $starttime) . " +{$i} day"));
                    self::reStatByDay($day, $line_id, $fc_type);
                }
            }
        }
        $redis->del('bet_restat_task_is_running');
        $connection->send(json_encode($result));
    }

    public static function reStatByDay($addday, $line_id = '', $fc_type = '') {
        $addday = preg_replace('/[-\/]+/', '', $addday);
        //会员重新统计
        self::userAccount($addday, $line_id, $fc_type);
        //线路重新统计
        self::lineAccount($addday, $line_id, $fc_type);
    }

    //会员统计
    public static function userAccount($addday, $line_id, $fc_type) {
        $where = "`js`=2 and `addday` = {$addday}";
        if ($line_id) {
            $where .= " AND `line_id` = '{$line_id}'";
        }
        if ($fc_type) {
            $where .= " AND `fc_type` = '{$fc_type}'";
        }
        $has = Db::instance('lottery')->select('COUNT(id)')->from('my_bet_record')->where($where)->single();
        if (!empty($has)) {
            $sql = " SELECT
                    `fc_type` ,
                    `uid` ,
                    min(`sh_id`) as sh_id,
                    min(`ua_id`) as ua_id,
                    min(`at_id`) as at_id ,
                    min(`line_id`) as line_id ,
                    min(`at_username`) as at_username ,
                    min(`uname`) as uname,
                    min(`addday`) as addday,
                    count(1) as bet_count,
                    sum(bet) as bet,
                    sum(case when status in (2, 3,6,7) then 1 else 0 end) as valid_bet_count,
                    sum(case when status in (2, 3,6,7) then valid_bet else 0 end) as valid_bet,
                    sum(case when status in (2,6,7) then 1 else 0 end) as win_count,
                    sum(case when status in (2,6,7) then win else 0 end) as win
                    FROM
                    my_bet_record
                    WHERE
                    $where
                    GROUP BY
                    fc_type ,
                    uid 
                    limit $has";
            $data = Db::instance('lottery')->query($sql);

            //获取统计那一天记录表里面所有uid
            $uids = self::getUids($addday, $line_id, $fc_type);
            //变一维数组
            if (!empty($uids)) {
                foreach ($uids as $k => $v) {
                    $uids_arr[] = $v['unique_key'];
                }
            } else {
                $uids_arr = array();
            }
            $insert_sql = 'insert into my_stat_data_bet 
                            (`fc_type`,`line_id` ,`sh_id`,`ua_id`,  `at_id`,`uid`, `at_username`, `uname`,
                            `bet_count`,`bet` ,`valid_bet_count` , `valid_bet` ,`win_count` , `win` ,  
                                 `addday`,`addtime`,`updatetime`)
                                values ';
            $if_insert = false;
            $if_update = false;
            if (is_array($data) && !empty($data)) {
                foreach ($data as $k => $v) {
                    if (!in_array($v['uid'] . $v['fc_type'], $uids_arr)) {
                        $if_insert = true;
                        //拼接插入sql
                        $insert_sql .= '("' . $v['fc_type'] . '","' . $v['line_id'] . '",' . $v['sh_id'] . ',' . $v['ua_id']
                                . ',' . $v['at_id'] . ',' . $v['uid'] . ',"' . $v['at_username'] . '","' . $v['uname'] . '",'
                                . $v['bet_count'] . ',' . $v['bet'] . ',' . $v['valid_bet_count'] .
                                ',' . $v['valid_bet'] . ',' . $v['win_count'] . ',' . $v['win'] . ','
                                . $addday . ',' . time() . ',' . time() . ')' . ',';
                    } else {
                        $if_update = true;
                        $v['updatetime'] = time();
                        $update_arr[] = $v;
                    }
                }
            }

            //定义要更新的字段 方便做循环
            $update_cloumns = array('bet_count', 'bet', 'valid_bet_count', 'valid_bet', 'win_count', 'win', 'updatetime');
            //分组update 一次update3000条数据
            $update_arr_group = array_chunk($update_arr, 3000);
            $updatesql_arr = array();
            foreach ($update_arr_group as $key => $val) {
                $update = 'update my_stat_data_bet set ';
                $set = '';
                foreach ($update_cloumns as $cloumn) {
                    $where_in = '';
                    $case = '';
                    $set .= $cloumn . '= case concat(uid,fc_type) ';
                    foreach ($val as $k => $v) {
                        $case .= 'when "' . $v['uid'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                        $where_in .= '"' . $v['uid'] . $v['fc_type'] . '",';
                    }
                    $case .= ' end ,';
                    $set .= $case;
                }
                //去拼接的最后的一个逗号
                $where_in = substr($where_in, 0, -1);
                $set = substr($set, 0, -1);
                //更新
                //$update_sql .= $update . $set . ' where day_type=' . $datetype . ' and uid in (' . $where_in . ')' . ';';
                $update_sql = $update . $set . ' where addday=' . $addday . ' and concat(uid,fc_type) in (' . $where_in . ')';
                $updatesql_arr[] = $update_sql;
            }

            //插入
            $insert_sql = substr($insert_sql, 0, -1);
            //执行sql
            if ($if_insert) {
                Db::instance('manage')->query($insert_sql);
            }
            if ($if_update) {
                foreach ($updatesql_arr as $k => $v) {
                    Db::instance('manage')->query($v);
                }
            }
        }
    }

    //获取会员统计表当天所有会员id
    public static function getUids($addday, $line_id, $fc_type) {
        $where = 'where addday =' . $addday;
        if (!empty($line_id)) {
            $where .= ' and line_id="' . $line_id . '"';
        }
        if (!empty($fc_type)) {
            $where .= ' and fc_type="' . $fc_type . '"';
        }
        $sql = "SELECT DISTINCT(concat(uid,fc_type)) as unique_key  from my_stat_data_bet $where ";
        return Db::instance('manage')->query($sql);
    }

    //线路统计
    public static function lineAccount($addday, $line_id, $fc_type) {
        $where = 'where addday=' . $addday;
        if ($line_id) {
            $where .= " and `line_id` = '{$line_id}'";
        }
        if ($fc_type) {
            $where .= " and `fc_type` = '{$fc_type}'";
        }
        $sql = "SELECT 
                line_id,fc_type,min(addday) as addday,
                sum(bet_count) as bet_count,
                sum(bet) as bet ,
                sum(valid_bet_count) as valid_bet_count ,
                sum(valid_bet) as valid_bet,
                sum(win_count) as win_count,
                sum(win) as win
                from my_stat_data_bet
                $where
                GROUP BY line_id,fc_type";
        $data = Db::instance('manage')->query($sql);

        //获取统计那一天记录表里面所有line_id
        $lineIds = self::getLineIds($addday, $line_id, $fc_type);
        //变一维数组
        if (!empty($lineIds)) {
            foreach ($lineIds as $k => $v) {
                $lineIds_arr[] = $v['unique_key'];
            }
        } else {
            $lineIds_arr = array();
        }

        $insert_sql = 'insert into my_stat_data_line_bet 
                            (`fc_type`,`line_id` ,
                            `bet_count`,`bet` ,`valid_bet_count` , `valid_bet` ,`win_count` , `win` ,  
                                 `addday`,`addtime`,`updatetime`)
                                values ';
        $if_insert = false;
        $if_update = false;
        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                if (!in_array($v['line_id'] . $v['fc_type'], $lineIds_arr)) {
                    $if_insert = true;
                    //拼接插入sql
                    $insert_sql .= '("' . $v['fc_type'] . '","' . $v['line_id'] . '",'
                            . $v['bet_count'] . ',' . $v['bet'] . ',' . $v['valid_bet_count'] .
                            ',' . $v['valid_bet'] . ',' . $v['win_count'] . ',' . $v['win'] . ','
                            . $addday . ',' . time() . ',' . time() . ')' . ',';
                } else {
                    $if_update = true;
                    $v['updatetime'] = time();
                    $update_arr[] = $v;
                }
            }
        }

        //定义要更新的字段 方便做循环
        $update_cloumns = array('bet_count', 'bet', 'valid_bet_count', 'valid_bet', 'win_count', 'win', 'updatetime');
        //分组update 一次update3000条数据
        $update_arr_group = array_chunk($update_arr, 3000);
        $updatesql_arr = array();
        foreach ($update_arr_group as $key => $val) {
            $update = 'update my_stat_data_line_bet set ';
            $set = '';
            foreach ($update_cloumns as $cloumn) {
                $where_in = '';
                $case = '';
                $set .= $cloumn . '= case concat(line_id,fc_type) ';
                foreach ($val as $k => $v) {
                    $case .= 'when "' . $v['line_id'] . $v['fc_type'] . '" then "' . $v[$cloumn] . '" ';
                    $where_in .= '"' . $v['line_id'] . $v['fc_type'] . '",';
                }
                $case .= ' end ,';
                $set .= $case;
            }
            //去拼接的最后的一个逗号
            $where_in = substr($where_in, 0, -1);
            $set = substr($set, 0, -1);
            //更新
            //$update_sql .= $update . $set . ' where day_type=' . $datetype . ' and uid in (' . $where_in . ')' . ';';
            $update_sql = $update . $set . ' where addday=' . $addday . ' and concat(line_id,fc_type) in (' . $where_in . ')';
            $updatesql_arr[] = $update_sql;
        }

        //插入
        $insert_sql = substr($insert_sql, 0, -1);
        //执行sql
        if ($if_insert) {
            Db::instance('manage')->query($insert_sql);
        }
        if ($if_update) {
            foreach ($updatesql_arr as $k => $v) {
                Db::instance('manage')->query($v);
            }
        }
    }

    //获取会员统计表当天所有会员id
    public static function getLineIds($addday, $line_id, $fc_type) {
        $where = 'where addday =' . $addday;
        if (!empty($line_id)) {
            $where .= ' and line_id="' . $line_id . '"';
        }
        if (!empty($fc_type)) {
            $where .= ' and fc_type="' . $fc_type . '"';
        }
        $sql = "SELECT DISTINCT(concat(line_id,fc_type)) as unique_key from my_stat_data_line_bet $where ";
        return Db::instance('manage')->query($sql);
    }
  
}
