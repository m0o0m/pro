<?php
namespace Applications\Common\Config;

class Config
{
    public static $prefix = 'my_';
    public static $lottery = array(
        'host' => '127.0.0.1',
        'port' => '3306',
        'user' => 'root',
        'password' => '',
        'dbname' => 'lottery'
    );
    public static $manage = array(
        'host' => '127.0.0.1',
        'port' => '3306',
        'user' => 'root',
        'password' => '',
        'dbname' => 'manage'
    );
    public static $redis = array(
        'host' => '127.0.0.1',
        'port' => '6379',
        'password' => ''
    );
    public static $upload = array(
        'basepath' => '/date/wwwuser/public_html', // 上传文件基本路径 // 线上使用
        // 'basepath' => '/Users/frank', // 上传文件基本路径 // 本地测试
        'upfolder' => 'upload', // 上传文件夹
        'max_file_size' => 210241024,//2M // 上传文件大小限制, 单位BYTE
        'uptypes' => [ // 上传文件类型列表
            'image/jpg',
            'image/jpeg',
            'image/png',
            'image/pjpeg',
            'image/gif',
            'image/bmp',
            'image/x-png',
            'application/x-shockwave-flash'
        ],
        'overwrite' => false // 同名是否覆盖
    );
}
