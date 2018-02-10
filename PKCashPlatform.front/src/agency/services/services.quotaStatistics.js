angular.module("services.quotaStatistics", [])
    .service("QuotaStatisticsService", QuotaStatisticsService);

QuotaStatisticsService.$inject = ['APP_CONFIG', 'httpSvc'];

//资金管理--会员返佣接口
function QuotaStatisticsService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,
        getQuotaList: getQuotaList,
        getQuotaRecharge: getQuotaRecharge,
        getTypeList:getTypeList,
        getDropList:getDropList,
        getSubDropList:getSubDropList,
        getRecordOrdernum: getRecordOrdernum,
        getRecordCardBank: getRecordCardBank,
        getThreeBank: getThreeBank,
        getBankOrdernum: getBankOrdernum,
        getThreeSub: getThreeSub,
        getBankSub: getBankSub,
        getQuotaRecord: getQuotaRecord

    };

    //站点下拉框
    function getSiteSelect () {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //额度统计
    function getQuotaList (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.QUOTA_LIST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //充值记录
    function getQuotaRecharge (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.QUOTA_RECHARGE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //掉单列表
    function getDropList (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_SINGLE_RECORD,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //转入转出平台
    function getTypeList () {
        return httpSvc.get(APP_CONFIG.apiUrls.TYPE_LIST).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //掉单申请--提交
    function getSubDropList (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SUB_SINGLE_RECORD,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--第三方--订单号
    function getRecordOrdernum () {
        return httpSvc.get(APP_CONFIG.apiUrls.RECORD_ORDERNUM).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--银行卡--订单号
    function getBankOrdernum () {
        return httpSvc.get(APP_CONFIG.apiUrls.BANK_ORDERNUM).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--银行卡--收款银行
    function getRecordCardBank () {
        return httpSvc.get(APP_CONFIG.apiUrls.RECORD_CARD_BANK).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--第三方--支付银行
    function getThreeBank () {
        return httpSvc.get(APP_CONFIG.apiUrls.THREE_BANK).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--第三方--提交
    function getThreeSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.THREE_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度充值--银行卡--提交
    function getBankSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.BANK_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //额度记录
    function getQuotaRecord (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.QUOTA_RECORD,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }

}