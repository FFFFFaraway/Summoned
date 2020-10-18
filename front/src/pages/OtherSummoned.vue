<template>
  <div>
    <h1>Others released Summoned</h1>
    <list :summoneds="summoneds"/>
  </div>
</template>
<script>
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