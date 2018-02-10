<?php

namespace backend\models;

use Yii;

/**
 * Login form
 */
class AdminModel extends \yii\base\Model {

    public static function login($params) {
        $user = (new \yii\db\Query())
            ->select('*')
            ->from($params['conn']->tablePrefix . $params['table'])
            ->where('login_user=:username', [':username' => $params['username']])
            ->andWhere('login_pwd=:password', [':password' => md5(md5($params['password']))])
            ->one($params['conn']);

        if (!empty($user)) {
            //查看帐号是否禁用
            if($user['is_delete'] == 2){
                echo json_encode(['code' => 300, 'msg' => '帐号被禁用!']);
                die;
            }
            //如果是代理或子帐号，session中要存储user_agent表的信息
            if($user['user_type'] == 3){
                $agent_id =  self::getAgentIdByLoginUser($user['login_user']);
                if(!$agent_id) return false;
                $agent_name = $user['login_user'];
            }elseif($user['user_type'] == 4){
                $agent_name = (new \yii\db\Query())
                    ->select('login_user')
                    ->from($params['conn']->tablePrefix . $params['table'])
                    ->where(['id'=>$user['pid']])
                    ->scalar($params['conn']);
                if(!$agent_name) return false;
                $agent_id =  self::getAgentIdByLoginUser($agent_name);
                if(!$agent_id) return false;
            }

            // 查用户权限
            if (in_array($user['user_type'], [1, 2, 3])) {
                $role_table = \Yii::$app->manage_db->tablePrefix . 'admin_role';
                $role = self::getOneRole($role_table, ['id'=>$user['role_id'], 'is_delete'=>1]);
            } elseif($user['user_type'] == 4) {//代理子帐号没有角色
                //查询父级代理权限
                $agent_role = self::getAccessByLine($user['line_id']);
                if(!$agent_role) return false;
                $role = array();
                //保持权限与代理同步 计算交集
                $agent_role_arr = explode(',', $agent_role['access_id']);
                $self_role_arr = explode(',', $user['son_role']);
                $same_role_arr = array_intersect($agent_role_arr, $self_role_arr);
                if(!empty($same_role_arr)){
                    $role['access_id'] = implode(',', $same_role_arr);
                }else{
                    $role['access_id'] = '';
                }
                $role['id'] = '';
                $role['role_name'] = '';
                $role['role'] = '';
            } else{
                return false;
            }

            //超管获取系统所有路由
            if ($user['user_type'] == 1) {
                $where = [];
            } else {
                $where = ['in', 'id', explode(',', $role['access_id'])];
            }
            //查看权限所对应的菜单url
            $menu = (new \yii\db\Query())
                    ->select('*')
                    ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                    ->where($where)
                    ->andWhere('is_delete=:is_delete', [':is_delete' => 1])
                    ->orderBy(['level' => SORT_ASC, 'sort' => SORT_DESC])
                    ->all(Yii::$app->manage_db);
        } else {
            return false;
        }

        //获取登录者账号身份
        $user = [
            'line_id' => $user['line_id'],
			'login_ip' => $user['login_ip'],
            'uid' => $user['id'],
            'pid' => $user['pid'],
            'login_user' => $user['login_user'],
            'login_name' => $user['login_name'],
            'user_type' => $user['user_type'],
            'role' => $role['role'],
            'role_id' => $role['id'],
            'role_name' => $role['role_name'],
            'role_access' => $role['access_id'],
            'menu' => $menu
        ];
        if(isset($agent_name)) $user['agent_name'] = $agent_name;
        if(isset($agent_id)) $user['agent_id'] = $agent_id;

        return $user;
    }

    //获取角色条数
    public static function getRoleCount($table, $where) {
        return (new \yii\db\Query())
                        ->from($table)
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    //获取角色列表
    public static function getRoleList($table, $where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from($table)
                        ->where($where)
                        ->offset($offset)
                        ->limit($limit)
                        ->orderBy(['id' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //获取系统所有路由
    public static function getAllRoute($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_menu')
                        ->where($where)
                        ->orderBy(['level' => SORT_ASC, 'sort' => SORT_DESC])
                        ->all(\Yii::$app->manage_db);
    }

    //新增角色
    public static function addRole($table, $arr) {
        return Yii::$app->manage_db->createCommand()->insert($table, $arr
                )->execute();
    }

    //修改角色信息
    public static function updateRole($table, $id, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update($table, $updateArr, ['id' => $id])->execute();
    }

    //获取一条角色信息
    public static function getOneRole($table, $where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from($table)
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }
    //获取代理角色id
    public static function getRoleId($table, $where){
         return (new \yii\db\Query())
                        ->select('id')
                        ->from($table)
                        ->where($where)
                        ->scalar(\Yii::$app->manage_db);
    }
    //获取管理员条数
    public static function getAdminCount($where) {
        return (new \yii\db\Query())
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                        ->where($where)
                        ->count('id', \Yii::$app->manage_db);
    }

    //获取管理员列表
    public static function getAdminList($where, $offset, $limit) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
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
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

    //获取系统所有角色
    public static function getAllRole($table, $where = []) {
        if(empty($where)) $where = ['is_delete' => 1];
        return (new \yii\db\Query())
                        ->select('*')
                        ->from($table)
                        ->where($where)
                        ->all(\Yii::$app->manage_db);
    }

    //新增角色
    public static function addAdmin($arr) {
        return Yii::$app->manage_db->createCommand()->insert(
                        \Yii::$app->manage_db->tablePrefix . 'agent_admin', $arr
                )->execute();
    }

    //修改角色信息
    public static function updateAdmin($id, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->db->tablePrefix . 'agent_admin', $updateArr, ['id' => $id])->execute();
    }

    //修改密码
    public static function updatePwd($uid, $updateArr) {
        return Yii::$app->manage_db->createCommand()->update(
                        \Yii::$app->manage_db->tablePrefix . 'agent_admin', $updateArr, ['id' => $uid])->execute();
    }

    //验证原始密码
    public static function checkPwd($where) {
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'agent_admin')
                        ->where($where)
                        ->one(\Yii::$app->manage_db);
    }

    //查询余额
    public static function queryMoney($table, $where) {
        return (new \yii\db\Query())
                        ->select('money')
                        ->from(\Yii::$app->db->tablePrefix . $table)
                        ->where($where)
                        ->scalar();
    }

    //根据帐号查询代理id
    public static function getAgentIdByLoginUser($login_user){
        return (new \yii\db\Query())
                        ->select('id')
                        ->from(\Yii::$app->db->tablePrefix . 'user_agent')
                        ->where(['login_user'=>$login_user])
                        ->scalar();
    }

    //因为每个线路只有一个角色是代理，所以可以根据线路id直接查询代理角色信息
    public static function getAccessByLine($line_id){
        return (new \yii\db\Query())
                        ->select('*')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'admin_role')
                        ->where(['line_id'=>$line_id, 'role'=>'agent'])
                        ->one(\Yii::$app->manage_db);
    }

   //公告信息
    public static function getNoticeInfo($map) {
        return (new \yii\db\Query())
                        ->select('id,content')
                        ->from(\Yii::$app->manage_db->tablePrefix . 'backend_notice')
                        ->where($map)
                        ->limit(1)
                        ->orderby('addtime desc')
                        ->one(\Yii::$app->manage_db);
    }
}
