<?php
namespace libraries\auto;
use \helper\Common_helper;

class Gd_11Auto{
	//$result 开奖结果，从第一球到第五球
	//玩法：（对应数据库字段）共有大玩法9种
	//任选组选直，只取 random_choose 一个gameplay id
	public static function get_auto($result){
		if(count($result) != 5 )return array();
		$values = $balls_arr = self::get_balls_arr($result);//一到五球
		$rx = implode(',', $result); //任选组选直选
		$total_sum = self::get_sum($result);//总和龙虎

		$gameplay = Common_helper::getGameplay('gd_11');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'total_sum':
					$values['ball_' . $val['id']] = $total_sum;//总和
				break;
				case 'random_choose':
					$values['ball_' . $val['id']] = $rx ;//组选直选
				break;
			}
		}

		return $values;
	}

/**
	  ***********************************************************
	  *  第一球-第五球数字大小单双和    @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_balls_arr($result){
		//11不计算输赢
		$balls_arr = array();
		foreach($result as $key=>$val){
			$res = array();
			$res[0] = $val;
			if($val >= 6)
				$res[1] = 'big';
			else
				$res[1] = 'small';
			if($val % 2 == 0)
				$res[2] = 'double';
			else
				$res[2] = 'single';

			if($val == 11){
				$balls_arr['ball_' . ($key+1)] = $val . ',' . 'sum,sum';
			}else{
				$balls_arr['ball_' . ($key+1)] = implode(',', $res);
			}
		}
		return $balls_arr;
	}
/**
	  ***********************************************************
	  *  总和大小单双 尾大小 龙虎           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_sum($result){
		$sum = array_sum($result);
		$sum_tail = $sum % 10;
		$res = array();
		$res[0] = $sum;
		//和大和小
		if($sum > 30){
			$res[1] = 'sum_big';
		}elseif($sum < 30){
			$res[1] = 'sum_small';
		}else{
			$res[1] = 'total_sum';
		}
		//和单和双
		if($sum%2 == 0)
			$res[2] = 'sum_double';
		else
			$res[2] = 'sum_single';
		//尾大小
		if($sum_tail >= 5)
			$res[3] = 'end_big';
		else
			$res[3] = 'end_small';
		//龙虎
		if($result[0] > $result[4])
			$res[4] = 'dragon';
		else
			$res[4] = 'tiger';

		return implode(',', $res);
	}

}