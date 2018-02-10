<?php
namespace our_backend\models;

use Yii;

class MaintainModel extends \yii\base\Model {

    private static function getCount($table, $where) {
        return (new \yii\db\Query())
            ->from($table)
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
    }

    private static function getList($table, $where, $offset, $limit, $orderby = '') {
        return (new \yii\db\Query())
            ->from($table)
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy($orderby)
            ->all(\Yii::$app->manage_db);
    }

    private static function getItems($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->all(\Yii::$app->manage_db);
    }

    private static function getOne($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }

    private static function insert($table, $values) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->insert($table, $values)
            ->execute();
    }

    private static function update($table, $set,$where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->update($table, $set, $where)
            ->execute();
    }

    private static function del($table, $where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->delete($table, $where)
            ->execute();
    }

    public static function getTableName($type = ''){
        switch($type){
            default:
                $tab = 'maintain_data';
                break;
        }
        return \Yii::$app->manage_db->tablePrefix . $tab;
    }

    public static function get_count($where = []) {
        return self::getCount(self::getTableName(), $where);
    }

    public static function get_list($where, $offset, $limit, $orderby = '') {
        return self::getList(self::getTableName(), $where, $offset, $limit, $orderby);
    }

    public static function get_all($where = [], $select = '*') {
        return self::getItems(self::getTableName(), $where, $select);
    }

    public static function get_one($where = [], $select = '*') {
        return self::getOne(self::getTableName(), $where, $select);
    }

    public static function _insert($values) {
        return self::insert(self::getTableName(), $values);
    }

    public static function _update($set, $where) {
        return self::update(self::getTableName(), $set, $where);
    }

    public static function _del($where) {
        return self::del(self::getTableName(), $where);
    }

}
