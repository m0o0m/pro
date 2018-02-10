<?php
/* Smarty version 3.1.31, created on 2018-02-10 11:00:02
  from "/Users/frank/www/newproject/trunk/php-yii/backend/views/admin/center.html" */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '3.1.31',
  'unifunc' => 'content_5a7ed0b2678e90_12274114',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    'b96bc72b91625528932892fde2f13ea28097b835' => 
    array (
      0 => '/Users/frank/www/newproject/trunk/php-yii/backend/views/admin/center.html',
      1 => 1517801681,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
  ),
),false)) {
function content_5a7ed0b2678e90_12274114 (Smarty_Internal_Template $_smarty_tpl) {
?>

<div class="row">
    <div class="col-xs-12">
        <div class="row">
            <div class="col-xs-12 col-sm-12">
                <h4 class="blue">
                    <span class="middle"><i class="ace-icon glyphicon glyphicon-user light-blue bigger-110"></i></span>
                    账号信息
                </h4>
                <div class="profile-user-info">
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 账号  </div>
                        <div class="profile-info-value">
                            <span><?php if (isset($_smarty_tpl->tpl_vars['data']->value['login_user'])) {
echo $_smarty_tpl->tpl_vars['data']->value['login_user'];
}?></span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 昵称  </div>
                        <div class="profile-info-value">
                            <span><?php if (isset($_smarty_tpl->tpl_vars['data']->value['login_name'])) {
echo $_smarty_tpl->tpl_vars['data']->value['login_name'];
}?></span>
                        </div>
                    </div>
                    <?php if (isset($_smarty_tpl->tpl_vars['data']->value['user_type']) && $_smarty_tpl->tpl_vars['data']->value['user_type'] == 2) {?>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 角色  </div>
                        <div class="profile-info-value">
                            <span>
                                <?php if (isset($_smarty_tpl->tpl_vars['data']->value['role_name']) && !empty($_smarty_tpl->tpl_vars['data']->value['role_name'])) {?>
                                <?php echo $_smarty_tpl->tpl_vars['data']->value['role_name'];?>

                                <?php } else { ?>
                                暂无角色
                                <?php }?>
                            </span>
                        </div>
                    </div>
                    <?php }?>
                    <div class="profile-info-row">
                        <div class="profile-info-name"> 帐号类型  </div>
                        <div class="profile-info-value">
                            <span><?php ob_start();
echo $_smarty_tpl->tpl_vars['data']->value['user_type_txt'];
$_prefixVariable1=ob_get_clean();
if (isset($_prefixVariable1)) {
echo $_smarty_tpl->tpl_vars['data']->value['user_type_txt'];
}?></span>
                        </div>
                    </div>
                    <div class="profile-info-row">
                        <div class="profile-info-name">
                            余额
                        </div>
                        <div class="profile-info-value">
                            <span><?php ob_start();
echo $_smarty_tpl->tpl_vars['data']->value['money'];
$_prefixVariable2=ob_get_clean();
if (isset($_prefixVariable2)) {
echo $_smarty_tpl->tpl_vars['data']->value['money'];
}?></span>
                        </div>
                    </div>
                </div>
                <div class="hr hr16 dotted"></div>
            </div>
        </div>                  
    </div>
</div>
<?php }
}
