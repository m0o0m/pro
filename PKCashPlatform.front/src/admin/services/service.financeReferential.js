angular.module("services.financeReferential", [])
    .service("financeReferentialService", financeReferentialService);

financeReferentialService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeReferentialService(APP_CONFIG,httpSvc) {
    return {
        getList: getList
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_REFERENTIAL,postData)
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