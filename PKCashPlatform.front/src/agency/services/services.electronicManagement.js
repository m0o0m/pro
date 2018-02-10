angular.module("services.electronicManagement", [])
    .service("electronicManagementService", electronicManagementService);

electronicManagementService.$inject = ['APP_CONFIG', 'httpSvc'];

function electronicManagementService(APP_CONFIG,httpSvc) {
    return {
        getTheme: getTheme,
        getList: getList,
        initialization: initialization,
        modifyTheme: modifyTheme,
        modifyOrder: modifyOrder
    };
    //获取主题配置信息
    function getTheme() {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_ELECTRONICS_THEME, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_ELECTRONICS,postData)
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
    function modifyOrder(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_ELECTRONICS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改主题配置
    function modifyTheme(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_ELECTRONICS_THEME,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //初始化
    function initialization(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_ELECTRONICS_INITIALIZATION,postData)
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