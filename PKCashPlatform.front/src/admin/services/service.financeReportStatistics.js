angular.module("services.financeReportStatistics", [])
    .service("financeReportStatisticsService", financeReportStatisticsService);

financeReportStatisticsService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeReportStatisticsService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getWay: getWay,
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_REPORTSTATISTICS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取类型下拉
    function getWay(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_REPORTSTATISTICS_TYPE,postData)
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