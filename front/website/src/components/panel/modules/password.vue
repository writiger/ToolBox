<template>
    <div class="findPwd">
        <n-grid :cols="5">
            <n-grid-item :span="2">
                <n-dropdown trigger="hover" placement="bottom-start" :options="op" @select="handleSelect">
                    <n-button size="tiny"> 你要查询的密码用途 </n-button>
                </n-dropdown>
                <br>
                <br>
                <n-button size="tiny" @click="showModal = true">添加</n-button>
                <n-modal v-model:show="showModal">
                    <n-card title="保存密码" style="width:600px;text-align: center;">
                        <n-space vertical>
                            <n-input v-model:value="newpwd1" placeholder="输入你的密码"></n-input>
                            <n-input v-model:value="newpwd2" placeholder="再次输入你的密码"></n-input>
                            <n-input v-model:value="newpwdFor" placeholder="输入这个密码的用途"></n-input>
                            <n-button size="tiny">上传</n-button>
                        </n-space>
                    </n-card>
                </n-modal>
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
    NDropdown,
    NModal,
    NCard,
    NSpace
} from "naive-ui"
@Options({
    components: {
        NGrid,
        NGridItem,
        NInput,
        NButton,
        NModal,
        NDropdown,
        NCard,
        NSpace
    }
})
export default class password extends Vue {
    value = ""
    showModal = false
    newpwd1 = ""
    newpwd2 = ""
    newpwdFor = ""
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