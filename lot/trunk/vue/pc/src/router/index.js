import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/view/Home'
import all_menu from '@/view/all_menu'
import pk_content from '@/view/pk_content'
import error from '@/view/error'
import wei_hu from '@/view/wei_hu'
//快乐彩
const happy = () => import(/* webpackChunkName: "klc" */ '@/view/lottery_kind/happy/index')
const happy_one = () => import(/* webpackChunkName: "klc" */ '@/view/lottery_kind/happy/or')
const happy_two = () => import(/* webpackChunkName: "klc" */ '@/view/lottery_kind/happy/selectOne_five')
//时时彩
const SscIndex = () => import (/* webpackChunkName: "ssc" */ '@/view/lottery_kind/often/index') //动态加载组件
const integrate = () => import (/* webpackChunkName: "ssc" */ '@/view/lottery_kind/often/integrate')
const or = () => import (/* webpackChunkName: "ssc" */ '@/view/lottery_kind/often/or')
//排列三，福彩3D
const yb = () => import (/* webpackChunkName: "pl3" */ '@/view/lottery_kind/common/Column3'); //动态加载组件
const yb_one = () => import (/* webpackChunkName: "pl3" */ '@/view/lottery_kind/common/Column-1-3');
const yb_two = () => import (/* webpackChunkName: "pl3" */ '@/view/lottery_kind/common/ColumnIntegration');
//快三
const happy_three = () => import (/* webpackChunkName: "k3" */ '@/view/lottery_kind/happy_three/index'); //动态加载组件
const happyThree_one = () => import (/* webpackChunkName: "k3" */ '@/view/lottery_kind/happy_three/integrate');
//11选5
const elevenIndex = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/index'); //动态加载组件
const e_s_f = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/eleven_select_five');
const eleven_too = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/eleven_too');
const eleven_opt = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/eleven_opt');
const eleven_group = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/eleven_group');
const eleven_just = () => import (/* webpackChunkName: "eleven_xuan5" */ '@/view/lottery_kind/eleven_select_five/eleven_just');
//高频彩
const bj_index = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/index'); //动态加载组件
const bj_2 = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/bj_2');
const bj = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/bj');
const bj_kemp = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/bj_kemp');
const bj_tooface = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/bj_tooface');
//分分彩
const ffIndex = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/hoping/index') //动态加载组件
const ff_integrate = () => import (/* webpackChunkName: "bj" */ '@/view/lottery_kind/bjVue/hoping/integrate')
//快乐十分
const happyTenindex = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/index'); //动态加载组件
const happy_too = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/happy_too');
const happy_sum = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/happy_sum');
const ball_one = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/happy_ball');
const ball_two = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_two');
const ball_three = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_three');
const ball_four = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_four');
const ball_five = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_five');
const ball_six = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_six');
const ball_seven = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_seven');
const ball_eight = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/ball_eight');
const happy_ten = () => import (/* webpackChunkName: "hp10" */ '@/view/lottery_kind/happy_ten/happy_ten');
//幸运彩
const luckyiottery = () => import (/* webpackChunkName: "lucky" */ '@/view/lottery_kind/lucky_lottery/luckyIottery'); //动态加载组件
const pcegg = () => import (/* webpackChunkName: "lucky" */ '@/view/lottery_kind/lucky_lottery/pcEggEgg');
const pkbj28 = () => import (/* webpackChunkName: "lucky" */ '@/view/lottery_kind/lucky_lottery/PKBj28');
const luckyindex = () => import (/* webpackChunkName: "lucky" */ '@/view/lottery_kind/lucky_lottery/luck_another/another_index');
const another_lucky = () => import (/* webpackChunkName: "lucky" */ '@/view/lottery_kind/lucky_lottery/luck_another/another_page');
//六合彩
const mark_six = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/index'); //动态加载组件
const six_one = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/one');
const six_two = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/two');
const six_three = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/three');
const six_four = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/four');
const six_five = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/five');
const six_six = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/six');
const six_seven = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/seven');
const six_eight = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/eight');
const six_nine = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/nine');
const six_ten = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/ten');
const six_eleven = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/eleven');
const six_twelve = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/twelve');
const six_thirteen = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/thirteen');
const six_row = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/row');
const sevencode = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/sevencode');
const totalshaw = () => import (/* webpackChunkName: "liuhecai1" */ '@/view/lottery_kind/mark_six/totalshaw');
const positiveshaw = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/positiveshaw');
const specialhead = () => import (/* webpackChunkName: "liuhecai2" */ '@/view/lottery_kind/mark_six/specialhead');
//投注记录
const record_index = () => import (/* webpackChunkName: "record" */ '@/view/record/index'); //动态加载组件
const record = () => import (/* webpackChunkName: "record" */ '@/view/record/record');
const cash = () => import (/* webpackChunkName: "record" */ '@/view/record/cash');
//报表统计
const report_index = () => import (/* webpackChunkName: "report" */ '@/view/report/index'); //动态加载组件
const report = () => import (/* webpackChunkName: "report" */ '@/view/report/report');
const report_one = () => import (/* webpackChunkName: "report" */ '@/view/report/report_one');

//消息中心
const message_index = () => import (/* webpackChunkName: "message" */ '@/view/message/index'); //动态加载组件
const notice = () => import (/* webpackChunkName: "message" */ '@/view/message/notice');
const personal = () => import (/* webpackChunkName: "message" */ '@/view/message/personal');
//规则
const game_rule = () => import (/* webpackChunkName: "game_rule" */ '@/share_components/game_rule');

Vue.use(Router);

export default new Router({
  routes: [
    {
    path: '/',
    component: Home,
    children: [
      {
        path: '/',
        name: 'home',
        component: all_menu
      },
      {
        path: '/pk_content',
        name: 'pk_content',
        component: pk_content,
        children: [
          {
            path: '/pk_content/message_index',
            name: 'message_index',
            component: message_index,
            children: [{
                path: 'notice',
                name: 'notice',
                component: notice,
              },
              {
                path: 'personal',
                name: 'personal',
                component: personal,
              }
            ]
          },
          {
            //报表统计
            path: '/pk_content/record_index',
            name: 'record_index',
            component: record_index,
            children: [{
                path: 'record',
                name: 'record',
                component: record,
              },
              {
                path: 'cash',
                name: 'cash',
                component: cash,
              },
            ]
          },
          {
            path: '/pk_content/report_index',
            name: 'report_index',
            component: report_index,
            children: [{
                path: 'report',
                name: 'report',
                component: report,
              },
              {
                path: 'report_one',
                name: 'report_one',
                component: report_one,
              },
            ]
          },
          {
            path: '/pk_content/bj_10',
            name: 'bj_10_index',
            component: bj_index,
            children: [{
                path: 'bj_tooface_1',
                name: 'bj_10',
                component: bj_tooface
              },
              {
                path: 'bj_kemp_2',
                name: 'bj_kemp',
                component: bj_kemp
              },
              {
                path: 'bj_1_5_3',
                name: 'bj_1_5',
                component: bj
              },
              {
                path: 'bj_6_10_4',
                name: 'bj_6_10',
                component: bj_2
              },
            ]
          },
          {
            path: '/pk_content/happy',
            name: 'happy',
            component: happy,
            children: [
              {
                path: 'klc_one_1',
                name: 'bj_kl8',
                component: happy_one
              },
              {
                path: 'klc_two_2',
                name: 'or',
                component: happy_two
              },

            ]
          },
          {
            path: '/pk_content/yb',
            name: 'yb_index',
            component: yb,
            children: [
              {
                path: 'yb_one_1',
                name: 'pl_3',
                component: yb_one
              },
              {
                path: 'yb_two_2',
                name: 'yb_two',
                component: yb_two
              },
            ]
          },
          {
            path: '/pk_content/select_11',
            name: 'select_index',
            component: elevenIndex,
            children: [{
                path: 'index_11_1',
                name: 'gd_11',
                component: eleven_too,
              },
              {
                path: 'e_s_f_2',
                name: '11_e_s_f',
                component: e_s_f,
              },
              {
                path: '11_opt_3',
                name: '11_opt',
                component: eleven_opt,
              },
              {
                path: '11_group_4',
                name: '11_group',
                component: eleven_group,
              },
              {
                path: '11_just_5',
                name: '11_just',
                component: eleven_just,
              }
            ]
          },
          {
            path: '/pk_content/happyTen',
            name: 'happyTenindex',
            component: happyTenindex,
            children: [
              {
                path: 'happy_too_1',
                name: 'gd_ten',
                component: happy_too,
              },
              {
                path: 'happy_sum_10',
                name: 'happy_sum',
                component: happy_sum
              },
              {
                path: 'ball_one_2',
                name: 'ball_one',
                component: ball_one
              },
              {
                path: 'ball_two_3',
                name: 'ball_two',
                component: ball_two
              },
              {
                path: 'ball_three_4',
                name: 'ball_three',
                component: ball_three
              },
              {
                path: 'ball_four_5',
                name: 'ball_four',
                component: ball_four
              },
              {
                path: 'ball_five_6',
                name: 'ball_five',
                component: ball_five
              },
              {
                path: 'ball_six_7',
                name: 'ball_six',
                component: ball_six
              },
              {
                path: 'ball_seven_8',
                name: 'ball_seven',
                component: ball_seven
              },
              {
                path: 'ball_eight_9',
                name: 'ball_eight',
                component: ball_eight
              },
              {
                path: 'happy_ten_11',
                name: 'happy_ten',
                component: happy_ten
              },
            ]
          },
          {
            path: '/pk_content/ssc',
            name: 'Ssc_Index',
            component: SscIndex,
            children: [{
                path: 'ssc_index_1',
                name: 'cq_ssc',
                component: integrate
              },
              {
                path: 'ssc_or_2',
                name: 'ssc_or',
                component: or
              },
            ]
          },
          {
            path: '/pk_content/ff_c',
            name: 'ffIndex',
            component: ffIndex,
            children: [{
              path: 'ffc_one_1',
              name: 'ffc_o',
              component: ff_integrate
            },
            ]
          },
          {
            path:'/pk_content/luck_another',
            name:'luck_another',
            component: luckyindex,
            children:[
              {
                path:'luck_28_1',
                name:'dm_28',
                component:another_lucky
              }
            ]
          },
          {
            path: '/pk_content/luckyiottery',
            name: 'luckyiottery',
            component: luckyiottery,
            children: [
              {
                path: 'pc_28_1',
                name: 'pc_28',
                component: pcegg
              },
              {
                path: 'bj_28_1',
                name: 'bj_28',
                component: pkbj28
              },
            ]
          },
          {
            path: '/pk_content/mark_six',
            name: 'mark_six',
            component: mark_six,
            children: [{
                path: 'six_one_1',
                name: 'liuhecai',
                component: six_one
              },
              {
                path: 'six_two_2',
                name: 'six_two',
                component: six_two
              },
              {
                path: 'six_three_3',
                name: 'six_three',
                component: six_three
              },
              {
                path: 'six_four_4',
                name: 'six_four',
                component: six_four
              },
              {
                path: 'six_five_5',
                name: 'six_five',
                component: six_five
              },
              {
                path: 'six_six_6',
                name: 'six_six',
                component: six_six
              },
              {
                path: 'six_seven_7',
                name: 'six_seven',
                component: six_seven
              },
              {
                path: 'six_eight_8',
                name: 'six_eight',
                component: six_eight
              },
              {
                path: 'six_nine_9',
                name: 'six_nine',
                component: six_nine
              },
              {
                path: 'six_ten_10',
                name: 'six_ten',
                component: six_ten
              },
              {
                path: 'six_eleven_11',
                name: 'six_eleven',
                component: six_eleven
              },
              {
                path: 'six_twelve_12',
                name: 'six_twelve',
                component: six_twelve
              },
              {
                path: 'six_thirteen_13',
                name: 'six_thirteen',
                component: six_thirteen
              },
              {
                path: 'six_row_14',
                name: 'six_row',
                component: six_row
              },
              {
                path: 'sevencode_17',
                name: 'sevencode',
                component: sevencode
              },
              {
                path: 'totalshaw_18',
                name: 'totalshaw',
                component: totalshaw
              },
              {
                path: 'positiveshaw_15',
                name: 'positiveshaw',
                component: positiveshaw
              },
              {
                path: 'specialhead_16',
                name: 'specialhead',
                component: specialhead
              },
            ]
          },
          {
            path: '/pk_content/happy_three',
            name: 'k3_index',
            component: happy_three,
            children: [{
                path: 'k3_one_1',
                name: 'gx_k3',
                component: happyThree_one
              },
            ]
          },
        ]
      },
      {
        path: 'error',
        name: 'error',
        component: error
      },
      {
        path: 'game_rule',
        name: 'game_rule',
        component: game_rule
      },
    ]
  },
    {
      path: '/wei_hu',
      name: 'wei_hu',
      component: wei_hu
    },
  ]
})
