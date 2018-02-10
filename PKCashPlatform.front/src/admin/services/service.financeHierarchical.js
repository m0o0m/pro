angular.module("services.financeHierarchical", [])
    .service("financeHierarchicalService", financeHierarchicalService);

financeHierarchicalService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeHierarchicalService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getHierarchi: getHierarchi,
        getDrop: getDrop,
        modify: modify,
        modifyHierarchi: modifyHierarchi,
        add: add,
        del: del
    };

    //获取层级设定列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_HIERARCHICAL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取层级管理列表
    function getHierarchi(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_HIERARCHICALMANAGER,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取层级下拉
    function getDrop(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_HIERARCHY_DROP,postData)
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
        return httpSvc.del(APP_CONFIG.apiUrls.FINANCE_HIERARCHICAL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改层级设定
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.FINANCE_HIERARCHICAL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改层级
    function modifyHierarchi(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.FINANCE_HIERARCHICALMANAGER,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //新增
    function add(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.FINANCE_HIERARCHICAL,postData)
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