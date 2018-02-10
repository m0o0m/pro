'use strict';

angular.module('app.dashboard', [
    'ui.router',
    'ngResource'
])

.config(function ($stateProvider) {
    $stateProvider
        .state('app.dashboard', {
            url: '/dashboard',
            views: {
                "content@app": {
                    controller: 'DashboardCtrl',
                    templateUrl: 'app/views/dashboard/dashboard.html'
                }
            },
            data:{
                title: 'Dashboard'
            }
        })
        .state('app.dashboard-social', {
            url: '/dashboard-social',
            views: {
                "content@app": {
                    templateUrl: 'app/views/dashboard/social-wall.html'
                }
            },
            data:{
                title: 'Dashboard Social'
            }
        });
});
