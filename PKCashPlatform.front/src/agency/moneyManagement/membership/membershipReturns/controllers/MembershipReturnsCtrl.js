angular.module('app.membershipReturns').controller('MembershipReturnsCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService){
    $scope.sitId = function (site_index_id) {
        MembershipReturnsService.getSiteSelect(site_index_id).then(function (response) {
            console.log(response);
            $scope.names = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        $scope.sitId(user.site_index_id);
    }
    $scope.json = APP_CONFIG.option;

//昨天
    $scope.yesTerDay = function(){
        var day1 = new Date();
        day1.setTime(day1.getTime()-24*60*60*1000);
        var s1 = day1.getFullYear()+"-" + (day1.getMonth()+1) + "-" + day1.getDate();
        console.log(s1);
        $scope.startTime = s1;
        $scope.endTime = s1;

    };
    //获得本周的开始日期
    $scope.thisWeek = function(n){
        var now=new Date();
        var year=now.getFullYear();
        //因为月份是从0开始的,所以获取这个月的月份数要加1才行
        var month=now.getMonth()+1;
        var date=now.getDate();
        var s1=year+"-"+(month<10?('0'+month):month)+"-"+(date<10?('0'+date):date);
        var day=now.getDay();
        console.log(date);
        console.log(day);
        //判断是否为周日,如果不是的话,就让今天的day-1(例如星期二就是2-1)
        if(day!==0){
            n = n+(day-1);
        }
        else{
            n = n+day;
        }
        if(day){
        //这个判断是为了解决跨年的问题
            if(month>1){
                month=month;
            }
        //这个判断是为了解决跨年的问题,月份是从0开始的
            else{
                year=year-1;
                month=12;
            }
        }
        now.setDate(now.getDate()-n);
        year=now.getFullYear();
        month=now.getMonth()+1;
        date=now.getDate();
        console.log(date);
        console.log(n);
        var s=year+"-"+(month<10?('0'+month):month)+"-"+(date<10?('0'+date):date);
        console.log(s);
        $scope.startTime = s;
        $scope.endTime = s1;
    };
    var nowDate= new Date(); //当天日期
    var nowDayOfWeek= nowDate.getDay();
    var nowDay = nowDate.getDate();
    var nowMonth = nowDate.getMonth();
    var nowYear = nowDate.getFullYear();
    nowYear += (nowYear < 2000) ? 1900 : 0; //
    var lastMonthDate = new Date(); //上月日期
    lastMonthDate.setDate(1);
    lastMonthDate.setMonth(lastMonthDate.getMonth()-1);
    var lastYear = lastMonthDate.getYear();
    var lastMonth = lastMonthDate.getMonth();

    $scope.formatDate = function(date){
        var myyear = date.getFullYear();
        var mymonth = date.getMonth()+1;
        var myweekday = date.getDate();

        if(mymonth < 10){
            mymonth = "0" + mymonth;
        }
        if(myweekday < 10){
            myweekday = "0" + myweekday;
        }

        return (myyear+"-"+mymonth + "-" + myweekday);
    };
    $scope.lastWeek = function(){
        //获得上周的开始日期
        var getUpWeekStartDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek -6);
        var getUpWeekStartDate =  $scope.formatDate(getUpWeekStartDate);
        $scope.startLastWeek = getUpWeekStartDate;
        console.log($scope.startLastWeek);
        $scope.startTime = $scope.startLastWeek;
        //获得上周的结束日期
        var getUpWeekEndDate = new Date(nowYear, nowMonth, nowDay + (6 - nowDayOfWeek - 6));
        var getUpWeekEndDate = $scope.formatDate(getUpWeekEndDate);
        $scope.endLastWeek = getUpWeekEndDate;
        console.log($scope.endLastWeek);
        $scope.endTime = $scope.endLastWeek;
    };
    $scope.check = function () {
        if($scope.site_index_id==undefined||$scope.startTime==undefined||$scope.endTime==undefined||$scope.condition==undefined){
            popupSvc.smallBox("fail","请输入完整！")
        }else{
            $state.go('app.MembershipReturns.TranslationResult',{
                site:$scope.site_index_id,
                startTime:$scope.startTime,
                endTime:$scope.endTime,
                condition:$scope.condition
            })
        }



    }
});