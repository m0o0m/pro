

'use strict';

angular.module('SmartAdmin.Layout').directive('smartPageTitle', function ($rootScope, $timeout) {
    return {
        restrict: 'A',
        compile: function (element, attributes) {
            element.removeAttr('smart-page-title data-smart-page-title');

            var defaultTitle = attributes.smartPageTitle;
            // var listener = function(event, toState, toParams, fromState, fromParams) {
            //     var title = defaultTitle;
            //     if (toState.data && toState.data.title) title = toState.data.title + ' | ' + title;
            //     // Set asynchronously so page changes before title does
            //     $timeout(function() {
            //         $('html head title').text(title);
            //     });
            // };

            //$rootScope.$on('$stateChangeStart', listener);
            $rootScope.$on('$stateChangeStart');
        }
    }
});


//
//
//
// 'use strict';
//
// angular.module('SmartAdmin.Layout').directive('smartPageTitle', function ($rootScope, $timeout,$compile) {
//     return {
//         restrict: 'EA',
//         controller: function ($scope, $element) {
//             $scope.getWord = function (key) {
//                 if (angular.isDefined($rootScope.lang[key])) {
//                     return $rootScope.lang[key];
//                 }
//                 else {
//                     return key;
//                 }
//             }
//
//         },
//         compile: function (element, attributes) {
//             element.removeAttr('smart-page-title data-smart-page-title');
//
//             var defaultTitle = attributes.smartPageTitle;
//             var listener = function(event, toState, toParams, fromState, fromParams) {
//                 var title = defaultTitle;
//                 if (toState.data && toState.data.title) title = toState.data.title + ' | ' + title;
//                 // Set asynchronously so page changes before title does
//                 // var span = $('<li />');
//                 // li.append(' {{getWord(\'' + crumb + '\')}}');
//                 $timeout(function() {
//                     // $('html head title').text(title);
//                     // $('html head title').text(' {{getWord(\'' + title + '\')}}');
//                     // var $scope = $rootScope.$new();
//                     // var html = $('html head title').text();
//                     // var linkingFunction = $compile(html);
//                     //
//                     // var _element = linkingFunction($scope);
//                     // // element.replaceWith(_element);
//                     // console.log(_element);
//                     // element.append(_element);
//                 });
//
//             };
//
//
//             $rootScope.$on('$stateChangeStart', listener);
//
//         }
//     }
// });