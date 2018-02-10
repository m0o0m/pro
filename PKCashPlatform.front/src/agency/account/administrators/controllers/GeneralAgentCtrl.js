angular.module('app.administrators').controller('GeneralAgentCtrl',
    function(httpSvc, popupSvc,$stateParams,DTColumnBuilder,$http, $scope,$rootScope, APP_CONFIG,$state){
        $scope.account_id=6;
        $scope.check=2;
        $scope.site_index_id="";
        console.log($stateParams.form_value);
        //获取JSON

        httpSvc.getJson("/select.json").then(function (data) {
            $scope.json = data[0];
            console.log($scope.json);
        });
        $scope.sitId = function (site_index_id) {
            httpSvc.get("/agent/second/drop",{
                site_index_id:site_index_id,
            }).then(function (response) {
                console.log(response);
                $scope.sharedJson=response.data;
            });
        }
        // var user=JSON.parse(resourceSvc.getSession("user"));
         $scope.isSuperAdmin=user.site_index_id==='';
        if($scope.isSuperAdmin === true){
            //获取站点
            httpSvc.get("/agent/first/drop").then(function (response) {
                console.log(response);
                $scope.siteJson=response.data;
            });
        }else {

        }

        //接收参数
        if($stateParams.form_value==null&&$stateParams.Shareholderid==null){

            var GetAllEmployee = function () {
                var postData = {
                    page: $scope.paginationConf.currentPage,
                    page_size: $scope.paginationConf.itemsPerPage,
                    form_value:$scope.idinfo,
                    isvague: $scope.check,
                    is_online:$scope.online,
                    account_name:$scope.acounted,
                    site_index_id:$scope.site_index_id,
                    status:$scope.statusd,
                    order_by:$scope.levels,
                    desc:$scope.big
                };

                httpSvc.get("/agent/second", postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }, function (error) {

                });
            };
        }else if($stateParams.form_value!==null) {
            console.log(23456);
            var GetAllEmployee = function () {
                var postData = {
                    page: $scope.paginationConf.currentPage,
                    page_size: $scope.paginationConf.itemsPerPage,
                    form_value:$stateParams.form_value,
                    isvague: 1,
                    is_online:$scope.online,
                    account_name:$scope.acounted,
                    site_index_id:$scope.site_index_id,
                    status:$scope.statusd,
                    order_by:$scope.levels,
                    desc:$scope.big
                };

                httpSvc.get("/agent/second", postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }, function (error) {

                })
            };
        }else if($stateParams.Shareholderid!==null){
            var GetAllEmployee = function () {
                var postData = {
                    page: $scope.paginationConf.currentPage,
                    page_size: $scope.paginationConf.itemsPerPage,
                    isvague: $scope.check,
                    is_online:$scope.online,
                    account_name:$scope.acounted,
                    form_value:$stateParams.Shareholderid,
                    site_index_id:$scope.site_index_id,
                    status:$scope.statusd,
                    order_by:$scope.levels,
                    desc:$scope.big
                };

                httpSvc.get("/agent/second", postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }, function (error) {

                });
            };
        }

        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        //点击搜索
        $scope.search = function () {
            $scope.idinfo = $("#seled_1").val();
            $scope.site_index_id = $("#site").val();
            var checkbox = document.getElementById('test');
            if(checkbox.checked){
                $scope.check=1;
            }else{
                $scope.check=0;
            };
            console.log($scope.acounted);
            httpSvc.get("/agent/second",{
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                isvague: $scope.check,
                is_online:$scope.online,
                account_name:$scope.acounted,
                form_value:$scope.idinfo,
                site_index_id:$scope.site_index_id,
                status:$scope.statusd,
                order_by:$scope.levels,
                desc:$scope.big
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
        //修改股东账号
        //优惠设定获取ID获取优惠设定;
        $scope.ids="";
        $scope.sitid="";
        $scope.Discount = function (ids,sitid) {
            $scope.ids = ids;
            $scope.sitid=sitid;

            httpSvc.get("/agent/second/discount", {
                acount_id:$scope.ids,
                site_index_id:$scope.sitid
            }).then(function (response) {
                console.log(response);
                $scope.modifyData = response.data
            }, function (error) {

            })

        };
        $scope.DiscountSubmit = function () {
            $scope.modifyData_1={
                account_id: $scope.ids,
                site_id:$scope.sitid,
                dismoney:$scope.modifyData.discount_money,
                dismultiple:$scope.modifyData.discount_multiple,
                is_limitip:$scope.modifyData.IsLimitIp,
                id: $scope.Id
            };
            httpSvc.post("/agent/second/discount", {
                site_index_id:$scope.sitid,
                agency_id:$scope.ids,
                offer:$scope.modifyData.offer*1,
                add_mosaic:$scope.modifyData.add_mosaic*1,
                is_ip:$scope.modifyData.is_ip*1
            })
                .then(function (response) {
                    GetAllEmployee();
                    popupSvc.smallBox("success","修改成功")
                    // if(response.data==1){
                    //
                    // }else {
                    //
                    // }

                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                })
        };
        //add
        $scope.add = function () {
            //获取站点
            httpSvc.get("/agent/first/drop").then(function (response) {
                console.log(response);
                $scope.siteJson_1=response.data;
            });
            //获取站点Id来获取股东下拉框
            $scope.parent = function (site_index_id) {
                httpSvc.get("/agent/second/drop",{
                    site_index_id:site_index_id,
                }).then(function (response) {
                    console.log(response);
                    $scope.sharedJson_1=response.data;
                });
            }
        }

        //新增总代
        $scope.submited = function () {
            $scope.formData.parent_id = $("#parent").val();
            $scope.formData.site_index_id = $("#sitindex").val();
            httpSvc.post("/second_agency", {
                site_index_id:$scope.formData.site_index_id,
                confirm_password:$scope.formData.confirm_password,
                account:$scope.formData.account,
                password:$scope.formData.password,
                parent_id:$scope.formData.parent_id*1,
                status:$scope.formData.status,
                username:$scope.formData.user_name
            })
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","添加成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                    // console.log(response);
                    // GetAllEmployee();
                    // popupSvc.smallBox("success","添加成功")
                }, function (data) {

                })
        };
        //修改
        $scope.modifys={
            account: '',
            username: ''
        };
        $scope.id_m='';
        //修改先获取当前id数据遍历;
        $scope.Modify = function (id) {
            $scope.id_m=id;
            httpSvc.get("/second_agency/info",{
                id:id
            }).then(function (response) {
                console.log(response);
                $scope.modifys.account = response.data.account;
                $scope.modifys.username = response.data.username
            }, function (data) {
            })

        };
        //获取数据完成修改后提交
        $scope.modifyssubmit = function () {
            $scope.modifys.site_index_id = JSON.parse(resourceSvc.getSession("user")).site_index_id;
            $scope.modifys.id = $scope.id_m*1;
            httpSvc.put("/second_agency", $scope.modifys)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","修改成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                })
        };
        // 更改状态
        $scope.disable=function (id) {
            var sure = function () {
                httpSvc.put("/second_agency/status",{
                    id:id
                }).then(function (response) {
                    popupSvc.smallBox("success","修改成功");
                    GetAllEmployee();
                });
            }
            popupSvc.smartMessageBox("确定修改?",sure);
        };
        //点击跳转
        $scope.General = function (ids) {
            $state.go('app.administrators.agent',{
                form_value:ids
            });
        };
        //点击跳转上级
        $scope.Share = function (Shareid) {
            $state.go('app.administrators.shareholders',{
                shareid:Shareid
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
