<?php

namespace common\helpers;

use yii\helpers\ArrayHelper;

/*
 * 明细 中英文对照
 *
 */
/*
  fc_3d(福彩3D)       pl_3(排列三)
  dm_28(丹麦28)      jnd_28(加拿大28)       bj_28(北京28)
  dm_klc(丹麦快乐彩) bj_kl8(北京快乐8)  jnd_bs(加拿大卑斯)
  tj_ssc(天津时时彩) cq_ssc(重庆时时彩) xj_ssc(新疆时时彩)
  gx_k3(广西快三)    js_k3(江苏快三)    ah_k3(安徽快三)
  pc_28(PC蛋蛋)
  gd_11(广东11选5) sd_11(山东11选5) jx_11(江西11选5)
  cq_ten(重庆快乐十分)  gd_ten(广东快乐十分)
  bj_10(北京pk拾)
 */

class lotteryPlayer {

    public static function returnPlayer($type, $gameplay = '', $pankou = 'a') {
        if (empty($type)) {
            return false;
        }
        switch ($type) {
            case 'fc_3d':
            case 'pl_3':
                $action = 'three_lot';
                break;
            case 'tj_ssc':
            case 'cq_ssc':
            case 'xj_ssc':
                $action = 'ssc_lot';
                break;
            case 'gx_k3':
            case 'js_k3':
            case 'ah_k3':
                $action = 'quickThree_lot';
                break;
            case 'bj_kl8':
            case 'dm_klc':
            case 'jnd_bs':
                $action = 'happy_lot';
                break;
            case 'jnd_28':
                $action = 'jnd_28';
                break;
            case 'dm_28':
                $action = 'dm_28';
                break;
            case 'bj_28':
                $action = 'bj_28';
                break;
            case 'pc_28':
                $action = 'pc_28';
                break;
            case 'bj_10':
                $action = 'bj_10';
                break;
            case 'cq_ten':
            case 'gd_ten':
                $action = 'happy_ten_lot';
                break;
            case 'gd_11':
            case 'sd_11':
            case 'jx_11':
                $action = 'ten_five_lot';
                break;
            default:
                $action = 'liuhecai';
                break;
        }
        if ($action == 'liuhecai') {
            return self::liuhecai($gameplay, $pankou);
        } else {
            return self::$action();
        }
    }

    //福彩3D,排列3
    public static function three_lot() {
        $ball = array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9);

        $balls = array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双'
        );
        $arr = array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9);
        $res = array('0' => '總和,龍虎', '1' => array('total_sum_big' => '总和大',
                'total_sum_small' => '总和小',
                'total_sum_single' => '总和单',
                'total_sum_double' => '总和双',
                'dragon' => '龙',
                'tiger' => '虎',
                'sum' => '和')
        );
        $result = array('0' => '3连', '1' => array(
                'leopard' => '豹子',
                'straight' => '顺子',
                'pairs' => '对子',
                'Half_suitable' => '半顺',
                'Miscellaneous_six' => '杂六')
        );

        return array(
            //第一球 双  first_ball double
            'first_ball' => array('0' => '第一球', '1' => $balls),
            //第二球 大  second_ball big
            'second_ball' => array('0' => '第二球', '1' => $balls),
            //第三球 双  third_ball small
            'third_ball' => array('0' => '第三球', '1' => $balls),
            //跨度 1   span_ball1
            'span_ball' => array('0' => '跨度', '1' => $ball),
            //独胆 1   gallbladder_ball 1
            'gallbladder_ball' => array('0' => '独胆', '1' => $ball),
            //總和,龍虎 总和大  triple_ball total_sum_big
            'triple_ball' => $result,
            //3连 顺子  tiger_ball straight
            'tiger_ball' => $res
        );
    }

    //tj_ssc(天津时时彩) cq_ssc(重庆时时彩) xj_ssc(新疆时时彩)
    public static function ssc_lot() {
        $balls = array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双'
        );
        $res = array('0' => '總和,龍虎', '1' => array(
                'total_sum_big' => '总和大',
                'total_sum_small' => '总和小',
                'total_sum_single' => '总和单',
                'total_sum_double' => '总和双',
                'dragon' => '龙',
                'tiger' => '虎',
                'sum' => '和'));
        $result = array(
            'leopard' => '豹子',
            'straight' => '顺子',
            'pairs' => '对子',
            'Half_suitable' => '半顺',
            'Miscellaneous_six' => '杂六'
        );
        $arr = array('0' => '斗牛', '1' => array(
                'not_cow' => '没牛',
                'cow_one' => '牛1',
                'cow_two' => '牛2',
                'cow_three' => '牛3',
                'cow_four' => '牛4',
                'cow_five' => '牛5',
                'cow_six' => '牛6',
                'cow_seven' => '牛7',
                'cow_eight' => '牛8',
                'cow_nine' => '牛9',
                'cow_cow' => '牛牛',
                'cow_big' => '牛大',
                'cow_small' => '牛小',
                'cow_single' => '牛单',
                'cow_double' => '牛双')
        );
        $arr_sh = array('0' => '梭哈', '1' => array(
                'cow_five' => ' 五条',
                'cow_four' => '四条',
                'cow_cow' => '葫芦',
                'cow_big' => '顺子',
                'cow_small' => '三条',
                'cow_single' => '两对',
                'cow_double' => '一对',
                'cow_powder' => '散号')
        );
        return array(
            //第一球 大 first_ball big
            'first_ball' => ['0' => '第一球', '1' => $balls],
            //第二球 小 second_ball small
            'second_ball' => ['0' => '第二球', '1' => $balls],
            //第三球 单  third_ball single
            'third_ball' => ['0' => '第三球', '1' => $balls],
            //第四球 双 fourth_ball double
            'fourth_ball' => ['0' => '第四球', '1' => $balls],
            //第五球 1  fifth_ball 1
            'fifth_ball' => ['0' => '第五球', '1' => $balls],
            //總和,龍虎 龙  tiger_ball dragon
            'tiger_ball' => $res,
            //前三球 顺子  before_three_ball  straight
            'before_three_ball' => ['0' => '前三球', '1' => $result],
            //中三球 半顺  middle_three_ball  Half_suitable
            'middle_three_ball' => ['0' => '中三球', '1' => $result],
            //后三球 杂六  after_three_ball  Miscellaneous_six
            'after_three_ball' => ['0' => '后三球', '1' => $result],
            //斗牛 牛1 Bullfighting cow_one
            'Bullfighting' => $arr,
            //梭哈 四条  poker cow_nine
            'poker' => $arr_sh
        );
    }

    //js_k3(江苏快三)    gx_k3(广西快三)    ah_k3(安徽快三)
    public static function quickThree_lot() {
        $balls = array('3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双'
        );
        $result = array('0' => '豹子', '1' => array(
                '1,1,1' => '1,1,1',
                '2,2,2' => '2,2,2',
                '3,3,3' => '3,3,3',
                '4,4,4' => '4,4,4',
                '5,5,5' => '5,5,5',
                '6,6,6' => '6,6,6',
                'random_leopard' => '任意豹子')
        );
        $arr = array('0' => '两连', '1' => array(
                '1,2' => '1,2',
                '1,3' => '1,3',
                '1,4' => '1,4',
                '1,5' => '1,5',
                '1,6' => '1,6',
                '2,3' => '2,3',
                '2,4' => '2,4',
                '2,5' => '2,5',
                '2,6' => '2,6',
                '3,4' => '3,4',
                '3,5' => '3,5',
                '3,6' => '3,6',
                '4,5' => '4,5',
                '4,6' => '4,6',
                '5,6' => '5,6')
        );
        $data = array('0' => '对子', '1' => array(
                '1,1' => '1,1',
                '2,2' => '2,2',
                '3,3' => '3,3',
                '4,4' => '4,4',
                '5,5' => '5,5',
                '6,6' => '6,6')
        );
        return array(
            //和值 大   sum big
            'sum' => array('0' => '和值', '1' => $balls),
            //独胆 1  gallbladder_ball 1
            'gallbladder_ball' => array('0' => '独胆', '1' => array('1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6)),
            //豹子 任意豹子  leopard random_leopard
            //豹子 1,1,1  leopard 1,1,1
            'leopard' => $result,
            //两连 5,6  two_even 5,6
            'two_even' => $arr,
            //对子 1,1 pairs 1,1
            'pairs' => $data
        );
    }

    //六合彩
    public static function liuhecai($gameplay, $pankou = 'a') {
        $balls = array(
            '0' => '特码',
            '1' => array(
                '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20, '21' => 21, '22' => 22, '23' => 23, '24' => 24, '25' => 25, '26' => 26, '27' => 27, '28' => 28, '29' => 29, '30' => 30, '31' => 31, '32' => 32, '33' => 33, '34' => 34, '35' => 35, '36' => 36, '37' => 37, '38' => 38, '39' => 39, '40' => 40, '41' => 41, '42' => 42, '43' => 43, '44' => 44, '45' => 45, '46' => 46, '47' => 47, '48' => 48, '49' => 49, '1-10' => '1-10', '11-20' => '11-20', '21-30' => '21-30', '31-40' => '31-40', '41-49' => '41-49', 'big' => '大', 'wild_animals' => '野兽', 'small_double' => '小双', 'sum_double' => '和双', 'double' => '双', 'blue_wave' => '蓝波', 'Poultry' => '家禽', 'big_double' => '大双', 'sum_single' => '和单', 'single' => '单', 'green_wave' => '绿波', 'end_small' => '尾小', 'small_single' => '小单', 'sum_small' => '和小', 'small' => '小', 'red_wave' => '红波', 'end_big' => '尾大', 'big_single' => '大单', 'sum_big' => '和大', 'big' => '大')
        );
        $arr = array(
            '0' => '正码',
            '1' => array(
                '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20, '21' => 21, '22' => 22, '23' => 23, '24' => 24, '25' => 25, '26' => 26, '27' => 27, '28' => 28, '29' => 29, '30' => 30, '31' => 31, '32' => 32, '33' => 33, '34' => 34, '35' => 35, '36' => 36, '37' => 37, '38' => 38, '39' => 39, '40' => 40, '41' => 41, '42' => 42, '43' => 43, '44' => 44, '45' => 45, '46' => 46, '47' => 47, '48' => 48, '49' => 49, 'total_single' => '总单', 'total_double' => '总双', 'total_big' => '总大', 'total_small' => '总小', 'total_end_big' => '总尾大', 'total_end_small' => '总尾小', 'tiger' => '虎', 'dragon' => '龙'));
        $result = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20, '21' => 21, '22' => 22, '23' => 23, '24' => 24, '25' => 25, '26' => 26, '27' => 27, '28' => 28, '29' => 29, '30' => 30, '31' => 31, '32' => 32, '33' => 33, '34' => 34, '35' => 35, '36' => 36, '37' => 37, '38' => 38, '39' => 39, '40' => 40, '41' => 41, '42' => 42, '43' => 43, '44' => 44, '45' => 45, '46' => 46, '47' => 47, '48' => 48, '49' => 49
        );
        $arr_zm = array(
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双',
            'red_wave' => '红波',
            'green_wave' => '绿波',
            'blue_wave' => '蓝波',
            'sum_big' => '合大',
            'sum_small' => '合小',
            'sum_single' => '合单',
            'sum_double' => '合双',
            'end_big' => '尾大',
            'end_small' => '尾小'
        );
        $arr_gg = array(
            'single' => '单',
            'double' => '双',
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双',
            'red_wave' => '红波',
            'green_wave' => '绿波',
            'blue_wave' => '蓝波',
        );
        $arr_bb = array(
            'red_single' => '红单',
            'red_double' => '红双',
            'red_big' => '红大',
            'red_small' => '红小',
            'red_sum_single' => '红合单',
            'red_sum_double' => '红合双',
            'green_single' => '绿单',
            'green_double' => '绿双',
            'green_big' => '绿大',
            'green_small' => '绿小',
            'green_sum_single' => '绿合单',
            'green_sum_double' => '绿合双',
            'blue_single' => '蓝单',
            'blue_double' => '蓝双',
            'blue_big' => '蓝大',
            'blue_small' => '蓝小',
            'blue_sum_single' => '蓝合单',
            'blue_sum_double' => '蓝合双',
        );
        $arr_sx = array(
            'mouse' => array('0' => '鼠', array(10, 22, 34, 46)),
            'cattle' => array('0' => '牛', array(9, 21, 33, 45)),
            'tiger' => array('0' => '虎', array(8, 20, 32, 44)),
            'rabbit' => array('0' => '兔', array(7, 19, 31, 43)),
            'dragon' => array('0' => '龙', array(6, 18, 30, 42)),
            'snake' => array('0' => '蛇', array(5, 17, 29, 41)),
            'horse' => array('0' => '马', array(4, 16, 28, 40)),
            'sheep' => array('0' => '羊', array(3, 15, 27, 39)),
            'monkey' => array('0' => '猴', array(2, 14, 26, 38)),
            'chicken' => array('0' => '鸡', array(1, 13, 25, 37, 49)),
            'dog' => array('0' => '狗', array(12, 24, 36, 48)),
            'pig' => array('0' => '猪', array(11, 23, 35, 47))
        );
        $arr_ws = array(
            'one_end' => array('0' => '1尾', array(1, 11, 21, 31, 41)),
            'two_end' => array('0' => '2尾', array(2, 12, 22, 32, 42)),
            'three_end' => array('0' => '3尾', array(3, 13, 23, 33, 43)),
            'four_end' => array('0' => '4尾', array(4, 14, 24, 34, 44)),
            'five_end' => array('0' => '5尾', array(5, 15, 25, 35, 45)),
            'six_end' => array('0' => '6尾', array(6, 16, 26, 36, 46)),
            'seven_end' => array('0' => '7尾', array(7, 17, 27, 37, 47)),
            'eight_end' => array('0' => '8尾', array(8, 18, 28, 38, 48)),
            'nine_end' => array('0' => '9尾', array(9, 19, 29, 39, 49)),
            'zero_end' => array('0' => '0尾', array(10, 20, 30, 40))
        );
        $data_zt = array(
            '0' => '正码特',
            '1' => array(
                'JustCode_Te_one' => array('0' => '正1特', '1' => $result),
                'JustCode_Te_two' => array('0' => '正2特', '1' => $result),
                'JustCode_Te_three' => array('0' => '正3特', '1' => $result),
                'JustCode_Te_four' => array('0' => '正4特', '1' => $result),
                'JustCode_Te_five' => array('0' => '正5特', '1' => $result),
                'JustCode_Te_six' => array('0' => '正6特', '1' => $result))
        );
        $data_zm = array(
            '0' => '正码1-6',
            '1' => array(
                'JustCode_one' => array('0' => '正码1', '1' => $arr_zm),
                'JustCode_two' => array('0' => '正码2', '1' => $arr_zm),
                'JustCode_three' => array('0' => '正码3', '1' => $arr_zm),
                'JustCode_four' => array('0' => '正码4', '1' => $arr_zm),
                'JustCode_five' => array('0' => '正码5', '1' => $arr_zm),
                'JustCode_six' => array('0' => '正码6', '1' => $arr_zm))
        );
        $data_gg = array(
            '0' => '过关',
            '1' => array(
                'JustCode_one' => array('0' => '正码1', '1' => $arr_gg),
                'JustCode_two' => array('0' => '正码2', '1' => $arr_gg),
                'JustCode_three' => array('0' => '正码3', '1' => $arr_gg),
                'JustCode_four' => array('0' => '正码4', '1' => $arr_gg),
                'JustCode_five' => array('0' => '正码5', '1' => $arr_gg),
                'JustCode_six' => array('0' => '正码6', '1' => $arr_gg))
        );
        $data_lm = array(
            '0' => '连码',
            '1' => array(
                'second_full' => array('0' => '二全中', '1' => $result),
                'in_te' => array('0' => '二中特', '1' => $result),
                'Special_series' => array('0' => '特串', '1' => $result),
                'third_full' => array('0' => '三全中', '1' => $result),
                'third_in_second' => array('0' => '三中二', '1' => $result),
                'fourth_full' => array('0' => '四全中', '1' => $result))
        );
        $data_bb = array(
            '0' => ' 半波',
            '1' => array(
                'red_single' => array('0' => '红单', '1' => array(1, 7, 13, 19, 23, 29, 35, 45)),
                'red_double' => array('0' => '红双', '1' => array(2, 8, 12, 18, 24, 30, 34, 40, 46)),
                'red_big' => array('0' => '红大', '1' => array(29, 30, 34, 35, 40, 45, 46)),
                'red_small' => array('0' => '红小', '1' => array(1, 2, 7, 8, 12, 13, 18, 19, 23, 24)),
                'red_sum_single' => array('0' => '红合单', '1' => array(1, 7, 23, 29, 45, 12, 18, 30, 34)),
                'red_sum_double' => array('0' => '红合双', '1' => array(13, 19, 35, 2, 8, 24, 40, 46)),
                'green_single' => array('0' => '绿单', '1' => array(5, 11, 17, 21, 27, 33, 39, 43)),
                'green_double' => array('0' => '绿双', '1' => array(6, 16, 22, 28, 32, 38, 44)),
                'green_big' => array('0' => '绿大', '1' => array(27, 28, 32, 33, 38, 39, 43, 44)),
                'green_small' => array('0' => '绿小', '1' => array(5, 6, 11, 16, 17, 21, 22)),
                'green_sum_single' => array('0' => '绿合单', '1' => array(5, 16, 21, 27, 32, 38, 43)),
                'green_sum_double' => array('0' => '绿合双', '1' => array(6, 11, 17, 22, 28, 33, 39, 44)),
                'blue_single' => array('0' => '蓝单', '1' => array(3, 9, 15, 25, 31, 37, 41, 47)),
                'blue_double' => array('0' => '蓝双', '1' => array(4, 10, 14, 20, 26, 36, 42, 48)),
                'blue_big' => array('0' => '蓝大', '1' => array(25, 26, 31, 36, 37, 41, 42, 47, 48)),
                'blue_small' => array('0' => '蓝小', '1' => array(3, 4, 9, 10, 14, 15, 20)),
                'blue_sum_single' => array('0' => '蓝合单', '1' => array(3, 9, 10, 14, 25, 36, 41, 47)),
                'blue_sum_double' => array('0' => '蓝合双', '1' => array(4, 15, 20, 26, 31, 37, 42, 48)))
        );
        $data_xw = array(
            '0' => '一肖',
            '1' => $arr_sx
        );
        $data_ws = array(
            '0' => '尾数',
            '1' => $arr_ws
        );
        $data_tx = array(
            '0' => '特码生肖',
            '1' => $arr_sx
        );
        $data_hx = array(
            '0' => '合肖',
            '1' => array(
                'two_Animal' => array('0' => '二肖', '1' => $arr_sx),
                'three_Animal' => array('0' => '三肖', '1' => $arr_sx),
                'four_Animal' => array('0' => '四肖', '1' => $arr_sx),
                'five_Animal' => array('0' => '五肖', '1' => $arr_sx),
                'six_Animal' => array('0' => '六肖', '1' => $arr_sx),
                'seven_Animal' => array('0' => '七肖', '1' => $arr_sx),
                'eight_Animal' => array('0' => '八肖', '1' => $arr_sx),
                'nine_Animal' => array('0' => '九肖', '1' => $arr_sx),
                'ten_Animal' => array('0' => '十肖', '1' => $arr_sx),
                'elven_Animal' => array('0' => '十一肖', '1' => $arr_sx))
        );
        $data_sxl = array(
            '0' => '生肖连',
            '1' => array(
                'two_Animal_in' => array('0' => '二肖连中', '1' => $arr_sx),
                'three_Animal_in' => array('0' => '三肖连中', '1' => $arr_sx),
                'four_Animal_in' => array('0' => '四肖连中', '1' => $arr_sx),
                'two_Animal_not_in' => array('0' => '二肖连不中', '1' => $arr_sx),
                'three_Animal_not_in' => array('0' => '三肖连不中', '1' => $arr_sx),
                'four_Animal_not_in' => array('0' => '四肖连不中', '1' => $arr_sx))
        );
        $data_wsl = array(
            '0' => '尾数连',
            '1' => array(
                'two_end_in' => array('0' => '二尾连中', '1' => $arr_ws),
                'three_end_in' => array('0' => '三尾连中', '1' => $arr_ws),
                'four_end_in' => array('0' => '四尾连中', '1' => $arr_ws),
                'two_end_not_in' => array('0' => '二尾连不中', '1' => $arr_ws),
                'three_end_not_in' => array('0' => '三尾连不中', '1' => $arr_ws),
                'four_end_not_in' => array('0' => '四尾连不中', '1' => $arr_ws))
        );
        $data_qbz = array(
            '0' => '全不中',
            '1' => array(
                'five_not_in' => array('0' => '五不中', '1' => $result),
                'six_not_in' => array('0' => '六不中', '1' => $result),
                'seven_not_in' => array('0' => '七不中', '1' => $result),
                'eight_not_in' => array('0' => '八不中', '1' => $result),
                'nine_not_in' => array('0' => '九不中', '1' => $result),
                'ten_not_in' => array('0' => '十不中', '1' => $result),
                'elven_not_in' => array('0' => '十一不中', '1' => $result),
                'twelve_not_in' => array('0' => '十二不中', '1' => $result))
        );

        //五行
        $data_wx = array(
            '0' => '五行',
            '1' => array(
                'metal' => array('0' => '金', '1' => array(1, 6, 11, 16, 21, 26, 31, 36, 41, 46)),
                'wood' => array('0' => '木', '1' => array(2, 7, 12, 17, 22, 27, 32, 37, 42, 47)),
                'water' => array('0' => '水', '1' => array(3, 8, 13, 18, 23, 28, 33, 38, 43, 48)),
                'fire' => array('0' => '火', '1' => array(4, 9, 14, 19, 24, 29, 34, 39, 44, 49)),
                'earth' => array('0' => '土', '1' => array(5, 10, 15, 20, 25, 30, 35, 40, 45))
            )
        );
        //七码
        $data_qm = array(
            '0' => '七码',
            '1' => array(
                'seven_double' => '单0 . 双7',
                'six_double' => '单1 . 双6',
                'five_double' => '单2 . 双5',
                'four_double' => '单3 . 双4',
                'three_double' => '单4 . 双3',
                'two_double' => '单5 . 双2',
                'one_double' => '单6 . 双1',
                'seven_single' => '单7 . 双0',
                'seven_small' => '大0 . 小7',
                'six_small' => '大1 . 小6',
                'five_small' => '大2 . 小5',
                'four_small' => '大3 . 小4',
                'three_small' => '大4 . 小3',
                'two_small' => '大5 . 小2',
                'one_small' => '大6 . 小1',
                'seven_big' => '大7 . 小0',
            )
        );
        //正肖
        $data_zx = array(
            '0' => '正肖',
            '1' => $arr_sx,
        );
        //总肖
        $data_zox = array(
            '0' => '总肖',
            '1' => array(
                'all_Animal_single' => '总肖单',
                'all_Animal_double' => '总肖双',
                'all_Animal_shunzi' => '234肖',
                'all_Animal_six' => '6肖',
                'all_Animal_five' => '5肖',
                'all_Animal_seven' => '7肖',
            )
        );
        //特码头数
        $data_tm_first = array(
            '0' => '特码头数',
            '1' => array(
                'First_num_zero' => array('0' => '头0', '1' => array(1, 2, 3, 4, 5, 6, 7, 8, 9)),
                'First_num_one' => array('0' => '头1', '1' => array(11, 12, 13, 14, 15, 16, 17, 18, 19)),
                'First_num_two' => array('0' => '头2', '1' => array(21, 22, 23, 24, 25, 26, 27, 28, 29)),
                'First_num_three' => array('0' => '头3', '1' => array(31, 32, 33, 34, 35, 36, 37, 38, 39)),
                'First_num_four' => array('0' => '头4', '1' => array(41, 42, 43, 44, 45, 46, 47, 48, 49))
            )
        );
        //特码尾数
        $data_tm_last = array(
            '0' => '特码尾数',
            '1' => array(
                'Last_num_zero' => array('0' => '尾0', '1' => array(10, 20, 30, 40,)),
                'Last_num_one' => array('0' => '尾1', '1' => array(1, 11, 21, 31, 41)),
                'Last_num_two' => array('0' => '尾2', '1' => array(2, 12, 22, 32, 42)),
                'Last_num_three' => array('0' => '尾3', '1' => array(3, 13, 23, 33, 43)),
                'Last_num_four' => array('0' => '尾4', '1' => array(4, 14, 24, 34, 44)),
                'Last_num_five' => array('0' => '尾5', '1' => array(5, 15, 25, 35, 45)),
                'Last_num_six' => array('0' => '尾6', '1' => array(6, 16, 26, 36, 46)),
                'Last_num_seven' => array('0' => '尾7', '1' => array(7, 17, 27, 37, 47)),
                'Last_num_eight' => array('0' => '尾8', '1' => array(8, 18, 28, 38, 48)),
                'Last_num_nine' => array('0' => '尾9', '1' => array(9, 19, 29, 39, 49))
            )
        );

        $data_all = array(
            'Tema' => $balls, //特码 2
            'JustCode' => $arr, //正码  30
            'JustCode_Te' => $data_zt, //正码特  正1特 30
            'JustCode_one_six' => $data_zm, //正码1-6 正码1 红波
            'PassTest' => $data_gg, //过关 正码1,2 单,双
            'JointMark' => $data_lm, //连码 二全中  20,21
            'HalfWave' => $data_bb, //半波 红小  1,2,7,8,12,13,18,19,23,24
            'Animal' => $data_xw, //一肖  兔
            'EndNum' => $data_ws, //尾数  0尾
            'Te_Animal' => $data_tx, //特码生肖 兔
            'SumAnimal' => $data_hx, //合肖 二肖 虎,兔
            'AnimalEven' => $data_sxl, //生肖连  二肖连中  牛,羊
            'EndNum_Even' => $data_wsl, //尾数连 二尾连中  2,4
            'AllMiss' => $data_qbz, //全不中 五不中 21,22,23,34,35
            'Five_elements' => $data_wx, //五行 金 木 水 火 土
            'Seven_code' => $data_qm, //七码
            'Just_Animal' => $data_zx, //正肖
            'All_Animal' => $data_zox, //总肖
            'Te_First_num' => $data_tm_first, //特码头数
            'Te_Last_num' => $data_tm_last, //特码尾数
        );

        if ($pankou == 'a') {
            return array($gameplay => $data_all[$gameplay]);
        } else {
            return array('Tema' => $data_all['Tema']);
        }
    }

    //北京快乐8,丹麦快乐彩,加拿大卑斯
    public static function happy_lot() {
        $balls = [
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20, '21' => 21, '22' => 22, '23' => 23, '24' => 24, '25' => 25, '26' => 26, '27' => 27, '28' => 28, '29' => 29, '30' => 30, '31' => 31, '32' => 32, '33' => 33, '34' => 34, '35' => 35, '36' => 36, '37' => 37, '38' => 38, '39' => 39, '40' => 40, '41' => 41, '42' => 42, '43' => 43, '44' => 44, '45' => 45, '46' => 46, '47' => 47, '48' => 48, '49' => 49, '50' => 50, '51' => 51, '52' => 52, '53' => 53, '54' => 54, '55' => 55, '56' => 56, '57' => 57, '58' => 58, '59' => 59, '60' => 60, '61' => 61, '62' => 62, '63' => 63, '64' => 64, '65' => 65, '66' => 66, '67' => 67, '68' => 68, '69' => 69, '70' => 70, '71' => 71, '72' => 72, '73' => 73, '74' => 74, '75' => 75, '76' => 76, '77' => 77, '78' => 78, '79' => 79, '80' => 80
        ];
        $arr_szx = array(
            'up_disc' => '上盘',
            'middle_disc' => '中盘',
            'down_disc' => '下盘'
        );
        $arr_jop = array('odd_disc' => '奇盘',
            'sum_disc' => '和盘',
            'even_disc' => '偶盘'
        );

        $arr_hz = array('total_sum_big' => '总和大',
            'total_sum_small' => '总和小',
            'total_sum_single' => '总和单',
            'total_sum_double' => '总和双',
            'total_sum_810' => '总和810'
        );
        return array(
            //选一 20  choose_one  20
            'choose_one' => array('0' => '选一', '1' => $balls),
            // //选二 二中二 21,22
            'choose_two' => array('0' => '选二', '1' => $balls),
            //选三 21,22,23 三中三:20,三中二:3
            'choose_three' => array('0' => '选三', '1' => $balls),
            //选四 21,22,23,24 四中四:50,四中三:5,四中二:3
            'choose_four' => array('0' => '选四', '1' => $balls),
            //选五 21,22,23,24,25 五中五:250,五中四:20,五中三:5
            'choose_five' => array('0' => '选五', '1' => $balls),
            //上中下 上盘  up_middle_down up_disc
            'up_middle_down' => array('0' => '上中下', '1' => $arr_szx),
            //奇和偶 偶盘  odd_and_even even_disc
            'odd_and_even' => array('0' => '奇和偶', '1' => $arr_jop),
            //和值 总和大    sum total_sum_big
            'sum' => array('0' => '和值', '1' => $arr_hz)
        );
    }

    //通用28 //加拿大28 丹麦28
    public static function jnd_28() {
        //加拿大28 双数    jnd_28 double_num
        //丹麦28 大单     dm_28  big_single
        $balls = array(
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'big_num' => '大数', 'big_single' => '大单', 'max_small' => '极小', 'small_num' => '小数', 'big_double' => '大双', 'max_big' => '极大', 'single_num' => '单数', 'small_single' => '小单', 'double_num' => '双数', 'small_double' => '小双'
        );
        return array(
            'jnd_28' => array('0' => '加拿大28', '1' => $balls)
        );
    }

    //通用28 //加拿大28 丹麦28
    public static function dm_28() {
        //加拿大28 双数    jnd_28 double_num
        //丹麦28 大单     dm_28  big_single
        $balls = array(
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'big_num' => '大数', 'big_single' => '大单', 'max_small' => '极小', 'small_num' => '小数', 'big_double' => '大双', 'max_big' => '极大', 'single_num' => '单数', 'small_single' => '小单', 'double_num' => '双数', 'small_double' => '小双'
        );
        return array(
            'dm_28' => array('0' => '丹麦28', '1' => $balls)
        );
    }

    public static function ty_28() {
        //加拿大28 双数    jnd_28 double_num
        //丹麦28 大单     dm_28  big_single
        return array(
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 'big_num' => '大数', 'big_single' => '大单', 'max_small' => '极小', 'small_num' => '小数', 'big_double' => '大双', 'max_big' => '极大', 'single_num' => '单数', 'small_single' => '小单', 'double_num' => '双数', 'small_double' => '小双'
        );
    }

    //北京28
    public static function bj_28() {
        //北京28  总和大 bj_28 total_sum_big
        $balls = array(
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
            'total_sum_big' => '总和大',
            'total_sum_small' => '总和小',
            'total_sum_single' => '总和单'
        );
        return array(
            'bj_28' => array('0' => 'PK北京28', '1' => $balls)
        );
    }

    //pc_dd(PC蛋蛋)
    public static function pc_28() {
        //北京28  特码包三 bj_28 Tema_in_Three
        $balls = array(
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
            'total_sum_big' => '总和大',
            'total_sum_small' => '总和小',
            'total_sum_single' => '总和单',
            'total_sum_double' => '总和双',
            'big_single' => '大单',
            'big_double' => '大双',
            'small_single' => '小单',
            'small_double' => '小双',
            'max_small' => '极小',
            'max_big' => '极大',
            'red_wave' => '红波',
            'green_wave' => '绿波',
            'blue_wave' => '蓝波',
            'leopard' => '豹子',
            'Tema_in_Three' => '特码包三'
        );
        return array(
            'pc_28' => array('0' => 'PC蛋蛋', '1' => $balls)
        );
    }

    //bj_10(北京pk拾)
    public static function bj_10() {
        $data_zh = array(
            '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19,
            'first_second_big' => '冠亚大',
            'first_second_small' => '冠亚小',
            'first_second_single' => '冠亚单',
            'first_second_double' => '冠亚双'
        );
        $data = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双',
            'dragon' => '龙',
            'tiger' => '虎'
        );
        $data_ = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双'
        );
        return array(
            //冠亚军和 冠亚小 first_second_sum first_second_small
            'first_second_sum' => array('0' => '冠亚军和', '1' => $data_zh),
            //冠军 小  first small
            //冠军 虎  龙虎 虎 1V10龙虎  dragon_tiger tiger first
            'first' => array('0' => '冠军', '1' => $data),
            //亚军 大  second big
            //亚军 虎  龙虎 虎 2V9龙虎  dragon_tiger tiger second
            'second' => array('0' => '亚军', '1' => $data),
            //亚军 大  second big
            //第三名 虎  龙虎 虎 3V8龙虎  dragon_tiger tiger third
            'third' => array('0' => '第三名', '1' => $data),
            //亚军 大  second big
            //第四名 虎  龙虎 虎 4V7龙虎  dragon_tiger tiger fourth
            'fourth' => array('0' => '第四名', '1' => $data),
            //亚军 大  second big
            //第五名 龙  龙虎 虎 5V6龙虎  dragon_tiger 龙 fifth
            'fifth' => array('0' => '第五名', '1' => $data),
            //第六名 大 sixth big
            'sixth' => array('0' => '第六名', '1' => $data_),
            //第七名 小 seventh small
            'seventh' => array('0' => '第七名', '1' => $data_),
            //第八名 大 eighth big
            'eighth' => array('0' => '第八名', '1' => $data_),
            //第九名 大 ninth big
            'ninth' => array('0' => '第九名', '1' => $data_),
            //第十名 1 tenth 1
            'tenth' => array('0' => '第十名', '1' => $data_)
        );
    }

    //cq_kl10(重庆快乐十分)  gd_kl10(广东快乐十分)
    public static function happy_ten_lot() {
        $arr_zh = array(
            'total_sum_big' => '总和大',
            'total_sum_small' => '总和小',
            'total_sum_single' => '总和单',
            'total_sum_double' => '总和双',
            'total_sum_end_big' => '总和尾大',
            'total_sum_end_small' => '总和尾小',
            'dragon' => '龙',
            'tiger' => '虎',
        );
        $arr_ball = array(
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双',
            'sum_single' => '和数单',
            'sum_double' => '和数双',
            'end_big' => '尾大',
            'end_small' => '尾小',
        );
        $arr_lm = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20
        );
        $data_ball = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11, '12' => 12, '13' => 13, '14' => 14, '15' => 15, '16' => 16, '17' => 17, '18' => 18, '19' => 19, '20' => 20,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双',
            'sum_single' => '和数单',
            'sum_double' => '和数双',
            'end_big' => '尾大',
            'end_small' => '尾小',
            'east' => '东',
            'south' => '南',
            'west' => '西',
            'north' => '北',
            'middle' => '中',
            'hair' => '发',
            'white' => '白'
        );
        $data_lm = array(
            'random_choose_two' => array('0' => '任选二', '1' => $arr_lm),
            'random_choose_two_group' => array('0' => '任选二组', '1' => $data_ball),
            'random_choose_three' => array('0' => '任选三', '1' => $arr_lm),
            'random_choose_four' => array('0' => '任选四', '1' => $arr_lm),
            'random_choose_five' => array('0' => '任选五', '1' => $arr_lm),
        );
        return array(
            // 第一球 7 first_ball 7
            'first_ball' => array('0' => '第一球', '1' => $data_ball),
            //第二球  7 second_ball 7
            'second_ball' => array('0' => '第二球', '1' => $data_ball),
            //第三球  7  third_ball 7
            'third_ball' => array('0' => '第三球', '1' => $data_ball),
            //第四球  7  fourth_ball 7
            'fourth_ball' => array('0' => '第四球', '1' => $data_ball),
            //第五球  7   fifth_ball 7
            'fifth_ball' => array('0' => '第五球', '1' => $data_ball),
            //第六球  7  six_ball 7
            'six_ball' => array('0' => '第六球', '1' => $data_ball),
            //第七球  7  seven_ball 7
            'seven_ball' => array('0' => '第七球', '1' => $data_ball),
            //第八球  7  eight_ball 7
            'eight_ball' => array('0' => '第八球', '1' => $data_ball),
            //总和  尾大  tiger_ball end_big
            'tiger_ball' => array('0' => '总和', '1' => $arr_zh),
            //连码
            //任选二  1,2 random_choose_two 1,2
            'random_choose_two' => array('0' => '任选二', '1' => $arr_lm),
            //任选二组 1,2  random_choose_two_group 1,2
            'random_choose_two_group' => array('0' => '任选二组', '1' => $arr_lm),
            //任选三 1,2,3  random_choose_three  1,2,3
            'random_choose_three' => array('0' => '任选三', '1' => $arr_lm),
            //任选四 1,2,3,4   random_choose_four  1,2,3,4
            'random_choose_four' => array('0' => '任选四', '1' => $arr_lm),
            //任选五 1,2,3,4,5  random_choose_five  1,2,3,4,5
            'random_choose_five' => array('0' => '任选五', '1' => $arr_lm)
        );
    }

    //gd_11(广东11选5) sd_11(山东11选5) jx_11(江西11选5)
    public static function ten_five_lot() {
        $arr_zh = array(
            'sum_big' => '和大',
            'sum_small' => '和小',
            'sum_single' => '和单',
            'sum_double' => '和双',
            'end_big' => '尾大',
            'end_small' => '尾小',
            'dragon' => '龙',
            'tiger' => '虎'
        );
        $arr_ball = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11,
            'big' => '大',
            'small' => '小',
            'single' => '单',
            'double' => '双'
        );
        $balls = array(
            '1' => 1, '2' => 2, '3' => 3, '4' => 4, '5' => 5, '6' => 6, '7' => 7, '8' => 8, '9' => 9, '10' => 10, '11' => 11
        );
        $arr_zhx_2 = array(
            'first_ball' => array('0' => '第一球', '1' => $balls),
            'second_ball' => array('0' => '第二球', '1' => $balls)
        );
        $arr_zhx_3 = array(
            'first_ball' => array('0' => '第一球', '1' => $balls),
            'second_ball' => array('0' => '第二球', '1' => $balls),
            'third_ball' => array('0' => '第三球', '1' => $balls)
        );
        $data_rx = array(
            'one_in_one' => array('0' => '一中一', '1' => $balls),
            'two_in_two' => array('0' => '二中二', '1' => $balls),
            'three_in_three' => array('0' => '三中三', '1' => $balls),
            'four_in_four' => array('0' => '四中四', '1' => $balls),
            'five_in_five' => array('0' => '五中五', '1' => $balls),
            'six_in_five' => array('0' => '六中五', '1' => $balls),
            'seven_in_five' => array('0' => '七中五', '1' => $balls),
            'eight_in_five' => array('0' => '八中五', '1' => $balls)
        );
        $data_zx = array(
            'before_two' => array('0' => '前二', '1' => $balls),
            'before_three' => array('0' => '前三', '1' => $balls)
        );
        $data_zhx = array(
            'before_two' => array('0' => '前二', '1' => $arr_zhx_2),
            'before_three' => array('0' => '前三', '1' => $arr_zhx_3)
        );
        return array(
            //总和 尾大 total_sum end_big
            'total_sum' => array('0' => '总和', '1' => $arr_zh),
            //第一球 单   first_ball single
            'first_ball' => array('0' => '第一球', '1' => $arr_ball),
            //第二球 双   second_ball double
            'second_ball' => array('0' => '第二球', '1' => $arr_ball),
            //第三球 大  third_ball big
            'third_ball' => array('0' => '第三球', '1' => $arr_ball),
            //第四球 小   fourth_ball small
            'fourth_ball' => array('0' => '第四球', '1' => $arr_ball),
            //第五球 1  fifth_ball 1
            'fifth_ball' => array('0' => '第五球', '1' => $arr_ball),
            //任选 二中二 1,2   random_choose  two_in_two 1,2
            'random_choose' => array('0' => '任选', '1' => $data_rx),
            //组选 前三 1,2,3   group_choose before_three 1,2,3
            'group_choose' => array('0' => '组选', '1' => $data_zx),
            //直选 前二 1,2   vertical_choose before_two 1,2
            'vertical_choose' => array('0' => '直选', '1' => $data_zhx)
        );
    }

}
