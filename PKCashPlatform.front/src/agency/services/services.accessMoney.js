/**
 * Created by apple on 17/12/9.
 */
/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.accessMoney", [])
    .service("AccessMoneyService", AccessMoneyService);

AccessMoneyService.$inject = ['APP_CONFIG', 'httpSvc'];

function AccessMoneyService(APP_CONFIG,httpSvc) {
    return {
        getDropSelect: getDropSelect,
        refuseMoney: refuseMoney,
        getMoney: getMoney,
        cancleMoney: cancleMoney,
        confirmMoney: confirmMoney,
        monitorDeposit:monitorDeposit,
        MonitorOnline:MonitorOnline,
        MonitorMoney:MonitorMoney,
        memberSearch:memberSearch,
        manualAccess:manualAccess,
        manualAccessBatch:manualAccessBatch,
        getLevel:getLevel,
        manualWithdrawal:manualWithdrawal,
        companyIncome:companyIncome,
        getAgencySelect:getAgencySelect,
        depositList:depositList,
        getAccountSelect:getAccountSelect,
        manualAccessRecord:manualAccessRecord,
        manualAccessCollect:manualAccessCollect,
        typeList:typeList,
        memberBlance:memberBlance,
        quotqSumbit:quotqSumbit,
        balanceConversion:balanceConversion,
        auditLog:auditLog,
        actualPurchase:actualPurchase,
        memberAuditnow:memberAuditnow
    };

    function getDropSelect() {

        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    function getAgencySelect () {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_BUY_AGENCY_SELECT).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    function getAccountSelect () {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRDPAID_LIST).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    function typeList () {
        return httpSvc.get(APP_CONFIG.apiUrls.TYPE_LIST).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    function memberBlance (account,for_type) {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_BALANCE,{
            account:account,
            for_type:for_type
        }).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取实际购买
    function actualPurchase (id) {
        return httpSvc.get(APP_CONFIG.apiUrls.ACTUAL_PURCHASE,{
            id:id
        }).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //稽核日志
    function memberAuditnow (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ACTUAL_MEMBER_AUDITNOW,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }


    function refuseMoney(id, agency_id,reason) {
        return httpSvc.put(APP_CONFIG.apiUrls.REFUSE_MONEY, {
            id: id,
            agency_id:agency_id,
            reason:reason
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    function getMoney(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.OUT_MONEY, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取层级
    function getLevel() {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_DROP)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    function cancleMoney(mid) {
        return httpSvc.put(APP_CONFIG.apiUrls.CANCLE_COMPANY, {
            id: mid
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function monitorDeposit(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MONITOR_DEPOSIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function MonitorOnline(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MONITOR_ONLINE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function MonitorMoney(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MONITOR_MOENY, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //人工存款-账号搜索
    function memberSearch(accountkey) {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMVER_BASEINFO,{
            account:accountkey
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //人工存款-存入
    function manualAccess(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MANUAL_ACCESS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    };
    //人工存款-批量存款
    function manualAccessBatch(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MANUAL_ACCESS_BATCH,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    };
    //人工取款-存入
    function manualWithdrawal(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MANUAL_WITHDRAWAL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //公司入款
    function companyIncome(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COMPANY_INCOME, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //线上公司入款
    function depositList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DEPOSIT_LIST, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //出入账目记录
    function manualAccessRecord(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MANUAL_ACCESS_RECORD, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function auditLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.AUDIT_AUDITLOG, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //出入账目记录
    function manualAccessCollect(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MANUAL_ACCESS_COLLECT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //额度转换记录列表
    function balanceConversion(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BALANCE_CONVERSION, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //额度转换提交
    function quotqSumbit(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.QUOTA_SUMBIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    function confirmMoney(id,agency_id) {
        return httpSvc.put(APP_CONFIG.apiUrls.CONFIRM_MOENY, {
            id:id,
            agency_id:agency_id

        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
}