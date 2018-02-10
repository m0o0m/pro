angular.module("services.financeCash", [])
    .service("financeCashService", financeCashService);

financeCashService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeCashService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getWay: getWay,
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_CASH,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取方式下拉
    function getWay(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_CASH_TYPE,postData)
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