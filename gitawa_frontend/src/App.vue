<template>
    <div class="common-layout">
        <el-container style="min-height: 95vh;">
            <el-header height="1vh">
                <el-text size="large">HeaderOfGitAwA</el-text>
            </el-header>
            <el-divider />

            <el-main>
                <RouterView />
            </el-main>

            <el-divider style="margin-bottom: 1vh;" />
            <el-footer  height="1vh">
                <el-text size="large">FooterOfGitAwA</el-text>
            </el-footer>

            
        </el-container>
    </div>
</template>

<script setup>
    import { RouterView } from 'vue-router'
</script>

<style scoped>

</style>
