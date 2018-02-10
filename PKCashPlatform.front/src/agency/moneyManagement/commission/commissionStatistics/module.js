angular.module("app.CommissionStatistics",['ui.router', 'datatables', 'datatables.bootstrap']).config(function ($stateProvider) {
    $stateProvider
        .state('app.CommissionStatistics',{
            abstract:true,
            data:{
                title:"CommissionStatistics"
            }
        })
        .state('app.CommissionStatistics.CommissionStatistics',{
            url: '/CommissionStatistics/CommissionStatistics',
            data: {
                title: 'CommissionStatistics'
            },
            views:{
                'content@app':{
                    controller:'CommissionStatCtrl',
                    templateUrl: 'views/moneyManagement/commission/CommissionStatistics/views/CommissionStatistics.html'
                }
            }
        })
        .state('app.CommissionStatistics.list',{
            url: '/CommissionStatistics/list',
            data: {
                title: 'CommissionStatistics'
            },
            views:{
                'content@app':{
                    controller:'listCtrl',
                    templateUrl: 'views/moneyManagement/commission/CommissionStatistics/views/list.html'
                }
            }
        })
        .state('app.CommissionStatistics.AgentSetting',{
            url: '/CommissionStatistics/AgentSetting',
            data: {
                title: 'AgentRetirementSetting'
            },
            views:{
                'content@app':{
                    controller:'AgentSettingCtrl',
                    templateUrl: 'views/moneyManagement/commission/CommissionStatistics/views/AgentSetting.html'
                }
            }
        })
        .state('app.CommissionStatistics.modify',{
            url: '/CommissionStatistics/modify',
            data: {
                title: 'RebateSetting'
            },
            params:{id:null},
            views:{
                'content@app':{
                    controller:'modifyCtrl',
                    templateUrl: 'views/moneyManagement/commission/CommissionStatistics/views/modify.html'
                }
            }
        })
        .state('app.CommissionStatistics.PeriodManagement', {
            url: '/CommissionStatistics/PeriodManagement',
            data: {
                title: 'PeriodManagement'
            },
            views: {
                "content@app": {
                    controller: 'PeriodManagementCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/PeriodManagement.html"
                }
            }
        })
        .state('app.CommissionStatistics.NewlyAdd', {
            url: '/CommissionStatistics/NewlyAdd',
            data: {
                title: 'MemberPromotionSettings'
            },
            views: {
                "content@app": {
                    controller: 'NewlyAddCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/NewlyAdd.html"
                }
            }
        })
        .state('app.CommissionStatistics.FeeSetting', {
            url: '/CommissionStatistics/FeeSetting',
            data: {
                title: 'FeeSetting'
            },
            views: {
                "content@app": {
                    controller: 'FeeSettingCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/FeeSetting.html"
                }
            }
        })
        .state('app.CommissionStatistics.AddFeeSetting', {
            url: '/CommissionStatistics/AddFeeSetting',
            data: {
                //123 title: '添加返佣设定'
                title: 'RebateSetting'
            },
            views: {
                "content@app": {
                    controller: 'AddFeeSettingCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/AddFeeSetting.html"
                }
            }
        })
        .state('app.CommissionStatistics.ModifyFeeSetting', {
            url: '/CommissionStatistics/ModifyFeeSetting/:id',
            data: {
                //123  title: '修改返佣设定'
                title: 'RebateSetting'
            },
            views: {
                "content@app": {
                    controller: 'ModifyFeeSettingCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/ModifyFeeSetting.html"
                }
            }
        })
        .state('app.CommissionStatistics.ModificationFee', {
            url: '/CommissionStatistics/ModificationFee?id',
            data: {
                title: 'FeeSetting'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'ModificationFeeCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/ModificationFee.html"
                }
            }
        })
        .state('app.CommissionStatistics.Commission', {
            url: '/CommissionStatistics/Commission',
            data: {
                title: 'CommissionInquiry'
            },
            views: {
                "content@app": {
                    controller: 'CommissionCtrl',
                    templateUrl: "views/moneyManagement/commission/CommissionStatistics/views/Commission.html"
                }
            }
        })

});