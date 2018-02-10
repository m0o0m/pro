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
    //
    // Smartadmin Angular Common Module
    'SmartAdmin',
    //框架
    'app.layout',
    'app.filter',
    //登陆验证的
    'app.dashboard',
    'app.directive',
    'app.auth',
    'services.auth',
    'app.site',
    'services.site',
    'app.CopyTemplate',
    'services.copyTemplate',
    'app.ReportForm',
    'services.common',
    'services.financeHierarchical',
    'services.financeCash',
    'services.financeDataCenter',
    'services.financeReferential',
    'services.financeIncome',
    'services.financeSummary',
    'services.financeArrears',
    'services.financeReport',
    'services.financeReportStatistics',
    'services.financeQuotaNum',
    'services.financeQuotaRecord',
    'services.financeRechargeRecord',
    'services.financeMoneyManagement',
    'services.customer',
    'app.customer',
    'app.Platform',
    'services.platform',
    "app.GraphicEditor",
    'services.logo',
    'app.copyEditor',
    'services.copyEditor',
    'services.attachment',
    'services.swiper',
    'services.customerVideo',
    'services.customerExceptionMember',
    'services.customerCommonbullet',
    'app.announcement',
    'services.announce',
    'services.advertisement',
    'services.configurationsetting',
    'services.customerPreferentialQuery',
    'services.customerApplicationInquiry',
    //"FileManagerApp",

]).config(function ($provide, $httpProvider, RestangularProvider) {

    // Intercept http calls.
    $provide.factory('ErrorHttpInterceptor', function ($q) {
        var errorCounter = 0;

        function notifyError(rejection) {
            console.log("=======",rejection);
            $.bigBox({
                title: "网络错误",
                content: "网络错误，请检查网络环境。。。",
                color: "#C46A69",
                icon: "fa fa-warning shake animated",
                timeout: 2000
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


