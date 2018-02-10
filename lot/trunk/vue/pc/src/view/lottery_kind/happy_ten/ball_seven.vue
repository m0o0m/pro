<template>
  <div class="happy_ball">
    <div style="width: 880px;
                padding: 5px;
                background: #eee;
                margin: 0 auto;">
      第七球
    </div>
    <div class="ball_box">
      <ul v-for="item in sevenlist">
        <li class="thend">
          <span class="one">号码</span>
          <span class="tow">赔率</span>
          <span class="three">金额</span>
        </li>
        <li v-for="(key,index) in item.object" :class="[key.flag?'styleclect':'']">
           <span :class="{'bor_none':!key.num}" class="one" @click="myclick(key,index)">
             <i v-if="key.bg" :style="{background:key.bg}">{{key.input_name}}</i>
             <span v-else :style="{color:key.styletxt}">{{key.num}}</span>
           </span>
          <span :class="{'bor_none':!key.num}" class="tow" @click="myclick(key,index)">{{key.odd}}</span>
          <span class="three" @click.self="myclick(key,index)">
             <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength="9" v-if="key.num" ref="myfocus" size="small" style="width: 60px;" v-model="key.money"
                    @on-keyup="gogo(key)" @on-afterpaste="gogo(key)" @on-focus="onfocus(key)"></I-Input>
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
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-focus="onfocus_top(1)" @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength="9" style="width: 100px" size="small"
                                                                            v-model="money" @on-keyup="push_money()"
                                                                            @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
  import api from '../../../api/config'
  import MeModal from '../../../share_components/bet'
  //import back from './ten_data'
  import hint from '../../../share_components/hint_msg'
  import share from '../../../share_components/share'
  import {Input, Modal} from 'iview';
  export default {
    components: { MeModal,'I-Input':Input,Modal},
    props: {
      sevenlist: {
        type: Array,
      },
      c_data: {
        type: Object,
      },
    },
    data() {
      return {
        money: '',
        modal: false,
        a:''
      }
    },
    created() {
      this.reset();
      this.$root.$on('success', (e) => {
        if (e) {
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
    watch: {
      // 如果路由有变化，会再次执行该方法
      '$route': 'reset',   // 只有这个页面初始化之后，这个监听事件才开始生效
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
      myclick: function(item, i) {
        console.log(item);
        console.log(item.flag);
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.myfocus[item.index].blur();
          this.$refs.myfocus[item.index].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.myfocus[item.index].focus();
        }else if (item.flag == true && item.money == ''){
          item.flag = false;
          this.$refs.myfocus[item.index].$refs.input.value = '';
        }else if (item.flag == true){
          item.money = this.money;
        }
      },
      onfocus: function(item){
        this.$refs.myfocus[item.index].$refs.input.data_onoff = 'true';
        this.a = item.index;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.index) {
                dom[i].data_onoff = 'false';
            }
        }
        if(item.flag == false && item.money == ''){
          item.flag = true;
          item.money = this.money;
        }
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
      push_money: function () {
        this.money = this.money.replace(/\D/g, "");
        this.computed_money();
      },
      change_money: function () { this.computed_money() },
      gogo(key) {
        key.money = key.money.replace(/\D/g, "");
        ////console.log(key);
        if (!key.money) {
          key.flag = false
        } else {
          key.flag = true
        }
      },
      go_to() {
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        var kk = 0;
        var is_select = false;
        for (let i = 0; i < this.sevenlist.length; i++) {
          for (let j = 0; j < this.sevenlist[i].object.length; j++) {
            let b = this.sevenlist[i].object[j].money + 'b';
            this.sevenlist[i].object[j].money = b.replace(/\D/g, "");
            kk += Number(this.sevenlist[i].object[j].money);
            if(this.sevenlist[i].object[j].flag){
              is_select = true
            }
          }
        }
        if(is_select){
          if (kk != 0) {
            this.modal = true;
//            document.querySelector('body').style.overflow = 'hidden';
            this.$root.$emit('id-selected', this.sevenlist);
            this.$root.$emit('c_data', this.c_data);
          } else if (kk == 0) {
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
      cancel: function (item) {
        this.modal = false;
        document.querySelector('body').style.overflow = 'auto'
      },
      reset: function () {
        for (let i = 0; i < this.sevenlist.length; i++) {
          for (let j = 0; j < this.sevenlist[i].object.length; j++) {
            this.sevenlist[i].object[j].money = '';
            this.sevenlist[i].object[j].flag = '';
          }
        }
        this.money = '';
        this.$root.$emit('reset', '');
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            dom[i].value = '';
            dom[i].data_onoff = 'false';
        }
        this.$root.$emit('clear_key_number','')
      },
      add_money: function (type) {
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money();
      },
      computed_money: function () {
        for (let i = 0; i < this.sevenlist.length; i++) {
          for (let j = 0; j < this.sevenlist[i].object.length; j++) {
            //添加金额参数入对象
            if (this.sevenlist[i].object[j].flag) {
              this.sevenlist[i].object[j].money = this.money
            } else if (!this.sevenlist[i].object[j].flag) {
              this.sevenlist[i].object[j].money = ''
            }
          }
        }
      },
    }
  }
</script>

<style lang="scss" scoped src="../../../assets/css/happy_ball.scss"></style>
