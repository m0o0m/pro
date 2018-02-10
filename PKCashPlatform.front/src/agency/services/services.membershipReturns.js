angular.module("services.membershipReturns", [])
    .service("MembershipReturnsService", MembershipReturnsService);

AnalysisExitService.$inject = ['APP_CONFIG', 'httpSvc'];

//资金管理--会员返佣接口
function MembershipReturnsService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,
        getRebateList: getRebateList,
        getRebateDetail: getRebateDetail,
        getRebateWriteoff: getRebateWriteoff,
        getRebateSetGetAll: getRebateSetGetAll,
        getRebateSetAddAll: getRebateSetAddAll,
        getRebateSetSubmit: getRebateSetSubmit,
        getRebateSetGetOne: getRebateSetGetOne,
        getUserRebateSearch: getUserRebateSearch,
        getRebateSetDel: getRebateSetDel,
        getUserRebateDeposit: getUserRebateDeposit,
        getSpreadInfo: getSpreadInfo,
        getSpreadNumInfo: getSpreadNumInfo,
        getSpreadList: getSpreadList,
        getSpreadEdit: getSpreadEdit,
        getSpreadAdd: getSpreadAdd
    };


    //站点下拉框
    function getSiteSelect () {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣查询
    function getRebateList (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_LIST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣查询--明细
    function getRebateDetail (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_DETAILS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣查询--明细--冲销
    function getRebateWriteoff(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REBATE_WRITEOFF, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣优惠设定
    function getRebateSetGetAll(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_SET_GET_ALL, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣优惠设定-新增详情
    function getRebateSetAddAll(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_SET_ADD_ALL, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //返佣优惠设定-新增/修改-提交
    function getRebateSetSubmit(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REBATE_SET_SUBMIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣优惠设定-修改详情
    function getRebateSetGetOne(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REBATE_SET_GET_ONE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员返佣-搜索
    function getUserRebateSearch(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.USER_REBATE_SEARCH, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //返佣优惠设定--删除
    function getRebateSetDel(postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.REBATE_SET_DEL, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员返佣-搜索页面-存入
    function getUserRebateDeposit(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.USER_REBATE_DEPOSIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员推广查询
    function getSpreadInfo(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SPREAD_INFO, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员推广查询--推荐会员数
    function getSpreadNumInfo(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SPREAD_NUM_INFO, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //会员推广设定
    function getSpreadList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SPREAD_LIST, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员推广设定--修改
    function getSpreadEdit(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SPREAD_EDIT, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //会员推广设定--添加
    function getSpreadAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SPREAD_ADD, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
}