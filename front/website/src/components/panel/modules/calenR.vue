<template>
    <n-card :bordered="false">
        <v-chart class="chart" :option="option" />
    </n-card>
</template>

<style scoped>
.chart {
    width: 100%;
}

.n-card {
    margin-top: -10px;
    height: 120px;
}
</style>

<script>
import { NCard } from 'naive-ui'
import { use } from 'echarts/core';
import { PieChart } from 'echarts/charts';
import { TitleComponent } from 'echarts/components';
import { LabelLayout } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import VChart, { THEME_KEY } from 'vue-echarts';
import { reactive, defineComponent } from 'vue';
use([TitleComponent, PieChart, LabelLayout, CanvasRenderer])

const option = reactive({
    title: {
        text: 'Work OR Rleax',
        left: 'center',
        textStyle: {
            fontSize: 8
        }
    },
    series: [
        {
            name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
                borderRadius: 10,
                borderColor: '#fff',
                borderWidth: 2
            },
            labelLine: {
                show: false
            },
            data: [
                { value: 1048 },
                { value: 735 },
            ]
        }
    ]
})

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
    },
    setup() {
        return { option };
    },
});
</script>