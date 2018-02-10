<?php
namespace our_backend\models;

use Yii;

class SendmsgModel extends \yii\base\Model {
/**
     * **********************************************************
     *  获取表名                    @author ruizuo qiyongsheng    *
     * **********************************************************
*/	
 public static function getTableName(){
        return \Yii::$app->manage_db->tablePrefix . 'send_msg';
 }
/**
     * **********************************************************
     *  查询总条数           		  @author ruizuo qiyongsheng    *
     * **********************************************************
*/
	public static function getCount($where){
		 return (new \yii\db\Query())
            ->from(self::getTableName())
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
	}

/**
     * **********************************************************
     *  查询列表数据            	  @author ruizuo qiyongsheng    *
     * **********************************************************
*/
	public static function getList($where, $offset, $limit, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from(self::getTableName())
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->all(\Yii::$app->manage_db);
    }


/**
     * **********************************************************
     *  添加推送消息            	  @author ruizuo qiyongsheng    *
     * **********************************************************
*/
	public static function insert($data){
		return \Yii::$app->manage_db
            ->createCommand()
            ->insert(self::getTableName(), $data)
            ->execute();
	}
}
