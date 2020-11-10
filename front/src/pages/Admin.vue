<template>
  <div>
    <h1>Users</h1>
    <el-table :data="users">
      <div v-for="name in userOptions" :key="name">
        <el-table-column :prop="name" :label="name"> </el-table-column>
      </div>
    </el-table>

    <h1>Summoneds</h1>
    <label for="name">Search by Name</label>
    <input type="text" v-model="filterName" id="name" />
    <br />
    <label>Search by Type</label>
    <el-select v-model="filterType">
      <el-option v-for="item in options" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <el-table :data="summoneds">
      <div v-for="name in summonedOptions" :key="name">
        <el-table-column :prop="name" :label="name"> </el-table-column>
      </div>
      <el-table-column fixed="right" label="owner">
        <template slot-scope="scope">
          <router-link :to="'/profile/' + scope.row.user_id">info</router-link>
        </template>
      </el-table-column>
    </el-table>
    
    <h1>Requests</h1>
    <el-table :data="requests">
      <div v-for="name in requestOptions" :key="name">
        <el-table-column :prop="name" :label="name"> </el-table-column>
      </div>
      <el-table-column fixed="right" label="taker">
        <template slot-scope="scope">
          <router-link :to="'/profile/' + scope.row.user_id">info</router-link>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      summonedOptions: [
        "status",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "user_id",
        "type",
        "name",
        "desc",
        "people",
        "ddl",
        "Img",
      ],
      userOptions: [
        "phone",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "username",
        "password",
      ],
      requestOptions: [
        "status",
        "ID",
        "CreatedAt",
        "UpdatedAt",
        "summoned_id",
        "user_id",
        "desc",
      ],
      options: [
        "",
        "技术交流",
        "学业探讨",
        "社会实践",
        "公益志愿者",
        "游玩",
        "其它",
      ],
      allSummoneds: [],
      users: [],
      requests: [],
      filterType: "",
      filterName: "",
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
    this.$axios
      .get("users")
      .then(function (response) {
        that.users = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("requestsAll")
      .then(function (response) {
        that.requests = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>