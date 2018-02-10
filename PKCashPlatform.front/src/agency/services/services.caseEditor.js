angular.module("services.caseEditor", [])
    .service("CaseEditorService", CaseEditorService);

CaseEditorService.$inject = ['APP_CONFIG', 'httpSvc'];

//网站资讯管理--案件编辑
function CaseEditorService(APP_CONFIG,httpSvc) {
    return {
        getPendingCase: getPendingCase,
        getPendingCaseSend: getPendingCaseSend,
        getPendingCaseDel: getPendingCaseDel,
        getAuditCase:getAuditCase,
        getAuditCaseDel:getAuditCaseDel,
        getThroughCase:getThroughCase,
        getRevokeCase: getRevokeCase
    };
    //待审案件
    function getPendingCase () {
        return httpSvc.get(APP_CONFIG.apiUrls.PENDING_CASE).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //待审案件--发送审核
    function getPendingCaseSend (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PENDING_CASE_SEND,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //待审案件--删除
    function getPendingCaseDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.PENDING_CASE_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //待审中案件
    function getAuditCase () {
        return httpSvc.get(APP_CONFIG.apiUrls.AUDIT_CASE).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //待审中案件--删除
    function getAuditCaseDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.AUDIT_CASE_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }

    //通过案件
    function getThroughCase () {
        return httpSvc.get(APP_CONFIG.apiUrls.THROUGH_CASE).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    //撤销案件
    function getRevokeCase () {
        return httpSvc.get(APP_CONFIG.apiUrls.REVOKE_CASE).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
}