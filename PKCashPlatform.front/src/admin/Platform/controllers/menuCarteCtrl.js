angular.module('app.Platform').controller('menuCarteCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state,$stateParams) {
    //获取数据
    var postData = {
        id:$stateParams.idese
    };
    PlatformService.getRoleMenuGet(postData).then(function (response) {
        console.log(response);
        $scope.listd1 = response.data.data.menu_list;
        console.log(response.data.data.menu_list)
    });
//展开收缩数据
    $scope.showFace = function($index,$event){
        if(angular.element($event.target).hasClass("fa-minus-circle")){
            $event.target.className="fa fa-lg fa-plus-circle";
            angular.element($event.target).parent("li").siblings("ul").hide();
            angular.element($event.target).parent("li").siblings("ul").find("ul").hide();
            angular.element($event.target).parent("li").siblings("ul").find("li").find("i").removeClass("fa-minus-circle");
            angular.element($event.target).parent("li").siblings("ul").find("li").find("i").addClass("fa-plus-circle");
        }else{
            $event.target.className="fa fa-lg fa-minus-circle";
            angular.element($event.target).parent("li").siblings("ul").show();
        }
    };
    //点击提交
    $scope.sub = function () {

        var check_val = [];
        $('input[name="checkbox"]:checked').each(function(){
            check_val.push($(this).val()*1);
        });
        var postData = {
            role_id:$stateParams.idese*1,
            menu_id:check_val
        };
        PlatformService.getRoleMenuPost(postData).then(function (response) {
            if(response.data.data===null){
                popupSvc.smallBox("success","操作成功");
                $state.go('app.Platform.role')
            }else{
                popupSvc.smallBox("fail", response.msg);
            }
        });
    };
    //全选
    $scope.selectAll= function ($event) {
        var target = $event.target;
        var isCheck=$(target).is(':checked');
      var a = $(target).parent().parent()[0];
        var a_1 = $(a).parent();
        var a_2 = $(a_1).parent();
        var child = $(a_2).children().children().children().find('.checkbox')
        for(var i=0;i<child.length;i++){
            if(isCheck){
                $(child[i]).prop("checked",true);
            }else {
                $(child[i]).prop("checked",false);
            }
        }
    };
});