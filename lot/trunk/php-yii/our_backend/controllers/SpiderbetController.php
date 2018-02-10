<?php

namespace our_backend\controllers;

use Yii;
use our_backend\controllers\Controller;
use our_backend\models\SpiderbetModel as selfModel;

class SpiderbetController extends Controller {
/**
	  ***********************************************************
	  *  注单补采纪录列表页           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public function actionIndex(){
		$get = Yii::$app->request->get();
        $line_list = $this->getLines();
        $games = $this->getAllFcTypes();
        if (isset($get['_pjax']))
            unset($get['_pjax']);
        if (empty($get)) {
            $render = [
                'data' => array(),
                'lines' => $line_list,
                'games'=>$games,
                'starttime' => date('Y-m-d'),
                'endtime' => date('Y-m-d'),
                'pagecount' => 1,
                'page' => 1
            ];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html', $render);
            } else {
                return $this->render('index.html', $render);
            }
        }
        $fc_type = isset($get['fc_type']) ? $get['fc_type'] : null;
        $addday = isset($get['addday']) ? $get['addday'] : null;
        $line_id = isset($get['line_id']) ? $get['line_id'] : null;
        $status = isset($get['status']) ? $get['status'] : null;
        $page = isset($get['page']) ? $get['page'] : 1;
        $pagenum = isset($get['pagenum']) ? $get['pagenum'] : 100;
        $offset = ($page - 1) * $pagenum;
        $starttime = isset($get['starttime']) ? $get['starttime'] : date('Y-m-d');
        $endtime = isset($get['endtime']) ? $get['endtime'] : date('Y-m-d');
        if (!empty($starttime)) {
            $starttime = strtotime($starttime) ? strtotime($starttime . '00:00:00') : null;
        }
        if (!empty($endtime)) {
            $endtime = strtotime($endtime) ? strtotime($endtime . ' 23:59:59') : null;
        }

        $where = array();
        if ($line_id)
            $where['line_id'] = $line_id;
        if ($fc_type)
            $where['fc_type'] = $fc_type;
        if ($status)
            $where['status'] = $status;

        if ($starttime && $endtime) {
            $where = ['and', $where, array('between', 'addtime', $starttime, $endtime)];
        } elseif (!empty($starttime)) {
            $where = ['and', $where, 'addtime>=' . $starttime];
        } elseif (!empty($endtime)) {
            $where = ['and', $where, array('<=', 'addtime', $endtime)];
        }
        $count = selfModel::getCount($where);
        $pagecount = ceil($count / $pagenum);
        $pagecount = empty($pagecount) ? 1 : $pagecount;
        if ( $count ) {
            $data = selfModel::getData('*', $where, $offset, $pagenum);
        	$trans = array('', '正在处理中', '完成');
        	$game_arr = array();
            if(!empty($data)){
            	foreach($games as $val){
            		$game_arr[$val['type']] = $val['name'];
            	}
            	foreach($data as $key=>$val){
            		$data[$key]['addday'] = date('Y-m-d',strtotime($val['addday']));
            		$data[$key]['addtime'] = date('Y-m-d H:i:s',$val['addtime']);
            		$data[$key]['statusTxt'] = isset($trans[$val['status']]) ? $trans[$val['status']] : '未知';
            		$data[$key]['fc_type'] = isset($game_arr[$val['fc_type']]) ? $game_arr[$val['fc_type']] : '全部';
                    if(!empty($val['remark'])){
                        $tmp = '';
                        $tmp = str_replace('total_num', '采集注单总条数', $val['remark']);
                        $tmp = str_replace('fail_num', '插入失败条数', $tmp);
                        $tmp = str_replace(',', '条 ', $tmp) . '条';
                        $data[$key]['remark'] = $tmp;
                    }
            	}
            }
        } else {
            $data = array();
        }

        $render = [
            'data' => $data,
            'lines' => $line_list,
            'games'=>$games,
            'starttime' => date('Y-m-d',$starttime),
            'endtime' => date('Y-m-d',$endtime),
            'pagecount' => 1,
            'page' => 1
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
	}
	
/**
	  ***********************************************************
	  *  注单补采集           @author ruizuo qiyongsheng    *
	  ***********************************************************
*/
	public function actionGetdata(){
		$post = Yii::$app->request->post();
		$line_id = isset($post['line_id']) ? $post['line_id'] : '';
		$addday = isset($post['addday']) ? $post['addday'] : '';
		$fc_type = isset($post['fc_type']) ? $post['fc_type'] : '';
		
		$result = array();
		$result['ErrorCode'] = 2;
		if(empty($line_id)){
			$result['ErrorMsg'] = '请选择线路';
			echo json_encode($result);die;
		}
        if(empty($addday)){
            $result['ErrorMsg'] = '请选择日期';
            echo json_encode($result);die;
        }
		$addday = date('Ymd',strtotime($addday));

        $redis = Yii::$app->redis;
		$redis_key = $line_id . '_' . $addday . '_' . $fc_type;
		$is_exists = $redis->exists($redis_key);
        if(!empty($fc_type)){
            $is_all_exists = $redis->exists($line_id . '_' . $addday . '_');
        }else{
            $is_all_exists = false;
        }
        if($is_exists || $is_all_exists){
        	$result['ErrorMsg'] = '24小时内同样的操作只能进行一次';
			echo json_encode($result);die;
        }

        $data = array();
        $session = Yii::$app->session;
        $data['uid'] = $session['uid'];
        $data['uname'] = $session['login_user'];
        $data['line_id'] = $line_id;
        $data['fc_type'] = $fc_type;
        $data['addday'] = $addday;
        $data['addtime'] = time();
        $data['status'] = 1;
        $res = selfModel::insert($data);
        if(!$res){
        	$result['ErrorMsg'] = '添加补采数据失败';
			echo json_encode($result);die;
        }

        $redis->setex($redis_key, 86400 , 1);
        $redis->lpush('spiderBetForApi_list',json_encode(['addtime'=>$data['addtime'], 'line_id'=>$line_id, 'addday'=>$addday, 'fc_type'=>$fc_type]));

		$result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '发送请求成功,请稍后查看结果';
		echo json_encode($result);die;
	}
	
	


}