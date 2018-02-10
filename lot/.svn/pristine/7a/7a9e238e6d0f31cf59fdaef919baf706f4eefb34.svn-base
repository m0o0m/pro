<template>
  <div class="integrate">
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span>
        <I-Input @on-change="change_money()" :value="money" @on-blur="onblur_top(0)"  @on-focus="onfocus_top(0)" :maxlength="9" style="width: 100px" size="small" v-model="money"
                 @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="box">
      <div class='clearfix'>
        <ul v-show="i<5" v-for="(item,i) in Ilists" class='uls'>
          <li class="thend">{{item.name}}</li>
          <li>
            <ul>
              <li v-for="(key,k) in item.object" :class="[key.flag?'styleclect':'']" class='list clearfix'>
                <span class="one fl" @click="myclick(key)">
                  <i :class="{'ball':k<10&&i<5}">
                  {{key.name}}
                  </i>
                  </span>
                <span class="tow fl" @click="myclick(key)">{{key.odd}}</span>
                <span class="three" @click.self="myclick(key)">
                   <I-Input @on-keydown="tab_now(key)" @on-change="onchange(key)" :value="key.money" @on-blur="onblur(key)" :maxlength="9" ref="myfocus"
                            style="width: 55px" @on-focus="onfocus(key)" v-model="key.money" size="small"
                            @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
                </span>
              </li>
            </ul>
          </li>
        </ul>
      </div>
      <p style="background: #eee;padding: 10px 0;">总和</p>
      <ul v-show="i>4" v-for="(item,i) in Ilists" class='uls1'>
        <li>
          <ul>
            <li v-for="(key,k) in item.object" :class="[key.flag?'styleclect':'']" class='list clearfix'>
                <span class="one fl" @click="myclick1(key)">
                  {{key.name}}
                  </span>
              <span class="tow fl" @click="myclick1(key)">{{key.odd}}</span>
              <span class="three" @click.self="myclick1(key)">
                   <I-Input @on-keydown="tab_now(key)" @on-change="onchange(key)" :value="key.money" @on-blur="onblur1(key)" :maxlength="9" ref="myfocus1"
                            style="width: 55px" @on-focus="onfocus1(key)" v-model="key.money" size="small"
                            @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
              </span>
            </li>
          </ul>
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
      <div class="one"><span class="left">金额￥</span>
        <I-Input @on-change="change_money()" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" :maxlength="9" style="width: 100px" size="small" v-model="money"
                 @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
  import api from "../../../../api/config";
  import MeModal from "../../../../share_components/bet";
  import {Input, Modal} from 'iview';
  import hint from '../../../../share_components/hint_msg'
  export default {
    components: {MeModal, 'I-Input': Input, Modal},
    data() {
      return {
        money: "",
        modal: false,
        a:''
      };
    },
    created() {
      this.$root.$on("success", e => {
        if (e) {
          this.modal = false;
          this.reset();
        }
      });
      this.reset();
      this.$root.$on('this_money',(e)=>{
          this.money = e
      });
    },
    mounted(){
      this.$root.$emit('no_top', false);
    },
    watch: {
      'route': 'reset',
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
      myclick: function (item) {
        console.log(item);
        console.log(item.flag);
        console.log(this.$refs.myfocus1);
        if (item.money && item.flag == true) {
          item.flag = false;
          item.money = '';
          this.$refs.myfocus[item.li_id].blur();
          this.$refs.myfocus[item.li_id].$refs.input.value = '';
        } else if (item.money == '' && item.flag == false) {
          this.$refs.myfocus[item.li_id].focus();
        } else if (item.flag == true && item.money == '') {
          item.flag = false;
          this.$refs.myfocus[item.li_id].$refs.input.value = '';
        } else if (item.flag == true) {
          item.money = this.money;
        }
      },
      myclick1: function(item) {
          console.log(this.$refs.myfocus1);
          console.log(item);
          if(item.money && item.flag == true){
              item.flag = false;
              item.money = '';
              this.$refs.myfocus1[item.li_id].blur();
              this.$refs.myfocus1[item.li_id].$refs.input.value = '';
          }else if(item.money == '' && item.flag == false){
              this.$refs.myfocus1[item.li_id].focus();
          }else if (item.flag == true && item.money == ''){
              item.flag = false;
              this.$refs.myfocus1[item.li_id].$refs.input.value = '';
          }else if (item.flag == true){
              item.money = this.money;
          }
      },
      onfocus: function (item) {
        this.$refs.myfocus[item.li_id].$refs.input.data_onoff = 'true';
        this.a = item.li_id;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.li_id+1) {
                dom[i].data_onoff = 'false';
            }
        }
        if (item.flag == false && item.money == '') {
          item.flag = true;
          item.money = this.money;
        }
      },
        onfocus1: function (item) {
            this.$refs.myfocus1[item.li_id].$refs.input.data_onoff = 'true';
            this.a = item.li_id;
            let k = 0;
            switch (item.li_id){
                case 70:
                    k = 145;
                    break;
                case 71:
                    k = 146;
                    break;
                case 72:
                    k = 147;
                    break;
                case 73:
                    k = 148;
                    break;
            }
            console.log('索引：'+item.li_id);
            let dom = document.querySelectorAll('input');
            for(let i = 0;i < dom.length;i++){
                if(i != k) {
                    dom[i].data_onoff = 'false';
                }
            }
            if (item.flag == false && item.money == '') {
                item.flag = true;
                item.money = this.money;
            }
        },
        onblur: function (key) {
            key.money = this.$refs.myfocus[key.li_id].$refs.input.value;
            console.log('选中后的金额：'+key.money);
        },
        onblur1: function (key) {
            key.money = this.$refs.myfocus1[key.li_id].$refs.input.value;
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
      cancel: function (item) {
        this.modal = false;
        document.querySelector("body").style.overflow = "auto";
      },
      push_money: function () {
        this.money = this.money.replace(/\D/g, "");
        this.computed_money();
      },
      change_money: function () {
        this.computed_money()
      },
      add_money: function (type) {
        let money = this.money;
        this.money = Number(money) + type;
        this.computed_money();
      },
      //处理金额
      computed_money: function () {
        for (let i = 0; i < this.Ilists.length; i++) {
          for (let j = 0; j < this.Ilists[i].object.length; j++) {
            //添加金额参数入对象
            if (this.Ilists[i].object[j].flag) {
              this.Ilists[i].object[j].money = this.money;
            } else if (!this.Ilists[i].object[j].flag) {
              this.Ilists[i].object[j].money = "";
            }
          }
        }
      },
      go_to: function () {
        let a = this.money + 'a';
        this.money = a.replace(/\D/g, "");
        var kk = 0;
        var is_select = false;
        for (let i = 0; i < this.Ilists.length; i++) {
          for (let j = 0; j < this.Ilists[i].object.length; j++) {
            let b = this.Ilists[i].object[j].money + 'b';
            this.Ilists[i].object[j].money = b.replace(/\D/g, "");
            kk += Number(this.Ilists[i].object[j].money);
            if(this.Ilists[i].object[j].flag){
              is_select = true
            }
          }
        }
        // console.log("kk:" + kk);
        if(is_select){
          if (kk != 0) {
            this.modal = true;
            this.$root.$emit("c_data", this.cdata);
            this.$root.$emit("id-selected", this.Ilists);
//            document.querySelector("body").style.overflow = "hidden";
          } else if (kk == 0) {
            this.$Modal.warning({
              content: hint.money_null
            });
            window.setTimeout(() => {
              this.$Modal.remove()
          }, 500)
          }
        }else{
          this.$Modal.warning({
            content: hint.all_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
        }, 500)
        }
      },
      reset: function () {
        for (let i = 0; i < this.Ilists.length; i++) {
          for (let j = 0; j < this.Ilists[i].object.length; j++) {
            this.Ilists[i].object[j].money = "";
            this.Ilists[i].object[j].flag = "";
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
      tab_now: function (key) {
        if (!key.money) {
          key.flag = false;
        } else if (key.money) {
          key.flag = true;
        }
      },
      //监听input值得变化
      onchange: function (key) {
        if (!key.money) {
          key.flag = false;
        } else if (key.money) {
          key.flag = true;
        }
      },
      //每个球对应的输入框
      gogo: function (key) {
        key.money = key.money.replace(/\D/g, "");
      }
    },
    props: ["Ilists", "cdata"]
  };
</script>

<style lang="scss" src="../../../../assets/css/ffc_one.scss" scoped></style>
