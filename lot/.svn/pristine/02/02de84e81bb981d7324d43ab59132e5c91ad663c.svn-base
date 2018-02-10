<?php

namespace our_backend\models;

use Yii;
use yii\db\Query;

class LineModel extends \yii\base\Model {

    /**
     * **********************************************************
     *  获取总条数               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function countLin($where) {
        return (new \yii\db\Query())
                        ->select('count(id)')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where($where)
                        ->scalar();
    }

    /**
     * **********************************************************
     * 查询所有线路信息          @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function queryLin($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['updatetime' => SORT_DESC])
                        ->all();
    }

    /**
     * **********************************************************
     *  查询原有额度             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function oldmoney($id) {
        return (new \yii\db\Query())
                        ->select('money')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where(['id' => $id])
                        ->scalar();
    }

    /**
     * **********************************************************
     *  单字段更改               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function updateone($field, $val, $where) {
        return (new \yii\db\Query())
                        ->createCommand()
                        ->update(\Yii::$app->db->tablePrefix . 'sys_line_list', array($field => $val), $where)
                        ->execute();
    }

    /**
     * **********************************************************
     *  线路基本信息(页面表单显示)@author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function lineForm($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where($where)
                        ->one();
    }

    /**
     * **********************************************************
     *              线路基本信息 修改( 提交表单)                *
     * **********************************************************
     */
    public static function baseUpdate($where, $data) {
        return (new \yii\db\Query())
                        ->createCommand()
                        ->update(\Yii::$app->db->tablePrefix . 'sys_line_list', $data, $where)
                        ->execute();
    }

    /**
     * **********************************************************
     *                             新增					                *
     * **********************************************************
     */
    public static function add($table, $data) {
        // $sql = \common\models\CommonModel::getRawSqlForMyCat(\Yii::$app->db->tablePrefix . $table, $data);
        // if($sql){
        //     return (new \yii\db\Query())->createCommand($sql)->execute();
        // }

        return (new \yii\db\Query())
                        ->createCommand()
                        ->insert(\Yii::$app->db->tablePrefix . $table, $data)
                        ->execute();
    }

}
