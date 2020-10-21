<template>
  <div>
    <div v-if="this.signed >= 0">
    <h1>Others released Summoned</h1>
    <list :summoneds="summoneds"/>
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
  }
}
</script>