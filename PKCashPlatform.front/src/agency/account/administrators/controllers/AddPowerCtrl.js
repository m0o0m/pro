angular.module('app.administrators').controller('AddPowerCtrl', function(DTOptionsBuilder,DTColumnBuilder,$http,$scope,$compile,APP_CONFIG,$state){
    var vm = this;
    activate();
    function activate() {

        vm.callback = function () {
            $state.go('app.administrators.power');
        };
        vm.see = function () {
            $state.go('app.administrators.accounts');
        };
    }
});
