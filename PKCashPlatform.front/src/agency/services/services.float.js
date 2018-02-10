angular.module("services.float", [])
    .service("floatService", floatService);

floatService.$inject = ['APP_CONFIG', 'httpSvc'];

function floatService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        modify: modify
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GRAPHIC_FLOAT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_FLOAT,postData)
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