<?php
namespace libraries\balance;
use \helper\GetQishuOpentime;
class Bj_10Balance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "bj_10";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 240;        ///下期间隔
    
    public static function Balance_bet($result,$info){
        switch($info[0]){
            case 'first_second_sum'://冠亚军和
                return static::check_sum($result,$info[1]);
            break;
            case 'first'://冠军
                return static::check_one($result[0],$info[1]);
            break;
            case 'second'://亚军
                return static::check_one($result[1],$info[1]);
            break;
            case 'third'://第三名
                return static::check_one($result[2],$info[1]);
            break;
            case 'fourth'://第四名
                return static::check_one($result[3],$info[1]);
            break;
            case 'fifth'://第五名
                return static::check_one($result[4],$info[1]);
            break;
            case 'sixth'://第六名
                return static::check_one($result[5],$info[1]);
            break;
            case 'seventh'://第七名
                return static::check_one($result[6],$info[1]);
            break;
            case 'eighth'://第八名
                return static::check_one($result[7],$info[1]);
            break;
            case 'ninth'://第九名
                return static::check_one($result[8],$info[1]);
            break;
            case 'tenth'://第十名
                return static::check_one($result[9],$info[1]);
            break;
            // case 'dragon_tiger'://龍虎
            //     return static::check_lh($result,$result);
            // break;
        }
        return static::FAIL;
    }
    
    ////计算1-10名 $res开奖号 $content明细2 $type比较龙虎的开奖结果
    private static function check_one($result,$num){
        $result = explode(',', $result);
        if(in_array($num, $result)){
            return static::WIN;
        }else{
            return static::FAIL;
        }
    }
    
    
    private static function check_sum($result,$num){
        $result = explode(',', $result);
        if(in_array($num, $result)){
            return static::WIN;
        }else{
            return static::FAIL;
        }
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


?>