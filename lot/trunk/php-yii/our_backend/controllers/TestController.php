<?php

namespace our_backend\controllers;

class TestController extends Controller {

	public function actionIndex(){
		$insertId = $this->insertLine();
	}

	public function insertLine(){
		$cols['line_id'] = $where['line_id'] = 'aaa';
		$cols['line_name'] = '';
		$cols['money'] = '0.00';
		$cols['url'] = '';
		$cols['addtime'] = time();
		$cols['updatetime'] = 0;
		$cols['type'] = 1;
		$cols['status'] = 1;
		$cols['md5key'] = '';
		$cols['deskey'] = '';
		$cols['is_shiwan'] = 1;
		if (empty($cols['line_id'])) {
			die('为空');
		}
		$has = \common\models\LineModel::getCount($where);
		if ($has > 0) {
			die('重名');
		}
		\common\models\LineModel::insert($cols);

		return \Yii::$app->db->getLastInsertID();
	}

}