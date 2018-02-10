<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AdminModel;
use common\models\AgentModel;

class AgentController extends Controller {

    /**
     * **********************************************************
     *  股东列表                             *
     * **********************************************************
     */
    public function actionShindex() {
        $get = Yii::$app->request->get();
        $tab_arr = self::get_tab('user_sh');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = !empty($post['line_id']) ? $post['line_id'] : null; //站点id
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database, $tab, $where);
        $data = AgentModel::getAgentList($database, $tab, $where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $lines = $this->getLines(); //获取线路
        if (empty($get)) {
            $data = array();
            $pagecount = 1;
            $page = 1;
        }
        $render = [
            'data' => $data,
            'lines' => $lines,
            'pagecount' => $pagecount,
            'page' => $page
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('sh_index.html', $render);
        } else {
            return $this->render('sh_index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  股东编辑页面                                       *
     * **********************************************************
     */
    public function actionSh_edit() {
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        $line_id = isset($get['line_id']) ? $get['line_id'] : '';
        $tab_arr = self::get_tab('user_sh');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = [];
        $sh_data = [];
        if ($type == 'update') {
            $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);
            if (empty($data)) {
                return $this->redirect('/agent/shindex');
            }
        }
        $lines = $this->getLines(); //获取线路
        // $ua_data = AgentModel::getMoreAgentByCondition(['is_delete' => 1, 'user_type' => 3, 'agent_type' => 2]);
        $ua_data = [];
        $render = [
            'ua_data' => $ua_data,
            'data' => $data,
            'lines' => $lines,
            'sh_data' => $sh_data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('sh_form.html', $render);
        } else {
            return $this->render('sh_form.html', $render);
        }
    }

    /**
     * **********************************************************
     *  股东保存新增和修改                          *
     * **********************************************************
     */
    public function actionSh_save() {
        $session = Yii::$app->session;
        $cur_user_type = $session->get('user_type');
        $tab_arr = self::get_tab('user_sh');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        //1.管理员 2.管理员子帐号 3.股东子帐号 4.总代子帐号 5.代理子帐号 6.股东 7.总代 8.代理
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $line_id = isset($post['line_id']) ? $post['line_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : 0;
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }

        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        $old = [];
        $str = '';
        if (!empty($id)) {
            unset($data['line_id']);
            $old = AgentModel::getOneAgentByCondition($database, $tab, array('id' => $id));
            //修改
            $data['updatetime'] = time();
            $res = AgentModel::updateAgent($database, $tab, $data, ['id' => $id]);
            $str = '更新';
        } else {
            //新增
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $res = AgentModel::insertAgent($database, $tab, $data);
            $str = '新增';
        }

        if ($res) {
            $data['line_id'] = $line_id;
            self::mongo_log($old, $data, $str . '股东:' . $login_user);
            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }

    /**
     * **********************************************************
     *  总代列表                             *
     * **********************************************************
     */
    public function actionUaindex() {
        $get = Yii::$app->request->get();
        $tab_arr = self::get_tab('user_ua');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = !empty($post['line_id']) ? $post['line_id'] : null; //站点id
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database, $tab, $where);
        $data = AgentModel::getAgentList($database, $tab, $where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $lines = $this->getLines(); //获取线路
        if (empty($get)) {
            $data = array();
            $pagecount = 1;
            $page = 1;
        }
        $render = [
            'data' => $data,
            'lines' => $lines,
            'pagecount' => $pagecount,
            'page' => $page
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('ua_index.html', $render);
        } else {
            return $this->render('ua_index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  总代编辑页面                                       *
     * **********************************************************
     */
    public function actionUa_edit() {
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        $line_id = isset($get['line_id']) ? $get['line_id'] : '';
        $tab_arr = self::get_tab('user_ua');
        $sh_table_arr = self::get_tab('user_sh');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];

        $data = [];
        $ua_data = [];
        $sh_data = [];
        if ($type == 'update') {
            $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);
            if (empty($data)) {
                return $this->redirect('/agent/uaindex');
            }
        }
        $lines = $this->getLines(); //获取线路
        if (isset($data['pid'])) {
            $sh_data = AgentModel::getOneAgentByCondition($sh_table_arr['database'], $sh_table_arr['tab'], ['id' => $data['pid']]);
        }
        $render = [
            'ua_data' => $ua_data,
            'data' => $data,
            'lines' => $lines,
            'sh_data' => $sh_data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('ua_form.html', $render);
        } else {
            return $this->render('ua_form.html', $render);
        }
    }

    /**
     * **********************************************************
     *  总代保存新增和修改                          *
     * **********************************************************
     */
    public function actionUa_save() {
        $session = Yii::$app->session;
        $cur_user_type = $session->get('user_type');
        $tab_arr = self::get_tab('user_ua');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        //1.管理员 2.管理员子帐号 3.股东子帐号 4.总代子帐号 5.代理子帐号 6.股东 7.总代 8.代理
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $line_id = isset($post['line_id']) ? $post['line_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : 0;
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }

        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        $old = [];
        $str = '';
        if (!empty($id)) {
            unset($data['line_id']);
            unset($data['pid']); //不能修改上级
            $old = AgentModel::getOneAgentByCondition($database, $tab, array('id' => $id));
            //修改
            $data['updatetime'] = time();
            $res = AgentModel::updateAgent($database, $tab, $data, ['id' => $id]);
            $str = '更新';
        } else {
            //新增
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $res = AgentModel::insertAgent($database, $tab, $data);
            $str = '新增';
        }

        if ($res) {
            $data['line_id'] = $line_id;
            self::mongo_log($old, $data, $str . '总代：' . $login_user);
            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }

    /**
     * **********************************************************
     *  代理列表                             *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $tab_arr = self::get_tab('user_agent');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = !empty($post['line_id']) ? $post['line_id'] : null; //站点id
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $count = AgentModel::getAgentCount($database, $tab, $where);
        $data = AgentModel::getAgentList($database, $tab, $where, $offset, $pagenum);
        $data = $this->trans($data);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $lines = $this->getLines(); //获取线路
        if (empty($get)) {
            $data = array();
            $pagecount = 1;
            $page = 1;
        }
        $render = [
            'data' => $data,
            'lines' => $lines,
            'pagecount' => $pagecount,
            'page' => $page
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  代理编辑页面                                       *
     * **********************************************************
     */
    public function actionEdit() {
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        $line_id = isset($get['line_id']) ? $get['line_id'] : '';
        $tab_arr = self::get_tab('user_agent');
        $sh_table_arr = self::get_tab('user_sh');
        $ua_table_arr = self::get_tab('user_ua');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];

        $data = [];
        $ua_data = [];
        $sh_data = [];
        if ($type == 'update') {
            $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);
            if (empty($data)) {
                return $this->redirect('/agent/index');
            }
        }
        $lines = $this->getLines(); //获取线路
        if (isset($data['pid'])) {
            $ua_data = AgentModel::getOneAgentByCondition($ua_table_arr['database'], $ua_table_arr['tab'], ['id' => $data['pid']]);
        }
        if (isset($ua_data['pid'])) {
            $sh_data = AgentModel::getOneAgentByCondition($sh_table_arr['database'], $sh_table_arr['tab'], ['id' => $ua_data['pid']]);
        }
        $render = [
            'data' => $data,
            'lines' => $lines,
            'sh_data' => $sh_data,
            'ua_data' => $ua_data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form.html', $render);
        } else {
            return $this->render('form.html', $render);
        }
    }

    /**
     * **********************************************************
     *  代理新增和修改                          *
     * **********************************************************
     */
    public function actionSave() {
        $session = Yii::$app->session;
        $cur_user_type = $session->get('user_type');
        $tab_arr = self::get_tab('user_agent');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        //子账号不能添加代理,即使被业主赋予了操作权限
        //1.管理员 2.管理员子帐号 3.股东子帐号 4.总代子帐号 5.代理子帐号 6.股东 7.总代 8.代理
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $line_id = isset($post['line_id']) ? $post['line_id'] : 0;
        $pid = isset($post['pid']) ? $post['pid'] : '';
        $login_user = isset($post['login_user']) ? $post['login_user'] : '';
        $login_name = isset($post['login_name']) ? $post['login_name'] : '';
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $is_delete = isset($post['is_delete']) ? $post['is_delete'] : '';
        //参数验证
        if (empty($login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号不能为空!']);
            die;
        } else if (!preg_match('/^[A-Za-z0-9_]*$/', $login_user)) {
            echo json_encode(['code' => 400, 'msg' => '账号只能为数字字母下划线!']);
            die;
        } else if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        } else if (empty($login_name)) {
            echo json_encode(['code' => 400, 'msg' => '昵称不能为空!']);
            die;
        } elseif (empty($pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd) && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd && empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }

        $data = [
            'line_id' => $line_id,
            'login_user' => $login_user,
            'login_name' => $login_name,
            'is_delete' => $is_delete,
            'pid' => $pid
        ];

        $old = [];
        $str = '';
        if (!empty($id)) {
            unset($data['line_id']);
            unset($data['pid']); //不能修改上级
            $old = AgentModel::getOneAgentByCondition($database, $tab, array('id' => $id));
            //修改
            $data['updatetime'] = time();
            $res = AgentModel::updateAgent($database, $tab, $data, ['id' => $id]);
            $str = '更新';
        } else {
            //新增
            $data['addtime'] = time();
            $data['login_pwd'] = md5(md5($pwd));
            $res = AgentModel::insertAgent($database, $tab, $data);
            $str = '新增';
        }

        if ($res) {
            $data['line_id'] = $line_id;
            self::mongo_log($old, $data, $str . '代理：' . $login_user);
            echo json_encode(['code' => 200, 'msg' => '保存成功!']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '保存失败!']);
            die;
        }
    }

    //查询条件
    public function where($get) {
        $where = [];
        if (isset($get['login_user']) && !empty($get['login_user'])) {
            $where['login_user'] = trim($get['login_user']);
        }
        if (isset($get['login_name']) && !empty($get['login_name'])) {
            $where['login_name'] = trim($get['login_name']);
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where['is_delete'] = trim($get['status']);
        }

        if (isset($get['line_id']) && !empty($get['line_id'])) {
            $line_id = trim($get['line_id']);
            $where['line_id'] = $line_id;
        }
        return $where;
    }

    //翻译
    public function trans($data) {
        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                //状态
                if ($v['is_delete'] == 1) {
                    $data[$k]['deleteTxt'] = '有效';
                } else if ($v['is_delete'] == 2) {
                    $data[$k]['deleteTxt'] = '无效';
                }
                //时间
                if (!empty($v['addtime'])) {
                    $data[$k]['addDate'] = date('Y-m-d H:i:s', $v['addtime']);
                }
                if (!empty($v['updatetime'])) {
                    $data[$k]['updateDate'] = date('Y-m-d H:i:s', $v['updatetime']);
                }
            }
        }
        return $data;
    }

    //详情页面
    public function actionDetail() {
        $get = Yii::$app->request->get();
        $type = isset($get['type']) ? $get['type'] : '';
        $id = isset($get['id']) ? $get['id'] : '';
        if (empty($type) || !in_array($type, ['user_sh', 'user_ua', 'user_agent']) || empty($id)) {
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }
        $tab_arr = self::get_tab($type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);

        if (empty($data)) {
            echo '<script>alert("获取信息失败"); history.back();</script>';
            exit;
        }

        if ($type == 'user_sh') {
            $data['agent_type_txt'] = '股东';
        } elseif ($type == 'user_ua') {
            $data['agent_type_txt'] = '总代';
            $ptab = \Yii::$app->db->tablePrefix . 'user_sh'; //股东表
            $parent = AgentModel::getOneAgentByCondition($database, $ptab, ['id' => $data['pid']]);
        } elseif ($type == 'user_agent') {
            $ptab = \Yii::$app->db->tablePrefix . 'user_ua'; //总代表
            $parent = AgentModel::getOneAgentByCondition($database, $ptab, ['id' => $data['pid']]);
            $data['agent_type_txt'] = '代理';
        }

        $data['parent'] = isset($parent['login_user']) ? $parent['login_user'] : "";
        $data['parent_name'] = isset($parent['login_name']) ? $parent['login_name'] : "";

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('detail.html', ['data' => $data]);
        } else {
            return $this->render('detail.html', ['data' => $data]);
        }
    }

    //获取线路下的股东
    public function actionGetlines() {
        $post = Yii::$app->request->post();
        $tab_arr = self::get_tab('user_sh');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $line_id = isset($post['line_id']) ? $post['line_id'] : '';
        if (empty($line_id)) {
            echo json_encode(['code' => 400, 'msg' => '线路id不能为空!']);
            die;
        }
        $where = [
            'is_delete' => 1,
            'line_id' => $line_id
        ];
        $data = AgentModel::getMoreAgentByCondition($database, $tab, $where);
        echo json_encode(['data' => $data, 'code' => 200]);
        die;
    }

    //获取股东下的总代
    public function actionGetagents() {
        $post = Yii::$app->request->post();
        $tab_arr = self::get_tab('user_ua');
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $sh_id = isset($post['sh_id']) ? $post['sh_id'] : '';
        if (empty($sh_id)) {
            echo json_encode(['code' => 400, 'msg' => '股东id不能为空!']);
            die;
        }
        $where = [
            'is_delete' => 1,
            'pid' => $sh_id,
        ];
        $data = AgentModel::getMoreAgentByCondition($database, $tab, $where);
        echo json_encode(['data' => $data, 'code' => 200]);
        die;
    }

    //角色表单
    public function actionRole() {
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : null;
        if (empty($type) || !in_array($type, ['user_sh', 'user_ua', 'user_agent']) || empty($id)) {
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }
        $tab_arr = self::get_tab($type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);

        if (empty($data)) {
            echo '<script>alert("获取信息失败"); history.back();</script>';
            exit;
        }
        //所有角色
        $roles = AdminModel::getAllRole();
        $render = [
            'roles' => $roles,
            'data' => $data
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('role.html', $render);
        } else {
            return $this->render('role.html', $render);
        }
    }

    //密码表单
    public function actionPassword() {

        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : null;
        if (empty($type) || !in_array($type, ['user_sh', 'user_ua', 'user_agent']) || empty($id)) {
            echo '<script>alert("id或者type参数丢失"); history.back();</script>';
            exit;
        }
        $tab_arr = self::get_tab($type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $data = AgentModel::getOneAgentByCondition($database, $tab, ['id' => $id]);

        if (empty($data)) {
            return $this->redirect('/agent/index');
        }

        $render = [
            'data' => $data
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('password.html', $render);
        } else {
            return $this->render('password.html', $render);
        }
    }

    //密码修改(入库)
    public function actionSavepwd() {
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $pwd = isset($post['pwd']) ? $post['pwd'] : '';
        $conf_pwd = isset($post['conf_pwd']) ? $post['conf_pwd'] : '';
        $type = isset($post['type']) ? $post['type'] : null;
        if (empty($type) || !in_array($type, ['user_sh', 'user_ua', 'user_agent'])) {
            echo json_encode(['code' => 400, 'msg' => '缺失重要参数!']);
            die;
        }
        if (empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '代理不能为空!']);
            die;
        } else if (empty($pwd)) {
            echo json_encode(['code' => 400, 'msg' => '密码不能为空!']);
            die;
        } elseif (!preg_match('/^[A-Za-z0-9_]*$/', $pwd)) {
            echo json_encode(['code' => 400, 'msg' => '密码只能为数字字母下划线!']);
            die;
        } else if ($pwd != $conf_pwd) {
            echo json_encode(['code' => 400, 'msg' => '两次密码输入不一致!']);
            die;
        }
        $tab_arr = self::get_tab($type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $res = AgentModel::updateAgent($database, $tab, ['login_pwd' => md5(md5($pwd))], ['id' => $id]);

        echo json_encode(['code' => 200, 'msg' => '保存成功!']);
        die;
    }

    //额度表单
    public function actionMoney() {

        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        if (empty($id)) {
            return $this->redirect('/agent/index');
        }
        $data = AgentModel::getOneAgentByCondition(['id' => $id]);

        if (empty($data)) {
            return $this->redirect('/agent/index');
        }

        $render = [
            'data' => $data
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('money.html', $render);
        } else {
            return $this->render('money.html', $render);
        }
    }

    //角色修改(入库)
    public function actionSaverole() {
        $post = Yii::$app->request->post();
        $agent_id = isset($post['agent_id']) ? $post['agent_id'] : '';
        $role_id = isset($post['role_id']) ? $post['role_id'] : '';
        $type = isset($post['type']) ? $post['type'] : null;
        if (empty($type) || !in_array($type, ['user_sh', 'user_ua', 'user_agent'])) {
            echo json_encode(['code' => 400, 'msg' => '缺失重要参数!']);
            die;
        }

        if (empty($agent_id)) {
            echo json_encode(['code' => 400, 'msg' => '代理不能为空']);
            die;
        } else if (empty($role_id)) {
            echo json_encode(['code' => 400, 'msg' => '角色不能为空']);
            die;
        }

        //入库
        $tab_arr = self::get_tab($type);
        $database = $tab_arr['database'];
        $tab = $tab_arr['tab'];
        $res = AgentModel::updateAgent($database, $tab, ['role_id' => $role_id], ['id' => $agent_id]);

        echo json_encode(['code' => 200, 'msg' => '保存成功!']);
        die;
    }

    //额度分配 入库
    public function actionSetmoney() {
        $post = Yii::$app->request->post();
        $id = isset($post['agent_id']) ? $post['agent_id'] : 0;
        $type = isset($post['type']) ? $post['type'] : 0;
        $money = isset($post['money']) ? $post['money'] : 0;
        if (empty($id)) {
            echo json_encode(['code' => 400, 'msg' => '代理不能为空!']);
            die;
        } else if (!in_array($type, [1, 2])) {
            echo json_encode(['code' => 400, 'msg' => '交易模式错误!']);
            die;
        } else if (empty($money)) {
            echo json_encode(['code' => 400, 'msg' => '交易金额不能为空!']);
            die;
        } else if (!preg_match('/^[+]{0,1}(\d+)$|^[+]{0,1}(\d+\.\d+)$/', $money)) {
            echo json_encode(['code' => 400, 'msg' => '交易金额只能为正数!']);
        }


        $arr = [
            'type' => $type,
            'money' => $money,
            'agent_id' => $id
        ];

        $res = AgentModel::setAgentMoney($arr);

        echo json_encode($res);
        die;
    }

    /**
     * **********************************************************
     *  获取数据库及表名           @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function get_tab($type = 'user_sh') {
        $data = array();
        if ($type == 'user_sh') { //股东
            $database = \Yii::$app->db;
            $tab = \Yii::$app->db->tablePrefix . 'user_sh';
        } elseif ($type == 'user_ua') {//总代
            $database = \Yii::$app->db;
            $tab = \Yii::$app->db->tablePrefix . 'user_ua';
        } elseif ($type == 'user_agent') {//代理
            $database = \Yii::$app->db;
            $tab = \Yii::$app->db->tablePrefix . 'user_agent';
        }
        $data['database'] = $database;
        $data['tab'] = $tab;
        return $data;
    }

    //插入mongo日志
    public static function mongo_log($old, $data, $remark) {
        parent::insertOperateLog(json_encode($old), json_encode($data), $remark);
    }

}
