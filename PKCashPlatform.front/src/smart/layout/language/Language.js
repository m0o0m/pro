"use strict";

angular.module('app').factory('Language', function ($http, APP_CONFIG) {

    function getLanguage(key, callback) {

        $http.get(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.LANG[key]).success(function (data) {

            callback(data);

        }).error(function () {
            callback([]);

        });

    }

    function getLanguages(callback) {

        $http.get(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.LANG["index"]).success(function (data) {

            callback(data);

        }).error(function () {
            callback([]);

        });

    }

    return {
        getLang: function (type, callback) {
            getLanguage(type, callback);
        },
        getLanguages: function (callback) {
            getLanguages(callback);
        }
    }

});