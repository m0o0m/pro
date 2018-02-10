angular.module("services.lotterySort", [])
    .service("lotterySortService", lotterySortService);

lotterySortService.$inject = ['APP_CONFIG', 'httpSvc'];

function lotterySortService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite,
        getSource: getSource,
        getList: getList,
        modify: modify
    };

    //获取站点下拉
    function getSite() {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    };
    //获取来源下拉
    function getSource() {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_LOTTERY_HALL_SOURCE_DROP, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_LOTTERY_HALL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改顺序
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_LOTTERY_HALL,postData)
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