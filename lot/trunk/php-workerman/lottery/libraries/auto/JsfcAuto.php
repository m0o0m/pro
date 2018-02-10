<?php

namespace libraries\auto;

use \libraries\auto\Bj_10Auto;

/**
* 极速飞车 开奖结果 计算
*/
class JsfcAuto{
    public static function get_auto($auto)
    {
        return Bj_10Auto::get_auto($auto);
       
    }
}