<template>
  <div>
    <ul>
      <li v-for="s in summoneds" :key="s.ID">
        <p>ID: {{s.ID}}, Creator ID: {{s.UserID}}, Name: {{s.name}}</p>
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