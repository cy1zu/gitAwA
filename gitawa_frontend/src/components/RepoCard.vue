<template>
    <div style="margin: 1vw;">
        <el-card style="width: 25vw" shadow="never">
            <el-text size="large"> 
                <el-icon ><Collection /></el-icon> 
                <el-link type="primary" :herf="repoLink" style="margin-left: 1.5vh;" >{{title}} </el-link> 
            </el-text>
            <div style="margin-top: 1vh;">
                <el-text type="info" truncated> {{description}} </el-text>
            </div>
            <div style="margin-top: 2vh;">
                <el-text size="small">{{mainLanguage}}</el-text>
                <el-text size="small" style="margin-left: 1vw;"> <el-icon><Star /></el-icon> {{stars}} </el-text>
                <el-text size="small" style="margin-left: 1vw;"> <el-icon><Plus /></el-icon> {{cons}} </el-text>
            </div>

        </el-card>
    </div>

</template>

<script setup>
    import { ref,reactive } from 'vue'
    const props = defineProps({
        title:{
            type: String,
            default: 'Title'
        },
        repoLink:{
            type: String,
            default: ''
        },
        description:{
            type: String,
            default: 'desc...desc...desc...desc...desc...desc...desc...desc...desc...desc...'
        },
        mainLanguage:{
            type: String,
            default: 'Python3'
        },
        stars:{
            type: Number,
            default: 150000
        },
        cons:{
            type: Number,
            default: 0.987654321
        }
    })

    const stars = ref(props.stars)
    if (stars.value > 1000 && stars.value < 1000000) {
        stars.value = stars.value / 1000 + 'k'
    } else if (stars.value >= 1000000) {
        stars.value = stars.value / 1000000 + 'M'
    }

    const cons = ref(props.cons)
    cons.value = (cons.value * 100).toFixed(2) + '%'

    // const description = ref(props.description)
    // if (description.value.length > 64) {
    //     description.value = description.value.substring(0, 64) + '...'
    // }


</script>