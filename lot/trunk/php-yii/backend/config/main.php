<?php

$params = require(__DIR__ . '/params.php');
return [
    'id' => 'agent_backend',
    'name' => 'Agent-Admin',
    'basePath' => dirname(__DIR__),
    'controllerNamespace' => 'backend\controllers',
    'bootstrap' => ['log'],
    'modules' => [
        'base' => [
            'class' => 'backend\modules\base\BaseModule',
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
    'defaultRoute' => 'login/login', //通过域名访问默认路由
];
