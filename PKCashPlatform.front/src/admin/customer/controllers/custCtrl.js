angular.module('app.customer').controller('custCtrl',function ($scope,$http,DTOptionsBuilder,DTColumnBuilder,$compile) {
  
    var page;
    $scope.pageSize = 1;
    function page() {

        //分页数
        $scope.pages = Math.ceil($scope.count / $scope.pageSize);
        console.log($scope.pages);
        $scope.newPages = $scope.pages > 5 ? 5 : $scope.pages;
        console.log($scope.newPages);
        $scope.pageList = [];
        $scope.selPage = 1;
        //设置表格数据源(分页)
        $scope.setData = function () {
            $scope.items = $scope.json.slice(($scope.pageSize * ($scope.selPage - 1)), ($scope.selPage * $scope.pageSize));//通过当前页数筛选出表格当前显示数据
            console.log($scope.items);

        }
       $scope.items = $scope.json.slice(0, $scope.pageSize);
        //分页要repeat的数组
        for (var i = 0; i < $scope.newPages; i++) {
            $scope.pageList.push(i + 1);

        }
        //打印当前选中页索引
        $scope.selectPage = function (page) {
            $http({
                url:'http://192.168.8.175:10086'+'/client/'+'?page='+page+'&row='+$scope.pageSize+'&return=client_nickname,client_source,order_num',
                method:'get'
            }).success(function (data) {
                console.log(data);
                $scope.json=data.data.CliList
                console.log($scope.json);
                $scope.items = $scope.json;
                //不能小于1大于最大
                if (page < 1 || page > $scope.pages) return;
                //最多显示分页数5
                if (page > 2) {
                    //因为只显示5个页数，大于2页开始分页转换
                    var newpageList = [];
                    for (var i = (page - 3) ; i < ((page + 2) > $scope.pages ? $scope.pages : (page + 2)) ; i++) {
                        newpageList.push(i + 1);
                    }
                    $scope.pageList = newpageList;
                    return  $scope.selectPage($scope.selPage );

                }
                $scope.selPage = page;
                $scope.setData();
                $scope.isActivePage(page);
                console.log("选择的页：" + page);
            })

        };
        //设置当前选中页样式
        $scope.isActivePage = function (page) {
            return $scope.selPage == page;
        };
        //上一页

        $scope.Previous = function () {
            if ($scope.selPage >1) {
                return $scope.selectPage($scope.selPage - 1);
            } else {
                $.bigBox({
                    title: "已是第一页!",
                    content: "",
                    color: "#C79121",
                    //timeout: 8000,
                    icon: "fa fa-shield fadeInLeft animated",

                });

            }
        }
        //下一页

        $scope.Next = function () {
            if($scope.selPage < $scope.pages){
                return $scope.selectPage($scope.selPage + 1);
            }else{
                $.bigBox({
                    title: "已是最后一页!",
                    content: "",
                    color: "#C79121",
                    //timeout: 8000,
                    icon: "fa fa-shield fadeInLeft animated",

                });
            }

        };
    };

    //请求数据
    $http({
        url:'http://192.168.8.175:10086'+'/client/'+'?page='+1+'&row='+1+'&return=client_nickname,client_source,order_num',
        method:'GET'
    }).success(function (data) {
        $scope.count =data.data.count;
        console.log(data);
        $scope.json = data.data.CliList;
        page();
    });
});
