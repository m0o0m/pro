'use strict';

angular.module('app.administrators').directive('repeatDone', function () {
    return {
        link: function(scope, element, attrs) {
            if (scope.$last) {                   // 这个判断意味着最后一个 OK
                scope.$eval(attrs.repeatDone)    // 执行绑定的表达式
            }
        }
    }
});