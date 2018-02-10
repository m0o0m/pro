"use strict";

angular.module('app').controller("LanguagesCtrl",  function LanguagesCtrl($scope, $rootScope, $log, Language){

    
    Language.getLanguages(function(data){

        $rootScope.currentLanguage = data[0];

        $rootScope.languages = data;

        Language.getLang(data[0].key,function(data){

            $rootScope.lang = data;
        });

    });

    $scope.selectLanguage = function(language){
        $rootScope.currentLanguage = language;

        Language.getLang(language.key,function(data){

            $rootScope.lang = data;

        });
    }
});