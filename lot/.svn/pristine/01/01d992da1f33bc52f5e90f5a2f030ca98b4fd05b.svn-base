<?php

namespace helper;
use \helper\RedisConPool;
use \helper\MysqlPdo as pdo;
use \helper\Common_helper as common;

class SpiderCommon{

  /**
     * **********************************************************
     *  通用接口 开奖数据采集                               *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param array 检测开奖结果合法性 数组
     * @param string 采集链接一 开奖网
     * @param string 采集链接二 彩票控
     * @param int 1为正常采集，2为复彩
     * @return array 采集成功返回数据，不成功返回空数组
     * **********************************************************
     */
    public static function spiderData($check_arr, $url, $url2, $type = 1) {
        $fc_type = $check_arr['type'];
        $redis = RedisConPool::getInstace();
        $redis_key = 'spider_lot_data_';
        $data = $cpk_data = $new_arr = array(); //有可能只使用一个接口采集
        //采集一 开奖网
        if($url){
            $data = array();
            $content = Curl::run($url, 'get');
            $data_arr = json_decode($content, TRUE);
            if (isset($data_arr['data'])) {
                $data = $data_arr['data'];
            }
        }
        // 采集二 彩票控
        if($url2){
            $cpk_content = Curl::run($url2, 'get');
            $cpk_data = json_decode($cpk_content, true);
        }

        if ($data) {
            //处理采集数据格式有差异的彩种
            $data = self::actionSpider($fc_type, $data, 1);
        } 

        if ($cpk_data) {
            $new_arr = self::actionSpider($fc_type, $cpk_data, 2);
        }

        //如果是复彩，直接返回数据
        if ($type == 2) {
            if (!empty($data))
                return $data;
            if (!empty($new_arr))
                return $new_arr;
        }


        //如果是正常采集，并且两个采集都能采集到数据，一个存mongo, 一个存mysql
        $data_a = array();
        $data_b = array();
        if (isset($data[0]))
            $data_a = $data[0]; //采集一为开奖网存mongo表
        if (isset($new_arr[0]))
            $data_b = $new_arr[0]; //采集二为彩票控存mysql

        //开奖网数据存mongo $data_a
        if (isset($data_a['expect']) && isset($data_a['opencode'])) {
            $mongo = new MongoDBManager(); //操作mongo数据库
            $data_a['fc_type'] = $fc_type;

            switch ($fc_type) {
                case 'bj_28':
                case 'jnd_28':
                    $data_a['opencode'] = self::regroup_auto_interval($data_a['opencode'], 3);
                    break;
                case 'pc_28':
                case 'dm_28':
                    $data_a['opencode'] = self::regroup_auto_continuous($data_a['opencode'], 3);
                    break;
                case 'els_o':
                case 'mg_o':
                    $data_a['opencode'] = self::regroup_auto_continuous($data_a['opencode'], 5);
                    break;
            }

            //查询mogo中的最后一期，如果当前期数大于最后一期就存入
            $mongo_data = $mongo->executeQuery('auto', ['fc_type' => $fc_type], ['sort' => ['expect' => -1], 'limit' => 1]);
            if (!empty($mongo_data)) {
                $mongo_data = $mongo_data[0];
                $mongo_last_qishu = $mongo_data->expect;
                if ($data_a['expect'] > $mongo_last_qishu) {
                    $mongo->insertData($data_a, 'auto', true);
                }
            } else {
                $mongo->insertData($data_a, 'auto', true);
            }
        }

        //检查三小时内是否有漏彩的期数，有的话进行补采集
        self::getLostData($fc_type, $new_arr);

        //彩票控数据存mysql  $data_b
        if (isset($data_b['expect']) && isset($data_b['opencode'])) {
            $check_arr['expect'] = $data_b['expect']; //期数
            if(in_array($fc_type, ['els_o', 'mg_o'])){
                $data_b['opencode'] = self::regroup_auto_continuous($data_b['opencode'], 5);
            }
            $check_arr['ball'] = $data_b['opencode']; //开奖结果 逗号拼接
            $res = self::checkSpider($check_arr);
            if ($res) {
                $result[0] = $data_b;
                // var_dump($result);
                return $result;
            } else {
                return array();
            }
        } else {
            return array();
        }
    }

    /**
     * **********************************************************
     * 处理不同接口采集结果中特珠部分    @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function actionSpider($fc_type, $data, $type = 1) {
        $new_data = array();
        //采集链接一 开奖网
        if ($type == 1) {
            foreach ($data as $key => $val) {
                if (isset($val['opencode']) && isset($val['expect'])) {
                    $new_data[$key] = $val;
                    if ($fc_type == 'bj_kl8' || $fc_type == 'bj_28' || $fc_type == 'pc_28') {
                        //采集共21个球，采集一最后两球格式80+2
                        $tmp_ball = array();
                        $tmp_ball = explode(',', $val['opencode']);
                        $tmp_sum = 0;
                        foreach ($tmp_ball as $k => $v) {
                            if ($k == 19) {
                                $last_ball = explode('+', $v);
                                if (isset($last_ball[0])) {
                                    $last_ball = $last_ball[0];
                                } else {
                                    $last_ball = $v;
                                }
                            }
                        }
                        $tmp_ball[19] = $last_ball;
                        $new_data[$key]['opencode'] = implode(',', $tmp_ball);
                    }

                    if ($fc_type == 'gd_ten') {
                        //采集到的是20170906027 修正成期数为两位
                        $tmp_qishu = $val['expect'];
                        $length = strlen($tmp_qishu);
                        $tmp_date = substr($tmp_qishu, 0, 8);
                        $tmp_qishu = substr($tmp_qishu, 8, $length - 8);
                        $tmp_qishu = common::func_BuLing(intval($tmp_qishu));
                        $tmp_qishu = $tmp_date . $tmp_qishu;
                        $new_data[$key]['expect'] = $tmp_qishu;
                    }

                    if ($fc_type == 'cq_ten') {
                        //采集到的是20170906027 修正成170906027
                        $new_data[$key]['expect'] = substr($data[0]['expect'], 2, 9);
                    }
                } else {
                    unset($data[$key]);
                }
            }
            return $new_data;
        } 

		//采集链接二  彩票控
        if ($type == 2) {
            foreach ($data as $key => $val) {
                if (isset($val['dateline']) && isset($val['number'])) {
                    $tmp_data = array();
                    $tmp_data['expect'] = $key;
                    $tmp_data['opencode'] = $val['number'];
                    $tmp_data['opentime'] = $val['dateline'];
                    $tmp_data['opentimestamp'] = strtotime($val['dateline']);
                    $new_data[] = $tmp_data;
                } else {
                    unset($key);
                }
            }

            if (empty($new_data)) {
                return array();
            }

            foreach ($new_data as $key => $val) {
                if (isset($val['opencode']) && isset($val['expect'])) {
                    if ($fc_type == 'bj_kl8' || $fc_type == 'bj_28' || $fc_type == 'pc_28') {
                        //采集共21个球，最后两球相加
                        $tmp_ball = array();
                        $tmp_ball = explode(',', $val['opencode']);
                        foreach ($tmp_ball as $k => $v) {
                            if ($k == 20) {
                                unset($tmp_ball[$k]);
                            }
                        }
                        $new_data[$key]['opencode'] = implode(',', $tmp_ball);
                    }

                    if ($fc_type == 'gx_k3' || $fc_type == 'js_k3' || $fc_type == 'gd_11') {
                        //采集数据期数开头为17 补成2017
                        $new_data[$key]['expect'] = '20' . $val['expect'];
                    }

                    if ($fc_type == 'xj_ssc') {
                        //采集数据期数格式：日期+两位期数 改成 日期+三位期数
                        $qishu = $val['expect'];
                        $str_count = strlen($qishu);
                        $date = substr($qishu, 0, 8);
                        $str = substr($qishu, 8, $str_count - 8);
                        $str = common::func_BuLings(intval($str));
                        $new_data[$key]['expect'] = $date . $str;
                    }
                } else {
                    unset($new_data[$key]);
                }
            }
            return $new_data;
        }
    }

/**
      ***********************************************************
      *  计算幸运彩开奖结果           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function get_28_auto($fc_type, $old_auto){
        $new_auto = '';
        if ($fc_type == 'bj_28' || $fc_type == 'jnd_28') {
            $ball_arr = explode(',', $old_auto);
            sort($ball_arr); //从低到高排序
            $ball_1 = $ball_2 = $ball_3 = null;
            $ball_1 = $ball_arr[1] + $ball_arr[4] + $ball_arr[7] + $ball_arr[10] + $ball_arr[13] + $ball_arr[16]; // 第一区
            $ball_2 = $ball_arr[2] + $ball_arr[5] + $ball_arr[8] + $ball_arr[11] + $ball_arr[14] + $ball_arr[17]; //第二区
            $ball_3 = $ball_arr[3] + $ball_arr[6] + $ball_arr[9] + $ball_arr[12] + $ball_arr[15] + $ball_arr[18]; //第三区 
            //取相加和的尾数
            $ball_1 = $ball_1 % 10;
            $ball_2 = $ball_2 % 10;
            $ball_3 = $ball_3 % 10;
            $new_auto = $ball_1 . ',' . $ball_2 . ',' . $ball_3;
        } elseif ($fc_type == 'pc_28') {
            $ball_arr = explode(',', $old_auto);
            sort($ball_arr); //从低到高排序
            $ball_1 = $ball_2 = $ball_3 = null;
            $ball_1 = $ball_arr[0] + $ball_arr[1] + $ball_arr[2] + $ball_arr[3] + $ball_arr[4] + $ball_arr[5]; // 第一区
            $ball_2 = $ball_arr[6] + $ball_arr[7] + $ball_arr[8] + $ball_arr[9] + $ball_arr[10] + $ball_arr[11]; //第二区
            $ball_3 = $ball_arr[12] + $ball_arr[13] + $ball_arr[14] + $ball_arr[15] + $ball_arr[16] + $ball_arr[17]; //第三区 
            //取相加和的尾数
            $ball_1 = $ball_1 % 10;
            $ball_2 = $ball_2 % 10;
            $ball_3 = $ball_3 % 10;
            //返回尾数以便存入数据库
            $new_auto = $ball_1 . ',' . $ball_2 . ',' . $ball_3;
        } elseif ($fc_type == 'dm_28') {
            $ball_arr = explode(',', $old_auto);
            sort($ball_arr); //从低到高排序
            $ball_1 = $ball_2 = $ball_3 = null;
            $ball_1 = $ball_arr[0] + $ball_arr[1] + $ball_arr[2] + $ball_arr[3] + $ball_arr[4] + $ball_arr[5];
            $ball_2 = $ball_arr[6] + $ball_arr[7] + $ball_arr[8] + $ball_arr[9] + $ball_arr[10] + $ball_arr[11];
            $ball_3 = $ball_arr[12] + $ball_arr[13] + $ball_arr[14] + $ball_arr[15] + $ball_arr[16] + $ball_arr[17];
            $ball_1 = $ball_1 % 10;
            $ball_2 = $ball_2 % 10;
            $ball_3 = $ball_3 % 10;
            $new_auto = $ball_1 . ',' . $ball_2 . ',' . $ball_3;
        }

        return $new_auto;
    }

    /**
      ***********************************************************
      *  计算高频彩结果           @author ruizuo qiyongsheng    *
      ***********************************************************
    */
    public static function get_gpc_auto($balls, $ball_count = 5) {
        if (!is_array($balls)) {
            $balls = explode(',', $balls);
        }

        sort($balls); // 由小到大排序

        $balls = array_chunk($balls, count($balls) / $ball_count);
        foreach ($balls as $key => $ball) {
            $result[] = array_sum($ball) % 10;
        }
        $result = implode(',', $result);

        return $result;
    }

    /**
     * 重组开奖结果 号码间隔
     * @param balls 原开奖号码
     * @param ball_count 最终球的个数
     * @author Frank
     */
    public static function regroup_auto_interval($balls, $ball_count = 3){
        if (!is_array($balls)) {
            $balls = explode(',', $balls);
        }
        sort($balls); // 由小到大排序

        array_shift($balls); // 去除多余球
        array_pop($balls); // 去除多余球
        foreach ($balls as $key => $ball) {
            $tmp[$key % $ball_count][] = $ball;
        }
        foreach ($tmp as $key => $ball) {
            $result[] = array_sum($ball) % 10;
        }
        $result = implode(',', $result);

        return $result;
    }

    /**
     * 重组开奖结果 号码连续
     * @param balls 原开奖号码
     * @param ball_count 最终球的个数
     * @author Frank
     */
    public static function regroup_auto_continuous($balls, $ball_count = 5) {
        if (!is_array($balls)) {
            $balls = explode(',', $balls);
        }
        sort($balls); // 由小到大排序

        $balls = array_chunk($balls, intval(count($balls) / $ball_count));
        foreach ($balls as $key => $ball) {
            if ($key < $ball_count) // 去除多余球
                $result[] = array_sum($ball) % 10;
        }
        $result = implode(',', $result);

        return $result;
    }

    /**
     * **********************************************************
     *  检测采集结果是否遗漏           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getLostData($fc_type, $data) {
        $not_arr = array(
            'fc_3d', //福彩3d
            'pl_3', //排列3
            'dm_28', //丹麦28
            'dm_klc', //丹麦快乐彩
            'liuhecai',//六合彩
            'ffc_o', //分分彩
            'lfc_o', //两分彩
            'jsliuhecai', //极速六合彩
            'jsfc' //极速飞车å
        );
        if (in_array($fc_type, $not_arr)) {
            return;
        }
        $redis = RedisConPool::getInstace();
        $redis_key = 'check_lost_spider_' . $fc_type; //每十分钟检查一次
        $is_ok = $redis->get($redis_key);
        if ($is_ok)
            return;
        if (empty($fc_type) || empty($data))
            return;
        if (count($data) < 3)
            return;
        $manage = pdo::instance('manage');
        $count = count($data) - 1;
        $small_qishu = $data[$count]['expect'];
        $big_qishu = $data[0]['expect'];
        $auto_table = common::GetAutoTableNameByType($fc_type);

        $qishu_arr = array();
        $sql = 'select qishu from ' . $auto_table . ' where qishu between ' . $small_qishu . ' and ' . $big_qishu;
        $qishu_arr = $manage->query($sql);
        if (empty($qishu_arr))
            return;
        foreach ($qishu_arr as $key => $val) {
            $qishu_arr[$key] = $val['qishu'];
        }
        $insert = array();
        foreach ($data as $key => $val) {
            if ($qishu_arr) {//检测数据库中该期结果是否存在
                if (in_array($val['expect'], $qishu_arr)) {
                    continue;
                }
            }
            $open_time = strtotime($val["opentime"]);
            $tmp_opencode = $val['opencode'];
            switch ($fc_type) {
                case 'bj_28':
                case 'jnd_28':
                    $tmp_opencode = self::regroup_auto_interval($val['opencode'], 3);
                    break;
                case 'pc_28':
                case 'dm_28':
                    $tmp_opencode = self::regroup_auto_continuous($val['opencode'], 3);
                    break;
                case 'els_o':
                case 'mg_o':
                    $tmp_opencode = self::regroup_auto_continuous($val['opencode'], 5);
                    break;
            }

            $open_res = explode(',', $tmp_opencode);
            $class = 'libraries\auto\\' . ucfirst($fc_type) . 'Auto';
            $auto_res = array();
            $auto_res = $class::get_auto($open_res);
            if (empty($auto_res)) {
                // $result['ErrorMsg'] = '获取玩法数组失败';
                return;
            }
            $auto_res['qishu'] = $val['expect'];
            $auto_res['datetime'] = $open_time;
            $auto_res['addtime'] = time();
            $insert[] = $auto_res;
        }
        if (empty($insert)) {
            return;
        }
        //插入数据库
        foreach ($insert as $val) {
            $field = '';
            $str = '';
            foreach ($val as $k => $v) {
                $field .= $k . ',';
                $str .= "'{$v}'" . ',';
            }
            $sql = "insert into $auto_table (" . rtrim($field, ',') . ") values(" . rtrim($str, ',') . ")";
            //采集结果插入数据库
            $res = $manage->query($sql);
            if ($res) {//插入成功，触发结算
                $port = common::getPort($fc_type);
                Curl::run('http://127.0.0.1:' . $port[1], 'post', array(
                    'todo' => 'balance',
                    'type' => $fc_type,
                    'qishu' => $val['qishu'] //期数
                ));
            }
        }
        $redis->setex($redis_key, 6000, 'ok');
        return;
    }

  /**
     * **********************************************************
     *  简单验证spider采集开奖结果的合法性                      *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param  array [maxexpect]最大期数[minexpect]最小期数[maxball minball]期数[expect] 最大球最小球[ballcount]球个数 [ball]球
     * @return boolean 合法返回true 不合法返回false
     * **********************************************************
   */
    public static function checkSpider($arr) {
        $class = '\libraries\balance\\' . ucfirst($arr['type']) . 'Balance';
        $now_qishu = $class::get_qishu(); //获取该彩种的当前期数
        
        if (($arr['expect'] < $arr['minexpect']) || ($arr['expect'] > $arr['maxexpect'])) {
            // echo 'The num is error';
            common::write_error_log('The spider expect is error type=' . $arr['type'] . ' qishu=' . $arr['expect'], 1, $arr['type'] . '.txt');
            return false;
        }
        //如果采集期数大于当前期数
        if ($arr['expect'] > $now_qishu) {
            // echo 'The expect is illegal!';
            common::write_error_log('Expect is illegal type=' . $arr['type'] . ' qishu=' . $arr['expect'] . 'local_qishu=' . $now_qishu, 1, $arr['type'] . '.txt');
            // return false;
        }

        //验证开奖结果合法性
        $ball = explode(',', $arr['ball']);
        if (count($ball) != $arr['ballcount']) {  //球的个数
            // echo 'The ball is error';
            common::write_error_log('The ball is error type=' . $arr['type'] . ' qishu=' . $arr['expect'] . ' ball=' . $arr['ball'], 1, $arr['type'] . '.txt');
            return false;
        }

        foreach ($ball as $val) { //球的大小范围
            if ((intval($val)) > $arr['maxball'] || (intval($val) < $arr['minball'])) {
                return false;
            }
        }

        return true;
    }

}