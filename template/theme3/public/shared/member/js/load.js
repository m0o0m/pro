
// 初始化弹框

;(function($){

    var imgUrl=cdnUrl+'/shared/member/image/Reload.gif';
    var loadDiv='<div class="loadDiv"></div>';
    var img='<img src='+imgUrl+' class="load_img">';


    $.fn.extend({
        //开启加载功能
        "loading":function(value){

            if(value =='screen'){
                $('body').append(loadDiv);
                $('.loadDiv').css({
                    'position':'fixed'
                });

            }else{

                //添加加载div
                this.append(loadDiv);
                $('.loadDiv').css({
                    "position":"absolute"
                })

            }

            $('.loadDiv').append(img);


            $('.loadDiv').css({
                'top':0,
                'left':0,
                "width":"100%",
                "height":"100%",
                'background': 'rgba(78, 74, 74, 0.5)'
            });


            //加载图片样式
            $('.load_img').css({
                "position":"absolute",
                'top':'50%',
                'left':'50%',
                'transform':'translate(-50%,-50%)'
            });

        },

        //关闭加载
        'hideLoad':function(){
            $('.loadDiv').remove();
        }
    })

})(jQuery);