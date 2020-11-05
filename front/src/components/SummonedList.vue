<template>
  <div>
    <ul>
      <li v-for="s in summoneds" :key="s.ID">
        <p>ID: {{s.ID}}, Creator ID: {{s.UserID}}, Type: {{s.type}}, Name: {{s.name}}, People: {{s.people}}, DDL: {{s.ddl}}</p>
        <router-link :to="'summoned/'+s.ID">details</router-link>
      </li>
    </ul>
  </div>
</template>

<script>

export default {
  props: {
    summoneds: Array,
  }
}
</script>