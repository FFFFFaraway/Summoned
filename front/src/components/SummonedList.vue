<template>
  <div>
    <ul>
      <li v-for="s in summoneds" :key="s.ID">
        <p>
          ID: {{ s.ID }}, Creator ID: {{ s.user_id }}, Type: {{ s.type }}, Name:
          {{ s.name }}, People: {{ s.people }}, DDL: {{ s.ddl }}, Status: {{s.status}}
        </p>
        <router-link :to="'profile/' + s.user_id">creator info</router-link> |
        <router-link :to="'summoned/' + s.ID">details</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    summoneds: Array,
  },
};
</script>