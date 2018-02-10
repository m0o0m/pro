angular.module("services.financeReport", [])
    .service("financeReportService", financeReportService);

financeReportService.$inject = ['APP_CONFIG', 'httpSvc'];

function financeReportService(APP_CONFIG,httpSvc) {
    return {
        getModule: getModule,
        getGame: getGame,
        reportquery:reportquery,
        hareholdersStatement:hareholdersStatement,
        generalGenerationReport:generalGenerationReport,
        proxyReport:proxyReport,
        getReport:getReport,
        getBill:getBill,
        modifyBill:modifyBill,
        delBill:delBill,
        issuedBill:issuedBill,
        membershipReport:membershipReport
    };

    //获取模块
    function getModule(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_REPORT_MODULE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取游戏
    function getGame(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FINANCE_REPORT_GAME,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //报表查询
    function reportquery(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REPORTQUERY,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取股东报表
    function hareholdersStatement(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.SHAREHOLDERS_STATEMENT,{
            id:id
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
    //获取总代
    function generalGenerationReport(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.GENERAL_GENERATION_REPORT,{
            id:id
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
    //获取代理
    function proxyReport(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.PROXY_REPORT,{
            id:id
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
    //获取会员
    function membershipReport(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBERSHIP_REPORT,{
            id:id
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

    //报表统计
    function getReport(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REPORT,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //账单查询
    function getBill(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BILLQUERY,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //账单查询--修改
    function modifyBill(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.BILLQUERY,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //账单查询--删除
    function delBill(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.BILLQUERY,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //账单查询--下发账单
    function issuedBill(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.BILLQUERY_ISSUED,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
}