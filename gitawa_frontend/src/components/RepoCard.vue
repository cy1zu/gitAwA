<template>
    <div style="margin: 1vw;">
        <el-card style="width: 25vw" shadow="never">{{title}}</el-card>
    </div>

</template>

<script setup>
    import { ref,reactive } from 'vue'
    const input = defineProps({
        title:{
            type: String,
            default: 'Title'
        },
        desc,
        stars,
        cons,
    })
</script>