<template>
  <div class="default_right">
    <div class="right_top">
      <div class="d_r_c">
        <div class="r_user">
          <div class="r_id">
            <span>账号:</span>
            <span class="s_txt">{{user_msg.user_name}}</span>
            <!--<span @mouseenter="enter" @mouseleave="leave" style="position: relative">hover_test<span v-show="show" style="position: absolute">qishu</span></span>-->
            <!--<span @click="send_huygo()">点我发射负数</span>-->
            <!--<span @click="send_huygo1()">点我发射正数</span>-->
            <!--<span @click="send_huygo2()">点我发射等数</span>-->
          </div>
          <div class="money">
            <span>余额:</span>
            <span v-show="!refech_now" class="s_txt">{{user_msg.money}}</span>
            <span v-show="refech_now" class="s_txt"><img src="../assets/img/refresh.gif" alt="刷新小图标"/></span>
            <span @click="upadte_money(1)" class="balance">
              <i class="iconfont pk-shuaxin">刷新</i>
            </span>
          </div>
        </div>
        <!--公告栏-->
        <div class="center_ri">
          <div class="laba_box fl">
            <i class="iconfont pk-laba laba"></i>
          </div>
          <div ref="demo" class="qimo8">
            <div class="qimo">
              <div ref="demo1" @mouseover="onmouseover()" @mouseout="onmouseout()">
                <ul>
                  <li>
                    <a ref="bb1" href='#' @click="message(item)" v-for="item in top_notice">{{item.content}}</a>
                  </li>
                </ul>
              </div>
              <div ref="demo2" @mouseover="onmouseover()" @mouseout="onmouseout()" v-show="demo2">
                <ul>
                  <li>
                    <a ref="bb2" href='#' @click="message(item)" v-for="item in top_notice">{{item.content}}</a>
                  </li>
                </ul>
              </div>
            </div>
            <!--<marquee scrollamount="6" scrolldelay="150" direction="left" @mouseover="onmouseover();" @mouseout="onmouseout();" style="cursor: pointer;">-->
              <!--<span id="bulletinMsg" data="left" onclick="message(item);" v-for="item in top_notice">{{item.content}}</span>-->
            <!--</marquee>-->
            <!--<text-scroll :dataList="top_notice" scrollType="scroll-up"></text-scroll>-->
          </div>
        </div>
      </div>
      <div class="d_r_r">
        <ul>
          <i class="i1">
            <!--<div class="msg_num" v-if="had_msg == 2"></div>-->
            <img src="../assets/img/massage.png" alt=""></i>
          <li @click="message()">消息中心</li>
          <i><img src="../assets/img/report.png" alt=""></i>
          <li @click="path()" :class="flag?'click':''">报表统计</li>
          <i><img src="../assets/img/record.png" alt=""></i>
          <li @click="betting()">投注记录</li>
          <!--<li class="out">退出登陆</li>-->
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
  import api from '../api/config'
  import cm_cookie from '../assets/js/com_cookie'
  export default {
    data() {
      return {
        show:false,

        demo2:false,
        refech_now: false,
        referch_gif:require('../assets/img/refresh.gif'),
        myvar: null,
        flag: false,
        //用户名or余额
        user_msg: {
          user_name: '',
          money: ''
        },
        //公告
        top_notice: [
          {content:''}
        ],
        a:0,
        b:5
      }
    },
    mounted() {
//        api.test(this, {}, (res) => {
//           console.log(123123);
//           console.log(res)
//        });

      this.$Message.config({
        top: 50,
        duration: 1
      });
      this.getajax();
      this.$root.$on('bet_success', (e) => {
        if (e) {
          this.upadte_money(2)
        }
      });
    },
    methods: {
//      enter: function(){
//          this.show = true
//      },
//      leave: function(){
//          this.show = false;
//      },
      send_huygo(){
        let test_data = {
          close_time:10,
          now_time:11
        };
        let num = String(this.$route.query.page+'lefttime');
        this.$root.$emit(num,test_data)
      },
      send_huygo1(){
        let test_data = {
          close_time:1516007380,
          now_time:1516007260
        };
        let num = String(this.$route.query.page+'lefttime');
        this.$root.$emit(num,test_data)
      },
//      send_huygo2(){
//          let test_data = {
//              close_time:10,
//              now_time:10
//          };
//          let num = String(this.$route.query.page+'lefttime');
//          this.$root.$emit(num,test_data)
//      },

      onmouseover(){
        window.clearInterval(this.myvar)
      },
      onmouseout(){
        this.scrollLeft()
      },
      scrollLeft() {
        let tab1 = this.$refs.demo1;
        this.demo2 = true;
        this.myvar = window.setInterval(this.out, 30);
        window.myvar = this.myvar;
      },
      out() {
        let tab = this.$refs.demo;
        let tab1 = this.$refs.demo1;
        let tab2 = this.$refs.demo2;
        if (tab2.offsetWidth-tab.scrollLeft<=0) {
          tab.scrollLeft -= tab1.offsetWidth;
        } else {
          tab.scrollLeft += 1;
        }
      },
      balance() {
        this.getajax();
      },
      message(item) {
        if(item){
          this.$router.push({
            name: 'notice',
            query:{id:item.id}
          })
        }else{
          this.$router.push({
            name: 'notice',
          })
        }
      },
      getajax() {
        this.refech_now = true;
        this.money_refech();
        api.notice(this, {}, (res) => {
          if (res.data.ErrorCode == 1) {
            console.log(res.data.Data);
            this.top_notice = res.data.Data;
            if(this.top_notice){
              this.scrollLeft();
            }
          }
        })
      },
      upadte_money: function (type) {
        if(type == 1){
          this.a += 1;
          if(this.a >= 2){
            this.$Message.warning({content:'您的操作太频繁,请'+this.b+'秒后再试！',duration: 1,top: 100});
          }else if(this.a == 1){
            this.b = 5;
            window.setTimeout(() => {
              this.a = 0;
            }, 5000);
            var iitmer = window.setInterval(()=>{
              this.b -= 1;
              console.log(this.b);
              if(this.b == 0){
                window.clearInterval(iitmer);
              }
            },1000);
            this.money_refech();
          }
        }else if(type == 2){
          this.money_refech()
        }
      },
      money_refech: function () {
        this.refech_now = true;
        api.userbalance(this, {}, (res) => {
          //console.log(res.data.Data);
          if (res.data.ErrorCode == 1) {
            this.user_msg.user_name = res.data.Data.uname;
            cm_cookie.setCookie('user_name',this.user_msg.user_name);
            this.user_msg.money = res.data.Data.money;
            this.refech_now = false
          }
        });
      },
      path() {
        this.$router.push({
          name: 'report'
        })
      },
      betting() {
        this.$router.push({
          name: 'record',
          params: {page: this.$route.query.page}
        })
      }
    },
  }
</script>

<style lang="scss" src="../assets/css/default_right.scss" scoped></style>
