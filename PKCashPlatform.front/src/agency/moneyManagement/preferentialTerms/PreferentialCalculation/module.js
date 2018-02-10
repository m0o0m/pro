angular.module('app.Precalcula', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.Precalcula', {
            abstract: true,
            data: {
                title: 'PreferentialCalculation'
            }
        })
        .state('app.Precalcula.Precalcula', {
            url: '/Precalcula/Precalcula',
            data: {
                title: 'SetPreferential'
            },
            views: {
                "content@app": {
                    controller: 'PrecalculaCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/Precalcula.html"
                }
            }
        })
        .state('app.Precalcula.NewOffer', {
            url: '/Precalcula/NewOffer',
            data: {
                title: 'SetPreferential'
            },
            views: {
                "content@app": {
                    controller: 'NewOfferCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/NewOffer.html"
                }
            }
        })
        .state('app.Precalcula.PreferentialQuiry', {
            url: '/Precalcula/PreferentialQuiry',
            data: {
                title: 'PreferentialInquiry'
            },
            views: {
                "content@app": {
                    controller: 'PreferentialQuiryCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/PreferentialQuiry.html"
                }
            }
        })
        .state('app.Precalcula.PreferentialStatistics', {
            url: '/Precalcula/PreferentialStatistics',
            data: {
                title: 'PreferentialStatistics'
            },
            views: {
                "content@app": {
                    controller: 'PreferentialStatisticsCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/PreferentialStatistics.html"
                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ]);
                }
            }
        })
        .state('app.Precalcula.query', {
            url: '/Precalcula/query',
            data: {
                title: 'PreferentialStatistics'
            },
            params:{
                site_index_id:null,
                rtype:null,
                v_type:null,
                typeed:null,
                start_time:null,
                end_time:null
            },
            views: {
                "content@app": {
                    controller: 'queryCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/query.html"
                },
                resolve: {
                    srcipts: function(lazyScript){
                        return lazyScript.register([
                            'vendor.ui.js'
                        ]);
                    }
                }
            }
        })
        .state('app.Precalcula.SelfQuery', {
            url: '/Precalcula/SelfQuery',
            data: {
                title: 'Self service return water'
            },
            views: {
                "content@app": {
                    controller: 'SelfQueryCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/SelfQuery.html"
                }
            }
        })
        .state('app.Precalcula.Detailed', {
            url: '/Precalcula/Detailed?id',
            data: {
                title: 'PreferentialInquiry'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'DetailedCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/Detailed.html"
                }
            }
        })
        .state('app.Precalcula.waterdetails', {
            url: '/Precalcula/waterdetails/:Id/:start/:end',
            data: {
                //123 title: '返水明细'
                title: 'returnDetail'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'waterdetailsCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/waterdetails.html"
                }
            }
        })
        .state('app.Precalcula.modiOffer', {
            url: '/Precalcula/modiOffer/:id',
            data: {
                title: 'SetPreferential'
            },
            views: {
                "content@app": {
                    controller: 'modiOfferCtrl',
                    templateUrl: "views/moneyManagement/preferentialTerms/PreferentialCalculation/views/modiOffer.html"
                }
            }
        });

});
