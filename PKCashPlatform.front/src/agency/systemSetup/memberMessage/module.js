angular.module('app.memberMessage', [ 'ui.router']);

angular.module('app.memberMessage').config(function ($stateProvider) {

    $stateProvider
        .state('app.memberMessage', {
            abstract: true,
            data: {
                // title: '会员消息'
                title: 'MemberMessage'
            }
        })
        .state('app.memberMessage.memberMessage', {
            url: '/memberMessage/memberMessage',
            data: {
                title: 'MemberMessage'
            },
            views: {
                "content@app": {
                    controller: 'memberMessageCtrl',
                    templateUrl: "views/systemSetup/memberMessage/views/memberMessage.html"
                }
            }
        })
        .state('app.memberMessage.releaseNews', {
            url: '/memberMessage/releaseNews',
            data: {
                title: 'ReleaseNewMessages'
            },
            views: {
                "content@app": {
                    controller: 'releaseNewsCtrl',
                        templateUrl: "views/systemSetup/memberMessage/views/releaseNews.html"
                }
            }
        })
        .state('app.memberMessage.registrationPrivilege', {
            url: '/memberMessage/registrationPrivilege',
            data: {
                title: 'RegistrationCouponTemplates'
            },
            views: {
                "content@app": {
                    controller: 'registrationPrivilegeCtrl',
                    templateUrl: "views/systemSetup/memberMessage/views/registrationPrivilege.html"
                }
            }
        });
});