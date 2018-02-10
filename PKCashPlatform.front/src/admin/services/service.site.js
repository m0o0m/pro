angular.module("services.site", [])
    .service("siteService", siteService);

siteService.$inject = ['APP_CONFIG','httpSvc'];
function siteService(APP_CONFIG,httpSvc) {

    return {
        mainteaceList:mainteaceList,
        mainTain:mainTain,
        mainTainProject:mainTainProject,
        thirdDropf:thirdDropf,
        copywariting:copywariting,
        copywaritingHistory:copywaritingHistory,
        copywaritingDerail:copywaritingDerail,
        copywaritingDel:copywaritingDel,
        sitepassword:sitepassword,
        sitepasswordDetail:sitepasswordDetail,
        sitepasswordModify:sitepasswordModify,
        linkDownload:linkDownload,
        addVideo:addVideo,
        downloadlinksModify:downloadlinksModify,
        downloadModify:downloadModify,
        downloadStatus:downloadStatus,
        jsWep:jsWep,
        JSTable:JSTable,
        JSPc:JSPc,
        JSDetail:JSDetail,
        JSDel:JSDel,
        tableModify:tableModify,
        genrateWep:genrateWep,
        genratePc:genratePc,
        IPSwitching:IPSwitching,
        IPWhiteList:IPWhiteList,
        IPSwitchAdd:IPSwitchAdd,
        IPSwitchModify:IPSwitchModify,
        siteManagemnet:siteManagemnet,
        moduleManagemnet:moduleManagemnet,
        negative:negative,
        multistation:multistation,
        hierchicalData:hierchicalData,
        hierchicalDataModify:hierchicalDataModify,
        hierchicalDataDel:hierchicalDataDel,
        hierchicalDataAdd:hierchicalDataAdd,
        proxydata:proxydata,
        proxydataModify:proxydataModify,
        addlowerLevel:addlowerLevel,
        agentEel:agentEel,
        videoConfifurtion:videoConfifurtion,
        adtaAdmin:adtaAdmin,
        adtaAdminModify:adtaAdminModify,
        adtaAdminAdd:adtaAdminAdd,
        adtaAdminDel:adtaAdminDel,
        maintencesttings:maintencesttings,
        savesettings:savesettings,
        siteAdd:siteAdd,
        quotaoperation:quotaoperation,
        modeular:modeular,
        goOniline:goOniline,
        negativeAdd:negativeAdd,
        multistationAdd:multistationAdd,
        multistationAddAgent:multistationAddAgent,
        oneTochAgent:oneTochAgent,
        multistationModify:multistationModify,
        siteModify:siteModify,
        copywritingStatus:copywritingStatus

    };

    //维护管理列表
    function  mainteaceList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.MAINTENANCE, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }

    //维护-全网维护/维护
    function  mainTain(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MAIN_TAIN, postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //维护项目列表
    function  mainTainProject() {
        return httpSvc.get(APP_CONFIG.apiUrls.MAINTAIN_PROJECT)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点下拉
    function  thirdDropf() {
        return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //
    function  copywariting(postData) {
    return httpSvc.get(APP_CONFIG.apiUrls.COPYWRITING,postData)
        .then(getDataComplete)
        .catch(getDataFailed);

        function getDataComplete(response) {
         return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //文案历史
    function  copywaritingHistory(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.COPYWRITING_HISTORY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取文案详情
    function  copywaritingDerail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.COPYWRITING_DETAIL,{
            id:id
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
    //删除文案
    function  copywaritingDel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.COPYWRITING_DEL,{
            id:id
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
    //站点口令设置列表
    function  sitepassword(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITEPASSWORD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取单个口令设置
    function  sitepasswordDetail(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITEPASSWORD_DETALI,{
            id:id
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
    //修改单个口令
    function  sitepasswordModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITEPASSWORD_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //下载链接地址列表
    function  linkDownload() {
        return httpSvc.get(APP_CONFIG.apiUrls.LINKDOWNLOAD)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //添加地址
    function  addVideo(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ADD_VIDEO,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取单个下载链接
    function  downloadlinksModify(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.DOWNLOADLINKS_MODIFY,{
            id:id
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
    //修改单个下载链接
    function  downloadModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.DOWNLOAD_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //修改单个下载链接状态
    function  downloadStatus(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.DOWNLOAD_STATUS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //Js版本号-wep
    function  jsWep() {
        return httpSvc.get(APP_CONFIG.apiUrls.JS_WEP)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //Js版本号-总表
    function  JSTable() {
        return httpSvc.get(APP_CONFIG.apiUrls.JS_TABLE)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //js版本号-pc
    function  JSPc() {
        return httpSvc.get(APP_CONFIG.apiUrls.JS_PC)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取 详情
    function  JSDetail(id,type_id) {
        return httpSvc.get(APP_CONFIG.apiUrls.JS_DETAIL,{
            id:id,
            type_id:type_id
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
    //删除wep/pc
    function  JSDel(id,type_id) {
        return httpSvc.del(APP_CONFIG.apiUrls.JS_del,{
            id:id,
            type_id:type_id
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
    //编辑总表

    function  tableModify(quipment) {
        return httpSvc.put(APP_CONFIG.apiUrls.TABLE_MODIFY,{
            quipment:quipment
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
    //生成wep版本号

    function  genrateWep(count) {
        return httpSvc.post(APP_CONFIG.apiUrls.GENERATE_WEP,{
            count:count
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
    //生成PC版本号

    function  genratePc(count) {
        return httpSvc.post(APP_CONFIG.apiUrls.GENERATE_PC,{
            count:count
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
    //IP开关操作
    function  IPSwitching(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.IP_SWITCHING,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //ip白名单吧

    function  IPWhiteList(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.IP_WHITE_LIST,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //添加ip操作

    function  IPSwitchAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.IP_SWITCH_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //修改Ip开关操作

    function  IPSwitchModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.IP_SWITCH_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理
    function  siteManagemnet(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.SITE_MANAGEMENT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //模块管理
    function  moduleManagemnet(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.MODULE_MANAGEMENT,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取负数
    function  negative(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.NEGATIVE,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取多站点

    function  multistation(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.MULTISTATION,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //获取层级数据
    function  hierchicalData(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.HIERrchicaldata,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //层级数据-修改

    function  hierchicalDataModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.HIERrchicaldataModify,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //层级数据-删除

    function  hierchicalDataDel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.HIERrchicaldataDel,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //层级数据添加

    function  hierchicalDataAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.HIERrchicaldataAdd,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //代理数据

    function  proxydata(postData) {
        return httpSvc.get(APP_CONFIG.apiUrls.PROXYDATA,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //编辑代理数据

    function  proxydataModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.PROXYDATA_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //添加代理下级

    function  addlowerLevel(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ADDLOWERLEVEL,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //删除

    function  agentEel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.AGENT_DEL,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //数据-视讯配置

    function  videoConfifurtion(id) {
        return httpSvc.put(APP_CONFIG.apiUrls.VIDEOCONFIFURATION,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //管理元列表
    function  adtaAdmin(id) {
        return httpSvc.get(APP_CONFIG.apiUrls.DATA_ADMIN,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //数据-管理员-修改

    function  adtaAdminModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.DATA_ADMIN_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //数据-管理员-添加

    function  adtaAdminAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.DATA_ADMIN_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //数据-管理员-删除

    function  adtaAdminDel(id) {
        return httpSvc.del(APP_CONFIG.apiUrls.DATA_ADMIN_DEL,{
            id:id
        })
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //全站维护设置

    function  maintencesttings() {
        return httpSvc.get(APP_CONFIG.apiUrls.MAINTENANCESETTINGS)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //保存设置

    function  savesettings(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SAVESETTINGS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理（添加站点）
    function  siteAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.SITE_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理(修改额度操作)

    function  quotaoperation(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.QUOTAOPERATION,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理(模块管理)

    function  modeular(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.MODULAR,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理(上线)

    function  goOniline(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.GOONLINE,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理(负数管理)-添加

    function  negativeAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.NEGATIVE_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理（多站点）-添加

    function  multistationAdd(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MULTISTATION_ADD,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理(多站点)-添加代理

    function  multistationAddAgent(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.MULTISTATION_ADDAGENT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理（一键生成三级管理添加）

    function  oneTochAgent(postData) {
        return httpSvc.post(APP_CONFIG.apiUrls.ONETOCH_AGENT,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理（多站点编辑）

    function  multistationModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.MULTISTATION_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //站点管理（站点编辑）

    function  siteModify(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.SITE_MODIFY,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }
    //文案管理（拒绝/通过）

    function  copywritingStatus(postData) {
        return httpSvc.put(APP_CONFIG.apiUrls.COPYWRITING_STATUS,postData)
            .then(getDataComplete)
            .catch(getDataFailed);

        function getDataComplete(response) {
            return response;
        }

        function getDataFailed(error) {
            console.error('XHR Failed for getAvengers.' + error);
        }
    }


}