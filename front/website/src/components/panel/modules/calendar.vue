<template>
    <n-grid :cols="2">
        <n-grid-item>
            <calen-l></calen-l>
        </n-grid-item>
        <n-grid-item>
            <calen-r></calen-r>
        </n-grid-item>
    </n-grid>
</template>

<style scoped>

</style>

<script lang='ts'>
import { NGrid, NGridItem } from 'naive-ui'
import CalenL from './calenL.vue'
import CalenR from './calenR.vue'
import { Options, Vue } from "vue-class-component"
@Options({
    components: {
        NGrid,
        NGridItem,
        CalenL,
        CalenR
    }
})
export default class calendar extends Vue {
}

</script>