<?php

$params = require(__DIR__ . '/params.php');
return [
    'id' => 'our_backend',
    'name' => 'PK-Admin',
    'basePath' => dirname(__DIR__),
    'controllerNamespace' => 'our_backend\controllers',
    'language' => 'zh-CN',
    'components' => [
        'request' => [
            // !!! insert a secret key in the following (if it is empty) - this is required by cookie validation
            'enableCookieValidation' => false,
            'enableCsrfValidation' => false,
        ],
    ],
    'params' => $params,
    'defaultRoute' => 'login/login', //通过域名访问默认路由
];
