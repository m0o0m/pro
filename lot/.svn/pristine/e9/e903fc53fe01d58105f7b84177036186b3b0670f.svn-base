<template lang="html">
  <div class="pk_time">
    <div class="time_left">
      <div class="left_logo">
        <img :src="left_data.img_path" alt="">
      </div>
      <div class="left_text">
        <p>{{left_data.fc_name}}</p>
        <p>第 {{left_data.qishu}} 期></p>
      </div>
    </div>
    <div class="time_center">
      <p class="center_top">投注剩余时间</p>
      <div class="center_bottom">
        <div class="fl time_content" v-if="h < 10">0{{h}}</div>
        <div class="fl time_content" v-else>{{h}}</div>
        <div class="fl fs">时</div>
        <div class="fl time_content" v-if="m < 10">0{{m}}</div>
        <div class="fl time_content" v-else>{{m}}</div>
        <div class="fl fs">分</div>
        <div class="fl time_content" v-if="s < 10">0{{s}}</div>
        <div class="fl time_content" v-else>{{s}}</div>
        <div class="fl fs">秒</div>
      </div>
    </div>
    <div class="time_bottom">
      <div class="bottom_left">第{{bottom_data.qishu}}期开奖</div>
      <div v-if="this.$route.query.page == 'liuhecai'" class='bottom_center6'>
        <div class="bottom_content">
          <div class="content_list" v-for="item in bottom_data.ball">
            <p :class="[item.color,'box']">{{item.number}}</p>
            <p class="animal">{{item.animal}}</p>
          </div>
        </div>
      </div>
      <div v-else :class="[bottom_data.ball.length >= 10?'bottom_center_other':'bottom_center']">
        <div :class="[bottom_data.ball.length == 10?'other_content':'bottom_content']">
          <div class="content_list" v-for="item in bottom_data.ball">
            <p class="box blue">{{item}}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="bottom_bottom">
      <p>历史结果</p>
      <p>开奖走势</p>
      <p @click="show_rule()">玩法规则</p>
    </div>
  </div>
</template>

<script>
  import {Modal} from 'iview';
  import api from '../api/config'
  export default {
    data() {
      return {
        // page:'',
        timer: null,
        // endTime: '2018/1/11 10:00:00',
        h: 0,
        m: 0,
        s: 0,
      }
    },
    components: {Modal},
    props:{
      left_data:{
        type:Object
      },
      center_data:{
        type:Object
      },
      bottom_data:{
        type:Object
      },
      status:{
        type:Boolean
      }
    },
    created() {
      //进入home 页面把定时器关闭
      this.$root.$on('now_home',(e)=>{
        if(e){
          for(var i = 0; i < 9999; i++) {
            window.clearInterval(i)
          }
        }
      });
    },
    watch: {
      // 如果路由有变化，会再次执行该方法
//     '$route.query.page': 'init' // 只有这个页面初始化之后，这个监听事件才开始生效
    },
    mounted() {
      this.$root.$emit('time_out', false);
      this.init()
    },
    methods: {
      show_rule: function(){
        this.$root.$emit('rule_show',true);
        this.$root.$emit('now_page',this.$route.query.page)
      },
      run_time: function() {
        this.timer = window.setInterval(this.getRTime, 1000);
        console.log('定时器id（run-time）：'+this.timer);
      },
      getRTime: function() {
        // var EndTime= new Date(); //截止时间
        // var NowTime = new Date();//现在的时间
        console.log('封盘时间：'+this.center_data.fengpan);
        console.log('当前时间：'+this.center_data.now_time);
        this.center_data.now_time += 1;
        var t1 = this.center_data.fengpan * 1000 - this.center_data.now_time * 1000;
        var t2 = this.center_data.kaijiang * 1000 - this.center_data.now_time * 1000;
        // var d=Math.floor(t/1000/60/60/24);
        this.h = Math.floor(t1 / 1000 / 60 / 60 % 24);
        this.m = Math.floor(t1 / 1000 / 60 % 60);
        this.s = Math.floor(t1 / 1000 % 60);
        // console.log('开奖时间：'+'时：'+h+'；分：'+m+'；秒：'+s);
        console.log('封盘时间：' + '时：' + this.h + '；分：' + this.m + '；秒：' + this.s);
        console.log('定时器id(time_out)：'+this.timer);
        if(this.h < 0 && this.m < 0 && this.s < 0){
          this.$root.$emit('time_out', false);
//          this.$root.$emit('success', true);
          this.$Modal.warning({
            content: '获取的时间为负值',
            onOk: () => {
              this.$router.push({name:'error'})
            },
          });
          window.clearInterval(this.timer);
          return
        }else if(this.h == 0 && this.m == 0 && this.s == 0){
          if(this.status){
            console.log('请求成功！');
            this.$root.$emit('time_out', false);
            this.$root.$off('time_out');
//            this.init();
          }else{
            this.$root.$emit('time_out', true);
            this.$root.$emit('success', true);
          }
        }
      },
      init: function() {
        console.log('定时器id(init)：'+this.timer);
        for(var i = 0; i < 9999; i++) {
          window.clearInterval(i)
        }
        this.run_time()
      },
    }
  }
</script>
<style lang="scss" src="../assets/css/pk_time.scss" scoped></style>
