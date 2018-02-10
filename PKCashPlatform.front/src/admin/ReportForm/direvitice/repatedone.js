/**
 * Created by apple on 17/11/8.
 */
'use strict';

angular.module('app.ReportForm').directive('repeatDone', function () {
    return {
        link: function(scope, element, attrs) {
            if (scope.$last) {                   // 这个判断意味着最后一个 OK
                scope.$eval(attrs.repeatDone)    // 执行绑定的表达式
            }
        }
    }
});