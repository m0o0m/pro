<template lang="html">
  <div class="nine">
    <div class="footer1 clearfix">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-blur="onblur_top(0)" @on-focus="onfocus_top(0)" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
        </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="content" style="margin-top: 0px">
      <ul v-for="(item,i) in lists" class='list5030 clearfix fl'>
        <li class="top">
          <ul class='clearfix'>
            <li class="one fl">生肖</li>
            <li class="two fl">赔率</li>
            <li class="three fl">金额</li>
            <!-- <li class="fl four">球号</li> -->
          </ul>
        </li>
        <li v-for="(key,k) in item.object" class="cen" :class="{'table-current':key.flag}">
          <ul class='clearfix'>
            <li class="one fl" @click="numClick(key)">
              {{key.lname}}
            </li>
            <li class="two fl" @click="numClick(key)">
              {{key.odd}}
            </li>
            <li class="three fl" @click.self="numClick(key)">
              <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength="9" ref="myfocus" style="width: 65px" @on-focus="onfocus(key)" v-model="key.money" size="small" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
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
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money"  @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
        </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-Modal :modal="modal" @cancel="cancel"></Me-Modal>
  </div>
</template>

<script>
import {Modal,Input} from 'iview';
import api from "../../../api/config";
import MeModal from "../../../share_components/bet";
import '../../../assets/css/six_twelve.scss'
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
export default {
  components: {
    MeModal,Modal,'I-Input':Input
  },
  props: {},
  data() {
    return {
      cdata: null,
      money: "",
      modal: false,
      a:'',
      arr: [],
      lists: [
        {
          object: [
            {
              lname: "单1双6",
              index: 0,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单2双5",
              index: 1,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单3双4",
              index: 2,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单4双3",
              index: 3,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单5双2",
              index: 4,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单6双1",
              index: 5,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单7双0",
              index: 6,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大0小7",
              index: 7,
              odd: null,
              flag: false,
              money: ""
              // box: []
            }
          ]
        },
        {
          object: [
            {
              lname: "大1小6",
              index: 8,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大2小5",
              index: 9,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大3小4",
              index: 10,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大4小3",
              index: 11,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大5小2",
              index: 12,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大6小1",
              index: 13,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "大7小0",
              index: 14,
              odd: null,
              flag: false,
              money: ""
              // box: []
            },
            {
              lname: "单0双7",
              index: 15,
              odd: null,
              flag: false,
              money: ""
              // box: []
            }
          ]
        }
      ]
    };
  },
  created() {
    this.fetchData();
    this.$root.$on("success", e => {
      if (e) {
        this.modal = false;
        this.reset();
      }
    });
    this.$root.$on('this_money',(e)=>{
        this.money = e
    });
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    $route: "fetchData", // 只有这个页面初始化之后，这个监听事件才开始生效
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
  mounted() {
    this.$root.$emit('no_top',false);
    this.$root.$emit("child_change", 0);
    this.$root.$on("time_out", e => {
      if (e) {
        this.fetchData(2);
      }
    });
  },
  destroyed() {
    this.$root.$off("time_out");
  },
  methods: {
    sortNumber: function(a, b) {
      return a.sort - b.sort;
    },
    fetchData: function(type) {
      this.reset();
      type == 2
        ? this.$root.$emit("loading", true, true)
        : this.$root.$emit("loading", true);
      let body = {
        fc_type: this.$route.query.page,
        gameplay: 182
      };
      api.getgameindex(this, body, res => {
          this.$root.$emit('only_back',res,type);
        if (res.data.ErrorCode == 1) {
          this.cdata = res.data.Data.c_data;
          this.$root.$emit("auto", res.data.Data.auto);
          this.$root.$emit("closetime", res.data.Data.closetime);
          this.$root.$emit("c_data", res.data.Data.c_data);
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          this.computed(back_data);
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
            }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
//          this.$root.$emit("get_closetime", res.data.Data.closetime);
        }
      });
    },
    cancel: function(item) {
      this.modal = false;
      document.querySelector("body").style.overflow = "auto";
    },
    computed: function(data, shengxiao) {
      this.$set(this.lists, this.lists);
      var k = 0;
      for (let i = 0; i < this.lists.length; i++) {
        for (let l = 0; l < this.lists[i].object.length; l++, k++) {
          Object.assign(this.lists[i].object[l], data[k]);
          var lname = this.lists[i].object[l].remark;
          lname = lname.replace(/七码#/, "");
          lname = lname.replace(/#/, "");
          this.lists[i].object[l].lname = lname;
        }
      }
    },
    numClick: function(item) {
      if (item.money && item.flag == true) {
        item.flag = false;
        item.money = "";
        this.$refs.myfocus[item.index].blur();
        this.$refs.myfocus[item.index].$refs.input.value = '';
      } else if (item.money == "" && item.flag == false) {
        this.$refs.myfocus[item.index].focus();
      } else if (item.flag == true && item.money == "") {
        item.flag = false;
        this.$refs.myfocus[item.index].$refs.input.value = '';
      } else if (item.flag == true) {
        item.money = this.money;
      }
    },
    onfocus: function(item) {
      this.$refs.myfocus[item.index].$refs.input.data_onoff = 'true';
      this.a = item.index;
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          if(i != item.index+1) {
              dom[i].data_onoff = 'false';
          }
      }
      if (item.flag == false && item.money == "") {
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
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          let b = this.lists[i].object[j].money + 'b';
          this.lists[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.lists[i].object[j].money);
          if(this.lists[i].object[j].flag){
            is_select = true
          }
        }
      }
      if(is_select){
        if (kk != 0) {
          this.modal = true;
          this.$root.$emit("c_data", this.cdata);
//          document.querySelector("body").style.overflow = "hidden";
          this.$root.$emit("id-selected", this.lists);
        } else if (kk == 0) {
          this.$Modal.warning({
            content: hint.money_null
          });
          window.setTimeout(() => {
            this.$Modal.remove();
          }, share.Prompt);
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
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          if (this.lists[i].object[j].flag) {
            this.lists[i].object[j].money = this.money;
          } else {
            this.lists[i].object[j].money = "";
          }
        }
      }
    },
    reset: function() {
      for (let i = 0; i < this.lists.length; i++) {
        for (let j = 0; j < this.lists[i].object.length; j++) {
          this.lists[i].object[j].money = "";
          this.lists[i].object[j].flag = false;
        }
      }
      this.money = "";
      this.$root.$emit("reset", "");
      let dom1 = document.querySelectorAll('input');
      for(let i = 0;i < dom1.length;i++){
          dom1[i].value = '';
          dom1[i].data_onoff = 'false';
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
    //筹码的输入事件
    add_money: function(type) {
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money();
    }
  }
};
</script>
