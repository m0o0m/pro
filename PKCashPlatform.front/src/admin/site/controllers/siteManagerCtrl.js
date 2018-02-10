angular.module('app.site').controller('siteManagerCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        $scope.option_package = APP_CONFIG.option_package;
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
            siteService.siteManagemnet(postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.data.meta.count;
                    $scope.list = response.data.data;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

       //搜索
       $scope.search = function () {
           GetAllEmployee()
       } ;
       //添加站点
       $scope.add = function () {
         var postData ={
             site_name:$scope.site_name,
             SITEID:$scope.SITEID,
             INDEX_ID:$scope.INDEX_ID,
             domain:$scope.domain,
             agent_domain:$scope.agent_domain,
             Backstage:$scope.Backstage,
             wap_domain:$scope.wap_domain,
             domainCount:$scope.domainCount,
             status:$scope.status,
             packge:$scope.packge,
             typeId:$scope.typeId,
             line:$scope.line,
             Mark:$scope.Mark
         };
         siteService.siteAdd(postData).then(function (data) {
             if(data.data==null){
                 popupSvc.smallBox("success",$rootScope.getWord('success'));
             }else {
                 popupSvc.smallBox("fail",data.msg);
             };
         });
       };
       //获取站点管理详情
         $scope.de = function (item) {
             $scope.datas = item;
         };
       //编辑
        $scope.mod = function () {
            var postData ={
                site_name:$scope.datas.site_name,
                SITEID:$scope.datas.SITEID,
                INDEX_ID:$scope.datas.INDEX_ID,
                domain:$scope.datas.domain,
                agent_domain:$scope.datas.agent_domain,
                Backstage:$scope.datas.Backstage,
                wap_domain:$scope.datas.wap_domain,
                domainCount:$scope.datas.domainCount,
                status:$scope.datas.status,
                packge:$scope.datas.packge,
                typeId:$scope.datas.typeId,
                line:$scope.datas.line,
                Mark:$scope.datas.Mark
            };
            siteService.siteModify(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        }

       //修改额度
        $scope.money = function () {
            var postData = {
                type_id:$scope.type_id,
                Current_quota:$scope.Current_quota,
                Operation:$scope.Operation,
                remark:$scope.content,
                choice:$scope.choice
            };
            siteService.quotaoperation(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };
        //上线
        $scope.online = function () {
            var postData={
                time:$scope.time
            }
            siteService.goOniline(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };




    });