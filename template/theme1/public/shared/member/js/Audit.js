$(function () {
   var logindata= loginId();
   if (logindata){
       $.ajax({
           url:'/ajax/draw/write',
           type:'get',
           data:{},
           headers: {
               'Authorization': 'bearer ' + getCookie('loginBack'),
               'Content-Type': 'application/json',
               'Accept': 'application/json',
               'platform': "pc"
           },
           success:function (data,info,xhr) {
               if(data){
                   if(data.code){
                       alert(data.msg)
                   }
               }
               console.log(xhr)
                if (xhr.status==204){
                    alert('出款成功，请耐心等待审核');
                    location.herf='/member/account';
                }
           }
       })
   }
})