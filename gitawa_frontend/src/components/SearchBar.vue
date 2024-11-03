<template>
    <div style="text-align: center;">
        <el-text size="large">awa</el-text>
        <div class="mt-4">
            <el-input
            size="large"
            v-model="githubLogin"
            style="max-width: 65vw"
            class="input-with-select"
            >
                <template #prepend>
                    <el-select v-model="select" placeholder="Select" style="width: 15vw" size="large" >
                        <el-option label="Developers" value="1" />
                        <el-option label="Languages" value="2" />
                        <el-option label="Natons" value="3" />
                    </el-select>
                </template>
                <template #append>
                    <el-button :icon="Search" @click="search" />
                </template>
            </el-input>
        </div>
    </div>
</template>

<script setup>
    import { useRouter } from 'vue-router'
    import { ref } from 'vue'
    import { Search } from '@element-plus/icons-vue'

    const select = ref('')
    const router = useRouter()
    const githubLogin = ref('')
    const search = () => {
        console.log(githubLogin.value)
        router.replace({
            path: '/result',
            query: {
                login: githubLogin.value
            }
        }).then(() => {
            // 重新刷新页面
            location.reload()
        })
    }
</script>