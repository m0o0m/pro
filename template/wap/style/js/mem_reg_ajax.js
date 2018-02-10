if (getCookie('loginBack')) {
    window.location.href = '/m';
} else {
    $.ajax({
        'url': '/m/ajax/register/set',
        'type': 'get',
        'data': {},
        'cache': false,
        'timeout': 30000,
        async: false,
        'error': function (e, textStatus) {
            if (textStatus == 'timeout') {
                alert('网路品质不佳!!');
            }
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            var htmldata = '';
            var datalist = data.data;
            var agencyId = datalist.agencyId;
            var agreement = datalist.agreement,
                data = datalist.regSet[0];
            // if (data.isShowName == 1) {
            //     htmldata += '<ul><li><i class="iconfont wap-privilege"></i><input type="text" value=' + agencyId + ' name="introducer_member" readonly="readonly"/></li><p></p></ul>';
            // }
            if (data.isCode == 1) {
                htmldata += '<ul> <li><i class="iconfont wap-yanzhengma"></i><input type="text" name="code" placeholder="请输入验证码"><img id="vImg" src="' + cdnUrl + '/wap/style/images/no.png" onclick="getYzm();" style="height: 20px; "></ul>';
            }
            $('#is_need').html(htmldata);
            $('#agreement').html(agreement)
        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    });
}
$('input[name=agree]').change(function () {
    if ($(this).is(':checked')) {
        $(this).val(1)
    } else {
        $(this).val(2)
    }
});

$('.big-shade').click(function () {
    e.stopPropagation()
    $('.big-shade').hide()
    $('.pop').slideToggle()
});

$("#regSub").click(function () {
    var doajax = 0;
    var regName = $('input[name=register-name]').val();
    doajax = checkAccount(regName, doajax);
    if (doajax == 1) {
        doajax = CheckPostData(doajax);
        if (doajax == 1) {
            var code = $('input[name=code]').val();
            var codestr;
            var regPass = $('input[name=register-password]').val();
            var regPass1 = $('input[name=password1]').val();
            var agree = $('input[name=agree]').val();
            var intro = $('input[name=introducer_member]').val();
            if (code) {
                codestr = _Code
            } else {
                codestr = '';
                code='';
            }
            if (agree == 2) {
                $.toast('请仔细阅读注册协议，并同意', 'text');
                return false;
            }
            var pass_str1=$(".pass-str1").text();
            // console.log(pass_str1);
            if (pass_str1=="弱") {
                $.toast('密码过弱，请输入由英文字母和数字组合的密码', 'text');
                return
            }
            // return
            $.ajax({
                url: '/m/register/reg',
                type: 'post',
                contentType: "application/json",
                dataType: 'json',
                async: false,
                data: JSON.stringify({
                    introducer_member: intro,
                    account: regName,
                    password: regPass,
                    confirm_password: regPass1,
                    isAgreeDeal: agree,
                    code: code
                }),
                headers: {
                    'Code': codestr,
                    'platform': 'wap'
                },
                cache: false,
                timeout: 30000,
                beforeSend: function () {
                    Loading();
                    $(this).attr({disabled: "disabled"});
                },
                success: function (data) {
                    // console.log(data);
                    if (data.code) {
                        $.toast(data.msg);
                    } else {
                        if (data.data) {
                            setCookie("loginBack", data['data'].token);    //将登陆信息写入缓存
                            $.toast('注册成功',function () {
                                window.location.href = '/m/account';
                            });
                            
                            return;
                        } else {
                            $.toast('注册失败');
                        }
                    }
                },
                complete: function () {
                    LoadingClose();
                    $(this).removeAttr("disabled");
                }
            });
        }
    }
});

$('.m-login .o2 dt .consent a').click(function () {
    $('.big-shade').show(0)
    $('.pop').show(0)
});

// $('.h-head .iconfont').click(function () {
//     window.history.go(-1);
// });

$('#register-name').change(function () {
    var account = $(this).val();
    var doajax = 0;
    if(account==''){
        $('#register-name').parent().next('p').html('账号只能为4-11位数字和字母组合');
        return false;
    }
    doajax = checkAccount(account, doajax);
    if (doajax == 1) {
        $.ajax({
            url: '/wap/ajax/reg/repeat',
            type: 'get',
            async: false,
            headers: {
                'platform': 'wap'
            },
            data: {ajax: 'CheckUser', account: account},
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                // Loading();
            },
            success: function (data, info, xhr) {
                if (data.code) {
                    $('#register-name').parent().next('p').html(data.msg);
                    return false;
                }
                if (data.data === false) {
                    $('#register-name').parent().next('p').html('账号可用');
                    $('#register-name').parent().next('p').css('color', 'green');
                }
            },
            complete: function () {
                $(this).removeAttr("disabled");
                // LoadingClose();
            }
        })
    }

});
$('#register-password').change(function () {
    var nameVlue = $(this).val();
    var account=$('#register-name').val();
    if(nameVlue==account){
        $.toast('密码不能和账号相同','text');
        $('#register-password').parent().next('p').html('*密码不能和账号相同');
        return false;
    }
    var tel = /^[A-Za-z0-9]{6,11}$/;
    if (!tel.test(nameVlue)) {
        $.toast('密码只能为数字和字母组合的6-11位','text');
        return false;
    } else {
        if (parseInt(nameVlue.length) < 6 || parseInt(nameVlue.length) > 11) {
            $('.password-p').css('opacity', 1);
            return false;
        } else {
            $('.password-p').css('opacity', 0);
        }
    }

});
$('input[name=password1]').change(function () {
    determinePass();
})

function password(x) {
    var nameVlue = $('#' + x).val();
    if (parseInt(nameVlue.length) < 6 || parseInt(nameVlue.length) > 12) {
        $('#register-password').parent().next('p').html('* 密码长度请控制在6~12位数之间');
        $('#register-password').parent().next('p').css('color', 'red');
        $('.pass-str1').html("");
        $('.pass-str1').removeClass("curr");
        $(".strong").hide();
        return
    }
    var o = VerifPassword(nameVlue);
    $('#register-password').parent().next('p').html('');
    // console.log(o)
    if (o == 1) {
        $(".strong").show();
        $('.pass-str1').html("弱");
        $('.pass-str1').addClass('curr');

    } else if (o == 2) {
        $(".strong").show();
        $('.pass-str1').addClass('curr');
        $('.pass-str1').html("中");
    } else if (o >=4) {
        $(".strong").show();
        $('.pass-str1').addClass('curr');
        $('.pass-str1').html("强");
    }
}

function VerifPassword(string) {

    return checkStrong(string);

    //判断输入密码的类型  
    function CharMode(iN) {
        if (iN >= 48 && iN <= 57) //数字    
            return 1;
        if (iN >= 65 && iN <= 90) //大写    
            return 2;
        if (iN >= 97 && iN <= 122) //小写    
            return 4;
        else
            return 8;
    }

    //计算密码模式 
    function bitTotal(num) {
        modes = 0;
        for (i = 0; i < 4; i++) {
            if (num & 1) modes++;
            num >>>= 1;
        }
        return modes;
    };


    function checkStrong(sPW) {
        if (sPW.length < 6)
            return 0; //密码太短，不检测级别  
        Modes = 0;
        for (i = 0; i < sPW.length; i++) {
            //密码模式    
            Modes |= CharMode(sPW.charCodeAt(i));
        }
        return bitTotal(Modes);
    };

}

function determinePass() {

    var v1 = $("#register-password").val();
    var v2 = $("#register-password1").val();
    if (v1 == v2) {
        $('#register-password1').parent().next('p').html('');
    } else {
        $('#register-password1').parent().next('p').html('*密码不一致');
        $('#register-password1').parent().next('p').css('color', 'red');
        return 0;
    }
}

function checkAccount(account, doajax) {
    var msg = '';

    if (account.length < 4) {
        msg = '账号不小于4位，为数字和字母组合';
    } else if (account.length > 11) {
        msg = '账号大小于11位，为数字和字母组合';
    } else if (account.length == 0) {
        msg = '账号不能为空';
    }
    var tel = /^[A-Za-z0-9]{4,11}$/;
    if (!tel.test(account)) {
        $.toast('账号只能为4-11位数字和字母组合', 'text');
    } else {
        if (msg != '') {
            doajax = 0;
            $.toast(msg, 'text');
            $(this).val('');
            return false;
        } else {
            doajax = 1;
        }
    }
    return doajax;
}

function CheckPostData(doajax) {
    var v1 = $("#register-password").val();
    var v2 = $("#register-password1").val();
    var account = $('#register-name').val();
    var tel = /^[A-Za-z0-9]{6,11}$/;
    var msg = '';
    if (!tel.test(v1)) {
        doajax = 0;
        $.toast('密码不符合规则', 'text');
        return doajax;
    }
    if (!tel.test(v2)) {
        doajax = 0;
        $.toast('重复密码不符合规则', 'text');
        return doajax;
    }
    if (parseInt(v1.length) < 6 || parseInt(v1.length) > 11) {
        msg = '密码长度为6-11位数字和字母组合';
    }
    if (parseInt(v2.length) < 6 || parseInt(v2.length) > 11) {
        msg = '密码长度为6-11位数字和字母组合';
    }
    if (v1 != v2) {
        msg = '两次密码不一致';
    }
    if (v1 == '') {
        msg = '密码不能为空';

    }
    if (v2 == '') {
        msg = '请确认密码';
    }
    if (account == '') {
        msg = '请输入账号';
    }
    if (msg != '') {
        $.toast(msg, 'text');
        doajax = 0;
    } else {
        doajax = 1;
    }
    return doajax;
}

 var mescroll = new MeScroll("protocol", {
    down:{
        use: false, //是否初始化下拉刷新; 默认true
        auto: false, //是否在初始化完毕之后自动执行下拉回调callback; 默认true
        hardwareClass: "mescroll-hardware", //硬件加速样式;解决iOS下拉因隐藏进度条而闪屏的问题,参见mescroll.css
        isBoth: false, //下拉刷新时,如果滑动到列表底部是否可以同时触发上拉加载;默认false,两者不可同时触发;
    }
 })