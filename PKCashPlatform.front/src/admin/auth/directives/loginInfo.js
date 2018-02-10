angular.module('app.auth').directive('loginInfo', function(){

    return {
        restrict: 'A',
        templateUrl: 'views/auth/directives/login-info.tpl.html',
        link: function(scope, element){
            scope.user = {
                username: "Agency",
                picture: "img/avatars/sunny.png"
            };
        }
    }
});
