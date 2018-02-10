angular.module("services.financeSummary", [])
    .service("financeSummaryService", financeSummaryService);

financeSummaryService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeSummaryService(APP_CONFIG,httpSvc) {
    return {
        getData: getData
    };

    //获取数据
    function getData(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_SUMMARY,postData)
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