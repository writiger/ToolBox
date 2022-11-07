<template>
    <div class="page">
        <div class="container">
            <div class="right">
                <div class="form">

                    <label for="email"><i>
                            <n-icon>
                                <MdPerson />
                            </n-icon>
                        </i> Email
                    </label>
                    <input type="email" id="email">
                    <div class="Line"></div>
                    <label for="password1"><i>
                            <n-icon>
                                <LockOpen />
                            </n-icon>
                        </i> 输入密码</label>
                    <input type="password" id="password1">
                    <div class="Line"></div>
                    <label for="password2"><i>
                            <n-icon>
                                <LockOpen />
                            </n-icon>
                        </i> 再次输入密码</label>
                    <input type="password" id="password2">
                    <div class="Line"></div>
                    <label for="code"><i>
                            <n-icon>
                                <MessageCircle />
                            </n-icon>
                        </i> 验证码</label>
                    <div style="width: 100px;">
                        <n-grid x-gap="12" :col="24">
                            <n-grid-item>
                                <input style="width: 180px;" maxlength="6" type="text" id="code">
                            </n-grid-item>
                            <n-grid-item :offset="7">
                                <n-button quaternary type="info" class="button">
                                    <template #icon>
                                        <n-icon>
                                            <cash-icon />
                                        </n-icon>
                                    </template>获取
                                </n-button>
                            </n-grid-item>
                        </n-grid>
                    </div>
                    <div class="Line"></div>
                    <input class="sub" type="submit" id="submit" value="注册">
                </div>
            </div>
            <div class="left">
                <div class="login">Register</div>
                <div class="eula">这里将会展示我收集的地狱笑话（你自己显示点别的吧 :-D
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
@import url("https://rsms.me/inter/inter-ui.css");

::selection {
    background: #2d2f36;
}

::-webkit-selection {
    background: #2d2f36;
}

::-moz-selection {
    background: #2d2f36;
}

.button {
    height: 25px;
}

.Line {
    width: 160px;
    height: 1px;
    background: linear-gradient(244deg, rgba(0, 255, 255, 1) 50%, rgba(255, 255, 255, 0) 100%);
}

.page {
    background: #e2e2e5;
    display: flex;
    flex-direction: column;
    height: calc(90%);
    position: absolute;
    place-content: center;
    width: calc(100%);
}

@media (max-width: 767px) {
    .page {
        height: auto;
        margin-bottom: 20px;
        padding-bottom: 20px;
    }
}

.container {
    display: flex;
    height: 420px;
    margin: 0 auto;
    width: 640px;
}

@media (max-width: 767px) {
    .container {
        flex-direction: column;
        height: 630px;
        width: 320px;
    }
}

.left {
    background: white;
    height: calc(100% - 40px);
    top: 20px;
    position: relative;
    width: 50%;
}

@media (max-width: 767px) {
    .left {
        height: 100%;
        left: 20px;
        width: calc(100% - 40px);
        max-height: 270px;
    }
}

.login {
    font-size: 50px;
    font-weight: 900;
    margin: 50px 40px 40px;
}

.eula {
    color: #999;
    font-size: 14px;
    line-height: 1.5;
    margin: 40px;
}

.right {
    background: #474a59;
    box-shadow: 0px 0px 40px 16px rgba(0, 0, 0, 0.22);
    color: #f1f1f2;
    position: relative;
    width: 50%;
}

@media (max-width: 767px) {
    .right {
        flex-shrink: 0;
        height: 100%;
        width: 100%;
        max-height: 350px;
    }
}


path {
    fill: none;
    stroke: url(#linearGradient);
    stroke-width: 4;
    stroke-dasharray: 240 1386;
}

.form {
    margin: 40px;
    position: absolute;
    width: 250px;
}

label {
    color: #c2c2c5;
    display: block;
    font-size: 14px;
    height: 16px;
    margin-top: 20px;
    margin-bottom: 5px;
}

input {
    background: transparent;
    border: 0;
    color: #f2f2f2;
    font-size: 16px;
    height: 15px;
    line-height: 30px;
    outline: none !important;
    width: 100%;
}

input::-moz-focus-inner {
    border: 0;
}

#submit {
    color: #707075;
    margin-top: 40px;
    transition: color 300ms;
}

#submit:hover {
    color: #f2f2f2;
}

#submit:active {
    color: #d0d0d2;
}

.sub {
    text-shadow: 0 0 1px rgba(195, 241, 241, 0.521);
    border-bottom: 1px solid rgba(0, 0, 0, 0.233);
    height: 30px;
}
</style>

<script lang='ts'>
import { Options, Vue } from 'vue-class-component';
import { NButton, NGrid, NGridItem, NIcon } from 'naive-ui'
import { CashOutline as CashIcon, LockOpen } from '@vicons/ionicons5'
import { MdPerson } from '@vicons/ionicons4'
import { MessageCircle } from '@vicons/tabler';
@Options({
    components: {
        NButton,
        NGrid,
        NGridItem,
        NIcon,
        CashIcon,
        MdPerson,
        LockOpen,
        MessageCircle,
    }
})
export default class register extends Vue {

}

</script>