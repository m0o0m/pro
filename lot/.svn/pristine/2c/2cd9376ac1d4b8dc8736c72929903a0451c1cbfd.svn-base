<?php
namespace our_backend\models;

use Yii;

class OtherModel extends \yii\base\Model {
/**
     * **********************************************************
     *  获取表名                    @author ruizuo qiyongsheng    *
     * **********************************************************
*/	
     public static function getTableName(){
            return \Yii::$app->manage_db->tablePrefix . 'api_whitelist';
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
     *  查询一条数据                @author ruizuo qiyongsheng    *
     * **********************************************************
*/
    public static function getOneData($where) {
        return (new \yii\db\Query())
            ->select('*')
            ->from(self::getTableName())
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }
/**
     * **********************************************************
     *  添加IP白名单            	  @author ruizuo qiyongsheng    *
     * **********************************************************
*/
	public static function insert($data){
		return \Yii::$app->manage_db
            ->createCommand()
            ->insert(self::getTableName(), $data)
            ->execute();
	}
/**
      ***********************************************************
      *     修改IP白名单                  @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function update( $set, $where) {
        return (new \yii\db\Query())
                        ->createCommand(\Yii::$app->manage_db)
                        ->update(self::getTableName(), $set, $where)
                        ->execute();
    }
}
