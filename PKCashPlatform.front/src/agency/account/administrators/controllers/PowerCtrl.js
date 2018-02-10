angular.module('app.administrators').controller('PowerCtrl', function(DTOptionsBuilder,DTColumnBuilder,$http,$scope,$compile,APP_CONFIG,$state){
    var vm = this;
    activate();
    function activate() {
        vm.name1 = "";

        vm.page = function () {
            vm.pageSize = 1;
            //分页数
            vm.pages = Math.ceil(vm.count / vm.pageSize);
            console.log(vm.pages);
            vm.newPages = vm.pages > 5 ? 5 : vm.pages;
            console.log(vm.newPages);
            vm.pageList = [];
            vm.selPage = 1;
            //设置表格数据源(分页)
            vm.setData = function () {
                vm.items = vm.dataJson.slice((vm.pageSize * (vm.selPage - 1)), (vm.selPage * vm.pageSize));//通过当前页数筛选出表格当前显示数据
                console.log(vm.items);
            }
            vm.items = vm.dataJson.slice(0, vm.pageSize);
            //分页要repeat的数组
            for (var i = 0; i < vm.newPages; i++) {
                vm.pageList.push(i + 1);
            }
            //打印当前选中页索引
            vm.selectPage = function (page) {
                // vm.selPage = vm.selPage+1;
                console.log(vm.selPage);
                $http({
                    url:APP_CONFIG.apiRootUrl_2+'manager-group/'+'?page='+page+'&pageSize='+1+'&count='+true,
                    method:"get"
                }).success(function (data) {
                    console.log(data);
                    vm.count = data.data.count;
                    vm.dataJson = data.data.data;
                    console.log(vm.dataJson);
                    for (var i = 0; i < vm.dataJson.length; i++) {
                        if(vm.dataJson[i].Status==0){
                            vm.dataJson[i].Status = "禁用"
                        }else if(vm.dataJson[i].Status==1){
                            vm.dataJson[i].Status = "启用"
                        }
                    }
                    //转化时间
                    for (var i = 0; i < vm.dataJson.length; i++){
                        function getLocalTime(nS) {
                            return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
                        };
                        vm.dataJson[i].CreatTime = getLocalTime(vm.dataJson[i].CreatTime);

                    };
                    if (page < 1 || page > vm.pages) return;
                    //最多显示分页数5
                    if (page > 2) {
                        //因为只显示5个页数，大于2页开始分页转换
                        var newpageList = [];
                        for (var i = (page - 3) ; i < ((page + 2) > vm.pages ? vm.pages : (page + 2)) ; i++) {
                            newpageList.push(i + 1);
                        }
                        vm.pageList = newpageList;
                    }
                    vm.selPage = page;
                    vm.setData();
                    vm.isActivePage(page);
                    console.log("选择的页：" + page);
                    console.log(vm.selPage);
                });


            };

            //设置当前选中页样式
            vm.isActivePage = function (page) {
                return vm.selPage == page;
            };

            //上一页
            vm.Previous = function () {
                console.log(vm.selPage);
                console.log(vm.pages);
                if(vm.selPage>1){
                    return vm.selectPage(vm.selPage - 1);
                }else{
                    $.bigBox({
                        title: "已是第一页!",
                        content: "",
                        color: "#C79121",
                        timeout: 4000,
                        icon: "fa fa-shield fadeInLeft animated",

                    });
                }
            }
            //下一页
            vm.Next = function () {
                console.log(vm.selPage);
                console.log(vm.pages);
                if(vm.selPage < vm.pages){
                    return vm.selectPage(vm.selPage + 1);
                }else{
                    $.bigBox({
                        title: "已是最后一页!",
                        content: "",
                        color: "#C79121",
                        timeout: 4000,
                        icon: "fa fa-shield fadeInLeft animated",

                    });
                }
            };

        };
        $http({
            url:APP_CONFIG.apiRootUrl_2+'manager-group/?return=group_name'+',id'+'&page='+0+'&pageSize='+0,
            method:"get"
        }).success(function (data) {
            vm.select = data.data.data;
            console.log(987654321);
            console.log(vm.select);
        });
        $http({
            url:APP_CONFIG.apiRootUrl_2+'manager-group/'+'?page='+vm.selPage+'&pageSize='+1+'&count='+true,
            method:"get"
        }).success(function (data) {
           console.log(data);
           vm.count = data.data.count;
           vm.dataJson = data.data.data;
           console.log(vm.dataJson);
            for (var i = 0; i < vm.dataJson.length; i++) {
                if(vm.dataJson[i].status==0){
                    vm.dataJson[i].status = "禁用";
                }else if(vm.dataJson[i].status==1){
                    vm.dataJson[i].status = "启用";
                }
            };
           vm.page();
        });
        vm.store = function () {
            console.log(223344);
            console.log(vm.dataJson);
            for (var k = 0;k < vm.dataJson.length; k++){
                var sid = document.getElementsByClassName("statusID")[k];
                if(vm.dataJson[k].status=="禁用"){
                    sid.innerHTML = "启用";
                }else if(vm.dataJson[k].status=="启用"){
                    sid.innerHTML = "禁用";
                }
            }
        }
        //搜索
        vm.search = function () {
            for(var k = 0 ; k < vm.select.length ; k++){
                var name = vm.select[k].group_name;
                if(vm.hidden == name){
                    var id = vm.select[k].id;
                }
            }
            console.log(vm.status);
            if(id==undefined){
                id = "";
            }
            if(vm.status==undefined){
                vm.status = "";
            }
            $http({
                url:APP_CONFIG.apiRootUrl_2+'manager-group/'+'?status='+vm.status+'&groupID='+id+'&page='+vm.selPage+'&pageSize='+1+'&count='+true,
                method:"get"
            }).success(function (data) {
                console.log(data);
                vm.dataJson = data.data.data;
                vm.count = data.data.count;
                console.log(vm.dataJson);
                vm.page();
                if(vm.dataJson == null){
                    console.log(223344);
                    $.smallBox({
                        title: "没有这条数据",
                        content: "<i class='fa fa-clock-o'></i> <i></i>",
                        color: "#C46A69",
                        iconSmall: "fa fa-check fa-2x fadeInRight animated",
                        timeout: 4000
                    });
                }
            })
        }


        //删除弹框
        vm.remove =  function ($index) {
            vm.text = vm.dataJson[$index].id;
            // var td = $event.target.parentNode;
            // console.log(td);
            // var tr = td.parentNode;
            // console.log(tr);
            // var td1 = tr.childNodes[1];
            // console.log(td1.innerText);
            // vm.text = td1.innerHTML;
            $.SmartMessageBox({
                title: "删除权限",
                content: "确定删除此条权限？",
                buttons: '[No][Yes]'
            }, function (ButtonPressed) {
                if (ButtonPressed === "Yes") {
                    $http({
                        url:APP_CONFIG.apiRootUrl_2+'manager-group/'+vm.text,
                        method:"DELETE"
                    }).success(function (data) {
                        console.log(data);
                        $http({
                            url:APP_CONFIG.apiRootUrl_2+'manager-group/'+'?page='+vm.selPage+'&pageSize='+1+'&count='+true,
                            method:"get"
                        }).success(function (data) {
                            console.log(data);
                            vm.count = data.data.count;
                            vm.dataJson = data.data.data;
                            console.log(vm.dataJson);
                            vm.page();
                            $.smallBox({
                                title: "删除成功",
                                content: "<i class='fa fa-clock-o'></i> <i></i>",
                                color: "#659265",
                                iconSmall: "fa fa-check fa-2x fadeInRight animated",
                                timeout: 4000
                            });
                        });
                    });
                }
                if (ButtonPressed === "No") {
                    $.smallBox({
                        title: "删除失败",
                        content: "<i class='fa fa-clock-o'></i> <i></i>",
                        color: "#C46A69",
                        iconSmall: "fa fa-times fa-2x fadeInRight animated",
                        timeout: 4000
                    });
                }

            });
        };

        //是否启用
        vm.disable =  function ($index) {
            var id = vm.dataJson[$index].id;
            if(vm.dataJson[$index].status == "禁用"){
                var status = 1;
            }else if(vm.dataJson[$index].status == "启用"){
                var status = 0;
            }
            $.SmartMessageBox({
                title: "停用权限",
                content: "确定停用此条权限？",
                buttons: '[No][Yes]'
            }, function (ButtonPressed) {
                if (ButtonPressed === "Yes") {
                    $http({
                        url:APP_CONFIG.apiRootUrl_2+'manager/status/'+id+'?status='+status,
                        method:"PUT"
                    }).success(function (data) {
                        console.log(data);
                        $http({
                            url:APP_CONFIG.apiRootUrl_2+'manager/?return=account,remark,create_time,status,id'+'&page='+vm.selPage+'&pageSize='+3+'&count='+true,
                            method:"get"
                        }).success(function (data) {
                            console.log(data);
                            vm.count = data.data.count;
                            console.log(vm.count);
                            vm.dataJson = data.data.user;
                            console.log(vm.dataJson);
                            for (var i = 0; i < vm.dataJson.length; i++) {
                                if(vm.dataJson[i].status==0){
                                    vm.dataJson[i].status = "禁用";
                                    vm.status_1 = "启用";
                                }else if(vm.dataJson[i].status==1){
                                    vm.dataJson[i].status = "启用";
                                    vm.status_1 = "禁用";
                                }
                            };
                            //转化时间
                            for (var i = 0; i < vm.dataJson.length; i++){
                                function getLocalTime(nS) {
                                    return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
                                };
                                vm.dataJson[i].create_time = getLocalTime(vm.dataJson[i].create_time);
                            };
                            vm.page();
                        });
                        $.smallBox({
                            title: "停用成功",
                            content: "<i class='fa fa-clock-o'></i> <i></i>",
                            color: "#659265",
                            iconSmall: "fa fa-check fa-2x fadeInRight animated",
                            timeout: 4000
                        });
                    })

                }
                if (ButtonPressed === "No") {
                    $.smallBox({
                        title: "停用失败",
                        content: "<i class='fa fa-clock-o'></i> <i></i>",
                        color: "#C46A69",
                        iconSmall: "fa fa-times fa-2x fadeInRight animated",
                        timeout: 4000
                    });
                }

            });
        };
        //点击添加跳转页面
        vm.add = function () {
            console.log(111);
            $state.go('app.administrators.addPower');
        }
        //编辑
        vm.modify = function ($index) {
            console.log(123);
            console.log($index)
            // vm.id = vm.dataJson[$index].id;
            // console.log(vm.id);
            // $http({
            //     url:APP_CONFIG.apiRootUrl_2+'manager-group/'+vm.id+'?return=*',
            //     method:"get"
            // }).success(function (data) {
            //     console.log(data);
            //     vm.groupName = data.data.groupinfo.group_name;
            // });
            $state.go('app.administrators.modifyPower');
        };
        //点击查看跳转页面
        vm.see = function ($index) {
            $state.go('app.administrators.seePower');
        };
    }
});
