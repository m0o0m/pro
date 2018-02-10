<?php

namespace common\models;

use Yii;

class CashModel extends \yii\base\Model {

/**
	  ***********************************************************
	  *  查询总条数               @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getCount($table,$where = array()){
		return (new \yii\db\Query())
                    ->select('count(*)')
                    ->from(\Yii::$app->db->tablePrefix . $table)
                    ->where($where)
                    ->scalar();
	}


/**
      ***********************************************************
      * 查询数据		          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getData($table, $field, $where, $offset = 1, $limit = 100) {
        return (new \yii\db\Query())
                ->select($field)
                ->from(\Yii::$app->db->tablePrefix . $table)
                ->where($where)
                ->offset($offset)
                ->limit($limit)
                ->orderBy(['id' => SORT_DESC])
                ->all();
    }


/**
      ***********************************************************
      *  查询mange库中的数据     @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getManageCount($table,$where = array()){
        return (new \yii\db\Query())
                    ->select('count(*)')
                    ->from(\Yii::$app->manage_db->tablePrefix . $table)
                    ->where($where)
                    ->scalar(\Yii::$app->manage_db);
    }

    public static function getManageData($table, $field, $where, $offset = 1, $limit = 100) {
        return (new \yii\db\Query())
                ->select($field)
                ->from(\Yii::$app->manage_db->tablePrefix . $table)
                ->where($where)
                ->offset($offset)
                ->limit($limit)
                ->orderBy(['id' => SORT_DESC])
                ->all(\Yii::$app->manage_db);
    }


// /lotteryApi/controllers/UserController.php中使用下面两个model
    public static function getCashRecord($where, $offset, $limit, $orderby = '') {
        return (new \yii\db\Query())
            ->from(\Yii::$app->db->tablePrefix  . 'user_cash_record')
            ->where($where)
            ->offset($offset)
            ->orderBy(['id' => SORT_DESC])
            ->limit($limit)
            ->orderby($orderby)
            ->all(\Yii::$app->db);
    }

    public static function getCashRecordCount($where) {
        return (new \yii\db\Query())
            ->from(\Yii::$app->db->tablePrefix  . 'user_cash_record')
            ->where($where)
            ->count('id', \Yii::$app->db);
    }

}