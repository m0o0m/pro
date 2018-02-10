<?php

namespace frontend_api\models;

use Yii;
use yii\db\Query;

class UserModel extends \yii\base\Model {

    /**
     * 取会员信息
     */
    public static function getUserInfo($map) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'user')
                        ->where($map)
                        ->one(\Yii::$app->db);
    }

    /**
     * 取用户现金记录
     */
    public static function getUserCashRecord($map) {
        return (new \yii\db\Query)
                        ->select('addtime,cash_do_type,cash_num')
                        ->from(\Yii::$app->db->tablePrefix . 'user_cash_record')
                        ->where($map)
                        ->all(\Yii::$app->db);
    }

    /**
     * **********************************************************
     * 查询会员注单统计表            @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getReportInfo($filed = '*', $where = []) {
        return (new \yii\db\Query())
                        ->select($filed)
                        ->from(\Yii::$app->manage_db->tablePrefix . 'user_bet')
                        ->where($where)
                        ->all(\Yii::$app->manage_db);
    }

    /**
     * 插入用户反馈问题
     */
    public static function insertUserProblem($data) {
        return Yii::$app->db->createCommand()
                        ->insert(\Yii::$app->db->tablePrefix . 'user_problem', $data)
                        ->execute();
    }

    /**
     * **********************************************************
     *  更新我喜欢的彩种             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function updateMyLoves($update, $where) {
        return \Yii::$app->manage_db
                        ->createCommand()
                        ->update(\Yii::$app->manage_db->tablePrefix . 'user_like', $update, $where)
                        ->execute();
    }

    /**
     * **********************************************************
     *  添加我喜欢的彩种                                           *
     * **********************************************************
     */
    public static function addMyLoves($data) {
        return Yii::$app->manage_db
                        ->createCommand()
                        ->insert(\Yii::$app->manage_db->tablePrefix . 'user_like', $data)
                        ->execute();
    }

}