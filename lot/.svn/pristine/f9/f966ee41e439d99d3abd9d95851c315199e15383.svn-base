<?php
namespace backend\controllers\gameapi;

use Yii;
use backend\controllers\Controller;
use common\models\AutoModel as selfModel;
use common\models\LogModel;
use common\helpers\lotteryOrm;

class AutoController extends Controller {

    public function actionIndex() {
        $type_arr = $this->getAllFcTypes();
        $data = array();
        foreach ($type_arr as $key => $value) {
            $data[$key]['fc_type'] = $value['type'];
            $data[$key]['fc_name'] = $value['name'];
        }
        $data['games'] = $data;

        $data['trans'] = $this->transData();
        $result = $this->search();

        $render = [
            'data' => $data,
            'result' => $result
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }


    /*
     * ******************************************************************************
     */

    public function search() {
        $query = $request = Yii::$app->request->get();
        unset($query['_pjax']);
        if(empty($query)){
            return;
        }

        $fc_type = !empty($request['fc_type']) ? trim($request['fc_type']) : '';
        $qishu = isset($request['qishu']) ? floatval($request['qishu']) : 0;
        $starttime = isset($request['starttime']) ? trim($request['starttime']) : '';
        $endtime = isset($request['endtime']) ? trim($request['endtime']) : '';
        $status = isset($request['status']) ? trim($request['status']) : 0;
        $page = isset($request['page']) ? intval($request['page']) : 1;
        $pagenum = isset($request['pagenum']) ? intval($request['pagenum']) : 10;

        if(!$fc_type){
            $result['ErrorCode'] = 2;
            $result['ErrorMsg'] = '请选择彩种';
            return $result;
        }

        $where = ['and'];
        // $where[] = ['status'=>1];
        if ($qishu) {
            $where[] = ['like', 'qishu', $qishu];
        }
        // 时间 处理
        $starttime = strtotime($starttime) ? strtotime($starttime) : $starttime;
        $endtime = strtotime($endtime) ? strtotime($endtime) : $endtime;
        // 时间筛选
        if ($starttime && $endtime) {
            $where[] = ['between', 'addtime', $starttime, $endtime];
        } elseif (!empty($starttime)) {
            $where[] = ['>', 'addtime', $starttime];
        } elseif (!empty($endtime)) {
            $where[] = ['<', 'addtime', $endtime];
        }
        if ($status) {
            $where[] = ['=', 'status', $status];
        }

        // 分页
        $total_count = selfModel::getCount($fc_type, $where);
        $total_page = ceil($total_count / $pagenum);
        if ($page > $total_page && $total_page != 0)
            $page = $total_page;
        if ($page <= 0)
            $page = 1;
        $offset = ($page - 1) * $pagenum;
        $limit = $pagenum;

        $rows = selfModel::getList($fc_type, $where, $offset, $limit);

        $rows = $this->trans($rows);

        $lotteryOrm = new lotteryOrm();
        $ball_num = $lotteryOrm->getBallNum($fc_type);
        //取出开奖结果
        $tmp_arr = array();
        for($i = 1; $i <= $ball_num ; $i++){
            $tmp_arr[] = 'ball_' . $i;
        }
        foreach($rows as $key=>$val){
            foreach($val as $k=>&$v){
                $ball = array();
                if(in_array($k, $tmp_arr)){
                    $ball = explode(',', $v);
                    $rows[$key][$k] = $ball[0];
                }
            }
        }

        $oauto = LogModel::getLogs('auto',['fc_type'=>$fc_type],0,999,[]);
        $diff = [];
        foreach($rows as &$item){
            foreach($oauto as $key => $auto){
                if($item['qishu'] == $auto['expect']){
                    $opencode = explode(',', $auto['opencode']);
                    $item['oauto'] = $opencode;
                    for ($i=0; $i < $ball_num; $i++) {
                        if ($item['ball_' . ($i+1)] != $opencode[$i]) {
                            $item['diff'][] = $i;
                        }
                    }
                }
            }
        }unset($item);

        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $key => $val) {
            $tmp[$val['type']] = $val;
        }
        $games = $tmp;
        $result['ErrorCode'] = 1;
        $result['fc_type'] = $games[$fc_type]['name'];
        $result['ball_num'] = $ball_num;
        $result['data'] = $rows;
        $result['diff'] = $diff;
        $result['page'] = $page;
        $result['pagenum'] = $pagenum;
        $result['totalpage'] = $total_page;
        return $result;
    }

    public function transData(){

        $result['status'] = [1=>'未结算',2=>'已结算',3=>'正在结算'];

        return $result;
    }

    public function trans($data) {
        foreach ($data as $k => &$v) {
            $v['datedate'] = date('Y-m-d H:i:s', $v['datetime']);
            $v['adddate'] = date('Y-m-d H:i:s', $v['addtime']);
            $v['datedatetime'] = date('Y-m-d H:i:s', $v['datetime']);
        }unset($v);
        return $data;
    }

}
