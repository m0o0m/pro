angular.module('app.ReportForm').controller('ReDetailCtrl',
    function(httpSvc,popupSvc,$scope,financeReportService,APP_CONFIG,$rootScope,$stateParams,$state,$timeout){
        //获取参数
        var postData = {
            ogStartTime:$stateParams.ogStartTime,
            ogEndTime:$stateParams.ogEndTime,
            otherStartTime:$stateParams.otherStartTime,
            otherEndTime:$stateParams.otherEndTime,
            module: $stateParams.module
        };

        //获取详情列表
        financeReportService.reportquery(postData).then(function (res) {
            $scope.data = res;
        });


        $scope.aa = function (event) {
            var target=event.target;
            var flag=$(target).attr('myAttr')==1?2:1;
            $(target).attr('myAttr',flag);
            if (flag==2){
                var id = $(target).attr('myIds');
                financeReportService.hareholdersStatement(id).then(function (res) {
                    console.log(res.data.length);
                    $scope.list1 = res.data;
                    var list=$scope.list1;
                    var html="";
                    for (var i in list){
                        var tr="<tr class='a2' style='background:#CCCCCC'>"+"<td>&nbsp;&nbsp;"+list[i].v1+"</td>"+"<td class='bb' ng-click='bb($event)' myAttr='1' myID="+list[i].v1+" style='cursor:pointer;color:#990000;'>"+list[i].v2+"</td>"+"<td>"+list[i].v3+"</td>"+"<td>"+list[i].Total_effective_bets+"</td><td>"+list[i].payout+"</td><td>"+list[i].JACK+"</td><td>"+list[i].Tip+"</td><td>"+list[i].meber+"</td><td>"+list[i].Agent+"</td><td>"+list[i].generation+"</td><td>"+list[i].Shareholder+"</td><td>"+list[i].company+"</td><td>"+list[i].profit_and_loss+"</td><td>"+list[i].company+"</td></tr>";
                        html+=tr;
                    }
                    $(target).parent().after(html);
                    $(target).parent().attr("data-num",list.length);
                    $(".bb").click(function(event){
                        var target=event.target;
                        var flag=$(target).attr('myAttr')==1?2:1;
                        $(target).attr('myAttr',flag);
                        if(flag==2) {
                            var iD = $(target).attr('myID');
                            financeReportService.generalGenerationReport(iD).then(function (res) {
                                $scope.list2=res.data;
                                var list = $scope.list2;
                                var html = "";
                                for (var i in list) {
                                    var tr = "<tr class='a3' style='background:#CCCCFF'>" + "<td>&nbsp;&nbsp;&nbsp;&nbsp;" + list[i].v1 + "</td>" + "<td class='cc' myAttr='1' myid="+ list[i].v1 +"  style='cursor:pointer;color:#CC0000;' >" + list[i].v2 + "</td>" + "<td>" + list[i].v3 + "</td>"+"<td>"+list[i].Total_effective_bets+"</td><td>"+list[i].payout+"</td><td>"+list[i].JACK+"</td><td>"+list[i].Tip+"</td><td>"+list[i].meber+"</td><td>"+list[i].Agent+"</td><td>"+list[i].generation+"</td><td>"+list[i].Shareholder+"</td><td>"+list[i].company+"</td><td>"+list[i].profit_and_loss+"</td><td>"+list[i].company+"</td></tr>";
                                    html += tr;
                                };
                                $(target).parent().after(html);
                                $(target).parent().attr("data-num",list.length);
                                $(".cc").click(function(event){
                                    var target=event.target;
                                    var flag=$(target).attr('myAttr')==1?2:1;

                                    $(target).attr('myAttr',flag);
                                    if(flag==2) {
                                        var id = $(target).attr('myid');
                                        financeReportService.proxyReport(id).then(function (res) {
                                            $scope.list3 = res.data;
                                            var list = $scope.list3;
                                            var html = "";
                                            for (var i in list) {
                                                var tr = "<tr style='background: #CCCC99'>" + "<td>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;" + list[i].v1 + "</td>" + "<td class='c1c' myAttr='1' myids=" + list[i].v1 + " style='cursor:pointer;color:#FF0000;'>" + list[i].v2 + "</td>" + "<td>" + list[i].v3 + "</td>"+"<td>"+list[i].Total_effective_bets+"</td><td>"+list[i].payout+"</td><td>"+list[i].JACK+"</td><td>"+list[i].Tip+"</td><td>"+list[i].meber+"</td><td>"+list[i].Agent+"</td><td>"+list[i].generation+"</td><td>"+list[i].Shareholder+"</td><td>"+list[i].company+"</td><td>"+list[i].profit_and_loss+"</td><td>"+list[i].company+"</td></tr>";
                                                html += tr;
                                            };
                                            $(target).parent().after(html);
                                            $(target).parent().attr("data-num",list.length);
                                            $('.c1c').click(function ($event) {
                                                var target = $event.target;
                                                var id = $(target).attr("myids");
                                                $state.go('app.ReportForm.Memberdetails',{
                                                    id:id
                                                });
                                            });
                                        });
                                    }else{
                                        var num=$(target).parent().attr('data-num');
                                        $(target).parent().nextUntil(".a3").remove();
                                    }
                                });

                            });

                        }else{
                            var num=$(target).parent().attr('data-num');
                            $(target).parent().nextUntil(".a2").remove();
                        }

                    });
                });
            }else{
                var num=$(target).parent().attr('data-num');
                $(target).parent().nextUntil(".a1").remove();
            }
        };




    });
