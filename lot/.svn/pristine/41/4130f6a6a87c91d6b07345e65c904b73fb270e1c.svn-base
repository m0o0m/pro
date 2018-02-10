<?php

namespace our_backend\controllers;

use Yii;

class RedisController extends Controller {

    public function actionIndex() {
        if (Yii::$app->request->isAjax) {
            return $this->renderAjax('index.html');
        } else {
            return $this->render('index.html');
        }
    }

    public function actionTodo() {
        $post = \Yii::$app->request->post();
        $get = \Yii::$app->request->get();
        $result = array();
        $redis_type = isset($post['redis']) ? $post['redis'] : '';
        if($redis_type == 'pullredis'){
            $redis = \Yii::$app->pullredis;
        }else{
           $redis = \Yii::$app->redis;
        }
        $todo = isset($post['todo']) ? $post['todo'] : null;
        $result['ErrorCode'] = 2;
        switch ($todo) {
            case 'getkey': //获取所有的Key
                return self::getKey($redis);
            break;

            case 'getval': //根据键获取值
                return self::getVal($redis);
            break;
            case 'getlen'://获取长度
                $res = self::getLen($redis);
                echo json_encode($res);die;
            break;
            case 'del'://清除值
                return self::delKey($redis);
            break;
            default:
                $result['ErrorCode'] = 2;
                $result['ErrorMsg'] = '处理方式不明确';
                return json_encode($result);
                break;
        }

        // var_dump($result);
        return json_encode($result);
    }
/**
      ***********************************************************
      *  获取redis中所有key         @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getKey($redis){
        $result['ErrorCode'] = 2;
        $post = \Yii::$app->request->post();
        $get = \Yii::$app->request->get();
        $type = isset($post['type']) ? $post['type'] : 'all';
        $pix = isset($post['pix']) ? $post['pix'] : null;
        // 获取Redis   Key
        if ($type == 'all') {
            $keys = $redis->keys('*');
        } elseif ($type == 'before' && $pix) {
            $keys = $redis->keys($pix . '*');
        } elseif ($type == 'after' && $pix) {
            $keys = $redis->keys('*' . $pix);
        } else {
            $result['ErrorMsg'] = '参数不完整';
            echo json_encode($result);die;
        }
        if (!empty($keys)) {
            foreach ($keys as $k => $v) {
                $keys[$k] = $v . ' ' . ' （' . $redis->type($v) . ')';
            }
            $result['ErrorCode'] = 1;
            $result['ErrorMsg'] = $keys;
        } else {
            $result['ErrorMsg'] = '当前redis缓存中没有查到符合条件的数据';
        }

        return json_encode($result);
    }
/**
      ***********************************************************
      *  根据键获取值           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getVal($redis){
        $result['ErrorCode'] = 2;
        $post = \Yii::$app->request->post();
        $get = \Yii::$app->request->get();
        $key = isset($post['key']) ? trim($post['key']) : null;
        $r_type = $redis->type($key);
        if ($r_type == 'none') {
            $result['ErrorMsg'] = '无符合条件的键名';
            echo json_encode($result);die;
        }
        $type = isset($post['val_type']) ? $post['val_type'] : $r_type;
        if($r_type == 'zset' && $type == 'set') $type = 'zset';
        if (empty($type))
            $type = $r_type;

        if ($type != $r_type) {
            $result['ErrorMsg'] = '查询不到结果，该key的类型是' . $r_type . ',' . '您选择的是' . $type;
            echo json_encode($result);die;
        }
        $is_json = isset($post['is_json']) ? $post['is_json'] : false;
        if (empty($key)) {
            $result['ErrorMsg'] = 'key不能为空！';
            echo json_encode($result);die;
        }
        if ($type == 'string') { //字符串
            $val = $redis->get($key);
        } elseif ($type == 'hash') { //哈希表
            $len = $redis->hlen($key);
            $keys = $redis->hkeys($key);
            $val = array();
            if($len >= 300){//限制显示300条
                for($i = 0; $i <= 299; $i++){
                    $val[] = $redis->hget($key,$keys[$i]);
                }
            }else{
                $val = $redis->hgetall($key);
            }
        } elseif ($type == 'list') {//列表
            $val = $redis->lrange($key, 0, 500);
        } elseif ($type == 'set') {//集合
            $val = $redis->smembers($key);//不能限定数量，只能全部展示
        } elseif($type == 'zset'){
            $val = $redis->zrange($key, 0 , 500);
        }
        if (empty($val)) {
            $result['ErrorMsg'] = '当前redis缓存中没有查到符合条件的数据';
            echo json_encode($result);die;
        }
        if ($is_json) {
            if (is_array($val)) {
                foreach ($val as $k => $v) {
                    $val[$k] = json_decode($v, true);
                }
            } else {
                $val = json_decode($val, true);
            }
        }
        $val = self::dohtml($val, $is_json);
        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = $val;
        return json_encode($result);
    }
/**
      ***********************************************************
      *  获取redis值的数量           @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function getLen($redis){
        $result['ErrorCode'] = 2;
        $post = \Yii::$app->request->post();
        $get = \Yii::$app->request->get();
        $key = isset($post['key']) ? trim($post['key']) : null;
        $r_type = $redis->type($key);
        $type = isset($post['len_type']) ? $post['len_type'] : $r_type;
        if (empty($type))
            $type = $r_type;
        if($r_type == 'none'){
            $result['ErrorMsg'] = '无符合条件的键名';
            echo json_encode($result);die;
        }
        if($r_type == 'zset' && $type == 'set') $type = 'zset';
        if ($type != $r_type) {
            $result['ErrorMsg'] = '查询不到结果，该key的类型是' . $r_type . ',' . '您选择的是' . $type;
            echo json_encode($result);die;
        }
        $len = 0;
        if ($type == 'string') { //字符串
            $len = $redis->strlen($key);
            $result['ErrorMsg'] = '字符串长度为' . $len . '字节';
        } elseif ($type == 'hash') { //哈希表
            $len = $redis->hlen($key);
            $result['ErrorMsg'] = '哈希表中域的数量是' . $len . '个';
        } elseif ($type == 'list') {//列表
            $len = $redis->llen($key); //获取列表长度
            $result['ErrorMsg'] = '列表的长度是' . $len . '条';
        } elseif ($type == 'set') {//集合
            $len = $redis->scard($key);
            $result['ErrorMsg'] = '集合中元素数量为' . $len . '个';
        } elseif($type == 'zset'){
            $len = $redis->zcard($key);//有序集合
            $result['ErrorMsg'] = '有序集合中元素数量为' . $len . '个';
        }else{
            $result['ErrorMsg'] = '未知';
        }
        $result['ErrorCode'] = 1;
        $result['len'] = $len;
        return $result;
    }
    
/**
      ***********************************************************
      *  清除redis             @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function delKey($redis){
        $result['ErrorCode'] = 2;
        $post = \Yii::$app->request->post();
        $get = \Yii::$app->request->get();
        $type = isset($post['type']) ? $post['type'] : 'all';
        $pix = isset($post['pix']) ? $post['pix'] : null;
        if (empty($pix)) {
            $result['ErrorMsg'] = '请输入要清除的Key';
            return json_encode($result);
        }

        // 获取Redis   Key
        if ($type == 'before' && $pix) {
            $keys = $redis->keys($pix . '*');
        } elseif ($type == 'after' && $pix) {
            $keys = $redis->keys('*' . $pix);
        } else {
            $keys = $pix;
            $is_exists = $redis->exists($keys);
        }
        if (!isset($is_exists) && !is_array($keys)) {
            $result['ErrorMsg'] = '键名不存在';
            return json_encode($result);
        }
        if(is_array($keys)){
            foreach($keys as $val){
                $res = $redis->del($val);
            }
        }else{
            $res = $redis->del($keys);
        }
        if ($res) {
            $result['ErrorCode'] = 1;
        } else {
            $result['ErrorMsg'] = '清除失败！';
        }
        return json_encode($result);
    }
/**
      ***********************************************************
      *  将查询出的redis值拼装成html格式 @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public static function dohtml($val, $is_json = true) {
        //拼接html,最多支持三维数组
        $tmp_str = '<div class="arr_main">';
        $width = '';
        if(!$is_json) $width = '100%';
        if (is_array($val)) {

            foreach ($val as $k => $v) {
                $tmp_str .= '<div class="arr_one" style="word-wrap:break-word;width:' . $width . ';">';
                if ((!empty($v)) && is_array($v)) {
                    foreach ($v as $k2 => $v2) {
                        $tmp_str .= '<div class="arr_two" style="border:.5px #ddd dotted;margin-top:1px;">';
                        if ((!empty($v2)) && is_array($v2)) {
                            foreach ($v2 as $k3 => $v3) {
                                $tmp_str .= ' &nbsp;' . $k3 . ' => ' . $v3 . ' &nbsp;';
                            }
                        } else {
                            $tmp_str .= $k2 . ' => ' . $v2;
                        }
                        $tmp_str .= '</div>'; //two
                    }
                } else {
                    $tmp_str .= $k . ' => ' . $v;
                }

                $tmp_str .= '</div>'; //one
            }
        } else {
            $tmp_str = $val;
        }
        $tmp_str .= '<div>'; //main

        return $tmp_str;
    }

/**
      ***********************************************************
      *  查看队列中剩余数据条数        @author ruizuo qiyongsheng    *
      ***********************************************************
*/
    public function actionPush_count(){
        $post = \Yii::$app->request->post();
        $result = array();
        $result['ErrorCode'] = 2;

        $name = isset($post['push_name']) ? trim($post['push_name']) : null;
        if(empty($name)){
            $result['ErrorMsg'] = '请输入队列名称';
            return json_encode($result);
        }
        // if($name == 'balance') $name = 'balance_sql';//结算sql队列
        if($name == 'total') $name = 'total_bet';//统计队列
        //读取队列数据
        $httpsqs= Yii::$app->httpsqs;
        $data = $httpsqs->getStatusJson($name);
        if($data){
            $data = json_decode($data,true);
        }else{
            $result['ErrorMsg'] = '暂时无法连接到队列，请稍候重试';
            return json_encode($result);
        }
        $unread = $data['unread'];
        $result['ErrorCode'] = 1;
        $result['ErrorMsg'] = '队列中现有未读取数据：<font color="red">' . $unread . ' </font>条';
        return json_encode($result);
    }
    
    
}
