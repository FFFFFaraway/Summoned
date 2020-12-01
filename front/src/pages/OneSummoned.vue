<template>
  <div>
    <!-- error message -->
    <div v-if="this.message != null">
      {{message}}
    </div>
    <div v-else>
      <sumDtl :summoned="summoned"/>
      <div v-if="this.summoned.status == 'Waiting'">
        <br>
        <!-- 令主 -->
        <div v-if="this.summoned.user_id == this.signed">
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span v-if="this.waitingRequests.length > 0">Exist Waiting Requests for this summoned</span>
              <span v-else>No Waiting Requests for this summoned</span>
            </div>
            <ul>
              <li v-for="r in waitingRequests" :key="r.ID">
                <p>ID: {{ r.user_id }}, Desc: {{ r.desc }}</p>
                <router-link :to="'../profile/' + r.user_id">User Profile</router-link>
                <br>
                <el-button type="primary" @click="acc(r.user_id)">Accept</el-button>
                <el-button type="primary" @click="rej(r.user_id)">Reject</el-button>
              </li>
            </ul>
          </el-card>
        </div>
        <br>
        <!-- 未登录 -->
        <div v-if="this.signed == -1">
          <p>Please sign in for more information</p>
        </div>
        <div v-else>
          <!-- 令主，修改召集令 -->
          <div v-if="this.summoned.user_id == this.signed">
            <div v-if="this.requests.length == 0">
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span>Update Your Summoned</span>
                </div>
              <updSum :summoned="this.summoned" />
              </el-card>
            </div>
            <div v-else>
              <el-card>
                <p>Someone has sent request, can't update anymore</p>
              </el-card>
            </div>
          </div>
          <!-- 不是令主，请求 -->
          <div v-else>
            <el-card class="box-card">
              <div slot="header" class="clearfix">
                <span>{{requestStatusMessage[requestStatus]}}</span>
              </div>
              <reqSum v-if="this.requestStatus == 'Not'" :summoned="this.summoned" />
            </el-card>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  components: {
    updSum: () => import("../components/UpdateSummoned"),
    reqSum: () => import("../components/RequestSummoned"),
    sumDtl: () => import("../components/SummonedDetail"),
  },
  data() {
    return {
      id: this.$route.params.id,
      summoned: {},
      message: null,
      requestStatus: false,
      requestStatusMessage: {
        "Not": "Take This Summoned",
        "Waiting": "Waiting for response, you can still change your request",
        "Accepted": "Your request has been Accepted",
        "Rejected": "Your request has been Rejected",
      },
      requests: [],
    };
  },
  computed: {
    ...mapState("auth", ["signed"]),
    waitingRequests: function () {
      return this.requests.filter((r) => r.status === "Waiting");
    },
  },
  methods: {
    set(status, user_id) {
      let that = this;
      var formData = new FormData();
      formData.append("user_id", user_id);
      formData.append("status", status);
      formData.append("people", this.summoned.people);
      this.$axios
        .put("requestStatus/" + this.id, formData)
        .then(function () {
          alert("Successfully " + status + " the request");
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    },
    acc(user_id) {
      this.set("Accepted", user_id);
    },
    rej(user_id) {
      this.set("Rejected", user_id);
    },
  },
  mounted() {
    let that = this;
    this.$axios
      .get("summoned/" + this.id)
      .then(function (response) {
        that.message = response.data.message;
        that.summoned = response.data;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("requestStatus/" + this.id)
      .then(function (response) {
        that.requestStatus = response.data.requestStatus;
      })
      .catch(function (response) {
        console.log(response);
      });
    this.$axios
      .get("request/" + this.id)
      .then(function (response) {
        that.requests = response.data.requests;
      })
      .catch(function (response) {
        console.log(response);
      });
  },
};
</script>

