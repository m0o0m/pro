<?php
namespace libraries\auto;
use \helper\Common_helper;

class Dm_28Auto{
	//$result 开奖结果，从第一球到第三球
	//玩法：（对应数据库字段）共有大玩法1种
	public static function get_auto($result){
		if(count($result) != 3 )return array();
		$values = $balls_arr = self::get_balls_arr($result);//一到三球
		$total_sum = self::get_sum($result);//总和

		$gameplay = Common_helper::getGameplay('dm_28');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'dm_28':
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
	  *  总和大小单双          @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_sum($result){
		$sum = array_sum($result);
		$res = array();
		$res[0] = $sum;
		//和大和小
		if($sum > 13){
			$res[1] = 'big_num';
		}else{
			$res[1] = 'small_num';
		}
		//和单和双
		if($sum%2 == 0)
			$res[2] = 'double_num';
		else
			$res[2] = 'single_num';
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

		return implode(',', $res);
	}
}