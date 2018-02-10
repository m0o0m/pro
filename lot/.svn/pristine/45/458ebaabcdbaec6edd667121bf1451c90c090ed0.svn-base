<template lang="html">
  <div class="pk_time_1">
    <div class="pk_time">
      <div class="time_left">
        <div class="left_logo">
          <img :src="c_data.img_path" alt="">
        </div>
        <div class="left_text">
          <p>{{c_data.fc_name}}</p>
          <p><span>第</span> {{c_data.qishu}} <span>期</span></p>
        </div>
      </div>
      <div class="time_center">
        <p class="center_top">投注剩余时间</p>
        <div class="center_bottom">
          <div class="fl time_content" v-if="h < 10 && h >= 0">0{{h}}</div>
          <div class="fl time_content" v-else-if="h < 0">--</div>
          <div class="fl time_content" v-else>{{h}}</div>
          <div class="fl fs">时</div>
          <div class="fl time_content" v-if="m < 10 && m >= 0">0{{m}}</div>
          <div class="fl time_content" v-else-if="m < 0">--</div>
          <div class="fl time_content" v-else>{{m}}</div>
          <div class="fl fs">分</div>
          <div class="fl time_content" v-if="s < 10 && s >= 0">0{{s}}</div>
          <div class="fl time_content" v-else-if="s < 0">--</div>
          <div class="fl time_content" v-else>{{s}}</div>
          <div class="fl fs">秒</div>
        </div>
      </div>
      <div class="time_bottom">
        <div class="bottom_left"><span>第</span> {{auto.qishu}} <span>期开奖</span></div>
        <div v-if="this.$route.query.page == 'liuhecai'" class='bottom_center6'>
          <div class="bottom_content">
            <div class="content_list" v-for="item in auto.ball">
              <p :class="[item.color,'box']">{{item.number}}</p>
              <p class="animal">{{item.animal}}</p>
            </div>
          </div>
        </div>
        <div v-else :class="[auto.ball.length >= 10?'bottom_center_other':'bottom_center']">
          <div :class="[auto.ball.length == 10?'other_content':'bottom_content']">
            <div class="content_list" v-for="item in auto.ball">
              <p class="box blue">{{item}}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="bottom_bottom">
        <p @click="history()">历史结果</p>
        <p @click="open_way()">开奖走势</p>
        <p @click="show_rule()">玩法规则</p>
      </div>
    </div>
    <Nav-top :lists="menus" :isActive="margin" v-on:menu="go_child"></Nav-top>
    <router-view :one="one" :two="two" :c_data="c_data"></router-view>
    <Nav-bottom ref="dewdrop_map" :back_data="auto_list" :nav_top="bottom_nav"></Nav-bottom>
  </div>
</template>

<script type="text/ecmascript-6">
function one() {
  return [
    {
      name: "第一球",
      object: [{
        num: "0",
        number: "9.700",
        flag: false,
        money: '',
        id: 0
      },
        {
          num: "2",
          number: "9.700",
          flag: false,
          money: '',
          id: 1
        },
        {
          num: "3",
          number: "9.700",
          flag: false,
          money: '',
          id: 2
        },
        {
          num: "4",
          number: "9.700",
          flag: false,
          money: '',
          id: 3
        },
        {
          num: "5",
          number: "9.700",
          flag: false,
          money: '',
          id: 4
        },
        {
          num: "6",
          number: "9.700",
          flag: false,
          money: '',
          id: 5
        },
        {
          num: "7",
          number: "9.700",
          flag: false,
          money: '',
          id: 6
        },
        {
          num: "8",
          number: "9.700",
          flag: false,
          money: '',
          id: 7
        },
        {
          num: "9",
          number: "9.700",
          flag: false,
          money: '',
          id: 8
        },
        {
          num: "10",
          number: "9.700",
          flag: false,
          money: '',
          id: 9
        },
        {
          num: "大",
          number: "1.930",
          flag: false,
          money: '',
          id: 10
        },
        {
          num: "小",
          number: "1.930",
          flag: false,
          money: '',
          id: 11
        },
        {
          num: "单",
          number: "1.930",
          flag: false,
          money: '',
          id: 12
        },
        {
          num: "双",
          number: "1.930",
          flag: false,
          money: '',
          id: 13
        }
      ]
    },
    {
      name: "第二球",
      object: [{
        num: "0",
        number: "9.700",
        flag: false,
        money: '',
        id: 14
      },
        {
          num: "1",
          number: "9.700",
          flag: false,
          money: '',
          id: 15
        },
        {
          num: "2",
          number: "9.700",
          flag: false,
          money: '',
          id: 16
        },
        {
          num: "3",
          number: "9.700",
          flag: false,
          money: '',
          id: 17
        },
        {
          num: "4",
          number: "9.700",
          flag: false,
          money: '',
          id: 18
        },
        {
          num: "5",
          number: "9.700",
          flag: false,
          money: '',
          id: 19
        },
        {
          num: "6",
          number: "9.700",
          flag: false,
          money: '',
          id: 20
        },
        {
          num: "7",
          number: "9.700",
          flag: false,
          money: '',
          id: 21
        },
        {
          num: "8",
          number: "9.700",
          flag: false,
          money: '',
          id: 22
        },
        {
          num: "9",
          number: "9.700",
          flag: false,
          money: '',
          id: 23
        },
        {
          num: "大",
          number: "1.930",
          flag: false,
          money: '',
          id: 24
        },
        {
          num: "小",
          number: "1.930",
          flag: false,
          money: '',
          id: 25
        },
        {
          num: "单",
          number: "1.930",
          flag: false,
          money: '',
          id: 26
        },
        {
          num: "双",
          number: "1.930",
          flag: false,
          money: '',
          id: 27
        }
      ]
    },
    {
      name: "第三球",
      object: [{
        num: "0",
        number: "9.700",
        flag: false,
        money: '',
        id: 28
      },
        {
          num: "1",
          number: "9.700",
          flag: false,
          money: '',
          id: 29
        },
        {
          num: "2",
          number: "9.700",
          flag: false,
          money: '',
          id: 30
        },
        {
          num: "3",
          number: "9.700",
          flag: false,
          money: '',
          id: 31
        },
        {
          num: "4",
          number: "9.700",
          flag: false,
          money: '',
          id: 32
        },
        {
          num: "5",
          number: "9.700",
          flag: false,
          money: '',
          id: 33
        },
        {
          num: "6",
          number: "9.700",
          flag: false,
          money: '',
          id: 34
        },
        {
          num: "7",
          number: "9.700",
          flag: false,
          money: '',
          id: 35
        },
        {
          num: "8",
          number: "9.700",
          flag: false,
          money: '',
          id: 36
        },
        {
          num: "9",
          number: "9.700",
          flag: false,
          money: '',
          id: 37
        },
        {
          num: "大",
          number: "1.930",
          flag: false,
          money: '',
          id: 38
        },
        {
          num: "小",
          number: "1.930",
          flag: false,
          money: '',
          id: 39
        },
        {
          num: "单",
          number: "1.930",
          flag: false,
          money: '',
          id: 40
        },
        {
          num: "双",
          number: "1.930",
          flag: false,
          money: '',
          id: 41
        }
      ]
    }
  ]
}
function two() {
  return [
    {
      num: "独胆",
      object: [
        {
          num: "0",
          number: "3.500",
          flag: false,
          id: 0,
          index: 0,
          money:''
        },
        {
          num: "1",
          number: "3.500",
          flag: false,
          id: 1,
          index: 1,
          money:''
        },
        {
          num: "2",
          number: "3.500",
          flag: false,
          id: 2,
          index: 2,
          money:''
        },
        {
          num: "3",
          number: "3.500",
          flag: false,
          id: 3,
          index: 3,
          money:''
        },
        {
          num: "4",
          number: "3.500",
          flag: false,
          id: 4,
          index: 4,
          money:''
        },
        {
          num: "5",
          number: "3.500",
          flag: false,
          id: 5,
          index: 5,
          money:''
        },
        {
          num: "6",
          number: "3.500",
          flag: false,
          id: 6,
          index: 6,
          money:''
        },
        {
          num: "7",
          number: "3.500",
          flag: false,
          id: 7,
          index: 7,
          money:''
        },
        {
          num: "8",
          number: "3.500",
          flag: false,
          id: 8,
          index: 8,
          money:''
        },
        {
          num: "9",
          number: "3.500",
          flag: false,
          id: 9,
          index: 9,
          money:''
        }
      ]
    },
    {
      num: "跨度",
      object: [
        {
          num: "0",
          number: "86.00",
          flag: false,
          id: 10,
          index: 10,
          money:''
        },
        {
          num: "1",
          number: "16.700",
          flag: false,
          id: 11,
          index: 11,
          money:''
        },
        {
          num: "2",
          number: "9.400",
          flag: false,
          id: 12,
          index: 12,
          money:''
        },
        {
          num: "3",
          number: "7.200",
          flag: false,
          id: 13,
          index: 13,
          money:''
        },
        {
          num: "4",
          number: "6.200",
          flag: false,
          id: 14,
          index: 14,
          money:''
        },
        {
          num: "5",
          number: "6.000",
          flag: false,
          id: 15,
          index: 15,
          money:''
        },
        {
          num: "6",
          number: "6.200",
          flag: false,
          id: 16,
          index: 16,
          money:''
        },
        {
          num: "7",
          number: "7.200",
          flag: false,
          id: 17,
          index: 17,
          money:''
        },
        {
          num: "8",
          number: "9.400",
          flag: false,
          id: 18,
          index: 18,
          money:''
        },
        {
          num: "9",
          number: "16.700",
          flag: false,
          id: 19,
          index: 19,
          money:''
        }
      ]
    },
    {
      num: "总和龙虎",
      object: [
        {
          num: "总和大",
          number: "1.930",
          flag: false,
          id: 20,
          index: 20,
          money:''
        },
        {
          num: "总和小",
          number: "1.930",
          flag: false,
          id: 21,
          index: 21,
          money:''
        },
        {
          num: "总和单",
          number: "1.930",
          flag: false,
          id: 22,
          index: 22,
          money:''
        },
        {
          num: "总和双",
          number: "1.930",
          flag: false,
          id: 23,
          index: 23,
          money:''
        },
        {
          num: "龙",
          number: "1.930",
          flag: false,
          id: 24,
          index: 24,
          money:''
        },
        {
          num: "虎",
          number: "1.930",
          flag: false,
          id: 25,
          index: 25,
          money:''
        },
        {
          num: "和",
          number: "9.000",
          flag: false,
          id: 26,
          index: 26,
          money:''
        }
      ]
    },
    {
      num: "3连",
      object: [
        {
          num: "豹子",
          number: "70.00",
          flag: false,
          id: 27,
          index: 27,
          money:''
        },
        {
          num: "顺子",
          number: "13.000",
          flag: false,
          id: 28,
          index: 28,
          money:''
        },
        {
          num: "对子",
          number: "2.800",
          flag: false,
          id: 29,
          index: 29,
          money:''
        },
        {
          num: "半顺",
          number: "2.000",
          flag: false,
          id: 30,
          index: 30,
          money:''
        },
        {
          num: "杂六",
          number: "2.200",
          flag: false,
          id: 31,
          index: 31,
          money:''
        }
      ]
    }
  ]
}
import api from "../../../api/config"
import NavTop from "../../../share_components/default_nav";
import ws from '../../../assets/js/socket'
import NavBottom from '../../../share_components/dewdrop_map'
import cm_cookie from '../../../assets/js/com_cookie'
export default {
  components: {
    NavTop, NavBottom
  },
  data() {
    return {
      auto_list:[],
      bottom_nav:[
          {name:'百位'},
          {name:'拾位'},
          {name:'个位'},
      ],
      timer: null,
      is_wh: false,
      h: 0,
      m: 0,
      s: 0,
      c_data:{
        fc_name:'',
        img_path:'',
        qishu:''
      },
      auto:{
        qishu:null,
        datetime:'',
        ball:[]
      },
      close_time:{
        fengpan:'',
        kaijiang:'',
        kaipan:'',
        now_time:''
      },
      one: [
          {
          name: "第一球",
          object: [{
              num: "0",
              number: "9.700",
              flag: false,
              money: '',
              id: 0
            },
            {
              num: "2",
              number: "9.700",
              flag: false,
              money: '',
              id: 1
            },
            {
              num: "3",
              number: "9.700",
              flag: false,
              money: '',
              id: 2
            },
            {
              num: "4",
              number: "9.700",
              flag: false,
              money: '',
              id: 3
            },
            {
              num: "5",
              number: "9.700",
              flag: false,
              money: '',
              id: 4
            },
            {
              num: "6",
              number: "9.700",
              flag: false,
              money: '',
              id: 5
            },
            {
              num: "7",
              number: "9.700",
              flag: false,
              money: '',
              id: 6
            },
            {
              num: "8",
              number: "9.700",
              flag: false,
              money: '',
              id: 7
            },
            {
              num: "9",
              number: "9.700",
              flag: false,
              money: '',
              id: 8
            },
            {
              num: "10",
              number: "9.700",
              flag: false,
              money: '',
              id: 9
            },
            {
              num: "大",
              number: "1.930",
              flag: false,
              money: '',
              id: 10
            },
            {
              num: "小",
              number: "1.930",
              flag: false,
              money: '',
              id: 11
            },
            {
              num: "单",
              number: "1.930",
              flag: false,
              money: '',
              id: 12
            },
            {
              num: "双",
              number: "1.930",
              flag: false,
              money: '',
              id: 13
            }
          ]
        },
        {
          name: "第二球",
          object: [{
              num: "0",
              number: "9.700",
              flag: false,
              money: '',
              id: 14
            },
            {
              num: "1",
              number: "9.700",
              flag: false,
              money: '',
              id: 15
            },
            {
              num: "2",
              number: "9.700",
              flag: false,
              money: '',
              id: 16
            },
            {
              num: "3",
              number: "9.700",
              flag: false,
              money: '',
              id: 17
            },
            {
              num: "4",
              number: "9.700",
              flag: false,
              money: '',
              id: 18
            },
            {
              num: "5",
              number: "9.700",
              flag: false,
              money: '',
              id: 19
            },
            {
              num: "6",
              number: "9.700",
              flag: false,
              money: '',
              id: 20
            },
            {
              num: "7",
              number: "9.700",
              flag: false,
              money: '',
              id: 21
            },
            {
              num: "8",
              number: "9.700",
              flag: false,
              money: '',
              id: 22
            },
            {
              num: "9",
              number: "9.700",
              flag: false,
              money: '',
              id: 23
            },
            {
              num: "大",
              number: "1.930",
              flag: false,
              money: '',
              id: 24
            },
            {
              num: "小",
              number: "1.930",
              flag: false,
              money: '',
              id: 25
            },
            {
              num: "单",
              number: "1.930",
              flag: false,
              money: '',
              id: 26
            },
            {
              num: "双",
              number: "1.930",
              flag: false,
              money: '',
              id: 27
            }
          ]
        },
        {
          name: "第三球",
          object: [{
              num: "0",
              number: "9.700",
              flag: false,
              money: '',
              id: 28
            },
            {
              num: "1",
              number: "9.700",
              flag: false,
              money: '',
              id: 29
            },
            {
              num: "2",
              number: "9.700",
              flag: false,
              money: '',
              id: 30
            },
            {
              num: "3",
              number: "9.700",
              flag: false,
              money: '',
              id: 31
            },
            {
              num: "4",
              number: "9.700",
              flag: false,
              money: '',
              id: 32
            },
            {
              num: "5",
              number: "9.700",
              flag: false,
              money: '',
              id: 33
            },
            {
              num: "6",
              number: "9.700",
              flag: false,
              money: '',
              id: 34
            },
            {
              num: "7",
              number: "9.700",
              flag: false,
              money: '',
              id: 35
            },
            {
              num: "8",
              number: "9.700",
              flag: false,
              money: '',
              id: 36
            },
            {
              num: "9",
              number: "9.700",
              flag: false,
              money: '',
              id: 37
            },
            {
              num: "大",
              number: "1.930",
              flag: false,
              money: '',
              id: 38
            },
            {
              num: "小",
              number: "1.930",
              flag: false,
              money: '',
              id: 39
            },
            {
              num: "单",
              number: "1.930",
              flag: false,
              money: '',
              id: 40
            },
            {
              num: "双",
              number: "1.930",
              flag: false,
              money: '',
              id: 41
            }
          ]
        }
      ],
      two: [
        {
          num: "独胆",
          object: [
            {
              num: "0",
              number: "3.500",
              flag: false,
              id: 0,
              index: 0,
              money:''
            },
            {
              num: "1",
              number: "3.500",
              flag: false,
              id: 1,
              index: 1,
              money:''
            },
            {
              num: "2",
              number: "3.500",
              flag: false,
              id: 2,
              index: 2,
              money:''
            },
            {
              num: "3",
              number: "3.500",
              flag: false,
              id: 3,
              index: 3,
              money:''
            },
            {
              num: "4",
              number: "3.500",
              flag: false,
              id: 4,
              index: 4,
              money:''
            },
            {
              num: "5",
              number: "3.500",
              flag: false,
              id: 5,
              index: 5,
              money:''
            },
            {
              num: "6",
              number: "3.500",
              flag: false,
              id: 6,
              index: 6,
              money:''
            },
            {
              num: "7",
              number: "3.500",
              flag: false,
              id: 7,
              index: 7,
              money:''
            },
            {
              num: "8",
              number: "3.500",
              flag: false,
              id: 8,
              index: 8,
              money:''
            },
            {
              num: "9",
              number: "3.500",
              flag: false,
              id: 9,
              index: 9,
              money:''
            }
          ]
        },
        {
          num: "跨度",
          object: [
            {
              num: "0",
              number: "86.00",
              flag: false,
              id: 10,
              index: 10,
              money:''
            },
            {
              num: "1",
              number: "16.700",
              flag: false,
              id: 11,
              index: 11,
              money:''
            },
            {
              num: "2",
              number: "9.400",
              flag: false,
              id: 12,
              index: 12,
              money:''
            },
            {
              num: "3",
              number: "7.200",
              flag: false,
              id: 13,
              index: 13,
              money:''
            },
            {
              num: "4",
              number: "6.200",
              flag: false,
              id: 14,
              index: 14,
              money:''
            },
            {
              num: "5",
              number: "6.000",
              flag: false,
              id: 15,
              index: 15,
              money:''
            },
            {
              num: "6",
              number: "6.200",
              flag: false,
              id: 16,
              index: 16,
              money:''
            },
            {
              num: "7",
              number: "7.200",
              flag: false,
              id: 17,
              index: 17,
              money:''
            },
            {
              num: "8",
              number: "9.400",
              flag: false,
              id: 18,
              index: 18,
              money:''
            },
            {
              num: "9",
              number: "16.700",
              flag: false,
              id: 19,
              index: 19,
              money:''
            }
          ]
        },
        {
          num: "总和龙虎",
          object: [
            {
              num: "总和大",
              number: "1.930",
              flag: false,
              id: 20,
              index: 20,
              money:''
            },
            {
              num: "总和小",
              number: "1.930",
              flag: false,
              id: 21,
              index: 21,
              money:''
            },
            {
              num: "总和单",
              number: "1.930",
              flag: false,
              id: 22,
              index: 22,
              money:''
            },
            {
              num: "总和双",
              number: "1.930",
              flag: false,
              id: 23,
              index: 23,
              money:''
            },
            {
              num: "龙",
              number: "1.930",
              flag: false,
              id: 24,
              index: 24,
              money:''
            },
            {
              num: "虎",
              number: "1.930",
              flag: false,
              id: 25,
              index: 25,
              money:''
            },
            {
              num: "和",
              number: "9.000",
              flag: false,
              id: 26,
              index: 26,
              money:''
            }
          ]
        },
        {
          num: "3连",
          object: [
            {
              num: "豹子",
              number: "70.00",
              flag: false,
              id: 27,
              index: 27,
              money:''
            },
            {
              num: "顺子",
              number: "13.000",
              flag: false,
              id: 28,
              index: 28,
              money:''
            },
            {
              num: "对子",
              number: "2.800",
              flag: false,
              id: 29,
              index: 29,
              money:''
            },
            {
              num: "半顺",
              number: "2.000",
              flag: false,
              id: 30,
              index: 30,
              money:''
            },
            {
              num: "杂六",
              number: "2.200",
              flag: false,
              id: 31,
              index: 31,
              money:''
            }
          ]
        }
      ],
      margin: false,
      menus: [
        {name:'1-3球',item:'pl_3'},
        {name:'整合',item:'yb_two'}
      ]
    };
  },
  created(){
  },
  mounted(){
    this.fetchData();
    this.socket_change(this.$route.query.page);
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route.query.page':function(to,from) {
      this.$root.$off(from);
      this.$root.$off(from+'lefttime');
      this.fetchData();
      this.socket_change(to);
    }
  },
  destroyed(){
    console.log('清除定时器：'+this.timer);
//    window.clearInterval(this.timer);
    if(this.timer){
        window.clearTimeout(this.timer);
        this.timer = null;
    }
    if(!this.isIE9){
      ws.close_ws(false);
    }
    this.$root.$off(this.$route.query.page);
    this.$root.$off(this.$route.query.page+'lefttime');
    cm_cookie.delCookie("top_nav")
  },
  methods:{
    socket_change: function(to){
//      console.log('是否在维护：'+this.is_wh);
      if(!this.isIE9){
//        console.log('我进来了！！');
        let self = this;
        ws.createWebSocket(to,self,true);
        this.$root.$on(to,(e)=>{
          this.auto = e;
        });
        this.$root.$on(to+'lefttime',(e)=>{
          this.c_data.qishu = e.qishu;
          this.close_time.fengpan = e.close_time;
          this.close_time.now_time = e.now_time;
          let t1 = e.close_time - e.now_time;
//          console.log('是否在维护中：'+this.is_wh);
          if(t1 == 0 && !this.is_wh){
              this.fetchData(2)
          }else if(t1 > 0  && !this.is_wh){
              if(this.timer){
                  window.clearTimeout(this.timer);
                  this.timer = null;
              }
              this.init();
              this.$root.$emit('wh_modal',false);
          }else if(t1 < 0  && !this.is_wh){
              this.h = -1;
              this.m = -1;
              this.s = -1;
              if(this.timer){
                  window.clearTimeout(this.timer);
                  this.timer = null;
              }
              this.$root.$emit('wh_modal',true,true)
          }
        })
      }
    },
    open_way: function (){
      let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=1';
      window.open(page)
    },
    history: function () {
      let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=3';
      window.open(page)
    },
    //aaaaa
    show_rule: function(){
      this.$root.$emit('rule_show',true);
      this.$root.$emit('now_page',this.$route.query.page)
    },
    getRTime: function() {
      this.close_time.now_time += 1;
      var t1 = this.close_time.fengpan * 1000 - this.close_time.now_time * 1000;
      // var d=Math.floor(t/1000/60/60/24);
      this.h = Math.floor(t1 / 1000 / 60 / 60 % 24);
      this.m = Math.floor(t1 / 1000 / 60 % 60);
      this.s = Math.floor(t1 / 1000 % 60);
      if (this.h == 0 && this.m == 0 && this.s == 0) {
        this.fetchData(2);
        this.$root.$emit("success", true);
      }else if(this.h < 0 && this.m < 0 && this.s < 0){
        if(this.timer){
            window.clearTimeout(this.timer);
            this.timer = null;
        }
        this.h = -1;
        this.m = -1;
        this.s = -1;
      }
    },
    init: function() {
      this.getRTime();
      this.timer = window.setTimeout(this.init,1000);
    },
    //aaaaaa
    sortNumber: function(a,b){
      return a.sort - b.sort
    },
    fetchData: function(type){
      this.$root.$emit('wh_modal',false);
      if(this.timer){
          window.clearTimeout(this.timer);
          this.timer = null;
      }
      type==2?this.$root.$emit('loading',true,true):this.$root.$emit('loading',true);
      let body = {
        'fc_type': this.$route.query.page,
      };
     api.dewdrop(this, body, (res) => {
         if (res.data.ErrorCode == 1) {
//             console.log(res);
             this.auto_list = res.data.Data;
//             console.log(this.$refs.dewdrop_map);
            api.getgameindex(this, body, (res) => {
                if (res.data.ErrorCode == 1) {
                if(type == 2){
                    window.setTimeout(() => {
                        this.$root.$emit("loading", false);
                }, 1000)
                }else{
                    this.$root.$emit("loading", false);
                }
                if(res.data.is_wh == 2){
                    this.is_wh = true;
                    this.$root.$emit('wh_modal',true)
                }else if(res.data.is_wh == 1){
                    this.is_wh = false;
                    this.$root.$emit('wh_modal',false)
                }
                let back_data = res.data.Data.odds;
                back_data.sort(this.sortNumber);
                // console.log(back_data);
                this.computed(back_data);
                this.computed_or(back_data);
                this.$refs.dewdrop_map.top_go(0);//点击触发露珠图组件头部选中事件
                this.$refs.dewdrop_map.left_go(0);//点击触发露珠图组件左侧选中事件
                this.auto = res.data.Data.auto;
                this.close_time = res.data.Data.closetime;
                if(this.close_time.fengpan - this.close_time.now_time < 0){
                    this.$root.$emit('wh_modal',true,true)
                }
                this.c_data = res.data.Data.c_data;
                if(!this.is_wh){
                    if(this.close_time.fengpan){
                        if(this.timer){
                            window.clearTimeout(this.timer);
                            this.timer = null;
                        }
                        if(type == 2){
                            window.setTimeout(() => {
                                this.init();
                            }, 1000)
                        }else{
                            this.init();
                        }
                    }
                }
            }
        })
         }
     });

    },
    computed: function(data) {
      this.$set(this.one, this.one);
      this.one = one();
      for (let j = 0; j < this.one[0].object.length; j++) {
        Object.assign(this.one[0].object[j], data[j]);
        let name = data[j].remark.slice(
          data[j].remark.search("#") + 1,
          data[j].remark.length
        );
        this.one[0].object[j].num = name;
      }
      for (let l = 14, k = 0; l < this.one[1].object.length, k < this.one[1].object.length; l++, k++) {
        Object.assign(this.one[1].object[k], data[l]);
        let name = data[l].remark.slice(
          data[l].remark.search("#") + 1,
          data[l].remark.length
        );
        this.one[1].object[k].num = name;
      }
      for (let n = 28, m = 0; n < this.one[2].object.length, m < this.one[2].object.length; n++, m++) {
        Object.assign(this.one[2].object[m], data[n]);
        let name = data[n].remark.slice(
          data[n].remark.search("#") + 1,
          data[n].remark.length
        );
//        console.log(name);
        this.one[2].object[m].num = name;
      }
    },
    computed_or: function(data) {
      this.$set(this.two, this.two);
      this.two = two();
      for (let j = 42, i = 0; j < this.two[0].object.length,i < this.two[0].object.length; j++,i++) {
        Object.assign(this.two[0].object[i], data[j]);
        let name = data[j].remark.slice(
          data[j].remark.search("#") + 1,
          data[j].remark.length
        );
//        console.log(name);
        this.two[0].object[i].num = name;
      }
      for (let l = 52, k = 0; l < this.two[1].object.length, k < this.two[1].object.length; l++, k++) {
        Object.assign(this.two[1].object[k], data[l]);
        let name = data[l].remark.slice(
          data[l].remark.search("#") + 1,
          data[l].remark.length
        );
//        console.log(name);
        this.two[1].object[k].num = name;
      }
      for (let n = 62, m = 0; n < this.two[2].object.length, m < this.two[2].object.length; n++, m++) {
        Object.assign(this.two[2].object[m], data[n]);
        let name = data[n].remark.slice(
          data[n].remark.search("#") + 1,
          data[n].remark.length
        );
//        console.log(name);
        this.two[2].object[m].num = name;
      }
      for (let n = 69, m = 0; n < this.two[3].object.length, m < this.two[3].object.length; n++, m++) {
        Object.assign(this.two[3].object[m], data[n]);
        let name = data[n].remark.slice(
          data[n].remark.search("#") + 1,
          data[n].remark.length
        );
//        console.log(name);
        this.two[3].object[m].num = name;
      }
    },
    go_child: function(child){
      this.$router.push({
        name:child,
        query: {page: this.$route.query.page}
      });
    }
  }
};
</script>
<style lang="scss" scoped>
@import '../../../assets/css/function.scss';
.pk_time_1{
  overflow: hidden;
  width: 920px
}
</style>
