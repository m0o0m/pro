<?php

use \GatewayWorker\Lib\Gateway;
use \Applications\Push\Lib\Common;

class Events {

    public static function onWorkerStart($worker) {
        Common::init($worker); // 初始化
        Common::lefttime($worker); // 封盘时间推送
    }

    public static function onMessage($client_id, $message) {
        if (isset($message['get']) && !empty($message['get'])) {
            $message_data = $message['get']; // http
        } elseif (isset($message['post']) && !empty($message['post'])) {
            $message_data = $message['post']; // http
        } else {
            $message_data = json_decode($message, true); // ws
        }
        $cmd = isset($message_data['cmd']) ? $message_data['cmd'] : '';
        switch ($cmd) {
            case 'join': // 用户分组
                Common::join($client_id, $message_data);
                break;
            case 'auto': // 开奖结果推送
                Common::auto($client_id, $message_data);
                break;
            case 'online': // 在线人数推送
                Common::online($client_id, $message_data);
                break;
            default:
                break;
        }
    }

}
