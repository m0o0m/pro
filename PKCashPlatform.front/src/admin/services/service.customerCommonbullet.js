angular.module("services.customerCommonbullet", [])
    .service("customerCommonbulletService", customerCommonbulletService);

customerCommonbulletService.$inject = ['APP_CONFIG', 'httpSvc'];

function customerCommonbulletService(APP_CONFIG,httpSvc) {
    return {
        getBullet: getBullet,
        getAnimation: getAnimation,
        disable: disable,
    };

    //获取公告弹框
    function getBullet(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_COMMONBULLET,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取h5动画
    function getAnimation(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CUSTOMER_ANIMATION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //h5动画--启用停用
    function disable(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CUSTOMER_ANIMATION,postData)
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