<template>
    <div class="page">
        <div class="container">
            <div class="left">
                <div class="login">Login</div>
                <div class="eula">这里将会展示我收集的地狱笑话（你自己显示点别的吧 :-D
                </div>
            </div>
            <div class="right">
                <div class="form">
                    <label for="email"><i>
                            <i>
                                <n-icon>
                                    <MdPerson />
                                </n-icon>
                            </i>
                        </i> Email</label>
                    <input type="email" id="email">
                    <div class="Line"></div>
                    <label for="password"><i>
                            <n-icon>
                                <LockOpen />
                            </n-icon>
                        </i> Password</label>
                    <input type="password" id="password">
                    <div class="Line"></div>
                    <input class="sub" type="submit" id="submit" value="登录">
                    <a href="">忘记密码？</a>
                    &nbsp;
                    <router-link to="/register">注册</router-link>
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

a {
    font-size: 10px;
    text-decoration: none;
    color: #d0d0d2;
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
    height: 320px;
    margin: auto;
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
    font-size: 20px;
    height: 30px;
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
}
</style>

<script lang='ts'>
import { Options, Vue } from 'vue-class-component';
import { MdPerson } from '@vicons/ionicons4'
import { LockOpen } from '@vicons/ionicons5'
import { NIcon } from 'naive-ui';

@Options({
    components: {
        MdPerson,
        LockOpen,
        NIcon
    }
})
export default class Login extends Vue {

}

</script>