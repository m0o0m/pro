/**
 * Created by apple on 17/10/14.
 */
angular.module('app.site').controller('urlCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
    $scope.option_video_type = APP_CONFIG.option.option_video_type;
    var GetAllEmployee = function () {
        siteService.linkDownload().then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.list;
        });
    };

    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
     $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


    //添加
    $scope.add = function () {
        var postData = {
            type_id:$scope.add.typeed,
            ios:$scope.add.ios,
            andriod:$scope.add.andriod
        };
        siteService.addVideo(postData).then(function (data) {
            if(data.data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
            }else {
                popupSvc.smallBox("fail",data.msg);
            };
        });
    };



    // 更改状态
    $scope.disables=function (status,id,item) {
        var status = 2;
        if (item.status === 2 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        }
       var postData={
           status:status,
           id:id
       }
        var sure = function () {
            siteService.downloadStatus(postData).then(function (response) {
                if(response.data===null){
                    item.status = status;
                    popupSvc.smallBox("success",$rootScope.getWord("success"));

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });

        }
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
    };
        //获取单个详情
        $scope.getID = function (id) {
            siteService.downloadlinksModify(id).then(function (res) {
                $scope.data = res.data;
                console.log($scope.data);
            });

        };
    //修改完成点击提交
    $scope.submit = function () {
        var postData = {
            type_id:$scope.data.type_id,
            ios:$scope.data.ios,
            andriod:$scope.data.andriod
        };
        siteService.downloadModify(postData).then(function (response) {
            if(response.data===null){
                popupSvc.smallBox("success","修改状态成功");
                GetAllEmployee();

            }else {
                popupSvc.smallBox("fail",response.msg);
            }
        });

    };




});