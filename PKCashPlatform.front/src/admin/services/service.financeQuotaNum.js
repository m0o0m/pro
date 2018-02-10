angular.module("services.financeQuotaNum", [])
    .service("financeQuotaNumService", financeQuotaNumService);

financeQuotaNumService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeQuotaNumService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_QUOTANUM,postData)
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