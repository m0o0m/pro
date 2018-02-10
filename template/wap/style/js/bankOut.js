$(function () {
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
            if (data) {
                if (data.code) {
                    $.toast(data.msg, 'text');
                }
            }
            if (data.data) {
                var showdata = data.data;
                var showTime = strToDate(showdata.CreateTime);
                // var showmoney = showdata.admin_money + showdata.deposit_money + showdata.charge;
                var htmldata = '<section class="m-withdrawas">\n' +
                    '    <dl class="plan">\n' +
                    '    \t<dt class="line">\n' +
                    '        \t<div>\n' +
                    '            \t<i class="iconfont curr wap-red-dot"></i>\n' +
                    '                <i class="curr">————</i>\n' +
                    '                <i class="iconfont curr wap-dian2"></i>\n' +
                    '                <i>————</i>\n' +
                    '                <i class="iconfont wap-red-dot"></i>\n' +
                    '            </div>\n' +
                    '            <div>\n' +
                    '            \t<i class="curr">提交成功</i>\n' +
                    '                <i class="curr">正在稽核</i>\n' +
                    '                <i>待出款</i>\n' +
                    '            </div>\n' +
                    '        </dt>\n' +
                    '    </dl>\n' +
                    '</section><section class="m-Inspecti">\n' +
                    '\t<ul class="btn">\n' +
                    '    \t<a href="javascript:;" id="show-list">查看提交</a>\n' +
                    '    </ul>';
                htmldata += '<p class="detail"><i>稽查明细</i></p>\n' +
                    '    <table class="table">' +
                    '<tr> ' +
                    '           <td>存款日期</td>\n' +
                    '            <td>手续费</td>\n' +
                    '            <td>出款金额</td>\n' +
                    '            <td>优惠稽查</td>\n' +
                    '            <td>常态稽核</td>\n' +
                    '        </tr>' +
                    '<tr>\n' +
                    '        \t<td>' + showTime[0] + '<br/>' + showTime[1] + '</td>\n' +
                    '            <td>' + showdata.Charge + '</td>\n' +
                    '            <td>' + showdata.OutwardNum + '</td>\n' +
                    '            <td>' + showdata.FavourableMoney + '</td>\n' +
                    '            <td>' + showdata.ExpeneseMoney + '</td>\n' +
                    '        </tr>\n' +
                    '    </table> <li>*出款金额:<i class="fred"> ' + showdata.OutwardMoney + '元</i></li>';
                htmldata += '    <dl class="order">\n' +
                    '        <li class="fred">满足出款条件</li>\n' +
                    '</dl>';

                htmldata += '</section>';
                $('#scroller').html(htmldata);
                $('#show-list').click(function () {
                    setCookie("isOutCharge", 1);
                    location.href = "/m/record";
                })
            }
        },
        complete: function () {
            LoadingClose();
        }
    });
    if (getCookie('loginBack')) {
        $("#name").html(loginData.account);
        $("#money").html(loginData.balance);
        $.ajax({
            type: "get",
            url: "/m/member/bank",
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': 'wap'
            },
            data: {},
            dataType: 'json',
            beforeSend: function () {
                $(this).attr({ disabled: "disabled" });
                Loading();
            },
            success: function (data) {
                // console.log(data);
                var datalist = data.data;
                var cardnum, cardName;
                var realNameList = datalist.real_name;
                if (data.code) {
                    LocationTips(data.msg)
                } else {
                    var $sdata = datalist.bank_list;
                    // console.log($sdata);
                    if (!$sdata) {
                        LocationTips('', '', 'pass', 1);
                        return false;
                    }
                    if (realNameList.draw_password == '') {
                        LocationTips('请设置取款密码,再进行操作');
                        return false;
                    }

                    var htmldata = '';
                    if ($sdata) {
                        for (var i = 0; i < $sdata.length; i++) {
                            htmldata += ' <option value="' + $sdata[i].id + '">' + $sdata[i].title + '</option>';
                        }
                        cardnum = $sdata[0].card;
                        cardName = $sdata[0].card_name;
                    } else {
                        cardnum = '';
                        htmldata = ' <option value="">添加卡号后再操作</option>';
                        LocationTips('请添加银行卡号之后再进行操作', '/m/bankCard')
                    }
                    $('#card').html(cardnum);
                    $('#cardName').html(cardName);
                    $('select[name=bankList]').html(htmldata);
                }
            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        });
    } else {
        LocationTips('用户登录信息已失效， 请重新登录！', '/m/login')
    }
    if (getCookie('loginBack')) {
        $('select[name=bankList]').change(function () {
            var id = $('select[name=bankList] option:selected').val();
            $.ajax({
                type: "get",
                url: "/m/member/getCard",
                async: false,
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': platform
                },
                data: { id: id },
                dataType: 'json',
                beforeSend: function () {
                    $(this).attr({ disabled: "disabled" });
                    Loading();
                },
                success: function (data) {
                    if (data.code) {
                        LoadingClose();
                        $.toast(data.msg, 'text');
                    } else {
                        if (data.data) {
                            var getData = data.data;
                            $('#cardName').html(getData.card_name);
                            $('#card').html(getData.card);
                        }
                    }
                    // console.log(data)
                }, complete: function () {
                    $(this).removeAttr("disabled");
                    LoadingClose();
                }
            })
        })
    } else {
        LocationTips('用户登录信息已失效， 请重新登录！', '/m/login')
    }
    $('input[name=balance]').change(function () {
        checkBalance(this);

    });
    $('input[name=drawPassword]').change(function () {

    });
    $('#bankOut').click(function () {
        var doajax = 0;
        doajax = checkBalance('input[name=balance]', doajax);
        if (doajax == 1) {
            doajax = checkDrawPassword('input[name=drawPassword]', doajax);
        }
        if (doajax == 1) {
            postData();
        }
    });

    function LocationTips(mes, url, parameter, val) {
        if (mes == '' || mes == undefined) {
            mes = '请完善个人取款资料(真实姓名、取款银行卡、取款密码)再进行取款！';
        }
        if (url == '' || url == undefined) {
            url = '/m/bankCard';
        }
        if (parameter && val) {
            url += '?' + parameter + '=' + val;
        }
        LoadingClose();
        $.alert({
            title: '温馨提示',
            text: mes,
            onOK: function () {
                //点击确认
                Loading();
                window.location.href = url;
                LoadingClose();
            }
        });
    }

    function checkBalance(id, ajaxdo) {
        var money = $(id).val();
        var balance = loginData.balance;
        var mes = '';
        if (money > balance) {
            mes = '取款金额大于账户余额';
        }
        if (money < 1) {
            mes = '取款金额不得低于1元';
        }
        if (mes != '') {
            $(id).val(0);
            ajaxdo = 0;
            $.toast(mes, 'text');
        } else {
            ajaxdo = 1;
        }
        return ajaxdo;
    }

    function checkDrawPassword(id, doajax) {
        var drawPassword = $(id).val();
        var mes = '';
        $.ajax({
            url: '/m/check/drawPass',
            type: 'get',
            data: { draw_password: drawPassword },
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': platform
            },
            beforeSend: function () {
                $(this).attr({ disabled: "disabled" });
                Loading();
            },
            success: function (data, info, xhr) {
                if (data) {
                    if (data) {
                        if (data.code) {
                            $.toast(data.msg, 'text');
                            doajax = 0;
                        }
                    }
                    if (xhr.status == 204) {
                        doajax = 1;
                    }
                }
            }, complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        })
        return doajax;
    }

    function postData() {
        if (getCookie('loginBack')) {
            var balance = $('input[name=balance]').val();
            var password = $('input[name=drawPassword]').val();
            var bank_id = $('select[name=bankList] option:selected').val();
            $.ajax({
                type: "post",
                url: "/wap/draw/money/h",
                async: false,
                headers: {
                    'Authorization': 'bearer ' + getCookie('loginBack'),
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'platform': platform
                },
                data: JSON.stringify({
                    money: balance,
                    drawPassword: password,
                    bankId: bank_id
                }),
                dataType: 'json',
                beforeSend: function () {
                    Loading();
                },
                success: function (data, info, xhr) {
                    // console.log(data)
                    if (data) {
                        if (data.code) {
                            $.toast(data.msg, 'text');
                        }
                    }
                    if (data.data) {
                        var showdata = data.data;
                        var showTime = strToDate(showdata.create_time);
                        // var showmoney = showdata.admin_money + showdata.deposit_money + showdata.charge;
                        var htmldata = '<section class="m-withdrawas">\n' +
                            '    <dl class="plan">\n' +
                            '    \t<dt class="line">\n' +
                            '        \t<div>\n' +
                            '            \t<i class="iconfont curr wap-red-dot"></i>\n' +
                            '                <i class="curr">——————</i>\n' +
                            '                <i class="iconfont curr wap-dian2"></i>\n' +
                            '                <i>——————</i>\n' +
                            '                <i class="iconfont wap-red-dot"></i>\n' +
                            '            </div>\n' +
                            '            <div>\n' +
                            '            \t<i class="curr">提交成功</i>\n' +
                            '                <i class="curr">提交成功</i>\n' +
                            '                <i>提交成功</i>\n' +
                            '            </div>\n' +
                            '        </dt>\n' +
                            '    </dl>\n' +
                            '</section><section class="m-Inspecti">\n' +
                            '\t<ul class="btn">\n' +
                            '    \t<a href="javascript:;" id="show-h-list">查看提交</a>\n' +
                            '        <a href="/m/withdraw">重新提交</a>\n' +

                            '    </ul>';
                        htmldata += '<p class="detail"><i>稽查明细</i></p>\n' +
                            '    <table class="table">' +
                            '<tr> ' +
                            '           <td>存款日期</td>\n' +
                            '            <td>手续费</td>\n' +
                            '            <td>有效投注</td>\n' +
                            '            <td>优惠稽查</td>\n' +
                            '            <td>常态稽核</td>\n' +
                            '        </tr>' +
                            '<tr>\n' +
                            '        \t<td>' + showTime[0] + '<br/>' + showTime[1] + '</td>\n' +
                            '            <td>' + showdata.charge + '</td>\n' +
                            '            <td>' + showdata.bet_valid.bet_valid + '</td>\n' +
                            '            <td>' + showdata.deposit_money + '</td>\n' +
                            '            <td>' + showdata.admin_money + '</td>\n' +
                            '        </tr>\n' +
                            '    </table> <li>*出款金额:<i class="fred"> ' + showdata.out_charge + '元</i></li>';
                        if (data.out_status == 0) {
                            htmldata += '    <dl class="order">\n' +
                                '        <li class="fred">额度小于费用，无法出款</li>\n' +
                                '        <button id="back">返回取款页</button>\n' +
                                '    </dl>';
                        } else {
                            htmldata += '    <dl class="order">\n' +
                                '        <li class="fred">满足出款条件</li>\n' +
                                '        <button id="continue-out">继续出款</button>\n' +
                                '    </dl>';
                        }
                        htmldata += '</section>';
                        $('#scroller').html(htmldata);
                        //出款提交
                        $('#continue-out').click(function () {
                            location.href = "/m/Withdrawal/write"
                        });
                        $('#show-h-list').click(function () {
                            localStorage.wapWriteBankInfo = 1;
                            location.href = "/m/record";
                        })

                    }

                }, complete: function () {
                    $(this).removeAttr("disabled");
                    LoadingClose();
                }
            })
        } else {
            LocationTips('用户登录信息已失效， 请重新登录！', '/m/login')
        }
    }

    $('#money').click(function () {
        hideBalance($(this));
    })
});

$(document).ready(function () {
    var mescroll = new MeScroll("mescroll", {
        //下拉刷新的所有配置项
        down: {
            use: false, //是否初始化下拉刷新; 默认true
            auto: false, //是否在初始化完毕之后自动执行下拉回调callback; 默认true
            hardwareClass: "mescroll-hardware", //硬件加速样式;解决iOS下拉因隐藏进度条而闪屏的问题,参见mescroll.css
            isBoth: false, //下拉刷新时,如果滑动到列表底部是否可以同时触发上拉加载;默认false,两者不可同时触发;
        }
    });
})

