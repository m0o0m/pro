<?php
namespace libraries\auto;
use \helper\Common_helper;

class Gd_tenAuto{
	//$result 开奖结果，从第一球到第五球
	//玩法：（对应数据库字段）共有大玩法14种
	//任选组选全部取'random_choose_two'的gameplay id
	public static function get_auto($result){
	// 第一球 ~ 第八球,总和龙虎，任选玩法四种 组选一种
		if(count($result) != 8 )return array();
		$values = $ball_arr = self::get_ball_arr($result); //一到八球
		$tiger_ball = self::get_tiger_ball($result); //总和龙虎
		$rx = $zx = implode(',', $result); //任选组选
		//组成数据库字段
		$gameplay = Common_helper::getGameplay('gd_ten');
        if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'tiger_ball':
					$values['ball_' . $val['id']] = $tiger_ball;//总和龙虎
				break;
				case 'random_choose_two':
					$values['ball_' . $val['id']] = $rx ;//任选组选
				break;
			}
		}

		return $values;		
	}

/**
     ***********************************************************
     * 计算 大 小 单 双  合数单双 尾大尾小 东南西北中发白      *
     ***********************************************************
*/
	private static function get_ball_arr($result){
        //方位数组
        $address_arr = array();
        $middle_arr = array();
        $address_arr['east'] = array(1,5,9,13,17); //东
        $address_arr['south']= array(2,6,10,14,18); //南
        $address_arr['west'] = array(3,7,11,15,19); //西
        $address_arr['north'] = array(4,8,12,16,20); //北
        $middle_arr['middle'] = array(1,2,3,4,5,6,7); //中
        $middle_arr['hair']  = array(8,9,10,11,12,13,14);  //发
        $middle_arr['white'] = array(15,16,17,18,19,20);  //白
		$arr = array();
		foreach($result as $key=>$val){
			$he = $one = $ten = 0;
			//头尾 合大小
			$one = $val % 10;//个位数
			$ten = floor($val/10);//十位数
			$he = $one + $ten;

			$res = array();
			//大小
			if($val >= 11)
				$res[0] = 'big';
			else
				$res[0] = 'small';
			//单双
			if($val % 2 == 0)
				$res[1] = 'double';
			else
				$res[1] = 'single';
			//尾大小
			if($one >= 5)
				$res[2] = 'end_big';
			else
				$res[2] = 'end_small';
			//合单双
			if($he % 2 == 0 )
				$res[3] = 'sum_double';
			else
				$res[3] = 'sum_single';

			foreach($address_arr as $k1=>$v1){
				if(in_array($val, $v1)){
					$res[4] = $k1;
				}
			}
			foreach($middle_arr as $k2=>$v2){
				if(in_array($val, $v2)){
					$res[5] = $k2;
				}
			}

			$arr['ball_' . ($key+1)] = $val . ',' . implode(',', $res);
		}

		return $arr;
	}
/**
	  ***********************************************************
	  *  总和大小单双尾大小 龙虎           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_tiger_ball($result){
        $sum = array_sum($result); 
		$sum_tail = $sum % 10;
		$res = array();
		//总和大小
		if($sum >= 85 && $sum <= 132){
			$res[0] = 'total_sum_big';
		}elseif($sum >= 37 && $sum <= 83){
			$res[0] = 'total_sum_small';
		}elseif($sum == 84){
			$res[0] = 'total_sum';
		}
		//总和单双
		if($sum % 2 == 0)
			$res[1] = 'total_sum_double';
		else
			$res[1] = 'total_sum_single';
		//尾大小
		if($sum_tail >= 5)
			$res[2] = 'total_sum_end_big';
		else
			$res[2] = 'total_sum_end_small';
		//龙虎
		if($result[0] > $result[7]){
			$res[3] = 'dragon';
		}elseif($result[0] < $result[7]){
			$res[3] = 'tiger';
		}else{
	
		}

		return $sum . ',' . implode(',', $res);
	}
	
	

}