$(function(){
	$(document).ready(function(){
		setTimeout(function(){
			$(".yx-img1").addClass("animated bounceInDown").css("opacity","1");
		},1000);
		
		setTimeout(function(){
			$(".yx-img2").addClass("animated fadeInDown").css("opacity","1");
		},2200);
		
		setTimeout(function(){
			$(".yx-img3").addClass("animated rollIn").css("opacity","1");
		},3500);
		
		setTimeout(function(){
			$(".yx-img3").removeClass("animated rollIn");
			$(".yx-img3").addClass("animated flipOutY");
		},7500);
		
		setTimeout(function(){
			$(".yx-img4").addClass("animated zoomIn").css("opacity","1");
		},2800);
		
		setTimeout(function(){
			$(".yx-img5").addClass("animated rubberBand").css("opacity","1");
		},3200);
		
		setTimeout(function(){
			$(".yx-img6").addClass("animated bounceInDown").css("opacity","1");
		},6500);
		
		setTimeout(function(){
			$(".yx-img7").addClass("animated pulse").css("opacity","1");
		},3000);
		
		setTimeout(function(){
			$(".yx-img8").addClass("animated pulse").css("opacity","1");
		},3000);
		setTimeout(function(){
			$(".yx-img9").addClass("animated bounceInUp").css("opacity","1");
		},3000);
		setTimeout(function(){
			$(".yx-colse").addClass("animated bounceInRight").css("opacity","1");
		},1500);
		
		$(".yx-colse").on("click",function(){
			$(".yx-div").remove();
			$(".yx-divBg").remove();
		});
		
		setTimeout(function(){
			$(".yx-div").remove();
			$(".yx-divBg").remove();
		},9000)
	});
});