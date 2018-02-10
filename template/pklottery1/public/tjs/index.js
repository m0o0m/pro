define(["jquery", "/public/tjs/lottery.js"], function(e, t) {
	var n, r = "button-current",
		i = "table-current";

	return n = t.extend({
		initialize: function(t) {

			var r = this,
				i = _.extend({
					hln: !0,
					lotteryId: 'bj_10'
				}, t);
			n.superclass.initialize.call(this, i), this.isOpen = !0, this._np = 0, this.refreshTimer = new this.CountDown(this.refreshDuration), this.cache = {
				lz: null,
				ranking: null
			}, e(function() {
				r.setLiveUrl(i.lotteryId), r.ballTpl = e("#tpl-prev-balls").html(), r.$selectedCount = e("#j-selected-count"), r._events()
			})
		},
		close: function() {
			n.superclass.close.call(this),this.$play.find(".zhu_none").removeClass("zhu_none"), this.$play.find(".input").val(""), this._setSelectedCount(0), this.isLianMa && (this.$ctn.find(":checked").prop("checked", !1), this._LianMaConfig && this._LianMaConfig.balls && (this._LianMaConfig.balls.length = 0))
		},
		open: function(e) {
			if (this.isLianMa)
				if (e) {
					var t = this.$play.find(":checkbox");
					t._luzCurrentTabgth == t.filter(":disabled").length ? n.superclass.open.call(this) : this.$play.find("input:not(:checkbox),button").prop("disabled", !1)
				} else n.superclass.open.call(this);
			else n.superclass.open.call(this);
			!e && this.refresh()
		},
		getCategory: function(e) {
			var t = e.closest(".j-betting");
			return this.$ctn.find(".h-category").val() || t.find("th").eq(0).text() || t.closest("tr").prevAll(".thead").eq(0).find("th").text()
		},
		reset: function() {
			n.superclass.reset.call(this), this._setSelectedCount(0), this.$ctn.find("." + i).removeClass(i), this.$ctn.find(".input").val(""), (this.isRengXuan||this.isZhiXuan)?(e(".button-secondary-group,.rx_select").hide(),$(".zhu_none").removeClass("zhu_none"),this.$doc.find(".quick-form").show(),this.$doc.find(".normal-form").hide()) : this.isLianMa ? (e(".button-secondary-group,.selected-ball,.removable").hide(), e("#j-normal-form").show(), this._resetLianMa(), this.$ctn.find("." + i).removeClass(i)) : e(".button-secondary-group,.selected-ball,.removable").show()
		},
		getData: function() {
            ///点击事件处理数据
			var t = this,
				n, r;
                //console.log(t.options);
			if (this.isLianMa) {
				this._LianMaConfig.lines === "0" && (this._LianMaConfig.lines = e(".j-current-odds").find("span").text()), t.$doc.find(".lm-total").text(0);
				return
			}
			if (this.isRengXuan) {
				var tab_list = $(".select_zhu");
				var total =	$(this).attr("data-total");
				tab_list.each(function(){
					var ts = $(this).find(".table-current").length;
				});
				this._RengXuanConfig.lines === "0" && (this._RengXuanConfig.lines = e(".j-current-odds").find("span").text()), t.$doc.find(".lm-total").text(0);
				return
			}
			if (this.isZhiXuan) {
				this._ZhiXuanConfig.lines === "0" && (this._ZhiXuanConfig.lines = e(".j-current-odds").find("span").text()), t.$doc.find(".lm-total").text(0);
				return
			}
            ////$ctn id j_all
            if(this.options.lotteryId == "pc_28"){
                r = this.isQuickMode ? this.$ctn.find("tr." + i) : this.$ctn.find(".j-betting .fb"), this.records = r.map(function(n, r)
                {////处理点击
                    var i = e(this),
                        s, o, u, a;
                    t.isQuickMode ? a = t.amount : (i = e(this).closest("tr"), a = parseInt(e.trim(i.find("input").val()), 10));
                    if (t.isQuickMode || a) return u = t.getCategory(i), a > 0 && (a = a.toFixed(2)), {
                        category: u?u:i.attr('data-type'),
                        id: parseInt(i.attr("data-id"), 10),
                        name: i.find(".table-odd").text(),
                        odds: i.attr("data-odds"),
                        amount: a
                    }
                }).get(), this.isQuickMode && t._setSelectedCount(), this.getPostData();
            }else{
                r = this.isQuickMode ? this.$ctn.find(".j-betting tr." + i) : this.$ctn.find(".j-betting .fb"), this.records = r.map(function(n, r)
                {////处理点击
                    var i = e(this),
                        s, o, u, a;
                    t.isQuickMode ? a = t.amount : (i = e(this).closest("tr"), a = parseInt(e.trim(i.find("input").val()), 10));
                    if (t.isQuickMode || a) return u = t.getCategory(i), a > 0 && (a = a.toFixed(2)), {
                        category: u?u:i.attr('data-type'),
                        id: parseInt(i.attr("data-id"), 10),
                        name: i.find(".table-odd").text(),
                        odds: i.attr("data-odds"),
                        amount: a
                    }
                }).get(), this.isQuickMode && t._setSelectedCount(), this.getPostData();
            }
		},
		afterBet: function() {
			this.reset(), this.isLianMa && this._resetLianMa(), this.refresh()
		},
		afterGetBet: function(e) {
			if (this.isLianMa) {
				var t = this.$play.find(":radio");
				t.filter(":checked").length || t.eq(0).prop("checked", !0).trigger("change")
			}
		},
		beforeConfirm: function() {
			var e = this;
			return this.isLianMa && !e.amount ? (e.ui.msg(e.lang.msg.emptyAmount), !1) : !0
		},
		valid: function() {
			var e = this;
			if (this.isLianMa) {
				var t = this._LianMaConfig;
				return t.id ? t.balls.length < t.min ? (e.ui.msg(e.utils.format(e.lang.msg.minNumbers, t.min)), !1) : !0 : (e.ui.msg(e.lang.msg.empty), !1)
			}
			if (this.isRengXuan) {
				var tType = $("#play-tab").find("li.active").attr("data-target");
				var tbl = $(tType).find(".select_zhu");
				var ny = true;
				var e_total;
				var e_title;
				var ths_tr = $("tr.table-current");
				if(ths_tr.length <= 0){
					return e.ui.msg("请选择球号!"), !1;
				}
				tbl.each(function(){
					var total = $(this).attr("data-total");

					if ($(this).find(".table-current").length > 0 && $(this).find(".table-current").length != total) {
						ny = false;
						e_total = total;
						e_title = $("#"+$(this).closest("table.zhu-table").attr("data-tabletype")+"Title_"+total).find("b").text();
						return false; 
					}
				});
						// console.log(total);
				// this._RengXuanConfig.money=e.amount*e._RengXuanConfig.balls.length;
				this._RengXuanConfig.money = e.amount*e._RengXuanConfig.balls.length;
				this._RengXuanConfig.abde = e.amount;
				this._RengXuanConfig.gclass = $("#play-tab").find("li.active").find("a").text();
				
				return !e.amount ? (e.ui.msg(e.lang.msg.emptyAmount), !1) : (!ny ? (e.ui.msg(e_title+"组"+e.utils.format(e.lang.msg.minNumbers, e_total)), !1) : !0);
			}
			if (this.isZhiXuan) {
				var tbl = $(".zhu-balls");
				var ny = true;
				var e_total;
				var e_title;
				var ths_tr = $("tr.table-current");
				if(ths_tr.length <= 0){
					return e.ui.msg("请选择球号!"), !1;
				}
				tbl.each(function(){
					var total = $(this).attr("data-bnumber");
					if ($(this).find(".table-current").length > 0 && $(this).find(".table-current").length != total) {
						ny = false;
						e_total = total;
						e_title = $("#zixTitle_"+total).find("b").text();
						
						return false; 
					}
				});
				this._ZhiXuanConfig.money = e.amount*e._ZhiXuanConfig.balls.length;
				this._ZhiXuanConfig.abde = e.amount;
				this._ZhiXuanConfig.gclass = $("#play-tab").find("li.active").find("a").text();
				return !e.amount ? (e.ui.msg(e.lang.msg.emptyAmount), !1) : (!ny ? (e.ui.msg(e_title+"组"+e.utils.format(e.lang.msg.minNumbers, e_total)), !1) : !0);
			}
			return n.superclass.valid.call(e)
		},
		afterRefresh: function(e) {
            ////ajax刷新
			var t = this;
			this._updateRanking(e.ChangLong), this._updateSummary(e), this._updateLuZhu(), this._updateHitBalls(e.ZongchuYilou), this._renderBall(e), this.isOpen && this._updateLines(e.Lines), this.refreshTimer.update = function(e) {}, this.refreshTimer.done = function() {
				t.refresh()
			}, t.refreshTimer.restart()
		},
		getRecordsHtml: function() {
			var t, r, i, s = this;
			if (this.isRengXuan) {
				return i = e("#j-rx-tpl").html(),t = s._RengXuanConfig, r = this.tpl.render(i,t), r;
			}else if(this.isZhiXuan){
				return i = e("#j-zx-tpl").html(),t = s._ZhiXuanConfig, r = this.tpl.render(i,t), r;				
			}
			// return this.isRengXuan ? (i = e("#j-rx-tpl").html(),t = s._RengXuanConfig, r = this.tpl.render(i,t), r) : n.superclass.getRecordsHtml.call(this);
			// return this.isZhiXuan ? (i = e("#j-rx-tpl").html(),t = s._RengXuanConfig, r = this.tpl.render(i,t), r) : n.superclass.getRecordsHtml.call(this);
			return this.isLianMa ? (i = e("#j-lm-tpl").html(), t = s._LianMaConfig, t.group = s.utils.combination(t.balls, t.min).length, t.total = 0, t.money = 0, r = s._.template(i)(t), t.total == 0 && e(r).find(".lm-total").text(0), r) : n.superclass.getRecordsHtml.call(this);
		},
		beforeBet: function(e) {
            ////下注前信息预处理
			var t, n = this;

			this.isLianMa && (t = n._LianMaConfig, n.data = [{

				BetContext: t.balls.join(","),
				min:t.min,
				max:t.max,
				gname:t.title,
				Money: n.amount,
				Lines: t.lines,
				Id: t.id,
				BetType: 5
			}]);

			this.isRengXuan && (
                t = n._RengXuanConfig, n.data = _.map(t.balls,function(e){

                    return {
                        BetContext: e.name.join(","),
                        Lines:e.odds,
                        mingxi_3:e.category,
                        Money: t.abde,
                        gname:t.gclass
                    };
                })
			);
            this.isZhiXuan && (t = n._ZhiXuanConfig, n.data = _.map(t.balls,function(e){

				return {
					BetContext: e.name.join(","),
					Lines:e.odds,
					mingxi_3:e.category,
					Money: t.abde,
					gname:t.gclass
				};
			})
			);
            
            ///尝试过滤特定彩种的注单内空格
            //var data = JSON.parse(JSON.stringify(this.data));
            //if(data[0].gname == "PC蛋蛋"){
            //    $.each(data,function(p1,p2){
            //        data[p1].BetContext = p2.BetContext.trim();
            //    });
            //}
            //
            //this.data = data;
            //console.log(JSON.stringify(this.data))
		},
		setLiveUrl: function(t) {
			var n = e(".live"),
				r = "http://www.133918.com/";
			switch (t) {
				case 2:
					r += "gdkl10/shipin/";
					break;
				case 3:
					r += "shishicai/shipin/";
					break;
				case 4:
					r += "pk10/shipin/";
					break;
				case 5:
					r += "xync/shipin/";
					break;
				case 6:
					r += "xyft/shipin";
					break;
				case 7:
					r += "sdc/";
					break;
				default:
			}
			n.attr("href", r)
		},
		_setSelectedCount: function(e) {
			e == undefined && (e = this.records.length), this.$selectedCount.find(".j-selected-count").text(e)
		},
		_afterTabChange: function(e) {
			this._luzCurrentTab = 0
		},
		_tabs: function() {
            /////加载调用  捕获点击对象
			var t = this,
				n = e("#play-tab");
			n.on("click", "li", function(n) {

				n.preventDefault();

				var r = e(this),
					i = e(r.data("target"));
				t.isLianMa = r.data("type") === "lm", t.disabledHits = r.data("hits") === "yes", t.disabledTrends = r.data("trends") === "yes";
				t.isRengXuan = r.data("type") === "rx", t.disabledHits = r.data("hits") === "yes", t.disabledTrends = r.data("trends") === "yes";
				t.isZhiXuan = r.data("type") === "zx", t.disabledHits = r.data("hits") === "yes", t.disabledTrends = r.data("trends") === "yes";
				if(t.isRengXuan){
					$(".button-secondary-group button[data-mode='quick']").trigger("click");
					t._RengXuan();
				}
				if(t.isZhiXuan){
					$(".button-secondary-group button[data-mode='quick']").trigger("click");
					t._ZhiXuan();
				}
				if (i.length) {
					t.$ctn = i;
					var s = t.isQuickMode ? 0 : 1;
                    ///mark
					t.$doc.find(".button-secondary-group button").eq(s).trigger("click"), r.addClass("active").siblings("li").removeClass("active"), t.NumberPostion = r.data("np"), t._np = r.index(), i.show().siblings().hide(), t._afterTabChange.call(t, r), t.refresh(), t.isLianMa && t._LianMa()
				}
				// autoHeight();
			}), n.find("li").eq(t._np).trigger("click")
		},
        ////鼠标点击事件
		_events: function() {
			var t = this;

			this.$doc.on("click", ".button-secondary-group button", function(n) {
				n.preventDefault();
				var i = e(this),
					s = e(this).data("mode");

				i.addClass(r).siblings().removeClass(r), t.reset();
				if (!s) return;
				t.setQuickMode(s === "quick"), t._toggleElements()
			}), this.$doc.on("click", ".j-betting td", function(n) {
				n.preventDefault();
				//console.log(t);
				var r = e(this).parent("tr");
				if (t.isQuickMode) {
					if(t.isOpen==true){
					var s = r.find("td"),
						o = e.trim(r.find(".table-odd").text());
					if (!o) return;
					s.toggleClass(i), r.toggleClass(i), t.getData()
					}
				} else r.find(".input").focus()
			}), this._tabs()
		},
		_toggleElements: function() {
			var e = this.$ctn.find(".j-odds"),
				t = this.$doc.find(".quick-form"),
				n = this.$doc.find(".normal-form"),
				r = "is-quick-mode",
				i = "is-normal-mode";
			(this.isRengXuan||this.isZhiXuan)?(t.show(), n.hide(), e.hide(), this.$selectedCount.show(), this.$play.removeClass(i).addClass(r)):(this.isQuickMode ? (t.show(), n.hide(), e.hide(), this.$selectedCount.show(), this.$play.removeClass(i).addClass(r)) : (t.hide(), e.show(), n.show(), this.$play.removeClass(r).addClass(i), this.$selectedCount.hide()))
		},
		_updateLines: function(t) {
			var n = this;
			for(var i in t){
				ii = (t[i].lastIndexOf(",") > 0 ? t[i].lastIndexOf(","):t[i].length),
				t[i] = t[i].substring(0, ii);
                
		    }
			if (this.isLianMa) {
				e('label[for^="j"]').each(function() {
					var r = e(this),
						i = r.attr("for"),
						s = t[i],
						o = n.$ctn.find("#" + i);
					s ? (o.prop("disabled", !1), e(this).find("span").text(s)) : o.prop("disabled", !0)
				});
				return
			}
			this.$play.find("tr[data-id]").each(function() {
				var n = e(this),
					r = "j" + n.attr("data-id"),
					i = t[r];
				i = i > 0 ? i : 0, n.attr({
					"data-odds": i
				}).find(".odds-text").text(parseFloat(i).toFixed(3))
			})
		},
		_updateRanking: function(t) {
			var n, r, i = _.map(t, function(e, t) {
				var n = "";
				return t % 2 == 0 && (n = "table-odd"), {
					name: e[0],
					issue: e[1],
					odds: n
				}
			});
			if (_.isEqual(i, this.cache.ranking)) return;
			this.cache.ranking = i, n = '{{#items}}<tr><td class="{{odds}} tal">{{name}}</td><td class="{{odds}} td-issue number">{{issue}}' + this.lang.msg.issue + "</td></tr>{{/items}}", r = this.tpl.to_html(n, {
				items: i
			}), e("#changelong").find("tbody").html(r)
		},
		_updateSummary: function(t) {

			var n = this,
				r = t.CloseCount;
			//console.log(t);
//			if(!t.IsLogin){
//				n.close(),this.ui.msg('请登录');// e("#close-timer").text("未登录");
//			}else{
			r > 0 && n.open(!0), n.closeTimer != null && n.closeTimer.stop(), n.closeTimer = new this.CountDown(r),
			n.closeTimer.update = function(t) {
	            function s(t,timer) {
	                var r, s,i = n.lang.date;
	                s = n.utils.secondsFormat(t),        
	                _.map(i,
	                function(t, timer) {
	                    var n = s[timer];
	                   // s[timer] = s[timer] + t
	                    n > 0 ? s[timer] = s[timer] + t: s[timer] = null
	                }),
	                r = n.tpl.to_html("{{days}}{{hours}}{{minutes}}{{seconds}}", s),	            
	                timer.text(r)
	            }
				s(t,e("#close-timer"));
				//e("#close-timer").text(n.utils.secondsFormat(t));
				
			},
			n.closeTimer.done = function() {
				n.close(),e("#close-timer").text(n.lang.msg.closedGate);
			}, n.closeTimer.start(), n.awarTimer != null && n.awarTimer.stop(), n.awarTimer = new this.CountDown(Math.abs(t.OpenCount) + 5), n.awarTimer.update = function(t) {
				//e("#award-timer").text(n.utils.secondsFormat(t, !0));
	            function s(t,timer) {
	                var r, s,i = n.lang.date;
	                s = n.utils.secondsFormat(t),
	          
	                _.map(i,
	                function(t, timer) {
	                    var n = s[timer];
	                   // s[timer] = s[timer] + t
	                    n > 0 ? s[timer] = s[timer] + t: s[timer] = null
	                }),
	                r = n.tpl.to_html("{{days}}{{hours}}{{minutes}}{{seconds}}", s),
		            
	                timer.text(r)
	            }
				s(t,e("#award-timer"));
			}, n.awarTimer.done = function() {
				n.open()
			}, n.awarTimer.start(), e("#current-issue").text(t.CurrentPeriod + n.lang.msg.issue), e("#win-lose").text(t.WinLoss), e("#prev-issue").text(t.PrePeriodNumber)
			
			},
		_updateHitBalls: function(t) {
			var n = this,
				r = e("#trends"),
				i = e("#tpl-hit-miss").html();
			r.show();
			if (n.disabledHits) {
				r.hide();
				return
			}
			t.hit = t.hit["n" + n.NumberPostion], r.html(this.tpl.to_html(i, t))
		},
		_updateLuZhu: function(t) {

			var n = e("#luzhu"),
				r = e("#tpl-luzhu").html(),
				i = this,
				s = i._luzCurrentTab || 0;
			n.show();

			var o = t || _.where(this.betInfo.LuZhu, {
				p: i.NumberPostion
			});

			if (!o.length || i.disabledTrends) {
				n.hide();
				return
			}
			if (_.isEqual(o, this.cache.lz)) return;
			this.cache.lz = o;
			var u = _.map(o, function(e) {
					var t = _.map(e.c.split(","), function(e) {
							var t = e.split(":"),
								n = t[0],
								r = t[1],
								i = r > 1 ? _.times(r, function() {
									return n
								}) : [n];
							return {
								item: i
							}
						}),
						n = 30 - t.length;
					return n > 0 && _.times(n, function() {
						t.push({
							item: []
						})
					}), {
						hd: e.n,
						bd: {
							items: t.reverse()
						}
					}
				}),
				a = {
					hd: _.pluck(u, "hd"),
					bd: _.pluck(u, "bd")
				};

			n.html(this.tpl.to_html(r, a)).tab({
				mouseover: !0,
				current: s,
				selected: function(e, t, n) {
					i._luzCurrentTab = n, i._renderIframe()
				}
			})
		},
		_renderBall: function(t) {
            ///ajax刷新  样式
			var n = this.ballTpl;
			var sum = 0;
			///并不知道js渲染是哪种  故而这样判断
            if(this.options.lotteryId == "pc_28"){
                t = t.PreResult.split(",");
                
                for(var i in t){
                    sum += parseInt(t[i],10);
                }
                if(!sum){
                    sum = 0;
                }
                var html = "";
                var pre = '<i class="icon_2 ball_add ball_28">';
                var next = '</i>';
                $.each(t,function(k,v){
                    html += pre+v+next;
                });
                html += '<span class="ball_txt">=</span>';
                var style =  "";
                switch(sum){
                    case 3:
                    case 6:
                    case 9:
                    case 12:
                    case 15:
                    case 18:
                    case 21:
                    case 24:
                        style = '<i class="icon_2 b-red" style="text-align:center;">';
                    break;
                    case 1:
                    case 4:
                    case 7:
                    case 10:
                    case 16:
                    case 19:
                    case 22:
                    case 25:
                        style = '<i class="icon_2 b-green" style="text-align:center;">';
                    break;
                    case 2:
                    case 5:
                    case 8:
                    case 11:
                    case 17:
                    case 20:
                    case 23:
                    case 26: 
                        style = '<i class="icon_2 b-blue" style="text-align:center;">';
                    break;
                    default:
                        style = '<i class="icon_2 ball_28" style="text-align:center;">';
                }
                html += style+sum+'</i>';
                r = html;
            }else{
                t = t.PreResult.split(",");
                
                for(var i in t){
                    sum += parseInt(t[i],10);
                }
                if(!sum){
                    sum = 0;
                }
                var r = this.tpl.to_html(n, {
                    balls: t,
                    balls_sum: sum
                });
            }
			r && e("#prev-bs").html(r)
		},
		_resetLianMa: function() {
			this._genderedLianMaBalls(!0).filter(":checked").prop("checked", !1).trigger("change"), this.$ctn.find(":radio:checked").prop("checked", !1), this.amount = 0
		},
		_resetRengXuan: function() {
			this._genderedRengXuanBalls(!0).filter(".table-current").prop("checked", !1).trigger("change"), this.$ctn.find(":radio:checked").prop("checked", !1), this.amount = 0
		},
		_genderedLianMaBalls: function(t) {
			var n = this.$ctn.find(":checkbox");
			return t ? n : n.filter(":checked").map(function() {
				return e(this).attr("id").replace("b-", "")
			}).get()
		},
		_LianMa: function() {
			var t = this;
			this._LianMaConfig = {
				max: 6,
				total: 0,
				min: 2,
				money: 0,
				lines: 0,
				title: "",
				balls: [],
				_balls: [],
				id: 0,
				group: 0
			};
			var n = this._LianMaConfig,
				r = "j-current-odds";
			t.$ctn.off("change", 'input[name="lm"]').on("change", 'input[name="lm"]', function() {
				$("input[type='checkbox']").each(function(index, el) {
					$(this).attr('checked', false);
					$(this).attr('disabled', false);
					$(this).parent("td").parent("tr").removeClass(i);
				});
				//console.log(t);
				var i = e(this),
				s = i.data("min"),
				m=0,
				o = e('label[for="' + this.id + '"]');
				if(t.lotteryId=='bj_8' || t.lotteryId=='jnd_bs'){
				    switch (i.data("min")) {
			        case 1:
			        	m=10;
			        break;
			        case 2:
			        	m=6;
			        break;
			        case 3:
			        	m=6;
			        break;
			        case 4:
			        	m=7;
			        break;
			        case 5:
			        	m=8;
			        break;
			        default:
			            ;
			        break;
			    }
				}else{
					m = i.data("min")+3
				}

//					if(s>1){
//						m=6;
//					}
					// if(s==2){
					// 	m = 6;
					// }else if(s==3){
					// 	m = 7;
					// }else if(s==4){
					// 	m = 8;
					// }else if(s==5){
					// 	m = 8;
					// }

				e("." + r).removeClass(r), o.addClass(r);
				var u = o.find("span").text();
				n.title = o.find("b").text(), n.id = i.val(), u && (n.lines = u), n.min = s,n.max = m, t.$ctn.find(":checkbox:checked:first").trigger("change")
			}), t.$ctn.off("change", ":checkbox").on("change", ":checkbox", function() {

				//t.ui.msg(t.lang.msg.limit);
			//	alert(e(this).lang.msg.emptyAmount);
				//console.log(t.lang.msg.limit);
				var langtext = "只允许选择"+n.max+"个号码";
				t._genderedLianMaBalls().length > n.max ?(t.ui.msg(langtext), $(this).prop("checked", !1)) : e(this).parent('td').parent('tr').toggleClass(i), n.balls = t._genderedLianMaBalls(), n._balls = [{
					id: parseInt(n.id, 10),
					category: n.title
				}], t.records = {
					category: n.title,
					id: n.id
				}
			}), t.$doc.off("keyup blur", ".single-bet").on("keyup blur", ".single-bet", function() {
				var r = e(this),
					i = r.val();
				i ? (n.money = (i * n.group).toFixed(2), n.total = n.money, t.amount = i) : t.amount = 0, t.$doc.find(".lm-total").text(n.total)
			}), t.$doc.off("click", ".j-highlights-tb tr").on({
				mouseenter: function() {
					//e(this).addClass(i)
				},
				mouseleave: function() {
					//e(this).removeClass(i)
				},
				click: function(t) {
					var n, r;
					n = e(this).find("input");
					if (n.prop("disabled")) return !1;
					e(t.target).is(n) || (t.preventDefault(), r = n.prop("checked"), n.prop("checked", !r).trigger("change"))
				}
			}, ".j-highlights-tb tr")
		},		
		_genderedRengXuanBalls: function(t) {
			var n = this.$ctn.find(".table-current");
			if (t == 0){
				var oba = [];
				var i = 0;
				var oi = -1;
				if (n.length > 0){
					n.each(function(){
						var thf = $(this).closest(".select_zhu");
						var tot = thf.attr("data-total");
						var tht = $(this).closest(".zhu-table").attr("data-tabletype");			
						var bn = $(this).find("i").attr("data-number");						
						var tit = $("#"+tht+"Title_"+tot).find("b").text();
						var ods = $("#"+tht+"Title_"+tot).find("span").text();
						
						if (tot != i) {
							oba.push({
								category:tit,
								name:[bn],
								odds:ods,
								money:0,
							});
							i = tot;
							oi++;
						}else{
							oba[oi].name.push(bn);
						}
					});
				}
				return oba;
			}else{
				return n;
			}


		},
		_genderedZhiXuanBalls: function(t) {
			var n = this.$ctn.find(".table-current");
			if (t == 0){
				var oba = [];
				var i = 0;
				var oi = -1;
				var nb = 2;
				if (n.length > 0){
					n.each(function(){
						var thf = $(this).closest(".zhu-balls");
						var tot = thf.attr("data-bnumber");

						var bn = $(this).find("i").attr("data-number");
						var tit = $("#zixTitle_"+tot).find("b").text();
						var ods = $("#zixTitle_"+tot).find("span").text();

						if (tot != i) {
							oba.push({
								category:tit,
								name:[bn],
								odds:ods,
								bname:["【第一球:"+bn+"】"],
							});
							i = tot;
							oi++;
							nb = 2;
						}else{
							oba[oi].name.push(bn);
							switch(nb){
								case 2:
									oba[oi].bname.push("【第二球:"+bn+"】");
									nb++;
									break;
								case 3:
									oba[oi].bname.push("【第三球:"+bn+"】");
									nb++;
									break;
								case 4:
									break;
							}
						}
					});
				}
				return oba;
			}else{
				return n;
			}


		},
		_RengXuan: function() {

			var t = this;
			var ts = this;

			this._RengXuanConfig = {
				max: 6,
				total: 0,
				min: 2,
				money: 0,
				lines: 0,
				title: "",
				balls: [],
				_balls: [],
				id: 0,
				lengths: 0,
				abde : 0,
				gclass:"任选"
			};
			var n = this._RengXuanConfig,
				r = "j-current-odds";

			 t.$doc.off("click", ".j-highlights-tb tr").on({
				click: function(t) {
					var ths = e(this);
					var thf = ths.closest("table");
					var total = thf.attr("data-total");
					var tr_list = thf.find("tr");
					
					if (!ths.hasClass("table-current")) {
						if (thf.find("."+i).length < total) {
							ths.addClass(i);
							n.balls = ts._genderedRengXuanBalls(0);
							n.lengths = n.balls.length;
							if (thf.find("."+i).length >= total) {
								tr_list.each(function(){
									if (!$(this).hasClass(i)){
										$(this).addClass("zhu_none");
									}
								});
							}

						}else{

						}

						
					}else{
						ths.removeClass(i);
						n.balls = ts._genderedRengXuanBalls(0);
						n.lengths = n.balls.length;
						if(thf.find(".zhu_none").length > 0){
							tr_list.removeClass("zhu_none");
						}
					}
					// t._RengXuanConfig._balls = [{}]
					
				}
			}, ".j-highlights-tb tr")
		},
		_ZhiXuan: function() {

			var t = this;
			var ts = this;

			this._ZhiXuanConfig = {
				max: 6,
				total: 0,
				min: 2,
				money: 0,
				lines: 0,
				title: "",
				balls: [],
				_balls: [],
				id: 0,
				lengths: 0,
				abde : 0,
			};
			var n = this._ZhiXuanConfig;

			 t.$doc.off("click", ".j-highlights-tb tr").on({
				click: function(t) {
					var ths = e(this);
					var thf = ths.closest("table");
					var total = thf.attr("data-total");
					var tr_list = thf.find("tr");
					var thff = ths.closest(".zhu-balls");
					var ftotal = thff.attr("data-bnumber");
					var flist = thff.find(".select_zhu");
					var nub = ths.find("i").attr("data-number");

					if (!ths.hasClass("table-current")) {
						if (thf.find("."+i).length < total && !ths.hasClass("zhu_none")) {
							ths.addClass(i);
							
							n.balls = ts._genderedZhiXuanBalls(0);
							n.lengths = n.balls.length;
								flist.each(function(){
									var a = $(this).find("i[data-number='"+nub+"']").closest("tr");//当前选区当前号码对象
									if(!a.hasClass("table-current")){//如果当前号码没有被选中
										a.addClass("zhu_none");
										if(a.closest(".select_zhu").find("."+i).length < total){//判断当前选区号码是否选满
											a.closest(".select_zhu").addClass("ky_zhu");
										}
										
									}
								});
							if (thf.find("."+i).length >= total) {
								tr_list.each(function(){
									if (!$(this).hasClass(i)){
										$(this).addClass("zhu_none");
									}
								});
								ths.closest(".zhu-balls .ky_zhu").removeClass("ky_zhu");
							}

						}else{

						}

						
					}else{
						ths.removeClass(i);
						ths.closest(".zhu-balls").find(".ky_zhu .zhu_none").removeClass("zhu_none");
						ths.closest(".zhu-balls").find(".ky_zhu").removeClass("ky_zhu");
						n.balls = ts._genderedZhiXuanBalls(0);
						n.lengths = n.balls.length;
						var on_l = ths.closest(".zhu-balls").find(".table-current");//当前选区被选中的球list
						var fList = thff.find(".select_zhu");//组列表
						fList.each(function(){
							var thisZhu = $(this);
							var zhuList = thisZhu.find("tr");
							if(on_l.length == 0){
								thff.find(".zhu_none").removeClass("zhu_none");
							}else{
								zhuList.each(function(){
									
									if(thisZhu.find(".table-current").length != 0){
										
									}else{
									var ny = true;
									var tis = $(this);
									on_l.each(function(){
										var sn = $(this).find("i").attr("data-number");//当前球号
										if(tis.find("i").attr("data-number")== sn){
											ny = false;
										}
									});
									
									if(ny){
										$(this).removeClass("zhu_none");
									}
									}
								});
							}
							
						});
						
							
					}
					// t._RengXuanConfig._balls = [{}]
					
				}
			}, ".j-highlights-tb tr")
		}
	}), n
});