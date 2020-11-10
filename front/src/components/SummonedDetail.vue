<template>
  <div>
    <h1>One Summoned details</h1>
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
      <img :src="this.$axios.defaults.baseURL + '/img/' + summoned.Img" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    summoned: {},
  },
};
</script>