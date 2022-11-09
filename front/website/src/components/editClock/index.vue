<template>
    <div class="panel">
        <n-carousel effect="card" prev-slide-style="transform: translateX(-150%) translateZ(-800px);"
            next-slide-style="transform: translateX(50%) translateZ(-800px);" style="height: 700px;" :show-dots="false">

            <n-carousel-item v-for="(day,index1) in whatDay" :key="index1" :style="{ width: '50%' }">
                <n-card class="dayCard" :title="day">
                    <div class="perEvent" v-for="(value,index2) in eventNum[index1]" :key="index2">
                        <div v-if="showEvents[index1]<=index2 && index2<showEvents[index1]+5">
                            <n-grid :x-grap="16" :cols="14">
                                <n-grid-item :span="1">
                                    {{showEvents[index1]<=index2 && index2<showEvents[index1]+5}} <n-popconfirm
                                        negative-text="算了" @positive-click="handleDeleteClick(index1,index2)"
                                        positive-text="确认">

                                        <template #trigger>
                                            <n-button text style="font-size: 28px" type="error">
                                                <n-icon>
                                                    <Delete16Regular />
                                                </n-icon>
                                            </n-button>
                                        </template>
                                        确认要删除第{{index2+1}}个事件么?
                                        </n-popconfirm>
                                </n-grid-item>
                                <n-grid-item :span="4">
                                    <n-input maxlength="10" show-count clearable v-model:value="value.info" />
                                </n-grid-item>
                                <n-grid-item :span="1">
                                    <n-button @click="addEvent(index1,index2)" small style="margin-top:7px" text
                                        type="success">
                                        <n-icon>
                                            <Add28Filled />
                                        </n-icon>
                                    </n-button>
                                </n-grid-item>
                                <n-grid-item :span="8">
                                    <n-space vertical size="small">
                                        <n-slider v-model:value="value.time" :step="5" size="tiny" />
                                        <n-input-number size="tiny" v-model:value="value.time" />
                                    </n-space>
                                </n-grid-item>
                            </n-grid>
                            <n-divider />
                        </div>
                    </div>
                    <template #footer>
                        <n-button text color="#8a2be2" @click="moveShowUp(index1)">
                            <n-icon>
                                <ArrowUp12Filled />
                            </n-icon>
                        </n-button>&nbsp;
                        <n-button size="tiny" strong secondary type="warning">SaveAll</n-button>&nbsp;
                        <n-button text color="#ff69b4" @click="moveShowDown(index1)">
                            <n-icon>
                                <ArrowDown12Filled />
                            </n-icon>
                        </n-button>
                    </template>
                </n-card>
            </n-carousel-item>
        </n-carousel>
    </div>
</template>

<style scoped>
.panel {
    margin: 40px 68px 0 68px;
}

.dayCard {
    text-align: center;
    margin: 0 auto;
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.perEvent {
    margin-bottom: 3px;
}
</style>

<script lang='ts'>
import { Options, Vue } from "vue-class-component"
import { Delete16Regular, Add28Filled, ArrowUp12Filled, ArrowDown12Filled } from "@vicons/fluent"
import {
    NSpace,
    NCarousel,
    NCarouselItem,
    NCard,
    NButton,
    NGrid,
    NGridItem,
    NInput,
    NIcon,
    NPopconfirm,
    NSlider,
    NInputNumber,
    NDivider
} from 'naive-ui'
@Options({
    components: {
        NSpace,
        NCarousel,
        NCarouselItem,
        NCard,
        NButton,
        NGrid,
        NGridItem,
        NInput,
        NIcon,
        Delete16Regular,
        NPopconfirm,
        Add28Filled,
        NSlider,
        NInputNumber,
        NDivider,
        ArrowUp12Filled,
        ArrowDown12Filled
    }
})
export default class editClock extends Vue {
    addEvent(index1, index2) {
        this.eventNum[index1].splice(index2 + 1, 0, {
            "info": "",
            "time": 0
        })
        if (this.eventNum[index1]['length'] > 5) {
            this.moveShowDown(index1)
        }
    }
    handleDeleteClick(i, j) {
        this.eventNum[i].splice(j, 1)
    }
    //为了解决轮播图高度问题不得不出此下策
    showEvents = [0, 0, 0, 0, 0, 0, 0]
    moveShowUp(index1) {
        var temp = this.showEvents[index1]
        if (temp == 0) {
            return
        }
        this.showEvents[index1] -= 1
    }
    moveShowDown(index1) {
        var temp = this.showEvents[index1]
        if (temp + 6 > this.eventNum[index1]['length']) {
            return
        }
        this.showEvents[index1] += 1

    }
    eventNum = [
        [
            { "info": "1", "time": 40 },
            { "info": "2", "time": 40 },
            { "info": "3", "time": 40 }
        ],
        [
            { "info": "4", "time": 40 },
            { "info": "5", "time": 40 },
            { "info": "6", "time": 40 }
        ],
        [
            { "info": "", "time": 0 },
        ],
        [
            { "info": "", "time": 0 },
        ],
        [
            { "info": "", "time": 0 },
        ],
        [
            { "info": "", "time": 0 },
        ],
        [
            { "info": "", "time": 0 },
        ],
    ]
    whatDay = ["星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期七",]
}

</script>