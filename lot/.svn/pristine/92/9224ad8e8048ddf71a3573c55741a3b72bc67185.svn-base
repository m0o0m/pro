<?php

namespace libraries\spider;

use \helper\GetQishuOpentime;

class Ffc_o {

    static $qishu = "";
    static $wait = [];
    public static function getData() {
      $data = array();
      $now_qishu = GetQishuOpentime::get_qishu('ffc_o');
      $qishu = $now_qishu - 1;
      $time_arr = GetQishuOpentime::getOpentime('ffc_o', $qishu);
      if(!isset($time_arr['fengpan'])) return array();
      $kaijiang = $time_arr['fengpan'];
      if(time() > $kaijiang){
         $arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
         shuffle($arr);
         $data['expect'] = $qishu;
         $data['opentime'] = date('Y-m-d H:i:s');
         $data['opencode'] = implode(',', array_slice($arr,0,5)) ;
         $new_arr = array();
         $new_arr[] = $data;
         return $new_arr;
      }
    }
}
