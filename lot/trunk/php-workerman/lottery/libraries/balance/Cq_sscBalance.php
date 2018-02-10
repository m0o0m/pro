<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Cq_sscBalance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "cq_ssc";
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
           
            case 'tiger_ball':   //总和龙虎
                return static::check_sum($result,$info[1]); 
                break;

            case 'before_three_ball': //前三球
            case 'middle_three_ball': //中三球
            case 'after_three_ball': //后三球
                    return static::check_three($result,$info[1]);
                break; 

            case 'Bullfighting':   //斗牛
                return static::check_dn($result,$info[1]);
                break;
            case 'poker': //梭哈
                return static::check_sh($result,$info[1]);
                break;
           

        }
        return static::FAIL;
    }


/**
     ***********************************************************
     * 计算 大 小 单 双                                        *
     ***********************************************************
*/
    private static function check_one($res,$content){
        $res = explode(',', $res);
        if(in_array($content, $res))
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
        $result = explode(',', $result);
        if(in_array($info, $result))
            return static::WIN;
        else
            return static::FAIL;
    }

/**
      ***********************************************************
      * 计算前三中三后三 返回类型                               *
      ***********************************************************
*/
    private static function check_three($result,$info){
        $result = explode(',', $result);
        if(in_array($info, $result))
            return static::WIN;
        else
            return static::FAIL;
    }


/**
      ***********************************************************
      * 计算斗牛 没牛-牛9-牛牛 牛大小单双                       *
      ***********************************************************
*/
    private static function check_dn($result,$info){
        $result = explode(',', $result);
        if(in_array($info, $result))
            return static::WIN;
        else
            return static::FAIL;
            
    }
 
 /**
      ***********************************************************
      * 计算梭哈（五条 四条 葫芦 三条 一对 两对 散号 顺子）  *
      ***********************************************************
*/
    private static function check_sh($result,$info){
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