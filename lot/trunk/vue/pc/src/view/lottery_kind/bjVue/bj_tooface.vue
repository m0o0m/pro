<template>
  <div class="bj_tooface">
    <div class="footer1">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :maxlength='9' style="width: 100px" size="small" v-model="money" :value="money" @on-blur="onblur_top(0)"  @on-focus="onfocus_top(0)" @on-keyup="push_money()" @on-change="change_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="bj_too_box">
      <ul class="ul1" v-for="item in toplist">
        <li class="thend">{{item.name}}</li>
        <li v-for="(key,index) in item.object" :class="[key.flag?'styleclect':'']">
          <span @click="numClick(key,index)" class="one">{{key.num}}</span>
          <span @click="numClick(key,index)" class="tow">{{key.odd}}</span>
          <span class="three" @click.self="numClick(key,index)">
            <I-Input @on-keydown="tab_now(key)" :value="key.money" @on-blur="onblur(key)" @on-change="onchange(key)" :maxlength='9' ref="myfocus" style="width: 60px;" size="small" v-model="key.money" @on-focus="onfocus(key)" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
          </span>
        </li>
      </ul>
      <div class="box2">
        <ul v-for="item in centerlist" class="ul2">
          <li class="thend2">{{item.name}}</li>
          <li  v-for="(key,index) in item.object" :class="[key.flag?'styleclect':'']">
            <span @click="numClick(key,index)" class="one2">{{key.num}}</span>
            <span @click="numClick(key,index)" class="tow2">{{key.odd}}</span>
            <span class="three2" @click.self="numClick(key,index)">
              <I-Input @on-keydown="tab_now(key)" :value="key.money" @on-blur="onblur(key)" @on-change="onchange(key)" :maxlength='9' ref="myfocus" style="width: 50px;" size="small" v-model="key.money" @on-focus="onfocus(key)" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
            </span>
          </li>
        </ul>
      </div>
      <div class="box2">
        <ul v-for="item in buttomlist" class="ul3">
          <li class="thend2">{{item.name}}</li>
          <li v-for="(key,index) in item.object" :class="[key.flag?'styleclect':'']">
            <span @click="numClick(key,index)" class="one2">{{key.num}}</span>
            <span @click="numClick(key,index)" class="tow2">{{key.odd}}</span>
            <span class="three2" @click.self="numClick(key,index)">
              <I-Input @on-keydown="tab_now(key)" :value="key.money" @on-blur="onblur(key)" @on-change="onchange(key)" :maxlength='9' ref="myfocus" style="width: 45px;" size="small" v-model="key.money" @on-focus="onfocus(key)" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
            </span>
          </li>
        </ul>
      </div>
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
      <div class="one"><span class="left">金额￥</span><I-Input :maxlength='9' style="width: 100px" size="small" v-model="money" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" @on-change="change_money()"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
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
  import {Input, Modal} from 'iview';
  import hint from '../../../share_components/hint_msg'
  import share from '../../../share_components/share'
  export default {
    components: { MeModal,'I-Input':Input,Modal},
    props: {
      toplist: {
        type: Array,
      },
      centerlist: {
        type: Array,
      },
      buttomlist: {
        type: Array,
      },
      c_data:{
        type:Object,
      }
    },
    data() {
      return {
        money: '',
        modal: false,
        margin: false,
        a:''
      }
    },
    created() {
      //console.log("路由ID:" + this.$route.params.id);
      this.$root.$on('success',(e)=>{
        if(e){
          this.modal = false;
          this.reset()
        }
      });
      this.$root.$on('this_money',(e)=>{
          this.money = e
      });
      this.fetchData();
    },
    mounted(){
      this.$root.$emit('no_top', false);
    },
    watch: {
      'route': 'fetchData',
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
      },
      numClick: function(item, i) {
        console.log(item);
        console.log(item.flag);
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.myfocus[item.index].blur();
          this.$refs.myfocus[item.id].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.myfocus[item.index].focus();
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
      add_money: function(type){
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money()
      },
      push_money: function () {
        this.money = this.money.replace(/\D/g, "");
        this.computed_money()
      },
      change_money: function () {
        this.computed_money()
      },
      computed_money: function () {
        for (let i = 0; i < this.toplist.length; i++) {
          for (let j = 0; j < this.toplist[i].object.length; j++) {
            //添加金额参数入对象
            if (this.toplist[i].object[j].flag) {
              this.toplist[i].object[j].money = this.money
            }
          }
        }
        //list2
        for (let i = 0; i < this.centerlist.length; i++) {
          for (let j = 0; j < this.centerlist[i].object.length; j++) {
            //添加金额参数入对象
            if (this.centerlist[i].object[j].flag) {
              this.centerlist[i].object[j].money = this.money
            }
          }
        }
        //list3
        for (let i = 0; i < this.buttomlist.length; i++) {
          for (let j = 0; j < this.buttomlist[i].object.length; j++) {
            //添加金额参数入对象
            if (this.buttomlist[i].object[j].flag) {
              this.buttomlist[i].object[j].money = this.money
            }
          }
        }
      },
      go_to: function () {
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        var kk = 0;
        var is_select = false;
        for (let i = 0; i < this.toplist.length; i++) {
          for (let j = 0; j < this.toplist[i].object.length; j++) {
            let b = this.toplist[i].object[j].money + 'b';
            this.toplist[i].object[j].money = b.replace(/\D/g, "");
            kk += Number(this.toplist[i].object[j].money);
            if(this.toplist[i].object[j].flag){
              is_select = true
            }
          }
        }
        //list2
        for (let i = 0; i < this.centerlist.length; i++) {
          for (let j = 0; j < this.centerlist[i].object.length; j++) {
            let c = this.centerlist[i].object[j].money + 'c';
            this.centerlist[i].object[j].money = c.replace(/\D/g, "");
            kk += Number(this.centerlist[i].object[j].money);
            if(this.centerlist[i].object[j].flag){
              is_select = true
            }
          }
        }
        //list3
        for (let i = 0; i < this.buttomlist.length; i++) {
          for (let j = 0; j < this.buttomlist[i].object.length; j++) {
            let d = this.buttomlist[i].object[j].money + 'd';
            this.buttomlist[i].object[j].money = d.replace(/\D/g, "");
            kk += Number(this.buttomlist[i].object[j].money);
            if(this.buttomlist[i].object[j].flag){
              is_select = true
            }
          }
        }
        if(is_select){
          if (kk != 0) {
            this.modal = true;
//            document.querySelector('body').style.overflow = 'hidden';
            var arr = this.toplist.concat(this.centerlist, this.buttomlist);
            this.$root.$emit('c_data',this.c_data);
            this.$root.$emit('id-selected', arr);
          } else if (kk == 0) {
            this.$Modal.warning({
              content: hint.money_null,
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
        for (let i = 0; i < this.toplist.length; i++) {
          for (let j = 0; j < this.toplist[i].object.length; j++) {
            this.toplist[i].object[j].money = '';
            this.toplist[i].object[j].flag = '';
          }
        }
        for (let i = 0; i < this.centerlist.length; i++) {
          for (let j = 0; j < this.centerlist[i].object.length; j++) {
            this.centerlist[i].object[j].money = '';
            this.centerlist[i].object[j].flag = '';
          }
        }
        for (let i = 0; i < this.buttomlist.length; i++) {
          for (let j = 0; j < this.buttomlist[i].object.length; j++) {
            this.buttomlist[i].object[j].money = '';
            this.buttomlist[i].object[j].flag = '';
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
      }
    }
  }
</script>

<style lang="scss" src="../../../assets/css/bj_tooface.scss" scoped></style>
