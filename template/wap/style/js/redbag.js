var ishttps = 'https:' == document.location.protocol ? true: false;
if (ishttps) {
    var ptoto = 'https://';
}else{
    var ptoto = 'http://';
}
var red_site_domain=ptoto+window.location.host;

function red_bag_html_(red_site_domain,pic){
    if(pic == '1'){
        pic = '';
    }
    // red_bag_html='<link rel="stylesheet" href="'+cdnUrl+'/wap/style/css/red_wap.css">';
    // red_bag_html+="<div id=\"wrapper_redbag\">";
    // red_bag_html+='    <a class="red-page">';
    // red_bag_html+='        <img src="'+cdnUrl+'/wap/style/images/red.png"/ style="width:75px;z-index:40">';
    // red_bag_html+='    </a>';
    // red_bag_html+='    <div class="red-page-box red-transparent">';
    // red_bag_html+='        <a class="red-closeBtn">&times;</a>';
    // red_bag_html+='        <div class="modal-box red-transparent red-rotate">';
    // red_bag_html+='            <div class="modal-head">';
    // red_bag_html+='                <img src="'+cdnUrl+'/wap/style/images/redHead.png"/>';
    // red_bag_html+='            </div>';
    // red_bag_html+='            <div class="modal-body">';
    // red_bag_html+='                <h3>送现金红包</h3>';
    // red_bag_html+='                <p id="djs_redbag">距离活动结束还剩:19时03分46秒</p>';
    // red_bag_html+='                <a class="submit-btn" id="red_sp">打开查看</a>';
    // red_bag_html+='                <p class="red_ping hide tix">派奖中...请稍后</p>';
    // red_bag_html+="                <p id=\"qd\" class=\"red_p hide\" style=\"border:2px dashed #fff;border-radius:7px;padding:30px 10px;width:86%;margin:20px auto;color:#fff;font-falimy:microsoft yahei;line-height:25px;font-weight:bold;\"></p>";
    // red_bag_html+='            </div>';
    // red_bag_html+='        </div>';
    // red_bag_html+='    </div>';
    // red_bag_html+="</div>";
    // red_bag_html+='<script type="text/javascript" src="'+cdnUrl+'/wap/style/js/jquery-ui.min.js"></script>';
    // red_bag_html+='<script type="text/javascript" src="'+cdnUrl+'/wap/style/js/jquery.ui.touch-punch.min.js"></script>';



    red_bag_html='<link rel="stylesheet" href="'+cdnUrl+'/wap/style/css/red_wap.css">';
    red_bag_html+="<div id=\"wrapper_redbag\">";
    red_bag_html+='    <a class="red-page">';
    red_bag_html+='        <img src="'+cdnUrl+'/wap/style/images/redFront.png"/ style="width:75px;z-index:40">';
    red_bag_html+='    </a>';
    red_bag_html+='    <div class="red-page-box red-transparent">';
    red_bag_html+='        <div class="modal-box red-transparent red-rotate" style="background:url('+cdnUrl+'/wap/style/images/redBg.png) no-repeat;background-size: 100% 100%">';
    red_bag_html+='       <a class="red-closeBtn">&times;</a>';
    red_bag_html+='             <div class="modal-head">';
    red_bag_html+='                 <p id="djs_redbag_t">距离活动结束还剩</p>';
    red_bag_html+='                 <p id="djs_redbag_b">19时03分46秒</p>';
    red_bag_html+='             </div>';
    red_bag_html+='              <div class="modal-body">';
    red_bag_html+='                <button class="submit-btn" id="red_sp" style="background: url('+ cdnUrl+'/wap/style/images/redBtn.png) no-repeat;background-size: 100% 100%"></button>';
    red_bag_html+='                <p class="red_ping hide tix" style="margin-top:10px;">派奖中...请稍后</p>';
    red_bag_html+="                <p id=\"qd\" class=\"red_p hide\" style=\"border:2px dashed #fff;border-radius:7px;padding:30px 10px;width:70%;margin:-58px auto;color:#fff;font-falimy:microsoft yahei;line-height:25px;font-weight:bold;\"></p>";
    red_bag_html+='            </div>';
    red_bag_html+='        </div>';
    red_bag_html+='    </div>';
    red_bag_html+="</div>";
    red_bag_html+='<script type="text/javascript" src="'+cdnUrl+'/wap/style/js/jquery-ui.min.js"></script>';
    red_bag_html+='<script type="text/javascript" src="'+cdnUrl+'/wap/style/js/jquery.ui.touch-punch.min.js"></script>';
    return red_bag_html;
}
$('body').append("<div class='hdddddddddd'></div>");
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

var imgPanel;
//redbag
(function($) {

    //红包js
    var rbh = "";
    rbh+='    <div class="red-page-box red-transparent">';

    rbh+='        <div class="modal-box red-transparent red-rotate" style="background:url('+cdnUrl+'/wap/style/images/redBg.png) no-repeat;background-size: 100% 100%">';
    rbh+='        <a class="red-closeBtn">&times;</a>';
    rbh+='            <div class="modal-head">';
    // rbh+='                <img src="'+cdnUrl+'/wap/style/images/redHead.png"/>';
    rbh+='                <p id="djs_redbag_t">距离活动结束还剩</p>';
    rbh+='                <p id="djs_redbag_b">19时03分46秒</p>';
    rbh+='            </div>';
    rbh+='            <div class="modal-body">';
    // rbh+='                <h3>送现金红包</h3>';
    rbh+='                <button class="submit-btn" id="red_sp" style="background: url('+ cdnUrl+'/wap/style/images/redBtn.png) no-repeat;background-size: 100% 100%"></button>';
    rbh+='                <p class="red_ping hide tix" style="margin-top: 10px">派奖中...请稍后</p>';
    rbh+="                <p id=\"qd\" class=\"red_p hide\" style=\"border:2px dashed #fff;border-radius:7px;padding:30px 10px;width:70%;margin:-58px auto;color:#fff;font-falimy:microsoft yahei;line-height:25px;font-weight:bold;\"></p>";
    rbh+='            </div>';
    rbh+='        </div>';
    rbh+='    </div>';


    var iii = 0;
    var ckclick;
    imgPanel = $(rbh);
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

                   ckclick = function (){
                    imgPanel.find("#red_sp").addClass("ABCIMG");
                    imgPanel.find("#red_sp").click(function(event){
                        // $("#red_sp").hide();
                        // $('.red_ping').show();
                        $("#red_sp").attr('disabled',true);
                        $("#qd").hide();
                        $("#red_sp").stop().addClass('btnRotate');
                        $(".red_ping").show();
                        if(getCookie('loginBack')){
                            var loginBack = getCookie('loginBack');
                            $.ajax({
                                type: "GET",
                                url: "/m/snatch",
                                headers: {
                                    'Authorization': 'bearer ' + loginBack,
                                    'Content-Type': 'application/json',
                                    'Accept': 'application/json',
                                    'platform': platform
                                },
                                data: {"rid":Number(_this.gameslist[0].id)},
                                dataType: 'json',
                                success: function (data) {
                                        if(!data.code){
                                            $('#qd').show()
                                            $("#qd").html('<span>恭喜抢到</span><span id="redmoney">'+parseFloat(data.data.Money).toFixed(2)+'</span><span>元</span>');
                                        }else if(data.code == '71110'){
                                            $('#qd').show()
                                            $("#qd").html('活动尚未开始!<br>请稍后再来!');
                                        }else if(data.code == '71109'){
                                            $('#qd').show()
                                            $("#qd").html('您已经抢过了!<br>请明天再来!');
                                        }else if(data.code == '91201'){
                                            $('#qd').show()
                                            $("#qd").html('您没有登录，请登录后再抢!');
                                        }else if(data.code == '71102'){
                                            $('#qd').show()
                                            $("#qd").html('试玩用户无法参与!<br>请注册真实账户!');
                                        }else if(data.code == '71105'){
                                            $('#qd').show()
                                            $("#qd").html('Sorry，权限不够!<br>请咨询在线客服!');
                                        }else if(data.code == '71108'){
                                            $('#qd').show()
                                            $("#qd").html('该IP已领取此红包!<br>敬请期待，下一轮!');
                                        }else if(data.code == 23){
                                            $('#qd').show()
                                            $("#qd").html('您未达到存款领取资格!<br>请咨询在线客服!');
                                        }else if(data.code == 24){
                                            $('#qd').show()
                                            $("#qd").html('您未达到有效打码领取资格!<br>请咨询在线客服!');
                                        }else if(data.code == 25){
                                            $('#qd').show()
                                            $("#qd").html('<span>恭喜抢到</span><span id="redmoney">'+parseFloat(data.data.Money).toFixed(2)+'</span><span>元</span><br>金额三到五分钟后到账!');
                                        }else if(data.code == 26){
                                            $('#qd').show()
                                            $("#qd").html('加入红包队列失败!<br>请咨询在线客服!');
                                        }else if(data.code == 27){
                                            $('#qd').show()
                                            $("#qd").html('红包派发中,请耐心等待!');
                                        }else{
                                            $('#qd').show()
                                            $("#qd").html('Sorry，来晚一步!');
                                        }

                                    $("#red_sp").removeClass('btnRotate');
                                        $(".red_sp").show();
                                        $(".red_ping").hide();
                                    $("#red_sp").removeAttr('disabled');

                                }
                            });
                        }else{
                            $("#red_sp").removeClass('btnRotate');
                            $(".red_ping").hide();
                            $("#qd").html('您没有登录，请登录后再抢!').show();
                            $("#red_sp").removeAttr('disabled');
                        }

                    });
                    }
                    // $(".red_a").click(function(){
                    //     //查看记录
                    //     $.get(red_site_domain+"/index.php/games/red/snatch_info", {"rid":_this.gameslist[0].id},
                    //         function(data, status){
                    //         if("success" === status ){
                    //            if(data.code == 0 && data.List.length > 0){
                    //            var  txt = "";
                    //             data.List.forEach(function(e){
                    //                 txt += "<p><span>"+ e.username +": </span><span>"+ parseFloat(e.money).toFixed(2)+"元</span></p>";
                    //             });
                    //             $(".jilu .jl_text").children().html(txt)
                    //            }else{
                    //             $(".jilu .jl_text").children().html("<p><span>暂无记录</span><span></span></p>");
                    //            }
                    //         }
                    //     });
                    //
                    //     $("#ren_box").removeClass("show");
                    //     $(".jilu").addClass("show_2");
                    // });
                }
            }else{
                _this.intDiff = _this.gameslist[0].closecount;
               //_this.dtimer(gameslist[0].closecount,2);
                if(!_this.isopen){
                    //抢红包事件

                   ckclick = function (){
                    imgPanel.find("#red_sp").addClass("ABCIMG");
                    imgPanel.find("#red_sp").click(function(event){
                        // $("#red_sp").hide();
                        // $('.red_ping').show();
                        $("#red_sp").attr('disabled',true);

                        $("#qd").hide();
                        $("#red_sp").stop().addClass('btnRotate');
                        $(".red_ping").show();

                        if(getCookie('loginBack')){
                            var loginBack = getCookie('loginBack');
                            $.ajax({
                                type: "GET",
                                url: "/m/snatch",
                                headers: {
                                    'Authorization': 'bearer ' + loginBack,
                                    'Content-Type': 'application/json',
                                    'Accept': 'application/json',
                                    'platform': platform
                                },
                                data: {"rid":Number(_this.gameslist[0].id)},
                                dataType: 'json',
                                success: function(data,status){
                                        // $(".red_a").show();
                                        // console.log(data)
                                        if(data.code==undefined){
                                            console.log(data,"111111111")
                                          $('#qd').show()
                                          
                                            $("#qd").html('<span>恭喜抢到</span><span id="redmoney">'+parseFloat(data.red.money).toFixed(2)+'</span><span>元</span>');
                                        }else
                                         if(data.code == '71110'){
                                           $('#qd').show()
                                          
                                            $("#qd").html('活动尚未开始!<br>请稍后再来!');
                                        }else if(data.code == '71109'){
                                           $('#qd').show()
                                          
                                            $("#qd").html('您已经抢过了!<br>请明天再来!');
                                        }else if(data.code == '91201'){
                                           $('#qd').show()
                                          
                                            $("#qd").html('您没有登录，请登录后再抢!');
                                        }else if(data.code == '71102'){
                                           $('#qd').show()
                                          
                                            $("#qd").html('试玩用户无法参与!<br>请注册真实账户!');
                                        }else if(data.code == '71105'){
                                           $('#qd').show()
                                          
                                            $("#qd").html('Sorry，权限不够!<br>请咨询在线客服!');
                                        }else if(data.code == '71108'){
                                           $('#qd').show()
                                          
                                            console.log(data.code,"66666")
                                            $("#qd").html('该IP已领取此红包!<br>敬请期待，下一轮!');
                                        }else if(data.code == 23){
                                          
                                           $('#qd').show()
                                            $("#qd").html('您未达到存款领取资格!<br>请咨询在线客服!');
                                        }else if(data.code == 24){
                                          
                                           $('#qd').show()
                                            $("#qd").html('您未达到有效打码领取资格!<br>请咨询在线客服!');
                                        }else if(data.code == 25){
                                         $('#qd').show()
                                            $("#qd").html('<span>恭喜抢到</span><span id="redmoney">'+parseFloat(data.red.money).toFixed(2)+'</span><span>元</span><br>金额三到五分钟后到账!');
                                        }else if(data.code == 26){
                                         $('#qd').show()
                                            $("#qd").html('加入红包队列失败!<br>请咨询在线客服!');
                                        }else if(data.code == 27){
                                         $('#qd').show()
                                            $("#qd").html('红包派发中,请耐心等待!');
                                        }else{
                                         $('#qd').show()
                                            $("#qd").html('Sorry，来晚一步!');
                                        }
                                    $("#red_sp").removeClass('btnRotate');
                                        $(".red_ping").hide();
                                        $(".red_sp").show();
                                    $("#red_sp").removeAttr('disabled');

                                }
                            });
                        }else{
                         // $('#qd').fadeIn(1000);
                            $("#red_sp").removeClass('btnRotate');
                            $(".red_ping").hide();
                            $("#qd").html('您没有登录，请登录后再抢!').show();
                            $("#red_sp").removeAttr('disabled');

                        }   
                        
                    });
                    }
                    // $(".red_a").click(function(){
                    //     //查看记录
                    //     $.get(red_site_domain+"/index.php/games/red/snatch_info", {"rid":_this.gameslist[0].id},
                    //         function(data, status){
                    //         if("success" === status ){
                    //            if(data.code == 0 && data.List.length > 0){
                    //            var  txt = "";
                    //             data.List.forEach(function(e){
                    //                 txt += "<p><span>"+ e.username +": </span><span>"+ parseFloat(e.money).toFixed(2)+"元</span></p>";
                    //             });
                    //             $(".jilu .jl_text").children().html(txt)
                    //            }else{
                    //             $(".jilu .jl_text").children().html("<p><span>暂无记录</span><span></span></p>");
                    //            }
                    //         }
                    //     });
                    //
                    //     $("#ren_box").removeClass("show");
                    //     $(".jilu").addClass("show_2");
                    // });
                }
                _this.isopen = true;
            }

        },
        refresh:function(){
            var _this = this;
                //红包活动
            $.get(red_site_domain+"/m/redLog", {},
                function(data, status){
                if("success" === status ){
                    if(!data.code && data.data.length > 0){
                        _this.gameslist = data.data;
                         _this.initbox();
                        var newpic = _this.gameslist[0].pic;
                        if(iii==0) {
                            var redBagDom = $(red_bag_html_(red_site_domain,newpic));
                            // imgPanel = redBagDom.find(".red-page-box.red-transparent").clone(true);
                            redBagDom.find(".red-page-box.red-transparent").remove();
                            // redBagDom.find("#wrapper_redbag");
                            $('.hdddddddddd').append(redBagDom);
                            ckclick();
                            clickfunction();
                            document.cookie="pic="+newpic;
                            iii++;
                        }
                        var cookiepic = getCookie('pic');
                        if(cookiepic != newpic){
                            var redBagDom = $(red_bag_html_(red_site_domain,newpic));
                            // imgPanel = redBagDom.find(".red-page-box.red-transparent").clone(true);
                            redBagDom.find(".red-page-box.red-transparent").remove();
                            // redBagDom.find("#wrapper_redbag");
                            $('.hdddddddddd').append(redBagDom);
                            ckclick();
                            clickfunction();
                            document.cookie="pic="+newpic;
                        }
                        $("#wrapper_redbag .box").show();

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
               txt =  type==2 ? "离活动开始还差:":"离活动结束还差:";
                $("#djs_redbag_t").html(txt);
                $("#djs_redbag_b").html(day+'天'+hour+'时'+minute+'分'+second+'秒');
                if(type==2){
                        _this.intDiff++;
                    }else{
                        _this.intDiff--;
                    }
            }, 1000);
        },
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
    $('#wrapper_redbag').width(w).height(h);
}

window.onresize = function(){
    getSrceenWH();
}
$(window).resize();

function clickfunction(){
    getSrceenWH();
    var newPanel = imgPanel.clone(true);
    $('.red-page').draggable({containment:'#wrapper_redbag'});
    var sto;
    $("a.red-page").click(function(){
        var isLogin = true;//是否登录
        var userData = null;//用户信息
            clearTimeout(sto);
            if(!newPanel.hasClass("amd")){newPanel.addClass("amd");}
            $("#wrapper_redbag").append(newPanel);
            setTimeout(function(){

                $(".red-transparent").removeClass("red-transparent");
                // $(".modal-box").removeClass("red-rotate");

            },10);
            newPanel.find(".red-closeBtn").click(function(){

            clearTimeout(sto);
            $(".red-page-box").addClass("red-transparent");
            sto = setTimeout(function(){
                $(".red-page-box .modal-box").addClass("red-transparent red-rotate");
                setTimeout(function(){
                    $(".red-page-box.red-transparent").remove();
                    newPanel = imgPanel.clone(true);
                },200);
            },500);
            setTimeout(function(){ $("#red_sp").show(),$("#qd").hide()},500);
        });
    });


}
$(document).ready(function(){
    $.redbag.refresh();
    $.redbag.dtimer();
    // $.redbag.testuse();


})

// var get_redbag = function(){
//     $.redbag.refresh();
// }
// $.timer(get_redbag, 120000);
    