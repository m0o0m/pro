angular.module('app.site')
    .controller('modifySiteCtrl', function($scope,DTOptionsBuilder,DTColumnBuilder,$state,$compile,$http,$stateParams){
        var vm = this;
        activate();
        function activate() {
            vm.id = $stateParams.id;
            console.log(vm.id);
            console.log(112233);
            var apiURL = "http://192.168.8.173:10080/";
            $http({
                url:apiURL+'site/?return=site_name,id,site_agency,creat_time,status,site_service_info',
                method:"get"
            }).success(function (data) {
                console.log(data);
                vm.dataJson = data.data;
                console.log(vm.dataJson);
                vm.data = vm.dataJson[vm.id];
                console.log(vm.data.ID);
                console.log(vm.data.Status);
                var sel = document.getElementById('sel');
                var val = sel.options;
                console.log(val);
                if(vm.data.Status == 1){
                    val[0].selected = true;
                }else{
                    val[1].selected = true;
                }
            });
            vm.submit = function () {
                var formData = new FormData($("#formid")[0]);
                var apiURL = apiURL+'site/';
                $.ajax({
                    url: apiURL + vm.data.ID,
                    type: 'PUT',
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

        }
    });


















