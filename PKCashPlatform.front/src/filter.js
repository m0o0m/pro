
"use strict";
angular.module('app.filter', []);
angular.module('app.filter')
    .filter('fiterTime',function(){

        return function(nS) {
            return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ');
        }
    })

    .filter('outStatus',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="已出款";
            }else if(value===2){
                status="预备出款";
            }else if(value===3){
                status="已取消";
            }else if(value===4){
                status="已拒绝";
            }else if(value===5){
                //status="待审核";
                status="未出款";
            }
            return status;
        }
    })

    .filter('noticeType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="类型一";
            }else if(value===2){
                status="类型二";
            }else if(value===3){
                status="类型三";
            }else if(value===4){
                status="类型四";
            }
            return status;
        }
    })

    .filter('fiterStatused',function () {

        return function (value) {
            var statused = "";
            if(value=== 1){
                statused = "启用"
            }else if(value===2){
                statused = "禁用"
            }
            return statused;
        }
    })
    .filter('fiterCX',function () {

        return function (value) {
            var status = "";
            if(value=== 1){
                status = "返佣冲销"
            }else if(value===2){
                status="未冲销"
            }
            return status;
        }
    })
    .filter('YorN',function () {

        return function (value) {
            var statused = "";
            if(value=== 1){
                statused = "是"
            }else if(value===2){
                statused="否"
            }
            return statused;
        }
    })
    .filter('currentstate',function () {
        return function (value) {
            var statused = "";
            if(value=== 1){
                statused = "启用"
            }else if(value=== 2) {
                statused = "禁用"
            }
            return statused;
        }
    })
    .filter('fiterStatuseds',function () {
        return function (value) {
            var statuseds = "";
            if(value=== 1){
                statuseds = "在线"
            }else {
                statuseds="离线"
            }
            return statuseds;
        }
    })
    .filter('submit',function () {
        return function (Submit) {
            var Submit = "";
            if(Submit=== 1){
                Submit = "提交"
            }else {
                Submit="未提交"
            }
            return Submit;
        }
    })
    .filter('fiterApp',function () {
        return function (value) {
            var status = "";
            if(value===1){
                status = "已通过"
            }else {
                status="未处理"
            }
            return status;
        }
    })
    .filter('formatDate',function () {
        return function (value) {
            var myDate = new Date(value*1000);
            var year = myDate.getFullYear().toString();
            var month = myDate.getMonth() + 1;
            var day = myDate.getDate();
            if(month<10){
                month='0'+month;
            }
            if(day<10){
                day='0'+day;
            }
            return year + '-' + month + '-' + day;
        };
    })
    .filter('favourable',function () {
        return function (value) {
            var statused = "";
            if(value=== 0){
                statused = "否"
            }else if(value=== 1) {
                statused = "是"
            }
            return statused;
        }
    })
    .filter('fiterStatusBtn',function () {
        return function (value) {
            var statused = "";
            if(value=== 1){
                statused = "禁用"
            }else if(value=== 2) {
                statused = "启用"
            }
            return statused;
        }
    })
    .filter('logoType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="类型一";
            }else if(value===2){
                status="类型二";
            }else if(value===3){
                status="类型三";
            }else if(value===4){
                status="类型四";
            }
            return status;
        }
    })
    .filter('floatType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="左浮动";
            }else if(value===2){
                status="右浮动";
            }
            return status;
        }
    })
    .filter('imgStatus',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="禁用";
            }else if(value===2){
                status="启用";
            }
            return status;
        }
    })
    .filter('noticeWay',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="方式一";
            }else if(value===2){
                status="方式二";
            }else if(value===3){
                status="方式三";
            }else if(value===4){
                status="方式四";
            }
            return status;
        }
    })
    .filter('noticeType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="类型一";
            }else if(value===2){
                status="类型二";
            }else if(value===3){
                status="类型三";
            }else if(value===4){
                status="类型四";
            }
            return status;
        }
    })
    //公司入款方式
    .filter('incomeWay',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="公司入款";
            }else if(value===2){
                status="线上入款";
            }else if(value===3){
                status="人工入款";
            }
            return status;
        }
    })
    //注册文案模板类型
    .filter('registerType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="会员注册";
            }else if(value===2){
                status="代理注册";
            }else if(value===3){
                status="开户协议";
            }
            return status;
        }
    })
    .filter('applicationStatus',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="待审核";
            }else if(value===2){
                status="审核通过";
            }else if(value===3){
                status="审核不通过";
            }
            return status;
        }
    })
    .filter('activityType',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="类型一";
            }else if(value===2){
                status="类型二";
            }else if(value===3){
                status="类型三";
            }else if(value===4){
                status="类型四";
            }
            return status;
        }
    })
    .filter('isOpen',function(){

        return function(value) {
            var status="";
            if(value===1){
                status="开启";
            }else if(value===2){
                status="关闭";
            }
            return status;
        }
    });








