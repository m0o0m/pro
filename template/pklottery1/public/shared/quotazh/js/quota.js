$(document).ready(function() {
	$("p.flipquota").click(function() {
		var loginIn;
        loginIn = getCookie('loginBack');
        if(!loginIn){
            zhuModal.login();
        }else{
        	var imgPath = $("p.flipquota").css("background-image");
			$(".panel").slideToggle("fast");
			if(imgPath.indexOf("transfer-bar-blue.png") > 0) {
				$("p.flipquota").css("background","url("+cdnUrl+"/shared/quotazh/images/blue.png?v="+jsVersion+") 0 no-repeat");
			}else{
		   		$("p.flipquota").css("background","url("+cdnUrl+"/shared/quotazh/images/blue.png?v="+jsVersion+") 0 no-repeat");
			}
        }
	});
	//if ( $(window).height() < 10000 ) {
	    $('.input-container > .radio-tile > .icon').remove();
	    $('.input-container > .radio-tile').addClass('s-label');
  	//}
});
