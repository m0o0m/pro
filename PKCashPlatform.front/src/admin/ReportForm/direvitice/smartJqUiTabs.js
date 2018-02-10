
'use strict';

angular.module('app.ReportForm').directive('smartJquiTabs', function () {
    return {
        restrict: 'A',
        link: function (scope, element, attributes) {

            element.tabs();
        }
    }
});