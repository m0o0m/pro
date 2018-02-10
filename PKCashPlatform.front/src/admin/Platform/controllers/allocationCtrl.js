angular.module('app.Platform').controller('allocationCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state,$stateParams) {
    $scope.toggleAdd = function () {
        if (!$scope.newTodow) {
            $scope.newTodow = {
                state: 'Important'
            };
        } else {
            $scope.newTodow = undefined;
        }
    };
    var GetAllEmployee = function () {
        var postData = {
            combo_id:$stateParams.ids,
        };
        PlatformService.getComboPlatformGet(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data;
        })
    };
    GetAllEmployee();
    //点击搜索
    $scope.search = function () {
        var postData = {
            name:$scope.namese,
            combo_id:$stateParams.ids
        };
        PlatformService.getProductType_C(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data;
        })
    };
    // $scope.ckecked = function (event) {
    //     var parent = $(event.target).parent()[0];
    //     var input= $(parent).siblings(".inputsese")[0];
    //     console.log(input);
    //     if(event.target.checked){
    //         console.log($(input));
    //         $(input).removeClass('ng-hide');
    //         $(input).show();
    //     }else {
    //         // $(input).find('.inputse').attr("readonly","readonly");
    //     }
    // };
    // $scope.ckecked = function (event) {
    //     var parent = $(event.target).parent()[0];
    //     var input= $(parent).sibling(".inputsese")[0];
    //     if(event.target.checked){
    //         $(input).removeClass('ng-hide');
    //         $(input).show();
    //     }else {
    //         $(input).hide();
    //     }
    // };
    //点击提交
    $scope.sub = function () {
        var check_val = [];
        var test  = document.getElementsByClassName('test');
        for (var j=0;j<test.length;j++){
           if(test[j].checked) {
               var obj ={
                   platform_id:test[j].value*1,
                   proportion:$(test[j]).parent().siblings(".inputsese").find('.inputse')[0].value*1
               };
               check_val.push(obj);
           }

        }
        var postData = {
            combo_id:$stateParams.ids*1,
            params:check_val
        };
        PlatformService.getComboPlatformPost(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };

    $scope.inputBlur = function () {
        var inputVal = document.getElementById("inp_2").value;
        var searchBtn = document.getElementById("searchBtn");
        console.log(inputVal);
        if (!inputVal == ""){
            searchBtn.removeAttribute("disabled");
        }else {
            searchBtn.getAttribute("disabled");
            console.log(8989)
        }
    };

    $scope.$watch('namese',  function() {
        var inputVal = document.getElementById("inp_2").value;
        var searchBtn = document.getElementById("searchBtn");
            if (!inputVal == ""){
                searchBtn.removeAttribute("disabled");
            }else {
                searchBtn.setAttribute("disabled","disabled");
            }
    });

});