angular.module('app.site').controller('user_navCtrl', function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,CONFIG){
    console.log('zxhsb');

    httpSvc.getJson("/select.json").then(function (data) {
        $scope.json = data[0];
    });
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    //菜单列表(平台)
    var GetAllEmployee = function () {
        httpSvc.get("/menu/list/admin").then(function (response) {
            console.log(response);
            // $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
            console.log($scope.list);
        }, function (error) {

        })
    };
    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


    //菜单列表(菜单列表（代理）)
    var GetAllEmployeeDl = function () {
        httpSvc.get("/menu/list/agency").then(function (response) {
            console.log(response);
            // $scope.paginationConf.totalItems = response.meta.count;
            $scope.listdl = response.data;
            console.log($scope.listdl);
        }, function (error) {

        })
    };
    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployeeDl);

    $scope.add = function () {
        console.log(1122);
        $scope.menu_name = "";
        $scope.route = "";
        $scope.add_language_key = "";
        $scope.sort = "";
        $scope.icon = "";
    }

    //展开收缩数据
    $scope.showFace = function($index,$event){
        if(angular.element($event.target).hasClass("fa-minus-circle")){
            $event.target.className="fa fa-lg fa-plus-circle";
            angular.element($event.target).parent("li").siblings("ul").hide();
            angular.element($event.target).parent("li").siblings("ul").find("ul").hide();
            angular.element($event.target).parent("li").siblings("ul").find("li").find("i").removeClass("fa-minus-circle");
            angular.element($event.target).parent("li").siblings("ul").find("li").find("i").addClass("fa-plus-circle");
        }else{
            $event.target.className="fa fa-lg fa-minus-circle";
            angular.element($event.target).parent("li").siblings("ul").show();
        }
    }

    //禁用or启用
    $scope.paySetting=function (id,status) {
        console.log(id);
        console.log(status);
        var able = function () {
            httpSvc.put("/menu/put",{
                id: id*1,
                status: status*1,
            }).then(function (response) {
                console.log(response);
                if(response===null){
                    if(status===1){
                        popupSvc.smallBox("success","禁用成功");
                        GetAllEmployee();
                        GetAllEmployeeDl();
                    }else{
                        popupSvc.smallBox("success","启用成功");
                        GetAllEmployee();
                        GetAllEmployeeDl();
                    }

                }else{
                    popupSvc.smallBox("fail", response.msg);
                }
            })
        }
        if(status==1){
            popupSvc.smartMessageBox("确定禁用？",able);
        }else{
            popupSvc.smartMessageBox("确定启用？",able);
        }

    };

    //删除
    $scope.delete = function (id) {
        var del = function () {
            httpSvc.del("/menu/delete",{
                id:id*1,
            }).then(function (response) {
                // console.log(1133311);
                console.log(response);
                if(response===null){
                    popupSvc.smallBox("success","删除成功");
                    GetAllEmployee();
                    GetAllEmployeeDl();
                }else{
                    popupSvc.smallBox("fail", response.msg);
                }

            })
        }
        popupSvc.smartMessageBox("确定删除？",del);
    }

    //修改平台一级菜单
    $scope.mod_show = 1;
    $scope.modify_1 = function ($index,$event) {
        $scope.mod_show = 1;
        $scope.m1_id = document.getElementsByClassName("m1_id")[$index].innerHTML;
        console.log($scope.m1_id);
        httpSvc.get("/menu/detail",{
            id:$scope.m1_id,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        }, function (error) {

        })
        httpSvc.get("/menu/drop",{
            id:$scope.m1_id,
        }).then(function (response) {
            console.log(response);
        }, function (error) {

        })
    }
    //修改平台二级菜单
    $scope.modify_2 = function($event){
        $scope.mod_show = 2;
        var m2 = $event.target.parentNode;
        $scope.m2_id = m2.previousSibling.previousSibling.innerHTML;
        console.log(65656);
        // $scope.m2_id = document.getElementsByClassName("m2_id")[$index].innerHTML;
        console.log($scope.m2_id);
        httpSvc.get("/menu/detail",{
            id:$scope.m2_id,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;

        }, function (error) {

        })
        httpSvc.get("/menu/drop",{
            level:1,
            type:"admin"
        }).then(function (response) {
            console.log(response);
            $scope.select_1 = response.data;

        }, function (error) {

        })
    }
    //修改平台三级菜单
    $scope.modify_3 = function($event){
        $scope.mod_show = 3;
        // $scope.m3_id = document.getElementsByClassName("m3_id")[$index].innerHTML;
        // console.log($scope.m3_id)
        var m3 = $event.target.parentNode;
        $scope.m3_id = m3.previousSibling.previousSibling.innerHTML;
        console.log( $scope.m3_id)
        httpSvc.get("/menu/drop",{
            level:2,
            type:"admin"
        }).then(function (response) {
            console.log(response);
            $scope.select_1 = response.data;
        })

        httpSvc.get("/menu/detail",{
            id:$scope.m3_id,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;

        }, function (error) {

        })

    }
    //平台修改提交
    $scope.submitpt = function () {
        if($scope.mod_show==1){
            $scope.mid = $scope.m1_id;
            $scope.addP = 0;
        }else if($scope.mod_show==2){
            console.log(7878951212)
            $scope.mid = $scope.m2_id;
            $scope.addP = $scope.menu_select_2;
        }else if($scope.mod_show==3){
            $scope.mid = $scope.m3_id;
            // $scope.select_Menu = $scope.menu_select_1;
            $scope.addP = $scope.menu_select_2;
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
            language_key:$scope.first_data.LanguageKey,
            parent_id:$scope.addP*1,
        }
        httpSvc.put("/menu",data).then(function (response) {
            console.log(response);
            if(response===null){
                popupSvc.smallBox("success","修改成功");
                GetAllEmployee();
                GetAllEmployeeDl();
            }else{
                popupSvc.smallBox("fail",response.msg);
            }
        }, function (error) {

        })
    }



    //修改代理一级菜单
    $scope.mod_show = 1;
    $scope.modify_d1 = function ($index,$event) {
        $scope.mod_show = 1;
        $scope.m1_idd = document.getElementsByClassName("m1_idd")[$index].innerHTML;
        console.log($scope.m1_idd);
        httpSvc.get("/menu/detail",{
            id:$scope.m1_idd,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;
        }, function (error) {

        })
        httpSvc.get("/menu/drop",{
            id:$scope.m1_idd,
        }).then(function (response) {
            console.log(response);
        }, function (error) {

        })
    }
    //修改代理二级菜单
    $scope.modify_d2 = function($event){
        // console.log($index);
        $scope.mod_show = 2;
        var md = $event.target.parentNode;
        $scope.m2_idd = md.previousSibling.previousSibling.innerHTML;
        // $scope.m2_id = document.getElementsByClassName("m2_id")[$index].innerHTML;
        console.log($scope.m2_idd);
        httpSvc.get("/menu/detail",{
            id:$scope.m2_idd,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;

        }, function (error) {

        })
        httpSvc.get("/menu/drop",{
            level:1,
            type:"agency"
        }).then(function (response) {
            console.log(response);
            $scope.select_1 = response.data;

        }, function (error) {

        })
    }
    //修改代理三级菜单
    $scope.modify_d3 = function($event){
        $scope.mod_show = 3;
        // $scope.m3_id = document.getElementsByClassName("m3_id")[$index].innerHTML;
        // console.log($scope.m3_id)
        console.log($event.target.parentNode);
        var m3 = $event.target.parentNode;
        $scope.m3_idd = m3.previousSibling.previousSibling.innerHTML;
        httpSvc.get("/menu/drop",{
            level:2,
            type:"agency"
        }).then(function (response) {
            console.log(response);
            $scope.select_1 = response.data;
        })

        httpSvc.get("/menu/detail",{
            id:$scope.m3_idd,
        }).then(function (response) {
            console.log(response.data);
            $scope.first_data = response.data;

        }, function (error) {

        })

    }
    $scope.selectM = function () {
        httpSvc.get("/menu/drop",{
            //id:$scope.menu_select_1,
            level:$scope.level,
            Type:$scope.Type
        }).then(function (response) {
            console.log(response);
            $scope.select_2 = response.data;
        })
    }
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
            language_key:$scope.first_data.LanguageKey,
            parent_id:$scope.addP*1,
        }
        httpSvc.put("/menu",data).then(function (response) {
            console.log(response);
            if(response===null){
                popupSvc.smallBox("success","修改成功");
                GetAllEmployee();
                GetAllEmployeeDl();
            }else{
                popupSvc.smallBox("fail",response.msg);
            }
        }, function (error) {

        })
    }



    //新增平台菜单
    $scope.addMenu = function () {
        var postData = {
            menu_name:$scope.menu_name,
            sort:$scope.sort*1,
            route:$scope.route,
            icon:$scope.icon,
            first_id:0,
            second_id:0,
            level:1,
            type:"admin",
            language_key:$scope.add_language_key,
        }
        httpSvc.post("/menu",postData).then(function (response) {
            console.log(response);
            if(response===null){
                popupSvc.smallBox("success","添加成功");
                GetAllEmployee();
                GetAllEmployeeDl();
            }else{
                popupSvc.smallBox("fail",response.msg);
            }
        }, function (error) {

        })
    }

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
            language_key:$scope.add_language_key,
        }
        httpSvc.post("/menu",postData).then(function (response) {
            console.log(response);
            if(response===null){
                popupSvc.smallBox("success","添加成功");
                GetAllEmployee();
                GetAllEmployeeDl();
            }else{
                popupSvc.smallBox("fail",response.msg);
            }
        }, function (error) {

        })
    }

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
        httpSvc.get("/menu/detail",{
            id:$scope.mid,
        }).then(function (response) {
            console.log(response.data);
            // $scope.add_data = response.data;
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.Type;
        }, function (error) {

        })
    }
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
        httpSvc.get("/menu/detail",{
            id:m2_1,
        }).then(function (response) {
            console.log(response.data);
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.Type;

        }, function (error) {

        })
    }

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
        httpSvc.get("/menu/detail",{
            id:$scope.mid,
        }).then(function (response) {
            console.log(response.data);
            // $scope.add_data = response.data;
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.Type;
        }, function (error) {

        })
    }
    $scope.addmd2 = function ($event) {
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
        httpSvc.get("/menu/detail",{
            id:m2_1,
        }).then(function (response) {
            console.log(response.data);
            $scope.add_name = response.data.menu_name;
            $scope.add_id = response.data.id;
            $scope.addF_type = response.data.Type;

        }, function (error) {

        })
    }


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
            parent_id:$scope.add_id,
        }
        httpSvc.post("/menu",mData).then(function (response) {
            console.log(response);
            if(response===null){
                popupSvc.smallBox("success","添加下级成功");
                GetAllEmployee();
                GetAllEmployeeDl();
            }else{
                popupSvc.smallBox("fail",response.msg);
            }
        }, function (error) {

        })
    }



});