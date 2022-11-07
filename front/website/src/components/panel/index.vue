<template>
    <div class="panel">
        <n-grid x-gap="48" :cols="3">
            <n-grid-item>
                <n-card class="firstThree" title="打卡" hoverable>
                    卡片内容
                </n-card>
            </n-grid-item>
            <n-grid-item>
                <n-card class="firstThree" title="查询密码" hoverable>
                    卡片内容
                </n-card>
            </n-grid-item>
            <n-grid-item>
                <n-card class="firstThree" title="番茄钟（最少20分钟" hoverable>
                    <Tomato></Tomato>
                </n-card>
            </n-grid-item>
        </n-grid>
        <br>
        <n-grid x-gap="48" :cols="4">
            <n-grid-item :span="1">
                <n-card class="calendar" title="日程" hoverable>
                    卡片内容
                </n-card>
            </n-grid-item>
            <n-grid-item :span="3">
                <n-carousel effect="card" prev-slide-style="transform: translateX(-150%) translateZ(-800px);"
                    next-slide-style="transform: translateX(50%) translateZ(-800px);" style="height: 240px"
                    :show-dots="false">
                    <n-carousel-item :style="{ width: '70%' }">
                        <img class="carousel-img"
                            src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg">
                    </n-carousel-item>
                    <n-carousel-item :style="{ width: '70%' }">
                        <img class="carousel-img"
                            src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg">
                    </n-carousel-item>
                    <n-carousel-item :style="{ width: '70%' }">
                        <img class="carousel-img"
                            src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg">
                    </n-carousel-item>
                    <n-carousel-item :style="{ width: '70%' }">
                        <img class="carousel-img"
                            src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg">
                    </n-carousel-item>
                </n-carousel>
                <chart></chart>
            </n-grid-item>
        </n-grid>
    </div>
</template>

<style scoped>
.panel {
    margin: 40px 68px 0 68px;
}

.firstThree {
    height: 150px;
}

.carousel-img {
    margin: 0 auto;
    width: 100%;
    height: 80%;
    object-fit: cover;
}

.calendar {
    min-height: 500px;
}
</style>

<script lang='ts'>
import {
    NGrid,
    NGridItem,
    NCard,
    NCarousel,
    NCarouselItem,
} from 'naive-ui'
import Tomato from './modules/tomato.vue'
import Chart from './modules/chart.vue'
import { Options, Vue } from "vue-class-component"
@Options({
    components: {
        NGrid,
        NGridItem,
        NCard,
        NCarousel,
        NCarouselItem,
        Chart,
        Tomato
    }
})
export default class Panel extends Vue {
}

</script>