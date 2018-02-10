$(document).ready(function(){
  setTimeout(function()
  {
    $(".midAutumn_container").fadeOut();
    $('#content_close_midAutumn').fadeOut();
    $("#content_bg_midAutumn").fadeOut();
  },18000);



});


function zhongqiujie(i) {
  if (i == 1) {
    $.cookie('PKBET_midAutumn', 'Q', {path: '/', expires: ''});
  }
  $('.midAutumn_container').remove();
  $('#content_close_midAutumn').remove();
  $("#content_bg_midAutumn").remove();
}


window.onload=function() {
  if ($.cookie('PKBET_midAutumn')) {
    $(".midAutumn_container").hide();
    $('#content_close_midAutumn').hide();
    $("#content_bg_midAutumn").hide();
    return;
  } else {

    $(".midAutumn_container").show();
    $('#content_close_midAutumn').show();
    $("#content_bg_midAutumn").show();
  }

}