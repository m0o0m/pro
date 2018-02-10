//登陆判断
function loginId() {
    var loginData;
    if (getCookie('loginBack')){
        $.ajax({
            type: 'GET',
            url: '/ajax/login/in',
            data: {},
            async: false,
            dataType: 'json',
            success: function (msg) {
                if (msg['data']) {
                    loginData = msg['data'];
                }else{
                    delCookie("loginBack");
                }
            },
            error: function(){
                delCookie("loginBack");
            },
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            }
        });
    }else{
        delCookie("loginBack");
    }
    return loginData;
}

var loginData = loginId();
$(function(){
    userInfo(loginData);
})

function userInfo(loginData) {
    if ( getCookie('loginBack') ) {
        $(".pk-login-before").hide();
        $(".pk-login-after").show();

        $("div.pk-login-after div.ele-acc-name strong").html(loginData["account"]);
        $("div.pk-login-after .ele-first-balance strong").html(loginData["balance"]);
        var obalanceStr = '';
        for (var k in loginData.TBalance) {
            var v = loginData.TBalance[k];
            if (v["name"] == '账户总余额' || v["name"] == '账户余额') {
                obalanceStr += '<div class="ele-obalance"><span>' + v["name"].toUpperCase() + '：</span><strong>' + v["balance"] + '</strong></div>';
            }else{
                obalanceStr += '<div class="ele-obalance"><span>' + v["name"].toUpperCase() + '余额：</span><strong>' + v["balance"] + '</strong></div>';
            }
        }
        $("div.pk-login-after div.ele-obalance-item").html(obalanceStr);
        $("div.pk-login-after span.MsgNotReadCount").html(loginData["count"]);
    } else {
        $(".pk-login-before").show();
        $(".pk-login-after").hide();
    }
}


/***********************************会员中心 我的账号******************************/
//获取个人信息
function Info(){
    var infoData;
    if (getCookie('loginBack')){
        $.ajax({
            type: 'GET',
            url: '/base/info/info',
            data: {},
            async: false,
            dataType: 'json',
            success: function (msg) {
                if (!msg['code']) {
                    var sdata = msg.data;
                    infoData = sdata;
                    var mobileObj = $("div.one-right-bottom .right-b-l .icon-shouji").siblings('input');
                    var emailObj = $("div.one-right-bottom .right-b-l .icon-xiaoxi").siblings('input');
                    var birthdayObj = $("div.one-right-bottom .right-b-l .icon-liwu").siblings('input');
                    if(sdata.mobile != ''){
                        mobileObj.val(sdata.mobile);
                    }else{
                        mobileObj.val("请尽快添加手机号信息");
                    }
                    if(sdata.email != ''){
                        emailObj.val(sdata.email);
                    }else{
                        emailObj.val("请尽快添加邮箱信息");
                    }
                    if(sdata.birthday != ''){
                        birthdayObj.val(date("Y-m-d",sdata.birthday));
                    }else{
                        birthdayObj.val("请尽快添加生日信息");
                    }

                    $("div.one-right-top p").find('span').eq(1).html(date("Y-m-d h:i:s",sdata.create_time)); 
                    
                    if (sdata.last_login_time != '') {
                        $("div.one-right-top p").find('span').eq(3).html(date("Y-m-d h:i:s",sdata.last_login_time));
                    }else{   
                        $("div.one-right-top p").find('span').eq(3).html(date("Y-m-d h:i:s",sdata.create_time));
                    }

                    $("div.one-top-left div span").eq(1).html(sdata.realname);
                    var topLeftP = $("div.one-top-left p");
                    topLeftP.eq(0).find('span').html(sdata.account);
                    topLeftP.eq(1).find('span').html(sdata.balance);

                }else{
                    alert(msg.Msg)
                }
            },
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            }
        });
    }
    return infoData;
}

//获取今日交易信息
function InfoRecord(){
    if (getCookie('loginBack')){
        var timeNum = Date.parse(new Date())/1000; //因为jquery的时间戳是充1970年开始的14位数的时间戳 所以要除以1000
        $.ajax({
            type: 'GET',
            url: '/base/info/record',
            data: {'start_time':date('Y-m-d',timeNum),'end_time':date('Y-m-d',timeNum)},
            async: false,
            dataType: 'json',
            success: function (msg) {
                if (!msg['code']) {
                    var sdata = msg.data;
                    var htmlStr = '';
                    if (sdata) {
                        for (var k in sdata) {
                            var v = sdata[k];
                            htmlStr += '<li>'+
                                '<p class="fl sbxl">'+date('Y-m-d', v.create_time)+'<br/>'+date('H:i:s', v.create_time)
                                '</p>'+
                                '<div class="one-bottom-hl fl font-size-red">'+v.balance+'</div>'+
                                '<div class="one-bottom-hl fl">'+sourceType(v.source_type)+'</div>'+
                                '<div class="one-bottom-hl fl">'+types(v.types)+'</div>'+
                                '<div class="one-bottom-hl fl">'+v.after_balance+'</div>'+
                                '<p class="fl one-bottom-hr">'+v.remark+
                                '</p></li>';
                        }
                    }else{
                        htmlStr += '<li style="text-align:center;"><span>暂无数据</span></li>';
                    }
                    
                    $("div.one-bottom-c-e div.one-bottom-c").eq(0).find('ul.one-bottom-f').html(htmlStr);
                }else{
                    alert(msg.msg)
                    return false;
                }
            },
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            }
        });
    }
}

$(function(){
    $("div.my-dialog").on('click', '.dialog-submit', function(){
        if(getCookie('loginBack')){
            var $formObj = $(this).parent('.dialog-bottom-c');
            var $titleName = $formObj.parent("div")[0].className;

            //修改密码
            if ($titleName == 'one-pass-dialog'){
                $passType = $formObj.find("select[name='select-passward']").val();
                if($passType == 'userpwd' || $passType == 'moneypwd'){
                    if($passType == 'userpwd'){
                        $types = 1;
                    }else if ($passType == 'moneypwd'){
                        $types = 2;
                    }
                    $beforePassword = $formObj.find("input[name='beforePassword']").val();
                    $password = $formObj.find("input[name='password']").val();
                    $confirmPassword = $formObj.find("input[name='confirmPassword']").val();
                    if($beforePassword.length == 0){
                        alert("原密码不能为空！");
                        return false;
                    }else{
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
                    if($password.length == 0){
                        alert("新密码不能为空！");
                        return false;
                    }else{
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
                            alert("原密码和新密码不能想同");
                            return false;
                        }
                    }

                    if($confirmPassword != $password){
                        alert("两次输入的新密码不一致");
                        return false;
                    }
                    var $success = function (msg) {
                            if (msg) {
                                alert(msg.msg);
                                $("input[name='beforePassword']").select();
                                $("input[name='password']").select();
                                $("input[name='confirmPassword']").select();
                            } else {
                                if ($types == 1) {
                                    delCookie("loginBack");
                                    alert("密码修改成功, 请重新登录！");
                                    setTimeout("window.location.href = '/m/login'; ",1000);//延迟1秒执行
                                }else{
                                    alert("取款密码修改成功");
                                }
                            }
                            LoadingClose();
                        };
                    var data = JSON.stringify({'types':$types,'beforePassword':$beforePassword,'password':$password,'confirmPassword':$confirmPassword});
                    $('.my-dialog').hide();
                    ajaxObj('PUT', "/member/password", $success, data);
                }
            }

            //修改会员资料
            if ($titleName == 'one-info-dialog') {
                var $birthday_num = $formObj.find("input[name='birthday_num']").val();
                var $email_num = $formObj.find("input[name='email_num']").val();
                var $phone_num = $formObj.find("input[name='phone_num']").val();
                var $qq_num = $formObj.find("input[name='qq_num']").val();
                var $wechat = $formObj.find("input[name='wechat']").val();
                var $card = $formObj.find("input[name='card']").val();
                var $local_code = '+86';
                var $remark = $formObj.find("textarea[name='remark']").val();
                var $realname = $formObj.find("input[name='realname']").val();
                if (loginData.realname == "") {
                    if(!(/^([\u4E00-\uFA29]|[\uE7C7-\uE7F3])*$/).test($realname)){
                        alert('真实姓名格式不正确');return false;
                    }
                }
                if($birthday_num != ''){
                    if($birthday_num.length >10){
                        alert('会员生日格式不正确');return false;
                    }
                }
                if ($email_num != '') {
                    var regEmail = /^([a-zA-Z0-9_.-]+)@([a-z0-9_.-]+).([a-z.]{2,6})$/;
                    if (!regEmail.test($email_num)) {
                        alert('邮箱格式不正确！');return false;
                    }
                }
                if ($phone_num != '') {
                    if (!(/(1[3-9]\d{9}$)/).test($phone_num)) {
                        alert('手机格式不正确!');return false;
                    }
                }
                if ($qq_num != '') {
                    if (!(/^[1-9][0-9]{4,}$/).test($qq_num)) {
                        alert('QQ格式不正确');return false;
                    }
                }
                if ($wechat != '') {
                    if ($wechat.length > 20) {
                        alert('微信账号长度太长,请不要超过20！');return false;
                    }
                }
                if ($card != '') {
                    if (!(/^([0-9]){17}([0-9]|X)$/).test($card)) {
                        alert('身份证号格式不正确！');return false;
                    }
                }

                if ( ($birthday_num+$email_num+$phone_num+$qq_num+$wechat+$card+$remark).length == 0 ) {
                    return false;
                }
                var data = JSON.stringify({
                    'realname':$realname,
                    'birthday_num':$birthday_num,
                    'email_num':$email_num,
                    'phone_num':$phone_num,
                    'qq_num':$qq_num,
                    'wechat':$wechat,
                    'card':$card,
                    'local_code':$local_code,
                    'remark':$remark
                });
                var $success = function(msg){
                    if (msg) {
                        alert(msg.msg);
                        LoadingClose();
                        return false;
                    } else {
                        alert('会员资料修改成功！');
                        if (loginData.realname == '') {
                            loginData = loginId();
                        }
                        Info();
                        LoadingClose();
                        return
                    }
                };
                $('.my-dialog').hide();
                ajaxObj('PUT', "/base/info/means", $success, data);
            }

            //添加出款银行
            if ($titleName == 'one-bank-dialog') {
                var $bankId = Number($formObj.find("select[name='bank']").val());
                var $card = $formObj.find("input[name='card']").val();
                var $cardName = $formObj.find("input[name='cardName']").val();
                var $cardAddress = $formObj.find("input[name='cardAddress']").val();
                if ($bankId == '') {
                    alert('请选择出款银行');
                    return
                }
                if ($card == '') {
                    alert('银行卡号不能为空');
                    return
                }else if( !(/^\d{16,19}$/).test($card) ){
                    alert('银行卡号格式不正确， 应该由16-19位的数字组成');
                    return
                }
                if ($cardName == '') {
                    alert('开户人不能为空')
                    return
                }
                if ($cardAddress == '') {
                    alert('开户地址不能为空');
                    return
                }

                if ($("ul.bank-ul li.iconjia").length == 3) {
                    var $password = $formObj.find("input[name='password']").val();
                    var $comPassword = $formObj.find("input[name='comPassword']").val();
                    if($password.length == 0){
                        alert("取款密码不能为空！");
                        return false;
                    }else{
                        if ( !(/^\d{4}$/).test($password) ) {
                            alert("取款密码为4位数的数字");
                            return false;
                        }
                    }

                    if($comPassword != $password){
                        alert("两次输入的密码不一致");
                        return false;
                    }
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
                var $success = function (data) {
                    if (data) {
                        if(data.code){
                            alert(data.msg);
                            LoadingClose();
                            return;
                        }
                    }else{
                        alert("添加成功");
                        bankPaymentList();
                        if (loginData.realname == '') {
                            loginData = loginId();
                        }
                        LoadingClose();
                        return;
                    }
                }
                $('.my-dialog').hide();
                ajaxObj('POST', "/base/info/add", $success, data);
            }
        }
    });
    
})


/***********************************END 会员中心 我的账号******************************/

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
            },   D: function () {
                return f.l().substr(0, 3)
            },   j: function () {
                return jsdate.getDate()
            },   l: function () {
                return txt_weekdays[f.w()]
            },   N: function () {
                return f.w() + 1
            },   S: function () {
                return txt_ordin[f.j()] ? txt_ordin[f.j()] : 'th'
            },   w: function () {
                return jsdate.getDay()
            },   z: function () {
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
            },   m: function () {
                return pad(f.n(), 2)
            },   M: function () {
                return f.F().substr(0, 3)
            },   n: function () {
                return jsdate.getMonth() + 1
            },   t: function () {   
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
            },   y: function () {
                return (jsdate.getFullYear() + "").slice(2)
            },       // Time 
              a: function () {
                return jsdate.getHours() > 11 ? "pm" : "am"
            },   A: function () {
                return f.a().toUpperCase()
            },   B: function () {    // peter paul koch: 
                   
                var off = (jsdate.getTimezoneOffset() + 60) * 60;   
                var theSeconds = (jsdate.getHours() * 3600) + (jsdate.getMinutes() * 60) + jsdate.getSeconds() + off;   
                var beat = Math.floor(theSeconds / 86.4);   
                if (beat > 1000) beat -= 1000;   
                if (beat < 0) beat += 1000;   
                if ((String(beat)).length == 1) beat = "00" + beat;   
                if ((String(beat)).length == 2) beat = "0" + beat;   
                return beat;  
            },   g: function () {
                return jsdate.getHours() % 12 || 12
            },   G: function () {
                return jsdate.getHours()
            },   h: function () {
                return pad(f.g(), 2)
            },   H: function () {
                return pad(jsdate.getHours(), 2)
            },   i: function () {
                return pad(jsdate.getMinutes(), 2)
            },   s: function () {
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
            },   P: function () {
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

/***********************************会员出款银行******************************/
function bankPaymentList(){
    var $success = function(data, code){
        if (data) {
            if (data.code) {
                alert(data.msg);
                LoadingClose();
                return
            }
            var $sdata = data.data;
            var $htmlStr = '';
            if ($sdata){
                if($sdata.length < 3){
                    $htmlStr += '<li class="iconjia"><div><p class="icon-iconjia icon iconfont "><span>添加银行卡</span></p></div></li>';
                    if ($sdata.length < 2) {
                       $htmlStr += '<li class="iconjia"><div><p class="icon-iconjia icon iconfont "><span>添加银行卡</span></p></div></li>';
                    }
                }
                for(var k in $sdata){
                    var v = $sdata[k];
                    $htmlStr += '<li style="background:url('+cdnUrl+'/shared/sitepublic/images/bank/'+v.bank_id+'.png) 0% 0% / 100% 100%"><div>'+
                    '<p><span class="ng-binding">'+v.card+'</span></p>'+
                    '</div></li>'
                }
                $("div.one-bank-dialog div.margin-bank.pass").hide();
                $('.one-bank-dialog').find("input[name='cardName']").attr({'disabled':"disabled"}).val(loginData.realname);
            }else{
                $htmlStr += '<li class="iconjia"><div><p class="icon-iconjia icon iconfont "><span>添加银行卡</span></p></div></li>';
                $htmlStr = $htmlStr + $htmlStr + $htmlStr;
                $("div.one-bank-dialog div.margin-bank.pass").show();
            }
            
            $(".bank-ul").html($htmlStr);
            var $width1 = $(".bank-ul").width();
            var $width2 = $(".bank-ul li").width();
            var num = ($width1 - $width2)/2;
            $.each($(".bank-ul li"), function(index, e){
                $(this).css({'left':index*num});
            })

            //从取款页面跳转过来的直接弹出添加框
            var $pass = getQueryString('pass');
            if ($pass) {
                $(".bank-ul .icon-iconjia:eq(0)").click();
            }
        }else{
            alert('获取会员银行卡信息失败！');
        }
        LoadingClose();
    }
    ajaxObj('GET', "/base/info/list", $success);
}

$(function(){
    $('div.right-b-r').on('click','p.icon-iconjia',function(){
        var $layerObj = $('.my-dialog');
        var $obj = $('.one-bank-dialog');
        $obj.find("input[name='realname']").attr({'disabled':"disabled"}).val(loginData.realname);
        var $success = function(msg){
            if (msg) {
                if(msg.code){
                    alert(msg.msg);
                    return false;
                }
                var sdata = msg.data;
                var $htmlStr = '';
                for (var k in sdata) {
                    var v = sdata[k];
                    $htmlStr += '<option value="'+v.id+'">'+v.title+'</option>'
                }
                $obj.find("select[name='bank']").html($htmlStr);
            } else {
                alert('获取会员详细信息失败');
            }
            LoadingClose();
            return
        }

        ajaxObj('GET', "/base/addInfo", $success);
        $layerObj.show();
        $obj.show();
        $obj.siblings('div').hide();

        var $height1 = $layerObj.height();
        var $height2 = $obj.height();
        $('.my-dialog div.dialog-c').css({'margin-top':"100px"});
    })
})
/***********************************END会员出款银行******************************/



// 数据来源类型
function sourceType(type){
    switch(type){
        case 1:
            return '公司入款';
        case 2:
            return '线上入款';
        case 3:
            return '人工取出';
        case 4:
            return '线上取款';
        case 5:
            return '出款';
        case 6:
            return '注册优惠';
        case 7:
            return '下单';
        case 8:
            return '取消出款';
        default:
            return '人工存入';
    }
}

function types(type){
    switch(type){
        case 1:
            return '存入';
        case 2:
            return '取出';
    }
}

/** 
 * ajax 封装
 * @param {string} type ajaxtype类型 默认GET 
 * @param {string} _url ajax提交的路径
 * @return {function} _success   回调函数
 * @return {json} data 传递的参数
 */
function ajaxObj(type="GET", _url, _success, data={}){
    $.ajax({
        type: type,
        url: _url,
        data: data,
        async: false,
        dataType: 'json',
        beforeSend : function(){
            Loading();
        },
        success: _success,
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "pc"
        },
        error:function(){
            alert("请求失败");
            LoadingClose();
        }
    });
}

$(function(){
    $("i.icon-f14").on("click",function(){
        if (!$(this).hasClass('Refresh')) {
            var _this = $(this);
            if ($(this)) {}
            _this.addClass("iconRefresh");
            _this.addClass("Refresh");
            setTimeout(function () {
                _this.removeClass("iconRefresh");
                _this.removeClass("Refresh");

            }, 1500) 
        }
    })
})

$(function(){
    $("i.RefreshButton").on("click",function(){
        if (!$(this).hasClass('Refresh')) {
            var _this = $(this);
            if ($(this)) {}
            _this.addClass("iconRefresh");
            _this.addClass("Refresh");
            setTimeout(function () {
                _this.removeClass("iconRefresh");
                _this.removeClass("Refresh");

            }, 1500) 
        }
    })
})
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

$('#_bbsportBalance').click(function () {
    hideBalance($(this).find('strong'))
})

//获取url参数
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;

}