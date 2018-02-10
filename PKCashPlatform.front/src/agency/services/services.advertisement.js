angular.module("services.advertisement", [])
    .service("advertisementService", advertisementService);

advertisementService.$inject = ['APP_CONFIG', 'httpSvc'];

function advertisementService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        sort: sort,
        disable: disable,
        getDetail: getDetail,
        modify: modify
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GRAPHIC_ADVERTISEMENT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //排序
    function sort(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.GRAPHIC_ADVERTISEMENT_SORT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //启用禁用
    function disable(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_ADVERTISEMENT_STATUS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取广告详情
    function getDetail(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GRAPHIC_ADVERTISEMENT_DETAIL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改LOGO
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_ADVERTISEMENT,postData)
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