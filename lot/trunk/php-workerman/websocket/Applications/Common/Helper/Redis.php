<?php

namespace Applications\Common\Helper;

use \Applications\Common\Config\Config;

class Redis {

    private static $Instace = "";           ///对象
    private $sock = "";                     ///Redis连接池链接
    private $type = 1;                      ///连接池模式
    private static $config = array(
    );

    private function __construct($database) {
        if ($database == 'pull') {
            self::$config = config::$pullredis;
        } else {
            self::$config = config::$redis;
        }
        if (self::$config['host'] && self::$config['port']) {
            $this->SockForRedis($database);
        } else {
            @$this->sock = fsockopen(self::$config['host'], self::$config['port'], $errno, $errstr, 1);

            if (!$this->sock) {
                $this->SockForRedis($database);
                ///error("$errstr ($errno)\n",3,'./cyj_web/cache/error.log');       ////连接连接池失败
            } else {
                $this->type = 1;
            }
        }
    }

    public function __call($func, $args) {
        if ($this->type == 1) {
            if ($this->BaseWrite(json_encode(['a' => $func, 'p' => $args]))) {
                $data = $this->BaseRead();
                if ($data !== FALSE) {
                    $data = json_decode($data, TRUE);
                    if ($data)
                        return $data['i'];
                    else
                        return FALSE;
                }else {
                    return FALSE;
                }
            } else {
                return FALSE;
            }
        } else {

            return call_user_func_array(array($this->sock, $func), $args);
        }
    }

    private function BaseWrite($data) {
        $data .= "\n";
        $sumlen = strlen($data);
        $len = fwrite($this->sock, $data);
        if ($sumlen == $len) {
            return TRUE;
        } else {
            return FALSE;
        }
    }

    private function BaseRead() {
        $buffer = '';

        while (!feof($this->sock)) {
            $buffer .= fgets($this->sock, 128);
            $pos = strpos($buffer, "\n");
            if ($pos !== false) {
                $buffer = trim($buffer);
                break;
            }
        }

        if (strlen($buffer) > 0) {
            return $buffer;
        } else {
            return FALSE;
        }
    }

    private function SockForRedis($database) {                                   ////直接调用redis
        $this->sock = new \Redis();
        $this->sock->connect(self::$config['host'], self::$config['port']);
        $this->sock->auth(self::$config['password']);
        if ($database) {
            $this->sock->select($database);
        }
        $this->type = 0;
    }

    public static function instance($database = '') {
        if (empty(static::$Instace)) {
            static::$Instace = new Redis($database);
        }
        return static::$Instace;
    }

    public static function delete($key) {
        return static::$Instace->del($key);
    }

    public static function lsize($key) {
        return static::$Instace->llen($key);
    }

    //压缩
    public function set($key, $value) {
        return $this->sock->set($key, gzcompress($value));
    }

    //解压
    public function get($key) {
        $res = $this->sock->get($key);
        if ($res === NULL) {
            return false;
        }
        if ($res === false) {
            return false;
        }
        return gzuncompress($res);
    }

    //压缩
    public function setex($key, $time, $value) {
        return $this->sock->setex($key, $time, gzcompress($value));
    }

    public function close() {
        if ($this->type == 1) {
            fclose($this->sock);
        } else {
            $this->sock->close();
        }
        static::$Instace = "";
    }

    public function __destruct() {
        $this->close();
    }

}
