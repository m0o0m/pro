(function () {
    "use strict";
    //TODO 修改菜单
    angular.module('SmartAdmin.Layout').directive('smartMenuItems', function ($http, $rootScope, $compile,httpSvc) {
        return {
            restrict: 'A',
            //template: '<p>Hello {{number}} {{getWord(\'Dashboard\')}}!</p>',
            controller: function ($scope, $element) {
                $scope.getWord = function (key) {
                    if (angular.isDefined($rootScope.getWord(key))) {
                        return $rootScope.getWord(key);
                    }
                    else {
                        return key;
                    }
                };
            },
            template: function (element, attrs) {
                function createItem(item, parent, level) {
                    var li = $('<li />', {'ui-sref-active': "active"})
                    var a = $('<a />');
                    var i = $('<i />');

                    li.append(a);

                    if (item.route)
                        a.attr('ui-sref', item.route);
                    // if (item.href)
                    a.attr('href', "#");
                    if (item.icon) {
                        //'fa fa-lg fa-fw fa-' +
                        i.attr('class',  'icon iconfont ' + item.icon);
                        a.append(i);
                    }
                    if (item.menu_name) {
                        a.attr('title', item.menu_name);
                        if (level > 1) {
                            a.append(' {{getWord(\'' + item.language_key + '\')}}');
                        } else {
                            a.append(' <span class="menu-item-parent">{{getWord(\'' + item.language_key + '\')}}</span>');
                        }
                    }

                    if (item.Children.length>0) {
                        var ul = $('<ul />');
                        li.append(ul);
                        li.attr('data-menu-collapse', '');
                        _.forEach(item.Children, function (child) {
                            createItem(child, ul, level + 1);
                        })
                    }

                    parent.append(li);
                }

                httpSvc.get("/menu/role").then(function (res) {
                    var ul = $('<ul />', {
                        'smart-menu': ''
                    });
                    _.forEach(res.data.data, function (item) {

                        createItem(item, ul, 1);
                    });

                    var $scope = $rootScope.$new();
                    var html = $('<div>').append(ul).html();
                    var linkingFunction = $compile(html);

                    var _element = linkingFunction($scope);

                    element.replaceWith(_element);
                });
            }
        };
    });
})();
