angular.module("services.financeArrears", [])
    .service("financeArrearsService", financeArrearsService);

financeArrearsService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeArrearsService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        add: add,
        modify: modify
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_ARREARS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //新增催款
    function add(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.FINANCE_ARREARS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //催款
    function modify(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.FINANCE_ARREARS_PRESS,postData)
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