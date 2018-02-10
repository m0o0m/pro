/**
 * Created by Administrator on 2017/1/3.common_forpage
 */

/***********************************js跳转页面******************************/
//open
function getPager(type, mo, me) {
    if (type.charAt(0) == "-") {
        if (type.charAt(1) != '') {
            if (me == '') {
                alert("待添加路径");
            } else {
            }
        } else {
            if (mo == "iword") {
                window.location.href = "/"+mo + '?id=' + me;
            } else {
                if (mo == "video") {
                    window.location.href = "/livetop";
                } else if (mo == "fc") {
                    window.location.href = "/lottery";
                } else if (mo == "dz") {
                    if(me == 'm'){
                        window.location.href = "/egame";
                    }else{
                        var mes = [];
                        mes = me.split("_");
                        window.location.href = "/egame?type="+mes[0].toUpperCase();
                    }
                } else if (mo == "sp") {
                    window.location.href = "/sports"
                } else {
                    window.location.href = "/"+mo;
                }
            }
        }
    } else if (type.charAt('_')) {
        alert("仍待添加路径");
    }
    // if(mo == 'shiwan_reg'){
    //     joinDemoDo();
    //     return false;
    // }
    // if(!arguments[2] || me == 'm' || mo == 'dailishenqing') me = '';
    // if (mo == 'lottery' && me == '')
    //     me = 'liuhecai';
    // if (mo == 'pk_lottery' && me == '')
    //     me = 'liuhecai';
    // var siteUrl = '';
    // var pathName = window.location.pathname;
    // if (pathName.indexOf('viewcache') > 0) {
    //     if (me == '')
    //         siteUrl = hex_md5(mo) + '.html?v=' + parent.Version;
    //     else
    //         siteUrl = hex_md5(mo + me) + '.html?v=' + parent.Version;
    // }else{
    //     siteUrl = mo;
    // }
    //
    // if(type.charAt(0) == "-") {
    //     if (type.charAt(1) != '') {
    //         var adid = type.substr(1);
    //         if (me == '') {
    //             window.location.href = siteUrl + '?advtype='+adid;
    //         }else{
    //             window.location.href = siteUrl + '?metype='+me + '&advtype='+adid;
    //         }
    //     }else{
    //         if (siteUrl.indexOf('.html') > 0 || me == '') {
    //             window.location.href = siteUrl;
    //         }else{
    //             window.location.href = siteUrl + '?metype='+me;
    //         }
    //     }
    // }else if(type.charAt('_')){
    //     if(type == '_bank') {
    //         mo = 'http://'+mo;
    //         window.open(mo, null);
    //     } else if(type == '_self') {
    //         location.href = mo;
    //     }
    // }
}

/***********************************END js跳转页面******************************/

/***********************************会员中心******************************/
function openmember(id) {
    // window.location.href = '/index.php/index/new_member_main?url=' + id;
    switch (Number(id)){
        case 1: //存款
            window.location.href = '/member/bank';
            break;
        case 2: //取款
            window.location.href = '/member/withdraw';
            break;
        case 3: //额度转换
            window.location.href = '/member/convert';
            break;
        case 4: //我的账户
            window.location.href = '/member/account';
            break;
        case 5: //交易记录
            window.location.href = '/member/record';
            break;
        case 6: //报表统计
            window.location.href = '/member/report';
            break;
        case 7: //我要推广
            window.location.href = '/member/spread';
            break;
        case 8: //消息中心
            window.location.href = '/member/mescenter';
            break;
    }
}

//前台会员中心打开
function openHelp(url) {
    id = url.split("=");//兼容老版本
    openmember(id[1]);
}

function open_new_member(mo) {
    window.location.href = '/index.php/index/new_member_main?url=' + mo;
}

/***********************************END会员中心******************************/

/***********************************回到顶部按钮******************************/
if ('undefined' != typeof($)) {
    $(function () {
        var btnNum = $('#ele-float-top').children().length,
            wrap = $('#ele-float-top-wrap'),
            wrapHeight = (btnNum - 1) * (40 + 2),
            gotop = $('#ele-float-top-up'),
            speedSet = 300,
            thebox = $('.ele-float-box-wrap'),
            boxwrap = '';

        wrap.height(wrapHeight);
        if (wrap.height() == wrapHeight) {
            $('#ele-float-top').show();
        }

        $('.ele-float-top-code').hover(function () {
            $(this).children(thebox).stop(true, true).fadeIn(speedSet);
        }, function () {
            $(this).children(thebox).stop(true, true).fadeOut(speedSet);
        });

        $("#ele-float-top-up").click(function () {
            $('html,body').animate({scrollTop: 0}, 1000, 'easeOutExpo');
        });
        $(window).scroll(function () {
            if (navigator.userAgent.indexOf("MSIE") != -1) {
                var fadeSec = 200;
            } else {
                var fadeSec = 300;
            }
            if ($(this).scrollTop() > 300) {
                $('#ele-float-top-up').fadeIn(fadeSec);
            } else {
                $('#ele-float-top-up').stop().fadeOut(fadeSec);
            }

        });
    });
}
/***********************************END回到顶部按钮******************************/

//文案页面跳转
function iworkJump(str){
    $("div.bottom a.js-article-color").each(function(e){
        var $title = $(this).text().replace(/[^\u4e00-\u9fa5]/gi,"")
        if ( $title == str ) {
            $(this).click();
        }
    });
}