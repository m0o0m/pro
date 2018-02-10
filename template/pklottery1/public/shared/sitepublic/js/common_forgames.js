/**
 * Created by Administrator on 2017/1/3.common_forgames
 */

//前端非电子页面指定游戏进入 电子游戏
function gameLink(uid,id,type) {
    if (uid == '') {
        alert('请先登录再进行游戏！');
        return false;
    }
    url = "/dz?v_type="+type
    newWin=window.open(url,'','width=900,height=600,fullscreen=1,scrollbars=0,location=no');
    window.opener=null;//出掉关闭时候的提示窗口
    window.open('','_self'); //ie7
    window.close();

}

//视讯
function opengeme(gameid,url,vtype){
    if (url == '/rule' || url == 'rule') {
        url = '/rule';
    } else if (url == 'video') {//兼容老版本
        url = '/video/login?vType='+ vtype + "&gameId="+gameid;
    }else {
        url = '/video/login?vType='+ vtype + "&gameId="+gameid;
    }

    newWin=window.open(url,'','width=900,height=600,fullscreen=1,scrollbars=0,location=no');
    window.opener=null;//出掉关闭时候的提示窗口
    window.open('','_self'); //ie7
    window.close();
}


    $(function(){
        $("#lotteryhallSelect").on('click','div', function(){
            var calssName = $(this)[0].className
            console.log(calssName);
        })
    })

function openGame(moduleName, vType){
    var loginIn = getCookie('loginBack');
    if(moduleName == 'video'){
        if(loginIn){
            opengeme('','video',vType)
        }else{
            zhuModal.login();
        }
    }else if(moduleName == 'fc'){
        if (vType == 'pk_fc'){
            getPager('-', moduleName, vType);
        }else if(vType == 'eg_fc' || vType == 'cs_fc'){
            if(loginIn){
                opengeme2('', vType);
            }else{
                var config = zhuModalConfig();
                config.DemoHref = 'opengeme("","video","'+vType+'");';
                zhuModal.init({loginConfig:config});
                zhuModal.login("sw");
            }
        }else{
            if(loginIn){
                opengeme('','video',vType)
            }else{
                zhuModal.login();
            }
        }
    }else if(moduleName == 'dz'){
        getPager('-',moduleName, vType);
    }else if(moduleName == 'sp'){
        if(loginIn){
            if(vType == 'bbin_sp'){
                opengeme('','video','bbin');
            }else{
                getPager('-', vType,'m');
            } 
        }else{
            zhuModal.login();
        }
    }
}