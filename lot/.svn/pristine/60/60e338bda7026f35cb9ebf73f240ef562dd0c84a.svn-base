<?php

namespace common\helpers;

class mongoTables {

    //获取表名
    public static function getTable($key) {
        $allTables = self::allTables();

        return $allTables[$key];
    }

    //所有mongodb表必须从这添加
    public static function allTables() {
        //取消分表
        // return [
        //     'httpRequest' => 'requestLog_' . date('Ymd'), //请求日志表
        //     'historyLogin' => 'histonryLoginLog_' . date('Ym'), //历史登陆
        //     'operate' => 'operateLog_' . date('Ym'), //操作日志
        // ];
        return [
            'httpRequest' => 'request_Log', //请求日志表
            'historyLogin' => 'histonryLogin_Log', //历史登陆
            'operate' => 'operate_Log', //操作日志
            'auto'    => 'auto', //开奖网开奖结果
            'bets'    => 'bets', //注单备份
        ];
    }

}

?>