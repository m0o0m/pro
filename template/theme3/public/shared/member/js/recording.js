//查询投注记录
function GetCashTimeList(page) {
    if (getCookie('loginBack')) {
        var gameType = $('.navs_active').attr('stat');
        var OrderNum = $('input[name=orderNum]').val();
        var startTime = $('.choose .inp').val();
        var endTime = $('.choose .inp2').val();
        if(startTime==endTime){
            startTime = DateToStr(0);
            endTime = DateToStr(1);
        }
        if (startTime == null || startTime == "") {
            startTime = DateToStr(0);
        } else {
            startTime = DateToStr(startTime);
        }

        if (endTime == null || endTime == "") {
            endTime = DateToStr(1);
        } else {
            endTime = DateToStr(endTime);
        }

        var gameName = $("#source_type .checkedDiv .checked").attr("name") || '';

        var gameonetpe = $("#source_type_one .checkedDiv .checked").attr("name") || '';

        gameType = Number(gameType) || 0;
        OrderNum = Number(OrderNum) || 0;

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
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {
                // console.log(data)
                var htmData = '';
                if (data.data != null) {
                    datlist = data.data;
                    var types, ctimedata;
                    for (var i = 0; i < datlist.length; i++) {
                        ctimedata = strToDate(datlist[i].bet_timeline);
                        //ctime,types,data.Balance,data.DisBalance,AfterBalance,Remark
                        htmData += '<ul class="ul1f clearfix">\n' +
                            '                                        <li>' + ctimedata[0] + ctimedata[1] + '</li>\n' +
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
                var totalPage = metadata.page_count;
                //分页
                GetPageData(GetCashTimeList, totalPage);
            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        })
    } else {
        window.location.href = '/m/login'
    }
}

//获取游戏分类列表
function GetGameTypesList() {
    if (getCookie('loginBack')) {
        var gameType = $('.navs_active').attr('stat');
        $.ajax({
            url: '/ajax/get/gameList',
            type: 'get',
            async: false,
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'platform': 'pc'
            },
            data: {typeId: gameType},
            success: function (data) {
                if (data) {
                    if (data.data) {
                        new MySelect({
                            domId: "#source_type",
                            liArr: data.data,
                            liclick: liclick,
                            VType: "VType",
                            ProductName: 'ProductName'
                        })
                    }
                } else {
                    new MySelect({
                        domId: "#source_type",
                        liArr: []
                    })
                }

                showGameName(gameType);

            }
        })
    } else {
        window.location.href = '/m/login'
    }
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

    $(".bettime >input").each(function () {
        $(this).val("")
    })

    var stat = $(this).attr('stat');
    $('.navs').attr('stat', stat);
    showGameName(stat);
    $('#source_type_one').hide().html('');
    GetGameTypesList();
    var nowdata=GetNowDate();
    $('.choose .inp').val(nowdata);
    $('.choose .inp2').val(nowdata);
});
//投注记录检索
$('#btn_search').click(function () {
    GetCashTimeList();
})
//模块下分类
$(function () {
    var nowdata=GetNowDate();
    $('.choose .inp').val(nowdata);
    $('.choose .inp2').val(nowdata);
    $('.choose .inp3').val(nowdata);
    $('.choose .inp4').val(nowdata);
    $('.choose .inp5').val(nowdata);
    $('.choose .inp6').val(nowdata);
    if (isOther == 1) {
        GetCashList();

        $('.r_center .record_item').eq(0).hide();
        $('.r_center .record_item').eq(1).show().siblings().hide();
    }
});

function liclick() {
    var gametype = $("#source_type").children().eq(0).children().attr("name")

    if (gametype == "pk_fc" || gametype == "eg_fc" || gametype == "cs_fc") {
        GetFcList(gametype);
    } else if (gametype == "pk_sp" || gametype == "im_sp" || gametype == "sb_sp" || gametype == "bbin_sp") {
        let arr = [
            {
                VType: 'single',
                ProductName: '体育单式'
            },
            {
                VType: 'more',
                ProductName: '体育串式'
            }
        ];

        new MySelect({
            domId: "#source_type_one",
            liArr: arr,
            VType: "VType",
            ProductName: 'ProductName'
        })
        $('#source_type_one').show();

    } else {
        $('#source_type_one').hide();
    }
}

$('input[name=orderNum]').keyup(function () {
    var val = $(this).val();
    val = val.replace(/[^\d.]/g, "");  //清除“数字”和“.”以外的字符   
    $(this).val(val);
})
/***************时间插件*********************** start *************************/
$('.choose .inp').click(function () {
    $('#schedule-box').show();
    $('#schedule-box2').hide();
})
$('.choose .inp2').click(function () {
    $('#schedule-box2').show();
    $('#schedule-box').hide();
})
$('.choose .inp3').click(function () {
    $('#schedule-box3').show();
    $('#schedule-box4').hide();
})
$('.choose .inp4').click(function () {
    $('#schedule-box4').show();
    $('#schedule-box3').hide();
});
$('.choose .inp5').click(function () {
    $('#schedule-box6').hide();
    $('#schedule-box5').show();
})
$('.choose .inp6').click(function () {
    $('#schedule-box6').show();
    $('#schedule-box5').hide();

});
$('.timebox').each(function (i) {
    $(this).mouseleave(function () {
        $(this).hide();
    })
});
/***************时间插件*********************** end *************************/
//现金流水
$('#cash-data-list').click(function () {
    // GetCashList()
});
$('.txt_l').click(function () {
    GetGameTypesList(4);
})
//现金记录检索查询
$('#btn_search_cash').click(function () {
    GetCashList()
})

//查询模块下游戏列表
function GetFcList(gametype) {

    if (getCookie('loginBack')) {
        $.ajax({
            url: '/ajax/get/game/list',
            type: 'get',
            async: false,
            data: {game_type: gametype},
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'platform': 'pc'
            },
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {

                if (data.data) {
                    new MySelect({
                        domId: "#source_type_one",
                        liArr: data.data,
                        VType: "game_type",
                        ProductName: 'game_name'
                    })
                } else {
                    new MySelect({
                        domId: "#source_type",
                        liArr: []
                    })
                }
                $("#source_type_one").show()
            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        })
    } else {
        window.location.href = '/m/login'
    }
    // console.log(gametype);

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
    if (getCookie('loginBack')) {
        var orderNum = $('input[name=cash-order-num]').val();
        var startTime = $('.choose .inp3').val();
        var endTime = $('.choose .inp4').val();

        var cashType = $("#cash_source_type .checkedDiv .checked").attr("name") || 99;
        if(startTime==endTime){
            startTime = DateToStr(0);
            endTime = DateToStr(1);
        }
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
                    cashType = Number(cashType);
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
                                '<li title="' + dataList[i].OrderNum + '">' + dataList[i].OrderNum + '</li>' +
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
                                    typeName = '入款';
                                    break;
                                case 1:
                                    typeName = '公司入款';
                                    break;
                                case 2:
                                    typeName = '线上入款';
                                    break;
                                case 3:
                                    typeName = '取款';
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
                var totalPage = metadata.page_count;
                //分页
                GetPageData(GetCashTimeList, totalPage);
            },
            complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        })
    }

}

/************** 时间插件 ****** start **********/
new Schedule({
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
new Schedule({
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
new Schedule({
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
new Schedule({
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
new Schedule({
    el: '#schedule-box5',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp5').value = y + '-' + m + '-' + d;
        $('#schedule-box5').hide();
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
new Schedule({
    el: '#schedule-box6',
    // date: '2018-9-20',
    clickCb: function (y, m, d) {
        //点击日期回调（可选）
        document.querySelector('.choose .inp6').value = y + '-' + m + '-' + d;
        $('#schedule-box6').hide();
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
/************* 优惠申请查询 start ***************/
$('#apply-data-list').click(function () {
    GetTitleList();
    GetApplyData();
});
$('#btn_search_apply').click(function () {
    GetApplyData();
})

//优惠申请查询
function GetApplyData(page) {
    var startTime = $('.choose .inp5').val();
    var endTime = $('.choose .inp6').val();
    var orderId = $("#apply-title-list .checkedDiv .checked").attr("name");
    orderId = Number(orderId) || 0;
    if(startTime==endTime){
        startTime = DateToStr(0);
        endTime = DateToStr(1);
    }
    if (startTime == null || startTime == "") {
        startTime = DateToStr(0);
    } else {
        startTime = DateToStr(startTime);
    }
    if (endTime == null || endTime == "") {
        endTime = DateToStr(1);
    } else {
        endTime = DateToStr(endTime);
    }
    $.ajax({
        url: '/ajax/apply/list',
        type: 'get',
        async: false,
        data: {
            start_time: startTime,
            end_time: endTime,
            order_num: orderId,
            page: page
        },
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'pc'
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            console.log(data);
            var htmData = '';
            if (data.data) {
                var datalist = data.data;
                for (var i = 0; i < datalist.length; i++) {
                    var creatTime = datalist[i].createtime;
                    var applyType = '';
                    creatTime = strToDate(creatTime);
                    switch (datalist[i].status) {
                        case 1:
                            applyType = '待审核';
                            break;
                        case 2:
                            applyType = '审核通过';
                            break;
                        case 3:
                            applyType = '审核不通过';
                            break;
                    }
                    htmData += '<ul class="ul1f clearfix">\n' +
                        '                                        <li>' + datalist[i].promotionTitle + '</li>\n' +
                        '                                        <li>' + creatTime[0] + creatTime[1] + '</li>\n' +
                        '                                        <li>' + datalist[i].applyreason + '</li>\n' +
                        '                                        <li>' + datalist[i].denyreason + '</li>\n' +
                        '                                        <li>' + datalist[i].applyMoney + '</li>\n' +
                        '                                        <li>' + datalist[i].giveMoney + '</li>\n' +
                        '                                        <li>' + applyType + '</li>\n' +
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
            // console.log(htmData)
            var metadata = data.meta;
            var links = data.links;
            if (metadata) {
                if (metadata.count != 0) {
                    htmData += pageList(metadata, links);
                }
            }
            $('#show-aply-list').html(htmData);
            if (metadata) {
                if (metadata.count != 0) {
                    //分页
                    GetPageData(GetApplyData, metadata.page_count)
                }
            }


        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

//优惠申请活动列表
function GetTitleList() {
    $.ajax({
        url: '/ajax/apply/config',
        type: 'get',
        async: false,
        data: {},
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'pc'
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            if (data) {
                if (data.data) {
                    new MySelect({
                        domId: "#apply-title-list",
                        liArr: data.data,
                        liclick: liclick,
                        VType: "id",
                        ProductName: 'proTitle'
                    })
                }
            } else {
                new MySelect({
                    domId: "#apply-title-list",
                    liArr: []
                })
            }

        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

function MySelect(obj) {

    var domId = obj.domId;
    var liArr = obj.liArr;
    var VType = obj.VType;
    var val = obj.ProductName;

    var my_Select = document.querySelector(domId);

    if (my_Select.children.length) {
        for (var i = my_Select.children.length - 1; i >= 0; i--) {
            my_Select.removeChild(my_Select.children[i]);
        }
    }


    var checkedDiv = document.createElement("div");
    checkedDiv.setAttribute("class", "checkedDiv");

    var checked = document.createElement("input");
    checked.setAttribute("class", "checked");
    checked.setAttribute('value', "全部");
    checked.setAttribute('type', "text");
    checked.setAttribute('readonly', "readonly");

    var selectUldiv = document.createElement("div");
    selectUldiv.setAttribute("class", "selectUldiv");

    var ul = document.createElement("ul");
    ul.setAttribute("class", "selectUl");

    // let li = document.createElement("li");
    // $(li).text("全部");
    // $(li).attr("vType","");        
    // ul.appendChild($(li)[0])


    if (liArr) {
        let index = 1;

        for (const i of liArr) {
            if (i[val] == "全部") {
                index = 0;
                break
            }
        }
        if (index) {
            let li = document.createElement("li");
            $(li).text("全部");
            $(li).attr("vType", "");
            ul.appendChild($(li)[0])
        }

        for (const i of liArr) {
            let li = document.createElement("li");
            $(li).text(i[val]);
            $(li).attr("vType", i[VType]);
            ul.appendChild($(li)[0])
        }
    }

    checkedDiv.appendChild(checked);
    selectUldiv.appendChild(ul);

    my_Select.appendChild(checkedDiv);
    my_Select.appendChild(selectUldiv);

    my_Select = $(domId);

    my_Select.on("click", ".checkedDiv", function () {
        $(this).siblings().show();
    })
    my_Select.on("click", ".selectUl li", function () {
        var input = my_Select.children().eq(0).children().eq(0);
        input.prop('value', $(this).text());
        input.prop('name', $(this).attr("vType"));
        $(this).parent().parent().hide();
        if (obj.liclick) {
            obj.liclick()
        }
    })

    my_Select.mouseleave(function () {
        $(this).children().eq(1).hide();
    })
}

new MySelect({
    domId: "#cash_source_type",
    liArr: [
        {
            vType: "99",
            ProductName: '全部'
        },
        {
            vType: "0",
            ProductName: '存款'
        }, {
            vType: "1",
            ProductName: '会员返佣'
        }, {
            vType: "2",
            ProductName: '线上入款'
        }, {
            vType: "3",
            ProductName: '人工取出'
        }, {
            vType: 4,
            ProductName: '取款'
        }, {
            vType: "5",
            ProductName: '线上出款'
        }, {
            vType: "6",
            ProductName: '注册优惠'
        }, {
            vType: 7,
            ProductName: '下单'
        }, {
            vType: "8",
            ProductName: '额度转换'
        }, {
            vType: "9",
            ProductName: '优惠返水'
        }, {
            vType: "10",
            ProductName: '自助返水'
        }, {
            vType: "11",
            ProductName: '会员返佣'
        }, {
            vType: "12",
            ProductName: '红包'
        },
        {
            vType: "13",
            ProductName: '入款监控'
        },
    ],
    VType: "vType",
    ProductName: 'ProductName',
})
