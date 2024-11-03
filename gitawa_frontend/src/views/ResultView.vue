<template>
    <div v-loading="loading">
        <SearchBar />
        
        <div style="display: flex; margin: 3%; margin-top: 5%;">
            <div style="width: 25vw;">
                <div style="text-align: center; margin-bottom: 1.5vh">   
                    <el-avatar :size="sizeOfWidth*0.18" :src="'https://avatars.githubusercontent.com/u/11942395?v=4'" style="background-color: transparent;"/>
                </div>
                <h3 style="margin-bottom: 1vh;">{{ nickname }}</h3>
                <div style="opacity: 40%; margin-bottom: 2vh;">{{ login }}</div>
                <div> <el-icon style="opacity: 80%; margin-bottom: 2vh;" :size="20"><LocationInformation /></el-icon> <el-text>{{ localtion }} </el-text> </div>
                <div><el-text>Score: {{ talentRank }}</el-text>  </div>
            </div>
            <el-divider direction="vertical" style="margin: 3.5vw;" />
            <div style="width: 70vw;">
                Content
                <el-row :gutter="24" justify="space-around">
                    
                    <RepoCard :title="'title1'" />
                    <RepoCard />
                    <RepoCard />
                    <RepoCard />
                    
                    
                </el-row>
            </div>
        </div>
    </div>
</template>
  
<script setup>
    import { ref } from 'vue'
    import axios from 'axios'
    import SearchBar from '@/components/SearchBar.vue';
    import RepoCard from '@/components/RepoCard.vue';
    const loading = ref(true)
    const sizeOfWidth = ref(window.innerWidth)
    const nickname = ref('')
    const talentRank = ref('')
    const localtion = ref('')
    const param = new URLSearchParams(window.location.search)
    const login = ref(param.get('login'))
    setTimeout(() => {
    console.log('/api/developers/'+login.value)
    const userData = ref()
    axios.get('api/developers/'+login.value).then(res => {
        console.log(res.data)
        userData.value = res.data

        nickname.value = userData.value.name
        talentRank.value = Number(userData.value.talent_rank).toFixed(3).toString()
        localtion.value = userData.value.localtion
        if (userData.value.localtion == null) {
            localtion.value = 'N/A'
        }

    })
        loading.value = false
    }, 1000)
</script>
  
<style>

</style>
  