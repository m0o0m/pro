<?php

namespace common\models;

use Yii;
use yii\db\Query;

class AgentModel extends \yii\base\Model {

    //获取代理条数
    public static function getAgentCount($database,$tab,$where) {
        return (new Query())
                        ->from($tab)
                        ->where($where)
                        ->count('id', $database);
    }

    //获取代理列表
    public static function getAgentList($database,$tab,$where, $offset, $limit) {
        return (new Query())
                        ->select('*')
                        ->from($tab)
                        ->where($where)
                        ->offset($offset)
                        ->orderBy(['id' => SORT_DESC])
                        ->limit($limit)

                        ->all($database);
    }

    //获取单个代理
    public static function getOneAgentByCondition($database,$tab,$where) {
        return (new Query())
                        ->select('*')
                        ->from($tab)
                        ->where($where)
                        ->one($database);
    }

    //获取多个代理
    public static function getMoreAgentByCondition($database,$tab,$where) {
        return (new Query())
                        ->select('*')
                        ->from($tab)
                        ->where($where)
                        ->all($database);
    }

    /*
     * 新增
     * $insert= []
     *
     */

    public static function insertAgent($database,$tab,$insert) {
        $sql = \common\models\CommonModel::getRawSqlForMyCat($tab, $insert);
        if(!empty($sql)){
            return $database->createCommand($sql)->execute();
        }

        return $database->createCommand()
                        ->insert($tab, $insert)
                        ->execute();
    }

    /*
     * 修改
     * $update= []
     * $where=[]
     */

    public static function updateAgent($database,$tab,$update, $where) {
        return $database->createCommand()
                        ->update($tab, $update, $where)
                        ->execute();
    }
    //获取线路信息
    public static function getlineinfo($line_id){
        $tablePrefix = Yii::$app->db->tablePrefix;
        return (new Query())
                ->select('*')
                ->from($tablePrefix . 'sys_line_list')
                ->where(['line_id' => $line_id, 'status' => 1])
                ->one(Yii::$app->db);
    }
    //更改线路金额
    public static function updatelinemoney($field, $val, $where) {
         return (new \yii\db\Query())
                ->createCommand()
                ->update(\Yii::$app->db->tablePrefix . 'sys_line_list', array($field => $val), $where)
                ->execute();
    }

    /* add code by frank start */
    public static function getAgent($agent_type = '', $line_id = '') {
        $where = [];
        if($line_id){
            $where['line_id'] = $line_id;
        }
        $atypes = ['sh' => 'sh', 'ua' => 'ua', 'at' => 'agent'];
        if($agent_type){
            if ( key_exists($agent_type, $atypes) ) $agent_type = $atypes[$agent_type];
            $result = self::getMoreAgentByCondition(\Yii::$app->db, \Yii::$app->db->tablePrefix . 'user_' . $agent_type, $where);
        }else{
            $result = [];
            foreach($atypes as $atype){
                $result = array_merge($result, self::getMoreAgentByCondition(\Yii::$app->db, \Yii::$app->manage_db->tablePrefix . 'user_' . $atype, $where));
            }
        }
        return $result;
    }
    public static function getAids($agent_type, $where) {
        $atypes = ['sh' => 'sh', 'ua' => 'ua', 'at' => 'agent'];
        if ( key_exists($agent_type, $atypes) ) $agent_type = $atypes[$agent_type];
        return (new Query())
            ->select('id,line_id')
            ->from(\Yii::$app->db->tablePrefix . 'user_' . $agent_type)
            ->where($where)
            ->all();
    }

    public static function insert($sql){
        return \Yii::$app->db
                ->createCommand($sql)
                ->execute();
    }
    /* add code by frank end */

}
