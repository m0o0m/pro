/**
 * Created by Administrator on 2017/1/3.common_helper
 */

/***********************************在线客服******************************/
function OnlineService(url) {
    newWin = window.open(url, '', 'width=900,height=600,top=0,left=0,status=no,toolbar=no,scrollbars=yes,resizable=no,personalbar=no');
    window.opener = null;//出掉关闭时候的提示窗口
    window.open('', '_self'); //ie7
    window.close();
}

/***********************************弹出历史消息******************************/
function notice_data() {
    window.open('/notice/data', "History", "width=816,height=500,top=0,left=0,status=no,toolbar=no,scrollbars=yes,resizable=no,personalbar=no");
}

/***********************************代理联盟table选项卡******************************/
$.fn.mtab2 = function (posType) {
    var area = this, bgTop = '', bgBottom = '';
    var posType = (typeof posType !== 'undefined' ? posType : 'l');
    switch (posType) {
        case 'c':
            bgTop = 'top center';
            bgBottom = 'bottom center';
            break;
        case 'r':
            bgTop = 'top right';
            bgBottom = 'bottom right';
            break;
        default:
            bgTop = 'top left';
            bgBottom = 'bottom left'
    }
    $.each(area.find('li[id^=#]'), function (i) {
        if (i != 0) {
            area.find(this.id)[0].style.display = 'none';
        }
    });
    area.find('li[id^=#]').click(function () {
        var self = this;
        $.each(area.find('li[id^=#]'), function (i) {
            if (self.id != this.id) {
                area.find(this.id)[0].style.display = 'none';
                $(this)[0].style.backgroundPosition = bgTop;
                $(this).removeClass('mtab');
            } else {
                area.find(this.id)[0].style.display = 'block';
                $(this)[0].style.backgroundPosition = bgBottom;
                $(this).addClass('mtab');
            }
        });
    });
};

/***********************************日历******************************/
function _getYear(d) {
    var yr = d.getYear();
    if (yr < 1000) yr += 1900;
    return yr;
}

function tick() {
    function initArray() {
        for (i = 0; i < initArray.arguments.length; i++) this[i] = initArray.arguments[i];
    }

    var isnDays = new initArray("星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日");
    var today = new Date();
    var hrs = today.getHours();
    var _min = today.getMinutes();
    var sec = today.getSeconds();
    var clckh = "" + ((hrs > 12) ? hrs - 12 : hrs);
    var clckm = ((_min < 10) ? "0" : "") + _min;
    clcks = ((sec < 10) ? "0" : "") + sec;
    var clck = (hrs >= 12) ? "下午" : "上午";

    //document.getElementById("t_2_1").innerHTML = _getYear(today)+"/"+(today.getMonth()+1)+"/"+today.getDate()+"&nbsp;"+clckh+":"+clckm+":"+clcks+"&nbsp;"+clck+"&nbsp;"+isnDays[today.getDay()];
    document.getElementById("t_2_1").innerHTML = _getYear(today) + "/" + (today.getMonth() + 1) + "/" + today.getDate() + "&nbsp;" + clckh + ":" + clckm + ":" + clcks;

    window.setTimeout("tick()", 100);
}

/***********************************收藏及设为首页******************************/
//用法
//onclick="AddFavorite(window.location, document.title)"
//onclick="SetHome(this, top.location)"

/**
 *加入收藏
 */
function AddFavorite(sURL, sTitle) {
    try {
        window.external.addFavorite(sURL, sTitle);
    }
    catch (e) {
        try {
            window.sidebar.addPanel(sTitle, sURL, "");
        }
        catch (e) {
            alert("加入收藏失败，请使用Ctrl+D进行添加");
        }
    }
}

/**
 *设为首页
 */
function SetHome(obj, vrl) {
    try {
        obj.style.behavior = 'url(#default#homepage)';
        obj.setHomePage(vrl);
    }
    catch (e) {
        if (window.netscape) {
            try {
                netscape.security.PrivilegeManager.enablePrivilege("UniversalXPConnect");
            }
            catch (e) {
                alert("此操作被浏览器拒绝！\n请在浏览器地址栏输入“about:config”并回车\n然后将 [signed.applets.codebase_principal_support]的值设置为'true',双击即可。");
            }
            var prefs = Components.classes['@mozilla.org/preferences-service;1'].getService(Components.interfaces.nsIPrefBranch);
            prefs.setCharPref('browser.startup.homepage', vrl);
        } else {
            alert("您的浏览器不支持，请按照下面步骤操作：1.打开浏览器设置。2.点击设置网页。3.输入：" + vrl + "点击确定。");
        }
    }
}

/***********************************A标签文字闪烁-无需调用******************************/
function toggleColor(id, arr, s) {
    var self = this;
    self._i = 0;
    self._timer = null;

    self.run = function () {
        if (arr[self._i]) {
            $(id).css('color', arr[self._i]);
        }
        self._i == 0 ? self._i++ : self._i = 0;
        self._timer = setTimeout(function () {
            self.run(id, arr, s);
        }, s);
    }
    self.run();
}

//讀取文案連結  data-color
$(function () {
    $('a.js-article-color').each(function () {
        var color_arr = $(this).data('color');

        if ('undefined' == typeof color_arr) return;

        color_arr = color_arr.split('|');

        // 確認顏色數量  2=>閃爍   1=>單一色  0=>跳過
        if (color_arr.length == 2) {
            new toggleColor(this, [color_arr[0], color_arr[1]], 500);
        } else if (color_arr.length == 1 && color_arr[0] != '') {
            $(this).css('color', color_arr[0]);
        }
    });
});


/********************************** IE低版本支持placeholder属性 *****************************/
var JPlaceHolder = {
    //检测
    _check: function () {
        return 'placeholder' in document.createElement('input');
    },
    //初始化
    init: function () {
        if (!this._check()) {
            this.fix();
        }
    },
    //修复
    fix: function () {
        jQuery(':input[placeholder]').each(function (index, element) {

            var self = $(this), txt = self.attr('placeholder');
            self.wrap($('<div></div>').css({
                position: 'relative',
                zoom: '1',
                border: 'none',
                background: 'none',
                padding: 'none',
                margin: 'none'
            }));
            var pos = self.position(), h = self.outerHeight(true), paddingleft = self.css('padding-left');

            var holder = $('<span></span>').text(txt).css({
                position: 'absolute',
                left: pos.left,
                top: pos.top,
                height: h,
                lienHeight: h,
                paddingLeft: paddingleft,
                color: '#aaa'
            }).appendTo(self.parent());
            self.focusin(function (e) {
                holder.hide();
            }).focusout(function (e) {
                if (!self.val()) {
                    holder.show();
                }
            });
            holder.click(function (e) {
                holder.hide();
                self.focus();
            });
        });
    }
};
//执行
$(function () {
    JPlaceHolder.init();
});