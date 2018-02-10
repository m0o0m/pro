angular.module('app.copyEditor', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.copyEditor').config(function ($stateProvider) {

    $stateProvider
        .state('app.CopyEditor', {
            abstract: true,
            data: {
                // title: '文案编辑'
                title: 'CopyEditor'
            }
        })
        .state('app.CopyEditor.HomeCopy', {
            url: '/CopyEditor/HomeCopy',
            data: {
                // title: '首页文案'
                title: 'HomeCopy'
            },
            views: {
                "content@app": {
                    controller: 'HomeCopyCtrl',
                    templateUrl: "views/copyEditor/views/HomeCopy.html"
                }
            }
        })
        .state('app.CopyEditor.HomeEditor', {
            url: '/CopyEditor/HomeEditor',
            data: {
                //123 title: '首页文案编辑'
                title: '首页文案编辑'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'HomeEditorCtrl',
                    templateUrl: "views/copyEditor/views/HomeEditor.html"
                }
            },
            resolve: {
                srcipts: function (lazyScript) {
                    return lazyScript.register([
                        "vendor.ui.js"
                    ])
                }
            }
        })
        .state('app.CopyEditor.DepositCopy', {
            url: '/CopyEditor/DepositCopy',
            data: {
                // title: '存款文案'
                title: 'DepositCopy'
            },
            views: {
                "content@app": {
                    controller: 'DepositCopyCtrl',
                    templateUrl: "views/copyEditor/views/DepositCopy.html"
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
        .state('app.CopyEditor.DepositEditor', {
            url: '/CopyEditor/DepositEditor',
            data: {
                //123 title: '存款文案编辑'
                title: '存款文案编辑'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'DepositEditorCtrl',
                    templateUrl: "views/copyEditor/views/DepositEditor.html"
                }
            }
        })
        .state('app.CopyEditor.Mould', {
            url: '/CopyEditor/Mould',
            data: {
                //123 title: '模板选择'
                title: '模板选择'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'MouldCtrl',
                    templateUrl: "views/copyEditor/views/Mould.html"
                }
            }
        })
        .state('app.CopyEditor.RegisterCopy', {
            url: '/CopyEditor/RegisterCopy',
            data: {
                // title: '注册文案'
                title: 'RegisterCopy'
            },
            views: {
                "content@app": {
                    controller: 'RegistCopyCtrl',
                    templateUrl: "views/copyEditor/views/RegisterCopy.html"
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
        .state('app.CopyEditor.RegisterEditor', {
            url: '/CopyEditor/RegisterEditor',
            data: {
                // title: '注册文案编辑'
                title: '注册文案编辑'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'RegisterEditorCtrl',
                    templateUrl: "views/copyEditor/views/RegisterEditor.html"
                }
            }
        })
        .state('app.CopyEditor.RegisterMould', {
            url: '/CopyEditor/RegisterMould',
            data: {
                //123 title: '注册模板选择'
                title: '注册模板选择'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'RegisterMouldCtrl',
                    templateUrl: "views/copyEditor/views/RegisterMould.html"
                }
            }
        })
        .state('app.CopyEditor.Discount', {
            url: '/CopyEditor/Discount',
            data: {
                // title: '优惠活动'
                title: 'Discount'
            },
            views: {
                "content@app": {
                    controller: 'DiscountCtrl',
                    templateUrl: "views/copyEditor/views/Discount.html"
                }
            }
        })
        .state('app.CopyEditor.discountContent', {
            url: '/CopyEditor/discountContent',
            data: {
                // title: '优惠活动'
                title: 'DiscountContent'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'discountContentCtrl',
                    templateUrl: "views/copyEditor/views/discountContent.html"
                }
            }
        })
        .state('app.CopyEditor.wapDiscountContent', {
            url: '/CopyEditor/wapDiscountContent',
            data: {
                // title: '优惠活动'
                title: 'wapDiscountContent'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'wapDiscountContentCtrl',
                    templateUrl: "views/copyEditor/views/wapDiscountContent.html"
                }
            }
        })
        .state('app.CopyEditor.contentModifyDiscount', {
            url: '/CopyEditor/contentModifyDiscount',
            data: {
                // title: '优惠活动'
                title: 'contentModifyDiscount'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'contentModifyDiscountCtrl',
                    templateUrl: "views/copyEditor/views/contentModifyDiscount.html"
                }
            }
        })
        .state('app.CopyEditor.wapContentModifyDiscount', {
            url: '/CopyEditor/wapContentModifyDiscount',
            data: {
                // title: '优惠活动'
                title: 'wapContentModifyDiscount'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'wapContentModifyDiscountCtrl',
                    templateUrl: "views/copyEditor/views/wapContentModifyDiscount.html"
                }
            }
        })
        .state('app.CopyEditor.ModifyDiscount', {
            url: '/CopyEditor/ModifyDiscount',
            data: {
                //123 title: '编辑优惠活动'
                title: '编辑优惠活动'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'ModifyDiscountCtrl',
                    templateUrl: "views/copyEditor/views/ModifyDiscount.html"
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
        .state('app.CopyEditor.WapDiscount', {
            url: '/CopyEditor/WapDiscount',
            data: {
                // title: 'Wap优惠活动'
                title: 'WapDiscount'
            },
            views: {
                "content@app": {
                    controller: 'WapDiscountCtrl',
                    templateUrl: "views/copyEditor/views/WapDiscount.html"
                }
            }
        })
        .state('app.CopyEditor.WapModifyDiscount', {
            url: '/CopyEditor/WapModifyDiscount',
            data: {
                //123 title: '编辑Wap优惠活动'
                title: '编辑Wap优惠活动'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'WapModifyDiscountCtrl',
                    templateUrl: "views/copyEditor/views/WapModifyDiscount.html"
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
        .state('app.CopyEditor.LineDetection', {
            url: '/CopyEditor/LineDetection',
            data: {
                // title: '线路检测'
                title: 'LineDetection'
            },
            views: {
                "content@app": {
                    controller: 'LineDetectionCtrl',
                    templateUrl: "views/copyEditor/views/LineDetection.html"
                }
            }
        })
});