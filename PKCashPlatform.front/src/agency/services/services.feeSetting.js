angular.module("services.feeSetting", [])
    .service("feeSettingService", feeSettingService);

feeSettingService.$inject = ['APP_CONFIG', 'httpSvc'];

function feeSettingService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        del: del,
        modify: modify,
        getInfo: getInfo
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.POUNDAGE_LISTSET,postData)
            .then(getDataComplete)
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
        return httpSvc.del(APP_CONFIG.apiUrls.POUNDAGE_DEL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取详情
    function getInfo(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.POUNDAGE_GETLIST,postData)
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
        return httpSvc.put(APP_CONFIG.apiUrls.POUNDAGE_UPDATE,postData)
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