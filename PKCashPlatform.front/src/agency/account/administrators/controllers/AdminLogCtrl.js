angular.module('app.administrators').controller('AdminLogCtrl', function(DTOptionsBuilder,DTColumnBuilder,$http,$scope,$compile){
    var vm = this;
    activate();
    function activate() {
        //TODO 管理日志
        vm.standardOptions = DTOptionsBuilder
            .fromSource('server/tables/datatables.standard.json')
            .withPaginationType('full_numbers')
            //Add Bootstrap compatibility
            .withDOM("<'dt-toolbar'<'col-xs-12 col-sm-6'><'col-sm-6 col-xs-12 hidden-xs'>r>" +
                "t" +
                "<'dt-toolbar-footer'<'col-sm-6 col-xs-12 hidden-xs'><'col-xs-12 col-sm-6'p>>")
            .withBootstrap()
            .withOption('createdRow', createdRow)
            .withOption('responsive', true)

        vm.standardColumns = [
            DTColumnBuilder.newColumn('name').withClass('text-danger'),
            DTColumnBuilder.newColumn('phone'),
            DTColumnBuilder.newColumn('company')
            // DTColumnBuilder.newColumn(null).withTitle('管理').notSortable().renderWith(actionDelete)
        ];
        function createdRow(row, data, dataIndex) {
            $compile(angular.element(row).contents())($scope);
        };

        // //添加按钮
        // function actionDelete(data, type, full, meta) {
        //     return '<button  data-smart-jqui-dialog-launcher="#dialog_simple"  ng-click="sub(' + data.id + ')"  class="ids btn btn-warning" style="margin: 5px">日志</button><button  data-smart-jqui-dialog-launcher="#dialog_simple"  ng-click="sub(' + data.id + ')"  class="ids btn btn-warning" style="margin: 5px">编辑</button><button  data-smart-jqui-dialog-launcher="#dialog_simple"  ng-click="sub(' + data.id + ')"  class="ids btn btn-warning" style="margin: 5px">停用</button><button  style="margin-right: 5px"  ng-click="smartModEg1()"  class="ids btn btn-danger">删除</button>';
        // };

        //点击下拉框 状态选择
        vm.select = function ($event) {
            document.getElementsByClassName("selectValue")[0].innerText = $event.target.innerHTML;
        };
        //点击下拉框 分组选择
        vm.select_1 = function ($event) {
            document.getElementsByClassName("selectValue_1")[0].innerText = $event.target.innerHTML;
        };
        // 点击search
        $('.search').click(function () {
            $('.inp_1').val();
            $('.inp_2').val();
            $('.inp_3').val();
            console.log($('.inp_1').val());
            console.log($('.inp_2').val());
            console.log($('.inp_3').val());
        });
        // vm.search = function () {
        //     console.log(999000);
        // };
        //点击添加
        $('.add').click(function () {
            console.log(300);
        });

    }
});
