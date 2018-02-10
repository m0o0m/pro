angular.module('app.membershipReturns').controller('PromotionSettingCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService){
        var extend = document.getElementsByName('extend');
        var ip = document.getElementsByName('ip');
        var users = document.getElementsByName('user');
        var code = document.getElementsByName('code');
    $scope.sitId = function (site_index_id) {
        MembershipReturnsService.getSiteSelect(site_index_id).then(function (response) {
            $scope.names = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        $scope.sitId(user.site_index_id);
    }

        $scope.search = function (site) {
            var postData = {
                site_index_id:site
            }
            MembershipReturnsService.getSpreadList(postData).then(function (response) {
                console.log(response);
                if (response.data == null) {
                    $scope.requst = "post";
                    $("input:radio[name='extend']").removeAttr("checked");
                    $("input:radio[name='ip']").removeAttr("checked");
                    $("input:radio[name='user']").removeAttr("checked");
                    $("input:radio[name='code']").removeAttr("checked");
                    $scope.people = "";
                    $scope.money = "";
                } else {
                    $scope.requst = "put";
                    var value1 = response.data[0].is_open;
                    var value2 = response.data[0].is_ip;
                    var value3 = response.data[0].is_mate_agency;
                    var value4 = response.data[0].is_code;
                    $scope.people = response.data[0].ranking_num;
                    $scope.money = response.data[0].ranking_money;
                    console.log(user);
                    for (var i = 0; i < 2; i++) {
                        if (extend[i].value == value1) {
                            extend[i].checked = 'checked';
                        }
                        if (ip[i].value == value2) {
                            ip[i].checked = 'checked';
                        }
                        if (users[i].value == value3) {
                            users[i].checked = 'checked';
                        }
                        if (code[i].value == value4) {
                            code[i].checked = 'checked';
                        }
                    }
                }
            })
        }
        $scope.submit = function () {
            console.log($scope.requst);
            console.log($scope.site_index_id);
            console.log($scope.money);
            console.log($scope.people);
            $scope.users = $("input[name='user']:checked").val();
            $scope.code = $("input[name='code']:checked").val();
            $scope.ip = $("input[name='ip']:checked").val();
            $scope.extend = $("input[name='extend']:checked").val();

            var registsetting = {
                ranking_money:$scope.money*1,
                ranking_num:$scope.people*1,
                is_mate_agency:$scope.users*1,
                is_code:$scope.code*1,
                is_ip:$scope.ip*1,
                is_open:$scope.extend*1,
                site_index_id:$scope.site_index_id,
            }
            if($scope.money==''||$scope.people==''||$scope.users==undefined||$scope.code==undefined||$scope.ip==undefined||$scope.extend==undefined||$scope.site_index_id==undefined){
                popupSvc.smallBox("fail","请输入完整！")
            }else{
                if($scope.requst == 'post'){
                    MembershipReturnsService.getSpreadAdd(registsetting).then(function (response) {
                        console.log(response);
                        if(response.data==null){
                            popupSvc.smallBox("success", $rootScope.getWord("success"));
                        } else {
                            popupSvc.smallBox("fail", response.data.msg);
                        }
                    })
                }else{
                    MembershipReturnsService.getSpreadEdit(registsetting).then(function (response) {
                        console.log(response);
                        if(response.data==null){
                            popupSvc.smallBox("success", $rootScope.getWord("success"));
                        } else {
                            popupSvc.smallBox("fail", response.data.msg);
                        }
                    })
                }

            }
        }
});
