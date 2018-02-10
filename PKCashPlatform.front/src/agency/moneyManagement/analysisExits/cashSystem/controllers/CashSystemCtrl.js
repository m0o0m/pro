angular.module('app.CashSystem').controller('CashSystemCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,CashSystemService){
    $scope.sitId = function (site_index_id) {
        CashSystemService.getSiteSelect(site_index_id).then(function (response) {
            $scope.names = response.data.data;
        });
    };
    var user=JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin=user.site_index_id==='';
    if($scope.isSuperAdmin){
        //获取站点
        $scope.sitId();
    }else{
        $scope.site_index_id=user.site_index_id;
    }
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        if($scope.start_time===null){
            $scope.start_time="";
        }
        if($scope.end_time===null){
            $scope.end_time="";
        }
        if($scope.order_num===null){
            $scope.order_num="";
        }
        if($scope.refresh===null){
            $scope.refresh="";
        }
        if($scope.type===null){
            $scope.type="";
        }
        if($scope.date===null){
            $scope.date="";
        }
        if($scope.account===null){
            $scope.account="";
        }
        var postData = {
            site_index_id:$scope.site_index_id,
            start_time:$scope.start_time,
            end_time:$scope.end_time,
            type:$scope.type,
            date:$scope.date,
            account:$scope.account,
            refresh:$scope.refresh,
            order_num:$scope.order_num,
        };
        CashSystemService.getCashSystem(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
            $scope.total = response.total;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 20
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.doallcheck = function(){
        var allche = document.getElementById("all");
        var che = document.getElementsByClassName("test");
        if(allche.checked === true){
            for(var i=0;i<che.length;i++){
                che[i].checked="checked";
            }
        }else{
            for(var i=0;i<che.length;i++){
                che[i].checked=false;
            }
        }
    };

    // $(".test :checkbox").click(function(){
    //     var chknum = $(".test :checkbox").size();//选项总个数
    //     console.log(chknum);
    //     var chk = 0;
    //     $(".test :checkbox").each(function () {
    //         if($(this).attr("checked")==true){
    //
    //             chk++;
    //         }
    //     });
    //     if(chknum==chk){//全选
    //         $(".test").attr("checked",true);
    //     }else{//不全选
    //         $(".test").attr("checked",false);
    //     }
    // });


    // $scope.noallcheck = function () {
    //     var allche = document.getElementById("all");
    //     var che = document.getElementsByClassName("test");
    //     for(var k=0;k<che.length;k++){
    //         console.log(che[k].checked);
    //         if(che[k].checked == false){
    //             console.log(112233);
    //             console.log(allche.checked);
    //             $("#all").attr("checked",false);
    //         }else{
    //             console.log(565656);
    //             $("#all").attr("checked",true);
    //         }
    //     }
    // };

    $scope.check = function () {
        var check_val = [];
        var test  = document.getElementsByClassName('test');
        for (var j=0;j<test.length;j++){
            if(test[j].checked) {
                var obj ={
                    id:test[j].value*1
                };
                check_val.push(obj);
            }

        }
        console.log(check_val);
        var postData = {
            delete_id:check_val
        };
        CashSystemService.delCashSystem(postData).then(function (response) {
            console.log(response);
            if(response.data===null){
                GetAllEmployee();
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
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


