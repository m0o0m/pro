angular.module('app').factory('activityService', function ($http, $log, APP_CONFIG) {

    function getActivities(callback) {

        $http.get(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.ACTIVITY["index"]).success(function (data) {

            callback(data);

        }).error(function () {

            $log.log('Error');
            callback([]);

        });

    }

    function getActivitiesByType(type, callback) {

        $http.get(APP_CONFIG.apiUrls.HOST + APP_CONFIG.apiUrls.ACTIVITY[type]).success(function (data) {

            callback(data);

        }).error(function () {

            $log.log('Error');
            callback([]);

        });

    }

    return {
        get: function (callback) {
            getActivities(callback);
        },
        getbytype: function (type, callback) {
            getActivitiesByType(type, callback);
        }
    }
});