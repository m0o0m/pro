angular.module('app.administrators').controller('RegistSettingCtrl', function(BusinessService,httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,resourceSvc){
    var Email = document.getElementsByName('Email');
    var Wechat = document.getElementsByName('Wechat');
    var Passport = document.getElementsByName('Passport');
    var Mobile = document.getElementsByName('Mobile');
    var Birthday = document.getElementsByName('Birthday');
    var IsReg = document.getElementsByName('IsReg');
    var IsName = document.getElementsByName('IsName');
    var IsIntroduce = document.getElementsByName('IsIntroduce');
    var IsBankNum = document.getElementsByName('IsBankNum');
    var IsTel = document.getElementsByName('IsTel');
    var IsEmail = document.getElementsByName('IsEmail');
    var IsQq = document.getElementsByName('IsQq');
    var IsWechat = document.getElementsByName('IsWechat');
    var IsCode = document.getElementsByName('IsCode');
    var IsIp = document.getElementsByName('IsIp');
    var Qq = document.getElementsByName('Qq');
    var is_wap_single = document.getElementsByName('is_wap_single');
    var try_play = document.getElementsByName('try_play');


    httpSvc.get("/agent/first/drop").then(function (response) {
        console.log(response);
        $scope.names = response.data;
        $scope.site_length = $scope.names.length;
        console.log($scope.site_length);
        $scope.siteid = $scope.site_index_id;
        console.log($scope.siteid);
        // console.log($scope.selectedName);
        // if($scope.selectedName==undefined){
        //     $scope.selectedName = $scope.siteid;
        //     console.log($scope.selectedName);
        // }
        if($scope.site_index_id==undefined){
            $scope.site_index_id = "";
        }
        httpSvc.get("/member/register/setting",{
            site_index_id:$scope.site_index_id
        }).then(function (response) {
            console.log(response);
            if(response.data==null){
                $("input:radio[name='Email']").removeAttr("checked");
                $("input:radio[name='Wechat']").removeAttr("checked");
                $("input:radio[name='Passport']").removeAttr("checked");
                $("input:radio[name='Mobile']").removeAttr("checked");
                $("input:radio[name='Birthday']").removeAttr("checked");
                $("input:radio[name='IsReg']").removeAttr("checked");
                $("input:radio[name='IsName']").removeAttr("checked");
                $("input:radio[name='IsIntroduce']").removeAttr("checked");
                $("input:radio[name='IsBankNum']").removeAttr("checked");
                $("input:radio[name='IsTel']").removeAttr("checked");
                $("input:radio[name='IsEmail']").removeAttr("checked");
                $("input:radio[name='IsQq']").removeAttr("checked");
                $("input:radio[name='IsWechat']").removeAttr("checked");
                $("input:radio[name='IsCode']").removeAttr("checked");
                $("input:radio[name='IsIp']").removeAttr("checked");
                $("input:radio[name='Qq']").removeAttr("checked");
                $("input:radio[name='is_wap_single']").removeAttr("checked");
                $("input:radio[name='try_play']").removeAttr("checked");


                $scope.offer = "";
                $scope.add_mosaic = "";
            }else{
                // console.log(response.data.email);
                var value1 = response.data.email;
                var value2 = response.data.wechat;
                var value3 = response.data.passport;
                var value15 = response.data.mobile;
                var value4 = response.data.birthday;
                var value16 = response.data.is_reg;
                var value5 = response.data.is_name;
                var value6 = response.data.is_show_name;
                var value7 = response.data.is_card_reply;
                var value8 = response.data.is_tel;
                var value9 = response.data.is_email;
                var value10 = response.data.is_qq;
                var value11 = response.data.is_wechat;
                var value12 = response.data.is_code;
                var value13 = response.data.is_ip;
                var value14 = response.data.qq;
                var value17 = response.data.is_wap_single;
                var value18 = response.data.try_play;


                console.log(value9);
                $scope.offer = response.data.offer;
                $scope.add_mosaic = response.data.add_mosaic;
                for(var i = 0;i < 2;i++){
                    if(Email[i].value == value1){
                        Email[i].checked =  'checked';
                    }
                    if(Wechat[i].value == value2){
                        Wechat[i].checked =  'checked';
                    }
                    if(Passport[i].value == value3){
                        Passport[i].checked =  'checked';
                    }
                    if(Mobile[i].value == value15){
                        Mobile[i].checked =  'checked';
                    }
                    if(Birthday[i].value == value4){
                        Birthday[i].checked =  'checked';
                    }
                    if(IsReg[i].value == value16){
                        IsReg[i].checked =  'checked';
                    }
                    if(IsName[i].value == value5){
                        IsName[i].checked =  'checked';
                    }
                    if(IsIntroduce[i].value == value6){
                        IsIntroduce[i].checked =  'checked';
                    }
                    if(IsBankNum[i].value == value7){
                        IsBankNum[i].checked =  'checked';
                    }
                    if(IsTel[i].value == value8){
                        IsTel[i].checked =  'checked';
                    }
                    if(IsEmail[i].value == value9){
                        IsEmail[i].checked =  'checked';
                    }
                    if(IsQq[i].value == value10){
                        IsQq[i].checked =  'checked';
                    }
                    if(IsWechat[i].value == value11){
                        IsWechat[i].checked =  'checked';
                    }
                    if(IsCode[i].value == value12){
                        IsCode[i].checked =  'checked';
                    }
                    if(IsIp[i].value == value13){
                        IsIp[i].checked =  'checked';
                    }
                    if(Qq[i].value == value14){
                        Qq[i].checked =  'checked';
                    }
                    if(is_wap_single[i].value == value14){
                        is_wap_single[i].checked =  'checked';
                    }
                    if(try_play[i].value == value14){
                        try_play[i].checked =  'checked';
                    }
                }
            }
        });
    });

    $scope.search = function () {
        console.log($scope.site_index_id);
        console.log($scope.siteid);
        if($scope.site_index_id == undefined){
            $scope.site_index_id = "";
        }
        // else{
        //     var siteindex = $scope.selectedName;
        // }
        httpSvc.get("/member/register/setting",{
            site_index_id:$scope.site_index_id
        }).then(function (response) {
            console.log(response);
            if(response.data == null){
                console.log(123321);
                $("input:radio[name='Email']").removeAttr("checked");
                $("input:radio[name='Wechat']").removeAttr("checked");
                $("input:radio[name='Passport']").removeAttr("checked");
                $("input:radio[name='Mobile']").removeAttr("checked");
                $("input:radio[name='Birthday']").removeAttr("checked");
                $("input:radio[name='IsReg']").removeAttr("checked");
                $("input:radio[name='IsName']").removeAttr("checked");
                $("input:radio[name='IsIntroduce']").removeAttr("checked");
                $("input:radio[name='IsBankNum']").removeAttr("checked");
                $("input:radio[name='IsTel']").removeAttr("checked");
                $("input:radio[name='IsEmail']").removeAttr("checked");
                $("input:radio[name='IsQq']").removeAttr("checked");
                $("input:radio[name='IsWechat']").removeAttr("checked");
                $("input:radio[name='IsCode']").removeAttr("checked");
                $("input:radio[name='IsIp']").removeAttr("checked");
                $("input:radio[name='Qq']").removeAttr("checked");
                $("input:radio[name='is_wap_single']").removeAttr("checked");
                $scope.quota = "";
                $scope.offer = "";
                $scope.add_mosaic = "";
                $scope.request = "post";
            }else{
                $scope.request = "put";
                var value1 = response.data.email;
                var value2 = response.data.wechat;
                var value3 = response.data.passport;
                var value15 = response.data.mobile;
                var value4 = response.data.birthday;
                var value16 = response.data.is_reg;
                var value5 = response.data.is_name;
                var value6 = response.data.is_show_name;
                var value7 = response.data.is_card_reply;
                var value8 = response.data.is_tel;
                var value9 = response.data.is_email;
                var value10 = response.data.is_qq;
                var value11 = response.data.is_wechat;
                var value12 = response.data.is_code;
                var value13 = response.data.is_ip;
                var value14 = response.data.qq;
                var value17 = response.data.is_wap_single;
                var value18 = response.data.try_play;

                $scope.quota = response.data.quota;
                $scope.offer = response.data.offer;
                $scope.add_mosaic = response.data.add_mosaic;
                for(var i = 0;i < 2;i++){
                    if(Email[i].value == value1){
                        Email[i].checked =  'checked';
                    }
                    if(Wechat[i].value == value2){
                        Wechat[i].checked =  'checked';
                    }
                    if(Passport[i].value == value3){
                        Passport[i].checked =  'checked';
                    }
                    if(Mobile[i].value == value15){
                        Mobile[i].checked =  'checked';
                    }
                    if(Birthday[i].value == value4){
                        Birthday[i].checked =  'checked';
                    }
                    if(IsReg[i].value == value16){
                        IsReg[i].checked =  'checked';
                    }
                    if(IsName[i].value == value5){
                        IsName[i].checked =  'checked';
                    }
                    if(IsIntroduce[i].value == value6){
                        IsIntroduce[i].checked =  'checked';
                    }
                    if(IsBankNum[i].value == value7){
                        IsBankNum[i].checked =  'checked';
                    }
                    if(IsTel[i].value == value8){
                        IsTel[i].checked =  'checked';
                    }
                    if(IsEmail[i].value == value9){
                        IsEmail[i].checked =  'checked';
                    }
                    if(IsQq[i].value == value10){
                        IsQq[i].checked =  'checked';
                    }
                    if(IsWechat[i].value == value11){
                        IsWechat[i].checked =  'checked';
                    }
                    if(IsCode[i].value == value12){
                        IsCode[i].checked =  'checked';
                    }
                    if(IsIp[i].value == value13){
                        IsIp[i].checked =  'checked';
                    }
                    if(Qq[i].value == value14){
                        Qq[i].checked =  'checked';
                    }
                    if(is_wap_single[i].value == value14){
                        is_wap_single[i].checked =  'checked';
                    }
                    if(try_play[i].value == value14){
                        try_play[i].checked =  'checked';
                    }
                }
            }
            console.log($scope.request);
        })
    }
    $scope.submit = function () {
        $scope.Email = $("input[name='Email']:checked").val();
        $scope.Wechat = $("input[name='Wechat']:checked").val();
        $scope.Passport = $("input[name='Passport']:checked").val();
        $scope.Mobile = $("input[name='Mobile']:checked").val();
        $scope.Birthday = $("input[name='Birthday']:checked").val();
        $scope.IsReg = $("input[name='IsReg']:checked").val();
        $scope.IsName = $("input[name='IsName']:checked").val();
        $scope.IsIntroduce = $("input[name='IsIntroduce']:checked").val();
        $scope.IsBankNum = $("input[name='IsBankNum']:checked").val();
        $scope.IsTel = $("input[name='IsTel']:checked").val();
        $scope.IsEmail = $("input[name='IsEmail']:checked").val();
        $scope.IsQq = $("input[name='IsQq']:checked").val();
        $scope.IsWechat = $("input[name='IsWechat']:checked").val();
        $scope.IsCode = $("input[name='IsCode']:checked").val();
        $scope.IsIp = $("input[name='IsIp']:checked").val();
        $scope.Qq = $("input[name='Qq']:checked").val();
        $scope.is_wap_single = $("input[name='is_wap_single']:checked").val();
        $scope.try_play = $("input[name='try_play']:checked").val();

        console.log($scope.is_wap_single);
        console.log($scope.Email);
        console.log($scope.Wechat);
        console.log($scope.Passport);
        console.log($scope.Mobile);
        console.log($scope.Birthday);
        console.log($scope.IsReg);
        console.log($scope.IsName);
        console.log($scope.IsIntroduce);
        console.log($scope.IsBankNum);
        console.log($scope.IsTel);
        console.log($scope.IsEmail);
        console.log($scope.IsQq);
        console.log($scope.IsWechat);
        console.log($scope.IsCode);
        console.log($scope.IsIp);
        console.log($scope.Qq);
        console.log($scope.offer);
        console.log($scope.add_mosaic);
        console.log($scope.site_index_id);
        console.log($scope.siteid);
        if($scope.site_index_id == undefined){
            var siteindexid = $scope.siteid;
        }else{
            var siteindexid = $scope.site_index_id;
        }
        $scope.registsetting = {
            email:$scope.Email*1,
            wechat:$scope.Wechat*1,
            passport:$scope.Passport*1,
            mobile:$scope.Mobile*1,
            birthday:$scope.Birthday*1,
            is_reg:$scope.IsReg*1,
            is_name:$scope.IsName*1,
            is_show_name:$scope.IsIntroduce*1,
            is_card_reply:$scope.IsBankNum*1,
            is_tel:$scope.IsTel*1,
            is_email:$scope.IsEmail*1,
            is_qq:$scope.IsQq*1,
            is_wechat:$scope.IsWechat*1,
            is_code:$scope.IsCode*1,
            offer:$scope.offer*1,
            add_mosaic:$scope.add_mosaic*1,
            is_ip:$scope.IsIp*1,
            qq:$scope.Qq*1,
            site_index_id:siteindexid,
            is_wap_single:$scope.is_wap_single*1,
            try_play:$scope.try_play*1,
            quota:$scope.quota*1
        }
        if($scope.offer==''||$scope.offer==''||$scope.add_mosaic==''||$scope.Email==undefined||$scope.Wechat==undefined||$scope.Passport==undefined||$scope.Mobile==undefined||$scope.Birthday==undefined||$scope.IsReg==undefined||$scope.IsName==undefined||$scope.IsIntroduce==undefined||$scope.IsBankNum==undefined||$scope.IsTel==undefined||$scope.IsEmail==undefined||$scope.IsQq==undefined||$scope.IsWechat==undefined||$scope.IsCode==undefined||$scope.offer==undefined||$scope.add_mosaic==undefined||$scope.IsIp==undefined||$scope.Qq==undefined||$scope.is_wap_single==undefined||$scope.try_play==undefined){
            popupSvc.smallBox("fail","请输入完整！");
        }else{

            httpSvc.post("/member/register/setting",$scope.registsetting).then(function (response) {
                console.log(response);
                if(response==null){
                    popupSvc.smallBox("success","操作成功");
                    // GetAllEmployee();
                }else{
                    $scope.msg = response.msg;
                    popupSvc.smallBox("fail",$scope.msg);
                }
            })
        }
    }
});