<?php

return [
    'components' => [
        // 数据库配置
        'db' => [
            'class' => 'yii\db\Connection',
            'dsn' => 'mysql:host=127.0.0.1:3306;dbname=lottery',
            'username' => 'root',
            'password' => '123456',
            'charset' => 'utf8',
            'tablePrefix' => 'my_',
        ],
        'manage_db' => [
            'class' => 'yii\db\Connection',
            'dsn' => 'mysql:host=127.0.0.1:3306;dbname=manage',
            'username' => 'root',
            'password' => '123456',
            'charset' => 'utf8',
            'tablePrefix' => 'my_',
        ],
        //redis配置
        'redis' => [
            'class' => 'yii\redis\Connection',
            'hostname' => 'localhost',
            'port' => 6379,
        //'database' => 0,
        ],
        //redis配置
        'pullredis' => [
            'class' => 'yii\redis\Connection',
            'hostname' => 'localhost',
            'port' => 6379,
            //'database' => 0,
        ],
        'mongodb' => [
            'class' => '\yii\mongodb\Connection',
            'dsn' => 'mongodb://127.0.0.1:27017/admin',
        ],
        //session配置
        'session' => [
            'class' => 'yii\redis\Session',
            'redis' => [
                'hostname' => '127.0.0.1',
                'port' => 6379,
                'database' => 1,
            ],
            'timeout' => 3600,
        ],
    ],
];
