$(function(){
    if(loginData){
        Loading();
        alert('已登陆账户不能进行注册！');
        window.location.href='/index';
    }
})

var LANGX = 'gb', H = '75', ajaxDoing = false, query = {};
$.extend({
    /**/
    _Dialog: function (Title, Data, Width, Height) {
        $('#Dialog').dialog({
            'title': Title,
            'width': Width,
            'minWidth': Width,
            'height': Height,
            'minHeight': Height,
            'modal': true,
            'bgiframe': true,
            'show': 'blind',
            'hide': 'blind'
        });
        $('#Dialog').html(Data);
    },
    /**/
    _BuildPrompt: function (Name, PromptText, o) {
        var options = {
            'showArrow': true,
            'positionType': 'topRight',
            'width': 100,
            'top': -28,
            'left': -27,
            'opacity': 0.8,
            'AMarginLeft': 13
        };
        options = $.extend(options, o);

        if (LANGX != 'gb' && LANGX != 'big5')
            options.width += 50;

        /**/
        var prompt = $('<div>');
        prompt.addClass("FormError");
        /**/
        var promptContent = $('<div>').addClass("FormErrorC").css('width', options.width).html(PromptText).appendTo(prompt);
        /**/
        if (options.showArrow) {
            var arrow = $('<div>').addClass("FormErrorA").css('marginLeft', options.AMarginLeft);
            switch (options.positionType) {
                case "bottomLeft":
                case "bottomRight":
                    prompt.find(".FormErrorC").before(arrow);
                    arrow.addClass("FormErrorABottom").html('<div style="width:1px;border:none;background: #DDDDDD;"><!-- --></div><div style="width:3px;border:none;background:#DDDDDD;"><!-- --></div><div style="width:1px;border-left:2px solid #DDDDDD;border-right:2px solid #ddd;border-bottom:0 solid #ddd;"><!-- --></div><div style="width:3px;"><!-- --></div><div style="width:5px;"><!-- --></div><div style="width:7px;"><!-- --></div><div style="width:9px;"><!-- --></div><div style="width:11px;"><!-- --></div><div style="width:13px;border:none;"><!-- --></div><div style="width:15px;border:none;"><!-- --></div>');
                    break;
                case "topLeft":
                case "topRight":
                    arrow.html('<div style="width:15px;border:none;"><!-- --></div><div style="width:13px;border:none;"><!-- --></div><div style="width:11px;"><!-- --></div><div style="width:9px;"><!-- --></div><div style="width:7px;"><!-- --></div><div style="width:5px;"><!-- --></div><div style="width:3px;"><!-- --></div><div style="width:1px;border-left:2px solid #ddd;border-right:2px solid #ddd;border-bottom:0 solid #DDDDDD;"><!-- --></div><div style="width:3px;border:none;background:#DDDDDD;"><!-- --></div><div style="width:1px;border:none;background: #DDDDDD;"><!-- --></div>');
                    prompt.append(arrow);
                    break;
            }
        }
        /**/
        prompt.css({
            "top": options.top,
            "left": options.left,
            "opacity": options.opacity
        });
        return $('<span>').addClass("Error_" + Name).css('position', 'relative').css('vertical-align', 'top').append(prompt.css('position', 'absolute'));
    },
    /**/
    _CheckIDCard: function (num) {
        var len = num.length, re;
        if (len == 15)
            re = new RegExp(/^(\d{6})()?(\d{2})(\d{2})(\d{2})(\d{3})$/);
        else if (len == 18)
            re = new RegExp(/^(\d{6})()?(\d{4})(\d{2})(\d{2})(\d{3})(\d)$/);
        else {
            return false;
        }
        var a = num.match(re);
        if (a != null) {
            if (len == 15) {
                var D = new Date("19" + a[3] + "/" + a[4] + "/" + a[5]);
                var B = D.getYear() == a[3] && (D.getMonth() + 1) == a[4] && D.getDate() == a[5];
            } else {
                var D = new Date(a[3] + "/" + a[4] + "/" + a[5]);
                var B = D.getFullYear() == a[3] && (D.getMonth() + 1) == a[4] && D.getDate() == a[5];
            }
            if (!B) return false;
        }
        return true;
    },
    /**/
    _InArray: function (stringToSearch, arrayToSearch) {
        for (s = 0; s < arrayToSearch.length; s++) {
            thisEntry = arrayToSearch[s].toString();
            if (thisEntry == stringToSearch) {
                return true;
            }
        }
        return false;
    }
});
// $(function () {
//     var mbodyheight = document.body.scrollHeight;
//     window.parent.document.getElementsByName("mbody")[0].style.height = mbodyheight + 'px';
//     console.log(mbodyheight);
//     console.log(window.parent.document.getElementsByName("mbody")[0].style.height);
// });




// 防被砍
if (top.location.hostname != self.location.hostname) {
    location = '/';
}
$(".password_adv").passStrength({
    userid: "#zcname",
    shortPass_txt: '密码强度：太短',
    badPass_txt: '密码强度：弱',
    goodPass_txt: '密码强度：很好',
    strongPass_txt: '密码强度：强',
    samePassword_txt: '帐号与密码不能相同'
});

$("#zcname").keyup(function () {

    var zh_pattern = /[a-z0-9]/g;
    if (!zh_pattern.test($(this).val())) {
        $("#zcname").val('');
    }
    /**/


})

$("#zcname").focus(function () {
    var vdefault = this.value;   //获得焦点时，如果值为默认值，则设置为空
    if (vdefault != "") {
        value = $('div').is('.FormErrorC');
        if (value) {
            $("#zcname").val("");
        }
    }
});
//账号只能输入数字和小写字母
jQuery.validator.addMethod("chenckzcname", function (value, element) {
    var tel = /[^0-9a-z]/g;
    return this.optional(element) || (!tel.test(value));
}, "只能输入数字和小写字母！");

//密码只能输入数字和小写字母
jQuery.validator.addMethod("chenckpsw", function (value, element) {
    var tel = /[^0-9a-z]/g;
    return this.optional(element) || (!tel.test(value));
}, "只能输入数字和小写字母！");



/**/
$.validator.addMethod('ajaxCheckData', function (value, element) {
    var id = element.id;
    if ('undefined' != typeof (query[id]) && query[id].value == value) {
        return true;
    }
    //'real_name': {'ajax' : 'checkdata', 'a_real_name': value},
    var data = {
        'zcname': {'ajax': 'CheckUser', 'account': value},
        'zcyzm': {'ajax': 'CheckCode', 'code': value},
        'zcturename': {'ajax': 'CheckRealName', 'real_name': value}
    };

    query[id] = {'status': false, 'value': value};
    ajaxDoing = true;
    $(element).siblings('.loading_pic').fadeIn();
    $.ajax({
        'url': '/ajax/register',
        'type': 'get',
        'data': data[id],
        headers: {
            'Code': _Code,
            'platform': 'pc'
        },
        'cache': false,
        'timeout': 30000,
        'error': function (e, textStatus) {
            if (textStatus == 'timeout') {
                alert('网路品质不佳!!');
                $(element).siblings('.loading_pic').fadeOut();
            }
        },
        'success': function (data) {
            console.log(data)
           if(data.code){
               var status = false;
            } else {
               status = true;
            }
            //var status = (data == '0') ? false : true;
            //var status = true;
            query[id].status = status;
            showError(id, status);
            ajaxDoing = false;
            $(element).siblings('.loading_pic').fadeOut();
        }
    });
    return true;
}, '');

/**/
$.validator.addMethod('CheckNameRule', function (value, element) {
    var Ch = /^[\u4e00-\u9fa5]+$/;
    var KRW = /^([\uAC00-\uD7AF])*$/gi;

    var En = /^([a-zA-Z]+)$/;

    // 韓幣特例
    var currency = 'RMB'; // $('#myFORM [name=currency]').val()
    if (currency == 'KRW') {
        return this.optional(element) || (Ch.test(value)) || (En.test(value)) || (KRW.test(value));
    } else {
        return this.optional(element) || (Ch.test(value)) || (En.test(value));
    }
}, '取款密码!!');

/**/
$.validator.addMethod('equalToUsername', function (value, element) {
    return (value == $('#zcname').val()) ? false : true;
}, '帐号与密码不能相同');


/**/
$.validator.addMethod('CheckPWDStrength', function (value, element) {
    return ($.fn.checkstrength(value) < 26) ? false : true;
}, '密码强度：弱');


/**/
$.validator.addMethod('CheckrmNum', function (value, element) {
    return value != '请点击' && value;
}, '验证码请务必输入!!');

/**/
var validator = $("#myFORM").validate({
    'onkeyup': false,
    'focusCleanup': true,
    'focusInvalid': false,
    'errorElement': 'span',
    /**/
    'rules': {
        'zcname': {
            'required': true,
            'minlength': 4,
            'ajaxCheckData': true,
            'chenckzcname': true,
            'maxlength': 11,
        },
        'zcpwd1': {
            'required': true,
            'minlength': 6,
            'chenckpsw': true,
            'equalToUsername': true,
            'CheckPWDStrength': true
        },
        'zcpwd2': {
            'required': true,
            'equalTo': '#zcpwd1'
        },
        'agree': {'required': true}
        ,
        'zcyzm': {
            'required': true,
            'ajaxCheckData': true,
            'chenckzcname': true
        }
    },
    /**/
    'messages': {
        'zcname': {
            'required': $._BuildPrompt('zcname', '✖ 请输入帐号!!'),
            'minlength': $._BuildPrompt('zcname', '✖ 帐号：请输入4-11个字元, 仅可输入英文字母以及数字的组合!!', {'width': 360}),
            'ajaxCheckData': $._BuildPrompt('zcname', '✖ 此帐号已经有人使用了！', {'width': 160}),
            'chenckzcname': $._BuildPrompt('zcname', '✖ 只能输入数字和小写字母！！', {'width': 180}),
        },
        'zcturename': {
            'required': $._BuildPrompt('zcturename', '✖ 请输入真实姓名!!', {'width': 160}),
            'CheckChinaName': $._BuildPrompt('zcturename', '✖ 请输入真实姓名!!', {'width': 160}),
            'maxlength': $._BuildPrompt('zcturename', '✖ 姓名：最多输入10个汉字!!', {'width': 160})
        },
        'zcpwd1': {
            'required': $._BuildPrompt('zcpwd1', '✖ 请输入密码!!'),
            'chenckpsw': $._BuildPrompt('zcpwd1', '✖ 密码：请输入6~11位英文或数字且符合0~9或a~z字元!', {'width': 310}),
            'minlength': $._BuildPrompt('zcpwd1', '✖ 请输入6-11个字元的密码!!', {'width': 170}),
            'equalToUsername': $._BuildPrompt('zcpwd1', '✖ 帐号与密码不能相同!!', {'width': 170}),
            'CheckPWDStrength': $._BuildPrompt('zcpwd1', '✖ 密码强度：弱')
        },

        'zcpwd2': {
            'required': $._BuildPrompt('zcpwd2', '✖ 确认密码!!'),
            'equalTo': $._BuildPrompt('zcpwd2', '✖ 确认密码错误！请重新输入!!', {'width': 140})
        },
        'agree': {
            'required': $._BuildPrompt('agree', '✖ 请勾选同意条款!!', {'width': 120, 'top': -60, 'left': -28})
        },
        'zcyzm': {
            'required': $._BuildPrompt('zcyzm', '✖ 请输入验证码!!', {'width': 120}),
            'ajaxCheckData': $._BuildPrompt('zcyzm', '✖ 验证码错误！', {'width': 160}),
            'chenckzcname': $._BuildPrompt('zcyzm', '✖ 只能输入数字和小写字母！', {'width': 160})
        }
    },
    /**/
    'submitHandler': function () {
        for (var i in query) {
            if (!query[i]['status']) {
                return false;
            }
        }
        if (confirm("是否确定写入?")) {
            $.ajax({
                    url: '/register',
                    type: 'post',
                    headers: {
                        'Code': _Code,
                        'platform': 'pc'
                    },
                    contentType: "application/json",
                    data: JSON.stringify({
                        introducer_member: $('input[name=introducer_member]').val(),
                        account: $('input[name=zcname]').val(),
                        password: $('input[name=zcpwd1]').val(),
                        confirm_password: $('input[name=zcpwd2]').val(),
                        code: $('input[name=zcyzm]').val(),
                        isAgreeDeal: $('input[name=agree]').val()
                    }),
                    dataType: 'json',
                    timeout: 30000,
                    error: function (e, textStatus) {
                        if (textStatus == 'timeout') {
                            alert('网路品质不佳!!');
                            $(element).siblings('.loading_pic').fadeOut();
                        }
                    },
                    success: function (data) {
                        //Todo:需要加载动画
                        console.log(data)
                        if (data){
                            if (data.data) {
                                setCookie("loginBack", data['data'].token);    //将登陆信息写入缓存
                                window.location.href = '/login/info';
                                return;
                            }
                        }else{
                            alert(data.msg)
                        }
                    }
                }
            )
        }
    }
});



/**/
$('input[name=password], input[name=passwd]').keyup(function () {
    this.value = this.value.replace(/[^a-z0-9]/g, '');
});

/**/
$('#zcname, #zcturename, #zcyzm').focus(function () {
    showError(this.id, true);
}).blur(function () {
    var id = this.id;
    if ('undefined' != typeof (query[id]) && !query[id].status && !ajaxDoing && $('.Error_' + id + ':visible').length == 0) {
        showError(id, false);
    }
})


function showError(name, status) {
    var errMes = $('.Error_ajax_' + name);
    if (status) {
        errMes.hide();
    } else {
        if (errMes.length > 0) {
            errMes.show();
            return;
        }
        switch (name) {
            case 'zcname':
                $('#zcname').after($._BuildPrompt('ajax_zcname', '✖ 此帐号已经有人使用了！', {'showArrow': false, 'width': 160}));
                break;
            case 'zcturename':
                $('#zcturename').after($._BuildPrompt('ajax_zcturename', '✖ 真实姓名重复！', {
                    'showArrow': false,
                    'width': 150
                }));
                break;
            case 'real_name':
                $('#real_name').after($._BuildPrompt('ajax_real_name', '✖ 姓名已注册, 请洽客服人员!!', {'width': 180}));
                break;
            case 'zcyzm':
                $('#zcyzm').after($._BuildPrompt('ajax_zcyzm', '✖ 验证码错误!', {'width': 90, 'top': -25, 'left': -25}));
                break;
        }
    }
}



