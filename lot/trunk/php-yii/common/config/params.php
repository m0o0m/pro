<?php

// 公共配置 Yii::$app->params[''];
return [
    'tablePrefix' => 'my_',
    'hostdir'  => 'http://113.10.246.106/date/wwwuser/public_html/', // 日志文件目录
    'basename' => 'pklottery',
    // 钱包地址
    'tcphost' => '113.10.246.105',
    'tcpport' => '9998',
    // 彩票图片地址
    'cdn_href' => 'http://pkcdn.pk1358.com',
    // 判断是不是mycat
    'is_mycat' => false,
    // 刷新线路缓存(refresh/key) 外部接入(golang接口)
    'golangApi' => 'http://113.10.246.105:8084',
    // 刷新线路缓存(refresh/key) 外部采集(golang接口)
    'golangSpider' => 'http://113.10.246.105:9898',
    // 用于处理福彩3D和排列3的期数及开封盘时间获取，如果在农历初一之前，为false,农历初七之后改成true 结算模块到配置文件同步更改,初一到初七挂维护
    'fc_3d_pl_3' => false,

    // Worker App
    'app_stat_host_http'        => '113.10.246.106:2346',    // 注单统计
    'app_push_host_ws'          => '113.10.246.106:9527',    // 消息推送
    'app_push_host_http'        => '113.10.246.106:9528',    // 消息推送
    'app_lottery_host_http'     => '113.10.246.106:10002',   // 注单结算
    'app_spider_host_http'      => '113.10.246.106:10004',   // 注单采集
];
