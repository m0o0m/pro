<?php

namespace common\models;

use Yii;
use yii\db\Query;
use common\helpers\lotteryOrm;
use common\helpers\Helper;

class BetModel extends \yii\base\Model {

    /**
     * **********************************************************
     *  获取用户简单信息         @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getUserInfo($uid) {
        return (new \yii\db\Query())
                        ->select('line_id,uname,addtime,status')
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where(['uid' => $uid])
                        ->one();
    }

    /**
     * **********************************************************
     *  获取注单简单信息         @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getJoinOneSql($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                        ->where($where)
                        ->one();
    }

    /**
     * 列表数据
     * @param type $condition
     * @return type array
     */
    public static function index($condition, $where) {
        return (new \yii\db\Query())
                        ->select($condition['field'])
                        ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                        ->where($where)
                        ->orderBy(['id'=>SORT_DESC])
                        ->offset($condition['offset'])
                        ->limit($condition['limit'])
                        ->all();
    }

    /**
     * **********************************************************
     * 根据彩种查询未结算注单    @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getNoBalanceBet($sql) {
        return \Yii::$app->db
                ->createCommand($sql)
                ->queryAll();
        // return (new \yii\db\Query())
        //                 ->select('fc_type,periods,count(id) as count')
        //                 ->from(\Yii::$app->db->tablePrefix . 'bet_record')
        //                 ->where($where)
        //                 ->groupBy('fc_type,periods')
        //                 ->all();
    }
   
        
    /**
     * **********************************************************
     *  根据期数获取开奖数据     @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getAutoData($type, $periods) {
        return (new \yii\db\Query())
                        ->select(['id', 'addtime', 'status'])
                        ->from(\Yii::$app->manage_db->tablePrefix . 'auto_' . $type)
                        ->where(['qishu' => $periods])
                        ->one(\Yii::$app->manage_db);
    }

    /**
     * **********************************************************
     *  获取所有线路             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getLines() {
        return (new \yii\db\Query())
                        ->select(['line_id', 'line_name'])
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->orderBy('line_id ASC')
                        ->where(['status' => 1])
                        ->all();
    }
    /**
          ***********************************************************
          *  获取所有代理           @author ruizuo qiyongsheng    *
          ***********************************************************
    */
    public static function getAllAgent(){
         return (new \yii\db\Query())
                    ->select(['id', 'login_name', 'line_id'])
                    ->from(\Yii::$app->db->tablePrefix . 'user_agent')
                    ->where(['is_delete' => 1])
                    ->all();
    }
        
    /**
     * **********************************************************
     *  批量设置无效及恢复       @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function invalid($type, $set, $where) {
        return Yii::$app->db
                        ->createCommand()
                        ->update(\Yii::$app->db->tablePrefix . 'bet_record', $set, $where)
                        ->execute();
    }

    /**
     * **********************************************************
     *  注单列表页总计数据         @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function totalcount($field, $where) {
        return (new \yii\db\Query())
                        ->select($field)
                        ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                        ->where($where)
                        ->one();
    }

    /**
     * 获取数据总条数
     * @param type $tabname
     */
    public static function getDataTotalNum($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                        ->where($where)
                        ->count();
    }

    /**
     * **********************************************************
     *  根据期数查询开封盘时间       @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getTimeByQishu($where, $type) {
        if ($type == 'jnd_28')
            $type = 'jnd_bs';
        if ($type == 'liuhecai' || $type == 'jnd_bs') {
            $table_name = \Yii::$app->manage_db->tablePrefix . $type . '_opentime'; //表名
        } else {
            $table_name = \Yii::$app->manage_db->tablePrefix . 'opentime';
        }
        return (new \yii\db\Query())
                        ->select('fengpan,kaijiang,kaipan,qishu')
                        ->from($table_name)
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

    // //查询是否开盘
    public static function get_is_open($type, $map = array(), $limit = 20) {
        if ($type == 'jnd_28')
            $type = 'jnd_bs';
        if ($type == 'liuhecai' || $type == 'jnd_bs') {
            $table_name = \Yii::$app->manage_db->tablePrefix . $type . '_opentime'; //表名
        } else {
            $table_name = \Yii::$app->manage_db->tablePrefix . 'opentime';
        }
        return (new \yii\db\Query())
                        ->select('fengpan,kaijiang,kaipan,qishu')
                        ->from($table_name)
                        ->where($map)
                        ->orderBy('kaijiang DESC')
                        ->limit($limit)
                        ->all(\Yii::$app->manage_db);
    }

    public static function getBetRecord($where, $offset = 0, $limit = 0) {
        $query = (new \yii\db\Query())
                ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                ->where($where)
                ->offset($offset)
                ->orderBy(['id' => SORT_DESC])
                ->limit($limit);
        return $query->all(\Yii::$app->db);
    }

    public static function getBetRecordCount($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->db->tablePrefix . 'bet_record')
                        ->where($where)
                        ->count('id', \Yii::$app->db);
    }

}
