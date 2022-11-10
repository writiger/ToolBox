<template>
    <div id="timedate">
        <n-grid col="10">
            <n-grid-item>
                <a id="mon" style="margin-left: -120px;">January</a>
                <a id="d">1</a>,
                <a id="y">0</a>
            </n-grid-item>
            <n-grid-item :span="9">
                <a id="h">12</a> :
                <a id="m">00</a>:
                <a id="s">00</a>
            </n-grid-item>
            <n-grid-item :offset="12">
                <div style="margin-bottom: -10px;">
                    <n-avatar style="margin: 0;" :size="36" round :src=avatar></n-avatar>
                </div>
            </n-grid-item>
        </n-grid>
    </div>
</template>

<style scoped>
#timedate {
    background-color: #3b3b3b;
    font: small-caps lighter 18px/150% "Segoe UI", Frutiger, "Frutiger Linotype",
        "Dejavu Sans", "Helvetica Neue", Arial, sans-serif;
    text-align: left;
    width: auto;
    margin: 1px 20px -10px 20px;
    color: #fff;
    border-left: 3px solid #ed1f24;
    padding: 16px 16px 16px 50%;
}
</style>

<script lang='ts'>
import { Options, Vue } from "vue-class-component"
import { NGrid, NGridItem, NAvatar } from "naive-ui"
@Options({
    components: {
        NGrid,
        NGridItem,
        NAvatar,
    }
})
export default class Header extends Vue {
    avatar = require('@/assets/avatars/1.png')
    months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    tags = ["mon", "d", "y", "h", "m", "s"]
    mounted() {
        this.updateTime()
        window.setInterval(this.updateTime, 1000)
    }
    updateTime() {
        var d = new Date()
        var now: number[] = [
            d.getMonth(),
            d.getDay(),
            d.getFullYear(),
            d.getHours(),
            d.getMinutes(),
            d.getSeconds()
        ]
        for (var i = 0; i < this.tags.length; i++) {
            var temp = document.getElementById(this.tags[i])!
            if (i != 0) {
                if (now[i] < 10) {
                    temp.innerHTML = "0" + String(now[i])
                } else {
                    temp.innerHTML = String(now[i])
                }
            } else {
                temp.innerHTML = this.months[now[i]]
            }

        }
    }

}

</script>