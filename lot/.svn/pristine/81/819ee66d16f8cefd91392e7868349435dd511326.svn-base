<?php

namespace our_backend\models;

use Yii;

/**
 * Login form
 */
class AdminModel extends \yii\base\Model {

    public static function login($params) {
        $user = (new \yii\db\Query())
                ->select('*')
                ->from('my_sys_admin')
                ->where('login_user=:username', [':username' => $params['username']])
                ->andWhere('login_pwd=:password', [':password' => md5(md5($params['password']))])
                ->one(Yii::$app->manage_db);

        if (!empty($user)) {
            // 查用户权限
            $role = (new \yii\db\Query())
                    ->select('*')
                    ->from('my_sys_role')
                    ->where('id=:id', [':id' => $user['role_id']])
                    ->andWhere('is_delete=:is_delete', [':is_delete' => 1])
                    ->one(Yii::$app->manage_db);

            //超管获取系统所有路由
            if ($role['role'] == 'super') {
                $where = [];
            } else {
                $where = ['in', 'id', explode(',', $role['access_id'])];
            }
            //查看权限所对应的菜单url
            $menu = (new \yii\db\Query())
                    ->select('*')
                    ->from('my_sys_menu')
                    ->where($where)
                    ->andWhere('is_delete=:is_delete', [':is_delete' => 1])
                    ->orderBy(['level' => SORT_ASC, 'sort' => SORT_DESC])
                    ->all(Yii::$app->manage_db);
            //请求路由对应中文名称
            $menu_name = (new \yii\db\Query())
                ->select('menu,menu_name')
                ->from('my_sys_menu')
                ->where(['<>','level',1])
//                ->where('level=:level', [':level' => 1])
                ->all(Yii::$app->manage_db);
        } else {
            return false;
        }
        $user = [
            'uid' => $user['id'],
            'login_ip' => $user['login_ip'],
            'login_user' => $user['login_user'],
            'login_name' => $user['login_name'],
            'role' => $role['role'],
            'role_id' => $role['id'],
            'role_name' => $role['role_name'],
            'role_access' => $role['access_id'],
            'menu' => $menu,
            'menu_name' => $menu_name
        ];

        return $user;
    }

    //获取角色条数
    public static function getRoleCount($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_role')
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    //获取角色列表
    public static function getRoleList($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_role')
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['id' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //获取系统所有路由
    public static function getAllRoute() {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_menu')
                        ->where(['is_delete' => 1])
                        ->orderBy(['level' => SORT_ASC, 'sort' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //新增角色
    public static function addRole($arr) {
        return Yii::$app->manage_db->createCommand()->insert(
                        \Yii::$app->manage_db->tablePrefix . 'sys_role', $arr
                )->execute();
    }

    //修改角色信息
    public static function updateRole($id, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->manage_db->tablePrefix . 'sys_role', $updateArr, ['id' => $id])->execute();
    }

    //获取一条角色信息
    public static function getOneRole($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_role')
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

    //获取管理员条数
    public static function getAdminCount($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_admin')
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    //获取管理员列表
    public static function getAdminList($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_admin')
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['id' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //获取一条管理员信息
    public static function getOneAdmin($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_admin')
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

    //获取系统所有角色
    public static function getAllRole() {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_role')
                        ->where(['is_delete' => 1])
                        ->all(\Yii::$app->manage_db);
    }

    //新增角色
    public static function addAdmin($arr) {
        return Yii::$app->manage_db->createCommand()->insert(
                        \Yii::$app->manage_db->tablePrefix . 'sys_admin', $arr
                )->execute();
    }

    //修改角色信息
    public static function updateAdmin($id, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->manage_db->tablePrefix . 'sys_admin', $updateArr, ['id' => $id])->execute();
    }

    //修改密码
    public static function updatePwd($uid, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->manage_db->tablePrefix . 'sys_admin', $updateArr, ['id' => $uid])->execute();
    }

    //验证原始密码
    public static function checkPwd($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'sys_admin')
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

}
