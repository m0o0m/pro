'use strict';

/**
 * @ngdoc overview
 * @name app [smartadminApp]
 * @description
 * # app [smartadminApp]
 *
 * Main module of the application.
 */
//配置文件
window.appConfig = appConfig;

$(function () {
  // moment.js default language
  moment.locale('en');
  angular.bootstrap(document, ['app']);
});


angular.module('app', [
    'ngSanitize',
    'ngAnimate',
    'restangular',
    'ui.router',
    'ui.bootstrap',

    // Smartadmin Angular Common Module
    'SmartAdmin',
    //框架
    'app.layout',
    'app.tables',
    'app.filter',
    //登陆验证的
    'app.auth',
    "services.auth",
    "services.account",
    "services.analysisExit",
    "services.cashSystem",
    "services.balanceStatistics",
    'app.administrators',
    'app.directive',
    'app.resource',
    'app.CashManagement',
    'app.analysisExit',
    'app.balanceStatistics',
    'app.CashSystem',
    'services.accessMoney',
    'app.AccessMoney',
    'app.Income',
    'app.Summary',
    'app.notice',
    'services.notice',
    'app.memberMessage',
    'services.memberMessage',
    'services.accessMoney',
    'app.announcement',
    'services.announce',
    'app.operation',
    'services.operation',
    'app.membershipReturns',
    'services.membershipReturns',
    'app.quotaStatistics',
    'services.quotaStatistics',
    'app.websiteManagement',
    "services.websiteInformation",
    "services.videoManagement",
    "services.sportManagement",
    "services.electronicManagement",
    "services.lotteryManagrment",
    "services.lotterySort",
    "services.modularManagement",
    "services.preferencesSettings",
    "services.cacheManagement",
    'app.AuditQuery',
    'app.AuditLog',
    'app.PaymentSetting',
    'services.paymentSetting',
    'app.GraphicEditor',
    "services.site",
    "services.logo",
    "services.float",
    "services.floatManagrment",
    "services.attachment",
    "services.swiper",
    "services.announcement",
    "services.caseEditor",
    'app.caseEditor',
    'app.copyEditor',
    'services.copyEditor',
    'app.ReportForm',
    'services.reportform',
    "app.CommissionStatistics",
    "services.commission",
    "services.agentSetting",
    "services.period",
    "services.feeSetting",
    "app.Precalcula",
    "services.precalcula",
    "services.preferentialQuiry",
    "services.advertisement"

]).config(function ($provide, $httpProvider, RestangularProvider) {

  // Intercept http calls.
  $provide.factory('ErrorHttpInterceptor', function ($q) {
    var errorCounter = 0;

    function notifyError(rejection) {
      console.log(rejection);
      $.bigBox({
        title: rejection.status + ' ' + rejection.statusText,
        content: rejection.data,
        color: "#C46A69",
        icon: "fa fa-warning shake animated",
        number: ++errorCounter,
        timeout: 6000
      });
    }

    return {
      // On request failure
      requestError: function (rejection) {
        // show notification
        notifyError(rejection);

        // Return the promise rejection.
        return $q.reject(rejection);
      },

      // On response failure
      responseError: function (rejection) {
        // show notification
        notifyError(rejection);
        // Return the promise rejection.
        return $q.reject(rejection);
      }
    };
  });

  // Add the interceptor to the $httpProvider.
  $httpProvider.interceptors.push('ErrorHttpInterceptor');

  RestangularProvider.setBaseUrl(location.pathname.replace(/[^\/]+?$/, ''));

}).constant('APP_CONFIG', window.appConfig)
  .run(function ($rootScope, $state, $stateParams, Language) {
    $rootScope.lang = {};
    $rootScope.$state = $state;
    $rootScope.$stateParams = $stateParams;

    $rootScope.getWord = function (key) {
      if (angular.isDefined($rootScope.lang[key])) {
        return $rootScope.lang[key];
      }
      else {
        return key;
      }
    };
  });


