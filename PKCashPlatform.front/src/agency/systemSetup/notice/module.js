angular.module('app.notice', [ 'ui.router']);

angular.module('app.notice').config(function ($stateProvider) {

    $stateProvider
        .state('app.notice', {
            abstract: true,
            data: {
                // title: '上级公告'
                title: 'AnnouncementInformation'
            }
        })
        .state('app.notice.notice', {
            url: '/notice/notice',
            data: {
                title: 'AnnouncementInformation'
            },
            views: {
                "content@app": {
                    controller: 'noticeCtrl',
                    templateUrl: "views/systemSetup/notice/views/notice.html"
                }
            }
        })
        .state('app.notice.systemLog', {
            url: '/notice/systemLog',
            data: {
                title: 'SystemAnnouncement'
            },
            views: {
                "content@app": {
                    controller: 'systemLogCtrl',
                    templateUrl: "views/systemSetup/notice/views/systemLog.html"
            }
        }
    })
});