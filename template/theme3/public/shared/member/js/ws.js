// TODO 推送取款消息的websocket服务
var withdrawWsService = (function () {

    //参数：传递给单例的一个参数集合
    function Singleton(orderNumber) {
        if (orderNumber != null) {
            var loc = window.location;
            var uri = 'ws:';

            if (loc.protocol === 'https:') {
                uri = 'wss:';
            }
            uri += '//' + loc.host;
            uri += '/ws/withdraw';
            console.log(uri);
            ws = new WebSocket(uri);

            ws.onopen = function () {
                console.log('Connected');
                console.log(orderNumber);
                ws.send(String(orderNumber));
            };
            ws.onmessage = function (evt) {
                if (f !== undefined) {
                    this.f(evt.data);
                }
                console.log(evt.data);
                ws.close();
                instance = undefined;
                //     var out = document.getElementById('output');
                //     out.innerHTML += evt.data + '<br>';
            };
        }
    }

    var f;

    function Notify(f) {
        this.f = f;
    }


    //实例容器
    var instance;

    var _static = {
        name: 'SingletonTester',
        //获取实例的方法
        //返回Singleton的实例
        Subscribe: function (args) {
            if (instance === undefined) {
                instance = new Singleton(args);
            }
            return instance;
        },
        Notify(f) {
            if (instance !== undefined) {
                instance.Notify(f)
            }
        }
    };
    return _static;
})();

// TODO 推送存款消息的websocket服务
var depositWsService = (function () {

    //参数：传递给单例的一个参数集合
    function Singleton(orderNumber) {
        if (orderNumber != null) {
            var loc = window.location;
            var uri = 'ws:';

            if (loc.protocol === 'https:') {
                uri = 'wss:';
            }
            uri += '//' + loc.host;
            uri += '/ws/deposit';
            console.log(uri);
            ws = new WebSocket(uri);

            ws.onopen = function () {
                console.log('Connected');
                console.log(orderNumber);
                ws.send(String(orderNumber));
            };
            ws.onmessage = function (evt) {
                if (f !== undefined) {
                    this.f(evt.data);
                }
                console.log(evt.data);
                ws.close();
                instance = undefined;
                //     var out = document.getElementById('output');
                //     out.innerHTML += evt.data + '<br>';
            };
        }
    }

    var f;

    function Notify(f) {
        this.f = f;
    }


    //实例容器
    var instance;

    var _static = {
        name: 'SingletonTester',
        //获取实例的方法
        //返回Singleton的实例
        Subscribe: function (args) {
            if (instance === undefined) {
                instance = new Singleton(args);
            }
            return instance;
        },
        Notify(f) {
            if (instance !== undefined) {
                instance.Notify(f)
            }
        }
    };
    return _static;
})();