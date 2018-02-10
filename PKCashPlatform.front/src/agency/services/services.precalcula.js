angular.module("services.precalcula", [])
    .service("precalculaService", precalculaService);

precalculaService.$inject = ['APP_CONFIG', 'httpSvc'];

function precalculaService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getLevel: getLevel,
        del: del,
        getLevel:getLevel,
        statistics:statistics,
        depositDiscouny:depositDiscouny,
        retreatWaterSelfSearch:retreatWaterSelfSearch,
        retreatWaterDetail:retreatWaterDetail,
        retreatWaterEdit:retreatWaterEdit,
        comboPlatform:comboPlatform,
        retreatWaterSetAdd:retreatWaterSetAdd,
        retreatWaterSetDetail:retreatWaterSetDetail,
        RETREAT_WATER_SET_EDIT:RETREAT_WATER_SET_EDIT,
        overrideAddone:overrideAddone,
        overrideDetail:overrideDetail,
        overrideModify:overrideModify
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_WATER_SET_LIST,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取列表
    function getLevel(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETURN_LEVEL,postData)
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
        return httpSvc.del(APP_CONFIG.apiUrls.RETREAT_WATER_SET_DEL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取层级
    function getLevel() {
        return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_DROP)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //优惠统计列表

    function statistics(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.STATISTICS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //存入
    function depositDiscouny(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.DEPOSIT_DISCOUNY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //自助返水列表
    function retreatWaterSelfSearch(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_WATER_SELF_SEARCH,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //明细
    function retreatWaterDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_DETAIL,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //优惠查询(明细-冲销)
    function retreatWaterEdit(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_WATER_EDIT,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取商品
    function comboPlatform() {
        return httpSvc.get(APP_CONFIG.apiUrls.COMBO_PLATFORM,{
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //添加返点优惠设定
    function retreatWaterSetAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.RETREAT_WATER_SET_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取返点优惠详情
    function retreatWaterSetDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_WATER_SET_DETAIL,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //修改返点优惠设定
    function RETREAT_WATER_SET_EDIT(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.RETREAT_WATER_SET_EDIT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //添加代理退佣设定
    function overrideAddone(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.OVERRIDE_ADDONE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取代理退佣详情
    function overrideDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.OVERRIDE_DETAIL,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //代理退佣修改
    function overrideModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.OVERRIDE_MODIFY,postData)
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