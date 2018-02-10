angular.module("services.preferencesSettings", [])
    .service("preferencesSettingsService", preferencesSettingsService);

preferencesSettingsService.$inject = ['APP_CONFIG', 'httpSvc'];

function preferencesSettingsService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite,
        getList: getList,
        modify: modify,
        add: add
    };
    //获取站点下拉
    function getSite() {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_APPLICATION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改状态
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_APPLICATION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //修改状态
    function add(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_APPLICATION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
}