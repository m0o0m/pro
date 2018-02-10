angular.module('app.administrators').controller('AccountsCtrl',
  function ( AccountService, $stateParams, DTColumnBuilder, $scope, $rootScope, APP_CONFIG, $state, $LocalStorage) {
    $scope.list = [];

    $scope.toggleAdd = function () {
      if (!$scope.newTodo) {
        $scope.newTodo = {
          state: 'Important'
        };
      } else {
        $scope.newTodo = undefined;
      }
    };
    //初始化下拉状态
    /*
     */
    $scope.stateJson = APP_CONFIG.option.option_status;
    $scope.sourceJson = APP_CONFIG.option.option_reg;
    $scope.isOnlineJson = APP_CONFIG.option.option_online;
    $scope.memberTypeJson = APP_CONFIG.option.option_type;

    //获取代理
    $scope.sitId = function (site_index_id) {
      AccountService.getThirdAgent(site_index_id).then(function (response) {
        $scope.sharedJson = response.data;
      });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
      //获取全部站点
      $scope.sitId("");
    } else {
      $scope.sitId(user.site_index_id);
    }

    // 停用
    $scope.disable = function (item) {
      var status = 2;
      if (item.status === 0 || item.status === 1) {
        status = 2;
      } else {
        status = 1;
      }
      //1正常2禁用
      var sure = function () {
        AccountService.setMemberStatus(item.id, status).then(function (response) {
          if (response.status == 200 || response.status == 204) {
            item.status = status;
            popupSvc.smallBox("success", $rootScope.getWord("success"));
          } else {
            popupSvc.smallBox("fail", response.data.msg);
          }
        });
      };
      popupSvc.smartMessageBox($rootScope.getWord("确定更改状态") + "?", sure);
    };


    $scope.paginationConf = {
      currentPage: 1,
      itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };

    var GetAllEmployee = function () {
      var postData = {
        page: $scope.paginationConf.currentPage,
        page_size: $scope.paginationConf.itemsPerPage,
        status: $scope.statusd,
        key: $scope.type,
        type_value: $scope.IP,
        start_time: $scope.startTime,
        end_time: $scope.endTime,
        source: $scope.reg,
        online: $scope.online,
        type: $scope.type,
        site_index_id: $scope.site_index_id,
        agency_id: $scope.accountse,
        first_id: $stateParams.first_id,
        second_id: $stateParams.second_id
      };
      AccountService.getMembers(postData).then(function (response) {
        $scope.paginationConf.totalItems = response.meta.count;
        $scope.list = response.data;
      })
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.member = {
      is_edit_password: "",
      reply_password: "",
      new_password: "",
      user_name: "",
      id: ""
    };

    //修改会员密码
    $scope.getID = function (mid) {
      AccountService.getMembersInfo(mid).then(function (data) {
        $scope.member.id = data.data.id;
        $scope.member.user_name = data.data.account;
        $scope.member.is_edit_password = data.data.is_edit_password;
        $scope.member.realname = data.data.realname;
      });
    };
    //修改后提交
    $scope.submit = function () {
      AccountService.setMembersInfo($scope.member).then(
        function (data) {
          if (data == null) {
            popupSvc.smallBox("success", $rootScope.getWord("success"));
          } else {
            popupSvc.smallBox("fail", data.msg);
          }
        })
    };
    //修改资料
    $scope.Info = function (id) {
      $state.go('app.administrators.accountInfo', {
          ids: id
        }
      );
    };
    //点击搜索
    $scope.search = function () {
      $scope.site_index_id = $('#site_index_id').val();
      $scope.accountse = $('#accountsd').val();
      GetAllEmployee();

    };
    //银行
    $scope.bank = function (bankid) {
      $state.go('app.administrators.bank', {
          ids: bankid
        }
      );

    };
    //点击跳转代理
    $scope.agents = function (agentsid) {
      $state.go('app.administrators.agent', {
        gene: agentsid
      });
    };
  });



