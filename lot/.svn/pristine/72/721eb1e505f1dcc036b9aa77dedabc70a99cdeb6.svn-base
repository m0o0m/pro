<template lang="html">
  <div class="six_one">
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-blur="onblur_top(0)" @on-focus="onfocus_top(0)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="content">
      <div class="top">
        <ul>
          <li class="only">项目</li>
          <li class="one">正码1</li>
          <li class="one">正码2</li>
          <li class="one">正码3</li>
          <li class="one">正码4</li>
          <li class="one">正码5</li>
          <li class="one">正码6</li>
        </ul>
      </div>
      <div class="table">
        <div class="bottom" v-for="item in list">
          <ul v-for="key in item.object">
            <li @click="pour(key)" class="one" v-show="key.num">{{key.num}}</li>
            <li @click="pour(key)" :class="[key.flag?'bg_color':'','two1']">{{key.odd}}</li>
            <li @click.self="pour(key)" :class="[key.flag?'bg_color':'','three1']">
              <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength='9' ref="input" @on-focus="onfocus(key)" style="width: 45px" v-model="key.money" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)" size="small"></I-Input>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
import api from '../../../api/config'
import {Modal,Input} from 'iview';
import MeModal from '../../../share_components/bet'
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
export default {
  components: {
    MeModal,Modal,'I-Input':Input
  },
  watch:{
    '$route': 'fetchData',
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
  created(){
    this.fetchData();
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
    this.$root.$emit('no_top',false);
    this.$root.$emit("child_change", 0);
    this.$root.$on('time_out',(e)=>{
      if(e){
        this.fetchData(2)
      }
    })
  },
  destroyed(){
    this.$root.$off('time_out')
  },
  data() {
    return {
      a:'',
      money: '',
      modal: false,
      list:[
        {
          name:'1',object:[
            {
              index:0,num:'大',number:'48.8',flag:false,money:''
            },
            {
              index:1,num:'小',number:'48.8',flag:false,money:''
            },
            {
              index:2,num:'单',number:'48.8',flag:false,money:''
            },
            {
              index:3,num:'双',number:'48.8',flag:false,money:''
            },
            {
              index:4,num:'合大',number:'48.8',flag:false,money:''
            },
            {
              index:5,num:'合小',number:'48.8',flag:false,money:''
            },
            {
              index:6,num:'合单',number:'48.8',flag:false,money:''
            },
            {
              index:7,num:'合双',number:'48.8',flag:false,money:''
            },
            {
              index:8,num:'尾大',number:'48.8',flag:false,money:''
            },
            {
              index:9,num:'尾小',number:'48.8',flag:false,money:''
            },
            {
              index:10,num:'红波',number:'48.8',flag:false,money:''
            },
            {
              index:11,num:'绿波',number:'48.8',flag:false,money:''
            },
            {
              index:12,num:'蓝波',number:'48.8',flag:false,money:''
            },
          ]
        },
        {
          name:'2',object:[
            {
              index:13,number:'48.8',flag:false,money:''
            },
            {
              index:14,number:'48.8',flag:false,money:''
            },
            {
              index:15,number:'48.8',flag:false,money:''
            },
            {
              index:16,number:'48.8',flag:false,money:''
            },
            {
              index:17,number:'48.8',flag:false,money:''
            },
            {
              index:18,number:'48.8',flag:false,money:''
            },
            {
              index:19,number:'48.8',flag:false,money:''
            },
            {
              index:20,number:'48.8',flag:false,money:''
            },
            {
              index:21,number:'48.8',flag:false,money:''
            },
            {
              index:22,number:'48.8',flag:false,money:''
            },
            {
              index:23,number:'48.8',flag:false,money:''
            },
            {
              index:24,number:'48.8',flag:false,money:''
            },
            {
              index:25,number:'48.8',flag:false,money:''
            },
          ]
        },
        {
          name:'2',object:[
            {
              index:26,number:'48.8',flag:false,money:''
            },
            {
              index:27,number:'48.8',flag:false,money:''
            },
            {
              index:28,number:'48.8',flag:false,money:''
            },
            {
              index:29,number:'48.8',flag:false,money:''
            },
            {
              index:30,number:'48.8',flag:false,money:''
            },
            {
              index:31,number:'48.8',flag:false,money:''
            },
            {
              index:32,number:'48.8',flag:false,money:''
            },
            {
              index:33,number:'48.8',flag:false,money:''
            },
            {
              index:34,number:'48.8',flag:false,money:''
            },
            {
              index:35,number:'48.8',flag:false,money:''
            },
            {
              index:36,number:'48.8',flag:false,money:''
            },
            {
              index:37,number:'48.8',flag:false,money:''
            },
            {
              index:38,number:'48.8',flag:false,money:''
            },
          ]
        },
        {
          name:'2',object:[
            {
              index:39,number:'48.8',flag:false,money:''
            },
            {
              index:40,number:'48.8',flag:false,money:''
            },
            {
              index:41,number:'48.8',flag:false,money:''
            },
            {
              index:42,number:'48.8',flag:false,money:''
            },
            {
              index:43,number:'48.8',flag:false,money:''
            },
            {
              index:44,number:'48.8',flag:false,money:''
            },
            {
              index:45,number:'48.8',flag:false,money:''
            },
            {
              index:46,number:'48.8',flag:false,money:''
            },
            {
              index:47,number:'48.8',flag:false,money:''
            },
            {
              index:48,number:'48.8',flag:false,money:''
            },
            {
              index:49,number:'48.8',flag:false,money:''
            },
            {
              index:50,number:'48.8',flag:false,money:''
            },
            {
              index:51,number:'48.8',flag:false,money:''
            },
          ]
        },
        {
          name:'2',object:[
            {
              index:52,number:'48.8',flag:false,money:''
            },
            {
              index:53,number:'48.8',flag:false,money:''
            },
            {
              index:54,number:'48.8',flag:false,money:''
            },
            {
              index:55,number:'48.8',flag:false,money:''
            },
            {
              index:56,number:'48.8',flag:false,money:''
            },
            {
              index:57,number:'48.8',flag:false,money:''
            },
            {
              index:58,number:'48.8',flag:false,money:''
            },
            {
              index:59,number:'48.8',flag:false,money:''
            },
            {
              index:60,number:'48.8',flag:false,money:''
            },
            {
              index:61,number:'48.8',flag:false,money:''
            },
            {
              index:62,number:'48.8',flag:false,money:''
            },
            {
              index:63,number:'48.8',flag:false,money:''
            },
            {
              index:64,number:'48.8',flag:false,money:''
            },
          ]
        },
        {
          name:'2',object:[
            {
              index:65,number:'48.8',flag:false,money:''
            },
            {
              index:66,number:'48.8',flag:false,money:''
            },
            {
              index:67,number:'48.8',flag:false,money:''
            },
            {
              index:68,number:'48.8',flag:false,money:''
            },
            {
              index:69,number:'48.8',flag:false,money:''
            },
            {
              index:70,number:'48.8',flag:false,money:''
            },
            {
              index:71,number:'48.8',flag:false,money:''
            },
            {
              index:72,number:'48.8',flag:false,money:''
            },
            {
              index:73,number:'48.8',flag:false,money:''
            },
            {
              index:74,number:'48.8',flag:false,money:''
            },
            {
              index:75,number:'48.8',flag:false,money:''
            },
            {
              index:76,number:'48.8',flag:false,money:''
            },
            {
              index:77,number:'48.8',flag:false,money:''
            },
          ]
        },
      ],
    }
  },
  methods:{
    fetchData: function(type){
      this.reset();
      type==2?this.$root.$emit('loading',true,true):this.$root.$emit('loading',true);
      let body = {
        'fc_type': this.$route.query.page,
        'gameplay':168,
        'pankou':'A'
      };
      api.getgameindex(this, body, (res) => {
        console.log('errorCode:' + res.data.ErrorCode);
          this.$root.$emit('only_back',res,type);
        if (res.data.ErrorCode == 1) {
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
          }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
          this.$root.$emit('c_data',res.data.Data.c_data);
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          this.computed(back_data);
//          this.$root.$emit('get_closetime',res.data.Data.closetime);
        }
      })
    },
    computed: function(data){
      this.$set(this.list, this.list);
      for (let j = 0; j < this.list[0].object.length; j++) {
        Object.assign(this.list[0].object[j], data[j]);
      }
      for (let l = 13, k = 0; l < this.list[1].object.length, k < this.list[1].object.length; l++, k++) {
        Object.assign(this.list[1].object[k], data[l]);
      }
      for (let n = 26, m = 0; n < this.list[2].object.length, m < this.list[2].object.length; n++, m++) {
        Object.assign(this.list[2].object[m], data[n]);
      }
      for (let n = 39, m = 0; n < this.list[3].object.length, m < this.list[3].object.length; n++, m++) {
        Object.assign(this.list[3].object[m], data[n]);
      }
      for (let n = 52, m = 0; n < this.list[4].object.length, m < this.list[4].object.length; n++, m++) {
        Object.assign(this.list[4].object[m], data[n]);
      }
      for (let n = 65, m = 0; n < this.list[5].object.length, m < this.list[5].object.length; n++, m++) {
        Object.assign(this.list[5].object[m], data[n]);
      }
    },
    pour: function(item) {
        console.log(item);
        console.log(item.flag);
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.input[item.index].blur();
          this.$refs.input[item.index].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.input[item.index].focus();
        }else if (item.flag == true && item.money == ''){
          item.flag = false;
          this.$refs.input[item.index].$refs.input.value = '';
        }else if (item.flag == true){
          item.money = this.money;
        }
      },
      onfocus: function(item){
        this.$refs.input[item.index].$refs.input.data_onoff = 'true';
        this.a = item.index;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.index+1) {
                dom[i].data_onoff = 'false';
            }
        }
        if(item.flag == false && item.money == ''){
          item.flag = true;
          item.money = this.money;
        }
      },
      onblur: function (key) {
          key.money = this.$refs.input[key.index].$refs.input.value;
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
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.list.length; i++) {
        for (let j = 0; j < this.list[i].object.length; j++) {
          let b = this.list[i].object[j].money + 'b';
          this.list[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.list[i].object[j].money);
          if(this.list[i].object[j].flag){
            is_select = true
          }
        }
      }
      console.log('kk:' + kk);
      if(is_select){
        if (kk != 0) {
          this.modal = true;
//          document.querySelector('body').style.overflow='hidden';
          this.$root.$emit('id-selected', this.list);
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
    push_money: function() {
      // this.disabled = true;
      this.money = this.money.replace(/\D/g, "");
      this.computed_money()
    },
    change_money: function () { this.computed_money() },
    add_money: function(type){
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money()
    },
    //处理金额
    computed_money: function () {
      for (let i = 0; i < this.list.length; i++) {
        for (let j = 0; j < this.list[i].object.length; j++) {
          //添加金额参数入对象
          if (this.list[i].object[j].flag) {
            this.list[i].object[j].money = this.money
          }else if(!this.list[i].object[j].flag){
            this.list[i].object[j].money = ''
          }
        }
      }
    },
    cancel: function(item) {
      this.modal = false;
      document.querySelector('body').style.overflow='auto'
    },
    reset: function() {
      for (let i = 0; i < this.list.length; i++) {
        for (let j = 0; j < this.list[i].object.length; j++) {
          this.list[i].object[j].money = '';
          this.list[i].object[j].flag = '';
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
<style lang="scss" src="../../../assets/css/six_four.scss" scoped></style>
