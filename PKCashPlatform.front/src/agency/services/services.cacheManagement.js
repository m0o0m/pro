angular.module("services.cacheManagement", [])
    .service("cacheManagementService", cacheManagementService);

cacheManagementService.$inject = ['APP_CONFIG', 'httpSvc'];

function cacheManagementService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite,
        getPage: getPage,
        submit: submit
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
    };
    //获取界面下拉
    function getPage() {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_CACHE_PAGE_DROP, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //生成缓存
    function submit(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_CACHE,postData)
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