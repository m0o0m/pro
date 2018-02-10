<?php

namespace our_backend\controllers;

use Yii;
use common\models\LogModel;
use common\helpers\mongoTables;
use our_backend\controllers\Controller;
use common\helpers\Curl;

class LogController extends Controller {

    //请求日志
    public function actionRequest() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html');
            } else {
                return $this->render('index.html');
            }
        }

        $page_data = $this->getPage($get);
        $where = [];
        //php路由
        if (isset($get['route']) && !empty($get['route'])) {
            $where['route'] = trim($get['route']);
        }
        $where = $this->getWhere($get, $where);
        //获取要查询的表名
        $mongoTable = mongoTables::getTable('httpRequest');
        $count = LogModel::getLogsCount($mongoTable, $where);
        if ($count <= 0) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('index.html');
            } else {
                return $this->render('index.html');
            }
        }

        $orderby = ['addtime' => SORT_DESC];
        $data = LogModel::getLogs($mongoTable, $where, $page_data['offset'], $page_data['pageNum'], $orderby);
        $data = $this->trans($data);
        $pagecount = ceil($count / $page_data['pageNum']);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page_data['page']
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html', $render);
        } else {
            return $this->render('index.html', $render);
        }
    }

    //翻译
    public function trans($data) {
        foreach ($data as $k => $v) {
            if (isset($v['old']) && is_array($v['old']))
                $data[$k]['old'] = json_encode($v['old']);
            if (isset($v['new']) && is_array($v['new']))
                $data[$k]['new'] = json_encode($v['new']);
            if (isset($v['ptype'])) {
                switch ($v['ptype']) {
                    case 1:
                        $data[$k]['ptypeTxt'] = '总后台';
                        break;
                    case 2:
                        $data[$k]['ptypeTxt'] = '业主后台';
                        break;
                    case 3:
                        $data[$k]['ptypeTxt'] = '前台PC';
                        break;
                    case 4:
                        $data[$k]['ptypeTxt'] = '前台WAP';
                        break;
                    case 5:
                        $data[$k]['ptypeTxt'] = 'APP';
                        break;
                    default :
                        $data[$k]['ptypeTxt'] = '未知';
                        break;
                }
            }
            if (isset($v['addtime'])) {
                $data[$k]['addtime'] = date('Y-m-d H:i:s', $v['addtime']);
            }
            //登录状态
            if (isset($v['state'])) {
                if ($v['state'] == 1)
                    $data[$k]['newState'] = '登录成功';
                if ($v['state'] == 2)
                    $data[$k]['newState'] = '登录失败';
            }
            //中文名字
            if (isset($v['menu_name'])) {
                $data[$k]['menu_name'] = (string) $v['menu_name'];
            }
        }
        return $data;
    }

    /**
     * **********************************************************
     *  获取条件               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function getWhere($get, $where = array()) {
        //平台
        if (isset($get['ptype']) && !empty($get['ptype'])) {
            $where['ptype'] = (int) trim($get['ptype']);
        }
        //日期
        $starttime = isset($get['starttime']) ? trim($get['starttime']) . ' 00:00:00' : '';
        $endtime = isset($get['endtime']) ? trim($get['endtime']) . '23:59:59' : '';
        $starttime = strtotime($starttime) ? strtotime($starttime) : '';
        $endtime = strtotime($endtime) ? strtotime($endtime) : '';
        if ($starttime && $endtime) {
            $where['addtime'] = ['$gt' => $starttime, '$lt' => $endtime];
        } elseif ($starttime) {
            $where['addtime'] = ['$gt' => $starttime];
        } elseif ($endtime) {
            $where['addtime'] = ['$lt' => $endtime];
        }
        return $where;
    }

    /**
     * **********************************************************
     *  获取页码                  @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function getPage($get) {
        $result = array();
        $result['pageNum'] = isset($get['pageNum']) && !empty($get['pageNum']) ? $get['pageNum'] : 100;
        $result['page'] = isset($get['page']) && !empty($get['page']) ? $get['page'] : 1;
        $result['offset'] = ( $result['page'] - 1 ) * $result['pageNum'];
        return $result;
    }

    /**
     * **********************************************************
     *  历史登录日志             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionHistory() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('history.html');
            } else {
                return $this->render('history.html');
            }
        }

        $page_data = $this->getPage($get);
        $where = [];
        //会员帐号
        if (isset($get['uname']) && !empty($get['uname'])) {
            $where['uname'] = trim($get['uname']);
        }
        $where = $this->getWhere($get, $where);
        //获取要查询的表名
        $mongoTable = mongoTables::getTable('historyLogin');
        $count = LogModel::getLogsCount($mongoTable, $where);
        if ($count <= 0) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('history.html');
            } else {
                return $this->render('history.html');
            }
        }

        $orderby = ['addtime' => SORT_DESC];
        $data = LogModel::getLogs($mongoTable, $where, $page_data['offset'], $page_data['pageNum'], $orderby);
        $data = $this->trans($data);
        $pagecount = ceil($count / $page_data['pageNum']);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page_data['page']
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('history.html', $render);
        } else {
            return $this->render('history.html', $render);
        }
    }

    //其它操作日志
    public function actionOperate() {
        $query = $get = Yii::$app->request->get();
        unset($query['_pjax']);
        if (empty($query)) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('operate.html');
            } else {
                return $this->render('operate.html');
            }
        }

        $page_data = $this->getPage($get);
        $where = [];
        //操作人
        if (isset($get['uname']) && !empty($get['uname'])) {
            $where['uname'] = trim($get['uname']);
        }
        $where = $this->getWhere($get, $where);
        //获取要查询的表名
        $mongoTable = mongoTables::getTable('operate');
        $count = LogModel::getLogsCount($mongoTable, $where);
        if ($count <= 0) {
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('operate.html');
            } else {
                return $this->render('operate.html');
            }
        }

        $orderby = ['addtime' => SORT_DESC];
        $data = LogModel::getLogs($mongoTable, $where, $page_data['offset'], $page_data['pageNum'], $orderby);
        $data = $this->trans($data);
        $pagecount = ceil($count / $page_data['pageNum']);
        $pagecount = empty($pagecount) ? 1 : $pagecount;

        $render = [
            'data' => $data,
            'pagecount' => $pagecount,
            'page' => $page_data['page']
        ];
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('operate.html', $render);
        } else {
            return $this->render('operate.html', $render);
        }
    }

    /**
     * **********************************************************
     *  读取日志文件             @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public function actionLogfile() {

        $post = Yii::$app->request->post();
        $page = isset($post['page']) ? $post['page'] : 1;
        if ($page < 1)
            $page = 1;
        $pagenum = 14; //默认每页行数

        $config = Yii::$app->params;
        $server = $config['hostdir'];
        $preg = '/.*log.*/'; //日志文件规则
        $preg_php = '/.*\.php/'; //php文件后缀

        $result = array();
        $result['ErrorCode'] = 2;
        //获取文件夹根目录
        $base_dir = isset($post['basename']) ? (base64_decode($post['basename'])) : $server . $config['basename'] . '/';
        //当前目录上层目录
        $this_path = str_replace($server, ' ', $base_dir);
        $is_base = dirname($this_path);
        $before_path = $is_base . '/';
        if ($before_path == './') {
            $before_path = $server . $config['basename'] . '/';
        }
        $before_path = base64_encode($server . trim($before_path));
        //获取要执行的指令
        $type = isset($post['type']) ? $post['type'] : 'dir';
        //数据类型
        $data_type = isset($post['data_type']) ? $post['data_type'] : 'html';
        //检测是不是目录或者文件
        $tmp_file = rtrim($base_dir, '/');
        if ((!is_dir($base_dir)) && (!file_exists($tmp_file))) {
            $result['ErrorMsg'] = '非法目录或文件';
            if ($is_base == '.') {
                $result['ErrorMsg'] = '已经是根目录';
            }
            return json_encode($result);
        }

        //如果是目录则遍历,是文件则读取
        if ($type == 'dir') {
            $file = self::getfile($base_dir);
            //检测是不是日志文件  过滤所有非日志文件
            if (!empty($file['file'])) {
                foreach ($file['file'] as $key => $val) {
                    if ((!preg_match($preg, $val)) || (preg_match($preg_php, $val))) {
                        unset($file['file'][$key]);
                    }
                }
            }

            if (empty($file['dir']) && empty($file['file'])) {
                $result['ErrorMsg'] = '已经是最后一层目录，目录里没有日志文件';
                return json_encode($result);
            }

            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = array_merge($file['file'], $file['dir']);
            $result['path'] = $this_path; //当前目录
            $result['beforePath'] = $before_path; //上层目录
            $result['is_base'] = $is_base; //是否根目录
        } elseif ($type == 'file') {
            if (!file_exists($tmp_file)) {
                $result['ErrorMsg'] = '非法文件';
                return json_encode($result);
            }
            //获取文件总行数
            @$handle = fopen($tmp_file, 'r'); //只读模式打开文件，指针指向文件起始位置
            if (!$handle) {
                $result['ErrorMsg'] = '读取文件失败';
                return json_encode($result);
            }
            fseek($handle, 0, SEEK_SET); //指针设置在文件开头
            $lineCount = 0;
            while (!feof($handle)) {
                $emails = fgets($handle); //注意不能不要变量,否则指针不会移动到下行.
                $lineCount++;
            }

            rewind($handle); //重置指针
            //关闭文件
            fclose($handle);
            if ($lineCount < 1) {
                $result['ErrorMsg'] = '空白文件';
                return json_encode($result);
            }
            $total_page = ceil($lineCount / $pagenum); //总页码数
            if ($page > $total_page && $total_page != 0)
                $page = $total_page;
            $offset = ($page - 1) * $pagenum + 1;  //开始条数
            $limit = $offset + $pagenum - 1;       //每页显示条数
            $content = self::getFileLines($tmp_file, $offset, $limit);
            $result['ErrorCode'] = 3;
            $result['ErrorMsg'] = $content;
            $result['totalpage'] = $total_page;
            $result['page'] = $page;
            if (empty($content)) {
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '此文件不存在或者没有读取权限';
            }
            return json_encode($result);
        }

        if ($data_type == 'html') {
            return $this->render('logfile.html', [
                        'data' => $result['ErrorMsg'],
                        'path' => $this_path,
                        'before' => $before_path
                            ]
            );
        } elseif ($data_type == 'ajax') {
            return json_encode($result);
        }
        // echo '<pre>';
        // var_dump($file);
    }

    /**
     * **********************************************************
     *  获取目录下所有目录和文件      @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getfile($dir) {
        $dir_list = array();
        $dir_list['file'] = array();
        $dir_list['dir'] = array();
        if (is_dir($dir)) {
            if ($source = opendir($dir)) {
                while (($file = readdir($source)) !== false) {
                    //加密路径
                    $path = base64_encode($dir . $file . '/');
                    if ((is_dir($dir . '/' . $file)) && $file != '.' && $file != '..') {
                        $dir_list['dir'][] = '<img src="/public/images/folder.gif" width="20"> <a href="javascript:void(0);" onclick="getmore(' . "'dir',this)" . '" rel="' . $path . '" >' . $file . '</a>';
                    } else {
                        if ($file != '.' && $file != '..') {
                            $dir_list['file'][] = '<img src="/public/images/page.gif" width="20"> <a href="javascript:void(0);" onclick="getmore(' . "'file',this)" . '" rel="' . $path . '" >' . $file . '</a>';
                        }
                    }
                }

                closedir($source);
            }
        }

        return $dir_list;
    }

    /**
     * **********************************************************
     *  读取日志文件               @author ruizuo qiyongsheng    *
     * **********************************************************
     */
    public static function getFileLines($filename, $startLine = 1, $endLine = 50, $method = 'rb') {
        $content = array();
        $count = $endLine - $startLine;
        // 判断php版本（因为要用到SplFileObject，PHP>=5.1.0）
        if (version_compare(PHP_VERSION, '5.1.0', '>=')) {
            $fp = new \SplFileObject($filename, $method);
            $fp->seek($startLine - 1); // 转到第N行, seek方法参数从0开始计数
            for ($i = 0; $i <= $count; ++$i) {
                $content[] = '<p>' . $fp->current() . '</p>'; // current()获取当前行内容
                $fp->next(); // 下一行
            }
        } else {//PHP<5.1
            $fp = fopen($filename, $method);
            if (!$fp)
                return 'error:can not read file';
            for ($i = 1; $i < $startLine; ++$i) {// 跳过前$startLine行
                fgets($fp);
            }
            for ($i; $i <= $endLine; ++$i) {
                $content[] = fgets($fp); // 读取文件行内容
            }
            fclose($fp);
        }
        return array_filter($content); // array_filter过滤：false,null,''
    }


 /**
          ***********************************************************
          *  mongo表及redis数据预警   @author ruizuo qiyongsheng    *
          ***********************************************************
*/

    public function actionMongocount(){
        $games = $this->getAllFcTypes();
        $tmp = [];
        foreach ($games as $key => $val) {
            $tmp[$val['type']] = $val;
        }
        $games = $tmp;

        $redis = \Yii::$app->redis;
        $session = Yii::$app->session;
        $uid = $session['uid'];
        $redis_key = 'getRedisCount_' . $uid ;
        if($redis->get($redis_key)){
            $render = ['data'=>array(), 'arr'=> array(), 'games'=>$games, 'wait'=> true];
            if (Yii::$app->request->isAjax) {
                return $this->renderAjax('count.html', $render);
            } else {
                return $this->render('count.html', $render);
            }
        }
        //mongo
        $mongo_arr = mongoTables::allTables();
        $type_arr = array_keys($mongo_arr);
        $table_arr = array_values($mongo_arr);
        $count = count($table_arr);
        $data = array();
        for($i = 0; $i < $count; $i++){
            $data[$i]['type'] = isset($type_arr[$i]) ? $type_arr[$i] : '';
            $data[$i]['table'] = isset($table_arr[$i]) ? $table_arr[$i] : '';
            switch ($type_arr[$i]) {
                case 'httpRequest':
                        $data[$i]['name'] = '请求日志';
                    break;
                case 'historyLogin':
                        $data[$i]['name'] = '历史登陆';
                    break;
                case 'operate':
                     $data[$i]['name'] = '操作日志';
                    break;
                case 'auto':
                     $data[$i]['name'] = '开奖网开奖结果';
                break;
                case 'bets':
                    $data[$i]['name'] = '注单备份';
                break;
                default:
                    $data[$i]['name'] = '未知表名';
                    break;
            }
        }

        foreach($data as $key=>$val){
            $data[$key]['count'] = LogModel::getLogsCount($val['table'], array());
        }
        //redis
        $redis_arr = $this->getRedisLen();
        
        $redis->setex($redis_key, 5, 1);

        $render = ['data'=>$data, 'arr'=> $redis_arr,  'games'=>$games];


        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('count.html', $render);
        } else {
            return $this->render('count.html', $render);
        }
    }

    public function getRedisLen(){
        $arr = array();
        $pullredis = \Yii::$app->pullredis;
        $redis = \Yii::$app->redis;
        //获取下注 结算时的缓存
        $games = $this->getAllFcTypes();
        $type_arr = $new_games = array();
        foreach($games as $val){
            $type_arr[] = $val['type'];
            $new_games[$val['type']] = $val['name'];
        }
        $balance_key_pix = 'for_balance-';
        $balance_keys = $redis->keys($balance_key_pix . '*');
        if($balance_keys){
            foreach($balance_keys as $k=>$keyname){
                $tmp = explode('-', $keyname);
                if(count($tmp) == 3 && in_array($tmp[1], $type_arr) && is_numeric($tmp[2])){
                    $arr[$k]['key'] = $keyname;
                    $arr[$k]['fc_type'] = $new_games[$tmp[1]];
                    $arr[$k]['qishu'] = $tmp[2];
                    $arr[$k]['count'] = $redis->hlen($keyname);
                }
            }
        }

        //api采集缓存
        $line_arr = $this->getLines();
        $lines = $spider_arr =  array();
        foreach($line_arr as $key=> $line){
            $tmp_key = 'spider_' . $line['line_id'] . '_data';
            $spider_arr[$key]['line_id'] = $line['line_id'];
            $spider_arr[$key]['key'] = $tmp_key;
            $spider_arr[$key]['count'] = $pullredis->zcard($tmp_key);
        }

        //统计队列
        $total_bet_key = 'AccountFromRedis';
        $total_count = $redis->llen($total_bet_key);

        $return = array();
        $return['balance'] = $arr;
        $return['spider'] = $spider_arr;
        $return['totalbet'] = $total_count;
        return $return;
    }

    /**
      ***********************************************************
      *  请求查看结算进程内存数据条数     @author ruizuo qiyongsheng    *
      ***********************************************************
    */
        public function actionMemorycount(){
            $post = Yii::$app->request->post();
            $type = isset($post['type']) ? $post['type'] : 'fc_3d';
            $result = [];
            $result['ErrorCode'] = 2;
            if (empty($host = Yii::$app->params['app_lottery_host_http'])) {
                $result['ErrorMsg'] = '未配置 : app_lottery_host_http';
                return $result;
             }
            // $host = '127.0.0.1:10002';
            $res = Curl::run($host, 'post', array(
                'todo' => 'getMemoryCount', 
                'fc_type' => $type, //彩种名称
            ));

            if($res === false){
                $result['ErrorMsg'] = 'common进程未开启';
                echo json_encode($result); die;
            }

            if(empty($res)){
                $result['ErrorCode'] = 1;
                $result['ErrorMsg'] = '结算内存中目前没有数据';
                echo json_encode($result); die;
            }

            $res = json_decode($res, true);
            if(isset($res['ErrorCode']) && $res['ErrorCode'] == 2){
                echo json_encode($res); die;
            }

            $str = '';
            foreach($res as $val){
                $str .= '期数:' . $val['qishu'] . ' 条数: ' . $val['count']  .  '<br/>';
            }
                
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = $str;
            echo json_encode($result); die;   
        }
        
}
