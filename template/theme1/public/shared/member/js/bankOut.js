$(function () {

    var ajaxdo = 0;
    if (loginData) {
        $.ajax({
            type: "get",
            url: "/ajax/get/drawList",
            data: {},
            dataType: 'json',
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
                    var realNameList=data.data.real_name;
                    var banklist = sdata.bank_list;
                    $("#name").html(loginData.account);
                    $("#money").html(loginData.balance);
                    $('#withdrawal-cap').html(sdata.poundage.outPoundageUp+'元');
                    if(realNameList.realname==''){
                        LocationTips();
                        return false;
                    }
                    if(realNameList.draw_password==''){
                        LocationTips('请设置取款密码,再进行操作');
                        return false;
                    }
                    var htmldata = '';
                    if (banklist == null) {
                        cardnum = '未添加出款银行卡号';
                        htmldata = ' <option value="">添加卡号后再操作</option>';
                        LocationTips();
                        return false;
                    } else {
                        if (banklist.length != 0) {
                            for (var i = 0; i < banklist.length; i++) {
                                htmldata += '  <option value="' + banklist[i].id + '">' + banklist[i].title + '</option>';
                            }
                            cardnum = banklist[0].card
                        } else {
                            cardnum = '未添加出款银行卡号';
                            htmldata = ' <option value="">添加卡号后再操作</option>';
                            LocationTips();
                            return false;
                        }
                    }
                    $('#bank-card').html(cardnum);
                    $('#bank-list').html(htmldata);
                }
            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        });
    } else {
        LocationTips('用户登录信息已失效， 请重新登录！','/index');
        return false;
    }
    if (loginData) {
        $('select[name=bankList]').change(function () {
            var id = $('select[name=bankList] option:selected').val();
            $.ajax({
                type: "get",
                url: "/ajax/get/oneBank",
                data: {id: id},
                dataType: 'json',
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
    } else {
        LocationTips('用户登录信息已失效， 请重新登录！','/index');
        return false;

    }
    function LocationTips(mes,url) {
        if (mes==''||mes==undefined){
            mes='请完善个人取款资料再进行取款！';
        }
        if(url==''||url==undefined){
            url='/member/account';
        }
        alert(mes);
        window.location.href =url;
    }
    $('input[name=drawMoney]').change(function () {
        var mes;
        mes = checkBalance(this);
        if (mes) {
            alert(mes);
            return false;
        }
    });
    $('input[name=drawPassword]').change(function () {
        checkDrawPassword('input[name=drawPassword]');
    });
    $('#draw-submit').click(function () {
        checkBalance('input[name=drawMoney]');
        if (ajaxdo == 1) {
            checkDrawPassword('input[name=drawPassword]');
        }
        if (ajaxdo == 1) {
            postData();
        }

    });
    $('#fresh-money').click(function () {
       $.ajax({
           url:'/ajax/get/balance',
           type:'get',
           data:{},
           headers: {
               'Authorization': 'bearer ' + getCookie('loginBack'),
               'Content-Type': 'application/json',
               'Accept': 'application/json',
               'platform': "pc"
           },
           success:function (data) {
               if(data.code){
                   alert(data.msg);
                   return false;
               }
               if (data.data){
                   $('#money').html(data.data)
               }

           }
       })
    });
    function checkBalance(id) {
        var money = $(id).val();
        var balance = loginData.balance;
        var mes = '';
        if (money > balance) {
            mes = '取款金额大于账户余额';
        }
        if (money < 10) {
            mes = '取款金额不得低于10元';
        }
        if (mes != '') {
            $(id).val(0);
            alert(mes);
            ajaxdo = 0;
            return false;
        } else {
            ajaxdo = 1;
        }
    }

    function checkDrawPassword(id) {
        var drawPassword = $(id).val();
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
                        alert(data.msg);
                        ajaxdo = 0;
                    }
                }
                if (xhr.status == 204) {
                    ajaxdo = 1;
                }
            }
        })
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
                $('#draw-submit').attr('disabled','disabled');
            },
            success: function (data, info, xhr) {
                console.log(data);
                if (data) {
                    if (data.code) {
                        alert(data.msg);
                        return false;
                    }
                    if (data.data) {
                        $('.audit').hide();
                        $('.mycenter').hide();
                        $('.smallamount').show();
                        var showdata = data.data;
                        var showTime = strToDate(showdata.create_time);
                        var showmoney = showdata.admin_money + showdata.deposit_money + showdata.charge;
                        var htmldata, normal, mui;
                        var htmldata1 = '';
                        var showresult = '';
                        if (showdata.mui_status == 1) {
                            mui = '满足';
                            showresult += '<span class="span-1"><i class="myd-red">*</i>满足综合稽核，无需扣除优惠金额</span>';
                        } else {
                            mui = '不满足';
                            showresult += '<span class="span-1"><i class="myd-red">*</i>不满足综合稽核，扣除优惠金</span>';
                        }
                        if (showdata.normal_status == 1) {
                            normal = '满足';
                            showresult += '<span class="span-1"><i class="myd-red">*</i>满足常态稽核，无需扣除50%行政优惠</span>';
                        } else {
                            normal = '不满足';
                            showresult += '<span class="span-1"><i class="myd-red">*</i>不满足常态稽核，需扣除50%行政优惠</span>';
                        }
                        showresult += '<span class="span-1"><i class="myd-red">*</i>综合稽核+常态稽核+手续费，共需扣除：<span class="myd-red">' + showmoney + '</span>元</span>';
                        htmldata = '<li class="fl">' + showTime[0]+showTime[1]+ '</li>' +
                            '<li class="fl"></li>' +
                            '<li class="fl">' + showdata.out_money + '</li>' +
                            '<li class="fl">' + showdata.admin_money + '</li>' +
                            '<li class="fl">' + showdata.deposit_money + '</li>' +
                            '<li class="fl">' + showdata.charge + '</li>' +
                            '<li class="fl">' + normal + '</li>' +
                            '<li class="fl">' + mui + '</li>';
                        $('#show-charge-data').html(htmldata);

                        $('#order-num').html(showdata.order_num);
                        $('#draw-money').html(showdata.out_money);
                        $('#deduction-money').html(showmoney);
                        $('#draw-out-money').html(showdata.out_charge);
                        $('#show-result').html(showresult);

                        if (data.data.out_status == 0) {
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
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            data: {id: bankId},
            success: function (data) {
                $('#bank-card').html(data.data.card)
            }
        })
    })

    $("#getMoney").keyup(function () {
        var val = $(this).val();
        $(this).val(val.replace(/[^\d^\.]+/g, ''))
    })

    $("#getPwd").keyup(function () {
        var val = $(this).val();
        $(this).val(val.replace(/[^\d^\.]+/g, ''))
    })

});