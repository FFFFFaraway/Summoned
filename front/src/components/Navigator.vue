<template>
  <div>
    <router-link to="/home">Home</router-link> |
    <router-link to="/signup" v-show="!signedIn">Sign up</router-link> |
    <router-link to="/login" v-show="!signedIn">Log in</router-link> |
    <router-link to="/profile" v-show="signedIn">Profile</router-link> |
    <router-link to="/mysummoned" v-show="signedIn">My summoned</router-link> |
    <router-link to="/othersummoned" v-show="signedIn">Others summoned</router-link> |
    <router-link to="/summoned">Search summoned</router-link> |
    <button @click="logout" v-show="signedIn">Log out</button>
  </div>
</template>

<script>
import {mapState} from "vuex";
export default {
  computed: {
    ...mapState("auth", ["signed"]),
    signedIn() {
      return this.signed >= 0
    }
  },
  methods: {
    logout() {
      let that = this
      this.$axios.post('signed')
      .then(function() {
        alert("successfully log out.")
        that.$store.commit("auth/set", -1)
      })
      .catch(function(response) {
        console.log(response)
      })
    }
  },
  mounted() {
    let that = this
    this.$axios.get('signed')
    .then(function(response) {
        that.$store.commit("auth/set", response.data.signed)
    })
    .catch(function(response) {
      console.log(response)
    })
  }
}
</script>
