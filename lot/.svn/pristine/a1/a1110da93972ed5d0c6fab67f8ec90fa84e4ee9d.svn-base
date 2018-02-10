<?php
namespace libraries\auto;
use \helper\Common_helper;

class Cq_sscAuto{
	//$result 开奖结果，从第一球到第五球
	//玩法：（对应数据库字段）共有大玩法11种
	public static function get_auto($result){
		// 第一球 ~ 第五球,单双大小，斗牛，梭哈 总和龙虎 前中后三球
		if(count($result) != 5 )return array();
		$values = $ball_arr = self::get_ball_arr($result); //一到五球
		$tiger_ball = self::get_sum($result); //总和龙虎
		$cattle = self::get_cattle($result); //斗牛
		$sh = self::get_sh($result); //梭哈
		$before_three = self::get_three($result,1); //前三
		$mid_three = self::get_three($result,2); //中三
		$after_three = self::get_three($result,3); //后三

		//组成数据库字段
		$gameplay = Common_helper::getGameplay('cq_ssc');
        if(empty($gameplay)) return array();
		foreach($gameplay as $val){
			switch($val['gameplay']){
				case 'tiger_ball':
					$values['ball_' . $val['id']] = $tiger_ball;//总和龙虎
				break;
				case 'Bullfighting':
					$values['ball_' . $val['id']] = $cattle ;//斗牛
				break;
				case 'poker':
					$values['ball_' . $val['id']] = $sh ;//梭哈
				break;
				case 'before_three_ball':
					$values['ball_' . $val['id']] = $before_three ;//前三
				break;
				case 'middle_three_ball':
					$values['ball_' . $val['id']] = $mid_three ;//中三
				break;
				case 'after_three_ball':
					$values['ball_' . $val['id']] = $after_three ;//后三
				break;
			}
		}

		return $values;
	}

/**
	  ***********************************************************
	  *  单双大小           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	private static function get_ball_arr($result){
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
     *    总和大小 总和单双 龙 虎  和                    *
     ***********************************************************
*/
    private static function get_sum($result){
    	$sum = array_sum($result);
    	$res = array();
    	//总和大小
    	if($sum >= 23)
    		$res[0] = 'total_sum_big';
    	else
    		$res[0] = 'total_sum_small';
    	//总和单双
    	if($sum % 2 == 0)
    		$res[1] = 'total_sum_double';
    	else
    		$res[1] = 'total_sum_single';
    	//龙虎和
    	if($result[0] > $result[4]){
    		$res[2] = 'dragon';
    	}elseif($result[0] < $result[4]){
    		$res[2] = 'tiger';
    	}else{
    		$res[2] = 'sum';
    	}

    	return $sum . ',' .  implode(',', $res);
    }	
	
/**
      ***********************************************************
      * 计算斗牛 没牛-牛9-牛牛 牛大小单双                       *
      ***********************************************************
*/
    private static function get_cattle($result){
         $arr = $result;
         for($i = 0;$i < 5;$i++){
           
            for( $j = $i+1; $j<5; $j++ ){

                for( $k = $j + 1; $k < 5; $k++ ){
                    $zh = $arr[$i]+$arr[$j]+$arr[$k];
                    if($zh == 0 || $zh%10 == 0){
                        $zh2 = 0;
                        foreach($arr as $key=>$val){
                            if( ($key != $i)&&( $key != $j ) && ( $key != $k ) ){
                                $zh2 += $val;
                            }
                        }
                        break;
                    }

                }
             if($zh == 0 || $zh%10 == 0){break;}

            }
             if($zh == 0 || $zh%10 == 0){break;}
        }

        if(!isset($zh2)) return 'not_cow';//没牛
        while($zh2 >10){
            $zh2 -= 10;
        }
        //牛牛 牛1-牛8
        if(array_sum($result) == 0 || $zh2 == 0 || $zh2 == 10){
         	$res[0] = 'cow_cow';
        }
        if($zh2 == 1){
        	$res[0] = 'cow_one';
        }elseif($zh2 == '2'){
        	$res[0] = 'cow_two';
        }elseif($zh2 == '3'){
        	$res[0] = 'cow_three';
        }elseif($zh2 == '4'){
        	$res[0] = 'cow_four';
        }elseif($zh2 == '5'){
        	$res[0] = 'cow_five';
        }elseif($zh2 == '6'){
        	$res[0] = 'cow_six';
        }elseif($zh2 == '7'){
        	$res[0] = 'cow_seven';
        }elseif($zh2 == '8'){
        	$res[0] = 'cow_eight';
        }elseif($zh2 == '9'){
        	$res[0] = 'cow_nine';
        }
        //大小单双
        if($zh2 >= 6)
        	$res[1] = 'cow_big';
        else
        	$res[1] = 'cow_small';

        if($zh2 % 2 == 0)
        	$res[2] = 'cow_double';
        else
        	$res[2] = 'cow_single';

        return implode(',', $res);

    }

/**
      ***********************************************************
      * 计算梭哈（五条 四条 葫芦 三条 一对 两对 散号 顺子）  *
      ***********************************************************
*/
    private static function get_sh($result){
        $new_arr = $result;
        $new_arr = array_count_values($new_arr);
        if( in_array(5, $new_arr) ){
            return 'cow_five'; //五条
        }elseif( in_array(4, $new_arr) ){
            return 'cow_four'; //四条
        }elseif( in_array(3, $new_arr) ){
            $three = 0; 
            foreach($new_arr as $k=>$v){
                if($v != 3){
                    $three += 1;
                }
            }
            if($three == 1){
                return 'cow_cow'; //葫芦
            }else{
                return 'cow_small'; //三条
            }
        }elseif( in_array(2, $new_arr) ){
            $res = 0;
            foreach($new_arr as $key=>$val){
                if($val == 2){
                    $res += 1;
                }
            }
            if($res == 2){
                return 'cow_double'; //两对
            }
            if($res == 1){
                return 'cow_single'; //一对
            }
        }else{
            if($result[4] != 0){ $j =4; }else{ $j = 3;}
            $shun = 0;
            for( $i = 0 ; $i < $j; $i++){
                    if( ($result[$i]+1) == $result[$i+1] ){
                        $shun += 1;
                    }   
            }

            if($shun == $j){
                return 'cow_big'; //顺子
            }else{
                return 'cow_powder'; //散号
             }   
            
        }
    }

/**
      ***********************************************************
      * 计算前三中三后三 返回类型                               *
      ***********************************************************
*/
    private static function get_three($result,$type){
        //$result 开奖结果
        // 返回类型 1豹子 2对子 3顺子 4半顺 5杂‘6’
        if($type == 1){
            $num1 = $result[0];
            $num2 = $result[1];
            $num3 = $result[2];
        }elseif($type == 2){
            $num1 = $result[1];
            $num2 = $result[2];
            $num3 = $result[3];
        }elseif($type == 3){
            $num1 = $result[2];
            $num2 = $result[3];
            $num3 = $result[4];
        }

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
}