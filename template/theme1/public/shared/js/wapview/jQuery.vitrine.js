/**
 * Created by liuwentong on 7/18/15.
 */

;
!
    function($) {

        'use strict';

        // constructor
        var Vitrine = function(ele, opt) {
            this.$element = ele;
            this.defaults = {
                start: null,
                narrow: false,
                speed: 400,
                loop: false,
                title: '',
                color: '#000'
            };
            this.options = $.extend({}, this.defaults, opt);
            this.init(this.options);
        };

        // prototype
        Vitrine.prototype = {

            init: function(options) {
                this.$items = this.$element.children();

                this.length = this.$items.length;
                this.current = Math.min((this.options.start || Math.floor(this.length / 2)), this.length - 1);

                if(this.length < 5) options.narrow = true;

                this.duration = options.speed / 1000;

                // init basic css
                this.$items.find('img').css('width', '100%');

                this.$title = $('<div id="vitrine-title">' + options.title + '</div>');
                this.$desc = $('<div id="vitrine-desc"></div>');

                var titleCSS = {
                    'width': '100%',
                    'font-size': '30px',
                    'z-index': '1',
                    'text-align': 'center',
                    'position': 'absolute',
                    'left': '0',
                    'color': options.color,
                    'top': '15%'
                };
                var descCSS = {
                    'position': 'absolute',
                    'left': '50%',
                    'z-index': '100',
                    'bottom': '-320px',
                    'color':'white',
                    'font-size': '25px',
                    '-webkit-transform': 'translateX(-50%)',
                    'transform': 'translateX(-50%)',
                    'text-align': 'center'
                    // '-webkit-writing-mode': 'vertical-rl'
                };

                this.$title.css(titleCSS);
                this.$desc.css(descCSS);

                this.$element.prepend(this.$title);
                this.$element.append(this.$desc);

                for (var i = 0; i < this.$items.length; i++) {
                    i < this.current && this.$items.eq(i).css(this._getCSS(0));
                    i > this.current && this.$items.eq(i).css(this._getCSS(4));
                }

                // init layout
                this.layout();

            },

            layout: function() {

                this.$showItems = [];

                var positions = [];

                positions[0] = this.options.narrow ? null : this.current - 1 > 0 ? this.current - 2 : (this.options.loop ? this.length - 2 + this.current : null);
                positions[1] = this.current > 0 ? this.current - 1 : (this.options.loop ? this.length - 1 : null);
                positions[2] = this.current;
                positions[3] = this.current < this.length - 1 ? this.current + 1 : (this.options.loop ? 0 : null);
                positions[4] = this.options.narrow ? null : this.current + 1 < this.length - 1 ? this.current + 2 : (this.options.loop ? 2 - this.length + this.current : null);

                this.$items.css({
                    'opacity': '0',
                    'position': 'absolute',
                    'top': '50%',
                    'z-index': '6',
                    'transition': this.duration + 's ease-in-out',
                    '-webkit-transform': 'translate(-50%,-50%)',
                    'transform': 'translate(-50%,-50%)',
                    'cursor': 'pointer'
                });

                var curDesc;

                for (var i = 0; i < 5; i++) {
                    if (positions[i] !== null) {
                        this.$showItems[i] = this.$items.eq(positions[i]);
                        this.$showItems[i].css(this._getCSS(i));
                        if (i == 2) curDesc = this.$showItems[2].attr('data-desc');
                    }
                }

                this.$desc.fadeOut(this.options.speed / 2, function() {
                    $(this).html(curDesc);
                }).fadeIn(this.options.speed / 2);

                this._setEvent();

            },

            _getCSS: function(position) {

                switch (position) {
                    case 0:
                        return {
                            'width': '100px',
                            'left': '15%',
                            'z-index': '7',
                            'opacity': '0.3'
                        };
                    case 1:
                        return {
                            'width': '200px',
                            'left': '30%',
                            'z-index': '8',
                            'opacity': '0.7'
                        };
                    case 2:
                        return {
                            'width': '250px',
                            'left': '50%',
                            'z-index': '9',
                            'opacity': '1'
                        };
                    case 3:
                        return {
                            'width': '200px',
                            'left': '70%',
                            'z-index': '8',
                            'opacity': '0.7'
                        };
                    case 4:
                        return {
                            'width': '100px',
                            'left': '85%',
                            'z-index': '7',
                            'opacity': '0.3'
                        };
                }
            },

            _setEvent: function() {

                this._cancelEvent();

                for (var i = 0; i < 5; i++) {
                    i < 2 && this.$showItems[i] && this.$showItems[i].on('click', this._prev());
                    i > 2 && this.$showItems[i] && this.$showItems[i].on('click', this._next());
                }
            },

            _prev: function(){
                var _self = this;
                return function(){
                    _self.current = _self.current === 0 ? (_self.options.loop ? _self.length - 1 : 0) : _self.current - 1;
                    _self.layout();
                };
            },

            _next: function(){
                var _self = this;
                return function() {
                    _self.current = _self.current === _self.length - 1 ? (_self.options.loop ? 0 : _self.current) : _self.current + 1;
                    _self.layout();
                };
            },

            _cancelEvent: function() {
                for (var i = 0; i < 5; i++)
                    this.$showItems[i] && this.$showItems[i].off('click');
            }

        };

        // Plugin
        $.fn.vitrine = function(options) {
            var virtine = new Vitrine(this, options);
            return this;
        };
    }(jQuery);
