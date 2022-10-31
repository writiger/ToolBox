<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>{{test}}</h2>
    <n-button v-on:click="up" tertiary circle type="error">
      <template #icon>
        <n-icon>
          <Money16Regular />
        </n-icon>
      </template>
    </n-button>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import { NSpace, NButton, NIcon } from 'naive-ui'
import Money16Regular from '@vicons/fluent/Money16Regular'

@Options({
  props: {
    msg: String
  },
  components: {
    NSpace,
    NButton,
    NIcon,
    Money16Regular
  },
})

export default class HelloWorld extends Vue {

  test = 1
  msg!: string
  up() {
    this.test++
  }
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
