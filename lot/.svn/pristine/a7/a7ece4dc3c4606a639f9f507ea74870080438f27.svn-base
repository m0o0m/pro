<?php

namespace common\models;

use Yii;

class CommonModel extends \yii\base\Model {

    //获取所有彩种
    public static function getAllFcTypes() {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'fc_games')
                        ->where((['in', 'state', [1, 3]]))
                        ->orderby('is_hot desc,order_by desc')
                        ->all(\Yii::$app->manage_db);
    }

    //获取所有彩种分类
    public static function getAllFcCates() {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'fc_sub')
                        ->orderby('order_by desc')
                        ->all(\Yii::$app->manage_db);
    }

    /**
     * **********************************************************
     * 查询开奖结果数据              @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getAutoData($fc_type, $where, $orderby = 'id desc', $offset = 0, $limit = 100) {
        $table = self::GetAutoTableNameByType($fc_type);
        return (new \yii\db\Query())
                        ->select('*')
                        ->from($table)
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderby($orderby)
                        ->all(\Yii::$app->manage_db);
    }

    /**
     * **********************************************************
     * 根据彩种获取开奖结果表名        @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function GetAutoTableNameByType($type) {
        return \Yii::$app->manage_db->tablePrefix . 'auto_' . $type;
    }

    /**
     * **********************************************************
     *  获取初始限额           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getInitLimitByFcType($fc_type) {
        return (new \yii\db\Query)
                        ->select('id,fc_type,gameplay,name,limit_min,single_field_max,single_note_max')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'fc_games_set')
                        ->where([
                            'fc_type' => $fc_type,
                            'status' => 1
                        ])
                        ->all(\Yii::$app->manage_db);
    }
    /**
          ***********************************************************
          *  获取通用 开封盘时间           @author ruizuo qiyongsheng    *
          ***********************************************************
    */
        public static function getAllOpentime(){
             return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'opentime')
                        ->where(['status'=>1])
                        ->all(\Yii::$app->manage_db);
        }
    /**
          ***********************************************************
          *  加拿大  六合彩开封盘时间     @author ruizuo qiyongsheng    *
          ***********************************************************
    */
        public static function getOhterOpentime($type,$where){
            if($type == 'jnd_28') $type = 'jnd_bs';
            $table = \Yii::$app->manage_db->tablePrefix . $type . '_opentime';
            return (new \yii\db\Query())
                        ->select('*')
                        ->from( $table)
                        ->where($where)
                        ->orderBY('qishu DESC')
                        ->one(\Yii::$app->manage_db);
        }

    public static function getRawSqlForMyCat($table, $data) {
        if($table == 'my_line_cash_record') return array();
        if(Yii::$app->params['is_mycat']){
            $fields = '`id`';
            $values = '';
            switch($table){
                case 'my_user_sh':
                case 'my_user_ua':
                case 'my_user_agent':
                    $values = 'next value for MYCATSEQ_AGENT';
                    break;
                case 'my_sh_cash_record':
                case 'my_ua_cash_record':
                case 'my_agent_cash_record':
                    $values = 'next value for MYCATSEQ_AGENTCASHRECORD';
                    break;
                case 'my_user':
                    $values = 'next value for MYCATSEQ_USER';
                    break;
                case 'my_user_cash_record':
                    $values = 'next value for MYCATSEQ_USERCASHRECORD';
                    break;
                case 'my_bet_record':
                    $values = 'next value for MYCATSEQ_BETRECORD';
                    break;
                case 'my_sys_line_list':
                    // $values = 'next value for MYCATSEQ_GLOBA';
                    return false;
                    break;
                default:
                    return false;
                    break;
            }
            foreach($data as $k => $v){
                $fields .= ", `{$k}`";
                $values .= is_numeric($v) ? ", {$v}" : ", '{$v}'";
            }
            return "insert into `{$table}` ({$fields}) values ({$values})";
        }else{
            return false;
        }
    }

     /**
     * **********************************************************
     *  获取总条数               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
     public static function count($where,$table) {
        return (new \yii\db\Query())
                        ->select('count(id)')
                        ->from(\Yii::$app->db->tablePrefix . $table)
                        ->where($where)
                        ->scalar();
    }  
    /**
        ***********************************************************
        *  查询数据               @author ruizuo qiyongsheng    *
        ***********************************************************
    */
    public static function getData($select, $where = array(), $offset = 0 , $limit = 100, $orderby = 'id DESC', $table) {
        return (new \yii\db\Query())
                        ->select($select)
                        ->from(\Yii::$app->db->tablePrefix . $table)
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy($orderby)
                        ->all();
    }

}
