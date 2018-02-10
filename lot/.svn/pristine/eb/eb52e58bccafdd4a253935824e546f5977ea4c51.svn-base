<template lang="html">
  <div class="nine">
    <div class="nav">
      <ul>
        <li v-for="(item,i) in menus" :class="{'item_class':selectItem==i}" @click="type(item,i)">{{item.name}}</li>
      </ul>
    </div>
    <div class="content">
      <ul v-for="(item,i) in lists[selectItem]" class='list clearfix fl'>
        <li class="top">
          <ul class='clearfix'>
            <li class="fl">{{item.name}}</li>
            <li class="fl">赔率</li>
            <li class="fl three">金额</li>
            <li class="fl four">球号</li>
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
              <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength="9" ref="myfocus" style="width: 45px" @on-focus="onfocus(key)" v-model="key.money" size="small" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)"></I-Input>
            </li>
            <li class="four fl" @click="numClick(key)"><span v-for="yan in key.box" class="ball" :style="{background:yan.color}">{{yan.num}}</span></li>
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
// import navCenter from "./components/nav";
import api from "../../../api/config";
import MeModal from "../../../share_components/bet";
import {Modal,Input} from 'iview';
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
      a:'',
      showA: true,
      showB: false,
      // name: null,
      cdata: null,
      money: "",
      modal: false,
      // leng: 2,
      selectItem: 0,
      // arr: [],
      // thisYear: null,
      menus: [
        {
          name: "一肖",
          type: 172
        },
        {
          name: "尾数",
          type: 173
        }
      ],
      lists: [
        [
          {
            name: "生肖",
            object: [
              {
                lname: "鼠",
                index: 0,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "牛",
                index: 1,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "虎",
                index: 2,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "兔",
                index: 3,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "龙",
                index: 4,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "蛇",
                index: 5,
                odd: null,
                flag: false,
                money: "",
                box: []
              }
            ]
          },
          {
            object: [
              {
                lname: "马",
                index: 6,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "羊",
                index: 7,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "猴",
                index: 8,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "鸡",
                index: 9,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "狗",
                index: 10,
                odd: null,
                flag: false,
                money: "",
                box: []
              },
              {
                lname: "猪",
                index: 11,
                odd: null,
                flag: false,
                money: "",
                box: []
              }
            ]
          }
        ],
        [
          {
            name: "号码",
            object: [
              {
                lname: "0尾",
                name: "0",
                index: 0,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 10, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 20, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 30, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 40, color: "linear-gradient(#ee4c19, #bd2706)" }
                ]
              },
              {
                lname: "1尾",
                name: "1",
                index: 1,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 1, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 11, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 21, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 31, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 41, color: "linear-gradient(#2991d2, #1c6196)" }
                ]
              },
              {
                lname: "2尾",
                name: "2",
                index: 2,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 2, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 12, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 22, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 32, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 42, color: "linear-gradient(#2991d2, #1c6196)" }
                ]
              },
              {
                lname: "3尾",
                name: "3",
                index: 3,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 3, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 13, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 23, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 33, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 43, color: "linear-gradient(#3ec948, #026d09)" }
                ]
              },
              {
                lname: "4尾",
                name: "4",
                index: 4,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 4, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 14, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 24, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 34, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 44, color: "linear-gradient(#3ec948, #026d09)" }
                ]
              }
            ]
          },
          {
            object: [
              {
                lname: "5尾",
                name: "5",
                index: 5,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 5, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 15, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 25, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 35, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 45, color: "linear-gradient(#ee4c19, #bd2706)" }
                ]
              },
              {
                lname: "6尾",
                name: "6",
                index: 6,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 6, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 16, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 26, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 36, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 46, color: "linear-gradient(#ee4c19, #bd2706)" }
                ]
              },
              {
                lname: "7尾",
                name: "7",
                index: 7,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 7, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 17, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 27, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 37, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 47, color: "linear-gradient(#2991d2, #1c6196)" }
                ]
              },
              {
                lname: "8尾",
                name: "8",
                index: 8,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 8, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 18, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 28, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 38, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 48, color: "linear-gradient(#2991d2, #1c6196)" }
                ]
              },
              {
                lname: "9尾",
                name: "9",
                index: 9,
                odd: null,
                flag: false,
                money: "",
                box: [
                  { num: 9, color: "linear-gradient(#2991d2, #1c6196)" },
                  { num: 19, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 29, color: "linear-gradient(#ee4c19, #bd2706)" },
                  { num: 39, color: "linear-gradient(#3ec948, #026d09)" },
                  { num: 49, color: "linear-gradient(#3ec948, #026d09)" }
                ]
              }
            ]
          }
        ]
      ],
      gettypa: 172
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
    this.$root.$emit('no_top',true);
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
      if(type == 2){
         this.gettypa = 172
      }
      let body = {
        fc_type: this.$route.query.page,
        gameplay: this.gettypa
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
          this.gettypa == 172
            ? this.computed(back_data, res.data.shengxiao)
            : this.computed1(back_data);
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
            }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
        }
      });
    },
    computed: function(data, shengxiao) {
      this.$set(this.lists, this.lists);
      var k = 0;
      for (var i = 0; i < this.lists[0].length; i++) {
        //  console.log(k)
        for (var l = 0; l < this.lists[0][i].object.length; l++, k++) {
          Object.assign(this.lists[0][i].object[l], data[k]);
          switch (this.lists[0][i].object[l].lname) {
            case "鼠":
              this.lists[0][i].object[l].name = "mouse";
              this.lists[0][i].object[l].box = shengxiao.mouse;
              break;
            case "牛":
              this.lists[0][i].object[l].name = "cattle";
              this.lists[0][i].object[l].box = shengxiao.cattle;
              break;
            case "虎":
              this.lists[0][i].object[l].name = "tiger";
              this.lists[0][i].object[l].box = shengxiao.tiger;
              break;
            case "兔":
              this.lists[0][i].object[l].name = "rabbit";
              this.lists[0][i].object[l].box = shengxiao.rabbit;
              break;
            case "龙":
              this.lists[0][i].object[l].name = "dragon";
              this.lists[0][i].object[l].box = shengxiao.dragon;
              break;
            case "蛇":
              this.lists[0][i].object[l].name = "snake";
              this.lists[0][i].object[l].box = shengxiao.snake;
              break;
            case "马":
              this.lists[0][i].object[l].name = "horse";
              this.lists[0][i].object[l].box = shengxiao.horse;
              break;
            case "羊":
              this.lists[0][i].object[l].name = "sheep";
              this.lists[0][i].object[l].box = shengxiao.sheep;
              break;
            case "猴":
              this.lists[0][i].object[l].name = "monkey";
              this.lists[0][i].object[l].box = shengxiao.monkey;
              break;
            case "鸡":
              this.lists[0][i].object[l].name = "chicken";
              this.lists[0][i].object[l].box = shengxiao.chicken;
              break;
            case "狗":
              this.lists[0][i].object[l].name = "dog";
              this.lists[0][i].object[l].box = shengxiao.dog;
              break;
            case "猪":
              this.lists[0][i].object[l].name = "pig";
              this.lists[0][i].object[l].box = shengxiao.pig;
              break;
          }
          for (var v = 0; v < this.lists[0][i].object[l].box.length; v++) {
            switch (this.lists[0][i].object[l].box[v].color) {
              case "blue":
                this.lists[0][i].object[l].box[v].color =
                  "linear-gradient(#2991d2, #1c6196)";
                break;
              case "green":
                this.lists[0][i].object[l].box[v].color =
                  "linear-gradient(#3ec948, #026d09)";
                break;
              case "red":
                this.lists[0][i].object[l].box[v].color =
                  "linear-gradient(#ee4c19, #bd2706)";
                break;
            }
          }
        }
      }
    },
    computed1: function(data) {
      this.$set(this.lists, this.lists);

      var k = 0;
      for (let i = 0; i < this.lists.length; i++) {
        for (let l = 0; l < this.lists[1].length; l++) {
          for (var n = 0; n < this.lists[1][l].object.length; n++, k++) {
            Object.assign(this.lists[1][l].object[n], data[k]);
          }
        }
      }
    },
    cancel: function(item) {
      this.modal = false;
      document.querySelector("body").style.overflow = "auto";
    },
    numClick: function(item) {
      this.$refs.myfocus[item.index].focus();
      item.flag = !item.flag;
      if (item.flag) {
        item.money = this.money;
      } else {
        item.money = "";
        this.$refs.myfocus[item.index].blur();
      }
    },
    type(key, i) {
      this.reset();
      this.selectItem = i;
      this.gettypa = key.type;
      this.fetchData();
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
          if(i != item.index) {
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
      for (let l = 0; l < this.lists[this.selectItem].length; l++) {
        for (var n = 0; n < this.lists[this.selectItem][l].object.length; n++) {
          let b = this.lists[this.selectItem][l].object[n].money + 'b';
          this.lists[this.selectItem][l].object[n].money = b.replace(/\D/g, "");
          kk += Number(this.lists[this.selectItem][l].object[n].money);
          if(this.lists[this.selectItem][l].object[n].flag){
            is_select = true
          }
        }
      }
      if(is_select){
        if (kk != 0) {
          this.modal = true;
          this.$root.$emit("c_data", this.cdata);
//          document.querySelector("body").style.overflow = "hidden";
          this.$root.$emit("id-selected", this.lists[this.selectItem]);
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
      for (let i = 0; i < this.lists[this.selectItem].length; i++) {
        for (let j = 0; j < this.lists[this.selectItem][i].object.length; j++) {
          if (this.lists[this.selectItem][i].object[j].flag) {
            this.lists[this.selectItem][i].object[j].money = this.money;
          } else {
            this.lists[this.selectItem][i].object[j].money = "";
          }
        }
      }
    },
    reset() {
      for (let l = 0; l < this.lists[this.selectItem].length; l++) {
        for (var n = 0; n < this.lists[this.selectItem][l].object.length; n++) {
          this.lists[this.selectItem][l].object[n].money = "";
          this.lists[this.selectItem][l].object[n].flag = false;
        }
      }
      this.money = "";
      this.$root.$emit('reset', '');
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
