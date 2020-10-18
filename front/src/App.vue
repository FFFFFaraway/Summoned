<template>
  <div>
  <navi/>
  <router-view></router-view>
  </div>
</template>

<script>
import navi from './components/Navigator'
export default {
  name: 'App',
  components: {
    'navi': navi,
  }
}
</script>

