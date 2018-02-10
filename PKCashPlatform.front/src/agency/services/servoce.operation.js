angular.module("services.operation", [])
    .service("operationService", operationService);

operationService.$inject = ['APP_CONFIG', 'httpSvc'];

function operationService(APP_CONFIG,httpSvc) {
    return {
        getDropSelect:getDropSelect,
        setSystemMemberLogin:setSystemMemberLogin,
        setSystemAdminLogin:setSystemAdminLogin,
        setSystemLog:setSystemLog,
        setAutomaticLog:setAutomaticLog
    };

    //获取站点下拉
    function getDropSelect(site_index_id) {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
            site_index_id: site_index_id
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }


    //会员消息列表
    function  setSystemMemberLogin(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTEM_MEMBER_LOGIN_LIST, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //管理员列表
    function  setSystemAdminLogin(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTEM_ADMIN_LOGIN, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //操作日志列表
    function  setSystemLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTEM_LOG_LIST, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //自动稽核列表
    function  setAutomaticLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTEM_AUDIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
}