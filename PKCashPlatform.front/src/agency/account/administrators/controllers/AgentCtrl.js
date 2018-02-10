angular.module('app.administrators').controller('AgentCtrl',
    function (httpSvc, popupSvc, $scope, APP_CONFIG, $state, resourceSvc, $stateParams) {
        $scope.account_id = 2;
        $scope.check = 2;
        //获取JSON
        $scope.json = APP_CONFIG.option;

        var user = JSON.parse(resourceSvc.getSession("user"));
        $scope.isSuperAdmin = user.site_index_id === '';

        if ($scope.isSuperAdmin === true) {
            //获取站点
            AccountService.getFirstDropSelect().then(function (response) {
                $scope.siteJson=response;
            });
            //获取站点Id来获取股东下拉框
            $scope.sitId = function (site_index_id) {
                $scope.site_index_id = site_index_id;
                AccountService.getSecondDropSelect(site_index_id)
                .then(function (response) {
                    $scope.sharedJson = response.data;
                });
            };
            $scope.sharedId = function (shared) {
                $scope.shared = shared;
                AccountService.getThirdDropSelect(site_index_id,shared)
                .then(function (response) {
                    console.log(response);
                    $scope.third = response.data;
                });
            };
        }else {

        }
        // 停用
        $scope.disable = function () {
            popupSvc.smartMessageBox("确定停用此账户权限?", "停用成功", "停用失败");
        };
        // 踢线
        $scope.kick = function () {
            popupSvc.smartMessageBox("确定踢线？", "踢线成功", "踢线失败");
        };
        //接收参数
        $scope.idinfo = $stateParams.form_value;
        $scope.shared = $stateParams.first_id;
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                isvague: 1,
                is_online: $scope.online,
                account_name: $stateParams.gene,
                form_value: $scope.idinfo,
                site_index_id: $scope.site_index_id,
                first_id: $scope.shared
            };

            httpSvc.get("/agent/third", postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            }, function (error) {

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
            var checkbox = document.getElementById('test');
            if (checkbox.checked) {
                $scope.check = 1;
            } else {
                $scope.check = 0;
            }
            GetAllEmployee();

        };
        //修改股东账号
        //优惠设定获取ID获取优惠设定;
        $scope.ids = "";
        $scope.sitid = "";
        $scope.Discount = function (ids, sitid) {
            $scope.ids = ids;
            $scope.sitid = sitid;

            httpSvc.get("/agent/third/discount", {
                account_id: $scope.ids,
                site_index_id: $scope.sitid
            }).then(function (response) {
                console.log(response);
                $scope.modifyData = response.data;

            }, function (error) {

            });

        };

        $scope.DiscountSubmit = function () {
            $scope.modifyData_1 = {
                account_id: $scope.ids,
                site_id: $scope.sitid,
                dismoney: $scope.modifyData.discount_money,
                dismultiple: $scope.modifyData.discount_multiple,
                is_limitip: $scope.modifyData.is_limit_ip,
                id: $scope.Id
            };
            httpSvc.post("/agent/third/discount", {
                    site_index_id: $scope.sitid,
                    agency_id: $scope.ids,
                    offer: $scope.modifyData.offer * 1,
                    add_mosaic: $scope.modifyData.add_mosaic * 1,
                    is_ip: $scope.modifyData.is_ip * 1
                })
                .then(function (response) {
                    if (response == null) {
                        console.log(response);
                        GetAllEmployee();
                        popupSvc.smallBox("success", "修改成功")
                    } else {
                        popupSvc.smallBox("fail", response.msg)
                    }

                }, function (data) {
                    popupSvc.smallBox("fail", data.msg)
                });
        };
        //add
        $scope.add = function () {
            //获取站点
            httpSvc.get("/agent/first/drop").then(function (response) {
                console.log(response);
                $scope.siteJson_1_a = response.data;
            });
            $scope.sitId = function (site_index_id_1) {
                $scope.site_index_id = site_index_id_1;
                httpSvc.get("/agent/second/drop", {
                    site_index_id: site_index_id_1,
                }).then(function (response) {
                    console.log(response);
                    $scope.sharedJson_1_a = response.data;
                });
            };
            $scope.sharedId = function (shared_1) {
                $scope.shared = shared_1;
                httpSvc.get("/agent/third/drop", {
                    site_index_id: $scope.site_index_id,
                    first_id: shared_1
                }).then(function (response) {
                    console.log(response);
                    $scope.third_1_a = response.data;
                });
            };

        };

        //新增总代
        $scope.submited = function () {
            $scope.site_index_id_1 = $('#sites').val();
            $scope.idinfo_1 = $('#parent').val();
            httpSvc.post("/third_agency", {
                    site_index_id: $scope.site_index_id_1,
                    account: $scope.formData_1.account,
                    password: $scope.formData_1.password,
                    confirm_password: $scope.formData_1.confirm_password,
                    username: $scope.formData_1.user_name,
                    status: $scope.formData_1.status,
                    parent_id: $scope.idinfo_1 * 1
                })
                .then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", "添加成功");
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                }, function (data) {

                });
        };
        //修改
        $scope.modifys = {
            account: '',
            username: ''
        };
        $scope.id_m = '';
        //修改先获取当前id数据遍历;
        $scope.Modify = function (id) {
            $scope.id_m = id;
            httpSvc.get("/third_agency/info", {
                id: id
            }).then(function (response) {
                console.log(response);
                $scope.modifys = response.data;

            }, function (data) {});

        };
        //获取数据完成修改后提交
        $scope.modifyssubmit = function () {
            httpSvc.put("/third_agency", {
                    site_index_id: JSON.parse(resourceSvc.getSession("user")).site_index_id,
                    id: $scope.id_m,
                    username: $scope.modifys.username,
                    confirm_password: $scope.modifys.confirm_password,
                    password: $scope.modifys.password,
                    account: $scope.modifys.account,
                })
                .then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", "修改成功");
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                }, function (data) {
                    popupSvc.smallBox("fail", data.msg)
                });
        };
        //点击跳转域名
        $scope.Domain = function (Domainid) {
            $state.go('app.administrators.agentDomain', {
                Domainid: Domainid
            });
        };
        //点击修改资料
        $scope.edit = function (editid) {
            console.log(editid);
            $state.go('app.administrators.agentEdit', {
                editid: editid
            });
        };
        // 更改状态
        $scope.disable = function (id) {
            var sure = function () {
                httpSvc.put("/third_agency/status", {
                    id: id
                }).then(function (response) {
                    popupSvc.smallBox("success", "更改成功");
                    GetAllEmployee();
                });
            };
            popupSvc.smartMessageBox("确定更改状态?", sure);
        };
        //点击跳转
        $scope.Proxy = function (ids) {
            $state.go('app.administrators.accounts', {
                agency_id: ids
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