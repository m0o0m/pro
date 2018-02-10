/**
 * Created by huygo on 2018/1/5.
 */
//JS操作cookies方法!

//写cookies
const com_cookie = {
  setCookie: function (name, value) {
    var Days = 30;
    var exp = new Date();
    exp.setTime(exp.getTime() + Days * 24 * 60 * 60 * 1000);
    document.cookie = name + "=" + escape(value) + ";expires=" + exp.toGMTString()+ ";path=/";
  },
//读取cookies
  getCookie: function (name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
      return unescape(arr[2]);
    else
      return null;
  },

//删除cookies
  delCookie: function (name) {
    var exp = new Date();
    exp.setTime(exp.getTime() - 1);
    var cval = com_cookie.getCookie(name);
    if (cval != null){
      document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString()+ ";path=/";
    }
  }
};
export default com_cookie;
//使用示例
// setCookie("name","hayden");
// alert(getCookie("name"));

