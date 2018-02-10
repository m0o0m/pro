<?php

namespace common\models;

use Yii;
use yii\mongodb\Query;
use common\helpers\mongoTables;

/**
 * 日志model(总后台,业主后台,前台)  共用
 */
class LogModel extends \yii\base\Model {

    //插入
    public static function insertLog($table, $insert) {
        return Yii::$app->mongodb->getCollection($table)->insert($insert);
    }

    public static function getLogs($table, $where = [], $offset = 0, $limit = 100, $orderby, $select = []) {

        $query = new Query;
        $query->select($select)
                ->from($table)
                ->where($where)
                ->offset($offset)
                ->limit($limit)
                ->orderBy($orderby);
        return $query->all();
    }

    //获取某用户最近的登录信息
    public static function getLoginLogs($table,$uid,$limit,$select = []) {
        $query = new Query;
        $query->select($select)
                ->from($table)
                ->where(['uid'=>$uid])
                ->orderby(['id' => SORT_DESC])
                ->limit($limit);
        return $query->all();
    }

    public static function getLogsCount($table, $where) {
        $query = new Query;
        return $query->from($table)
                        ->where($where)
                        ->count('_id');
    }

    public static function getPtype(){
        switch(Yii::$app->id){
            case 'our_backend':
                $ptype = 1;
                break;
            case 'agent_backend':
                $ptype = 2;
                break;
            case 'pc_frontend':
                $ptype = 3;
                break;
            case 'wap':
                $ptype = 4;
                break;
            case 'app':
                $ptype = 5;
                break;
            default:
                break;
        }
        return $ptype;
    }

          //获取异常住单信息
    public static function getErrorbets($table, $where = [], $offset = 0, $limit = 100, $select = []) {
        $query = new Query;
        $query->select($select)
                ->from($table)
                ->where($where)
                ->offset($offset)
                ->limit($limit);
        return $query->all();
    }

    //获取总条数
    public static function getCount($table, $where) {
        $query = new Query;
        return $query->from($table)
                        ->where($where)
                        ->count('_id');
    }

}
