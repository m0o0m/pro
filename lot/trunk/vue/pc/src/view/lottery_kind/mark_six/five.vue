<template lang="html">
  <div class="six_one">
    <div class="content">
      <div class="top">
        <ul>
          <li class="only">项目</li>
          <li class="one">大</li>
          <li class="one">小</li>
          <li class="one">单</li>
          <li class="one">双</li>
          <li class="one">红波</li>
          <li class="one">蓝波</li>
          <li class="one">绿波</li>
        </ul>
      </div>
      <div class="table">
        <div class="bottom" v-for="item in lists">
          <ul v-for="key in item.object">
            <li class="one"  @click="pour(key)" v-show="key.num">{{key.num}}</li>
            <li  @click="pour(key)" :class="[key.number?'two1':'border2_none']">{{key.odd}}</li>
            <li  @click.self="pour(key)" :class="[key.number?'three1':'border3_none']">
              <Radio size="large" name="key.index" v-model="key.flag"></Radio>
              <div class="nothing_do" @click.self="pour(key)"></div>
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
    <Me-modal :modal='modal' @cancel="cancel"></Me-modal>
  </div>
</template>

<script>
import {Modal,Input,Radio} from 'iview';
import api from '../../../api/config'
import MeModal from '../../../share_components/bet_five'
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
export default {
  components: {
    MeModal,Modal,'I-Input':Input,Radio
  },
  data() {
    return {
      a:'',
      modal: false,
      money:'',
      picked: false,
      click_number:0,
      lists:[
        {
          name:'1',object:[
            {
              index:0,num:'正码1',flag:false,type:1,only:1,number:'48.8',money:''
            },
            {
              index:1,num:'正码2',flag:false,type:2,only:1,number:'48.8',money:''
            },
            {
              index:2,num:'正码3',flag:false,type:3,only:1,number:'48.8',money:''
            },
            {
              index:3,num:'正码4',flag:false,type:4,only:1,number:'48.8',money:''
            },
            {
              index:4,num:'正码5',flag:false,type:5,only:1,number:'48.8',money:''
            },
            {
              index:5,num:'正码6',flag:false,type:6,only:1,number:'48.8',money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:6,number:'48.8',type:1,only:1,flag:false,money:''
            },
            {
              index:7,number:'48.8',type:2,only:1,flag:false,money:''
            },
            {
              index:8,number:'48.8',type:3,only:1,flag:false,money:''
            },
            {
              index:9,number:'48.8',type:4,only:1,flag:false,money:''
            },
            {
              index:10,number:'48.8',type:5,only:1,flag:false,money:''
            },
            {
              index:11,number:'48.8',type:6,only:1,flag:false,money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:12,number:'48.8',type:1,only:2,flag:false,money:''
            },
            {
              index:13,number:'48.8',type:2,only:2,flag:false,money:''
            },
            {
              index:14,number:'48.8',type:3,only:2,flag:false,money:''
            },
            {
              index:15,number:'48.8',type:4,only:2,flag:false,money:''
            },
            {
              index:16,number:'48.8',type:5,only:2,flag:false,money:''
            },
            {
              index:17,number:'48.8',type:6,only:2,flag:false,money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:18,number:'48.8',type:1,only:2,flag:false,money:''
            },
            {
              index:19,number:'48.8',type:2,only:2,flag:false,money:''
            },
            {
              index:20,number:'48.8',type:3,only:2,flag:false,money:''
            },
            {
              index:21,number:'48.8',type:4,only:2,flag:false,money:''
            },
            {
              index:22,number:'48.8',type:5,only:2,flag:false,money:''
            },
            {
              index:23,number:'48.8',type:6,only:2,flag:false,money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:24,number:'48.8',type:1,only:3,flag:false,money:''
            },
            {
              index:25,number:'48.8',type:2,only:3,flag:false,money:''
            },
            {
              index:26,number:'48.8',type:3,only:3,flag:false,money:''
            },
            {
              index:27,number:'48.8',type:4,only:3,flag:false,money:''
            },
            {
              index:28,number:'48.8',type:5,only:3,flag:false,money:''
            },
            {
              index:29,number:'48.8',type:6,only:3,flag:false,money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:30,number:'48.8',type:1,only:3,flag:false,money:''
            },
            {
              index:31,number:'48.8',type:2,only:3,flag:false,money:''
            },
            {
              index:32,number:'48.8',type:3,only:3,flag:false,money:''
            },
            {
              index:33,number:'48.8',type:4,only:3,flag:false,money:''
            },
            {
              index:34,number:'48.8',type:5,only:3,flag:false,money:''
            },
            {
              index:35,number:'48.8',type:6,only:3,flag:false,money:''
            }
          ]
        },
        {
          name:'2',object:[
            {
              index:36,number:'48.8',type:1,only:3,flag:false,money:''
            },
            {
              index:37,number:'48.8',type:2,only:3,flag:false,money:''
            },
            {
              index:38,number:'48.8',type:3,only:3,flag:false,money:''
            },
            {
              index:39,number:'48.8',type:4,only:3,flag:false,money:''
            },
            {
              index:40,number:'48.8',type:5,only:3,flag:false,money:''
            },
            {
              index:41,number:'48.8',type:6,only:3,flag:false,money:''
            }
          ]
        },
      ],
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
  watch: {
      // 如果路由有变化，会再次执行该方法
      '$route': 'fetchData', // 只有这个页面初始化之后，这个监听事件才开始生成
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
  mounted(){
    this.$root.$emit('no_top',true);
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
  methods:{
    sortNumber: function(a,b){
      return a.sort - b.sort
    },
    fetchData: function(type){
      this.reset();
      type==2?this.$root.$emit('loading',true,true):this.$root.$emit('loading',true);
      this.$root.$emit('loading',true);
      let body = {
        'fc_type': this.$route.query.page,
        'gameplay':169
      };
      api.getgameindex(this, body, (res) => {
          this.$root.$emit('only_back',res,type);
        if (res.data.ErrorCode == 1) {
          console.log('success');
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
            }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          this.computed(back_data);
          this.$root.$emit('c_data',res.data.Data.c_data);
//          this.$root.$emit('get_closetime',res.data.Data.closetime);
        }
      });
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
    go_to:function () {
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      var le = 0;//选择了多少项
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          let b = this.lists[i].object[j].money + 'b';
          this.lists[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.lists[i].object[j].money);
          if(this.lists[i].object[j].flag){
            le += 1;
            is_select = true
          }
        }
      }
      if(is_select){
        if (kk != 0 && le >= 2 && le <= 8) {
          this.$root.$emit('id-selected-five',this.lists,this.money);
          this.modal = true;
//          document.querySelector('body').style.overflow='hidden';
        }else if (kk == 0) {
          this.$Modal.warning({
            content: hint.money_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
        }, share.Prompt)
        }else if(le < 2){
          this.$Modal.warning({
            content: '请选择2-8组玩法，若只要单一下注请前往正特投注！'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
        }, share.Prompt)
        }else if(le > 8){
          this.$Modal.warning({
            content: '最多只能选8组玩法!'
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
    cancel: function(item) {
      this.modal = false;
      document.querySelector('body').style.overflow='auto'
    },
    add_money: function(type){
      let money = this.money;
      this.money = Number(money)+type;
      this.computed_money()
    },
    push_money: function() {
      // this.disabled = true;
      this.money = this.money.replace(/\D/g, "");
      this.computed_money()
    },
    change_money: function () { this.computed_money() },
    computed_money: function(){
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          //添加金额参数入对象
          if (this.lists[i].object[j].flag) {
            this.lists[i].object[j].money = this.money
          }else if(!this.lists[i].object[j].flag){
            this.lists[i].object[j].money = ''
          }
        }
      }
    },
    computed: function(data){
      this.$set(this.lists,this.lists);
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[0],data[i])
      }
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[1],data[7+i])
      }
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[2],data[14+i])
      }
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[3],data[21+i])
      }
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[4],data[28+i])
      }
      for(let i = 0;i<this.lists.length;i++){
        Object.assign(this.lists[i].object[5],data[35+i])
      }
    },
    //重置
    reset: function () {
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          this.lists[i].object[j].money = '';
          this.lists[i].object[j].flag = false;
        }
      }
      this.money = '';
      this.click_number = 0;
      this.$root.$emit('reset', '');
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          dom[i].value = '';
          dom[i].data_onoff = 'false';
      }
      this.$root.$emit('clear_key_number','')
    },
    pour: function(item){
      console.log(item);
      this.click_number = 0;
      let remark1 = item.remark.split('#')[1];//项目名
      for(let i = 0;i<this.lists.length;i++){
        for(let j = 0;j<this.lists[i].object.length;j++){
          let remark = this.lists[i].object[j].remark.split('#')[1];//项目名
          if(item.only == this.lists[i].object[j].only && remark1 == remark){
            this.lists[i].object[j].flag = false;
            item.flag = true;
            item.money = this.money;
            this.lists[i].object[j].money = '';
          }else if(this.lists[i].object[j].flag){
            this.click_number += 1
          }else{
            item.flag = true;
            item.money = this.money;
          }
        }
      }
//      this.unique(this.click_number);
      console.log('选中了多少个：'+this.click_number);
      if(this.click_number > 7){
        item.flag = false;
        this.$Modal.warning({
          content: '最多只能选8组玩法!'
        });
        window.setTimeout(() => {
          this.$Modal.remove()
        }, share.Prompt)
      }
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
<style lang="scss" src="../../../assets/css/six_five.scss" scoped></style>
