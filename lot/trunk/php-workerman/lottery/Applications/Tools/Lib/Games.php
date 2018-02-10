<?php

namespace Applications\Tools\Lib;

class Games {

    private $UserAgent = "Lottery_Api_Go_";
    private $MD5Key = '';
    private $DESKey = '';
    private $site_id = 'bbb';
    private $apiurl = 'http://127.0.0.1:8084';

    function __construct() {
        //bbb
//        $this->DESKey = 'OLKIUJYH';
//        $this->MD5Key = 'wdjikoljhgftyuips';
        //aaa
        $this->DESKey = 'AGTHIJUL';
        $this->MD5Key = 'adikjujkolpkoiuhd';
    }

    function jie(){
        $str='ZBToIQ4AmmTLV5sRfjb37kXzXYupyRnAAyRSIYBm0eyqVkzU0sSK7afk897dJhwk7MuvMrZagrve0C/HSGfnPFiEuKpChSWQkqXAZ7+qqq1sfWXSZFo3uhef+yP4uQW4';
        $data= $this->decrypt($str);
        var_dump($data);
        return;
    }
            
    function login() {
        //$this->apiurl = 'http://113.10.246.105:8084';
        $this->apiurl = 'http://127.0.0.1:8084';
        $crypt = $this->DES1($this->DESKey);
        $siteid = 'bbb';
        $loginname = 't_phpzero6';
        $pwd = '123456';
        $agent_name = 'pk_t_a_dl';
        $ip = '127.0.0.1';

        $params = $this->encrypt(
                "site_id=" . $siteid .
                "/\\\\/uname=" . $loginname .
                "/\\\\/pwd=" . $pwd .
                "/\\\\/agent_name=" . $agent_name .
                "/\\\\/create_ip=" . $ip
        );
        $Key = $this->getKey($params);
        $url = $this->apiurl . "/account/login?params=" . $params . "&key=" . $Key;
        $result = $this->web_curl($url);
        print_r($result);
    }

    function spider() {
        $this->apiurl = 'http://127.0.0.1:9898';
        $crypt = $this->DES1($this->DESKey);
        $siteid = 'aaa';
        //$loginname = 't_phpzero6';
        //$pwd = '123456';
        //$agent_name = 'pk_t_a_dl';
        //$ip = '127.0.0.1';

        $params = $this->encrypt(
                "site_id=" . $siteid
        );

        $Key = $this->getKey($params);
        $url = $this->apiurl . "/getList?params=" . $params . "&key=" . $Key;
        $result = $this->web_curl($url);
        return $result;
    }

    function delkeys($keys) {
        $this->apiurl = 'http://127.0.0.1:9898';
        $crypt = $this->DES1($this->DESKey);
        $siteid = 'aaa';
        //$loginname = 't_phpzero6';
        //$pwd = '123456';
        //$agent_name = 'pk_t_a_dl';
        //$ip = '127.0.0.1';

        $params = $this->encrypt(
                "site_id=" . $siteid .
                "/\\\\/keys=" . $keys
        );

        $Key = $this->getKey($params);
        $url = $this->apiurl . "/delList?params=" . $params . "&key=" . $Key;
        $result = $this->web_curl($url);
        return $result;
    }

    function getKey($params) {
        return md5($params . $this->MD5Key);
    }

    //返回数据
    function web_curl($url) {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_TIMEOUT, 30);
        curl_setopt($ch, CURLOPT_USERAGENT, $this->UserAgent . $this->site_id);
        //执行命令
        $data = curl_exec($ch);
        //关闭URL请求
        curl_close($ch);
        return $data;
    }

    //返回数据
    function web_curllong($url) {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_TIMEOUT, 30);
        curl_setopt($ch, CURLOPT_USERAGENT, $this->UserAgent . $this->siteid);
        //执行命令
        $data = curl_exec($ch);
        //关闭URL请求
        curl_close($ch);
        return $data;
    }

    //des加密
    function DES1($key) {
        $this->key = $key;
    }

    function encrypt($input) {
        $size = mcrypt_get_block_size('des', 'ecb');
        $input = $this->pkcs5_pad($input, $size);
        $key = $this->key;
        $td = mcrypt_module_open('des', '', 'ecb', '');
        $iv = @mcrypt_create_iv(mcrypt_enc_get_iv_size($td), MCRYPT_RAND);
        @mcrypt_generic_init($td, $this->DESKey, $iv);
        $data = mcrypt_generic($td, $input);
        mcrypt_generic_deinit($td);
        mcrypt_module_close($td);
        $data = base64_encode($data);
        return preg_replace("/\s*/", '', $data);
    }

    function decrypt($encrypted) {
        $this->DES1($this->DESKey);
        $encrypted = preg_replace('/[\s　]/', '+', $encrypted);
        $encrypted = base64_decode($encrypted);
        $key = $this->key;
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
        $y = $this->pkcs5_unpad($decrypted);
        return $y;
    }

    function pkcs5_pad($text, $blocksize) {
        $pad = $blocksize - (strlen($text) % $blocksize);
        return $text . str_repeat(chr($pad), $pad);
    }

    function pkcs5_unpad($text) {
        $pad = ord($text{strlen($text) - 1});
        if ($pad > strlen($text)) {
            return false;
        }

        if (strspn($text, chr($pad), strlen($text) - $pad) != $pad) {
            return false;
        }

        return substr($text, 0, -1 * $pad);
    }

    //获取IP
    public function get_ip() {
        $realip = '';
        $unknown = 'unknown';
        if (isset($_SERVER)) {
            if (isset($_SERVER['HTTP_X_FORWARDED_FOR']) && !empty($_SERVER['HTTP_X_FORWARDED_FOR']) && strcasecmp($_SERVER['HTTP_X_FORWARDED_FOR'], $unknown)) {
                $arr = explode(',', $_SERVER['HTTP_X_FORWARDED_FOR']);
                foreach ($arr as $ip) {
                    $ip = trim($ip);
                    if ($ip != 'unknown') {
                        $realip = $ip;
                        break;
                    }
                }
            } else if (isset($_SERVER['HTTP_CLIENT_IP']) && !empty($_SERVER['HTTP_CLIENT_IP']) && strcasecmp($_SERVER['HTTP_CLIENT_IP'], $unknown)) {
                $realip = $_SERVER['HTTP_CLIENT_IP'];
            } else if (isset($_SERVER['REMOTE_ADDR']) && !empty($_SERVER['REMOTE_ADDR']) && strcasecmp($_SERVER['REMOTE_ADDR'], $unknown)) {
                $realip = $_SERVER['REMOTE_ADDR'];
            } else {
                $realip = $unknown;
            }
        } else {
            if (getenv('HTTP_X_FORWARDED_FOR') && strcasecmp(getenv('HTTP_X_FORWARDED_FOR'), $unknown)) {
                $realip = getenv("HTTP_X_FORWARDED_FOR");
            } else if (getenv('HTTP_CLIENT_IP') && strcasecmp(getenv('HTTP_CLIENT_IP'), $unknown)) {
                $realip = getenv("HTTP_CLIENT_IP");
            } else if (getenv('REMOTE_ADDR') && strcasecmp(getenv('REMOTE_ADDR'), $unknown)) {
                $realip = getenv("REMOTE_ADDR");
            } else {
                $realip = $unknown;
            }
        }
        $realip = preg_match("/[\d\.]{7,15}/", $realip, $matches) ? $matches[0] : $unknown;
        return $realip;
    }

}

?>