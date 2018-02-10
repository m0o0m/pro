<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Pc_28Balance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "pc_28";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 240;        ///下期间隔

/**
      ***********************************************************
      *  总和大小单双 总和大单 大双 小单 小双  极小 极大 红绿蓝波 豹子        *
      ***********************************************************
*/
    public static function Balance_bet($result,$info){
        if($info[0] != 'pc_28'){return static::FAIL;}
        $result = explode(',', $result);
        if($info[1] == 'Tema_in_Three'){//特码包三
            $bet_arr = explode(',', $info[2]);
            array_unique($bet_arr);
            if(count($bet_arr) != 3){return static::FAIL;}
            foreach($bet_arr as $val){
                if(in_array($val, $result))return static::WIN;
            }
            return static::FAIL;
        }
        if(in_array($info[1], $result))
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