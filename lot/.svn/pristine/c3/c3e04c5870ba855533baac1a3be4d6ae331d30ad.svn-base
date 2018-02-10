<?php
namespace libraries\auto;
use \helper\Common_helper;

class Bj_28Auto{
	//$result 开奖结果，从第一球到第三球
	//玩法：（对应数据库字段）共有大玩法1种
	public static function get_auto($result){
		if(count($result) != 3 )return array();
		$values = $balls_arr = self::get_balls_arr($result);//一到三球
		$total_sum = self::get_sum($result);//总和

		$gameplay = Common_helper::getGameplay('bj_28');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'bj_28':
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
			$res[1] = 'total_sum_big';
		}else{
			$res[1] = 'total_sum_small';
		}
		//和单和双
		if($sum%2 == 0)
			$res[2] = 'total_sum_double';
		else
			$res[2] = 'total_sum_single';

		return implode(',', $res);
	}
}