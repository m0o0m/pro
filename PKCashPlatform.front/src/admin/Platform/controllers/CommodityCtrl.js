angular.module('app.Platform').controller('CommodityCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state) {
    //获取JSON
    $scope.json = APP_CONFIG.option;

    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
    //获取商品类型
    PlatformService.getProductType().then(function (response) {
        console.log(response);
        $scope.types=response.data.data;
    });
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            product_name:$scope.commodity,
            product_id:$scope.commodity_id,
            status:$scope.statuse,
            title:$scope.typeed
        };
        PlatformService.getProductList(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    };

    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


    //点击搜索
    $scope.search = function () {
        if($scope.type=="product_name"){
            $scope.commodity = $scope.IP_name;
            $scope.commodity_id=""

        }else if($scope.type=="product_id") {
            $scope.commodity_id = $scope.IP_name;
            $scope.commodity=""
        }
        GetAllEmployee();
    };
    //添加
   $scope.addcommodity  = function () {
       var postData = {
           type_id:$scope.add.typeed*1,
           product_name:$scope.add.name,
           api:$scope.add.api,
           status:$scope.add.status
       };
       PlatformService.getProductPost(postData).then(function (response) {
           if (response.data.data === null) {
               popupSvc.smallBox("success", $rootScope.getWord("success"));
               GetAllEmployee();
           } else {
               popupSvc.smallBox("fail", response.data.msg);
           }
       });
   };
    // 更改状态
    $scope.disables=function (item) {
        var able = function () {
            var postData = {
                product_id:item.id
            };
            PlatformService.getProductStatus(postData).then(function (response) {
                if (response.data.data === null) {
                    if(item.status==1){
                        item.status = 2;
                    }else{
                        item.status = 1;
                    }
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), able);
    };
    //获取单个详情
    $scope.getID = function (ids) {
        var postData = {
            product_id:ids
        };
        PlatformService.getProductGet(postData).then(function (response) {
           $scope.modifyes = response.data.data;
        })
    };
    //修改完成点击提交
    $scope.submit = function () {
        var data = {
            type_id:$scope.modifyes.type_id,
            product_id:$scope.modifyes.id,
            product_name:$scope.modifyes.product_name,
            api:$scope.modifyes.api,
            status:$scope.modifyes.status
        };
        PlatformService.getProductPut(data).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
   // 删除
    $scope.del=function (id) {
        var del = function () {
            var postData = {
                product_id:id
            };
            PlatformService.getProductDel(postData).then(function (response) {
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

//点击跳转页面
    $scope.typse = function () {
        $state.go('app.Platform.Managementtype')
    };

});