<template>
    <div class="container">
        <div class="clock" v-for="(value,index) in clocks" :key="index">
            <n-card>
                | <span>{{value.name}}</span> |<span v-if="value.achieve_now >= value.achieve_target">
                    <n-icon :size="22">
                        <StarEmphasis24Filled style="margin-bottom: -5px;" color="rgb(251, 213, 60)" />
                    </n-icon>
                </span>
                <n-divider></n-divider>
                目标: <span :style="getColor(index)">{{value.achieve_now}}/{{parseAT(value.achieve_target)}}</span>
                <br>
                循环: {{getTimes(index)}}
                <br>
                <n-button size="tiny" type="error" dashed>删除</n-button>&nbsp;
                <n-button size="tiny" type="success">打卡</n-button>
            </n-card>
        </div>
    </div>
</template>

<style scoped>
.clock {
    float: left;
    width: 17%;
    box-sizing: border-box;
    padding: 10px;
    min-width: 150px;
    text-align: center;
    margin: 20px;
}

.container {
    width: 100%;
}

@media (max-width:615px) {
    .clock {
        float: left;
        width: 33%;
        box-sizing: border-box;
        padding: 10px;
        min-width: 150px;
    }
}

@media (max-width:465px) {
    .clock {
        float: left;
        width: 40%;
        box-sizing: border-box;
        padding: 10px;
        min-width: 150px;
    }
}

@media (max-width:315px) {
    .clock {
        float: left;
        width: 100%;
        box-sizing: border-box;
        padding: 10px;
    }
}
</style>

<script lang='ts'>
import { NCard, NDivider, NSpace, NButton, NIcon } from "naive-ui";
import { Options, Vue } from "vue-class-component"
import { StarEmphasis24Filled } from "@vicons/fluent"
@Options({
    components: {
        NCard,
        NDivider,
        NSpace,
        NButton,
        StarEmphasis24Filled,
        NIcon
    }
})
export default class Clock extends Vue {
    parseAT(num) {
        if (num == -1) {
            return "∞"
        }
        return String(num)
    }
    getTimes(num) {
        var str = this.clocks[num].interval
        if (str == "-1") {
            return "不用"
        }
        // var res = ""
        var min = str[0] + str[1]
        var h = str[2] + str[3]
        var day = str[4] + str[5]
        var mon = str[6] + str[7]
        var week = str[9]
        switch (week) {
            case "1":
                return "每周一"
            case "2":
                return "每周二"
            case "3":
                return "每周三"
            case "4":
                return "每周四"
            case "5":
                return "每周五"
            case "6":
                return "每周六"
            case "7":
                return "每周日"
        }
        var res = "每"
        if (mon != "00") {
            if (mon[0] != "0")
                res += "每" + mon + "月"
            else
                res += "每" + mon[1] + "月"
        }
        if (day != "00") {
            if (day[0] != "0")
                res += day
            else
                res += day[1]
        }
        if (day !== "00" || mon != "00")
            return res + "号"
        if (h != "00") {
            if (h[0] != "0")
                res += h + "小时"
            else
                res += h[1] + "小时"
        }
        if (min != "00") {
            if (min[0] != "0")
                res += min + "分钟"
            else
                res += min[1] + "分钟"
        }
        return res
    }
    getColor(num) {
        // 完成越多越绿
        var span = this.clocks[num]
        if (span.achieve_target == -1) {
            return "color:green"
        }
        if (span.achieve_now >= span.achieve_target) {
            return "color:green"
        }
        var g = span.achieve_now / span.achieve_target * 255
        return "color:rgb(" + String(255 - g) + "," + String(g) + ",0)"
    }
    clocks = [
        {
            "name": "吃饭szdas2",
            "kind": 3,
            "start": 1667113518,
            "achieve_target": 6,
            "achieve_now": 7,
            "interval": "-1",
            "status": 2
        },
        {
            "name": "吃饭",
            "kind": 1,
            "start": 1667113518,
            "achieve_target": 5,
            "achieve_now": 0,
            "interval": "3000060001",
            "status": 1
        },
        {
            "name": "吃饭",
            "kind": 1,
            "start": 1667113518,
            "achieve_target": 10,
            "achieve_now": 5,
            "interval": "0000060000",
            "status": 1
        },
        {
            "name": "吃饭",
            "kind": 2,
            "start": 1667113518,
            "achieve_target": -1,
            "achieve_now": 0,
            "interval": "0100000000",
            "status": 1
        },
    ]
}

</script>