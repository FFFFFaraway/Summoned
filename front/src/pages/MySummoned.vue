<template>
  <div>
    <div v-if="this.signed >= 0">
      <h1>My released Summoneds</h1>
      <list :summoneds="summoneds" />
      <hr />
      <h1>Release a NEW Summoned</h1>
      <creSum />
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
