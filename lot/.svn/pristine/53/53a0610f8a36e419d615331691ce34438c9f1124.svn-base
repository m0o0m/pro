<?php

namespace our_backend\models;

use Yii;

class AgentmenuModel extends \yii\base\Model {

    //新增
    public static function add($params) {
        return Yii::$app->manage_db->createCommand()->insert(
                        \Yii::$app->manage_db->tablePrefix . 'agent_menu', $params
                )->execute();
    }

    //修改
    public static function update($id, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->manage_db->tablePrefix . 'agent_menu', $updateArr, ['id' => $id])->execute();
    }

    //获取一级菜单
    public static function getFirstMenu() {
        return (new \yii\db\Query())
                        ->select('id, menu,menu_name')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where('level=:level', [':level' => 1])
                        ->andWhere('is_delete=:is_delete', [':is_delete' => 1])
                        ->andWhere('is_url=:is_url', [':is_url' => 1])
                        ->all(Yii::$app->manage_db);
    }

    //获取菜单条数
    public static function getMenuCount($where) {
        return (new \yii\db\Query())
                        ->select('id, menu,menu_name')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    //获取菜单列表
    public static function getMenuList($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['level' => SORT_ASC, 'sort' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //单条数据
    public static function getOneData($id) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where(['id' => $id])
                        ->one(\Yii::$app->manage_db);
    }

    //获取子类菜单
    public static function getSubMenu($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where($where)
                        ->all(\Yii::$app->manage_db);
    }

}
