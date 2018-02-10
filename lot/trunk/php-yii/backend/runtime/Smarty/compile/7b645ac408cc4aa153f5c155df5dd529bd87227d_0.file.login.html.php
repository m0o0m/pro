<?php
/* Smarty version 3.1.31, created on 2018-02-10 10:59:50
  from "/Users/frank/www/newproject/trunk/php-yii/backend/views/login/login.html" */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '3.1.31',
  'unifunc' => 'content_5a7ed0a681d8e7_36682396',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    '7b645ac408cc4aa153f5c155df5dd529bd87227d' => 
    array (
      0 => '/Users/frank/www/newproject/trunk/php-yii/backend/views/login/login.html',
      1 => 1517801681,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
  ),
),false)) {
function content_5a7ed0a681d8e7_36682396 (Smarty_Internal_Template $_smarty_tpl) {
?>

<style>
    #vcode{max-width:100px;float: left}
    #captchaimg{cursor:pointer; height: 32px;margin-top:1px;margin-left:10px;width:100px}
    .error{color:red}
</style>
<div class="position-relative">
    <div id="login-box" class="login-box visible widget-box no-border">
        <div class="widget-body">
            <div class="widget-main">
                <h4 class="header blue bigger">
                    请输入您的登录信息
                </h4>
                <div class="space-6"></div>
                <form id="login-form"  role="form">
                    <fieldset>
                        <label class="block clearfix">
                            <span class="block input-icon input-icon-right">
                                <div class="form-group ">
                                    <input type="text" id="username" class="form-control" name="username" placeholder="管理员账号">                                   
                                </div>           
                            </span>
                        </label>
                        <label class="block clearfix">
                            <span class="block input-icon input-icon-right">
                                <div class="form-group ">
                                    <input type="password" id="password" class="form-control" name="password" placeholder="管理员密码">                                   
                                </div>           
                            </span>
                        </label>
                        <label class="block clearfix">
                            <span class="block input-icon input-icon-right">
                                <div class="form-group">
                                    <input type="text" id="vcode" class="form-control" name="vcode" placeholder="验证码">
                                    <?php echo yii\captcha\Captcha::widget(array('name'=>'captchaimg','captchaAction'=>'captcha','imageOptions'=>array('id'=>'captchaimg','title'=>'换一个','alt'=>'换一个'),'template'=>'{image}'));?>

                                </div>           
                            </span>
                        </label>
                        <div class="clearfix">
                            <button id='login-submit' type="button" class="btn bg-olive btn-block width-100 pull-right btn btn-sm btn-primary">登录</button>    </div>
                        <div class="space-4"></div>
                    </fieldset>
                </form>                               
            </div>
        </div>
    </div>
</div>
<?php echo '<script'; ?>
>
    var submit = true;
    //点击刷新验证码
    $("#captchaimg").click(function () {
        refresh();
    })

    //验证码刷新
    function refresh() {
        $.ajax({
            //使用ajax请求login/captcha方法，加上refresh参数，接口返回json数据
            url: "/login/captcha?refresh",
            dataType: 'json',
            cache: false,
            success: function (data) {
                $("#captchaimg").attr('src', data['url']);
            }
        })
    }

    //表单提交
    $('#login-submit').click(function () {
        var logintype = $.trim($('#logintype').val());
        var uname = $.trim($('#username').val());
        var pwd = $.trim($('#password').val());
        var vcode = $.trim($('#vcode').val());
        var rem = $('#remember').is(':checked') === true ? 1 : 0;
        var errorState = false;
        var errorMsg = '';
        var regUname = /^[A-Za-z0-9_]*$/g;
        var regPwd = /^[A-Za-z0-9_]*$/g;
        //验证
        if (uname == '') {
            errorState = true;
            errorMsg = '用户名不能为空!';
        } else if (uname.length > 20 || uname.length < 4) {
            errorState = true;
            errorMsg = '用户名长度只能为4-20位!';
        } else if (!regUname.test(uname)) {
            errorState = true;
            errorMsg = '用户名只能为数字字母下划线!';
        } else if (pwd == '') {
            errorState = true;
            errorMsg = '密码不能为空!';
        } else if (pwd.length > 20 || pwd.length < 6) {
            errorState = true;
            errorMsg = '密码长度只能为6-20位!';
        } else if (!regPwd.test(pwd)) {
            errorState = true;
            errorMsg = '密码只能为数字字母下划线!';
        } else if (vcode == '') {
            errorState = true;
            errorMsg = '验证码不能为空!';
        } else if (vcode.length != 4) {
            errorState = true;
            errorMsg = '请输入4位验证码!';
        }
        //展示错误信息
        if (errorState) {
            $("#login-box").find('.bigger').addClass('red');
            $("#login-box").find('.bigger').removeClass('blue');
            $("#login-box").find('.bigger').text(errorMsg);
            return false;
        } else {
            $("#login-box").find('.bigger').text('正在登录,请稍等!');
        }

        //提交
        if (submit) {
            submit = false;
            $.ajax({
                url: "/login/logindo",
                method: 'post',
                dataType: 'json',
                data: {
                    'logintype': logintype,
                    "uname": uname,
                    "pwd": pwd,
                    "vcode": vcode,
                    "rem": rem
                },
                success: function (data) {
                    if (data.code == 400) {
                        submit = true;
                        $("#login-box").find('.bigger').text(data.msg);
                    } else if (data.code == 300) {
                        submit = true;
                        $("#login-box").find('.bigger').text(data.msg);
                        refresh();
                    } else if (data.code == 200) {
                        window.location.href = '/admin/center';
                    }
                },
                error: function () {
                    submit = true;
                    $("#login-box").find('.bigger').text('网络异常,请稍后再试!');
                    refresh();
                }
            })
        }
    })

    $(document).keyup(function (event) {
        if (event.keyCode == 13) {
            $("#login-submit").trigger("click");
        }
    });
<?php echo '</script'; ?>
>


<?php }
}
