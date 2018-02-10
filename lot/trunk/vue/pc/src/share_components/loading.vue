<template lang="html">
  <div class="" v-show="loading_modal">
    <div class="modal_content">
      <h2 style="color:#FFCC00" v-show="go_next">即将进入下一期</h2>
      <vue-loading type="bars" color="red" :size="{width: '50px', height: '50px'}"></vue-loading>
    </div>
    <div class="bg_modal">
    </div>
  </div>
</template>

<script>
import vueLoading from 'vue-loading-template'
export default {
   components:{vueLoading},
   props:{
     loading_modal:{
       type:Boolean,
       default: false
     },
     go_next:{
       type:Boolean,
       default: false
     },
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
   .modal_content{
     position:absolute;top:30%;left:45%;z-index:98
   }
   .bg_modal{
     position: absolute;
     top: 0;
     bottom: 0;
     left: 0;
     right: 0;
     border-radius:5px;
     background-color: rgba(55,55,55,.6);
     opacity: 0.5;
     width:100%;
     height: 100%;
     z-index: 80;
   }
 </style>
