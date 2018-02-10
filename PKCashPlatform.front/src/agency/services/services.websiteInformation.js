angular.module("services.websiteInformation", [])
    .service("websiteInformationService", websiteInformationService);

websiteInformationService.$inject = ['APP_CONFIG', 'httpSvc'];

function websiteInformationService(APP_CONFIG,httpSvc) {
    return {
        getColorSelect: getColorSelect,
        getInformation: getInformation,
        modifyInformation: modifyInformation
    };
    //获取颜色下拉
    function getColorSelect() {
        return httpSvc.get(APP_CONFIG.apiUrls.COLOR_DROP, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取网站信息
    function getInformation(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_WEBSITE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改网站信息
    function modifyInformation(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_WEBSITE,postData)
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