<template>
  <div class="report">
    <div class="r_box">
      <div class="table">
        <ul class="thend clearfix">
          <li>日期</li>
          <li>彩票类型</li>
          <li>注单笔数</li>
          <li>下注金额</li>
          <li>有效下注总额</li>
          <li>盈利总额</li>
        </ul>
        <ul v-for="(key,index) in lists" class="from">
          <li  @click="myclick(key,index)" class="li1">
            <ul class="ul1 clearfix">
              <li>{{key.addday}}</li>
              <li>
                <div class='rightbtn'>
                  <i :class="[key.flag?'pk-arrowDown':'pk-jiantou']" class='iconfont'></i>
                </div>
              </li>
              <li>{{key.bet_count}}</li>
              <li>{{key.bet}}</li>
              <li>{{key.valid_bet}}</li>
              <li :class="{'red':key.win!=0}">{{key.win}}</li>
            </ul>
          </li>
          <li class="li2" ref="myref">
            <ul @click="tiaozhuan(item)" v-for="item in key.list" class="ul2 clearfix">
              <li></li>
              <li>{{item.fc_typeTxt}}</li>
              <li>{{item.bet_count}}</li>
              <li>{{item.bet}}</li>
              <li>{{item.valid_bet}}</li>
              <li class="red" v-if="item.win!=0">{{item.win}}</li>
              <li v-else>未中奖</li>
            </ul>
          </li>
        </ul>
        <ul class="ul3 clearfix" v-if="this.lists.length">
          <li style="color: red;">总计</li>
          <li></li>
          <li>{{totals.bet_count}}</li>
          <li>{{totals.bet}}</li>
          <li>{{totals.valid_bet}}</li>
          <li class="red" v-if="totals.win!=0">{{totals.win}}</li>
          <li v-else>未中奖</li>
        </ul>
        <ul class="ulfor" v-else="this.lists.length==0">
          <li>暂无数据</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
  import api from "../../api/config";
  //import back from "./back";
  import animate from "../../assets/js/animate";

  export default {
    components: {},
    props: {},
    data() {
      return {
        show: null,
        index: null,
        lists: [],
        totals: [{}],
        ulls: null,
        this: null
      };
    },
    created() {
      //
      // var d = new Date().getDay();
      // console.log(d);
      this.getlist();
    },
    mounted() {
    },
    methods: {
      tiaozhuan(item) {
//        console.log(item);
        //        this.$root.$emit('send_data',item.fc_type);
        this.$router.push({
          name: "record",
          params: { one: item.fc_type, two: "已结算", date: item.addday }
        });
      },
      myclick(key, i) {
        key.flag = !key.flag;
        // console.log(key);

        if (this.ulls == null) {
          this.ulls = this.$refs.myref;
        }
        this.myAccordion(this.ulls, key, i, this.lists);
      },

      myAccordion(ul, item, i, lists) {
        for (let index = 0; index < lists.length; index++) {
          if (index == i) {
            if (lists[index].flag) {
              ul[index].style.display = "block";
              animate(ul[index], {
                height: 60 * ul[index].children.length,
                opacity: 1
              });
            } else {
              animate(
                ul[index],
                {
                  height: 0,
                  opacity: 0
                },
                () => {
                  ul[index].style.display = "none";
                }
              );
            }
          } else {
            lists[index].flag = false;
            animate(
              ul[index],
              {
                height: 0,
                opacity: 0
              },
              () => {
                ul[index].style.display = "none";
              }
            );
          }
        }
      },

      report_one() {
        this.$router.push({
          name: "report_one"
        });
      },
      getlist() {
        this.$root.$emit("loading", true);
        let body = {
          week: "this"
        };
        api.report(this, body, res => {
          if (res.data.ErrorCode == 1) {
            this.$root.$emit("loading", false);
            // this.$set(this.lists, this.lists);
            this.lists = res.data.Data;
            console.log(res.data.Data);
//            for (let i = 0; i < this.lists.length; i++) {
//              Object.assign(this.lists[i], res.data.Data[i]);
//            }
            this.totals = res.data.Total;
            this.lists.length = res.data.Data.length;
            for (let i = 0; i < this.lists.length; i++) {
              this.lists[i].flag = false;
              console.log(this.lists[i]);
            }
          }
        });
      }
    }
  };
</script>
<style lang="scss" scoped src="../../assets/css/report.scss"></style>
