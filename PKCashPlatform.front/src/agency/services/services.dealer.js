/**
 * Created by mebar on 08/12/2017.
 * 代理管理 dealer的services
 */
angular.module("services.account", [])
  .service("DealerService", DealerService);

DealerService.$inject = ['APP_CONFIG', 'httpSvc'];

function DealerService(APP_CONFIG, httpSvc) {
  return {
    getThirdAgent: getThirdAgent,
    getFirstDropSelect: getFirstDropSelect,
    getSecondDropSelect:getSecondDropSelect,
    getThirdDropSelect:getSecondDropSelect,
    setMemberStatus: setMemberStatus,
    getMembers: getMembers,
    getMembersInfo: getMembersInfo,
    setMembersInfo: setMembersInfo,
    getMembersDetail: getMembersDetail,
    putMembersDetail: putMembersDetail,
    putMemberLevel: putMemberLevel
  };
  //获取站点id 代理列表
  function getDealerList() {
    return httpSvc.get(APP_CONFIG.apiUrls.FIRST_DROPF)
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      console.log(response)
      return response.data;
    }

    function getDataFailed(error) {
      console.log('XHR Failed for getAvengers.' + error);
    }
  }

  

}