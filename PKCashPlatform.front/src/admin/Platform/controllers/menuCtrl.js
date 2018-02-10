angular.module('app.Platform').controller('menuCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state) {
    //菜单列表(平台)
    var GetAllEmployee = function () {
        PlatformService.getMenuAdmin().then(function (response) {
            console.log(response);
            $scope.list = response.data.data;
        })
    };
   GetAllEmployee();

    // //菜单列表(菜单列表（代理）)
    var GetAllEmployeeDl = function () {
        PlatformService.getMenuAgency().then(function (response) {
            console.log(response);
            $scope.listdl = response.data.data;
        })
    };
    GetAllEmployeeDl();

    $scope.add = function () {
        console.log(1122);
        $scope.menu_name = "";
        $scope.route = "";
        $scope.add_language_key = "";
        $scope.sort = "";
        $scope.icon = "";
    };
    $scope.showFace = function($index,$event){
        if(angular.element($event.target).hasClass("fa-minus-circle")){
            $event.target.className="fa fa-lg fa-plus-circle";
            angular.element($event.target).parent("td").parent("tr").siblings("tr").hide();
            angular.element($event.target).parent("td").parent("tr").siblings("tr").find("td").hide();
            angular.element($event.target).parent("td").siblings("tr").find("tr").find("i").removeClass("fa-minus-circle");
            angular.element($event.target).parent("td").siblings("tr").find("tr").find("i").addClass("fa-plus-circle");
        }else{
            $event.target.className="fa fa-lg fa-minus-circle";
            angular.element($event.target).parent("td").parent("tr").siblings("tr").show();
            angular.element($event.target).parent("td").parent("tr").siblings("tr").find("td").show();
        }
    };

    //禁用or启用
    $scope.paySetting=function (item) {
        var able = function () {
            var postData = {
                id: item.id*1,
                status: item.status*1
            };
            PlatformService.getMenuStatus(postData).then(function (response) {
                if (response.data.data === null) {
                    if(item.status==1){
                        item.status = 2;
                    }else{
                        item.status = 1;
                    }
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), able);
    };

    //删除
    $scope.delete = function (id,status) {
        var del = function () {
            var postData = {
                id: id*1,
                status: status*1
            };
            PlatformService.getMenuDel(postData).then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                    GetAllEmployeeDl();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), del);
    };
    //修改平台一级菜单
    $scope.mod_show = 1;
    $scope.modify_1 = function ($index) {
        $scope.mod_show = 1;
        $scope.m1_id = document.getElementsByClassName("m1_id")[$index].innerHTML;
        var postData = {
            id:$scope.m1_id*1,
            type:"admin"
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
        var drop = {
            id:$scope.m1_id,
            type:"admin"
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
    };
    //修改平台二级菜单
    $scope.modify_2 = function($event){
        $scope.test ="12312313";
        $scope.mod_show = 2;
        var m2 = $event.target.parentNode;
        $scope.menuItem = {};
        $scope.m2_id = m2.previousSibling.previousSibling.innerHTML;
        var postData = {
            id:$scope.m2_id,
            type:"admin"
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            $scope.menuItem = response.data;
            $scope.test ="aaaaaa";
        });
        var drop = {
            level:1,
            type:"admin"
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            $scope.select_1 = response.data;
        });
    };
    //修改平台三级菜单
    $scope.modify_3 = function($index,$event){
        $scope.mod_show = 3;
        $scope.m3_id = angular.element($event.target).parent("td").prev(".m3_id").text();
        var postData = {
            id:$scope.m3_id
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
        var drop = {
            level:2,
            type:"admin"
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response.data);
            $scope.select_1 = response.data;
        });
    };
    //平台修改提交
    $scope.submitpt = function () {
        if($scope.mod_show==1){
            $scope.mid = $scope.m1_id;
            $scope.addP = 0;
        }else if($scope.mod_show==2){
            $scope.mid = $scope.m2_id;
            $scope.addP = $scope.menu_select_2;
        }else if($scope.mod_show==3){
            $scope.mid = $scope.m3_id;
            $scope.addP = $scope.menu_select_2;
        }
        console.log($scope.mid);
        var data = {
            id:$scope.mid*1,
            menu_name:$scope.first_data.menu_name,
            route:$scope.first_data.route,
            sort:$scope.first_data.sort*1,
            icon:$scope.first_data.icon,
            type:$scope.first_data.type,
            language_key:$scope.first_data.language_key,
            parent_id:$scope.addP*1
        };
        PlatformService.getMenuPut(data).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    //修改代理一级菜单
    $scope.mod_show = 1;
    $scope.modify_d1 = function ($index) {
        $scope.mod_show = 1;
        $scope.m1_idd = document.getElementsByClassName("m1_idd")[$index].innerHTML;
        console.log($scope.m1_idd);

        var postData = {
            id:$scope.m1_idd
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
        var drop = {
            id:$scope.m1_idd
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response);
        });
    };
    //修改代理二级菜单
    $scope.modify_d2 = function($index,$event){
        $scope.mod_show = 2;
        $scope.m2_idd = angular.element($event.target).parent("td").prev(".m1_iddff").text();
        var postData = {
            id:$scope.m2_idd
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
        var drop = {
            level:1,
            type:"agency"
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response.data);
            $scope.select_1 = response.data;
        });
    }
    //修改代理三级菜单
    $scope.modify_d3 = function($event){
        $scope.mod_show = 3;
        console.log($event.target.parentNode);
        var m3 = $event.target.parentNode;
        $scope.m3_idd = m3.previousSibling.previousSibling.innerHTML;
        var postData = {
            id:$scope.m3_idd
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        });
        var drop = {
            level:2,
            type:"agency"
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response.data);
            $scope.select_1 = response.data;
        });
    };
    $scope.selectM = function () {
        var drop = {
            id:$scope.menu_select_2,
            type:$scope.first_data.type
        };
        PlatformService.getMenuDrop(drop).then(function (response) {
            console.log(response.data);
            $scope.select_2 = response.data;
        });
    };
    //修改提交
    $scope.submit = function () {
        if($scope.mod_show==1){
            $scope.mid = $scope.m1_idd;
            $scope.addP = 0;
        }else if($scope.mod_show==2){
            $scope.mid = $scope.m2_idd;
            $scope.addP = $scope.menu_select_1;
        }else if($scope.mod_show==3){
            $scope.mid = $scope.m3_idd;
            // $scope.select_Menu = $scope.menu_select_1;
            $scope.addP = $scope.menu_select_1;
            // if($scope.select_2==null){
            //     $scope.select_Menu2 = 0;
            // }else{
            //     $scope.select_Menu2 = $scope.menu_select_2;
            // }
        }
        console.log($scope.mid);
        var data = {
            id:$scope.mid*1,
            menu_name:$scope.first_data.menu_name,
            route:$scope.first_data.route,
            // first_id:$scope.select_Menu*1,
            // second_id:$scope.select_Menu2*1,
            sort:$scope.first_data.sort*1,
            icon:$scope.first_data.icon,
            type:$scope.first_data.Type,
            language_key:$scope.first_data.language_key,
            parent_id:$scope.addP*1
        };

        PlatformService.getMenuPut(data).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                // GetAllEmployee();
                GetAllEmployeeDl();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };

    //新增平台菜单
    $scope.addMenu = function () {
        var postData = {
            menu_name: $scope.menu_name,
            sort: $scope.sort * 1,
            route: $scope.route,
            icon: $scope.icon,
            first_id: 0,
            second_id: 0,
            level: 1,
            type: "admin",
            language_key: $scope.add_language_key
        };
        PlatformService.getMenuAdd(postData).then(function (response) {
            // httpSvc.post("/menu",postData).then(function (response) {
            console.log(response);
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                // GetAllEmployee();
                GetAllEmployeeDl();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    //
    //新增代理菜单
    $scope.addMenuDl = function () {
        var postData = {
            menu_name:$scope.menu_name,
            sort:$scope.sort*1,
            route:$scope.route,
            icon:$scope.icon,
            first_id:0,
            second_id:0,
            level:1,
            type:"agency",
            language_key:$scope.add_language_key
        };
        PlatformService.getMenuAdd(postData).then(function (response) {
            console.log(response);
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                // GetAllEmployee();
                GetAllEmployeeDl();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    //平台添加下级菜单
    $scope.addm1 = function ($index) {
        $scope.addname = "";
        $scope.addroute = "";
        $scope.addicon = "";
        $scope.addsort = "";
        $scope.addF_language_key = "";
        $scope.addF_type = "";
        $scope.addF_level = 2;
        console.log($index);
        $scope.m1id = document.getElementsByClassName("m1_id")[$index].innerHTML;
        $scope.mid = $scope.m1id;
        var postData = {
            id:$scope.mid,
            type:"admin"
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.type;
        });
    };
    $scope.addm2 = function ($event) {
        $scope.addname = "";
        $scope.addroute = "";
        $scope.addicon = "";
        $scope.addsort = "";
        $scope.addF_language_key = "";
        $scope.addF_type = "";
        $scope.addF_level = 3;
        console.log($event);
        var m2 = $event.target.parentNode;
        var m2_1 = m2.previousSibling.previousSibling.innerHTML;
        console.log(m2_1);
        var postData = {
            id:m2_1,
            type:"admin"
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.type;
        })
    };
    //
    //代理添加下级菜单
    $scope.addmd1 = function ($index) {
        $scope.addname = "";
        $scope.addroute = "";
        $scope.addicon = "";
        $scope.addsort = "";
        $scope.addF_language_key = "";
        $scope.addF_type = "";
        $scope.addF_level = 2;
        console.log($index);
        $scope.m1id = document.getElementsByClassName("m1_idd")[$index].innerHTML;
        $scope.mid = $scope.m1id;
        var postData = {
            id:$scope.mid
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.type;
        })
    };
    $scope.addmd2 = function ($index,$event) {
        $scope.addname = "";
        $scope.addroute = "";
        $scope.addicon = "";
        $scope.addsort = "";
        $scope.addF_language_key = "";
        $scope.addF_type = "";
        $scope.addF_level = 3;
        console.log($event);
        var m2_1 = angular.element($event.target).parent("td").prev(".m1_iddff").text();
        console.log(m2_1);
        var postData = {
            id:m2_1
        };
        PlatformService.getMenuDetail(postData).then(function (response) {
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.type;
        })
    };

    $scope.add_M = function () {
    console.log($scope.add_id);
        var mData = {
            menu_name:$scope.addname,
            sort:$scope.addsort*1,
            route:$scope.addroute,
            icon:$scope.addicon,
            first_id:$scope.add_id*1,
            second_id:0,
            level:$scope.addF_level,
            type:$scope.addF_type,
            language_key:$scope.addF_language_key,
            parent_id:$scope.add_id
        };
        PlatformService.getMenuAdd(mData).then(function (response) {
            console.log(response);
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
                GetAllEmployeeDl();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }
});