<?php
namespace libraries\auto;
use \helper\Common_helper;

class Bj_kl8Auto{
	//$result 开奖结果，从第一球到第二十球
	//玩法：（对应数据库字段）共有大玩法8种
	//选一到选五 和值 奇偶盘 上中下盘
	public static function get_auto($result){
		if(count($result) != 20 )return array();
		$he = self::get_he($result); //和值
		$three_pan = self::get_three_pan($result); //上中下盘
		$two_pan = self::get_two_pan($result); //奇偶盘
		//组成数据库字段
		$values = array();
		$gameplay = Common_helper::getGameplay('bj_kl8');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'sum':
					$values['ball_' . $val['id']] = $he;//和值
				break;
				case 'up_middle_down':
					$values['ball_' . $val['id']] = $three_pan ;//上中下盘
				break;
				case 'odd_and_even':
					$values['ball_' . $val['id']] = $two_pan ;//奇偶盘
				break;
				case 'choose_one':
					$values['ball_' . $val['id']] = implode(',', $result) ;//结果集合
				break;
			}
		}
		foreach($result as $key=>$val){
			$values['ball_' . ($key+1)] = $val;
		}

		return $values;
	}

/**
	  ***********************************************************
	  *  和值           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_he($result){
		$sum = array_sum($result);
		$res = array();
		//总和大小
		if($sum > 810){
			$res[0] = 'total_sum_big';
		}elseif($sum < 810){
			$res[0] = 'total_sum_small';
		}else{
			$res[0] = 'total_sum_810';
		}
		//总和单双
		if($sum % 2 == 0){
			$res[1] = 'total_sum_double';
		}else{
			$res[1] = 'total_sum_single';
		}

		return $sum . ',' . implode(',', $res);
	}
	
/**
	  ***********************************************************
	  *  上中下盘          @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_three_pan($result){
		$up = 0;//上盘
        $down = 0;//下盘
        foreach( $result as $val ){
            if($val <= 40){
                $up += 1;
            }else{
                $down += 1;
            }
        }

        $res = array();
        if($up > $down){
        	return 'up_disc';
        }elseif($up < $down){
        	return 'down_disc';
        }else{
        	return 'middle_disc';
        }
	}
	
/**
	  ***********************************************************
	  *  奇偶盘           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_two_pan($result){
		$odd = 0; //寄盘
        $even = 0; //偶盘

        foreach( $result as $val ){
            if($val%2 == 0){
                $even += 1;
            }else{
                $odd += 1;
            }
        }
        if($odd > $even){
        	return 'odd_disc';
        }elseif($odd < $even){
        	return 'even_disc';
        }else{
        	return 'sum_disc';
        }
	}
	
	
	
}