<?php
namespace libraries\auto;
use \helper\Common_helper;

class LiuhecaiAuto{
	//$result 开奖结果，从第一球到第七球
	//玩法：（对应数据库字段）共有大玩法19种
	//正码1-6不存入数据库 gameplay:168
	//过关根据正码计算，也不存入gameplay:169

	public static function get_auto($result){
		if(count($result) != 7 )return array();
		
		$tema = self::get_tema_str($result[6]); //特码
		$zm = self::get_zm_auto($result); //正码
		$zm_one_six = self::get_zm_str($result);//正码1-6 结果是数组
		//依次是大小 单双 合大小 合单双 尾大小 波色
		$zm_t = self::zm_t($result);//正码特

		$half_wave = self::get_harf_wave($result[6]);//半波
		$one_animal = self::get_one_animal($result); //一肖==正肖
		$tail = self::get_tail($result);//尾数 === 尾数连
		$tema_animal = self::tema_animal($result[6]);//特码生肖
		$he_animal = self::tema_animal($result[6],2);//合肖
		$lian_animal = $one_animal . ',' . $tema_animal; //生肖连
		$all_miss = implode(',', $result); //全不中 === 连码
		$wx = self::get_wx($result[6]); //五行
		$tema_head_tail = self::get_tema_head_tail($result[6]);//特码头尾
		$qima = self::get_qima($result); //七码
		$total_animal = self::get_total_animal($result);//总肖
		//组成数据库字段
		$values = array();
		$gameplay = Common_helper::getGameplay('liuhecai');
		if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'Tema':
					$values['ball_' . $val['id']] = $tema; //特码
				break;
				case 'JustCode':
					$values['ball_' . $val['id']] = $zm; //正码
				break;
				case 'JustCode_Te':
					$values['ball_' . $val['id']] = $zm_t;//正码特
				break;
				case 'JointMark':
				case 'AllMiss':
					$values['ball_' . $val['id']] = $all_miss; //全不中 === 连码
				break;
				case 'HalfWave':
					$values['ball_' . $val['id']] = $half_wave;//半波
				break;
				case 'EndNum_Even':
				case 'EndNum':
					$values['ball_' . $val['id']] = $tail;//尾数 === 尾数连
				break;
				case 'Te_Animal':
					$values['ball_' . $val['id']] = $tema_animal;//特码生肖
				break;
				case 'SumAnimal':
					$values['ball_' . $val['id']] = $he_animal;//合肖
				break;
				case 'AnimalEven':
					$values['ball_' . $val['id']] = $lian_animal; //生肖连
				break;
				case 'Five_elements':
					$values['ball_' . $val['id']] = $wx; //五行
				break;
				case 'Te_First_num':
					$values['ball_' . $val['id']] = $tema_head_tail;//特码头尾
				break;
				case 'Seven_code':
					$values['ball_' . $val['id']] = $qima; //七码
				break;
				case 'All_Animal':
					$values['ball_' . $val['id']] = $total_animal; //总肖
				break;
				case 'Animal':
				case 'Just_Animal':
					$values['ball_' . $val['id']] = $one_animal; //一肖==正肖
				break;

			}
		}
		//ball_1 - ball_6
		foreach($zm_one_six as $key=>$val){
			$values['ball_' . ($key + 1)] = $val;
		}
		$values['ball_7'] = $tema;
		return $values;
		

	}


/**
	  ***********************************************************
	  *  处理正码字符串              @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_zm_str($result){
		$new_arr = array();
		foreach($result as $key=>$val){
			$tmp_arr = array();
			if($key != 6){
				$tmp_arr = self::get_tema_auto($val);
				$new_arr[] = implode(',', $tmp_arr);
			}
		}
		return $new_arr;
	}
	
/**
	  ***********************************************************
	  *  处理特码字符串              @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_tema_str($tema){
		$tmp_arr = array();
		$tmp_arr = self::get_tema_auto($tema,true);
		$str = '';//拼接字段
		$str .= implode(',', $tmp_arr);
		return $str;
	}
/**
	  ***********************************************************
	  *  特码  正码1-6         	   @author ruizuo qiyongsheng   *
	  ***********************************************************
*/
	private static function get_tema_auto($ball,$tema = false){

		$res = array();
		if($ball < 0 || $ball > 49){
			$res[] = 'wrong';
			return $res;
		}
		$res[0] = $ball;
		if($ball == 49){
			$res[1] = 'sum';//和
		}else{
			//大小
			if($ball >= 25){
				$res[1] = 'big';
			}else{
				$res[1] = 'small';
			}

			//单双
			if($ball % 2 == 0){
				$res[2] = 'double';
			}else{
				$res[2] = 'single';
			}

			//合大小
			$one = $ball % 10;//个位数
			$ten = floor($ball/10);//十位数
			$he = $one + $ten;
			if($he >= 7){
				$res[3] = 'sum_big';
			}else{
				$res[3] = 'sum_small';
			}

			//合单双
			if($he % 2 == 0){
				$res[4] = 'sum_double';
			}else{
				$res[4] = 'sum_single';
			}

			//尾大小
			if($one >= 5){
				$res[5] = 'end_big';
			}else{
				$res[5] = 'end_small';
			}
		}

		//波色
		$tmp_wave = self::wave($ball);
		$res[6] = $tmp_wave . '_wave';
		//生肖
		$animal = self::get_animal();
		foreach($animal as $key=>$val){
			if(in_array($ball, $animal)){
				$res[6] = $key;
			}
		}
		
		if($tema){
			//大单大双 小单小双
				if($ball == 49){
					$res[7] = 'sum';
				}else{
					$res[7] = $res[1] . '_' . $res[2];
				}

			//家禽野兽
			 if(in_array($ball,array_merge($animal['cattle'],$animal['horse'],$animal['sheep'],$animal['chicken'],$animal['dog'],$animal['pig']))){

			 	$res[8] = 'wild_animals';
			 }else{
			 	$res[8] = 'Poultry';
			 }
		}
		
		return $res;
	}

/**
	  ***********************************************************
	  *  正码 和大小单双龙虎          @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_zm_auto($result){
		$res = '';
		$sum = array_sum($result);
		if($sum >= 100){
			$sum_end = $sum%100;
		}elseif($sum < 100){
			$sum_end = $sum%10;
		}
		if($sum >= 175){
			$res[0] = 'total_big';
		}else{
			$res[0] = 'total_small';
		}

		if($sum%2 == 0){
			$res[1] = 'total_double';
		}else{
			$res[1] = 'total_single';
		}

		if($result[0] > $result[6]){
			$res[2] = 'dragon';
		}else{
			$res[2] = 'tiger';
		}
		if($sum_end >= 5){
			$res[4] = 'total_end_big';
		}else{
			$res[4] = 'total_end_small';
		}
		unset($result[6]);
		$res[3] = implode(',', $result);
		return implode(',', $res);
	}
	
/**
	  ***********************************************************
	  *  正码特                    @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function zm_t($result){
		$str = '';
		foreach($result as $key => $val){
			if($key != 6){
				$str .= $val . ',';
			}
		}
		return rtrim($str,',');
	}
	
/**
	 ***********************************************************
		  *  半波           @author ruizuo qiyongsheng    *
	 ***********************************************************
*/
	private static function get_harf_wave($tema){
		$res = array();
		if($tema == 49){
			return 49;
		}
		$wave = self::wave($tema);
		$result = self::get_tema_auto($tema);
		$res[0] = $wave . '_' . $result[1];
		$res[1] = $wave . '_' . $result[2];
		$res[2] = $wave . '_' . $result[4];
		$str = implode(',', $res);
		return $str;
	}
/**
	  ***********************************************************
	  *  特码生肖   合肖        @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function tema_animal($tema,$type = 1){
		if($type == 2) return $tema;
		$animal = self::get_animal_str($tema);
		return $animal;
	}
/**
	  ***********************************************************
	  *  正码生肖           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_one_animal($result){
		unset($result[6]);
		$animal_str = '';
		foreach($result as $key=>$val){
			$tmp_str = '';
			$tmp_str = self::get_animal_str($val);
			$animal_str .= $tmp_str .  ',';
		}
		return rtrim($animal_str , ',');
	}

/**
	  ***********************************************************
	  *  尾数               @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_tail($result){
		$tail_str = '';
		foreach($result as $val){
			$tmp_str = $str =  0;
			$tmp_str = $val%10;
			$tail_str .= $tmp_str . ',';
		}

		return rtrim($tail_str , ',');
	}
	

/**
	  ***********************************************************
	  *  获取生肖           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_animal_str($ball){
		$animal_arr = self::get_animal();
		foreach($animal_arr as $animal=>$val){
			if(in_array($ball, $val)){
				return $animal;
			}
		}
	}
	
			
/**
	  ***********************************************************
	  *  波色        		    @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function wave($num){
		$ball = [];
        $ball['red'] =[1,2,7,8,12,13,18,19,23,24,29,30,34,35,40,45,46];
        $ball['blue'] =[3,4,9,10,14,15,20,25,26,31,36,37,41,42,47,48];
        $ball['green'] =[5,6,11,16,17,21,22,27,28,32,33,38,39,43,44,49];
        foreach($ball as $wave=>$val){
        	if(in_array($num, $val)){
        		return $wave;
        	}
        }
        return 'wrong';
	}
/**
	  ***********************************************************
	  *  五行             @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_wx($tema){
		$arr = array();
		$arr['metal'] = array(1,6,11,16,21,26,31,36,41,46);//金
        $arr['wood']= array(2,7,12,17,22,27,32,37,42,47);//木
        $arr['water'] = array(3,8,13,18,23,28,33,38,43,48);//水
        $arr['fire']= array(4,9,14,19,24,29,34,39,44,49);//火
        $arr['earth'] = array(5,10,15,20,25,30,35,40,45);//土

        foreach($arr as $key=>$val){
        	if(in_array($tema, $val)){
        		return $key;
        	}
        }

        return 'wrong';
	}
	
/**
	  ***********************************************************
	  *  特码头尾                  @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_tema_head_tail($tema){
		$head = floor($tema/10);
		$tail = $tema%10;
		switch($head){
			case '0':
				$head = 'First_num_zero';
			break;
			case '1':
				$head = 'First_num_one';
			break;
			case '2':
				$head = 'First_num_two';
			break;
			case '3':
				$head = 'First_num_three';
			break;
			case '4':
				$head = 'First_num_four';
			break;
		}
		switch($tail){
			case '0':
				$tail = 'Last_num_zero';
			break;
			case '1':
				$tail = 'Last_num_one';
			break;
			case '2':
				$tail = 'Last_num_two';
			break;
			case '3':
				$tail = 'Last_num_three';
			break;
			case '4':
				$tail = 'Last_num_four';
			break;
			case '5':
				$tail = 'Last_num_five';
			break;
			case '6':
				$tail = 'Last_num_six';
			break;
			case '7':
				$tail = 'Last_num_seven';
			break;
			case '8':
				$tail = 'Last_num_eight';
			break;
			case '9':
				$tail = 'Last_num_nine';
			break;
		}

		return $head . ',' . $tail;
	}

/**
	  ***********************************************************
	  *  七码           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_qima($result){
		$single = 0;
		$big = 0;
		foreach($result as $val){
			if($val >= 25){
				$big += 1;
			}
			if($val%2 ==1){
				$single += 1;
			}
		}
		$big_str = '';
		switch($big){
			case '0':
				$big_str = 'seven_small';
			break;
			case '1':
				$big_str = 'six_small';
			break;
			case '2':
				$big_str = 'five_small';
			break;
			case '3':
				$big_str = 'four_small';
			break;
			case '4':
				$big_str = 'three_small';
			break;
			case '5':
				$big_str = 'two_small';
			break;
			case '6':
				$big_str = 'one_small';
			break;
			case '7':
				$big_str = 'seven_big';
			break;
		}

		$single_str = '';
		switch($single){
			case '0':
				$single_str = 'seven_double';
			break;
			case '1':
				$single_str = 'six_double';
			break;
			case '2':
				$single_str = 'five_double';
			break;
			case '3':
				$single_str = 'four_double';
			break;
			case '4':
				$single_str = 'three_double';
			break;
			case '5':
				$single_str = 'two_double';
			break;
			case '6':
				$single_str = 'one_double';
			break;
			case '7':
				$single_str = 'seven_single';
			break;
		}

		return $big_str . ',' . $single_str;
	}
	
/**
	  ***********************************************************
	  *  总肖            @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_total_animal($result){
		$res = array();
		foreach($result as $val){
			$res[] = self::get_animal_str($val);
		}
		$res = array_unique($res);
		$count =  count($res);
		$str = '';
		if(in_array($count, [2,3,4])){
			$str = 'all_Animal_shunzi';//二三四肖
		}elseif($count == 5){
			$str = 'all_Animal_five';
		}elseif($count == 6){
			$str = 'all_Animal_six';
		}elseif($count == 7){
			$str = 'all_Animal_seven';//七肖
		}elseif($count == 1){
			$str = 'all_Animal_one';
		}else{
			$str = 'all_Animal_zero';
		}

		if($count % 2 == 0){
			$str .= ',all_Animal_double';
		}else{
			$str .= ',all_Animal_single';
		}

		return $str;

	}
	
	
//根据传入指定生肖获取相对应的数组  此数组在这里专用于处理六合彩的正肖玩法
    public static function get_animal(){
    	$animal_arr = array('mouse','cattle','tiger','rabbit','dragon','snake','horse','sheep','monkey','chicken','dog','pig');
        $year = date('Y');
        $year = $animal_arr[(($year-4)%12)];//获取本命年生肖
        $shenxiao_arr = array('pig','dog','chicken','monkey','sheep','horse','snake','dragon','rabbit','tiger','cattle','mouse');

        $num_arr  =  array(
                 array( 1, 13, 25, 37, 49 ),
                 array( 2, 14, 26, 38 ),
                 array( 3, 15, 27, 39 ),
                 array( 4, 16, 28, 40 ),
                 array( 5, 17, 29, 41 ),
                 array( 6, 18, 30, 42 ),
                 array( 7, 19, 31, 43 ),
                 array(8, 20, 32, 44 ),
                 array( 9, 21, 33, 45 ),
                 array( 10, 22, 34, 46 ),
                 array( 11, 23, 35, 47 ),
                 array( 12, 24, 36, 48 )
             );


        $key =  array_search($year,$shenxiao_arr);

        $begin_arr = array_splice($shenxiao_arr,$key);

        $end_arr = array_splice($shenxiao_arr,$key-12);

        $new_shenxiao_arr = array_merge($begin_arr,$end_arr);

        $shenxiao_num_arr = array();
        foreach($new_shenxiao_arr as $k=>$v){
            $shenxiao_num_arr[$v] = $num_arr[$k];
        }

        return $shenxiao_num_arr; //返回新的生肖数组
    }
}

