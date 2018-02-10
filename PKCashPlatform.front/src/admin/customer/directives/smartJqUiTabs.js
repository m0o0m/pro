/**
 * Created by apple on 17/8/18.
 */
'use strict';

angular.module('app.customer').directive('smartJquiTabs', function () {
    return {
        restrict: 'A',
        link: function (scope, element, attributes) {

            element.tabs();
        }
    }
});