<template>
  <div class="record">
    <div class="dates">
      <span>
        <span class="titles">注单号:</span>
        <I-Input style="width: 125px" v-model="order_num"></I-Input>
      </span>
      <span>
      <span class="titles">投注时间:</span>
      <DatePicker :options="data_options" :value="day" type="date" placeholder="请输入日期" style="width: 190px" @on-change="timego"></DatePicker>
      </span>
      <span>
        <span class="titles">彩票种类:</span>
      <I-Select v-model="fc_type" style="width:178px">
        <I-Option v-for="item in cityList" :value="item.value" :key="item.value">{{ item.label }}</I-Option>
      </I-Select>
      </span>
      <I-Select v-model="model1" style="width:80px;margin-left: 10px;">
        <I-Option v-for="item in lists" :value="item.value" :key="item.value">{{ item.label }}</I-Option>
      </I-Select>
      <span style="margin-left: 10px">
        <I-Button type="primary" style="width: 78px" @click="topsearch">搜索</I-Button>
      </span>
    </div>
    <div class="table">
      <div class='top'>
        <ul class="clearfix thend">
          <li class="thend_date fl" v-for="item in list">{{item}}</li>
        </ul>
      </div>
      <ul v-for="(item,i) in addlist" class="clearfix forms">
        <li class="thend_date fl">{{item.adddate}}</li>
        <li class="thend_date fl">{{item.order_num}}</li>
        <li class="thend_date fl">{{item.periods}}</li>
        <li class="thend_date fl">{{item.fc_typeTxt}}</li>
        <li class="thend_date fl">{{item.gameplayTxt}}</li>
        <li class="thend_date fl">
          <Poptip trigger="hover"  :content="item.input_nameTxt">
            <p class="input_nameTxt">{{item.input_nameTxt}}</p>
          </Poptip>
        </li>
        <li class="thend_date fl">{{item.valid_bet}}</li>
        <li class="thend_date fl" >{{item.can_win}}</li>
        <li class="thend_date fl" :class="{'red':item.win=='未结算'}">{{item.win}}</li>
      </ul>
      <ul class="ulfor" v-if="addlist.length==0">
        <li>暂无数据</li>
      </ul>
    </div>
    <div class="footer clearfix">
      <div class="top_txt">共搜索到{{Recordcount}}条数据，共{{Pagecount}}页</div>
      <div class="clearfix">
        <Page class="fl" :current="page" :total="Recordcount" show-sizer placement="top" @on-change="gogo" @on-page-size-change="change"></Page>
        <span class="txt fl">跳至</span>
        <input type="text" class="inputTxt fl" v-model="numpage" onkeyup="this.value=this.value.replace(/\D/g,'')" onafterpaste="this.value=this.value.replace(/\D/g,'')">
        <span class="txt fl">页</span>
        <I-Button style="height: 30px ; width: 30px" class="fl" type="primary" shape="circle" icon="ios-search" @click="search"></I-Button>
      </div>
    </div>
  </div>
</template>

<script>
import httpApi from "../../api/config";
import {Input,DatePicker,Poptip,Select,Option,Page,Button, Modal} from 'iview';
export default {
  components: {'I-Input':Input,DatePicker,Poptip,'I-Select':Select,'I-Option':Option,Page,'I-Button':Button,Modal},
  props: {},
  data() {
    return {
      data_options:{
        disabledDate (date) {
          return date.valueOf() > Date.now();
        }
      },
      lists: [
        {
          value: "全部",
          label: "全部"
        },
        {
          value: "已结算",
          label: "已结算"
        },
        {
          value: "未结算",
          label: "未结算"
        }
      ],
      model1: '全部',
      list: ["投注日期", "注单号", "期数", "投注类型", "玩法", "投注内容", "投注金额", "可赢金额", "派彩金额"],
      cityList: [],
      addlist: [{}, {}, {}, {}, {}, {}, {}, {}],
      //仓库里有多少条数据
      Recordcount: null,
      //数据库中的总页数
      Pagecount: null,
      //第几页
      page: 1,
      //每页显示数量
      pagenum: 10,
      //注单号
      order_num: null,
      //投资彩种
      fc_type: "全部",
      //日期
      day: null,
      // endtime: null,
      flag: true,
      numpage: 1
    };
  },
  // watch: {
  //   // 如果路由有变化，会再次执行该方法
  //   $route: "getlist" // 只有这个页面初始化之后，这个监听事件才开始生效
  // },
  created() {
    Date.prototype.Format = function(fmt) {
      var o = {
        "M+": this.getMonth() + 1, //月份
        "d+": this.getDate(), //日
        "h+": this.getHours(), //小时
        "m+": this.getMinutes(), //分
        "s+": this.getSeconds(), //秒
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度
        S: this.getMilliseconds() //毫秒
      };
      if (/(y+)/.test(fmt))
        fmt = fmt.replace(
          RegExp.$1,
          (this.getFullYear() + "").substr(4 - RegExp.$1.length)
        );
      for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt))
          fmt = fmt.replace(
            RegExp.$1,
            RegExp.$1.length == 1
              ? o[k]
              : ("00" + o[k]).substr(("" + o[k]).length)
          );
      return fmt;
    };
    if (this.$route.params.date) {
      this.fc_type = this.$route.params.one;
      this.model1 = this.$route.params.two;
      this.day = this.$route.params.date;
      if (this.fc_type && this.day) {
        this.flag = true;
        this.getlist();
      }
    }else{
      this.day = new Date().Format("yyyy-MM-dd");
      if(this.$route.params.page){
        this.fc_type = this.$route.params.page;
      }
      this.getlist();
    }
  },
   mounted() {
//     if (this.$route.params.page) {
//       this.fc_type = this.$route.params.page;
//       this.getlist();
//     }
   },
  methods: {
    cash() {
      this.$router.push({
        name: "cash"
      });
    },
    gogo(e) {
      this.page = e;
      this.getlist();
    },
    change(e) {
      this.pagenum = e;
      this.getlist();
    },
    timego(e) {
      this.day = e;
    },
    topsearch() {
      this.page = 1;
      this.numpage = 1;
      this.getlist();
    },
    search() {
      if (this.numpage > this.Pagecount) {
        this.numpage = this.Pagecount;
      }
      if (this.numpage != 0 && this.page != this.numpage) {
        this.page = Number(this.numpage);
        this.getlist();
      }
    },
    getlist() {
      if (this.flag) {
        this.flag = false;

        var type=this.fc_type;

        if (type=='全部') {
          type=''
        }
        let wind_data = '';
        if(this.model1=='已结算'){
          wind_data = 2
        }else if(this.model1 == '未结算'){
          wind_data = 1
        }else{
          wind_data = ''
        }
        let body = {
          fc_type: type,
          order_num: this.order_num,
          day: this.day,
          // endtime: this.endtime,
          wind: wind_data,
          page: this.page,
          pagenum: this.pagenum
        };
        this.$root.$emit("loading", true);
        httpApi.bets(this, body, res => {
          if (res.data.ErrorCode == 1) {
            this.$set(this.addlist, this.addlist);
            this.addlist = res.data.Data;
            this.Pagecount = res.data.Pagecount;
            this.cityList = res.data.type_list;
            this.cityList.unshift({
              value: "全部",
              label: "全部"
            });
            this.Recordcount = Number(res.data.Recordcount);
            for (let i = 0; i < this.addlist.length; i++) {
              // var l = null;
              // var a = this.addlist[i].odds;
              // var b = this.addlist[i].bet;
              // l = a * b;
              this.addlist[i].bets = this.addlist[i].odds * this.addlist[i].bet;
            }
            // console.log(Recordcount);
            this.$root.$emit("loading", false);
            this.flag = true;
          }
        });
      }
    }
  }
};
</script>

<style lang="scss" src="../../assets/css/record.scss" scoped>

</style>
