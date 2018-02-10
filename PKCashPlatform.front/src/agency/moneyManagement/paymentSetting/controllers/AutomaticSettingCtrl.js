angular.module('app.PaymentSetting').controller('AutomaticSettingCtrl', function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,$state){
    var GetAllEmployee = function () {
        var postData = {
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        }

        httpSvc.get("/report.json", postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data[0].logo.length;
            $scope.bank = response.data[0].logo;
            console.log($scope.bank);
        })
    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 20
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.set = function () {
        $state.go('app.PaymentSetting.NewSettings')
    }

    // var page;
    // function page() {
    //     //分页总数
    //     $scope.pageSize = 2;
    //     //分页数
    //     $scope.pages = Math.ceil($scope.json.length / $scope.pageSize);
    //     console.log($scope.pages);
    //     $scope.newPages = $scope.pages > 5 ? 5 : $scope.pages;
    //     console.log($scope.newPages);
    //     $scope.pageList = [];
    //     $scope.selPage = 1;
    //     //设置表格数据源(分页)
    //     $scope.setData = function () {
    //         $scope.items = $scope.json.slice(($scope.pageSize * ($scope.selPage - 1)), ($scope.selPage * $scope.pageSize));//通过当前页数筛选出表格当前显示数据
    //         console.log($scope.items);
    //     }
    //     $scope.items = $scope.json.slice(0, $scope.pageSize);
    //     //分页要repeat的数组
    //     for (var i = 0; i < $scope.newPages; i++) {
    //         $scope.pageList.push(i + 1);
    //     }
    //     //打印当前选中页索引
    //     $scope.selectPage = function (page) {
    //         //不能小于1大于最大
    //         if (page < 1 || page > $scope.pages) return;
    //         //最多显示分页数5
    //         if (page > 2) {
    //             //因为只显示5个页数，大于2页开始分页转换
    //             var newpageList = [];
    //             for (var i = (page - 3) ; i < ((page + 2) > $scope.pages ? $scope.pages : (page + 2)) ; i++) {
    //                 newpageList.push(i + 1);
    //             }
    //             $scope.pageList = newpageList;
    //         }
    //         $scope.selPage = page;
    //         $scope.setData();
    //         $scope.isActivePage(page);
    //         console.log("选择的页：" + page);
    //     };
    //     //设置当前选中页样式
    //     $scope.isActivePage = function (page) {
    //         return $scope.selPage == page;
    //     };
    //     //上一页
    //     $scope.Previous = function () {
    //         $scope.selectPage($scope.selPage - 1);
    //     }
    //     //下一页
    //     $scope.Next = function () {
    //         $scope.selectPage($scope.selPage + 1);
    //     };
    // }





});

