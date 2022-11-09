<template>
    <div>
        <n-grid :cols="5">
            <n-grid-item :span="4">
                <n-space vertical>
                    <n-slider :disabled="isDisabled" v-model:value="stu" :step="5" :min="0" :max="35" />
                    <n-slider :disabled="isDisabled" v-model:value="rlx" :step="1" :min="0" :max="10" />
                </n-space>
            </n-grid-item>
            <n-grid-item>
                <n-button class="but" circle @click="startTomato">
                    <template #icon>
                        <n-icon :size="26" :color="isStart">
                            <CaretRight20Regular />
                        </n-icon>
                    </template>
                </n-button>
            </n-grid-item>
        </n-grid>
    </div>
</template>

<style scoped>
.but {
    margin: 5px 0 0 15px;
}
</style>

<script lang='ts'>
import {
    NIcon,
    NSpace,
    NSlider,
    NGrid,
    NGridItem,
    useMessage,
    NButton,
} from 'naive-ui'
import { Options, Vue } from "vue-class-component"
import { CaretRight20Regular } from '@vicons/fluent'

@Options({
    components: {
        NIcon,
        NSlider,
        NSpace,
        NGrid,
        NGridItem,
        CaretRight20Regular,
        NButton,
    }
})
export default class tomato extends Vue {
    message = useMessage()
    stu = 25
    rlx = 5
    isStart = "green"
    isRlx = false
    targetS = 0
    targetR = 0
    timer = 0
    pause = false
    isDisabled = false
    tiktik() {
        if (!this.isRlx) {
            if (this.stu != 0) {
                this.stu--
            } else if (this.stu == 0) {
                this.message.info(
                    "休息！",
                    { duration: 5000 }
                )
                this.stu--
                this.isRlx = true
            }
        } else {
            this.rlx--
            if (this.rlx < 0) {
                this.isRlx = false
                this.stu = this.targetS
                this.rlx = this.targetR
                this.message.info(
                    "休息结束",
                    { duration: 5000 }
                )
            }
        }
    }
    checkTime() {
        this.targetR = this.rlx
        this.targetS = this.stu
        if (this.targetS < 20) {
            this.stu = 20
            this.message.info(
                "好歹20分钟吧",
                { duration: 5000 }
            )
            return
        }
        if (this.targetR < 3) {
            this.rlx = 3
            this.message.info(
                "多歇一会？",
                { duration: 5000 }
            )
            return
        }
    }
    startTomato() {
        if (this.rlx == this.stu && this.stu == 0) {
            this.rlx = 5
            this.stu = 25
        }
        if (this.isStart == "green") {
            if (!this.pause) {
                this.checkTime()
            }
            this.isStart = "red"
            this.timer = window.setInterval(this.tiktik, 60000)
            this.isDisabled = true
        } else {
            this.pause = true
            this.isStart = "green"
            this.isDisabled = false
            clearInterval(this.timer)
        }
    }
}

</script>