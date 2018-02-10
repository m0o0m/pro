$(function () {
    loginId();
    var loginBack = getCookie('loginBack');
    if(loginBack){
        $.ajax({
            url: '/wap/draw/write',
            data: {},
            type: 'get',
            headers: {
                'Authorization': 'bearer ' + loginBack,
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': platform
            },
            beforeSend: function () {
                Loading();
            },
            success: function (data, info, xhr) {
                var mes = '';
                if(data){
                    if (data.code) {
                        mes= data.msg;
                    }
                    if(data.data==0){
                        location.href='/m/withdraw';
                    }
                }

                if (xhr.status == 204) {
                    mes = '出款申请提交成功，系统正在为您处理！';
                }
                if (mes != '') {
                    $.toast(mes,'text',function () {
                        location.href='/m/withdraw';
                    });
                }


            }, complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }

        })
    }

});
$(document).ready(function(){
    var mescroll = new MeScroll("mescroll", {
            //下拉刷新的所有配置项
            down:{
                use: false, //是否初始化下拉刷新; 默认true
                auto: false, //是否在初始化完毕之后自动执行下拉回调callback; 默认true
                hardwareClass: "mescroll-hardware", //硬件加速样式;解决iOS下拉因隐藏进度条而闪屏的问题,参见mescroll.css
                isBoth: false, //下拉刷新时,如果滑动到列表底部是否可以同时触发上拉加载;默认false,两者不可同时触发;
            }
        });
})