angular.module('app.auth').controller('LoginCtrl',
    function ($scope, $timeout, $LocalStorage, AuthService,APP_CONFIG, $state, popupSvc) {
        var winHeight = $("html").height();
        $("#login-box").height(winHeight);
        $LocalStorage.setItem("version", APP_CONFIG.version);
        $scope.getCode = function () {
            AuthService.getCaptcha().then(function (data) {
                console.log(data);
                $scope.header_code = data.code;
                $("#codeImg").attr("src", data.blob);
            });
        };
        $scope.getCode();
        $scope.login = function () {
            AuthService.login($scope.account, $scope.password, 
                $scope.header_code, $scope.code, $scope.verify_code).then(function (response) {
                if (response.code) {
                    if (response.code == 20022) {
                        popupSvc.smallBox("fail", response.msg);
                        $scope.klyz = true;
                        $scope.getCode();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                        $scope.getCode();
                    }
                    console.info(response);
                } else {
                    var obj = JSON.stringify(response.data);
                    $LocalStorage.setItem(APP_CONFIG.tokenKey, obj);
                    $state.go('app.dashboard');
                }
            });
        };
        if (APP_CONFIG.debugState) {
            $scope.header_code = "111111111111";
            $("#codeImg").attr("src", "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==");
        } else {
            $scope.getCode();
        }
    });