<template>
  <div>
    <el-select v-model="summoned.type">
      <el-option v-for="item in options" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <br />
    <label for="name">Name: </label>
    <el-input type="text" v-model="summoned.name" id="name" ></el-input>
    <br />
    <label for="desc">Description: </label>
    <el-input type="text" v-model="summoned.desc" id="desc" ></el-input>
    <br />
    <label for="people">People Needed: </label>
    <el-input type="number" v-model.number="summoned.people" id="people" ></el-input>
    <br />
    <label for="ddl">Deadline: </label>
    <el-input type="text" v-model="summoned.ddl" id="ddl" ></el-input>
    <br />
    <label for="img">Image: </label>
    <el-input type="file" @change="processImg($event)" id="img" ></el-input>
    <br />
    <el-button type="primary" @click="submit">Update</el-button>
    <el-button type="primary" @click="deleteSummoned">Delete</el-button>
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
      formData.append("img", this.summoned.Img);

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
      this.summoned.Img = event.target.files[0];
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



