
angular.module('app.Platform').controller('DomainCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,$stateParams){
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        // 删除
        $scope.delete = function (id, site_id, site_index_id) {
            var del = function () {
                httpSvc.del("/site/domain_id",{
                    site: site_id,
                    site_index: site_index_id,
                    id: id
                }).then(function (response) {
                    console.log(typeof (response));
                    if(JSON.parse(response)===null){
                        popupSvc.smallBox("success","删除成功");
                        GetAllEmployee();
                    }else{
                        popupSvc.smallBox("fail", response.msg);
                    }
                })
            }
            popupSvc.smartMessageBox("确定删除？",del);
        }

        var csrExt = new Array(".crt");//.csr文件的后缀名
        var keyExt = new Array(".key");//.key文件的后缀名

        //获取文件名后缀名
        String.prototype.extension = function(){
            var ext = null;
            var name = this.toLowerCase();
            var i = name.lastIndexOf(".");
            if(i > -1){
                var ext = name.substring(i);
            }
            return ext;
        }
        //判断Array中是否包含某个值
        Array.prototype.contain = function(obj){
            for(var i=0; i<this.length; i++){
                if(this[i] === obj)
                    return true;
            }
            return false;
        };

        function typeMatch(type, filename){
            var ext = filename.extension();
            if(type.contain(ext)){
                return true;
            }
            return false;
        }

        //新增
        $scope.add = function () {
            var fd = new FormData();
            var file1 =$("#exampleInputFile1").get(0).files[0];
            var file2 =$("#exampleInputFile2").get(0).files[0];
            if(file1&&!typeMatch(csrExt, file1.name)){
                popupSvc.smallBox("fail","文件格式有误");
            }else{
                if(file2&&!typeMatch(keyExt, file2.name)){
                    popupSvc.smallBox("fail","文件格式有误");
                }
            }

            fd.append('ssl_key_csr', file1);
            fd.append('ssl_key_file', file2);
            fd.append('site', $stateParams.id);
            fd.append('site_index', $stateParams.site);
            fd.append('pc_domain', $scope.pc_domain);
            fd.append('wap_domain', $scope.wap_domain);
            httpSvc.file("/site/domain", fd).then(function (response) {
                if(!response.code){
                    popupSvc.smallBox("success","添加成功");
                    GetAllEmployee();
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        }

        $('#exampleInputFile3').change(function(){
            $scope.modifyData.ssl_key_csr=$("#exampleInputFile3").get(0).files[0].name;
        });
        $('#exampleInputFile4').change(function(){
            $scope.modifyData.ssl_key_file=$("#exampleInputFile4").get(0).files[0].name;
        });

        //修改
        $scope.modify=function (id, site_id, site_index_id) {
            $scope.modifyData={};
            httpSvc.get("/site/domain_id",{
                site: site_id,
                site_index: site_index_id,
                id: id
            }).then(function (response) {
                if (!response.code) {
                    $scope.modifyData = response.data;
                }
            })

        }
        
        //提交修改
        $scope.modifySubmit=function () {
            var fd = new FormData();
            var file1 =$("#exampleInputFile1").get(0).files[0];
            var file2 =$("#exampleInputFile2").get(0).files[0];
            var file3 =$("#exampleInputFile3").get(0).files[0];
            var file4 =$("#exampleInputFile4").get(0).files[0];
            if(file1&&!typeMatch(csrExt, file3.name)){
                popupSvc.smallBox("fail","文件格式有误");
            }else{
                if(file2&&!typeMatch(keyExt, file4.name)){
                    popupSvc.smallBox("fail","文件格式有误");
                }
            }
            var is_change=!file3||!file4?2:1;
            fd.append('ssl_key_csr', file3);
            fd.append('ssl_key_file', file4);
            fd.append('is_change', is_change);
            fd.append('site', $scope.modifyData.site_id);
            fd.append('id', $scope.modifyData.id);
            fd.append('site_index', $scope.modifyData.site_index_id);
            fd.append('pc_domain', $scope.modifyData.pc_domain);
            fd.append('wap_domain', $scope.modifyData.wap_domain);
            httpSvc.file("/site/domain_id", fd).then(function (response) {
                if(!response.code){
                    popupSvc.smallBox("success","修改成功");
                    GetAllEmployee();
                }else{
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        }




        //点击搜索
        $scope.search = function () {
            GetAllEmployee();
        }


        var GetAllEmployee = function () {

            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site: $stateParams.id,
                site_index: $stateParams.site,
                domain: $scope.domain,
                get_count: '',
                desc: '',
                order_by: '',
            }

            httpSvc.get("/site/domain", postData).then(function (response) {
                if (!response.code) {
                    $scope.paginationConf.totalItems = response.meta.count;
                    $scope.list = response.data;
                }
            })

        }

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);



    });
