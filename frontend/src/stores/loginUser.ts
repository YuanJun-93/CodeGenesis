import { ref } from 'vue'
import { defineStore } from 'pinia'
import { getLoginUser } from '@/api/user'

/**
 * 登录用户信息
 */
export const useLoginUserStore = defineStore('loginUser', () => {
  // 默认值
  const loginUser = ref<API.LoginUserResponse>({
    userName: '未登录',
  } as API.LoginUserResponse)

  // 获取登录用户信息
  async function fetchLoginUser() {
    try {
      const res = await getLoginUser()
      if (res.data.code === 0 && res.data.data) {
        loginUser.value = res.data.data
      }
    } catch (e) {
      // 捕获异常，比如 401 未登录，避免阻断页面加载
      loginUser.value = {
        userName: '未登录',
      } as API.LoginUserResponse
    }
  }

  // 更新登录用户信息
  function setLoginUser(newLoginUser: any) {
    loginUser.value = newLoginUser
  }

  return { loginUser, fetchLoginUser, setLoginUser }
})
