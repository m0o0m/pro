<template>
    <div class="often">
        <div class="pk_time">
            <div class="time_left">
                <div class="left_logo">
                    <img :src="c_data.img_path" alt="">
                </div>
                <div class="left_text">
                    <p>{{c_data.fc_name}}</p>
                    <p><span>第</span> {{c_data.qishu}} <span>期</span></p>
                </div>
            </div>
            <div class="time_center">
                <p class="center_top">投注剩余时间</p>
                <div class="center_bottom">
                    <div class="fl time_content" v-if="h < 10 && h >= 0">0{{h}}</div>
                    <div class="fl time_content" v-else-if="h < 0">--</div>
                    <div class="fl time_content" v-else>{{h}}</div>
                    <div class="fl fs">时</div>
                    <div class="fl time_content" v-if="m < 10 && m >= 0">0{{m}}</div>
                    <div class="fl time_content" v-else-if="m < 0">--</div>
                    <div class="fl time_content" v-else>{{m}}</div>
                    <div class="fl fs">分</div>
                    <div class="fl time_content" v-if="s < 10 && s >= 0">0{{s}}</div>
                    <div class="fl time_content" v-else-if="s < 0">--</div>
                    <div class="fl time_content" v-else>{{s}}</div>
                    <div class="fl fs">秒</div>
                </div>
            </div>
            <div class="time_bottom">
                <div class="bottom_left"><span>第</span> {{auto.qishu}} <span>期开奖</span></div>
                <div v-if="this.$route.query.page == 'liuhecai'" class='bottom_center6'>
                    <div class="bottom_content">
                        <div class="content_list" v-for="item in auto.ball">
                            <p :class="[item.color,'box']">{{item.number}}</p>
                            <p class="animal">{{item.animal}}</p>
                        </div>
                    </div>
                </div>
                <div v-else :class="[auto.ball.length >= 10?'bottom_center_other':'bottom_center']">
                    <div :class="[auto.ball.length == 10?'other_content':'bottom_content']">
                        <div class="content_list" v-for="item in auto.ball">
                            <p class="box blue">{{item}}</p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="bottom_bottom">
                <p @click="history()">历史结果</p>
                <p @click="open_way()">开奖走势</p>
                <p @click="show_rule()">玩法规则</p>
            </div>
        </div>
        <Nav-top :lists="menus" @menu="go_child"></Nav-top>
        <router-view :Ilists="integrateLists" :cdata="c_data"></router-view>
        <Nav-bottom ref="dewdrop_map" :back_data="auto_list" :nav_top="bottom_nav"></Nav-bottom>
    </div>
</template>

<script type="text/ecmascript-6">
    function sscInitial() {
        return [
            {
                name: "第一球",
                object: [
                    {
                        money: "",
                        li_id: 0,
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 1,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 2,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 3,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 4,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 5,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 6,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 7,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 8,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 9,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 10,
                        name:'大',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 11,
                        name:'小',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 12,
                        name:'单',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 13,
                        name:'双',
                        flag: false
                    },
                ]
            },
            {
                name: "第二球",
                object: [
                    {
                        money: "",
                        li_id: 14,
                        flag: false
                        //
                    },
                    {
                        money: "",
                        li_id: 15,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 16,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 17,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 18,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 19,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 20,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 21,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 22,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 23,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 24,
                        name:'大',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 25,
                        name:'小',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 26,
                        name:'单',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 27,
                        name:'双',
                        flag: false
                    },
                ]
            },
            {
                name: "第三球",
                object: [
                    {
                        money: "",
                        li_id: 28,
                        flag: false
                        //
                    },
                    {
                        money: "",
                        li_id: 29,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 30,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 31,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 32,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 33,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 34,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 35,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 36,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 37,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 38,
                        name:'大',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 39,
                        name:'小',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 40,
                        name:'单',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 41,
                        name:'双',
                        flag: false
                    },
                ]
            },
            {
                name: "第四球",
                object: [
                    {
                        money: "",
                        li_id: 42,
                        flag: false
                        //
                    },
                    {
                        money: "",
                        li_id: 43,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 44,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 45,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 46,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 47,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 48,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 49,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 50,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 51,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 52,
                        name:'大',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 53,
                        name:'小',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 54,
                        name:'单',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 55,
                        name:'双',
                        flag: false
                    },
                ]
            },
            {
                name: "第五球",
                object: [
                    {
                        money: "",
                        li_id: 56,
                        flag: false
                        //
                    },
                    {
                        money: "",
                        li_id: 57,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 58,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 59,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 60,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 61,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 62,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 63,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 64,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 65,
                        flag: false
                        // name: "00"
                    },
                    {
                        money: "",
                        li_id: 66,
                        name:'大',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 67,
                        name:'小',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 68,
                        name:'单',
                        flag: false
                    },
                    {
                        money: "",
                        li_id: 69,
                        name:'双',
                        flag: false
                    },
                ]
            },
            {
                name: "",
                object: [
                    {
                        money: "",
                        li_id: 70,
                        name:'大',
                        flag: false
                    },
                ]
            },
            {
                name: "",
                object: [
                    {
                        money: "",
                        li_id: 71,
                        name:'小',
                        flag: false
                    },
                ]
            },
            {
                name: "",
                object: [
                    {
                        money: "",
                        li_id: 72,
                        name:'单',
                        flag: false
                    },
                ]
            },
            {
                name: "",
                object: [
                    {
                        money: "",
                        li_id: 73,
                        name:'双',
                        flag: false
                    },
                ]
            }
        ];
    }
    import api from "../../../../api/config";
    import NavTop from "../../../../share_components/default_nav";
    import ws from '../../../../assets/js/socket'
    import NavBottom from '../../../../share_components/dewdrop_map'
    import cm_cookie from '../../../../assets/js/com_cookie'
    export default {
        components: {
            NavTop,NavBottom
        },
        data() {
            return {
                auto_list:[],
                bottom_nav:[
                    {name:'万位'},
                    {name:'千位'},
                    {name:'百位'},
                    {name:'十位'},
                    {name:'个位'},
                ],
                margin: false,
                menus: [
                    {
                        name: "整合",
                        item: 'ffc_o'
                    },
                ],
                routePage: null,
                integrateLists:[
                    {
                        name: "第一球",
                        object: [
                            {
                                money: "",
                                li_id: 0,
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 1,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 2,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 3,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 4,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 5,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 6,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 7,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 8,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 9,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 10,
                                name:'大',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 11,
                                name:'小',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 12,
                                name:'单',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 13,
                                name:'双',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "第二球",
                        object: [
                            {
                                money: "",
                                li_id: 14,
                                flag: false
                                //
                            },
                            {
                                money: "",
                                li_id: 15,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 16,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 17,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 18,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 19,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 20,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 21,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 22,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 23,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 24,
                                name:'大',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 25,
                                name:'小',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 26,
                                name:'单',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 27,
                                name:'双',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "第三球",
                        object: [
                            {
                                money: "",
                                li_id: 28,
                                flag: false
                                //
                            },
                            {
                                money: "",
                                li_id: 29,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 30,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 31,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 32,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 33,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 34,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 35,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 36,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 37,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 38,
                                name:'大',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 39,
                                name:'小',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 40,
                                name:'单',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 41,
                                name:'双',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "第四球",
                        object: [
                            {
                                money: "",
                                li_id: 42,
                                flag: false
                                //
                            },
                            {
                                money: "",
                                li_id: 43,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 44,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 45,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 46,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 47,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 48,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 49,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 50,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 51,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 52,
                                name:'大',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 53,
                                name:'小',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 54,
                                name:'单',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 55,
                                name:'双',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "第五球",
                        object: [
                            {
                                money: "",
                                li_id: 56,
                                flag: false
                                //
                            },
                            {
                                money: "",
                                li_id: 57,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 58,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 59,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 60,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 61,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 62,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 63,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 64,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 65,
                                flag: false
                                // name: "00"
                            },
                            {
                                money: "",
                                li_id: 66,
                                name:'大',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 67,
                                name:'小',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 68,
                                name:'单',
                                flag: false
                            },
                            {
                                money: "",
                                li_id: 69,
                                name:'双',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "",
                        object: [
                            {
                                money: "",
                                li_id: 70,
                                name:'大',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "",
                        object: [
                            {
                                money: "",
                                li_id: 71,
                                name:'小',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "",
                        object: [
                            {
                                money: "",
                                li_id: 72,
                                name:'单',
                                flag: false
                            },
                        ]
                    },
                    {
                        name: "",
                        object: [
                            {
                                money: "",
                                li_id: 73,
                                name:'双',
                                flag: false
                            },
                        ]
                    }
                ],
                timer: '',
                h: 0,
                m: 0,
                s: 0,
                c_data: {
                    fc_name: "",
                    img_path: "",
                    qishu: ""
                },
                auto: {
                    qishu: null,
                    datetime: "",
                    ball: []
                },
                close_time: {
                    fengpan: "",
                    kaijiang: "",
                    kaipan: "",
                    now_time: ""
                },
                is_wh: false,
            };
        },
        created() {
            this.fetchData();
        },
        watch: {
            // 如果路由有变化，会再次执行该方法
            '$route.query.page':function(to,from) {
                this.$root.$off(from);
                this.$root.$off(from+'lefttime');
                this.fetchData();
                this.socket_change(to);
            }
        },
        mounted(){
            this.socket_change(this.$route.query.page);
        },
        destroyed(){
            console.log('清除定时器：'+this.timer);
            if(this.timer){
                window.clearTimeout(this.timer);
                this.timer = null;
            }
            if(!this.isIE9){
                ws.close_ws(false);
            }
            this.$root.$off(this.$route.query.page);
            this.$root.$off(this.$route.query.page+'lefttime');
            cm_cookie.delCookie("top_nav")
        },
        methods:{
            socket_change: function(to){
                if(!this.isIE9){
                    let self = this;
                    ws.createWebSocket(to,self,true);
                    this.$root.$on(to,(e)=>{
                        this.auto = e;
                    });
                    this.$root.$on(to+'lefttime',(e)=>{
                    this.c_data.qishu = e.qishu;
                    this.close_time.fengpan = e.close_time;
                    this.close_time.now_time = e.now_time;
                    let t1 = e.close_time - e.now_time;
                    console.log('是否在维护中：'+this.is_wh);
                    if(t1 == 0 && !this.is_wh){
                        this.fetchData(2)
                    }else if(t1 > 0  && !this.is_wh){
                        if(this.timer){
                            window.clearTimeout(this.timer);
                            this.timer = null;
                        }
                        this.init();
                        this.$root.$emit('wh_modal',false);
                    }else if(t1 < 0  && !this.is_wh){
                        this.h = -1;
                        this.m = -1;
                        this.s = -1;
                        if(this.timer){
                            window.clearTimeout(this.timer);
                            this.timer = null;
                        }
                        this.$root.$emit('wh_modal',true,true)
                    }
                })
                }
            },
            open_way: function (){
                let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=1';
                window.open(page)
            },
            history: function () {
                let page = 'trend_chart/chart-lotteryId='+this.$route.query.page+'.html'+'?tab=3';
                window.open(page)
            },
            //aaaaa
            show_rule: function() {
                this.$root.$emit("rule_show", true);
                this.$root.$emit("now_page", this.$route.query.page);
            },
            getRTime: function() {
                this.close_time.now_time += 1;
                var t1 = this.close_time.fengpan * 1000 - this.close_time.now_time * 1000;
                // var d=Math.floor(t/1000/60/60/24);
                this.h = Math.floor((t1 / 1000 / 60 / 60) % 24);
                this.m = Math.floor((t1 / 1000 / 60) % 60);
                this.s = Math.floor((t1 / 1000) % 60);
                // console.log('开奖时间：'+'时：'+h+'；分：'+m+'；秒：'+s);
                console.log('封盘时间：' + '时：' + this.h + '；分：' + this.m + '；秒：' + this.s);
                console.log('定时器id(time_out)：'+this.timer);
                if (this.h == 0 && this.m == 0 && this.s == 0) {
                    this.fetchData(2);
                    this.$root.$emit("success", true);
                }else if(this.h < 0 && this.m < 0 && this.s < 0){
                    if(this.timer){
                        window.clearTimeout(this.timer);
                        this.timer = null;
                    }
                    this.h = -1;
                    this.m = -1;
                    this.s = -1;
                }
            },
            init: function() {
               this.getRTime();
               this.timer = window.setTimeout(this.init,1000); //time是指本身,延时递归调用自己,1000为间隔调用时间,单位毫秒
            },
            sortNumber: function(a, b) {
                return a.sort - b.sort;
            },
            go_child: function(child) {
                this.$router.push({ name: child, query: { page: this.$route.query.page} });
            },
            fetchData(type) {
                this.$root.$emit('wh_modal',false);
                if(type == 2){
                    this.$root.$emit('loading',true,true);
                }else{
                    this.$root.$emit('loading',true);
                }
                if(this.timer){
                    window.clearTimeout(this.timer);
                    this.timer = null;
                }
                let body = {
                    fc_type: this.$route.query.page
                };
                api.dewdrop(this, body, (res) => {
                    if (res.data.ErrorCode == 1) {
                        this.auto_list = res.data.Data;
                        api.getgameindex(this, body, res => {
                            if (res.data.ErrorCode == 1) {
                                if(type == 2){
                                    window.setTimeout(() => {
                                        this.$root.$emit("loading", false);
                                    }, 1000)
                                }else{
                                    this.$root.$emit("loading", false);
                                }
                                if(res.data.is_wh == 2){
                                    this.$root.$emit('wh_modal',true);
                                    this.is_wh = true;
                                }else if(res.data.is_wh == 1){
                                    this.$root.$emit('wh_modal',false);
                                    this.is_wh = false;
                                }
                                this.auto = res.data.Data.auto;
                                this.close_time = res.data.Data.closetime;
                                if(this.close_time.fengpan - this.close_time.now_time < 0){
                                   this.$root.$emit('wh_modal',true,true)
                                }
                                this.c_data = res.data.Data.c_data;
                                let data = res.data.Data.odds;
                                data.sort(this.sortNumber);
                                this.computed(data);
                                this.$refs.dewdrop_map.top_go(0);//点击触发露珠图组件头部选中事件
                                this.$refs.dewdrop_map.left_go(0);//点击触发露珠图组件左侧选中事件
                                if(!this.is_wh){
                                    if(this.close_time.fengpan){
                                        if(this.timer){
                                            window.clearTimeout(this.timer);
                                            this.timer = null;
                                        }
                                        if(type == 2){
                                            window.setTimeout(() => {
                                                this.init();
                                            }, 1000)
                                        }else{
                                            this.init();
                                        }
                                    }
                                }
                            }
                        })
                    }
                });
            },

            computed(data) {
                this.$set(this.integrateLists, this.integrateLists);
                this.integrateLists = sscInitial();
                for(let i = 0;i<this.integrateLists.length;i++){
                    for (let j = 0; j < this.integrateLists[i].object.length;j++){
                        if(i <= 4){
                            Object.assign(this.integrateLists[i].object[j],data[(i*14)+j]);
                            let name = data[(i*14)+j].remark.slice(
                                data[(i*14)+j].remark.search("#") + 1,
                                data[(i*14)+j].remark.length
                            );
                            this.integrateLists[i].object[j].name = name;
                        }else if(i >= 5){
                            Object.assign(this.integrateLists[i].object[j], data[65+i]);
                            let name = data[65+i].remark.slice(
                                data[65+i].remark.search("#") + 1,
                                data[65+i].remark.length
                            );
                            this.integrateLists[i].object[j].name = name;
                        }
                    }
                }
            },
        }
    };
</script>
