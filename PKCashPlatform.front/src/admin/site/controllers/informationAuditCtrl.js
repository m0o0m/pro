angular.module('app.site').controller('informationAuditCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        //获取站点
        siteService.thirdDropf().then(function (response) {
            $scope.siteJson = response.data.data;
        });
        //获取下拉框类型选择
        $scope.information = APP_CONFIG.option.option_information_type;

        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            }
            siteService.copywariting(postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.data.meta.count;
                    $scope.list = response.data.list;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        //获取详情
        $scope.detail = function (id) {
            siteService.copywaritingDerail(id).then(function (res) {
                $scope.data = res.data;
            })
        };
        //状态修改
        $scope.status=function (item) {
            var postData={
                id:item.id,
                status:item.status
            }
            var del = function () {
                siteService.copywritingStatus(postData).then(function (response) {
                    if(response.data===null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
        };

        //删除
        $scope.del=function (ids) {
            var del = function () {
                siteService.copywaritingDel(ids).then(function (response) {
                    if(response.data===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
        };


    });