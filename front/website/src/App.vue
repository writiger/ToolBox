<template>
  <Header></Header>
  <RouterView />
</template>
<script lang="ts">
import { Vue, Options } from 'vue-class-component';
import Header from '@/components/header/index.vue'
@Options({
  components: {
    Header
  }
})
export default class App extends Vue {
}
</script>

<style>
body {
  margin: 0;
  padding: 0;
  background: #e2e2e5;
}
</style>
