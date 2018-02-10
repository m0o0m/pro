/**
 * Created by apple on 17/11/20.
 */
/**
 * Created by apple on 17/11/20.
 */
angular.module('app.site').controller('ProxynavigationCtrl', function(CONFIG,$scope,httpSvc,$stateParams,popupSvc){
    //获取数据
    httpSvc.getJson("/06.json").then(function (response) {
        $scope.listdl = response.data;
        console.log(response.data.menu_list)
    })

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
        console.log(check_val);
        httpSvc.post("/role/menu",{
            role_id:$stateParams.idese*1,
            menu_id:check_val
        }).then(function (response) {
            if(response===null){
                popupSvc.smallBox("success","操作成功");
                $state.go('app.Platform.role')
            }else{
                popupSvc.smallBox("fail", response.msg);
            }
        })


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

//禁用or启用
    $scope.paySetting=function (id,status) {
        console.log(id);
        console.log(status);
        var able = function () {
            popupSvc.smallBox("success","禁用成功");
        }

            popupSvc.smartMessageBox("确定禁用？",able);




    };


//删除
    $scope.delete = function () {
        var del = function () {
            popupSvc.smallBox("success","删除成功");
        }
        popupSvc.smartMessageBox("确定删除？",del);
    }


    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
});