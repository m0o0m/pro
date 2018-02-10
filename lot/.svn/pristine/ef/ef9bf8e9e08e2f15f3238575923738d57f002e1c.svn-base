<?php
namespace libraries\auto;
use \helper\Common_helper;

class Fc_3dAuto{
	//$result 开奖结果，从第一球到第三球
	//玩法：（对应数据库字段）共有大玩法7种
	//1-3球 跨度 三连 独胆 总和龙虎
	public static function get_auto($result){
		if(count($result) != 3 )return array();
		$first_ball = self::get_ball($result[0]);
		$second_ball = self::get_ball($result[1]);
		$third_ball = self::get_ball($result[2]);
		$span_ball = self::get_span_ball($result);//跨度
		$dd = self::get_dd($result); //独胆
		$sl = self::get_sl($result); //三连
		$tiger = self::get_tiger($result); //总和龙虎

		//组成数据库字段
		$values = array();
		$gameplay = Common_helper::getGameplay('fc_3d');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'span_ball':
					$values['ball_' . $val['id']] = $span_ball;//跨度
				break;
				case 'gallbladder_ball':
					$values['ball_' . $val['id']] = $dd ;//独胆
				break;
				case 'triple_ball':
					$values['ball_' . $val['id']] = $sl ;//三连
				break;
				case 'tiger_ball':
					$values['ball_' . $val['id']] = $tiger ;//总和龙虎
				break;
			}
		}

		$values['ball_1'] = $first_ball;
		$values['ball_2'] = $second_ball;
		$values['ball_3'] = $third_ball;
		return $values;
	}

/**
	  ***********************************************************
	  *  计算第一到第三球            @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_ball($num){
		if($num > 9 || $num < 0) return 'wrong';
		$res = array();
		//单双
		if($num%2 == 0)
			$res[0] = 'double';
		else
			$res[0] = 'single';

		//大小
		if($num >= 5)
			$res[1] = 'big';
		else
			$res[1] = 'small';

		return $num . ',' . implode(',', $res);
	}

/**
	  ***********************************************************
	  *  跨度           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_span_ball($result){
		sort($result); //排序
        $res = $result[2] - $result[0]; //最大值-最小值
        return $res;
	}
		
/**
	  ***********************************************************
	  *  独胆           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_dd($result){
		return implode(',', $result);
	}
/**
	  ***********************************************************
	  *  三连           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_sl($result){
		$num1 = $result[0];
        $num2 = $result[1];
        $num3 = $result[2];
        $new_arr = array($num1,$num2,$num3);
        sort($new_arr); //排序
        $num_new1 = $new_arr[0] + 1; //方便处理顺子
        $num_new2 = $new_arr[1] + 1;

        if(($num1 == $num2) && ($num2 == $num3)){
            return 'leopard';//豹子
        }elseif( ($num1 == $num2) || ($num2 == $num3) || ($num1 == $num3) ){
            return 'pairs'; //对子
        }elseif( in_array(1, $new_arr) && in_array(9, $new_arr) && in_array(0, $new_arr) ){
            return 'straight'; //顺子
        }elseif( ($new_arr[1] == $num_new1) && ($new_arr[2] == $num_new2) ){
            return 'straight'; //顺子
        }elseif( (abs($num1 - $num2) == 1) || (abs($num2 - $num3) == 1) ||(abs($num1 - $num3) == 1) ){
            return 'Half_suitable'; //半顺
        }else{
            return 'Miscellaneous_six'; //杂6
        }
	}
	
/**
	  ***********************************************************
	  *  总和大小单双龙虎           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	
	private static function get_tiger($result){
		$sum = array_sum($result);
		$res = array();
		//总和大小
		if($sum >= 14)
            $res[0] = 'total_sum_big';
        else
            $res[0] = 'total_sum_small';
        //总和单双  
        if($sum % 2 == 0)
        	$res[1] = 'total_sum_double';
        else
        	$res[1] = 'total_sum_single';
        //龙虎和
        if($result[0] > $result[2]){
        	$res[2] = 'dragon';
        }elseif($result[0] < $result[2]){
        	$res[2] = 'tiger';
        }else{
        	$res[2] = 'sum';
        }
                
        return implode(',', $res);
	}	
	
	
	
}

