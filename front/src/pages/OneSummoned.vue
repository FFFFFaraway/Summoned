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

    <div v-if="this.summoned.UserID==this.signed">
      <hr>
      <h1>Exist Waiting Requests for this summoned</h1>
      <ul>
        <li v-for="r in waitingRequests" :key="r.ID">
          <p>ID: {{r.UserID}}, Desc: {{r.desc}}</p>
          <button @click="acc(r.UserID)">Accept</button>
          <button @click="rej(r.UserID)">Reject</button>
        </li>
      </ul>
    </div>

    <hr>
    <div v-if="this.signed == -1">
      <p>Please sign in for more information</p>
    </div>
    <div v-else>
      <div v-if="this.summoned.UserID==this.signed">
        <div v-if="this.requests.length == 0">
          <h1>Update Your Summoned</h1>
          <updSum :summoned="this.summoned"/>
        </div>
        <div v-else>
          <p>Someone has sent request, can't update anymore</p>
        </div>
      </div>
      <div v-else>
        <div v-if="this.requestStatus=='Not'">
          <h1>Take This Summoned</h1>
          <reqSum :summoned="this.summoned"/>
        </div>
        <div v-if="this.requestStatus=='Waiting'">
          <h1>Waiting for response, you can still change your request</h1>
        </div>
        <div v-if="this.requestStatus=='Accepted'">
          <h1>Your request has been Accepted</h1>
        </div>
        <div v-if="this.requestStatus=='Rejected'">
          <h1>Your request has been Rejected</h1>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {mapState} from "vuex";
import updSum from '../components/UpdateSummoned'
import reqSum from '../components/RequestSummoned'
export default {
  components: {
    'updSum': updSum,
    'reqSum': reqSum,
  },
  data() {
    return {
      id: this.$route.params.id,
      summoned: {},
      message: null,
      requestStatus: false,
      requests: [],
    }
  },
  computed: {
    ...mapState("auth", ["signed"]),
    waitingRequests: function() {
      return this.requests.filter(r => r.Status === 'Waiting')
    }
  },
  methods: {
    set(status, UserID) {
      let that = this
      var formData = new FormData();
      formData.append('UserID', UserID);
      formData.append('status', status);
      this.$axios.put('requestStatus/'+this.id, formData)
      .then(function() {
        alert("Successfully " + status + " the request")
        that.$router.go()
      })
      .catch(function(response) {
        console.log(response)
      })
    },
    acc(UserID) {
      this.set("Accepted", UserID)
    },
    rej(UserID) {
      this.set("Rejected", UserID)
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
    this.$axios.get('requestStatus/'+this.id)
    .then(function(response) {
      that.message = response.data.message
      that.requestStatus = response.data.requestStatus
    })
    .catch(function(response) {
      console.log(response)
    })
    this.$axios.get('request/'+this.id)
    .then(function(response) {
      that.message = response.data.message
      that.requests = response.data.requests
    })
    .catch(function(response) {
      console.log(response)
    })
    
  }
}
</script>

