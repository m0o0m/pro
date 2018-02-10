<?php
namespace our_backend\models;
use Yii;

class AdminbackModel extends \yii\base\Model {
/**
	  ***********************************************************
	  *  查询总条数               @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getCount($where = array()){
		return (new \yii\db\Query())
                    ->select('count(*)')
                    ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                    ->where($where)
                    ->scalar(\Yii::$app->manage_db);
	}
	
/**
	  ***********************************************************
	  *  查询数据           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getData($where = [], $offset = 0, $limit = 50){
		return (new \yii\db\Query())
                    ->select('*')
                    ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                    ->where($where)
                    ->offset($offset)
                    ->limit($limit)
                    ->orderBy(['id' => SORT_DESC])
                    ->all(\Yii::$app->manage_db);
	}
/**
	  ***********************************************************
	  *  获取一条数据           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function getOneData($where){
		return (new \yii\db\Query())
                    ->select('*')
                    ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                    ->where($where)
                    ->one(\Yii::$app->manage_db);
	}
	
/**
	  ***********************************************************
	  *  添加           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function insert($data){
		 return \Yii::$app->manage_db
		 		->createCommand()
                ->insert(\Yii::$app->manage_db->tablePrefix . 'agent_admin', $data)
                ->execute();
	}
/**
	  ***********************************************************
	  *  修改           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public static function update($data, $where){
		 return \Yii::$app->manage_db
		 		->createCommand()
	            ->update(\Yii::$app->manage_db->tablePrefix . 'agent_admin', $data, $where)
	            ->execute();
	}
	
	
	
	
	

}