angular.module('app.site')
    .controller('SiteCtrl', function($scope,DTOptionsBuilder,DTColumnBuilder,$state,$compile,$http){
        var vm = this;
        activate();
        function activate() {
            var apiURL = "http://192.168.8.173:10080/";
            vm.page = function () {
                vm.pageSize = 5;
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
                    console.log(page);
                    $http({
                        url:apiURL+'site/?return=site_name,id,site_agency,creat_time,status,site_service_info,order_num,trade_total_money,distribution_agency_num'+'&page='+page+'&pageSize='+5,
                        method:"get"
                    }).success(function (data) {
                        console.log(data);
                        vm.count = data.count;
                        vm.dataJson = data.data;
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
                url:apiURL+'site/?return=site_name,id,site_agency,creat_time,status,site_service_info,order_num,trade_total_money,distribution_agency_num'+'&page='+vm.selPage+'&pageSize='+5,
                method:"get"
            }).success(function (data) {
                console.log(data);
                vm.count = data.count;
                vm.dataJson = data.data;
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
                vm.page();
            });

            $http({
                url:apiURL+'agency/?return=agency_name',
                method:"get"
            }).success(function (data) {
                console.log(data);
                vm.count = data.count;
                vm.select = data.data;
            })
            vm.search = function () {
                console.log(vm.siteName);
                console.log(vm.status);
                if(vm.siteName==undefined){
                    vm.siteName = "";
                }
                if(vm.status==undefined){
                    vm.status = "";
                }
                $http({
                    url:apiURL+'site/?return=site_name,id,site_agency,creat_time,status,site_service_info,order_num,trade_total_money,distribution_agency_num'+'&siteName='+vm.siteName+'&agencyName='+vm.status,
                    method:'get'
                }).success(function (data) {
                    console.log(data);
                    vm.dataJson = data.data;
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
                    vm.page();
                })
            }



            // 点击search
            $('.search').click(function () {
                $('.inp_1').val();
                $('.inp_2').val();
                console.log($('.inp_1').val());
                console.log($('.inp_2').val());
            });
            //点击添加站点
            vm.addSite = function () {
                console.log(1);
                $.SmartMessageBox({
                    title: "提示",
                    content: "请选择站点设置？",
                    buttons: '[HTTP][HTTPS]'
                }, function (ButtonPressed) {
                    if (ButtonPressed === "HTTP") {
                        $state.go('app.site.addSite1');
                    }
                    if (ButtonPressed === "HTTPS") {
                        $state.go('app.site.addSite');
                    }

                });
            };

            //点击修改站点
            vm.sub = function ($index) {
                console.log($index);
                var index = $index;
                $state.go('app.site.modifySite',{id:index});
            };
            vm.see = function ($index) {
                $state.go('app.site.seeSite');
            }
        }
    });

