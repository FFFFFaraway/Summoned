<template>
  <div>
    <multiselect v-model="summoned.type" :options="options"></multiselect>
    <label for="name">Name: </label>
    <input type="text" v-model="summoned.name" id="name" />
    <br />
    <label for="desc">Description: </label>
    <input type="text" v-model="summoned.desc" id="desc" />
    <br />
    <label for="people">People Needed: </label>
    <input type="number" v-model.number="summoned.people" id="people" />
    <br />
    <label for="ddl">Deadline: </label>
    <input type="text" v-model="summoned.ddl" id="ddl" />
    <br />
    <label for="img">Image: </label>
    <input type="file" @change="processImg($event)" id="img" />
    <br />
    <button @click="submit">Update</button>
    <button @click="deleteSummoned">Delete</button>
  </div>
</template>

<script>
export default {
  props: {
    summoned: {},
  },
  data() {
    return {
      options: [
        "技术交流",
        "学业探讨",
        "社会实践",
        "公益志愿者",
        "游玩",
        "其它",
      ],
    };
  },
  methods: {
    submit() {
      let that = this;
      var formData = new FormData();
      if (
        this.summoned.type == null ||
        this.summoned.name == "" ||
        this.summoned.desc == "" ||
        this.summoned.people == 0 ||
        this.summoned.ddl == ""
      ) {
        alert("Please fill all information before " + this.type);
        return;
      }
      formData.append("ID", this.summoned.ID);
      formData.append("type", this.summoned.type);
      formData.append("name", this.summoned.name);
      formData.append("desc", this.summoned.desc);
      formData.append("people", this.summoned.people);
      formData.append("ddl", this.summoned.ddl);
      formData.append("img", this.summoned.img);

      this.$axios
        .put("mysummoned", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        })
        .then(function () {
          alert("Successfully updated the summoned");
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    },
    processImg(event) {
      this.summoned.img = event.target.files[0];
    },
    deleteSummoned() {
      let that = this;
      this.$axios
        .delete("mysummoned/" + this.summoned.ID)
        .then(function () {
          alert("Successfully delete the summoned");
          that.$router.go();
        })
        .catch(function (response) {
          console.log(response);
        });
    },
  },
};
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>



