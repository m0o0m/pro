angular.module("app.membershipReturns",['ui.router', 'datatables', 'datatables.bootstrap']).config(function ($stateProvider) {
    $stateProvider
        .state('app.MembershipReturns',{
            abstract:true,
            data:{
                title:"MembershipReturns"
            }
        })
        .state('app.MembershipReturns.MembershipReturns',{
            url: '/MembershipReturns/MembershipReturns',
            data: {
                title: 'MembershipReturns'
            },
            views:{
                'content@app':{
                    controller:'MembershipReturnsCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/MembershipReturns.html'
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
        .state('app.MembershipReturns.TranslationResult',{
            url: '/MembershipReturns/TranslationResult',
            data: {
                title: 'MembershipReturns'
            },
            params:{
                site:null,startTime:null,endTime:null,condition:null
            },
            views:{
                'content@app':{
                    controller:'TranslationResultCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/TranslationResult.html'
                }
            }
        })
        .state('app.MembershipReturns.MemberPromotion',{
            url: '/MembershipReturns/MemberPromotion',
            data: {
                title: 'MemberPromotionInquiries'
            },
            views:{
                'content@app':{
                    controller:'MemberPromotionCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/MemberPromotion.html'
                }
            }
        })
        .state('app.MembershipReturns.setting',{
            url: '/MembershipReturns/setting',
            data: {
                title: 'CommissionSetting'
            },
            views:{
                'content@app':{
                    controller:'settingCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/setting.html'
                }
            }
        })
        .state('app.MembershipReturns.Discount',{
            url: '/MembershipReturns/Discount',
            data: {
                title: 'RebateSetting'
            },
            views:{
                'content@app':{
                    controller:'DiscountsCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/Discount.html'
                }
            }
        })
        .state('app.MembershipReturns.SetUp',{
            url: '/MembershipReturns/SetUp',
            data: {
                title: 'RebateSetting'
            },
            views:{
                'content@app':{
                    controller:'SetUpCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/SetUp.html'
                }
            }
        })
        .state('app.MembershipReturns.Detail',{
            url: '/MembershipReturns/Detail',
            data: {
                title: 'MembershipReturns'
            },
            params:{detailedID:null},
            views:{
                'content@app':{
                    controller:'DetailCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/Detail.html'
                }
            }
        })
        .state('app.MembershipReturns.settingAdd',{
            url: '/MembershipReturns/settingAdd',
            data: {
                title: 'RebateSetting'
            },
            views:{
                'content@app':{
                    controller:'settingAddCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/settingAdd.html'
                }
            }
        })
        .state('app.MembershipReturns.settingModify',{
            url: '/MembershipReturns/settingModify',
            data: {
                title: 'RebateSetting'
            },
            params:{discountID:null},
            views:{
                'content@app':{
                    controller:'settingModifyCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/settingModify.html'
                }
            }
        })
        .state('app.MembershipReturns.PromotionSetting',{
            url: '/MembershipReturns/PromotionSetting',
            data: {
                title: 'MemberPromotionSettings'
            },
            views:{
                'content@app':{
                    controller:'PromotionSettingCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/PromotionSetting.html'
                }
            }
        })
        .state('app.MembershipReturns.numDetail',{
            url: '/MembershipReturns/Detail',
            data: {
                title: 'Recommended number'
            },
            params:{id:null},
            views:{
                'content@app':{
                    controller:'DetailCtrl',
                    templateUrl: 'views/moneyManagement/membership/membershipReturns/views/numDetail.html'
                }
            }
        })
});