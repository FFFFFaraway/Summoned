<template>
  <div>
    <div v-if="this.signed >= 0">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>My released Summoneds</span>
        </div>
        <list :summoneds="summoneds" />
      </el-card>
      <br>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>Release a NEW Summoned</span>
        </div>
      <creSum />
      </el-card>
    </div>
    <div v-else>
      <p>please sign in for more information</p>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  components: {
    list: () => import("../components/SummonedList"),
    creSum: () => import("../components/CreateSummoned"),
  },
  data() {
    return {
      summoneds: [],
    };
  },
  computed: {
    ...mapState("auth", ["signed"]),
  },
  mounted() {
    let that = this;
    this.$axios
      .get("mysummoned")
      .then(function (response) {
        that.summoneds = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>
