<template>
    <div style="margin: 1vw;">
        <el-card style="width: 15vw;" shadow="never">
            <div style="display: flex; justify-content: space-around;">
                <el-avatar :size="80" :src="avatar_url" style="background-color: transparent; margin-bottom: 1vw;" />
                <RankIconSmall :score="Number(scores)" />
            </div>
            <el-link type="primary" :href="detailPage"> {{ login }} </el-link><br>
            <el-text truncated> <el-icon style="margin-right: 1vw;"><CoffeeCup /></el-icon> {{ language }} </el-text><br>
            <el-text > <el-icon style="margin-right: 1vw;"><Location /></el-icon> {{ nation }} </el-text>
        </el-card>
    </div>

</template>

<script setup>
    import { ref,reactive } from 'vue'
    import axios from 'axios'
    import RankIconSmall from './RankIconSmall.vue';
    const desc = ref('desc...')
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
    const avatar_url = ref('')
    const props = defineProps({
        login:{
            type: String,
            default: ''
        },
        language:{
            type: String,
            default: 'N/A'
        },
        nation:{
            type: String,
            default: 'N/A'
        },
        scores: {
            type: String,
            default: '9'
        }
    })
    const login = ref(props.login)
    const detailPage = ref('result?login='+login.value)
    const language = ref(props.language)
    const nation = ref(props.nation)
    const scores = ref(props.scores)


    if (login.value != null && login.value != '') {
        axios.get('https://api.github.com/users/'+login.value).then(res => {
            avatar_url.value = res.data.avatar_url
        })
    }
</script>

<style scoped>

</style>