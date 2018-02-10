angular.module('app.announcement', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.Announcement', {
            abstract: true,
            data: {
                title: 'Announcement'
            }
        })

        .state('app.Announcement.Announcement', {
            url: '/Announcement/Announcement',
            data: {
                title: 'Announcement'
            },
            views: {
                "content@app": {
                    controller: 'announceCtrl',
                    templateUrl: "views/announcement/views/Announcement.html"
                }
            }
        })
});
