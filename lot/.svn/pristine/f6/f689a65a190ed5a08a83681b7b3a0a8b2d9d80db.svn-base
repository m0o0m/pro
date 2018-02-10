<template lang="html">
  <div class="right_content">
    <ul v-for="(item,num) in mores">
      <!--<li class="three" v-if="key < 6">-->
        <!--<span ref="top_config" :style="[color==key.index?bg_color:'']" :class="[kk==1?'pd5':'pd10','top_config']" @click="config_num(key)" v-for="(key,kk) in item.arr">{{key.one}}</span>-->
      <!--</li>-->
      <!--<li class="four" v-if="key >= 6">-->
        <!--<span class="pd7 top_config" :style="[color==key.index?bg_color:'']" @click="config_num(key)" v-for="key in item.arr">{{key.four}}</span>-->
      <!--</li>-->
      <li class="top_style top_config" ref="top_config" v-if="num < 6" :style="[color==key.index?bg_color:'']" @click="config_num(key)" v-for="key in item.arr">{{key.one}}</li>
      <li class="li_style top_config" v-if="num >= 6" :style="[color==key.index?bg_color:'']" @click="config_num(key)" v-for="key in item.arr">{{key.four}}</li>
    </ul>
    <ul v-for="(item,key) in bottom_mores">
      <!--<li class="four">-->
        <!--<span class="pd7" :style="[key.state?bg_color:'']" @click="config_num1(key)" v-for="key in item.arr">{{key.four}}</span>-->
      <!--</li>-->
      <li class="li_style" :style="[key.state?bg_color:'']" @click="config_num1(key)" v-for="key in item.arr">{{key.four}}</li>
    </ul>
  </div>
</template>

<script>
export default {
  props:{
    mores:{
      type:Array
    },
    bottom_mores:{
      type:Array
    }
  },
  data(){
    return{
      color:null,
      bottom_color:false,
      bg_color:{
        color:'red',
        backgroundColor:'rgba(82, 210, 246, 0.16)'
      }
    }
  },
  methods:{
    config_num: function (item,key) {
      this.color = item.index;
      this.$emit('config_num',item)
    },
    config_num1: function (item,key) {
      item.state = !item.state;
      this.$emit('config_num1',item)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../../../assets/css/function.scss';
.right_content{
  width: 198px;
  float: left;
  ul{
    width: 160px;
    border: 1px solid $border_color;
    color: #999999;
    /*padding: 7px 10px;*/
    border-radius: 20px;
    overflow: hidden;
    margin: 0 auto;
    margin-bottom: 10px;
    .three{
      width: 100%;
      cursor: pointer;
      float: left;
      .pd5{
        padding: 5px;
      }
      .pd10{
        padding: 10px;
      }
    }
    .four{
      width: 100%;
      cursor: pointer;
      float: left;
      .pd7{
        padding: 7px;
      }
    }
    .top_style{
      cursor: pointer;
      width:33.3%;
      display:inline-block;
      float: left;
      padding:8px 0;
    }
    .top_style:hover{
      background-color: $bg_select;
      color:red;
    }
    .li_style{
      cursor: pointer;
      width:25%;
      display:inline-block;
      float: left;
      padding:8px 0;
    }
    .li_style:hover{
      background-color: $bg_select;
      color:red;
    }
  }
}
</style>
