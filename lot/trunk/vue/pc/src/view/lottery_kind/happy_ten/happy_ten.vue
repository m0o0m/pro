<template>
  <div class="happy_ten">
    <div class="h_t_box">
      <ul v-for="item in tenlist">
        <li class="thend1"  @click="choose(item)">
          <input type="radio" ref="radio_choose" name="contact" @click="choose(item)" :value="item.index" v-model="picked">
        </li>
        <li class="thend2" @click="choose(item)">{{item.name}}</li>
        <li class="thend3" @click="choose(item)">{{item.odd}}</li>
        <li @click.prevent="myclick(key)"  v-for="(key,i) in item.object" :class="{'styleclect':key.state}">
          <span class="one"><i>{{key.input_name}}</i></span>
          <span class="tow">
            <input type="checkbox" v-model="key.state"></input>
          </span>
        </li>
      </ul>
    </div>
    <div class="footer1" style="margin-bottom: 20px">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-focus="onfocus_top(1)" @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength="9" style="width: 100px" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
  import api from '../../../api/config'
  import MeModal from '../../../share_components/bet_happy.vue'
  import hint from '../../../share_components/hint_msg'
  import {Input, Modal} from 'iview';
  import share from '../../../share_components/share'
  export default {
    components: { MeModal,'I-Input':Input,Modal},
    data() {
      return {
        money: '',
        modal: false,
//        my_data:'',
        arr:[],//个数计算数据,及选中保存下来的数组
        picked:0,//当选框默认选定
        a:''
      }
    },
    created() {
      this.reset();
      this.$root.$on('success',(e)=>{
        if(e){
          this.modal = false;
          this.reset()
        }
      });
      this.$root.$on('this_money',(e)=>{
          this.money = e
      });
    },
    mounted(){
      this.$root.$emit('no_top', true);
    },
    props:{
      tenlist:{
        type:Array,
      },
      c_data:{
        type:Object,
      },
      only_data:{
        type:Array
      }
    },

    watch:{
      '$route': 'fetchData',   // 只有这个页面初始化之后，这个监听事件才开始生效
      a: function (new_val,old_val) {
          if(new_val != old_val){
              this.$root.$emit('clear_key_number','')
          }
      },
      money: function(new_val,old_val){
          if(new_val != old_val){
              this.computed_money()
          }
      }
    },
    methods: {
      fetchData: function(){
        this.reset();
        this.first_choose()
      },
      go_to: function() {
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        for(let i = 0;i < this.tenlist.length;i++){
          for(let j = 0;j < this.tenlist[i].object.length;j++){
            if(this.tenlist[i].object[j].state){
              let b = this.tenlist[i].object[j].money + 'b';
              this.tenlist[i].object[j].money = b.replace(/\D/g, "");
            }
          }
        }
        this.$root.$emit('c_data',this.c_data);
        if(this.picked == 0 && this.arr.length < 2){
          this.$Modal.warning({
            content: '至少选2个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }else if(this.picked == 1 && this.arr.length < 2){
          this.$Modal.warning({
            content: '至少选2个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }else if(this.picked == 2 && this.arr.length < 3){
          this.$Modal.warning({
            content: '至少选3个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }else if(this.picked == 3 && this.arr.length < 4){
          this.$Modal.warning({
            content: '至少选4个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }
        else if(this.picked == 4 && this.arr.length < 5) {
          this.$Modal.warning({
            content: '至少选5个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }else{
          this.$root.$emit('id-selected-h',this.tenlist);
          if(this.money){
            this.modal = true;
            //console.log(this.list);
//            document.querySelector('body').style.overflow='hidden';
//            this.push_money();
          }else{
            this.$Modal.warning({
              content: hint.money_null
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }
        }
      },
      push_money: function(){
        this.money = this.money.replace(/\D/g, "");
        this.computed_money()
      },
        onblur: function (key) {
            key.money = this.$refs.myfocus[key.index].$refs.input.value;
            console.log('选中后的金额：'+key.money);
        },
        onfocus_top: function(index){
            let dom = document.querySelectorAll('input');
            this.a = 99;
            if(index == 0){
                index=0
            }else{
                index=dom.length-1;
            }
            for(let i = 0;i < dom.length;i++){
                if(i != index) {
                    dom[i].data_onoff = 'false';
                }else{
                    dom[i].data_onoff = 'true'
                }
            }
        },
        onblur_top: function(index){
            let dom = document.querySelectorAll('input');
            if(index == 0){
                index=0
            }else{
                index=dom.length-1;
            }
          if(dom[index].value != 'on'){
            this.money = dom[index].value;
          }
        },
      change_money: function () { this.computed_money() },
      add_money: function(type){
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money()
      },
      computed_money: function(){
        for(let i = 0;i < this.tenlist.length;i++){
          for(let j = 0;j < this.tenlist[i].object.length;j++){
            if(this.tenlist[i].object[j].state){
              Object.assign(this.tenlist[i].object[j],{'money':this.money});
            }
          }
        }
      },
      cancel: function(item) {
        this.modal = false;
        document.querySelector('body').style.overflow='auto'
      },
      myclick: function(key){
        console.log(key);
        key.state = !key.state;
        if(key.state){
          this.computed_money()
        }
        //console.log(key);
        var max_limit = [];
        if(this.$route.query.page == 'cq_ten'){
          max_limit = [5,5,5,7,8];
        }else if (this.$route.query.page == 'gd_ten') {
          max_limit = [5,5,5,7,8];
        }

        if(key.state){
          this.arr.push(key);
          this.arr = this.unique(this.arr);
          if(this.picked == 0 && this.arr.length>max_limit[0]){
            this.arr.pop();
            //console.log(this.arr);
            key.state = false;
            //console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[0]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 1 && this.arr.length > max_limit[1]){
            this.arr.pop();
            //console.log(this.arr);
            key.state = false;
            //console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[1]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 2 && this.arr.length > max_limit[2]){
            this.arr.pop();
            //console.log(this.arr);
            key.state = false;
            //console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[2]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 3 && this.arr.length >  max_limit[3]){
            this.arr.pop();
            //console.log(this.arr);
            key.state = false;
            //console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+ max_limit[3] +'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 4 && this.arr.length >  max_limit[4]){
            this.arr.pop();
            //console.log(this.arr);
            key.state = false;
            //console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+ max_limit[4]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }
        }else{
          for(let i = 0;i<this.arr.length;i++){
            //console.log('数组的所有对象的index：'+this.arr[i].index);
            //console.log('取消的index：'+key.index);
            if(key.index == this.arr[i].index){
              this.arr.splice(i,1);
              //console.log(this.arr)
              break;
            }
          }
        }
      },
      choose: function(key){
        this.reset();
        console.log('index:'+key.index);
        this.picked = key.index;
        if(this.picked == 0){
          for(let i= 0;i<this.tenlist[0].object.length;i++){
            Object.assign(this.tenlist[0].object[i],this.only_data[i])
          }
          for(let i= 0;i<this.tenlist[1].object.length;i++){
            Object.assign(this.tenlist[1].object[i],this.only_data[4+i])
          }
          for(let i= 0;i<this.tenlist[2].object.length;i++){
            Object.assign(this.tenlist[2].object[i],this.only_data[8+i])
          }
          for(let i= 0;i<this.tenlist[3].object.length;i++){
            Object.assign(this.tenlist[3].object[i],this.only_data[12+i])
          }
          for(let i= 0;i<this.tenlist[4].object.length;i++){
            Object.assign(this.tenlist[4].object[i],this.only_data[16+i])
          }
        }else if(this.picked == 1){
          for(let i= 0;i<this.tenlist[0].object.length;i++){
            Object.assign(this.tenlist[0].object[i],this.only_data[20+i])
          }
          for(let i= 0;i<this.tenlist[1].object.length;i++){
            Object.assign(this.tenlist[1].object[i],this.only_data[24+i])
          }
          for(let i= 0;i<this.tenlist[2].object.length;i++){
            Object.assign(this.tenlist[2].object[i],this.only_data[28+i])
          }
          for(let i= 0;i<this.tenlist[3].object.length;i++){
            Object.assign(this.tenlist[3].object[i],this.only_data[32+i])
          }
          for(let i= 0;i<this.tenlist[4].object.length;i++){
            Object.assign(this.tenlist[4].object[i],this.only_data[36+i])
          }
        }else if(this.picked == 2){
          for(let i= 0;i<this.tenlist[0].object.length;i++){
            Object.assign(this.tenlist[0].object[i],this.only_data[40+i])
          }
          for(let i= 0;i<this.tenlist[1].object.length;i++){
            Object.assign(this.tenlist[1].object[i],this.only_data[44+i])
          }
          for(let i= 0;i<this.tenlist[2].object.length;i++){
            Object.assign(this.tenlist[2].object[i],this.only_data[48+i])
          }
          for(let i= 0;i<this.tenlist[3].object.length;i++){
            Object.assign(this.tenlist[3].object[i],this.only_data[52+i])
          }
          for(let i= 0;i<this.tenlist[4].object.length;i++){
            Object.assign(this.tenlist[4].object[i],this.only_data[56+i])
          }
        }else if(this.picked == 3){
          for(let i= 0;i<this.tenlist[0].object.length;i++){
            Object.assign(this.tenlist[0].object[i],this.only_data[60+i])
          }
          for(let i= 0;i<this.tenlist[1].object.length;i++){
            Object.assign(this.tenlist[1].object[i],this.only_data[64+i])
          }
          for(let i= 0;i<this.tenlist[2].object.length;i++){
            Object.assign(this.tenlist[2].object[i],this.only_data[68+i])
          }
          for(let i= 0;i<this.tenlist[3].object.length;i++){
            Object.assign(this.tenlist[3].object[i],this.only_data[72+i])
          }
          for(let i= 0;i<this.tenlist[4].object.length;i++){
            Object.assign(this.tenlist[4].object[i],this.only_data[76+i])
          }
        }else if(this.picked == 4){
          for(let i= 0;i<this.tenlist[0].object.length;i++){
            Object.assign(this.tenlist[0].object[i],this.only_data[80+i])
          }
          for(let i= 0;i<this.tenlist[1].object.length;i++){
            Object.assign(this.tenlist[1].object[i],this.only_data[84+i])
          }
          for(let i= 0;i<this.tenlist[2].object.length;i++){
            Object.assign(this.tenlist[2].object[i],this.only_data[88+i])
          }
          for(let i= 0;i<this.tenlist[3].object.length;i++){
            Object.assign(this.tenlist[3].object[i],this.only_data[92+i])
          }
          for(let i= 0;i<this.tenlist[4].object.length;i++){
            Object.assign(this.tenlist[4].object[i],this.only_data[96+i])
          }
        }
        for(let i = 0;i < this.tenlist.length;i++){
          for(let j = 0;j < this.tenlist[i].object.length;j++){
            if(!key.state){
              this.tenlist[i].object[j].state = false
            }else{
              this.tenlist[i].object[j].state = true;
            }
//            Object.assign(this.tenlist[i].object[j], this.only_data[key.index]);
          }
        }
      },
      first_choose: function(){
        this.picked = 0;
        for(let i= 0;i<this.tenlist[0].object.length;i++){
          Object.assign(this.tenlist[0].object[i],this.only_data[288+i])
        }
        for(let i= 0;i<this.tenlist[1].object.length;i++){
          Object.assign(this.tenlist[1].object[i],this.only_data[292+i])
        }
        for(let i= 0;i<this.tenlist[2].object.length;i++){
          Object.assign(this.tenlist[2].object[i],this.only_data[296+i])
        }
        for(let i= 0;i<this.tenlist[3].object.length;i++){
          Object.assign(this.tenlist[3].object[i],this.only_data[300+i])
        }
        for(let i= 0;i<this.tenlist[4].object.length;i++){
          Object.assign(this.tenlist[4].object[i],this.only_data[304+i])
        }
      },
      reset: function(){
        this.money = '';
        this.arr = [];
        this.$root.$emit("reset","");
        for(let i = 0;i < this.tenlist.length;i++){
          for(let j = 0;j<this.tenlist[i].object.length;j++){
            this.tenlist[i].object[j].state = false;
            this.tenlist[i].object[j].money = '';
          }
        }
      },
      unique: function(arr){
        var newArr = [];
        for(var i in arr) {
          if(newArr.indexOf(arr[i]) == -1) {
            newArr.push(arr[i])
          }
        }
        return newArr;
      },
    }
  }
</script>

<style lang="scss" src="../../../assets/css/happy_ten.scss" scoped></style>
