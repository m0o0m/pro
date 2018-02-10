/**
 * Created by apple on 17/12/13.
 */
angular.module("services.reportform", [])
    .service("reportformService", reportformService);
reportformService.$inject = ['APP_CONFIG', 'httpSvc'];
function reportformService(APP_CONFIG,httpSvc) {
    return{
        getLevel:getLevel,
        getDropSelect:getDropSelect,
        reportquery:reportquery,
        hareholdersStatement:hareholdersStatement,
        generalGenerationReport:generalGenerationReport,
        proxyReport:proxyReport,
        membershipReport:membershipReport
    };

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
    //站点
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
    //报表查询
    function reportquery(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.REPORTQUERY,postData).then(getDataComplete)
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
}
