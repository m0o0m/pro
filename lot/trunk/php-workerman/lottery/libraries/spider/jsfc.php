<?php

namespace libraries\spider;

use \helper\GetQishuOpentime;

class Jsfc {

    static $qishu = "";
    static $wait = [];
    public static function getData() {
      $data = array();
      $now_qishu = GetQishuOpentime::get_qishu('jsfc');
      $qishu = $now_qishu - 1;
      $time_arr = GetQishuOpentime::getOpentime('jsfc', $qishu);
      if(!isset($time_arr['fengpan'])) return array();
      $kaijiang = $time_arr['fengpan'];
      if(time() > $kaijiang){
         $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
         shuffle($arr);
         $data['expect'] = $qishu;
         $data['opentime'] = date('Y-m-d H:i:s');
         $data['opencode'] = implode(',', array_slice($arr,0,10)) ;
         $new_arr = array();
         $new_arr[] = $data;
         return $new_arr;
      }
    }
}
