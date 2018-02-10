<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;

class Ah_k3Balance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "ah_k3";
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
            case 'sum'://和值
                return static::check_sum($result,$info[1]);
            break;
            case 'two_even'://两连
                return static::check_two_even($result,$info[1]);
            break;
            case 'gallbladder_ball'://独胆
                return static::check_gallbladder_ball($result,$info[1]);
            break;
            case 'leopard'://豹子
            case 'pairs'://对子
                return static::check_leopard($result,$info[1]);
            break;
        }
        return static::FAIL;
    }

    //豹子对子
    private static function check_leopard($result,$play){
        if($play == $result)
            return static::WIN;
        else
            return static::FAIL;
    }

    //独胆
    private static function check_gallbladder_ball($result,$content){
        $result = explode(',', $result);
        if(in_array($content,$result)){
            return static::WIN;
        }
        return static::FAIL;
    }

    //两连
    private static function check_two_even($result,$content){
        $result = explode(',', $result);
        $cbetball = explode(',',$content);
        if(count($cbetball) != 2) return static::FAIL;
        $i = 0;
        foreach ($cbetball as $v) {
            if(in_array($v, $result)){
                $i++;
            }
        }
        if($i >= 2){
            return static::WIN;
        }
        return static::FAIL;
    }

    //和值玩法
    private static function check_sum($result,$play){
        if($result == '') return static::FAIL;
        $result = explode(',', $result);
        if(in_array($play, $result))
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


?>