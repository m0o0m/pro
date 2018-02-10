<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Bj_kl8Balance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "bj_kl8";
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
            case 'choose_one'://选一
                 return static::check_xz($result,$info[1],1);
                 break;
            case 'choose_two'://选二
                return static::check_xz($result,$info[1],2);
                break;
            case 'choose_three'://选三
                return static::check_xz($result,$info[1],3);
                break;
            case 'choose_four'://选四
                return static::check_xz($result,$info[1],4);
                break;
            case 'choose_five'://选五
                return static::check_xz($result,$info[1],5);
                break;
           
            case 'sum':   //和值
                return static::check_sum($result,$info[1]); 
                break;
            case 'up_middle_down': //上中下盘
                return static::check_disc($result,$info[1]);
                break;
            case 'odd_and_even': //奇偶盘
                return static::check_odd_and_even($result,$info[1]);
                break;


        }
        return static::FAIL;
    }



/**
     ***********************************************************
     *  计算 和 总和大小 总和单双  和                          *
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
     *  计算 上盘 中盘 下盘                                    *
     ***********************************************************
*/
    private static function check_disc($result,$info){
        if($info == $result)
            return static::WIN;
        else
            return static::FAIL;
          
    }

/**
     ***********************************************************
     *  计算 奇盘 偶盘 和盘                                    *
     ***********************************************************
*/
     private static function check_odd_and_even($result,$info){
        if($info == $result)
            return static::WIN;
        else
            return static::FAIL;
     }

/**
      ***********************************************************
      *  选一 选二 选三 选四 选五                               *
      ***********************************************************
*/
    private static function check_xz($result,$info,$type){
        //$result 结果 //$info 投注数字 $type选X
        $result = explode(',', $result);
        $u_arr = explode(',', $info); //用户投注数字
        $u_arr = array_unique($u_arr); //去重复数字
        if( count($u_arr) != $type ){return static::FAIL;} //投注数字个数不对
        
        $res = 0; //储存投中数字
        foreach($u_arr as $val){
            if( in_array($val, $result) ){
                $res += 1;
            }
        }
        switch ($type) {
            case '1': //选一
                if( $res == 1)
                    return static::WIN;
                else
                    return static::FAIL;
            case '2': //选二
                if( $res == 2 )
                    return static::WIN;
                else
                    return static::FAIL;
                break;
            
            case '3': //选三
                if( $res == 3 ){ return static::WIN; } //三中三
                if( $res == 2 ){ return static::ONE; } //三中二
                return static::FAIL;
            break;

            case '4': //选四
                if( $res == 4 ){ return static::WIN; } //四中四
                if( $res == 3 ){ return static::ONE; } //四中三
                if( $res == 2 ){ return static::TWO; } //四中二
                return static::FAIL;
            break;

             case '5': //选五
                if( $res == 5 ){ return static::WIN; } //五中五
                if( $res == 4 ){ return static::ONE; } //五中四
                if( $res == 3 ){ return static::TWO; } //五中三
                return static::FAIL;
            break;

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