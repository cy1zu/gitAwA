<template>
    <div style="margin: 1vw;">
        <el-card style="width: 25vw" shadow="never">
            <el-text size="large"> 
                <el-icon ><Collection /></el-icon> 
                <el-link type="primary" :href="repoLink" target="_blank" style="margin-left: 1.5vh;" >{{showTitle}} </el-link> 
            </el-text>
            <div style="margin-top: 1vh;">
                <el-text size="small" type="info" truncated> {{description}} </el-text>
            </div>
            <div style="margin-top: 1vh;">
                <el-text size="small"> <el-icon><Star /></el-icon> {{stars}} </el-text>
                <el-text size="small" style="margin-left: 2vw;"> <el-icon><Plus /></el-icon> {{cons}} </el-text>
            </div>

        </el-card>
    </div>

</template>

<script setup>
    import { ref,reactive } from 'vue'
    import axios from 'axios'

    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')

    const props = defineProps({
        title:{
            type: String,
            default: 'Title'
        },
        repoLink:{
            type: String,
            default: ''
        },
        fork:{
            type: Boolean,
            default: true
        },
        cons:{
            type: Number,
            default: 0.987654321
        }
    })
    const title = ref(props.title)
    const cons = ref(props.cons)
    cons.value = (cons.value * 100).toFixed(2) + '%'

    const stars = ref('')
    const description = ref('')
    
    axios.get('https://api.github.com/repos/'+ title.value).then((res) => {
        description.value = res.data.description
        if (Number(res.data.stargazers_count) > 1000 && Number(res.data.stargazers_count) < 1000000) {
            stars.value = (Number(res.data.stargazers_count) / 1000).toFixed(1) + 'k'
        } else if (Number(res.data.stargazers_count) >= 1000000) {
            stars.value = (Number(res.data.stargazers_count) / 1000000).toFixed(1) + 'M'
        } else {
            stars.value = res.data.stargazers_count
        }
    })

    const repoLink = ref(props.repoLink)
    repoLink.value = 'https://github.com/' + title.value

    const showTitle = ref(title.value)
    if (props.fork == false) {
        showTitle.value = title.value.split('/')[1]
    }
    
    


    // const description = ref(props.description)
    // if (description.value.length > 64) {
    //     description.value = description.value.substring(0, 64) + '...'
    // }


</script>