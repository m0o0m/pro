angular.module('app.administrators').controller('ChildAccountPermissionsCtrl',
    function(httpSvc,popupSvc,$http,$scope,$stateParams,$state){
        httpSvc.get("/agent/sub/permission",{
            id: $stateParams.id,
            type:1
        })
        .then(function (response) {
            $scope.account=response.data.account;
            $scope.list=response.data.permission_list;
        });

        $scope.selectAll= function ($event) {
            var target = $event.target;
            var isCheck=$(target).is(':checked');
            var list=$(target).parents(".checkall").find('.check');
            for(var i=0;i<list.length;i++){
                if(isCheck){
                    $(list[i]).prop("checked",true);
                }else {
                    $(list[i]).prop("checked",false);
                }
            }
        };

        $scope.submit=function(){
            var status=$('input:radio[name="status"]:checked').val();
            var arr=[];
            var list=$('.check');
            for(var i=0;i<list.length;i++){
                var id=$(list[i]).val();
                if($(list[i]).is(':checked')){
                    arr.push(id*1);
                }
            }
            httpSvc.post("/agent/sub/permission",{
                id: $stateParams.id*1,
                permission_id: arr
            }).then(function (response) {
                //popupSvc.smallBox("success","操作成功");
                if(response===null){
                    popupSvc.smallBox("success","操作成功");
                    $state.go('app.administrators.childAccount');
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        }

        //筛选展开
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
