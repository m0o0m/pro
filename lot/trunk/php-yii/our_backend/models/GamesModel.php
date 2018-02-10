<?php
namespace our_backend\models;

use Yii;
use yii\db\Query;

class GamesModel extends \yii\base\Model {
//此model用于彩种分类管理和采种管理
/**
	  ***********************************************************
	  *  获取表名      			  @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	 public static function getTableName(){
	        return \Yii::$app->manage_db->tablePrefix . 'fc_games';
	    }

   	 public static function getSubTableName(){
        return \Yii::$app->manage_db->tablePrefix . 'fc_sub';
     }
/**
	  ***********************************************************
	  *  查询总条数               @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getCount($table,$where = array()){
		return (new \yii\db\Query())
                    ->select('count(*)')
                    ->from($table)
                    ->where($where)
                    ->scalar(\Yii::$app->manage_db);
	}


/**
      ***********************************************************
      * 查询数据		          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
	public static function getData($table,$orderby,$field, $where, $offset = 0, $limit = 100) {
        return (new \yii\db\Query())
                ->select($field)
                ->from($table)
                ->where($where)
                ->offset($offset)
                ->limit($limit)
                ->orderby($orderby)
                ->all(\Yii::$app->manage_db);
    }
/**
      ***********************************************************
      *  	插入数据		          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function insert($table, $data) {
        return (new \yii\db\Query())
                        ->createCommand(\Yii::$app->manage_db)
                        ->insert($table, $data)
                        ->execute();
    }
/**
      ***********************************************************
      *  	更新数据		          @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function update($table, $set, $where) {
        return (new \yii\db\Query())
                        ->createCommand(\Yii::$app->manage_db)
                        ->update($table, $set, $where)
                        ->execute();
    }


}