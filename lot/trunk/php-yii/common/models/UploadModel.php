<?php
namespace common\models;

use Yii;

class UploadModel extends \yii\base\Model {
    public static function getCount($where) {
        return (new \yii\db\Query())
            ->from(self::getTableName())
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
    }

    public static function getList($where, $offset, $limit, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy(['id'=>SORT_DESC])// 默认id倒序，如需其它非索引字段的排序 另行设计优化。
            ->all(\Yii::$app->manage_db);
    }

    public static function getItems($where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->all(\Yii::$app->manage_db);
    }

    public static function getOne($where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }

    public static function insert($values) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->insert(self::getTableName(), $values)
            ->execute();
    }

    public static function update($set,$where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->update(self::getTableName(), $set, $where)
            ->execute();
    }

    public static function delete($where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->delete(self::getTableName(), $where)
            ->execute();
    }

    public static function getTableName(){
        return \Yii::$app->manage_db->tablePrefix . 'upload_list';
    }
}
