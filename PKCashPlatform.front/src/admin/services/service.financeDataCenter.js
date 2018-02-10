angular.module("services.financeDataCenter", [])
    .service("financeDataCenterService", financeDataCenterService);

financeDataCenterService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeDataCenterService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getWay: getWay,
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_DATACENTER,postData)
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
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_DATACENTER_TYPE,postData)
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