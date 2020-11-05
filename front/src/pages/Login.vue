<template>
  <div>
    <h1>Login</h1>
    <label for="username">Username: </label
    ><input type="text" v-model="username" id="username" />
    <label for="password">Password: </label
    ><input type="password" v-model="password" id="password" />
    <button @click="submit">Log in</button>
    <button @click="reset">Reset</button>
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