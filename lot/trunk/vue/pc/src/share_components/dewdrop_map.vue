<template>
    <div style="border-top: 20px solid #fff9ee;float: left;">
        <div class="top">
            <p class="top_title">露珠图</p>
            <p @click="top_go(i)" :class="[i == top_select?'top_select':'','top_nav']" v-for="(item,i) in nav_top">{{item.name}}</p>
        </div>
        <div class="bottom">
            <div class="bottom_left">
                <div @click="left_go(i)" :class="[i == left_select?'left_topS':'','left_top']" v-for="(item,i) in left_nav">
                    <p class="boll mb color1">{{item.one_name}}</p>
                    <p class="text mb">{{item.one_number}}</p>
                    <p class="boll mb color2">{{item.two_name}}</p>
                    <p class="text">{{item.two_number}}</p>
                </div>
            </div>
            <div class="bottom_right" :style="fuck_div" ref="fuck_div">
                <div :style="styleObject" v-show="left_select == 0">
                    <div class="right_div" v-for="key in body_content">
                        <div class="right_one" v-for="item in key.bolls">
                            <div v-if="item.ball">
                                <Tooltip :transfer="true" placement="left-start" v-if="Number(item.ball.split(',')[top_select]) >= max_lint">
                                    <div class="boll_box"><p class="boll boll_color1">大</p></div>
                                    <div slot="content">
                                        <p>期数：{{item.qishu}}</p>
                                        <p><i>开奖号码：{{item.ball}}</i></p>
                                    </div>
                                </Tooltip>
                                <Tooltip :transfer="true" placement="left-start" v-if="Number(item.ball.split(',')[top_select]) < max_lint">
                                    <div class="boll_box"><p class="boll boll_color2">小</p></div>
                                    <div slot="content">
                                        <p>期数：{{item.qishu}}</p>
                                        <p><i>开奖号码：{{item.ball}}</i></p>
                                    </div>
                                </Tooltip>
                            </div>
                            <div v-else>
                                <div class="boll_box" style="width: 34px;"><p class="boll"></p></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div :style="styleObject" v-show="left_select == 1">
                    <div class="right_div" v-for="key in body_content">
                        <div class="right_one" v-for="item in key.bolls">
                            <div v-if="item.ball">
                                <Tooltip :transfer="true" placement="left-start" v-if="Number(item.ball.split(',')[top_select]) %2 == 1">
                                    <div class="boll_box"><p class="boll boll_color1">单</p></div>
                                    <div slot="content">
                                        <p>期数：{{item.qishu}}</p>
                                        <p><i>开奖号码：{{item.ball}}</i></p>
                                    </div>
                                </Tooltip>
                                <Tooltip :transfer="true" placement="left-start" v-if="Number(item.ball.split(',')[top_select]) %2 == 0">
                                    <div class="boll_box"><p class="boll boll_color2">双</p></div>
                                    <div slot="content">
                                        <p>期数：{{item.qishu}}</p>
                                        <p><i>开奖号码：{{item.ball}}</i></p>
                                    </div>
                                </Tooltip>
                            </div>
                            <div v-else>
                                <div class="boll_box" style="width: 34px;"><p class="boll"></p></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
    import {Tooltip} from 'iview'
    export default{
        props:{
            back_data:{
                type:Array
            },
            nav_top:{
                type:Array
            },
            max_lint:{
                type:Number,
                default:5
            }
        },
        components:{
          Tooltip
        },
        data(){
            return{
                fuck_div:{
                    overflowX:'scroll',
                },
                styleObject:{
                    width:'1300px',
                    height:'328px',
                    overflow: 'hidden'
                },
                body_content:[],
                left_nav:[
                    {status:false,one_name:'大',one_number:0,two_name:'小',two_number:0},
                    {status:false,one_name:'单',one_number:0,two_name:'双',two_number:0},
                ],
                top_select:0,
                left_select:0,
            }
        },
        methods:{
            //渲染左侧总计的球数
            computed_boll:function(){
                this.left_nav[0].one_number = 0;
                this.left_nav[0].two_number = 0;
                this.left_nav[1].one_number = 0;
                this.left_nav[1].two_number = 0;
                for(let j = 0;j<this.body_content.length;j++){
                    for(let k = 0;k<this.body_content[j].bolls.length;k++){
                        if(this.body_content[j].bolls[k].ball){
                            if(Number(this.body_content[j].bolls[k].ball.split(',')[this.top_select]) < 5){
                                this.left_nav[0].two_number += 1
                            }else if(Number(this.body_content[j].bolls[k].ball.split(',')[this.top_select]) >= 5){
                                this.left_nav[0].one_number += 1
                            }
                            if(Number(this.body_content[j].bolls[k].ball.split(',')[this.top_select]) % 2 == 1){
                                this.left_nav[1].one_number += 1;
                            }else if(Number(this.body_content[j].bolls[k].ball.split(',')[this.top_select]) % 2 == 0){
                                this.left_nav[1].two_number += 1
                            }
                        }
                    }
                }
            },
            //渲染右侧历史球号
            computed_right: function(k){
                var first_build = [];
                for(let l = 0;l<this.back_data.length;l++){
                   first_build.push({bolls:[]});
                }
                if(this.left_select == 0){
                    var a = 0;
                    for(let i = 0;i<this.back_data.length;i++){
                        if (Number(this.back_data[i].ball.split(',')[k]) < this.max_lint) {//小
                            first_build[a].bolls.push(this.back_data[i]);
                            if(typeof this.back_data[i+1] != 'undefined') {
                                if (Number(this.back_data[i + 1].ball.split(',')[k]) >= this.max_lint) {
                                    a += 1;
                                }
                            }
                        } else if (Number(this.back_data[i].ball.split(',')[k]) >= this.max_lint) {//大
                            first_build[a].bolls.push(this.back_data[i]);
                            if(typeof this.back_data[i+1] != 'undefined') {
                                if (Number(this.back_data[i + 1].ball.split(',')[k]) < this.max_lint) {
                                    a += 1;
                                }
                            }
                        }
                     }
                    console.log('huygo_look_hear!');
                    let b  = this.new_obj(first_build);
                    if(b.length<25){
                        const gap_number1 = 25 - b.length;
                        for(let n = 0;n<gap_number1;n++){
                            b.push({bolls:[]})
                        }
                    }
                    this.computed_list(b);
                }else if(this.left_select == 1){
                    var c = 0;
                    for(let i = 0;i<this.back_data.length;i++) {
                        //单双处理
                        if (Number(this.back_data[i].ball.split(',')[k]) % 2 == 1) {
                            first_build[c].bolls.push(this.back_data[i]);
                            if(typeof this.back_data[i+1] != 'undefined') {
                                if (Number(this.back_data[i + 1].ball.split(',')[k]) % 2 == 0) {
                                    c += 1;
                                }
                            }
                        } else if (Number(this.back_data[i].ball.split(',')[k]) % 2 == 0) {
                            first_build[c].bolls.push(this.back_data[i]);
                            if(typeof this.back_data[i+1] != 'undefined') {
                                if (Number(this.back_data[i + 1].ball.split(',')[k]) % 2 == 1) {
                                    c += 1;
                                }
                            }
                        }
                    }
                    let d = this.new_obj(first_build);
                    if(d.length<25){
                        const gap_number2 = 25 - d.length;
                        for(let n = 0;n<gap_number2;n++){
                           d.push({bolls:[]})
                        }
                    }
                    this.computed_list(d);
                }
            },
            top_go:function (i) {
               this.top_select = i;
               this.computed_right(i);
               this.computed_boll();
               console.log('read_now!!!!!'+i)
            },
            left_go:function (i) {
                this.left_select = i;
                this.computed_right(this.top_select);
                this.computed_boll();
            },
            computed_list: function (back) {
                let is_long = false;
                let long_item = 0;
                for(let k = 0;k<back.length;k++){
                    if(back[k].bolls.length>10){
                        is_long = true;
                        long_item = back[k].bolls.length;
                        break;
                    }
                }
                for(let i = 0;i<back.length;i++){
                    if(is_long){
                        let num = String(328+(long_item-10)*35);
                        this.styleObject.height = num+'px';
                        let add_length = (long_item-back[i].bolls.length);
                        for(let j = 0;j<add_length;j++){
                            back[i].bolls.push({ball:''})
                        }
                    }else{
                        this.styleObject.height = '328px';
                        let add_length1 = (10-back[i].bolls.length);
                        for(let j = 0;j<add_length1;j++){
                            back[i].bolls.push({ball:''})
                        }
                    }
                }
                console.log('长度：'+back.length);
                if(back.length<26){
                    const number1= String(33*26);
                    this.styleObject.width = number1+'px';
                    this.fuck_div.overflowX = 'hidden';
                    this.$refs.fuck_div.scrollLeft = 0
                }else{
                    const number2= String(33*back.length);
                    this.styleObject.width = number2+'px';
                    this.fuck_div.overflowX = 'scroll';
                    this.$refs.fuck_div.scrollLeft = 0
                }
                this.body_content = back
            },
            new_obj: function(e){
                let new_arr = [];
                for(let i = 0;i<e.length;i++){
                    if(e[i].bolls.length != 0){
                        new_arr.push(e[i])
                    }
                }
                return new_arr
            }
        }
    }
</script>
<style lang="scss" scoped>
    @import "../assets/css/function.scss";
    .top{
        overflow: hidden;width: 920px;color: #fff;margin: 0 auto;background: $bg_color5;
        .top_title{
            float: left;
            padding-top: 7px;
            color: #FFCC00;
            font-weight: bold;
            width: 70px;
            font-size: 16px;
        }
        .top_nav{
            float: left;
            padding: 10px 15px;
            font-size: 13px;
        }
        .top_nav:hover{
            background-color: #FFF;
            border-top:1px solid $bg_color3;
            font-weight: bold;
            color: #000;
            padding: 9px 15px 10px 15px;
            cursor: pointer;
        }
        .top_select{
            background-color: #FFF;
            border-top:1px solid $bg_color3;
            font-weight: bold;
            color: #000;
            padding: 9px 15px 10px 15px;
        }
    }
    .bottom{
        overflow: hidden;width:920px;color: #fff;margin: 0 auto;
        .bottom_left{
            float: left;
            width:70px;
            /*margin-top: 20px;*/
            /*background: $bg_color5;*/
            .left_top{
                padding: 10px 0;
                .boll {
                    -webkit-border-radius: 20px;
                    -moz-border-radius: 20px;
                    border-radius: 20px;
                    width: 30px;
                    height: 30px;
                    line-height: 30px;
                    margin: 0 auto;
                    font-weight: bold;
                    font-size: 16px;
                }
                .text{
                    color: #000;
                }
                .mb{
                    margin-bottom: 5px;
                }
                .color1{
                    background: $lg_blue2
                }
                .color2{
                    background: $lg_orange;
                }
            }
            .left_topS{
                width:70px;
                border-left:2px solid #FFCC00;
                background-color: rgba(82, 210, 246, 0.16)
            }
        }
        .bottom_right{
            float: left;
            color: #fff;
            width:850px;
            /*height: 400px;*/
            /*overflow:hidden;*/
            /*overflow-x:scroll;*/
            .right_div{
                float: left;
                width:33px;
                /*height: 400px;*/
                .right_one{
                    width: 35px;
                    .boll_box {
                        padding:5px;
                        border-bottom:1px solid #eee;
                        border-right:1px solid #eee;
                        .boll {
                            -webkit-border-radius: 20px;
                            -moz-border-radius: 20px;
                            border-radius: 20px;
                            width: 22px;
                            height: 22px;
                            line-height: 22px;
                            margin: 0 auto;
                            font-weight: bold;
                            font-size: 10px;
                            cursor: Default;
                        }
                        .boll_color1{
                            background: $lg_blue2;
                        }
                        .boll_color2{
                            background: $lg_orange;
                        }
                    }
                }
            }

        }
    }
</style>