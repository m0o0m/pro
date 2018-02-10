angular.module("services.configurationsetting", [])
    .service("configurationsettingService", configurationsettingService);

configurationsettingService.$inject = ['APP_CONFIG', 'httpSvc'];

function configurationsettingService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getAd: getAd,
        mdifyAd: mdifyAd,
        modify: modify,
        del: del,
        getContent: getContent,
        editContent: editContent,
        add: add
    };
    //获取列表
    function getList() {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_CONFIGURATIONSETTING, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取弹框广告信息
    function getAd() {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_NOTICEAD_SETTING, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改弹框广告信息
    function mdifyAd(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CUSTOMER_NOTICEAD_SETTING, postData).then(getDataComplete)
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
        return httpSvc.put(APP_CONFIG.apiUrls.CUSTOMER_CONFIGURATIONSETTING, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //删除
    function del(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.CUSTOMER_CONFIGURATIONSETTING, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取广告内容
    function getContent(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_NOTICEAD_CONTENT, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //编辑广告内容
    function editContent(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CUSTOMER_NOTICEAD_CONTENT, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //新增广告
    function add(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.CUSTOMER_CONFIGURATIONSETTING, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
}