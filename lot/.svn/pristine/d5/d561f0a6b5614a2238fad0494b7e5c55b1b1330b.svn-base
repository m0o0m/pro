  <template>
  <div class='denmark'>
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
        <input @blur="onblur_top(0)" @focus="onfocus_top(0)" @change="change_money()" @keyup="push_money" @afterpaste="push_money" v-model="money" class="money">
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class='box'>
      <div class='classli'>{{cdata.fc_name}}</div>
      <ul class='list2 clearfix'>
        <li v-for="(item,i) in luckyLists" class='lis_2 fl'>
          <ul>
            <li v-for="(key,k) in item.object" class='tr_2' :class="[key.flag?'table-current':'']">
              <ul class='clearfix'>
                <li class='fl cla' @click='numClick(key)'>
                  <i :class="{'ball':k<7}">
                    {{key.name}}
                    </i>
                </li>
                <li class='fl cla2' @click='numClick(key)'>
                  {{key.odd}}
                </li>
                <li class='fl cla3'>
                  <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" ref="myfocus" style="width: 45px" v-model="key.money" size="small" @on-focus="onfocus(key)" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
                </li>
              </ul>
            </li>
          </ul>
        </li>
      </ul>
    </div>
    <lucky-Right :lists='luckyLists' :money='money'></lucky-Right>
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
        <input @focus="onfocus_top(1)"  @blur="onblur_top(1)" @keyup="push_money" @afterpaste="push_money" v-model="money" class="money">
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
import api from "../../../../api/config";
import {Input, Modal} from 'iview';
import MeModal from "../../../../share_components/bet";
import luckyRight from "../luckyRight";
import hint from '../../../../share_components/hint_msg'
import share from '../../../../share_components/share'
export default {
  components: {
    MeModal,
    luckyRight,
    'I-Input':Input,
    Modal
  },
  data() {
    return {
      money: "",
      modal: false,
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
  props: ["cdata", "luckyLists"],
  created() {
    this.$root.$on("success", e => {
      if (e) {
        this.modal = false;
        this.reset();
        this.money = null;
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
    numClick: function(item) {
      if (item.money && item.flag == true) {
        item.flag = false;
        item.money = "";
        this.$refs.myfocus[item.li_id].blur();
        this.$refs.myfocus[item.li_id].$refs.input.value = '';
      } else if (item.money == "" && item.flag == false) {
        this.$refs.myfocus[item.li_id].focus();
      } else if (item.flag == true && item.money == "") {
        item.flag = false;
        this.$refs.myfocus[item.li_id].$refs.input.value = '';
      } else if (item.flag == true) {
        item.money = this.money;
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
    onfocus: function(item) {
      this.$refs.myfocus[item.li_id].$refs.input.data_onoff = 'true';
      this.a = item.li_id;
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          if(i != item.li_id+1) {
              dom[i].data_onoff = 'false';
          }
      }
      if (item.flag == false && item.money == "") {
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
    go_to: function() {
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.luckyLists.length; i++) {
        for (let j = 0; j < this.luckyLists[i].object.length; j++) {
          let b = this.luckyLists[i].object[j].money + 'b';
          this.luckyLists[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.luckyLists[i].object[j].money);
          if(this.luckyLists[i].object[j].flag){
            is_select = true
          }
        }
      }
      if(is_select){
        if (kk != 0) {
          this.modal = true;
          this.$root.$emit("c_data", this.cdata);
//          document.querySelector("body").style.overflow = "hidden";
          this.$root.$emit("id-selected", this.luckyLists);
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
      for (let i = 0; i < this.luckyLists.length; i++) {
        for (let j = 0; j < this.luckyLists[i].object.length; j++) {
          //添加金额参数入对象
          if (this.luckyLists[i].object[j].flag) {
            this.luckyLists[i].object[j].money = this.money;
          } else {
            this.luckyLists[i].object[j].money = "";
          }
        }
      }
    },
    reset: function() {
      for (let i = 0; i < this.luckyLists.length; i++) {
        for (let j = 0; j < this.luckyLists[i].object.length; j++) {
          this.luckyLists[i].object[j].money = "";
          this.luckyLists[i].object[j].flag = false;
        }
      }
      this.money = '';
      this.$root.$emit('right_config',true);
      this.$root.$emit('reset', '');
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

<style lang="scss" scoped>
@import "../../../../assets/css/function.scss";
div,
ul,
li,
span,
i {
  box-sizing: border-box;
}
.footer1 {
  width: 920px;
  float: left;
  margin-top: 10px;
  padding-right: 20px; // padding-bottom: 45px;
  .input {
    width: 85px;
    padding: 5px;
    background-color: #eeeeee;
    margin-right: 10px;
    border-radius: 5px;
    float: right;
  }
  button {
    width: 85px;
    padding: 5px;
    float: right;
    outline: none 0;
    border-radius: 5px;
    margin-top: 12px;
    cursor: pointer;
  }
  .img {
    margin-left: 136px;
    float: left;
    span {
      //margin-left: 30px;
      //height: 60px;
      float: left;
      display: block;
      cursor: pointer;
    }
    .s1 {
      height: 60px;
      width: 64px;
      background: url("../../../../assets/img/Chips.png")no-repeat 0 0px;
      background-size: 368px;
    }
    .s2 {
      height: 60px;
      width: 64px;
      background: url("../../../../assets/img/Chips.png")no-repeat -64px 0px;
      background-size: 368px;
    }
    .s3 {
      height: 60px;
      width: 64px;
      background: url("../../../../assets/img/Chips.png")no-repeat -128px 0px;
      background-size: 368px;
    }
    .s4 {
      height: 60px;
      width: 64px;
      background: url("../../../../assets/img/Chips.png")no-repeat -192px 0px;
      background-size: 368px;
    }
    .s5 {
      height: 60px;
      width: 64px;
      background: url("../../../../assets/img/Chips.png")no-repeat -256px 0px;
      background-size: 368px;
    }
    .s6 {
      height: 60px;
      width: 73px;
      background: url("../../../../assets/img/Chips.png")no-repeat -320px 0px;
      background-size: 368px;
    }
  }
  .one {
    width: 160px;
    float: left;
    background-color: #eee;
    padding: 5px;
    border-radius: 10px;
    margin-top: 12px;
    border: none;
    .money {
      width: 100px;
      height: 24px;
      border-radius: 3px;
      outline: none;
      /*color: #999;*/
      padding-left: 5px;
      border: 1px solid transparent;
    }
    .money:hover {
      border-color: #2d8cf0;
    }
    .left {
      color: #999999;
    }
  }
  .two {
    float: right;
    background-color: $bg_color;
    color: #fff;
  }
}
.box {
  width: 722px;
  padding-left: 20px;
  float: left;
  overflow: hidden;
}
.classli {
  color: #999;
  background-color: #eee;
  height: 38px;
  line-height: 38px;
  text-align: left;
  text-indent: 26px;
  border: 1px solid $border_color;
}
.lis_2 {
  width: 25%;
}
.list2 {
  border-left: 1px solid $border_color;
  border-right: 1px solid $border_color;
  border-bottom: 1px solid $border_color;
}
.tr_2 {
  border-bottom: 1px solid $border_color;
  // font-weight: 700;
  li {
    line-height: 30px;
    height: 39px;
    padding-top: 4px;
    padding-bottom: 5px;
  }
}
.cla {
  display: inline-block;
  width: 55px;
  border-right: 1px solid $border_color;
  cursor: pointer;
  color: #000;
  /*font-size: 14px;*/
}
.ball {
  display: inline-block;
  width: 30px;
  height: 30px;
  border-radius: 30px;
  line-height: 30px;
  background: $ball;
  color: #fff;
}
.cla2 {
  width: 59px;
  display: inline-block;
  margin: 0 auto;
  color: rgb(64, 64, 64);
  border-right: 1px solid $border_color;
  cursor: pointer;
}
.cla3 {
  width: 60px;
  display: inline-block;
}
.lis_2 {
  border-right: 1px solid $border_color;
}
.lis_2:last-child {
  border-right: none;
}
.lis_2:nth-child(1) .tr_2:last-child,
.lis_2:nth-child(2) .tr_2:last-child {
  border-bottom: none;
}
.table-current {
  background-color: $bg_select;
}
</style>
