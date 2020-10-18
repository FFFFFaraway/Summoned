<template>
  <div>
    <h1>My released Summoneds</h1>
    <list :summoneds="summoneds"/>
    <hr>
    <h1>Release a NEW Summoned</h1>
    <multiselect v-model="newSummoned.type" :options="options"></multiselect>
    <label for="name">Name: </label>
    <input type="text" v-model="newSummoned.name" id="name">
    <br>
    <label for="desc">Description: </label>
    <input type="text" v-model="newSummoned.desc" id="desc">
    <br>
    <label for="people">People Needed: </label>
    <input type="number" v-model.number="newSummoned.people" id="people">
    <br>
    <label for="ddl">Deadline: </label>
    <input type="text" v-model="newSummoned.ddl" id="ddl">
    <br>
    <label for="img">Image: </label>
    <input type="file" @change="processImg($event)" id="img">
    <br>
    <button @click="submit" multiple>Add</button>
  </div>
</template>

<script>
import list from '../components/SummonedList'
export default {
  components: {
    'list': list,
  },
  data() {
    return {
      summoneds: [],
      options: ['技术交流','学业探讨','社会实践','公益志愿者','游玩','其它'],
      newSummoned: {
        type: null,
        name: "",
        desc: "",
        people: 0,
        ddl: "",
        img: null,
      },
      message: "",
    }
  },
  methods: {
    submit(){
      let that = this
      var formData = new FormData();
      if(this.newSummoned.type == null || this.newSummoned.name == "" ||
      this.newSummoned.desc == "" || this.newSummoned.people == 0 ||
      this.newSummoned.ddl == "" || this.newSummoned.img == null){
        alert("Please fill all information before add")
        return
      }
      formData.append('type', this.newSummoned.type);
      formData.append('name', this.newSummoned.name);
      formData.append('desc', this.newSummoned.desc);
      formData.append('people', this.newSummoned.people);
      formData.append('ddl', this.newSummoned.ddl);
      formData.append('img', this.newSummoned.img);
      this.$axios.post('mysummoned', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
      })
      .then(function() {
        alert("Successfully released the summoned")
        that.$router.go()
      })
      .catch(function(response) {
        console.log(response)
      })
    },
    processImg(event){
      this.newSummoned.img = event.target.files[0]
    }
  },
  mounted() {
    let that = this
    this.$axios.get('mysummoned')
    .then(function(response) {
      that.summoneds = response.data
    })
    .catch(function(response) {
        console.log(response)
    })
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>