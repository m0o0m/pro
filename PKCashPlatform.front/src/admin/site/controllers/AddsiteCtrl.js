/**
 * Created by 俊 on 2017/7/25.
 */
angular.module('app.site').controller('AddsiteCtrl', function(DTOptionsBuilder, DTColumnBuilder,$http,$scope,$state,$compile){
    var vm = this;
    activate();
    function activate() {
        var apiURL = "http://192.168.8.173:10080/";
        $http({
            url:apiURL+'agency/?return=agency_name&agency_parent_id=0',
            method:"get"
        }).success(function (data) {
            console.log(data);
            vm.dataJson = data.data;
            console.log(vm.dataJson);
        });
        vm.submit = function () {
            console.log(vm.file);
            var formData = new FormData($("#formid")[0]);
            $.ajax({
                url: 'http://192.168.8.173:10080/site/',
                type: 'POST',
                data: formData,
                async: false,
                cache: false,
                contentType: false,
                processData: false,
                success: function(data) {
                    console.log(data);
                    $state.go('app.site.siteList');
                },  
                error: function(data) {
                    console.log(data);

                }
            });
            console.log(formData);
        };
        vm.add = function ($event) {
            var html = "<div class='form-group'> <label class='col-md-2 control-label'>PC域名:</label> <div class='col-md-3'> <input class='form-control' placeholder='请输入PC域名' name='siteName' type='text'> </div> <label class='col-md-2 control-label'>WAP域名:</label> <div class='col-md-3'> <input class='form-control' placeholder='请输入WAP域名' name='siteName' type='text'> </div> <div class='form-group'></div> <label class='col-md-2 control-label'>SSI证书:</label> <div class='col-md-3'> <input type='file' class='btn btn-default form-control' value='FileUploader' name='siteLogo'> </div> <label class='col-md-2 control-label'></label> <div class='col-md-3'> <input type='file' class='btn btn-default form-control' value='FileUploader' name='siteLogo'> </div>                                 <div class='col-md-1' ng-click='addsite.remove($event)'><i class='fa fa-times relmove'></i></div> <div class='form-group'></div></div>";
            // var html = "<div class='form-group'> <br><label class='col-sm-2 control-label'></label><div class='col-sm-3'><input class='form-control' type='text' name='activitydigest' ng-model='release.activitydigest' placeholder='中文：苹果'/></div><div class='col-sm-1 reltop'><a href='javascript:void(0);' class='button-icon jarviswidget-delete-btn relmar' rel='tooltip' title='' data-placement='bottom' data-original-title='Delete' ng-click='release.add($event)'><i class='fa fa-plus reladd'></i></a><a href='javascript:void(0);' class='button-icon jarviswidget-delete-btn' ng-click='release.move($event)' rel='tooltip' title='' data-placement='bottom' data-original-title='Delete'><i class='fa fa-times relmove'></i></a></div><div class='col-sm-1 relect'><select class='form-control' ><option value='0'>全部</option><option value='1'>类一</option><option value='2'>类二</option><option value='3'>类三</option><option value='4'>类四</option></select></div></div>"
            var template = angular.element(html);
            var mobile = $compile(template)($scope);
            angular.element(document.getElementsByClassName('addInp')).append(mobile);
            // $event.target.remove();
        };
        vm.remove = function ($event) {
            var parent = $event.target.parentNode;
            parent.parentNode.remove();
        };
        vm.submit = function(){
            var formData = new FormData($("#formdata")[0]);
            console.log(formData);
            $.ajax({
                url: 'http://192.168.8.152:10080/site/',
                type: 'POST',
                data: formData,
                async: false,
                cache: false,
                contentType: false,
                processData: false,
                success: function(data) {
                    console.log(data);
                },
                error: function(data) {
                    console.log(data);

                }
            });
        }
        // $scope.submit = function(){
        //     var formData = new FormData($("#formid")[0]);
        //     $.ajax({
        //         url: APP_MEDIAQUERY.apiUrl+'/pc/v1/addindexcdata' + '/' + localStorage.token +'/'+ localStorage.id,
        //         type: 'POST',
        //         data: formData,
        //         async: false,
        //         cache: false,
        //         contentType: false,
        //         processData: false,
        //         success: function(data) {
        //             console.log(data);
        //             if(data.code==1){
        //                 $state.go('app.activity');
        //             }else{
        //                 return false;
        //             }
        //         },
        //         error: function(data) {
        //             console.log(data);
        //
        //         }
        //     });
        // };


    }
});