angular.module('app.caseEditor').controller('RevokeCaseCtrl',
    function($scope,APP_CONFIG,CaseEditorService){
        var GetAllEmployee = function () {
            CaseEditorService.getRevokeCase().then(function (response) {
                console.log(response.data.list);
                $scope.list = response.data.list;
                $scope.paginationConf.totalItems = response.data.meta.count;
            })
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    });