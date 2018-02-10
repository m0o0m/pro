<template lang="html">
  <div class="bet" v-show="modal">
    <div class="pk_bet">
    </div>
    <div class="modal">
      <div class="header">
        <h3>下注金额</h3>
      </div>
      <p>当前投注：</p>
      <div class="bet_center">
        <p class="hp_center_top">
          <span class="one">类型</span>
          <span class="two">号码</span>
          <span class="three">赔率</span>
        </p>
        <div ref="tail_modal" class="tail_modal" style="overflow:auto;height: 145px;">
          <p class="hp_center_content" v-for="item in lists" v-if="item.money">
            <span class="one">{{item.remark.split('#')[0]}}</span>
            <span class="two" v-if="item.gameplay == 'AllMiss'">
              <i class="box">{{item.lname}}</i>
            </span>
            <span class="two" v-else>
              {{item.lname}}
            </span>
            <span class="three">{{item.odd}}</span>
          </p>
        </div>
      </div>
      <div style="text-align: center;padding: 10px 0">
        <span style="margin-right:10px">共选：{{number}}个号码</span>
        <span style="margin-right:10px">复式共分为：{{team}}组</span>
        <span>金额合计：{{money}}</span>
      </div>
      <div style="padding: 10px 0">
        <button type="button" v-if="!loading" class="bet_bottom color1" @click="pour()">下注</button>
        <button type="button" v-else class="bet_bottom color1">下单中...</button>
        <button type="button" class="bet_bottom color2 font" @click="cancel()">取消</button>
      </div>
    </div>
  </div>
</template>

<script>
import api from "../api/config";
import algorithm from "../assets/js/split";
import deepcopy from "../assets/js/deepcopy";
import share from './share'
import mymousewheel from '../assets/js/mousewheel'
import {Modal} from 'iview';
import '../assets/css/bet.scss';
export default {
  components: {Modal},
  props: {
    modal: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      number: "", //注单量
      loading: false,
      lists: [],
      arr: [],
      qishu: "",
      money: 0,
      team: 0,
      go_pour: []
      // indexxxx:0
    };
  },
  created() {
    this.$root.$on("id-selected-tail", (e, fn) => {
      // console.log(e);
      for (let i = 0; i < e.length; i++) {
        for (let j = 0; j < e[i].object.length; j++) {
          //添加金额参数入对象
          //          console.log( e[i].object);
          let money = e[i].object[j].money;
          this.arr.push(
            Object.assign(e[i].object[j], {
              money: money
            })
          );
        }
      }
      if (this.arr) {
        this.lists = this.unique(this.arr);
        let arr = [];
        for (let i = 0; i < this.lists.length; i++) {
          if (this.lists[i].flag) {
            arr.push(this.lists[i]);
          }
        }
        this.lists = arr;
        this.number = arr.length;
        var moj = 0;
        //计算金额
        for (let l = 0; l < this.lists.length; l++) {
          if (this.lists[l].money) {
            moj = Number(this.lists[l].money);
            this.money = moj;
          }
        }
        this.computed(fn);
      }
    });
    this.$root.$on("reset", e => {
      this.money = e;
    });
    this.$root.$on("c_data", e => {
      this.qishu = e.qishu;
    });
  },
  mounted() {
    let tail_modal = document.querySelector('.tail_modal');
    mymousewheel(tail_modal);
  },
  methods: {
    choose_box: function() {
      var box_arr = [];
      for (let i = 0; i < this.lists.length; i++) {
        if(this.lists[i].flag) {
          box_arr.push(this.lists[i].name);
        }
      }
      return box_arr;
    },
    //排列组合输出二维数组
    computed: function(fn) {
      console.log(this.lists);
//      var money_data = {};
//      for(let i=0;i<this.lists.length;i++){
//        if(this.lists[i].box.length>4){
//          money_data = deepcopy(this.lists[i], {});
//          break;
//        }else{
//          money_data = deepcopy(this.lists[0], {});
//        }
//      }
//      console.log(money_data);
      var money_data = deepcopy(this.lists[0], {});
//      money_data.box = [];
      //抽取选中的那种类型
      let select_type = 0;
      switch (this.lists[0].mingxi) {
        case "two_end_in":
          select_type = 2;
          break;
        case "three_end_in":
          select_type = 3;
          break;
        case "four_end_in":
          select_type = 4;
          break;
        case "two_end_not_in":
          select_type = 2;
          break;
        case "three_end_not_in":
          select_type = 3;
          break;
        case "four_end_not_in":
          select_type = 4;
          break;
        case "two_Animal_in":
          select_type = 2;
          break;
        case "three_Animal_in":
          select_type = 3;
          break;
        case "four_Animal_in":
          select_type = 4;
          break;
        case "five_Animal_in":
          select_type = 5;
          break;
        case "two_Animal_not_in":
          select_type = 2;
          break;
        case "three_Animal_not_in":
          select_type = 3;
          break;
        case "four_Animal_not_in":
          select_type = 4;
          break;

        case "two_Animal":
          select_type = 2;
          break;
        case "three_Animal":
          select_type = 3;
          break;
        case "four_Animal":
          select_type = 4;
          break;
        case "five_Animal":
          select_type = 5;
          break;
        case "six_Animal":
          select_type = 6;
          break;
        case "seven_Animal":
          select_type = 7;
          break;
        case "eight_Animal":
          select_type = 8;
          break;
        case "nine_Animal":
          select_type = 9;
          break;
        case "ten_Animal":
          select_type = 10;
          break;
        case "elven_Animal":
          select_type = 11;
          break;

        case "five_not_in":
          select_type = 5;
          break;
        case "six_not_in":
          select_type = 6;
          break;
        case "seven_not_in":
          select_type = 7;
          break;
        case "eight_not_in":
          select_type = 8;
          break;
        case "nine_not_in":
          select_type = 9;
          break;
        case "ten_not_in":
          select_type = 10;
          break;
        case "elven_not_in":
          select_type = 11;
          break;
        case "twelve_not_in":
          select_type = 12;
          break;
      }
      //二维数组转换
      const newGroup = algorithm.groupSplit(this.choose_box(), select_type); //select_type是选中的是类型,比如2
      // console.log(newGroup);
      this.team = newGroup.length;
      this.money = this.money * this.team;
      let checkedArr = [];
      newGroup.forEach((item, index) => {
        checkedArr[index] = JSON.parse(JSON.stringify(money_data));
        checkedArr[index].input_name = item.join();
      });

      this.go_pour = checkedArr;
      fn(this.go_pour);
    },
    pour: function() {
      window.pour_status = 1;
      const body = {
        fc_type: JSON.stringify(this.$route.query.page),
        qishu: JSON.stringify(this.qishu),
        data: JSON.stringify(this.go_pour)
      };
      this.loading = true;
      api.addbet(this, body, res => {
        if (res.data.ErrorCode == 1) {
          window.pour_status = 2;
          this.loading = false;
          this.$Modal.success({
            content: "下注成功",
            onOk: () => {
              this.$root.$emit("success", true);
              this.$root.$emit('bet_success',true);
              window.clearTimeout(time);
            }
          });
          var time = window.setTimeout(() => {
            this.$root.$emit('success',true);
            this.$root.$emit('bet_success',true);
            this.$Modal.remove()
          }, share.bet_time)
        } else if (res.data.ErrorCode == 2) {
          this.loading = false;
        }
      });
    },
    reset: function() {
      // this.modal = false;
      this.lists = [];
      this.$root.$emit("reset", "");
    },
    cancel: function() {
      this.loading = false;
      this.$emit("cancel", false);
    },
    //数组去重
    unique: function(arr) {
      var newArr = [];
      for (var i in arr) {
        if (newArr.indexOf(arr[i]) == -1) {
          newArr.push(arr[i]);
        }
      }
      return newArr;
    }
  }
};
</script>
