<template>
  <div class="rightbox clearfix">
     <ul class="clearfix uls" v-for="item in list">
         <li v-for="key in item" :class="[key.flag?'table-current':'']" @click="myclick(key)">{{key.name}}</li>
     </ul>
  </div>
</template>
<script>
export default {
  data() {
    return {
      flag: false,
      list: [
        [
          {
            name: "大",
            flag: false
          },
          {
            name: "小",
            flag: false
          },
          {
            name: "单",
            flag: false
          },
          {
            name: "双",
            flag: false
          }
        ],
        [
          {
            name: "中",
            flag: false
          },
          {
            name: "边",
            flag: false
          },
          {
            name: "大单",
            flag: false
          },
          {
            name: "小单",
            flag: false
          }
        ],
        [
          {
            name: "大双",
            flag: false
          },
          {
            name: "小双",
            flag: false
          },
          {
            name: "大边",
            flag: false
          },
          {
            name: "小边",
            flag: false
          }
        ],
        [
          {
            name: "0尾",
            flag: false
          },
          {
            name: "1尾",
            flag: false
          },
          {
            name: "2尾",
            flag: false
          },
          {
            name: "3尾",
            flag: false
          }
        ],
        [
          {
            name: "4尾",
            flag: false
          },
          {
            name: "5尾",
            flag: false
          },
          {
            name: "6尾",
            flag: false
          },
          {
            name: "7尾",
            flag: false
          }
        ],
        [
          {
            name: "8尾",
            flag: false
          },
          {
            name: "9尾",
            flag: false
          },
          {
            name: "全包",
            flag: false
          }
        ],
        [
          {
            name: "3余0",
            flag: false
          },
          {
            name: "3余1",
            flag: false
          },
          {
            name: "3余2",
            flag: false
          },
          {
            name: "4余0",
            flag: false
          }
        ],
        [
          {
            name: "4余1",
            flag: false
          },
          {
            name: "4余2",
            flag: false
          },
          {
            name: "4余3",
            flag: false
          },
          {
            name: "5余0",
            flag: false
          }
        ],
        [
          {
            name: "5余1",
            flag: false
          },
          {
            name: "5余2",
            flag: false
          },
          {
            name: "5余3",
            flag: false
          },
          {
            name: "5余4",
            flag: false
          }
        ]
      ],
      select: null
    };
  },
  props: ["lists", "money"],
  created() {
    this.$root.$on("success", e => {
      if (e) {
        for (let i = 0; i < this.list.length; i++) {
          for (var l = 0; l < this.list[i].length; l++) {
            this.list[i][l].flag = false;
          }
        }
      }
    });
    this.$root.$on("right_config", e => {
      if (e) {
        for (let i = 0; i < this.list.length; i++) {
          for (var l = 0; l < this.list[i].length; l++) {
            this.list[i][l].flag = false;
          }
        }
      }
    });
  },
  methods: {
    myclick(key) {
      if (
        this.lists[0].object[this.lists[0].object.length - 1].name == "特码包三"
      ) {
        var ta_flag = this.lists[0].object[this.lists[0].object.length - 1]
          .flag;
        var ta_money = this.lists[0].object[this.lists[0].object.length - 1]
          .money;
      }

      for (let i = 0; i < this.lists.length; i++) {
        for (let l = 0; l < this.lists[i].object.length; l++) {
          this.lists[i].object[l].flag = false;
          this.lists[i].object[l].money = "";
        }
      }

      if (key == this.select) {
        key.flag = false;
        this.select = null;

        if (
          this.lists[0].object[this.lists[0].object.length - 1].name == "特码包三"
        ) {
          this.lists[0].object[this.lists[0].object.length - 1].flag = ta_flag;
          this.lists[0].object[
            this.lists[0].object.length - 1
          ].money = ta_money;
        }
      } else {
        this.select = key;

        for (let i = 0; i < this.list.length; i++) {
          for (var l = 0; l < this.list[i].length; l++) {
            this.list[i][l].flag = false;
          }
        }

        key.flag = true;

        if (key.name == "全包") {
          for (let i = 0; i < this.lists.length; i++) {
            for (let l = 0; l < this.lists[i].object.length; l++) {
              this.lists[i].object[l].flag = true;
              this.lists[i].object[l].money = this.money;
            }
          }
        } else {
          for (let i = 0; i < this.lists.length; i++) {
            for (let l = 0; l < 7; l++) {
              this.lists[i].object[l].flag = false;
              for (var k = 0; k < this.lists[i].object[l].class.length; k++) {
                if (this.lists[i].object[l].class[k] == key.name) {
                  this.lists[i].object[l].flag = true;
                  this.lists[i].object[l].money = this.money;
                  break;
                }
              }
            }
          }
        }

        if (
          this.lists[0].object[this.lists[0].object.length - 1].name == "特码包三"
        ) {
          this.lists[0].object[this.lists[0].object.length - 1].flag = ta_flag;
          this.lists[0].object[
            this.lists[0].object.length - 1
          ].money = ta_money;
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped>
@import "../../../assets/css/function.scss";
.rightbox {
  width: 170px;
  float: right;
  margin-right: 14px;
  .uls {
    height: 36px;
    border: 1px solid $border_color;
    border-radius: 18px;
    overflow: hidden;
    margin-bottom: 5px;
    li {
      float: left;
      line-height: 34px;
      width: 25%;
      cursor: pointer;
    }
    li:hover {
      background-color: $bg_select;
      color:red;
    }
  }
  .table-current {
    background-color: $bg_select;
    color:red;
  }
}
</style>
