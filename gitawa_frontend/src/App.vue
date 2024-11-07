<template>
    <div class="common-layout">
        <el-container style="min-height: 95vh;">
            <el-header height="1vh">
                <el-dropdown :hide-on-click="false" @visible-change="setToken()"
                style="margin-right: 2vw; margin-top: 0.33vh;">
                    <span style="color: var(--el-color-primary); cursor: pointer; outline: 0; align-items: center;">
                        <el-icon size="large"><Setting /></el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item>
                                <el-input
                                    v-model="githubTokenInput"
                                    style="width: 30vw"
                                    placeholder="Your Github Access Token"
                                />
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
                <el-text size="large">HeaderOfGitAwA</el-text>
            </el-header>
            <el-divider />

            <el-main style="padding-bottom: 0;">
                <RouterView />
            </el-main>

            <el-divider style="margin-bottom: 1vh;" />
            <el-footer  height="0">
                <el-text size="large" >FooterOfGitAwA. </el-text> <el-link :href="'https://icons8.com/'"> Icon by Icons8</el-link>
            </el-footer>

            
        </el-container>
    </div>
</template>

<script setup>
    import { RouterView } from 'vue-router'
    import { ref } from 'vue'
    import axios from 'axios'
    const githubTokenInput = ref('')
    if (window.localStorage['token'] != null) {
        githubTokenInput.value = window.localStorage['token']
    }
    
    function setToken() {
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + githubTokenInput.value
        window.localStorage['token'] = githubTokenInput.value
    }

    
    // useTokenStore.token.value = githubTokenInput.value
</script>

<style scoped>
</style>
