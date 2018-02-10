angular.module('app.administrators').controller('bankCtrl',
    function(BusinessService,httpSvc,popupSvc,resourceSvc,DTColumnBuilder,$http,$scope,$stateParams,CONFIG,APP_CONFIG){
        //接受参数
        $scope.bankId = $stateParams.ids;
        console.log($scope.bankId);

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                id:$scope.bankId
            };

            httpSvc.get("/member/bank", postData).then(function (response) {
                console.log(response);
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            }, function (error) {

            });
        };

        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        // 删除
        $scope.disable=function (bankid) {
            var del = function () {
                httpSvc.del("/member/bank",{
                    id:bankid
                }).then(function (response) {
                    popupSvc.smallBox("success","删除成功");
                    GetAllEmployee();
                });
            }
            popupSvc.smartMessageBox("删除会员银行?",del);
        };
        $scope.modifys = {
            card_name:"",
            card:"",
            card_address:"",
            bank_id:"",
            member_id:$scope.bankId,
            id:""

        }

        //获取单个银行信息
        $scope.Discount = function (id) {
            httpSvc.get("/member/bank/info",{
                card:id
            }).then(function (response) {
               console.log(response);
               $scope.modifys.card_name = response.data.card_name;
               $scope.modifys.card = response.data.card;
               $scope.modifys.card_address = response.data.card_address;
               $scope.modifys.bank_id = response.data.bank_id;
               $scope.modifys.id = response.data.id;
            });
        };

        //修改后提交
        $scope.modifyssubmit = function () {
            httpSvc.put("/member/bank",$scope.modifys)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","修改成功");
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }

                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                });
        };

    });