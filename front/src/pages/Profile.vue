<template>
  <div>
    <div v-if="this.message != null">
      {{ message }}
    </div>
    <div v-else>
      <div v-if="this.signed == -1">
        <p>please sign in for more information</p>
      </div>
      <div v-else>
        <div v-if="this.signed == this.id">
          <h1>Profile</h1>
          <p>Hello {{ user.Username }}</p>
          <p>Your Phone number: {{ user.Phone }}</p>
          <p>ID: {{ user.ID }}</p>
          <p>CreatedAt: {{ user.CreatedAt }}</p>
          <p>UpdatedAt: {{ user.UpdatedAt }}</p>

          <label for="password">Change Password: </label>
          <input type="password" v-model="password" id="password" />

          <label for="phone">Change Phone number: </label>
          <input type="number" v-model="phone" id="phone" />

          <button @click="submit">Save</button>
        </div>
        <div v-else>
          <h1>Other Profile</h1>
          <p>Username {{ user.Username }}</p>
          <p>Phone number: {{ user.Phone }}</p>
          <p>ID: {{ user.ID }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  data() {
    return {
      id: this.$route.params.id,
      user: {
        Username: "",
        Phone: "",
        ID: -1,
        CreatedAt: "-1",
        UpdatedAt: "-1",
      },
      password: "",
      phone: "",
      message: null,
    };
  },
  computed: {
    ...mapState("auth", ["signed"]),
  },
  methods: {
    submit() {
      var formData = new FormData();
      formData.append("password", this.password);
      formData.append("phone", this.phone);
      let that = this;
      this.$axios
        .post("profile", formData)
        .then(function (response) {
          alert(response.data.message);
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    },
  },
  mounted() {
    let that = this;
    this.$axios
      .get("profile/" + this.id)
      .then(function (response) {
        that.user = response.data;
        that.message = response.data.message
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>