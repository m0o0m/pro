<?php
namespace libraries\auto;
use \helper\Common_helper;

class Pc_28Auto{
	//$result 开奖结果，从第一球到第三球
	//玩法：（对应数据库字段）共有大玩法1种
	public static function get_auto($result){
		if(count($result) != 3 )return array();
		$values = $balls_arr = self::get_balls_arr($result);//一到三球
		$total_sum = self::get_sum($result);//总和

		$gameplay = Common_helper::getGameplay('pc_28');
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'pc_28':
					$values['ball_' . $val['id']] = $total_sum;//总和
				break;
			}
		}
		return $values;
	}
/**
	  ***********************************************************
	  *  1-3球           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_balls_arr($result){
		$balls_arr = array();
		foreach($result as $key=>$val){
			$balls_arr['ball_' . ($key+1)] = $val;
		}
		return $balls_arr;
	}
/**
      ***********************************************************
      *  总和大小单双 总和大单 大双 小单 小双  极小 极大 红绿蓝波 豹子        *
      ***********************************************************
*/
	private static function get_sum($result){
		$sum = array_sum($result);
		$wave_arr = array();
		$wave_arr['red_wave'] = array(3,6,9,12,15,18,21,24); //红波
        $wave_arr['green_wave'] = array(1,4,7,10,16,19,22,25); //绿波
        $wave_arr['blue_wave'] = array(2,5,8,11,17,20,23,26); //蓝波
        // $fail_arr = array(0,27,13,14); //波色含有此号码为输
		$res = array();
		$res[0] = $sum;
		//和大和小
		if($sum > 13){
			$res[1] = 'total_sum_big';
		}else{
			$res[1] = 'total_sum_small';
		}
		//和单和双
		if($sum%2 == 0)
			$res[2] = 'total_sum_double';
		else
			$res[2] = 'total_sum_single';

		//大双大单小双小单
		 if( ($sum%2 == 1) && ($sum >= 15) && ($sum <= 27) ){
		 	$res[3] = 'big_single';
		 }elseif( ($sum%2 == 0) && ($sum >= 14) && ($sum <= 26) ){
		 	$res[3] = 'big_double';
		 }elseif( ($sum%2 == 1) && ($sum >= 1) && ($sum <= 13) ){
		 	$res[3] = 'small_single';
		 }elseif( ($sum%2 == 0) && ($sum >= 0) && ($sum <= 12) ){
		 	$res[3] = 'small_double';
		 }
		 //极小极大
		 if( ($sum >= 0) && ($sum <= 5) ) $res[4] = 'max_small';
		 if( ($sum >= 22) && ($sum <= 27) ) $res[4] = 'max_big';
		 //豹子
		 if($result[0] == $result[1] && $result[0] == $result[2])$res[5] = 'leopard';
		 //波色
		foreach($wave_arr as $wave=>$val){
			if(in_array($sum, $val)){
				$res[6] = $wave;
			}
		}
		return implode(',', $res);

	}
}