angular.module("services.memberMessage", [])
    .service("memberMessageService", memberMessageService);

memberMessageService.$inject = ['APP_CONFIG', 'httpSvc'];

function memberMessageService(APP_CONFIG,httpSvc) {
    return {
        setMemberNews: setMemberNews,
        getDropSelect: getDropSelect,
        getTypeSelect: getTypeSelect,
        deleteNews:deleteNews,
        postPreferencesNews:postPreferencesNews,
        getSystemSelect:getSystemSelect,
        postMemberNews:postMemberNews
    };

    //会员消息列表
    function  setMemberNews(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_NEWS, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //获取站点下拉
    function getDropSelect(site_index_id) {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
            site_index_id: site_index_id
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取类型下拉
    function getTypeSelect() {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_STATUS, {
            site_index_id: site_index_id
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

    //删除
    function deleteNews(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DELETE_NEWS, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //会员获取体系下拉框
    function getSystemSelect() {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_SYSTEM).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //发布新消息
    function postMemberNews(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SYSTEM_MEMBER_NEWS, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }


    //注册优惠消息
    function postPreferencesNews(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SYSTEM_PREFERENCES_NEWS, postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response);
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
}