"use strict";

angular.module('app').directive('languageSelector', function(Language){
    return {
        restrict: "EA",
        replace: true,
        templateUrl: "views/layout/language/language-selector.tpl.html",
        scope: true
    }
});