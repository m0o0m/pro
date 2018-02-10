angular.module('app.administrators', ['ui.router']);

angular.module('app.administrators').config(function ($stateProvider) {

    $stateProvider
        .state('app.administrators', {
            abstract: true,
            data: {
                // title: '账号管理'
                title: 'Account management'
            }
        })

        .state('app.administrators.power', {
            url: '/administrators/power',
            data: {
                // title: '权限分组'
                title: 'Permission Grouping'
            },
            views: {
                "content@app": {
                    controller: 'PowerCtrl as power',
                    templateUrl: "views/account/administrators/views/power.html"
                }
            }

        })
        .state('app.administrators.addPower', {
            url: '/administrators/addPower',
            data: {
                // title: '创建分组'
                title: 'CreatePacket'
            },
            views: {
                "content@app": {
                    controller: 'AddPowerCtrl as addPower',
                    templateUrl: "views/account/administrators/views/addPower.html"
                }
            }

        })
        .state('app.administrators.modifyPower', {
            url: '/administrators/modifyPower',
            data: {
                // title: '编辑分组'
                title: 'EditGroup'
            },
            views: {
                "content@app": {
                    controller: 'ModifyPowerCtrl as ModifyPower',
                    templateUrl: "views/account/administrators/views/modifyPower.html"
                }
            }

        })
        .state('app.administrators.seePower', {
            url: '/administrators/seePower',
            data: {
                // title: '查看权限分组'
                title: 'LookPower'
            },
            views: {
                "content@app": {
                    controller: 'SeePowerCtrl as SeePower',
                    templateUrl: "views/account/administrators/views/seePower.html"
                }
            }

        })
        .state('app.administrators.adminLog', {
            url: '/administrators/adminLog',
            data: {
                // title: '管理日志'
                title: 'Admin Log'
            },
            views: {
                "content@app": {
                    // controller: 'AdminLogCtrl as adminLog',
                    templateUrl: "views/account/administrators/views/adminLog.html"
                }
            }

        })
        .state('app.administrators.accounts', {
            url: '/administrators/accounts',
            data: {
                // title: '会员管理——正式账号'
                title: 'Customer Management'
            },
            params:{accountsid:null,first_id:null,second_id:null,agency_id:null},
            views: {
                "content@app": {
                    controller: 'AccountsCtrl',
                    templateUrl: "views/account/administrators/views/accounts.html"
                }
            }

        })
        .state('app.administrators.playAccounts', {
            url: '/administrators/playAccounts',
            data: {
                // title: '会员管理——带玩账号'
                title: 'Customer Management'
            },
            views: {
                "content@app": {
                    controller: 'PlayAccountsCtrl',
                    templateUrl: "views/account/administrators/views/playAccounts.html"
                }
            }

        })
        .state('app.administrators.addAccounts', {
            url: '/administrators/addAccounts',
            data: {
                // title: '会员管理——添加带玩'
                title: 'Customer Management'
            },
            views: {
                "content@app": {
                    controller: 'AddAccountsCtrl',
                    templateUrl: "views/account/administrators/views/addAccounts.html"
                }
            }

        })
        .state('app.administrators.accountInfo', {
            url: '/administrators/accountInfo',
            data: {
                // title: '会员资料'
                title: 'Customer Info'
            },
            params:{ids:null},
            views: {
                "content@app": {
                    controller: 'AccountInfoCtrl',
                    templateUrl: "views/account/administrators/views/accountInfo.html"
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
        .state('app.administrators.shareholders', {
            url: '/administrators/views/shareholders.html',
            data: {
                title: 'A_2'
            },
            params:{shareid:null},
            views: {
                "content@app": {
                    controller: 'ShareholdersCtrl',
                    templateUrl: "views/account/administrators/views/shareholders.html"
                }
            }

        })
        .state('app.administrators.generalAgent', {
            url: '/administrators/generalAgent',
            data: {
                // title: '总代理管理'
                title: 'A_3'
            },
            // params:{form_value:null},
            params:{Shareholderid:null,form_value:null,first_id:null},
            views: {
                "content@app": {
                    controller: 'GeneralAgentCtrl',
                    templateUrl: "views/account/administrators/views/generalAgent.html"
                }
            }

        })
        .state('app.administrators.agent', {
            url: '/administrators/agent',
            data: {
                // title: '代理管理'
                title: 'A_4'
            },
            params:{first_id:null,form_value:null,gene:null},
            views: {
                "content@app": {
                    controller: 'AgentCtrl',
                    templateUrl: "views/account/administrators/views/agent.html"
                }
            }

        })
        .state('app.administrators.agentEdit', {
            url: '/administrators/agentEdit',
            data: {
                // title: '修改代理资料'
                title: 'Modify Agent'
            },
            params:{editid:null},
            views: {
                "content@app": {
                    controller: 'AgentEditCtrl',
                    templateUrl: "views/account/administrators/views/agentEdit.html"
                }
            }

        })
        .state('app.administrators.agentInfo', {
            url: '/administrators/agentInfo',
            data: {
                // title: '代理资料'
                title: 'Modify Agent Info'
            },
            views: {
                "content@app": {
                    controller: 'AgentInfoCtrl',
                    templateUrl: "views/account/administrators/views/agentInfo.html"
                }
            }
        })
        .state('app.administrators.agentDomain', {
            url: '/administrators/agentDomain',
            data: {
                // title: '代理域名'
                title: 'Agent Domain'
            },
            params:{Domainid:null},
            views: {
                "content@app": {
                    controller: 'AgentDomainCtrl',
                    templateUrl: "views/account/administrators/views/agentDomain.html"
                }
            }

        })
        .state('app.administrators.childAccount', {
            url: '/administrators/childAccount',
            data: {
                // title: '子账号'
                title: 'ChildAccount'
            },
            views: {
                "content@app": {
                    controller: 'ChildAccountCtrl',
                    templateUrl: "views/account/administrators/views/childAccount.html"
                }
            }
        })
        .state('app.administrators.childAccountPermissions', {
            url: '/administrators/childAccountPermissions?id&site_id',
            data: {
                // title: '子账号权限'
                title: 'ChildAccount Power'
            },
            views: {
                "content@app": {
                    controller: 'ChildAccountPermissionsCtrl',
                    templateUrl: "views/account/administrators/views/childAccountPermissions.html"
                }
            }
        })
        .state('app.administrators.hierarchicalManag', {
            url: '/administrators/hierarchicalManag',
            data: {
                // title: '层级管理'
                title: 'A_9'
            },
            views: {
                "content@app": {
                    controller: 'hierarchicalManagCtrl',
                    templateUrl: "views/account/administrators/views/hierarchicalManag.html"
                }
            }

        })
        
        .state('app.administrators.registSetting', {
            url: '/administrators/registSetting',
            data: {
                // title: '会员注册设定'
                title: 'Member registration settings'
            },
            views: {
                "content@app": {
                    controller: 'RegistSettingCtrl',
                    templateUrl: "views/account/administrators/views/registSetting.html"
                }
            }
        })
        .state('app.administrators.MembershipDetails', {
            url: '/administrators/MembershipDetails?site_index_id&level_id',
            data: {
                // title: '会员详情'
                title: 'Membership details'
            },
            views: {
                "content@app": {
                    controller: 'MembershipDetailsCtrl',
                    templateUrl: "views/account/administrators/views/MembershipDetails.html"
                }
            }

        })

        .state('app.administrators.applicationManagement', {
            url: '/administrators/applicationManagement',
            data: {
                // title: '代理申请管理'
                title: 'A_7'
            },
            views: {
                "content@app": {
                    controller: 'ApplicationManagementCtrl',
                    templateUrl: "views/account/administrators/views/applicationManagement.html"
                }
            }
        })
        .state('app.administrators.AddHierarchy', {
            url: '/administrators/AddHierarchy',
            data: {
                // title: '添加层级'
                title: 'Add Heriy'
            },
            views: {
                "content@app": {
                    controller: 'AddHierarchyCtrl',
                    templateUrl: "views/account/administrators/views/AddHierarchy.html"
                }
            }

        })

        .state('app.administrators.systemQuery', {
            url: '/administrators/systemQuery',
            data: {
                // title: '体系查询'
                title: 'A_8'
            },
            views: {
                "content@app": {
                    controller: 'SystemQueryCtrl',
                    templateUrl: "views/account/administrators/views/systemQuery.html"
                }
            }
        })
        .state('app.administrators.ModifyHierarchy', {
            url: '/administrators/ModifyHierarchy?site_index_id&level_id',
            data: {
                // title: '修改层级信息'
                title: 'Modify Heriy'
            },
            views: {
                "content@app": {
                    controller: 'ModifyHierarchyCtrl',
                    templateUrl: "views/account/administrators/views/ModifyHierarchy.html"
                }
            }

        })
        .state('app.administrators.bank', {
            url: '/administrators/bank',
            data: {
                // title: '出款银行'
                title: 'Bank'
            },
            params:{ids:null},
            views: {
                "content@app": {
                    controller: 'bankCtrl',
                    templateUrl: "views/account/administrators/views/bank.html"
                }
            }

        })
});