/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.analysisExit", [])
    .service("AnalysisExitService", AnalysisExitService);

AnalysisExitService.$inject = ['APP_CONFIG', 'httpSvc'];

function AnalysisExitService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,
        getTypeSelect: getTypeSelect,
        getAgencySelect: getAgencySelect,
        getPurchaseAnalysis: getPurchaseAnalysis,
        getAnalysisExit: getAnalysisExit,
        getPreferentialAnalysis: getPreferentialAnalysis,
        getValidList: getValidList

    };

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
    function getTypeSelect () {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_BUY_TYPE_SELECT).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
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
    function getPreferentialAnalysis(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_RETURN_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function getPurchaseAnalysis(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_BUY_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function getAnalysisExit(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_INCOME_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function getValidList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
}