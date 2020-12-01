<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>One Summoned details</span>
      </div>
      <ul>
        <li>summoned ID: {{ summoned.ID }}</li>
        <li>type: {{ summoned.type }}</li>
        <li>name: {{ summoned.name }}</li>
        <li>desc: {{ summoned.desc }}</li>
        <li>people: {{ summoned.people }}</li>
        <li>ddl: {{ summoned.ddl }}</li>
        <li>Status: {{ summoned.status }}</li>
      </ul>
      <div v-if="summoned.Img">
        <el-image
            style="width: 200px; height: 200px"
            :src="this.$axios.defaults.baseURL + '/img/' + summoned.Img"
            fit="contain"></el-image>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  props: {
    summoned: {},
  },
};
</script>