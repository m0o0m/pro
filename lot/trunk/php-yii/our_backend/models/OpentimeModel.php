<?php

namespace our_backend\models;

use Yii;
use yii\db\Query;

class OpentimeModel extends \yii\base\Model {

    /**
     * **********************************************************
     *  开盘时间表查询 Lottery库my_ah_k3_opentime表    	    *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param string 表名
     * @param int page 当前页码 default 1
     * @param int pagenum 每页显示条数 default 200
     * @return array 处理好的数组，包含错误码与查询出的数据
     * **********************************************************
     */
    public static function index($fc_type, $page = 1, $pagenum = 200) {
        $query = new Query;
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $fc_type; //表名	
        $new_data = array(); //返回数据数组
        //判断数据表是否存在
        $table_array = $connection->createCommand("show tables")->queryAll();
        $new_table_array = array();
        foreach ($table_array as $val) {
            foreach ($val as $v) {
                $new_table_array[] = $v;
            }
        }
        if (!in_array($table_name, $new_table_array)) {
            $new_data['ErrorCode'] = 2;
            $new_data['data'] = array();
            $new_data['page'] = 1;
            $new_data['totalpage'] = 0;
            return $new_data;
        }
        //判断分页是否正确
        if (!is_int($page) || $page <= 0) {
            $page = 1;
        }
        if (!is_int($pagenum) || $pagenum == 0) {
            $pagenum = 100;
        } else {
            $pagenum = abs($pagenum);
        }

        //查询满足条件的数据总条数
        $total_count = $query->from($table_name)->count();
        if ($total_count != 0) {
            //根据总条数判断传来的页码数是否合法
            $total_page = ceil($total_count / $pagenum); //总页码数
            if ($page > $total_page && $total_page != 0)
                $page = $total_page;
            $offset = ($page - 1) * $pagenum;  //开始条数
            $limit = $pagenum;       //每页显示条数
            //过滤字段
            $field = array('id', 'qishu', 'kaipan', 'fengpan', 'kaijiang', 'status');
            $query->select($field)->from($table_name)->offset($offset)->limit($limit);
            $data = $query->all();
        }else {
            $data = array();
            $total_page = 0;
            $page = 1;
        }
        $new_data = array();
        $new_data['ErrorCode'] = 1;  //错误号
        $new_data['data'] = $data; //查询出的数据
        $new_data['page'] = $page; //当前页码
        $new_data['totalpage'] = $total_page; //总页码

        return $new_data;
    }

    /**
     * **********************************************************
     *  开封盘时间添加(提交表单)                     *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param string 表名
     * @param int 期数
     * @param string 开奖时间
     * @param string 开盘时间
     * @param string 封盘时间
     * @return array 处理好的结果数组
     * **********************************************************
     */
    public static function open_add($fc_type, $qishu, $kaipan, $fengpan, $kaijiang) {
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $fc_type; //表名	
        $new_data = array(); //返回数据数组
        //验证数据完整性
        if (empty($fc_type) || empty($qishu) || empty($kaipan) || empty($fengpan) || empty($kaijiang)) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据不完整';
            return $new_data;
        }
        //验证数据类型
        if (!is_int($qishu) || $qishu == 0) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据类型不正确';
            return $new_data;
        }
        $time_preg = '/^([0-1]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$/';
        $kaipan_p = preg_match($time_preg, $kaipan);
        $fengpan_p = preg_match($time_preg, $fengpan);
        $kaijiang_p = preg_match($time_preg, $kaijiang);
        if ($kaipan_p == 0 || $fengpan_p == 0 || $kaijiang_p == 0) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据类型不正确';
            return $new_data;
        }

        //判断数据表是否存在
        $table_array = $connection->createCommand("show tables")->queryAll();
        $new_table_array = array();
        foreach ($table_array as $val) {
            foreach ($val as $v) {
                $new_table_array[] = $v;
            }
        }
        if (!in_array($table_name, $new_table_array)) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据表不存在';
            return $new_data;
        }

        //处理数组
        $data = array();
        $data['qishu'] = $qishu;
        $data['kaipan'] = $kaipan;
        $data['fengpan'] = $fengpan;
        $data['kaijiang'] = $kaijiang;
        $row = $connection->createCommand()->insert($table_name, $data)->execute();
        if ($row) {
            $new_data['ErrorCode'] = 1;
            $new_data['ErrorMsg'] = '增加成功';
            return $new_data;
        } else {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '增加失败';
            return $new_data;
        }
    }

    /**
     * **********************************************************
     *      开封盘时间 编辑页面                    *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param int 主键
     * @param string 表名
     * @return array 处理好的数组，包含错误号 查询出的数据 
     * **********************************************************
     */
    public static function open_edit_info($id, $fc_type) {
        $query = new Query;
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $fc_type . '_opentime'; //表名	
        $new_data = array(); //返回数据数组
        //验证数据完整性
        if (empty($id) || empty($fc_type) || $id == 0) {
            $new_data['ErrorCode'] = 2;
            $new_data['data'] = array();
            return $new_data;
        }
        //判断数据表是否存在
        $table_array = $connection->createCommand("show tables")->queryAll();
        $new_table_array = array();
        foreach ($table_array as $val) {
            foreach ($val as $v) {
                $new_table_array[] = $v;
            }
        }
        if (!in_array($table_name, $new_table_array)) {
            $new_data['ErrorCode'] = 2;
            $new_data['data'] = array();
            return $new_data;
        }

        //过滤字段
        $field = array('qishu', 'kaipan', 'fengpan', 'kaijiang');
        //查询数据
        $query->select($field)->where(['id' => $id])->from($table_name);
        $data = $query->all();
        if (!empty($data)) {
            $data['ErrorCode'] = 1;
            $data[0]['fc_type'] = $fc_type;
        }
        return $data;
    }

    /**
     * **********************************************************
     *      开封盘时间 修改（提交表单）                     *****
     * **********************************************************
     * @author ruizuo qiyongsheng
     * @param int 主键
     * @param string 表名
     * @param int 期数
     * @param string 开盘时间
     * @param string 封盘时间
     * @param string 开奖时间
     * @return array 处理好的数组 包含错误号和错误信息
     * **********************************************************
     */
    public static function open_update($id, $fc_type, $qishu, $kaipan, $fengpan, $kaijiang) {
        $connection = \Yii::$app->db;
        $prefix = $connection->tablePrefix; //表前缀
        $table_name = $prefix . $fc_type; //表名	
        $new_data = array(); //返回数据数组
        //验证数据完整性
        if (empty($fc_type) || empty($qishu) || empty($kaipan) || empty($fengpan) || empty($kaijiang) || empty($id)) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据不完整';
            return $new_data;
        }
        //验证数据类型
        if (!is_int($qishu) || $qishu == 0 || !is_int($id) || $id == 0) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据类型不正确';
            return $new_data;
        }
        $time_preg = '/^([0-1]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$/';
        $kaipan_p = preg_match($time_preg, $kaipan);
        $fengpan_p = preg_match($time_preg, $fengpan);
        $kaijiang_p = preg_match($time_preg, $kaijiang);
        if ($kaipan_p == 0 || $fengpan_p == 0 || $kaijiang_p == 0) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据类型不正确';
            return $new_data;
        }

        //判断数据表是否存在
        $table_array = $connection->createCommand("show tables")->queryAll();
        $new_table_array = array();
        foreach ($table_array as $val) {
            foreach ($val as $v) {
                $new_table_array[] = $v;
            }
        }
        if (!in_array($table_name, $new_table_array)) {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据表不存在';
            return $new_data;
        }

        //处理数组
        $data = array();
        $data['qishu'] = $qishu;
        $data['kaipan'] = $kaipan;
        $data['fengpan'] = $fengpan;
        $data['kaijiang'] = $kaijiang;
        $where = array();
        $where['id'] = $id;

        //执行修改
        $row = $connection->createCommand()->update($table_name, $data, $where)->execute();

        if ($row) {
            $new_data['ErrorCode'] = 1;
            $new_data['ErrorMsg'] = '修改成功';
            return $new_data;
        } else {
            $new_data['ErrorCode'] = 2;
            $new_data['ErrorMsg'] = '数据未变动或格式不正确';
            return $new_data;
        }
    }

}
