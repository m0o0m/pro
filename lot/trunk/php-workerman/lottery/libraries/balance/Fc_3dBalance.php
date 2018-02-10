<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Fc_3dBalance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "fc_3d";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 85320;        ///下期间隔


    public static function Balance_bet($result,$info){

        switch($info[0]){
            case 'first_ball'://第一球
                 return static::check_one($result,$info[1]);
                 break;
            case 'second_ball'://第二球
                 return static::check_one($result,$info[1]);
                 break;
            case 'third_ball'://第三球
                 return static::check_one($result,$info[1]);
                 break;
            case 'tiger_ball'://总和龙虎
                return static::check_sum($result,$info[1]);
                break;
            case 'triple_ball'://三连
                    return static::check_three($result,$info[1]);
                break;
            case 'gallbladder_ball': //独胆
                    return static::dd($result,$info[1]);
                break;
            case 'span_ball': //跨度
                    return static::check_span($result,$info[1]);
                break;
        }
        return static::FAIL;
    }

/**
     ***********************************************************
     * 计算 大 小 单 双                                        *
     ***********************************************************
*/
    private static function check_one($result,$info){
       $res = explode(',', $result);
       if(in_array($info, $res))
            return static::WIN;
        else
            return static::FAIL;
    }

/**
     ***********************************************************
     *  计算 和 总和大小 总和单双 龙 虎  和                    *
     ***********************************************************
*/
    private static function check_sum($result,$info){
       $res = explode(',', $result);
       if(in_array($info, $res))
            return static::WIN;
        else
            return static::FAIL;
    
    }
    
/**
      ***********************************************************
      * 计算三连                                  *
      ***********************************************************
*/
    private static function check_three($result,$info){
        if($info == $result)
            return static::WIN;
        else
            return static::FAIL;
        
    }    

                  
/**
      ***********************************************************
      *  计算跨度                                               *
      ***********************************************************
*/
    private static function check_span($result,$info){
        if($info == $result)
            return static::WIN;
        else
            return static::FAIL;
    }
/**
      ***********************************************************
      *  计算独胆                                               *
      ***********************************************************
*/
    private static function dd($result,$info){
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
    
}