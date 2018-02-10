angular.module('app.PaymentSetting', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.PaymentSetting', {
            abstract: true,
            data: {
                title: 'PaymentSetting'
            }
        })
        .state('app.PaymentSetting.BankElimination', {
            url: '/PaymentSetting/BankElimination',
            data: {
                title: 'BankElimination'
            },
            views: {
                "content@app": {
                    controller: 'BankEliminationCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/BankElimination.html"
                }
            }
        })
        .state('app.PaymentSetting.DepositBank', {
            url: '/PaymentSetting/DepositBank',
            data: {
                title: 'DepositBankSet'
            },
            views: {
                "content@app": {
                    controller: 'DepositBankCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/DepositBank.html"
                }
            }
        })
        .state('app.PaymentSetting.DepositRecord', {
            url: '/PaymentSetting/DepositRecord/:Id',
            data: {
                title: 'DepositBankSet'
            },
            params:{value1:null,value2:null,value3:null,value4:null,value5:null,value6:null},
            views: {
                "content@app": {
                    controller: 'DepositRecordCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting//views/DepositRecord.html"
                }
            }
        })
        .state('app.PaymentSetting.ModifySettings', {
            url: '/PaymentSetting/ModifySettings',
            data: {
                title: 'DepositBankSet'
            },
            params:{value1:null,value2:null,value3:null,value4:null,value5:null,value6:null},
            views: {
                "content@app": {
                    controller: 'ModifySettingsCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/ModifySettings.html"
                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ])

                }
            }
        })
        .state('app.PaymentSetting.Modify', {
            url: '/PaymentSetting/Modify/:Id',
            data: {
                title: 'DepositBankSet'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'ModifyCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/Modify.html"
                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ])

                }
            }
        })
        .state('app.PaymentSetting.ThirdBankElimin', {
            url: '/PaymentSetting/ThirdBankElimin',
            data: {
                title: 'bankElimination'
            },
            views: {
                "content@app": {
                    controller: 'ThirdBankEliminCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/ThirdBankElimin.html"
                }
            }
        })
        .state('app.PaymentSetting.OnlinePayment', {
            url: '/PaymentSetting/OnlinePayment',
            data: {
                title: 'OnlinePaymentSetup'
            },
            views: {
                "content@app": {
                    controller: 'OnlinePaymentCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/OnlinePayment.html"
                }
            }
        })
        .state('app.PaymentSetting.NewSettings', {
            url: '/PaymentSetting/NewSettings',
            data: {
                title: 'OnlinePaymentSetup'
            },
            views: {
                "content@app": {
                    controller: 'NewSettingsCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/NewSettings.html"
                }
            }
        })
        .state('app.PaymentSetting.PayParameter', {
            url: '/PaymentSetting/PaymentParameter',
            data: {
                title: 'PayParameter'
            },
            views: {
                "content@app": {
                    controller: 'PayParameterCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/PayParameter.html"
                }
            }
        })
        .state('app.PaymentSetting.DetailsSetting', {
            url: '/PaymentSetting/DetailsSetting/:id',
            data: {
                title: 'PaymentParameterSetting'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'DetailsSettingCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/DetailsSetting.html"
                }
            }
        })
        .state('app.PaymentSetting.AutomaticSetting', {
            url: '/PaymentSetting/AutomaticSetting',
            data: {
                title: 'AutomaticCashSetting'
            },
            views: {
                "content@app": {
                    controller: 'AutomaticSettingCtrl',
                    templateUrl: "app/views/PaymentSetting/views/AutomaticSetting.html"
                }
            }
        })
        .state('app.PaymentSetting.excludingBank', {
            url: '/PaymentSetting/excludingBank',
            data: {
                title: 'IncomeExcludingBank'
            },
            views: {
                "content@app": {
                    controller: 'excludingBankCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/excludingBank.html"
                }
            }
        })
        .state('app.PaymentSetting.modifyOnline', {
            url: '/PaymentSetting/modifyOnline/:Id',
            data: {
                title: 'OnlinePaymentSetup'
            },
            views: {
                "content@app": {
                    controller: 'modifyOnlineCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/modifyOnline.html"
                }
            }
        })
        .state('app.PaymentSetting.seedetail', {
            url: '/PaymentSetting/seedetail/:Id',
            data: {
                title: 'Detailed content'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'seedetailCtrl',
                    templateUrl: "views/moneyManagement/PaymentSetting/views/seedetail.html"
                }
            }
        })

});
