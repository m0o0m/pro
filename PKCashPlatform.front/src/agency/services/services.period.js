angular.module("services.period", [])
    .service("periodService", periodService);

periodService.$inject = ['APP_CONFIG', 'httpSvc'];

function periodService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        del: del,
        commission: commission,
        modify: modify,
        add: add,
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PERIODS,postData)
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
        return httpSvc.del(APP_CONFIG.apiUrls.PERIODS_DEL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //退佣冲销
    function commission(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PERIODS_COMMISSION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //新增期数
    function add(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PERIODS_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改期数
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PERIODS_MODIFY,postData)
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