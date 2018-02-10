angular.module("services.platform", [])
    .service("PlatformService", PlatformService);

PlatformService.$inject = ['APP_CONFIG', 'httpSvc'];

//管理员
function PlatformService(APP_CONFIG,httpSvc) {
    return {
        getSiteSelect: getSiteSelect,                   // 站点下拉框
        getGetAdmin: getGetAdmin,                       // 平台管理	GET/admin
        getPostAdmin: getPostAdmin,                     // 平台管理--添加	POST/admin
        getAdminInfo: getAdminInfo,                     // 平台管理--获取详细	GET/admin/info
        getPutAdmin:getPutAdmin,                        // 平台管理--修改	PUT/admin
        getAdminStatus:getAdminStatus,                  // 平台管理--修改状态	PUT/admin/status
        getAdminDel:getAdminDel,                        // 平台管理--删除	DELETE/admin
        getRoleDrop:getRoleDrop,                        // 角色下拉框
        getComboDrop:getComboDrop,                      // 套餐下拉框
        getHolderList:getHolderList,                    // 开户人管理	GET /holder/list
        getHolderAdd:getHolderAdd,                      // 开户人管理--添加	POST /holder/add
        getHolder:getHolder,                            // 开户人管理--修改	GET /holder
        getHolderUpdata:getHolderUpdata,                // 开户人管理--修改--提交	PUT /holder/updata
        getHolderDisable:getHolderDisable,              // 开户人管理--修改状态	PUT /holder/disable
        getHolderDel:getHolderDel,                      // 开户人管理--删除	DELETE
        getRole:getRole,                                // 角色管理	GET/role
        getRoleStatus:getRoleStatus,                    // 角色管理--修改状态	PUT /role/status
        getRoleDel:getRoleDel,                          // 角色管理--删除	DELETE/role
        getRoleAdd:getRoleAdd,                          // 角色管理--添加	POST/role
        getRolePermissionGet:getRolePermissionGet,      // 角色管理--权限配置--修改	GET/role/permission
        getRolePermissionPost:getRolePermissionPost,    // 角色管理--权限配置--修改	POST/role/permission
        getRoleMenuGet:getRoleMenuGet,                  // 角色管理--菜单	GET/role/menu
        getRoleMenuPost:getRoleMenuPost,                // 角色管理--菜单--修改	POST/role/menu
        getMenuAdmin:getMenuAdmin,                      // 菜单管理(平台)admin	GET/menu_admin/list
        getMenuDel:getMenuDel,                          // 菜单管理--删除	DELETE/menu/delete
        getMenuDrop:getMenuDrop,                        // 菜单管理--修改--根据id取一级二级菜单	GET/menu/drop
        getMenuPut:getMenuPut,                          // 菜单管理--修改提交	PUT/menu/put
        getMenuAdd:getMenuAdd,                          // 菜单管理--添加	POST/menu/add
        getMenuStatus:getMenuStatus,                    // 菜单管理--状态修改	PUT/menu/status
        getMenuDetail:getMenuDetail,                    // 菜单管理--详情	GET/menu/detail
        getMenuAgency:getMenuAgency,                    // 菜单管理(代理)agency	GET/menu_agency/list
        getProductList:getProductList,                  // 商品管理	GET/product/list
        getProductDel:getProductDel,                    // 商品管理--删除	DELETE/product
        getProductStatus:getProductStatus,              // 商品管理--修改状态	PUT/product/status
        getProductPut:getProductPut,                    // 商品管理--修改提交	PUT/product
        getProductGet:getProductGet,                    // 商品管理--获取商品详情	GET/product
        getProductPost:getProductPost,                  // 商品管理--添加	POST/product
        getProductType:getProductType,                  // 商品管理--商品类型	GET/product/type/infos
        getProductTypeDel:getProductTypeDel,            // 商品管理--类型管理--删除	DELETE/product/type
        getProductTypeStatus:getProductTypeStatus,      // 商品管理--类型管理--修改状态	PUT/product/type/status
        getProductTypePut:getProductTypePut,            // 商品管理--类型管理--修改提交	PUT/product/type
        getProductTypeGet:getProductTypeGet,            // 商品管理--类型管理--获取商品详情	GET/product/type
        getProductTypePost:getProductTypePost,          // 商品管理--类型管理--添加	POST/product/type
        getProductTypeInfo:getProductTypeInfo,          // 商品管理--类型管理	GET/product/types/infos
        getPermissionPost:getPermissionPost,            // 功能管理--添加	POST/permission
        getPermissionDel:getPermissionDel,              // 功能管理--删除	DELETE/permission
        getPermissionStatus:getPermissionStatus,        // 功能管理--修改状态	PUT/permission/status
        getPermissionPut:getPermissionPut,              // 功能管理--修改提交	PUT/permission
        getPermissionGet:getPermissionGet,              // 功能管理	GET/permission
        getPermissionInfo:getPermissionInfo,            // 功能管理--详情	GET/permission/info
        getProductType_C:getProductType_C,              // 套餐管理--配置--搜索	GET/product_type
        getComboPlatformPost:getComboPlatformPost,      // 套餐管理--配置--提交	POST/combo_platform
        getComboPlatformGet:getComboPlatformGet,        // 套餐管理--获取配置详情	GET/combo_platform
        getComboDel:getComboDel,                        // 套餐管理--删除	DELETE/combo
        getComboPut:getComboPut,                        // 套餐管理--修改提交	PUT/combo
        getComboInfo:getComboInfo,                      // 套餐管理--详情	GET/combo/info
        getComboStatus:getComboStatus,                  // 套餐管理--修改状态	PUT/combo/status
        getComboPost:getComboPost,                      // 套餐管理--新增	POST/combo
        getComboGet:getComboGet,                        // 套餐管理	GET/combo
        getSiteCloumnPrivate:getSiteCloumnPrivate,      // 站点栏目--私有	POST/siteCloumn/private
        getSiteCloumnPost:getSiteCloumnPost,            // 站点栏目--添加	POST/siteCloumn
        getSiteCloumnSynchro:getSiteCloumnSynchro,      // 站点栏目--栏目同步	POST/siteCloumn/synchro
        getSiteCloumnDel:getSiteCloumnDel,              // 站点栏目--删除	DELETE/siteCloumn
        getSiteCloumnPut:getSiteCloumnPut,              // 站点栏目--修改提交	PUT/siteCloumn
        getSiteCloumnGet:getSiteCloumnGet,              // 站点栏目	GET/siteCloumn
        getLog:getLog,                                  // 日志管理	GET/log
        getOperation:getOperation                       // 操作记录	GET/operation
    };
    // 操作记录	GET/operation
    function getOperation (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.OPERATION,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 日志管理	GET/log
    function getLog (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.LOG,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目	GET/siteCloumn
    function getSiteCloumnGet () {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_CLOUMN_GET).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目--修改提交	PUT/siteCloumn
    function getSiteCloumnPut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_CLOUMN_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目--删除	DELETE/siteCloumn
    function getSiteCloumnDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.SITE_CLOUMN_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目--栏目同步	POST/siteCloumn/synchro
    function getSiteCloumnSynchro () {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_CLOUMN_SYNCHRO).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目--添加	POST/siteCloumn
    function getSiteCloumnPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_CLOUMN_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 站点栏目--私有	POST/siteCloumn/private
    function getSiteCloumnPrivate (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_CLOUMN_PRIVATE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理	GET/combo
    function getComboGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COMBO_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--新增	POST/combo
    function getComboPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.COMBO_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--修改状态	PUT/combo/status
    function getComboStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.COMBO_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--详情	GET/combo/info
    function getComboInfo (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COMBO_INFO,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--修改提交	PUT/combo
    function getComboPut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.COMBO_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--删除	DELETE/combo
        function getComboDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.COMBO_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }

    // 套餐管理--配置--搜索	GET/product_type
    function getProductType_C (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_TYPE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--获取配置详情	GET/combo_platform
    function getComboPlatformGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COMBO_PLATFORM_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 套餐管理--配置--提交	POST/combo_platform
    function getComboPlatformPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.COMBO_PLATFORM_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理--详情	GET/permission/info
    function getPermissionInfo (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PERMISSION_INFO,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理	GET/permission
    function getPermissionGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PERMISSION_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理--修改提交	PUT/permission
    function getPermissionPut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PERMISSION_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理--修改状态	PUT/permission/status
    function getPermissionStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PERMISSION_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理--删除	DELETE/permission
    function getPermissionDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.PERMISSION_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 功能管理--添加	POST/permission
    function getPermissionPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PERMISSION_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理--添加	POST/product/type
    function getProductTypePost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PRODUCT_TYPE_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理--获取商品详情	GET/product/type
    function getProductTypeGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_TYPE_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理--修改提交	PUT/product/type
    function getProductTypePut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PRODUCT_TYPE_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理--修改状态	PUT/product/type/status
    function getProductTypeStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PRODUCT_TYPE_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理--删除	DELETE/product/type
    function getProductTypeDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.PRODUCT_TYPE_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--商品类型	GET/product/type/infos
    function getProductType () {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_TYPE_INFO).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--类型管理	GET/product/types/infos
    function getProductTypeInfo (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_TYPES_INFO,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理	GET/product/list
    function getProductList (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_LIST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--添加	POST/product
    function getProductPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.PRODUCT_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--获取商品详情	GET/product
    function getProductGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PRODUCT_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--删除	DELETE/product
    function getProductDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.PRODUCT_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--修改状态	PUT/product/status
    function getProductStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PRODUCT_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 商品管理--修改提交	PUT/product
    function getProductPut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PRODUCT_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }


    // 菜单管理(代理)agency	GET/menu_agency/list
    function getMenuAgency () {
        return httpSvc.get(APP_CONFIG.apiUrls.MENU_AGENCY).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--详情	GET/menu/detail
    function getMenuDetail (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MENU_DETAIL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response.data;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--添加	POST/menu/add
    function getMenuAdd (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MENU_ADD,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--修改提交	PUT/menu/put
    function getMenuPut (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.MENU_PUT,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--状态修改	PUT/menu/status
    function getMenuStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.MENU_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }

    // 菜单管理(平台)admin	GET/menu_admin/list
    function getMenuAdmin () {
        return httpSvc.get(APP_CONFIG.apiUrls.MENU_ADMIN).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--删除	DELETE/menu/delete
    function getMenuDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.MENU_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 菜单管理--修改--根据id取一级二级菜单	GET/menu/drop
    function getMenuDrop (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MENU_DROP,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response.data;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }




    // 角色管理--菜单	GET/role/menu
    function getRoleMenuGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ROLE_MENU_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--菜单--修改	POST/role/menu
    function getRoleMenuPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ROLE_MENU_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--权限配置--修改	GET/role/permission
    function getRolePermissionGet (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ROLE_PERMISSION_GET,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--权限配置--修改	POST/role/permission
    function getRolePermissionPost (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ROLE_PERMISSION_POST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理	GET/role
    function getRole (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GET_ROLE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--添加	POST/role
    function getRoleAdd (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.POST_ROLE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--修改状态	PUT /role/status
    function getRoleStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.ROLE_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色管理--删除	DELETE/role
    function getRoleDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DEL_ROLE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理--删除	DELETE
    function getHolderDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.HOLDER_DEL,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理--修改状态	PUT /holder/disable
    function getHolderDisable (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.HOLDER_DISABLE,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理--修改--提交	PUT /holder/updata
    function getHolderUpdata (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.HOLDER_UPDATA,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理--修改	GET /holder
    function getHolder (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.HOLDER,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理--添加	POST /holder/add
    function getHolderAdd (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.HOLDER_ADD,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 开户人管理	GET /holder/list
    function getHolderList (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.HOLDER_LIST,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 角色下拉框
        function getRoleDrop () {
            return httpSvc.get(APP_CONFIG.apiUrls.ROLE_DROP).then(getDataComplete)
                .catch(getDataFailed);
            function getDataComplete(response) {
                return response;
            }
            function getDataFailed(error) {
                console.log('XHR Failed for getAvengers.' + error);
            }
        }
    // 套餐下拉框
    function getComboDrop () {
        return httpSvc.get(APP_CONFIG.apiUrls.COMBO_DROP).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
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
    // 平台管理	GET/admin
    function getGetAdmin (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.GET_ADMIN,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 平台管理--添加	POST/admin
    function getPostAdmin (postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.POST_ADMIN,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 平台管理--获取详细	GET/admin/info
    function getAdminInfo (postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.ADMIN_INFO,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 平台管理--修改	PUT/admin
    function getPutAdmin (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PUT_ADMIN,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);
        }
    }
    // 平台管理--修改状态	PUT/admin/status
    function getAdminStatus (postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.ADMIN_STATUS,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }
    // 平台管理--删除	DELETE/admin
    function getAdminDel (postData) {
        return httpSvc.del(APP_CONFIG.apiUrls.DEL_ADMIN,postData).then(getDataComplete)
            .catch(getDataFailed);
        function getDataComplete(response) {
            return response;
        }
        function getDataFailed(error) {
            console.log('XHR Failed for getAvengers.' + error);

        }
    }


}