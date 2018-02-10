<template lang="html">
  <div class="six_one">
    <nav-center @go_child="go_child" :menus="lists" :margin="false"></nav-center>
    <div class="footer1 clearfix ">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-blur="onblur_top(0)" @on-focus="onfocus_top(0)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <div class="content">
      <div class="top">
        <ul>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">金额</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">金额</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">金额</li>
          <li class="one">号码</li>
          <li class="two">赔率</li>
          <li class="three">金额</li>
        </ul>
      </div>
      <div class="table">
        <div class="bottom" v-for="item in lists1">
          <ul v-for="key in item.object" :class="[key.flag?'table-current':'']">
            <li @click="pour(key)" :class="[key.num?'one':'border1_none']">
              <span :class="[key.color?key.color:'']">{{key.input_name}}</span>
            </li>
            <li @click="pour(key)" :class="[key.num?'two':'border2_none']">{{key.odd}}</li>
            <li @click.self="pour(key)" :class="[key.num?'three':'border3_none']">
              <I-Input :value="key.money" @on-blur="onblur(key)" @on-keydown="tab_now(key)" @on-change="onchange(key)" :maxlength='9' ref="input" @on-focus="onfocus(key)" style="width: 45px"  v-if="key.num" v-model="key.money" @on-keyup="gogo(key)" @on-afterpaste="gogo(key)" size="small"></I-Input>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <right-config @config_num="config_num" @config_num1="config_num1" :mores="mores" :bottom_mores="bottom_mores"></right-config>
    <div class="footer1 clearfix" style="float:left;margin-top: 10px">
      <div class="img">
        <span class="s1" @click="add_money(10)"></span>
        <span class="s2" @click="add_money(20)"></span>
        <span class="s3" @click="add_money(50)"></span>
        <span class="s4" @click="add_money(100)"></span>
        <span class="s5" @click="add_money(500)"></span>
        <span class="s6" @click="add_money(1000)"></span>
      </div>
      <div class="one"><span class="left">金额￥</span><I-Input :value="money" @on-focus="onfocus_top(1)"  @on-blur="onblur_top(1)" @on-change="change_money()" :maxlength='9' style="width: 100px" size="small" v-model="money" @on-keyup="push_money()" @on-afterpaste="push_money()"></I-Input>
      </div>
      <button type="button" class="two" style="padding: 8px;" @click="go_to()">立即下注</button>
      <button type="button" class="two" style="padding: 8px;margin-right:20px;" @click="reset()">重置</button>
    </div>
    <Me-modal :modal='modal' @cancel="cancel"></Me-modal>
  </div>
</template>

<script>
function get_lists1() {
  return [
    {
      name:'1',object:[
      {
        index:0,num:'1',flag:false,number:'48.8',money:''
      },
      {
        index:1,num:'2',flag:false,number:'48.8',money:''
      },
      {
        index:2,num:'3',flag:false,number:'48.8',money:''
      },
      {
        index:3,num:'4',flag:false,number:'48.8',money:''
      },
      {
        index:4,num:'5',flag:false,number:'48.8',money:''
      },
      {
        index:5,num:'6',flag:false,number:'48.8',money:''
      },
      {
        index:6,num:'7',flag:false,number:'48.8',money:''
      },
      {
        index:7,num:'8',flag:false,number:'48.8',money:''
      },
      {
        index:8,num:'9',flag:false,number:'48.8',money:''
      },
      {
        index:9,num:'10',flag:false,number:'48.8',money:''
      },
      {
        index:10,num:'11',flag:false,number:'48.8',money:''
      },
      {
        index:11,num:'12',flag:false,number:'48.8',money:''
      },
      {
        index:12,num:'13',flag:false,number:'48.8',money:''
      },
    ]
    },
    {
      name:'2',object:[
      {
        index:13,num:'14',flag:false,number:'48.8',money:''
      },
      {
        index:14,num:'15',flag:false,number:'48.8',money:''
      },
      {
        index:15,num:'16',flag:false,number:'48.8',money:''
      },
      {
        index:16,num:'17',flag:false,number:'48.8',money:''
      },
      {
        index:17,num:'18',flag:false,number:'48.8',money:''
      },
      {
        index:18,num:'19',flag:false,number:'48.8',money:''
      },
      {
        index:19,num:'20',flag:false,number:'48.8',money:''
      },
      {
        index:20,num:'21',flag:false,number:'48.8',money:''
      },
      {
        index:21,num:'22',flag:false,number:'48.8',money:''
      },
      {
        index:22,num:'23',flag:false,number:'48.8',money:''
      },
      {
        index:23,num:'24',flag:false,number:'48.8',money:''
      },
      {
        index:24,num:'25',flag:false,number:'48.8',money:''
      },
      {
        index:25,num:'26',flag:false,number:'48.8',money:''
      },
    ]
    },
    {
      name:'2',object:[
      {
        index:26,num:'27',flag:false,number:'48.8',money:''
      },
      {
        index:27,num:'28',flag:false,number:'48.8',money:''
      },
      {
        index:28,num:'29',flag:false,number:'48.8',money:''
      },
      {
        index:29,num:'30',flag:false,number:'48.8',money:''
      },
      {
        index:30,num:'31',flag:false,number:'48.8',money:''
      },
      {
        index:31,num:'32',flag:false,number:'48.8',money:''
      },
      {
        index:32,num:'33',flag:false,number:'48.8',money:''
      },
      {
        index:33,num:'34',flag:false,number:'48.8',money:''
      },
      {
        index:34,num:'35',flag:false,number:'48.8',money:''
      },
      {
        index:35,num:'36',flag:false,number:'48.8',money:''
      },
      {
        index:36,num:'37',flag:false,number:'48.8',money:''
      },
      {
        index:37,num:'38',flag:false,number:'48.8',money:''
      },
      {
        index:38,num:'39',flag:false,number:'48.8',money:''
      },
    ]
    },
    {
      name:'2',object:[
      {
        index:39,num:'40',flag:false,number:'48.8',money:''
      },
      {
        index:40,num:'41',flag:false,number:'48.8',money:''
      },
      {
        index:41,num:'42',flag:false,number:'48.8',money:''
      },
      {
        index:42,num:'43',flag:false,number:'48.8',money:''
      },
      {
        index:43,num:'44',flag:false,number:'48.8',money:''
      },
      {
        index:44,num:'45',flag:false,number:'48.8',money:''
      },
      {
        index:45,num:'46',flag:false,number:'48.8',money:''
      },
      {
        index:46,num:'47',flag:false,number:'48.8',money:''
      },
      {
        index:47,num:'48',flag:false,number:'48.8',money:''
      },
      {
        index:48,num:'49',flag:false,number:'48.8',money:''
      },
      {
        index:'',num:'',flag:false,number:'',money:''
      },
      {
        index:'',num:'',flag:false,number:'',money:''
      },
      {
        index:'',num:'',flag:false,number:'',money:''
      },
    ]
    },
  ]
}
import api from '../../../api/config'
import MeModal from '../../../share_components/bet'
import navCenter from './components/nav'
import rightConfig from './components/right_config'
import {Modal,Input} from 'iview';
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
export default {
  components: {
    navCenter,
    rightConfig,
    MeModal,
    Modal,'I-Input':Input
  },
  data() {
    return {
      money:'',
      back_data:'',
      modal: false,
      animal_config:{},
      mores:[
        {
          arr:[
            {one:'红大',index:0},
            {one:'红合单',index:1},
            {one:'红双',index:2},
          ],
        },
        {
          arr:[
            {one:'红小',index:3},
            {one:'红合双',index:4},
            {one:'红单',index:5},
          ],
        },
        {
          arr:[
            {one:'绿大',index:6},
            {one:'绿合单',index:7},
            {one:'绿双',index:8},
          ],
        },
        {
          arr:[
            {one:'绿单',index:9},
            {one:'绿合双',index:10},
            {one:'绿小',index:11},
          ],
        },
        {
          arr:[
            {one:'蓝大',index:12},
            {one:'蓝合单',index:13},
            {one:'蓝双',index:14},
          ],
        },
        {
          arr:[
            {one:'蓝单',index:15},
            {one:'蓝合双',index:16},
            {one:'蓝小',index:17},
          ],
        },
        {
          arr:[
            {four:'大',index:18},
            {four:'小',index:19},
            {four:'单',index:20},
            {four:'双',index:21},
          ],
        },
      ],
      bottom_mores:[
        {
          arr:[
            {four:'鼠',index:22,state:false},
            {four:'牛',index:23,state:false},
            {four:'虎',index:24,state:false},
            {four:'兔',index:25,state:false},
          ],
        },
        {
          arr:[
            {four:'龙',index:26,state:false},
            {four:'蛇',index:27,state:false},
            {four:'马',index:28,state:false},
            {four:'羊',index:29,state:false},
          ],
        },
        {
          arr:[
            {four:'猴',index:30,state:false},
            {four:'鸡',index:31,state:false},
            {four:'狗',index:32,state:false},
            {four:'猪',index:33,state:false},
          ],
        },
      ],
      lists:[
        {name:'正1特',item:'0'},
        {name:'正2特',item:'1'},
        {name:'正3特',item:'2'},
        {name:'正4特',item:'3'},
        {name:'正5特',item:'4'},
        {name:'正6特',item:'5'},
      ],
      lists1:[
        {
          name:'1',object:[
          {
            index:0,num:'1',flag:false,number:'48.8',money:''
          },
          {
            index:1,num:'2',flag:false,number:'48.8',money:''
          },
          {
            index:2,num:'3',flag:false,number:'48.8',money:''
          },
          {
            index:3,num:'4',flag:false,number:'48.8',money:''
          },
          {
            index:4,num:'5',flag:false,number:'48.8',money:''
          },
          {
            index:5,num:'6',flag:false,number:'48.8',money:''
          },
          {
            index:6,num:'7',flag:false,number:'48.8',money:''
          },
          {
            index:7,num:'8',flag:false,number:'48.8',money:''
          },
          {
            index:8,num:'9',flag:false,number:'48.8',money:''
          },
          {
            index:9,num:'10',flag:false,number:'48.8',money:''
          },
          {
            index:10,num:'11',flag:false,number:'48.8',money:''
          },
          {
            index:11,num:'12',flag:false,number:'48.8',money:''
          },
          {
            index:12,num:'13',flag:false,number:'48.8',money:''
          },
        ]
        },
        {
          name:'2',object:[
          {
            index:13,num:'14',flag:false,number:'48.8',money:''
          },
          {
            index:14,num:'15',flag:false,number:'48.8',money:''
          },
          {
            index:15,num:'16',flag:false,number:'48.8',money:''
          },
          {
            index:16,num:'17',flag:false,number:'48.8',money:''
          },
          {
            index:17,num:'18',flag:false,number:'48.8',money:''
          },
          {
            index:18,num:'19',flag:false,number:'48.8',money:''
          },
          {
            index:19,num:'20',flag:false,number:'48.8',money:''
          },
          {
            index:20,num:'21',flag:false,number:'48.8',money:''
          },
          {
            index:21,num:'22',flag:false,number:'48.8',money:''
          },
          {
            index:22,num:'23',flag:false,number:'48.8',money:''
          },
          {
            index:23,num:'24',flag:false,number:'48.8',money:''
          },
          {
            index:24,num:'25',flag:false,number:'48.8',money:''
          },
          {
            index:25,num:'26',flag:false,number:'48.8',money:''
          },
        ]
        },
        {
          name:'2',object:[
          {
            index:26,num:'27',flag:false,number:'48.8',money:''
          },
          {
            index:27,num:'28',flag:false,number:'48.8',money:''
          },
          {
            index:28,num:'29',flag:false,number:'48.8',money:''
          },
          {
            index:29,num:'30',flag:false,number:'48.8',money:''
          },
          {
            index:30,num:'31',flag:false,number:'48.8',money:''
          },
          {
            index:31,num:'32',flag:false,number:'48.8',money:''
          },
          {
            index:32,num:'33',flag:false,number:'48.8',money:''
          },
          {
            index:33,num:'34',flag:false,number:'48.8',money:''
          },
          {
            index:34,num:'35',flag:false,number:'48.8',money:''
          },
          {
            index:35,num:'36',flag:false,number:'48.8',money:''
          },
          {
            index:36,num:'37',flag:false,number:'48.8',money:''
          },
          {
            index:37,num:'38',flag:false,number:'48.8',money:''
          },
          {
            index:38,num:'39',flag:false,number:'48.8',money:''
          },
        ]
        },
        {
          name:'2',object:[
          {
            index:39,num:'40',flag:false,number:'48.8',money:''
          },
          {
            index:40,num:'41',flag:false,number:'48.8',money:''
          },
          {
            index:41,num:'42',flag:false,number:'48.8',money:''
          },
          {
            index:42,num:'43',flag:false,number:'48.8',money:''
          },
          {
            index:43,num:'44',flag:false,number:'48.8',money:''
          },
          {
            index:44,num:'45',flag:false,number:'48.8',money:''
          },
          {
            index:45,num:'46',flag:false,number:'48.8',money:''
          },
          {
            index:46,num:'47',flag:false,number:'48.8',money:''
          },
          {
            index:47,num:'48',flag:false,number:'48.8',money:''
          },
          {
            index:48,num:'49',flag:false,number:'48.8',money:''
          },
          {
            index:'',num:'',flag:false,number:'',money:''
          },
          {
            index:'',num:'',flag:false,number:'',money:''
          },
          {
            index:'',num:'',flag:false,number:'',money:''
          },
        ]
        },
      ],
      a:''
    }
  },
  created(){
    this.fetchData();
    this.$root.$on('success',(e)=>{
      if(e){
        this.modal = false;
        this.reset()
      }
    });
    this.$root.$on('this_money',(e)=>{
        this.money = e
    });
  },
  mounted(){
    this.$root.$emit('no_top',false);
    this.$root.$emit('child_change',0);
    this.$root.$on('time_out',(e)=>{
      if(e){
        this.fetchData(2)
      }
    })
  },
  destroyed(){
    this.$root.$off('time_out')
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route': 'fetchData',  // 只有这个页面初始化之后，这个监听事件才开始生效
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
  methods:{
    go_child: function(e){
      console.log(e);
      this.reset();
      this.computed(this.back_data,e);
    },
    config_num: function (e) {
      this.clear();
      for(let i = 0;i < this.bottom_mores.length;i++){
        for(let j = 0;j<this.bottom_mores[i].arr.length;j++){
          this.bottom_mores[i].arr[j].state = false
        }
      }
      for (let i = 0; i < this.lists1.length; i++) {
          for (let j = 0; j < this.lists1[i].object.length; j++) {
            let input_name = Number(this.lists1[i].object[j].input_name);
            //合值
            var c = input_name%10+parseInt(input_name/10 % 10);
            if(e.index == 0){
              if(this.lists1[i].object[j].color == 'red'&& input_name > 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 1){
              if(this.lists1[i].object[j].color == 'red'&& c%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 2){
              if(this.lists1[i].object[j].color == 'red'&& input_name%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 3){
              if(this.lists1[i].object[j].color == 'red'&& input_name <= 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 4){
              if(this.lists1[i].object[j].color == 'red'&& c%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 5){
              if(this.lists1[i].object[j].color == 'red'&& input_name%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 6){
              if(this.lists1[i].object[j].color == 'green'&& input_name > 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 7){
              if(this.lists1[i].object[j].color == 'green'&& c%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 8){
              if(this.lists1[i].object[j].color == 'green'&& input_name%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 9){
              if(this.lists1[i].object[j].color == 'green'&& input_name%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 10){
              if(this.lists1[i].object[j].color == 'green'&& c%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 11){
              if(this.lists1[i].object[j].color == 'green'&& input_name <= 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 12){
              if(this.lists1[i].object[j].color == 'blue'&& input_name > 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 13){
              if(this.lists1[i].object[j].color == 'blue'&& c%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 14){
              if(this.lists1[i].object[j].color == 'blue'&& input_name%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 15){
              if(this.lists1[i].object[j].color == 'blue'&& input_name%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 16){
              if(this.lists1[i].object[j].color == 'blue'&& c%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 17){
              if(this.lists1[i].object[j].color == 'blue'&& input_name <= 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 18){
              if(input_name > 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 19){
              if(input_name <= 24){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 20){
              if(input_name%2 == 1){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }else if(e.index == 21){
              if(input_name%2 == 0){
                this.lists1[i].object[j].flag = true;
                this.lists1[i].object[j].money = this.money;
              }
            }
          }
        }
    },
    config_num1: function (e) {
//      console.log(this.$refs.top_config);
      var dom = document.querySelectorAll('.top_config');
      for(let i = 0;i < dom.length;i++){
//        console.log(dom[i].style);
        if(dom[i].style.color == 'red'){
          dom[i].style = '';
          this.clear();
        }
      }
      for (let i = 0; i < this.lists1.length; i++) {
          for (let j = 0; j < this.lists1[i].object.length; j++) {
            let input_name = Number(this.lists1[i].object[j].input_name);
            if(e.index == 22){//鼠
              for(let b=0;b<this.animal_config.mouse.length;b++){
                if(input_name == this.animal_config.mouse[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 23){//牛
              for(let b=0;b<this.animal_config.cattle.length;b++){
                if(input_name == this.animal_config.cattle[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 24){//虎
              for(let b=0;b<this.animal_config.tiger.length;b++){
                if(input_name == this.animal_config.tiger[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 25){//兔
              for(let b=0;b<this.animal_config.rabbit.length;b++){
                if(input_name == this.animal_config.rabbit[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 26){//龙
              for(let b=0;b<this.animal_config.dragon.length;b++){
                if(input_name == this.animal_config.dragon[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 27){//蛇
              for(let b=0;b<this.animal_config.snake.length;b++){
                if(input_name == this.animal_config.snake[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 28){//马
              for(let b=0;b<this.animal_config.horse.length;b++){
                if(input_name == this.animal_config.horse[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 29){//羊
              for(let b=0;b<this.animal_config.sheep.length;b++){
                if(input_name == this.animal_config.sheep[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 30){//猴
              for(let b=0;b<this.animal_config.monkey.length;b++){
                if(input_name == this.animal_config.monkey[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 31){//鸡
              for(let b=0;b<this.animal_config.chicken.length;b++){
                if(input_name == this.animal_config.chicken[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 32){//狗
              for(let b=0;b<this.animal_config.dog.length;b++){
                if(input_name == this.animal_config.dog[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }else if(e.index == 33){//猪
              for(let b=0;b<this.animal_config.pig.length;b++){
                if(input_name == this.animal_config.pig[b]){
                  this.lists1[i].object[j].flag = e.state;
                  if(e.state){
                    this.lists1[i].object[j].money = this.money;
                  }else{
                    this.lists1[i].object[j].money = '';
                  }
                }
              }
            }
          }
        }
    },
    //清空每个球的颜色和金额
    clear: function () {
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          this.lists1[i].object[j].flag = '';
          this.lists1[i].object[j].money = ''
        }
      }
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
    //点击下注
    go_to:function () {
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          let b = this.lists1[i].object[j].money + 'b';
          this.lists1[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.lists1[i].object[j].money);
          if(this.lists1[i].object[j].flag){
            is_select = true
          }
        }
      }
      console.log('kk:' + kk);
      if(is_select){
        if (kk != 0) {
          this.modal = true;
//          document.querySelector('body').style.overflow='hidden';
          this.$root.$emit('id-selected',this.lists1);
        }else if (kk == 0) {
          this.$Modal.warning({
            content: hint.money_null
          });
          window.setTimeout(() => {
            this.$Modal.remove()
        }, share.Prompt)
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
    cancel: function(item) {
      this.modal = false;
      document.querySelector('body').style.overflow='auto'
    },
    add_money: function(type){
      let money = this.money;
      this.money = Number(money)+type;
      this.computed_money()
    },
    push_money: function() {
      // this.disabled = true;
      this.money = this.money.replace(/\D/g, "");
      this.computed_money()
    },
    change_money: function () { this.computed_money() },
    computed_money: function(){
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          //添加金额参数入对象
          if (this.lists1[i].object[j].flag) {
            this.lists1[i].object[j].money = this.money
          }else if(!this.lists1[i].object[j].flag){
            this.lists1[i].object[j].money = ''
          }
        }
      }
    },
    fetchData: function(type){
      this.reset();
      type==2?this.$root.$emit('loading',true,true):this.$root.$emit('loading',true);
      console.log('this_page:'+this.$route.query.page);
      this.$root.$emit('child_change',0);
      let body = {
        'fc_type':this.$route.query.page,
        'gameplay':167,
        'pankou':'A'
      };
      api.getgameindex(this,body,(res)=>{
          this.$root.$emit('only_back',res,type);
        if(res.data.ErrorCode == 1){
          this.animal_config = res.data.shengxiao;
          if(type == 2){
            window.setTimeout(() => {
              this.$root.$emit("loading", false);
          }, 1000)
          }else{
            this.$root.$emit("loading", false);
          }
          this.$root.$emit('c_data',res.data.Data.c_data);
          let back_data = res.data.Data.odds;
          back_data.sort(this.sortNumber);
          this.back_data = back_data;
          this.computed(this.back_data,'0');
        }
      })
    },
    sortNumber: function(a,b){
      return a.sort - b.sort
    },
    //对象的拼接
    extend: function(target, source) {
      for (var obj in source) {
        target[obj] = source[obj];
      }
      return target;
    },
    computed: function(data,type){
      switch(type){
        case '0':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[13+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[26+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[39+i]);
            }
          }
          break;
        case '1':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[49+i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[62+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[75+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[88+i]);
            }
          }
          break;
        case '2':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[98+i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[111+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[124+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[137+i]);
            }
          }
          break;
        case '3':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[147+i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[160+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[173+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[186+i]);
            }
          }
          break;
        case '4':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[196+i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[209+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[222+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[235+i]);
            }
          }
          break;
        case '5':
          this.lists1 = get_lists1();
          for(let i = 0;i < this.lists1[0].object.length;i++){
            this.extend(this.lists1[0].object[i],data[245+i]);
          }
          for(let i = 0;i < this.lists1[1].object.length;i++){
            this.extend(this.lists1[1].object[i],data[258+i]);
          }
          for(let i = 0;i < this.lists1[2].object.length;i++){
            this.extend(this.lists1[2].object[i],data[271+i]);
          }
          for(let i = 0;i < this.lists1[3].object.length;i++){
            if(i == 10){
              break;
            }else{
              this.extend(this.lists1[3].object[i],data[284+i]);
            }
          }
          break;
      }
    },
    pour: function(item) {
        console.log(item);
        console.log(item.flag);
        if(item.money && item.flag == true){
          item.flag = false;
          item.money = '';
          this.$refs.input[item.index].blur();
          this.$refs.input[item.index].$refs.input.value = '';
        }else if(item.money == '' && item.flag == false){
          this.$refs.input[item.index].focus();
        }else if (item.flag == true && item.money == ''){
          item.flag = false;
          this.$refs.input[item.index].$refs.input.value = '';
        }else if (item.flag == true){
          item.money = this.money;
        }
      },
      onfocus: function(item){
        this.$refs.input[item.index].$refs.input.data_onoff = 'true';
        this.a = item.index;
        let dom = document.querySelectorAll('input');
        for(let i = 0;i < dom.length;i++){
            if(i != item.index+1) {
                dom[i].data_onoff = 'false';
            }
        }
        if(item.flag == false && item.money == ''){
          item.flag = true;
          item.money = this.money;
        }
      },
      onblur: function (key) {
          key.money = this.$refs.input[key.index].$refs.input.value;
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
    //重置
    reset: function () {
      //清空生肖控制状态
      for(let i = 0;i < this.bottom_mores.length;i++){
        for(let j = 0;j<this.bottom_mores[i].arr.length;j++){
          this.bottom_mores[i].arr[j].state = false
        }
      }
      //清空颜色，数字控制状态
      var dom = document.querySelectorAll('.top_config');
      for(let i = 0;i < dom.length;i++){
        if(dom[i].style.color == 'red'){
          dom[i].style = '';
        }
      }
      for (let i = 0; i < this.lists1.length; i++) {
        for (let j = 0; j < this.lists1[i].object.length; j++) {
          this.lists1[i].object[j].money = '';
          this.lists1[i].object[j].flag = '';
        }
      }
      this.money = '';
      this.$root.$emit('reset', '');
      let dom1 = document.querySelectorAll('input');
      for(let i = 0;i < dom1.length;i++){
          dom1[i].value = '';
          dom1[i].data_onoff = 'false';
      }
      this.$root.$emit('clear_key_number','')
    }
  }
}
</script>
<style lang="scss" src="../../../assets/css/six_three.scss" scoped></style>
