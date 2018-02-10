<?php
namespace common\models;

use Yii;

class LineModel extends \yii\base\Model {
    public static function getTableName(){
        return \Yii::$app->db->tablePrefix . 'sys_line_list';
    }

    public static function getCount($where) {
        return (new \yii\db\Query())
            ->from(self::getTableName())
            ->where($where)
            ->count('id', \Yii::$app->db);
    }

    public static function getList($where, $offset, $limit, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy(['id'=>SORT_DESC])// 默认id倒序，如需其它非索引字段的排序 另行设计优化。
            ->all(\Yii::$app->db);
    }

    public static function getItems($where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->all(\Yii::$app->db);
    }

    public static function getOne($where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->one(\Yii::$app->db);
    }

    public static function insert($values) {
        $sql = \common\models\CommonModel::getRawSqlForMyCat(self::getTableName(), $values);
        if(!empty($sql)){
            return (new \yii\db\Query())->createCommand($sql)->execute();
        }

        return (new \yii\db\Query())
            ->createCommand()
            ->insert(self::getTableName(), $values)
            ->execute();
    }

    public static function update($set,$where) {
        return (new \yii\db\Query())
            ->createCommand()
            ->update(self::getTableName(), $set, $where)
            ->execute();
    }

    public static function del($where) {
        return (new \yii\db\Query())
            ->createCommand()
            ->delete(self::getTableName(), $where)
            ->execute();
    }
}
