angular.module('app.administrators').controller('AgentDomainCtrl',
    function(httpSvc, popupSvc, $scope, CONFIG,$stateParams){
        $scope.ids = $stateParams.Domainid;
        if($scope.ids==null){
            var btns =document.getElementsByClassName("btns")[0];
            btns.style.display="none";
            $scope.ids="";
        }

        // 踢线
        $scope.kick=function () {
            popupSvc.smartMessageBox("确定踢线？","踢线成功","踢线失败");
        }


        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                agency_id: $scope.ids,
                domain:$scope.domain
            };

            httpSvc.get("/agent/domain",postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });

        };

        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.modifys={
            domain: '',
            agency_id:"",
            id:""
        };
        //点击修改获取当前id域名
        $scope.Modify = function (seeDomainid,seeID) {
            $scope.modifys.domain_id = seeID;
            httpSvc.get("/agent/domain",{
                domain:seeDomainid
            }).then(function (response) {
                console.log(response.data[0].domain);
                $scope.modifys.domain =response.data[0].domain;
                $scope.modifys.agency_id = response.data[0].agency_id;
                $scope.modifys.id = response.data[0].id

            });

        };
        //修改后点击提交
        $scope.modifyssubmit = function () {
            httpSvc.put("/agent/domain",$scope.modifys)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","修改成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                    // GetAllEmployee();
                    // popupSvc.smallBox("success","修改成功")
                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                });
        };
        // 删除
        $scope.disable=function (mainid) {
           var del = function () {
               httpSvc.del("/agent/domain",{
                   id:mainid
               }).then(function (response) {
                   popupSvc.smallBox("success","删除成功");
                   GetAllEmployee();
               });
           };
            popupSvc.smartMessageBox("确定停用此账户权限?",del);
        };
        //添加
        $scope.adddomain = function () {
            httpSvc.post("/agent/domain", {
                agency_id:$scope.ids,
                domain:$scope.fromData.domain
            })
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","添加成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                    // GetAllEmployee();
                    // popupSvc.smallBox("success","添加成功")
                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                });
        };

        //搜索
        $scope.search = function () {
            httpSvc.get("/agent/domain",{
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                id: $scope.ids,
                domain:$scope.domain

            }).then(function (response) {
                console.log(response);
                if(response.meta.count>0){
                    $scope.paginationConf.totalItems = response.meta.count;
                }else{
                    $scope.paginationConf.totalItems = 0;
                }
                $scope.list = response.data;
            });
        };

        //筛选展开
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
