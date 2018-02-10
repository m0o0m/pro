<template lang="html">
  <div class="six_six">
    <nav-center :menus="lists" :margin="false" @go_child="go_child"></nav-center>
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input @on-blur="onblur_top(0)" @on-focus="onfocus_top(0)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="content">
      <div class="top">
        <ul>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">勾选</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">勾选</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">勾选</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">勾选</li>
        </ul>
      </div>
      <div class="table">
        <div class="bottom" v-for="item in lists1">
          <ul v-for="key in item.object" :class="[key.flag?'table-current':'']">
            <li @click="pour(key)" :class="[key.boll_name?'one':'border1_none']">
              <!--<span>{{key.boll_name}}</span>-->
              <span :class="[key.color?key.color:'']">{{key.boll_name}}</span>
            </li>
            <li  @click="pour(key)" v-if="key.boll_name" :class="[key.boll_name?'two':'border2_none']">
              {{key.odd}}
            </li>
            <li v-else :class="[key.boll_name?'two':'border2_none']">

            </li>
            <li @click.self="pour(key)" :class="[key.boll_name?'three':'border3_none']">
               <Checkbox v-if="key.boll_name" v-model="key.flag" ref="one_click" :disabled="key.config_disabled"></Checkbox>
              <!--<input v-if="key.number" type="checkbox" @click="pour(key)" v-model="key.item">-->
              <p @click="pour(key)" class="three_nono"></p>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <!--底部控制tabs-->
    <div :class="[bottom_width_config?'bottom_width':'','bottom_config']">
      <button type="button" :ref="item.item" class="buttom" @click="config_type(item,i)" :class="{'item_class':selectItem==i}" v-for="(item,i) in bottom_list">
        {{item.name}}
      </button>
    </div>
    <div class="bottom_no" v-if="one">
      <div class="dan">
        <div class="list">
          <p>胆1</p>
          <p class="dan_list">{{dan1}}</p>
          <!--<I-Input style="width: 60px" size="small"></I-Input>-->
        </div>
        <div class="list">
          <p>胆2</p>
          <p class="dan_list">{{dan2}}</p>
          <!--<I-Input style="width: 60px" size="small"></I-Input>-->
        </div>
      </div>
    </div>
    <div class="bottom_no1" v-if="two">
      <div style="overflow: hidden;padding-left: 10px;margin-bottom: 15px">
        <div class="list" v-for="item in shengxiao1" @click="get_shengxiao1(item)">
          <Radio size="large" name="item.item" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
          <div class="select"></div>
        </div>
      </div>
      <div style="overflow: hidden;padding-left: 10px">
        <div class="list" v-for="item in shengxiao2" @click="get_shengxiao2(item)">
          <Radio size="large" name="item.item" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
          <div class="select"></div>
        </div>
      </div>
    </div>
    <div class="bottom_no1" v-if="three">
      <div style="overflow: hidden;padding-left: 80px;margin-bottom: 15px">
        <div class="list" v-for="item in weishu1" @click="get_weishu1(item)">
          <Radio size="large" name="concat1" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
          <div class="select"></div>
        </div>
      </div>
      <div style="overflow: hidden;padding-left: 80px">
        <div class="list" v-for="item in weishu2" @click="get_weishu2(item)">
          <Radio size="large" name="concat2" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
          <div class="select"></div>
        </div>
      </div>
    </div>
    <div class="bottom_no1" v-if="four">
      <div style="overflow: hidden;padding-left: 10px;margin-bottom: 15px">
        <div class="list" v-for="item in shengwei1" @click="get_shengwei1(item)">
          <Radio size="large" name="item.item" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
        </div>
      </div>
      <div style="overflow: hidden;padding-left: 80px">
        <div class="list" v-for="item in shengwei2" @click="get_shengwei2(item)">
          <Radio size="large" name="concat2" v-model="item.flag" style="margin-right: 0"></Radio>
          <span>{{item.name}}</span>
        </div>
      </div>
    </div>
    <div class="bump" v-if="five">
        <span v-for="(item,key) in bump1">
          <I-Input :maxlength='9' style="width: 60px" ref="five_input1" @on-focus="onfocus_five(1,key,item)" class="mr" size="small" v-model="item.input_name" @on-blur="five_out(item,key,1)" @on-keyup="five_verify(item)" @on-afterpaste="five_verify(item)"></I-Input>
        </span>
      <div class="center">碰</div>
        <span v-for="(val,key) in bump2">
          <I-Input :maxlength='9' style="width: 60px" ref="five_input2" @on-focus="onfocus_five(2,key,val)" class="mr" size="small" v-model="val.input_name" @on-blur="five_out(val,key,2)" @on-keyup="five_verify(val)" @on-afterpaste="five_verify(val)"></I-Input>
        </span>
    </div>
    <six_modal :modal='modal' @cancel="cancel"></six_modal>
  </div>
</template>

<script>
function get_lists1() {
  return [{
      name: '1',
      object: [{
          name: '1',
          index: 0,
          boll_name: '1',
          color: 'red',
          number: '48.8',
          flag: false
        },
        {
          index: 1,
          boll_name: '2',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 2,
          boll_name: '3',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 3,
          boll_name: '4',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 4,
          boll_name: '5',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 5,
          boll_name: '6',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 6,
          boll_name: '7',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 7,
          boll_name: '8',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 8,
          boll_name: '9',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 9,
          boll_name: '10',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 10,
          boll_name: '11',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 11,
          boll_name: '12',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 12,
          boll_name: '13',
          number: '48.8',
          color: 'red',
          flag: false
        },
      ]
    },
    {
      name: '2',
      object: [{
          index: 13,
          boll_name: '14',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 14,
          boll_name: '15',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 15,
          boll_name: '16',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 16,
          boll_name: '17',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 17,
          boll_name: '18',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 18,
          boll_name: '19',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 19,
          boll_name: '20',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 20,
          boll_name: '21',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 21,
          boll_name: '22',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 22,
          boll_name: '23',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 23,
          boll_name: '24',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 24,
          boll_name: '25',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 25,
          boll_name: '26',
          number: '48.8',
          color: 'blue',
          flag: false
        },
      ]
    },
    {
      name: '2',
      object: [{
          index: 26,
          boll_name: '27',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 27,
          boll_name: '28',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 28,
          boll_name: '29',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 29,
          boll_name: '30',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 30,
          boll_name: '31',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 31,
          boll_name: '32',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 32,
          boll_name: '33',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 33,
          boll_name: '34',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 34,
          boll_name: '35',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 35,
          boll_name: '36',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 36,
          boll_name: '37',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 37,
          boll_name: '38',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 38,
          boll_name: '39',
          number: '48.8',
          color: 'green',
          flag: false
        },
      ]
    },
    {
      name: '2',
      object: [{
          index: 39,
          boll_name: '40',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 40,
          boll_name: '41',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 41,
          boll_name: '42',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 42,
          boll_name: '43',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 43,
          boll_name: '44',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: 44,
          boll_name: '45',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 45,
          boll_name: '46',
          number: '48.8',
          color: 'red',
          flag: false
        },
        {
          index: 46,
          boll_name: '47',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 47,
          boll_name: '48',
          number: '48.8',
          color: 'blue',
          flag: false
        },
        {
          index: 48,
          boll_name: '49',
          number: '48.8',
          color: 'green',
          flag: false
        },
        {
          index: '',
          boll_name: '',
          number: '',
          flag: ''
        },
        {
          index: '',
          boll_name: '',
          number: '',
          flag: ''
        },
        {
          index: '',
          boll_name: '',
          number: '',
          flag: ''
        },
      ]
    },
  ]
}
function get_bottom1() {
  return [{
      name: '正常',
      item: 1
    },
    {
      name: '胆拖',
      item: 2
    },
    {
      name: '生肖对碰',
      item: 3
    },
    {
      name: '尾数对碰',
      item: 4
    },
    {
      name: '生尾对碰',
      item: 5
    },
    {
      name: '任意对碰',
      item: 6
    },
  ]
}
function get_bottom2() {
  return [
    {
      name: '正常',
      item: 1
    },
    {
      name: '胆拖',
      item: 2
    }
  ]
}
import api from '../../../api/config'
import split from '../../../assets/js/delete_object'
import navCenter from './components/nav'
import six_modal from '../../../share_components/bet_six'
import {Modal,Input,Radio,Checkbox} from 'iview';
import share from '../../../share_components/share'
import hint from '../../../share_components/hint_msg'
export default {
  components: {
    navCenter,
    six_modal,
    Modal,'I-Input':Input,Radio,Checkbox
  },
  props: {},
  data() {
    return {
      a:'',
      now_setItem: 0,//当前处于顶部控制栏的哪个
      ok: false,//是否处于重复号码状态
      bump_arr: [], //任意组合数组
      bump1: [
        {
          name: 'bump1',
          index: 0,
          input_name: ''
        },
        {
          name: 'bump1',
          index: 1,
          input_name: ''
        },
        {
          name: 'bump1',
          index: 2,
          input_name: ''
        },
        {
          name: 'bump1',
          index: 3,
          input_name: ''
        },
        {
          name: 'bump1',
          index: 4,
          input_name: ''
        },
      ], //任意碰1
      bump2: [
        {
          name: 'bump2',
          index: 5,
          input_name: ''
        },
        {
          name: 'bump2',
          index: 6,
          input_name: ''
        },
        {
          name: 'bump2',
          index: 7,
          input_name: ''
        },
        {
          name: 'bump2',
          index: 8,
          input_name: ''
        },
        {
          name: 'bump2',
          index: 9,
          input_name: ''
        },
      ], //任意碰2
      bottom_width_config: false, //底部控制的宽度
      top_config: 0, //顶部tab识别码
      dan1: null, //胆1
      dan2: null, //胆2
      arr: [], //点击选择统计下来的数组 胆拖
      config_shengxiao1: [], //生肖碰1
      config_shengxiao2: [], //生肖碰2
      config_weishu1: [], //尾数碰2
      config_weishu2: [], //尾数碰2
      selectItem: 0,
      back_data: [],
      money: '',
      modal: false,
      one: false,
      two: false,
      three: false,
      four: false,
      five: false,
      shengxiao_data: [], //接口获取出来的生肖对应的号码
      weishu_data: {
        zero: [10, 20, 30, 40],
        one: [1, 11, 21, 31, 41],
        two: [2, 12, 22, 32, 42],
        three: [3, 13, 23, 33, 43],
        four: [4, 14, 24, 34, 44],
        five: [5, 15, 25, 35, 45],
        six: [6, 16, 26, 36, 46],
        seven: [7, 17, 27, 37, 47],
        eight: [8, 18, 28, 38, 48],
        night: [9, 19, 29, 39, 49],
      }, //本地数据的尾数对应的号码
      shengxiao1: [
        {
          name: '鼠',
          item: 0,
          index: 0,
          flag: false
        },
        {
          name: '牛',
          item: 0,
          index: 1,
          flag: false
        },
        {
          name: '虎',
          item: 0,
          index: 2,
          flag: false
        },
        {
          name: '兔',
          item: 0,
          index: 3,
          flag: false
        },
        {
          name: '龙',
          item: 0,
          index: 4,
          flag: false
        },
        {
          name: '蛇',
          item: 0,
          index: 5,
          flag: false
        },
        {
          name: '马',
          item: 0,
          index: 6,
          flag: false
        },
        {
          name: '羊',
          item: 0,
          index: 7,
          flag: false
        },
        {
          name: '猴',
          item: 0,
          index: 8,
          flag: false
        },
        {
          name: '鸡',
          item: 0,
          index: 9,
          flag: false
        },
        {
          name: '狗',
          item: 0,
          index: 10,
          flag: false
        },
        {
          name: '猪',
          item: 0,
          index: 11,
          flag: false
        },
      ],
      shengxiao2: [
        {
          name: '鼠',
          item: 1,
          index: 12,
          flag: false
        },
        {
          name: '牛',
          item: 1,
          index: 13,
          flag: false
        },
        {
          name: '虎',
          item: 1,
          index: 14,
          flag: false
        },
        {
          name: '兔',
          item: 1,
          index: 15,
          flag: false
        },
        {
          name: '龙',
          item: 1,
          index: 16,
          flag: false
        },
        {
          name: '蛇',
          item: 1,
          index: 17,
          flag: false
        },
        {
          name: '马',
          item: 1,
          index: 18,
          flag: false
        },
        {
          name: '羊',
          item: 1,
          index: 19,
          flag: false
        },
        {
          name: '猴',
          item: 1,
          index: 20,
          flag: false
        },
        {
          name: '鸡',
          item: 1,
          index: 21,
          flag: false
        },
        {
          name: '狗',
          item: 1,
          index: 22,
          flag: false
        },
        {
          name: '猪',
          item: 1,
          index: 23,
          flag: false
        },
      ],
      weishu1: [
        {
          name: '0尾',
          item: 0,
          index: 0,
          flag: false
        },
        {
          name: '1尾',
          item: 0,
          index: 1,
          flag: false
        },
        {
          name: '2尾',
          item: 0,
          index: 2,
          flag: false
        },
        {
          name: '3尾',
          item: 0,
          index: 3,
          flag: false
        },
        {
          name: '4尾',
          item: 0,
          index: 4,
          flag: false
        },
        {
          name: '5尾',
          item: 0,
          index: 5,
          flag: false
        },
        {
          name: '6尾',
          item: 0,
          index: 6,
          flag: false
        },
        {
          name: '7尾',
          item: 0,
          index: 7,
          flag: false
        },
        {
          name: '8尾',
          item: 0,
          index: 8,
          flag: false
        },
        {
          name: '9尾',
          item: 0,
          index: 9,
          flag: false
        },
      ],
      weishu2: [
        {
          name: '0尾',
          item: 1,
          index: 0,
          flag: false
        },
        {
          name: '1尾',
          item: 1,
          index: 1,
          flag: false
        },
        {
          name: '2尾',
          item: 1,
          index: 2,
          flag: false
        },
        {
          name: '3尾',
          item: 1,
          index: 3,
          flag: false
        },
        {
          name: '4尾',
          item: 1,
          index: 4,
          flag: false
        },
        {
          name: '5尾',
          item: 1,
          index: 5,
          flag: false
        },
        {
          name: '6尾',
          item: 1,
          index: 6,
          flag: false
        },
        {
          name: '7尾',
          item: 1,
          index: 7,
          flag: false
        },
        {
          name: '8尾',
          item: 1,
          index: 8,
          flag: false
        },
        {
          name: '9尾',
          item: 1,
          index: 9,
          flag: false
        },
      ],
      shengwei1: [
        {
          name: '鼠',
          item: 0,
          index: 0,
          flag: false
        },
        {
          name: '牛',
          item: 0,
          index: 1,
          flag: false
        },
        {
          name: '虎',
          item: 0,
          index: 2,
          flag: false
        },
        {
          name: '兔',
          item: 0,
          index: 3,
          flag: false
        },
        {
          name: '龙',
          item: 0,
          index: 4,
          flag: false
        },
        {
          name: '蛇',
          item: 0,
          index: 5,
          flag: false
        },
        {
          name: '马',
          item: 0,
          index: 6,
          flag: false
        },
        {
          name: '羊',
          item: 0,
          index: 7,
          flag: false
        },
        {
          name: '猴',
          item: 0,
          index: 8,
          flag: false
        },
        {
          name: '鸡',
          item: 0,
          index: 9,
          flag: false
        },
        {
          name: '狗',
          item: 0,
          index: 10,
          flag: false
        },
        {
          name: '猪',
          item: 0,
          index: 11,
          flag: false
        },
      ],
      shengwei2: [
        {
          name: '0尾',
          item: 0,
          index: 0,
          flag: false
        },
        {
          name: '1尾',
          item: 0,
          index: 1,
          flag: false
        },
        {
          name: '2尾',
          item: 0,
          index: 2,
          flag: false
        },
        {
          name: '3尾',
          item: 0,
          index: 3,
          flag: false
        },
        {
          name: '4尾',
          item: 0,
          index: 4,
          flag: false
        },
        {
          name: '5尾',
          item: 0,
          index: 5,
          flag: false
        },
        {
          name: '6尾',
          item: 0,
          index: 6,
          flag: false
        },
        {
          name: '7尾',
          item: 0,
          index: 7,
          flag: false
        },
        {
          name: '8尾',
          item: 0,
          index: 8,
          flag: false
        },
        {
          name: '9尾',
          item: 0,
          index: 9,
          flag: false
        },
      ],
      bottom_list: [
        {
          name: '正常',
          item: 1
        },
        {
          name: '胆拖',
          item: 2
        },
        {
          name: '生肖对碰',
          item: 3
        },
        {
          name: '尾数对碰',
          item: 4
        },
        {
          name: '生尾对碰',
          item: 5
        },
        {
          name: '任意对碰',
          item: 6
        },
      ],
      bottom_list1: [
        {
          name: '正常',
          item: 1
        },
        {
          name: '胆拖',
          item: 2
        },
      ],
      lists: [
        {
          name: '二全中',
          item: 0
        },
        {
          name: '二中特',
          item: 1
        },
        {
          name: '特串',
          item: 2
        },
        {
          name: '三全中',
          item: 3
        },
        {
          name: '三中二',
          item: 4
        },
        {
          name: '四全中',
          item: 5
        }
      ],
      lists1: [
          {
          name: '1',
          object: [{
              config_disabled: false,
              index: 0,
              boll_name: '1',
              color: 'red',
              number: '48.8',
              flag: false
            },
            {
              config_disabled: false,
              index: 1,
              boll_name: '2',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 2,
              boll_name: '3',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 3,
              boll_name: '4',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 4,
              boll_name: '5',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 5,
              boll_name: '6',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 6,
              boll_name: '7',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 7,
              boll_name: '8',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 8,
              boll_name: '9',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 9,
              boll_name: '10',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 10,
              boll_name: '11',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 11,
              boll_name: '12',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 12,
              boll_name: '13',
              number: '48.8',
              color: 'red',
              flag: false
            },
          ]
        },
        {
          name: '2',
          object: [{
              config_disabled: false,
              index: 13,
              boll_name: '14',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 14,
              boll_name: '15',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 15,
              boll_name: '16',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 16,
              boll_name: '17',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 17,
              boll_name: '18',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 18,
              boll_name: '19',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 19,
              boll_name: '20',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 20,
              boll_name: '21',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 21,
              boll_name: '22',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 22,
              boll_name: '23',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 23,
              boll_name: '24',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 24,
              boll_name: '25',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 25,
              boll_name: '26',
              number: '48.8',
              color: 'blue',
              flag: false
            },
          ]
        },
        {
          name: '2',
          object: [{
              config_disabled: false,
              index: 26,
              boll_name: '27',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 27,
              boll_name: '28',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 28,
              boll_name: '29',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 29,
              boll_name: '30',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 30,
              boll_name: '31',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 31,
              boll_name: '32',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 32,
              boll_name: '33',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 33,
              boll_name: '34',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 34,
              boll_name: '35',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 35,
              boll_name: '36',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 36,
              boll_name: '37',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 37,
              boll_name: '38',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 38,
              boll_name: '39',
              number: '48.8',
              color: 'green',
              flag: false
            },
          ]
        },
        {
          name: '2',
          object: [{
              config_disabled: false,
              index: 39,
              boll_name: '40',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 40,
              boll_name: '41',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 41,
              boll_name: '42',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 42,
              boll_name: '43',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 43,
              boll_name: '44',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              config_disabled: false,
              index: 44,
              boll_name: '45',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 45,
              boll_name: '46',
              number: '48.8',
              color: 'red',
              flag: false
            },
            {
              config_disabled: false,
              index: 46,
              boll_name: '47',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 47,
              boll_name: '48',
              number: '48.8',
              color: 'blue',
              flag: false
            },
            {
              config_disabled: false,
              index: 48,
              boll_name: '49',
              number: '48.8',
              color: 'green',
              flag: false
            },
            {
              index: '',
              boll_name: '',
              number: '',
              flag: ''
            },
            {
              index: '',
              boll_name: '',
              number: '',
              flag: ''
            },
            {
              index: '',
              boll_name: '',
              number: '',
              flag: ''
            },
          ]
        },
      ]
    }
  },
  created() {
    this.fetchData();
    this.$root.$on('success', (e) => {
      if (e) {
        this.modal = false;
        this.reset()
      }
    });
    this.$root.$on('this_money',(e)=>{
        this.money = e
    });
  },
  mounted() {
    this.$root.$emit('no_top',false);
    this.$root.$emit('child_change', 0);
    this.$root.$on('time_out',(e)=>{
      if(e){
        this.fetchData(2)
      }
    })
  },
  destroyed(){
    this.$root.$off('time_out')
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route': 'fetchData', // 只有这个页面初始化之后，这个监听事件才开始生效
    a: function (new_val,old_val) {
        if(new_val != old_val){
            this.$root.$emit('clear_key_number','')
        }
    },
//    money: function(new_val,old_val){
//        if(new_val != old_val){
//            this.computed_money()
//        }
//    }
  },
  methods: {
//    onblur_five: function(type,key){
//        if(type == 1){
//            key.input_name = this.$refs.five_input1[key].$refs.input.value;
//        }else if(type == 2){
//            key.input_name = this.$refs.five_input2[key].$refs.input.value;
//        }
//    },
    onfocus_five: function(type,item,key){
//        this.a = 99;
        if(type == 1){
            console.log('当前索引：'+item);
            console.log(this.$refs.five_input1);
            this.$refs.five_input1[item].$refs.input.data_onoff = 'true';
            this.a = item+51;
            let dom = document.querySelectorAll('input');
            for(let i = 0;i < dom.length;i++){
                if(i != item+51) {
                    dom[i].data_onoff = 'false';
                }
            }
        }else{
            this.$refs.five_input2[item].$refs.input.data_onoff = 'true';
            this.a = item+56;
            let dom = document.querySelectorAll('input');
            for(let i = 0;i < dom.length;i++){
                if(i != item+56) {
                    dom[i].data_onoff = 'false';
                }
            }
        }
        if(type == 1){
          console.log('输入的数字是什么？：'+key.input_name);
          let dom = document.querySelectorAll('input');
          for(let i = 51;i < dom.length;i++) {
            if (dom[i].value > 49 || dom[i].value == 0) {
              dom[i].value = '';
            }
          }
          key.input_name = this.$refs.five_input1[item].$refs.input.value;
          if (key.input_name > 49 || key.input_name == 0) {
            key.input_name = '';
          }
        }else{
          let dom = document.querySelectorAll('input');
          for(let i = 51;i < dom.length;i++) {
            if (dom[i].value > 49 || dom[i].value == 0) {
              dom[i].value = '';
            }
          }
          key.input_name = this.$refs.five_input2[item].$refs.input.value;
          if (key.input_name > 49 || key.input_name == 0) {
            key.input_name = '';
          }
        }
        let dom = document.querySelectorAll('input');
//        var arr = [9, 9, 111, 2, 3, 4, 4, 5, 7];
//        console.log(dom);
//        let dom_value = [];
//        for(let i = 51;i < dom.length-1;i++){
//            if(dom[i].value != ''){
//              dom_value.push(dom[i].value)
//            }
//        }
//        let sortedArr = dom_value.sort();
//        for (let i = 0; i < dom_value.length - 1; i++) {
//          if (sortedArr[i + 1] == sortedArr[i]) {
//            sortedArr = sortedArr.splice(i,1);
//            console.log('重复了')
//          }
//        }
//        console.log(dom_value);
//        let newArr = [];
//        for (let k = 51;k < dom.length;k++) {
//            console.log('myfriend????????'+newArr.indexOf(dom[k].value));
//          if (newArr.indexOf(dom[k].value) == -1) {
//            newArr.push(dom[k].value);
//            dom[k].value = '';
//            console.log('重复了');
////            alert('您上次输入的数值重复了')
//          }
//        }
        let hash = {};
        for(let i = 51;i < dom.length;i++) {
          if(dom[i].value != ''){
            if(hash[dom[i].value]){
              dom[i].value = '';
              key.input_name = '';
              console.log(key.input_name);
              console.log('重复了');
              this.$Modal.warning({
                content: '抱歉不允许输入重复号码！'
              });
              window.setTimeout(() => {
                this.$Modal.remove()
              }, share.Prompt);
            }
          }
          console.log('wowowoowoowowo!:'+dom[i].value);
          hash[dom[i].value] = true;
        }
      console.log('wowowoowoowowo11111!:'+key.input_name);
    },
    onfocus_top: function(index){
        let dom = document.querySelectorAll('input');
        console.log('是在最后一个tab：？'+this.five);
        if(this.five){
            if(index == 0){
                index = 0
            }else{
                index = dom.length - 11
            }
        }else if(this.four || this.three || this.two){
          if(index == 0){
            index = 0
          }else{
            index = 50;
          }
        }else{
          if(index == 0){
              index = 0
          }else{
              index = dom.length-1;
          }
        }
        console.log('索引：'+ index);
        this.a = index;
        for(let i = 0;i < dom.length;i++){
            if(i != index) {
                dom[i].data_onoff = 'false';
            }else{
                dom[i].data_onoff = 'true'
            }
        }
    },
    onblur_top: function(index){
        let dom = document.querySelectorAll('input');
        if(this.five){
          if(index == 0){
            index = 0
          }else{
            index = dom.length - 11
          }
        }else if(this.four || this.three || this.two){
          if(index == 0){
            index = 0
          }else{
            index = 50;
          }
        }else{
          if(index == 0){
            index = 0
          }else{
            index = dom.length-1;
          }
        }
        if(dom[index].value != 'on'){
          this.money = dom[index].value;
        }
    },
    get_shengwei1: function(item) {
      for (let i = 0; i < this.shengwei1.length; i++) {
        this.shengwei1[i].flag = false
      }
      item.flag = true;
      switch (item.name) {
        case '鼠':
          this.config_shengxiao1 = this.shengxiao_data.mouse;
          break;
        case '牛':
          this.config_shengxiao1 = this.shengxiao_data.cattle;
          break;
        case '虎':
          this.config_shengxiao1 = this.shengxiao_data.tiger;
          break;
        case '兔':
          this.config_shengxiao1 = this.shengxiao_data.rabbit;
          break;
        case '龙':
          this.config_shengxiao1 = this.shengxiao_data.dragon;
          break;
        case '蛇':
          this.config_shengxiao1 = this.shengxiao_data.snake;
          break;
        case '马':
          this.config_shengxiao1 = this.shengxiao_data.horse;
          break;
        case '羊':
          this.config_shengxiao1 = this.shengxiao_data.sheep;
          break;
        case '猴':
          this.config_shengxiao1 = this.shengxiao_data.monkey;
          break;
        case '鸡':
          this.config_shengxiao1 = this.shengxiao_data.chicken;
          break;
        case '狗':
          this.config_shengxiao1 = this.shengxiao_data.dog;
          break;
        case '猪':
          this.config_shengxiao1 = this.shengxiao_data.pig;
          break;
      }
      console.log(this.config_shengxiao1);
    },
    get_shengwei2: function(item) {
      for (let i = 0; i < this.shengwei2.length; i++) {
        this.shengwei2[i].flag = false
      }
      item.flag = true;
      switch (item.name) {
        case '0尾':
          this.config_weishu1 = this.weishu_data.zero;
          break;
        case '1尾':
          this.config_weishu1 = this.weishu_data.one;
          break;
        case '2尾':
          this.config_weishu1 = this.weishu_data.two;
          break;
        case '3尾':
          this.config_weishu1 = this.weishu_data.three;
          break;
        case '4尾':
          this.config_weishu1 = this.weishu_data.four;
          break;
        case '5尾':
          this.config_weishu1 = this.weishu_data.five;
          break;
        case '6尾':
          this.config_weishu1 = this.weishu_data.six;
          break;
        case '7尾':
          this.config_weishu1 = this.weishu_data.seven;
          break;
        case '8尾':
          this.config_weishu1 = this.weishu_data.eight;
          break;
        case '9尾':
          this.config_weishu1 = this.weishu_data.night;
          break;
      }
      console.log(this.config_weishu1);
    },
    get_shengxiao1: function(item) {
      console.log(item);
      for (let i = 0; i < this.shengxiao1.length; i++) {
        this.shengxiao1[i].flag = false;
      }
      for (let i = 0; i < this.shengxiao2.length; i++) {
        if (item.name === this.shengxiao2[i].name && this.shengxiao2[i].flag) {
          item.flag = false;
          this.$Modal.warning({
            content: '请重新选择两个不一样的！',
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt);
          this.config_shengxiao1 = [];
          break
        } else {
          item.flag = true;
          switch (item.name) {
            case '鼠':
              this.config_shengxiao1 = this.shengxiao_data.mouse;
              break;
            case '牛':
              this.config_shengxiao1 = this.shengxiao_data.cattle;
              break;
            case '虎':
              this.config_shengxiao1 = this.shengxiao_data.tiger;
              break;
            case '兔':
              this.config_shengxiao1 = this.shengxiao_data.rabbit;
              break;
            case '龙':
              this.config_shengxiao1 = this.shengxiao_data.dragon;
              break;
            case '蛇':
              this.config_shengxiao1 = this.shengxiao_data.snake;
              break;
            case '马':
              this.config_shengxiao1 = this.shengxiao_data.horse;
              break;
            case '羊':
              this.config_shengxiao1 = this.shengxiao_data.sheep;
              break;
            case '猴':
              this.config_shengxiao1 = this.shengxiao_data.monkey;
              break;
            case '鸡':
              this.config_shengxiao1 = this.shengxiao_data.chicken;
              break;
            case '狗':
              this.config_shengxiao1 = this.shengxiao_data.dog;
              break;
            case '猪':
              this.config_shengxiao1 = this.shengxiao_data.pig;
              break;
          }
          console.log(this.config_shengxiao1);
        }
      }
    },
    get_shengxiao2: function(item) {
      for (let i = 0; i < this.shengxiao2.length; i++) {
        this.shengxiao2[i].flag = false;
      }
      for (let i = 0; i < this.shengxiao1.length; i++) {
        if (item.name === this.shengxiao1[i].name && this.shengxiao1[i].flag) {
          item.flag = false;
          this.$Modal.warning({
            content: '请重新选择两个不一样的！',
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt);
          this.config_shengxiao2 = [];
          break
        } else {
          item.flag = true;
          switch (item.name) {
            case '鼠':
              this.config_shengxiao2 = this.shengxiao_data.mouse;
              break;
            case '牛':
              this.config_shengxiao2 = this.shengxiao_data.cattle;
              break;
            case '虎':
              this.config_shengxiao2 = this.shengxiao_data.tiger;
              break;
            case '兔':
              this.config_shengxiao2 = this.shengxiao_data.rabbit;
              break;
            case '龙':
              this.config_shengxiao2 = this.shengxiao_data.dragon;
              break;
            case '蛇':
              this.config_shengxiao2 = this.shengxiao_data.snake;
              break;
            case '马':
              this.config_shengxiao2 = this.shengxiao_data.horse;
              break;
            case '羊':
              this.config_shengxiao2 = this.shengxiao_data.sheep;
              break;
            case '猴':
              this.config_shengxiao2 = this.shengxiao_data.monkey;
              break;
            case '鸡':
              this.config_shengxiao2 = this.shengxiao_data.chicken;
              break;
            case '狗':
              this.config_shengxiao2 = this.shengxiao_data.dog;
              break;
            case '猪':
              this.config_shengxiao2 = this.shengxiao_data.pig;
              break;
          }
          console.log(this.config_shengxiao2);
        }
      }
    },
    get_weishu1: function(item) {
      for (let i = 0; i < this.weishu1.length; i++) {
        this.weishu1[i].flag = false;
      }
      for (let i = 0; i < this.weishu2.length; i++) {
        if (item.name === this.weishu2[i].name && this.weishu2[i].flag) {
          item.flag = false;
          this.$Modal.warning({
            content: '请重新选择两个不一样的！',
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt);
          this.config_weishu1 = [];
          break
        } else {
          item.flag = true;
          switch (item.name) {
            case '0尾':
              this.config_weishu1 = this.weishu_data.zero;
              break;
            case '1尾':
              this.config_weishu1 = this.weishu_data.one;
              break;
            case '2尾':
              this.config_weishu1 = this.weishu_data.two;
              break;
            case '3尾':
              this.config_weishu1 = this.weishu_data.three;
              break;
            case '4尾':
              this.config_weishu1 = this.weishu_data.four;
              break;
            case '5尾':
              this.config_weishu1 = this.weishu_data.five;
              break;
            case '6尾':
              this.config_weishu1 = this.weishu_data.six;
              break;
            case '7尾':
              this.config_weishu1 = this.weishu_data.seven;
              break;
            case '8尾':
              this.config_weishu1 = this.weishu_data.eight;
              break;
            case '9尾':
              this.config_weishu1 = this.weishu_data.night;
              break;
          }
          console.log(this.config_weishu1);
        }
      }
    },
    get_weishu2: function(item) {
      for (let i = 0; i < this.weishu2.length; i++) {
        this.weishu2[i].flag = false;
      }
      for (let i = 0; i < this.weishu1.length; i++) {
        if (item.name === this.weishu1[i].name && this.weishu1[i].flag) {
          item.flag = false;
          this.$Modal.warning({
            content: '请重新选择两个不一样的！',
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt);
          this.config_weishu2 = [];
          break
        } else {
          item.flag = true;
          switch (item.name) {
            case '0尾':
              this.config_weishu2 = this.weishu_data.zero;
              break;
            case '1尾':
              this.config_weishu2 = this.weishu_data.one;
              break;
            case '2尾':
              this.config_weishu2 = this.weishu_data.two;
              break;
            case '3尾':
              this.config_weishu2 = this.weishu_data.three;
              break;
            case '4尾':
              this.config_weishu2 = this.weishu_data.four;
              break;
            case '5尾':
              this.config_weishu2 = this.weishu_data.five;
              break;
            case '6尾':
              this.config_weishu2 = this.weishu_data.six;
              break;
            case '7尾':
              this.config_weishu2 = this.weishu_data.seven;
              break;
            case '8尾':
              this.config_weishu2 = this.weishu_data.eight;
              break;
            case '9尾':
              this.config_weishu2 = this.weishu_data.night;
              break;
          }
          console.log(this.config_weishu2);
        }
      }
    },
    //切换底部控制tab
    config_type: function(item, i) {
      this.reset_bottom();
      this.selectItem = i;
      console.log(item.name);
      let dom = item.item;
      this.$refs[dom][0].setAttribute('disabled',"true");
      if(!this.bottom_width_config){
        for(let j=1;j<7;j++){
          if(this.$refs[j][0].hasAttribute('disabled')){
            if(j != dom){
              this.$refs[j][0].removeAttribute('disabled')
            }
          }
        }
      }else if(this.bottom_width_config){
        for(let j=1;j<3;j++){
          if(this.$refs[j][0].hasAttribute('disabled')){
            if(j != dom){
              this.$refs[j][0].removeAttribute('disabled')
            }
          }
        }
      }
      switch (item.item) {
        case 1:
          this.one = false;
          this.two = false;
          this.three = false;
          this.four = false;
          this.five = false;
          this.$root.$emit('no_top',false,0);
          break;
        case 2:
          this.one = true;
          this.two = false;
          this.three = false;
          this.four = false;
          this.five = false;
          this.$root.$emit('no_top',false,1);
          break;
        case 3:
          this.close_checked();
          this.one = false;
          this.two = true;
          this.three = false;
          this.four = false;
          this.five = false;
          this.$root.$emit('no_top',false,2);
          break;
        case 4:
          this.close_checked();
          this.one = false;
          this.two = false;
          this.three = true;
          this.four = false;
          this.five = false;
          this.$root.$emit('no_top',false,3);
          break;
        case 5:
          this.close_checked();
          this.one = false;
          this.two = false;
          this.three = false;
          this.four = true;
          this.five = false;
          this.$root.$emit('no_top',false,4);
          break;
        case 6:
          this.close_checked();
          this.one = false;
          this.two = false;
          this.three = false;
          this.four = false;
          this.five = true;
          this.$root.$emit('no_top',false,5);
          break;
      }
    },
    //禁止勾选上方的球号
    close_checked: function() {
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          this.lists1[i].object[j].config_disabled = true;
        }
      }
    },
    //头部控制栏tab
    go_child: function(e) {
      console.log(e);
      for(let j=1;j<7;j++){
        if(this.$refs[j] != ''){
          this.$refs[j][0].removeAttribute('disabled')
        }
      }
      if (e >= 3) {
        this.bottom_width_config = true;
        this.bottom_list = get_bottom2();
        this.bottom_list = this.bottom_list1;
        //清除disable
        console.log(this.$refs);
      }else {
        this.bottom_width_config = false;
        this.bottom_list = get_bottom1();
        this.bottom_list = this.bottom_list;
        //清除disable
      }
      this.top_config = e;
      this.reset_bottom();
      this.computed(this.back_data, e);
    },
    //重置
    reset: function() {
      for (let i = 0; i < this.shengxiao1.length; i++) {
        this.shengxiao1[i].flag = false
      }
      for (let i = 0; i < this.shengxiao2.length; i++) {
        this.shengxiao2[i].flag = false
      }
      for (let i = 0; i < this.weishu1.length; i++) {
        this.weishu1[i].flag = false
      }
      for (let i = 0; i < this.weishu2.length; i++) {
        this.weishu2[i].flag = false
      }
      for (let i = 0; i < this.shengwei1.length; i++) {
        this.shengwei1[i].flag = false
      }
      for (let i = 0; i < this.shengwei2.length; i++) {
        this.shengwei2[i].flag = false
      }
      this.dan1 = null; //胆1
      this.dan2 = null; //胆2
      this.arr = [];
      this.bump_arr = [];
      for (let i in this.bump1) {
        this.bump1[i].input_name = ''
      }
      for (let i in this.bump2) {
        this.bump2[i].input_name = ''
      }
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          this.lists1[i].object[j].money = '';
          this.lists1[i].object[j].flag = false;
          if(this.two || this.three || this.four || this.five) {
            this.lists1[i].object[j].config_disabled = true;
          }else{
            this.lists1[i].object[j].config_disabled = false;
          }
        }
      }
      this.money = '';
      this.$root.$emit('reset', '');
      let dom1 = document.querySelectorAll('input');
      for(let i = 0;i < dom1.length;i++){
          dom1[i].value = '';
          dom1[i].data_onoff = 'false';
      }
      this.$root.$emit('clear_key_number','')
    },
    reset_bottom: function() {
      for (let i = 0; i < this.shengxiao1.length; i++) {
        this.shengxiao1[i].flag = false
      }
      for (let i = 0; i < this.shengxiao2.length; i++) {
        this.shengxiao2[i].flag = false
      }
      for (let i = 0; i < this.weishu1.length; i++) {
        this.weishu1[i].flag = false
      }
      for (let i = 0; i < this.weishu2.length; i++) {
        this.weishu2[i].flag = false
      }
      for (let i = 0; i < this.shengwei1.length; i++) {
        this.shengwei1[i].flag = false
      }
      for (let i = 0; i < this.shengwei2.length; i++) {
        this.shengwei2[i].flag = false
      }
      this.dan1 = null; //胆1
      this.dan2 = null; //胆2
      this.selectItem = 0;
      this.one = false;
      this.two = false;
      this.three = false;
      this.four = false;
      this.five = false;
      this.arr = [];
      this.bump_arr = [];
      for (let i in this.bump1) {
        this.bump1[i].input_name = ''
      }
      for (let i in this.bump2) {
        this.bump2[i].input_name = ''
      }
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          this.lists1[i].object[j].money = '';
          this.lists1[i].object[j].flag = false;
          if(this.two || this.three || this.four || this.five) {
            this.lists1[i].object[j].config_disabled = true;
          }else{
            this.lists1[i].object[j].config_disabled = false;
          }
          this.money = '';
          this.$root.$emit('reset', '');
          let dom1 = document.querySelectorAll('input');
          for(let i = 0;i < dom1.length;i++){
            dom1[i].value = '';
            dom1[i].data_onoff = 'false';
          }
          this.$root.$emit('clear_key_number','')
        }
      }
    },
    //disable状态重置
    reset_disable: function(){
      let dom = 0;
      this.$refs[dom][0].setAttribute('disabled',"true");
      if(!this.bottom_width_config){
        for(let j=1;j<7;j++){
          if(this.$refs[j][0].hasAttribute('disabled')){
            if(j != dom){
              this.$refs[j][0].removeAttribute('disabled')
            }
          }
        }
      }else if(this.bottom_width_config){
        for(let j=1;j<3;j++){
          if(this.$refs[j][0].hasAttribute('disabled')){
            if(j != dom){
              this.$refs[j][0].removeAttribute('disabled')
            }
          }
        }
      }
    },
    sortNumber: function(a, b) {
      return a.sort - b.sort
    },
    //点击选择球
    pour: function(item) {
      console.log(item);
      if (item.config_disabled) {

      } else {
        item.flag = !item.flag;
      }
      if (item.flag) {
        item.money = this.money;
        this.arr.push(item);
        this.arr = this.unique(this.arr);
        if (this.one && this.top_config < 3) {
          this.dan1 = this.arr[0].boll_name;
          this.arr[0].config_disabled = true;
          this.arr[0].flag = true
        } else if (this.one && this.top_config >= 3) {
          this.arr[0].config_disabled = true;
          this.arr[0].flag = true;
          this.dan1 = this.arr[0].boll_name;
          if (this.arr.length >= 2) {
            this.arr[1].config_disabled = true;
            this.dan2 = this.arr[1].boll_name;
            this.arr[1].flag = true
          }
        }
        //验证
        if (this.arr.length > 5 && this.arr[0].mingxi == 'second_full') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择2-5个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        } else if (this.arr.length > 5 && this.arr[0].mingxi == '2/2:30,2/1:50') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择2-5个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        } else if (this.arr.length > 5 && this.arr[0].mingxi == 'Special_series') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择2-5个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        } else if (this.arr.length > 6 && this.arr[0].mingxi == 'third_full') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择3-6个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        } else if (this.arr.length > 5 && this.arr[0].mingxi == '3/3:10,3/2:20') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择3-6个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        } else if (this.arr.length > 7 && this.arr[0].mingxi == 'fourth_full') {
          this.arr.pop();
          item.flag = false;
          this.$Modal.warning({
            content: '只能选择4-7个号码'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }
      } else {
        split.splice_arr(this.arr, item);
        console.log(this.arr);
        item.money = '';
      }
    },
    //点击下注
    go_to: function() {
        //键盘输入的值如下
      let dom = document.querySelectorAll('input');
      let hash = {};
      for(let i = 51;i<dom.length;i++){
        if(this.five && dom[i].value != ''){
          //处理重复号码
          if(hash[dom[i].value]){
            dom[i].value = '';
            console.log('重复了');
            this.$Modal.warning({
              content: '抱歉不允许输入重复号码！'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt);
            for(let j = 0;j<this.bump1.length;j++){
              this.bump1[j].input_name = dom[51+j].value
            }
            for(let k = 0;k<this.bump2.length;k++){
              this.bump2[k].input_name = dom[56+k].value
            }
            return false;
          }
          hash[dom[i].value] = true;
          //添加键盘输入的号码到input_name中
          for(let j = 0;j<this.bump1.length;j++){
            this.bump1[j].input_name = dom[51+j].value
          }
          for(let k = 0;k<this.bump2.length;k++){
            this.bump2[k].input_name = dom[56+k].value
          }
        }
      }

      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          let b = this.lists1[i].object[j].money + 'b';
          this.lists1[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.lists1[i].object[j].money);
          if(this.lists1[i].object[j].flag){
            is_select = true
          }
        }
      }
      //任意碰验证
      var zxc = false;
      var asd = false;
      for (let i = 0; i < this.bump1.length; i++) {
        if (this.bump1[i].input_name != '') {
          zxc = true
        }
      }
      for (let j = 0; j < this.bump2.length; j++) {
        if (this.bump2[j].input_name != '') {
          asd = true
        }
      }
      //生肖验证
      var shengxiao1 = false;
      var shengxiao2 = false;
      for (let j = 0; j < this.shengxiao1.length; j++) {
        if (this.shengxiao1[j].flag) {
          shengxiao1 = true
        }
      }
      for (let j = 0; j < this.shengxiao2.length; j++) {
        if (this.shengxiao2[j].flag) {
          shengxiao2 = true
        }
      }
      //尾数验证
      var weishu1 = false;
      var weishu2 = false;
      for (let j = 0; j < this.weishu1.length; j++) {
        if (this.weishu1[j].flag) {
          weishu1 = true
        }
      }
      for (let j = 0; j < this.weishu2.length; j++) {
        if (this.weishu2[j].flag) {
          weishu2 = true
        }
      }
      //生尾验证
      var shengwei1 = false;
      var shengwei2 = false;
      for (let j = 0; j < this.shengwei1.length; j++) {
        if (this.shengwei1[j].flag) {
          shengwei1 = true
        }
      }
      for (let j = 0; j < this.shengwei2.length; j++) {
        if (this.shengwei2[j].flag) {
          shengwei2 = true
        }
      }

      console.log('kk:' + kk);
      console.log(this.arr.length);
      if(this.two == true || this.three == true || this.four == true || this.five == true){
        this.arr = [];
        this.arr.push({'mingxi':'null'})
      }
      if (is_select && this.arr.length < 2 && this.arr[0].mingxi == 'second_full') {
        this.$Modal.warning({
          content: '请选择2-5个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.arr.length < 2 && this.arr[0].mingxi == '2/2:30,2/1:50') {
        this.$Modal.warning({
          content: '请选择2-5个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.arr.length < 2 && this.arr[0].mingxi == 'Special_series') {
        this.$Modal.warning({
          content: '请选择2-5个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.arr.length < 3 && this.arr[0].mingxi == 'third_full') {
        this.$Modal.warning({
          content: '请选择3-6个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.arr.length < 3 && this.arr[0].mingxi == '3/3:10,3/2:20') {
        this.$Modal.warning({
          content: '请选择3-6个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.arr.length < 4 && this.arr[0].mingxi == 'fourth_full') {
        this.$Modal.warning({
          content: '请选择4-7个号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (kk != 0 && this.arr.length >= 2 && this.one) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        this.$root.$emit('id-selected-dan', this.arr, this.top_config);
        console.log(1)
      } else if (kk != 0 && this.arr.length >= 2) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        this.$root.$emit('id-selected-normal', this.lists1);
      } else if (this.money != 0 && this.two && shengxiao1 && shengxiao2) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        console.log(2);
        this.$root.$emit('id-selected-sw', this.config_shengxiao1, this.config_shengxiao2, this.lists1, this.money)
      } else if (this.money != 0 && this.three && weishu1 && weishu2) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        this.$root.$emit('id-selected-sw', this.config_weishu1, this.config_weishu2, this.lists1, this.money)
      } else if (this.money != 0 && this.four && shengwei1 && shengwei2) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        this.$root.$emit('id-selected-sw', this.config_shengxiao1, this.config_weishu1, this.lists1, this.money);
        console.log(4)
      } else if (this.money != 0 && this.five && this.ok == true && zxc == false) {
        console.log(this.bump);
        this.$Modal.warning({
          content: '请选择对碰的任意号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (this.money != 0 && this.five && this.ok == true && asd == false ) {
        console.log(this.bump);
        this.$Modal.warning({
          content: '请选择对碰的任意号码'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.two && shengxiao1 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的生肖'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.two && shengxiao2 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的生肖'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.three && weishu1 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的尾数'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.three && weishu2 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的尾数'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.four && shengwei1 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的生肖'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }else if (this.money != 0 && this.four && shengwei2 == false ) {
        this.$Modal.warning({
          content: '请选择对碰的尾数'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (this.money != 0 && this.five && zxc && asd) {
        this.modal = true;
//        document.querySelector('body').style.overflow = 'hidden';
        this.$root.$emit('id-selected-random', this.bump1, this.bump2, this.lists1, this.money);
        console.log(5)
      }else if (!is_select && this.money == 0 && !this.two && !this.three && !this.four && !this.five) {
        this.$Modal.warning({
          content: hint.all_null
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if (is_select && this.money == 0) {
        this.$Modal.warning({
          content: hint.money_null
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      } else if(this.money == 0){
          this.$Modal.warning({
            content: hint.money_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
      }
    },
    remove: function(val) {
      var index = this.indexOf(val);
      if (index > -1) {
        this.bump_arr.splice(index, 1);
      }
    },
    //任意碰输入框验证
    five_verify: function(e) {
      e.input_name = e.input_name.replace(/\D/g, "");
      if (e.input_name > 49 || e.input_name == 0) {
        e.input_name = ''
      }
    },
    five_out: function(e, key,type) {
      if(type == 1){
        console.log('输入的数字是什么？：'+e.input_name);
        let dom = document.querySelectorAll('input');
        for(let i = 51;i < dom.length;i++) {
            if (dom[i].value > 49 || dom[i].value == 0) {
                dom[i].value = '';
            }
        }
        e.input_name = this.$refs.five_input1[key].$refs.input.value;
        if (e.input_name > 49 || e.input_name == 0) {
          e.input_name = '';
        }
      }else{
        let dom = document.querySelectorAll('input');
        for(let i = 51;i < dom.length;i++) {
            if (dom[i].value > 49 || dom[i].value == 0) {
                dom[i].value = '';
            }
        }
        e.input_name = this.$refs.five_input2[key].$refs.input.value;
        if (e.input_name > 49 || e.input_name == 0) {
            e.input_name = '';
        }
      }
      this.ok = false;
      console.log(this.bump_arr);
      if (e.input_name != '') {
        for (let i = 0; i < this.bump_arr.length; i++) {
          if (this.bump_arr[i].index == e.index) {
            console.log('我的index:' + e.index);
            split.splice_arr(this.bump_arr, this.bump_arr[i]);
            break;
          }
        }
        this.bump_arr.push(e);
        if (this.bump_arr.length >= 2) {
          let a = this.bump_arr.pop();
          console.log('a:' + a);
          for (let i = 0; i < this.bump_arr.length; i++) {
            if (this.bump_arr[i].input_name == e.input_name) {
              e.input_name = '';
              if(type == 1){
                this.$refs.five_input1[key].$refs.input.value = '';
              }else{
                this.$refs.five_input2[key].$refs.input.value = '';
              }
              this.ok = true;
              this.$Modal.warning({
                content: '抱歉不允许输入重复号码！'
              });
              window.setTimeout(() => {
                this.$Modal.remove()
              }, share.Prompt);
              break;
            } else {
              this.ok = false;
            }
          }
        } else if (this.bump_arr.length == 1) {
          this.ok = true;
          console.log('nothing to happen');
        }
      } else if (e.input_name == '') {
        this.ok = true;
        for (let i = 0; i < this.bump_arr.length; i++) {
          if (this.bump_arr[i].index == e.index) {
            console.log('我的index:' + e.index);
            split.splice_arr(this.bump_arr, this.bump_arr[i])
          }
        }
      }
      if (!this.ok) {
        this.bump_arr.push(e);
      }
      console.log(this.bump_arr)
    },
    cancel: function() {
      this.modal = false;
//      document.querySelector('body').style.overflow = 'auto'
    },
    add_money: function(type) {
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money()
    },
    push_money: function() {
      this.money = this.money.replace(/\D/g, "");
      this.computed_money()
    },
    change_money: function () { this.computed_money() },
    computed_money: function() {
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          //添加金额参数入对象
          if (this.lists1[i].object[j].flag) {
            this.lists1[i].object[j].money = this.money
          } else if (!this.lists1[i].object[j].flag) {
            this.lists1[i].object[j].money = ''
          }
        }
      }
    },
    fetchData: function(type) {
      this.reset();
      type==2?this.$root.$emit('loading',true,true):this.$root.$emit('loading',true);
      this.$root.$emit('child_change',0);
      let body = {
        'fc_type': this.$route.query.page,
        'gameplay': 170,
        'pankou': 'A'
      };
      api.getgameindex(this, body, (res) => {
          this.$root.$emit('only_back',res,type);
        if (res.data.ErrorCode == 1) {
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
          }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
          this.$root.$emit('c_data', res.data.Data.c_data);
          this.shengxiao_data = res.data.shengxiao;
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          this.back_data = back_data;
          for (let i = 0; i < this.back_data.length; i++) {
            delete this.back_data[i].boll_name
          }
          console.log(this.back_data);
          this.computed(this.back_data, 0);
          this.go_child(0);
        }
      })
    },
    computed: function(data, type) {
      this.lists1 = get_lists1();
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          Object.assign(this.lists1[i].object[j], data[type])
        }
      }
    },
    //数组去重
    unique: function(arr) {
      var newArr = [];
      for (var i in arr) {
        if (newArr.indexOf(arr[i]) == -1) {
          newArr.push(arr[i])
        }
      }
      return newArr;
    }
  }
}
</script>
<style lang="scss" src="../../../assets/css/six_six.scss" scoped></style>
