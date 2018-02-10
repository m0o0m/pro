angular.module("services.customerVideo", [])
    .service("customerVideoService", customerVideoService);

customerVideoService.$inject = ['APP_CONFIG', 'httpSvc'];

function customerVideoService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getVideoType: getVideoType
    };

    //获取数据
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_VIDEO,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取视讯类别下拉
    function getVideoType(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_VIDEO_TYPE,postData)
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