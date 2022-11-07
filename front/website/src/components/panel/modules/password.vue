<template>
    <div class="findPwd">
        <n-grid :cols="5">
            <n-grid-item :span="2">
                <n-dropdown trigger="hover" placement="bottom-start" :options="op" @select="handleSelect">
                    <n-button> 你要查询的密码用途 </n-button>
                </n-dropdown>
            </n-grid-item>
            <n-grid-item :span="3">
                <n-input class="input" v-model:value="value" round type="textarea" placeholder="输入你的密钥" size="small" />
            </n-grid-item>
        </n-grid>
    </div>
</template>

<style scoped>
.findPwd {
    margin-top: -15px;
}
</style>

<script lang='ts'>
import { Options, Vue } from "vue-class-component"
import {
    NGrid,
    NGridItem,
    NInput,
    NButton,
    NDropdown
} from "naive-ui"
@Options({
    components: {
        NGrid,
        NGridItem,
        NInput,
        NButton,
        NDropdown
    }
})
export default class password extends Vue {
    value = ""
    op = [
        {
            label: '滨海湾金沙，新加坡',
            key: 'marina bay sands'
        },
        {
            label: '布朗酒店，伦敦',
            key: "brown's hotel, london"
        },
        {
            label: '亚特兰蒂斯巴哈马，拿骚',
            key: 'atlantis nahamas, nassau'
        },
        {
            label: '比佛利山庄酒店，洛杉矶',
            key: 'the beverly hills hotel, los angeles'
        }
    ]
    handleSelect(key: string | number) {
        alert(String(key) + this.value)
    }
}

</script>