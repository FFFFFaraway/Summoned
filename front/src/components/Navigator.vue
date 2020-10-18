<template>
  <div>
    <router-link to="/home">Home</router-link> |
    <router-link to="/signup" v-show="!signed">Sign up</router-link> |
    <router-link to="/login" v-show="!signed">Log in</router-link> |
    <router-link to="/profile" v-show="signed">Profile</router-link> |
    <router-link to="/mysummoned" v-show="signed">My summoned</router-link> |
    <router-link to="/othersummoned" v-show="signed">Others summoned</router-link> |
    <router-link to="/summoned">Search summoned</router-link> |
    <button @click="logout" v-show="signed">Log out</button>
  </div>
</template>

<script>
import {mapState} from "vuex";
export default {
  computed: {
    ...mapState("auth", ["signed"]),
  },
  methods: {
    logout() {
      let that = this
      this.$axios.post('signed')
      .then(function() {
        alert("successfully log out.")
        that.$store.commit("auth/set", false)
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
