/* ===========================================================
 * jquery-onepage-scroll.js v1
 * ===========================================================
 * Copyright 2013 Pete Rojwongsuriya.
 * http://www.thepetedesign.com
 *
 * Create an Apple-like website that let user scroll
 * one page at a time
 *
 * Credit: Eike Send for the awesome swipe event
 * https://github.com/peachananr/onepage-scroll
 *
 * ========================================================== */

!function($){
  
  var defaults = {
    sectionContainer: "section",
    easing: "ease",
    animationTime: 1000,
    pagination: true,
    updateURL: false
	};
	
	/*------------------------------------------------*/
	/*  Credit: Eike Send for the awesome swipe event */    
	/*------------------------------------------------*/
	
	$.fn.swipeEvents = function() {
      return this.each(function() {

        var startX,
            startY,
            $this = $(this);

        $this.bind('touchstart', touchstart);

        function touchstart(event) {
          var touches = event.originalEvent.touches;
          if (touches && touches.length) {
            startX = touches[0].pageX;
            startY = touches[0].pageY;
            $this.bind('touchmove', touchmove);
          }
          event.preventDefault();
        }

        function touchmove(event) {
          var touches = event.originalEvent.touches;
          if (touches && touches.length) {
            var deltaX = startX - touches[0].pageX;
            var deltaY = startY - touches[0].pageY;

            if (deltaX >= 50) {
              $this.trigger("swipeLeft");
            }
            if (deltaX <= -50) {
              $this.trigger("swipeRight");
            }
            if (deltaY >= 50) {
              $this.trigger("swipeUp");
            }
            if (deltaY <= -50) {
              $this.trigger("swipeDown");
            }
            if (Math.abs(deltaX) >= 50 || Math.abs(deltaY) >= 50) {
              $this.unbind('touchmove', touchmove);
            }
          }
          event.preventDefault();
        }

      });
    };
	

  $.fn.onepage_scroll = function(options){
    var settings = $.extend({}, defaults, options),
        el = $(this),
        sections = $(settings.sectionContainer)
        total = sections.length,
        status = "off",
        topPos = 0,
        lastAnimation = 0,
        quietPeriod = 500,
        paginationList = "";
    
    $.fn.transformPage = function(settings, pos) {
      $(this).css({
        "-webkit-transform": "translate3d(0, " + pos + "%, 0)", 
        "-webkit-transition": "all " + settings.animationTime + "ms " + settings.easing,
        "-moz-transform": "translate3d(0, " + pos + "%, 0)", 
        "-moz-transition": "all " + settings.animationTime + "ms " + settings.easing,
        "-ms-transform": "translate3d(0, " + pos + "%, 0)", 
        "-ms-transition": "all " + settings.animationTime + "ms " + settings.easing,
        "transform": "translate3d(0, " + pos + "%, 0)", 
        "transition": "all " + settings.animationTime + "ms " + settings.easing
      });
    }

    // 队列抓取动画加判断
    var moveList=function(){
      switch(next.data('index')){
        case 1:
          next.siblings().find('.animateMove,.animateMove1,.animateMove2,.animateMove3').removeClass('left_to_right bottom_to_top inAndOutChange1 inAndOutChange2 inAndOutChange3 fingerMove');
          $("#wapIntroduce_nav").show();
          break;
        case 2:
          next.find(".animateMove").addClass('bottom_to_top');
          next.siblings().find('.animateMove,.animateMove1,.animateMove2,.animateMove3').removeClass('left_to_right bottom_to_top inAndOutChange1 inAndOutChange2 inAndOutChange3 fingerMove');
         $("#wapIntroduce_nav").hide();
          break;
        case 3:
          next.find(".animateMove").addClass('left_to_right');
          next.siblings().find('.animateMove,.animateMove1,.animateMove2,.animateMove3').removeClass('left_to_right bottom_to_top inAndOutChange1 inAndOutChange2 inAndOutChange3 fingerMove');
          $("#wapIntroduce_nav").hide();
          break;
        case 4:
          next.find(".animateMove").addClass('fingerMove');
          next.siblings().find(".animateMove").removeClass('left_to_right bottom_to_top');
          $("#wapIntroduce_nav").hide();
          break;
        case 5:
          next.siblings().find('.animateMove,.animateMove1,.animateMove2,.animateMove3').removeClass('left_to_right bottom_to_top inAndOutChange1 inAndOutChange2 inAndOutChange3 fingerMove');
          $("#wapIntroduce_nav").hide();
          break;

      }
    }

    
    $.fn.moveDown = function() {
      var el = $(this)
      index = $(settings.sectionContainer +".active").data("index");
      if(index < total) {
        current = $(settings.sectionContainer + "[data-index='" + index + "']");
        next = $(settings.sectionContainer + "[data-index='" + (index + 1) + "']");
        if(next) {
          // 滚动下面添加section里面图片效果
          moveList();


          current.removeClass("active")
          next.addClass("active");
          if(settings.pagination == true) {
            $(".onepage-pagination li a" + "[data-index='" + index + "']").removeClass("active");
            $(".onepage-pagination li a" + "[data-index='" + (index + 1) + "']").addClass("active");
          }
          $("body")[0].className = $("body")[0].className.replace(/\bviewing-page-\d.*?\b/g, '');
          $("body").addClass("viewing-page-"+next.data("index"))
          
          if (history.replaceState && settings.updateURL == true) {
            var href = window.location.href.substr(0,window.location.href.indexOf('#')) + "#" + (index + 1);
            history.pushState( {}, document.title, href );
          }
        }
        pos = (index * 100) * -1;
        el.transformPage(settings, pos);
      }
    }
    
    $.fn.moveUp = function() {
      var el = $(this)
      index = $(settings.sectionContainer +".active").data("index");
      if(index <= total && index > 1) {
        current = $(settings.sectionContainer + "[data-index='" + index + "']");
        next = $(settings.sectionContainer + "[data-index='" + (index - 1) + "']");


        if(next) {
          // 滚动上面添加section里面图片效果
          moveList();


          current.removeClass("active")
          next.addClass("active")
          if(settings.pagination == true) {
            $(".onepage-pagination li a" + "[data-index='" + index + "']").removeClass("active");
            $(".onepage-pagination li a" + "[data-index='" + (index - 1) + "']").addClass("active");
          }
          $("body")[0].className = $("body")[0].className.replace(/\bviewing-page-\d.*?\b/g, '');
          $("body").addClass("viewing-page-"+next.data("index"))
          
          if (history.replaceState && settings.updateURL == true) {
            var href = window.location.href.substr(0,window.location.href.indexOf('#')) + "#" + (index - 1);
            history.pushState( {}, document.title, href );
          }

        }
        pos = ((next.data("index") - 1) * 100) * -1;
        el.transformPage(settings, pos);
      }
    }
    
    function init_scroll(event, delta) {
        deltaOfInterest = delta;
        var timeNow = new Date().getTime();
        // Cancel scroll if currently animating or within quiet period
        if(timeNow - lastAnimation < quietPeriod + settings.animationTime) {
            event.preventDefault();
            return;
        }

        if (deltaOfInterest < 0) {
          el.moveDown()
        } else {
          el.moveUp()
        }
        lastAnimation = timeNow;
    }
    
    // Prepare everything before binding wheel scroll
    
    el.addClass("onepage-wrapper").css("position","relative");
    $.each( sections, function(i) {
      $(this).css({
        position: "absolute",
        top: topPos + "%"
      }).addClass("section").attr("data-index", i+1);
      topPos = topPos + 100;
      if(settings.pagination == true) {
        switch(i){
          case 0:
            paginationList += "<li><a data-index='"+(1)+"' href='#" + (1) + "'>1.扫码下载</a></li>";
            break;
          case 1:
            paginationList += "<li><a data-index='"+(2)+"' href='#" + (2) + "'>2.下载中心</a></li>";
            break;
          case 2:
            paginationList += "<li><a data-index='"+(3)+"' href='#" + (3) + "'>3.选择下载</a></li>";
            break;
          case 3:
            paginationList += "<li><a data-index='"+(4)+"' href='#" + (4) + "'>4.安装步骤</a></li>";
            break;
          case 4:
            paginationList += "<li><a data-index='"+(5)+"' href='#" + (5) + "'>5.开始游戏</a></li>";
            break;
        }
      }
    });
    
    el.swipeEvents().bind("swipeDown",  function(){ 
      el.moveUp();
    }).bind("swipeUp", function(){ 
      el.moveDown(); 
    });
    
    // Create Pagination and Display Them
    if(settings.pagination == true) {
      $("<ul class='onepage-pagination'>" + paginationList + "</ul>").prependTo("body");
      posTop = (el.find(".onepage-pagination").height() / 2) * -1;
      el.find(".onepage-pagination").css("margin-top", posTop);
    }
    
    if(window.location.hash != "" && window.location.hash != "#1") {
      init_index =  window.location.hash.replace("#", "")
      $(settings.sectionContainer + "[data-index='" + init_index + "']").addClass("active")
      $("body").addClass("viewing-page-"+ init_index)
      if(settings.pagination == true) $(".onepage-pagination li a" + "[data-index='" + init_index + "']").addClass("active");
      
      next = $(settings.sectionContainer + "[data-index='" + (init_index) + "']");
      if(next) {
        next.addClass("active")
        if(settings.pagination == true) $(".onepage-pagination li a" + "[data-index='" + (init_index) + "']").addClass("active");
        $("body")[0].className = $("body")[0].className.replace(/\bviewing-page-\d.*?\b/g, '');
        $("body").addClass("viewing-page-"+next.data("index"))
        if (history.replaceState && settings.updateURL == true) {
          var href = window.location.href.substr(0,window.location.href.indexOf('#')) + "#" + (init_index);
          history.pushState( {}, document.title, href );
        }
      }
      pos = ((init_index - 1) * 100) * -1;
      el.transformPage(settings, pos);
      
    }else{
      $(settings.sectionContainer + "[data-index='1']").addClass("active")
      $("body").addClass("viewing-page-1")
      if(settings.pagination == true) $(".onepage-pagination li a" + "[data-index='1']").addClass("active");
    }
    if(settings.pagination == true)  {
      $(".onepage-pagination li a").click(function (){
        var page_index = $(this).data("index");
        if (!$(this).hasClass("active")) {
          current = $(settings.sectionContainer + ".active")
          next = $(settings.sectionContainer + "[data-index='" + (page_index) + "']");
          if(next) {
            //点击li里面a效果
            moveList();

            current.removeClass("active")
            next.addClass("active")
            $(".onepage-pagination li a" + ".active").removeClass("active");
            $(".onepage-pagination li a" + "[data-index='" + (page_index) + "']").addClass("active");
            $("body")[0].className = $("body")[0].className.replace(/\bviewing-page-\d.*?\b/g, '');
            $("body").addClass("viewing-page-"+next.data("index"))
          }
          pos = ((page_index - 1) * 100) * -1;
          el.transformPage(settings, pos);
        }
        if (settings.updateURL == false) return false;
      });
    }
    
    
    
    $(document).bind('mousewheel DOMMouseScroll', function(event) {
      event.preventDefault();
      var delta = event.originalEvent.wheelDelta || -event.originalEvent.detail;
      init_scroll(event, delta);
    });
    return false;
    
  }
  
}(window.jQuery);


