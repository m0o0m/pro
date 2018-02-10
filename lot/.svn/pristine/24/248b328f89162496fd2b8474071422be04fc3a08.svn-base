<?php 
namespace common\models;

use Yii;
use yii\db\Query;  

class OpentimeModel extends \yii\base\Model{
    /**
     * 插入数据
     * @param type $tabname
     * @param type $data
     * @return type
     */
    public static function addOpenData($tabname,$data){
        $rows = \Yii::$app->manage_db->createCommand()->insert($tabname, $data)->execute();
        return $rows;
    }
   /**
    * 删除
    * @param type $where
    * @param type $tabname
    */
   public static function delOpenData($where, $tabname){
       $rows = \Yii::$app->manage_db->createCommand()->delete($tabname, $where)->execute();
       return $rows;
   }
    /**
     * 获取数据总条数
     * @param type $tabname
     */
    public static function getDataTotalNum($tabname, $where) {
        return (new \yii\db\Query())
                ->from($tabname)
                ->where($where)
                ->count('id',\Yii::$app->manage_db);
    }
     /**
     * 列表数据
     * @param type $condition
     * @return type array
     */
     public static function index($condition, $where) {
        return (new \yii\db\Query())
                ->select($condition['field'])
                ->from($condition['tabname'])
                ->where($where)
                ->orderBy(['id' => SORT_DESC])
                ->offset($condition['offset'])
                ->limit($condition['limit'])
                ->all(\Yii::$app->manage_db);
    }
 /**
     * 获取单条数据
     * @param type $id
     * @param type $tabname
     * @param type $fields
     * @return type
     */
    public static function getOneSql($where, $tabname) {
        return (new \yii\db\Query())
                        ->select("*")
                        ->from($tabname)
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

 /**
     * 更新编辑数据
     * @param type $data
     * @param type $where
     * @param type $tabname
     * @return type
     */
    public static function upEditOpenData($data, $where, $tabname) {
        return \Yii::$app->manage_db
                        ->createCommand()
                        ->update($tabname, $data, $where)
                        ->execute();
    }
}