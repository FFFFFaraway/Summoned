<template>
  <div>
    <h1>Profile</h1>
    <p>Hello {{user.Username}}</p>
    <p>Your Phone number: {{user.Phone}}</p>
    <p>ID: {{user.ID}}</p>
    <p>CreatedAt: {{user.CreatedAt}}</p>
    <p>UpdatedAt: {{user.UpdatedAt}}</p>

    <label for="password">Change Password: </label>
    <input type="password" v-model="password" id="password">

    <label for="phone">Change Phone number: </label>
    <input type="number" v-model="phone" id="phone">

    <button @click="submit">Save</button>
  </div>
</template>

<script>

export default {
  data() {
    return {
      user: {
        Username: "",
        Phone: "",
        ID: -1,
        CreatedAt: "-1",
        UpdatedAt: "-1",
      },
      password: "",
      phone: "",
    }
  },
  methods: {
    submit(){
      var formData = new FormData();
      formData.append('password', this.password);
      formData.append('phone', this.phone);
      let that = this
      this.$axios.post('profile', formData)
      .then(function(response) {
        alert(response.data.message)
        that.$router.go()
      })
      .catch(function(response) {
        console.log(response)
      })
    },
  },
  mounted() {
    let that = this
    this.$axios.get('profile')
    .then(function(response) {
      that.user = response.data
    })
    .catch(function(response) {
      console.log(response)
    })
  }
}
</script>