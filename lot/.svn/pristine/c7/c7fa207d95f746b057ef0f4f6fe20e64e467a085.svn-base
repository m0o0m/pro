<?php

namespace frontend_api\controllers;

use Yii;
use frontend_api\controllers\Controller;
use frontend_api\models\SiteModel;

/**
 * pc wap 首页相关接口
 */
class SiteController extends Controller {
    
    public function actionIndex() {
        echo 'index';
    }

    // 获取维护状态
    public function actionMaintain() {
        $data['ErrorCode'] = 1;
        $data['ErrorMsg'] = 'OK';
        $clients = ['pc', 'wap'];
        foreach ($clients as $client) {
            $maintain = $this->game_is_maintain(1, $client, $this->user->line_id);
            if ($maintain['return'] == 1) {
                $data['Data'][$client] = 1;
            } else {
                $data['Data'][$client] = 2;
            }
        }
        return json_encode($data);
    }

    //wap首页轮播图
    public function actionSlideimg() {
        /*$line_id = $this->user->line_id;
        $upurl = isset(Yii::$app->params['upurl']) ? Yii::$app->params['upurl'] : '';
        $redis_key = 'wap' . $line_id . 'flash';
        $redis = Yii::$app->redis;
        $data = $redis->get($redis_key);
        if (empty($data)) {
            $where = array('line_id' => $line_id, 'enable' => 1);
            $order = 'sort';
            $data = SiteModel::get_banner_data($where, $order);
            $redis->set($redis_key, json_encode($data));
        } else {
            $data = json_decode($data, true);
        }
        if (!empty($data)) {
            foreach ($data as $k => $v) {
                $res = [];
                $res['sort'] = $v['sort'];
                $res['imgName'] = $v['filename'];
                $res['path'] = Yii::$app->params['upurl'] . $v['filepath'];
                $bannerlist[] = $res;
            }
        } else {
            //没有
            $bannerlist = [];
        }*/

        $banners[] = Yii::$app->params['cdn_href'] . '/images/banner/banner1.png';
        $banners[] = Yii::$app->params['cdn_href'] . '/images/banner/banner2.png';
        $banners[] = Yii::$app->params['cdn_href'] . '/images/banner/banner3.png';

        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '执行成功';
        $result['Data'] = $banners;
        return json_encode($result);
    }

    //公告消息
    public function actionNotice() {
        //取用户信息
        $line_id = $this->user->line_id;
        $ptype = $this->user->client;
        $redis = Yii::$app->redis;
        $redis_key = $line_id . $ptype . "_site_notice"; //线路公告信息的key
        $temp = $redis->get($redis_key);
        if ($temp) {
            $list = json_decode($temp, true);
        } else {
            $where[] = 'and';
            $where[] = ['or', ['=', 'line_id', $line_id], ['=', 'line_id', '']];
            $where[] = ['or', ['=', 'type', $ptype], ['=', 'type', '0']];
            $where[] = ['=', 'status', 1];
            $list = SiteModel::getNoticeInfo($where);
            $redis->set($redis_key, json_encode($list));
        }

        $data['Data'] = array();
        foreach ($list as $key => $val) {
            $data['Data'][$key] = $val;
            $data['Data'][$key]['addtime'] = date('Y-m-d', $val['addtime']);
        }
        $data['ErrorCode'] = 1;
        $data['Message'] = '执行成功';
        echo json_encode($data);
        die;
    }

    //首页所有彩种菜单
    public function actionAlltype() {
        $line_id = $this->user->line_id;
        //所有彩种
        $all_data = $this->getAllFcTypes();
        //彩种分类
        $cat_data = $this->getAllFcCates();
        //我的收藏
        $like_data = $this->getMyLovesFc();

        $data = $this->reGroupData($all_data, $cat_data, $like_data, $line_id);
        echo json_encode($data);
        die;
    }

    /**
     * **********************************************************
     * 处理成接口需要的数据          @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function reGroupData($data, $cat_data, $like_data, $line_id) {
        $alltype = $onetype = $hot = $recommend = $like = [];
        $like_arr = explode(',', $like_data['fc_types']);

        //所有彩种维护信息
        $ptype = $this->user->client;
        $maintain_data = $this->getAllTypeMaintain(2, $ptype, $line_id);

        $qishu_arr = array();//所有彩种期数

        foreach ($data as $key => $vv) {
            $qishu_arr[$vv['type']] = $this->get_fc_qishu($vv['type']); //当前期数

            //我的收藏 我喜欢的
            if (in_array($vv['type'], $like_arr)) {
                $data[$key]['is_like'] = 2;
            } else {
                $data[$key]['is_like'] = 1;
            }

            $is_maintain = $maintain_data[$vv['type']];
            $data[$key]['is_wh'] = $is_maintain['return'];
            $data[$key]['wh_content'] = $is_maintain['remark'];
            $data[$key]['wh_starttime'] = $is_maintain['starttime'];
            $data[$key]['wh_endtime'] = $is_maintain['endtime'];
            
            $push = [
                'type' => $vv['type'],
                'name' => $vv['name'],
                'ltype' => $vv['ltype'],
                'img_path' => $vv['img_path'],
                'is_hot' => $vv['is_hot'],
                'is_recom' => $vv['is_recom'],
                'is_like' => $data[$key]['is_like'],
                'is_wh' => $data[$key]['is_wh'],
                'template' => $vv['template'],
                'wh_content' => $data[$key]['wh_content'],
                'wh_starttime' => $data[$key]['wh_starttime'],
                'wh_endtime' => $data[$key]['wh_endtime']
            ];


            //热门
            if ($vv['is_hot'] == 2) {
                $hot[] = $push;
            }

            //推荐
            if ($vv['is_recom'] == 2) {
                $recommend[] = $push;
            }
            //收藏
            if ($data[$key]['is_like'] == 2) {
                $like[] = $push;
            }
        }

        //重组数据
        $last_auto_arr = $this->getLastAuto();
        foreach ($cat_data as $k => $v) {
            $onetype['ltype'] = $v['type'];
            $onetype['lname'] = $v['name'];
            $onetype['lsrc'] = $v['img_path'];
            $onetype['types'] = [];
            foreach ($data as $key => $val) {
                $qishu = $qishu_arr[$val['type']]; //当前期数
                $auto = $last_auto_arr[$val['type']];
                $auto_qishu = isset($auto['qishu']) ? $auto['qishu'] : 0;
                $auto = isset($auto['ball']) ? $auto['ball'] : '';
                if ($v['type'] == $val['ltype']) {
                    $tmp = [
                        'type' => $val['type'],
                        'name' => $val['name'],
                        'ltype' => $val['ltype'],
                        'img_path' => $val['img_path'],
                        'is_hot' => $val['is_hot'],
                        'is_recom' => $val['is_recom'],
                        'is_like' => $val['is_like'],
                        'is_wh' => $val['is_wh'],
                        'wh_content' => $val['wh_content'],
                        'template' => $val['template'],
                        'auto_qishu' => $auto_qishu,
                        'qishu' => $qishu,
                        'auto' => $auto
                    ];
                    $onetype['types'][] = $tmp;
                }
            }

            $alltype[] = $onetype;
        }


        //处理返回接口
        $return = array();
        $return['Data']['all_type'] = $alltype;
        $return['Data']['hot'] = $hot;
        $return['Data']['recommend'] = $recommend;
        $return['Data']['like'] = $like;
        $return['ErrorCode'] = 1;
        $return['ErrorMsg'] = '获取成功';
        return $return;
    }

    /**
     * **********************************************************
     * 返回维护中的所有彩种的数组     @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function get_is_wh($line_id, $type) {
        $redis = Yii::$app->redis;
        //全网维护key
        $key = "maintain_pc_all_line_ids";
        $keys = "maintain_pc_one_line_ids_" . $line_id;

        $data = $data2 = array();
        $res = $res1 = $res2 = $content = $content1 = array();

        $temp_data = $redis->get($key); //全网维护数据
        $data = json_decode($temp_data, true);
        $temp_data = $redis->get($keys); //单线维护数据
        $data2 = json_decode($temp_data, true);
        if (isset($data[0]['module'])) {//全网维护彩种
            $res1 = json_decode($data[0]['module'], true);
        }
        if (isset($data[0]['content'])) {//全网维护原因
            $content = $data[0]['content'];
        }
        if (isset($data2['module'])) {
            $res2 = json_decode($data2['module'], true);
        }
        if (isset($data2['content'])) {
            $content1 = $data2['content'];
        }
        $res['content'] = null;
        if (!empty($content) && in_array($type, $res1)) {
            $res['content'] = $content;
        } elseif (!empty($content1) && in_array($type, $res2)) {
            $res['content'] = $content1;
        }
        $type_res = array_merge($res1, $res2);
        if (in_array($type, $type_res)) {
            $msg = "全网维护,维护原因: " . $res['content'];
            $arr = array();
            $arr['is_wh'] = true;
            $arr['msg'] = $msg;
            return $arr;
        } else {
            return false;
        }
    }


}
