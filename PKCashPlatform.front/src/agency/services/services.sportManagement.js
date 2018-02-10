angular.module("services.sportManagement", [])
    .service("sportManagementService", sportManagementService);

sportManagementService.$inject = ['APP_CONFIG', 'httpSvc'];

function sportManagementService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        modify: modify
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_SPORTS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改顺序
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_SPORTS,postData)
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