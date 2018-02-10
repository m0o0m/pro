angular.module("services.copyEditor", [])
    .service("CopyEditorService", CopyEditorService);

CopyEditorService.$inject = ['APP_CONFIG', 'httpSvc'];

//网站资讯--文案编辑
function CopyEditorService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,                   // 站点下拉框
        getHomeCopy: getHomeCopy,                       // 首页文案
        getHomeCopyKeep: getHomeCopyKeep,               // 首页文案--储存案件
        getHomeCopySub: getHomeCopySub,                 // 首页文案--编辑--提交
        getHomeEditor:getHomeEditor,                    // 首页文案--编辑内容
        getHomeEditorSub:getHomeEditorSub,              // 首页文案--编辑内容--提交
        getDiscountDel:getDiscountDel,                  // 优惠活动--删除
        getLineModifySub: getLineModifySub,             // 线路检测--编辑--提交
        getLineAddSub: getLineAddSub,                   // 线路检测--新增--提交
        getLineDetection: getLineDetection,             // 线路检测
        getWapDiscount_C_Del: getWapDiscount_C_Del,     // WAP优惠活动--优惠内容--删除
        getDepositCopy: getDepositCopy,                 // 存款文案
        getDepositCopySub: getDepositCopySub,           // 存款文案--编辑--提交
        getDepositEditor: getDepositEditor,             // 存款文案--编辑内容
        getDepositEditorSub: getDepositEditorSub,       // 存款文案--编辑内容--提交
        getDepositCopyKeep: getDepositCopyKeep,         // 存款文案--储存案件
        getDepositCopyModule: getDepositCopyModule,     // 存款文案--模板选择
        getDepositCopyModule_C: getDepositCopyModule_C, // 存款文案--模板选择--模板选择
        getDepositCopyModule_k: getDepositCopyModule_k, // 存款文案--模板选择--储存案件
        getRegisterCopy: getRegisterCopy,               // 注册文案
        getDiscountSub: getDiscountSub,                 // 优惠活动--编辑--提交
        getRegisterEditor: getRegisterEditor,           // 注册文案--编辑内容
        getRegisterEditorSub: getRegisterEditorSub,     // 注册文案--编辑内容--提交
        getRegisterCopyKeep: getRegisterCopyKeep,       // 注册文案--储存案件
        getRegisterCopyModule: getRegisterCopyModule,   // 注册文案--模板选择
        getRegisterCopy_M_C: getRegisterCopy_M_C,       // 注册文案--模板选择--模板选择
        getRegisterCopy_M_Keep: getRegisterCopy_M_Keep, // 注册文案--模板选择--储存案件
        getDiscount: getDiscount,                       // 优惠活动
        getRegisterCopySub: getRegisterCopySub,         // 注册文案--编辑--提交
        getDiscount_C: getDiscount_C,                   // 优惠活动--优惠内容
        getDiscount_M_C: getDiscount_M_C,               // 优惠活动--编辑内容
        getDiscount_M_Sub: getDiscount_M_Sub,           // 优惠活动--编辑内容--提交
        getDiscountUpdate: getDiscountUpdate,           // 优惠活动--上传--提交
        getWapDiscount_C_M_Sub: getWapDiscount_C_M_Sub, // WAP优惠活动--优惠内容--编辑内容--提交
        getDiscountKeep: getDiscountKeep,               // 优惠活动--储存案件
        getDiscountAddSub: getDiscountAddSub,           // 优惠活动--新增--提交
        getDiscountWidth: getDiscountWidth,             // 优惠活动--优惠宽度编辑
        getDiscountWidthSub: getDiscountWidthSub,       // 优惠活动--优惠宽度编辑--提交
        getDiscount_C_AddSub: getDiscount_C_AddSub,     // 优惠活动--优惠内容--添加优惠活动--提交
        getDiscount_C_W_Sub: getDiscount_C_W_Sub,       // 优惠活动--优惠内容--优惠宽度编辑--提交
        getDiscount_C_Update: getDiscount_C_Update,     // 优惠活动--优惠内容--上传--提交
        getDiscount_C_M_Content: getDiscount_C_M_Content,// 优惠活动--优惠内容--编辑内容
        getDiscount_C_M_Sub: getDiscount_C_M_Sub,       // 优惠活动--优惠内容--编辑内容--提交
        getDiscount_C_Del: getDiscount_C_Del,           // 优惠活动--优惠内容--删除
        getDiscount_C_Keep: getDiscount_C_Keep,         // 优惠活动--优惠内容--储存案件
        getWapDiscount_C_Keep:getWapDiscount_C_Keep,    // WAP优惠活动--优惠内容--储存案件
        getWapDiscount: getWapDiscount,                 // WAP优惠活动
        getWapDiscountAddSub: getWapDiscountAddSub,     // WAP优惠活动--新增--提交
        getWapDiscountSub: getWapDiscountSub,           // WAP优惠活动--编辑--提交
        getWapDiscount_C: getWapDiscount_C,             // WAP优惠活动--优惠内容
        getWapDiscount_M_C: getWapDiscount_M_C,         // WAP优惠活动--编辑内容
        getWapDiscount_M_Sub: getWapDiscount_M_Sub,     // WAP优惠活动--编辑内容--提交
        getWapDiscountUpdate: getWapDiscountUpdate,     // WAP优惠活动--上传--提交
        getWapDiscountDel: getWapDiscountDel,           // WAP优惠活动--删除
        getWapDiscount_C_AddSub: getWapDiscount_C_AddSub,// WAP优惠活动--优惠内容--添加优惠活动--提交
        getWapDiscount_C_Update: getWapDiscount_C_Update,// WAP优惠活动--优惠内容--上传--提交
        getWapDiscount_C_M_C: getWapDiscount_C_M_C,      // WAP优惠活动--优惠内容--编辑内容
        getDiscount_C_M_S:getDiscount_C_M_S,             // 优惠活动--优惠内容--编辑--提交
        getWapDiscount_C_M_S:getWapDiscount_C_M_S,       // WAP优惠活动--优惠内容--编辑--提交
        getLineDetectionDel:getLineDetectionDel,         // 线路检测--删除


    };
    // 站点下拉框
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
    // 优惠活动--优惠内容--编辑--提交
    function getDiscount_C_M_S (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_M_SUBMIT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--编辑--提交
    function getWapDiscount_C_M_S (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_M_SUBMIT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 首页文案
    function getHomeCopy (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.HOME_COPY,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 首页文案--储存案件
    function getHomeCopyKeep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.HOME_COPY_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 首页文案--编辑--提交
    function getHomeCopySub (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.HOME_COPY_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 线路检测--删除
    function getLineDetectionDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.LINE_DETECTION_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 首页文案--编辑内容
    function getHomeEditor (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.HOME_EDITOR,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 首页文案--编辑内容--提交
    function getHomeEditorSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.HOME_EDITOR_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--删除
    function getDiscountDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DISCOUNT_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 线路检测--编辑--提交
    function getLineModifySub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.LINE_DETECTION_MODIFY_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 线路检测--新增--提交
    function getLineAddSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.LINE_DETECTION_ADD_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 线路检测
    function getLineDetection (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.LINE_DETECTION,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--删除
    function getWapDiscount_C_Del (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.WAP_DISCOUNT_CONTENT_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案
    function getDepositCopy (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DEPOSIT_COPY,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--编辑--提交
    function getDepositCopySub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DEPOSIT_COPY_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--编辑内容
    function getDepositEditor (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DEPOSIT_EDITOR,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--编辑内容--提交
    function getDepositEditorSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DEPOSIT_EDITOR_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--储存案件
    function getDepositCopyKeep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DEPOSIT_COPY_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--模板选择
    function getDepositCopyModule (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DEPOSIT_COPY_MODULE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--模板选择--模板选择
    function getDepositCopyModule_C (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DEPOSIT_COPY_MODULE_CHOICE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 存款文案--模板选择--储存案件
    function getDepositCopyModule_k (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DEPOSIT_COPY_MODULE_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案
    function getRegisterCopy (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REGISTER_COPY,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--编辑--提交
    function getDiscountSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--编辑内容
    function getRegisterEditor (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REGISTER_EDITOR,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--编辑内容--提交
    function getRegisterEditorSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REGISTER_EDITOR_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--储存案件
    function getRegisterCopyKeep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REGISTER_COPY_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--模板选择
    function getRegisterCopyModule (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.REGISTER_COPY_MODULE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--模板选择--模板选择
    function getRegisterCopy_M_C (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REGISTER_COPY_MODULE_CHOICE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--模板选择--储存案件
    function getRegisterCopy_M_Keep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REGISTER_COPY_MODULE_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动
    function getDiscount (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DISCOUNT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 注册文案--编辑--提交
    function getRegisterCopySub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.REGISTER_COPY_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容
    function getDiscount_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DISCOUNT_CONTENT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--编辑内容
    function getDiscount_M_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DISCOUNT_MODIFY_CONTENT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--编辑内容--提交
    function getDiscount_M_Sub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_MODIFY_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--上传--提交
    function getDiscountUpdate (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_UPDATE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--编辑内容--提交
    function getWapDiscount_C_M_Sub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_M_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--储存案件
    function getDiscountKeep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--新增--提交
    function getDiscountAddSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_ADD_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠宽度编辑
    function getDiscountWidth () {
        return httpSvc.get(APP_CONFIG.apiUrls.DISCOUNT_WIDTH).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠宽度编辑--提交
    function getDiscountWidthSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_WIDTH_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--添加优惠活动--提交
    function getDiscount_C_AddSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_ADD_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--优惠宽度编辑--提交
    function getDiscount_C_W_Sub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_WIDTH_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--上传--提交
    function getDiscount_C_Update (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_UPDATE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--编辑内容
    function getDiscount_C_M_Content (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.DISCOUNT_C_M_CONTENT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--编辑内容--提交
    function getDiscount_C_M_Sub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_M_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--删除
    function getDiscount_C_Del (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DISCOUNT_C_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 优惠活动--优惠内容--储存案件
    function getDiscount_C_Keep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DISCOUNT_C_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--储存案件
    function getWapDiscount_C_Keep (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_KEEP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动
    function getWapDiscount (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.WAP_DISCOUNT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--新增--提交
    function getWapDiscountAddSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_ADD_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--编辑--提交
    function getWapDiscountSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容
    function getWapDiscount_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.WAP_DISCOUNT_C,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--编辑内容
    function getWapDiscount_M_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.WAP_DISCOUNT_M_C,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--编辑内容--提交
    function getWapDiscount_M_Sub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_M_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--上传--提交
    function getWapDiscountUpdate (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_UPDATE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--删除
    function getWapDiscountDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.WAP_DISCOUNT_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--添加优惠活动--提交
    function getWapDiscount_C_AddSub (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_ADD_SUB,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--上传--提交
    function getWapDiscount_C_Update (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_UPDATE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // WAP优惠活动--优惠内容--编辑内容
    function getWapDiscount_C_M_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.WAP_DISCOUNT_C_M_C,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }

}