angular.module('app.auth', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider.state('realLogin', {
        url: '/real-login',
        views: {
            root: {
                templateUrl: "views/auth/login/login.html",
                controller: 'LoginCtrl'
            }
        }

    })

    .state('modifyPwd', {
        url: '/modifyPwd',
        views: {
            root: {
                templateUrl: 'app/views/auth/modifyPwd/modifyPwd.html',
                controller: 'modifyPwdCtrl'
            }
        },
        data: {
            title: 'Forgot Password',
            htmlId: 'extr-page'
        }
    })

    .state('login', {
        url: '/login',
        views: {
            root: {
                templateUrl: 'views/auth/views/login.html'
            }
        },
        data: {
            title: 'Login',
            htmlId: 'extr-page'
        },
        resolve: {
            srcipts: function(lazyScript){
                return lazyScript.register([
                    'vendor.ui.js'
                ])

            }
        }
    })

    .state('register', {
        url: '/register',
        views: {
            root: {
                templateUrl: 'app/views/auth/views/register.html'
            }
        },
        data: {
            title: 'Register',
            htmlId: 'extr-page'
        }
    })

    .state('forgotPassword', {
        url: '/forgot-password',
        views: {
            root: {
                templateUrl: 'app/views/auth/views/forgot-password.html'
            }
        },
        data: {
            title: 'Forgot Password',
            htmlId: 'extr-page'
        }
    })

    .state('lock', {
        url: '/lock',
        views: {
            root: {
                templateUrl: 'app/views/auth/views/lock.html'
            }
        },
        data: {
            title: 'Locked Screen',
            htmlId: 'lock-page'
        }
    })


}).constant('authKeys', {
    googleClientId: '',
    facebookAppId: ''
});
