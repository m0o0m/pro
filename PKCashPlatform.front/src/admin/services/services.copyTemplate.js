angular.module("services.copyTemplate", [])
    .service("copyTemplateService", copyTemplateService);

copyTemplateService.$inject = ['APP_CONFIG','httpSvc'];

function copyTemplateService(APP_CONFIG,httpSvc) {

    return {
        copyTemplate:copyTemplate,
        copyTemplateStatus:copyTemplateStatus,
        videoCopyTemplate:videoCopyTemplate,
        putCpyTemplateStatus:putCpyTemplateStatus,
        postCpyTemplateStatus:postCpyTemplateStatus,
        addRegister:addRegister,
        getAddRegister:getAddRegister,
        putAddRegister:putAddRegister
    };

    //注册文案列表
    function  copyTemplate(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COPYTEMPLATE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //注册文案状态修改
    function  copyTemplateStatus(id, status) {
        return httpSvc.post(APP_CONFIG.apiUrls.COPY_STATUS, {
            id:id,
            status:status
        })
        .then(getDataComplete)
        .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //注册文案模板--添加
    function  addRegister(posData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ADD_REGISTER, posData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //注册文案模板--编辑获取
    function  getAddRegister(posData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GET_ADD_REGISTER, posData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //注册文案模板--编辑获取
    function putAddRegister(posData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PUT_ADD_REGISTER, posData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }


    //视讯文案列表
    function  videoCopyTemplate() {
        return httpSvc.get(APP_CONFIG.apiUrls.COPYTEMPLATE_VIDEO)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //视讯文案状态修改提交
    function  putCpyTemplateStatus(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PUT_VIDEOCOPY, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //视讯文案状态添加
    function  postCpyTemplateStatus(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.POST_VIDEOCOPY, postData)
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