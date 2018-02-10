angular.module('app.dashboard', [
    'ui.router',
    'ngResource',
    "FileManagerApp"
])
.config( function (fileManagerConfigProvider,$stateProvider) {
    var defaults = fileManagerConfigProvider.$get();
    var manageUrl = window.appConfig.apiUrls.HOST + window.appConfig.apiUrls.FILE_MANAGER;
    fileManagerConfigProvider.set({
        appName: '文件管理系统',
        listUrl: manageUrl+"/list",
        uploadUrl: manageUrl,
        renameUrl: manageUrl,
        copyUrl: manageUrl,
        moveUrl: manageUrl,
        removeUrl: manageUrl,
        editUrl: manageUrl,
        getContentUrl: manageUrl,
        createFolderUrl: manageUrl,
        downloadFileUrl: manageUrl,
        downloadMultipleUrl: manageUrl,
        compressUrl: manageUrl,
        extractUrl: manageUrl,
        permissionsUrl: manageUrl,
        basePath: manageUrl,

        allowedActions: angular.extend(defaults.allowedActions, {
          edit: false,
          compress: false,
          compressChooseName: false,
          extract: false,
          download: false,
          preview: false,
          upload: false
        })
      });
    $stateProvider
        .state('app.dashboard', {
            url: '/dashboard',
            views: {
                "content@app": {
                    controller: 'DashboardCtrl',
                    templateUrl: 'views/dashboard/dashboard.html'
                }
            },
            data:{
                title: 'Dashboard'
            }
        })
        .state('app.filemanager', {
            url: '/filemanager',
            views: {
                "content@app": {
                    controller: 'DashboardCtrl',
                    templateUrl: 'views/dashboard/filemanager.html'
                }
            },
            data:{
                title: 'filemanager'
            }
        })
        .state('app.dashboard-social', {
            url: '/dashboard-social',
            views: {
                "content@app": {
                    templateUrl: 'views/dashboard/social-wall.html'
                }
            },
            data:{
                title: 'Dashboard Social'
            }
        });
});
