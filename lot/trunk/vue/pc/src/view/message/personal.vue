<template>
  <div class="notice">
    <div class="r_box">
      <div class="table">
        <ul class="thend clearfix">
          <li>日期</li>
          <li>标题</li>
        </ul>
        <ul v-if="lists.length!=0" v-for="(key,index) in lists" class="from">
          <li @click="myclick(key,index)" class="li1">
            <ul class="ul1 clearfix">
              <li>{{key.adddate}}</li>
              <li>{{key.title}}</li>
              <li style="width: 30px">
                <i :class="[key.flag?'pk-arrowDown':'pk-jiantou']" class='iconfont'></i>
              </li>
              <!--<li class="li11" style="cursor: pointer" @click.prevent="del(key.id)">删除</li>-->
            </ul>
          </li>
          <!--内容-->
          <li class="li2 clearfix" ref="myref">
            <ul class="ul2 clearfix">
              <li>{{key.content}}</li>
            </ul>
          </li>
        </ul>
        <ul class="ulfor" v-if="lists.length == 0">
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
  </div>
</template>

<script>
  import api from "../../api/config";
  //import back from "./back";
  import {Page,Button, Modal} from 'iview';
  import animate from "../../assets/js/animate";

  export default {
    components: { Page,'I-Button':Button,Modal},
    props: {},
    data() {
      return {
        value:null,
        show: null,
        index: null,
        lists: [
          {adddate:'',title:'',content:''}
        ],
        //仓库里有多少条数据
        Recordcount: null,
        //数据库中的总页数
        Pagecount: null,
        totals: [{}],
        page: 1,
        //每页显示数量
        pagenum: 10,
        ulls: null,
        numpage: 1
      };
    },
    created() {
    },
    mounted() {
      this.getlist();
    },
    methods: {
      myclick(key, i) {
        key.flag = !key.flag;
//        console.log(key);
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
      search() {
        if (this.numpage > this.Pagecount) {
          this.numpage = this.Pagecount;
        }
        if (this.numpage != 0 && this.page != this.numpage) {
          this.page = Number(this.numpage);
          this.getlist();
        }
      },
      gogo(e) {
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
      getlist() {
        this.$root.$emit("loading", true);
        let body = {
          page: 1,
          pagenum: 10,
        };
        api.message(this, body, res => {
          if (res.data.ErrorCode == 1) {
            this.$root.$emit("loading", false);
            this.$set(this.lists, this.lists);
            this.lists.length = res.data.Data.length;
            this.Pagecount = res.data.Pagecount;
            this.Recordcount = Number(res.data.Recordcount);
//            for (var i = 0; i < this.lists.length; i++) {
//              Object.assign(this.lists[i], res.data.Data[i]);
//            }
            this.lists = res.data.Data;

          }
        });
      }
    }
  };
</script>
<style lang="scss" scoped src="../../assets/css/notice.scss"></style>
