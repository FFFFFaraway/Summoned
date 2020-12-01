<template>
  <div>
    <div v-if="this.signed >= 0">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>Others released Summoned</span>
        </div>
      <list :summoneds="summoneds" />
      </el-card>
      
      <br>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>Your requests</span>
        </div>
        <ul>
          <li v-for="r in passedRequests" :key="r.ID">
            <p>
              Summoned ID: {{ r.summoned_id }}, Desc: {{ r.desc }}, Status: {{ r.status }}
            </p>
          </li>
          <li v-for="r in waitRequests" :key="r.ID">
            <p>
              Summoned ID: {{ r.summoned_id }},
              Desc: 
              <el-input type="text" v-model="r.desc"></el-input>
              Status: {{ r.status }}
              <el-button type="primary" @click="update(r)">update</el-button>
              <el-button type="primary" @click="cancel(r)">cancel</el-button>
            </p>
          </li>
        </ul>
      </el-card>
    </div>
    <div v-else>
      <p>please sign in for more information</p>
    </div>
  </div>
</template>
<script>
import { mapState } from "vuex";
export default {
  components: {
    list: () => import("../components/SummonedList"),
  },
  data() {
    return {
      summoneds: [],
      requests: [],
    };
  },
  computed: {
    ...mapState("auth", ["signed"]),
    passedRequests() {
      return this.requests.filter(r => r.status != 'Waiting')
    },
    waitRequests() {
      return this.requests.filter(r => r.status == 'Waiting')
    }
  },
  methods: {
    update(request) {
      let that = this;
      var formData = new FormData();
      formData.append("id", request.ID);
      formData.append("desc", request.desc);
      this.$axios
        .put("request", formData)
        .then(function () {
          alert("Successfully updated the request");
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    },
    cancel(request) {
      let that = this;
      this.$axios
        .delete("request/" + request.ID)
        .then(function () {
          alert("Successfully cancel the request");
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    }
  },
  mounted() {
    let that = this;
    this.$axios
      .get("othersummoned")
      .then(function (response) {
        that.summoneds = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("requestByUser")
      .then(function (response) {
        that.requests = response.data.requests;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>