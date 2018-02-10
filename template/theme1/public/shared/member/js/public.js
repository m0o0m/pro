/*加载动画*/
function Loading() {
    var createDiv = document.createElement("div");
    createDiv.id = "mask";
    createDiv.setAttribute('class', 'mask-loading');
    var bigImg = document.createElement("img");		//创建一个img元素
    bigImg.src = cdnUrl + "/wap/style/images/loading.gif";   //给img元素的src属性赋值
    createDiv.appendChild(bigImg);
    document.body.appendChild(createDiv);
}

/*关闭加载动画*/
function LoadingClose(time) {
    if (!time) {
        time = 500;
    } else {
        time = time * 1000;
    }
    setTimeout(function () {
        $('#mask').remove();
    }, time);
    return;
}


/*********暂无数据展示*********/
function ShowNoData(url, mes,clas='nodate') {

    var divll = document.createElement('div');
    divll.setAttribute('class', clas);
    var showimg = document.createElement('img');
    var showp = document.createElement('p');
    if (!mes) {
        mes = '暂无数据';
    }
    showp.innerHTML = mes;
    showp.style.color='#666';
    showp.style.textAlign='center';
    divll.appendChild(showimg);
    divll.appendChild(showp);
    return divll;
}
//日期处理
//时间戳转日期
function strToDate(data) {
    var nowdate, Y, M, D, h, m, s, data1, data2;
    var datatime ;
    if (data === 0 || data === -1) {
        nowdate = new Date();
        nowdate.setDate(nowdate.getDate() + parseInt(data));
    } else {
        nowdate = new Date(parseInt(data) * 1000);//时间戳转换为日期格式*1000
    }
    // console.log(nowdate);
    Y = nowdate.getFullYear() + '-';
    M = (nowdate.getMonth() + 1 < 10 ? '0' + (nowdate.getMonth() + 1) : nowdate.getMonth() + 1) + '-';
    D = nowdate.getDate() + ' ';
    h = nowdate.getHours() + ':';
    m = nowdate.getMinutes() + ':';
    s = nowdate.getSeconds();
    data1 = Y + M + D;
    data2 = h + m + s;
    datatime = [data1, data2];
    if (data === 0 || data === -1) {
        return data1 + data2;
    } else {
        return datatime;
    }

}

//日期转时间戳
function DateToStr(data) {
    var datatime, time1;
    if (data == 0 || data == -1||data==null||data==-7) {
        datatime = new Date();
        datatime.setHours(0);
        datatime.setMinutes(0);
        datatime.setSeconds(0);
        datatime.setMilliseconds(0);
    } else {
        datatime = new Date(data.replace(/-/g, '/'));// 有三种方式获取，在后面会讲到三种方式的区别
    }
    if (data == -1) {
        time1 = (Date.parse(datatime) / 1000) - 24 * 3600;
    }else if(data==-7){
        time1 = (Date.parse(datatime) / 1000) - 24 * 3600*7;
    }else {
        time1 = Date.parse(datatime) / 1000;
    }
    return time1;
}
//分页
function pageList(metalist, links) {
    console.log(metalist);
    var totalNum = metalist.count;
    var totalPage = metalist.page_count;
    var current_page = metalist.current_page;
    // var page_size = metalist.page_size;
    var htmldata = '<div class="one-foot"><p class="fl">共' + totalNum + '条，共计' + totalPage + '页</p><div class="fr">' +
        '<span class="iconfont icon-back"></span>';
    for (var i = 1; i <= totalPage; i++) {
        if (i == current_page) {
            htmldata += '<span class="one-foot-circle page-active">' + i + '</span>';
        } else {
            htmldata += '<span class="one-foot-circle">' + i + '</span>';
        }
    }
    htmldata += '<span class="iconfont icon-forward"></span>';
    htmldata += '<span>跳转至<select name="select"  class="xla_k ">';

    for (var j = 1; j <= totalPage; j++) {
        if (j == current_page) {
            htmldata += '<option value="'+j+'" selected>'+j+'</option>';
        }else{
            htmldata += '<option value="'+j+'" >'+j+'</option>';
        }

    }
    htmldata += '</select><span>页</span></span></div></div>';
    return htmldata;
}