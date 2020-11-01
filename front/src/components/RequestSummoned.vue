<template>
  <div>
    <label for="desc">Description: </label>
    <input type="text" v-model="newRequest.desc" id="desc">
    <br>
    <button @click="submit">Request</button>
  </div>
</template>

<script>

export default {
  props: {
    summoned: {},
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
      this.$axios.post('mysummoned', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
      })
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




