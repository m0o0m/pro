<?php
namespace libraries\balance;
use \helper\RedisConPool;
use \helper\Common_helper;
use \helper\GetQishuOpentime;
use \helper\MysqlPdo as pdo;
class Gd_11Balance{
	const WIN   = 2;
    const FAIL  = 3;
    const DRAW  = 4;
    const ONE   = 5;
	const TWO   = 6;
    static $type = "gd_11";
    static $count = [];      ///总用户数
    static $info = [];      ///用户输赢
    static $bets = [];      ///总注单
    static $run = [];        ///0待命 1读取就绪 2正在结算
    static $qishu       = "";               ///当前期数
    static $open        = 0;                ///当前开盘时间
    static $close       = 0;                ///封盘时间
    static $next_sec    = 540;        ///下期间隔


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
            case 'random_choose':   //任选
                return static::check_rx($result,$info);
                break;
            case 'vertical_choose': //直选
            case 'group_choose' :   //组选
                return static::check_zx($result,$info);
                break;

           

        }
        return static::FAIL;
    }


/**
     ***********************************************************
     * 计算 大 小 单 双                   *
     ***********************************************************
*/
    private static function check_one($res,$content){
        if($res == 11) return static::DRAW;
        $res = explode(',', $res);
        if(in_array($content, $res))
            return static::WIN;
        else
            return static::FAIL;
      
    }

/**
     ***********************************************************
     *  计算 和 和大 和小 和单 和双 龙 虎 尾大 尾小            *
     ***********************************************************
*/
    private static function check_sum($result,$info){
        $result = explode(',', $result);
        if($info == 'sum_big' || $info == 'sum_small'){
            if($result[0] == 30) return static::DRAW;  
        }
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
    private static function check_rx($result,$info){
        $result = explode(',', $result);
        //$info[0]:random_choose任选
        //$info[1]: X选X
        //$info[2]: 1,3,5,7 投注数字
        $num_arr = explode(',', $info[1]); 
        $arr_len = count($num_arr); //投注数字位数
        $type = '';   //每种玩法指定匹配次数
        switch ($info[2]) {
            case 'one_in_one': //一选一
                $type = 1;
                break;
            case 'two_in_two': //二选二
                $type = 2;
                break;
            case 'three_in_three': //三选三
                $type = 3;
                break;
            case 'four_in_four': //四选四
                $type = 4;
                break;
            case 'five_in_five': //五选五
                $type = 5;
                break;
            case 'six_in_five': //六选五
                $type = 6;
                break;
            case 'seven_in_five': //七选五
                $type = 7;
                break;
            case 'eight_in_five': //八选五
                $type = 8;
                break;
        }
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
        if($type <= 5){
            if($res != $type)
                return static::FAIL;
            else
                return static::WIN;
        }else{
            if($res != 5)
                return static::FAIL;
            else
                return static::WIN;   
        }



    }

/**
      ***********************************************************
      *  计算组选和直选                                         *
      ***********************************************************
*/
    private static function check_zx($result,$info){
        $result = explode(',', $result);
         //$info[0]:组选 直选
        //$info[1]: 前二 前三
        //$info[2]: 1,3,5,7 投注数字
        $num_arr = explode(',', $info[1]); 
        $arr_len = count($num_arr); //投注数字位数
        $type = '';     //指定玩法
        switch($info[2]){
            case 'before_two': //前二
                $type = 2;
            break;
            case 'before_three': //前三
                $type = 3;
            break;
        }
         //过滤投注数字中的重复值
        $num_arr = array_unique($num_arr);
        //判定投注数字位数是否合法
        if($arr_len != $type){
            return static::FAIL;
        }
        //指定匹配选号
        if($type == 2)$new_arr = array($result[0],$result[1]);
        if($type == 3)$new_arr = array($result[0],$result[1],$result[2]);

        $res = 0; //默认匹配结果
        //组选算法匹配  
        if($info[0] == 'group_choose'){    
            foreach($num_arr as $val){
                if(in_array(intval($val), $new_arr)){
                    $res += 1;
                }
            }
        }

        //直选算法匹配
        if($info[0] == 'vertical_choose'){
            $res = 0;
            for($i = 0; $i < $type ;$i++){
                if($num_arr[$i] == $new_arr[$i]){
                    $res += 1;
                }
            }
        }

        if($res != $type)
            return static::FAIL;
        else
            return static::WIN;
    
        
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