<?php
namespace libraries\balance;
use \helper\GetQishuOpentime;
class Ffc_oBalance{
	  const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	  const TWO   = 6;
    static $type = "ffc_o";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 60;        ///下期间隔

 
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
            case 'total_sum':   //总和
                return static::check_sum($result,$info[1]); 
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