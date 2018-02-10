var $loginBack = getCookie('loginBack');
var platform = 'wap';

$(document).ready(function (e) {
    // $('#mescroll').css({ 'top': $('.head-h').height()+60 });
    // $('#mescroll').css({ 'bottom': jQuery('.foot-h').height() });
    // setTimeout(function () {
    //     loaded();
    // }, 100)
});
/***********************************登陆验证******************************/
//登陆判断
function loginId() {
    if (getCookie('loginBack')) {   //判断是否登陆
        var $loginBack = getCookie('loginBack');
        var $token = $loginBack;
        $.ajax({
            type: 'GET',
            url: '/m/ajaxLoginVerify',
            data: {},
            async: false,
            dataType: 'json',
            success: function (msg) {
                if (!msg['code']) {
                    loginData = msg['data'];
                } else {
                    delCookie("loginBack");
                }
            },
            error: function (data) {
                delCookie("loginBack");
            },
            headers: {
                'Authorization': 'bearer ' + $token,
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "wap"
            }
        });
    } else {
        delCookie("loginBack");
    }
    
    return loginData;
}
var loginData = loginId();
/***********************************END登陆验证******************************/

var myScroll;

// function loaded() {
//     myScroll = new IScroll('#mescroll', {
//         //preventDefault为false这行就是解决onclick失效问题
//         //为true就是阻止事件冒泡,所以onclick没用
//         preventDefault: false,
//         scrollbars: true,
//         mouseWheel: true,
//         interactiveScrollbars: true,
//         shrinkScrollbars: 'scale',
//         fadeScrollbars: true,
//         /*	checkDOMChanges:true,*/
//     });
// }

// document.addEventListener('touchmove', function (e) {
//     e.preventDefault();
// }, false);
// $(function () {
//     loaded();
// });

/*myScroll.refresh()*/


/***************************************存取COOKIE******************************************/

//存cookie
function setCookie(name, value, day) {

    if (day) {
        var exp = new Date();
        exp.setTime(exp.getTime() + day * 24 * 60 * 60 * 1000);
        document.cookie = name + "=" + value + ";path=/;expires=" + exp.toGMTString();
    } else {
        document.cookie = name + "=" + value + ";path=/;";
    }

}
//取cookie
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
        return arr[2];
    else
        return null;
}

//删除cookie
function delCookie(name) {
    if (name) {
        var exp = new Date();
        exp.setTime(exp.getTime() - 1000);
        var cval = getCookie(name);
        if (cval != null)
            document.cookie = name + "=" + cval + ";path=/;expires=" + exp.toGMTString();
    } else {
        var keys = document.cookie.match(/[^ =;]+(?=\=)/g);
        if (keys) {
            for (var i = keys.length; i--;)
                document.cookie = keys[i] + '=0;path=/;expires=' + new Date(0).toUTCString();
        }
    }
}

function clearCache() {
    $.toast("清除成功", 'text');
    delCookie();
    window.location.href = '/m/index';

}
/***************************************存取COOKIE END******************************************/


//验证码
function getKey(url) {
    // var xmlhttp;
    // xmlhttp = new XMLHttpRequest();
    // xmlhttp.open("GET", "/m/code", true);
    // xmlhttp.responseType = "blob";
    // xmlhttp.onload = function () {
    //     if (this.status == 200) {
    //         var blob = this.response;
    //         $("#vImg").src = window.URL.createObjectURL(blob);
    //         $("#vImg").attr("src", window.URL.createObjectURL(blob))
    //         $("#vImg").show();
    //         _Code = xmlhttp.getResponseHeader("Code")
    //         setCookie("Code", xmlhttp.getResponseHeader("Code"));
    //     }
    // }
    // xmlhttp.send();
    $.get("/code",{},function(data){
        var sdata = data.data;
        if (sdata.code){
            $("#vImg").attr({"src":sdata.image});
            _Code = sdata.code;
            setCookie("Code", sdata.code);
        }
    },"json")
}

/***********************************会员登出******************************/
function logOut() {
    if (getCookie('loginBack')) {
        var loginBack = getCookie('loginBack');
        $.ajax({
            type: "put",
            url: "/m/logout",
            headers: {
                'Authorization': 'bearer ' + loginBack,
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "wap"
            },
            contentType: "application/json",
            data: {},
            dataType: 'json',
            success: function (data) {
                if (data) {
                    $.toast(data.msg, 'text');
                } else {
                    delCookie('loginBack');
                    delCookie("loginIn");
                    delCookie("loginData");
                    window.location.href = "/m/index"
                }
            }
        });
    } else {
        $.toast("用户未登录", 'text');
    }
}

/***********************************END会员登出******************************/

$(function () {
    var loginIn = getCookie('loginBack');
    if (loginIn) {
        // var loginData = JSON.parse(getCookie('loginData'));
        $(".pk-login-before").hide();
        $(".pk-login-after").show();

    } else {
        $(".pk-login-before").show();
        $(".pk-login-after").hide();
    }
})

/***********************************date时间处理******************************/
/**
  * 和PHP一样的时间戳格式化函数
  * @param {string} format 格式
  * @param {int} timestamp 要格式化的时间 默认为当前时间
  * @return {string}   格式化的时间字符串
  */
function date(format, timestamp) {
    var a, jsdate = ((timestamp) ? new Date(timestamp * 1000) : new Date());
    var pad = function (n, c) {
        if ((n = n + "").length < c) {
            return new Array(++c - n.length).join("0") + n;
        } else {
            return n;
        }
    };
    var txt_weekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    var txt_ordin = {
        1: "st",
        2: "nd",
        3: "rd",
        21: "st",
        22: "nd",
        23: "rd",
        31: "st"
    };
    var txt_months = ["", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    var f = {   // Day 

        d: function () {
            return pad(f.j(), 2)
        }, D: function () {
            return f.l().substr(0, 3)
        }, j: function () {
            return jsdate.getDate()
        }, l: function () {
            return txt_weekdays[f.w()]
        }, N: function () {
            return f.w() + 1
        }, S: function () {
            return txt_ordin[f.j()] ? txt_ordin[f.j()] : 'th'
        }, w: function () {
            return jsdate.getDay()
        }, z: function () {
            return (jsdate - new Date(jsdate.getFullYear() + "/1/1")) / 864e5 >> 0
        },       // Week
        W: function () {
            var a = f.z(),
                b = 364 + f.L() - a;
            var nd2, nd = (new Date(jsdate.getFullYear() + "/1/1").getDay() || 7) - 1;
            if (b <= 2 && ((jsdate.getDay() || 7) - 1) <= 2 - b) {
                return 1;
            } else {
                if (a <= 2 && nd >= 4 && a >= (6 - nd)) {
                    nd2 = new Date(jsdate.getFullYear() - 1 + "/12/31");
                    return date("W", Math.round(nd2.getTime() / 1000));
                } else {
                    return (1 + (nd <= 3 ? ((a + nd) / 7) : (a - (7 - nd)) / 7) >> 0);
                }
            }
        },       // Month
        F: function () {
            return txt_months[f.n()]
        }, m: function () {
            return pad(f.n(), 2)
        }, M: function () {
            return f.F().substr(0, 3)
        }, n: function () {
            return jsdate.getMonth() + 1
        }, t: function () {
            var n;
            if ((n = jsdate.getMonth() + 1) == 2) {
                return 28 + f.L();
            } else {
                if (n & 1 && n < 8 || !(n & 1) && n > 7) {
                    return 31;
                } else {
                    return 30;
                }
            }
        },       // Year
        L: function () {
            var y = f.Y();
            return (!(y & 3) && (y % 1e2 || !(y % 4e2))) ? 1 : 0
        },    //o not supported yet
        Y: function () {
            return jsdate.getFullYear()
        }, y: function () {
            return (jsdate.getFullYear() + "").slice(2)
        },       // Time
        a: function () {
            return jsdate.getHours() > 11 ? "pm" : "am"
        }, A: function () {
            return f.a().toUpperCase()
        }, B: function () {    // peter paul koch:

            var off = (jsdate.getTimezoneOffset() + 60) * 60;
            var theSeconds = (jsdate.getHours() * 3600) + (jsdate.getMinutes() * 60) + jsdate.getSeconds() + off;
            var beat = Math.floor(theSeconds / 86.4);
            if (beat > 1000) beat -= 1000;
            if (beat < 0) beat += 1000;
            if ((String(beat)).length == 1) beat = "00" + beat;
            if ((String(beat)).length == 2) beat = "0" + beat;
            return beat;
        }, g: function () {
            return jsdate.getHours() % 12 || 12
        }, G: function () {
            return jsdate.getHours()
        }, h: function () {
            return pad(f.g(), 2)
        }, H: function () {
            return pad(jsdate.getHours(), 2)
        }, i: function () {
            return pad(jsdate.getMinutes(), 2)
        }, s: function () {
            return pad(jsdate.getSeconds(), 2)
        },    //u not supported yet
        // Timezone
        //e not supported yet
        //I not supported yet
        O: function () {
            var t = pad(Math.abs(jsdate.getTimezoneOffset() / 60 * 100), 4);
            if (jsdate.getTimezoneOffset() > 0) t = "-" + t;
            else t = "+" + t;
            return t;
        }, P: function () {
            var O = f.O();
            return (O.substr(0, 3) + ":" + O.substr(3, 2))
        },    //T not supported yet
        //Z not supported yet
        // Full Date/Time
        c: function () {
            return f.Y() + "-" + f.m() + "-" + f.d() + "T" + f.h() + ":" + f.i() + ":" + f.s() + f.P()
        },    //r not supported yet
        U: function () {
            return Math.round(jsdate.getTime() / 1000)
        }
    };
    return format.replace(/[\\]?([a-zA-Z])/g, function (t, s) {
        if (t != s) {    // escaped 

            ret = s;
        } else if (f[s]) {    // a date function exists 

            ret = f[s]();
        } else {    // nothing special 

            ret = s;
        }
        return ret;
    });
}

/***********************************ENDdate时间处理******************************/

/***********************************修改密码******************************/
function eidtPas() {
    var $loginBack = getCookie('loginBack');
    var $beforePassword = $("input[name='beforePassword']").val();
    var $password = $("input[name='password']").val();
    var $confirmPassword = $("input[name='confirmPassword']").val();
    var $types = Number($('select[name="type"]').val());
    if ($beforePassword.length == 0) {
        $.toast("原密码不能为空！", 'text');
        return false;
    } else{
        if ($types == 1) {
            if ($beforePassword.length < 6 || $beforePassword.length > 12) {
                alert("原密码长度为6-12位的数字或字母组成！");
                return false;
            }
        }else{
            if ( !(/^\d{4}$/).test($beforePassword) ) {
                alert("取款密码为4位数的数字");
                return false;
            }
        }
    }
    if ($password.length == 0) {
        $.toast("原密码不能为空！", 'text');
        return fasle;
    } else{
        if ($types == 1) {
            if ($password.length < 6 || $password.length > 12) {
                alert("新密码长度为6-12位的数字或字母组成！");
                return false;
            }
        }else{
            if ( !(/^\d{4}$/).test($password) ) {
                alert("取款密码为4位数的数字");
                return false;
            }
        }

        if ($beforePassword == $password) {
            alert("原密码和新密码不能相同");
            return false;
        }
    }

    if ($confirmPassword != $password) {
        $.toast("两次输入的新密码不一致", 'text');
        return false;
    }
    $('dl.m-myaccount button').attr("disabled", true); //按钮失效
    $.ajax({
        type: "put",
        url: "/m/editPassword",
        headers: {
            'Authorization': 'bearer ' + $loginBack,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "wap"
        },
        data: JSON.stringify({
            beforePassword: $beforePassword,
            password: $password,
            confirmPassword: $confirmPassword,
            types: $types
        }),
        dataType: 'json',
        success: function (data) {
            if (data) {
                $.toast(data.msg, 'text');
                $("input[name='beforePassword']").val('');
                $("input[name='password']").val('');
                $("input[name='confirmPassword']").val('');
            } else {
                if ($types == 1) {
                    delCookie("loginBack");
                    $.toast("登录密码修改成功, 请重新登录！", 'text');
                    setTimeout("window.location.href = '/m/login'; ", 1000);//延迟1秒执行
                }else{
                    $.toast("取款密码修改成功, 请重新登录！", 'text');
                    history.back(-1);
                }
                return;
            }
            $('dl.m-myaccount button').attr("disabled", false); //按钮有效
        }
    });
}

/***********************************END修改密码******************************/

/***********************************银行列表******************************/
function bank() {
    var $loginBack = getCookie('loginBack');
    $.ajax({
        type: "get",
        url: "/m/back",
        headers: {
            'Authorization': 'bearer ' + $loginBack,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "wap"
        },
        data: {},
        dataType: 'json',
        success: function (data) {
            if (data) {
                if (!data.code) {
                    var $sdata = data.data;
                    var $htmlStr = '';
                    if ($sdata) {
                        for (var k in $sdata) {
                            var v = $sdata[k];
                            $htmlStr += '<li onclick="backClick(this)"><a style="background:url(' + cdnUrl + '/style/images/bank/'+v.bank_id+'.png) 0% 0% / 100% 100%">' + v.card + '</a><div class="col hide">' +
                                '<p><span>开户姓名</span>： <span class="ng-binding">' + v.card_name + '</span></p>' +
                                '<p><span>开户账号</span>： <span class="ng-binding">' + v.card + '</span></p>' +
                                '<p><span>银行名称</span>： <span class="ng-binding">' + v.title + '</span></p>' +
                                '<p><span>开户网点</span>： <span class="ng-binding">' + v.card_address + '</span></p>' +
                                '</div></li>';
                        }
                        if ($sdata.length >= 3) {
                            $("ul.addcord").addClass('hide');
                        }
                    }else{
                        $("ul.addcord a").attr({'href':"bankCardAdd?pass=1"});
                    }
                    $("ul.cardlist").html($htmlStr);

                } else {
                    $.toast(data.msg, 'text');
                }
            }
        }
    });
}

//点击银行列表时间
function backClick(obj) {
    $(obj).siblings('li').find('div').addClass('hide');
    var $devObj = $(obj).find('div')
    if ($devObj.hasClass('hide')) {
        $devObj.removeClass('hide');
    } else {
        $devObj.addClass('hide');
    }
}

/***********************************END银行列表******************************/

/***********************************添加银行******************************/
//获取可用的银行列表
function bankList() {
    var $loginBack = getCookie('loginBack');
    $.ajax({
        type: "get",
        url: "/m/backAddInfo",
        headers: {
            'Authorization': 'bearer ' + $loginBack,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "wap"
        },
        data: {},
        dataType: 'json',
        success: function (data) {
            if (data) {
                if (!data.code) {
                    var $sdata = data.data, $htmlStr = '';
                    for (var k in $sdata) {
                        var v = $sdata[k];
                        $htmlStr += '<option value="' + v.id + '">' + v.title + '</option>'
                    }
                    $("select[name='bank']").html($htmlStr);
                } else {
                    $.toast(data.msg, 'text');
                }
            }
        }
    });
}

function bankAdd() {
    var $loginBack = getCookie('loginBack');
    var $bankId = Number($("select[name='bank']").val());
    var $card = $("input[name='card']").val();
    var $cardName = $("input[name='cardName']").val();
    var $cardAddress = $("input[name='cardAddress']").val();
    if (!$("input[type='password']").is('hidden')) {
        var $password = $("input[name='password']").val();
        var $comPassword = $("input[name='comPassword']").val();
        var data = JSON.stringify({
            'bank_id' : Number($bankId),
            'card' : $card,
            'card_name' : $cardName,
            'card_address' : $cardAddress,
            'password' : $password
        });
    }else{
        var data = JSON.stringify({
            'bank_id' : Number($bankId),
            'card' : $card,
            'card_name' : $cardName,
            'card_address' : $cardAddress
        }); 
        }
    $.ajax({
        type: "post",
        url: "/m/bankAdd",
        headers: {
            'Authorization': 'bearer ' + $loginBack,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "wap"
        },
        data: data,
        dataType: 'json',
        success: function (data) {
            if (data) {
                if (data.code) {
                    $.toast(data.msg, 'text');
                } else {

                }
            } else {
                $.toast("添加成功", 'text');
                setTimeout("window.location.href = '/m/bankCard'; ", 1000);//延迟2秒执行
            }
        }
    });
}

/***********************************END添加银行******************************/

/***********************************优惠******************************/
function getActivity(fn) {
    $.ajax({
        type: "get",
        url: "/m/getYouhui",
        data: {},
        dataType: 'json',
        async: false,
        // beforeSend: function () {
        //     ajaxLoading($("#mescroll"));
        // },
        success: function (data) {
            if (data) {
                if (data.code) {
                    $.toast(data.msg, 'text');
                    return false
                } else {
                    var $yhTitleData = data.data.yhTitleData;
                    var $yhData = data.data.yhData;
                    var $YhWidth = data.data.YhWidth;
                    var $titleStr = '', $htmlStr = '';
                    $titleStr = '<li id="PT1" class="curr">所有活动</li>';
                    if ($yhTitleData.length >= 1) {
                        for (var k in $yhTitleData) {
                            var v = $yhTitleData[k];
                            $titleStr += '<li id=' + v.id + '>' + v.title + '</li>';
                        }
                    }
                    // console.log($yhTitleData)

                    for (var key in $yhData) {
                        var val = $yhData[key];
                        $htmlStr += '<section class="my_section" data-type="' + val.TopId + '" ><a class="memExclusive" href="javascript:;"><img src="' + val.Img + '" alt="优惠图片"></a>' +
                            '<div class="eventtext" style="display:none;">' + val.Content + '</div></section>'
                    }
                    $("div.active-index-c ul").html($titleStr);

                    $("#mescroll .m-activits").children().remove(".my_section")

                    var hh=$($htmlStr);
                    $("#mescroll .m-activits").append(hh)
                    
                    var num = $('div.active-index-c ul li').length;
                    if (num < 5) {
                        $("div.active-index-c ul").css({ 'width': '100%' }).find('li').css({ 'width': 100 / num + '%' });
                    }
                    // setTimeout('$("#mask").remove()', 500);
                    fn();
                }
            }
        }
    });
}

/***********************************END优惠******************************/

/***********************************电子******************************/
//电子导航
function getGameTitle(gtype) {
    $.ajax({
        type: "get",
        url: "/m/gameTitle",
        data: {},
        dataType: 'json',
        async: false,
        success: function (data) {
            if (data) {
                if (data.code) {
                    $.toast(data.msg, 'text');
                    return false
                } else {
                    console.log(data);
                    var $gameTitleStr = '';
                    var $sdata = data.data;
                    for (var k in $sdata) {
                        var v = $sdata[k];
                        if (v == gtype) {
                            $gameTitleStr += '<li class="egame-tab-active" lang="cn" data-type="' + v + '">' + v + '电子</li>';
                        } else {
                            $gameTitleStr += '<li lang="cn" data-type="' + v + '">' + v + '电子</li>';
                        }
                    }
                    $(".bank-tab .bank-nav").html($gameTitleStr);
                }
            }
        }
    });
}

$.EGame = {
    eData: {},
    EgDemoUrl: "/index.php/games/Logingame/EgDemo",
    PtDemoUrl: "http://cache.download.banner.longsnake88.com/flash/79/launchcasino.html?mode=offline&affiliates=1&language=ZH-CN&game=",
    BBinUrl: "/login?g_type=bbin",
    AllGameUrl: "/index.php/games/Logingame/index",
    _openGame: function (url) {
        window.open(url, '_blank', 'width=1000,height=800,top=0,left=0,status=no,toolbar=no,scrollbars=yes,resizable=no,personalbar=no');
    },
    //进入游戏
    inGame: function (code, g_type, sw) {
        var loginIn = getCookie('loginIn');
        if (loginIn) {
            var loginData = JSON.parse(getCookie('loginData'));
            var userid = loginData['id'];
        } else {
            var userid = "";
        }

        var _this = $.EGame;
        if (!arguments[2]) sw = 0;
        var param = "?gameid=" + code + "&g_type=" + g_type;

        if (g_type === 'eg' && (userid == "" || sw == 1)) {
            _this._openGame(_this.EgDemoUrl + param);
            return false;
        }

        if (g_type === 'pt' && sw == '1') {
            _this._openGame(_this.PtDemoUrl + code);
            return false;
        }

        if (userid == "") {
            $.toast("用户登录信息已过期,请重新登录！", 'text');
            setTimeout("window.location.href = '/m/login';", 1000);
        } else {
            if (g_type === 'bbin') {
                _this._openGame(_this.BBinUrl);
            } else {
                _this._openGame(_this.AllGameUrl + param + "&sw=" + sw);
            }
        }
    },
    getData: function (gtype) {
        if (!gtype) {
            gtype = 'EG';
        }
        $('#tab1').css('position', 'relative');
        $.ajax({
            type: "get",
            url: "/m/gameData",
            data: { type: gtype },
            dataType: 'json',
            async: false,
            beforeSend: function () {
                ajaxLoading($("#mescroll"));
            },
            success: function (data) {
                if (data) {
                    if (data.code) {
                        $.toast(data.msg, 'text');
                        return false
                    } else {
                        if (data.data.wh == 1) { //维护
                            // $.EGame.GetWhHtml(msg['wh'].content);
                            // return false;
                        }
                        var $gameDataStr = '';
                        var $sdata = data.data.data;
                        if ($sdata) {
                            for (var k in $sdata) {
                                var v = $sdata[k];
                                $gameDataStr += '<li><a href="javascript:$.EGame.inGame(\'' + v.gameid + '\', \'' + gtype.toLowerCase() + '\', \'1\');">';
                                if (gtype == "EG") {
                                    $gameDataStr += '<img src="' + cdnUrl + '/wap/style/official/images/pk/' + v.image + '" alt="">';
                                } else if (gtype == "HB") {
                                    $gameDataStr += '<img src="' + v.image + '" alt="">';
                                } else {
                                    if (v.image) {
                                        $gameDataStr += '<img src="' + cdnUrl + '/wap/style/official/images/' + gtype.toLowerCase() + '/' + v.image + '" alt="">';
                                    } else {
                                        $gameDataStr += '<img src="https://app-a.insvr.com/img/square/160/7d176542-0de1-4144-afbf-14d6998d2cbd_zh-CN.png" alt="">';
                                    }
                                }
                                $gameDataStr += '<span>' + v.name + '</span></a></li>';
                            }
                        } else {
                            $gameDataStr += '<li><a href="opengame()">暂无数据</a><li>';
                        }
                        $gameDataStr += '<script>$("input#game-search").quicksearch("#mescroll ul li");</script>';
                        $("#mescroll .egame-bar-item ul").html($gameDataStr);
                        console.log($gameDataStr)
                        $('#etype').val(gtype);
                        $("#mask").remove();
                        // loaded();
                    }
                }
            }
        });

    },
    GetWhHtml: function (whcontent) { //维护
        // var games = '';
        // games += "<h1 style='margin-left:20px;font-size:64px;'>此游戏正在维护，请稍后访问······</h1><br/>";
        // games += "<h1 style='margin-left:20px;font-size:64px;'>"+whcontent+"</h1>";
        // $('#ul_1,#con_one_1 div.search').hide();
        // $('#xxoo').remove();
        // $('#page_navigation>.btndiv').html('');
        // $('.games_menu').html(games);
    }
};

/***********************************END电子******************************/
//获取url参数
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;

}

/*加载动画*/
function Loading() {
    var createDiv = document.createElement("div");
    if (cdnUrl == undefined) {
        var cdnUrl = "/template";
    }
    createDiv.id = "mask";
    createDiv.style.position = 'fixed';
    createDiv.style.backgroundColor = 'black';
    createDiv.style.zIndex = '1002';
    createDiv.style.opacity = '0.5';
    createDiv.style.display = 'flex';
    createDiv.style.flex = 'fixed';
    createDiv.style.flexDirection = 'row';
    createDiv.style.justifyContent = 'center';
    createDiv.style.alignItems = 'center';
    createDiv.style.width = '100%';
    createDiv.style.top = '0';
    createDiv.style.left = '0';
    createDiv.style.height = document.documentElement.clientHeight + 'px';
    var bigImg = document.createElement("img");		//创建一个img元素
    bigImg.src = cdnUrl + "/wap/style/images/loading.gif";   //给img元素的src属性赋值
    bigImg.style.margin = '0 auto';
    bigImg.style.display = 'block';
    createDiv.appendChild(bigImg);
    document.body.appendChild(createDiv);
}

/*关闭加载动画*/
function LoadingClose(time) {
    if (!time) {
        time = 500;
    } else {
        time = time * 1000;
    }
    setTimeout(function () {
        $('#mask').remove();
    }, time);
    return;
}


/*********暂无数据展示*********/
function ShowNoData(url, mes) {
    if (cdnUrl == undefined) {
        var cdnUrl = "/template";
    }
    var divll = document.createElement('div');
    divll.setAttribute('class', 'nodate');
    divll.style.color='black';
    divll.style.textAlign='center';
    divll.style.height='50px';
    divll.style.lineHeight='50px';

    // var showimg = document.createElement('img');
    // if (!url) {
    //     url = '/wap/style/images/empty@2x.png';
    // }
    // showimg.src = cdnUrl + url;
    var showp = document.createElement('p');
    if (!mes) {
        mes = '暂无数据';
    }
    showp.style.display='block';
    showp.innerHTML = mes;

    // divll.appendChild(showimg);
    divll.appendChild(showp);
    return divll;
}

//ajax 页面加载动画
function ajaxLoading(obj) {
    obj.prepend('<div id="mask" class="mask"><img src="' + cdnUrl + '/wap/style/images/ajax-loader.gif" alt=""></div>');
    $("#mask").css({
        position: 'absolute',
        top: 0,
        bottom: 0,
        left: 0,
        right: 0,
        'background-color': '#777',
        'z-index': 1002,
        left: 0,
        opacity: 0.5,
        '-moz-opacity': 0.5,
        display: 'flex',
        'flex-direction': 'row',
        'justify-content': 'center',
        'align-items': 'center',
    });
    $("#mask img").css({ margin: '0 auto' });
}

//时间戳转日期
function strToDate(data) {
    var nowdate, Y, M, D, h, m, s, data1, data2;
    var datatime ;
    if (data === 0 || data === -1) {
        nowdate = new Date();
        nowdate.setDate(nowdate.getDate() + parseInt(data));
    } else {
        nowdate = new Date(parseInt(data) * 1000);//时间戳转换为日期格式*1000
    }
    // console.log(nowdate);
    Y = nowdate.getFullYear() + '-';
    M = (nowdate.getMonth() + 1 < 10 ? '0' + (nowdate.getMonth() + 1) : nowdate.getMonth() + 1) + '-';
    D = nowdate.getDate()< 10 ? '0'+nowdate.getDate() + ' ':nowdate.getDate() + ' ';
    h = nowdate.getHours()< 10 ? '0'+ nowdate.getHours()+ ':':nowdate.getHours()+ ':';
    m = nowdate.getMinutes()< 10 ? '0'+nowdate.getMinutes() + ':':nowdate.getMinutes()+ ':';
    s = nowdate.getSeconds()< 10 ? '0'+ nowdate.getSeconds():nowdate.getSeconds();
    data1 = Y + M + D;
    data2 = h + m + s;
    datatime = data1+ data2;   
    if (data === 0 || data === -1) {
        return data1 + data2;
    } else {
        return datatime;
    }

}

//日期转时间戳
function DateToStr(data) {
    var datatime, time1;
    if (data == 0 || data == -1||data==null||data==-7) {
        datatime = new Date();
        datatime.setHours(0);
        datatime.setMinutes(0);
        datatime.setSeconds(0);
        datatime.setMilliseconds(0);
    } else {
        datatime = new Date(data.replace(/-/g, '/'));// 有三种方式获取，在后面会讲到三种方式的区别
    }
    if (data == -1) {
        time1 = (Date.parse(datatime) / 1000) - 24 * 3600;
    }else if(data==-7){
        time1 = (Date.parse(datatime) / 1000) - 24 * 3600*7;
    }else {
        time1 = Date.parse(datatime) / 1000;
    }
    return time1;
}


let cent_flag = true;
$(document).on('click', 'nav.f-hmoe .cent span', function () {
    if (cent_flag) {
        cent_flag = false;
        $('section.p-semicircle').slideToggle(500, function () {
            cent_flag = true;
        });
        $('section.big-shade').fadeToggle(300);
    }
});
$(document).on('click', '.big-shade', function () {
    $('section.p-semicircle').slideToggle(500, function () {
        cent_flag = true;
    });
    $('section.big-shade').fadeToggle(300);
});

/***********************************在线客服******************************/
function OnlineService(url) {
    newWin = window.open(url, '', 'width=900,height=600,top=0,left=0,status=no,toolbar=no,scrollbars=yes,resizable=no,personalbar=no');
    window.opener = null;//出掉关闭时候的提示窗口
    window.open('', '_self'); //ie7
    window.close();
}

//头部栏的返回上一页
$('.head-h .h-head').on("click", "li .wap-fanhui", function () {
    window.history.go(-1);
})

function detailInfo() {
    $.ajax({
        type: "GET",
        url: '/m/detail',
        data: {},
        async: false,
        dataType: 'json',
        success: function (msg) {
            if (msg) {
                if (msg.code) {
                    $.toast(msg.msg, "text");
                    setTimeout('$("#mask").remove()', 500);
                    return false;
                }
                var sdata = msg.data;
                var $obj = $("#mescroll dl li");
                $obj.find("input[name='birthday_num']").val(date('Y-m-d', sdata.birthday));
                $obj.find("input[name='email_num']").val(sdata.email);
                $obj.find("input[name='phone_num']").val(sdata.mobile);
                $obj.find("input[name='qq_num']").val(sdata.qq);
                $obj.find("input[name='wechat']").val(sdata.wechat);
                $obj.find("input[name='card']").val(sdata.card);
                $obj.find("textarea[name='remark']").val(sdata.remark);
            } else {
                $.toast('获取会员详细信息失败', "text");
                setTimeout('$("#mask").remove()', 500);
                return
            }
            setTimeout('$("#mask").remove()', 500);
        },
        beforeSend: function () {
            ajaxLoading($("#mescroll"));
        },
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': platform
        }
    });
}

    // document.querySelector("#mescroll").addEventListener("touchmove", function(e){
    //     e.preventDefault(); 
    // }
    // , {
    //     passive: true,
    //     capture: false,
    // });

//上下滑动的事件
// function touchUpUnder(MYtouchUnder, MYtouchUp,MYtouchEnd) {
//     var startX, startY, endX, endY;
//     var getdom = document.querySelector("#mescroll");

//     getdom.addEventListener("touchstart", touchStart, {
//         passive: true,
//         capture: false,
//     });
//     getdom.addEventListener("touchmove", touchMove, {
//         passive: true,
//         capture: false,
//     });
//     getdom.addEventListener("touchend", touchEnd, {
//         passive: true,
//         capture: false,
//     });
//     function touchStart(event) {
//         var touch = event.touches[0];
//         startY = touch.pageY;
//     }
//     function touchMove(event) {
//         var touch = event.touches[0];
//         endY = startY - touch.pageY;
//         if (endY < 0) {
//             MYtouchUp()
//         } else if (endY > 0) {
//             MYtouchUnder()
//         }
//     }


//     function touchEnd(event) {

//     }
// }

// function footerClass(){
//     var currPage={{.CurrPage}};
//     $(".clear a").eq(currPage-1).addClass("")
// }
/***************分页 重要 勿动****************/
//分页
function pageList(metalist, links) {
    // console.log(metalist);
    var totalNum = metalist.count;
    var totalPage = metalist.page_count;
    var current_page = metalist.current_page;
    // var page_size = metalist.page_size;
    var htmldata = '<div class="one-foot clear"><p class="fl">共' + totalNum + '条，共计' + totalPage + '页</p><div class="fr">';
    for (var i = 1; i <= totalPage; i++) {
        if (i == current_page) {
            htmldata += '<span class="one-foot-circle page-active">' + i + '</span>';
        } else {
            htmldata += '<span class="one-foot-circle">' + i + '</span>';
        }
    }
    // htmldata += '<span class="iconfont icon-forward"></span>';
    htmldata += '<span>跳转至<select name="select"  class="xla_k ">';

    for (var j = 1; j <= totalPage; j++) {
        if (j == current_page) {
            htmldata += '<option value="'+j+'" selected>'+j+'</option>';
        }else{
            htmldata += '<option value="'+j+'" >'+j+'</option>';
        }

    }
    htmldata += '</select><span>页</span></span></div></div>';
    return htmldata;
}
//分页需要调用的事件
function GetPageData(FuncName,metadata) {
    $('.one-foot .fr').on('click', '.one-foot-circle', function () {
        var ind = $(this).index();
        $(this).addClass('page-active').siblings().removeClass('page-active');
        var page = Number($(this).text());
        FuncName(page);
    });
    $('.xla_k').change(function () {
        var page = $(this).val();
        FuncName(page);
    });
    $('.icon-back').click(function () {
        var page = $(' .xla_k option:selected').val();
        var getpage=Number(page)-1;
        if(getpage<1){
            $.toast('没有更多内容', 'text');
        }else{
            FuncName(getpage);
        }
    });
    $(' .icon-forward').click(function () {
        var page = $('.xla_k option:selected').val();
        var getpage=Number(page)+1;
        // console.log(getpage)
        if(page>(metadata-1)){
            $.toast('没有更多内容', 'text');
        }else{
            FuncName(getpage);
        }
    });
}
/********************分页 重要 勿动******************/

console.log(getQueryString('pass'));
/*********** 隐藏金额 start *******************/
function hideBalance(obj) {
    var str=obj.text();
    var hideStr='******';
    if(str!=hideStr){
        obj.attr('balance',str);
        obj.text(hideStr);
    }else{
        var balance= obj.attr('balance');
        obj.text(balance);
    }
}
/********************* 隐藏金额 end *********************************/