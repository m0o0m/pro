angular.module("services.customerPreferentialQuery", [])
    .service("customerPreferentialQueryService", customerPreferentialQueryService);

customerPreferentialQueryService.$inject = ['APP_CONFIG', 'httpSvc'];

function customerPreferentialQueryService(APP_CONFIG,httpSvc) {
    return {
        getBullet: getBullet,
        getList: getList
    };

    //获取优惠查询
    function getBullet(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_PREFERENTIALQUERY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }


    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_PREFERENTIAL_LIST,postData)
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