/**
 * Created by apple on 17/12/12.
 */
angular.module("services.paymentSetting", [])
    .service("paymentSettingService", paymentSettingService);

paymentSettingService.$inject = ['APP_CONFIG', 'httpSvc'];
function paymentSettingService(APP_CONFIG,httpSvc) {
    return{
        getDropSelect:getDropSelect,
        bankOutbank:bankOutbank,
        outbankStatus:outbankStatus,
        paymentList:paymentList,
        getLevel:getLevel,
        addPayment:addPayment,
        paymentDetail:paymentDetail,
        paymentPut:paymentPut,
        paymentDeposit:paymentDeposit,
        paymentStatus:paymentStatus,
        paymentDelete:paymentDelete,
        bankIncome:bankIncome,
        bankIncomeStatus:bankIncomeStatus,
        thirdPaidList:thirdPaidList,
        thirdBank:thirdBank,
        thirdStatus:thirdStatus,
        onlineSetup:onlineSetup,
        newOnlineSetup:newOnlineSetup,
        paidType:paidType,
        onlineSetupSingle:onlineSetupSingle,
        newOnlineSetupModify:newOnlineSetupModify,
        stopOnlineSstup:stopOnlineSstup,
        onlineSetupDel:onlineSetupDel,
        paysetList:paysetList,
        currency:currency,
        paysetAdd:paysetAdd,
        paysetPublic:paysetPublic,
        paysetDetail:paysetDetail,
        paysetModify:paysetModify,
        paysetDel:paysetDel,
        denomination:denomination,
        paysetOne:paysetOne,
        paysets:paysets,
        paysetPublicOne:paysetPublicOne
    };

    function getDropSelect() {

        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
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
    //出款银行剔除列表
    function bankOutbank(postData) {

        return httpSvc.get(APP_CONFIG.apiUrls.BANK_OUTBANK,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改出款银行剔除状态
    function outbankStatus(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.OUTBANK_STATUS,{
            id:id
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
    //入款银行设定列表
    function paymentList(postData) {

        return httpSvc.get(APP_CONFIG.apiUrls.PAYMENT_LIST,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
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
    //添加入款银行设定
    function addPayment(formData) {
        return httpSvc.form(APP_CONFIG.apiUrls.ADD_PAYMENT,formData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取单个银行入款设定
    function paymentDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYMENT,{
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
    //修改单个入款银行设定
    function paymentPut(formData) {
        return httpSvc.forms(APP_CONFIG.apiUrls.PAYMENT_PUT,formData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //存款记录
    function paymentDeposit(postData) {

        return httpSvc.get(APP_CONFIG.apiUrls.PAYMENT_DEPOSIT,postData).then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            console.log(response)
            return response.data;
        }

        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    //修改单个入款银行设定状态
    function paymentStatus(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.PAYMENT_STATUS,{
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
    //删除单个入款银行设定
    function paymentDelete(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.PAYMENT_DELETE,{
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
    //入款银行剔除列表
    function bankIncome(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BANK_INCOME,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //入款银行剔除状态
    function bankIncomeStatus(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.BANK_INCOME_STATUS,{
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
    //第三方下拉框
    function thirdPaidList() {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_PAID_LIST)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //第三方银行剔除列表

    function thirdBank(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.BANK_THIRD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //第三方银行状态
    function thirdStatus(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.THIRD_STATUS,{
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
    //线上支付设定列表
    function onlineSetup(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ONLINE_SETUP,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //线上支付设定（新增）
    function newOnlineSetup(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.NEWONLINE_SETUP,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取类型
    function paidType() {
        return httpSvc.get(APP_CONFIG.apiUrls.PAID_TYPE)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取线上支付设定详情
    function onlineSetupSingle(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.ONLINE_SETUP_SINGLE,{
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
    //修改线上支付设定
    function newOnlineSetupModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.NEWONLINE_STUP_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //修改线上支付状态
    function stopOnlineSstup(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.STOP_ONLINESETUP,{
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
    //删除线上支付设定
    function onlineSetupDel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.ONLINE_SETUP_DEL,{
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
    //获取支付参数设定(公司自定设置)
    function paysetList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYSET_LIST,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取币别
    function currency() {
        return httpSvc.get(APP_CONFIG.apiUrls.CURRENCY)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //支付参数设定(添加公司设定)
    function paysetAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PAYSET_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //支付参数设定(币别设定)
    function paysetPublic(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYSET_PUBLIC,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取单个公司设定
    function paysetDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYSET_DETAIL,{
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
    function paysetModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PAYSET_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //删除公司设定
    function paysetDel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.PAYSET_DELETE,{
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
    //删除币别
    function denomination(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.DENOMINATION,{
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
    //获取公司设定详情
    function paysetOne(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYSET_ONE,{
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
    //修改公司设定
    function paysets(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PAYSETES,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response.data;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //查看币别
    function paysetPublicOne(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.PAYSET_PUBLIC_ONE,{
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