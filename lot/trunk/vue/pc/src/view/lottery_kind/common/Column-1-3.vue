<template>
  <div class='box div_content'>
    <div class="footer2">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="money_button one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :value="money" @on-blur="onblur_top(0)"  @on-focus="onfocus_top(0)" :maxlength="9" style="width: 100px" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="money_button two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="money_button two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="box_table">
      <ul class='list' v-for="item in one">
        <li>
          {{item.name}}
        </li>
        <li>
          <ul>
            <li v-for="(key,index) in item.object" :class="['lis_2',key.flag?'table-current':'']">
              <ul class='clearfix'>
                <li class='fl cla' @click='numClick(key,index)'>
                  <!-- {{key.num}} -->
                  <span class='ball' v-if="index<10">{{key.input_name}}</span>
                  <span class="box_bg" v-else-if="index>=10">{{key.num}}</span>
                </li>
                <li class='fl cla2' @click='numClick(key,index)'>
                  {{key.odd}}
                </li>
                <li class='fl cla3' @click.self='numClick(key,index)'>
                  <!-- <I-Input type="text" v-model="key.item" ref="myfocus"> -->
                  <I-Input :maxlength="9"  ref="myfocus" style="width: 60px" v-model="key.money" :value="key.money" @on-blur="onblur(key)" size="small" @on-keydown="tab_now(key)"  @on-change="onchange(key)" @on-focus="onfocus(key)" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
                </li>
              </ul>
            </li>
          </ul>
        </li>
      </ul>
    </div>
    <div class="footer2" style="margin-top:10px;margin-bottom: 20px">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="money_button one"><span class="left">金额￥</span><I-Input style="width: 100px" @on-change="change_money()" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="money_button two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="money_button two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
  // import back from './back_data'
  import {Input, Modal} from 'iview';
  import api from '../../../api/config'
  import MeModal from '../../../share_components/bet'
  import hint from '../../../share_components/hint_msg'
  import share from '../../../share_components/share'
  export default {
    components: {
     MeModal,'I-Input':Input,Modal
    },
    data() {
      return {
        money:'',
        modal: false,
        a:''
      }
    },
    props:{
      one:{
        type:Array
      },
      c_data:{
        type:Object
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
      this.$root.$emit('no_top', false);
    },
    watch: {
      // 如果路由有变化，会再次执行该方法
      '$route': 'reset',// 只有这个页面初始化之后，这个监听事件才开始生成
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
      computed: function(data) {
        this.$set(this.one, this.one);
        for (let j = 0; j < this.one[0].object.length; j++) {
          Object.assign(this.one[0].object[j], data[j]);
        }
        for (let l = 14, k = 0; l < this.one[1].object.length, k < this.one[1].object.length; l++, k++) {
          Object.assign(this.one[1].object[k], data[l]);
        }
        for (let n = 28, m = 0; n < this.one[2].object.length, m < this.one[2].object.length; n++, m++) {
          Object.assign(this.one[2].object[m], data[n]);
        }
        // console.log(this.one);
      },
      numClick: function(item,i) {
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.myfocus[item.id].blur();
          this.$refs.myfocus[item.id].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.myfocus[item.id].focus();
        }else if (item.flag == true && item.money == ''){
          item.flag = false;
          this.$refs.myfocus[item.id].$refs.input.value = '';
        }else if (item.flag == true){
          item.money = this.money;
        }
      },
      onfocus: function(item){
        this.$refs.myfocus[item.id].$refs.input.data_onoff = 'true';
        this.a = item.id;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.id+1) {
                dom[i].data_onoff = 'false';
            }
        }
        if(item.flag == false && item.money == ''){
          item.flag = true;
          item.money = this.money;
        }
      },
      onblur: function (key) {
        key.money = this.$refs.myfocus[key.id].$refs.input.value;
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
      //点击下注
      go_to: function() {
        this.$root.$emit('c_data',this.c_data);
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        var kk = 0;
        var is_select = false;
        for (let i = 0; i < this.one.length; i++) {
          for (let j = 0; j < this.one[i].object.length; j++) {
            let b = this.one[i].object[j].money + 'b';
            this.one[i].object[j].money = b.replace(/\D/g, "");
            kk += Number(this.one[i].object[j].money);
            if(this.one[i].object[j].flag){
              is_select = true
            }
          }
        }
        console.log('kk:' + kk);
        if(is_select){
          if (kk != 0) {
            this.modal = true;
            this.$root.$emit('id-selected', this.one);
          }else if (kk == 0) {
            this.$Modal.warning({
              content: hint.money_null
            });
            window.setTimeout(() => {
              this.$Modal.remove()
          }, share.Prompt)
          }
        }else{
          this.$Modal.warning({
            content: hint.all_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }
      },
      //按键tab
      tab_now:function (key) {
        if(!key.money) {
          key.flag = false;
        }else if(key.money){
          key.flag = true;
        }
      },
      //监听input值得变化
      onchange: function (key) {
        console.log(key.money);
        if(!key.money) {
          key.flag = false;
        }else if(key.money){
          key.flag = true;
        }
      },
      //每个球对应的输入框
      gogo: function(key) {
        key.money = key.money.replace(/\D/g, "");
      },
      change_money: function () {
          console.log('change_money:'+this.money);
           this.computed_money() 
      },
      push_money: function() {
        // this.disabled = true;
        this.money = this.money.replace(/\D/g, "");
        this.computed_money()
      },
      add_money: function(type){
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money()
      },
      //处理金额
      computed_money: function () {
        for (let i = 0; i < this.one.length; i++) {
          for (let j = 0; j < this.one[i].object.length; j++) {
            //添加金额参数入对象
            if (this.one[i].object[j].flag) {
              this.one[i].object[j].money = this.money
            }else if(!this.one[i].object[j].flag){
              this.one[i].object[j].money = ''
            }
          }
        }
      },
      cancel: function(item) {
        this.modal = false;
        document.querySelector('body').style.overflow='auto';
      },
      reset: function() {
        this.money = '';
        this.$root.$emit('reset', '');
        for (let i = 0; i < this.one.length; i++) {
          for (let j = 0; j < this.one[i].object.length; j++) {
            this.one[i].object[j].money = '';
            this.one[i].object[j].flag = false;
          }
        }
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            dom[i].value = '';
            dom[i].data_onoff = 'false';
        }
       this.$root.$emit('clear_key_number','')
      }
    }
  };
</script>
<style lang="scss" src="../../../assets/css/pl3_one.scss" scoped></style>
