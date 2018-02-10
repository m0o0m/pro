<template>
  <div class="happy">
    <!--<iTime :fect_data="status" :left_data="c_data" :center_data="close_time" :bottom_data="auto"></iTime>-->
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
    <Nav-top :isActive='false' :lists="menus" v-on:menu="go_child"></Nav-top>
    <router-view :one="one" :two="two" :c_data="c_data" :only_data="only_data"></router-view>
  </div>
</template>

<script>
import api from "../../../api/config"
import NavTop from '../../../share_components/default_nav'
import ws from '../../../assets/js/socket'
import cm_cookie from '../../../assets/js/com_cookie'
export default {
  components:{
    NavTop,
  },
  data(){
    return{
      timer: null,
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
      only_data:[],
      one: [
        {
          name: '和值', object: [
          {num: '总和大', number: 165.000, money: '',id:0,index:0,flag:false},
          {num: '总和小', number: 165.000, money: '',id:1,index:1,flag:false},
          {num: '总和单', number: 165.000, money: '',id:2,index:2,flag:false},
          {num: '总和双', number: 165.000, money: '',id:3,index:3,flag:false},
          {num: '总和810', number: 165.000, money: '',id:4,index:4,flag:false},
        ]
        },
        {
          name: '上中下', object: [
          {num: '上盘', number: 165.000, money: '',id:5,index:5,flag:false},
          {num: '中盘', number: 165.000, money: '',id:6,index:6,flag:false},
          {num: '下盘', number: 165.000, money: '',id:7,index:7,flag:false},
        ]
        },
        {
          name: '奇和偶', object: [
          {num: '奇盘', number: 165.000, money: '',id:8,index:8,flag:false},
          {num: '和盘', number: 165.000, money: '',id:9,index:9,flag:false},
          {num: '偶盘', number: 165.000, money: '',id:10,index:10,flag:false},
        ]
        },
      ],
      two:[
        {
          name:'选一',txt:'3',index:0,object:[
          {num:'1',id:0,index:0,state:false,money:''},
          {num:'2',id:1,index:1,state:false,money:''},
          {num:'3',id:2,index:2,state:false,money:''},
          {num:'4',id:3,index:3,state:false,money:''},
          {num:'5',id:4,index:4,state:false,money:''},
          {num:'6',id:5,index:5,state:false,money:''},
          {num:'7',id:6,index:6,state:false,money:''},
          {num:'8',id:7,index:7,state:false,money:''},
          {num:'9',id:8,index:8,state:false,money:''},
          {num:'10',id:9,index:9,state:false,money:''},
          {num:'11',id:10,index:10,state:false,money:''},
          {num:'12',id:11,index:11,state:false,money:''},
          {num:'13',id:12,index:12,state:false,money:''},
          {num:'14',id:13,index:13,state:false,money:''},
          {num:'15',id:14,index:14,state:false,money:''},
          {num:'16',id:15,index:15,state:false,money:''},
        ]
        },
        {
          name:'选二',txt:'2/2:10',index:1,object:[
          {num:'17',id:16,index:16,state:false,money:''},
          {num:'18',id:17,index:17,state:false,money:''},
          {num:'19',id:18,index:18,state:false,money:''},
          {num:'20',id:19,index:19,state:false,money:''},
          {num:'21',id:20,index:20,state:false,money:''},
          {num:'22',id:21,index:21,state:false,money:''},
          {num:'23',id:22,index:22,state:false,money:''},
          {num:'24',id:23,index:23,state:false,money:''},
          {num:'25',id:24,index:24,state:false,money:''},
          {num:'26',id:25,index:25,state:false,money:''},
          {num:'27',id:26,index:26,state:false,money:''},
          {num:'28',id:27,index:27,state:false,money:''},
          {num:'29',id:28,index:28,state:false,money:''},
          {num:'30',id:29,index:29,state:false,money:''},
          {num:'31',id:30,index:30,state:false,money:''},
          {num:'32',id:31,index:31,state:false,money:''},
        ]
        },
        {
          name:'选三',txt:'3/3:12',index:2,object:[
          {num:'33',id:32,index:32,state:false,money:''},
          {num:'34',id:33,index:33,state:false,money:''},
          {num:'35',id:34,index:34,state:false,money:''},
          {num:'36',id:35,index:35,state:false,money:''},
          {num:'37',id:36,index:36,state:false,money:''},
          {num:'38',id:37,index:37,state:false,money:''},
          {num:'39',id:38,index:38,state:false,money:''},
          {num:'40',id:39,index:39,state:false,money:''},
          {num:'41',id:40,index:40,state:false,money:''},
          {num:'42',id:41,index:41,state:false,money:''},
          {num:'43',id:42,index:42,state:false,money:''},
          {num:'44',id:43,index:43,state:false,money:''},
          {num:'45',id:44,index:44,state:false,money:''},
          {num:'46',id:45,index:45,state:false,money:''},
          {num:'47',id:46,index:46,state:false,money:''},
          {num:'48',id:47,index:47,state:false,money:''},
        ]
        },
        {
          name:'选四',txt:'4/4:50、4/3:5、4/2:3',index:3,object:[
          {num:'49',id:48,index:48,state:false,money:''},
          {num:'50',id:49,index:49,state:false,money:''},
          {num:'51',id:50,index:50,state:false,money:''},
          {num:'52',id:51,index:51,state:false,money:''},
          {num:'53',id:52,index:52,state:false,money:''},
          {num:'54',id:53,index:53,state:false,money:''},
          {num:'55',id:54,index:54,state:false,money:''},
          {num:'56',id:55,index:55,state:false,money:''},
          {num:'57',id:56,index:56,state:false,money:''},
          {num:'58',id:57,index:57,state:false,money:''},
          {num:'59',id:58,index:58,state:false,money:''},
          {num:'60',id:59,index:59,state:false,money:''},
          {num:'61',id:60,index:60,state:false,money:''},
          {num:'62',id:61,index:61,state:false,money:''},
          {num:'63',id:62,index:62,state:false,money:''},
          {num:'64',id:63,index:63,state:false,money:''},
        ]
        },
        {
          name:'选五',txt:'5/5:250、5/4:20、5/3:5',index:4,object:[
          {num:'65',id:64,index:64,state:false,money:''},
          {num:'66',id:65,index:65,state:false,money:''},
          {num:'67',id:66,index:66,state:false,money:''},
          {num:'68',id:67,index:67,state:false,money:''},
          {num:'69',id:68,index:68,state:false,money:''},
          {num:'70',id:69,index:69,state:false,money:''},
          {num:'71',id:70,index:70,state:false,money:''},
          {num:'72',id:71,index:71,state:false,money:''},
          {num:'73',id:72,index:72,state:false,money:''},
          {num:'74',id:73,index:73,state:false,money:''},
          {num:'75',id:74,index:74,state:false,money:''},
          {num:'76',id:75,index:75,state:false,money:''},
          {num:'77',id:76,index:76,state:false,money:''},
          {num:'78',id:77,index:77,state:false,money:''},
          {num:'79',id:78,index:78,state:false,money:''},
          {num:'80',id:79,index:79,state:false,money:''},
        ]
        },
      ],
      menus: [
        {name:'选1-5',item:'bj_kl8'},
        {name:'整合',item:'or'}
      ],
      routePage: null,
      is_wh: false,
    }
  },
  created(){
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
      this.fetchData();
      this.socket_change(to);
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
      });
        this.$root.$on(to+'lefttime',(e)=>{
          console.log(e);
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
//        window.setTimeout(() => {
//          this.fetchData(2);
//        },60000)
      }
    },
    init: function() {
        this.getRTime();
        this.timer = window.setTimeout(this.init,1000); //time是指本身,延时递归调用自己,1000为间隔调用时间,单位毫秒
    },
    sortNumber: function(a,b){
      return a.sort - b.sort
    },
    fetchData: function(type){
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
        'fc_type': this.$route.query.page
      };
      api.getgameindex(this, body, (res) => {
        if (res.data.ErrorCode == 1) {
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
          }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
          // console.log('success');
          if(res.data.is_wh == 2){
            this.$root.$emit('wh_modal',true);
            this.is_wh = true
          }else if(res.data.is_wh == 1){
            this.$root.$emit('wh_modal',false);
            this.is_wh = false
          }
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          // console.log(back_data);
          this.computed(back_data);
          this.computed_or(back_data);
          this.auto = res.data.Data.auto;
          this.close_time = res.data.Data.closetime;
          if(this.close_time.fengpan - this.close_time.now_time < 0){
              this.$root.$emit('wh_modal',true,true)
          }
          this.c_data = res.data.Data.c_data;
          this.only_data = res.data.Data.odds;
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
    },
    computed: function(data) {
      this.$set(this.one, this.one);
      for (let j = 0,i = 5;i<this.one[0].object.length,j < this.one[0].object.length;i++,j++) {
        Object.assign(this.one[0].object[j], data[i]);
        let name = data[i].remark.slice(
          data[i].remark.search("#") + 1,
          data[i].remark.length
        );
        console.log(name);
        this.one[0].object[j].num = name;
      }
      for (let l = 10, k = 0; l < this.one[1].object.length, k < this.one[1].object.length; l++, k++) {
        Object.assign(this.one[1].object[k], data[l]);
        let name = data[l].remark.slice(
          data[l].remark.search("#") + 1,
          data[l].remark.length
        );
//        console.log(name);
        this.one[1].object[k].num = name;
      }
      for (let n = 13, m = 0; n < this.one[2].object.length, m < this.one[2].object.length; n++, m++) {
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
      //get_two
      this.$set(this.two, this.two);
      Object.assign(this.two[0], data[0]);
      Object.assign(this.two[1], data[1]);
      Object.assign(this.two[2], data[2]);
      Object.assign(this.two[3], data[3]);
      Object.assign(this.two[4], data[4]);
      for(let i = 0;i < this.two.length;i++){
        for(let j = 0;j < this.two[i].object.length;j++){
          Object.assign(this.two[i].object[j],data[0]);
        }
      }
    },
    go_child: function(child){
      // console.log(child);
      this.$router.push({
        name:child,
        query: {page: this.$route.query.page}
      })
    }
  }
}
</script>

