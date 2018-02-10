<template>
    <div>
        <div class="key_buttom" @click="key_show = !key_show,num_error = false">
            <i class="iconfont pk-jianpan"></i>
        </div>
        <div class="key_body" v-show="key_show">
            <div :class="[num_error?'wobble':'','key_of_number']">
               {{key_of_number}}
            </div>
            <div class="key_div">
                <p class="key_num" v-for="item in numbers" @click="go(item)">
                    <span v-if="item.code != 8">{{item.num}}</span>
                    <i v-else class="iconfont pk-shanchu">{{item.num}}</i>
                </p>
            </div>
        </div>
    </div>
</template>
<script>
    import {Input} from 'iview';
    export default{
        components: {
           'I-Input':Input
        },
        data(){
            return {
                key_show: false,
                key_of_number:'',
                num_error: false,
                numbers: [
                    {num: '1', code: 97},
                    {num: '2', code: 98},
                    {num: '3', code: 99},
                    {num: '4', code: 100},
                    {num: '5', code: 101},
                    {num: '6', code: 102},
                    {num: '7', code: 103},
                    {num: '8', code: 104},
                    {num: '9', code: 105},
                    {num: '0', code: 96},
                    {num: '', code: 8},
                    {num: '清空', code: 12},
                ],
                _no_top:false,
                five_show: 0
            }
        },
        created(){
            this.$root.$on('no_top',(e,k)=>{
                this._no_top = e;
                this.five_show = k
            })
        },
        mounted(){
            document.onselectstart = function(){return false;};
            this.$root.$on('clear_key_number',(e)=>{
                this.key_of_number = e;
            });
        },
        methods: {
            go: function (item) {
                let dom = document.querySelectorAll('input');
                this.num_error = false;
                if(item.code == 8){
                    this.key_of_number = this.key_of_number.substring(this.key_of_number.length-1,0);
                }else if(item.code == 12){
                    this.key_of_number = '';
                }else{
                    this.key_of_number = this.key_of_number + item.num;
                    if(this.key_of_number.length > 9){
                        this.key_of_number = this.key_of_number.substring(9,0);
                        this.num_error = true
                    }
                }
                console.log(dom);
                for(let i = 0;i < dom.length;i++){
                  if(dom[i].data_onoff === 'true'){
                     console.log(dom[i]);
                     dom[i].value = this.key_of_number;
                     console.log('当前的值：'+dom[i].value);
                     if(this._no_top){
                       dom[i].focus();
                       console.log('i-am-come-in');
                       if(i == dom.length-1){
                           this.$root.$emit('this_money',dom[i].value);
                       }
                     }else{
                         if(this.five_show == 5){
//                           event.preventDefault();
                           if(i == 0 || i == dom.length-11){
                             this.$root.$emit('this_money',dom[i].value);
                           }
                         }else if(this.five_show == 4 || this.five_show == 3 || this.five_show == 2){
                             if(i == 0 || i == 50){
                               this.$root.$emit('this_money',dom[i].value);
                             }
                         }else{
                           dom[i].focus();
                           if(i == 0 || i == dom.length-1){
                             this.$root.$emit('this_money',dom[i].value);
                           }
                         }
                     }
                  }
                }
            },
        }
    }
</script>
<style lang="scss" scoped>
    @keyframes wobble {
        from {
            transform: none;
        }
        15% {
            transform: translate3d(-25%, 0, 0) rotate3d(0, 0, 1, -5deg);
        }
        30% {
            transform: translate3d(20%, 0, 0) rotate3d(0, 0, 1, 3deg);
        }
        45% {
            transform: translate3d(-15%, 0, 0) rotate3d(0, 0, 1, -3deg);
        }
        60% {
            transform: translate3d(10%, 0, 0) rotate3d(0, 0, 1, 2deg);
        }
        75% {
            transform: translate3d(-5%, 0, 0) rotate3d(0, 0, 1, -1deg);
        }
        to {
            transform: none;
        }
    }
    .wobble {
        -webkit-transform: transition3d(0,0,0);
        animation-name: wobble;
        animation-duration: 1s;
        animation-fill-mode: both;
    }
    .key_buttom {
        position: fixed;
        right: 0px;
        top: 50%;
        /*height: 85px;*/
        border-radius: 5px;
        z-index: 97;
        color: #fff;
        background: #000;
        width: 35px;
    }
    .key_buttom:hover{
        cursor: pointer;
    }

    .key_body {
        width: 170px;
        position: fixed;
        padding-top: 15px;
        right: 0px;
        top: 50%;
        right: 40px;
        border-radius: 5px;
        z-index: 99;
        color: #fff;
        background-color:#000;
        .key_of_number{
            width: 80%;
            height: 40px;
            line-height: 40px;
            margin: 0px auto;
            background-color: rgb(255, 255, 255);
            border-radius: 10px;
            font-size: 20px;
            font-weight: bold;
            color: #000;
        }
        .key_div {
            overflow: hidden;
            padding: 30px 15px 0 15px;
            .key_num {
                box-shadow: 1px 1px 1px 1px;
                border: 1px solid #ddd;
                margin-right: 5px;
                /*background-color: #000000;*/
                margin-bottom: 20px;
                float: left;
                width: 30px;
                height: 30px;
                line-height: 30px;
                border-radius: 10px
            }
            .key_num:hover{
                cursor: pointer;
            }
        }
    }
</style>
