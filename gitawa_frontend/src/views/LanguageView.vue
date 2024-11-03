<template>
    <div v-loading="loading">
        <SearchBar :select="'2'"/>
        
        <div style="display: flex; margin: 3%; margin-top: 4.5vw; margin-left: 3.5vw" v-show="allpage">
            <el-row>
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
                <DevCard />
            </el-row>
            

        </div>
        
        

        
    </div>
</template>
  
<script setup>
    import { ref } from 'vue'
    import axios from 'axios'
    import SearchBar from '@/components/SearchBar.vue';
    import DevCard from '@/components/DevCard.vue';
    const allpage = ref(false)
    allpage.value = true

    const loading = ref(true)
    const param = new URLSearchParams(window.location.search)
    const q = ref(param.get('q'))
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

    setTimeout(() => {
        loading.value = false
    }, 1000)
</script>
  
<style>

</style>
  