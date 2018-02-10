<template>
  <div class="e_s_f">
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :maxlength="9" style="width: 100px" size="small" v-model="money" :value="money" @on-blur="onblur_top(0)"  @on-focus="onfocus_top(0)"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="e_s_f_box">
      <ul v-for="item in esfLists">
        <li class="thend">{{item.name}}</li>
        <li v-for="key in item.object" :class="key.flag?'styleclect':''">
          <span class="one fl" @click="myclick(key)">
            <i class='ball'>
            {{key.name}}
            </i>
          </span>
          <span class="tow fl" @click="myclick(key)">{{key.odd}}</span>
          <span class="three" @click.self="myclick(key)">
             <I-Input @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength="9" class="inp" ref="myfocus" style="width: 45px" @on-focus="onfocus(key)" :value="key.money" @on-blur="onblur(key)" v-model="key.money" size="small" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
            </span>
        </li>
      </ul>
    </div>
    <div class="footer1 clearfix" style="margin-bottom: 20px">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" :maxlength="9" style="width: 100px" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
// import api from "../../../api/config";
import {Input, Modal} from 'iview';
import MeModal from "../../../share_components/bet";
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
export default {
components: { MeModal,'I-Input':Input,Modal},
  data() {
    return {
      money: "",
      modal: false,
      disabled: false,
      money_disabled: false,
      a:''
    };
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
  props: ["cdata", "esfLists"],
  created() {
    this.$root.$on("success", e => {
      if (e) {
        this.modal = false;
        this.reset();
      }
    });
    this.$root.$on('this_money',(e)=>{
        this.money = e
    });
    this.reset();
  },
  mounted(){
    this.$root.$emit('no_top', false);
  },
  methods: {
    myclick: function(item, i) {
        // console.log(item);
        // console.log(item.flag);
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.myfocus[item.li_id].blur();
          this.$refs.myfocus[item.li_id].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.myfocus[item.li_id].focus();
        }else if (item.flag == true && item.money == ''){
          item.flag = false;
          this.$refs.myfocus[item.li_id].$refs.input.value = '';
        }else if (item.flag == true){
          item.money = this.money;
        }
      },
      onfocus: function(item){
        this.$refs.myfocus[item.li_id].$refs.input.data_onoff = 'true';
        this.a = item.li_id;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.li_id+1) {
                dom[i].data_onoff = 'false';
            }
        }
        if(item.flag == false && item.money == ''){
          item.flag = true;
          item.money = this.money;
        }
      },
      onblur: function (key) {
          key.money = this.$refs.myfocus[key.li_id].$refs.input.value;
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
    cancel: function(item) {
      this.modal = false;
      document.querySelector("body").style.overflow = "auto";
    },
    push_money: function() {
      this.money = this.money.replace(/\D/g, "");
      this.computed_money();
    },
    change_money: function () { this.computed_money() },
    go_to: function() {
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.esfLists.length; i++) {
        for (let j = 0; j < this.esfLists[i].object.length; j++) {
          let b = this.esfLists[i].object[j].money + 'b';
          this.esfLists[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.esfLists[i].object[j].money);
          if(this.esfLists[i].object[j].flag){
            is_select = true
          }
        }
      }
      // console.log("kk:" + kk);
      if(is_select){
        if (kk != 0) {
          this.modal = true;
          this.$root.$emit("c_data", this.cdata);
          this.$root.$emit("id-selected", this.esfLists);
//          document.querySelector("body").style.overflow = "hidden";
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
    computed_money() {
      for (let i = 0; i < this.esfLists.length; i++) {
        for (let j = 0; j < this.esfLists[i].object.length; j++) {
          if (this.esfLists[i].object[j].flag) {
            this.esfLists[i].object[j].money = this.money;
          } else {
            this.esfLists[i].object[j].money = "";
          }
        }
      }
    },
    reset: function() {
      for (let i = 0; i < this.esfLists.length; i++) {
        for (let j = 0; j < this.esfLists[i].object.length; j++) {
          this.esfLists[i].object[j].money = "";
          this.esfLists[i].object[j].flag = false;
        }
      }
      this.money = "";
      this.$root.$emit("reset", "");
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          dom[i].value = '';
          dom[i].data_onoff = 'false';
      }
      this.$root.$emit('clear_key_number','')
    },
    //按键tab
    tab_now: function(key) {
      if (!key.money) {
        key.flag = false;
      } else if (key.money) {
        key.flag = true;
      }
    },
    //监听input值得变化
    onchange: function(key) {
      if (!key.money) {
        key.flag = false;
      } else if (key.money) {
        key.flag = true;
      }
    },
    //每个球对应的输入框
    gogo: function(key) {
      key.money = key.money.replace(/\D/g, "");
    },
    add_money: function(type) {
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money();
    }
  }
};
</script>

<style lang="scss" src="../../../assets/css/eleven_select_five.scss" scoped>

</style>
