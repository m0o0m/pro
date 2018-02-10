
angular.module('app.Platform').controller('SiteCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,$stateParams,resourceSvc){
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        
        $scope.formData={
            site: $stateParams.site
        };
        $scope.modifyData={
            combo_id: ''
        };

        //套餐列表
        httpSvc.get("/combo/drop").then(function (response) {
            $scope.packageList=response.data;
        })

        //获取站点
        $scope.getSite=function(site,site_index){
            $scope.domlist={};
            httpSvc.get("/site/domlist",{
                site:site,
                site_index: site_index
            }).then(function (response) {
                $scope.domlist=response.data;
            })
        }

        //新增
        $scope.add = function () {
            $scope.formData.open_user=$stateParams.id*1;
            $scope.formData.combo_id=$scope.formData.combo_id*1;
            console.log($scope.formData)
            httpSvc.post("/site/add", $scope.formData).then(function (response) {
                if (response===null) {
                    popupSvc.smallBox("success","添加成功");
                    GetAllEmployee();
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        }

        // 停用
        $scope.disable=function (site_id, site_index_id) {
            var disable = function () {
                httpSvc.put("/site/status",{
                    site: site_id,
                    site_index: site_index_id
                }).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","操作成功");
                        GetAllEmployee();
                    }else{
                        popupSvc.smallBox("fail", response.msg);
                    }

                })
            }
            popupSvc.smartMessageBox("确定更改状态?", disable);
        }
        // 删除
        $scope.delete = function (site_id, site_index_id) {
            var del = function () {
                httpSvc.del("/site",{
                    site: site_id,
                    site_index: site_index_id
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
        }

        //修改站点
        $scope.modify=function (site_id, site_index_id) {
            $scope.site_id=site_id;
            $scope.site_index_id=site_index_id;
            httpSvc.get("/site",{
                site: $scope.site_id,
                site_index: $scope.site_index_id
            }).then(function (response) {
                if (!response.code) {
                    $scope.modifyData = response.data
                }
            })

        }
        //提交站点修改
        $scope.modifySubmit=function () {
            $scope.modifyData.site=$scope.site_id;
            $scope.modifyData.site_index=$scope.site_index_id;
            $scope.modifyData.combo_id=$scope.combo*1;
            delete $scope.modifyData.combo_name;
            delete $scope.modifyData.site_id;
            delete $scope.modifyData.site_index_id;
            httpSvc.put("/site",$scope.modifyData).then(function (response) {
                $(".modal-backdrop").hide();
                $("#myModal2").hide();
                if(response===null){
                    popupSvc.smallBox("success","修改成功");
                    GetAllEmployee();
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        }

        //修改优惠
        $scope.dis=function (site_id, site_index_id) {
            $scope.site_id=site_id;
            $scope.site_index_id=site_index_id;
            $scope.discount={
                offer: '',
                add_mosaic: '',
                is_ip: '',
                is_clear: '',
            }
            httpSvc.get("/holder/discount",{
                site: $scope.site_id,
                site_index: $scope.site_index_id
            }).then(function (response) {
                if (!response.code) {
                    $scope.discount = angular.copy(response.data);
                    $scope.discount.is_clear=2;
                }
            })

        }
        //提交优惠修改
        $scope.discountSubmit=function () {
            $scope.discount=angular.extend($scope.discount,{
                site: $scope.site_id,
                site_index: $scope.site_index_id
            })
            $scope.discount.offer=$scope.discount.offer*1;
            $scope.discount.add_mosaic=$scope.discount.add_mosaic*1;
            $scope.discount.is_ip=$scope.discount.is_ip*1;
            $scope.discount.is_clear=$scope.discount.is_clear*1;
            httpSvc.post("/holder/discount", $scope.discount).then(function (response) {
                $(".modal-backdrop").hide();
                $("#myModal2").hide();
                if(response===null){
                    popupSvc.smallBox("success","操作成功");
                    GetAllEmployee();
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        }

        //点击搜索
        $scope.search = function () {
            GetAllEmployee();
        }


        var GetAllEmployee = function () {

            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                open_user: $stateParams.id*1,
                status: $scope.status,
                site_name: $scope.site_name,
                combo_id: $scope.combo_id
            }

            httpSvc.get("/site/openuser", postData).then(function (response) {
                if (!response.code) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }
            })

        }

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);



    });
