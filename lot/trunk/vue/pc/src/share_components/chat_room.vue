<template>
    <div class="chat_box">
      <div class="new_num" :style="{width:newNum_wd}" v-show="new_num>0 && this.$route.query.page"><span v-if="new_num>99">99+</span><span v-else>{{new_num}}</span></div>
      <div @click="open_chat()" class="chart_button" v-show="this.$route.query.page">
        <i class="iconfont pk-liaotianshi"></i>
        <p class="chart_text">聊天室</p>
      </div>
      <div class="chat_body" :class="[zhankai_status?'zhan_kai':'wei_zhan']" v-show="show_chat && this.$route.query.page">
        <p class="top">
          <i class="iconfont pk-liaotianshi1"></i>
          <span class="top_text">当前房间(<span>{{now_page}}</span>)</span>
          <i @click="zhan_kai()" :class="[zhankai_status?'zhankai':'zhankai_defalut','pk-zhankai','iconfont']"></i>
          <i @click="close_chart()" class="iconfont pk-guanbi1"></i>
        </p>
        <div class="center" ref="center">
          <div class="Item type-left" v-for="item in user_msg">
            <div v-if="item.content != undefined" :class="[item.uid == 1?'me_lay-block':'lay-block']">
              <div class="avatar">
                <img :src="head_img" alt="管理员小马哥">
              </div>
              <div class="lay-content">
                <div class="msg-header">
                  <h4>{{item.username}}</h4>
                  <span class="MsgTime">{{item.sendtime}}</span>
                </div>
                <div class="Bubble">
                  <p><span class="msg" v-html="replaceFace(item.content)"></span></p>
                </div>
              </div>
            </div>
            <p v-else-if="item.content == undefined" class="join_exit">{{item.username}}{{item.cmd}}房间</p>
          </div>
        </div>
        <div class="bottom">
          <div class="bottom_top">
            <div class="emoji">
              <i :class="[showEmoji?'en_yl':'','iconfont','pk-biaoqing']" @click="showEmoji=!showEmoji"></i>
              <transition name="showbox">
                <div class="emojiBox" v-show="showEmoji">
                  <li v-for="(item, index) in emojis">
                    <img :src="emojs_img+item.file" :data="item.code" @click="push_emoji(item.code)">
                  </li>
                </div>
              </transition>
            </div>
          </div>
          <div class="bottom_bottom">
            <I-Input :class="[zhankai_status?'zhan_kai':'wei_zhan']" ref="msg_input" @keyup.enter.native="send()" v-model="message" :autosize="{minRows: 2,maxRows: 5}" placeholder="畅所欲言吧..."></I-Input>
            <I-Button type="primary" icon="android-send" @click="send()">发送</I-Button>
          </div>
        </div>
      </div>
    </div>
</template>
<script>
  import emojis from '../assets/js/emoji'
  import {Input,Button} from 'iview'
  import w_chat from '../assets/js/chat_socket'
  import mymousewheel from '../assets/js/mousewheel'
  import cm_cookie from '../assets/js/com_cookie'
  export default{
    components:{
      'I-Input':Input,
      'I-Button':Button,
    },
    data(){
      return{
        head_img:require('../assets/img/head_img.jpg'),
        emojs_img:'./static/emoji/',
        emojis:emojis,
        showEmoji:false,
        now_page:'',
        show_chat:false,
        zhankai_status:false,
        message:'',
        user_msg:[],
        new_num:0,
        newNum_wd:'18px'
      }
    },
    created(){
      this.$root.$on('get_fcName',(e)=>{
        this.now_page = e;
        cm_cookie.setCookie('now_page',e);
        console.log(e.length);
        console.log(e);
        if(e=='北京赛车pk拾'){
          this.now_page = '北赛pk拾'
        }else if(e=='广东快乐十分'){
          this.now_page = '广东快十'
        }else if(e=='重庆快乐十分'){
            this.now_page = '重庆快十'
        }else if(e.length >= 7) {
          this.now_page= this.now_page.substring(0,this.now_page.length-1)
        }
      });
//      this.$root.$on('chat_bet',(e)=>{
//        let ooo = [];
//        for(let i=0;i<e.length;i++){
//          ooo.push(e[i].remark.split('#')[0] + e[i].input_name+'号')
//        }
//        let aaa = '已投注：'+ooo;
//        console.log(JSON.stringify(e));
//        const bet_content = {
//          cmd: 'send',
//          fc_type: this.$route.query.page,
//          uid: 1, username: 'huygo', sendtime: this.now_time(), content: aaa
//        };
//        w_chat.send_ws(JSON.stringify(bet_content));
//      })
    },
    watch:{
      '$route.query.page':function(to,from) {
        this.user_msg = [];
        this.new_num = 0;
        if(!this.isIE9){
          let self = this;
          w_chat.createWebSocket(self.$route.query.page,self);
        }
      },
      new_num:function (new_val) {
        if(new_val < 100){
          this.newNum_wd = '18px'
        }else if(new_val > 99){
          this.newNum_wd = '25px';
        }
      },
    },
    mounted(){
      let center = document.querySelector(".center");
      mymousewheel(center);
      //创建websoket
      if(!this.isIE9){
        let self = this;
        w_chat.createWebSocket(self.$route.query.page,self);
        this.$root.$on('get_msgContent',(e,new_num)=>{
          if(e){
            if(e.cmd != undefined){
              if(e.cmd == 'join'){
                console.log('加入：'+e.data);
                e.data.cmd = '加入';
                this.user_msg.push(e.data)
              }else if(e.cmd == 'exit'){
                console.log('退出：'+e.data);
                e.data.cmd = '退出';
                this.user_msg.push(e.data);
              }
            }else{
              this.new_num = new_num;
              this.user_msg.push(e);
              this.message = ''
            }
            setTimeout(() => this.$refs.center.scrollTop = this.$refs.center.scrollHeight, 0);
          }
        });
      }
    },
    methods:{
      push_emoji: function (code) {
        this.message += code;
        this.$refs.msg_input.focus()
      },
      open_chat: function () {
        this.now_page = cm_cookie.getCookie('now_page');
        //聊天室切割标题
        if(this.now_page=='北京赛车pk拾'){
            this.now_page = '北赛pk拾'
        }else if( this.now_page=='广东快乐十分'){
            this.now_page = '广东快十'
        }else if( this.now_page=='重庆快乐十分'){
            this.now_page = '重庆快十'
        }else if(this.now_page.length >= 7) {
            this.now_page= this.now_page.substring(0,this.now_page.length-1)
        }
        console.log('dananndnansdna:'+this.now_page);
        if(this.now_page && !this.isIE9){
          this.show_chat = true;
          w_chat.clear_newNum();
          this.new_num = 0;
          setTimeout(() => this.$refs.center.scrollTop = this.$refs.center.scrollHeight, 0);
        }else if(this.now_page && this.isIE9){
          this.$Modal.warning(
            {content:'您的浏览器不支持websocket协议,建议使用新版谷歌、火狐等浏览器，请勿使用IE10以下浏览器，注意360浏览器不要使用兼容模式！'}
          )
        }else{
          this.$Message.warning({content:'正在加载中，暂不能打开。如等待太久，请刷新页面再试！',duration: 1,top: 100});
        }
      },
      send: function () {
        this.showEmoji = false;
        setTimeout(() => this.$refs.center.scrollTop = this.$refs.center.scrollHeight, 0);
        if(this.message != ""){
          const msg_content = {
            cmd: 'send',
            fc_type: this.$route.query.page,
            uid: 1, username: cm_cookie.getCookie('user_name'), sendtime: this.now_time(), content: this.message
          };
          w_chat.send_ws(JSON.stringify(msg_content));
//          w_chat.test();

        }else{//输入为空时的提示

        }
      },
      close_chart: function () {
        this.show_chat = false;
        this.new_num = 0;
      },
      zhan_kai: function () {
        this.zhankai_status = !this.zhankai_status
      },
      //  在发送信息之后，将输入的内容中属于表情的部分替换成emoji图片标签
      //  再经过v-html 渲染成真正的图片
      replaceFace: function(con) {
        if(con.includes('/:')) {
          var emojis=this.emojis;
          for(var i=0;i<emojis.length;i++){
            con = con.replace(emojis[i].reg, '<img src="./static/emoji/' + emojis[i].file +'" style="vertical-align: middle; width: 24px; height: 24px" />');
          }
          return con;
        }
        return con;
      },
      //处理时间
      now_time: function(){
        let date = new Date();
        let year=date.getFullYear();
        let month=date.getMonth()+1;
        let day = date.getDate();
        month =(month<10 ? "0"+month:month);
        day =(day<10 ? "0"+day:day);
        let mydate = (year.toString()+'-'+month.toString()+'-'+day.toString());
        if(date.getHours()<10 && date.getMinutes()>=10 && date.getMinutes()>=10){
          return '0' + date.getHours() + ':' + date.getMinutes() +':'+  date.getSeconds();
        }else if (date.getHours()>=10 && date.getMinutes()<10 && date.getSeconds()>=10){
          return date.getHours() + ':0' + date.getMinutes() +':'+  date.getSeconds();
        }else if (date.getHours()>=10 && date.getMinutes()>=10 && date.getSeconds()<10){
          return date.getHours() + ':' + date.getMinutes() +':0'+ date.getSeconds();
        }else if (date.getHours()<10 && date.getMinutes()<10 && date.getSeconds()>=10){
          return '0' + date.getHours() + ':0' + date.getMinutes() +':'+ date.getSeconds();
        }else if (date.getHours()<10 && date.getMinutes()>=10 && date.getSeconds()<10){
          return '0' + date.getHours() + ':' + date.getMinutes() +':0'+ date.getSeconds();
        }else if (date.getHours()>=10 && date.getMinutes()<10 && date.getSeconds()<10){
          return date.getHours() + ':0' + date.getMinutes() +':0'+ date.getSeconds();
        }else{
          return date.getHours() + ':' + date.getMinutes() +':'+ date.getSeconds();
        }
      }
    }
  }
</script>
<style lang="scss" src="../assets/css/chat_room.scss" scoped></style>
