angular.module("services.commission", [])
    .service("commissionService", commissionService);

commissionService.$inject = ['APP_CONFIG', 'httpSvc'];

function commissionService(APP_CONFIG,httpSvc) {
    return {
        getList: getList
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_SEARCH,postData)
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