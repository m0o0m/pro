angular.module("services.videoManagement", [])
    .service("videoManagementService", videoManagementService);

videoManagementService.$inject = ['APP_CONFIG', 'httpSvc'];

function videoManagementService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite,
        getType: getType,
        getStyle: getStyle,
        getList: getList,
        use: use,
        back: back,
        modify: modify
    };
    //获取站点下拉
    function getSite() {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取类型下拉
    function getType() {
        return httpSvc.get(APP_CONFIG.apiUrls.VIDEO_TYPE_DROP, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取风格下拉
    function getStyle() {
        return httpSvc.get(APP_CONFIG.apiUrls.VIDEO_STYLE_DROP, {
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
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_VIDEO,postData)
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
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_VIDEO,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //使用
    function use(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_VIDEO_USE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //还原老版本
    function back(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_VIDEO_BACK,postData)
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