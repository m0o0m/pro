<?php
namespace libraries\auto;
use \helper\Common_helper;

class Ah_k3Auto{
	//$result 开奖结果，从第一球到第三球
	//玩法：（对应数据库字段）共有大玩法5种
	//独胆和两连都是三个数字，取独胆gameplay id
	public static function get_auto($result){
	// 第一球 ~ 第三球 和值 豹子 两连 对子 独胆
		if(count($result) != 3 )return array();
		$values = array();
		$values['ball_1'] = $result[0];
		$values['ball_2'] = $result[1];
		$values['ball_3'] = $result[2];
		$sum = self::get_sum($result); //总和
		$dd = implode(',', $result); //独胆 两连
		$baozi = self::get_baozi($result); //豹子
		$duizi = self::get_duizi($result); //对子
		//组成数据库字段
		$gameplay = Common_helper::getGameplay('ah_k3');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'sum':
					$values['ball_' . $val['id']] = $sum;//总和
				break;
				case 'gallbladder_ball':
					$values['ball_' . $val['id']] = $dd ;//独胆 两连
				break;
				case 'leopard':
					$values['ball_' . $val['id']] = $baozi ;//豹子
				break;
				case 'pairs':
					$values['ball_' . $val['id']] = $duizi ;//对子
				break;
			}
		}

		return $values;
	}

/**
	  ***********************************************************
	  *  和值           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_sum($result){
		$sum = array_sum($result);
		$res = array();
		$res[0] = $sum;
		//三个号码相同时为输
		//大小单双
		//三个数组相同时的号码组合
		$same = ['1,1,1','2,2,2','3,3,3','4,4,4','5,5,5','6,6,6'];
		$str = $result[0] . ',' . $result[1] . ',' . $result[2];

		if(!in_array($str, $same)){
			$res[1] = $res[2] =  '';
		
			if($sum >= 4 && $sum <= 10)
				$res[1] = 'small';
			else
				$res[1] = 'big';

			if($sum % 2 == 0)
				$res[2] = 'double';
			else
				$res[2] = 'single';
		}

		return  implode(',', $res);
	}	
/**
	  ***********************************************************
	  *  豹子           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_baozi($result){
		if($result[0] == $result[1] && $result[1] == $result[2]){
			$res = implode(',', $result);
		}else{
			$res = '';
		}
		return $res;
	}
/**
	  ***********************************************************
	  *  对子           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_duizi($result){
		if($result[0] == $result[1] && $result[1] == $result[2]){
			return '';
		}
		if($result[0] == $result[1]){
			return $result[0] . ',' . $result[1];
		}elseif($result[1] == $result[2]){
			return $result[1] . ',' . $result[2];
		}else{
			return '';
		}
	}
			
}