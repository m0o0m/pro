<?php

namespace frontend_api\controllers;

use Yii;
use frontend_api\controllers\Controller;
use \common\helpers\Helper;

/**
 * Games controller
 */
class GamesController extends Controller {

    private $DESKey = 'DCEFDADB';
    private $MD5Key = 'cdcddsfggg';
    private $key = '';
    private $apiurl = 'http://127.0.0.1:8084';
    private $UserAgent = "Lottery_Api_Go_";
    private $siteid = "aaa";

    public function actionTest() {
        $this->CreateMemberAndForwardGame();
    }

    //注册
    public function CreateMemberAndForwardGame() {
        $crypt = $this->DES1($this->DESKey);

        $ip = $this->get_ip();

        $userinfo = [
            'site_id' => 'aaa',
            'uname' => 'test134',
            'pwd' => '123456',
            'agent_name' => 'at_01',
        ];

        //'http://127.0.0.1:8084/account/create?site_id=aaa&uname=bb6&pwd=123456&agent_name=at_01';

        $params = $this->encrypt(
                "site_id=" . $userinfo['site_id'] .
                "/\\\\/uname=" . $userinfo['uname'] .
                "/\\\\/pwd=" . $userinfo['pwd'] .
                "/\\\\/agent_name=" . $userinfo['agent_name']
        );

        $Key = $this->getKey($params);
        $url = $this->apiurl . "/account/create?params=" . $params . "&key=" . $Key;

        $result = $this->web_curl($url);

        var_dump($result);
        die;

        return $result;
    }

    //登录
    public function login() {
        
    }

    //获取余额
    public function GetBalance($loginname, $gtype) {
        $crypt = $this->DES1($this->DESKey);
        $params = $this->encrypt(
                "siteid=" . $this->siteid .
                "/\\\\/gtype=" . $gtype .
                "/\\\\/username=" . $loginname);
        $Key = $this->getKey($params);

        $url = $this->apiurl . "/GetBalance?params=" . $params . "&key=" . $Key;

        $result = $this->web_curl($url);
        return $result;
    }

    //额度转换
    public function TransferCredit($userinfo, $gtype, $type, $credit, $cur = "RMB", $lang = "zh", $media = "pc") {
        $crypt = $this->DES1($this->DESKey);
        if ($gtype == "lebo" || $gtype == "bbin") {
            $lang = "zh-cn";
        }
        if ($gtype == "bbin") {
            $gametype = "";
        }
        if ($gtype == 'sa') {
            $cur = "CNY";
        }
        if (empty($userinfo['username']) || empty($userinfo['agent_id']) ||
                empty($userinfo['index_id']) || empty($userinfo['sh_id']) ||
                empty($userinfo['ua_id'])) {
            return "用户信息错误";
        }
        $ip = $this->get_ip();

        $params = $this->encrypt(
                "siteid=" . $this->siteid .
                "/\\\\/username=" . $userinfo['username'] .
                "/\\\\/type=" . $type .
                "/\\\\/credit=" . $credit .
                "/\\\\/gtype=" . $gtype .
                "/\\\\/agent_id=" . $userinfo['agent_id'] .
                "/\\\\/index_id=" . $userinfo['index_id'] .
                "/\\\\/sh_id=" . $userinfo['sh_id'] .
                "/\\\\/ua_id=" . $userinfo['ua_id'] .
                "/\\\\/cur=" . $cur .
                "/\\\\/limit=" . $userinfo['limit'] .
                "/\\\\/lang=" . $lang .
                "/\\\\/ip=" . $ip .
                "/\\\\/sw=" . $userinfo['sw']);

        $Key = $this->getKey($params);

        $url = $this->apiurl . "/TransferCredit?params=" . $params . "&key=" . $Key;

        $result = $this->web_curl($url);
        return $result;
    }

    //返回{"result":true,"data":{"Code":0}}为修改成功
    public function EditAccountPwd($gtype, $loginname, $pwd) {
        $crypt = $this->DES1($this->DESKey);
        $params = $this->encrypt(
                "siteid=" . $this->siteid .
                "/\\\\/username=" . $loginname .
                "/\\\\/password=" . $pwd .
                "/\\\\/gtype=" . $gtype);

        $Key = $this->getKey($params);
        $url = $this->apiurl . "/EditUserPassword?params=" . $params . "&key=" . $Key;
        $result = $this->web_curl($url);
        return $result;
    }

    public function getKey($params) {
        return md5($params . $this->MD5Key);
    }

    //返回数据
    public function web_curl($url) {
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
    public function DES1($key) {
        $this->key = $key;
    }

    public function encrypt($input) {
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

    public function decrypt($encrypted) {
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

    public function pkcs5_pad($text, $blocksize) {
        $pad = $blocksize - (strlen($text) % $blocksize);
        return $text . str_repeat(chr($pad), $pad);
    }

    public function pkcs5_unpad($text) {
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