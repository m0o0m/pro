<template lang="html">
  <!-- <div class="error">
    <h3 class="title">error:{{errorCode}}</h3>
    <div class="">
      服务器出现异常
    </div>
  </div> -->
  <div class="wrap">
		<div class="logo">
			<h1>{{errorCode}}</h1>
			<p>抱歉，服务器出现异常!</p>
			<div class="sub">
			   <p><a @click="back_page()" href="javascript:;">返回</a></p>
			   <p><a @click="back_home()" href="javascript:;">返回首页</a></p>
			</div>
		</div>
    <div class="footer">

  	</div>
	</div>
</template>

<script>
import {Modal} from 'iview';
export default {
  components: {Modal},
  props: {},
  data() {
    return {
      errorCode:500
    }
  },
  created() {
    window.scrollTo(0, 0)
  },
  mounted() {
    this.$Modal.remove();
    this.$root.$emit('loading',false);
  },
  methods: {
    back_home: function(){
      this.$router.push({name:'home'})
    },
    back_page: function(){
      this.$router.go(-1)
    }
  },
}
</script>
<style scoped lang="css">
  /*// .error{
  //   .title{
  //     margin: 0 auto;
  //     font-weight: bold;
  //     font-size: 30px
  //   }
  // }*/
  .wrap{
  	margin:0 auto;
  	width:100%;
    margin-top:100px;
    /*background:#DAD6CC;*/
  }
  .logo h1{
  	font-size:200px;
  	color:#FF7A00;
  	text-align:center;
  	margin-bottom:1px;
  	text-shadow:4px 4px 1px white;
  }
  .logo p{
  	color:#B1A18D;;
  	font-size:20px;
  	margin-top:1px;
  	text-align:center;
  }
  .logo p span{
  	color:lightgreen;
  }
  .sub a{
  	color:#ff7a00;
  	text-decoration:none;
  	padding:5px;
  	font-size:13px;
  	font-family: arial, serif;
  	font-weight:bold;
  }
  .footer{
  	color:white;
  	position:absolute;
  	right:10px;
  	bottom:10px;
  }
  .footer a{
  	color:#ff7a00;
  }
</style>
