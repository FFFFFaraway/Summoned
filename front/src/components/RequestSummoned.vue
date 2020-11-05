<template>
  <div>
    <div v-if="haveRequested">
      <p>Waiting Response</p>
    </div>
    <div v-else>
      <label for="desc">Description: </label>
      <input type="text" v-model="newRequest.desc" id="desc">
      <br>
      <button @click="submit">Request</button>
    </div>
  </div>
</template>

<script>

export default {
  props: {
    summoned: {},
    haveRequested: Boolean,
  },
  data() {
    return {
      newRequest: {
        desc: "",
      }
    }
  },
  methods: {
    submit(){
      let that = this
      var formData = new FormData();
      if(this.newRequest.desc == null){
        alert("Please fill desc before request")
        return
      }
      formData.append('desc', this.newRequest.desc);
      formData.append('ID', this.summoned.ID);
      this.$axios.post('request', formData)
      .then(function() {
        alert("Successfully sent request, please wait for response")
        that.$router.go()
      })
      .catch(function(response) {
        console.log(response)
      })
    }
  }
}
</script>




