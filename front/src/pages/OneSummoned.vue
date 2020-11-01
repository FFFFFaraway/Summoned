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
      <h1>Update Your Summoned</h1>
      <updSum :summoned="this.summoned"/>
    </div>
    <div v-else>
      <hr>
      <h1>Take This Summoned</h1>
      <reqSum :summoned="this.summoned"/>
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
    }
  },
  computed: {
    ...mapState("auth", ["signed"]),
  },
  methods: {
    request() {

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

