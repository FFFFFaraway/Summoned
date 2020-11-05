<template>
  <div>
    <div v-if="this.signed >= 0">
      <h1>Others released Summoned</h1>
      <list :summoneds="summoneds"/>

      <hr>
      <h1>Your requests</h1>
      <ul>
        <li v-for="r in requests" :key="r.ID">
          <p>Summoned ID: {{r.ID}}, Desc: {{r.desc}}, Status: {{r.Status}}</p>
        </li>
      </ul>
    </div>
    <div v-else>
      <p>please sign in for more information</p>
    </div>
  </div>
</template>
<script>
import {mapState} from "vuex";
import list from '../components/SummonedList'
export default {
  components: {
    'list': list,
  },
  data() {
    return {
      summoneds: [],
      requests: [],
    }
  },
  computed: {
    ...mapState("auth", ["signed"]),
  },
  mounted() {
    let that = this
    this.$axios.get('othersummoned')
    .then(function(response) {
      that.summoneds = response.data
    })
    .catch(function(response) {
      console.log(response)
    })
    this.$axios.get('requestByUser')
    .then(function(response) {
      that.requests = response.data.requests
    })
    .catch(function(response) {
      console.log(response)
    })
  }
}
</script>