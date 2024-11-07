<template>
    <div v-loading="loading">
        <SearchBar :select="'2'"/>
        
        <div style="display: flex; margin: 3%; margin-top: 4.5vw; margin-left: 3.5vw; margin-bottom: 1%;">
            <el-row v-if="loaded">
                <DevCard v-for="v of devCardDataThisPage"
                :login="v.login" :nation="v.nation" :language="v.top_languages" :scores="v.talent_rank" />
            </el-row>
            

        </div>
        <div style="width: 90vw; margin-left: 1.3vw; display: flex; justify-content: space-around;">
            <el-pagination layout="prev, pager, next" 
            :total="devCardData.length" v-model:current-page="page" @current-change="pageChange"  />
        </div>
        
        

        
    </div>
</template>
  
<script setup>
    import { ref, computed } from 'vue'
    import axios from 'axios'
    import SearchBar from '@/components/SearchBar.vue';
    import DevCard from '@/components/DevCard.vue';
    const loading = ref(true)
    const loaded = ref(false)
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
    const page = ref(1)
    const devCardData = ref([])
    let devCardDataThisPage = computed(() => {
        return devCardData.value.slice((page.value-1)*10, page.value*10)
    })

    axios.get('api/search'+window.location.search).then(res => {
        devCardData.value = res.data
        // console.log(devCardData.value)
        loading.value = false
        loaded.value = true
    })

    setTimeout(() => {
    }, 150)

    function pageChange() {
        loaded.value = false
        loading.value = true
        setTimeout(() => {
            loading.value = false
            loaded.value = true
        }, 150)
    }

</script>
  
<style>

</style>
  