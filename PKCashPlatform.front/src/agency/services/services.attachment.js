angular.module("services.attachment", [])
    .service("attachmentService", attachmentService);

attachmentService.$inject = ['APP_CONFIG', 'httpSvc'];

function attachmentService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        modify: modify,
        del: del,
        upload: upload
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改
    function modify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.ENCLOSURE,postData)
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
    function del(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.ENCLOSURE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //上传附件
    function upload(postData) {
        return httpSvc.file(APP_CONFIG.apiUrls.ENCLOSURE,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

}