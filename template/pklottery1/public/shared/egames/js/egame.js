$.EGame = {
    Article : {
        five : 25,
        six : 24
    },
    eData : {},
    gameData : {},
    EgDemoUrl : "/index.php/games/Logingame/EgDemo",
    PtDemoUrl : "http://cache.download.banner.longsnake88.com/flash/79/launchcasino.html?mode=offline&affiliates=1&language=ZH-CN&game=",
    BBinUrl : "/login?g_type=bbin",
    AllGameUrl : "/index.php/games/Logingame/index",
    _openGame : function (url) {
        window.open(url, '_blank', 'width=1000,height=800,top=0,left=0,status=no,toolbar=no,scrollbars=yes,resizable=no,personalbar=no');
    },
    //进入游戏
    inGame : function (code, g_type,sw){
        var loginIn = getCookie('loginBack');
        if(loginIn){
            
            var userid = loginData['id'];
        }else{
            var userid = "";
        }
        
        var _this = $.EGame;
        if(!arguments[2]) sw = 0;
        // var param = "?gameid=" + code + "&g_type=" + g_type;
        // opengeme(code,'',g_type);
        // return;

        if (g_type === 'eg' && (userid == "" || sw == 1)) {
            // _this._openGame(_this.EgDemoUrl + param);
            opengeme(code,'',g_type);
            return false;
        }
        
        if (g_type === 'pt' && sw == '1') {
            // _this._openGame(_this.PtDemoUrl + code);
            opengeme(code,'',g_type);
            return false;
        }
        
        if (userid == "") {
            zhuModal.login();
        }else{
            opengeme(code,'',g_type);
            // if (g_type === 'bbin') {
            //     _this._openGame(_this.BBinUrl);
            // } else {
            //     _this._openGame(_this.AllGameUrl + param + "&sw=" + sw);
            // }
        }
    },
    //點擊后選中樣式
    setTab : function (e) {
        $('.off').removeClass('off');
        $('#one' + e).addClass('off');
    },
    /*********************** get game data **************************/
    getData : function (gtype) {
        $('#tab1').css('position','relative');
        if(gtype == 'GG') {$('#ul_1,#con_one_1 div.search').hide();}
        $.ajax({
            type : 'GET',
            url : '/ajax/egame',
            data : 'type='+gtype,
            dataType : 'json',
            beforeSend : function(){
                $('div#tab1').prepend('<div id="xxoo"><img src="'+ cdnUrl +'/shared/egames/images/ajax-loader-white.gif?v='+jsVersion+'" id="xxoo1" width="150" height="150"/></div>');
                $('#xxoo').css({
                    padding:        0,
                    margin:         0,
                    width:          '100%',
                    height:         '100%',
                    top:            '0',
                    left:           '0',
                    textAlign:      'center',
                    color:          '#000',
                    border:         'none',
                    "position":     "absolute",
                    "z-index":      1000,
                    "opacity":      0.7,
                    "background-color": "#000000"
                });
                $('#xxoo1').css({'margin-top': '23%'});
            },
            success: function(msg){
                $.EGame.eData = msg['data']['data'];
                if (msg['wh'] == '1') { //维护
                    $.EGame.GetWhHtml(msg['wh'].content);
                    return false;
                }else if(gtype != 'GG'){$('#ul_1,#con_one_1 div.search').show();}

                var top = '<li><a href="javascript:;" class="active" style="width: 94px;" onclick="$.EGame.GetTopGame(0);">所有游戏</a></li>'+
                    '<li><a href="javascript:;" style="width: 94px;" onclick="$.EGame.GetTopGame(1);" class="">拉霸</a></li>'+
                    '<li><a href="javascript:;" style="width: 94px;" onclick="$.EGame.GetTopGame(2);" class="">桌面游戏</a></li>'+
                    '<li><a href="javascript:;" style="width: 94px;" onclick="$.EGame.GetTopGame(3);" class="">视频扑克</a></li>'+
                    '<li><a href="javascript:;" style="width: 94px;" onclick="$.EGame.GetTopGame(4);" class="">其它游戏</a></li>';
                $('#ul_1').html(top);
                $('#etype').val(msg['data']['type']);
                $('#xxoo').remove();
                $('#tab1').css('position','static');
                $.EGame.GetGamePage(1);
            }
        });
    },
    //分頁
    GetGamePage : function (page,egameD){
        var _this = $.EGame;
        var type = $('#etype').val();
        if(!arguments[1]) _this.gameData = _this.eData;
        else _this.gameData = egameD;
        var games = "";
        var totalNum = Number(_this.gameData.length);          //总条数
        if (totalNum < 1) {
            games = "<h2 style='margin-left:20px;font-size:18px;'>此分类暂无数据</h2><br/><h3 style='margin-left:20px;font-size:18px;'>敬请期待！</h3>";
            $('.games_menu').html(games);
            return false;
        }
        var game = new Array();
        var PageHtml = "";
        var PageSize = (type == 'EG' || type == 'MG') ? _this.Article.five : _this.Article.six; //每页显示条数
        var offset = (page-1)*PageSize;                //偏移量
        var EndNum = 0;                                //key值
        var totalPage = Math.ceil(totalNum/PageSize);  //总页数
        if (page < 1 || page > totalPage) {alert('页数错误！');return false;}
        if(totalNum < PageSize || page == totalPage) EndNum = totalNum;
        else if(page < totalPage) EndNum = offset+PageSize;
        game = _this.gameData.slice(offset, EndNum);
        games = eval('_this.Get'+type+'html(game)');
        for (var i = 1; i <= totalPage; i++) {
            if (page == i) {
                PageHtml += '<li class="Dz-btn" style="background:#a42919">'+i+'</li>';
            }else{
                PageHtml += '<li class="Dz-btn" onclick="eval($.EGame.GetGamePage('+i+',$.EGame.gameData))">'+i+'</li>';
            }
        }
        $('.games_menu').html(games);
        if(type == 'GG'){
            $('#page_navigation>.btndiv').hide();
        }else{
            $('#page_navigation>.btndiv').show();
            $('#page_navigation>.btndiv').html(PageHtml);
        }
        _this.HoverEvent();
    },

    // 遊戲分類  如：拉霸
    GetTopGame : function (top) {
        var _this = $.EGame;
        $('.active').removeClass('active');
        $('#ul_1>li').eq(top).find("a").addClass("active");
        var game = new Array();
        if (top == 0) {
            _this.GetGamePage(1);
            return false;
        }
        $.each(_this.eData, function (index, item) {
            if (item.topid == top) {
                game.push(_this.eData[index]);
            }
        });
        _this.GetGamePage(1,game);
    },

    //遊戲搜索
    search : function (keywords) {
        var _this = $.EGame;
        var game = new Array();
        $.each(_this.eData, function (index, content) {
            var key = content.name;
            if (key.indexOf(keywords) + 1 > 0) {
                game.push(_this.eData[index]);
            }
        });
        if (game.length == 0) {
            $('.games_menu').html("<h3 style='margin-left:20px;font-size:18px;'>没有搜索到相关游戏</h3>");
        }else _this.GetGamePage(1,game);
    },

    //绑定 hover 事件
    HoverEvent : function () {
        $(".games_menu div.games_bravado_container").each(function(i){
            $(this).mouseover(function(){
                $(this).find("div.game_button_play").show();
                $(this).find("div.game_button_try").show();
            }).mouseout(function(){
                $(this).find("div.game_button_play").hide();
                $(this).find("div.game_button_try").hide();
            });
        });

        $(".video-con-bg").mousemove(function(){
            $(this).find('img').css("left" , "-145px");
            $(this).find(".video-btn").css("top" , "-50px");
        });
        $(".video-con-bg").mouseout(function(){
            $(this).find('img').css("left" , "0");
            $(this).find(".video-btn").css("top" , "0");
        });
    },

    //获取电子标题
    egameTitle : function (Ttype){
        var eTitle = '';
        if (Ttype == "AG" || Ttype == "EG") {
            eTitle = Ttype + '电子/捕鱼';
        }else if (Ttype == "GG") {
            eTitle = Ttype + '捕鱼';
        }else{
            eTitle = Ttype + '电子';
        }
        $("#etitle-"+Ttype).html(eTitle);
    },
    GetWhHtml : function (whcontent) { //维护
        var games = '';
        games += "<h1 style='margin-left:20px;font-size:64px;'>此游戏正在维护，请稍后访问······</h1><br/>";
        games += "<h1 style='margin-left:20px;font-size:64px;'>"+whcontent+"</h1>";
        $('#ul_1,#con_one_1 div.search').hide();
        $('#xxoo').remove();
        $('#page_navigation>.btndiv').html('');
        $('.games_menu').html(games);
    },
    GetEGhtml : function (e) {
        var games = '<div style="margin-bottom:20px;"><a href="javascript:$.EGame.inGame(\'10000\', \'eg\');"><img src="'+ cdnUrl +'/shared/egames/images/eg/egby.gif?v='+jsVersion+'" width="1000" /></a></div><ul id="ajax-content" style="display: block;">';
        $.each(e,function(i, v) {
            games += '<div class="video-bg"><div class="video-tit-bg"><span class="video-tit-col">'+v.name+'</span></div><div class="video-con-bg"><div class="video-logo"><a href="javascript:;"><img onerror="nofind();" src="'+ cdnUrl +'/shared/egames/images/pk/'+v.image +'?v='+jsVersion+'" style="left: 0px;"></a></div><div class="video-btn" style="top: 0px;"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'eg\', \'1\');" class="video-sw">免费试玩</a><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'eg\');" class="video-go">开始游戏</a></div></div></div>';
        });
        return games;
    },
    GetMGhtml : function (e) {
        var games = '<ul id="gamelist">';
        $.each(e,function(i, v) {
            games += '<li><div class="game_text" id="'+ v.gameid +'">'+ v.name +'</div><div class="game_logo"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'mg\');"><img onerror="nofind();" src="'+ cdnUrl +'/shared/egames/images/mg/'+ v.image +'?v='+jsVersion+'"></a></div></li>';
        });
        games += '</ul>';
        return games;
    },
    GetAGhtml : function (e) {
        var games = '<div style="margin-bottom:20px;"><a href="javascript:$.EGame.inGame(\'6\', \'ag\');"><img src="'+ cdnUrl +'/shared/egames/images/ag/byw_2.gif?v='+jsVersion+'" width="1000" /></a></div><ul id="ajax-content" style="display: block;">';
        $.each(e,function(i, v) {
            games += '<li class="game_item" style="display: list-item;"><div class="game_title"><div class="game_title_text">'+v.name+'</div><span class="game_star"><a class="star_favor" href="javascript:;"></a></span><div class="clear"></div></div><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'' + v.type + '\');"><img src="'+ cdnUrl +'/shared/egames/images/ag/'+v.image+'?v='+jsVersion+'" onerror="nofind();"></a><a class="enter-game" href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'ag\');" title="进入游戏"></a><div class="clear"></div></li>';
        });
        games += '</ul>';
        return games;
    },
    GetBBINhtml : function (e) {
        var games = '<div style="padding-left:10px">';
        $.each(e,function(i, v) {
            games += '<div class="game_bbin"><div class="bbin_bg"><div class="bbin_img"><a class="img_bg"><img src="'+ cdnUrl +'/shared/egames/images/bbin/'+v.image+'?v='+jsVersion+'" onerror="nofind();"/></a><div class="bbin_tit"><h3>'+ v.name +' </h3></div></div><div class="bbin_hide"><div class="bbin-game-ctl-links"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'bbin\');" class="bbin_jinru">进入游戏</a><a href="javascript:window.open(\''+ cdnUrl +'/shared/download/download.html?v='+jsVersion+'\',\'\',\'width=1040,height=706,fullscreen=1,scrollbars=0,location=no\');" class="bbin_shuom">下载专区</a></div></div></div></div>';
        });
        games += '</div>';
        return games;
    },
    GetPThtml : function (e) {
        var games = '';
        $.each(e,function(i, v) {
            games += '<div class="games_bravado_container"><div class="games"><div class="image">';
            games += '<img src="'+ cdnUrl +'/shared/egames/images/pt/'+v.image+'?v='+jsVersion+'" onerror="nofind();">';          
            games += '</div><div class="name"><div class="opacity_content"><div class="opacity_background"></div><div class="opacity_content"><div>'+ v.name +'</div></div></div></div><div class="game_button_play" ><div class="game_button_play_bg"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'pt\');">立即游戏 </a></div></div><div class="game_button_try"><div class="game_button_try_bg"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'pt\', \'1\');">立即试玩 </a></div><div class="game_button_play_bg"><a href="http://cdn.fruitfarm88.com/generic/d/setupglx.exe" target=\'_blank\'>客户端下载 </a></div></div></div></div>';
        });
        return games;
    },
    GetGPIhtml : function (e) {
        var games = '';
        $.each(e,function(i, v) {
            games += '<div class="games_bravado_container"><div class="gamesgpi" id="' + v.gameid + '"><div class="image"><img src="'+ cdnUrl +'/shared/egames/images/gpi/'+ v.image +'?v='+jsVersion+'" onerror="nofind();"></div><div class="name"><div class="opacity_content"><div class="opacity_background"></div><div class="opacity_content"><div>'+ v.name +'</div></div></div></div><div class="game_button_play" ><div class="game_button_play_bg"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'gpi\');">立即游戏 </a></div></div></div></div>';
        });
        return games;
    },
    GetGDhtml : function (e) {
        var games = '';
        $.each(e,function(i, v) {
            games += '<div class="games_bravado_container"><div class="gamesgpi" id="' + v.gameid + '"><div class="image"><img src="'+ cdnUrl +'/shared/egames/images/gd/'+ v.image +'?v='+jsVersion+'" onerror="nofind();"></div><div class="name"><div class="opacity_content"><div class="opacity_background"></div><div class="opacity_content"><div>'+ v.name +'</div></div></div></div><div class="game_button_play" ><div class="game_button_play_bg"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'gd\');">立即游戏 </a></div></div></div></div>';
        });
        return games;
    },
    GetHBhtml : function (e) {
        var games = '';
        $.each(e,function(i, v) {
            games += '<div class="games_bravado_container"><div class="gamesgpi" id="' + v.gameid + '"><div class="image">';
            if (v.image != '') {
                games += '<img src="'+ v.image +'?v='+jsVersion+'" onerror="nofind();">';
            }else{
                games += '<img src="'+ cdnUrl +'/shared/egames/images/pt/PT_bal.jpg?v='+jsVersion+'">';
            }
            games += '</div><div class="name"><div class="opacity_content"><div class="opacity_background"></div><div class="opacity_content"><div>'+ v.name +'</div></div></div></div><div class="game_button_play" ><div class="game_button_play_bg"><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'hb\');">立即游戏 </a></div></div></div></div>';
        });
        return games;
    },
    GetGGhtml : function (e) {
        if(loginData){
            userid = loginData['id'];
        }else{
            userid = '';
        }
        var games = '';
        var ggimg = new Array();
        ggimg = e[0].image.split(",");
        if(userid){
            var image = ggimg[1];
        }else{
            var image = ggimg[0];
        }
        games = '<div><div class="gamesgg" id="' +  e[0].gameid + '"><div ><a href="javascript:$.EGame.inGame(\'' +  e[0].gameid + '\', \'gg\');"><img onerror="nofind();" src="'+ cdnUrl +'/shared/egames/images/gg/'+ image +'?v='+jsVersion+'"></a></div><div class="name"><div class="opacity_content"><div class="opacity_background"></div><div class="opacity_content"></div></div></div><div class="game_button_play" ></div><div class="game_button_try"></div></div></div>';
        return games;
    }
};

//標題自動寬度
$(function(){
    var liwidth = 1000 / $('.divgmenu>.ul_ul>li').length;
    $('.divgmenu>.ul_ul>li').width(liwidth);

    //輪播
    var $index = 0;
    var $exdex = 0;
    $(".egamechoose span").mouseover(function() {
        $index = $(this).index();
        $(".egamechoose span").eq($index).addClass("egamered").siblings().removeClass("egamered");
        if ($index > $exdex) {
            next();
        } else if ($index < $exdex) {
            pre();
        }
        return $exdex = $index;
    });
    $(".egamenext").click(function() {
        $index++;
        if ($index > 4) {
            $index = 0
        }
        $(".egamechoose span").eq($index).addClass("egamered").siblings().removeClass("egamered");
        next();
        return $exdex = $index;
    });
    $(".egamepre").click(function() {
        $index--;
        if ($index < 0) {
            $index = 4
        };
        $(".egamechoose span").eq($index).addClass("egamered").siblings().removeClass("egamered");
        pre();
        return $exdex = $index;
    });
    var atime = setInterval(function() {
        $(".egamenext").click();
    }, 5);
    function next() {
        $(".egamebanner a").eq($index).stop(true, true).css("left", "100%").animate({
            "left": "0"
        });
        $(".egamebanner a").eq($exdex).stop(true, true).css("left", "0").animate({
            "left": "-100%"
        });
    }

    function pre() {
        $(".egamebanner a").eq($index).stop(true, true).css("left", "-100%").animate({
            "left": "0"
        });
        $(".egamebanner a").eq($exdex).stop(true, true).css("left", "0").animate({
            "left": "100%"
        });
    }

});

function getQueryString(name) { 
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i"); 
    var r = window.location.search.substr(1).match(reg); 
    if (r != null) return unescape(r[2]); return null; 
}

function nofind(){
    var img=event.srcElement; 
    img.src= cdnUrl +'/shared/egames/images/pt/PT_bal.jpg'; 
    img.onerror=null; 控制不要一直跳动 
}