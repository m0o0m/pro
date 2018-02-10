/**
 * Created by apple on 17/9/8.
 */
angular.module('app.administrators').controller('hierarchicalManagCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,$state,resourceSvc){
        var user=JSON.parse(resourceSvc.getSession("user"));

        $scope.isSuperAdmin=user.site_index_id==='';
        if($scope.isSuperAdmin){
            $scope.site_index_id=user.default_site;
            //获取站点
            httpSvc.get("/agent/first/drop").then(function (response) {
                $scope.siteJson=response.data;
                $scope.site_index_id=$scope.siteJson[0].site_index_id;
                GetAllEmployee();
            });
        }else{
            $scope.site_index_id=user.site_index_id;
        }

        //点击搜索
        $scope.search = function () {
            GetAllEmployee();
        };

        //回归
        $scope.regression=function (Id,site_index_id) {
            function f1() {
                httpSvc.put("/member/level/regress",{
                    site_index_id: site_index_id,
                    level_id: Id
                }).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","操作成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail","操作失败");
                    }
                })
            }

            popupSvc.smartMessageBox("确认要将该用户移到未分层?", f1);
        }

        $scope.Id='';
        $scope.SiteId='';

        //修改分层
        $scope.modifyHierarchy=function (Id,site_index_id) {
            $scope.Id=Id;
            httpSvc.get("/member/level/drop",{
                site_index_id: site_index_id
            }).then(function (response) {
                if(response.data){
                $scope.hierarchylist = response.data;
                console.log($scope.hierarchylist);
                //$scope.hierarchylist.shift();
                }
                // console.log($scope.hierarchylist)
            })

        }
        //提交分层
        $scope.modifyHierarchySubmit=function () {
           var id=$("input[name='hierarch']:checked").val();
            httpSvc.put("/member/level/move",{
                site_index_id: $scope.site_index_id,
                move_out: $scope.Id,
                move_in: id
            }).then(function (response) {
                $(".modal-backdrop").hide();
                $("#myModal2").hide();
                GetAllEmployee();
            })
        }

        //开启反水
        $scope.open=function (level_id,is_self_rebate) {
            httpSvc.put("/member/level/selfrebate",{
                level_id: level_id,
                site_index_id: $scope.site_index_id,
                is_self_rebate: is_self_rebate==1?2:1
            }).then(function (response) {
                GetAllEmployee();
            })
        }

        //支付设定
        $scope.paySetting=function (level,pay,site){
            $scope.level_id=level;
            $scope.pay_set_id=pay;
            $scope.site_index_id=site;
        }
        //支付设定确认
        $scope.paySettingSubmit=function () {
            httpSvc.put("/member/level/payset",{
                site_id: user.site_id,
                site_index_id: $scope.site_index_id,
                level_id: $scope.level_id,
                pay_set_id: $("#pay_set_id")[0].value*1
            }).then(function (response) {
                $("#myModal3").hide();
                $(".modal-backdrop").hide();
                popupSvc.smallBox("success","设置成功");
                GetAllEmployee();
            })
        }
        
        // 踢线
        $scope.kick=function () {
            popupSvc.smartMessageBox("确定踢线？","踢线成功","踢线失败");
        }

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id
            };

            httpSvc.get("/member/level/list", postData).then(function (response) {
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

        //点击搜索
        $scope.search = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id
            };

            httpSvc.get("/member/level/list", postData).then(function (response) {
                if (!response.code) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }
            });
        };

        //点击跳转会员详情
        $scope.Membership = function () {
            $state.go('app.administrators.MembershipDetails')
        };
        //新增
        $scope.add = function () {
            $state.go('app.administrators.AddHierarchy')
        };

        // //点击搜索
        // $scope.search = function () {
        //     console.log($scope.site_index_id);
        //     GetAllEmployee();
        // }

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
