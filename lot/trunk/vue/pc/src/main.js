// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import '././assets/css/Reset.css'
import '././assets/css/Default.css'
import 'iview/dist/styles/iview.css'//全局加载iview
import '././assets/css/pk_time.scss'
import {Modal} from 'iview';
import {Message} from 'iview';
// import iView from 'iview'//全局加载iView
import 'babel-polyfill'
Vue.use(VueAxios, axios);
axios.defaults.timeout =  30000;//设置axios超时时间
// Vue.use(iView);//全局加载iView
Vue.config.productionTip = false;
Vue.prototype.$Modal = Modal;
Vue.prototype.$Message = Message;
//全局配置请求前缀
// Vue.axios.defaults.baseURL = basePath;//设置跨域请求的时候要注释掉 打包时要解开
/* eslint-disable no-new */
new Vue({el: '#app', router,
 // data: {
 //    eventHub: new Vue()
 //  },
  template: '<App/>', components: {
    App
  }})
