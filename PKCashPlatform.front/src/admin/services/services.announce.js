angular.module("services.announce", [])
    .service("announceService", announceService);

announceService.$inject = ['APP_CONFIG', 'httpSvc'];

function announceService(APP_CONFIG,httpSvc) {
    return {
        getDropSelect:getDropSelect,
        setSystemNoticeList:setSystemNoticeList,
        postAddNotice:postAddNotice,
        setAddType:setAddType,
        getNoticeNews:getNoticeNews,
        putNoticeNews:putNoticeNews,
        delSystemNotice:delSystemNotice
    };

    //获取站点下拉
    function getDropSelect(site_index_id) {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF,{
            site_index_id:site_index_id
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }


    //会员消息列表
    function  setSystemNoticeList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SYSTEM_MOTICE_LIST, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //获取类型下拉框
    function setAddType() {
        return httpSvc.get(APP_CONFIG.apiUrls.ADD_TYPE_DROP)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //添加消息
    function  postAddNotice(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ADD_SYSTEM_MOTICE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //获取修改消息
    function  getNoticeNews(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GET_NOTICE_NEWS, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //提交修改消息
    function  putNoticeNews(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.POST_NOTICE_NEWS, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //删除信息
    function  delSystemNotice(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DEL_SYSTEM_NOTICE, postData)
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