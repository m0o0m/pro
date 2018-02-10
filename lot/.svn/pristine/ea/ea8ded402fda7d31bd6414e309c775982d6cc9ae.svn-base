<template lang="html">
  <div class="mark_six">
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
        <div class='bottom_center6'>
          <div class="bottom_content">
            <div class="content_list" v-for="item in auto.ball">
              <p :class="[item.color,'box']">{{item.number}}</p>
              <p class="animal">{{item.animal}}</p>
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
    <Nav-top :lists="menus" @menu="go_child" :isActive="margin"></Nav-top>
    <router-view></router-view>
  </div>
</template>

<script type="text/ecmascript-6">
import api from "../../../api/config";
import NavTop from "../../../share_components/default_nav";
import ws from '../../../assets/js/socket'
import cm_cookie from '../../../assets/js/com_cookie'
export default {
  components: {
    NavTop
  },
  data() {
    return {
      margin: false,
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
      menus: [
        {
          name: "特码",
          item: "liuhecai"
        },
        {
          name: "正码",
          item: "six_two"
        },
        {
          name: "正码特",
          item: "six_three"
        },
        {
          name: "正码1-6",
          item: "six_four"
        },
        {
          name: "过关",
          item: "six_five"
        },
        {
          name: "连码",
          item: "six_six"
        },
        {
          name: "半波",
          item: "six_seven"
        },
        {
          name: "一肖/尾数",
          item: "six_eight"
        },
        {
          name: "特码生肖",
          item: "six_nine"
        },
        {
          name: "合肖",
          item: "six_ten"
        },
        {
          name: "生肖连",
          item: "six_eleven"
        },
        {
          name: "尾数连",
          item: "six_twelve"
        },
        {
          name: "全不中",
          item: "six_thirteen"
        },
        {
          name: "五行",
          item: "six_row"
        },
        {
          name: "正肖",
          item: "positiveshaw"
        },
        {
          name: "特码头",
          item: "specialhead"
        },
        {
          name: "七码",
          item: "sevencode"
        },
        {
          name: "总肖",
          item: "totalshaw"
        }
      ],
      is_wh: false
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route.query.page':function(to,from) {
      this.$root.$off(from);
      this.$root.$off(from+'lefttime');
      this.$root.$off('only_back');
      this.fetchData();
      this.socket_change(to);
    }
  },
  mounted(){
    this.socket_change(this.$route.query.page);
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
    this.$root.$off('only_back');
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
        });
        this.$root.$on(to+'lefttime',(e)=>{
          this.c_data.qishu = e.qishu;
          this.close_time.fengpan = e.close_time;
          this.close_time.now_time = e.now_time;
          let t1 = e.close_time - e.now_time;
          console.log('是否在维护中：'+this.is_wh);
//          if(t1 == 0 && !this.is_wh){
//            this.$root.$emit("time_out",true)
//          }
//          if(this.m == -1 && !this.is_wh){
//            if(t1 > 0){
//              this.init();
//              this.$root.$emit('wh_modal',false);
//            }
//          }
          if(t1 == 0 && !this.is_wh){
              this.$root.$emit("time_out",true)
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
    go_child: function(child) {
      console.log(child);
      this.$router.push({ name: child, query: { page: this.$route.query.page } });
    },
    fetchData: function() {
      this.$root.$emit('wh_modal',false);
      let index = 0;
      if(this.timer){
          window.clearTimeout(this.timer);
          this.timer = null;
      }
      this.$root.$on('only_back',(e,type)=>{
        index+=1;
        console.log('传过来的次数：'+index);
        if (e.data.ErrorCode == 1) {
          this.$root.$emit('get_fcName', e.data.Data.c_data.fc_name);
          if(e.data.is_wh == 2){
            this.$root.$emit('wh_modal',true);
            this.is_wh = true
          }else if(e.data.is_wh == 1){
            this.$root.$emit('wh_modal',false);
            this.is_wh = false
          }
          this.auto = e.data.Data.auto;
          this.c_data = e.data.Data.c_data;
          this.close_time = e.data.Data.closetime;
          if(this.close_time.fengpan - this.close_time.now_time < 0){
              this.$root.$emit('wh_modal',true,true)
          }
          if(!this.is_wh){
              if(index == 1) {
                  if (this.close_time.fengpan) {
                      if (this.timer) {
                          window.clearTimeout(this.timer);
                          this.timer = null;
                      }
                      if (type == 2) {
                          window.setTimeout(() => {
                              this.init();
                          }, 1000)
                      } else {
                          this.init();
                      }
                  }
              }
            }
        }
      })
    },
    //aaaaa
    show_rule: function() {
      this.$root.$emit("rule_show", true);
      this.$root.$emit("now_page", this.$route.query.page);
    },
    getRTime: function() {
      this.close_time.now_time += 1;
      var t1 = this.close_time.fengpan * 1000 - this.close_time.now_time * 1000;
      var d= Math.floor(t1/1000/60/60/24);
      this.h = Math.floor((t1 / 1000 / 60 / 60) % 24 + d*24);
      this.m = Math.floor((t1 / 1000 / 60) % 60);
      this.s = Math.floor((t1 / 1000) % 60);
      console.log('单前时间h：'+this.s);
      if (this.h == 0 && this.m == 0 && this.s == 0) {
        this.$root.$emit("success",true);
        this.$root.$emit("time_out",true);
      }else if(this.h < 0 && this.m < 0 && this.s < 0){
        if(this.timer){
            window.clearTimeout(this.timer);
            this.timer = null;
        }
        console.log('时：'+this.h+'分：'+this.m+'秒：'+this.s);
        this.h = -1;
        this.m = -1;
        this.s = -1;
      }
    },
    init: function() {
        this.getRTime();
        this.timer = window.setTimeout(this.init,1000); //time是指本身,延时递归调用自己,1000为间隔调用时间,单位毫秒
    },
  }
};
</script>
<style lang="scss" src="../../../assets/css/mark_six.scss" scoped></style>
