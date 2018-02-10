angular.module('app.copyEditor').controller('LineDetectionCtrl', function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$state,attachmentService){
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
    var GetAllEmployee = function () {
        var postData = {
            site_index_id: $scope.site_index_id
        }
        CopyEditorService.getLineDetection(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        })
    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    
    $scope.submit = function () {
        $scope.HTTP = $("input[name='http']:checked").val();
        var postData = {
            http: $scope.HTTP,
            url: $scope.url
        }
        CopyEditorService.getLineAddSub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }
    $scope.modify = function(id,src,http){
        $scope.id_s = id;
        $scope.mUrl = src;
        var status_1 = document.getElementsByName('mHttp');
        console.log(status_1);
        for(var i = 0;i < 2;i++) {
            if (status_1[i].value == http) {
                status_1[i].checked = 'checked';
            }
        }
    }
    $scope.sub = function () {
        $scope.status_2 = $("input[name='mHttp']:checked").val();
        var postData = {
            url: $scope.mUrl,
            id: $scope.id_s*1,
            http: $scope.status_2*1
        }
        CopyEditorService.getLineModifySub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    $scope.del = function (id) {
        var del_s = function () {
            var postData = {
                id: id
            };
            CopyEditorService.getLineDetectionDel(postData).then(function (response) {
                if (response.data.data===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del_s);
    };
});