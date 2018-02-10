angular.module('app.administrators').controller('SystemQueryCtrl',
    function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,CONFIG,$state){
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };
        httpSvc.getJson("/select.json").then(function (data) {
            $scope.json=data[0];
        })
        if($scope.account == undefined){
            $scope.account = "";
        }
        if($scope.systemvague == undefined){
            $scope.systemvague = "";
        }
        if($scope.systemlevel == undefined){
            $scope.systemlevel = "";
        }

        //下拉框请求json

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                isvague: $scope.systemvague*1,
                level: $scope.systemlevel*1,
                account: $scope.account
            }
            // httpSvc.get("/agent/search",postData).then(function (response) {
            //     $scope.paginationConf.totalItems = response.meta.count;
            //     $scope.list = response.data;
            // }, function (error) {
            //
            // })
        }
        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        //搜索
        $scope.search = function () {
            $rootScope.searchData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                isvague: $scope.systemvague*1,
                level: $scope.systemlevel*1,
                account: $scope.account
            };

            httpSvc.get("/agent/search",$rootScope.searchData).then(function (response) {
                console.log(response);
                $scope.paginationConf.totalItems = response.meta.count;
                console.log($scope.systemlevel);
                $scope.level = $scope.systemlevel;
                $scope.list = response.data;

            })
        }

        //点击总代理跳转
        $scope.second = function ($index) {
            // console.log($scope.list[$index].agency_id);
            $scope.formvalue = $scope.list[$index].agency_id;
            console.log($scope.formvalue);
            $state.go('app.administrators.generalAgent',{
                form_value:$scope.formvalue
            })
        }

        //点击代理跳转
        $scope.third = function ($index) {
            if($scope.systemlevel==1){
                $scope.first_id = $scope.list[$index].agency_id;
                console.log($scope.first_id);
                $state.go('app.administrators.agent',{
                    first_id:$scope.first_id
                })
            }else if($scope.systemlevel==2){
                $scope.form_value_2 = $scope.list[$index].agency_id;
                console.log($scope.first_id);
                $state.go('app.administrators.agent',{
                    form_value:$scope.form_value_2
                })
            }
            console.log($scope.systemlevel);
            // // console.log($scope.list[$index].agency_id);

        }

        //点击会员跳转
        $scope.member = function ($index) {
            if($scope.systemlevel==1){
                $scope.first_id = $scope.list[$index].agency_id;
                console.log($scope.first_id);
                $state.go('app.administrators.accounts',{
                    first_id:$scope.first_id
                })
            }else if($scope.systemlevel==2){
                $scope.second_id = $scope.list[$index].agency_id;
                console.log($scope.second_id);
                $state.go('app.administrators.accounts',{
                    second_id:$scope.second_id
                })
            }else if($scope.systemlevel==3){
                $scope.agency_id = $scope.list[$index].agency_id;
                $state.go('app.administrators.accounts',{
                    agency_id:$scope.agency_id
                })
            }
            // console.log($scope.list[$index].agency_id);

        }
    });
