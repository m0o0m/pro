function urlparent(url){
    window.open(url,"newFrame");
}

function topMouseEvent(mi,ty,i){
    if(ty == "o" && i != disnum){
        mi.className = "homemenua";
    }else if(ty == "t" && i != disnum){
        mi.className = "alink";
    }
}


/***********************************memberUrl******************************/
function memberUrl(url) {
    art.dialog.open(url,{width:960,height:500});
}

function get_dled(){
    $.getJSON("getDLED.php?callback=?",function(json){
        $("#dled").html("("+json.dled+")");
    });
}

function navfocu(i){
    var as = document.getElementById("top_3").getElementsByTagName("a");
    for(var s=0; s<as.length; s++){
        if(s == (i-1)){
            as[s].className = "nav"+i+"_f";
        }else{
            as[s].className = "nav"+(s+1);
        }
    }
}
