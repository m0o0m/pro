"use strict";

angular.module('app.tables', [ 'ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.tables').config(function ($stateProvider) {

    $stateProvider
        .state('app.tables', {
            abstract: true,
            data: {
                title: 'Tables'
            }
        })
        .state('app.tables.normal', {
            url: '/tables/normal',
            data: {
                title: 'Normal Tables'
            },
            views: {
                "content@app": {
                    templateUrl: "app/views/tables/views/normal.html"

                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        "vendor.ui.js"
                    ])

                }
            }
        })

        .state('app.tables.datatables', {
            url: '/tables/datatables',
            data: {
                title: 'Data Tables'
            },
            views: {
                "content@app": { //view 的入口
                    controller: 'DatatablesCtrl as datatables',
                    templateUrl: "app/views/tables/views/datatables.html"
                }
            }
        })

        .state('app.tables.jqgrid', {
            url: '/tables/jqgrid',
            data: {
                title: 'Jquery Grid'
            },
            views: {
                "content@app": {
                    controller: 'JqGridCtrl',
                    templateUrl: "app/views/tables/views/jqgrid.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register([
                        'app/smartadmin-plugin/legacy/jqgrid/js/minified/jquery.jqGrid.min.js',
                        'app/smartadmin-plugin/legacy/jqgrid/js/i18n/grid.locale-cn.js'
                    ])

                }
            }
        })
});