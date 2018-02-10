angular.module("services.lotteryManagrment", [])
    .service("lotteryManagrmentService", lotteryManagrmentService);

lotteryManagrmentService.$inject = ['APP_CONFIG', 'httpSvc'];

function lotteryManagrmentService(APP_CONFIG,httpSvc) {
    return {
        getSite: getSite,
        getList: getList,
        modifyOrder: modifyOrder,
        getEnclosure: getEnclosure,
        modifyEnclosure: modifyEnclosure,
        deleteEnclosure: deleteEnclosure,
        selectEnclosure: selectEnclosure
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
    }

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_LOTTERY,postData)
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
    function modifyOrder(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_LOTTERY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取附件信息
    function getEnclosure(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_LOTTERY_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改附件信息
    function modifyEnclosure(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_LOTTERY_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //删除附件信息
    function deleteEnclosure(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.SITE_LOTTERY_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //选取附件
    function selectEnclosure(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_LOTTERY_ENCLOSURE_SELECT,postData)
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