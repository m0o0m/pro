<?php

namespace common\helpers;

class bytes {

    /**
     * 转换一个string字符串为byte数组
     * @param $str 需要转换的字符串
     * @param $bytes 目标byte数组
     * @author zikie
     */
    public static function getbytes($str) {

        $len = strlen($str);
        $bytes = array();
        for ($i = 0; $i < $len; $i++) {
            if (ord($str[$i]) >= 128) {
                $byte = ord($str[$i]) - 256;
            } else {
                $byte = ord($str[$i]);
            }
            $bytes[] = $byte;
        }
        return $bytes;
    }

    /**
     * 将字节数组转化为string类型的数据
     * @param $bytes 字节数组
     * @param $str 目标字符串
     * @return 一个string类型的数据
     */
    public static function tostr($bytes) {
        $str = '';
        foreach ($bytes as $ch) {
            $str .= chr($ch);
        }

        return $str;
    }

    /**
     * 转换一个int为byte数组
     * @param $byt 目标byte数组
     * @param $val 需要转换的字符串
     * @author zikie
     */
    public static function integertobytes($val) {
        $byt = array();
        $byt [0] = ($val >> 24 & 0xff);
        $byt [1] = ($val >> 16 & 0xff);
        $byt [2] = ($val >> 8 & 0xff);
        $byt [3] = ($val & 0xff);
        return $byt;
    }

    /**
      public static function integerToBytes($val) {
      $byt = array ();
      $byt [0] = ($val >> 24 & 0xff);
      $byt [1] = ($val >> 16 & 0xff);
      $byt [2] = ($val >> 8 & 0xff);
      $byt [3] = ($val & 0xff);
      return $byt;
      }
     * 从字节数组中指定的位置读取一个integer类型的数据
     * @param $bytes 字节数组
     * @param $position 指定的开始位置
     * @return 一个integer类型的数据
     */
    public static function bytesToInteger($bytes, $position) {
        $val = 0;
        $val |= $bytes[$position] & 0xff;
        $val <<= 8;
        $val |= $bytes[$position + 1] & 0xff;
        $val <<= 8;
        $val |= $bytes[$position + 2] & 0xff;
        $val <<= 8;
        $val |= $bytes[$position + 3] & 0xff;

        return $val;
    }

    /**
     * 转换一个shor字符串为byte数组
     * @param $byt 目标byte数组
     * @param $val 需要转换的字符串
     * @author zikie
     */
    public static function shorttobytes($val) {
        $byt = array();
        $byt[0] = ($val & 0xff);
        $byt[1] = ($val >> 8 & 0xff);
        return $byt;
    }

    /**
     * 从字节数组中指定的位置读取一个short类型的数据。
     * @param $bytes 字节数组
     * @param $position 指定的开始位置
     * @return 一个short类型的数据
     */
    public static function bytestoshort($bytes, $position) {
        $val = 0;
        $val = $bytes[$position + 1] & 0xff;
        $val = $val << 8;
        $val |= $bytes[$position] & 0xff;
        return $val;
    }

}
