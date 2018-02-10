<template>
  <div class="or">
    <div class="footer1">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :value="money" @on-blur="onblur_top(0)"  @on-focus="onfocus_top(0)" :maxlength="9" style="width: 100px" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="or_box">
      <ul v-for="item in two">
        <li class="thend1" @click="choose(item)">
        <input ref="radio_choose" type="radio" @click="choose(item)" name="contact" :value="item.index" v-model="picked"/>
        </li>
        <li class="thend2" @click="choose(item)">{{item.remark}}</li>
        <li class="thend3" @click="choose(item)">{{item.odd}}</li>
        <li @click="myclick(key)" :class="{'bg_color':key.state}" v-for="(key,i) in item.object">
          <span class="one"><i>{{key.num}}</i></span>
          <span class="tow" style="position:relative">
            <!-- <Checkbox style="height:40px" v-model="key.item" ref="myfocus">checkbox</Checkbox> -->
            <input type="checkbox" @click="myclick(key)" v-model="key.state"/>
            <div style="position: absolute;top: 10px;left: 15px;width: 20px;height: 20px;">
            </div>
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
      <div class="one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" :maxlength="9" style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
// import back_data from './back_data'
import {Input, Modal} from 'iview';
import api from '../../../api/config'
import MeModal from '../../../share_components/bet_happy'
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
  export default {
components: { MeModal,'I-Input':Input,Modal},
    data() {
      return {
        money: '',
        modal: false,
        // my_data:'',
        arr:[],//个数计算数据,及选中保存下来的数组
        picked:0,//当选框默认选定
        a:''
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
      this.$root.$emit('no_top', false);
    },
    props:{
      two:{
        type:Array
      },
      c_data:{
        type:Object
      },
      only_data:{
        type:Array
      }
    },
    watch: {
      // 如果路由有变化，会再次执行该方法
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
    methods:{
      fetchData: function(){
        this.reset();
        this.first_choose()
      },
      go_to: function() {
        this.$root.$emit('c_data',this.c_data);
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        for(let i = 0;i < this.two.length;i++){
          for(let j = 0;j < this.two[i].object.length;j++){
            if(this.two[i].object[j].state){
              let b = this.two[i].object[j].money + 'b';
              this.two[i].object[j].money = b.replace(/\D/g, "");
            }
          }
        }
        if(this.picked == 0 && this.arr.length < 1){
          this.$Modal.warning({
            content: '至少选1个'
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
        }else if(this.picked == 4 && this.arr.length < 5){
          this.$Modal.warning({
            content: '至少选5个'
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }else{
          this.$root.$emit('id-selected-h',this.two);
          if(this.money){
            this.modal = true;
//            document.querySelector('body').style.overflow='hidden';
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
      change_money: function () { this.computed_money() },
      add_money: function(type){
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money()
      },
      computed_money: function () {
        for(let i = 0;i < this.two.length;i++){
          for(let j = 0;j < this.two[i].object.length;j++){
            if(this.two[i].object[j].state){
              Object.assign(this.two[i].object[j],{'money':this.money});
            }
          }
        }
      },
      cancel: function(item) {
        this.modal = false;
        document.querySelector('body').style.overflow='auto'
      },
      myclick: function(key){
        key.state = !key.state;
        if(key.state){
          this.computed_money()
        }
        console.log(key);
        var max_limit = [];
        if(this.$route.query.page == 'bj_kl8'){
          max_limit = [10,6,6,7,8];
        }else if (this.$route.query.page == 'dm_klc') {
          max_limit = [4,5,6,7,8];
        }else if (this.$route.query.page == 'jnd_bs') {
          max_limit = [10,6,6,7,8];
        }
        if(key.state){
          this.arr.push(key);
          this.arr = this.unique(this.arr);
          if(this.picked == 0 && this.arr.length>max_limit[0]){
            this.arr.pop();
            // console.log(this.arr);
            key.state = false;
            // console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[0]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 1 && this.arr.length > max_limit[1]){
            this.arr.pop();
            // console.log(this.arr);
            key.state = false;
            // console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[1]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 2 && this.arr.length > max_limit[2]){
            this.arr.pop();
            // console.log(this.arr);
            key.state = false;
            // console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+max_limit[2]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 3 && this.arr.length >  max_limit[3]){
            this.arr.pop();
            // console.log(this.arr);
            key.state = false;
            // console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+ max_limit[3] +'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }else if(this.picked == 4 && this.arr.length >  max_limit[4]){
            this.arr.pop();
            // console.log(this.arr);
            key.state = false;
            // console.log('总数：'+this.arr.length);
            this.$Modal.warning({
              content: '只允许选中'+ max_limit[4]+'个'
            });
            window.setTimeout(() => {
              this.$Modal.remove()
            }, share.Prompt)
          }
        }else{
          for(let i = 0;i<this.arr.length;i++){
            // console.log('数组的所有对象的index：'+this.arr[i].index);
            // console.log('取消的index：'+key.index);
            if(key.index == this.arr[i].index){
              this.arr.splice(i,1);
              // console.log(this.arr)
              break;
            }
          }
        }
      },
      choose: function(key){
        this.reset();
        //console.log('index:'+key.index);
        this.picked = key.index;
        for(let i = 0;i < this.two.length;i++){
          for(let j = 0;j < this.two[i].object.length;j++){
            if(!key.state){
              this.two[i].object[j].state = false
            }else{
              this.two[i].object[j].state = true;
            }
            Object.assign(this.two[i].object[j], this.only_data[key.index]);
          }
        }
      },
      first_choose: function(){
        this.picked = 0;
        for(let i = 0;i < this.two.length;i++){
          for(let j = 0;j < this.two[i].object.length;j++){
            Object.assign(this.two[i].object[j], this.only_data[0]);
          }
        }
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
      reset: function(){
        this.money = '';
        this.arr = [];
        for(let i = 0;i < this.two.length;i++){
          for(let j = 0;j<this.two[i].object.length;j++){
            this.two[i].object[j].state = false;
            this.two[i].object[j].money = '';
          }
        }
        this.$root.$emit("reset","");
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
<style lang="scss" src="../../../assets/css/happy_or.scss" scoped></style>
