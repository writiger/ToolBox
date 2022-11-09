<template>
    <n-grid :cols="5">
        <n-grid-item :span="4">
            <div class="column">
                <div v-for="(value,index) in today" :key="index" class="event" :id="String(index)">
                    <span style=" mix-blend-mode:difference;">{{value.info}}</span>
                </div>
            </div>
        </n-grid-item>
        <n-grid-item>
            <n-slider style="margin:3px;height:95%" :value="now" :max="workTime*60" vertical reverse :tooltip="false"
                disabled>
                <template #thumb>⬅
                </template>
            </n-slider>
        </n-grid-item>
    </n-grid>
</template>

<style scoped>
.column {
    margin: 3px;
    width: 97%;
    height: 400px;
}

.event {
    width: auto;
    padding-left: 5px;
    border: 2px solid rgb(255, 255, 255);
    border-radius: 5px;
    text-align: center;
    min-height: 20px;
}
</style>

<script lang='ts'>
import { Options, Vue } from "vue-class-component"
import { NGrid, NGridItem, NSlider } from "naive-ui"
import { AccessTime20Regular } from "@vicons/fluent"
@Options({
    components: {
        NGrid,
        NGridItem,
        NSlider,
        AccessTime20Regular
    }
})
export default class calenL extends Vue {
    // 个人信息
    workTime = 9
    startTime = 9
    today = [
        { "info": "学习", "time": 40 },
        { "info": "学习", "time": 10 },
        { "info": "学习", "time": 100 },
        { "info": "空", "time": 390 },]
    culculateEvent() {
        for (var i = 0; i < this.today['length']; i++) {
            var useLen = this.today[i]["time"] / this.workTime / 60 * 100
            var div = document.getElementById(String(i))!
            div.style.height = String(useLen * 0.9) + "%"
            div.style.backgroundColor = this.getColor()
        }
    }
    now = 0
    color = [
        "#c1cbd7",
        "#c1cbd7",
        "#afb0b2",
        "#939391",
        "#bfbfbf",
        "#e0e5df",
        "#b5c4b1",
        "#8696a7",
        "#9ca8b8",
        "#ececea",
        "#fffaf4",
        "#96a48b",
        "#7b8b6f",
        "#dfd7d7",
        "#656565",
        "#d8caaf",
        "#c5b8a5",
        "#fdf9ee",
        "#f0ebe5",
        "#d3d4cc",
        "#e0cdcf",
        "#b7b1a5",
        "#a29988",
        "#dadad8",
        "#f8ebd8",
        "#965454",
        "#6b5152",
        "#f0ebe5",
        "#cac3bb",
        "#a6a6a8",
        "#7a7281",
        "#a27e7e",
        "#ead0d1",
        "#faead3",
        "#c7b8a1",
        "#c9c0d3",
    ]
    getColor() {
        var num = parseInt(Math.random() * (this.color['length'] - 1) + '')
        return this.color[num]
    }
    culculateNow() {
        var d = new Date()
        var m = d.getMinutes()
        var h = d.getHours()
        this.now = h * 60 + m - this.workTime * 60
        console.log(this.now)
    }
    mounted() {
        this.culculateNow()
        this.culculateEvent()
        window.setInterval(this.culculateNow, 1000)
    }

}

</script>