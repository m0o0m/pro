angular.module("services.customerExceptionMember", [])
    .service("customerExceptionMemberService", customerExceptionMemberService);

customerExceptionMemberService.$inject = ['APP_CONFIG', 'httpSvc'];

function customerExceptionMemberService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        deal: deal
    };

    //获取数据
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_EXCEPTIONMEMBER,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    function deal(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.CUSTOMER_EXCEPTIONMEMBER_HANDLE,postData)
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