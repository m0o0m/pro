<?php

$params = require(__DIR__ . '/params.php');
return [
    'id' => 'frontend_api',
    'name' => 'pc',
    'basePath' => dirname(__DIR__),
    'controllerNamespace' => 'frontend_api\controllers',
    'bootstrap' => ['log'],
    'modules' => [
        'base' => [
            'class' => 'frontend_api\modules\base\BaseModule',
        ],
    ],
    'language' => 'zh-CN',
    'components' => [
        'request' => [
            // !!! insert a secret key in the following (if it is empty) - this is required by cookie validation
            'enableCookieValidation' => false,
            'enableCsrfValidation' => false,
        ],
    ],
    'params' => $params,
    'defaultRoute' => 'site/index', //通过域名访问默认路由
];
