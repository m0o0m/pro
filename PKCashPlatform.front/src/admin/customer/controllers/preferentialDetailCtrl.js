angular.module('app.customer').controller('preferentialDetailCtrl',
    function ($scope, popupSvc, customerPreferentialQueryService, $rootScope, APP_CONFIG) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };

        var GetAllEmployee = function () {
            var postData = {
                date: $scope.date
            };

            customerPreferentialQueryService.getList(postData).then(function (response) {
                $scope.list = response.data;
            });

        };
        GetAllEmployee();

        $scope.search=function(){
            GetAllEmployee()
        };


    });