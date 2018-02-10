/**
 * Created by mebar on 08/12/2017.
 * 股东管理的 shareholder services
 */
angular.module("services.account", [])
  .service("ShareholderService", ShareholderService);

ShareholderService.$inject = ['APP_CONFIG', 'httpSvc'];

function ShareholderService(APP_CONFIG, httpSvc) {

  //获取站点id 
  function getFirstDropSelect() {
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

  function getSecondDropSelect() {
    return httpSvc.get(APP_CONFIG.apiUrls.SECOND_DROPF)
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      console.log(response);
      return response.data;
    }

    function getDataFailed(error) {
      console.log('XHR Failed for getAvengers.' + error);
    }
  }

  function getThirdDropSelect(site_index_id,first_id) {
    return httpSvc.get(APP_CONFIG.apiUrls.THIRD_DROPF, {
        site_index_id: site_index_id,
        first_id:first_id
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
  //获取低3级
  function getThirdAgent(site_index_id) {
    return httpSvc.get(APP_CONFIG.apiUrls.THIRD_AGENT_DROPF, {
        site_index_id: site_index_id
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

  function setMemberStatus(id, status) {
    return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_STATUS, {
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

  function getMembers(postData) {
    return httpSvc.get(APP_CONFIG.apiUrls.MEMBER, postData)
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response.data;
    }

    function getDataFailed(error) {
      console.error('XHR Failed for getAvengers.' + error);
    }
  }

  function getMembersInfo(mid) {
    return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_INFO, {
        id: mid
      }).then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response.data;
    }

    function getDataFailed(error) {
      console.error('XHR Failed for getAvengers.' + error);
    }
  }

  function setMembersInfo(member) {
    return httpSvc.put(APP_CONFIG.apiUrls.MEMBER_INFO, {
        id: member.id,
        is_edit_password: member.is_edit_password * 1,
        reply_password: member.reply_password,
        new_password: member.new_password,
        realname: member.realname

      }).then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response;
    }

    function getDataFailed(error) {
      console.log('XHR Failed for getAvengers.' + error);
    }
  }
  //获取用户详细信息MEMBER_DETAIL
  function getMembersDetail(id) {
    return httpSvc.get(APP_CONFIG.apiUrls.MEMBER_DETAIL, {
        id: id
      }).then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response.data;
    }

    function getDataFailed(error) {
      console.error('XHR Failed for getAvengers.' + error);
    }
  }
  //修改用户信息
  function putMembersDetail(detail) {
    return httpSvc.put(APP_CONFIG.apiUrls.MEMBER_DETAIL, {
        id: detail.id,
        realname: detail.realname,
        birthday: detail.birthday,
        card: detail.card,
        mobile: detail.mobile,
        email: detail.email,
        qq: detail.qq,
        wechat: detail.wechat,
        draw_password: detail.draw_password,
        remark: detail.remark
      }).then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response.data;
    }

    function getDataFailed(error) {
      console.error('XHR Failed for getAvengers.' + error);
    }
  }

  function putMemberLevel(formData) {
    return httpSvc.post(APP_CONFIG.apiUrls.MEMBER_LEVEL, formData)
      .then(getDataComplete)
      .catch(getDataFailed);

    function getDataComplete(response) {
      return response.data;
    }

    function getDataFailed(error) {
      console.error('XHR Failed for getAvengers.' + error);
    }
  }

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

}