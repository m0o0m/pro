<template lang="html">
  <div class="all_menu clearfix">
    <div class="menu_list" v-for="(item,i) in rightList">
      <div class="content">
        <div class="top_left">
          <img :src="item.img_path" alt="图标">
        </div>
        <div class="top_right">
          <h3>{{item.name}}</h3>
          <p>第 {{item.qishu}} 期</p>
        </div>
        <div class="center clearfix">
          <span class="center_left">近期开奖<br/>第{{item.auto_qishu}}期</span>
          <span v-if="item.ltype == 'k3'" class="center_right clearfix" >
            <span v-for="(key,k) in item.auto" :class="['k3','img'+key]">
              <i>
                {{key.number}}
              </i>
            </span>
          </span>
          <span v-else-if="item.ltype == 'yb'" class="center_right clearfix">
            <span v-for="(key,k) in item.auto" :class="[{'three':item.auto.length>14&&k>6&&k<14,'one':item.auto.length<9,'two':item.auto.length>8&&item.auto.length<17},key.color]">
              <i v-if="item.template == 'liuhecai'">
                {{key.number}}
              </i>
              <i v-else>
                {{key}}
              </i>
            </span>
          </span>
          <span v-else-if="item.ltype == 'gpc'" class="center_right clearfix">
            <span v-if="item.template != 'ffc_o'" v-for="(key,k) in item.auto" :class="[{'three':item.auto.length>14&&k>6&&k<14,'one':item.auto.length<9,'two':item.auto.length>8&&item.auto.length<17},key.color]">
              <i>
                {{key.number}}
              </i>
            </span>
            <span v-if="item.template == 'ffc_o'" v-for="(key,k) in item.auto" :class="[{'three':item.auto.length>14&&k>6&&k<14,'one':item.auto.length<9,'two':item.auto.length>8&&item.auto.length<17},key.color]">
              <i>
                {{key}}
              </i>
            </span>
          </span>
          <span v-else-if="item.ltype == 'xy'" class="center_right clearfix">
            <span v-for="(key,k) in item.auto" :class="[{'three':item.auto.length>14&&k>6&&k<14,'one':item.auto.length<9,'two':item.auto.length>8&&item.auto.length<17},'xyc_div']">
              <i class="boll">
                {{key}}
              </i>
              <i v-if="k<2">+</i>
              <i v-else-if="k==2">=</i>
            </span>
          </span>
          <span v-else class="center_right clearfix">
            <span v-for="(key,k) in item.auto" :class="{'three':item.auto.length>14&&k>6&&k<14,'one':item.auto.length<9,'two':item.auto.length>8&&item.auto.length<17}">
              <i>
                {{key}}
              </i>
            </span>
          </span>
        </div>
        <div class="bottom">
          <p class="bottom_one" @click="history(item)">历史开奖</p>
          <p class="bottom_two" @click="open_way(item)">开奖走势</p>
          <p class="bottom_three" @click="go(item)">立即投注</p>
        </div>
      </div>
      <!-- 维护遮罩 -->
      <div v-if="item.is_wh == 2" class="weihu">
        <img class="left_top" :src="item.img_path" alt="">
        <img class="weihu_logo" :src="weihu" alt="weihu_logo">
        <p>此彩票正在维护，请稍后再试</p>
      </div>
    </div>
    <Loading :loading_modal="loading_modal"></Loading>
  </div>
</template>

<script>
import Loading from '../share_components/loading'
export default {
  data() {
    return {
      weihu:require('../assets/img/weihu.png'),
//      loading_modal: true
    };
  },
  components: {
    Loading
  },
  mounted(){
//    this.$root.$on('loading_home',(e)=>{
//      this.loading_modal = e
//    })
  },
  methods: {
    open_way: function (v){
      let page = 'trend_chart/chart-lotteryId='+v.type+'.html'+'?tab=1';
      window.open(page)
    },
    history: function (v) {
      let page = 'trend_chart/chart-lotteryId='+v.type+'.html'+'?tab=3';
      window.open(page)
    },
    go: function(v) {
      this.$root.$emit("change_item", 0);
      this.$root.$emit("child_change", 0);
      let index_page = '';
      index_page = v.template;
      this.$router.push({ name: index_page, query: { page: v.type } });
    },
    menu_get_love(item) {
      this.$root.$emit("menu_love", item);
    }
  },
  props: {
    rightList:{
      type:Array
    },
    loading_modal:{
      type:Boolean
    }
  }
};
</script>
<style lang="scss" src="../assets/css/all_menu.scss" scoped></style>
