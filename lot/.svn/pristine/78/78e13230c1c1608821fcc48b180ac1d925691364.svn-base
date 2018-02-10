<?php

namespace our_backend\models;

use Yii;
use yii\db\Query;
use common\helpers\Helper; //

class BetModel extends \yii\base\Model {

    public static function get_user_info($uid) {
        return (new \yii\db\Query())
                        ->select('line_id,uname,money,agent_id,uid')
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where(['uid' => $uid])
                        ->one(yii::$app->db);
    }

    public static function get_auto_odds($type) {
        // return (new \yii\db\Query())
        //                 ->select('id,fc_type,odd,remark')
        //                 ->from(\Yii::$app->manage_db->tablePrefix . 'fc_games_type')
        //                 ->where(['status' => 1])
        //                 ->all(Yii::$app->manage_db);
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'fc_games_type')
                        ->where(['status' => 1])
                        ->andWhere('fc_type=:fc_type', [':fc_type' => $type])
                        ->all(Yii::$app->manage_db);
    }

    public static function get_all_games() {
        return (new \yii\db\Query())
                        ->select('type')
                        ->from(\Yii::$app->db->tablePrefix . 'fc_games')
                        ->where(['state' => 1])
                        ->all(Yii::$app->db);
    }

    public static function get_one_agent_info($agent_id) {
        return (new \yii\db\Query())
                        ->select('login_name,line_id,id')
                        ->from(\Yii::$app->db->tablePrefix . 'agent_admin')
                        ->where(['id' => $agent_id])
                        ->andWhere('is_delete=:is_delete', [':is_delete' => 1])
                        ->one(Yii::$app->db);
    }

    public static function get_line_agent( $agent_id,$type) {
        return (new \yii\db\Query())
                        ->select('play_id,odd,line_id,id')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'fc_games_odd')
                        ->where(['agent_id' => $agent_id])
                        ->andWhere('status=:status', [':status' => 1])
                        ->andWhere('fc_type=:fc_type', [':fc_type' => $type])
                        ->all(Yii::$app->manage_db);
    }

    public static function get_auto_limit() {
        return (new \yii\db\Query())
                        ->select('id,fc_type,gameplay,pankou,limit_min,single_field_max,single_note_max')
                        ->from(\Yii::$app->db->tablePrefix . 'fc_games_set')
                        ->where(['status' => 1])
                        ->all(Yii::$app->db);
    }

   public static function get_auto_limit_($type) {
        return (new \yii\db\Query())
                        ->select('id,fc_type,gameplay,limit_min,single_field_max,single_note_max')
                        ->from(\Yii::$app->db->tablePrefix . 'fc_games_set_')
                        ->where(['status' => 1])
                        ->andWhere('fc_type=:fc_type', [':fc_type' => $type])
                        ->all(Yii::$app->db);
    }


    public static function get_line_agent_limit($agent_id,$type) {
        return (new \yii\db\Query())
                        ->select('aid,limit_min,single_field_max,single_note_max,gameplay')
                        ->from(\Yii::$app->db->tablePrefix . 'agent_fc_set')
                        ->where(['aid' => $agent_id])
                        ->andWhere('status=:status', [':status' => 1])
                        ->andWhere('fc_type=:fc_type', [':fc_type' => $type])
                        ->all(Yii::$app->db);
    }

    //获取彩种一天所有的开奖时间和期数
    public static function get_last_kaijiang($type, $map=array()) {
        return (new \yii\db\Query())
                        ->select('kaijiang,qishu')
                        ->from(\Yii::$app->db->tablePrefix .$type.'_opentime')
                        ->where($map)
                        ->orderBy(['kaijiang' => SORT_DESC])
                        ->one(Yii::$app->db);
    }

    //获取当期开盘关盘时间
    public static function get_now_open($type, $now_time) {
        return (new \yii\db\Query())
                        ->select('aid,limit_min,single_field_max,single_note_max,play_id,gameplay')
                        ->from(\Yii::$app->db->tablePrefix . $type . '_opentime')
                        ->where(['>=', 'kaijiang', $now_time])
                        ->andWhere('ok=:ok', [':ok' => 0])
                        ->all(Yii::$app->db);
    }

    // //查询是否开盘
    public static function get_is_open($type, $map=array()) {
        $query = new Query;
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $type . '_opentime'; //表名
        $rows = $query->select('fengpan,kaijiang,kaipan,qishu')->from($table_name)->where($map)->orderBy('kaijiang DESC')->all();
        return $rows;
    }

    //查询是否开盘
    public static function _get_is_open($type, $map) {
        $limit = 20;
        $query = new Query;
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $type . '_opentime'; //表名
        $rows = $query->select('fengpan,kaijiang,kaipan,qishu')->from($table_name)->where($map)->orderBy('kaijiang asc')->limit($limit)->all();
        return $rows;
    }

    //查询是否开盘
    public static function get_is_open1($type, $map,$limit,$order) {
        $query = new Query;
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $type . '_opentime'; //表名
        $rows = $query->select('fengpan,kaijiang,kaipan,qishu')->from($table_name)->where($map)->orderBy($order)->limit($limit)->all();
        return $rows;
    }



    public static function get_now_bet_count($type,$qishu,$uid){
        return (new \yii\db\Query())
                        ->select(["SUM(money) as total_money"])
                        ->from(\Yii::$app->db->tablePrefix.'bet_'.$type)
                        ->where(['periods' => $qishu])
                        ->andWhere('uid=:uid', [':uid' => $uid])
                        ->all(Yii::$app->db);

    }



}
