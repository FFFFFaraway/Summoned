<template>
  <div>
    <span><router-link to="/home">Home</router-link> | </span>
    <span v-show="!signedIn"><router-link to="/signup">Sign up</router-link> | </span>
    <span v-show="!signedIn"><router-link to="/login">Log in</router-link> | </span>
    <span v-show="signedIn"><router-link :to="'/profile/' + this.signed">Profile</router-link> | </span>
    <span v-show="signedIn"><router-link to="/mysummoned">My summoned</router-link> | </span>
    <span v-show="signedIn"><router-link to="/othersummoned">Others summoned</router-link> | </span>
    <span><router-link to="/summoned">Search summoned</router-link> | </span>
    <el-button type="primary" @click="logout" v-show="signedIn">Log out</el-button>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  computed: {
    ...mapState("auth", ["signed"]),
    signedIn() {
      return this.signed >= 0;
    },
  },
  methods: {
    logout() {
      let that = this;
      this.$axios
        .post("signed")
        .then(function () {
          alert("successfully log out.");
          that.$store.commit("auth/set", -1);
        })
        .catch(function (response) {
          console.log(response);
        });
    },
  },
  mounted() {
    let that = this;
    this.$axios
      .get("signed")
      .then(function (response) {
        that.$store.commit("auth/set", response.data.signed);
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>
