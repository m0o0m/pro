angular.module("services.customerApplicationInquiry", [])
    .service("customerApplicationInquiryService", customerApplicationInquiryService);

customerApplicationInquiryService.$inject = ['APP_CONFIG', 'httpSvc'];

function customerApplicationInquiryService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getSwitch: getSwitch,
        modify: modify
    };
    //获取自助优惠申请列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_APPLICATIONINQUIRY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取自助优惠开关列表
    function getSwitch(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_APPLICATIONSWITCH,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //编辑
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CUSTOMER_APPLICATIONSWITCH,postData)
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