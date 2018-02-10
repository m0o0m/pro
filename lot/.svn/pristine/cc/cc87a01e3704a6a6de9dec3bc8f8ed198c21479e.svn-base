<template>
  <div class="often">
    <div class="pk_time">
      <div class="time_left">
        <div class="left_logo">
          <img :src="c_data.img_path" alt="">
        </div>
        <div class="left_text">
          <p>{{c_data.fc_name}}</p>
          <p><span>第</span> {{c_data.qishu}} <span>期</span></p>
        </div>
      </div>
      <div class="time_center">
        <p class="center_top">投注剩余时间</p>
        <div class="center_bottom">
          <div class="fl time_content" v-if="h < 10 && h >= 0">0{{h}}</div>
          <div class="fl time_content" v-else-if="h < 0">--</div>
          <div class="fl time_content" v-else>{{h}}</div>
          <div class="fl fs">时</div>
          <div class="fl time_content" v-if="m < 10 && m >= 0">0{{m}}</div>
          <div class="fl time_content" v-else-if="m < 0">--</div>
          <div class="fl time_content" v-else>{{m}}</div>
          <div class="fl fs">分</div>
          <div class="fl time_content" v-if="s < 10 && s >= 0">0{{s}}</div>
          <div class="fl time_content" v-else-if="s < 0">--</div>
          <div class="fl time_content" v-else>{{s}}</div>
          <div class="fl fs">秒</div>
        </div>
      </div>
      <div class="time_bottom">
        <div class="bottom_left"><span>第</span> {{auto.qishu}} <span>期开奖</span></div>
        <div v-if="this.$route.query.page == 'liuhecai'" class='bottom_center6'>
          <div class="bottom_content">
            <div class="content_list" v-for="item in auto.ball">
              <p :class="[item.color,'box']">{{item.number}}</p>
              <p class="animal">{{item.animal}}</p>
            </div>
          </div>
        </div>
        <div v-else :class="[auto.ball.length >= 10?'bottom_center_other':'bottom_center']">
          <div :class="[auto.ball.length == 10?'other_content':'bottom_content']">
            <div class="content_list" v-for="item in auto.ball">
              <p class="box blue">{{item}}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="bottom_bottom">
        <p @click="history()">历史结果</p>
        <p @click="open_way()">开奖走势</p>
        <p @click="show_rule()">玩法规则</p>
      </div>
    </div>
    <Nav-top :lists="menus" @menu="go_child"></Nav-top>
    <router-view :Ilists="integrateLists" :Olists="orLists" :cdata="c_data"></router-view>
    <Nav-bottom ref="dewdrop_map" :back_data="auto_list" :nav_top="bottom_nav"></Nav-bottom>
  </div>
</template>

<script>
function sscInitial() {
  return [
    {
      name: "第一球",
      object: [
        {
          money: "",
          li_id: 0,
          flag: false
        },
        {
          money: "",
          li_id: 1,
          flag: false
        },
        {
          money: "",
          li_id: 2,
          flag: false
        },
        {
          money: "",
          li_id: 3,
          flag: false
        },
        {
          money: "",
          li_id: 4,
          flag: false
        },
        {
          money: "",
          li_id: 5,
          flag: false
        },
        {
          money: "",
          li_id: 6,
          flag: false
        },
        {
          money: "",
          li_id: 7,
          flag: false
        },
        {
          money: "",
          li_id: 8,
          flag: false
        },
        {
          money: "",
          li_id: 9,
          flag: false
        },
        {
          money: "",
          li_id: 10,
          flag: false
        },
        {
          money: "",
          li_id: 11,
          flag: false
        },
        {
          money: "",
          li_id: 12,
          flag: false
        },
        {
          money: "",
          li_id: 13,
          flag: false
        }
      ]
    },
    {
      name: "第二球",
      object: [
        {
          money: "",
          li_id: 14,
          flag: false
        },
        {
          money: "",
          li_id: 15,
          flag: false
        },
        {
          money: "",
          li_id: 16,
          flag: false
        },
        {
          money: "",
          li_id: 17,
          flag: false
        },
        {
          money: "",
          li_id: 18,
          flag: false
        },
        {
          money: "",
          li_id: 19,
          flag: false
        },
        {
          money: "",
          li_id: 20,
          flag: false
        },
        {
          money: "",
          li_id: 21,
          flag: false
        },
        {
          money: "",
          li_id: 22,
          flag: false
        },
        {
          money: "",
          li_id: 23,
          flag: false
        },
        {
          money: "",
          li_id: 24,
          flag: false
        },
        {
          money: "",
          li_id: 25,
          flag: false
        },
        {
          money: "",
          li_id: 26,
          flag: false
        },
        {
          money: "",
          li_id: 27,
          flag: false
        }
      ]
    },
    {
      name: "第三球",
      object: [
        {
          money: "",
          li_id: 28,
          flag: false
        },
        {
          money: "",
          li_id: 29,
          flag: false
        },
        {
          money: "",
          li_id: 30,
          flag: false
        },
        {
          money: "",
          li_id: 31,
          flag: false
        },
        {
          money: "",
          li_id: 32,
          flag: false
        },
        {
          money: "",
          li_id: 33,
          flag: false
        },
        {
          money: "",
          li_id: 34,
          flag: false
        },
        {
          money: "",
          li_id: 35,
          flag: false
        },
        {
          money: "",
          li_id: 36,
          flag: false
        },
        {
          money: "",
          li_id: 37,
          flag: false
        },
        {
          money: "",
          li_id: 38,
          flag: false
        },
        {
          money: "",
          li_id: 39,
          flag: false
        },
        {
          money: "",
          li_id: 40,
          flag: false
        },
        {
          money: "",
          li_id: 41,
          flag: false
        }
      ]
    },
    {
      name: "第四球",
      object: [
        {
          money: "",
          li_id: 42,
          flag: false
        },
        {
          money: "",
          li_id: 43,
          flag: false
        },
        {
          money: "",
          li_id: 44,
          flag: false
        },
        {
          money: "",
          li_id: 45,
          flag: false
        },
        {
          money: "",
          li_id: 46,
          flag: false
        },
        {
          money: "",
          li_id: 47,
          flag: false
        },
        {
          money: "",
          li_id: 48,
          flag: false
        },
        {
          money: "",
          li_id: 49,
          flag: false
        },
        {
          money: "",
          li_id: 50,
          flag: false
        },
        {
          money: "",
          li_id: 51,
          flag: false
        },
        {
          money: "",
          li_id: 52,
          flag: false
        },
        {
          money: "",
          li_id: 53,
          flag: false
        },
        {
          money: "",
          li_id: 54,
          flag: false
        },
        {
          money: "",
          li_id: 55,
          flag: false
        }
      ]
    },
    {
      name: "第五球",
      object: [
        {
          money: "",
          li_id: 56,
          flag: false
        },
        {
          money: "",
          li_id: 57,
          flag: false
        },
        {
          money: "",
          li_id: 58,
          flag: false
        },
        {
          money: "",
          li_id: 59,
          flag: false
        },
        {
          money: "",
          li_id: 60,
          flag: false
        },
        {
          money: "",
          li_id: 61,
          flag: false
        },
        {
          money: "",
          li_id: 62,
          flag: false
        },
        {
          money: "",
          li_id: 63,
          flag: false
        },
        {
          money: "",
          li_id: 64,
          flag: false
        },
        {
          money: "",
          li_id: 65,
          flag: false
        },
        {
          money: "",
          li_id: 66,
          flag: false
        },
        {
          money: "",
          li_id: 67,
          flag: false
        },
        {
          money: "",
          li_id: 68,
          flag: false
        },
        {
          money: "",
          li_id: 69,
          flag: false
        }
      ]
    },
    {
      name: "前三球",
      object: [
        {
          money: "",
          li_id: 70,
          flag: false
        },
        {
          money: "",
          li_id: 71,
          flag: false
        },
        {
          money: "",
          li_id: 72,
          flag: false
        },
        {
          money: "",
          li_id: 73,
          flag: false
        },
        {
          money: "",
          li_id: 74,
          flag: false
        }
      ]
    },
    {
      name: "中三球",
      object: [
        {
          money: "",
          li_id: 75,
          flag: false
        },
        {
          money: "",
          li_id: 76,
          flag: false
        },
        {
          money: "",
          li_id: 77,
          flag: false
        },
        {
          money: "",
          li_id: 78,
          flag: false
        },
        {
          money: "",
          li_id: 79,
          flag: false
        }
      ]
    },
    {
      name: "后三球",
      object: [
        {
          money: "",
          li_id: 80,
          flag: false
        },
        {
          money: "",
          li_id: 81,
          flag: false
        },
        {
          money: "",
          li_id: 82,
          flag: false
        },
        {
          money: "",
          li_id: 83,
          flag: false
        },
        {
          money: "",
          li_id: 84,
          flag: false
        }
      ]
    },
    {
      name: "总和龙虎",
      object: [
        {
          money: "",
          li_id: 85,
          flag: false
        },
        {
          money: "",
          li_id: 86,
          flag: false
        },
        {
          money: "",
          li_id: 87,
          flag: false
        },
        {
          money: "",
          li_id: 88,
          flag: false
        }
      ]
    },
    {
      name: "总和龙虎",
      object: [
        {
          money: "",
          li_id: 89,
          flag: false
        },
        {
          money: "",
          li_id: 90,
          flag: false
        },
        {
          money: "",
          li_id: 91,
          flag: false
        }
      ]
    },
    {
      name: "斗牛",
      object: [
        {
          money: "",
          li_id: 92,
          flag: false
        },
        {
          money: "",
          li_id: 93,
          flag: false
        },
        {
          money: "",
          li_id: 94,
          flag: false
        },
        {
          money: "",
          li_id: 95,
          flag: false
        },
        {
          money: "",
          li_id: 96,
          flag: false
        }
      ]
    },
    {
      name: "斗牛",
      object: [
        {
          money: "",
          li_id: 97,
          flag: false
        },
        {
          money: "",
          li_id: 98,
          flag: false
        },
        {
          money: "",
          li_id: 99,
          flag: false
        },
        {
          money: "",
          li_id: 100,
          flag: false
        },
        {
          money: "",
          li_id: 101,
          flag: false
        }
      ]
    },
    {
      name: "斗牛",
      object: [
        {
          money: "",
          li_id: 102,
          flag: false
        },
        {
          money: "",
          li_id: 103,
          flag: false
        },
        {
          money: "",
          li_id: 104,
          flag: false
        },
        {
          money: "",
          li_id: 105,
          flag: false
        },
        {
          money: "",
          li_id: 106,
          flag: false
        }
      ]
    },
    {
      name: "梭哈",
      object: [
        {
          money: "",
          li_id: 107,
          flag: false
        },
        {
          money: "",
          li_id: 108,
          flag: false
        },
        {
          money: "",
          li_id: 109,
          flag: false
        },
        {
          money: "",
          li_id: 110,
          flag: false
        },
        {
          money: "",
          li_id: 111,
          flag: false
        }
      ]
    },
    {
      name: "梭哈",
      object: [
        {
          money: "",
          li_id: 112,
          flag: false
        },
        {
          money: "",
          li_id: 113,
          flag: false
        },
        {
          money: "",
          li_id: 114,
          flag: false
        }
      ]
    }
  ];
}
function SscOrInitial() {
  return [
    {
      name: "第一球",
      object: [
        {
          or_id: 0,
          money: "",
          flag: false
        },
        {
          or_id: 1,
          money: "",
          flag: false
        },
        {
          or_id: 2,
          money: "",
          flag: false
        },
        {
          or_id: 3,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "第二球",
      object: [
        {
          or_id: 4,
          money: "",
          flag: false
        },
        {
          or_id: 5,
          money: "",
          flag: false
        },
        {
          or_id: 6,
          money: "",
          flag: false
        },
        {
          or_id: 7,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "第三球",
      object: [
        {
          or_id: 8,
          money: "",
          flag: false
        },
        {
          or_id: 9,
          money: "",
          flag: false
        },
        {
          or_id: 10,
          money: "",
          flag: false
        },
        {
          or_id: 11,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "第四球",
      object: [
        {
          or_id: 12,
          money: "",
          flag: false
        },
        {
          or_id: 13,
          money: "",
          flag: false
        },
        {
          or_id: 14,
          money: "",
          flag: false
        },
        {
          or_id: 15,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "第五球",
      object: [
        {
          or_id: 16,
          money: "",
          flag: false
        },
        {
          or_id: 17,
          money: "",
          flag: false
        },
        {
          or_id: 18,
          money: "",
          flag: false
        },
        {
          or_id: 19,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "总和，龙虎",
      object: [
        {
          or_id: 20,
          money: "",
          flag: false
        },
        {
          or_id: 21,
          money: "",
          flag: false
        },
        {
          or_id: 22,
          money: "",
          flag: false
        },
        {
          or_id: 23,
          money: "",
          flag: false
        }
      ]
    },
    {
      name: "总和，龙虎",
      object: [
        {
          or_id: 24,
          money: "",
          flag: false
        },
        {
          or_id: 25,
          money: "",
          flag: false
        },
        {
          or_id: 26,
          money: "",
          flag: false
        }
      ]
    }
  ];
}
import api from "../../../api/config";
import NavTop from "../../../share_components/default_nav";
import ws from '../../../assets/js/socket'
import cm_cookie from '../../../assets/js/com_cookie'
import NavBottom from '../../../share_components/dewdrop_map'
export default {
  components: {
    NavTop,NavBottom
  },
  data() {
    return {
      auto_list:[],
      bottom_nav:[
          {name:'万位'},
          {name:'千位'},
          {name:'百位'},
          {name:'十位'},
          {name:'个位'},
      ],
      margin: false,
      menus: [
        {
          name: "整合",
          item: 'cq_ssc'
        },
        {
          name: "两面盘",
          item: 'ssc_or'
        }
      ],
      routePage: null,
      integrateLists: [
        {
          name: "第一球",
          object: [
            {
              money: "",
              li_id: 0,
              flag: false
              //
            },
            {
              money: "",
              li_id: 1,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 2,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 3,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 4,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 5,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 6,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 7,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 8,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 9,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 10,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 11,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 12,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 13,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "第二球",
          object: [
            {
              money: "",
              li_id: 14,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 15,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 16,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 17,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 18,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 19,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 20,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 21,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 22,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 23,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 24,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 25,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 26,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 27,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "第三球",
          object: [
            {
              money: "",
              li_id: 28,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 29,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 30,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 31,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 32,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 33,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 34,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 35,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 36,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 37,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 38,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 39,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 40,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 41,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "第四球",
          object: [
            {
              money: "",
              li_id: 42,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 43,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 44,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 45,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 46,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 47,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 48,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 49,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 50,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 51,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 52,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 53,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 54,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 55,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "第五球",
          object: [
            {
              money: "",
              li_id: 56,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 57,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 58,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 59,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 60,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 61,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 62,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 63,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 64,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 65,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 66,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 67,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 68,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 69,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "前三球",
          object: [
            {
              money: "",
              li_id: 70,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 71,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 72,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 73,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 74,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "中三球",
          object: [
            {
              money: "",
              li_id: 75,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 76,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 77,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 78,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 79,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "后三球",
          object: [
            {
              money: "",
              li_id: 80,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 81,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 82,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 83,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 84,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "总和龙虎",
          object: [
            {
              money: "",
              li_id: 85,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 86,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 87,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 88,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "总和龙虎",
          object: [
            {
              money: "",
              li_id: 89,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 90,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 91,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "斗牛",
          object: [
            {
              money: "",
              li_id: 92,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 93,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 94,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 95,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 96,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "斗牛",
          object: [
            {
              money: "",
              li_id: 97,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 98,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 99,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 100,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 101,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "斗牛",
          object: [
            {
              money: "",
              li_id: 102,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 103,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 104,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 105,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 106,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "梭哈",
          object: [
            {
              money: "",
              li_id: 107,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 108,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 109,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 110,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 111,
              flag: false
              // name: "00"
            }
          ]
        },
        {
          name: "梭哈",
          object: [
            {
              money: "",
              li_id: 112,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 113,
              flag: false
              // name: "00"
            },
            {
              money: "",
              li_id: 114,
              flag: false
              // name: "00"
            }
          ]
        }
      ],
      orLists: [
        {
          name: "第一球",
          object: [
            {
              //
              or_id: 0,
              money: "",
              flag: false
            },
            {
              //
              or_id: 1,
              money: "",
              flag: false
              // name: "00"
            },
            {
              //
              or_id: 2,
              money: "",
              flag: false
            },
            {
              //
              or_id: 3,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "第二球",
          object: [
            {
              //
              or_id: 4,
              money: "",
              flag: false
            },
            {
              //
              or_id: 5,
              money: "",
              flag: false
            },
            {
              //
              or_id: 6,
              money: "",
              flag: false
            },
            {
              //
              or_id: 7,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "第三球",
          object: [
            {
              //
              or_id: 8,
              money: "",
              flag: false
            },
            {
              //
              or_id: 9,
              money: "",
              flag: false
            },
            {
              //
              or_id: 10,
              money: "",
              flag: false
            },
            {
              //
              or_id: 11,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "第四球",
          object: [
            {
              //
              or_id: 12,
              money: "",
              flag: false
            },
            {
              //
              or_id: 13,
              money: "",
              flag: false
            },
            {
              //
              or_id: 14,
              money: "",
              flag: false
            },
            {
              //
              or_id: 15,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "第五球",
          object: [
            {
              //
              or_id: 16,
              money: "",
              flag: false
            },
            {
              //
              or_id: 17,
              money: "",
              flag: false
            },
            {
              //
              or_id: 18,
              money: "",
              flag: false
            },
            {
              //
              or_id: 19,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "总和，龙虎",
          object: [
            {
              //
              or_id: 20,
              money: "",
              flag: false
            },
            {
              //
              or_id: 21,
              money: "",
              flag: false
            },
            {
              //
              or_id: 22,
              money: "",
              flag: false
            },
            {
              //
              or_id: 23,
              money: "",
              flag: false
            }
          ]
        },
        {
          name: "总和，龙虎",
          object: [
            {
              or_id: 24,
              money: "",
              flag: false
            },
            {
              or_id: 25,
              money: "",
              flag: false
            },
            {
              or_id: 26,
              money: "",
              flag: false
            }
          ]
        }
      ],
      timer: null,
      // endTime: '2018/1/11 10:00:00',
      h: 0,
      m: 0,
      s: 0,
      c_data: {
        fc_name: "",
        img_path: "",
        qishu: ""
      },
      auto: {
        qishu: null,
        datetime: "",
        ball: []
      },
      close_time: {
        fengpan: "",
        kaijiang: "",
        kaipan: "",
        now_time: ""
      },
      is_wh: false
    };
  },
  created() {
    this.fetchData();
  },
  mounted(){
    this.socket_change(this.$route.query.page);
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route.query.page':function(to,from) {
      this.$root.$off(from);
      this.$root.$off(from+'lefttime');
      this.fetchData();
      this.socket_change(to);
    }
  },
  destroyed(){
    console.log('清除定时器：'+this.timer);
//    window.clearInterval(this.timer);
    if(this.timer){
        window.clearTimeout(this.timer);
        this.timer = null;
    }
    if(!this.isIE9){
      ws.close_ws(false);
    }
    this.$root.$off(this.$route.query.page);
    this.$root.$off(this.$route.query.page+'lefttime');
    cm_cookie.delCookie("top_nav")
  },
  methods:{
    socket_change: function(to){
      if(!this.isIE9){
        let self = this;
        ws.createWebSocket(to,self,true);
        this.$root.$on(to,(e)=>{
          this.auto = e;
        });
        this.$root.$on(to+'lefttime',(e)=>{
          console.log(e);
          this.c_data.qishu = e.qishu;
          this.close_time.fengpan = e.close_time;
          this.close_time.now_time = e.now_time;
          let t1 = e.close_time - e.now_time;
          console.log('是否在维护中：'+this.is_wh);
          if(t1 == 0 && !this.is_wh){
              this.fetchData(2)
          }else if(t1 > 0  && !this.is_wh){
              if(this.timer){
                  window.clearTimeout(this.timer);
                  this.timer = null;
              }
              this.init();
              this.$root.$emit('wh_modal',false);
          }else if(t1 < 0  && !this.is_wh){
              this.h = -1;
              this.m = -1;
              this.s = -1;
              if(this.timer){
                  window.clearTimeout(this.timer);
                  this.timer = null;
              }
              this.$root.$emit('wh_modal',true,true)
          }
        })
      }
    },
    open_way: function (){
      let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=1';
      window.open(page)
    },
    history: function () {
      let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=3';
      window.open(page)
    },
    //aaaaa
    show_rule: function() {
      this.$root.$emit("rule_show", true);
      this.$root.$emit("now_page", this.$route.query.page);
    },
    getRTime: function() {
      this.close_time.now_time += 1;
      var t1 = this.close_time.fengpan * 1000 - this.close_time.now_time * 1000;
      // var d=Math.floor(t/1000/60/60/24);
      this.h = Math.floor((t1 / 1000 / 60 / 60) % 24);
      this.m = Math.floor((t1 / 1000 / 60) % 60);
      this.s = Math.floor((t1 / 1000) % 60);
      // console.log('开奖时间：'+'时：'+h+'；分：'+m+'；秒：'+s);
      console.log('封盘时间：' + '时：' + this.h + '；分：' + this.m + '；秒：' + this.s);
      console.log('定时器id(time_out)：'+this.timer);
      if (this.h == 0 && this.m == 0 && this.s == 0) {
        this.fetchData(2);
        this.$root.$emit("success", true);
      }else if(this.h < 0 && this.m < 0 && this.s < 0){
//        window.clearInterval(this.timer);
        if(this.timer){
            window.clearTimeout(this.timer);
            this.timer = null;
        }
        this.h = -1;
        this.m = -1;
        this.s = -1;
      }
    },
    init: function() {
      this.getRTime();
      this.timer = window.setTimeout(this.init,1000);
    },
    //aaaaaa
    sortNumber: function(a, b) {
      return a.sort - b.sort;
    },
    go_child: function(child) {
      this.$router.push({ name: child, query: { page: this.$route.query.page } });
    },
    fetchData(type) {
      this.$root.$emit('wh_modal',false);
      if(type == 2){
        this.$root.$emit('loading',true,true);
      }else{
        this.$root.$emit('loading',true);
      }
      let body = {
        fc_type: this.$route.query.page
      };
      api.dewdrop(this, body, (res) => {
        if (res.data.ErrorCode == 1) {
        console.log(res);
        this.auto_list = res.data.Data;
        console.log(this.$refs.dewdrop_map);
            api.getgameindex(this, body, res => {
                if (res.data.ErrorCode == 1) {
                if(res.data.is_wh == 2){
                    this.$root.$emit('wh_modal',true);
                    this.is_wh = true
                }else if(res.data.is_wh == 1){
                    this.$root.$emit('wh_modal',false);
                    this.is_wh = false
                }
                this.auto = res.data.Data.auto;
                this.close_time = res.data.Data.closetime;
                if(this.close_time.fengpan - this.close_time.now_time < 0){
                    this.$root.$emit('wh_modal',true,true)
                }
                this.c_data = res.data.Data.c_data;
                let data = res.data.Data.odds;
                data.sort(this.sortNumber);
                this.computed(data);
                this.computedOr(data);
                this.$refs.dewdrop_map.top_go(0);//点击触发露珠图组件头部选中事件
                this.$refs.dewdrop_map.left_go(0);//点击触发露珠图组件左侧选中事件
                if(type == 2){
                    window.setTimeout(() => {
                        this.$root.$emit("loading", false);
                }, 1000)
                }else{
                    this.$root.$emit("loading", false);
                }
                if(!this.is_wh){
                    if(this.close_time.fengpan){
                        if(this.timer){
                            window.clearTimeout(this.timer);
                            this.timer = null;
                        }
                        if(type == 2){
                            window.setTimeout(() => {
                                this.init();
                        }, 1000)
                        }else{
                            this.init();
                        }
                    }
                }
            }
        });
        }
      });
    },

    computed(data) {
      this.$set(this.integrateLists, this.integrateLists);
      this.integrateLists = sscInitial();
      var k = 0;
      //处理一~五球
      for (let l = 0; l < this.integrateLists.length - 10; l++) {
        for (let i = 0; i < this.integrateLists[l].object.length; i++, k++) {
          Object.assign(this.integrateLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.integrateLists[l].object[i].name = name;
        }
      }
      k = 70;
      //处理三球
      for (let l = 5; l < this.integrateLists.length - 7; l++) {
        for (let i = 0; i < this.integrateLists[l].object.length; i++, k++) {
          Object.assign(this.integrateLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.integrateLists[l].object[i].name = name;
        }
      }
      //处理龙虎
      console.log(data);
      for (let l = 8; l < this.integrateLists.length - 5; l++) {
        for (let i = 0; i < this.integrateLists[l].object.length; i++, k++) {
          Object.assign(this.integrateLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.integrateLists[l].object[i].name = name;
        }
      }
      //处理斗牛
      for (let l = 10; l < this.integrateLists.length; l++) {
        for (let i = 0; i < this.integrateLists[l].object.length; i++, k++) {
          Object.assign(this.integrateLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.integrateLists[l].object[i].name = name;
        }
      }
    },
    computedOr(data) {
      this.$set(this.orLists, this.orLists);
      this.orLists = SscOrInitial();

      var k = 10;
      for (let l = 0; l < this.orLists.length - 2; l++, k += 10) {
        for (let i = 0; i < this.orLists[l].object.length; i++, k++) {
          Object.assign(this.orLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.orLists[l].object[i].name = name;
        }
      }

      k = 70;

      for (let l = 5; l < this.orLists.length; l++) {
        for (let i = 0; i < this.orLists[l].object.length; i++, k++) {
          Object.assign(this.orLists[l].object[i], data[k]);
          let name = data[k].remark.slice(
            data[k].remark.search("#") + 1,
            data[k].remark.length
          );
          this.orLists[l].object[i].name = name;
        }
      }
    }
  }
};
</script>

