<template lang="html">
  <div class="timelottery">
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
        <div class="bottom_center">
          <div class="bottom_content">
            <div class="content_list_xy" v-for="(item,key) in auto.ball">
              <p class="list">
                <span class="box blue">{{item}}</span>
                <span v-if="key<2" class="another">+</span>
                <span v-else-if="key==2" class="another">=</span>

              </p>
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
    <router-view :cdata="c_data" :luckyLists="luckyLists" :pcdd="pcdd"></router-view>
  </div>
</template>

<script>
  function luckyLists() {
    return [
      {
        object: [
          {
            flag: false,
            li_id: 0,
            money: '',
            class:['双','小','边','小双','小边','双边','小尾','0尾','3余0','4余0','5余0']
          },
          {
            flag: false,
            li_id: 1,
            money: '',
            class:['单','小','边','小单','小边','单边','小尾','1尾','3余1','4余1','5余1']
          },
          {
            flag: false,
            li_id: 2,
            money: '',
            class:['双','小','边','小双','小边','双边','小尾','2尾','3余2','4余2','5余2']
          },
          {
            flag: false,
            li_id: 3,
            money: '',
            class:['单','小','边','小单','小边','单边','小尾','3尾','3余0','4余3','5余3']
          },
          {
            flag: false,
            li_id: 4,
            money: '',
            class:['双','小','边','小双','小边','双边','小尾','4尾','3余1','4余0','5余4']
          },
          {
            flag: false,
            li_id: 5,
            money: '',
            class:['单','小','边','小单','小边','单边','大尾','5尾','3余2','4余1','5余0']
          },
          {
            flag: false,
            li_id: 6,
            money: '',
            class:['双','小','边','小双','小边','双边','大尾','6尾','3余0','4余2','5余1']
          },
          {
            flag: false,
            li_id: 7,
            money: ''
          },
          {
            flag: false,
            li_id: 8,
            money: ''
          },
          {
            flag: false,
            li_id: 9,
            money: ''
          }
        ]
      },
      {
        object: [
          {
            flag: false,
            money: '',
            li_id: 10,
            class:['单','小','边','小单','小边','单边','大尾','7尾','3余0','4余1','5余4']
          },
          {
            flag: false,
            money: '',
            li_id: 11,
            class:['双','小','边','小双','小边','双边','大尾','8尾','3余2','4余0','5余3']
          },
          {
            flag: false,
            money: '',
            li_id: 12,
            class:['单','小','边','小单','小边','单边','大尾','9尾','3余0','4余1','5余4']
          },
          {
            flag: false,
            money: '',
            li_id: 13,
            class:['双','小','中','小双','小尾','0尾','3余1','4余2','5余0']
          },
          {
            flag: false,
            money: '',
            li_id: 14,
            class:['单','小','中','小单','小尾','1尾','3余2','4余3','5余1']
          },
          {
            flag: false,
            money: '',
            li_id: 15,
            class:['双','小','中','小双','小尾','2尾','3余0','4余0','5余2']
          },
          {
            flag: false,
            money: '',
            li_id: 16,
            class:['单','小','中','小单','小尾','3尾','3余1','4余1','5余3']
          },
          {
            flag: false,
            money: '',
            li_id: 17
          },
          {
            flag: false,
            money: '',
            li_id: 18
          },
          {
            flag: false,
            money: '',
            li_id: 19
          }
        ]
      },
      {
        object: [
          {
            flag: false,
            money: '',
            li_id: 20,
            class:['双','大','中','大双','小尾','4尾','3余2','4余2','5余4']
          },
          {
            flag: false,
            money: '',
            li_id: 21,
            class:['单','大','中','大单','大尾','5尾','3余0','4余3','5余0']
          },
          {
            flag: false,
            money: '',
            li_id: 22,
            class:['双','大','中','大双','大尾','6尾','3余1','4余0','5余1']
          },
          {
            flag: false,
            money: '',
            li_id: 23,
            class:['单','大','中','大单','大尾','7尾','3余2','4余1','5余2']
          },
          {
            flag: false,
            money: '',
            li_id: 24,
            class:['双','大','边','大双','大边','双边','大尾','8尾','3余2','4余0','5余3']
          },
          {
            flag: false,
            money: '',
            li_id: 25,
            class:['单','大','边','大单','大边','单边','大尾','9尾','3余1','4余3','5余4']
          },
          {
            flag: false,
            money: '',
            li_id: 26,
            class:['双','大','边','大双','大边','双边','小尾','0尾','3余2','4余0','5余0']
          },
          {
            flag: false,
            money: '',
            li_id: 27
          },
          {
            flag: false,
            money: '',
            li_id: 28
          }
        ]
      },
      {
        object: [
          {
            flag: false,
            money: '',
            li_id: 29,
            class:['单','大','边','大双','大边','单边','小尾','1尾','3余0','4余1','5余1']
          },
          {
            flag: false,
            money: '',
            li_id: 30,
            class:['双','大','边','大双','大边','双边','小尾','2尾','3余1','4余2','5余2']
          },
          {
            flag: false,
            money: '',
            li_id: 31,
            class:['单','大','边','大单','大边','单边','小尾','3尾','3余2','4余3','5余3']
          },
          {
            flag: false,
            money: '',
            li_id: 32,
            class:['双','大','边','大双','大边','双边','小尾','4尾','3余0','4余0','5余4']
          },
          {
            flag: false,
            money: '',
            li_id: 33,
            class:['单','大','边','大单','大边','单边','大尾','5尾','3余1','4余1','5余0']
          },
          {
            flag: false,
            money: '',
            li_id: 34,
            class:['双','大','边','大双','大边','双边','大尾','6尾','3余2','4余2','5余1']
          },
          {
            flag: false,
            money: '',
            li_id: 35,
            class:['单','大','边','大单','大边','单边','大尾','7尾','3余0','4余3','5余2']
          },
          {
            flag: false,
            money: '',
            li_id: 36
          },
          {
            flag: false,
            money: '',
            li_id: 37
          }
        ]
      }
    ];
  }
  import NavTop from "../../../../share_components/default_nav";
  import api from "../../../../api/config";
  import ws from '../../../../assets/js/socket'
  import cm_cookie from '../../../../assets/js/com_cookie'
  export default {
    components: {
      NavTop
    },
    data() {
      return {
        margin: false,
        menus: [
          {
            name: "整合",
            item: null
          }
        ],
        timer: null,
        h: 0,
        m: 0,
        s: 0,
        c_data: {
          fc_name: "",
          img_path: "",
          qishu: ""
        },
        auto: {
          qishu: null,
          datetime: "",
          ball: []
        },
        close_time: {
          fengpan: "",
          kaijiang: "",
          kaipan: "",
          now_time: ""
        },
        luckyLists: [
          {
            object: [
              {
                flag: false,
                li_id: 0,
                money: null
              },
              {
                flag: false,
                li_id: 1,
                money: ''
              },
              {
                flag: false,
                li_id: 2,
                money: null
              },
              {
                flag: false,
                li_id: 3,
                money: null
              },
              {
                flag: false,
                li_id: 4,
                money: null
              },
              {
                flag: false,
                li_id: 5,
                money: null
              },
              {
                flag: false,
                li_id: 6,
                money: null
              },
              {
                flag: false,
                li_id: 7,
                money: null
              },
              {
                flag: false,
                li_id: 8,
                money: null
              },
              {
                flag: false,
                li_id: 9,
                money: null
              }
            ]
          },
          {
            object: [
              {
                flag: false,
                money: null,
                li_id: 10
              },
              {
                flag: false,
                money: null,
                li_id: 11
              },
              {
                flag: false,
                money: null,
                li_id: 12
              },
              {
                flag: false,
                money: null,
                li_id: 13
              },
              {
                flag: false,
                money: null,
                li_id: 14
              },
              {
                flag: false,
                money: null,
                li_id: 15
              },
              {
                flag: false,
                money: null,
                li_id: 16
              },
              {
                flag: false,
                money: null,
                li_id: 17
              },
              {
                flag: false,
                money: null,
                li_id: 18
              },
              {
                flag: false,
                money: null,
                li_id: 19
              }
            ]
          },
          {
            object: [
              {
                flag: false,
                money: null,
                li_id: 20
              },
              {
                flag: false,
                money: null,
                li_id: 21
              },
              {
                flag: false,
                money: null,
                li_id: 22
              },
              {
                flag: false,
                money: null,
                li_id: 23
              },
              {
                flag: false,
                money: null,
                li_id: 24
              },
              {
                flag: false,
                money: null,
                li_id: 25
              },
              {
                flag: false,
                money: null,
                li_id: 26
              },
              {
                flag: false,
                money: null,
                li_id: 27
              },
              {
                flag: false,
                money: null,
                li_id: 28
              }
            ]
          },
          {
            object: [
              {
                flag: false,
                money: null,
                li_id: 29
              },
              {
                flag: false,
                money: null,
                li_id: 30
              },
              {
                flag: false,
                money: null,
                li_id: 31
              },
              {
                flag: false,
                money: null,
                li_id: 32
              },
              {
                flag: false,
                money: null,
                li_id: 33
              },
              {
                flag: false,
                money: null,
                li_id: 34
              },
              {
                flag: false,
                money: null,
                li_id: 35
              },
              {
                flag: false,
                money: null,
                li_id: 36
              },
              {
                flag: false,
                money: null,
                li_id: 37
              }
            ]
          }
        ],
        pcdd: [
          {
            object: [
              {
                money: null,
                // name: "00",
                // odds: "888.000",
                flag: false,
                color: "",
                li_id: 0
              },
              {
                money: null,
                // name: "01",
                // odds: "300.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 1
              },
              {
                money: null,
                // name: "02",
                // odds: "150.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 2
              },
              {
                money: null,
                // name: "03",
                // odds: "80.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 3
              },
              {
                money: null,
                // name: "04",
                // odds: "60.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 4
              },
              {
                money: null,
                // name: "05",
                // odds: "30.000",
                color: "",
                flag: false,
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 5
              },
              {
                money: null,
                // name: "06",
                // odds: "25.000",
                color: "",
                flag: false,
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 6
              },
              {
                money: null,
                // name: "总和大",
                // odds: "1.930",
                flag: false,
                li_id: 7
              },
              {
                money: null,
                // name: "大单",
                // odds: "3.930",
                flag: false,
                li_id: 8
              },
              {
                money: null,
                // name: "极小",
                // odds: "10.000",
                flag: false,
                li_id: 9
              },
              {
                money: null,
                // name: "红波",
                // odds: "3.000",
                flag: false,
                li_id: 10
              },
              {
                money: null,
                name: "特码包三",
                odd: "3.50",
                flag: false,
                li_id: 42
              }
            ]
          },
          {
            object: [
              {
                money: null,
                // name: "07",
                // odds: "20.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 11
              },
              {
                money: null,
                // name: "08",
                // odds: "18.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 12
              },
              {
                money: null,
                // name: "09",
                // odds: "16.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 13
              },
              {
                money: null,
                // name: "10",
                // odds: "15.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 14
              },
              {
                money: null,
                // name: "11",
                // odds: "14.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 15
              },
              {
                money: null,
                // name: "12",
                // odds: "13.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 16
              },
              {
                money: null,
                // name: "13",
                // odds: "12.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_lightgreen.png"),
                li_id: 17
              },
              {
                money: null,
                // name: "总和小",
                // odds: "1.930",
                flag: false,
                li_id: 18
              },
              {
                money: null,
                // name: "大双",
                // odds: "3.470",
                flag: false,
                li_id: 19
              },
              {
                money: null,
                // name: "极大",
                // odds: "10.000",
                flag: false,
                li_id: 20
              },
              {
                money: null,
                // name: "豹子",
                // odds: "100.000",
                flag: false,
                li_id: 21
              }
            ]
          },
          {
            object: [
              {
                money: null,
                // name: "14",
                // odds: "12.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_lightgreen.png"),
                li_id: 22
              },
              {
                money: null,
                // name: "15",
                // odds: "13.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 23
              },
              {
                money: null,
                // name: "16",
                // odds: "14.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 24
              },
              {
                money: null,
                // name: "17",
                // odds: "15.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 25
              },
              {
                money: null,
                // name: "18",
                // odds: "16.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 26
              },
              {
                money: null,
                // name: "19",
                // odds: "18.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 27
              },
              {
                money: null,
                // name: "20",
                // odds: "20.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 28
              },
              {
                money: null,
                // name: "总和单",
                // odds: "1.930",
                flag: false,
                li_id: 29
              },
              {
                money: null,
                // name: "小单",
                // odds: "3.470",
                flag: false,
                li_id: 30
              },
              {
                money: null,
                // name: "绿波",
                // odds: "3.000",
                flag: false,
                li_id: 31
              }
            ]
          },
          {
            object: [
              {
                money: null,
                // name: "21",
                // odds: "25.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 32
              },
              {
                money: null,
                // name: "22",
                // odds: "30.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 33
              },
              {
                money: null,
                // name: "23",
                // odds: "60.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 34
              },
              {
                money: null,
                // name: "24",
                // odds: "80.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_red.png"),
                li_id: 35
              },
              {
                money: null,
                // name: "25",
                // odds: "150.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_yellogreen.png"),
                li_id: 36
              },
              {
                money: null,
                // name: "26",
                // odds: "300.000",
                flag: false,
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_blue.png"),
                li_id: 37
              },
              {
                money: null,
                // name: "27",
                // odds: "888.000",
                flag: false,
                color: "",
                // img_src: require("../../../assets/img/pcdd_lightgreen.png"),
                li_id: 38
              },
              {
                money: null,
                // name: "总和双",
                // odds: "1.930",
                flag: false,
                li_id: 39
              },
              {
                money: null,
                // name: "小双",
                // odds: "3.930",
                flag: false,
                li_id: 40
              },
              {
                money: null,
                // name: "蓝波",
                // odds: "3.000",
                flag: false,
                li_id: 41
              }
            ]
          }
        ],
        is_wh: false
      };
    },
    created() {
      this.fetchData();
    },
    mounted(){
      this.socket_change(this.$route.query.page);
    },
    watch: {
      // 如果路由有变化，会再次执行该方法
      '$route.query.page':function(to,from) {
        this.$root.$off(from);
        this.$root.$off(from+'lefttime');
        this.socket_change(to);
        this.fetchData();
      }
    },
    destroyed(){
      console.log('清除定时器：'+this.timer);
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
        if(!this.isIE9){
          let self = this;
          ws.createWebSocket(to,self,true);
          this.$root.$on(to,(e)=>{
            this.auto = e;
            let ball_sum = 0;
            for(let i=0;i<this.auto.ball.length;i++){
              ball_sum += Number(this.auto.ball[i])
            }
            this.auto.ball.push(ball_sum);
          });
          this.$root.$on(to+'lefttime',(e)=>{
            this.c_data.qishu = e.qishu;
            this.close_time.fengpan = e.close_time;
            this.close_time.now_time = e.now_time;
            let t1 = e.close_time - e.now_time;
            console.log('是否在维护中：'+this.is_wh);
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
      open_way: function() {
        let page =
          "trend_chart/chart-lotteryId=" +
          this.$route.query.page +
          ".html" +
          "?tab=1";
        window.open(page);
      },
      history: function() {
        let page =
          "trend_chart/chart-lotteryId=" +
          this.$route.query.page +
          ".html" +
          "?tab=3";
        window.open(page);
      },
      show_rule: function() {
        this.$root.$emit("rule_show", true);
        this.$root.$emit("now_page", this.$route.query.page);
      },
      getRTime: function() {
        this.close_time.now_time += 1;
        var t1 = this.close_time.fengpan * 1000 - this.close_time.now_time * 1000;
        this.h = Math.floor((t1 / 1000 / 60 / 60) % 24);
        this.m = Math.floor((t1 / 1000 / 60) % 60);
        this.s = Math.floor((t1 / 1000) % 60);
        if (this.h == 0 && this.m == 0 && this.s == 0) {
          this.fetchData(2);
          this.$root.$emit("success", true);
        } else if (this.h < 0 && this.m < 0 && this.s < 0) {
          if(this.timer){
              window.clearTimeout(this.timer);
              this.timer = null;
          }
          this.h = -1;
          this.m = -1;
          this.s = -1;
//          window.setTimeout(() => {
//            this.fetchData(2);
//          },60000)
        }
      },
      init: function() {
          this.getRTime();
          this.timer = window.setTimeout(this.init,1000); //time是指本身,延时递归调用自己,1000为间隔调用时间,单位毫秒
      },
      sortNumber: function(a, b) {
        return a.sort - b.sort;
      },
      go_child: function(child) {
        this.$router.push({ name: child, query: { page: this.$route.query.page } });
      },
      fetchData(type) {
        this.$root.$emit('wh_modal',false);
        if(type == 2){
          this.$root.$emit('loading',true,true);
        }else{
          this.$root.$emit('loading',true);
        }
        if(this.timer){
            window.clearTimeout(this.timer);
            this.timer = null;
        }
        let body = {
          fc_type: this.$route.query.page
        };
        api.getgameindex(this, body, res => {
          if (res.data.ErrorCode == 1) {
            if(res.data.is_wh == 2){
              this.$root.$emit('wh_modal',true);
              this.is_wh = true
            }else if(res.data.is_wh == 1){
              this.$root.$emit('wh_modal',false);
              this.is_wh = false
            }
            this.auto = res.data.Data.auto;
            let ball_sum = 0;
            for(let i=0;i<this.auto.ball.length;i++){
              ball_sum += Number(this.auto.ball[i])
            }
            this.auto.ball.push(ball_sum);
            this.close_time = res.data.Data.closetime;
            if(this.close_time.fengpan - this.close_time.now_time < 0){
                this.$root.$emit('wh_modal',true,true)
            }
            this.c_data = res.data.Data.c_data;
            var data = res.data.Data.odds;
            data.sort(this.sortNumber);
            this.computed(data);
            if(type == 2){
              window.setTimeout(() => {
                this.$root.$emit("loading", false);
              }, 1000)
            }else{
              this.$root.$emit("loading", false);
            }
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
        });
      },
      computed: function(data) {
        this.$set(this.luckyLists, this.luckyLists);
        this.luckyLists = [];
          this.luckyLists = luckyLists();
          var k = 0;
          for (let l = 0; l < this.luckyLists.length; l++) {
            for (let i = 0; i < 7; i++, k++) {
              Object.assign(this.luckyLists[l].object[i], data[k]);
              let name = data[k].remark.slice(
                data[k].remark.search("#") + 1,
                data[k].remark.length
              );
              this.luckyLists[l].object[i].name = name;
            }
          }
          for (let l = 0; l < this.luckyLists.length; l++) {
            k = 28 + l;
            for (let i = 7; i < this.luckyLists[l].object.length; i++, k += 4) {
              Object.assign(this.luckyLists[l].object[i], data[k]);
              let name = data[k].remark.slice(
                data[k].remark.search("#") + 1,
                data[k].remark.length
              );
              this.luckyLists[l].object[i].name = name;
            }
          }
      }
    },
  };
</script>
