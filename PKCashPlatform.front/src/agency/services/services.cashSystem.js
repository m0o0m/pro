/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.cashSystem", [])
    .service("CashSystemService", CashSystemService);

CashSystemService.$inject = ['APP_CONFIG', 'httpSvc'];

function CashSystemService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,
        getCashSystem: getCashSystem,
        delCashSystem: delCashSystem
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
    function getCashSystem(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_MONEY_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    function delCashSystem(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.USER_MONEY_DELETE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
};