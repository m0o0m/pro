$(function(){

	var modal = {},
		titStyle,
		PanBg,
		canStyle,
		modaConfig,
		info = {
			creator:"朱盛文",
			version:"1.0",
			updateTime:"2016-5-16",
			rely:"jQuery"
		},
		loginConfig,
		showConfig;

	window.zhuModal = modal;
	// 初始化
	modal.init = function(){
		
		if(!!arguments[0]){
			modaConfig = arguments[0];
			if(!!arguments[0].loginConfig)loginConfig = arguments[0].loginConfig;
		}

	}

	// 显示登录弹框
	modal.login = function(type){
		var initData = createModal(modaConfig);

		var modalTit = "登&nbsp;&nbsp;录",
		userNameTit = "用户名",
		passwordTit = "密&nbsp;&nbsp;码",
		codeTit = "验证码",
		userNameID = "zhu-modalUserName",
		passwordID = "zhu-modalPassword",
		codeID = "zhu-modalCode",
		userNamePlaceholder = "请输入用户名",
		passwordPlaceholder = "请输入密码",
		codePlaceholder = "请输入验证码",
		enterBtnText = modalTit,
		registHref = "javascript:;",
		DemoHref = "javascript:;",
		DemoStytle = "",
		registBtnText = "免费注册账号",
		codesImg,
		codesClick,
		enterBtnClick,
		enterBtnID;

		if(!!loginConfig){
			if(loginConfig.modalTit)modalTit = loginConfig.modalTit;
			if(loginConfig.userNameTit)userNameTit = loginConfig.userNameTit;
			if(loginConfig.passwordTit)passwordTit = loginConfig.passwordTit;
			if(loginConfig.codeTit)codeTit = loginConfig.codeTit;
			if(loginConfig.userNameID)userNameID = loginConfig.userNameID;
			if(loginConfig.passwordID)passwordID = loginConfig.passwordID;
			if(loginConfig.codeID)codeID = loginConfig.codeID;
			if(loginConfig.userNamePlaceholder)userNamePlaceholder = loginConfig.userNamePlaceholder;
			if(loginConfig.passwordPlaceholder)passwordPlaceholder = loginConfig.passwordPlaceholder;
			if(loginConfig.codePlaceholder)codePlaceholder = loginConfig.codePlaceholder;
			if(loginConfig.enterBtnText)enterBtnText = loginConfig.enterBtnText;
			if(loginConfig.registHref)registHref = loginConfig.registHref;
			if(loginConfig.DemoHref)DemoHref = loginConfig.DemoHref;
			if(loginConfig.DemoStytle)DemoStytle = loginConfig.DemoStytle;
			if(loginConfig.enterBtnText)enterBtnText = loginConfig.enterBtnText;
			if(loginConfig.codesImg)codesImg = loginConfig.codesImg;
			if(loginConfig.codesClick)codesClick = loginConfig.codesClick;
			if(loginConfig.enterBtnClick)enterBtnClick = loginConfig.enterBtnClick;
			if(loginConfig.enterBtnID)enterBtnID = loginConfig.enterBtnID;
		}

		if(!!arguments[0]){
			if(arguments[0].modalTit)modalTit = arguments[0].modalTit;
			if(arguments[0].userNameTit)userNameTit = arguments[0].userNameTit;
			if(arguments[0].passwordTit)passwordTit = arguments[0].passwordTit;
			if(arguments[0].codeTit)codeTit = arguments[0].codeTit;
			if(arguments[0].userNameID)userNameID = arguments[0].userNameID;
			if(arguments[0].passwordID)passwordID = arguments[0].passwordID;
			if(arguments[0].codeID)codeID = arguments[0].codeID;
			if(arguments[0].userNamePlaceholder)userNamePlaceholder = arguments[0].userNamePlaceholder;
			if(arguments[0].passwordPlaceholder)passwordPlaceholder = arguments[0].passwordPlaceholder;
			if(arguments[0].codePlaceholder)codePlaceholder = arguments[0].codePlaceholder;
			if(arguments[0].enterBtnText)enterBtnText = arguments[0].enterBtnText;
			if(arguments[0].registHref)registHref = arguments[0].registHref;
			if(arguments[0].DemoHref)DemoHref = arguments[0].DemoHref;
			if(arguments[0].DemoStytle)DemoStytle = arguments[0].DemoStytle;
			if(arguments[0].enterBtnText)enterBtnText = arguments[0].enterBtnText;
			if(arguments[0].codesImg)codesImg = arguments[0].codesImg;
			if(arguments[0].codesClick)codesClick = arguments[0].codesClick;
			if(arguments[0].enterBtnClick)enterBtnClick = arguments[0].enterBtnClick;
			if(arguments[0].enterBtnID)enterBtnID = arguments[0].enterBtnID;
		}

		var loginBox = $("<div class='login-box'></div>");
		var userName = $("<label><span>"+userNameTit+" : </span><input type='text' id="+userNameID+" placeholder="+userNamePlaceholder+" /></label>");
		var password = $("<label><span>"+passwordTit+" : </span><input type='password' id="+passwordID+" placeholder="+passwordPlaceholder+" /></label>");
		var codesInput = $("<label><span>"+codeTit+" : </span><input type='text' class='zhu-inputW40' id="+codeID+" placeholder="+codePlaceholder+" /><a class='zhu-codeImg' href='javascript:;' title='看不清?点击切换!'></a></label>");
		var regist = $("<label><span>&nbsp; </span><a href='javascript:;' onclick='"+registHref+"' class='zhu-registLink' >"+registBtnText+"</a></label>");
		var demo = $("<a href='javascript:;' class='zhu-enterBtn frt' onclick='"+DemoHref+"' style='"+DemoStytle+"'>试&nbsp;&nbsp;玩</a>");
		var codeBtn = codesInput.find(".zhu-codeImg");

		loginBox.append(userName,password,codesInput,regist);
		initData.modalBody.append(loginBox);
		initData.modalPanel.addClass("login");
		initData.modalEnterBtn.html(enterBtnText);
		initData.modalHead.html(modalTit);
		if (type == 'sw'){
			initData.modalFoot.append(initData.modalEnterBtn,initData.modalCancelBtn,demo, "<div class='clr'></div>");
		}else{
			initData.modalFoot.append(initData.modalEnterBtn,initData.modalCancelBtn, "<div class='clr'></div>");
		}
		if(codesImg)codeBtn.append(codesImg);
		if(codesClick)codeBtn.click(codesClick);		
		if(enterBtnClick){
			initData.modalEnterBtn.unbind("click");
			initData.modalEnterBtn.click(enterBtnClick);
		}
		if(enterBtnID)initData.modalEnterBtn.attr("id",enterBtnID);

		showModule(initData);
		
		
	};

	// 显示帮助文档
	modal.login_help = function(){
		var helpTxt = 	"zhuModal.login()不传参为默认登录弹框,传递对象可实现自定义input标签id,提示文字,输入框标题等信息,对象属性如下(属性无排序,不设置则为默认值):\n\n"+
						"modalTit:弹框标题名称,字符串类型\n\nuserNameTit:用户名输入框前标题名可按需求改成'账号'等,字符串类型\npasswordTit:同上,密码输入框标题,字符串类型\ncodeTit:同上,验证码,字符串类型\n\n"+
						"userNameID:用户名输入框id,字符串类型\npasswordID:密码id,字符串类型\ncodeID:验证码id,,字符串类型\r\n"+
						"userNamePlaceholder:用户名输入框提示语,字符串类型\npasswordPlaceholder:密码,字符串类型\ncodePlaceholder:验证码,字符串类型\n\n"+
						"egistHref:注册链接地址,字符串类型\n\n"+
						"enterBtnText:登录按钮文字,字符串类型\registBtnText:注册链接文字,字符串类型\n\n"+
						"codesImg:验证码图片jQuery对象,\ncodesClick:验证码点击事件,函数对象,\nenterBtnClick:登录按钮点击事件,函数对象\n\n"+
						"示例,\n"+
						"var loginObj = {\n"+
						"	modalTit:'用户登录',\n"+
						"	userNameTit:'账号/邮箱',\n"+
						"	userNamePlaceholder:'输入6-13位数字或字母组合用户名',\n"+
						"	...\n"+
						"};\n"+
						"zhuModal.login(loginObj);\n\n";
		console.log(helpTxt);
	};
	// 显示帮助文档
	modal.show_help = function(){
		var helpTxt = 	"zhuModal.show()传入字符串直接显示文字信息,第二个参数可传入对象来自定义标题文字,以及按钮文字:\n\n"+
						"modalTit:弹框标题名称,字符串类型\n\n"+
						"enterBtnText:登录按钮文字,字符串类型\n\n"+
						"示例,\n"+
						"var madalObj = {\n"+
						"	modalTit:'温馨提示',\n"+
						"	enterBtnText:'我知道了',\n"+
						"	...\n"+
						"};\n"+
						"zhuModal.show('您的账户余额不足,请充值!',madalObj);";
		console.log(helpTxt);
	};
	modal.init_help = function(){
		var helpTxt = 	"zhuModal.init()传入对象来定义标题背景色,字体颜色,背景色等等,对象属性格式#ffffff,具体如下:\n\n"+
						"mainBg:弹框标题背景色,字符串类型\n\n"+
						"titleCl:弹框标题文字颜色,字符串类型\n\n"+
						"panelBg:弹框主体背景色,字符串类型\n\n"+
						"cancelBg:取消按钮背景色,字符串类型\n\n"+
						"cancelCl:取消按钮文字颜色,字符串类型\n\n"+
						"loginConfig:配置登录弹框的相关信息,对象类型,\nshowConfig:配置登录弹框的相关信息,对象类型,\n\n"+
						"示例,\n"+
						"var madalObj = {\n"+
						"	mainBg:'#000',\n"+
						"	titleCl:'#fff',\n"+
						"	...\n"+
						"};\n"+
						"zhuModal.init(madalObj);\n"+
						"zhuModal.login();\n"+
						"...";
		console.log(helpTxt);
	};
	modal.help = function(){
		var helpTxt = 	"本插件提供登录弹框,文字弹框,以及自定义弹框等功能,可自定义标题栏,确认按钮文字,主体/背景颜色等.\n\n"+
						"zhuModal.init():自定义背景色,字体颜色等信息,详见zhuModal.init_help(),\n\n"+
						"zhuModal.login():登录弹框,详细信息用法请见zhuModal.login_help(),\n\n"+
						"zhuModal.show():消息弹框,详细信息用法请见zhuModal.show_help()";
		console.log(helpTxt);
	}

	// 显示文本或自定义主体信息
	modal.show = function(object){
		var initData = createModal(modaConfig),
			modalTit = "提&nbsp;&nbsp;示",
			enterBtnText = "确&nbsp;&nbsp;定";

		if(!!showConfig){
			if(showConfig.modalTit)modalTit = showConfig.modalTit;
			if(showConfig.enterBtnText)enterBtnText = showConfig.enterBtnText;
		}

		if(!!arguments[1]){
			if(arguments[1].modalTit)modalTit = arguments[1].modalTit;
			if(arguments[1].enterBtnText)enterBtnText = arguments[1].enterBtnText;
		}

		if(Object.prototype.toString.call(object) === "[object String]"){

			initData.modalHead.html(modalTit);
			initData.modalEnterBtn.html(enterBtnText);
			initData.modalFoot.append(initData.modalEnterBtn,"<div class='clr'></div>");
			initData.modalBody.append(initData.modalBodyText);
			initData.modalBodyText.text(object);

			if(arguments[2] == "frame"){
				showModuleCurrent(initData);
			}else{
				showModule(initData);
			}
		}else{
			initData.modalHead.html(modalTit);
			initData.modalEnterBtn.html(enterBtnText);
			initData.modalFoot.append(initData.modalEnterBtn,initData.modalCancelBtn,"<div class='clr'></div>");
			initData.modalBody.append(object);
			
			showModule(initData);
			
		}
	};

	modal.enterClick = function(){
		var modal = $(this).closest(".zhu-modal");
		modal.addClass("zhu-transparent");
		setTimeout(function(){modal.remove();},1000);
	};
	
	modal.cancelClick = function(){
		var modal = $(this).closest(".zhu-modal");
		modal.addClass("zhu-transparent");
		setTimeout(function(){modal.remove();},1000);
	}
	
	modal.info = function(){
		var helpTxt = 	"版本:"+info.version+"\n"+
						"更新时间:"+info.updateTime+"\n"+
						"依赖库:"+info.rely+"\n"+
						"作者:"+info.creator;
		console.log(helpTxt);
	}
	function err_init(){
		console.error("zhuModal未初始化,请用zhuModal.init();");
	}
	
	function createModal(){

		if(arguments[0]){
			if(arguments[0].mainBg || arguments[0].titleCl){
				var bg = !!arguments[0].mainBg ? "background:"+arguments[0].mainBg+";" : "";
				var cl = !!arguments[0].titleCl ? "color:"+arguments[0].titleCl+";" : "";
				titStyle = bg + cl;
			}
			if(arguments[0].panelBg)PanBg = "background:"+arguments[0].panelBg+";";
			if(arguments[0].cancelBg || arguments[0].cancelCl){
				var bg = !!arguments[0].cancelBg ? "background:"+arguments[0].cancelBg+";":"";
				var cl = !!arguments[0].cancelBg ? "color:"+arguments[0].cancelBg+";":"";
				canStyle = bg + cl;
			}
		}

		var modalBg = $("<div class='zhu-modal zhu-transparent zhu-amd'></div>");
		var modalPanel = $("<div class='zhu-modalPanel zhu-amd' style='"+PanBg+"'></div>");
		var modalHead = $("<div class='zhu-modalHead' style='"+titStyle+"'></div>");
		var modalBody = $("<div class='zhu-modalBody'></div>");
		var modalFoot = $("<div class='zhu-modalFoot'></div>");
		var modalEnterBtn = $("<a href='javascript:;' class='zhu-enterBtn frt' style='"+titStyle+"'></a>");
		var modalCancelBtn = $("<a href='javascript:;' class='zhu-cancelBtn frt' style='"+canStyle+"'>取&nbsp;&nbsp;消</a>");
		var modalDemoBtn = $("<a href='javascript:;' class='zhu-enterBtn frt' style='display:none;"+titStyle+"'>试&nbsp;&nbsp;玩</a>");
		var modalBodyText = $("<p></p>");
		
		modalEnterBtn.click(modal.enterClick);
		modalCancelBtn.click(modal.cancelClick);
		modalPanel.append(modalHead,modalBody,modalFoot);
		modalBg.append(modalPanel);
		
		var modalObj = {
			modalBg : modalBg,
			modalPanel : modalPanel,
			modalHead : modalHead,
			modalBody : modalBody,
			modalFoot : modalFoot,
			modalEnterBtn : modalEnterBtn,
			modalCancelBtn : modalCancelBtn,
			modalDemoBtn : modalDemoBtn,
			modalBodyText : modalBodyText
		};

		return modalObj;
	}

	function showModule(object){
		$("body").append(object.modalBg);
		// $(window.top.frames["mem_index"].document).find("body").append(object.modalBg);
		setTimeout(function(){object.modalBg.removeClass("zhu-transparent");},1);
	}

	function showModuleCurrent(object){
		$("body").append(object.modalBg);
		setTimeout(function(){object.modalBg.removeClass("zhu-transparent");},1);
	}

	

});

