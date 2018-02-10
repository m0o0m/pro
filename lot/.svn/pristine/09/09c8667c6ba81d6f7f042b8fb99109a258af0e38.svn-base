<?php

namespace our_backend\models;

use Yii;

class SpiderbetModel extends \yii\base\Model {

/**
	  ***********************************************************
	  *  查询总条数               @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getCount($where = array()){
		return (new \yii\db\Query())
                    ->select('count(id)')
                    ->from(\Yii::$app->manage_db->tablePrefix . 'spider_record')
                    ->where($where)
                    ->scalar(\Yii::$app->manage_db);
	}


/**
      ***********************************************************
      * 查询数据		          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getData($field, $where, $offset = 1, $limit = 100) {
        return (new \yii\db\Query())
                ->select($field)
                ->from(\Yii::$app->manage_db->tablePrefix . 'spider_record')
                ->where($where)
                ->offset($offset)
                ->limit($limit)
                ->orderBy(['id' => SORT_DESC])
                ->all(\Yii::$app->manage_db);
    }

/**
      ***********************************************************
      *  添加纪录           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
     public static function insert($arr) {
        return Yii::$app->manage_db->createCommand()->insert(
                        \Yii::$app->manage_db->tablePrefix . 'spider_record', $arr
                )->execute();
    }
    
    
}