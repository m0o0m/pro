/**
 * Created by Administrator on 2017/1/3.common_member
 */

//定义验证时的全局变量
var _User = '';
var _Pass = '';
var _Rnum = '';
var _UserId = '';
var _Code = '';

/***********************************js获取会员uid***************************/
function getUserId() {
    $.post("/index.php/Index/getUserOnline", {}, function (data) {
        _UserId = data;
    });
}

/***********************************js试玩注册***************************/
function joinDemoDo() {
    if (!getCookie('registerDemo') && !getCookie('uid')) {
        $.get("/index.php/webcenter/Register_web/getDemoUserToken", {}, function (data) {
            console.log(data);
            var checkCode = getCookie('registerDemo');
            if (checkCode) {
                joinDemoUser(checkCode);
            } else {
                alert("请勿连续注册试玩账号！");
            }
        });
    } else alert("登陆用户不能注册试玩账号！");
}

function joinDemoUser(checkCode) {
    $.post("/index.php/webcenter/Register_web/joinDemoUserDo", {'pk_token': checkCode}, function (data) {
        if (data.state == 1) {
            alert(data.msg);
            location.href = '/index.php/Index/N_index';
        } else {
            alert(data.msg);
        }
    }, 'json');
}

/***********************************END js试玩注册***************************/

/***************************************存取COOKIE******************************************/

//存cookie
function setCookie(name, value, day) {
    if(day){
        var exp = new Date();
        exp.setTime(exp.getTime() + day * 24 * 60 * 60 * 1000);
        document.cookie = name + "=" + value + ";path=/;expires=" + exp.toGMTString();
    }else{
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
        var cval=getCookie(name);
        if(cval!=null)
            document.cookie= name + "="+cval+";path=/;expires="+exp.toGMTString(); 
    }else{
        var keys=document.cookie.match(/[^ =;]+(?=\=)/g); 
        if (keys) { 
        for (var i = keys.length; i--;) 
            document.cookie=keys[i]+'=0;path=/;expires=' + new Date( 0).toUTCString();
        }
    }
}



//存intr代理推广
if (getCookie('intr') && !getCookie('setSon')) {
    setCookie('setSon', 1, 1);
    $.get("/index.php/Index/setSessionIntr", {'intr': getCookie('intr')}, function (e) {
        if (e == 1) {
            setCookie('intr', '', 1);
            setCookie('setSon', '', 1);
        }
    });
}
//存u会员推广
if (getCookie('u') && !getCookie('setSonU')) {
    setCookie('setSonU', 1, 1);
    $.get("/index.php/Index/setSessionU", {'u': getCookie('u')}, function (e) {
        if (e == 1) {
            setCookie('u', '', 1);
            setCookie('setSonU', '', 1);
        }
    });
}
/***************************************存取COOKIE END******************************************/

/***********************************验证码******************************/

//验证码
// function getKey(url) {
//     $("#vPic").attr("src", "/code?type="+Math.random()+(new Date).getTime());
//     $("#vPic").show();
// }

// function getYzm(url) {
//     $("#vImg").attr("src", "/code?type="+Math.random()+(new Date).getTime());

//     $("#vImg").show();
// }
/***********************************END验证码******************************/
function getYzm(url) {

    // $.ajax({
    //     type: "GET",
    //     url : "/code?type="+codetime,
    //     success : function (data, status, xhr) {
    //         //获取响应头全部参数信息
    //         //alert(xhr.getAllResponseHeaders());
    //         // 获取指定响应头参数信息
    //         //alert(xhr.getResponseHeader( "Code" ));
    //         //$("#vImg").attr("src", "/code?type="+codetime);
    //         _Code=xhr.getResponseHeader( "Code" )
    //         console.log(_Code)
    //         $("#vImg").show();
    //         return
    //         //return xhr.getResponseHeader( "Code" )
    //     }
    //
    // });

    // var codetime=Math.random()+(new Date).getTime()
    // $.ajax({
    //     url : "/code?type="+codetime,
    //     success : function (data, status, xhr) {
    //         //获取响应头全部参数信息
    //         //alert(xhr.getAllResponseHeaders());
    //         // 获取指定响应头参数信息
    //         //alert(xhr.getResponseHeader( "Code" ));
    //         //$("#vImg").attr("src", "/code?type="+codetime);
    //         _Code=xhr.getResponseHeader( "Code" )
    //         console.log(_Code)
    //         $("#vImg").show();
    //         return
    //         //return xhr.getResponseHeader( "Code" )
    //     }
    //
    // });

    // var xmlhttp;
    // xmlhttp = new XMLHttpRequest();
    // xmlhttp.open("GET", "/code", true);
    // xmlhttp.responseType = "blob";
    // xmlhttp.onload = function () {
    //     if (this.status == 200) {
    //         var blob = this.response;
    //         $("#vImg").src = window.URL.createObjectURL(blob);
    //         // $("#vImg").attr("src", window.URL.createObjectURL(blob))
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

function getKey(url) {
    // var xmlhttp;
    // xmlhttp = new XMLHttpRequest();
    // xmlhttp.open("GET", "/code", true);
    // xmlhttp.responseType = "blob";
    // xmlhttp.onload = function () {
    //     if (this.status == 200) {
    //         var blob = this.response;
    //         $("#vPic").src = window.URL.createObjectURL(blob);
    //         $("#vPic").attr("src", window.URL.createObjectURL(blob))
    //         $("#vPic").show();
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
            setCookie("Code", _Code);
        }
    },"json")
}


/***********************************会员登录******************************/
//登陆验证
function aLeftForm1Sub() {
    console.log(_Code);
    var un = $("#username").val();
    if (un == "" || un == _User) {
        $("#username").focus();
        $("#username").parent('p').css({'border-color':'red'});
        return false;
    }else{
        $("#username").parent('p').css({'border-color':'#92b1e0'});
    }
    var pw = $("#password").val();
    if (pw == "" || pw == _Pass) {
        $("#password").focus();
        $("#password").parent('p').css({'border-color':'red'});
        return false;
    }else{
        $("#password").parent('p').css({'border-color':'#92b1e0'});
    }
    var rmNum = $("#rmNum").val();
    if (rmNum == "" || rmNum == _Rnum) {
        $("#rmNum").focus();
        $("#rmNum").parent('p').css({'border-color':'red'});
        return false;
    }else{
        $("#rmNum").parent('p').css({'border-color':'#92b1e0'});
    }
    $("#submit").attr("disabled", true); //按钮失效
    $.ajax({
        type: "post",
        url: "/login",
        headers: {"Code": _Code, 'platform': 'pc'},
        contentType: "application/json",
        data: JSON.stringify({
            account: $('input[name="account"]').val(),
            password: $('input[name="password"]').val(),
            code: $('input[name="code"]').val()
        }),
        dataType: 'json',
        success: function (data) {
            if (data.code == 20021) {
                alert("验证码错误，请重新输入");
                $("#rmNum").select();
            } else if (data.code == 20001) {
                alert("账号密码不匹配,请重新输入!");
                $("#rmNum").val('');
                $("#password").val('');
                $("#username").select();
            } else if (data.code == 30138) {
                alert("账号不存在!");
                $("#rmNum").val('');
                $("#password").val('');
                $("#username").select();
            } else if (data.code == 20002) {
                alert("对不起，账户已暂停使用,请联系在线客服！");
            } else if (data.code == 60000) {
                alert("系统错误！");
            } else {
                setCookie("loginBack", data['data'].token);    //将登陆信息写入缓存
                window.location.href = '/login/info';
                return;
            }
            $("#submit").attr("disabled", false); //按钮有效
        }
    });


    // $.post("login",JSON.stringify({account:un,password:pw,code:rmNum}),function(data){
    //     if(data == '5'){
    //         alert("验证码错误，请重新输入");
    //         $("#rmNum").select();
    //     }else if(data == 4){
    //         alert("账号密码不匹配,请重新输入!");
    //         $("#rmNum").val('');
    //         $("#password").val('');
    //         $("#username").select();
    //     }else if(data == 3){
    //         alert("账号不存在!");
    //         $("#rmNum").val('');
    //         $("#password").val('');
    //         $("#username").select();
    //     }else if(data == 2){
    //         alert("对不起，账户已暂停使用,请联系在线客服！");
    //     }else if(data == 1){
    //         window.location.href= '/login/info';
    //         return;
    //     }
    //     alert(data.code);
    //     alert(data.msg);
    //     $("#submit").attr("disabled",false); //按钮有效
    // });
}

//老版本
function mem_login() {
    var uname = $("#username").val();
    if (uname == "" || uname == _User) {
        $("#username").focus();
        return false;
    }
    var pwd = $("#password").val();
    if (pwd == "" || pwd == _Pass) {
        $("#password").focus();
        return false;
    }
    var rmNum = $("#rmNum").val();
    if (rmNum == "" || rmNum == _Rnum) {
        $("#rmNum").focus();
        return false;
    }

    $("#submit").attr("disabled", true); //按钮失效
    $.post("../webcenter/Login/login_do", {
        r: Math.random(),
        action: "login",
        username: uname,
        password: pwd,
        vlcodes: rmNum
    }, function (data) {
        if (data == '5') {
            alert("验证码错误，请重新输入");
            $("#rmNum").select();
        } else if (data == 4) {
            alert("账号密码不匹配,请重新输入!");
            $("#rmNum").val('');
            $("#password").val('');
            $("#username").select();
        } else if (data == 3) {
            alert("账号不存在!");
            $("#rmNum").val('');
            $("#password").val('');
            $("#username").select();
        } else if (data == 2) {
            alert("对不起，账户已暂停使用,请联系在线客服！");
        } else if (data == 1) {
            window.location.href = '/login/info';
            return;
        }
        $("#submit").attr("disabled", false); //按钮有效
    });
}

/***********************************END会员登录******************************/

//输入框验证
$(document).ready(function () {

    _User = $('#username').val();
    _Pass = $('#password').val();
    _Rnum = $('#rmNum').val();

    $('#username').focus(function () {
        if ($(this).val() == _User) {
            $(this).val('');
        }
    }).blur(function () {
        if ($(this).val() == '') {
            $(this).val(_User);
        }
    });

    $('#password').focus(function () {
        if ($(this).val() == _Pass) {
            $(this).val('');
        }
    }).blur(function () {
        if ($(this).val() == '') {
            $(this).val(_Pass);
        }
    });
    $('#rmNum').focus(function () {
        if ($(this).val() == _Rnum) {
            $(this).val('');
        }
        getYzm('');
    }).blur(function () {
        if ($(this).val() == '') {
            $(this).val(_Rnum);
        }
    });

});

/***********************************会员登出******************************/
function logOut() {
    var loginBack = getCookie('loginBack');
    if (loginBack) {
        $.ajax({
            type: "put",
            url: "/logout",
            headers: {
                'Authorization': 'bearer ' + loginBack,
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': 'pc'
            },
            contentType: "application/json",
            data: {},
            dataType: 'json',
            success: function (data) {
                if (data) {
                    alert(data.msg);
                } else {
                    delCookie('loginBack');
                    alert("会员登出成功！");
                    window.location.href = "/index"
                }
            }
        });
    }else{
        alert("会员登出成功！");
        window.location.href = "/index"
    }
}

/***********************************END会员登出******************************/

