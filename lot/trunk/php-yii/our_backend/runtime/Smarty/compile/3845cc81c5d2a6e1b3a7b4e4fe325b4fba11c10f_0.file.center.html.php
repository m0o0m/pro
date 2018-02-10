<?php
/* Smarty version 3.1.31, created on 2018-02-10 10:57:06
  from "/Users/frank/www/newproject/trunk/php-yii/our_backend/views/admin/center.html" */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '3.1.31',
  'unifunc' => 'content_5a7ed0029441a1_36292391',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    '3845cc81c5d2a6e1b3a7b4e4fe325b4fba11c10f' => 
    array (
      0 => '/Users/frank/www/newproject/trunk/php-yii/our_backend/views/admin/center.html',
      1 => 1517801693,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
  ),
),false)) {
function content_5a7ed0029441a1_36292391 (Smarty_Internal_Template $_smarty_tpl) {
?>

<div class="row">
    <div class="col-xs-12">
        <div class="row">
            <div class="col-xs-12 col-sm-12">
                <h4 class="blue">
                    <span class="middle"><i class="ace-icon glyphicon glyphicon-user light-blue bigger-110"></i></span>
                    账号信息
                </h4>
                <div class="profile-user-info" style="margin-left:50px;">
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 账号  </div>
                        <div class="profile-info-value">
                            <span> <?php if (isset($_smarty_tpl->tpl_vars['data']->value['login_user'])) {?> <?php echo $_smarty_tpl->tpl_vars['data']->value['login_user'];?>
 <?php }?> </span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 昵称  </div>
                        <div class="profile-info-value">
                            <span> <?php if (isset($_smarty_tpl->tpl_vars['data']->value['login_name'])) {?> <?php echo $_smarty_tpl->tpl_vars['data']->value['login_name'];?>
 <?php }?> </span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 角色  </div>
                        <div class="profile-info-value">
                            <span> <?php if (isset($_smarty_tpl->tpl_vars['data']->value['role_name'])) {?> <?php echo $_smarty_tpl->tpl_vars['data']->value['role_name'];?>
 <?php }?> </span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 上次登录时间  </div>
                        <div class="profile-info-value">
                            <span> <?php if (isset($_smarty_tpl->tpl_vars['data']->value['adddate'])) {?> <?php echo $_smarty_tpl->tpl_vars['data']->value['adddate'];?>
 <?php }?> </span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 上次登录IP  </div>
                        <div class="profile-info-value">
                            <span> <?php if (isset($_smarty_tpl->tpl_vars['data']->value['ip'])) {?> <?php echo $_smarty_tpl->tpl_vars['data']->value['ip'];?>
 <?php }?> </span>
                        </div>
                    </div>
                </div>
                <div class="hr hr16 dotted"></div>

<!--                 <h4 class="blue">
                    <span class="middle"><i class="fa fa-desktop light-blue bigger-110"></i></span>
                    系统信息
                </h4>

                <div class="profile-user-info">
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 操作系统  </div>
                        <div class="profile-info-value">
                            <span>Darwin&nbsp;16.5.0</span>
                        </div>
                    </div>

                    <div class="profile-info-row">
                        <div class="profile-info-name"> 服务器软件 </div>

                        <div class="profile-info-value">
                            <span>nginx/1.10.3</span>
                        </div>
                    </div>

                    <div class="profile-info-row">
                        <div class="profile-info-name"> MySQL版本 </div>

                        <div class="profile-info-value">
                            <span>MySQL 5.6.35</span>
                        </div>
                    </div>

                    <div class="profile-info-row">
                        <div class="profile-info-name"> PHP版本 </div>

                        <div class="profile-info-value">
                            <span>PHP 5.5.38</span>
                        </div>
                    </div>

                    <div class="profile-info-row">
                        <div class="profile-info-name"> Yii版本 </div>
                        <div class="profile-info-value">
                            <span>Yii 2.0.9</span>
                        </div>
                    </div>

                    <div class="profile-info-row">
                        <div class="profile-info-name"> 上传文件 </div>
                        <div class="profile-info-value">
                            <span>2M</span>
                        </div>
                    </div>
                </div>
                <div class="hr hr-8 dotted"></div>
                <div class="profile-user-info">
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 个人主页 </div>
                        <div class="profile-info-value">
                            <a target="_blank" href="http://821901008.qzone.qq.com">http://821901008.qzone.qq.com</a>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name">
                            <i class="fa fa-github-square" aria-hidden="true"></i>
                            GitHub
                        </div>
                        <div class="profile-info-value">
                            <a href="https://github.com/myloveGy" target="_blank">https://github.com/myloveGy</a>
                        </div>
                    </div>
                </div> -->
                <div class="hr hr16 dotted"></div>
            </div>
        </div>                  
    </div>
</div>
<?php }
}
