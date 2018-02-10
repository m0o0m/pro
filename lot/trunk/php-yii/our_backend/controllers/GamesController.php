<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\GamesModel;

class GamesController extends Controller {

    /**
     * **********************************************************
     *  彩票列表页   			  @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        //查询彩种分类
        $redis = \Yii::$app->redis;
        $cat_arr = $redis->get('admin_games_cat');
        if ($cat_arr) {
            $cat_arr = json_decode($cat_arr, true);
        } else {
            $cat_tab = GamesModel::getSubTableName();
            $field = 'type,name';
            $cat_arr = GamesModel::getData($cat_tab, 'order_by', $field, array());
            $redis->setex('admin_games_cat', 600, json_encode($cat_arr));
        }

        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'cat_arr' => $cat_arr,
                'data' => array()
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $name = isset($get['name']) ? $get['name'] : null;
        $type = isset($get['type']) ? $get['type'] : null;
        $state = isset($get['state']) ? $get['state'] : null;
        $hot = isset($get['hot']) ? $get['hot'] : null;
        $order_by = isset($get['order_by']) ? $get['order_by'] : 2;
        $ltype = isset($get['ltype']) ? $get['ltype'] : null;
        $table = GamesModel::getTableName();
        //拼接条件
        $where = array();
        if (!empty($state))
            $where['state'] = $state;
        if (!empty($hot))
            $where['is_hot'] = $hot;
        if (!empty($name))
            $where = ['and', $where, ['like', 'name', $name]];
        if (!empty($type))
            $where = ['and', $where, ['like', 'type', $type]];
        if (!empty($ltype))
            $where = ['and', $where, ['like', 'ltype', $ltype]];
        //查询条数
        $count = GamesModel::getCount($table, $where);
        $data = array();
        if ($count != 0) {
            if ($order_by == 2) {
                $order_by = 'order_by desc';
            } else {
                $order_by = 'order_by asc';
            }
            $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 1000;
            $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
            $offset = ($page - 1) * $pagenum;
            $field = '*';
            $data = GamesModel::getData($table, $order_by, $field, $where, $offset, $pagenum);
            //翻译
            $state_arr = array('', '有效', '无效', '维护');
            foreach ($data as $key => $val) {
                if (in_array($val['state'], array(1, 2, 3))) {
                    $data[$key]['stateTxt'] = strtr($val['state'], $state_arr);
                }
                $data[$key]['is_hotTxt'] = self::is_hot($val['is_hot']);
                $data[$key]['is_recomTxt'] = self::is_hot($val['is_recom']);
            }
        }

        foreach ($data as $key => $value) {
            if (!empty($value['img_path'])) {
                $data[$key]['img_href'] = Yii::$app->params['cdn_href'] . '/images/lotterytype/' . $value['img_path'];
            } else {
                $data[$key]['img_href'] = '';
            }
        }
        $render = [
            'data' => $data,
            'cat_arr' => $cat_arr
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    /**
     * **********************************************************
     *  修改彩种   			  @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionEditgame() {
        $post = Yii::$app->request->post();
        $id = isset($post['id']) ? intval($post['id']) : null;
        $type = isset($post['type']) ? trim($post['type']) : null;
        $name = isset($post['name']) ? trim($post['name']) : null;
        $path = isset($post['path']) ? trim($post['path']) : null;
        $is_hot = isset($post['is_hot']) ? intval(trim($post['is_hot'])) : 1;
        $is_recom = isset($post['is_recom']) ? intval(trim($post['is_recom'])) : 1;
        $state = isset($post['state']) ? intval(trim($post['state'])) : 1;
        $ltype = isset($post['ltype']) ? $post['ltype'] : null;
        $lname = isset($post['lname']) ? $post['lname'] : null;
        $order_by = isset($post['order_by']) ? intval(trim($post['order_by'])) : 0;
        $template  = isset($post['template']) ? $post['template'] : '';
        $table = GamesModel::getTableName();
        $result = $map = array();
        $result['ErrorCode'] = 2;
        if (empty($ltype) || empty($lname)) {
            $result['ErrorMsg'] = '请选择彩种分类';
            return json_encode($result);
        }
        if (empty($name) || empty($type)) {
            $result['ErrorMsg'] = '请填写彩种名称';
            return json_encode($result);
        }

        if (!preg_match('/^[_0-9a-z]{1,16}$/', $type)) {
            $result['ErrorMsg'] = '非法彩种类型！';
            return json_encode($result);
        }

        if (!preg_match('/[a-zA-Z]{1}/', $type)) {
            $result['ErrorMsg'] = '非法彩种类型！';
            return json_encode($result);
        }

        if (empty($id)) {
            //添加新彩种
            $where = array();
            $where['type'] = $type;
            $count = GamesModel::getCount($table, $where);
            if ($count) {
                $result['ErrorMsg'] = '此彩种已存在';
                return json_encode($result);
            }

            $data = array();
            $data['ltype'] = $ltype;
            $data['type'] = $type;
            $data['name'] = $name;
            $data['state'] = $state;
            $data['is_hot'] = $is_hot;
            $data['is_recom'] = $is_recom;
            $data['order_by'] = $order_by;
            $data['ltype_name'] = $lname;
            $data['template'] = $template;
            if (!empty($path)) {
                $data['img_path'] = $path;
            }

            $res = GamesModel::insert($table, $data);
            if (!$res) {
                $result['ErrorMsg'] = '数据添加失败';
            } else {
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '数据添加成功';
            }
            return json_encode($result);
        } else {
            //修改新采种
            $where = array();
            $where['id'] = $id;
            $where['type'] = $type;

            $data = array();
            $data['ltype'] = $ltype;
            $data['name'] = $name;
            $data['state'] = $state;
            $data['is_hot'] = $is_hot;
            $data['is_recom'] = $is_recom;
            $data['order_by'] = $order_by;
            $data['ltype_name'] = $lname;
            $data['template'] = $template;
            if (!empty($path)) {
                $data['img_path'] = $path;
            }

            $res = GamesModel::update($table, $data, $where);

            if (!$res) {
                $result['ErrorMsg'] = '修改失败,数据未变更或其它原因';
            } else {
                $redis = Yii::$app->redis;
                $redis_key = 'c_lot_all_game_site';
                $temp = $redis->del($redis_key);
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '修改成功';
            }
            return json_encode($result);
        }
    }

    public static function is_hot($type) {
        switch ($type) {
            case '0':
                $result = '';
                break;
            case '1':
                $result = '否';
                break;
            case '2':
                $result = '是';
                break;
            default:
                $result = '';
                break;
        }
        return $result;
    }

    // 彩票分类列表
    public function actionCaigame() {
        $get = Yii::$app->request->get();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'cat_arr' => array()
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('caigame.html', $render);
            } else {
                return $this->render('caigame.html', $render);
            }
        }
        $name = isset($get['name']) ? $get['name'] : null;
        $type = isset($get['type']) ? $get['type'] : null;
        $order_by = isset($get['order_by']) ? $get['order_by'] : 2;
        $table = GamesModel::getSubTableName();
        //拼接条件
        $where = array();
        if (!empty($name))
            $where = ['like', 'name', $name];
        if (!empty($type))
            $where = ['and', $where, ['like', 'type', $type]];
        //查询条数
        $count = GamesModel::getCount($table, $where);
        $cat_data = $cat_arr = array();
        $data = array();
        if ($count != 0) {
            if ($order_by == 2) {
                $order_by = 'order_by desc';
            } else {
                $order_by = 'order_by asc';
            }
            $pagenum = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
            $page = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
            $offset = ($page - 1) * $pagenum;
            $field = 'id,name,type,img_path,order_by';
            $data = GamesModel::getData($table, $order_by, $field, $where, $offset, $pagenum);
        }
        foreach ($data as $key => $value) {
            if (!empty($value['img_path'])) {
                $data[$key]['img_href'] = Yii::$app->params['cdn_href'] . '/images/lotterytype/' . $value['img_path'];
            } else {
                $data[$key]['img_href'] = '';
            }
        }
        $render = [
            'data' => $data,
            'cat_arr' => $cat_data
        ];
        //print_r($render);die;
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('caigame.html', $render);
        } else {
            return $this->render('caigame.html', $render);
        }
    }

    public function actionAddgame() {
        $post = Yii::$app->request->post();
        $redis = Yii::$app->redis;

        $id = isset($post['id']) ? intval($post['id']) : null;
        $type = isset($post['type']) ? trim($post['type']) : null;
        $name = isset($post['name']) ? trim($post['name']) : null;
        $path = isset($post['path']) ? trim($post['path']) : null;
        $order_by = isset($post['order_by']) ? intval(trim($post['order_by'])) : 0;
        $table = GamesModel::getSubTableName();
        $result = $map = array();
        $result['ErrorCode'] = 2;
        if (empty($name) || empty($type)) {
            $result['ErrorMsg'] = '请填写彩票名称';
            return json_encode($result);
        }

        if (empty($id)) {
            //添加新彩票
            $where = array();
            $where['type'] = $type;
            $count = GamesModel::getCount($table, $where);
            if ($count) {
                $result['ErrorMsg'] = '此彩票已存在';
                return json_encode($result);
            }
            $data = array();
            $data['type'] = $type;
            $data['name'] = $name;
            $data['order_by'] = $order_by;
            if (!empty($path)) {
                $data['img_path'] = $path;
            }

            $res = GamesModel::insert($table, $data);
            if (!$res) {
                $result['ErrorMsg'] = '数据添加失败';
            } else {
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '数据添加成功';
            }
            return json_encode($result);
        } else {
            //查询彩票修改前的数据
            $where = array();
            $where['id'] = $id;
            $field = 'name,type';
            $data_res = GamesModel::getData($table, '', $field, $where, '', '');
            //修改彩种表中对应的彩票信息
            $tables = GamesModel::getTableName();
            $wheres = array();
            $wheres['ltype'] = $data_res[0]['type'];

            $datas = array();
            $datas['ltype_name'] = $name;
            $datas['ltype'] = $type;
            $res_d = GamesModel::update($tables, $datas, $wheres);
            $redis->del('admin_games_cat'); //清除缓存
            //修改彩票信息
            $data = array();
            $data['name'] = $name;
            $data['type'] = $type;
            $data['order_by'] = $order_by;
            $data['img_path'] = $path;

            $res = GamesModel::update($table, $data, $where);

            if (!$res) {
                $result['ErrorMsg'] = '修改失败,数据未变更或其它原因';
            } else {
                $redis = Yii::$app->redis;
                $redis_key = 'games_cat';
                $redis_keys = 'c_lot_all_game_site';
                $temps = $redis->del($redis_keys);
                $temp = $redis->del($redis_key);
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '修改成功';
            }
            return json_encode($result);
        }
    }

}
