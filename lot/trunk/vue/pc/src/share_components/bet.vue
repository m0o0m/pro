<template lang="html">
  <div class="bet" v-show="modal">
    <div class="pk_bet">
    </div>
    <div class="modal">
      <div class="header">
        <h3>下注金额</h3>
      </div>
        <p>当前投注：</p>
        <div class="bet_center">
          <p class="center_top">
            <span class="one">类型</span>
            <span class="two">号码</span>
            <span class="three">赔率</span>
            <span class="four">金额</span>
          </p>
          <div class="bet_modal" style="overflow:auto;height: 145px;">
            <p class="center_content" v-for="item in lists" v-if="item.money">
              <span class="one">{{item.remark.split('#')[0]}}</span>
              <span class="two" v-if="!isNaN(item.remark.split('#')[1])"><i class="box">{{item.remark.split('#')[1]}}</i></span>
              <span v-else class="two">{{item.remark.split('#')[1]}}</span>
              <span class="three">{{item.odd}}</span>
              <span class="four">{{item.money}}</span>
            </p>
          </div>
        </div>
        <div style="text-align: center;padding: 10px 0">
          <span style="margin-right:10px">共{{number}}注</span>
          <span>金额合计：{{money}}</span>
        </div>
        <div style="padding: 10px 0">
          <button type="button" v-if="!loading" class="bet_bottom color1" @click="pour()">下注</button>
          <button type="button" v-else class="bet_bottom color1">下单中...</button>
          <button type="button" class="bet_bottom color2 font" @click="cancel()">取消</button>
        </div>
    </div>
  </div>
</template>

<script>
import api from '../api/config'
import share from './share'
import mymousewheel from '../assets/js/mousewheel'
import {Modal} from 'iview';
import '../assets/css/bet.scss';
export default {
  components: {Modal},
  props:{
    modal:{
      type:Boolean,
      default: false
    }
  },
  data(){
    return{
      number:'',//注单量
      loading:false,
      lists:[],
      arr:[],
      qishu:'',
      money: 0
    }
  },
  created() {
    this.$root.$on('id-selected', (e) => {
      // console.log(e)
      for(let i = 0;i< e.length;i++){
        for(let j = 0;j<e[i].object.length;j++){
          //添加金额参数入对象
          let money = e[i].object[j].money;
          this.arr.push(Object.assign(e[i].object[j],{'money':money}));
        }
      }
      if(this.arr){
        this.lists = this.unique(this.arr);
        var moj = 0;
        var number = 0;
        var ooo = [];
          //计算金额
        for(let l = 0;l<this.lists.length;l++){
          if(this.lists[l].money){
            number += 1;
            moj += Number(this.lists[l].money);
            ooo.push(this.lists[l])
          }
        }
//        this.$root.$emit('chat_bet',ooo);
        this.money = moj;
        this.number = number;
      }
    });
    this.$root.$on('reset',(e)=>{
      this.money = e
    });
    this.$root.$on('c_data',(e)=>{
      this.qishu = e.qishu
    })
  },
  mounted(){
    let bet_modal = document.querySelector(".bet_modal");
    mymousewheel(bet_modal);
  },
  methods:{
    pour: function(){
      window.pour_status = 1;
      var send_data = this.lists;
      var send_op = [];
      for(let l = 0;l<send_data.length;l++){
        if(send_data[l].money){
          delete send_data[l].number;
          send_op.push(send_data[l]);
        }
      }
      const body = {
        fc_type:JSON.stringify(this.$route.query.page),
        qishu:JSON.stringify(this.qishu),
        data:JSON.stringify(send_op)
      };
      this.loading = true;
      api.addbet(this,body,(res)=>{
        console.log(res.data.ErrorCode);
        if(res.data.ErrorCode == 1){
          window.pour_status = 2;
          this.loading = false;
          this.$Modal.success({
            content: "下注成功",
            onOk: () => {
              this.$root.$emit('success',true);
              this.$root.$emit('bet_success',true)
            },
          });
          window.setTimeout(() => {
            this.$root.$emit('success',true);
            this.$root.$emit('bet_success',true);
            this.$Modal.remove()
          }, share.bet_time)
        }else if(res.data.ErrorCode == 2){
          this.loading = false;
        }
      })
      // console.log(param)
    },
    reset: function(){
      // this.modal = false;
      this.lists = [];
      this.$root.$emit('reset','');
    },
    cancel: function(){
      this.loading = false;
      this.$emit('cancel',false)
    },
    //数组去重
    unique: function(arr){
     var newArr = [];
     for(var i in arr) {
         if(newArr.indexOf(arr[i]) == -1) {
             newArr.push(arr[i])
         }
     }
     return newArr;
    }
  }
}
</script>
