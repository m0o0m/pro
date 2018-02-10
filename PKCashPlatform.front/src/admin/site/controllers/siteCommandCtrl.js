angular.module('app.site').controller('siteCommandCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        //获取站点
        siteService.thirdDropf().then(function (response) {
            $scope.siteJson = response.data.data;
        });
        $scope.option_onOff = APP_CONFIG.option.option_onOff

        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            };
            siteService.sitepassword(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.mate.cunt;
                $scope.list =response.data.list;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };
        //获取单个详情
        $scope.detali = function (site_id) {
          siteService.sitepasswordDetail(site_id).then(function (res) {
              $scope.data = res.data;
          }) ;
        };
        //修改
        $scope.sumbit = function () {
            var postData={
                site_id:$scope.data.site_index_id,
                status:$scope.data.status,
                key:$scope.data.key
            };
            siteService.sitepasswordModify(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };

    });