
<!--画面中间，云和国庆祝福，烟花-->
  var cs=document.getElementById('Nation_canvas');

  var ctx=cs.getContext('2d');
  cs.width=1000;
  cs.height=400;


  var lot=new Image();
  lot.src= <{$cdnpcurl}> + '/shared/Nindex/guoqing/image/lot_3.png';


  var cloud1=new Image();
  cloud1.src= <{$cdnpcurl}> + '/shared/Nindex/guoqing/image/cloud1.png';

  var cloud2=new Image();
  cloud2.src= <{$cdnpcurl}> + '/shared/Nindex/guoqing/image/cloud2.png';

  var font=new Image();
  font.src= <{$cdnpcurl}> + '/shared/Nindex/guoqing/image/font.png';






// 画骰子 移动轨迹

// 初始化数据
var pos_top = 0
var pos_left = 0
var pos_right = 660
var lot_pos = $('.box3_pos').position();
var dy = 2;
var dx1 = 1;
var dx2 = 1;
var alpha = 0;

lot.width = 40;
lot.height = 65;

cloud1.width = 340;
cloud1.height = 124;

cloud2.width = 340;
cloud2.height = 124;

// font.width = 392;
// font.height = 226;
setTimeout(drawScreen,2000)



window.requestAnimFrame = (function(){
  return  window.requestAnimationFrame       ||
    window.webkitRequestAnimationFrame ||
    window.mozRequestAnimationFrame    ||
    window.oRequestAnimationFrame      ||
    window.msRequestAnimationFrame     ||
    function( callback ){
      window.setTimeout(callback, 1000 / 60);
    };
})();


    function drawScreen () {
       requestAnimationFrame(drawScreen)

      // ctx.clearRect(lot_pos.left,0,lot.width,cs.height);
      ctx.clearRect(0,0,cs.width,cs.height)

        // 获取骰子的位置
        ctx.drawImage(lot,lot_pos.left+500,pos_top,lot.width,lot.height);

        pos_top+=dy;

        // 画云的位置
        // ctx.save()
        // ctx.globalAlpha=.4;
        ctx.drawImage(cloud1,pos_left,0,cloud1.width,cloud1.height);
        ctx.drawImage(cloud2,pos_right,0,cloud2.width,cloud2.height);
        // ctx.restore()
        // if(pos_left+cloud1.width>cs.width){
        //   dx1=-dx1;
        // }
        if(pos_left+dx1==150){
          dx1=0;
        }

        if(pos_right-dx2==500){
          dx2=0;
        }
        pos_left+=dx1;
        pos_right-=dx2;


    }



function drawFont(){
  requestAnimationFrame(drawFont)

  ctx.save()
  ctx.globalAlpha=alpha;
  ctx.drawImage(font,lot_pos.left+cs.width/3-10,0,font.width,font.height);
  alpha+=0.01;
  if(alpha>=0.6){
    alpha=1;
  }
  ctx.restore()

}

setTimeout(function () {
  drawFont();
  $('#content_bg_NationDay').css('background-color','black');
  drawFireworks();
},5000);





<!--画面底部柱子，箱子-->

  setTimeout(function(){
    var imgFist=$('.section_bottom li img:first-child');

    $(imgFist[0]).attr('src','image/open1.png')
    $(imgFist[1]).attr('src','image/open2.png')
    $(imgFist[2]).attr('src','image/open3.png')
    $(imgFist[3]).attr('src','image/open4.png')
    $(imgFist[4]).attr('src','image/open5.png')

  },5000)





// 绘制烟花



function drawFireworks () {
  var listFire = [];
  var listFirework = [];
    var fireNumber = 15;
  var center = { x: cs.width / 2, y: cs.height / 2 };
  var range = 150;
  for (var i = 0; i < fireNumber; i++) {
    var fire = {
      x: Math.random() *cs.width,
      y: Math.random() * range * 2 + cs.height,
      size: Math.random() + 2,
      fill: '#fd1',
      vx: Math.random() - 0.5,
      vy: -(Math.random() + 4),
      ax: Math.random() * 0.02 - 0.01,
      far: Math.random() * range + (center.y - range)
    };
    fire.base = {
      x: fire.x,
      y: fire.y,
      vx: fire.vx
    };
    //
    listFire.push(fire);
  }

  function randColor() {
    var r = Math.floor(Math.random() * 256);
    var g = Math.floor(Math.random() * 256);
    var b = Math.floor(Math.random() * 256);
    var color = 'rgb($r, $g, $b)';
    color = color.replace('$r', r);
    color = color.replace('$g', g);
    color = color.replace('$b', b);

    return color;
  }

  (function loop() {
    requestAnimationFrame(loop);
    update();
    draw();
  })();

  function update() {
    for (var i = 0; i < listFire.length; i++) {
      var fire = listFire[i];
      //
      if (fire.y <= fire.far) {
        // case add firework
        var color = randColor();
        for (var i = 0; i < fireNumber * 5; i++) {
          var firework = {
            x: fire.x,
            y: fire.y,
            size: Math.random() + 3,
            fill: color,
            vx: Math.random() * 5 - 2.5,
            vy: Math.random() * -5 + 1.5,
            ay: 0.05,
            alpha: 1,
            life: Math.round(Math.random() * range / 2) + range / 2
          };
          firework.base = {
            life: firework.life,
            size: firework.size
          };
          listFirework.push(firework);
        }
        // reset
        fire.y = fire.base.y;
        fire.x = fire.base.x;
        fire.vx = fire.base.vx;
        fire.ax = Math.random() * 0.02 - 0.01;
      }
      //
      fire.x += fire.vx;
      fire.y += fire.vy;
      fire.vx += fire.ax;
    }

    for (var i = listFirework.length - 1; i >= 0; i--) {
      var firework = listFirework[i];
      if (firework) {
        firework.x += firework.vx;
        firework.y += firework.vy;
        firework.vy += firework.ay;
        firework.alpha = firework.life / firework.base.life;
        firework.size = firework.alpha * firework.base.size;
        firework.alpha = firework.alpha > 0.6 ? 1 : firework.alpha;
        //
        firework.life--;
        if (firework.life <= 0) {
          listFirework.splice(i, 1);
        }
      }
    }
  }

  function draw() {
    // clear
    ctx.globalCompositeOperation = 'source-over';
    ctx.globalAlpha = 0.18;
    // ctx.fillStyle = '#000';
    // ctx.fillRect(0, 0, cs.width, cs.height);

    // re-draw
    ctx.globalCompositeOperation = 'screen';
    ctx.globalAlpha = 1;
    for (var i = 0; i < listFire.length; i++) {
      var fire = listFire[i];
      ctx.beginPath();
      ctx.arc(fire.x, fire.y, fire.size, 0, Math.PI * 2);
      ctx.closePath();
      ctx.fillStyle = fire.fill;
      ctx.fill();
    }

    for (var i = 0; i < listFirework.length; i++) {
      var firework = listFirework[i];
      ctx.globalAlpha = firework.alpha;
      ctx.beginPath();
      ctx.arc(firework.x, firework.y, firework.size, 0, Math.PI * 2);
      ctx.closePath();
      ctx.fillStyle = firework.fill;
      ctx.fill();
    }
  }
}


// 页面加载cookie,关闭按钮，自动关闭
  $(document).ready(function(){


    setTimeout(function()
    {
      $("#NationDay").fadeOut();
      $('#content_close_NationDay').fadeOut();
      $("#content_bg_NationDay").fadeOut();
    },18000);
  });

  function guoqinjie(i) {
    if (i == 1) {
      $.cookie('PKBET_NationalDay', 'Q', {path: '/', expires: ''});
    }
    $('#NationDay').remove();
    $('#content_close_NationDay').remove();
    $("#content_bg_NationDay").remove();
  }
  window.onload=function(){
      if ($.cookie('PKBET_NationalDay')){
        $("#NationDay").hide();
        $('#content_close_NationDay').hide();
        $('#content_bg_NationDay').hide();
        return;
      }else{

        $('#NationDay').show();
        $('#content_close_NationDay').show();
        $("#content_bg_NationDay").show();
      }
  };


