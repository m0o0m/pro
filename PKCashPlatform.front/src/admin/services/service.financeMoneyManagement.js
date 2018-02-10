angular.module("services.financeMoneyManagement", [])
    .service("financeMoneyManagementService", financeMoneyManagementService);

financeMoneyManagementService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeMoneyManagementService(APP_CONFIG,httpSvc) {
    return {
        getThird: getThird,
        getBank: getBank,
        getHierarchy: getHierarchy,
        addThird: addThird,
        modifyThird: modifyThird,
        addBank: addBank,
        modifyBank: modifyBank
    };

    //获取第三方列表
    function getThird(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_THIRD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取银行卡列表
    function getBank(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_BANK,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取层级列表
    function getHierarchy(postData) {
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
    //添加第三方
    function addThird(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.FINANCE_THIRD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改第三方
    function modifyThird(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.FINANCE_THIRD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //添加银行
    function addBank(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.FINANCE_BANK,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改银行
    function modifyBank(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.FINANCE_BANK,postData)
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