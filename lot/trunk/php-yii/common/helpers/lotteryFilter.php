<?php

namespace common\helpers;
/*
 * 是否是合法注单
 *
 */

class lotteryFilter {

    //福彩3D 排列3
    public static function fc_3d($betInfo) {

        switch ($betInfo['mingxi_1']) {
            case 'first_ball':
            case 'second_ball':
            case 'third_ball':
                if (!in_array($betInfo['mingxi_2'], [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 'big', 'small', 'single', 'double'])) {
                    return false;
                }
                break;
            case'span_ball':
            case'gallbladder_ball':
                if (!in_array($betInfo['mingxi_2'], [0, 1, 2, 3, 4, 5, 6, 7, 8, 9])) {
                    return false;
                }
                break;
            case 'triple_ball':
                if (!in_array($betInfo['mingxi_2'], ['pairs', 'Half_suitable', 'leopard', 'straight', 'Miscellaneous_six'])) {
                    return false;
                }
                break;
            case 'tiger_ball':
                if (!in_array($betInfo['mingxi_2'], ['total_sum_big', 'total_sum_small', 'total_sum_single', 'total_sum_double', 'dragon', 'tiger', 'sum'])) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }


    //六合彩
    public static function liuhecai($betInfo) {
        $betres = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, '1-10', '11-20', '21-30', '31-40', '41-49', 'big', 'wild_animals', 'small_double', 'sum_double', 'double', 'blue_wave', 'Poultry', 'big_double', 'sum_single', 'single', 'green_wave', 'end_small', 'small_single', 'sum_small', 'small', 'red_wave', 'end_big', 'big_single', 'sum_big', 'big');
        $betres_z = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 'total_single', 'total_double', 'total_big', 'total_small', 'total_end_big', 'total_end_small', 'dragon', 'tiger');
        $betres_t = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 'JustCode_Te_one', 'JustCode_Te_two', 'JustCode_Te_three', 'JustCode_Te_four', 'JustCode_Te_five', 'JustCode_Te_six');
        $betres_z_1 = array('big', 'small', 'single', 'double', 'red_wave', 'green_wave', 'blue_wave', 'sum_big', 'sum_small', 'sum_single', 'sum_double', 'end_big', 'end_small', 'JustCode_one', 'JustCode_two', 'JustCode_three', 'JustCode_four', 'JustCode_five', 'JustCode_six');
        switch ($betInfo['mingxi_1']) {
            case 'Tema':
                if (!in_array($betInfo['mingxi_2'], $betres)) {
                    return false;
                }
                break;
            case 'JustCode':
                if (!in_array($betInfo['mingxi_2'], $betres_z)) {
                    return false;
                }
                break;
            case 'JustCode_Te':
                if (!in_array($betInfo['mingxi_2'], $betres_t) || !in_array($betInfo['mingxi_3'], $betres_t)) {
                    return false;
                }
                break;
            case 'JustCode_one_six':
                if (!in_array($betInfo['mingxi_2'], $betres_z_1) || !in_array($betInfo['mingxi_3'], $betres_z_1)) {
                    return false;
                }
                break;
            case 'PassTest':
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                $betres_2 = explode(',', $betInfo['mingxi_3']);
                $betres_2 = array_unique($betres_2);
                $arr = ['single', 'double', 'big', 'red_wave', 'blue_wave', 'green_wave','small'];
                $arr_2 = ['JustCode_one', 'JustCode_two', 'JustCode_three', 'JustCode_four', 'JustCode_five', 'JustCode_six'];
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                foreach ($betres_2 as $value) {
                    if (!in_array($value, $arr_2)) {
                        return false;
                    }
                }
                break;
            case 'JointMark':
                $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49];
                // $arr_2 = array('second_full', 'Special_series', 'third_full', 'fourth_full','in_te','in_two','in_two_in_three', 'in_three');
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                break;
            case 'HalfWave':
                $arr = array('red_single', 'red_double', 'red_big', 'red_small', 'red_sum_single', 'red_sum_double', 'green_single', 'green_double', 'green_big', 'green_small', 'green_sum_single', 'green_sum_double', 'blue_single', 'blue_double', 'blue_big', 'blue_small', 'blue_sum_single', 'blue_sum_double');
                if (!in_array($betInfo['mingxi_2'], $arr)) {
                    return false;
                }
                break;
            case 'Animal':
            case 'SumAnimal':
            case 'Te_Animal':
                $arr = ['mouse', 'cattle', 'tiger', 'rabbit', 'dragon', 'snake', 'horse', 'sheep', 'monkey', 'chicken', 'dog', 'pig'];
                $arr_2 = ['Ashor', 'Te_Animal', 'two_Animal', 'three_Animal', 'four_Animal', 'five_Animal', 'six_Animal', 'seven_Animal', 'eight_Animal', 'nine_Animal', 'ten_Animal', 'elven_Animal'];
                if (!in_array($betInfo['mingxi_2'], $arr) || !in_array($betInfo['mingxi_3'], $arr_2)) {
                    return false;
                }
                break;
            case 'EndNum':
                $arr = ['zero_end', 'one_end', 'two_end','three_end','four_end','five_end','six_end','seven_end','eight_end', 'nine_end'];
                if (!in_array($betInfo['mingxi_2'], $arr)) {
                    return false;
                }
                break;
            case 'AnimalEven':
                $arr = ['mouse', 'cattle', 'tiger', 'rabbit', 'dragon', 'snake', 'horse', 'sheep', 'monkey', 'chicken', 'dog', 'pig'];
                $arr_2 = ['two_Animal_in', 'three_Animal_in', 'four_Animal_in', 'five_Animal_in', 'two_Animal_not_in', 'three_Animal_not_in', 'four_Animal_not_in'];
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                if (!in_array($betInfo['mingxi_3'], $arr_2)) {
                    return false;
                }
                break;
            case 'EndNum_Even':
                $arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
                $arr_2 = ['two_end_in', 'three_end_in', 'four_end_in', 'two_end_not_in', 'three_end_not_in', 'four_end_not_in'];
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                if (!in_array($betInfo['mingxi_3'], $arr_2)) {
                    return false;
                }
                break;
            case 'AllMiss':
                $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49];
                $arr_2 = ['five_not_in', 'six_not_in', 'seven_not_in', 'eight_not_in', 'nine_not_in', 'ten_not_in', 'elven_not_in', 'twelve_not_in'];
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                if (!in_array($betInfo['mingxi_3'], $arr_2)) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }

        return true;
    }

    //北京快乐8
    public static function bj_8($betInfo) {
        $betres = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80];
        switch ($betInfo['mingxi_1']) {
            case 'choose_one':
                if (!in_array($betInfo['mingxi_2'], $betres)) {
                    return false;
                }
                break;
            case 'choose_two':
                $betarr = explode(',', $betInfo['mingxi_2']);
                $betarr = array_unique($betarr);
                if (count($betarr) != 2) {
                    return false;
                }
                foreach ($betarr as $value) {
                    if (!in_array($value, $betres)) {
                        return false;
                    }
                }
                break;
            case 'choose_three':
                $betarr = explode(',', $betInfo['mingxi_2']);
                $betarr = array_unique($betarr);
                if (count($betarr) != 3) {
                    return false;
                }
                foreach ($betarr as $value) {
                    if (!in_array($value, $betres)) {
                        return false;
                    }
                }
                break;
            case 'choose_four':
                $betarr = explode(',', $betInfo['mingxi_2']);
                $betarr = array_unique($betarr);
                if (count($betarr) != 4) {
                    return false;
                }
                foreach ($betarr as $value) {
                    if (!in_array($value, $betres)) {
                        return false;
                    }
                }
                break;
            case 'choose_five':
                $betarr = explode(',', $betInfo['mingxi_2']);
                $betarr = array_unique($betarr);
                if (count($betarr) != 5) {
                    return false;
                }
                foreach ($betarr as $value) {
                    if (!in_array($value, $betres)) {
                        return false;
                    }
                }
                break;
            case 'sum':
                if (!in_array($betInfo['mingxi_2'], ['total_sum_big', 'total_sum_small', 'total_sum_single', 'total_sum_double', 'total_sum_810'])) {
                    echo 11;die;
                    return false;
                }
                break;
            case 'up_middle_down':
                if (!in_array($betInfo['mingxi_2'], ['up_disc', 'middle_disc', 'down_disc'])) {
                    return false;
                }
                break;
            case 'odd_and_even':
                if (!in_array($betInfo['mingxi_2'], ['odd_disc', 'sum_disc', 'even_disc'])) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //北京pk10
    public static function bj_10($betInfo) {
        switch ($betInfo['mingxi_1']) {
            case 'first_second_sum':
                if (!in_array($betInfo['mingxi_2'], [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 'first_second_big','first_second_single', 'first_second_double','first_second_small'])) {
                    return false;
                }
                break;
            case 'first':
            case 'second':
            case 'third':
            case 'fourth':
            case 'fifth':
                if (!in_array($betInfo['mingxi_2'], [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 'big', 'small', 'single', 'double', 'dragon', 'tiger'])) {
                    return false;
                }
                break;
            case 'sixth':
            case 'seventh':
            case 'eighth':
            case 'ninth':
            case 'tenth':
                if (!in_array($betInfo['mingxi_2'], [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 'big', 'small', 'single', 'double'])) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //时时彩
    public static function ssc($betInfo) {
        switch ($betInfo['mingxi_1']) {
            case 'first_ball':
            case 'second_ball':
            case 'third_ball':
            case 'fourth_ball':
            case 'fifth_ball':
                if (!in_array($betInfo['mingxi_2'], [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 'big', 'small', 'single', 'double'])) {
                    return false;
                }
                break;
            case 'before_three_ball':
            case 'middle_three_ball':
            case 'after_three_ball':
                if (!in_array($betInfo['mingxi_2'], ['leopard', 'straight', 'pairs', 'Half_suitable', 'Miscellaneous_six'])) {
                    return false;
                }
                break;
            case 'tiger_ball':
                if (!in_array($betInfo['mingxi_2'], ['total_sum_big', 'total_sum_small', 'total_sum_single', 'total_sum_double', 'dragon', 'tiger', 'sum'])) {
                    return false;
                }
                break;
            case 'Bullfighting':
                if (!in_array($betInfo['mingxi_2'], ['not_cow', 'cow_one', 'cow_two', 'cow_three', 'cow_four', 'cow_five', 'cow_six', 'cow_seven', 'cow_eight', 'cow_nine', 'cow_cow', 'cow_big', 'cow_small', 'cow_single', 'cow_double'])) {
                    return false;
                }
                break;
            case 'poker':
                if (!in_array($betInfo['mingxi_2'], ['cow_eight', 'cow_nine', 'cow_cow', 'cow_big', 'cow_small', 'cow_single', 'cow_double', 'cow_double','cow_five','cow_powder','cow_four'])) {
                    echo $betInfo['mingxi_2'];
                    return false;
                }
                break;
            default:
            echo 666;die;
                return false;
                break;
        }
        return true;
    }

    //快乐十分
    public static function kl_10($betInfo) {
        $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20];
        switch ($betInfo['mingxi_1']) {
            case 'tiger_ball':
                if (!in_array($betInfo['mingxi_2'], ['total_sum_big', 'total_sum_small', 'total_sum_single', 'total_sum_double', 'total_sum_end_big', 'total_sum_end_small', 'dragon', 'tiger'])) {
                    return false;
                }
                break;
            case 'first_ball':
            case 'second_ball':
            case 'third_ball':
            case 'fourth_ball':
            case 'fifth_ball':
            case 'six_ball':
            case 'seven_ball':
            case 'eight_ball':
                if (!in_array($betInfo['mingxi_2'], [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 'big', 'end_big', 'east', 'middle', 'small', 'end_small', 'south', 'hair', 'single', 'sum_single', 'west', 'double', 'sum_double', 'north', 'white'])) {
                    return false;
                }
                break;
            case 'random_choose_two':
            case 'random_choose_two_group':
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                if (count($betres) != 2) {
                    return false;
                }
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                break;
            case 'random_choose_three':
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                if (count($betres) != 3) {
                    return false;
                }
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                break;
            case 'random_choose_four':
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                if (count($betres) != 4) {
                    return false;
                }
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                break;
            case 'random_choose_five':
                $betres = explode(',', $betInfo['mingxi_2']);
                $betres = array_unique($betres);
                if (count($betres) != 5) {
                    return false;
                }
                foreach ($betres as $value) {
                    if (!in_array($value, $arr)) {
                        return false;
                    }
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //十一选五
    public static function ten_five($betInfo) {
        $arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
        switch ($betInfo['mingxi_1']) {
            case 'total_sum':
                if (!in_array($betInfo['mingxi_2'], ['sum_small', 'sum_big', 'sum_single', 'sum_double', 'end_big', 'end_small', 'dragon', 'tiger'])) {
                    return false;
                }
                break;
            case 'first_ball':
            case 'second_ball':
            case 'third_ball':
            case 'fourth_ball':
            case 'fifth_ball':
                if (!in_array($betInfo['mingxi_2'], [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 'big', 'small', 'single', 'double'])) {
                    return false;
                }
                break;
            case 'random_choose':
                switch ($betInfo['mingxi_3']) {
                    case 'one_in_one':
                        if (!in_array($betInfo['mingxi_2'], $arr)) {
                            return false;
                        }
                        break;
                    case 'two_in_two':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 2) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'three_in_three':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 3) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'four_in_four':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 4) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'five_in_five':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 5) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'six_in_five':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 6) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'seven_in_five':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 7) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'eight_in_five':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 8) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    default:
                        return false;
                        break;
                }
                break;
            case 'group_choose':
            case 'vertical_choose':
                switch ($betInfo['mingxi_3']) {
                    case 'before_two':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 2) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    case 'before_three':
                        $betres = explode(',', $betInfo['mingxi_2']);
                        $betres = array_unique($betres);
                        if (count($betres) != 3) {
                            return false;
                        }
                        foreach ($betres as $value) {
                            if (!in_array($value, $arr)) {
                                return false;
                            }
                        }
                        break;
                    default:
                        return false;
                        break;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //快乐3
    public static function kl_3($betInfo) {
        $betres = [0, 1, 2, 3, 4, 5, 6];
        switch ($betInfo['mingxi_1']) {
            case 'sum':
                if (!in_array($betInfo['mingxi_2'], [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 'big', 'small', 'single', 'double'])) {
                    return false;
                }
                break;
            case 'two_even':
                $betres = array('1,2', '1,3', '1,4', '1,5', '1,6', '2,3', '2,4', '2,5', '2,6', '3,4', '3,5', '3,6', '4,5', '4,6', '5,6');
                if(!in_array($betInfo['mingxi_2'], $betres)) {
                        return false;
                }
                break;
            case 'gallbladder_ball':
                if (!in_array($betInfo['mingxi_2'], $betres)) {
                    return false;
                }
                break;
            case 'leopard':
                $betres = array('1,1,1', '2,2,2', '3,3,3', '4,4,4', '5,5,5', '6,6,6', 'random_leopard');
                if(!in_array($betInfo['mingxi_2'], $betres)) {
                        return false;
                }
                break;
            case 'pairs':
                $betres = array('1,1', '2,2', '3,3', '4,4', '5,5', '6,6');
                if(!in_array($betInfo['mingxi_2'], $betres)) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //北京28
    public static function bj_28($betInfo) {
        $betres = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'total_sum_big', 'total_sum_small', 'total_sum_single', 'total_sum_double'];
        switch ($betInfo['mingxi_1']) {
            case 'bj_28':
                if (!in_array($betInfo['mingxi_2'], $betres)) {
                    return false;
                }
                break;
            default:
                return FALSE;
                break;
        }
        return true;
    }

    //PC蛋蛋
    public static function pc_28($betInfo) {
        $arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'total_sum_big', 'big_single', 'max_small', 'red_wave', 'total_sum_small', 'big_double', 'max_big', 'leopard', 'total_sum_single', 'small_single', 'green_wave', 'total_sum_double', 'small_double', 'blue_wave'];
        switch ($betInfo['mingxi_1']) {
            case 'pc_28':
                if ($betInfo['mingxi_3'] == 'Tema_in_Three') {
                    $betres = explode(',', $betInfo['mingxi_2']);
                    $betres = array_unique($betres);
                    if (count($betres) != 3) {
                        return false;
                    }
                    foreach ($betres as $value) {
                        if (!in_array($value, [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27])) {
                            return false;
                        }
                    }
                } else {
                    if (!in_array($betInfo['mingxi_2'], $arr)) {
                        return false;
                    }
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

    //通用28
    public static function ty_28($betInfo) {
        $arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'big_num', 'big_single', 'max_small', 'small_num', 'big_double', 'max_big', 'single_num', 'small_single', 'double_num', 'small_double'];
        switch ($betInfo['mingxi_1']) {
            case 'dm_28':
            case 'jnd_28':
                if (!in_array($betInfo['mingxi_2'], $arr)) {
                    return false;
                }
                break;
            default:
                return false;
                break;
        }
        return true;
    }

}

?>
