<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <n-button>
    </n-button>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { NSpace, NButton } from 'naive-ui'

@Options({
  props: {
    msg: String
  },
  components: {
    NSpace,
    NButton
  }
})
export default class HelloWorld extends Vue {
  msg!: string

}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
