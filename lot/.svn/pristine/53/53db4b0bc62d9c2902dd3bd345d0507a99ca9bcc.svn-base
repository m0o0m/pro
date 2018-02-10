<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\AgentmenuModel as selfModel;
use yii\helpers\ArrayHelper;

class AgentmenuController extends Controller {

    /**
     * 菜单列表
     * @menu      string  菜单地址
     * @menu_name string  菜单名称
     * @is_delete int     是否有效
     * @is_url    int     是否菜单栏显示
     * @level     int     菜单级别
     * @pagenum   int     每页显示条数
     * @page      int     第几页
     * @pagecount int     总页数
     * */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'page' => 1,
                'pagecount' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $where = $this->where($get);
        $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $offset = ($page - 1) * $pagenum;
        $data = selfModel::getMenuList($where, $offset, $pagenum);
        $count = selfModel::getMenuCount($where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        $data = $this->trans($data);

        $data = ArrayHelper::index($data, 'id');
        $tmp = [];
        foreach ($data as $val) {
            if (isset($data[$val['pid']])) {
                $data[$val['pid']]['children'][] = &$data[$val['id']];
            } else {
                $tmp[] = &$data[$val['id']];
            }
        }
        $data = $tmp;

        $render = [
            'data' => $data, 'pagecount' => $pagecount, 'page' => $page
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    public function where($get) {
        $where = [];
        if (isset($get['menu']) && !empty($get['menu'])) {
            $where['menu'] = trim($get['menu']);
        }
        if (isset($get['menuName']) && !empty($get['menuName'])) {
            $where['menu_name'] = trim($get['menuName']);
        }
        if (isset($get['status']) && !empty($get['status'])) {
            $where['is_delete'] = trim($get['status']);
        }
        if (isset($get['isUrl']) && !empty($get['isUrl'])) {
            $where['is_url'] = trim($get['isUrl']);
        }
        if (isset($get['menuLevel']) && !empty($get['menuLevel'])) {
            $where['level'] = trim($get['menuLevel']);
        }
        return $where;
    }

    //翻译
    public function trans($data) {
        if (is_array($data) && !empty($data)) {
            foreach ($data as $k => $v) {
                //级别
                if ($v['level'] == 1) {
                    $data[$k]['levelTxt'] = '一级';
                } else if ($v['level'] == 2) {
                    $data[$k]['levelTxt'] = '二级';
                } else if ($v['level'] == 3) {
                    $data[$k]['levelTxt'] = '三级';
                }
                //是否页面可见
                if ($v['is_url'] == 1) {
                    $data[$k]['urlTxt'] = '是';
                } else if ($v['is_url'] == 2) {
                    $data[$k]['urlTxt'] = '否';
                }
                //状态
                if ($v['is_delete'] == 1) {
                    $data[$k]['deleteTxt'] = '有效';
                } else if ($v['is_delete'] == 2) {
                    $data[$k]['deleteTxt'] = '无效';
                }
            }
        }
        return $data;
    }

    //获取单条数据
    public function actionOnedata() {
        $post = Yii::$app->request->post();
        if (!isset($post['id']) || empty($post['id'])) {
            echo json_encode(['code' => 400, 'msg' => 'id 不能为空']);
            die;
        }
        $data = selfModel::getOneData($post['id']);
        //获取一级菜单
        $data['oneLevels'] = selfModel::getSubMenu(['level' => 1, 'is_delete' => 1]);
        //获取二级菜单
        if ($data['level'] == 3) {
            $pdata = selfModel::getOneData($data['pid']);
            $data['ppid'] = $pdata['pid'];
            //获取二级菜单
            $data['twoLevels'] = selfModel::getSubMenu(['pid' => $pdata['pid'], 'is_delete' => 1]);
        } elseif ($data['level'] == 2) {
            //获取二级菜单
            $data['twoLevels'] = selfModel::getSubMenu(['pid' => $data['pid'], 'is_delete' => 1]);
        }
        echo json_encode(['code' => 200, 'data' => $data]);
        die;
    }

    //获取子类菜单
    public function actionSubmenu() {
        $post = Yii::$app->request->post();
        if (!isset($post['id']) || empty($post['id'])) {
            $where = ['level' => 1, 'is_delete' => 1];
        } else {
            $where = ['pid' => $post['id'], 'is_delete' => 1];
        }
        $data = selfModel::getSubMenu($where);
        echo json_encode(['code' => 200, 'data' => $data]);
        die;
    }

    //保存菜单(新增,修改)
    public function actionSavemenu() {
        $post = Yii::$app->request->post();
        $error = false;
        $msg;
        if (empty($post['menuName']) || !isset($post['menuName'])) {
            $error = true;
            $msg = '菜单名称不能为空!';
        } else if (empty($post['menuUrl']) || !isset($post['menuUrl'])) {
            $error = true;
            $msg = '菜单地址不能为空!';
        } else if (!preg_match('/^[A-Za-z0-9\/#]*$/', $post['menuUrl'])) {
            $error = true;
            $msg = '菜单地址只能为数字字母斜杠#';
        }

        if ($error) {
            echo json_encode(['code' => 400, 'msg' => $msg]);
            die;
        }

        //数据组装
        if (!empty($post['twoLevel'])) {
            $pid = $post['twoLevel'];
            $level = 3;
        } elseif (!empty($post['oneLevel'])) {
            $pid = $post['oneLevel'];
            $level = 2;
        } else {
            $pid = 0;
            $level = 1;
        }
        $data = [
            'pid' => $pid,
            'level' => $level,
            'is_delete' => $post['isUse'],
            'is_url' => $post['isSee'],
            'is_agent' => $post['is_agent'],
            'menu_name' => $post['menuName'],
            'menu' => $post['menuUrl'],
            'icon_class' => $post['iconClass'],
            'sort' => $post['sort']
        ];

        //入库
        if (!isset($post['menuId']) || empty($post['menuId'])) {
            //新增
            if (selfModel::add($data)) {
                echo json_encode(['code' => 200, 'msg' => '保存成功']);
                die;
            } else {
                echo json_encode(['code' => 400, 'msg' => '保存失败']);
                die;
            }
        } else {
            //修改
            selfModel::update($post['menuId'], $data);
            echo json_encode(['code' => 200, 'msg' => '保存成功']);
            die;
        }
    }

}
