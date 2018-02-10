angular.module("services.common", [])
    .service("commonService", commonService);

commonService.$inject = ['APP_CONFIG', 'httpSvc'];

function commonService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite
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
}