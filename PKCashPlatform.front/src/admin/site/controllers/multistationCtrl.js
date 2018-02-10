/**
 * Created by apple on 17/12/18.
 */
/**
 * Created by apple on 17/12/18.
 */
angular.module('app.site').controller('multistationCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService,$stateParams){

        siteService.moduleManagemnet($stateParams.id).then(function (response) {
            $scope.data =response.data.data;
            console.log($scope.data);
        });

        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            }
            siteService.multistation(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.mate.count;
                $scope.list = response.data.data;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        //新增站点
        $scope.add = function () {
            var postData={
                site_name:$scope.site_name,
                SITEID:$scope.SITEID,
                INDEX_ID:$scope.INDEX_ID,
                domain:$scope.domain,
                agent_domain:$scope.agent_domain,
                Backstage:$scope.Backstage,
                wap_domain:$scope.wap_domain,
                domainCount:$scope.domainCount,
                status:$scope.status,
            };
            siteService.multistationAdd(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });

        };
        //获取代理
        $scope.getGaent = function (item) {
            $scope.item = item;
        };
        //添加代理
        $scope.addAgent = function () {
            var postData={
                site_name:$scope.item.INDEX_ID,
                SITEID:$scope.item.SITEID,
                account:$scope.res.account,
                Category:$scope.Category,
                type_id:$scope.type_id,
                remark:$scope.remark
            }
            siteService.multistationAddAgent(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };

        //获取一键生成三级代理详情
        $scope.oneTouch = function (item) {
            $scope.resd = item;
        };
        //添加代理生成
        $scope.level = function () {
            var postData={
                default_name:$scope.default_name,
                default_account:$scope.resd.Module,
                default_remark:$scope.default_remark,
                default_totalgeneration:$scope.default_totalgeneration,
                default_totalgeneration_account:$scope.default_totalgeneration_account,
                default_totalgeneration_remark:$scope.default_totalgeneration_remark,
                default_agent_name:$scope.default_agent_name,
                default_agent_account:$scope.default_agent_account,
                default_agent_remark:$scope.default_agent_remark,
            };
            siteService.oneTochAgent(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            })

        };
        //获取多站点详情
        $scope.detail = function (item) {
            $scope.detail = item;

        };

        //编辑
        $scope.mod = function () {
            var postData={
                site_name:$scope.detail.site_name,
                SITEID:$scope.detail.SITEID,
                INDEX_ID:$scope.detail.INDEX_ID,
                domain:$scope.detail.domain,
                agent_domain:$scope.detail.agent_domain,
                Backstage:$scope.detail.Backstage,
                wap_domain:$scope.detail.wap_domain,
                domainCount:$scope.detail.domainCount,
                status:$scope.detail.status,
            };
            siteService.multistationModify(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };


    });