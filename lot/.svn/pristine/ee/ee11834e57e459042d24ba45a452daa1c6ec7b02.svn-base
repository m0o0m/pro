<template>
  <div class="eleven_opt">
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
    <div class="e_o_box">
      <ul v-for="(item,i) in OptLists">
        <li class="thend">{{item.name}}</li>
        <li v-for="(key,k) in item.object" @click="myclick(item,key,i)" :class="{'styleclect':key.flag}" class='num'>
          <i class="ball">{{key.name}}</i>
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
      <div class="one"><span class="left">金额￥</span><I-Input @on-change="change_money()" :maxlength="9" style="width: 100px" size="small" v-model="money" :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
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
      cc_time:share.Prompt,
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
  props: ["cdata", "OptLists", "OptType"],
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
  methods: {
    myclick(item, key, i) {
      this.columnClaic(item, key, i);
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
      var flag = true;
      var time = null;
      if (this.OptLists[0].arr.length == 1) {
        this.OptType[0].object[0].flag = true;
        this.OptType[0].object[0].money = this.money;
        this.OptType[0].object[0].input_name = this.OptLists[0].arr[0].name;
      }

      if (this.OptLists[1].arr.length == 2) {
        this.OptType[0].object[1].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[1].arr.length; l++) {
          input_name += this.OptLists[1].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[1].flag = true;
        this.OptType[0].object[1].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[1].input_name = input_name;
      } else if (this.OptLists[1].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[1].flag = false;
        this.$Modal.warning({
          content: "二中二少选了 1 个球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }

      if (this.OptLists[2].arr.length == 3) {
        this.OptType[0].object[2].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[2].arr.length; l++) {
          input_name += this.OptLists[2].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[2].flag = true;
        this.OptType[0].object[2].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[2].input_name = input_name;
      } else if (this.OptLists[2].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[2].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "三中三少选了 " + (3 - this.OptLists[2].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }

      if (this.OptLists[3].arr.length == 4) {
        this.OptType[0].object[3].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[3].arr.length; l++) {
          input_name += this.OptLists[3].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[3].flag = true;
        this.OptType[0].object[3].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[3].input_name = input_name;
      } else if (this.OptLists[3].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[3].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "四中四少选了 " + (4 - this.OptLists[3].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }

      if (this.OptLists[4].arr.length == 5) {
        this.OptType[0].object[4].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[4].arr.length; l++) {
          input_name += this.OptLists[4].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[4].flag = true;
        this.OptType[0].object[4].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[4].input_name = input_name;
      } else if (this.OptLists[4].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[4].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "五中五少选了 " + (5 - this.OptLists[4].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }

      if (this.OptLists[5].arr.length == 6) {
        this.OptType[0].object[5].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[5].arr.length; l++) {
          input_name += this.OptLists[5].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[5].flag = true;
        this.OptType[0].object[5].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[5].input_name = input_name;
      } else if (this.OptLists[5].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[5].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "六中五少选了 " + (6 - this.OptLists[5].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }
      if (this.OptLists[6].arr.length == 7) {
        this.OptType[0].object[6].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[6].arr.length; l++) {
          input_name += this.OptLists[6].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[6].flag = true;
        this.OptType[0].object[6].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[6].input_name = input_name;
      } else if (this.OptLists[6].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[6].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "七中五少选了 " + (7 - this.OptLists[6].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }

      if (this.OptLists[7].arr.length == 8) {
        this.OptType[0].object[7].input_name = null;
        var input_name = "";
        for (let l = 0; l < this.OptLists[7].arr.length; l++) {
          input_name += this.OptLists[7].arr[l].name;
          input_name += ",";
        }
        this.OptType[0].object[7].flag = true;
        this.OptType[0].object[7].money = this.money;
        input_name = input_name.substring(0, input_name.length - 1);
        this.OptType[0].object[7].input_name = input_name;
      } else if (this.OptLists[7].arr.length == 0) {
      } else {
        flag = false;
        this.OptType[0].object[7].flag = false;
        clearTimeout(time);
        this.$Modal.warning({
          content: "八中五少选了 " + (8 - this.OptLists[7].arr.length) + " 球",
          onOk: () => {
            clearTimeout(time);
          }
        });
        time = setTimeout(() => {
          this.$Modal.remove();
        }, this.cc_time);
      }
      // console.log(this.OptType[0].object);

      if (flag) {
        var kk = 0;
        var is_select = false;
        for (let i = 0; i < this.OptType.length; i++) {
          for (let j = 0; j < this.OptType[i].object.length; j++) {
            let b = this.OptType[i].object[j].money + 'b';
            this.OptType[i].object[j].money = b.replace(/\D/g, "");
            kk += Number(this.OptType[i].object[j].money);
            if(this.OptType[i].object[j].flag){
              is_select = true
            }
          }
        }
        // console.log("kk:" + kk);
        if(is_select){
          if (kk != 0) {
            this.modal = true;
            this.$root.$emit("c_data", this.cdata);
            this.$root.$emit("id-selected", this.OptType);
//            document.querySelector("body").style.overflow = "hidden";
          } else if (kk == 0) {
            this.$Modal.warning({
              content: hint.money_null
            });
            window.setTimeout(() => {
              this.$Modal.remove();
          }, this.cc_time);
          }
        }else{
          this.$Modal.warning({
            content: hint.all_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
          }, share.Prompt)
        }
      }
    },
    computed_money() {
      for (let i = 0; i < this.OptLists.length; i++) {
        for (let j = 0; j < this.OptLists[i].object.length; j++) {
          if (this.OptLists[i].object[j].flag) {
            this.OptLists[i].object[j].money = this.money;
          } else {
            this.OptLists[i].object[j].money = "";
          }
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
    reset: function() {
      for (let i = 0; i < this.OptLists.length; i++) {
        this.OptLists[i].arr = [];
        for (let j = 0; j < this.OptLists[i].object.length; j++) {
          this.OptLists[i].object[j].money = "";
          this.OptLists[i].object[j].flag = false;
        }
      }
      if(this.OptType[0]){
        for (let i = 0; i < this.OptType[0].object.length; i++) {
          this.OptType[0].object[i].flag = false;
          this.OptType[0].object[i].money = "";
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

    columnClaic(item, key, i) {
      var k = i;
      if (item.arr.length < k + 1) {
        var flag = true;
        for (let i = 0; i < item.arr.length; i++) {
          if (item.arr[i].name == key.name) {
            flag = !flag;
            item.arr.splice(i, 1);
            key.flag = false;
          }
        }
        if (flag) {
          item.arr.push(key);
          key.flag = true;
        }
      } else {
        var flag = true;
        for (let i = 0; i < item.arr.length; i++) {
          if (item.arr[i].name == key.name) {
            item.arr.splice(i, 1);
            key.flag = false;
            flag = !flag;
            break;
          }
        }
        if (flag) {
          this.$Modal.warning({
            content: "当前只能选 " + (k + 1) + " 个",
            onOk: () => {
              clearTimeout(time);
            }
          });
          var time = setTimeout(() => {
            this.$Modal.remove();
          }, this.cc_time);
        }
      }
    },
    add_money: function(type) {
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money();
    }
  }
};
</script>

<style lang="scss" src="../../../assets/css/eleven_opt.scss" scoped>

</style>
