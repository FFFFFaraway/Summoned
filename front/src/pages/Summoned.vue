<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Summoned</span>
      </div>
    <label for="name">Search by Name</label>
    <el-input type="text" v-model="filterName" id="name" ></el-input>
    <br />
    <label>Search by Type</label>
    <el-select v-model="filterType">
      <el-option v-for="item in options" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <list :summoneds="summoneds" />
    </el-card>
  </div>
</template>

<script>
import list from "../components/SummonedList";
export default {
  components: {
    list: list,
  },
  data() {
    return {
      options: [
        "",
        "技术交流",
        "学业探讨",
        "社会实践",
        "公益志愿者",
        "游玩",
        "其它",
      ],
      filterType: "",
      filterName: "",
      allSummoneds: [],
    };
  },
  computed: {
    summoneds() {
      if (this.filterName == "" && this.filterType == "")
        return this.allSummoneds;
      if (this.filterName == "")
        return this.allSummoneds.filter((s) => s.type == this.filterType);
      if (this.filterType == "")
        return this.allSummoneds.filter((s) =>
          s.name.includes(this.filterName)
        );
      return this.allSummoneds
        .filter((s) => s.name.includes(this.filterName))
        .filter((s) => s.type == this.filterType);
    },
  },
  mounted() {
    let that = this;
    this.$axios
      .get("summoned")
      .then(function (response) {
        that.allSummoneds = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>
