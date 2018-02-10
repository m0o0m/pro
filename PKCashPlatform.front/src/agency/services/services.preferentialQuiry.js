angular.module("services.preferentialQuiry", [])
    .service("preferentialQuiryService", preferentialQuiryService);

preferentialQuiryService.$inject = ['APP_CONFIG', 'httpSvc'];

function preferentialQuiryService(APP_CONFIG,httpSvc) {
    return {
        getList: getList,
        getDetail: getDetail,
        getLevel: getLevel,
        del: del,
        retreatWaterEdit:retreatWaterEdit
    };

    //获取列表
    function getList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PREFERENTIAL_INQUIRIES,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //获取明细
    function getDetail(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.RETREAT_WATER_DETAIL,postData)
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
    //优惠查询(明细-冲销)
    function retreatWaterEdit(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.RETREAT_WATER_EDIT,{
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

}