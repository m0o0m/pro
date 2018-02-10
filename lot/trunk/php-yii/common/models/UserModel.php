<?php

namespace common\models;

use Yii;

/**
 * 会员model(总后台,业主后台,前台)  共用
 */
class UserModel extends \yii\base\Model {

    //获取会员条数
    public static function getUserCount($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where($where)
                        ->count('uid', \Yii::$app->db);
    }

    //获取会员列表
    public static function getUserList($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['uid' => SORT_DESC])
                        ->all(\Yii::$app->db);
    }

    //获取一条会员信息
    public static function getOneUser($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where($where)
                        ->one(\Yii::$app->db);
    }

    //修改会员信息
    public static function updateUser($data,$where){
        return \Yii::$app->db
        ->createCommand()
        ->update(\Yii::$app->db->tablePrefix . 'user', $data, $where)
        ->execute();

    }

    //获取问题总数
    public static function getCount($table,$map) {
        return (new \yii\db\Query())
                        ->from($table)
                        ->where($map)
                        ->count('id', \Yii::$app->manage_db);
    }

    //更改状态
    public static function upState($table,$data,$where){
        return \Yii::$app->db
        ->createCommand()
        ->update($table, $data, $where)
        ->execute();
    }

    public static function getTableName() {
        return \Yii::$app->db->tablePrefix . 'user';
    }

    public static function getItems($table, $where) {
        return (new \yii\db\Query())
            ->select('*')
            ->from($table)
            ->where($where)
            ->orderBy(['uid' => SORT_DESC])
            ->all(\Yii::$app->db);
    }

    public static function get_items($where) {
        return self::getItems(self::getTableName(), $where);
    }

    //获取代理
    public static function getAgent($where, $offset = 0, $limit = 100) {
        return (new \yii\db\Query())
                        ->select('id,login_user')
                        ->from(\Yii::$app->db->tablePrefix . 'user_agent')
                        ->offset($offset)
                        ->limit($limit)
                        ->where($where)
                        ->all(\Yii::$app->db);
    }

    //获取所属线路金额
    public static function getLineMoney($where) {
        return (new \yii\db\Query())
                        ->select('money')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where($where)
                        ->scalar(\Yii::$app->db);
    }
    //更新线路金额
    public static function updateLineMoney($money, $where) {
        return (new \yii\db\Query())
                        ->createCommand()
                        ->update(\Yii::$app->db->tablePrefix . 'sys_line_list', array('money' => $money), $where)
                        ->execute();
    }

    //写入现金记录
    public static function insertCashRecord($data) {
        return Yii::$app->db
                        ->createCommand()
                        ->Insert(\Yii::$app->db->tablePrefix . 'user_cash_record', $data)
                        ->execute();
    }

    /**
      ***********************************************************
      *  原生sql                 @author ruizuo qiyongsheng    *
      ***********************************************************
    */
    public static function insert($sql){
        return \Yii::$app->db
                ->createCommand($sql)
                ->execute();
    }
}
