$(function () {
//main.html
    var main = {
        load: '',
        init: function () {
            this.menu();
            this.pwdForm();
            this.updatePwd();
            this.clock();
            this.renderPjax();
        },
        menu: function () {
            $('ul.submenu').find('li.active').parents('li.oneLevel').addClass('active');
        },
        clock: function () {
            var time = $('input[name="nowtime"]').val();
            clock(time);
        },
        renderPjax: function () {
            var that = this;
            $('.pjax-href').click(function () {
                $(this).parents('li.oneLevel').siblings('li.oneLevel').removeClass('active');
                $(this).parents('li.oneLevel').siblings('li.oneLevel').removeClass('open');
                $(this).parents('li.oneLevel').siblings('li.oneLevel').find('ul.submenu').hide();
                $('ul.submenu').find('li.active').removeClass('active');
                $(this).parents('li').addClass('active');
                var url = $(this).attr('data');
                $.pjax({
                    method: 'get',
                    url: url,
                    container: '#container',
                    timeout: 10000
                });
            });

            $(document).on('pjax:start', function () {
                that.load = layer.load(1, {
                    shade: [0.5, '#000000'] // loading...
                });
            });
            $(document).on('pjax:end', function () {
                layer.close(that.load);
            });
            $(document).on('pjax:error', function () {
                layer.close(that.load);
            })
        },
        pwdForm: function () {
            $('#pwdForm').click(function () {
                layer.open({
                    type: 1 //Page层类型
                    , area: ['604px', '330px']
                    , title: '修改密码'
                    , shade: 0.5 //遮罩透明度
                    , maxmin: true //允许全屏最小化
                    , anim: 1 //0-6的动画形式，-1不开启
                    , content: $('#pwd-layer')
                });
            });
        },
        updatePwd: function () {
            $('#updatePwd').click(function () {
                var index;
                var errorState = false;
                var errorMsg = '';
                var oldPwd = $.trim($('#oldPwd').val());
                var newPwd = $.trim($('#newPwd').val());
                var confirmPwd = $.trim($('#confirmPwd').val());
                if (oldPwd == '') {
                    errorState = true;
                    errorMsg = '原始密码不能为空!';
                } else if (!/^[A-Za-z0-9_]*$/g.test(oldPwd)) {
                    errorState = true;
                    errorMsg = '原始密码只能为数字字母下划线!';
                } else if (oldPwd.length < 6 || oldPwd.length > 20) {
                    errorState = true;
                    errorMsg = '原始密码长度只能为6-20位!';
                } else if (newPwd == '') {
                    errorState = true;
                    errorMsg = '新密码不能为空!';
                } else if (!/^[A-Za-z0-9_]*$/g.test(newPwd)) {
                    errorState = true;
                    errorMsg = '新密码只能为数字字母下划线!';
                } else if (newPwd.length < 6 || newPwd.length > 20) {
                    errorState = true;
                    errorMsg = '新密码长度只能为6-20位!';
                } else if (confirmPwd == '') {
                    errorState = true;
                    errorMsg = '确认密码不能为空!';
                } else if (!/^[A-Za-z0-9_]*$/g.test(confirmPwd)) {
                    errorState = true;
                    errorMsg = '确认密码只能为数字字母下划线!';
                } else if (confirmPwd.length < 6 || confirmPwd.length > 20) {
                    errorState = true;
                    errorMsg = '确认密码长度只能为6-20位!';
                } else if (newPwd != confirmPwd) {
                    errorState = true;
                    errorMsg = '新密码与确认密码不一致!';
                } else {
                    errorState = false;
                    errorMsg = '';
                }

                if (errorState) {
                    layer.alert(errorMsg, {icon: 2});
                    return false;
                }

                $.ajax({
                    type: "post",
                    url: '/admin/updatepwd',
                    data: {
                        "oldPwd": oldPwd,
                        "newPwd": newPwd,
                        "confirmPwd": confirmPwd
                    },
                    beforeSend: function () {
                        index = layer.load(1, {
                            shade: [0.5, '#000000'] //0.1透明度的白色背景
                        });
                    },
                    error: function () {
                        layer.alert('出错啦！', {icon: 2});
                        layer.close(index);
                    },
                    dataType: 'json',
                    success: function (res) {
                        if (res.code == 400) {
                            layer.alert(res.msg, {icon: 2});
                            layer.close(index);
                        } else if (res.code == 200) {
                            layer.alert(res.msg, {icon: 2});
                            window.location.href = '/login/loginout';
                        }

                    }
                });

            })

        }
    }

    //初始
    main.init();
})