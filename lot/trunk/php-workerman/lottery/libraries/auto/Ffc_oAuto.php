<?php
namespace libraries\auto;
use \helper\Common_helper;

class Ffc_oAuto{
	//$result 开奖结果，从第一球到第五球
	//玩法：（对应数据库字段）共有大玩法1种
	public static function get_auto($result){
		if(count($result) != 5 )return array();
		$values = $balls_arr = self::get_balls_arr($result);//一到五球
		$total_sum = self::get_sum($result);//总和
		//组成数据库字段
		$gameplay = Common_helper::getGameplay('ffc_o');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'total_sum':
					$values['ball_' . $val['id']] = $total_sum;//总和龙虎
				break;
			}
		}

		return $values;
	}
/**
	  ***********************************************************
	  *  1-5球 单双大小          @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_balls_arr($result){
		$arr = array();
		foreach($result as $key=>$val){
			$res = array();
			//大小
			if($val >= 5)
				$res[0] = 'big';
			else
				$res[0] = 'small';
			//单双
			if($val % 2 == 0)
				$res[1] = 'double';
			else
				$res[1] = 'single';

			$arr['ball_' . ($key+1)] = $val . ',' . implode(',', $res);
		}

		return $arr;
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
		if($sum > 22){
			$res[1] = 'total_sum_big';
		}else{
			$res[1] = 'total_sum_small';
		}
		//和单和双
		if($sum%2 == 0)
			$res[2] = 'total_sum_double';
		else
			$res[2] = 'total_sum_single';
		if($result[0] > $result[4]){
			$res[3] = 'dragon';
		}elseif($result[0] < $result[4]){
			$res[3] = 'tiger';
		}else{
			$res[3] = 'sum';
		}
		return implode(',', $res);
	}
}