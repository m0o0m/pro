<template>
  <div class="cash">
      <div class="bill">
        <span>
          <span class="titles">流水账单:</span>
          <I-Select v-model="cash_type" style="width:178px">
            <I-Option v-for="item in cityList" :value="item.value" :key="item.value">{{ item.label }}</I-Option>
          </I-Select>
        </span>
        <span>
          <span class="titles">交易时间:</span>
          <DatePicker :options="data_options" :value="day" type="date" placeholder="请输入日期" style="width: 190px" @on-change="timego"></DatePicker>
        </span>
        <span style="margin-left: 10px">
          <I-Button type="primary" style="width: 78px" @click="topsearch">搜索</I-Button>
        </span>
      </div>
      <div class="table">
        <ul class="thend">
          <li>交易日期</li>
          <li>流水项目</li>
          <li>流水类型</li>
          <li>现金金额</li>
          <li>现有金额</li>
          <li>备注</li>
        </ul>
        <ul v-for="item in list" class="forms clearfix">
          <li>{{item.adddate}}</li>
          <li>{{item.cash_do_typeTxt}}</li>
          <li>{{item.cash_typeTxt}}</li>
          <li>{{item.cash_num}}</li>
          <li>{{item.cash_balance}}</li>
          <li>
            <Poptip trigger="hover" :content="item.remark" placement="top-end">
              <p class="input_nameTxt">{{item.remark}}</p>
            </Poptip>
          </li>
        </ul>
        <ul class="ulfor" v-if="this.list.length==0">
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
import api from "../../api/config";
import {DatePicker,Poptip,Select,Option,Page,Button, Modal} from 'iview';
export default {
  components: {DatePicker,Poptip,'I-Select':Select,'I-Option':Option,Page,'I-Button':Button,Modal},
  props: {},
  data() {
    return {
      data_options:{
        disabledDate (date) {
          return date.valueOf() > Date.now();
        }
      },
      Pagecount: null,
      page: 1,
      pagenum: 10,
      day: null,
      cash_type: "全部",
      Recordcount: null,
      flag: true,
      numpage: 1,
      list: [{}, {}, {}, {}, {}, {}, {}],
      // lists: [
      //   {
      //     value: "全部",
      //     label: "全部"
      //   },
      //   {
      //     value: "已结算",
      //     label: "已结算"
      //   },
      //   {
      //     value: "未结算",
      //     label: "未结算"
      //   }
      // ],
      cityList: [
        {
          value: "全部",
          label: "全部"
        },
        {
          value: "1",
          label: "彩票下注"
        },
        {
          value: "2",
          label: "彩票派彩"
        },
        {
          value: "3",
          label: "彩票和局"
        },
        {
          value: "4",
          label: "额度转入"
        },
        {
          value: "5",
          label: "额度转出"
        },
        {
          value: "6",
          label: "注单取消"
        },
        {
            value: "8",
            label: "人工存入"
        },
        {
            value: "9",
            label: "人工取出"
        },
      ]
    };
  },
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

    this.day = new Date().Format("yyyy-MM-dd");
    this.getlist();
  },

  methods: {
    record() {
      this.$router.push({
        name: "record"
      });
    },
    gogo(e) {
      this.page = e;
      this.getlist();
    },
    myclick(e) {
      this.page = e;
      this.getlist();
    },
    change(e) {
      this.pagenum = e;
      this.getlist();
    },

    topsearch() {
      this.page = 1;
      this.numpage = 1;
      this.getlist();
    },

    search(e) {
      if (this.numpage > this.Pagecount) {
        this.numpage = this.Pagecount;
      }
      if (this.numpage != 0 && this.page != this.numpage) {
        this.page = Number(this.numpage);
        this.getlist();
      }
    },
    timego(e) {
      this.day = e;
    },
    getlist() {
      if (this.flag) {
        this.flag = false;

        var cash_type = this.cash_type;

        if (this.cash_type=='全部') {
          cash_type=''
        }

        let body = {
          cash_type: cash_type,
          day: this.day,
          // starttime: this.starttime,
          // endtime: this.endtime,
          page: this.page,
          pagenum: this.pagenum
        };

        this.$root.$emit("loading", true);
        api.cashs(this, body, res => {
          if (res.data.ErrorCode == 1) {
            this.$set(this.list, this.list);
            this.list = res.data.Data;
            this.Pagecount = res.data.Pagecount;
            this.Recordcount = Number(res.data.Recordcount);
            this.$root.$emit("loading", false);
            this.flag = true;
          }
        });
      }
    }
  }
};
</script>
<style lang="scss" src="../../assets/css/cash.scss" scoped></style>
