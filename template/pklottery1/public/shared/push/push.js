////push
if (typeof console == "undefined") { this.console = { log: function (msg) {  } };}
WEB_SOCKET_SWF_LOCATION = cdnUrl+"/shared/push/swf/WebSocketMain.swf";
WEB_SOCKET_DEBUG = false;
var ws={};

// 连接服务端
function connect() {
   // 创建websocket
   ws = new WebSocket("wss://link.pk051.com/");
   // 当socket连接打开时，根据ifream内容不同选择不同的处理方式
   ws.onopen = onopen;
   // 当有消息时根据消息类型显示不同信息
   ws.onmessage = onmessage;
   ws.init = socket_init;
   ws.onclose = function() {
        var callback = ws.callback;
        ///自动重连
		document.customize.ws = connect();
		document.customize.ws.init(callback,'onmessage');
   };
   ws.onerror = function() {
        console.log("出现错误");
   };
   return ws;
}

// 连接建立时
function onopen()
{


    ///如果header未共享
    var data = {};
    data.type = 'xxx';//随意定义
    data.group = 'fc';
    ws.send(JSON.stringify(data));

    // ///如果header未共享
    // var data = {};
    // var patten = /(?:\/([^\/]+)$)/;
    // ///检测ifream 内部调用的方法
    // var func = $('#iframepage').context.location.pathname.match(patten);

    // switch(func[1]){
    //     case 'lottery':
    //     break;
    // }
}

// 服务端发来消息时
function onmessage(e)
{
    var data = JSON.parse(e.data);
    switch(data.type){
        case 'ping':
            var data = {};
            data.type = 'pong';
            ws.send(JSON.stringify(data));
        break;

        default:   ///onmessage  自定义回调
            ws.callback(data);
        break;
    }
}

///初始化
function socket_init(callback,even_status){
    switch(even_status){
        case 'onmessage':
            ws.callback = callback;
        break;

        case 'onopen':
            ws.start = callback;
        break;
    }
}

$(document).ready(function(){
    document.customize = function(){};
    document.customize.ws = connect();
});

////push