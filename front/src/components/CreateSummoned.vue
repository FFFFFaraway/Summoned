<template>
  <div>
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
    <input type="date" v-model="newSummoned.ddl" id="ddl">
    <br>
    <label for="img">Image: </label>
    <input type="file" @change="processImg($event)" id="img">
    <br>
    <button @click="submit">Add</button>
  </div>
</template>

<script>

export default {
  data() {
    return {
      options: ['技术交流','学业探讨','社会实践','公益志愿者','游玩','其它'],
      message: "",
      newSummoned: {
        type: null,
        name: "",
        desc: "",
        people: 0,
        ddl: "",
        img: null,
      },
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
    },
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>



