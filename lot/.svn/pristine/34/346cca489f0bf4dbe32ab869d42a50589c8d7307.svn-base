<?php
namespace our_backend\models;

use Yii;

class TaocanModel extends \yii\base\Model {
    public static function getCount($table, $where) {
        return (new \yii\db\Query())
            ->from($table)
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
    }

    public static function getList($table, $where, $offset, $limit, $orderby = []) {
        return (new \yii\db\Query())
            ->from($table)
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy($orderby)
            ->all(\Yii::$app->manage_db);
    }

    public static function getItems($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->all(\Yii::$app->manage_db);
    }

    public static function getOne($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }

    public static function insert($table, $values) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->insert($table, $values)
            ->execute();
    }

    public static function update($table, $set,$where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->update($table, $set, $where)
            ->execute();
    }

    public static function del($table, $where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->delete($table, $where)
            ->execute();
    }

    public static function getTableName($type = ''){
        switch($type){
            case 'set':
                $tab = 'line_pack';
                break;
            default:
                $tab = 'pack_list';
                break;
        }
        return \Yii::$app->manage_db->tablePrefix . $tab;
    }

    public static function get_count($where) {
        return self::getCount(self::getTableName(), $where);
    }

    public static function get_set_count($where) {
        return self::getCount(self::getTableName('set'), $where);
    }

    public static function get_all($where) {
        return self::getItems(self::getTableName(), $where);
    }

    public static function get_set_all($where) {
        return self::getItems(self::getTableName('set'), $where);
    }

    public static function get_one($where) {
        return self::getOne(self::getTableName(), $where);
    }

    public static function get_set_one($where) {
        return self::getOne(self::getTableName('set'), $where);
    }

    public static function _insert($values) {
        return self::insert(self::getTableName(), $values);
    }

    public static function set_insert($values) {
        return self::insert(self::getTableName('set'), $values);
    }

    public static function _update($set, $where) {
        return self::update(self::getTableName(), $set, $where);
    }

    public static function set_del($where) {
        return self::del(self::getTableName('set'), $where);
    }
}
