angular.module('app.auth').controller('LoginCtrl',
    function ($scope, $timeout, $LocalStorage, AuthService,APP_CONFIG, $state, popupSvc) {
        var winHeight = $("html").height();
        $("#login-box").height(winHeight);
        $LocalStorage.setItem("version", APP_CONFIG.version);
        $scope.getCode = function () {
          AuthService.getCaptcha().then(function (data) {
            console.log(1111)
            console.log(data);
            $scope.header_code = data.code;
            $("#codeImg").attr("src", data.blob);
          });
          /*
            $http({
                url: APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.CAPTCHA,
                method: "GET",
                responseType: 'arraybuffer'
            }).success(function (data, status, headers) {
                console.log(1111)
                console.log(data);
                $scope.header_code = headers()["code"];
                var blob = new Blob([data]);
                var objectUrl = URL.createObjectURL(blob);
                $("#codeImg").attr("src", objectUrl);
            }).error(function (error) {
            });*/
        };
        $scope.getCode();
        $scope.login = function () {
          AuthService.login().then(function (response) {
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
              var obj = JSON.stringify(response);
              $LocalStorage.setItem("user", obj);
              $state.go('app.administrators.accounts');  //代理登录跳转
            }
          });
          /*
            $http.post(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.LOGIN, {
                account: $scope.account,
                password: $scope.password,
                code: $scope.code,
                verify_code: $scope.verify_code
            }, {
                headers: {'code': $scope.header_code}
            }).success(function (response) {
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
                    var obj = JSON.stringify(response);
                    $LocalStorage.setItem("user", obj);
                    $state.go('app.administrators.accounts');  //代理登录跳转
                }
            }).error(function (error) {
                popupSvc.smallBox("fail", "unknown error!");
            })*/
        }
        if (APP_CONFIG.debugState) {
            $scope.header_code = "111111111111";
            $("#codeImg").attr("src", "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==");
        } else {
            $scope.getCode();
        }
    });