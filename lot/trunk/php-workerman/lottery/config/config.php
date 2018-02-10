<?php

namespace config;

class config {

    //mysql 表前缀
    public static $tablePrefix = 'my_';
    /*     * *************** 内测-测试环境   ***************** */
    // 数据库实例1(彩票库)
    /* public static $lottery = array(
      'host' => '163.47.173.73',
      'user' => 'new',
      'password' => 'q#*(sfsd&#$*#Y435',
      'dbname' => 'lottery',
      'port' => '3309',
      'charset' => 'utf8',
      );
      // 数据库实例2(管理库)
      public static $manage = array(
      'host' => '163.47.173.73',
      'user' => 'new',
      'password' => 'q#*(sfsd&#$*#Y435',
      'dbname' => 'manage',
      'port' => '3309',
      'charset' => 'utf8',
      );
      //redis
      public static $redis = array(
      'host' => '127.0.0.1',
      'port' => '6379'
      );
      //mongodb
      public static $mongo = array(
      'user' => 'cpuser',
      'pass' => 'PK_CYJ2016163.COM',
      'host' => '163.47.173.73',
      'port' => '27017',
      'db' => 'caipiaodb'
      ); */

    /*     * *************** 公测-测试环境-Mycat   ***************** */
    // 数据库实例1(彩票库)
    // public static $lottery = array(
    // 'host' => '113.10.246.111',
    // 'user' => 'lottery',
    // 'password' => 'tjVBd&RfWX0Y',
    // 'dbname' => 'lottery',
    // 'port' => '3306',
    // 'charset' => 'utf8',
    // );
    // // 数据库实例2(管理库)
    // public static $manage = array(
    // 'host' => '113.10.246.111',
    // 'user' => 'manage',
    // 'password' => 'ATRwSxdn#8vV',
    // 'dbname' => 'manage',
    // 'port' => '3307',
    // 'charset' => 'utf8',
    // );
    // //redis
    // public static $redis = array(
    // 'host' => '127.0.0.1',
    // 'port' => '6379',
    // 'password' => ''
    // );
    // //mongodb
    // public static $mongo = array(
    // 'user' => 'cpuser',
    // 'pass' => 'PK_CYJ2016163.COM',
    // 'host' => '163.47.173.73',
    // 'port' => '27017',
    // 'db' => 'caipiaodb'
    // ); 

    /*     * *************** 本地*************** */
    // 数据库实例1(彩票库)
    public static $lottery = array(
        'host' => '127.0.0.1',
        'user' => 'root',
        'password' => '123456',
        'dbname' => 'lottery',
        'port' => '3306',
        'charset' => 'utf8',
    );
    // 数据库实例2(管理库)
    public static $manage = array(
        'host' => '127.0.0.1',
        'user' => 'root',
        'password' => '123456',
        'dbname' => 'manage',
        'port' => '3306',
        'charset' => 'utf8',
    );
    //redis
    public static $redis = array(
        'host' => '127.0.0.1',
        'port' => '6379',
        'password' => ''
    );
    //采集redis
    public static $pullredis = array(
        'host' => '127.0.0.1',
        'port' => '6379',
        'password' => '',
    );
    //mongodb
    public static $mongo = array(
        'user' => '',
        'pass' => '',
        'host' => '127.0.0.1',
        'port' => '27017',
        'db' => 'admin'
    );
    //tcp 钱包配置
    public static $wallet = array(
        'host' => '127.0.0.1',
        'port' => 9998
    );
    //判断是公测环境还是内测环境
    public static $is_mycat = false;
    //推送开奖结果的ip
    public static $sendAuto = '113.10.246.106:9528';
    //结算服务器ip
    public static $balance = '127.0.0.1';

}

?>