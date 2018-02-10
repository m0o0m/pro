<?php
namespace common\models;

use Yii;
use common\models\AgentModel;
use yii\helpers\ArrayHelper;

class OddsModel extends \yii\base\Model {

    public static function getTableName($type = ''){
        switch($type){
            case 'agent':
                $tab = 'agent_odds';
                break;
            case 'line':
                $tab = 'line_odds';
                break;
            default:
                $tab = 'fc_games_type';
                break;
        }
        return \Yii::$app->manage_db->tablePrefix . $tab;
    }

    public static function getCount($table, $where) {
        return (new \yii\db\Query())
            ->from($table)
            ->where($where)
            ->count('id', \Yii::$app->manage_db);
    }

    public static function getList($table, $where, $offset, $limit, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->offset($offset)
            ->limit($limit)
            ->orderBy(['id'=>SORT_DESC])// 默认id倒序，如需其它非索引字段的排序 另行设计优化。
            ->all(\Yii::$app->manage_db);
    }

    public static function getItems($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->all(\Yii::$app->manage_db);
    }

    public static function getOne($table, $where, $select = '*') {
        return (new \yii\db\Query())
            ->select($select)
            ->from($table)
            ->where($where)
            ->one(\Yii::$app->manage_db);
    }

    public static function insert($table, $values) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->insert($table, $values)
            ->execute();
    }

    public static function update($table, $set,$where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->update($table, $set, $where)
            ->execute();
    }

    public static function del($table, $where) {
        return \Yii::$app->manage_db
            ->createCommand()
            ->delete($table, $where)
            ->execute();
    }

    public static function agent($request) {
        $line_id = isset($request['line_id']) ? trim($request['line_id']) : null;
        $agent_type = isset($request['agent_type']) ? trim($request['agent_type']) : null;

        if (empty($line_id) || empty($agent_type)) {
            return [];
        }
        return AgentModel::getAgent($agent_type, $line_id);
    }

    public static function getOdds($request) {
        $line_id = isset($request['line_id']) ? trim($request['line_id']) : null;
        $agent_type = isset($request['agent_type']) ? trim($request['agent_type']) : null;
        $agent_id = isset($request['agent_id']) ? trim($request['agent_id']) : null;
        $fc_type = isset($request['fc_type']) ? trim($request['fc_type']) : null;

        $agentCustom = [];
        if ($agent_type && $agent_id) {
            $agentCustom = self::getAgentCustom($fc_type, $agent_type, $agent_id);
            $agentCustom = ArrayHelper::index($agentCustom, 'play_id');
        }
        $lineCustom = [];
        if ($line_id && !($agent_type && $agent_id)) {
            $lineCustom = self::getLineCustom($fc_type, $line_id);
            $lineCustom = ArrayHelper::index($lineCustom, 'play_id');
        }

        $def = self::getDef($fc_type);
        $def = ArrayHelper::index($def, 'id');

        $data = $agentCustom + $lineCustom + $def; // 代理 覆盖 线路 覆盖 初始
        ksort($data); // ID正序

        return $data;
    }

    // 代理赔率
    public static function getAgentCustom($fc_type = '', $agent_type = '', $agent_id = '') {
        $where['status'] = '1';
        if ($fc_type)
            $where['fc_type'] = $fc_type;
        if ($agent_type)
            $where['agent_type'] = $agent_type;
        if ($agent_id)
            $where['agent_id'] = $agent_id;

        $result = self::getItems(self::getTableName('agent'), $where);

        return $result;
    }

    // 线路赔率
    public static function getLineCustom($fc_type = '', $line_id = '') {
        $where['status'] = '1';
        if ($fc_type)
            $where['fc_type'] = $fc_type;
        if ($line_id)
            $where['line_id'] = $line_id;

        $result = self::getItems(self::getTableName('line'), $where);

        return $result;
    }

    // 初始赔率
    public static function getDef($fc_type = '') {
        $where['status'] = '1';
        if ($fc_type)
            $where['fc_type'] = $fc_type;

        return self::getItems(self::getTableName(), $where);
    }

}
