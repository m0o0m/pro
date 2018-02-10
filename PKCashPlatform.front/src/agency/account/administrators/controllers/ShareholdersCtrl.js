angular.module('app.administrators').controller('ShareholdersCtrl',
    function (httpSvc, popupSvc,$stateParams,DTColumnBuilder,$http, $scope,$rootScope, APP_CONFIG,$state) {
        //初始化值
        $scope.account_id = 2;
        $scope.online=0;
        $scope.check=2;
        var site_index  = "";
        //获取JSON
        httpSvc.getJson("/select.json").then(function (data) {
            $scope.json = data[0];
            console.log($scope.json);
        });

        // var user=JSON.parse(resourceSvc.getSession("user"));
        // $scope.isSuperAdmin=user.site_index_id==='';
        if($scope.isSuperAdmin === true){
            //获取站点
            httpSvc.get("/agent/first/drop").then(function (response) {
                console.log(response);
                $scope.siteJson=response.data;
            });
        }else {


        };



        // //获取站点
        // httpSvc.get("/agent/first/drop").then(function (response) {
        //     console.log(response);
        //     $scope.siteJson=response.data;
        // })


        // 更改状态
        $scope.disable=function (status,id) {
            var sure = function () {
                httpSvc.put("/first_agency/status",{
                    id:id
                }).then(function (response) {
                    popupSvc.smallBox("success","更改成功");
                    GetAllEmployee();
                });
            }
            popupSvc.smartMessageBox("确定更改状态?",sure);
        }

        // 踢线
        $scope.kick = function () {
            popupSvc.smartMessageBox("确定踢线？", "踢线成功", "踢线失败");
        };

        //获取name
        if($stateParams.shareid ==null){

            var GetAllEmployee = function () {
                var postData = {
                    page: $scope.paginationConf.currentPage,
                    page_size: $scope.paginationConf.itemsPerPage,
                    isvague: $scope.check,
                    is_online: $scope.online,
                    account_name: $scope.acounted,
                    form_value: site_index,
                    status: $scope.statusd,
                    order_by: $scope.levels,
                    desc: $scope.big
                };

                httpSvc.get("/agent/first", postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }, function (error) {

                })
            }
        }else {
            var GetAllEmployee = function () {
                var postData = {
                    page: $scope.paginationConf.currentPage,
                    page_size: $scope.paginationConf.itemsPerPage,
                    isvague: 0,
                    is_online:$scope.online,
                    account_name:$stateParams.shareid,
                    form_value:site_index,
                    status:$scope.statusd

                };

                httpSvc.get("/agent/first", postData).then(function (response) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }, function (error) {

                })
            }
        }


        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        //点击搜索
        $scope.search = function () {
            if($("#seled").val()==undefined){
                site_index =  user.site_index_id;
            }else {
                site_index = $("#seled").val();
            }

            var checkbox = document.getElementById('test');
            if(checkbox.checked){
                $scope.check=1;
            }else{
                $scope.check=0;
            };
            console.log($scope.acounted);

            httpSvc.get("/agent/first",{
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                isvague: $scope.check,
                is_online:$scope.online,
                account_name:$scope.acounted,
                form_value:site_index,
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
        //add
        $scope.add = function () {
            httpSvc.get("/agent/first/drop").then(function (response) {
                console.log(response);
                $scope.siteJson_1=response.data;
            })
        }
        //新增股东
        $scope.submited = function () {
            if($("#seled_1").val()==undefined){
                $scope.formData.site_index_id =  user.site_index_id;
            }else {
                $scope.formData.site_index_id = $("#seled_1").val();
            }

            httpSvc.post("/first_agency", $scope.formData)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","操作成功");
                        GetAllEmployee();

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                }, function (data) {

                })
        };
        //修改股东账号
        //优惠设定获取ID获取优惠设定;
        $scope.ids="";
        $scope.sitid="";
        $scope.Discount = function (ids,site_index_id) {
            $scope.ids = ids;
            $scope.site_index_id=site_index_id;
            httpSvc.get("/agent/first/discount", {
                acount_id:$scope.ids,
                site_index_id:site_index_id
            }).then(function (response) {
                console.log(response);
                $scope.modifyData = response.data

            }, function (error) {

            })

        };
        $scope.DiscountSubmit = function () {
            $scope.modifyData_1={
                agency_id: $scope.ids,
                site_index_id:$scope.site_index_id,
                offer:$scope.modifyData.offer,
                add_mosaic:$scope.modifyData.add_mosaic,
                is_ip:$scope.modifyData.is_ip,
                site_id:""
            }
            httpSvc.post("/agent/first/discount", {
                agency_id: $scope.ids,
                site_index_id:$scope.site_index_id,
                offer:$scope.modifyData.offer*1,
                add_mosaic:$scope.modifyData.add_mosaic*1,
                is_ip:$scope.modifyData.is_ip*1,
                site_id:""
            })
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","修改成功");
                        GetAllEmployee();

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }

                }, function (data) {

                })
        };
        $scope.modifys={
            account: '',
            username: ''
        };
        $scope.id_m='';
        //修改先获取当前id数据遍历;
        $scope.Modify = function (id) {
            $scope.id_m=id;
            httpSvc.get("/first_agency/info", {
                id:$scope.id_m
            }).then(function (response) {
                console.log(response);
                $scope.modifys.account = response.data.account;
                $scope.modifys.username = response.data.user_name
            }, function (data) {
            })

        };
        //获取数据完成修改后提交
        $scope.modifyssubmit = function () {
            // $scope.modifys.site_index_id = JSON.parse(resourceSvc.getSession("user")).site_index_id;
            $scope.modifys.id = $scope.id_m;
            httpSvc.put("/first_agency", $scope.modifys)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","修改成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                }, function (data) {
                });
        };
        //点击跳转总代
        $scope.Shareholder = function (idsa) {
            $state.go('app.administrators.generalAgent',{
                form_value:idsa
            })
        }
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

