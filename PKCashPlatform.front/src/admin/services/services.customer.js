angular.module("services.customer", [])
    .service("customerService", customerService);

customerService.$inject = ['APP_CONFIG','httpSvc'];

function customerService(APP_CONFIG,httpSvc) {

    return {
        getSite:getSite,
        getProxyList:getProxyList,
        getChildernSite:getChildernSite,
        getUserList:getUserList,
        setMemberStatus:setMemberStatus,
        setuserInforment:setuserInforment,
        getLoginLog:getLoginLog,
        getSiteLog:getSiteLog,
        getAutoAudit:getAutoAudit,
        getHierarchicalList:getHierarchicalList,
        getChilderList:getChilderList,
        getBankOut:getBankOut,
        getBankIn:getBankIn,
        getChilderStatus:getChilderStatus,                                   // 子账号管理--状态修改	PUT/child/status
        getChilderPut:getChilderPut,                                              // 子账号管理--修改提交	PUT/child
        getAuditLog:getAuditLog,                                         // 稽核日志--稽核日志	GET/audit/log
        getAuditRecord:getAuditRecord,                                    // 稽核日志--稽核列表	GET/audit/record
        getNoticeKey:getNoticeKey                                // 公告密钥管理	GET/customer/noticeKey

    };
    // 公告密钥管理	GET/customer/noticeKey
    function getNoticeKey(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.NOTICE_KEY,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 稽核日志--稽核日志	GET/audit/log
    function getAuditRecord(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.AUDIT_RECORD,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 稽核日志--稽核日志	GET/audit/log
    function getAuditLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.AUDIT_LOG,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 子账号管理--状态修改	PUT/child/status
    function getChilderStatus(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CHILD_STATUS,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    // 子账号管理--修改提交	PUT/child
    function getChilderPut(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.CHILD_PUT,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    // 出入款管理--出款管理	GET/bankin_out/out
    function getBankOut(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BANKIN_OUT,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 出入款管理--入款管理	GET/bankin_out/in
    function getBankIn(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BANKIN_IN,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
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

    //获取子站点
    function getChildernSite() {
        return httpSvc.get(APP_CONFIG.apiUrls.CHILDERN_SITE, {
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //代理列表
    function getProxyList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PROXY_LIST,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //获取用户列表
    function getUserList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_LIST,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //用户列表修改状态
    function setMemberStatus(id, status) {
        return httpSvc.put(APP_CONFIG.apiUrls.USER_STATUS, {
            id: id
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //用户会员资料
    function setuserInforment(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_INFORMENT, {
            id: id
        }).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //日志管理--登录日志列表
    function getLoginLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.LOGIN_LOG,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //日志管理--操作日志列表
    function getSiteLog(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_LOG,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //日志管理--自动稽核列表
    function getAutoAudit(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.AUTOAUDIT,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //层次列表
    function getHierarchicalList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.HIERARCHICAL_LIST,postData )
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }

    //子账号列表
    function getChilderList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.CHILD_LIST,postData )
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