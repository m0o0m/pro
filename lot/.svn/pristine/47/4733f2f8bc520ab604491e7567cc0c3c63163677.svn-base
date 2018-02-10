<?php

namespace common\helpers;

class Encrypt {

    private static $MD5Key = "qhsjbevrg773grvtee";
    private static $DESKey = "PK55HHEE";
    private static $key;

    /*  测试代码

      common中的params 添加两个配置
      //钱包地址
      'tcphost' => '127.0.0.1',
      'tcpport' => '9999',

      use common\helpers\TcpConPoll;
      use common\helpers\Encrypt;

      获取余额
      $str = Encrypt::GetBalance('owei',"CNY");
      $socket = TcpConPoll::getInstace();


      加减钱  只需要变化金额字段
      var_dump($socket::send($str));
      $strr = Encrypt::Transfer("owei","CNY",-20,"3241243243432","彩票注单2017092006228763~201709200622879,类型:重庆时时彩,共计:3单");
      var_dump($socket::send($strr));
     */

    /** 获取平台方余额
     * username string 彩票库用户名
     * cur      string 币种
     * siteid   string 平台方站点id
     * agentid  string 平台方代理id
     */
    //Encrypt::GetBalance('owei',"CNY");
    public static function GetBalance($username, $cur) {
        $crypt = static::DES1(static::$DESKey);
        $data = array();
        $data['cmd'] = "getbalance";
        $data['member']['username'] = $username;
        $data['member']['currency'] = $cur;
        $params = static::encrypt(json_encode($data));
        $Key = static::getKey($params);
        $result = ['data' => $params, 'key' => $Key];
        $info = json_encode($result);
        return $info;
    }

    /** 额度加减
     * username  string 彩票库的用户名
     * amount    int or float  加减钱的金额
     * cur    string 币种
     */
    //Encrypt::Transfer("owei","CNY",-20,"3241243243432","彩票注单2017092006228763~201709200622879,类型:重庆时时彩,共计:3单");
    //Encrypt::Transfer("owei","CNY",20,"4324324234234","彩票注单2017092006228763,类型:重庆时时彩,共计:1单");
    public static function Transfer($username, $cur, $amount,$requestId,$remark) {
        $crypt = static::DES1(static::$DESKey);
        $data = array();
        $data['cmd'] = "transfer";
        $data['requestId'] = $requestId;
        //彩票注单2017092006228763~201709200622879,类型:重庆时时彩,共计:3单
        $data['data'] = $remark;
        $data['member']['username'] = $username;
        $data['member']['currency'] = $cur;
        $data['member']['amount'] = $amount;
        $params = static::encrypt(json_encode($data));
        $Key = static::getKey($params);
        $result = ['data' => $params, 'key' => $Key];
        $info = json_encode($result);
        return $info;
    }

    //验证加减钱是否成功
    public static function checkTransfer($requestId) {
        $crypt = static::DES1(static::$DESKey);
        $data = array();
        $data['cmd'] = "checktransfer";
        $data['requestId'] = '1';
        $params = static::encrypt(json_encode($data));
        $Key = static::getKey($params);
        $result = ['data' => $params, 'key' => $Key];
        $info = json_encode($result);
        return $info;
    }

    public static function getKey($params) {
        return md5($params . static::$MD5Key);
    }

    //des加密
    public static function DES1($key) {
        static::$key = $key;
    }

    public static function encrypt($input) {
        $size = mcrypt_get_block_size('des', 'ecb');
        $input = static::pkcs5_pad($input, $size);
        $key = static::$key;
        $td = mcrypt_module_open('des', '', 'ecb', '');
        $iv = @mcrypt_create_iv(mcrypt_enc_get_iv_size($td), MCRYPT_RAND);
        @mcrypt_generic_init($td, static::$DESKey, $iv);
        $data = mcrypt_generic($td, $input);
        mcrypt_generic_deinit($td);
        mcrypt_module_close($td);
        $data = base64_encode($data);
        return preg_replace("/\s*/", '', $data);
    }

    public static function decrypt($encrypted) {
        static::DES1(static::$DESKey);
        $encrypted = preg_replace('/[\s　]/', '+', $encrypted);
        $encrypted = base64_decode($encrypted);
        $key = static::$key;
        $td = mcrypt_module_open('des', '', 'ecb', '');
        //使用MCRYPT_DES算法,cbc模式
        $iv = @mcrypt_create_iv(mcrypt_enc_get_iv_size($td), MCRYPT_RAND);
        $ks = mcrypt_enc_get_key_size($td);
        @mcrypt_generic_init($td, $key, $iv);
        //初始处理
        $decrypted = mdecrypt_generic($td, $encrypted);
        //解密
        mcrypt_generic_deinit($td);
        //结束
        mcrypt_module_close($td);
        $y = static::pkcs5_unpad($decrypted);
        return $y;
    }

    public static function pkcs5_pad($text, $blocksize) {
        $pad = $blocksize - (strlen($text) % $blocksize);
        return $text . str_repeat(chr($pad), $pad);
    }

    public static function pkcs5_unpad($text) {
        $pad = ord($text{strlen($text) - 1});
        if ($pad > strlen($text)) {
            return false;
        }

        if (strspn($text, chr($pad), strlen($text) - $pad) != $pad) {
            return false;
        }

        return substr($text, 0, -1 * $pad);
    }

}

?>