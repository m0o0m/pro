/**
 * Created by mebar on 08/12/2017.
 */
angular.module("services.notice", [])
    .service("noticeService", noticeService);

noticeService.$inject = ['APP_CONFIG', 'httpSvc'];

function noticeService(APP_CONFIG,httpSvc) {
    return {
        setSystermNotice: setSystermNotice,
        setSystermInformation: setSystermInformation
    };

    function  setSystermNotice(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTERM_NOTICE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    function  setSystermInformation() {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTERM_INFORMATION)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
}