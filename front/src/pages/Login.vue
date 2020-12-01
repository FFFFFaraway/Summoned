<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>Login</span>
      </div>
    <label for="username">Username: </label>
    <el-input type="text" v-model="username" id="username" ></el-input>
    <label for="password">Password: </label>
    <el-input type="password" v-model="password" id="password" ></el-input>
    <el-button type="primary" @click="submit">Log in</el-button>
    <el-button type="primary" @click="reset">Reset</el-button>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    submit() {
      var formData = new FormData();
      formData.append("username", this.username);
      formData.append("password", this.password);
      let that = this;
      this.$axios
        .post("login", formData, {
          withCredentials: true,
        })
        .then(function (response) {
          alert(response.data.message);
          if (response.data.message == "successfully log in") {
            that.$store.commit("auth/set", response.data.user_id);
            that.$router.push("home");
          } else return;
        })
        .catch(function (response) {
          console.log(response);
        });
    },
    reset() {
      this.username = "";
      this.password = "";
    },
  },
};
</script>