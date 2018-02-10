<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class LiuhecaiBalance{
        const WIN   = 2;
        const FAIL  = 3;
        const DRAW  = 4;
        const ONE   = 5;        ////一单三种结果第一种意外
		const TWO   = 6;	
        static $type = "liuhecai";
        static $count = [];      ///总用户数
        static $info = [];      ///用户输赢
        static $bets = [];      ///总注单
        static $run = [];        ///0待命 1读取就绪 2正在结算
        static $qishu       = "";               ///当前期数
        static $open        = 0;                ///当前开盘时间
        static $close       = 0;                ///封盘时间
        static $next_sec    = 169200;        ///下期间隔

        public static function Balance_bet($result,$info){
            switch($info[0]){
                case 'Tema'://特码
                    return static::check_tema($info[1],$result);
                break;
                case 'JustCode'://正码
                    return static::check_zm($info[1],$result);
                break;
                case 'JustCode_Te'://正码特
                    return static::check_zhengmate($info,$result);
                break;
                case 'JustCode_one_six'://正码1-6
                    return static::check_zhengma($info,$result);
                break;
                case 'PassTest'://过关
                    return static::check_guoguan($info,$result);
                break;
                case 'JointMark'://连码
                    return static::check_lianma($info,$result);
                break;
                case 'HalfWave'://半波
                    return static::check_banbo($info,$result);
                break;
                case 'SumAnimal'://合肖
                    return static::check_hx($info,$result);
                break;
                case 'AnimalEven'://生肖连
                    return static::check_lian($info,$result);
                break;
                case 'EndNum_Even'://尾数连
                    return static::check_weishus($info,$result);
                break;
                
                case 'AllMiss'://全不中
                    return static::check_buzhong($info,$result);
                break;
                case 'Te_Animal'://特肖
                case 'Animal'://一肖（生肖)
                case 'EndNum'://尾数
                case 'Five_elements': //五行
                case 'Just_Animal'://正肖
                case 'Te_First_num'://特码头
                case 'Te_Last_num'://特码尾
                case 'Seven_code'://七码 下注时:英文，one_small,two_double
                case 'All_Animal': //总肖
                    return static::check_zm($info[1],$result);
                break;

                default:
                    return static::FAIL;
                break;

            }
        }
        
/**
      ********************{***************************************
      *  特码 正码1-6          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
       private static function check_tema($info,$result){ 
            $result = explode(',', $result);
            $arr = ['Poultry','wild_animals','red_wave','green_wave','blue_wave'];//家禽野兽 红绿蓝波
            $arr2 = ['1-10','11-20','21-30','31-40','41-49'];
            if(in_array($info, $arr) && in_array($info, $result)){
                return static::WIN;
            }elseif(in_array($info, $arr2)){
                $multiple = static::multiple();
                foreach($arr2 as $key=>$val){
                    if($info == $val){
                        if(in_array($result[0],$multiple[$key])){
                             return static::WIN;
                        }
                    }
                }
                return static::FAIL;
            //除了上面的玩法，剩余的玩法：开出49为和
            }elseif(in_array($info, $result)){
                return static::WIN;
            }elseif(in_array('sum', $result)){
                return static::DRAW;
            }else{
                return static::FAIL;
            }

       }
/**
      ***********************************************************
      *  正码           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    private static function check_zm($info,$result){
        $result = explode(',', $result);
        if($result[0] == 'sum') return static::DRAW;
        if(in_array($info, $result))
            return static::WIN;
        else
            return static::FAIL;
    }
    //过关
    private static function check_guoguan($info,$result){
        $info2 = explode(',',$info[1]); //big small一类的
        $list = explode(',',$info[2]); //正码一 正码二
        $num = count($info2);
        if($num != count($list)){
            return static::FAIL;
        }else{
            for($i=0;$i<$num;$i++){
                $data = array();
                $data[0] = 0;
                $data[1] = $info2[$i]; //大小等
                $data[2] = $list[$i]; //正码英文数字
                if(static::check_zhengma( $data,$result ) != static::WIN){
                    return static::FAIL;
                }
            }
            return static::WIN;
        }
    }

    //正码1-6   
    private static function check_zhengma($info,$result){
        switch($info[2]){
            case 'JustCode_one'://正码1
                return static::check_tema($info[1],$result[0]);
            break;
            case 'JustCode_two'://正码2
                return static::check_tema($info[1],$result[1]);
            break;
            case 'JustCode_three'://正码3
                return static::check_tema($info[1],$result[2]);
            break;
            case 'JustCode_four'://正码4
                return static::check_tema($info[1],$result[3]);
            break;
            case 'JustCode_five'://正码5
                return static::check_tema($info[1],$result[4]);
            break;
            case 'JustCode_six'://正码6
                return static::check_tema($info[1],$result[5]);
            break;
            default:
                return static::FAIL;
            break;
        }
    }

    //正码特
    private static function check_zhengmate($info,$result){
        $result = explode(',', $result);
        switch($info[2]){
            case 'JustCode_Te_one'://正1特
                if($info[1] == $result[0]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            case 'JustCode_Te_two'://正2特
                if($info[1] == $result[1]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            case 'JustCode_Te_three'://正3特
                if($info[1] == $result[2]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            case 'JustCode_Te_four'://正4特
                if($info[1] == $result[3]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            case 'JustCode_Te_five'://正5特
                if($info[1] == $result[4]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            case 'JustCode_Te_six'://正6特
                if($info[1] == $result[5]){
                    return static::WIN;
                }else{
                    return static::FAIL;
                }
            break;
            default:
                return static::FAIL;
            break;
        }
    }    
   //连码
   private static function check_lianma($info,$result){
        $result = explode(',', $result);
        $infos = explode(',',$info[1]);
        $tema = $result[6];
        unset($result[6]);
        $num = 0;
        $count = count($infos);
        foreach($infos as $val){
            if(in_array($val, $result))$num += 1;
        }
        $tema_num = 0;
        if(in_array($tema, $infos)) $tema_num = $num + 1;
        switch($info[2]){
            case 'second_full'://二全中
                if($count != 2) return static::FAIL;
                if($num == 2)
                    return static::WIN;
                else
                    return static::FAIL;
            break;
            
            case 'Special_series'://特串
                if($count != 2) return static::FAIL;
                if($tema_num == 2 && in_array($tema, $infos))
                    return static::WIN;
                else
                    return static::FAIL;
            break;
            
            case 'third_full'://三全中
                if($count != 3) return static::FAIL;
                if($num == 3)
                    return static::WIN;
                else
                    return static::FAIL;
            break;
            case 'fourth_full'://四全中
                if($count != 4) return static::FAIL;
                if($num == 4)
                    return static::WIN;
                else
                    return static::FAIL;
            break;
        }
        
        $infos = explode(',',$info[1]);
      
        if(count($infos) == 2){//二中二
             if(stripos($info[2],'/')){
                if(in_array($infos[0],$result) && in_array($infos[1],$result)){         ///中二
                    return static::ONE;
                }else{                                                                  ///中特
                    if((in_array($infos[0],$result) && $infos[1] == $tema) || (in_array($infos[1],$result) && $infos[0] == $tema)){
                        return static::WIN;
                    }else{
                        return static::FAIL;
                    }
                }

                return static::FAIL;
            }
        }

        if(count($infos) == 3){//三中二
            if(stripos($info[2],'/')){
                if(in_array($infos[0],$result) && in_array($infos[1],$result) && in_array($infos[2],$result)){//中三
                    return static::WIN;
                }else{
                    if((in_array($infos[0],$result) && in_array($infos[1],$result)) || (in_array($infos[0],$result) && in_array($infos[2],$result)) || (in_array($infos[1],$result) && in_array($infos[2],$result))){//中二
                        return static::ONE;
                    }
                    return static::FAIL;
                }

                return static::FAIL;
            }
        }
        
        return static::FAIL;
    }
    //半波
    private static function check_banbo($info,$result){
        if($result == 49) return static::DRAW;
        $result = explode(',', $result);
        if(in_array($info[1], $result))
            return static::WIN;
        else
            return static::FAIL;
    }
    //合肖
    public static function check_hx($info,$result){
        $animals = $info[1];
        $num = 0;
        switch($info[2]){
            case 'two_Animal':
                $num = 2;
            break;
            case 'three_Animal':
                $num = 3;
            break;
            case 'four_Animal':
                $num = 4;
            break;
            case 'five_Animal':
                $num = 5;
            break;
            case 'six_Animal':
                $num = 6;
            break;
            case 'seven_Animal':
                $num = 7;
            break;
            case 'eight_Animal':
                $num = 8;
            break;
            case 'nine_Animal':
                $num = 9;
            break;
            case 'ten_Animal':
                $num = 10;
            break;
            case 'elven_Animal':
                $num = 11;
            break;
        }
        $animals = explode(',',$animals);
        if(count($animals) != $num){
            return static::FAIL;    
        }

        if($result == 49){
            return static::DRAW;
        }
        $year_animal = self::get_year_animal(); //本命年生肖
        $animal_list = static::get_animal_arr($year_animal);
        $winer = [];
       
        foreach($animals as $v){
            $tmp = isset($animal_list[$v]) ? $animal_list[$v] : array();
            $winer = array_merge($winer,$tmp);
        }
        if(in_array($result,$winer)){
            return static::WIN;
        }else{
            return static::FAIL;
        }
    }
    //生肖连
    private static function check_lian($info,$result){
            $res_animal = explode(',', $result);
            switch($info[2]){
                case 'two_Animal_in'://二肖连中
                case 'three_Animal_in'://三肖连中
                case 'four_Animal_in'://四肖连中
                case 'five_Animal_in'://五肖连中
                    $infos = explode(',',$info[1]);
                    $num = 0;
                    foreach($infos as $val){
                        if(in_array($val,$res_animal)){
                            $num++;
                        }
                    }
                    if($num < 2){
                        return static::FAIL;
                    }
                    if($num == count($infos)){
                        return static::WIN;
                    }else{
                        return static::FAIL;
                    }
                break;
                case 'two_Animal_not_in'://二肖连不中
                case 'three_Animal_not_in'://三肖连不中
                case 'four_Animal_not_in'://四肖连不中
                    $infos = explode(',',$info[1]);
                    if(count($infos)<2){
                        return static::FAIL;
                    }
                    foreach($infos as $val){
                        if(in_array($val,$res_animal)){
                            return static::FAIL;
                        }
                    }
                    return static::WIN;
                break;

            }
            return static::FAIL;
    }
    //尾数连
    private static function check_weishus($info,$result){
            $res_wei = explode(',', $result);

            switch ($info[2]) {
                case 'two_end_in'://二尾连中
                case 'three_end_in'://三尾连中
                case 'four_end_in'://四尾连中
                    $infos = explode(',',$info[1]);
                    $num = 0;
                    foreach($infos as $val){
                        if(in_array($val,$res_wei)){
                            $num++;
                        }
                    }
                    if($num == count($infos)){
                        return static::WIN;
                    }else{
                        return static::FAIL;
                    }
                    break;
                case 'two_end_not_in'://二尾连不中
                case 'three_end_not_in'://三尾连不中
                case 'four_end_not_in'://四尾连不中
                    $infos = explode(',',$info[1]);
                    foreach($infos as $val){
                        if(in_array($val,$res_wei)){
                            return static::FAIL;
                        }
                    }
                    return static::WIN;
                    break;
            }
            return static::FAIL;
    }
    //全不中
    private static function check_buzhong($info,$result){
            switch($info[2]){
                case 'five_not_in':
                    $num = 5;
                break;
                case 'six_not_in':
                    $num = 6;
                break;
                case 'seven_not_in':
                    $num = 7;
                break;
                case 'eight_not_in':
                    $num = 8;
                break;
                case 'nine_not_in':
                    $num = 9;
                break;
                case 'ten_not_in':
                    $num = 10;
                break;
                case 'elven_not_in':
                    $num = 11;
                break;
                case 'twelve_not_in':
                    $num = 12;
                break;
                default:
                    return static::FAIL;
                break;
            }
            $check = explode(',',$info[1]);
            if(is_array($result) || empty($result)) return static::FAIL;
            $result = explode(',', $result);
            foreach($result as $key=> $val){
                if(!is_numeric($val)) return static::FAIL;
            }
            if(count($check) != $num) return static::FAIL;
            foreach($check as $val){
                if(in_array($val,$result))
                    return static::FAIL;
            }
            return static::WIN;
    }

    ///综合玩法
    private static function multiple(){
        $multiple = [];
        $multiple[0] = [1,2,3,4,5,6,7,8,9,10];                ///1-10
        $multiple[1] = [11,12,13,14,15,16,17,18,19,20];       ///11-20
        $multiple[2] = [21,22,23,24,25,26,27,28,29,30];       ///21-30
        $multiple[3] = [31,32,33,34,35,36,37,38,39,40];       ///31-40
        $multiple[4] = [41,42,43,44,45,46,47,48,49];          ///41-49
        
        return $multiple;
    }
    //根据传入指定生肖获取相对应的数组
    private static function get_animal_arr($year){

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
/**
      ***********************************************************
      *  获取本命年生肖           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function get_year_animal(){
        //生肖数组
        $animal_arr = array('mouse','cattle','tiger','rabbit','dragon','snake','horse','sheep','monkey','chicken','dog','pig');
        $year = date('Y');
        $year_animal = $animal_arr[(($year-4)%12)];//获取本命年生肖
        return $year_animal;
    }
    
    
/**
      ***********************************************************
      *  获取期数                                               *
      ***********************************************************
*/
    public static function get_qishu(){
       return GetQishuOpentime::get_qishu(static::$type);
    }

 /**
       ***********************************************************
       *  获取期数起始时间         输入期数  返回下注美东时间    *
       ***********************************************************
 */
    public static function getTimeByQishu($qishu){
       return GetQishuOpentime::getOpentime(static::$type, $qishu);
    }
    

}