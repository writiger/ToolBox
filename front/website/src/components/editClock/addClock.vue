<template>
    <n-card title="新增打卡">
        <n-space>
            描述:
            <n-input size="tiny" v-model:value="newClock.name"></n-input>
        </n-space>
        <n-space>
            次数:
            <n-input-number :show-button="false" size="tiny" v-model:value="newClock.achieve_target" />
        </n-space>
        <n-space>
            循环:<n-input size="tiny" v-model:value="newClock.time"></n-input>

            <span style="font-size:10px"> 例:3001000000（每一小时30分钟）</span>
            <span style="font-size:10px"> 例:0000000003（每周三）</span>
            <span style="font-size:10px"> 例:0000011200（每12月1日）</span>
            <span style="font-size:10px;color:brown;"> 实在没想好咋处理这里，啥时候想到了啥时候换</span>
        </n-space>
        <n-space>
            <n-button size="tiny">新增</n-button>
        </n-space>
    </n-card>

</template>

<style scoped>

</style>

<script lang='ts'>
import {
    NCard,
    NSpace,
    NButton,
    NInput,
    NInputNumber,
    NCascader,
} from "naive-ui";
import { Options, Vue } from "vue-class-component"
@Options({
    components: {
        NCard,
        NSpace,
        NButton,
        NInput,
        NCascader,
        NInputNumber,
    }
})
export default class addClock extends Vue {
    newClock = {
        name: "",
        time: "",
        achieve_target: "",
    }
}

</script>