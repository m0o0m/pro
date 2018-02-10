angular.module('app.Platform').controller('SitecolumnCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
    var GetAllEmployee = function () {
        PlatformService.getSiteCloumnGet().then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta[0].count;
            $scope.list = response.data.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    //新增
    $scope.add = function (id,type) {
        $scope.ctype = type;
        $scope.addId = id;
    };
    $scope.addSub = function () {
        var postData = {
            id: $scope.addId,
            name: $scope.addName,
            url: $scope.addUrl,
            type: $scope.ctype,
            orderby: $scope.addSort*1
        };
        PlatformService.getSiteCloumnPost(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };
    $scope.se = function (item) {
        $scope.name = item.name;
        $scope.typeM = item.type;
        $scope.url = item.url;
        $scope.sort = item.sort;
        $scope.ids = item.id;
    };
    $scope.submit = function () {
        var postData = {
            id: $scope.ids,
            name: $scope.name,
            url: $scope.url,
            type: $scope.typeM,
            orderby: $scope.sort*1
        };
        PlatformService.getSiteCloumnPut(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };
    // 删除
    $scope.del=function (id) {
        var del = function () {
            var postData = {
                id:id
            };
            PlatformService.getSiteCloumnDel(postData).then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), del);
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
    // 全选
    var sels = document.getElementsByClassName('selected');
    $scope.all= function (m) {
        for(var i=0;i<sels.length;i++){
            if(m===true){
                sels[i].checked = true;
            }else {
                sels[i].checked = false;
            }
        }
    };
    $scope.private = function (id) {
        for(var i=0;i<sels.length;i++){
            sels[i].checked = false;
        }
        $scope.privateId = id;
    };
    $scope.privateSub = function () {
        var check_val = [];
        var test  = document.getElementsByClassName('selected');
        for (var j=0;j<test.length;j++){
            if(test[j].checked) {
                var obj ={
                    check_id:test[j].value*1
                };
                check_val.push(obj);
            }
        }
        var postData = {
            id:$scope.privateId,
            check:check_val
        };
        PlatformService.getSiteCloumnPrivate(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    $scope.disables = function () {
        var sure = function () {
            PlatformService.getSiteCloumnSynchro().then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), sure);
    }

});
