angular.module('app.administrators').controller('MembershipDetailsCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,$stateParams){

        // 搜索
        $scope.search=function () {
            GetAllEmployee()
        };

        $scope.IsLock='';
        $scope.lock=function ($event,id) {
            var is_lock=2;
            var isCheck=$($event.target).is(':checked');
            if(isCheck){
                is_lock = 1;
            }
            httpSvc.put("/member/level/lock",{
                lock: is_lock,
                member_id: id
            }).then(function (response) {
                GetAllEmployee();
            });
        };

        //修改分层
        $scope.modifyHierarchy=function (Id) {
            $scope.Id=Id;

            httpSvc.get("/member/level/drop",{
                site_index_id: $stateParams.site_index_id
            }).then(function (response) {
                $scope.hierarchylist = response.data;
                $scope.hierarchylist.shift();
            });

        };
        //提交分层
        $scope.modifyHierarchySubmit=function () {
            var id=$("input[name='hierarch']:checked").val();
            httpSvc.put("/member/level/move",{
                site_index_id: $stateParams.site_index_id,
                move_out: $scope.Id,
                move_in: id
            }).then(function (response) {
                GetAllEmployee();
                $(".modal-backdrop").hide();
                $("#myModal2").hide();

            });
        };

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                level_id: $stateParams.level_id,
                site_index_id: $stateParams.site_index_id,
                account: $scope.account
            };

            httpSvc.get("/member/level/memberlist",postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


    });
