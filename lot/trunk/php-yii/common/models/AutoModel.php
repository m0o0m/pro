<?php

namespace common\models;

use Yii;

class AutoModel extends \yii\base\Model {

    public static function getCount($fc_type, $where) {
        return (new \yii\db\Query())
                        ->from(self::getTableName($fc_type))
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    public static function getList($fc_type, $where, $offset, $limit, $orderby = '') {
        if(empty($orderby)) $orderby = ['id'=>SORT_DESC];
        return (new \yii\db\Query())
                        ->from(self::getTableName($fc_type))
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy($orderby)// 默认id倒序，如需其它非索引字段的排序 另行设计优化。
                        ->all(\Yii::$app->manage_db);
    }

    public static function getItems($fc_type, $where, $select = '*') {
        return (new \yii\db\Query())
                        ->select($select)
                        ->from(self::getTableName($fc_type))
                        ->where($where)
                        ->all(\Yii::$app->db);
    }

    public static function getOne($fc_type, $where, $select = '*') {
        return (new \yii\db\Query())
                        ->select($select)
                        ->from(self::getTableName($fc_type))
                        ->where($where)
                        ->one(\Yii::$app->db);
    }

    public static function insert($fc_type, $values) {
        return \Yii::$app->manage_db
                        ->createCommand()
                        ->insert(self::getTableName($fc_type), $values)
                        ->execute();
    }

    public static function update($fc_type, $set, $where) {
        return \Yii::$app->manage_db
                        ->createCommand()
                        ->update(self::getTableName($fc_type), $set, $where)
                        ->execute();
    }

    public static function delete($fc_type, $where) {
        return (new \yii\db\Query())
                        ->createCommand()
                        ->delete(self::getTableName($fc_type), $where)
                        ->execute();
    }

    public static function getBallNum($fc_type) {
        return \Yii::$app->db
                        ->createCommand("SELECT count(*)-5 ball_num FROM information_schema.`COLUMNS` WHERE TABLE_NAME='" . self::getTableName($fc_type) . "'")
                        ->queryScalar();
    }

    public static function getTableName($fc_type) {
        return \Yii::$app->manage_db->tablePrefix . 'auto_' . $fc_type; 
    }

    public static function getBetTableName($fc_type) {
        return \Yii::$app->db->tablePrefix . 'bet_record';
    }

    /**
     * **********************************************************
     *  查询已结算注单条数       @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getOKBetCount($table, $type, $qishu) {
        return (new \yii\db\Query())
                        ->select('count(*)')
                        ->from($table)
                        ->where(['js' => 2, 'periods' => "$qishu",'fc_type'=>$type])
                        ->scalar();
    }

    public static function getLastQishu($fc_type) {
        return (new \yii\db\Query())
                        ->select('qishu')
                        ->from(self::getTableName($fc_type))
                        ->offset(0)
                        ->limit(1)
                        ->orderBy(['qishu'=>SORT_DESC])
                        ->scalar(\Yii::$app->manage_db,'qishu');
    }
    // 获取线路基本信息（域名）
    public static function getLine($line_id = '') {
        $andwhere = '1';
        $params = [];
        if ($line_id) {
            $andwhere .= ' AND line_id = :line_id';
            $params[':line_id'] = $line_id;
        }
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where(['status' => 1])
                        ->andwhere($andwhere, $params)
                        ->all();
    }

    //维护项目列表
    public function maintain_item_list() {
        $andwhere = '1';
        $params = [];
        if ($line_id) {
            $andwhere .= ' AND line_id = :line_id';
            $params[':line_id'] = $line_id;
        }
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where(['status' => 1])
                        ->andwhere($andwhere, $params)
                        ->all();
    }

    //维护项目编辑
    public function maintain_item_edit($id, $arr) {
        // $db_model['tab'] = 'site_cate_module';
        // $db_model['type'] = 4;
        // if($id){
        //     return $this->M($db_model)->where(array('id'=>$id))->update($arr);
        //}

        $andwhere = '1';
        $params = [];
        if ($line_id) {
            $andwhere .= ' AND line_id = :line_id';
            $params[':line_id'] = $line_id;
        }
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->db->tablePrefix . 'sys_line_list')
                        ->where(['status' => 1])
                        ->andwhere($andwhere, $params)
                        ->all();
    }

    public static function get_count($fc_type, $where) {
        return self::getCount($fc_type, $where);
    }

    public static function get_list($fc_type, $where, $offset, $limit, $orderby = '') {
        return self::getList($fc_type, $where, $offset, $limit, $orderby);
    }

}
