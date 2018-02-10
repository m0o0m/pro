angular.module('app.copyEditor').controller('discountContentCtrl',
    function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$state,attachmentService,$stateParams){
        $scope.id = $stateParams.id;
        $scope.json = APP_CONFIG.option.option_discount[0];
        var GetAllEmployee = function () {
            var postData = {
                id: $scope.id
            }
            CopyEditorService.getDiscount_C(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.list;
            })
            CopyEditorService.getDiscountWidth(postData).then(function (response) {
                $scope.width = response.data.width;
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
        $scope.mWidth = function () {
            console.log($scope.width);
            var postData = {
                width:$scope.width
            };
            CopyEditorService.getDiscount_C_W_Sub(postData).then(function (response) {
                console.log(response);
                if(response.data.data==null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
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
            CopyEditorService.getDiscount_C_M_S(postData).then(function (response) {
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
            CopyEditorService.getDiscount_C_AddSub(postData).then(function (response) {
                console.log(response);
                if(response.data.data==null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        }

        $scope.mContent = function (id) {
            $state.go('app.CopyEditor.contentModifyDiscount',{
                id:id
            })
        }
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
            CopyEditorService.getDiscount_C_Update({
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
                CopyEditorService.getDiscount_C_Del(postData).then(function (response) {
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
        $scope.keep = function () {
            console.log($scope.site_index_id);
            if($scope.site_index_id==undefined){
                $scope.site_index_id = "";
            }
            var keep = function () {
                var postData = {
                    site_index_id: $scope.site_index_id
                };
                CopyEditorService.getDiscount_C_Keep(postData).then(function (response) {
                    if (response.data.data===null){
                        GetAllEmployee();
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),keep);
        }
    });