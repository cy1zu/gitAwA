import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useTokenStore = defineStore('githubAccessToken', () => {
    const token = ref('')

    function setToken(newToken) {
        token.value = newToken
    }
    function getToken() {
        return token.value
    }

    return { token, setToken, getToken }
}, {
    persist: {
        persist: true,
        storage: sessionStorage,
    },
})
