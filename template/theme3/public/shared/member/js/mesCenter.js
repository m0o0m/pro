var loginBack = getCookie('loginBack');
localStorage.setId = 7;
$(function () {
    GetMesList();
});

function GetMesList(page) {
    if (getCookie('loginBack')) {
        $.ajax({
            url: '/ajax/get/mes/list',
            type: "get",
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            data: {page: page, pageSize: 10},
            beforeSend: function () {
                $(this).attr({disabled: "disabled"});
                Loading();
            },
            success: function (data) {
                // $('.o2 dl ul').html()
                console.log(data)
                var metadata = data.meta;
                var links = data.links;
                var htmldata = '';
                var readli, mestime;
                data = data.data;
                if (data != null) {
                    for (var i = 0; i < data.length; i++) {

                        mestime = strToDate(data[i].create_time);
                        htmldata += '<ul class="clearfix date" value="' + data[i].id + '" stat="' + data[i].state + '"> ';
                        if (data[i].state == 2) {
                            readli = '<li class="read-state" style="color:#d1d1d1;">已读</li>';
                        } else {
                            readli = ' <li class="read-state">未读</li> ';
                        }
                        htmldata += readli + '<li>' + mestime[0] + mestime[1] + '</li> ' +
                            '<li class="pk_title" content="' + data[i].content + '">' + data[i].title + '  </li>' +
                            ' <li class="dd"> <span class="del">删除</span></li>' +
                            '</ul>';
                        // htmldata += '<div class="pk_content">' + data[i].content + '</div>';
                    }
                } else {
                    htmldata = ShowNoData('', '暂无任何消息');
                }


                if (metadata.count != 0) {
                    htmldata += pageList(metadata, links);
                }
                $('#person-msg').html(htmldata);
                var totalPage = metadata.page_count;
                //分页
                GetPageData(GetMesList, totalPage);
                $('.show-msg-block').each(function () {
                    var this_div = $(".date_nav div");
                    var pk_title = $(".dd .pk_title").index(this);
                    $(this).click(
                        function () {
                            this_div.eq(pk_title).slideToggle();
                            var state = $(this).attr('stat');
                            var id = $(this).attr('value');
                            var tstat = $('.ul1f .read-state');
                            var t = $(this);
                            id = Number(id);
                            if (state == 1) {
                                ChangeStatus(t, tstat, id)
                            }
                        }
                    );
                });
                $('.del').click(function () {
                    var id = $(this).parent().parent().find('.show-msg-block').attr('value');
                    var this_all = $(this).parent().parent();
                    id = Number(id);
                    console.log(id);
                    if (confirm('确认要删除吗？')) {
                        $.ajax({
                            url: '/ajax/delete/mes',
                            type: 'get',
                            data: {'id': id},
                            headers: {
                                'Authorization': 'bearer ' + getCookie('loginBack'),
                                'Content-Type': 'application/json',
                                'Accept': 'application/json',
                                'platform': "pc"
                            },
                            success: function (data, info, xhr) {
                                if (data) {
                                    if (data.code) {
                                        // alert(data.msg);
                                        My_pc_Modal({
                                            text: data.msg,
                                        })
                                    }
                                }
                                if (xhr.status == 204) {
                                    // alert('删除成功');
                                    My_pc_Modal({
                                        text: "删除成功",
                                    })
                                    this_all.remove();
                                    this_all.next('.pk_content').remove();
                                }

                            }
                        })
                    }
                });
            }, complete: function () {
                $(this).removeAttr("disabled");
                LoadingClose();
            }
        });

    } else {
        location.href = "/index"
    }
}

// $('.vip-tab-1').on('click', 'li', function () {
//     var index = $(this).index();
//     $(this).addClass('tab-1-active').siblings().removeClass('tab-1-active');
//     $('.vip-bar-1 .vip-bar-item').eq(index).show().siblings().hide();
// })

$("#person-msg").on("click", ".date", function () {
    var t = $(this);
    var content = t.children('.pk_title').attr('content')

    var title = t.children('.pk_title').text();

    $.Pop(content, {Title: title})

    var state = t.attr('stat');
    var id = Number(t.attr('value'));
    var tstat = t.children('.read-state')

    if (state == 1) {
        ChangeStatus(t, tstat, id)
    }
})

$("#person-msg").on("click", ".del", function (e) {
    e.stopPropagation()
    var id = $(this).parent().parent().attr('value');
    var this_all = $(this).parent().parent();
    id = Number(id);
    var t = $(this).parent().siblings(".pk_title")

    var content = t.attr('content')

    var _title = t.text();

    $.Pop(content, 'confirm', function () {
        $.ajax({
            url: '/ajax/delete/mes',
            type: 'get',
            data: {'id': id},
            headers: {
                'Authorization': 'bearer ' + getCookie('loginBack'),
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'platform': "pc"
            },
            success: function (data, info, xhr) {

                if (data) {
                    if (data.code) {
                        $.Pro(data.msg, {
                            Time: 2, StartOn: function () {
                                $("html").click(function () {
                                    $(".showAlert_Pro").hide()
                                })
                            }
                        })
                    }
                }
                if (xhr.status == 204) {
                    // alert('删除成功');
                    $.Pro("删除成功", {
                        Time: 2, StartOn: function () {
                            $("html").click(function () {
                                $(".showAlert_Pro").hide()
                            })
                        }
                    })
                    this_all.remove();
                }
            }
        })
    })
})


$('.r_name').on('click', 'li', function () {
    var index = $(this).index();
    $(this).addClass('txt_active').siblings().removeClass('txt_active');
    $('.r_center .record_item').eq(index).show().siblings().hide();
    var gen = $(this).attr('gen');
    var th = $(this);
    if (gen == 2) {
        GetNoticeData();
    }
});
$('.nav2 .show-notice-list').click(function () {
    var id = $(this).attr('id');
    $(this).addClass('notice-curr').siblings().removeClass('notice-curr');
    localStorage.setId = id;

    GetNoticeData()
})

function GetNoticeData(page) {
    var id = localStorage.setId;
    id = Number(id);
    $.ajax({
        url: '/ajax/noticeList',
        data: {id: id, page: page, pageSize: 10},
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
            console.log(data);
            var htmldata = '';
            var timestr;
            var metadata = data.meta;
            var links = data.links;
            if (data.data) {
                var showdata = data.data;
                if (showdata) {
                    for (var i = 0; i < showdata.length; i++) {
                        timestr = strToDate(showdata[i].notice_date);
                        htmldata += '<ul class="ul4" stat="' + showdata[i].id + '">';
                        htmldata += '<li>' + showdata[i].notice_title + '</li><li>' + timestr[0] + timestr[1] + '</li><li class="notice-content">' + showdata[i].notice_content + '</li>';
                        htmldata += '</ul>';
                    }
                }

            } else {
                htmldata = ShowNoData('', '暂无任何消息');
            }
            if (metadata.count != 0) {
                htmldata += pageList(metadata, links);
            }
            $('#notice-show').html(htmldata);
            if (metadata.count != 0) {
                var totalPage = metadata.page_count;
                //分页
                GetPageData(GetNoticeData, totalPage);
            }
        }, complete: function () {
            $(this).removeAttr("disabled");
            LoadingClose();
        }
    })
}

function ChangeStatus(t, r, id) {
    // console.log(id)
    $.ajax({
        url: '/mesAjax/mes/status',
        type: 'PUT',
        headers: {
            'Authorization': 'bearer ' + getCookie('loginBack'),
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'platform': "pc"
        },
        dataType: 'json',
        contentType: "application/json",
        data: JSON.stringify({'Id': id}),
        success: function (data, con, code) {
            if (code.status == 204) {
                t.attr('stat', '2');
                r.html('已读');
                r.css('color', '#d1d1d1')
            }
        }
    })
}
$('#notice-show').on("click",".notice-content",function(){
    var content = $(this).text();
    var title = $(this).siblings('li:first-child').text();


    $.Pop(content, {Title: title})
})
