<template>
    <div class="wh_body">
      <div class="wh_center">
        <div class="wh_left"><img :src="wh_img" alt="维护中的图片"></div>
        <div class="wh_right">
          <h1 class="one">网站维护中...</h1>
          <h4 class="two">Web site maintenance...</h4>
          <h4>很抱歉网站目前正在维护中，维护原因如下：</h4>
          <p>1.asdasdasdasdasd</p>
          <h4>维护时间：<span>{{wh_time}}</span>，请稍后再访问，感谢您的支持！如有其它问题，</h4>
          <h4>请联系客服 <span class="call_skype" @click="open_skype">000-000-000</span> 或发送邮件至 <span>asdasd@hotmail.com</span></h4>
          <h2 class="three">感谢您一直以来，一如既往的支持！！！</h2>
        </div>
      </div>
      <div class="footer">
        <p>Copyright © 欢迎来到PK家族 Reserved</p>
      </div>
    </div>
</template>
<script>
  import api from '../api/config'
  export default{
    data(){
      return{
        wh_img:require('../assets/img/wei_huA.png'),
        wh_time:'8:00-12:00'
      }
    },
    created(){
//      console.log(window.myvar);
      window.clearInterval(window.myvar);
      api.maintain(this,{},(e)=>{
//        console.log(e);
        if(e.data.ErrorCode == 1){
          if(e.data.pc == 1){
            this.$router.push({name:'home'})
          }
        }
      })
    },
    methods:{
      open_skype:function(){
        window.open('skype:pkbetcc?chat')
      }
    }
  }
</script>
<style lang="scss" scoped>
  .wh_body{
    background-color:#fff;
    width:100%;
    .wh_center{
      width:60%;
      margin:0 auto;
      overflow:hidden;
      margin-top:200px;
      .wh_left{
        float: left;
        img{
          width:400px;
        }
      }
      .wh_right{
        float: left;
        text-align:left;
        margin-left:20px;
        .call_skype{
          cursor: pointer;
        }
        h4,p{
          margin-bottom:5px;
          span{
            color: #ee8600;
          }
        }
        .one{
          color: #ee8600;
        }
        .two{
          margin-bottom: 15px;
        }
        .three{
          margin-top: 15px;
        }
      }
    }
    .footer{
      position: absolute;
      bottom:0;
      width:100%;
      p{
        width:50%;
        margin:0 auto;
      }
    }
  }

</style>
