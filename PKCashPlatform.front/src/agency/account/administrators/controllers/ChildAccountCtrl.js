angular.module('app.administrators').controller('ChildAccountCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,resourceSvc){
        var user=JSON.parse(resourceSvc.getSession("user"));
        var siteList=JSON.parse(resourceSvc.getSession("siteList"));
        $scope.isSuperAdmin=user.site_index_id==='';

        $scope.site_index_id=user.site_index_id;
        if($scope.isSuperAdmin){
            $scope.siteList=siteList;
        }

        var GetAllEmployee = function () {
            var postData = {
                // site_index_id: $scope.site_index_id,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                key: $scope.type,
                value: $scope.searchKey,
                is_login: $scope.lineState
            };
            
            httpSvc.get("/agent/sub/list", postData).then(function (response) {
                if (!response.code) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }
            }, function (error) {

            })

        }

        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


        $scope.openAdd=function(){
            $scope.formData={};
        };
        //点击添加
        $scope.addAcount = function () {
            httpSvc.post("/agent/sub", $scope.formData)
            .then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success","添加成功")
                    GetAllEmployee();
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        };


        $scope.modifyData={
            account: '',
            username: ''
        };

        $scope.id='';
        //点击修改时获取ID;
        $scope.getID = function (index) {
            $scope.id = index;
            console.log(index);
            httpSvc.get("/agent/sub/info",{
                id: $scope.id
            }).then(function (response) {
                $scope.modifyData.account = response.data.account;
                $scope.modifyData.username = response.data.username;
            })
        };
        //修改后提交
        $scope.submit = function () {
            $scope.modifyData.id=$scope.id
            httpSvc.put("/agent/sub", $scope.modifyData)
            .then(function (response) {
                if(response===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success","修改成功")
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        };
        //删除子账号
        $scope.kick=function (ids) {
            var del = function () {
                httpSvc.del("/agent/sub",{
                    id: ids
                }).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","删除成功");
                        GetAllEmployee();
                    }else{
                        popupSvc.smallBox("fail", response.msg);
                    }
                })
            }
            popupSvc.smartMessageBox("确定删除？",del);
        };

        //点击搜索
        $scope.search = function () {
            // $scope.site_index_id=$("#site_index_id")[0].value
            GetAllEmployee()
        }


        // 停用
        $scope.disable=function (status,id) {
            var sure = function () {
                httpSvc.put("/agent/sub/status",{
                    id: id
                }).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","操作成功");
                        GetAllEmployee();
                    }else{
                        popupSvc.smallBox("success","操作失败");
                    }

                });
            };
            popupSvc.smartMessageBox("确定更改状态?",sure);
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

        //子账号口令验证
        httpSvc.get("/agent/subtoken/info").then(function (response) {
            $scope.key = response.data.pass_key;
            $scope.coderadio =  response.data.status;
        });

        //子账号口令提交
        $scope.codeSubmit = function () {
            console.log($scope.coderadio)
            httpSvc.post("/agent/subtoken/status",{
                status:$scope.coderadio*1,
                pass_key:$scope.key
            }).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success","操作成功");
                }else{
                    popupSvc.smallBox("fail","操作失败");
                }
            })
        }

        //获取密钥
        $scope.codeKey = function () {
            httpSvc.get("/agent/subtoken",{
                len:16
            }).then(function (response) {
                console.log(response);
                $scope.key = response.data.key;
            })
        }

    });
