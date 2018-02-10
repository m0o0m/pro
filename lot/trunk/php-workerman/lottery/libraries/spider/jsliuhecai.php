<?php

namespace libraries\spider;

use \helper\GetQishuOpentime;

class Jsliuhecai {

    static $qishu = "";
    static $wait = [];

    public static function getData() {
        $data = array();
        $now_qishu = GetQishuOpentime::get_qishu('jsliuhecai');
        $qishu = $now_qishu - 1;
        $time_arr = GetQishuOpentime::getOpentime('jsliuhecai', $qishu);
        if(!isset($time_arr['fengpan'])) return array();
        $kaijiang = $time_arr['fengpan']; 
        if(time() > $kaijiang){
            for ($i=1; $i <= 49; $i++) { 
                $arr[] = $i;
            }
            shuffle($arr);
            $data['expect'] = $qishu;
            $data['opentime'] = date('Y-m-d H:i:s');
            $data['opencode'] = implode(',', array_slice($arr,0,7)) ;
            $new_arr = array();
            $new_arr[] = $data;
            return $new_arr;
        }
    }
}
