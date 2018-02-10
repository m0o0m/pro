/**
 * Created by Administrator on 2017/1/3.common_static
 */

/*************************************ajax获取公告***********************************/
function noticeType(path){
    if (path.indexOf('N_index') > 0)
        return 1;
    else if (path.indexOf('lottery') > 0)
        return 2;
    else if (path.indexOf('livetop') > 0)
        return 3;
    else if (path.indexOf('youhui') > 0)
        return 4;
    else if (path.indexOf('egame') > 0)
        return 7;
    else if (path.indexOf('sports') > 0)
        return 8;
    else if (path.indexOf('iword') > 0)
        return 10;
    else
        return 0;
}

$(function () {
     //获取用户u的参数
    var top_u = window.top.location.search;
    if(top_u){
        setCookie('top_u',top_u.substr(3));
    }
    // if($('#bulletinMsg').length > 0){
    //     var pathName = window.location.pathname;
    //     var $attr = $('#bulletinMsg').attr('data');
    //     var parame ={};
    //     parame.type = noticeType(pathName);
    //     parame.isUp = $attr;
    //     $.get("/notice",parame,function(data){
    //         $('#bulletinMsg').html(data);
    //     });
    // }
});

/*************************************END ajax获取公告***********************************/

/***********************************动态显示美东时间******************************/
var mddate = '';
var dd2 = 0;
function getMdTime() {
    $.get("/index.php/Index/getMdTime",{},function(data){
        mddate = data;
        dd2 = new Date(mddate);
        setInterval("RefTime()",1000);
    });
}

//美东时间
function RefTime(){
    dd2.setSeconds(dd2.getSeconds()+1);
    var myYears = ( dd2.getYear() < 1900 ) ? ( 1900 + dd2.getYear() ) : dd2.getYear();
    $("#vlock").html('美東時間'+'：'+myYears+'年'+fixNum(dd2.getMonth()+1)+'月'+fixNum(dd2.getDate())+'日 '+time(dd2));
}

function time(vtime){
    var s='';
    var d=vtime!=null?new Date(vtime):new Date();
    with(d){
        s=fixNum(getHours())+':'+fixNum(getMinutes())+':'+fixNum(getSeconds())
    }
    return(s);
}

function fixNum(num){
    return parseInt(num)<10?'0'+num:num;
}

/***********************************END 动态显示美东时间******************************/