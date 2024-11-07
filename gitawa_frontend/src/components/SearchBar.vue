<template>
    <div style="text-align: center;">
        <el-text size="large">awa</el-text>
        <div class="mt-4">
            <el-input
            size="large"
            v-model="devInput"
            style="max-width: 65vw"
            class="input-with-select"
            >
                <template #prepend>
                    <el-dropdown :hide-on-click="false" style="margin-top: 0.33vh;">
                        <span style="cursor: pointer; outline: 0; align-items: center;">
                            <el-icon size="large"><Setting /></el-icon>
                        </span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>
                                    <el-input
                                        v-model="langInput"
                                        style="width: 20vw; margin-right: 1vw;"
                                        placeholder="Languages"
                                        :disabled="!disabledLanguage"
                                    />
                                    <el-switch v-model="disabledLanguage" />
                                </el-dropdown-item>
                                <el-dropdown-item>
                                    <el-input
                                        v-model="nationInput"
                                        style="width: 20vw; margin-right: 1vw;"
                                        placeholder="Nation"
                                        :disabled="!disabledNation"
                                    />
                                    <el-switch v-model="disabledNation" />
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
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
    import axios from 'axios'

    const disabledLanguage = ref(false)
    const disabledNation = ref(false)

    const router = useRouter()

    const param = new URLSearchParams(window.location.search)


    const devInput = ref('')
    const langInput = ref('')
    const nationInput = ref('')

    devInput.value = param.get('dev')
    langInput.value = param.get('lang')
    nationInput.value = param.get('nation')

    if (langInput.value != '' && langInput.value!= null) {
        disabledLanguage.value = true
    }
    if (nationInput.value != '' && nationInput.value!= null) {
        disabledNation.value = true
    }


    const hasDev = ref(false)
    const search = () => {
        checkDev()
        async function checkDev() {
            if (devInput.value == '' || devInput.value == null || disabledLanguage.value || disabledNation.value) {
                searchBy()
            }
            await axios.get('https://api.github.com/users/'+devInput.value).then(res => {
                if (res.status == 200) {
                    hasDev.value = true
                }
            })
            if (hasDev.value && !disabledLanguage.value && !disabledNation.value) {
                router.replace({
                    path: '/result',
                    query: {
                        login: devInput.value,
                    }
                }).then(() => {
                    // 重新刷新页面
                    location.reload()
                })
            } else {
                searchBy()
            }

        }
        function searchBy() {
            if (disabledLanguage.value == false) {
                langInput.value = null
            }
            if (disabledNation.value == false) {
                nationInput.value = null
            }
            router.replace({
                path: '/search',
                query: {
                    dev: devInput.value,
                    lang: langInput.value,
                    nation: nationInput.value,
                }
            }).then(() => {
                // 重新刷新页面
                location.reload()
            })
        }
    }
</script>