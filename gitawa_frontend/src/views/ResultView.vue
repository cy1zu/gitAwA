<template>
    <div v-loading="loading">
        <SearchBar />
        
        <div style="display: flex; margin: 3%; margin-top: 4.5vw; margin-left: 5.5vw; margin-bottom: 0" v-show="allpage">
            <div style="width: 25vw;">
                <div style="text-align: center; margin-top: 0.7vh; margin-bottom: 1.5vh">   
                    <el-avatar :size="sizeOfWidth*0.18" :src="avatar_url" style="background-color: transparent;"/>
                </div>
                <h3 style="margin-bottom: 1vh;">{{ nickname }}</h3>
                <div style="opacity: 40%; margin-bottom: 2vh;">{{ login }}</div>
                <div> <el-icon style="opacity: 80%; margin-bottom: 2vh;" :size="20"><LocationInformation /></el-icon> <el-text>{{ nation }} </el-text> </div>
                <!-- <div><el-text>Score: {{ talentRank }}</el-text>  </div> -->
            </div>
            <el-divider direction="vertical" style="margin: 3.5vw;" />
            <div style="width: 70vw;" v-show="showDetail==true">
                <el-row>
                    <el-col :span="10" >
                        <el-text type="primary" size="large">Top Languages </el-text>
                        <el-row :span="24" style="margin: 1.5vw;">
                            <el-col :span="12">
                                <el-text type="info" v-for="v of languages">{{v[0]}} <br></el-text>
                            </el-col>
                            <el-col :span="8" style="text-align: center;">
                                <el-text type="info" v-for="v of languages">{{v[1]}}<br></el-text>
                            </el-col>
                        </el-row>
                    </el-col>
                    <el-col :span="12" style="text-align: center;">
                        <RankIcon :score="numTRank" v-if="loaded" />
                        <el-text type="warning">{{ talentRank }}</el-text>
                    </el-col>
                </el-row>
                <el-row :gutter="24" justify="start">
                    <RepoCard v-for="v of contributions" 
                    :title="v.repo_full_name" 
                    :fork="v.fork" 
                    :stars="v.stargazers_count"
                    :cons="Number(v.contributions)"/>

                </el-row>
            </div>
            <div style="width: 70vw;" v-show="showDetail==false">
                <el-empty :description="emptyDesc" />
            </div>
        </div>
    </div>
</template>
  
<script setup>
    import { ref } from 'vue'
    import axios from 'axios'
    import SearchBar from '@/components/SearchBar.vue';
    import RepoCard from '@/components/RepoCard.vue';
    import RankIcon from '@/components/RankIcon.vue';
    const loading = ref(true)
    const showDetail = ref(false)
    const allpage = ref(false)
    const emptyDesc = ref('')
    
    const sizeOfWidth = ref(window.innerWidth)

    window.addEventListener('resize', () => {
        sizeOfWidth.value = window.innerWidth
    })

    const nickname = ref('')
    const talentRank = ref('0pp')
    const nation = ref('')
    const param = new URLSearchParams(window.location.search)
    const login = ref(param.get('login'))
    const languages = ref([])
    const contributions = ref([])
    const avatar_url = ref('')
    const loaded = ref(false)

    import { LocationInformation } from '@element-plus/icons-vue'

    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
    
    setTimeout(() => {
        // console.log('/api/developers/'+login.value)
        axios.get('https://api.github.com/users/'+login.value).then(res => {
            allpage.value = true
            nickname.value = res.data.name
            avatar_url.value = res.data.avatar_url
            loading.value = false
        })
    }, 1000)

    const userData = ref()
    const numTRank = ref(0)
    axios.get('api/developers/'+login.value).then(res => {
        // console.log(res.data)
        userData.value = res.data
        if (res.data.login == null) {
            showDetail.value = false
            emptyDesc.value = res.data.msg + ', try again 30s later'
        } else {
            showDetail.value = true
            talentRank.value = Number(userData.value.talent_rank).toFixed(0).toString()+'pp'
            
            nation.value = userData.value.nation

            languages.value = userData.value.languages
            languages.value = Object.entries(languages.value).sort((a, b) => b[1] - a[1])
            let sumsOfAllLanguages = 0
            for (let i = 0; i < languages.value.length; i++) {
                sumsOfAllLanguages += languages.value[i][1]
            }
            for (let i = 0; i < languages.value.length; i++) {
                languages.value[i][1] = (languages.value[i][1] / sumsOfAllLanguages*100).toFixed(2)+'%'
            }
            languages.value = languages.value.slice(0, 3)
            // console.log(languages.value)
            contributions.value = userData.value.contributions
            loaded.value = true
            numTRank.value = Number(talentRank.value.substring(0, talentRank.value.length-2))
        }     
    })

</script>
  
<style>

</style>
  