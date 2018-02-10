/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.auth", [])
  .service("AuthService", AuthService);

AuthService.$inject = ['APP_CONFIG','$http'];
function AuthService(APP_CONFIG,$http) {

  return {
    login: login,
    getCaptcha :getCaptcha
  };

  //登陆用户名 密码 google code，验证码
  function login(username, password, header_code, code, captcha) {
    return $http.post(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.LOGIN, {
      account: username,
      password: password,
      code: code,
      verify_code: captcha
    }, {
      headers: {'code': header_code}
    })
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      // popupSvc.smallBox("fail", response.msg);
      return response.data;
    }

    function getDataFailed(error) {
      console.log('XHR Failed for getAvengers.' + error);
    }
  }

  function getCaptcha() {
    return $http({
      url: APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.CAPTCHA,
      method: "GET",
      responseType: 'arraybuffer'
    })
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      var header_code = response.headers()["code"];
      var blob = new Blob([response.data]);
      var objectUrl = URL.createObjectURL(blob);
      return {code: header_code, blob: objectUrl};
    }

    function getDataFailed(error) {
      console.log('XHR Failed for getAvengers.' + error);
    }
  }
}