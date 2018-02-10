angular.module('app.membershipReturns').controller('DetailCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService,$stateParams){
    $scope.detailedID = $stateParams.detailedID;
    console.log($scope.detailedID);
    $scope.json = APP_CONFIG.option;
    console.log($scope.json);
    if($scope.status==undefined){
        $scope.status = 2;
    }

    var GetAllEmployee = function () {
        console.log($scope.status);
        var postData = {
            periods_id:$scope.detailedID,
            status:$scope.status
        }
        MembershipReturnsService.getRebateDetail(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.list;
            $scope.name = response.data.product_name;
            $scope.total = response.data.total;
            console.log($scope.total);
        })
    };

   GetAllEmployee();

    $scope.search = function () {
         console.log($scope.status);
        if($scope.status==undefined){
            $scope.status = 2;
        }
        if($scope.status==2){
            $('.detail_C').show();
        }else{
            $('.detail_C').hide();
        }
        GetAllEmployee();
    };

    $scope.doallcheck = function(){
        var allche = document.getElementById("all");
        var che = document.getElementsByClassName("test");

        if(allche.checked == true){
            for(var i=0;i<che.length;i++){
                che[i].checked="checked";
            }
        }else{
            for(var i=0;i<che.length;i++){
                che[i].checked=false;
            }
        }
    };
    $scope.check = function () {
        var check_val = [];
        var test  = document.getElementsByClassName('test');
        for (var j=0;j<test.length;j++){
            if(test[j].checked == true) {
                var obj ={
                    id:test[j].value*1
                };
                check_val.push(obj);
                console.log(test[j].value);
            }

        }
        console.log(check_val);
        var postData= {
            record_ids:check_val
        };
        MembershipReturnsService.getRebateWriteoff(postData).then(function (response) {
            console.log(response);
            if(response.data==null){
                GetAllEmployee();
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
});

