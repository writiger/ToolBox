<template>
    <n-card>
        <v-chart class="chart" :option="option" />
    </n-card>
</template>
  
<script>
import { NCard } from 'naive-ui'
import { use } from 'echarts/core';
import { GridComponent, TitleComponent } from 'echarts/components';
import { LineChart } from 'echarts/charts';
import { UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import VChart, { THEME_KEY } from 'vue-echarts';
import { reactive, defineComponent } from 'vue';


use([
    GridComponent,
    LineChart,
    CanvasRenderer,
    UniversalTransition,
    TitleComponent
]);
const xData = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30]

const option = reactive({
    title: {
        text: '最近30天打卡情况',
        left: '1%'
    },
    animationDuration: 5000,
    xAxis: {
        type: 'category',
        data: xData
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [150, 230, 224, 218, 135, 147, 260],
            type: 'line'
        }
    ]
});
export default defineComponent({
    name: 'HelloWorld',
    components: {
        VChart,
        NCard
    },
    provide: {
        [THEME_KEY]: 'light',
    },
    mounted() {
        for (var i = 0; i < 30; i++) {
            option.series[0].data[i] = i
        }
    },
    setup() {
        return { option };
    },
});
</script>
  
<style scoped>
.chart {
    width: 100%;
}

.n-card {
    margin-top: -10px;
    height: 250px;
}
</style>
  