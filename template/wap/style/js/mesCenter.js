var loginBack = getCookie('loginBack');
$(function () {
    GetWapMes();
});
function GetWapMes(page) {
    if (loginBack) {
        $.ajax({
            url: '/ajax/get/mes/list',
            type: "get",
            headers: {
                'Authorization': 'bearer ' + loginBack,
                'platform': 'pc'
            },
            data: {page:page,pageSize:10},
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {
                // $('.o2 dl ul').html()
                var htmldata = '';
                var metadata = data.meta;
                var links = data.links;
                data = data.data;
                if (data != null) {
                    for (var i = 0; i < data.length; i++) {
                        var datetime = new Date(data[i].create_time);
                        var Y, M, D, h, m, s;
                        // console.log(data[i].create_time)
                        Y = datetime.getFullYear() + '-';
                        M = (datetime.getMonth() + 1 < 10 ? '0' + (datetime.getMonth() + 1) : datetime.getMonth() + 1) + '-';
                        D = datetime.getDate() + ' ';
                        h = datetime.getHours() + ':';
                        m = datetime.getMinutes() + ':';
                        s = datetime.getSeconds();
                        // console.log(Y+M+D+h+m+s);
                        var mestime = Y + M + D + h + m + s;
                        htmldata += '<li><h1 value="' + data[i].id + '" stat="' + data[i].state + '"><a>' + data[i].title + '</a><i>' + mestime + '</i>';
                        if (data[i].state == 2) {
                            htmldata += '<span style="float: right;color:#d1d1d1;">已读</span>';
                        }

                        htmldata += '</h1><p style="display:none;">' + data[i].content + '</p></li>';
                    }
                } else {
                    htmldata = ShowNoData('', '暂无任何消息');
                }
                if (metadata.count != 0) {
                    htmldata += pageList(metadata, links);
                }
                $('.o2 dl ul').html(htmldata);
                var totalPage = metadata.page_count;
                //分页
                GetPageData(GetWapMes, totalPage);
                $('.o2 dl ul li h1').click(function () {
                    $(this).siblings('p').slideToggle();
                    var state = $(this).attr('stat');
                    if (state == '1') {
                        var id = $(this).attr('value');
                        id = Number(id);
                        ChangeStatus($(this), id);
                    }
                });
            }, complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        });

    } else {
        location.href = "/m/login"
    }
}
$('.m-infornav .o1 dl dt').click(function () {
    if ($(this).next('ul').find('.game-show-no-notice')) {
        var gamenotice = $(this).next('ul').find('.game-show-no-notice');
        gamenotice.html(ShowNoData('', '暂无相关数据'));
    }
    $(this).next('ul').slideToggle().parents('dl').siblings('dl').children('ul').hide();

});

$('.m-infornav .o1 ul li h1').click(function () {
    $(this).next('p').slideToggle().parents('h1').siblings('h1').children('p').hide();

})
$('#mesCenter').click(function () {
    $('.o2').show();
    $('.o1').hide();
    GetWapMes();
})
$('.h-infornav li').click(function () {
    var index = $(this).index();
    $(this).addClass('curr').siblings().removeClass('curr');
    $('.m-infornav').find('div').eq(index).slideDown().siblings().slideUp();
});
// $('.h-head .iconfont').click(function () {
//     window.history.go(-1);
// });

function ChangeStatus(t, id) {
    // console.log(id)
    $.ajax({
        url: '/ajax/mes/status',
        type: 'put',
        headers: {
            'Authorization': 'bearer ' + loginBack,
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': 'wap'
        },
        dataType: 'json',
        contentType: "application/json",
        data: JSON.stringify({Id: id}),
        success: function (data, con, code) {
            if (code.status == 204) {
                t.attr('stat', '2');
                t.children('i').after('<span style="float: right;color:#d1d1d1;">已读</span>');
            }
        }
    })
}


$(document).ready(function(){
    var mescroll = new MeScroll("mescroll", {
            //下拉刷新的所有配置项
            down:{
                use: false, //是否初始化下拉刷新; 默认true
                auto: false, //是否在初始化完毕之后自动执行下拉回调callback; 默认true
                hardwareClass: "mescroll-hardware", //硬件加速样式;解决iOS下拉因隐藏进度条而闪屏的问题,参见mescroll.css
                isBoth: false, //下拉刷新时,如果滑动到列表底部是否可以同时触发上拉加载;默认false,两者不可同时触发;
            }
        });
})

function GetNoticeData(page) {
    var id=localStorage.setId;
    $.ajax({
        url: '/ajax/noticeList',
        data: {id: id,page:page,pageSize:10},
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
            var htmldata = '';
            var metadata = data.meta;
            var links = data.links;
            if (data) {
                var showdata=data.data;
                for (var i = 0; i < showdata.length; i++) {
                    timestr = strToDate(showdata[i].notice_date);
                    htmldata += '<ul class="ul4" stat="' + showdata[i].id + '">';
                    htmldata += '<li>' + showdata[i].notice_title + '</li><li>' + timestr[0] + timestr[1] + '</li><li class="notice-content">' + showdata[i].notice_content + '</li>';
                    htmldata += '</ul>';
                }
            } else {
                htmldata = '暂无数据';
            }
            if (metadata.count != 0) {
                htmldata += pageList(metadata, links);
            }
            $('#notice-show').html(htmldata);
            var totalPage = metadata.page_count;
            //分页
            GetPageData(GetNoticeData, totalPage);
        }, complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

