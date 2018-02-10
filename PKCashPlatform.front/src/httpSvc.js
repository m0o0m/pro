'use strict';

angular.module('app').factory('httpSvc', ["$q", "$http", 'APP_CONFIG', "$LocalStorage", function ($q, $http, APP_CONFIG, $LocalStorage) {
  var user = $LocalStorage.getItem(APP_CONFIG.tokenKey);
  var userObj = {};
  if(user && user != "undefined"){
     userObj = $.parseJSON(user);
  }
  var httpObj = {
    initParams: function (params) {
      var deferred = $q.defer();
      var defaultParams = angular.extend({}, params);
      deferred.resolve(angular.copy(defaultParams));
      return deferred.promise;
    },
    getHttpPromise: function (method, url, params, timeout) {
      var deferred = $q.defer();
      var token = "";
      if(userObj && userObj.token){
        token = userObj.token;
      }
      var config = angular.extend({}, {
        "method": method,
        "url": APP_CONFIG.apiUrls.HOST + url,
        "headers": {
          "Accept": "application/json, text/plain, */*",
          "Content-Type": "application/json;charset=utf-8;",
          "Lanugage": "zh",
          "Authorization": "Bearer " + token
        },
        "timeout": timeout || 100000,
        "responseType": "json"
      });

      if (angular.uppercase(method) === "GET") {
        config = angular.extend(config, {
          "params": params
        });
      } else if (angular.uppercase(method) === "POST") {
        config = angular.extend(config, {
          "data": params
        });
      } else if (angular.uppercase(method) === "PUT") {
        config = angular.extend(config, {
          "data": params
        });
      } else if (angular.uppercase(method) === "DELETE") {
        config = angular.extend(config, {
          "data": params
        });
      }

      return $http(config);
        /*.success(function (data) {
        deferred.resolve(data);
        console.info("请求成功:", data);
      }).error(function (data) {
        deferred.reject(data);
        console.log(config);
        console.info("请求失败:", data);
      });
      return deferred.promise;*/
    }
  };
  var factory = {
    post: function (url, params, timeout) {
      var deferred = $q.defer();
      // params = httpObj.initParams(params);
      httpObj.initParams(params).then(function (initParams) {
        httpObj.getHttpPromise("POST", url, initParams, timeout).then(function (data) {
          deferred.resolve(data);
          console.log(url, data);
        }, function (data) {
          deferred.reject(data);
        });
      });
      return deferred.promise;
    },
    get: function (url, params, timeout) {
      var deferred = $q.defer();
      httpObj.initParams(params).then(function (initParams) {
        httpObj.getHttpPromise("GET", url, initParams, timeout).then(function (data) {
          deferred.resolve(data);
          console.log(url, data);
        }, function (data) {
          deferred.reject(data);
        });
      });
      return deferred.promise;
    },
    put: function (url, params, timeout) {
      var deferred = $q.defer();
      httpObj.initParams(params).then(function (initParams) {
        httpObj.getHttpPromise("PUT", url, initParams, timeout).then(function (data) {
          deferred.resolve(data);
          console.log(url, data);
        }, function (data) {
          deferred.reject(data);
        });
      });
      return deferred.promise;
    },
    del: function (url, params, timeout) {
      var deferred = $q.defer();
      httpObj.initParams(params).then(function (initParams) {
        httpObj.getHttpPromise("DELETE", url, initParams, timeout).then(function (data) {
          deferred.resolve(data);
          console.log(url, data);
        }, function (data) {
          deferred.reject(data);
        });
      });
      return deferred.promise;
    },
    file: function (url, params, timeout) {
      var deferred = $q.defer();
      $http({
        method: 'POST',
        url: APP_CONFIG.apiUrls.HOST + url,
        data: params,
        headers: {
          "timeout": timeout || 100000,
          'Content-Type': undefined
        },
        transformRequest: angular.identity
      }).success(
        function (data, status, header, config) {
          deferred.resolve(data);
        }).error(
        function (data, status, header, config) {
          deferred.reject(data);
        });
      return deferred.promise;
    },
    form: function (url, params, timeout) {
      var deferred = $q.defer();
      $http({
        method: 'POST',
        url: APP_CONFIG.apiUrls.HOST + url,
        headers: {
          "Accept": "*/*",
          "Content-Type": "multipart/form-data; boundary=----WebKitFormBoundaryTMM4W2B1ALSDGYN6",
          "lanugage": "zh",
          "timeout": timeout || 100000
        },
        data: params
      }).success(
        function (data, status, header, config) {
          deferred.resolve(data);
        }).error(
        function (data, status, header, config) {
          deferred.reject(data);
        });
      return deferred.promise;
    },
      forms: function (url, params, timeout) {
          var deferred = $q.defer();
          $http({
              method: 'PUT',
              url: APP_CONFIG.apiUrls.HOST + url,
              headers: {
                  "Accept": "*/*",
                  "Content-Type": "multipart/form-data; boundary=----WebKitFormBoundaryTMM4W2B1ALSDGYN6",
                  "lanugage": "zh",
                  "timeout": timeout || 100000
              },
              data: params
          }).success(
              function (data, status, header, config) {
                  deferred.resolve(data);
              }).error(
              function (data, status, header, config) {
                  deferred.reject(data);
              });
          return deferred.promise;
      },
    //TODO:调取本地json
    getJson: function (url) {
      var deferred = $q.defer();
      $http.get(APP_CONFIG.apiUrls.APP_HOST + url).success(function (data) {
        deferred.resolve(data);
      });
      return deferred.promise;
    }
  };
  return factory;
}]).factory("$LocalStorage", ["$CookieStorage", function (e) {
  if (navigator.userAgent.indexOf('UCBrowser') > -1) {
    return e;
  }
  var t = {
    key: function (e) {
      return window.localStorage.key(e)
    },
    setItem: function (e, t) {
      return window.localStorage.setItem(e, t)
    },
    getItem: function (e) {
      return window.localStorage.getItem(e)
    },
    removeItem: function (e) {
      return window.localStorage.removeItem(e)
    },
    clear: function () {
      return window.localStorage.clear()
    },
    getAll: function () {
      return window.localStorage.valueOf();
    }
  };
  try {
    return null !== window.localStorage ? (window.localStorage.setItem("testkey", "foo"),
        window.localStorage.removeItem("testkey"),
        t) : e
  } catch (a) {
    return e
  }
}
]).factory("$CookieStorage", ['Cookies', function (Cookies) {
  return {
    key: function (t) {
      return Cookies.getItem(t)
    },
    setItem: function (t, a) {
      Cookies.setItem(t, a)
    },
    getItem: function (t) {
      return Cookies.getItem(t)
    },
    removeItem: function (t) {
      return Cookies.removeItem(t)
    },
    clear: function () {
      for (var t in Cookies.keys)
        Cookies.removeItem(t);
      return !0
    }
  }
}]).factory('Cookies', function () {
  return {
    getItem: function (sKey) {
      return decodeURI(document.cookie.replace(new RegExp('(?:(?:^|.*;)\\s*' + encodeURI(sKey).replace(/[\-\.\+\*]/g, '\\$&') + '\\s*\\=\\s*([^;]*).*$)|^.*$'), '$1')) || null;
    },
    setItem: function (sKey, sValue, vEnd, sPath, sDomain, bSecure) {
      if (!sKey || /^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)) {
        return false;
      }
      var sExpires = '';
      if (vEnd) {
        switch (vEnd.constructor) {
          case Number:
            sExpires = vEnd === Infinity ? '; expires=' + 'Fri, 31 Dec 9999 23:59:59 GMT' : '; max-age=' + vEnd;
            break;
          case String:
            sExpires = '; expires=' + vEnd;
            break;
          case Date:
            sExpires = '; expires=' + vEnd.toGMTString();
            break;
        }
      }
      document.cookie = encodeURI(sKey) + '=' + encodeURI(sValue) + sExpires + (sDomain ? '; domain=' + sDomain : '') + (sPath ? '; path=' + sPath : '') + (bSecure ? '; secure' : '');
      return true;
    },
    removeItem: function (sKey, sPath) {
      if (!sKey || !this.hasItem(sKey)) {
        return false;
      }
      document.cookie = encodeURI(sKey) + '=; expires=Thu, 01 Jan 1970 00:00:00 GMT' + (sPath ? '; path=' + sPath : '');
      return true;
    },
    hasItem: function (sKey) {
      return (new RegExp('(?:^|;\\s*)' + encodeURI(sKey).replace(/[\-\.\+\*]/g, '\\$&') + '\\s*\\=')).test(document.cookie);
    },
    keys: /* optional method: you can safely remove it! */ function () {
      var aKeys = document.cookie.replace(/((?:^|\s*;)[^\=]+)(?=;|$)|^\s*|\s*(?:\=[^;]*)?(?:\1|$)/g, '').split(/\s*(?:\=[^;]*)?;\s*/);
      for (var nIdx = 0; nIdx < aKeys.length; nIdx++) {
        aKeys[nIdx] = decodeURI(aKeys[nIdx]);
      }
      return aKeys;
    }

  }
}).factory('popupSvc', function () {
  return {
    smartMessageBox: function (title, f1) {
      $.SmartMessageBox({
        title: title,
        content: "",
        buttons: '[No][Yes]'
      }, function (ButtonPressed) {
        if (ButtonPressed === "Yes") {
          if (f1) {
            f1()
          }
          ;
        }
        if (ButtonPressed === "No") {

        }
      });
    },
    smallBox: function (type, title, timeout) {
      if (type === "success") {
        $.smallBox({
          title: title,
          content: "<i class='fa fa-clock-o'></i> <i></i>",
          color: "#659265",
          iconSmall: "fa fa-check fa-2x fadeInRight animated",
          timeout: timeout || 4000
        });
      }
      if (type === "fail") {
        $.smallBox({
          title: title,
          content: "<i class='fa fa-clock-o'></i> <i></i>",
          color: "#C46A69",
          iconSmall: "fa fa-times fa-2x fadeInRight animated",
          timeout: timeout || 4000
        });
      }

    }
  };
});

