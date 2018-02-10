<?php
namespace common\models;

use Yii;

class StatModel extends \yii\base\Model {
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
            case 'line':
                $tab = 'stat_data_line_bet';
                break;
            default:
                $tab = 'stat_data_bet';
                break;
        }
        return \Yii::$app->manage_db->tablePrefix . $tab;
    }

    public static function get_list($where, $offset, $limit, $orderby = '') {
        return self::getList(self::getTableName(), $where, $offset, $limit, $orderby);
    }

    public static function get_list_line($where, $offset, $limit, $orderby = '') {
        return self::getList(self::getTableName('line'), $where, $offset, $limit, $orderby);
    }

    public static function get_count($where) {
        return self::getCount(self::getTableName(), $where);
    }

    public static function get_count_line($where) {
        return self::getCount(self::getTableName('line'), $where);
    }

    public static function get_items($where) {
        return self::getItems(self::getTableName('line'), $where);
    }
/**
      ***********************************************************
      *  查询统计数据           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getSum($field, $where, $tab = ''){
        $res = (new \yii\db\Query())
            ->select($field)
            ->from(self::getTableName($tab))
            ->where($where)
            ->scalar(\Yii::$app->manage_db);
        if(!$res) $res = 0;
        return $res;
    }
    
    
}
