<template lang="html">
  <div class="pk_content">
    <Loading :loading_modal="loading_modal" :go_next='go_next'></Loading>
    <Wh :weihu_modal="weihu_modal" :time_fuck="time_fuck"></Wh>
    <router-view></router-view>
  </div>
</template>

<script>
import Loading from "../share_components/loading";
import Wh from "../share_components/weihu_modal";
export default {
  components:{
    Loading,
    Wh
  },
  data(){
    return{
      go_next: false,
      loading_modal: false,
      weihu_modal: false,
      time_fuck: false
    }
  },
  created(){
    this.$root.$on("loading",(e,k) => {
      this.loading_modal = e;
      this.go_next = k;
    });
    this.$root.$on("wh_modal",(e,k)=>{
      this.weihu_modal = e;
      this.time_fuck = k;
    })
  }
}
</script>
<style lang="scss" src="../assets/css/pk_content.scss" scoped></style>
