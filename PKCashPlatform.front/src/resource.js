/**
 * Created by gbh on 2017/09/19.
 */
"use strict";
angular.module('app.resource', []);
angular.module('app.resource')
.provider('resourceSvc', [
    function () {
        var resourceObj = {
            setItem: function (resource, key, value) {
                if (angular.isObject(value)) {
                    value = JSON.stringify(value);
                } else if (angular.isString(value) || angular.isNumber(value)) {
                    value = value.toString();
                }

                if (!angular.isString(key)) {
                    return false;
                }

                if (!angular.isString(value)) {
                    return false;
                }

                if (window[resource]) {
                    window[resource].setItem(key, value);
                }
            },
            getItem: function (resource, key, defaultValue) {
                if (!angular.isString(key)) {
                    return false;
                }

                if (window[resource]) {
                    return window[resource].getItem(key) || defaultValue || null;
                } else {
                    return defaultValue;
                }
            },
            removeItem: function (resource, key) {
                if (!angular.isString(key)) {
                    return false;
                }

                if (window[resource]) {
                    window[resource].removeItem(key);
                }
            },
            clear: function (resource) {
                if (window[resource]) {
                    window[resource].clear();
                }
            }
        };

        var that = this;

        this.setLocal = function (key, value) {
            resourceObj.setItem('localStorage', key, value);
        };
        this.getLocal = function (key, defaultValue) {
            return resourceObj.getItem('localStorage', key, defaultValue);
        };
        this.getLocalObj = function (key) {
            return JSON.parse(resourceObj.getItem('localStorage', key, "{}"));
        };
        this.removeLocal = function (key) {
            resourceObj.removeItem('localStorage', key);
        };
        this.clearLocal = function (args) {
            if(args && angular.isArray(args)){
                var localStorage = window.localStorage;
                var reg = new RegExp("^(" + args.join("|") + ")$");
                for (var key in localStorage) {
                    if (!reg.test(key)) {
                        that.removeLocal(key);
                    }
                }
            } else {
                resourceObj.clear('localStorage');
            }
        };
        this.setSession = function (key, value) {
            resourceObj.setItem('sessionStorage', key, value);
        };
        this.getSession = function (key, defaultValue) {
            return resourceObj.getItem('sessionStorage', key, defaultValue);
        };
        this.getSessionObj = function (key) {
            return JSON.parse(this.getSession(key, "{}"));
        };
        this.removeSession = function (key) {
            resourceObj.removeItem('sessionStorage', key);
        };
        this.clearSession = function (args) {
            if(args && angular.isArray(args)){
                var sessionStorage = window.sessionStorage;
                var reg = new RegExp("^(" + args.join("|") + ")$");
                resourceObj.clear('sessionStorage');
                for (var key in sessionStorage) {
                    if (!reg.test(key)) {
                        that.removeSession(key);
                    }
                }
            } else {
                resourceObj.clear('sessionStorage');
            }
        };
        this.localTemp = function (key) {
            return {
                query: function (valueKey, defaultVal) {
                    if (valueKey) {
                        var obj = that.getLocalObj(key);
                        return obj[valueKey] || defaultVal;
                    }
                    return angular.copy(that.getLocalObj(key));
                },
                save: function (value) {
                    var obj = that.getLocalObj(key);
                    angular.extend(obj, value);
                    that.setLocal(key, obj)
                },
                clear: function (valueKey) {
                    if (valueKey) {
                        var obj = that.getLocalObj(key);
                        delete obj[valueKey];
                        that.setLocal(key, obj)
                    } else {
                        that.removeLocal(key);
                    }
                }
            }
        };
        this.sessionTemp = function (key) {
            return {
                query: function (valueKey, defaultVal) {
                    if (valueKey) {
                        var obj = that.getSessionObj(key);
                        return obj[valueKey] || defaultVal;
                    }
                    return angular.copy(that.getSessionObj(key));
                },
                save: function (value) {
                    var obj = that.getSessionObj(key);
                    angular.extend(obj, value);
                    that.setSession(key, obj)
                },
                clear: function (valueKey) {
                    if (valueKey) {
                        var obj = that.getSessionObj(key);
                        delete obj[valueKey];
                        that.setSession(key, obj)
                    } else {
                        that.removeSession(key);
                    }
                }
            }
        };
        this.$get = function () {
            return this;
        };
    }]);



