angular.module('app.CopyTemplate', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.CopyTemplate').config(function ($stateProvider) {
    $stateProvider
        .state('app.CopyTemplate', {
            abstract: true,
            data: {
                title: '文案模板'
            }
        })
        .state('app.CopyTemplate.RegisterCopy', {
            url: '/CopyTemplate/RegisterCopy',
            data: {
                title: '注册文案模板'
            },
            views: {
                "content@app": {
                    controller: 'registerCopyCtrl',
                    templateUrl: "views/CopyTemplate/views/RegisterCopy.html"
                }
            }
        })
        .state('app.CopyTemplate.AddRegisterCopy', {
            url: '/CopyTemplate/AddRegisterCopy',
            data: {
                title: '添加注册文案模板'
            },
            views: {
                "content@app": {
                    controller: 'AddRegisterCopyCtrl',
                    templateUrl: "views/CopyTemplate/views/AddRegisterCopy.html"
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
        .state('app.CopyTemplate.ModifyRegisterCopy', {
            url: '/CopyTemplate/ModifyRegisterCopy',
            data: {
                title: '修改注册文案模板'
            },
            params:{id:null,data:null},
            views: {
                "content@app": {
                    controller: 'ModifyRegisterCopyCtrl',
                    templateUrl: "views/CopyTemplate/views/ModifyRegisterCopy.html"
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
        .state('app.CopyTemplate.VideoCopy', {
            url: '/CopyTemplate/VideoCopy',
            data: {
                title: '视讯文案模板'
            },
            views: {
                "content@app": {
                    controller: 'VideoCopyCtrl',
                    templateUrl: "views/CopyTemplate/views/VideoCopy.html"
                }
            }
        })

});