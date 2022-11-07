<template>
    <n-grid :cols="5">
        <n-grid-item :span="2">
            <n-statistic label="你一共使用了" tabular-nums>
                <n-number-animation ref="numberAnimationInstRef" :from="0" :to="12039" />
                <template #suffix>
                    次打卡
                </template>

            </n-statistic>
        </n-grid-item>
        <n-grid-item :span=3>
            花时间、花精力、花心思，愿你每一份努力都不负所望
        </n-grid-item>
    </n-grid>
</template>

<style scoped>

</style>

<script lang='ts'>
import {
    NStatistic,
    NNumberAnimation,
    NGrid,
    NGridItem
} from "naive-ui"
import { Options, Vue } from "vue-class-component"
@Options({
    components: {
        NStatistic,
        NNumberAnimation,
        NGrid,
        NGridItem
    }
})
export default class clock extends Vue {
}

</script>