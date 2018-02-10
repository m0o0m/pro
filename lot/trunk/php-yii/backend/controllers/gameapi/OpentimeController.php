<?php

namespace backend\controllers\gameapi;

use Yii;
use backend\controllers\Controller;
use common\helpers\Helper;
use common\models\OpentimeModel as otime;

class OpentimeController extends Controller {

    /**
     * **********************************************************
     *  开盘时间列表展示         @author  Rom  *
     * **********************************************************
     */
    public function actionIndex() {
        $get = Yii::$app->request->get();
        $page = isset($get['page']) ? intval($get['page']) : 1;
        $pagenum = isset($get['pageNum']) ? intval($get['pageNum']) : 100;
        $fc_type = isset($get['fc_type']) ? $get['fc_type'] : 'fc_3d';
        $qishu = isset($get['qishu']) ? $get['qishu'] : null;
        //判断分页是否正确
        if (!is_int($page) || $page <= 0) {
            $page = 1;
        }
        if (!is_int($pagenum) || $pagenum == 0) {
            $pagenum = 100;
        } else {
            $pagenum = abs($pagenum);
        }

        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $key => $val) {
            $tmp[$val['type']] = $val;
        }
        $games = $tmp;

       $tabname = Helper::GetOpenTimeTableNameByType($fc_type); //获取表名
        
        $where = array();
        if ($qishu && is_numeric($qishu) && $qishu > 0) {
            $where['qishu'] = $qishu;
        }
        if(!in_array($fc_type, ['liuhecai','jnd_bs','jnd_28'])){
            $where['fc_type'] = $fc_type;
        }
        $totalnum = otime::getDataTotalNum($tabname, $where); //获取当前彩种的总条数
        $total_page = ceil($totalnum / $pagenum); //总页码数
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        $condition['offset'] = ($page - 1) * $pagenum;  //开始条数
        $condition['limit'] = $pagenum;
        $condition['tabname'] = $tabname;
        $condition['field'] = array('id', 'qishu', 'kaipan', 'fengpan', 'kaijiang', 'status');
        $data = otime::index($condition, $where);

        foreach ($data as $k => $v) {
            $data[$k]['fc_type'] = $games[$fc_type]['name'];
        }
        $pagedata['page'] = $page;
        $pagedata['totalpage'] = $total_page;
        $pagedata['totalnum'] = $totalnum;
    
        $render = ['data' => $data, 'games' => $games, 'pagedata' => $pagedata, 'fc_type' => $fc_type];
        
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

}
