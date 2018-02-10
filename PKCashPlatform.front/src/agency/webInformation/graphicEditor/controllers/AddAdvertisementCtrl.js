angular.module('app.GraphicEditor').controller('AddAdvertisementCtrl',
    function ($scope, popupSvc, siteService, attachmentService, advertisementService, $rootScope, APP_CONFIG,$state,$stateParams) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });

        function formatDateTime(timeStamp) {
            var date = new Date();
            date.setTime(timeStamp * 1000);
            var y = date.getFullYear();
            var m = date.getMonth() + 1;
            m = m < 10 ? ('0' + m) : m;
            var d = date.getDate();
            d = d < 10 ? ('0' + d) : d;
            return y + '-' + m + '-' + d;
        }

        advertisementService.getDetail({
            id: $stateParams.id
        }).then(function (response) {
            $scope.data = response.data;
            $scope.data.startTime=formatDateTime($scope.data.startTime);
            $scope.data.endTime=formatDateTime($scope.data.endTime);
        });

        $scope.submitAd=function(){
            advertisementService.modify($scope.data).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };


        $scope.upload = function () {
            attachmentService.getList({
                id: $stateParams.id
            }).then(function (response) {
                $scope.enclosure = response.data;
            });
        };

        $scope.select = function (item) {
            $scope.title = item.title;
            $scope.url = item.url;
            $scope.enclosure_id = item.id;
        };
        $scope.modifyTitle = function (item) {
            attachmentService.modify({
                id: item.id,
                title: item.title
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
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
                    logoService.getEnclosure({
                        id: $scope.id
                    }).then(function (response) {
                        $scope.enclosure = response.data;
                    });
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.submit = function () {
            $scope.data.imgUrl=$scope.url
        }



    });