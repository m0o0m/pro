$(function () {
    withdrawWsService.Notify(function (data) {
        // alert(data);
        $.Pro(data, {
            Time: 2, StartOn: function () {
                $("html").click(function () {
                    $(".showAlert_Pro").hide()
                })
            }
        })
    })
    if (getCookie('loginBack')) {
        $.ajax({
            type: "get",
            url: "/ajax/draw/data",
            data: {},
            dataType: 'json',
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            beforeSend: function () {
                Loading();
            },
            success: function (data) {
                // console.log(data);
                if (data.data) {
                    $('.audit').show();
                    $('.mycenter').hide();
                    $('.smallamount').hide();
                    var showdata = data.data;
                    // console.log(showdata);
                    $('.round .round-2').click(function () {
                        var showTime = strToDate(showdata.CreateTime);
                        var outTime;

                        if($('.audit').attr('stat')==undefined||$('.audit').attr('stat')==0){
                            $('.audit').attr('stat',1);
                            if (showdata.OutTime != 0) {
                                outTime = strToDate(showdata.OutTime);
                            } else {
                                outTime = '';
                            }
                            var showmoney = Math.floor((showdata.FavourableMoney + showdata.ExpeneseMoney + showdata.Charge) * 100) / 100;
                            var  normal, mui,zhStr,yhStr;
                            if(showdata.FavourableMoney>0){
                                normal='不满足';
                                yhStr='不满足常态稽核，需扣除50%行政优惠'+showdata.FavourableMoney+'元';
                            }else{
                                normal='满足';
                                yhStr='满足常态稽核，无需扣除50%行政优惠';
                            }
                            if(showdata.ExpeneseMoney>0){
                                zhStr='不满足综合稽核，需扣除优惠金额'+showdata.ExpeneseMoney+'元';
                                mui='不满足';
                            }else{
                                mui='满足';
                                zhStr='满足综合稽核，无需扣除优惠金额';
                            }
                            var htmldata = '<div class="deposit-schedule-c">\n' +
                                '        <div class="center-3 clearfix">\n' +
                                '            <div class="center-3-t">出款常见问题\n' +
                                '                <i class="icon iconfont icon-wenhao"></i>\n' +
                                '            </div>\n' +
                                '            <h3>尊敬的会员，您好！出款时若出现以下情况将会扣除一定的费用</h3>\n' +
                                '            <ul class="fl fl-1 f1-font">\n' +
                                '                <li>1. 未满足常态稽核将扣除入款金额50%行政费用与优惠</li>\n' +
                                '                <li>2. 未满足综合稽核将会扣除优惠金额</li>\n' +
                                '                <li>3. 提交订单后注意查看稽核状态，请注意是否需要扣除费用</li>\n' +
                                '                <li>4. 有出款订单未处理时无法提交第二笔订单</li>\n' +
                                '            </ul>\n' +
                                '            <ul class="fl fl-2 f1-font" style="padding-left: 20px;">\n' +
                                '                <li class="myd-red">常态稽核=会员入款金额 * 常态稽核倍数</li>\n' +
                                '                <li class="myd-red">综合（优惠）稽核=（会员入款金额 + 优惠） * 综合稽核倍数</li>\n' +
                                '                <li class="myd-red">优惠即存款优惠，人工优惠等等</li>\n' +
                                '            </ul>\n' +
                                '        </div>\n' +
                                '        <div class="center-4-1">\n' +
                                '            <ul class="clearfix my-theme">';
                            htmldata+= '                <li class="fl">开始时间</li>\n' +
                                '                <li class="fl" style="line-height: 50px;">结束时间</li>\n' +
                                '                <li class="fl">存款金额</li>\n' +
                                '                <li class="fl">优惠稽核</li>\n' +
                                '                <li class="fl">综合稽核</li>\n' +
                                '                <li class="fl">手续费</li>\n' +
                                '                <li class="fl">常态稽核</li>\n' +
                                '                <li class="fl">综合稽核</li>\n' +
                                '            </ul>\n' +
                                '            <ul class="clearfix" id="show-charge-data">\n' +
                                '                <li class="fl">'+showTime[0]+showTime[1]+'</li>';
                            if(outTime!=''){
                                htmldata+= '                <li class="fl">'+outTime[0]+outTime[1]+'</li>';
                            }else{
                                htmldata+= '                <li class="fl">'+outTime+'</li>';
                            }
                            htmldata+=  '                <li class="fl">'+showdata.OutwardNum+'</li>\n' +
                                '                <li class="fl">'+showdata.FavourableMoney+'</li>\n' +
                                '                <li class="fl">'+showdata.ExpeneseMoney+'</li>\n' +
                                '                <li class="fl">'+showdata.Charge+'</li>\n' +
                                '                <li class="fl myd-red">'+normal+'</li>\n' +
                                '                <li class="fl myd-red">'+mui+'</li>\n' +
                                '            </ul>\n' +
                                '        </div>\n' +
                                '    </div>\n' +
                                '    <div class="center-5">\n' +
                                '        <div class="center-5-t">稽核详情</div>\n' +
                                '        <div class="center-5-c">\n' +
                                '            <span class="span-1"><span><i class=\'myd-red\'>*</i>订单号：</span><span class="myd-red" id="order-num">2017777777777777</span></span><br>\n' +
                                '            <span class="span-1"><span><i class=\'myd-red\'>*</i>取款金额:</span><span\n' +
                                '                    class="myd-red" id="draw-money">'+showdata.OutwardNum+'</span>元</span></span><br>\n' +
                                '            <span class="span-1"><span><i class=\'myd-red\'>*</i>扣除手续费：</span><span\n' +
                                '                    class="myd-red" id="deduction-money">'+showdata.Charge+'</span>元</span></span><br>\n' +
                                '            <span class="span-1"><span><i class=\'myd-red\'>*</i>最终取款金额：</span><span\n' +
                                '                    class="myd-red" id="draw-out-money">'+showdata.OutwardMoney+'</span>元</span></span><br>\n' +
                                '        </div>\n' +
                                '        <div class="center-5-b" id="show-result">\n' +
                                '            <span class="span-1"><i class=\'myd-red\'>*</i>'+zhStr+'</span><br>\n' +
                                '            <span class="span-1"><i class=\'myd-red\'>*</i>'+yhStr+'</span><br>\n' +
                                '            <span class="span-1"><i class=\'myd-red\'>*</i>综合稽核+常态稽核+手续费，共需扣除：<span class=\'myd-red\'>'+showmoney+'</span>元</span><br>\n' +
                                '        </div>\n' +
                                '    </div>\n' +
                                '    <div class="center-6" id="result-url"></div>';
                            $('.audit').append(htmldata);
                        }else{
                            $('.audit').attr('stat',0);
                            $('.audit').children('.deposit-schedule-c').remove();
                            $('.audit').children(' .center-5').remove();
                            $('.audit').children(' .center-6').remove();

                        }

                    })
                }
            },
            complete: function () {
                LoadingClose();
            }
        });
        var ajaxdo = 0;
        $.ajax({
            type: "get",
            url: "/ajax/get/drawList",
            data: {},
            dataType: 'json',
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {
                // console.log(data)
                var cardnum;
                if (data.code) {
                    LocationTips(data.msg);
                } else {
                    var sdata = data.data;
                    var realNameList = data.data.real_name;
                    var banklist = sdata.bank_list;
                    $("#name").html(loginData.account);
                    $("#money").html(loginData.balance);
                    $('#withdrawal-cap').html(sdata.poundage.outPoundageUp + '元');
                    if (banklist == null) {
                        setCookie("pass", 1);
                        LocationTips();
                        return false;
                    }
                    if (realNameList.draw_password == '') {
                        LocationTips('请设置取款密码,再进行操作');
                        return false;
                    }
                    var htmldata = '';
                    if (banklist.length != 0) {
                        if (banklist.length == 3) {//银行有三张的情况下不能添加
                            $('#page-body .myd-green').remove();
                        }
                        for (var i = 0; i < banklist.length; i++) {
                            htmldata += '  <option value="' + banklist[i].id + '">' + banklist[i].title + '</option>';
                        }
                        cardnum = banklist[0].card;
                        bankName = banklist[0].card_name;
                    } else {
                        cardnum = '未添加出款银行卡号';
                        htmldata = ' <option value="">添加卡号后再操作</option>';
                        LocationTips();
                        return false;
                    }
                }
                $('#bank-card').html(cardnum);
                $('#bank-name').html(bankName);
                $('#bank-list').html(htmldata);

            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        });
        $('select[name=bankList]').change(function () {
            var id = $('select[name=bankList] option:selected').val();
            $.ajax({
                type: "get",
                url: "/ajax/get/oneBank",
                data: {id: id},
                dataType: 'json',
                async: false,
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': "pc"
                },
                beforeSend: function () {
                    $(this).attr({disabled: "disabled"});
                    Loading();
                },
                success: function (data) {
                    console.log(data)
                    if (data.code) {
                        LocationTips(data.msg)
                    } else {
                        var getData = data.data;
                        $('#bank-card').html(getData.card);
                    }
                    // console.log(data)
                }, complete: function () {
                    $(this).removeAttr("disabled");
                    LoadingClose();
                }
            })
        })
        $('#fresh-money').click(function () {
            $.ajax({
                url: '/ajax/get/balance',
                type: 'get',
                data: {},
                async: false,
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': "pc"
                },
                beforeSend: function () {
                    $(this).attr({disabled: "disabled"});
                    Loading();
                },
                success: function (data) {
                    if (data.code) {
                        // alert(data.msg);
                        $.Pro(data.msg, {
                            Time: 2, StartOn: function () {
                                $("html").click(function () {
                                    $(".showAlert_Pro").hide()
                                })
                            }
                        })
                        return false;
                    }
                    if (data.data) {
                        $('#money').html(data.data)
                    }

                }, complete: function () {
                    $(this).removeAttr("disabled");
                    LoadingClose();
                }
            })
        });

        function checkDrawPassword(id, ajaxdo) {
            var drawPassword = $(id).val();
            var mes='';
            $(id).next('span').css('color','red');
            if(drawPassword==''){
                mes='*取款密码不能为空';
                $.Pro(mes, {
                    Time: 2, StartOn: function () {
                    }
                });
                $(id).next('span').html(mes);
                return false;
            }else if (drawPassword.length<4&&drawPassword.length>0){
                mes='*请输入4位取款密码';
                $.Pro(mes, {
                    Time: 2, StartOn: function () {
                    }
                });
                $(id).next('span').html(mes);
                return false;
            }else{
                $(id).next('span').html('');
            }
            $.ajax({
                url: '/ajax/check/drawPass',
                type: 'get',
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': "pc"
                },
                data: {draw_password: drawPassword},
                async: false,
                success: function (data, info, xhr) {
                    if (data) {
                        if (data.code) {
                            // alert(data.msg);
                            $.Pro(data.msg, {
                                Time: 2, StartOn: function () {
                                    $("html").click(function () {
                                        $(".showAlert_Pro").hide()
                                    })
                                }
                            })
                            ajaxdo = 0;
                        }
                    }
                    if (xhr.status == 204) {
                        ajaxdo = 1;
                    }
                }
            })
            return ajaxdo;
        }

        function postData() {
            var balance = $('input[name=drawMoney]').val();
            var password = $('input[name=drawPassword]').val();
            var bank_id = $("#bank-list option:selected").val();
            $.ajax({
                type: "post",
                url: "/member/draw/data",
                async: false,
                data: JSON.stringify({
                    money: balance,
                    drawPassword: password,
                    bankId: bank_id
                }),
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': "pc"
                },
                dataType: 'json',
                beforeSend: function () {
                    Loading();
                    $('#draw-submit').attr('disabled', 'disabled');
                },
                success: function (data, info, xhr) {
                    // console.log(data);
                    if (data) {
                        if (data.code) {
                            // alert(data.msg);
                            $.Pro(data.msg, {
                                Time: 2, StartOn: function () {
                                    $("html").click(function () {
                                        $(".showAlert_Pro").hide()
                                    })
                                }
                            })
                            return false;
                        }
                        if (data.data) {
                            $('.audit').hide();
                            $('.mycenter').hide();
                            $('.smallamount').show();
                            var showdata = data.data;
                            var showTime = strToDate(showdata.create_time);
                            var showmoney = Math.floor((showdata.admin_money + showdata.deposit_money + showdata.charge) * 100) / 100;
                            var htmldata, normal, mui;
                            var htmldata1 = '';
                            var showresult = '';
                            if (showdata.mui_status == 1) {
                                mui = '满足';
                                showresult += '<span class="span-1"><i class="myd-red">*</i>满足综合稽核，无需扣除优惠金额</span><br>';
                            } else {
                                mui = '不满足';
                                showresult += '<span class="span-1"><i class="myd-red">*</i>不满足综合稽核，扣除优惠金</span><br>';
                            }
                            if (showdata.normal_status == 1) {
                                normal = '满足';
                                showresult += '<span class="span-1"><i class="myd-red">*</i>满足常态稽核，无需扣除50%行政优惠</span><br>';
                            } else {
                                normal = '不满足';
                                showresult += '<span class="span-1"><i class="myd-red">*</i>不满足常态稽核，需扣除50%行政优惠</span><br>';
                            }
                            showresult += '<span class="span-1"><i class="myd-red">*</i>综合稽核+常态稽核+手续费，共需扣除：<span class="myd-red">' + showmoney + '</span>元</span><br>';
                            htmldata = '<li class="fl">' + showTime[0] + showTime[1] + '</li>' +
                                '<li class="fl"></li>' +
                                '<li class="fl">' + showdata.out_money + '</li>' +
                                '<li class="fl">' + showdata.admin_money + '</li>' +
                                '<li class="fl">' + showdata.deposit_money + '</li>' +
                                '<li class="fl">' + showdata.charge + '</li>' +
                                '<li class="fl">' + normal + '</li>' +
                                '<li class="fl">' + mui + '</li>';
                            $('#show-charge-data').html(htmldata);
                            var drawmoney = Math.floor(showdata.out_money * 100) / 100;
                            var outmoney = Math.floor(showdata.out_charge * 100) / 100;
                            $('#order-num').html(showdata.order_num);
                            $('#draw-money').html(drawmoney);
                            $('#deduction-money').html(showmoney);
                            $('#draw-out-money').html(outmoney);
                            $('#show-result').html(showresult);

                            if (data.data.out_status == 0 || showdata.out_charge <= 0) {
                                htmldata1 += ' <a href="/member/withdraw">返回取款页</a>';
                            } else {
                                htmldata1 += '<a href="/member/draw/write">继续出款</a>';
                            }
                            $('#result-url').html(htmldata1)
                        }
                    }
                }, complete: function () {
                    $('#draw-submit').removeAttr("disabled");
                    LoadingClose();
                }
            })
        }

        $('#bank-list').change(function () {
            var bankId = $('#bank-list option:selected').val();
            $.ajax({
                url: '/ajax/get/oneCard',
                type: 'get',
                async: false,
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': "pc"
                },
                data: {id: bankId},
                success: function (data) {
                    // console.log(data);
                    if (data.code) {
                        // alert(data.msg);
                        $.Pro(data.msg, {
                            Time: 2, StartOn: function () {
                            }
                        })
                    }
                    var datainfo = data.data;
                    if (datainfo) {
                        $('#bank-name').html(datainfo.card_name);
                        $('#bank-card').html(datainfo.card);
                    }
                }
            })
        });
    } else {
        LocationTips('用户登录信息已失效， 请重新登录！', '/index');
        return false;
    }


    function LocationTips(mes, url) {
        if (mes == '' || mes == undefined) {
            mes = '请完善个人取款资料再进行取款(真实姓名，出款银行，取款密码)！';
        }
        if (url == '' || url == undefined) {
            url = '/member/account';
        }
        // alert(mes);

        $.Pro(mes, {
            Time: 2, StartOn: function () {
                $("html").click(function () {
                    $(".showAlert_Pro").hide()
                })
            }
        })
        window.location.href = url;
    }

    $('input[name=drawMoney]').change(function () {
        var mes;
        mes = checkBalance(this);
        if (mes == 0) {

            return false;
        }
    });
    $('#draw-submit').click(function () {

        var ajaxdo = checkBalance('input[name=drawMoney]');
        if (ajaxdo == 1) {
            ajaxdo = checkDrawPassword('input[name=drawPassword]', ajaxdo);
        }
        if (ajaxdo == 1) {
            postData();
        }
    });

    function checkBalance(id) {
        var money = $(id).val();
        var balance = loginData.balance;
        var mes = '';
        if (money > balance) {
            mes = '取款金额大于账户余额';
        }
        if (money < 10&&money>0) {
            mes = '取款金额不得低于10元';
        }else if(money==''){
            mes='输入金额非法，请填写正确金额！';
        }
        if (mes != '') {
            $(id).val("");
            $(id).next('span').css('color','red');
            $(id).next('span').html(mes);
            // alert(mes);
            $.Pro(mes, {
                Time: 2, StartOn: function () {
                }
            })
            ajaxdo = 0;
            return false;
        } else {
            $(id).next('span').html('');
            ajaxdo = 1;
        }
        return ajaxdo;
    }


    $('#money').click(function () {
        hideBalance($(this));
    })

    function My_pc_Modal(obj) {
        var tet = obj.text;
        var time = obj.time || 2000;
        var otime = time - 500;
        var otime = (time - 500) <= 600 ? time : (time - 500);

        $(".my_pc_Modal").text(tet).show().css("opacity", '1');

        setTimeout(function () {
            $(".my_pc_Modal").animate({
                "opacity": "0"
            })
        }, otime)

        setTimeout(function () {
            $(".my_pc_Modal").text('').hide();
        }, time)
    }
});