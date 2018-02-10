var ishttps = 'https:' == document.location.protocol ? true: false;
if (ishttps) {
    var ptoto = 'https://';
}else{
    var ptoto = 'http://';
}
var red_site_domain=ptoto+window.location.host;
function red_bag_html_(red_site_domain,pic,skin){
    if(pic == '1'){
        pic = '';
    }

    if(skin){
        pic = '';
    }
    red_bag_html='<link href="'+cdnUrl+'/shared/red'+pic+'/css/red_pc.css?v='+jsVersion+'" rel="stylesheet" type="text/css">';
    red_bag_html+="<div id=\"wrapper_redbag\">";
    red_bag_html+="            <div class=\"box\" style=\"position: absolute; width: 100%; height: 100%; display:none\">";
    red_bag_html+="                <div class=\"demo\">";
    red_bag_html+="                    <div onclick=\"document.getElementById('hdddddddddd').style.display='none';\" style='position:absolute;z-index:900;background-repeat: no-repeat;width: 41px;height: 40px;left: 180px;background-image:url(\""+cdnUrl+"/shared/sitepublic/images/close-btn.png?v="+jsVersion+"\");zoom: 0.5;'></div>";
    red_bag_html+="                    <a href=\"javascript:;\" class=\"flipInX\"><img src=\""+cdnUrl+"/shared/red"+pic+"/images/rt-ad.gif?v="+jsVersion+"\"></a>";
    red_bag_html+="                </div>";
    red_bag_html+="";
    red_bag_html+="                <div id=\"dialogBg\"></div>";
    red_bag_html+="                <div id=\"dialog\" class=\"animated\">";
    if(skin){
        red_bag_html+="                    <div class=\"red_bg\" style='background-image:url("+skin.bgpicurl+");'>";
    }else{
        red_bag_html+="                    <div class=\"red_bg\">";
    }
    red_bag_html+="                        <div class=\"dialogTop\">";
    red_bag_html+="                            <a class=\"claseDialogBtn\"></a>";
    red_bag_html+="                        </div>";

    if(skin){
        red_bag_html+="                    <div id=\"djs_redbag\" class=\"djs\" style='top:"+skin.timecss+"'><strong></strong></div>";
    }else{
        red_bag_html+="                    <div id=\"djs_redbag\" class=\"djs\"><strong></strong></div>";
    }

    if(skin){
        red_bag_html+="                    <div class=\"red_but\" style='top:"+skin.clickcss+"'><span id=\"red_sp\"><img src=\""+skin.clickpic+"\" ></span></div>";
    }else{
        red_bag_html+="                    <div class=\"red_but \"><span id=\"red_sp\"><img src=\""+cdnUrl+"/shared/red"+pic+"/images/q_red.gif?v="+jsVersion+"\" ></span></div>";
    }
    red_bag_html+="                        <div class=\"red_ping hide\"><span id=\"red_sp\">派奖中...请稍后</span></div>";
    red_bag_html+="                        <div class=\"ren_box\" id=\"ren_box\">";
    red_bag_html+="                            <p id=\"qd\" class=\"red_p hide\"><span>恭喜抢到</span><span id=\"redmoney\">0</span><span>元</span></p>";
    red_bag_html+="                            <p id=\"late\" class=\"red_p hide\"><span>Sorry，来晚一步！</span></p>";
    red_bag_html+="                            <p id=\"is_ip\" class=\"red_p hide\"><span>该IP已领取此红包！<br><br>敬请期待，下一轮！</span></p>";
    red_bag_html+="                            <p id=\"is_sum\" class=\"red_p hide\"><span>您未达到存款领取资格！<br><br>详情请咨询客服！</span></p>";
    red_bag_html+="                            <p id=\"is_bet\" class=\"red_p hide\"><span>您未达到有效打码领取资格！<br><br>详情请咨询客服！</span></p>";
    red_bag_html+="                            <p id=\"cengji\" class=\"red_p hide\"><span>Sorry，权限不够！<br><br>请咨询在线客服！</span></p>";
    red_bag_html+="                            <p id=\"qg\" class=\"red_p hide\"><span>您已经抢过了!<br/>请明天再来!</span></p>";
    red_bag_html+="                            <p id=\"ks\" class=\"red_p hide\"><span>活动尚未开始!<br/>请稍后再来!</span></p>";
    red_bag_html+="                            <p id=\"needlogin\" class=\"red_p hide\">您没有登录，<br/>请登录后再抢!</p>";
    red_bag_html+="                            <p id=\"shiwan\" class=\"red_p hide\">试玩用户无法参与!<br/>请注册真实账户!</p>";
    red_bag_html+="                            <p id=\"wlfm1\" class=\"red_p hide\">网络繁忙!<br/>刷新后重新抢红包1!</p>";
    red_bag_html+="                            <p id=\"wlfm2\" class=\"red_p hide\">网络繁忙!<br/>刷新后重新抢红包2!</p>";
    red_bag_html+="                            <p id=\"wlfm3\" class=\"red_p hide\">网络繁忙!<br/>刷新后重新抢红包3!</p>";
    red_bag_html+="                            <p id=\"wlfm4\" class=\"red_p hide\">网络繁忙!<br/>刷新后重新抢红包4!</p>";
    red_bag_html+="                            <p id=\"wlfm5\" class=\"red_p hide\">网络繁忙!<br/>红包领取失败!</p>";
    red_bag_html+="                            <span class=\"red_a\" href=\"javascript:;\">查看记录</span>";
    red_bag_html+="                        </div>";
    red_bag_html+="                        <div class=\"jilu\">";
    red_bag_html+="                            <div class=\"jl_text str1 str_wrap\">";
    red_bag_html+="                                ";
    red_bag_html+="                        </div>";
    red_bag_html+="                    </div>";
    red_bag_html+="                </div>";
    red_bag_html+="</div>";
    red_bag_html+='<script type="text/javascript" src="'+cdnUrl+'/shared/red'+pic+'/js/jquery.liMarquee.js?v='+jsVersion+'"></scritp>';
    return red_bag_html;
}
$('body').append("<div id='hdddddddddd'></div>");
//$('body').append('<link href="'+cdnUrl+'/shared/red/css/red_pc.css" rel="stylesheet" type="text/css">');
(function ($) {
    $.extend({
        timer: function (action,time) {
            var _timer;
            if ($.isFunction(action)) {
                (function () {
                    _timer = setInterval(function () {
                        action();
                    }, time);
                })();
            }
        }
    });
})(jQuery);

//redbag
(function($) {
    //红包js
    var iii = 0;
    $.redbag = {
        gameslist:{},
        isopen:false,
        intDiff:0,
        initbox:function(){
            var _this = this;
            if(_this.gameslist[0].opencount < 0){
                _this.intDiff = _this.gameslist[0].opencount;
                 if(!_this.isopen){
                    //抢红包事件
                    $("#red_sp").click(function(){
                        $(".red_but").hide();
                         $(".red_ping").show();
                         var loginBack = getCookie('loginBack');
                         if(loginBack){
                            var $success = function(data, status){
                                if("success" === status ){
                                    $(".red_a").show();
                                   if(data.code == 0){
                                     $(".red_p").addClass("hide");
                                     $("#qd").removeClass("hide");
                                     $("#qd #redmoney").html(parseFloat(data.red.money).toFixed(2))
                                   }else if(data.code == '71106'){
                                     $("#qd").addClass("hide");
                                     $("#ks").removeClass("hide");//needlogin
                                     $("#ks").html(data.msg);
                                     $(".red_a").hide();
                                   }else if (data.code == '71110') {
                                     $("#qd").addClass("hide");
                                     $("#ks").removeClass("hide");//needlogin
                                     $(".red_a").hide();
                                   }else if(data.code == '71109'){
                                     $("#qd").addClass("hide");
                                     $("#qg").removeClass("hide");//needlogin
                                     $(".red_a").hide();
                                   }else if(data.code == '91201'){
                                     $("#needlogin").removeClass("hide");
                                     $(".red_a").hide();
                                   }else if(data.code == '71102'){
                                    $(".red_p").addClass("hide");
                                    $('#shiwan').removeClass("hide");
                                    $(".red_a").hide();
                                   }else if(data.code == '71105'){
                                    $(".red_p").addClass("hide");
                                    $('#cengji').removeClass("hide");
                                    $(".red_a").hide();
                                   }else if(data.code == '71108'){
                                    $(".red_p").addClass("hide");
                                    $('#is_ip').removeClass("hide");
                                    $(".red_a").hide();
                                   // }else if(data.code == 23){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#is_sum').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 24){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#is_bet').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 94){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#wlfm1').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 95){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#wlfm2').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 96){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#wlfm3').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 97){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#wlfm4').removeClass("hide");
                                   //  $(".red_a").hide();
                                   // }else if(data.code == 98){
                                   //  $(".red_p").addClass("hide");
                                   //  $('#wlfm5').removeClass("hide");
                                   //  $(".red_a").hide();
                                   }else{
                                    $("#late").removeClass("hide");
                                    $(".red_a").hide();
                                   }
                                   if(data.code != 0){
                                    $("#qd").addClass("hide");
                                   }
                                    $(".red_ping").hide();
                                    $(".red_but").show();
                                }
                                $("#ren_box").addClass("show");
                            }
                            var data = {"rid":_this.gameslist[0].id}
                            ajaxObj('GET', "/snatch", $success, data);


                         }else{
                            $("#needlogin").removeClass("hide");
                            $(".red_a").hide();
                            $("#ren_box").addClass("show");
                         }
                        
                    });
                    $(".red_a").click(function(){
                        //查看记录
                        $.get(red_site_domain+"/index.php/games/red/snatch_info", {"rid":_this.gameslist[0].id},
                            function(data, status){
                            if("success" === status ){
                               if(data.Code == 0 && data.List.length > 0){
                               var  txt = "";
                                data.List.forEach(function(e){
                                    txt += "<p><span>"+ e.username +": </span><span>"+ parseFloat(e.money).toFixed(2)+"元</span></p>";
                                });
                                $(".jilu .jl_text").children().html(txt)
                               }else{
                                $(".jilu .jl_text").children().html("<p><span>暂无记录</span><span></span></p>");
                               }
                            }
                        });

                        $("#ren_box").removeClass("show");
                        $(".jilu").addClass("show_2");
                    });
                }
                _this.isopen = true;
                //_this.dtimer(gameslist[0].opencount,1);
            }else{
                _this.intDiff = _this.gameslist[0].closecount;
               //_this.dtimer(gameslist[0].closecount,2);
                if(!_this.isopen){
                    //抢红包事件
                    $("#red_sp").click(function(){
                        $(".red_but").hide();
                        $(".red_ping").show();
                        var loginBack = getCookie('loginBack');
                        if(loginBack){
                            var $success = function(data, status){
                                    if("success" === status ){
                                        $(".red_a").show();
                                        if(!data.code){
                                            $(".red_p").addClass("hide");
                                            $("#qd").removeClass("hide");
                                            $("#qd #redmoney").html(parseFloat(data.data.money).toFixed(2))
                                        }else if(data.code == '71106'){
                                            $("#qd").addClass("hide");
                                            $("#ks").removeClass("hide");//needlogin
                                            $("#ks").html(data.msg);
                                            $(".red_a").hide();
                                        }else if (data.code == '71110') {
                                            $("#qd").addClass("hide");
                                            $("#ks").removeClass("hide");//needlogin
                                            $("#ks").html(data.msg);
                                            $(".red_a").hide();
                                        }else if(data.code == '71109'){
                                            $("#qd").addClass("hide");
                                            $("#qg").removeClass("hide");//needlogin
                                            $(".red_a").hide();
                                        }else if(data.code == '91201'){
                                            $("#needlogin").removeClass("hide");
                                            $(".red_a").hide();
                                        }else if(data.code == '71102'){
                                            $(".red_p").addClass("hide");
                                            $('#shiwan').removeClass("hide");
                                            $(".red_a").hide();
                                        }else if(data.code == '71105'){
                                            $(".red_p").addClass("hide");
                                            $('#cengji').removeClass("hide");
                                            $(".red_a").hide();
                                        }else if(data.code == '71108'){
                                            $(".red_p").addClass("hide");
                                            $('#is_ip').removeClass("hide");
                                            $(".red_a").hide();
                                        }else{
                                            $("#late").removeClass("hide");
                                            $(".red_a").hide();
                                        }
                                        if(data.code != 0){
                                            $("#qd").addClass("hide");
                                        }
                                        $(".red_ping").hide();
                                        $(".red_but").show();
                                    }
                                    $("#ren_box").addClass("show");
                                }
                            var data = {"rid":_this.gameslist[0].id}
                            ajaxObj('GET', "/snatch", $success, data);
                        }else{
                            $("#needlogin").removeClass("hide");
                            $(".red_a").hide();
                            $("#ren_box").addClass("show");
                        }
                    });
                    $(".red_a").click(function(){
                        //查看记录
                        $.get(red_site_domain+"/index.php/games/red/snatch_info", {"rid":_this.gameslist[0].id},
                            function(data, status){
                            if("success" === status ){
                                if(data.Code == 0 && data.List.length > 0){
                                    var  txt = "";
                                    data.List.forEach(function(e){
                                        txt += "<p><span>"+ e.username +": </span><span>"+ parseFloat(e.money).toFixed(2)+"元</span></p>";
                                    });
                                    $(".jilu .jl_text").children().html(txt)
                                }else{
                                    $(".jilu .jl_text").children().html("<p><span>暂无记录</span><span></span></p>");
                                }
                            }
                        });

                        $("#ren_box").removeClass("show");
                        $(".jilu").addClass("show_2");
                    });
                }
                _this.isopen = true;
            }

        },
        refresh:function(){
            var _this = this;
                //红包活动
            $.get(red_site_domain+"/red/log", {},
                function(data, status){
                if("success" === status ){
                    if(!data.Code && data.data.length > 0){
                        _this.gameslist = data.data;
                        var newpic = _this.gameslist[0].pic;
                        if(iii==0) {
                            $('#hdddddddddd').html(red_bag_html_(red_site_domain,newpic,_this.gameslist[0].skin));
                            clickfunction();
                            document.cookie="pic="+newpic;
                            iii++;
                        }
                        var cookiepic = getCookie('pic');
                        if(cookiepic != newpic){
                            $('#hdddddddddd').html(red_bag_html_(red_site_domain,newpic,_this.gameslist[0].skin));
                            clickfunction();
                            document.cookie="pic="+newpic;
                        }
                        $("#wrapper_redbag .box").show();
                        _this.initbox();
                    }else{
                        $("#wrapper_redbag .box").hide();
                    }
                }
            });
        },
        dtimer:function(){
            var _this = this;
            $.timer(function(){
                var day=0,  hour=0,  minute=0, second=0,type=1;//时间默认值
                if(_this.gameslist.length >0){
                    if(_this.gameslist[0].opencount > 0){
                        type = 1;
                    }else{
                        type = 2;
                    }
                }else{
                    return;
                }
               if(_this.intDiff > 0){
                    day = Math.floor(_this.intDiff / (60 * 60 * 24));
                    hour = Math.floor(_this.intDiff / (60 * 60)) - (day * 24);
                    minute = Math.floor(_this.intDiff / 60) - (day * 24 * 60) - (hour * 60);
                    second = Math.floor(_this.intDiff) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
                }else{
                    var shijian = Math.abs(_this.intDiff)
                    day = Math.floor(shijian / (60 * 60 * 24));
                    hour = Math.floor(shijian / (60 * 60)) - (day * 24);
                    minute = Math.floor(shijian / 60) - (day * 24 * 60) - (hour * 60);
                    second = Math.floor(shijian) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
                }
                if (minute <= 9) minute = '0' + minute;
                if (second <= 9) second = '0' + second;
                txt =  type==2 ? "离活动开始还差:":"离结束还差:";
                $("#djs_redbag").children().html(txt+day+'天'+hour+'时'+minute+'分'+second+'秒');

                  if(type==2){
                        _this.intDiff++;
                    }else{
                        _this.intDiff--;
                    }
            }, 1000);
        }
    }
})(jQuery);

function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
    if(arr=document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

var w,h,className;
function getSrceenWH(){
    w = $(window).width();
    h = $(window).height();
    $('#dialogBg').width(w).height(h);
}

window.onresize = function(){
    getSrceenWH();
}
$(window).resize();

function clickfunction(){
    getSrceenWH();
    //显示弹框
    $('.box a').click(function(){
        className = $(this).attr('class');
        $('#dialogBg').fadeIn(300);
        $('#dialog').removeAttr('class').addClass('animated '+className+'').fadeIn();
    });

    //关闭弹窗
    $('.claseDialogBtn').click(function(){
        $('#dialogBg').fadeOut(300,function(){
            $('#dialog').addClass('bounceOutUp').fadeOut();
        });
    });

    $(".claseDialogBtn").click(function(){
        $("#ren_box").removeClass("show");
        $(".jilu").removeClass("show_2");
    });
}
$(document).ready(function(){
    $.redbag.refresh();
    $.redbag.dtimer();
})
 /* var get_redbag = function(){
        $.redbag.refresh();
}
$.timer(get_redbag, 60000); */