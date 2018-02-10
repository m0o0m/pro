angular.module("services.floatManagrment", [])
    .service("floatManagrmentService", floatManagrmentService);

floatManagrmentService.$inject = ['APP_CONFIG', 'httpSvc'];

function floatManagrmentService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        deleteImg: deleteImg,
        modifyImg: modifyImg,
        addImg: addImg,
        getEnclosure: getEnclosure,
        modifyEnclosure: modifyEnclosure,
        deleteEnclosure: deleteEnclosure,
        selectEnclosure: selectEnclosure
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.FLOAT_IMG,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改图片信息
    function modifyImg(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GRAPHIC_LOGO,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //添加图片
    function addImg(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.GRAPHIC_LOGO,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //删除
    function deleteImg(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.FLOAT_IMG,postData)
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
        return httpSvc.get(APP_CONFIG.apiUrls.FLOAT_IMG_ENCLOSURE,postData)
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
        return httpSvc.put(APP_CONFIG.apiUrls.FLOAT_IMG_ENCLOSURE,postData)
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
        return httpSvc.del(APP_CONFIG.apiUrls.FLOAT_IMG_ENCLOSURE,postData)
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
        return httpSvc.post(APP_CONFIG.apiUrls.FLOAT_IMG_ENCLOSURE_SELECT,postData)
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