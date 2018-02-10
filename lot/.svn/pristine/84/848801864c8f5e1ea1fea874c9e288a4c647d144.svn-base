import api from '../../../api/config'
import MeModal from '../../../share_components/bet'
import hint from '../../../share_components/hint_msg'
import share from '../../../share_components/share'
// import back from './back_data'
import {Input, Modal} from 'iview'
export default {
components: { MeModal,'I-Input':Input,Modal},
  data() {
    return {
      money:'',
      modal:false,
      a:''
    }
  },
  props:{
    one:{
      type:Array,
    },
    c_data:{
      type:Object
    }
  },
  created() {
    this.reset();
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
    this.$root.$emit('no_top', true);
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route': 'reset',   // 只有这个页面初始化之后，这个监听事件才开始生效
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
  methods: {
    numClick: function(item, i) {
      console.log(item);
      console.log(item.flag);
      if(item.money && item.flag == true){
        item.flag = false;
        item.money = '';
        this.$refs.myfocus[item.index].blur();
        this.$refs.myfocus[item.index].$refs.input.value = '';
      }else if(item.money == '' && item.flag == false){
        this.$refs.myfocus[item.index].focus();
      }else if (item.flag == true && item.money == ''){
        item.flag = false;
        this.$refs.myfocus[item.index].$refs.input.value = '';
      }else if (item.flag == true){
        item.money = this.money;
      }
    },
    onfocus: function(item){
      console.log('索引：！！！！'+item.index);
      this.$refs.myfocus[item.index].$refs.input.data_onoff = 'true';
      this.a = item.index;
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          if(i != item.index) {
              dom[i].data_onoff = 'false';
          }
      }
      if(item.flag == false && item.money == ''){
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
    go_to: function() {
      this.$root.$emit('c_data',this.c_data);
      let a = this.money + 'a';
      this.money = a.replace(/\D/g, "");
      var kk = 0;
      var is_select = false;
      for (let i = 0; i < this.one.length; i++) {
        for (let j = 0; j < this.one[i].object.length; j++) {
          let b = this.one[i].object[j].money + 'b';
          this.one[i].object[j].money = b.replace(/\D/g, "");
          kk += Number(this.one[i].object[j].money);
          if(this.one[i].object[j].flag){
            is_select = true
          }
        }
      };
      //console.log('kk:' + kk);
      if(is_select){
        if (kk != 0) {
            this.modal = true;
            // document.querySelector('body').style.overflow='hidden';
            this.$root.$emit('id-selected', this.one);
        } else if (kk == 0) {
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
    push_money: function() {
      // this.disabled = true;
      this.money = this.money.replace(/\D/g, "");
      this.computed_money()
    },
    change_money: function () { this.computed_money() },
    add_money: function(type){
      let money = this.money;
      this.money = Number(money) + type;
      this.computed_money()
    },
    //处理金额
    computed_money: function () {
      for (let i = 0; i < this.one.length; i++) {
        for (let j = 0; j < this.one[i].object.length; j++) {
          //添加金额参数入对象
          if (this.one[i].object[j].flag) {
            this.one[i].object[j].money = this.money
          }else if(!this.one[i].object[j].flag){
            this.one[i].object[j].money = ''
          }
        }
      }
    },
    reset: function() {
      this.money = '';
      this.$root.$emit('reset', '');
      for (let i = 0; i < this.one.length; i++) {
        for (let j = 0; j < this.one[i].object.length; j++) {
          this.one[i].object[j].money = '';
          this.one[i].object[j].flag = '';
        }
      }
      let dom = document.querySelectorAll('input');
      for(let i = 0;i < dom.length;i++){
          dom[i].value = '';
          dom[i].data_onoff = 'false';
      }
      this.$root.$emit('clear_key_number','')
    }
  }
};
