$(document).ready(function (e) {
    var loginBack = getCookie('loginBack');
    //交易记录时间插件开始时间
    $("#my1").calendar({});
    //时间插件结束时间
    $("#my2").calendar({});
    //时间插件开始时间
    $("#my3").calendar({});
    //时间插件结束时间
    $("#my4").calendar({});
    if(getCookie('isOutCharge')==1){
       $('.h-infornav1 li').eq(1).click();
        $('.h-infornav1 li').eq(1).addClass('curr').siblings().removeClass('curr');
        $('.report-c .report-item').eq(1).show().siblings().hide();
        GetCashTimeList();
        localStorage.wapWriteBankInfo=0;
    }
    //选项卡切换
    $('.h-recora li').click(function () {
        var ind = $(this).index()
        $(this).addClass('curr').siblings().removeClass('curr');
        $('.record-content-c .record-main-c').eq(ind).show().siblings().hide();
    });
    $('.h-infornav1').on('click', 'li', function () {
        var index = $(this).index();
        $(this).addClass('curr').siblings().removeClass('curr');
        $('.report-c .report-item').eq(index).show().siblings().hide();
        if(index==1){
            GetCashTimeList();
        }
    });

    $('.record-ul').on('click', '.record-ul-p', function () {

        $(this).siblings('ul').stop(true).slideToggle(100, function () {
            // myScroll.refresh()
        });
        if ($(this).children('.arrow').hasClass('icon-return-copy')) {
            $(this).children('.arrow').removeClass('icon-return-copy').addClass('icon-xiafanhui')
        } else {
            $(this).children('.arrow').removeClass('icon-xiafanhui').addClass('icon-return-copy')
        }
        if (getCookie('loginBack')) {
            var gameType = $(this).attr('vType');
            var state = $(this).attr('stat');
            var t = $(this);
            if (state == 0) {
                GetRecordInfo(gameType, t)
            }
        }
    });
    // $('#source_type').change(function () {
    //     GetCashTimeList();
    // });

    //交易记录查询
    function GetCashTimeList() {
        if (getCookie('loginBack')) {
            var end_time, start_time;
            start_time = $('#my1').val();
            end_time = $('#my2').val();
            if (start_time == '' || start_time == null) {
                start_time = 0;
            } else {
                start_time = DateToStr(start_time);
            }
            if (end_time == '' || end_time == null) {
                end_time = 0;
            } else {
                end_time = DateToStr(end_time);
            }
            if (getCookie('loginBack')) {
                var gameType = $('#source_type option:selected').val();
                $.ajax({
                    url: '/ajax/cash/record',
                    type: 'get',
                    headers: {
                        'Authorization': 'bearer ' + getCookie('loginBack'),
                        'platform': 'wap'
                    },
                    data: {
                        source_type: gameType,
                        start_time: start_time,
                        end_time: end_time
                    },
                    beforeSend: function () {
                        $(this).attr({disabled: "disabled"});
                        Loading();
                    },
                    success: function (data) {
                        console.log(data);
                        var htmData = '';
                        var metadata = data.meta;
                        if (data.data != null) {
                            data = data.data;
                            var types, ctimedata;
                            if (gameType == 13) {
                                $('.record-total-h').html(' <li class="p-flex-l">时间</li>\n' +
                                    '                            <li>状态</li>\n' +
                                    '                            <li>存入金额</li>\n' +
                                    '                            <li>存入优惠</li>\n' +
                                    '                            <li>其他优惠</li>')
                                for (var i = 0; i < data.length; i++) {
                                    switch (data[i].Status) {
                                        // 1已确认2已取消,3未处理
                                        case 1:
                                            types = '已确认';
                                            break;
                                        case 2:
                                            types = '已取消';
                                            break;
                                        case 3:
                                            types = '未处理';
                                            break;
                                    }
                                    if (data[i].CreateTime != 0) {
                                        ctimedata = strToDate(data[i].CreateTime);
                                    } else {
                                        ctimedata = '';
                                    }
                                    if (ctimedata != '') {
                                        htmData += ' <li>' +
                                            '<p class="p-flex-l"><span>' + ctimedata[0] + '</span><br><span>' + ctimedata[1] + '</span></p>';
                                    } else {
                                        htmData += ' <li>' +
                                            '<p class="p-flex-l"><span>暂无数据</span><br><span></span></p>';
                                    }

                                    htmData += '<p>' + types + '</p><p>' + data[i].DepositCount + '</p><p>' + data[i].DepositDiscount + '</p><p>' + data[i].OtherDiscount + '</p></li>';
                                }
                            } else {
                                for (var i = 0; i < data.length; i++) {
                                    switch (data[i].Type) {
                                        case 1:
                                            types = '存入';
                                            break;
                                        case 2:
                                            types = '取出';
                                            break;
                                    }
                                    if (data[i].CreateTime != 0) {
                                        ctimedata = strToDate(data[i].CreateTime);
                                    } else {
                                        ctimedata = '';
                                    }
                                    if (ctimedata != '') {
                                        htmData += ' <li>' +
                                            '<p class="p-flex-l"><span>' + ctimedata[0] + '</span><br><span>' + ctimedata[1] + '</span></p>';
                                    } else {
                                        htmData += ' <li>' +
                                            '<p class="p-flex-l"><span>暂无数据</span><br><span></span></p>';
                                    }

                                    htmData += '<p>' + types + '</p><p>' + data[i].Balance + '</p><p>' + data[i].DisBalance + '</p><p>' + data[i].AfterBalance + '</p></li>';
                                }
                            }

                        } else {
                            var mes = "";
                            if (data.code) {
                                mes = data.msg;
                            }
                            htmData = ShowNoData("", mes);
                        }
                        if (metadata) {
                            var links = data.links;
                            $('#record-total-ul').html(htmData);
                            $('#record-total-ul').parent().siblings(".one-foot").remove()
                            if (metadata.count != 0) {
                                var thtmData = pageList(metadata, links);
                                $('#record-total-ul').parent().parent().append($(thtmData));
                            }
                            
                            
                            //分页
                            if (metadata.count != 0) {
                                GetPageData(GetApplyData, metadata.page_count)
                            }

                        }

                    },
                    complete: function () {
                        $(this).removeAttr("disabled");
                        LoadingClose();
                    }
                })
            } else {
                location.href = '/m/login';
            }
        }
    }

    $('#record-search').click(function () {
        GetCashTimeList();
    });
    $('#apply-search').click(function () {
        GetCashTimeList();
    })

    $(".tabs_center").on("mousedown", ".btn_search", function () {
        if ((!$.browser.msie && e.button == 0) || ($.browser.msie && e.button == 1)) {
            $(this).addClass("buttonshadowR")
        }
    })
    $(".tabs_center").on("mouseup", ".btn_search", function () {
        $(this).removeClass("buttonshadowR")
    })


    var mescroll = new MeScroll("mescroll", {
        //下拉刷新的所有配置项
        down:{
            use: false, //是否初始化下拉刷新; 默认true
            auto: false, //是否在初始化完毕之后自动执行下拉回调callback; 默认true
            hardwareClass: "mescroll-hardware", //硬件加速样式;解决iOS下拉因隐藏进度条而闪屏的问题,参见mescroll.css
            isBoth: false, //下拉刷新时,如果滑动到列表底部是否可以同时触发上拉加载;默认false,两者不可同时触发;
        }
    });

});

//投注记录查询
function GetRecordInfo(gameType, t) {
    $.ajax({
        url: '/ajax/record/info',
        type: 'get',
        data: {game_name: gameType},
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'wap'
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            // console.log(data)
            datalist = data.data;
            var htmlStr = '';
            if (data.code) {
                $.toast(data.msg, 'text');
                return false;
            }
            if (datalist !== null) {
                if (datalist.length > 0) {
                    for (var i = 0; i < datalist.length; i++) {
                        htmlStr += '<li><p class="time"><i>▍</i><span>' + datalist[i].bet_time + '</span></p><h2 class="six"><i>' + datalist[i].game_name + '</i><span>' + datalist[i].order_id + '</span></h2> <p class="any"><i>开奖结果：</i><span>' + datalist[i].game_result + '</span></p><dt class="bet"><p>总投注：<i>' + datalist[i].bet_all + '元</i></p> <p>有效投注：<i>' + datalist[i].bet_yx + '元</i></p><p>盈利：<i>' + datalist[i].win + '元</i></p></dt> </li>';
                    }
                }
            } else {
                htmlStr = ShowNoData();
            }
            t.attr('stat', '1');
            var metadata = data.meta;
            if (metadata) {
                var links = data.links;
                if (metadata.count != 0) {
                    htmlStr += pageList(metadata, links);
                }
                t.next('ul').append(htmlStr)
                //分页
                GetPageData(GetApplyData, metadata.page_count)
            }

        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

// $('.h-head .iconfont').click(function () {
//     window.history.go(-1);
// });
//优惠申请
$('#proApply').click(function () {
    GetTitleList();
    GetApplyData();
});

//获取优惠数据
function GetApplyData(page) {
    var startTime = $('#my3').val();
    var endTime = $('#my4').val();
    var orderId = $('#apply-title-list option:selected').val();
    // console.log(orderId)
    if (orderId == undefined || orderId == '') {
        orderId = 0;
    } else {
        orderId = Number(orderId);
    }
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
    if (page == undefined) {
        page = 0;
    }
    $.ajax({
        url: '/ajax/wap/apply/list',
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
            'platform': 'wap'
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            // console.log(data);
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
                    htmData += '<ul><li >' + datalist[i].promotionTitle + '</li><li><span>' + creatTime[0] + '<span><br><span>' + creatTime[1] + '</span></li><li>' + datalist[i].applyMoney + '元</li><li>' + datalist[i].giveMoney + '元</li><li>' + applyType + '</li></li></ul>';
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
            if (metadata) {
                var links = data.links;
                // $('#record-total-ul').html(htmData);
                // $('#record-total-ul').parent().siblings(".one-foot").remove()
                if (metadata.count != 0) {
                    htmData += pageList(metadata, links);
                }
                $('#apply-total-data').html(htmData);
                //分页
                GetPageData(GetApplyData, metadata.page_count)
            }
        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

//获取优惠活动列表
function GetTitleList() {
    $.ajax({
        url: '/ajax/wap/apply/config',
        type: 'get',
        async: false,
        data: {},
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'platform': 'wap'
        },
        beforeSend: function () {
            $(this).attr({disabled: "disabled"});
            Loading();
        },
        success: function (data) {
            // console.log(data);
            var htmData = '<option value="0">全部</option>';
            if (data.data) {
                var datalist = data.data;
                for (var i = 0; i < datalist.length; i++) {
                    htmData += '<option value="' + datalist[i].id + '">' + datalist[i].proTitle + '</option>';
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
            $('#apply-title-list').html(htmData);

        },
        complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

