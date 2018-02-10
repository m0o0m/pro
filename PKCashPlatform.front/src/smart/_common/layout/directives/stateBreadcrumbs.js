
'use strict';

angular.module('SmartAdmin.Layout').directive('stateBreadcrumbs', function ($rootScope, $state, $compile) {
    return {
        restrict: 'EA',
        controller: function ($scope, $element) {
            $scope.getWord = function (key) {
                if ( angular.isDefined($rootScope.lang) && angular.isDefined($rootScope.lang[key])) {
                    return $rootScope.lang[key];
                }
                else {
                    return key;
                }
            }

        },
        template: function (element, attrs) {
            function setBreadcrumbs(breadcrumbs) {
                var ol = $('<ol />');
                ol.attr('class',  'breadcrumb');
                angular.forEach(breadcrumbs, function (crumb) {
                    var li = $('<li />');
                    li.append(' {{getWord(\'' + crumb + '\')}}');
                    ol.append(li);
                });
                var $scope = $rootScope.$new();
                var html = $('<div>').append(ol).html();
                var linkingFunction = $compile(html);

                var _element = linkingFunction($scope);
                // element.replaceWith(_element);
                element.append(_element);
                if($(element[0]).find("ol").length>1){
                    var el=$(element[0]).find("ol")[0];
                    el.remove();
                }

            }

            function fetchBreadcrumbs(stateName, breadcrunbs) {

                var state = $state.get(stateName);
                if (state && state.data && state.data.title && breadcrunbs.indexOf(state.data.title) == -1) {
                    breadcrunbs.unshift(state.data.title);
                }

                var parentName = stateName.replace(/.?\w+$/, '');
                if (parentName) {
                    return fetchBreadcrumbs(parentName, breadcrunbs);
                } else {
                    return breadcrunbs;
                }
            }

            function processState(state) {
                var breadcrumbs;
                if (state.data && state.data.breadcrumbs) {
                    breadcrumbs = state.data.breadcrumbs;
                } else {
                    breadcrumbs = fetchBreadcrumbs(state.name, []);
                }
                setBreadcrumbs(breadcrumbs);
            }

            processState($state.current);

            $rootScope.$on('$stateChangeStart', function (event, state) {
                processState(state);
            })
        }
    }
});