<template>
  <div>
    <navi />
    <router-view></router-view>
  </div>
</template>

<script>
import navi from "./components/Navigator";
export default {
  name: "App",
  components: {
    navi: navi,
  },
};
</script>
<style>
  body {font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;}
</style>