<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AdminModel;

class RoleController extends Controller {

    public function actionIndex() {
        $get = Yii::$app->request->get();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'pagecount' => 1,
                'page' => 1
            ];

            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $where = ['and'];
        if (isset($get['roleName']) && !empty($get['roleName'])) {
            $where[] = 'role_name="' . trim($get['roleName']) . '"';
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where[] = 'is_delete=' . trim($get['status']);
        }
        //添加时间
        if (isset($get['startTime']) && !empty($get['startTime'])) {
            $where[] = ['>=', 'addtime', strtotime(trim($get['startTime']))];
        }
        if (isset($get['endTime']) && !empty($get['endTime'])) {
            $where[] = ['<=', 'addtime', strtotime(trim($get['endTime']))];
        }
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $data = AdminModel::getRoleList($where, $offset, $pagenum);
        $data = $this->trans($data);
        $count = AdminModel::getRoleCount($where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
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

    //权限编辑
    public function actionEdit() {
        $get = Yii::$app->request->get();
        $id = isset($get['id']) ? $get['id'] : '';
        $type = isset($get['type']) ? $get['type'] : '';
        $data = [];
        $access = [];
        if (!empty($id)) {
            $data = AdminModel::getOneRole(['id' => $id]);
            $data['addTimeTxt'] = !empty($data['addtime']) ? date('Y-m-d H:i:s', $data['addtime']) : '';
            $data['updateTimeTxt'] = !empty($data['updatetime']) ? date('Y-m-d H:i:s', $data['updatetime']) : '';
            $access = explode(',', $data['access_id']);
        }
        //所有路由
        $routes = AdminModel::getAllRoute();

        $oneLevel = [];
        $twoLevel = [];
        $threeLevel = [];
        if (!empty($routes)) {
            foreach ($routes as $k => $v) {
                if ($v['level'] == 1) {
                    $oneLevel[] = $v;
                } else if ($v['level'] == 2) {
                    $twoLevel[] = $v;
                } else if ($v['level'] == 3) {
                    $threeLevel[] = $v;
                }
            }
        }

        $orm = [
            'id' => $id,
            'type' => $type,
            'data' => $data,
            'access' => $access,
            'oneLevel' => $oneLevel,
            'twoLevel' => $twoLevel,
            'threeLevel' => $threeLevel
        ];

        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('form.html', $orm);
        } else {
            return $this->render('form.html', $orm);
        }
    }

    public function actionSave() {
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? $post['id'] : '';
        $role = isset($post['role']) ? $post['role'] : '';
        $roleName = isset($post['roleName']) ? $post['roleName'] : '';
        $status = isset($post['status']) ? $post['status'] : '';
        $access = isset($post['access']) ? $post['access'] : [];
        $accessStr = implode(',', $access);

        if (empty($role)) {
            echo json_encode(['code' => 400, 'msg' => '角色不能为空!']);
            die;
        }else if (!preg_match('/^[A-Za-z0-9_]*$/', $role)) {
            echo json_encode(['code' => 400, 'msg' => '角色只能为数字字母下划线!']);
            die;
        } else if (empty($roleName)) {
            echo json_encode(['code' => 400, 'msg' => '角色名称不能为空!']);
            die;
        }

        // 唯一
        $has = AdminModel::getRoleCount(['and', ['=', 'role', $role], ['<>', 'id', $id]]);
        if ($has > 0) {
            echo json_encode(['code' => 400, 'msg' => '角色重名!']);
            die;
        }

        $arr = [
            'role' => $role,
            'role_name' => $roleName,
            'access_id' => $accessStr,
            'is_delete' => $status,
            'handler' => Yii::$app->session->get('login_name')
        ];

        if (!empty($id)) {
            $arr['updatetime'] = time();
            $res = AdminModel::updateRole($id, $arr);
        } else {
            $arr['addtime'] = time();
            $res = AdminModel::addRole($arr);
        }

        if ($res) {
            echo json_encode(['code' => 200, 'msg' => '更新成功']);
            die;
        } else {
            echo json_encode(['code' => 400, 'msg' => '更新失败']);
            die;
        }
    }

}
