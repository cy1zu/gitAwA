<template>
    <div>
        <el-image style="width: 15vh; height: 15vh;" :src="ranks" fit="fill" />
    </div>

</template>

<script setup>
    import { ref } from 'vue'
    const props = defineProps({
        score:{
            type: Number,
            default: 0
        }
    })


    const score = ref(props.score)
    const ranks = ref('')
    if (score.value == 0 ) {
        ranks.value = '0.png'
    } else if (score.value < 50) {
        ranks.value = 'f.png'
    } else if (score.value < 1000) {
        ranks.value = 'c.png'
    } else if (score.value < 3000) {
        ranks.value = 'b.png'
    } else if (score.value < 6000) {
        ranks.value = 'a.png'
    } else if (score.value < 10000) {
        ranks.value = 'fakes.png'
    } else if (score.value < 20000) {
        ranks.value = 's.png'
    } else {
        ranks.value = 'godlike.png'
    }
</script>

<style scoped>

</style>