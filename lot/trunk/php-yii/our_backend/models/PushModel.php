<?php
namespace our_backend\models;

use Yii;

class PushModel extends \yii\base\Model {
    public static function getCount($where, $isqueue = false) {
        return (new \yii\db\Query())
            ->from(self::getTableName($isqueue))
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
    }

    public static function getList($where, $offset, $limit, $select = '*', $isqueue = false) {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName($isqueue))
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy(['id'=>SORT_DESC])// 默认id倒序，如需其它非索引字段的排序 另行设计优化。
            ->all(\Yii::$app->manage_db);
    }

    public static function getItems($where, $select = '*', $isqueue = false) {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName($isqueue))
            ->where($where)
            ->all(\Yii::$app->manage_db);
    }

    public static function getOne($where, $select = '*', $isqueue = false) {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName($isqueue))
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }

    public static function insert($values, $isqueue = false) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->insert(self::getTableName($isqueue), $values)
            ->execute();
    }

    public static function update($set,$where, $isqueue = false) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->update(self::getTableName($isqueue), $set, $where)
            ->execute();
    }

    public static function delete($where, $isqueue = false) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->delete(self::getTableName($isqueue), $where)
            ->execute();
    }

    public static function getTableName($isqueue){
        return $isqueue ? \Yii::$app->manage_db->tablePrefix . 'push_queue' : \Yii::$app->manage_db->tablePrefix . 'push_list';
    }
}
