<?php

namespace common\helpers;

use Yii;

class TcpConPoll {

    private static $Instace = "";           ///对象
    private static $socket = "";                     ///tcp连接池链接

    //构造函数

    private function __construct() {
        $host = isset(Yii::$app->params['tcphost']) ? Yii::$app->params['tcphost'] : "127.0.0.1";
        $port = isset(Yii::$app->params['tcpport']) ? Yii::$app->params['tcpport'] : "9998";
        static::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP)or die("Could not create  socket\n");
        $connection = socket_connect(static::$socket, $host, $port) or die("Could not connet server\n");
    }

    public static function send($str) {
        $length = bytes::integertobytes(intval(strlen($str)));
        $payitem = bytes::getbytes($str);
        $return_betys = array_merge($length, $payitem);
        $msg = bytes::tostr($return_betys);
        socket_write(static::$socket, $msg) or die("Write failed\n");
        //读取长度
        if ($buff = socket_read(static::$socket, 4, PHP_NORMAL_READ)) {
            $payitem1 = bytes::getbytes($buff);
            $len = bytes::bytesToInteger($payitem1, 0);
            if ($buff = socket_read(static::$socket, $len, PHP_NORMAL_READ)) {
                return $buff;
            }
        }
    }

    public static function getInstace() {
        if (empty(static::$Instace)) {
            static::$Instace = new TcpConPoll();
        }
        return static::$Instace;
    }

    public function close() {
        socket_close(static::$socket);
    }

    public function __destruct() {
        $this->close();
    }

}

?>