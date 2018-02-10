<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Gd_tenBalance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "gd_ten";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 480;        ///下期间隔

 
    public static function Balance_bet($result,$info){
        switch($info[0]){
            case 'first_ball'://第一球
                 return static::check_one($result[0],$info[1]);
                 break;
            case 'second_ball'://第二球
                return static::check_one($result[1],$info[1]);
                break;
            case 'third_ball'://第三球
                return static::check_one($result[2],$info[1]);
                break;
            case 'fourth_ball'://第四球
                return static::check_one($result[3],$info[1]);
                break;
            case 'fifth_ball'://第五球
                return static::check_one($result[4],$info[1]);
                break;
            case 'six_ball'://第六球
                return static::check_one($result[5],$info[1]);
                break;
            case 'seven_ball'://第七球
                return static::check_one($result[6],$info[1]);
                break;
            case 'eight_ball'://第八球
                return static::check_one($result[7],$info[1]);
                break;    
            case 'tiger_ball':   //总和
                return static::check_sum($result,$info[1]); 
                break;
            case 'random_choose_two':   //任选二
                return static::check_rx($result,$info[1],2);
                break;
            case 'random_choose_three': //任选三
                return static::check_rx($result,$info[1],3);
                break;
            case 'random_choose_four':  //任选四
                return static::check_rx($result,$info[1],4);
                break;
            case 'random_choose_five':  //任选五
                return static::check_rx($result,$info[1],5);
                break;
            case 'random_choose_two_group': //任选两组
                return static::check_zx($result,$info[1]);
                break;

           

        }
        return static::FAIL;
    }


/**
     ***********************************************************
     * 计算 大 小 单 双  合数单双 尾大尾小 东南西北中发白      *
     ***********************************************************
*/
    private static function check_one($res,$content){
        $result = explode(',', $res);
        if(in_array($content, $result))
            return static::WIN;
        else
            return static::FAIL;
    }

/**
     ***********************************************************
     *  计算 和 总和大小 总和单双 龙 虎 总和尾大尾小           *
     ***********************************************************
*/
    private static function check_sum($result,$info){
        $result = explode(',', $result);
        if(in_array($info, $result))
            return static::WIN;
        else
            return static::FAIL;
    }
/**
      ***********************************************************
      *  计算任选                                               *
      ***********************************************************
*/
    private static function check_rx($result,$info,$type){
        $result = explode(',', $result);
        //$result 开奖结果
        //$info 投注数字字符串
        //$type 任选数字个数
        $num_arr = explode(',', $info); 
        $arr_len = count($num_arr); //投注数字位数
        //过滤投注数字中的重复值
        $num_arr = array_unique($num_arr);
        //判定投注数字位数是否合法
        if($arr_len != $type){
            return static::FAIL;
        }
        //统计匹配次数
        $res = 0;
        foreach($num_arr as $val){
            if(in_array(intval($val), $result)){
                $res += 1;
            }
        }
        if($res != $type){
            return static::FAIL;
        }else{
            return static::WIN;
        }
    }

/**
      ***********************************************************
      *  计算任选两组                                           *
      ***********************************************************
*/
    private static function check_zx($result,$info){
        // $result 开奖结果
        // $info 投注数字字符串信息
        $result = explode(',', $result);
        $num_arr = explode(',', $info); 
        $arr_len = count($num_arr); //投注数字位数
         //过滤投注数字中的重复值
        $num_arr = array_unique($num_arr);
        //判定投注数字位数是否合法
        if($arr_len != 2){
            return static::FAIL;
        }
        
        //任选两组算法
        $num1 = intval($num_arr[0]); //投注第一位数字  
        $num2 = intval($num_arr[1]); //投注第二位数字
        $res_count = count($result) - 1;
        foreach($result as $key=>$val){
            if($val == $num1){
                if($key == 0){
                    if($result[$key + 1] == $num2)return static::WIN;
                }elseif($key == $res_count){
                    if($result[$key - 1] == $num2)return static::WIN;
                }else{
                     if($result[$key + 1] == $num2 || $result[$key - 1] == $num2 ){
                     return static::WIN;
                     }
                }
            }
        }

        return static::FAIL;
    
        
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
    
/**
      ***********************************************************
      *  当期数增加或减少到尽头时返回正确的期数                 *
      ***********************************************************
*/
    public static function jiajian($qishu,$jiajian = 1){
        return GetQishuOpentime::new_qishu(static::$type, $qishu, $jiajian);    
    }   
}