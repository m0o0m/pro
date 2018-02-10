<template lang="html">
  <div class="" v-show="weihu_modal">
    <div class="" style="position:absolute;top:30%;left:40%;z-index:98">
      <p class="wh_text" v-if="!time_fuck">正在维护中...</p>
      <p class="wh_text" v-if="time_fuck">封盘中...</p>
    </div>
    <div style="visibility: visible;
      position: absolute;
      left: 0px;
      top: 0px;
      z-index: 80;
      background: gray;
      zoom: 1;
      border-radius:5px;
      filter: alpha(opacity=80);
      opacity: 0.5;
      width: 100%;
      height: 100%;">
    </div>
  </div>
</template>

<script>
  import vueLoading from 'vue-loading-template'
  export default {
    components:{vueLoading},
    props:{
      weihu_modal:{
        type:Boolean,
        default: false
      },
      time_fuck:{
        type:Boolean,
        default: false
      }
    },
    data(){
      return{}
    },
    created(){},
    mounted(){},
    methods:{},
  }
</script>
 <style lang="scss" scoped>
   .wh_text{
     font-size: 30px;
     color: #eeef6f;
   }
 </style>
