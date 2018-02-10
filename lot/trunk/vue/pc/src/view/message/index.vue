<template>
  <div class="betting">
    <div class="r_box clearfix">
      <div class="betrecord">
        <ul>
          <!-- :class="{'nuv-ul-s-list-click':v.flag}" -->
          <li v-for="(item,i) in nav" @click='mypath(item,i)' :class="{'styleBg':setindex==i}">
            <span>{{item.name}}</span>
          </li>
        </ul>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
  //  import cash from './cash.vue'
  export default {
    components: {
      //     cash,
    },
    props: {},
    data() {
      return {
        setindex: 0,
        nav: [
          {
            name: "公告消息",
            item: "notice"
          },
          {
            name: "个人消息",
            item: "personal"
          }
        ]
      };
    },
    created() {
      this.$root.$emit('wh_modal',false);
      var str = this.$route.path;
      str = str.slice(str.lastIndexOf("/") + 1);

      for (let i = 0; i < this.nav.length; i++) {
        if (this.nav[i].item==str) {
          this.setindex=i;
          break
        }
      }

    },
    mounted() {},
    methods: {
      mypath(item, i) {
        this.setindex = i;
        // console.log(item);
        this.$router.push({
          name: item.item
        });
      }
    }
  };
</script>
<style lang="scss" src="../../assets/css/betting_index.scss" scoped></style>
