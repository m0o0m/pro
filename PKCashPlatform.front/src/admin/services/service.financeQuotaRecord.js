angular.module("services.financeQuotaRecord", [])
    .service("financeQuotaRecordService", financeQuotaRecordService);

financeQuotaRecordService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeQuotaRecordService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getTransactionType: getTransactionType,
        getVideoType: getVideoType,
        getTransactionCategory: getTransactionCategory
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_QUOTARECORD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取交易类型下拉
    function getTransactionType(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.TRANSACTION_TYPE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取视讯类型下拉
    function getVideoType(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.VIDEO_TYPE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取交易类别下拉
    function getTransactionCategory(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.TRANSACTION_CATEGORY,postData)
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