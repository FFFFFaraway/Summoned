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
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>Profile</span>
            </div>
            <p>Hello {{ user.username }}</p>
            <div v-if="user.is_admin">
              You are admin, <router-link to="../admin">admin page</router-link>
            </div>
            <p>Your Name: {{ user.name }}</p>
            <p>Your {{user.number_type}} number: {{user.number}}</p>
            <p>Your Phone number: {{ user.phone }}</p>
            <p>Your Rank: {{user.rank}}</p>
            <p>Your Introduction: {{user.introduction}}</p>
            <p>Your City: {{user.city}}</p>
            
            <p>Your UID: {{ user.ID }}</p>
            <p>CreatedAt: {{ user.CreatedAt }}</p>
            <p>UpdatedAt: {{ user.UpdatedAt }}</p>

            <label for="password">Change Password: </label>
            <el-input type="password" v-model="password" id="password" ></el-input>

            <label for="phone">Change Phone number: </label>
            <el-input type="number" v-model="phone" id="phone" ></el-input>

            <label for="introduction">Change Introduction: </label>
            <el-input type="text" v-model="introduction" id="introduction" ></el-input>

            <el-button type="primary" @click="submit">Save</el-button>
          </el-card>
        </div>
        <div v-else>
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>Other Profile</span>
            </div>
            <p>Username {{ user.username }}</p>
            <p>Phone number: {{ user.phone }}</p>
            <p>Rank: {{user.rank}}</p>
            <p>Introduction: {{user.introduction}}</p>
            <p>City: {{user.city}}</p>
          </el-card>
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
      user: {},
      password: "",
      phone: "",
      introduction: "",
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
      formData.append("introduction", this.introduction);
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