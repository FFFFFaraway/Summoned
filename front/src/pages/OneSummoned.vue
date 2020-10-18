<template>
  <div>
    <p v-if="message!=null">{{message}}</p>
    <div v-else>
      <h1>One Summoned details</h1>
      <ul>
        <li>summoned ID: {{id}}</li>
        <li>type: {{summoned.type}}</li>
        <li>name: {{summoned.name}}</li>
        <li>desc: {{summoned.desc}}</li>
        <li>people: {{summoned.people}}</li>
        <li>ddl: {{summoned.ddl}}</li>
        <li>Img: {{summoned.Img}}</li>
        <li>Status: {{summoned.Status}}</li>
      </ul>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      id: this.$route.params.id,
      summoned: {},
      message: null,
    }
  },
  mounted() {
    let that = this
    this.$axios.get('summoned/'+this.id)
    .then(function(response) {
      that.message = response.data.message
      that.summoned = response.data
    })
    .catch(function(response) {
      console.log(response)
    })
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>
