angular.module('app.administrators').controller('ApplicationManagementCtrl',
    function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,CONFIG){
        var arr = sessionStorage.getItem('user');
        var array = JSON.parse(arr);
        $scope.token = array.token;
        $scope.site_index_id = array.site_index_id;
        $scope.site_id = array.site_id;
        $scope.level = array.level;
        if($scope.site_index_id == ""){

        }else{
            var sitelength = document.getElementsByClassName('allsite').length;
            console.log(sitelength);
            for(var i=0;i<sitelength;i++){
                document.getElementsByClassName("allsite")[i].style.display="none";
            }
        }

        //下拉框请求json
        httpSvc.getJson("/select.json").then(function (data) {
            $scope.json=data[0];
        })


        httpSvc.get("/agent/first/drop").then(function (response) {
            console.log(response);
            // $scope.paginationConf.totalItems = response.meta.count;
            $scope.site = response.data;
            console.log($scope.site);
        })
        if($scope.searchterms == undefined){
            $scope.searchterms = "";
        }
        if($scope.appsite == undefined){
            $scope.appsite = "";
        }
        if($scope.applicationstatus == undefined){
            $scope.applicationstatus = "";
        }
        if($scope.applicationtype == undefined){
            $scope.applicationtype = "";
        }
        // if($scope.applicationorderby == undefined){
        //     $scope.applicationorderby = "";
        // }

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };
        var GetAllEmployee = function () {

            var searchData = {
                site_index_id:$scope.appsite,
                status:$scope.applicationstatus,
                key:$scope.applicationtype,
                // orderby:$scope.applicationtype,
                // desc:$scope.applicationorderby,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                value:$scope.searchterms
            }
            httpSvc.get("/agent/register",searchData).then(function (data) {
                console.log(data);
                if(data.data == null){
                    $scope.paginationConf.totalItems = 0;
                    $scope.list = data.data;
                }else if(data.meta.count>0){
                    $scope.paginationConf.totalItems = data.meta.count;
                    $scope.list = data.data;
                }
            })
        }

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        //搜索
        $scope.search = function () {
            console.log($scope.applicationstatus);
            var searchData = {
                site_index_id:$scope.appsite,
                status:$scope.applicationstatus,
                key:$scope.applicationtype,
                // orderby:"",
                // desc:$scope.applicationorderby,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                value:$scope.searchterms
            }
            httpSvc.get("/agent/register",searchData).then(function (data) {
                console.log(112233);
                console.log(data);
                if(data.data == null){
                    $scope.paginationConf.totalItems = 0;
                    $scope.list = data.data;
                }else if(data.meta.count>0){
                    $scope.paginationConf.totalItems = data.meta.count;
                    $scope.list = data.data;
                }
            });
        };
        //查看
        $scope.see = function ($index) {
            console.log($scope.list[$index]);
            $scope.detail = $scope.list[$index];
            // httpSvc.post("/agent/register/setting",{
            //     site_index_id:
            // }).then(function (response) {
            //
            // })
        }
        //删除
        $scope.del = function ($index) {
            // console.log($scope.searchData);
            var delData = {
                site_id:$scope.appsite,
                status:$scope.applicationstatus,
                key:$scope.applicationtype,
                // orderby:$scope.applicationtype,
                // desc:$scope.applicationorderby,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                value:$scope.searchterms
            };
            httpSvc.get("/agent/register",delData).then(function (response) {
                console.log(112211);
                console.log(response.data[$index].id);
                $scope.appID = response.data[$index].id;
                console.log($scope.appID);
            });
            $scope.delete = function () {

                httpSvc.del("/agent/register",{
                    register_id:$scope.appID
                }).then(function (response) {
                    console.log(1133311);
                    console.log(response);
                    if(response===null){
                        GetAllEmployee();
                        popupSvc.smallBox("success","删除成功！");
                    }else{
                        popupSvc.smallBox("fail",response.msg);
                    }


                });
            };
            popupSvc.smartMessageBox("确定删除？",$scope.delete);
        };
        //代理申请注册设定初始化
        $scope.addregister = function () {
            console.log($scope.site[0].site_index_id);
            console.log(112233);
            var addsite = document.getElementsByClassName('addsite')[0].value;
            var registerProxy = document.getElementsByName('registerProxy');
            var chineseNickname = document.getElementsByName('chineseNickname');
            var englishNickname = document.getElementsByName('englishNickname');
            var needCardnumber = document.getElementsByName('needCardnumber');
            var needEmail = document.getElementsByName('needEmail');
            var needQQ = document.getElementsByName('needQQ');
            var NeedPhone = document.getElementsByName('NeedPhone');
            var promoteWebsite = document.getElementsByName('promoteWebsite');
            var otherMethod = document.getElementsByName('otherMethod');
            var isMustChineseNickname = document.getElementsByName('isMustChineseNickname');
            var isMustEnglishNickname = document.getElementsByName('isMustEnglishNickname');
            var isMustEmail = document.getElementsByName('isMustEmail');
            var isMustIdentity = document.getElementsByName('isMustIdentity');
            var isMustQQ = document.getElementsByName('isMustQQ');
            var isMustPhone = document.getElementsByName('isMustPhone');
            var isMustPromoteWebsite = document.getElementsByName('isMustPromoteWebsite');
            var isMUstMethod = document.getElementsByName('isMUstMethod');
            //代理申请注册设定下拉框的值发生变化时
            $('.addsite').change(function () {
                console.log($('.addsite'));
                var sitevalue = $('.addsite')[1].value;
                console.log(sitevalue);
                if(sitevalue===undefined){
                    sitevalue = "";
                }
                httpSvc.get("/agent/register/setting",{
                    site_index_id:sitevalue
                }).then(function (response) {
                    console.log(response);
                    console.log(response.code);
                    var code = response.code;
                    // var msg = response.msg;
                    if(response.data===null||code === 10050){
                        $("input:radio[name='registerProxy']").removeAttr("checked");
                        $("input:radio[name='chineseNickname']").removeAttr("checked");
                        $("input:radio[name='englishNickname']").removeAttr("checked");
                        $("input:radio[name='needCardnumber']").removeAttr("checked");
                        $("input:radio[name='needEmail']").removeAttr("checked");
                        $("input:radio[name='needQQ']").removeAttr("checked");
                        $("input:radio[name='NeedPhone']").removeAttr("checked");
                        $("input:radio[name='promoteWebsite']").removeAttr("checked");
                        $("input:radio[name='otherMethod']").removeAttr("checked");
                        $("input:radio[name='isMustChineseNickname']").removeAttr("checked");
                        $("input:radio[name='isMustEnglishNickname']").removeAttr("checked");
                        $("input:radio[name='isMustEmail']").removeAttr("checked");
                        $("input:radio[name='isMustIdentity']").removeAttr("checked");
                        $("input:radio[name='isMustQQ']").removeAttr("checked");
                        $("input:radio[name='isMustPhone']").removeAttr("checked");
                        $("input:radio[name='isMustPromoteWebsite']").removeAttr("checked");
                        $("input:radio[name='isMUstMethod']").removeAttr("checked");
                    }else{
                        var value1 = response.data.site_index_id;
                        var value2 = response.data.register_proxy;
                        var value3 = response.data.chinese_nickname;
                        var value4 = response.data.english_nickname;
                        var value5 = response.data.need_card;
                        var value6 = response.data.need_email;
                        var value7 = response.data.need_qq;
                        var value8 = response.data.need_phone;
                        var value9 = response.data.promote_website;
                        var value10 = response.data.other_method;
                        var value11 = response.data.is_must_chinese_nickname;
                        var value12 = response.data.is_must_english_nickname;
                        var value13 = response.data.is_must_email;
                        var value14 = response.data.is_must_identity;
                        var value15 = response.data.is_must_qq;
                        var value16 = response.data.is_must_phone;
                        var value17 = response.data.is_must_promote_website;
                        var value18 = response.data.is_must_method;
                        for(var i = 0;i < 2;i++){
                            if(registerProxy[i].value == value2){
                                registerProxy[i].checked =  'checked';
                            }
                            if(chineseNickname[i].value == value3){
                                chineseNickname[i].checked =  'checked';
                            }
                            if(englishNickname[i].value == value4){
                                englishNickname[i].checked =  'checked';
                            }
                            if(needCardnumber[i].value == value5){
                                needCardnumber[i].checked =  'checked';
                            }
                            if(needEmail[i].value == value6){
                                needEmail[i].checked =  'checked';
                            }
                            if(needQQ[i].value == value7){
                                needQQ[i].checked =  'checked';
                            }
                            if(NeedPhone[i].value == value8){
                                NeedPhone[i].checked =  'checked';
                            }
                            if(promoteWebsite[i].value == value9){
                                promoteWebsite[i].checked =  'checked';
                            }
                            if(otherMethod[i].value == value10){
                                otherMethod[i].checked =  'checked';
                            }
                            if(isMustChineseNickname[i].value == value11){
                                isMustChineseNickname[i].checked =  'checked';
                            }
                            if(isMustEnglishNickname[i].value == value12){
                                isMustEnglishNickname[i].checked =  'checked';
                            }
                            if(isMustEmail[i].value == value13){
                                isMustEmail[i].checked =  'checked';
                            }
                            if(isMustIdentity[i].value == value14){
                                isMustIdentity[i].checked =  'checked';
                            }
                            if(isMustQQ[i].value == value15){
                                isMustQQ[i].checked =  'checked';
                            }
                            if(isMustPhone[i].value == value16){
                                isMustPhone[i].checked =  'checked';
                            }
                            if(isMustPromoteWebsite[i].value == value17){
                                isMustPromoteWebsite[i].checked =  'checked';
                            }
                            if(isMUstMethod[i].value == value18){
                                isMUstMethod[i].checked =  'checked';
                            }
                        }
                    }

                });
            });
        };
        //代理申请注册设定提交
        $scope.appsubmit = function () {
            $scope.addsite = document.getElementsByClassName('addsite')[1].value;
            $scope.registerProxy = $("input[name='registerProxy']:checked").val();
            $scope.chineseNickname = $("input[name='chineseNickname']:checked").val();
            $scope.englishNickname = $("input[name='englishNickname']:checked").val();
            $scope.needCardnumber = $("input[name='needCardnumber']:checked").val();
            $scope.needEmail = $("input[name='needEmail']:checked").val();
            $scope.needQQ = $("input[name='needQQ']:checked").val();
            $scope.NeedPhone = $("input[name='NeedPhone']:checked").val();
            $scope.promoteWebsite = $("input[name='promoteWebsite']:checked").val();
            $scope.otherMethod = $("input[name='otherMethod']:checked").val();
            $scope.isMustChineseNickname = $("input[name='isMustChineseNickname']:checked").val();
            $scope.isMustEnglishNickname = $("input[name='isMustEnglishNickname']:checked").val();
            $scope.isMustEmail = $("input[name='isMustEmail']:checked").val();
            $scope.isMustIdentity = $("input[name='isMustIdentity']:checked").val();
            $scope.isMustQQ = $("input[name='isMustQQ']:checked").val();
            $scope.isMustPhone = $("input[name='isMustPhone']:checked").val();
            $scope.isMustPromoteWebsite = $("input[name='isMustPromoteWebsite']:checked").val();
            $scope.isMUstMethod = $("input[name='isMUstMethod']:checked").val();
            console.log($scope.addsite);
            if($scope.addsite==''||$scope.registerProxy==undefined||$scope.chineseNickname==undefined||$scope.englishNickname==undefined||$scope.needCardnumber==undefined||$scope.needEmail==undefined||$scope.needQQ==undefined||$scope.NeedPhone==undefined||$scope.promoteWebsite==undefined||$scope.otherMethod==undefined||$scope.isMustChineseNickname==undefined||$scope.isMustEnglishNickname==undefined||$scope.isMustEmail==undefined||$scope.isMustIdentity==undefined||$scope.isMustQQ==undefined||$scope.isMustPhone==undefined||$scope.isMustPromoteWebsite==undefined||$scope.isMUstMethod==undefined){
                popupSvc.smallBox("fail","请输入完整！")
            }else{
                console.log($scope.registerProxy);
                console.log($scope.addsite);
                console.log($scope.registerProxy);
                console.log($scope.chineseNickname);
                console.log($scope.englishNickname);
                console.log($scope.needCardnumber);
                console.log($scope.needEmail);
                console.log($scope.needQQ);
                console.log($scope.NeedPhone);
                console.log($scope.promoteWebsite);
                console.log($scope.otherMethod);
                console.log($scope.isMustChineseNickname);
                console.log($scope.isMustEnglishNickname);
                console.log($scope.isMustEmail);
                console.log($scope.isMustIdentity);
                console.log($scope.isMustQQ);
                console.log($scope.isMustPhone);
                console.log($scope.isMustPromoteWebsite);
                console.log($scope.isMUstMethod);
                var addData = {
                    site_index_id:$scope.addsite,
                    register_proxy:$scope.registerProxy*1,
                    chinese_nickname:$scope.chineseNickname*1,
                    english_nickname:$scope.englishNickname*1,
                    need_card:$scope.needCardnumber*1,
                    need_email:$scope.needEmail*1,
                    need_qq:$scope.needQQ*1,
                    need_phone:$scope.NeedPhone*1,
                    promote_website:$scope.promoteWebsite*1,
                    other_method:$scope.otherMethod*1,
                    is_must_chinese_nickname:$scope.isMustChineseNickname*1,
                    is_must_english_nickname:$scope.isMustEnglishNickname*1,
                    is_must_email:$scope.isMustEmail*1,
                    is_must_identity:$scope.isMustIdentity*1,
                    is_must_qq:$scope.isMustQQ*1,
                    is_must_phone:$scope.isMustPhone*1,
                    is_must_promote_website:$scope.isMustPromoteWebsite*1,
                    is_must_method:$scope.isMUstMethod*1
                };
                httpSvc.post("/agent/register/setting",addData).then(function (response) {
                    console.log(response);
                    if(response == null){
                        if(response===null){
                            popupSvc.smallBox("success","操作成功");
                            GetAllEmployee();
                        }else {
                            popupSvc.smallBox("fail",response.msg);
                        }
                    }
                })
            }
        }
        //新增账号初始化
        $scope.addmodal =function ($index,$event) {
            $scope.modalsite="";
            $scope.modalsecond="";
            $scope.modalparent="";
            $scope.modalaccounts="";
            $scope.modalpassword="";
            $scope.modalpass="";
            angular.element("#select2-chosen-4").text("请选择");
            angular.element("#select2-chosen-6").text("请选择");
            var idkey = angular.element($event.target).parent().prev().text();
            console.log($index);
            console.log(idkey)
            //获取账号信息
            httpSvc.get("/agent/register/one",{
                id: idkey*1
            }).then(function (response) {
                console.log(response);
                $scope.modalaccounts = response.data.account;
                $scope.site_id = response.data.site_id;
                $scope.modalID = response.data.id;
                $scope.site_index_id = response.data.site_index_id;
                $scope.username = response.data.zh_name;
                console.log($scope.site_index_id);
                httpSvc.get("/agent/second/drop",{
                    site_index_id: $scope.site_index_id
                }).then(function (response) {
                    console.log(121212);
                    console.log(response);
                    $scope.second_id = response.data;
                    $scope.iddd = response.data[0].id;
                    console.log($scope.iddd);
                    httpSvc.get("/agent/third/drop",{
                        site_id:$scope.site_id,
                        first_id:$scope.iddd
                    }).then(function (response) {
                        console.log(121212);
                        console.log(response);
                        $scope.third_id = response.data;
                    })
                })

            })
        }
        //股东下拉改变是获取总代
        $scope.secondadd = function () {
            document.getElementById("select2-chosen-6").innerText="请选择";
            console.log($scope.site_id);
            console.log($scope.modalsecond);
            httpSvc.get("/agent/third/drop",{
                site_id:$scope.site_id,
                first_id:$scope.modalsecond
            }).then(function (response) {
                console.log(121212);
                console.log(response);
                $scope.third_id = response.data;
            });
        };
        //新增账号提交
        $scope.modalsubmit = function () {
            console.log($scope.userName);
            var modaladd = {
                register_id:$scope.modalID,
                parent_id:$scope.modalparent*1,
                account:$scope.modalaccounts,
                password:$scope.modalpass,
                confirm_password:$scope.modalpassword,
                site_index_id:$scope.site_index_id,
                username:$scope.username
            };
            // if($scope.modalsite==undefined||$scope.modalparent==undefined||$scope.modalaccounts==undefined||$scope.modalpass==undefined||$scope.modalpassword==undefined||$scope.modalsecond==undefined||$scope.modalsecond==''||$scope.modalparent==''){

            // if($scope.modalsite==''||$scope.modalparent==''||$scope.modalaccounts==''||$scope.modalpass==''||$scope.modalpassword==''||$scope.modalsecond==''){
            //     popupSvc.smallBox("fail","请输入完整！")
            // }else{
            //
            // }

            httpSvc.post("/agent/register",modaladd).then(function (response) {

                if(response===null){
                    popupSvc.smallBox("success","操作成功");
                    GetAllEmployee();

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
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
