
$(document).ready(function() {
	$.each(clients_group, function(i, val) {
		$('.btn-row').hide();
		cardHoverEff(val[0], val[1]);
		noticeHoverEff();
	});

});

//卡牌动画
function cardHoverEff(dom, imgTag) {
	var other_cards_arr = new Array;
	$('.cards-box .card.' + dom).mouseenter(function() {
		$(this).css('margin-top', '0');
		hoverImgEff(dom, imgTag);
		reloadBtnAinimation();
		$('.cards-box .card').not(this).css('margin-top', '65px');
	});
}

//当鼠标移至卡牌，改变初始页面效果
function noticeHoverEff() {
	$('.cards-box .card').mouseenter(function() {
		$('.btn-row').show();
		$('.ele-live-notice').css('height','0');
		$('.ele-live-align').css('background-image', 'url("/theme3/public/shared/video/images/live2/default-bg-blur.jpg")');
	});
}

//鼠标移至卡牌 => 切换
function hoverImgEff(dom, imgTag) {

	$("#egame_go").attr({onclick:"opengeme('','video','"+dom+"')"});
	if(dom == 'dg'){
		$("#egame_rule").attr({href:"https://f.dg99.info/home/rule/cn/index.html"});
		$("#egame_rule").attr({onclick:""});
	}else{
		$("#egame_rule").attr({onclick:"opengeme('','rule','')"});
        //$("#egame_rule").attr({href:"javascript:void(0);"});
		//$("#egame_rule").attr({onclick:"getPager('','rule','')"});
	}
	$thisDom = $('.gallery-picture.' + dom);
	$thisDom.append(imgTag);
	$notThisDom = $('.gallery-picture').not($thisDom);

	$thisDom.css('bottom', '0');
	$notThisDom.css('bottom', '-500px');
	setTimeout("$notThisDom.children().remove();", 250);
}

//鼠标移至卡牌 => 按钮动画效果
function reloadBtnAinimation() {
	var el = $('.btn-row'),
	newone = el.clone(true);
	el.before(newone);
	$("." + el.attr("class") + ":last").remove();
}