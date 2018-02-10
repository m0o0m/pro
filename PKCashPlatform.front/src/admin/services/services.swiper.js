angular.module("services.swiper", [])
    .service("swiperService", swiperService);

swiperService.$inject = ['APP_CONFIG', 'httpSvc'];

function swiperService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        storage: storage,
        disable: disable,
        getEnclosure: getEnclosure,
        modifyEnclosure: modifyEnclosure,
        deleteEnclosure: deleteEnclosure,
        selectEnclosure: selectEnclosure
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GRAPHIC_SWIPER,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //存储案件
    function storage(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_STORAGE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //启用禁用
    function disable(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_STATUS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取附件
    function getEnclosure(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改附件
    function modifyEnclosure(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //删除附件
    function deleteEnclosure(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //选择附件
    function selectEnclosure(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.GRAPHIC_SWIPER_ENCLOSURE_SELECT,postData)
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