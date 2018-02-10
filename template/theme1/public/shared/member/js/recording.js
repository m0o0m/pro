//查询投注记录
function GetCashTimeList(page) {
    var gameType = $('.navs_active').attr('stat');
    var OrderNum = $('input[name=orderNum]').val();
    var startTime = $('.choose .inp').val();
    var endTime = $('.choose .inp2').val();

    if (startTime == null || startTime == "") {
        startTime = 0;
    } else {
        startTime = DateToStr(startTime);
    }
    if (endTime == null || endTime == "") {
        endTime = 0;
    } else {
        endTime = DateToStr(endTime);
    }

    if (OrderNum != null) {
        OrderNum = Number(OrderNum);
    }
    if (gameType != null) {
        gameType = Number(gameType);
    } else {
        gameType = 0;
    }
    var gameName = $("#source_type option:selected").val();
    if (gameName == undefined) {
        gameName = '';
    }
    if (OrderNum == undefined) {
        OrderNum = 0;
    }
    var gameonetpe = $("#source_type_one option:selected").val();
    if (gameonetpe == undefined) {
        gameonetpe = '';
    }
    $.ajax({
        url: '/ajax/record/infoList',
        type: 'get',
        async: false,
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'pc'
        },
        data: {
            vType: gameType,
            start_time: startTime,
            end_time: endTime,
            order_num: OrderNum,
            game_name: gameName,
            game_one_type: gameonetpe,
            page: page
        },
        beforeSend: function () {
            $(this).attr({ disabled: "disabled" });
            Loading();
        },
        success: function (data) {
            console.log(data)
            var htmData = '';
            if (data.data != null) {
                datlist = data.data;
                var types, ctimedata;
                for (var i = 0; i < datlist.length; i++) {
                    switch (datlist[i].Type) {
                        case 1:
                            types = '存入';
                            break;
                        case 2:
                            types = '取出';
                            break;
                    }
                    ctimedata = strToDate(datlist[i].CreateTime);
                    //ctime,types,data.Balance,data.DisBalance,AfterBalance,Remark
                    htmData += '<ul class="ul1f clearfix">\n' +
                        '                                        <li>' + datlist[i].bet_time + '</li>\n' +
                        '                                        <li>' + datlist[i].order_id + '</li>\n' +
                        '                                        <li>' + datlist[i].game_name + '</li>\n' +
                        '                                        <li>' + datlist[i].win + '</li>\n' +
                        '                                        <li>' + datlist[i].bet_all + '</li>\n' +
                        '                                        <li>' + datlist[i].bet_yx + '</li>\n' +
                        '                                    </ul>';
                }
            } else {
                var mes = '';
                if (data.code) {
                    mes = data.msg;
                } else {
                    mes = '暂无数据';
                }
                htmData = ShowNoData("", mes);
            }
            var metadata = data.meta;
            var links = data.links;
            if (metadata.count != 0) {
                htmData += pageList(metadata, links);
            }
            $('#show-record-list').html(htmData);
            //分页
            $('#show-record-list .one-foot .fr').on('click', '.one-foot-circle', function () {
                var ind = $(this).index();
                $(this).addClass('page-active').siblings().removeClass('page-active');
                var page = Number($(this).text());
                GetCashTimeList(page);
            });
            $('#show-record-list .xla_k').change(function () {
                var page = $('#show-record-list .xla_k option:selected').val();
                GetCashTimeList(page);
            })
        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

//获取游戏分类列表
function GetGameTypesList() {
    var gameType = $('.navs_active').attr('stat');
    $.ajax({
        url: '/ajax/get/gameList',
        type: 'get',
        async: false,
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'pc'
        },
        data: { typeId: gameType },
        success: function (data) {
            var htmldata = '<option value="">全部</option>';
            if (data) {
                if (data.data) {
                    var showdata = data.data;
                    for (var i = 0; i < showdata.length; i++) {
                        htmldata += ' <option value="' + showdata[i].VType + '">' + showdata[i].ProductName + '</option>';
                    }
                }
            } else {
                htmldata = '暂无数据';
            }
            showGameName(gameType);
            $('#source_type').html(htmldata)
        }
    })
}

//现金流水和投注记录切换
$('.r_name').on('click', 'li', function () {
    var index = $(this).index();
    $(this).addClass('txt_active').siblings().removeClass('txt_active');
    $('.r_center .record_item').eq(index).show().siblings().hide();
})
//5大分类点击事件
$('.navs').on('click', 'li', function () {
    var index = $(this).index();
    $(this).addClass('navs_active').siblings().removeClass('navs_active');
    // $('.tabs_center .tabs_item').eq(index).show().siblings().hide();
    var stat = $(this).attr('stat');
    $('.navs').attr('stat', stat);
    showGameName(stat);
    $('.select1').hide();
    GetGameTypesList();
    GetCashTimeList();
});
//投注记录检索
$('#btn_search').click(function () {
    GetCashTimeList();
})
//模块下分类
$(function () {
    if (getCookie('loginBack')) {
        if (isOther == 1) {
            GetCashList();
        } else {
            GetCashTimeList();
            GetGameTypesList(4);
        }

    } else {
        location.herf = '/index'
    }
    $('#source_type').change(function () {
        var gametype = $('#source_type option:selected').val();
        console.log(gametype);
        if (gametype == "pk_fc" || gametype == "eg_fc" || gametype == "cs_fc") {
            $('.select1').show();
            GetFcList(gametype);
        } else if (gametype == "pk_sp" || gametype == "im_sp" || gametype == "sb_sp" || gametype == "bbin_sp") {
            $('.select1').show();
            $('#source_type_one').html(' <option value="single">体育单式</option><option value="more">体育串式</option>')
        } else {
            $('.select1').hide();
        }
    })

});
/***************时间插件*********************** start *************************/
$('.choose .inp').click(function () {
    $('#schedule-box').show();
})
$('.choose .inp2').click(function () {
    $('#schedule-box2').show();
})
$('.choose .inp3').click(function () {
    $('#schedule-box3').show();
})
$('.choose .inp4').click(function () {
    $('#schedule-box4').show();
});

$('.timebox').each(function (i) {
    $(this).mouseleave(function () {
        $(this).hide();
    })
});
/***************时间插件*********************** end *************************/
//现金流水
$('#cash-data-list').click(function () {
    GetCashList()
});
$('.txt_l').click(function () {
    GetCashTimeList();
    GetGameTypesList(4);
})
//现金记录检索查询
$('#btn_search_cash').click(function () {

    GetCashList()
})

//查询模块下游戏列表
function GetFcList(gametype) {
    console.log(gametype);
    $.ajax({
        url: '/ajax/get/game/list',
        type: 'get',
        async: false,
        data: { game_type: gametype },
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'pc'
        },
        success: function (data) {
            console.log(data);
            var htmldata = '<option value="">全部</option>';
            if (data.data) {
                var datalist = data.data;
                for (var i = 0; i < datalist.length; i++) {
                    htmldata += ' <option value="' + datalist[i].game_type + '">' + datalist[i].game_name + '</option>';
                }
            }
            $('#source_type_one').html(htmldata)
        }
    })
}

//展示分类名字
function showGameName(stat) {
    var typemes = '';
    stat = Number(stat);
    switch (stat) {
        case 1:
            typemes = '视讯种类：';

            break;
        case 2:
            typemes = '电子种类：';
            break;
        case 3:
            typemes = '捕鱼种类：';
            break;
        case 4:
            typemes = '彩票种类：';
            break;
        case 5:
            typemes = '体育种类：';
            break;
    }
    $('.s_type').children('span').text(typemes);
}


//现金流水数据
function GetCashList(page) {
    if(getCookie('loginBack')){
        var orderNum = $('input[name=cash-order-num]').val();
        var startTime = $('.choose .inp3').val();
        var endTime = $('.choose .inp4').val();
        var cashType = $('#cash_source_type option:selected').val();

        if (startTime == null || startTime == undefined || startTime == '') {
            startTime = 0;
        } else {
            startTime = DateToStr(startTime);
        }

        if (endTime == null || endTime == undefined || endTime == '') {
            endTime = 0;
        } else {
            endTime = DateToStr(endTime);
        }
        $.ajax({
            url: '/ajax/cashRecord',
            type: 'get',
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            data: {
                start_time: startTime,
                end_time: endTime,
                order_num: orderNum,
                source_type: cashType,
                page: page
            },
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {
                // console.log(data);
                var htmData = '';
                if (data.data != null) {
                    var dataList = data.data;
                    var types, ctimedata, typeName;
                    cashType=Number(cashType);
                    if (cashType == 13) {
                        // console.log("入款监控数据：", data);

                        for (var i = 0; i < dataList.length; i++) {
                            var datastatus = Number(dataList[i].Status)
                            ctimedata = strToDate(dataList[i].CreateTime);
                            switch (datastatus) {
                                case 1:
                                    typeName = '已确认';
                                    break;
                                case 2:
                                    typeName = '已取消';
                                    break;
                                case 3:
                                    typeName = '未处理';
                                    break;
                            }
                            htmData += '<ul class="ul1f clearfix">' +
                                '<li>' + ctimedata[0] + ctimedata[1] + '</li>' +
                                '<li title="'+dataList[i].OrderNum +'">' + dataList[i].OrderNum + '</li>' +
                                '<li>入款监控</li>' +
                                '<li>存入</li>' +
                                '<li>' + dataList[i].DepositMoney + '</li>' +
                                '<li>' + dataList[i].DepositCount + '</li>' +
                                '<li>' + typeName + '</li>' +
                                '</ul>';
                        }
                    } else {
                        // console.log("其他数据：", data);
                        for (var i = 0; i < dataList.length; i++) {
                            switch (dataList[i].Type) {
                                case 1:
                                    types = '存入';
                                    break;
                                case 2:
                                    types = '取出';
                                    break;
                            }
                            switch (dataList[i].SourceType) {
                                case 0:
                                    typeName = '人工存入';
                                    break;
                                case 1:
                                    typeName = '公司入款';
                                    break;
                                case 2:
                                    typeName = '线上入款';
                                    break;
                                case 3:
                                    typeName = '人工取出';
                                    break;
                                case 4:
                                    typeName = '线上取款';
                                    break;
                                case 5:
                                    typeName = '出款';
                                    break;
                                case 6:
                                    typeName = '注册优惠';
                                    break;
                                case 7:
                                    typeName = '下单';
                                    break;
                                case 8:
                                    typeName = '额度转换';
                                    break;
                                case 9:
                                    typeName = '优惠返水';
                                    break;
                                case 10:
                                    typeName = '自助返水';
                                    break;
                                case 11:
                                    typeName = '会员返佣';
                                    break;
                                case 12:
                                    typeName = '红包';
                                    break;
                            }
                            ctimedata = strToDate(dataList[i].CreateTime);
                            //ctime,types,data.Balance,data.DisBalance,AfterBalance,Remark
                            htmData += '<ul class="ul1f clearfix">' +
                                '<li>' + ctimedata[0] + ctimedata[1] + '</li>' +
                                '<li>' + dataList[i].TradeNo + '</li>' +
                                '<li>' + typeName + '</li>' +
                                '<li>' + types + '</li>' +
                                '<li>' + dataList[i].Balance + '</li>' +
                                '<li>' + dataList[i].AfterBalance + '</li>' +
                                '</ul>';
                        }
                    }

                } else {
                    var msg = '';
                    if (data.code) {
                        msg = data.msg;
                    } else {
                        msg = '';
                    }
                    htmData = ShowNoData("", msg);
                }
                var metadata = data.meta;
                var links = data.links;
                if (metadata.count != 0) {
                    htmData += pageList(metadata, links);
                }
                $('#show-cash-list').html(htmData);
                $(' .one-foot .fr').on('click', '.one-foot-circle', function () {
                    var ind = $(this).index();
                    $(this).addClass('page-active').siblings().removeClass('page-active');
                    var page = Number($(this).text());
                    GetCashList(page);
                });
                $('.xla_k').change(function () {
                    var page = $('.xla_k option:selected').val();
                    GetCashList(page);
                })

            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        })
    }

}

/************** 时间插件 ****** start **********/
var mySchedule = new Schedule({
    el: '#schedule-box',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp').value = y + '-' + m + '-' + d;
        $('#schedule-box').hide();
    },
    // nextMonthCb: function (y,m,d) {
    //   //点击下个月回调（可选）
    // },
    // nextYeayCb: function (y,m,d) {
    //   //点击下一年回调（可选）
    // },
    // prevMonthCb: function (y,m,d) {
    //   //点击上个月回调（可选）
    // },
    // prevYearCb: function (y,m,d) {
    //    //点击上一年月回调（可选）
    // }
});
var mySchedule = new Schedule({
    el: '#schedule-box2',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp2').value = y + '-' + m + '-' + d;
        $('#schedule-box2').hide();
    },
    // nextMonthCb: function (y,m,d) {
    //   //点击下个月回调（可选）
    // },
    // nextYeayCb: function (y,m,d) {
    //   //点击下一年回调（可选）
    // },
    // prevMonthCb: function (y,m,d) {
    //   //点击上个月回调（可选）
    // },
    // prevYearCb: function (y,m,d) {
    //    //点击上一年月回调（可选）
    // }
});
var mySchedule = new Schedule({
    el: '#schedule-box3',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp3').value = y + '-' + m + '-' + d;
        $('#schedule-box3').hide();
    },
    // nextMonthCb: function (y,m,d) {
    //   //点击下个月回调（可选）
    // },
    // nextYeayCb: function (y,m,d) {
    //   //点击下一年回调（可选）
    // },
    // prevMonthCb: function (y,m,d) {
    //   //点击上个月回调（可选）
    // },
    // prevYearCb: function (y,m,d) {
    //    //点击上一年月回调（可选）
    // }
});
var mySchedule = new Schedule({
    el: '#schedule-box4',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp4').value = y + '-' + m + '-' + d;
        $('#schedule-box4').hide();
    },
    // nextMonthCb: function (y,m,d) {
    //   //点击下个月回调（可选）
    // },
    // nextYeayCb: function (y,m,d) {
    //   //点击下一年回调（可选）
    // },
    // prevMonthCb: function (y,m,d) {
    //   //点击上个月回调（可选）
    // },
    // prevYearCb: function (y,m,d) {
    //    //点击上一年月回调（可选）
    // }
});


/************** 时间插件 ****** end **********/

//按钮颜色
function buttonMouse(dom, classss) {
    dom.mousedown(function () {
        $(this).addClass(classss)
    })
    dom.mouseup(function () {
        $(this).removeClass(classss)
    })
}

buttonMouse($("#btn_search"), "buttonshadowR");
buttonMouse($("#btn_search_cash"), "buttonshadowR");

