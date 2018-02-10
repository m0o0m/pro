/**
 * Created by griga on 2/9/16.
 */


angular.module('app.tables').controller('DatatablesCtrl', function(DTOptionsBuilder, DTColumnBuilder,$http,$scope){
    $http({
        url:'server/tables/datatables.standard.json',
        method:"get"
    }).success(function (data) {
        $scope.json=data;
        console.log(data);
    })
    // //s
    // this.standardOptions = DTOptionsBuilder
    //     .fromSource('server/tables/datatables.standard.json')
    //      //Add Bootstrap compatibility
    //     .withDOM("t<'dt-toolbar-footer'<'col-sm-6 col-xs-12 hidden-xs'l><'col-xs-12 col-sm-6'p>>")
     //   .withBootstrap();//.withOption('bInfo',false);;
    // this.standardColumns = [
    //     DTColumnBuilder.newColumn('id').withClass('text-danger'),
    //     DTColumnBuilder.newColumn('name'),
    //     DTColumnBuilder.newColumn('phone'),
    //     DTColumnBuilder.newColumn('company'),
    //     DTColumnBuilder.newColumn('zip'),
    //     DTColumnBuilder.newColumn('city'),
    //     DTColumnBuilder.newColumn('date')
    // ];


});