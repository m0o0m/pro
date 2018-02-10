<?php

namespace Applications\Common\Helper;

use \Applications\Common\Config\Config;
use \Applications\Common\Helper\Db;
use \Applications\Common\Helper\Redis;

class GetQishuOpentime {
/**
	  ***********************************************************
	  *  获取当前期数           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
  public static function get_qishu($type) {
        if (empty($type)) {
            return false;
        }
        switch ($type) {
            case 'fc_3d':
            case 'pl_3':
                $action = 'get_fc_3d_qishu';
                break;
            case 'ah_k3':
            case 'gx_k3':
            case 'js_k3':
            case 'cq_ssc':
            case 'tj_ssc':
            case 'xj_ssc':
                $action = 'get_all_qishu';
                break;
            case 'bj_10':
                $action = 'get_bj_10_qishu';
                break;
            case 'ffc_o':
            case 'lfc_o':
            case 'els_o':
            case 'jsfc':
            case 'mg_o':
            case 'xdl_10':
            case 'dj_o':
            case 'jsliuhecai':
            case 'mnl_o':
                $action = 'get_gpc_qishu';
                break;
            case 'bj_28':
            case 'bj_kl8':
            case 'pc_28':
                $action = 'get_bj_28_qishu';
                break;
            case 'cq_ten':
                $action = 'get_cq_ten_qishu';
                break;
            case 'dm_28':
            case 'dm_klc':
                $action = 'get_dm_28_qishu';
                break;
            case 'gd_11':
            case 'gd_ten':
            case 'jx_11':
            case 'sd_11':
                $action = 'get_gd_11_qishu';
                break;
            case 'jnd_28':
            case 'jnd_bs':
                $action = 'get_jnd_28_qishu';
                break;
            case 'liuhecai':
                $action = 'get_liuhecai_qishu';
            break;
            default:
                return 1;
            break;
        }
        return self::$action($type);
    }
/**
	  ***********************************************************
	  *  根据期数获取开封盘时间        @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getOpentime($type, $qishu){
		$closetime = self::getFengpanByQishu($type, $qishu);
        $new_closetime['kaipan'] = strtotime($closetime['kaipan']);
        $new_closetime['fengpan'] = strtotime($closetime['fengpan']);
        $new_closetime['kaijiang'] = strtotime($closetime['kaijiang']);
        $new_closetime['now_time'] = time();
		return $new_closetime;
	}

    public static function getGametime($type, $qishu){
        return self::getFengpanByQishu($type, $qishu);
    }

/**
	  ***********************************************************
	  *  获取通用开封盘时间            @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
   	private static function getAllOpentime($type, $get = 'all'){
   	
   		$tab = config::$prefix . 'opentime';
        $manage = Db::instance('manage');
        $redis = Redis::instance();
        $redis_key = 'opentime_list';
        $opentime = $redis->get($redis_key);
        if($opentime){
            $opentime = json_decode($opentime,true);
        }else{
        	$sql = 'select * from ' . $tab . ' where status=1';
            $opentime = $manage->query($sql);
            if(!$opentime) return false;
            $redis->set($redis_key,json_encode($opentime));
        }
        $new_arr = array();
        foreach($opentime as $val){
            if($val['fc_type'] == $type){
                $new_arr[] = $val; 
            }
        }
        if($get == 'all') return $new_arr; //所有
        if($get == 'one'){//第一期
            foreach($new_arr as $val){
                if($val['qishu'] == 1){
                    return $val;
                }
            }
        }
        return false;
   	}
/**
	  ***********************************************************
	  *  获取六合彩 加拿大开封盘时间   @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function getOhterOpentime($type,$where){
		if($type == 'jnd_28') $type = 'jnd_bs';
		$tab = config::$prefix . $type . '_opentime';
        $manage = Db::instance('manage');
        $sql = 'select * from ' . $tab . ' where ' . $where . ' order by qishu desc';
        $data = $manage->row($sql);
        return $data;
	}
		

  /**
 * **********************************************************
 *  各种获取期数   福彩3d 排列3                          *
 * **********************************************************
 */
    private static function get_fc_3d_qishu($type) {
        $opentime = self::getAllOpentime($type,'one');
        $kaijiang = strtotime($opentime['kaijiang']);
        $now_time = time();
        if ($type == 'pl_3') {
            $date_Y = date('y', $now_time);
        } else {
            $date_Y = date('Y', $now_time);
        }
        $qishu = date("z", $now_time) + 1;
        if($now_time >= $kaijiang && $now_time <= strtotime('23:59:59')){
            $qishu += 1;
        }
        //判断是初一前还是初七后
        $tmp = isset(config::$fc_3d_pl_3) ? config::$fc_3d_pl_3 : false;
        if($tmp){
            $qishu -= 7;
        }
        $qishu = $date_Y . substr(strval($qishu + 1000), 1, 3);
        return $qishu;
    }

    /**
     * **********************************************************
     *  获取通用期数*
     * **********************************************************
     */
    private static function get_all_qishu($type) {
        $result = self::getAllOpentime($type, 'all');
        $now_time = date("H:i:s");
        $data_time = array();
        foreach ($result as $val) {
            //当前时间<开奖时间 并且开盘时间<当前时间
            if ($val['kaijiang'] == '00:00:00') {
                if (strtotime($val['kaipan']) <= strtotime($now_time)) {
                    $data_time = $val;
                    break;
                }
            }else{
                if (strtotime($now_time) < strtotime($val['kaijiang']) && strtotime($val['kaipan']) <= strtotime($now_time)) {
                    $data_time = $val;
                    break;
                }
            }
        }

        $date = date("Ymd");
        if($type == 'cq_ssc'){//这个时间段空白。。。
            if(strtotime($now_time) < strtotime('09:50:00') && strtotime($now_time) >= strtotime('01:55:00')){
                 return $date . '024';
            }
        }
        if (empty($data_time)) {
            //如果当前时间在当天十点后，默认期数取第二天第一期
           if($type == 'ah_k3') $end = '22:00:00';
           if($type == 'gx_k3') $end = '22:27:00';
           if($type == 'js_k3') $end = '22:10:00';
           if($type == 'cq_ssc') $end = '23:59:59';
           if($type == 'tj_ssc') $end = '23:00:00';
           if($type != 'xj_ssc'){
               if($now_time >= $end){
                   $date = date("Ymd", strtotime("+1 day"));
               }
           }
           $data_time = self::getAllOpentime($type, 'one');
        }

        //判断是否跨天
        if ((strtotime($now_time) < strtotime('02:00:00')) && (strtotime($now_time) >= strtotime('00:00:00')) && $type == 'xj_ssc') {
            $date = date("Ymd", strtotime("-1 day"));
        }
        if ((strtotime($now_time) < strtotime('10:00:00')) && (strtotime($now_time) >= strtotime('02:00:00')) && $type == 'xj_ssc') {
            $data_time = self::getAllOpentime($type, 'one');
            $date = date("Ymd");
        }
        //返回ymd.补零后的期数
        return $date . str_pad($data_time['qishu'], 3, '0', STR_PAD_LEFT);
    }

    /**
     * **********************************************************
     *  获取北京pk10期数                                               *
     * **********************************************************
     */
    private static function get_bj_10_qishu() {
        $now = time();
        // bj_pk10初始数据
        $old_qishu = 489173;
        $old_lizi = 1431446220; // 固定不要动
        $left = 9 * 60 + 2; // 每天第一期开盘时间(分钟)
        $now_time = ceil(($now - $old_lizi - 3 * 60) / 60 % (60 * 24) + 0.1); // 已过去的当天分钟总数

        $time = $now - $old_lizi - 3 * 60; // 秒数
        $day = floor(($time / (60 * 60 * 24)) - (7 * 2)); // 天数
        // 判断期数
        if ($time > 0) {
            $today = ceil(($now_time - $left) / 5);
            if ($today > 179)
                $today = 179;
            if ($now_time >= $left) {
                $old_qishu += ($day * 179 + $today);
                return $old_qishu;
            } else {
                $old_qishu += ($day * 179); // 当天第一期
                return $old_qishu;
            }
        } else {
            return false;
        }
    }
    
     //获取高频彩期数
    public static function get_gpc_qishu($type){
        switch($type){
            case 'ffc_o':
            case 'jsfc':
                $date = date('ymd');
                $time = 60;
                $zero = 4;
            break;
            case 'lfc_o':
                $date = date('ymd');
                $time = 120;
                $zero = 3;
            break;
            case 'jsliuhecai':
                $date = date('Ymd');
                $time = 120;
                $zero = 3;
            break;
            case 'els_o':
            case 'xdl_10':
                $date = date('Ymd');
                $time = 90;
                $zero = 3;
            break;
            case 'mg_o':
                $date = date('Ymd');
                $time = 45;
                $zero = 4;
            break;
            case 'mnl_o':
                $date = date('ymd');
                $time = 45;
                $zero = 4;
            break;
            case 'dj_o':
                return self::get_dj_qishu();
            break;
        }
        $seconds = strtotime(date('Y-m-d'));//当天开始
        $now_seconds = time() - $seconds;
        $now_minute = ceil($now_seconds/$time);
        return $date . str_pad( $now_minute, $zero ,"0",STR_PAD_LEFT );
    }



     /**
      ***********************************************************
      *  获取东京1.5分彩期数        @author ruizuo qiyongsheng    *
      ***********************************************************
    */
    public static function get_dj_qishu(){
         $now_time = time();
         $dj_time = $now_time + 3600; //东京时间比北京时间快一小时
         $date = date('Ymd',$dj_time);
         $seconds = strtotime(date('Y-m-d', $dj_time));//当天开始
         $eight = $seconds + 8 * 3600; //八点钟
         $nine = $seconds + 9 * 3600; //九点钟
         
        if($dj_time <= $eight){
            $now_seconds = $dj_time - $seconds;
            $now_minute = ceil($now_seconds/90);
        }elseif( $dj_time > $eight && $dj_time <= $nine){
             $now_minute = (8 * 3600) / 90 + 1 ; //进入下一期
        }else{
            $now_seconds = $dj_time - $seconds;
            $now_minute = ceil($now_seconds/90);
            $now_minute -= 40;
        }
        return  $date . str_pad( $now_minute, 3 ,"0",STR_PAD_LEFT );
    }

    /**
     * **********************************************************
     *  获取北京28 北京快乐8期数                                             *
     * **********************************************************
     */
    private static function get_bj_28_qishu($now = '') {
        $now = time();
        // bj_kl8初始数据
        $old_qishu = 694602 - 1;
        ///$old_lizi = strtotime("2015-05-12 23:55:00"); // 固定不要动
        $old_lizi = 1431446100; // 固定不要动
        $left = 9 * 60; // 每天第一期开盘时间(分钟)
        $now_time = ceil(($now - $old_lizi - 5 * 60) / 60 % (60 * 24) + 0.1); // 已过去的当天分钟总数
        $time = $now - $old_lizi - 3 * 60; // 秒数
        $day = floor(($time / (60 * 60 * 24)) - (7 * 2)); // 天数
        // 判断期数
        if ($time > 0) {
            $today = ceil(($now_time - $left) / 5);
            if ($today > 179)
                $today = 179;
            if ($now_time >= $left) {
                $old_qishu += ($day * 179 + $today);
                return $old_qishu;
            } else {
                $old_qishu += ($day * 179); // 当天第一期
                return $old_qishu;
            }
        } else {
            return false;
        }
    }

    /**
     * **********************************************************
     *  获取重庆10分期数                                               *
     * **********************************************************
     */
    private static function get_cq_ten_qishu($type) {
        $result = self::getAllOpentime($type,'all');
        $now_time = date("H:i:s");
        $data_time = array();
        $date = date('ymd');
        $tomorrow =  date('ymd',strtotime("+1days"));
        //第一期 第十四期 特殊情况
        if(strtotime($now_time) >= strtotime('23:53:00') || strtotime($now_time) < strtotime('00:03:00')){
            return $tomorrow . '001';
        }
        if(strtotime($now_time) >= strtotime('01:53:00') && strtotime($now_time) < strtotime('09:53:00')){
            return $date . '014';
        }
        foreach ($result as $val) {
            //当前时间<开奖时间 并且开盘时间<当前时间
            if (strtotime($now_time) < strtotime($val['kaijiang']) && strtotime($val['kaipan']) <= strtotime($now_time)) {
                $data_time = $val;
                break;
            }
        }
        if (!$data_time) {
           $data_time = self::getAllOpentime($type, 'one');
        }
        //返回ymd.补零后的期数
        return $date . str_pad($data_time['qishu'], 3, '0', STR_PAD_LEFT);
    }

    /**
     * **********************************************************
     *  获取丹麦28丹麦快乐彩期数                    
     * **********************************************************
     */
    private static function get_dm_28_qishu($type) {
        $time = time() - 1481864680;
        return 1788567 + (($time - $time % 300) / 300);
    }

    /**
     * **********************************************************
     *  获取广东11选5期数                                               *
     * **********************************************************
     */
    private static function get_gd_11_qishu($type) {
        $result = self::getAllOpentime($type, 'all');
        $now_time = date("H:i:s");
        $data_time = array();
        $date = date("Ymd");
        foreach ($result as $val) {
            //当前时间<开奖时间 并且开盘时间<当前时间
            if (strtotime($now_time) < strtotime($val['kaijiang']) && strtotime($val['kaipan']) <= strtotime($now_time)) {
                $data_time = $val;
                break;
            }
        }
        if (empty($data_time)) {
           $end = '23:00:00';
           if($type == 'sd_11') $end = '22:55:00';
           if($now_time >= $end){
               $date = date("Ymd", strtotime("+1 day"));
           }
            
           $data_time = self::getAllOpentime($type, 'one');
        }
        //返回ymd.补零后的期数
        return $date . str_pad($data_time['qishu'], 2, '0', STR_PAD_LEFT);
    }

    /**
     * **********************************************************
     *  获取加拿大28期数                                               *
     * **********************************************************
     */
    private static function get_jnd_28_qishu($type) {
        //取redis 判断当前时间和取出的期数是否相符
        //相符Ok，不符删除更新 此redis时效五分钟
        $redis = Redis::instance();
        $now_time = date('Y-m-d H:i:s', time());
        $redis_key = 'c_lot_time_' . $type;
        $info = $redis->get($redis_key);
        $qishu = '';
        if ($info && $info !== 'false') {
            $info = json_decode($info, true);
            //查看当前时间是否在开封盘时间之间
            if (($info['kaipan'] <= $now_time ) && ($info['kaijiang']) > $now_time) {
                $qishu = $info['qishu'];
            } else {
                $where = 'status=1 and kaipan<=' . "'{$now_time}'" . ' and ' . "'{$now_time}'" . '<kaijiang';
                $result = self::getOhterOpentime($type,$where);
                if($result){
                    $qishu = $result['qishu'];
                    $redis->setex($redis_key, 300, json_encode($result)); //五分钟时效}
                }
            }
        } else {
            $where = 'status=1 and kaipan<=' . "'{$now_time}'" . ' and ' . "'{$now_time}'" . '<kaijiang';
            $result = self::getOhterOpentime($type,$where);
            if ($result) {
                $qishu = $result['qishu'];
                $redis->setex($redis_key, 300, json_encode($result)); //五分钟时效
            }
        }
        //如果查询不到取默认值
        if (empty($qishu)) {
           $data_time = self::getOhterOpentime($type,'status=1');
           if($data_time){
                $qishu = $data_time['qishu'];
           }
        }
        return $qishu;
    }

    private static function get_liuhecai_qishu($type) {
        $redis = Redis::instance();
        $hours = date("H");
        $now_time = date('Y-m-d H:i:s', time());
        $data_time = array();
        $redis_key = 'c_lot_time_' . $type;
        $info = $redis->get($redis_key);
        $qishu = '';
        if ($info && $info !== 'false') {
            $info = json_decode($info, true);
            //查看当前时间是否在开封盘时间之间
            if (($info['kaipan'] <= $now_time ) && ($info['fengpan']) >= $now_time) {
                $qishu = $info['qishu'];
            } else {
                $where = "status ='1' and ('{$now_time}'>kaipan) and ('{$now_time}'<kaijiang) "; 
                $result = self::getOhterOpentime($type,$where);
                $redis->setex($redis_key, 300, json_encode($result)); //五分钟时效}
            }
        } else {
            $where = "status ='1' and ('{$now_time}'>kaipan) and ('{$now_time}'<kaijiang) "; 
            $result = self::getOhterOpentime($type,$where);
            if ($result) {
	            $redis->setex($redis_key, 300, json_encode($result)); //五分钟时效
                $qishu = $result['qishu'];
            }
        }
        //如果查询不到取默认值
        if (empty($qishu)) {
            $data_time = self::getOhterOpentime($type,'status=1');
            $qishu = $data_time['qishu'];
        }
        return $qishu;
    }

/**
      ***********************************************************
      *  根据彩种和期数获取开封盘时间  @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    private static function getFengpanByQishu($type, $qishu) {
        if(empty($type) || empty($qishu)) return array();
        if($type == 'bj_28' || $type == 'pc_28') $type = 'bj_kl8';
        if($type == 'jnd_28') $type = 'jnd_bs';
        if($type == 'dm_28') $type = 'dm_klc';
        $redis = Redis::instance();
        $redis_key = 'fengpan_time_' . $type;
        $time_arr = $redis->get($redis_key);
        if($time_arr) $time_arr = json_decode($time_arr,true);
       
        if($type == 'liuhecai' || $type == 'jnd_bs'){
            $where = 'qishu=' . $qishu;
            if($time_arr && isset($time_arr['qishu']) && ($time_arr['qishu'] == $qishu) ){
                //缓存数据合法  什么也不干
            }else{
                $time_arr = self::getTimeByQishu($where, $type);
                if(!$time_arr) return false;
                $redis->set($redis_key,json_encode($time_arr));
            }
           
            return $time_arr;
        }

        if($type == 'fc_3d' || $type == 'pl_3'){//固定的开封盘
            if($type == 'pl_3') $qishu = '20' . $qishu;
            $where = 'qishu=1 and fc_type=' . "'{$type}'";
            if(!$time_arr){
                $time_arr = self::getTimeByQishu($where, $type);
                if(!$time_arr) return false;
                $redis->set($redis_key,json_encode($time_arr));
            }
            $kaijiang = $time_arr['kaijiang'];
            $fengpan = $time_arr['fengpan'];
            $kaipan = $time_arr['kaipan'];
            $str_len = strlen($qishu);
            $date = substr($qishu, 0,4);
            $qishu = substr($qishu,4,$str_len-4);
            $qishu = intval($qishu) - 1;
            //判断是初一前还是初七后
            $tmp = isset(config::$fc_3d_pl_3) ? config::$fc_3d_pl_3 : false;
            if($tmp){
                $qishu += 7;
            }
            $kaipan_day = date('Y-m-d',strtotime('+'. $qishu - 1 .' days',strtotime($date . '-01-01')));
            $fengpan_day =  date('Y-m-d',strtotime('+'. $qishu .' days',strtotime($date . '-01-01')));
            $data = array();
            $data['kaipan'] = $kaipan_day . ' ' . $kaipan;
            $data['fengpan'] = $fengpan_day . ' ' . $fengpan;
            $data['kaijiang'] = $fengpan_day . ' ' . $kaijiang;
            return $data;
        }
       //数据库中现有数据的彩种开封盘时间
       //gd_11 jx_11 9:00-23:00  sd_11 8:25-22:55
       //cq_ten 23:53-23:53 gd_ten 9:00-23:00
       //xj_ssc 10:00-2:00  tj_ssc 9:00-23:00 cq_ssc 00:00-23:00
       //js_k3 8:30-22:10 gx_k3 9:27-22:27 ah_k3 8:50-22:00

        $qishu_arr = ['gd_11','jx_11','sd_11','gd_ten','cq_ten','xj_ssc','tj_ssc','cq_ssc','js_k3','gx_k3','ah_k3'];
        if(in_array($type, $qishu_arr)){
            $str_len = strlen($qishu);
            if($type == 'cq_ten'){
                $date = '20' . substr($qishu, 0,6);
                $qishu = substr($qishu,6,$str_len-6);
            }else{
                $date = substr($qishu, 0,8);
                $qishu = substr($qishu,8,$str_len-8);
            }
            $qishu = intval($qishu);
            $where = 'qishu=' . $qishu . ' and fc_type=' . "'{$type}'";
      
            if($time_arr && isset($time_arr['qishu']) && ($time_arr['qishu'] == $qishu) ){
                //缓存数据合法  什么也不干
            }else{
                $time_arr = self::getTimeByQishu($where, $type);
                if(!$time_arr) return false;
                $redis->set($redis_key,json_encode($time_arr));
            }
            $data = array();
            $data['kaipan'] = date('Y-m-d H:i:s',strtotime($date . ' ' . $time_arr['kaipan']));
            $data['fengpan'] = date('Y-m-d H:i:s',strtotime($date . ' ' . $time_arr['fengpan']));
            $data['kaijiang'] = date('Y-m-d H:i:s',strtotime($date . ' ' . $time_arr['kaijiang']));
            if($type == 'cq_ten' && $qishu == 1){
                $yestoday = date('Y-m-d' ,strtotime('-1 days',strtotime($date)));
                $data['kaipan'] = $yestoday . ' ' . $time_arr['kaipan'];
            }
            if($type == 'xj_ssc' && $qishu >= 85){
                $tomorrow = date('Y-m-d' ,strtotime('+1 days',strtotime($date)));
                $data['kaipan'] = $tomorrow . ' ' . $time_arr['kaipan'];
                $data['fengpan'] = $tomorrow . ' ' . $time_arr['fengpan'];
                $data['kaijiang'] = $tomorrow . ' ' . $time_arr['kaijiang'];
            }
           
            return $data;
        }

        if($type == 'bj_10' || $type == 'bj_kl8'){//五分钟一期，4分封盘
            if($type == 'bj_kl8'){
                $begin_qishu = 694601;
                $begin_time = ' 09:00:00';
            } 
            if($type == 'bj_10'){
                $begin_qishu = 624497;
                $begin_time= ' 09:02:00';
            }
            if(strtotime(date('Y-m-d'). $begin_time)+24*3600 < time()){
                $first_time = strtotime(date('Y-m-d'). $begin_time)+24*3600;
            }else{
                $first_time = strtotime(date('Y-m-d') . $begin_time);
            } 
            if($qishu < $begin_qishu){
                return FALSE;
            }
            $ToDayNow = ($qishu -$begin_qishu)%179;                            ////今天多少期
            if($ToDayNow == 0){
                $ToDayNow = 179;
            }else{
                $ToDayNow--;
            }
            $data = array();
            $data['kaipan'] = date('Y-m-d H:i:s',$first_time + $ToDayNow*300);
            $data['fengpan'] = date('Y-m-d H:i:s', $first_time + $ToDayNow*300 + 240);
            $data['kaijiang'] =  date('Y-m-d H:i:s', $first_time + $ToDayNow*300 + 300);
            return $data;
        }

        if($type == 'dm_klc'){
             $data = array();
             $start_qishu = 1788567;  //初始期数
             $start_time = 1481864680; //初始时间
             $over_qishu = $qishu - $start_qishu; //已过去期数
             $over_time = $over_qishu * 5 *60; //已经过去时间 五分钟一期
             $data['kaipan'] = date('Y-m-d H:i:s',$start_time + $over_time);
             $data['fengpan'] = date('Y-m-d H:i:s',$start_time + $over_time + 240);
             $data['kaijiang'] =  date('Y-m-d H:i:s',$start_time + $over_time + 300);
             return $data;
        }

           //高频彩
        if($type == 'jsliuhecai'){
            return self::getGpctimeByQishu($qishu, 120, 8);
        }

        if(in_array($type, ['lfc_o'])){
            return self::getGpctimeByQishu($qishu, 120, 6);
        }

        if(in_array($type, ['els_o', 'xdl_10'])){
            return self::getGpctimeByQishu($qishu, 90, 8);
        }

        if($type == 'jsfc' || $type == 'ffc_o'){
            return self::getGpctimeByQishu($qishu, 60, 6);
        }

        if(in_array($type, ['mg_o'])){
            return self::getGpctimeByQishu($qishu, 45, 8);
        }

        if(in_array($type, ['mnl_o'])){
            return self::getGpctimeByQishu($qishu, 45, 6);
        }

        if($type == 'dj_o'){
            $data = array();
            $date = substr($qishu, 0, 8);
            $str_len = strlen($qishu);
            $qishu = substr($qishu, 8, $str_len - 8);
            $dj_time = time() + 3600; //东京时间比北京时间快一小时
            $seconds = strtotime(date('Y-m-d', $dj_time));//当天开始
            if($qishu <= 320){ //东京八点钟以前
               $seconds += intval($qishu) * 90 - 90;
            }else{//东京八点钟以后
               $seconds += intval($qishu + 40) * 90 - 90;
            }
            $seconds -= 3600; //北京时间
            $data['kaipan'] = date('Y-m-d H:i:s',$seconds);
            $data['fengpan'] = date('Y-m-d H:i:s',$seconds + 35);
            $data['kaijiang'] = date('Y-m-d H:i:s',$seconds + 45);
            return $data;
        }
    }

    
    public static function getGpctimeByQishu($qishu, $qishu_time, $strlen){
        $data = array();
        $date = substr($qishu, 0, $strlen);
        if($strlen == 6) $date = '20' . $date;
        $str_len = strlen($qishu);
        $qishu = substr($qishu, $strlen, $str_len - $strlen);
        //秒数
        $seconds = strtotime($date) + $qishu * $qishu_time - $qishu_time;
        $data['kaipan'] = date('Y-m-d H:i:s',$seconds);
        $data['fengpan'] = date('Y-m-d H:i:s', $seconds + $qishu_time - 10);
        $data['kaijiang'] = date('Y-m-d H:i:s',$seconds + $qishu_time);
        return $data;
    }

/**
	  ***********************************************************
	  *  根据期数获取开封盘时间       @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getTimeByQishu($where, $type){
		if($type == 'jnd_28') $type = 'jnd_bs';
		if($type == 'liuhecai' || $type == 'jnd_bs'){
   			$tab = config::$prefix . $type . '_opentime';
		}else{
   			$tab = config::$prefix . 'opentime';
		}
		$sql = 'select * from ' . $tab . ' where ' . $where;
        $manage = Db::instance('manage');
		$time_arr = $manage->row($sql);
		return $time_arr;
	}
	
	
   /**
     * **********************************************************
     *  期数自增自减          @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function new_qishu($fc_type, $qishu, $jiajian = 1) {
        $type_array = array('tj_ssc', 'cq_ssc', 'xj_ssc', 'gd_ten', 'cq_ten', 'js_k3', 'gx_k3', 'ah_k3', 'gd_11', 'sd_11', 'jx_11', 'ffc_o', 'lfc_o', 'els_o', 'jsfc', 'jsliuhecai', 'dj_o', 'mnl_o', 'xdl_10', 'mg_o');
        if (in_array($fc_type, $type_array)) {
            switch ($fc_type) {
                case 'tj_ssc':
                    $end_qishu = 84;
                    $zero = 2;
                    break;
                case 'cq_ssc':
                    $end_qishu = 120;
                    $zero = 2;
                    break;
                case 'xj_ssc':
                    $end_qishu = 96;
                    $zero = 2;
                    break;
                case 'gd_ten':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'cq_ten':
                    $end_qishu = 97;
                    $zero = 2;
                    break;
                case 'js_k3':
                    $end_qishu = 82;
                    $zero = 2;
                    break;
                case 'gx_k3':
                    $end_qishu = 78;
                    $zero = 2;
                    break;
                case 'ah_k3':
                    $end_qishu = 80;
                    $zero = 2;
                    break;
                case 'gd_11':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'sd_11':
                    $end_qishu = 87;
                    $zero = 1;
                    break;
                case 'jx_11':
                    $end_qishu = 84;
                    $zero = 1;
                    break;
                case 'ffc_o':
                case 'jsfc' :
                    $end_qishu = 1440;
                    $zero = 1;
                    break;
                case 'lfc_o':
                case 'jsliuhecai':
                    $end_qishu = 720;
                    $zero = 1;
                break;
                case 'els_o' :
                case 'xdl_10':
                    $end_qishu = 960;
                    $zero = 1;
                break;
                case 'dj_o'  :
                    $end_qishu = 920;
                    $zero = 1;
                break;
                case 'mnl_o':
                case 'mg_o' :
                    $end_qishu = 1920;
                    $zero = 1;
                break;
            }
            return self::jiajian($fc_type, $qishu, $end_qishu, $jiajian, $zero);
        } else {
        	if($jiajian == 1) return $qishu + 1;
        	if($jiajian == 2) return $qishu - 1;
        }
    }

 /**
     * **********************************************************
     *  当期数增加或减少到尽头时返回正确的期数                 *
     * **********************************************************
 */
    private static function jiajian($fc_type, $qishu, $end_qishu, $jiajian = 1, $zero = 1) {
        $start_qishu = 1;
        $old_qishu = $qishu;
        $str_len = strlen($qishu);

          //日期为六位的彩种
        $six_arr = ['cq_ten', 'ffc_o', 'jsfc', 'lfc_o', 'mnl_o'];
        if (in_array($fc_type, $six_arr)) {
            $date_len = 6;
        } else {
            $date_len = 8;
        }
        $date = substr($qishu, 0, $date_len);
        $qishu = substr($qishu, $date_len, $str_len - $date_len);
        $qishu = intval($qishu);

        if ($jiajian == 1 && $qishu == $end_qishu) { //期数自增
            $tomorrow = date('Ymd', strtotime('+ 1days', strtotime($date)));
            if(in_array($fc_type, $six_arr)){
                $tomorrow = date('ymd', strtotime('+ 1days', strtotime($date)));
            }

            if ($zero == 1) {
                $qishu = $tomorrow . str_pad($start_qishu, 2, '0', STR_PAD_LEFT);
            } else {
                $qishu = $tomorrow . str_pad($start_qishu, 3, '0', STR_PAD_LEFT);
            }
            if( in_array($fc_type, ['ffc_o', 'jsfc', 'mg_o', 'mnl_o']) ){
                $qishu = $tomorrow . '0001';
            }elseif( in_array($fc_type, ['lfc_o', 'els_o', 'xdl_10', 'dj_o', 'jsliuhecai']) ){
                $qishu = $tomorrow . '001';
            }
        } elseif ($jiajian == 2 && $qishu == $start_qishu) { //期数自减
            $yesterday = date('Ymd', strtotime('- 1days', strtotime($date)));
            if(in_array($fc_type, $six_arr)){
                $yesterday = date('ymd', strtotime('- 1days', strtotime($date)));
            }
            if ($zero == 1) {
                $qishu = $yesterday . str_pad($end_qishu, 2, '0', STR_PAD_LEFT);
            } else {
                $qishu = $yesterday . str_pad($end_qishu, 3, '0', STR_PAD_LEFT);
            }
            if( in_array($fc_type, ['ffc_o', 'jsfc', 'mg_o', 'mnl_o']) ){
                $qishu = $yesterday . $end_qishu;
            }elseif( in_array($fc_type, ['lfc_o', 'els_o', 'xdl_10', 'dj_o', 'jsliuhecai']) ){
                $qishu = $yesterday . $end_qishu;
            }

        } elseif ($jiajian == 1) {
            $qishu = $old_qishu + 1;
        } elseif ($jiajian == 2) {
            $qishu = $old_qishu - 1;
        }
        return $qishu;
    }

}