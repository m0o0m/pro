angular.module('app.copyEditor').controller('WapDiscountCtrl',
    function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$state,attachmentService){
        $scope.sitId = function (site_index_id) {
            CopyEditorService.getSiteSelect(site_index_id).then(function (response) {
                $scope.sharedJson = response.data.data;
            });
        };
        var user = JSON.parse($LocalStorage.getItem("user"));
        $scope.isSuperAdmin = user.site_index_id === '';
        if ($scope.isSuperAdmin === false) {
            //获取全部站点
            $scope.sitId();
        } else {
            $scope.sitId(user.site_index_id);
        }
        $scope.json = APP_CONFIG.option;
        var GetAllEmployee = function () {
            var postData = {
                site_index_id: $scope.site_index_id
            }
            CopyEditorService.getWapDiscount(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.list;
            })
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };
        $scope.modify = function (id,name,orderby,status,type) {
            console.log(name);
            console.log(orderby);
            console.log(status);
            console.log(type);
            console.log(id);
            $scope.id_s = id;
            $scope.name = name;
            $scope.orderby = orderby;
            $scope.type = type;
            $scope.status = status;
            var status_1 = document.getElementsByName('status');
            console.log(status_1);
            for(var i = 0;i < 2;i++) {
                if (status_1[i].value == status) {
                    status_1[i].checked = 'checked';
                }
            }
        };
        $scope.sub = function () {
            $scope.status_2 = $("input[name='status']:checked").val();
            var postData = {
                title: $scope.addName,
                orderby: $scope.addOrderby,
                status: $scope.status_2*1,
                type: $scope.addType*1
            }
            CopyEditorService.getWapDiscountSub(postData).then(function (response) {
                console.log(response);
                if(response.data.data==null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        $scope.addSub = function () {
            $scope.status_1 = $("input[name='addStatus']:checked").val();
            var postData = {
                id: $scope.id_s,
                title: $scope.name,
                orderby: $scope.orderby,
                status: $scope.status_1*1,
                type: $scope.type*1
            }
            CopyEditorService.getWapDiscountAddSub(postData).then(function (response) {
                console.log(response);
                if(response.data.data==null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };

        $scope.mContent = function (id) {
            $state.go('app.CopyEditor.WapModifyDiscount',{
                id:id
            })
        };
        $scope.dContent = function (id) {
            $state.go('app.CopyEditor.wapDiscountContent',{
                id:id
            })
        };
        $scope.update = function (id,name,url) {
            $scope.title = name;
            $scope.url = url;
            $scope.is_id = id;
            attachmentService.getList({

            }).then(function (response) {
                $scope.enclosure = response.data;
            });

        };

        $scope.select = function (item) {
            $scope.url = item.url;
        };
        $scope.modifyTitle = function (item) {
            attachmentService.modify({
                id: item.id,
                title: item.title
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (item) {
            attachmentService.del({
                id: item.id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.submit = function () {
            CopyEditorService.getWapDiscountUpdate({
                id: $scope.is_id,
                title: $scope.title,
                url: $scope.url
            }).then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }
        $scope.del = function (id) {
            var del_s = function () {
                var postData = {
                    id: id
                };
                CopyEditorService.getWapDiscountDel(postData).then(function (response) {
                    if (response.data.data===null){
                        GetAllEmployee();
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del_s);
        }
    });