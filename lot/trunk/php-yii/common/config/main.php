<?php

$params = require(__DIR__ . '/params.php');
return [
    'vendorPath' => dirname(dirname(__DIR__)) . '/vendor',
    'bootstrap' => ['debug'],
    'modules' => [
        'debug' => [
            'class' => 'yii\debug\Module',
            'allowedIPs' => ['127.0.0.1','96.9.85.113','96.9.85.114','96.9.85.74','96.9.85.75'],
        ],
    ],
    'components' => [
        // 缓存配置
        'cache' => [
            'class' => 'yii\caching\FileCache',
        ],
        // 路由配置
        'urlManager' => [
            'enablePrettyUrl' => true,
            'showScriptName' => false,
            'rules' => [
                'debug/<controller>/<action>' => 'debug/<controller>/<action>',
                '<controller:\w+>/<id:\d+>' => '<controller>/view',
                '<controller:\w+>/<action:\w+>/<id:\d+>' => '<controller>/<action>',
                '<controller:\w+>/<action:\w+>' => '<controller>/<action>',
            ],
        ],
        // smarty 模板
        'view' => [
            'renderers' => [
                'html' => [
                    'class' => 'yii\smarty\ViewRenderer',
                    'options' => [
                        'left_delimiter' => '<{',
                        'right_delimiter' => '}>',
                    ],
                ],
            ],
        ],
        // 多语言配置
        'i18n' => [
            'translations' => [
                'app' => [
                    'class' => 'yii\i18n\PhpMessageSource',
                    'basePath' => '@common/messages',
                    'fileMap' => [
                        'app' => 'app.php',
                        'app/error' => 'error.php',
                    ],
                ],
            ],
        ],
        // 日志
        'log' => [
            'traceLevel' => YII_DEBUG ? 3 : 0,
            'targets' => [
                [
                    'class' => 'yii\log\FileTarget',
                    'levels' => ['error', 'warning'],
                ],
            ],
        ],
    ],
    'params' => $params
];
