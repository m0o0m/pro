<?php

namespace libraries\auto;

use \helper\Common_helper;

/**
* 北京PK拾 开奖结果 计算
*/
class Bj_10Auto{

    private static $js;
    private static $max=10;
    private static $min=1;

    public static function get_auto(array $auto) :array
    {
        $balls = self::js_balls($auto);
        $first_second_sum = self::js_first_second_sum($auto);

        $gameplay = Common_helper::getGameplay('bj_10');
        if(empty($gameplay)) return array();
        foreach($gameplay as $val){
            switch($val['gameplay']){
                case 'first_second_sum':
                    $balls['ball_' . $val['id']] = $first_second_sum;//总和
                break;
            }
        }
        return $balls;
    }

    private static function js_balls(array $auto) :array
    {
        $dragon_tiger = self::dragon_tiger($auto, false);
        foreach($auto as $key => $num)
        {
            $ball = 'ball_' . ($key + 1);

            $result[$ball][] = $num;
            $result[$ball][] = self::big_small($num);
            $result[$ball][] = self::single_double($num);

            if( key_exists($key, $dragon_tiger) )
            {
                $result[$ball][] = $dragon_tiger[$key];
            }

            $result[$ball] = implode(',', $result[$ball]);
        }

        return $result;
    }

    private static function js_first_second_sum(array $auto) :string
    {
        $result[] = self::first_second_sum($auto);
        $result[] = self::first_second_sum_big_small($auto);
        $result[] = self::first_second_sum_single_double($auto);

        return implode(',', $result);
    }

    /**
    * 冠亚和 大小
    */
    private static function first_second_sum_big_small(array $auto) :string
    {
        $first_second_num = self::first_second_sum($auto);
        $first_second_mid = 11;

        $result = self::sum_big_small($first_second_num, $first_second_mid, true);

        return 'first_second_' . $result;
    }

    /**
    * 大小
    */
    public static function big_small(int $num) :string
    {
        $mid = ( self::$max - self::$min ) / 2 + self::$min;
        if( $num >= $mid )
        {
            $result = 'big';
        }
        else
        {
            $result = 'small';
        }
        return $result;
    }

    /**
    * 单双
    */
    public static function single_double(int $num) :string
    {
        if( $num %2 == 0 )
        {
            $result = 'double';
        }
        else
        {
            $result = 'single';
        }
        return $result;
    }

    /**
    * 龙 虎 平局
    */
    public static function dragon_tiger(array $data, bool $single = true, &$result = []) /* :string|:array */
    {
        $first = array_shift($data);
        $last = array_pop($data);

        if ($first > $last)
        {
            $result[] = 'dragon';
        }
        elseif ($first < $last)
        {
            $result[] = 'tiger';
        }
        else
        {
            $result[] = 'draw';
        }

        if (count($data) >= 2)
        {
            self::dragon_tiger($data, true, $result);
        }

        return $single ? $result[0] : $result;
    }

    /**
    * 和 大小
    */
    public static function sum_big_small(int $num, int $mid, bool $draw = false) :string
    {
        if( $num >= $mid )
        {
            $result = 'big';
            if( $num == $mid && $draw === true) $result = 'draw';
        }
        else
        {
            $result = 'small';
        }
        return $result;
    }

    /**
    * 冠亚和
    */
    public static function first_second_sum(array $auto) :string
    {
        list($first, $second) = $auto;

        $result = $first + $second;

        return $result;
    }

    /**
    * 冠亚和 单双
    */
    public static function first_second_sum_single_double(array $auto) :string
    {
        $first_second_num = self::first_second_sum($auto);

        $result = self::single_double($first_second_num);

        return 'first_second_' . $result;
    }
}