<template>
  <div>
    <div v-if="this.signed >= 0">
      <h1>Others released Summoned</h1>
      <list :summoneds="summoneds" />

      <hr />
      <h1>Your requests</h1>
      <ul>
        <li v-for="r in passedRequests" :key="r.ID">
          <p>
            Summoned ID: {{ r.SummonedID }}, Desc: {{ r.desc }}, Status: {{ r.Status }}
          </p>
        </li>
        <li v-for="r in waitRequests" :key="r.ID">
          <p>
            Summoned ID: {{ r.SummonedID }},
            Desc: 
            <input type="text" v-model="r.desc">,
            Status: {{ r.Status }}
            <button @click="update(r)">update</button>
            <button @click="del(r)">delete</button>
          </p>
        </li>
      </ul>
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
      return this.requests.filter(r => r.Status == 'Accepted' || r.Status == 'Rejected')
    },
    waitRequests() {
      return this.requests.filter(r => r.Status == 'Waiting')
    }
  },
  methods: {
    update(request) {
      let that = this;
      var formData = new FormData();
      formData.append("ID", request.ID);
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
    del(request) {
      let that = this;
      this.$axios
        .delete("request/" + request.ID)
        .then(function () {
          alert("Successfully delete the request");
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