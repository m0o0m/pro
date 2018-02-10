/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.balanceStatistics", [])
    .service("BalanceStatisticsService", BalanceStatisticsService);

BalanceStatisticsService.$inject = ['APP_CONFIG', 'httpSvc'];

function BalanceStatisticsService(APP_CONFIG,httpSvc) {
    return {
        getAgencySelect: getAgencySelect,
        getBalance: getBalance
    };
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
    function getBalance(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_BALANCE_SEARCH, postData)
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