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
      <li>Status: {{ summoned.Status }}</li>
    </ul>
    <img :src="this.$axios.defaults.baseURL + '/img/' + summoned.Img" />
  </div>
</template>

<script>
export default {
  props: {
    summoned: {},
  },
};
</script>