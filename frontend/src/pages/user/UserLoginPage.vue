<template>
  <div id="userLoginPage">
    <h2 class="title">CodeGenesis - User Login</h2>
    <div class="desc">Create full-stack applications without writing code</div>
    <a-form :model="formState" name="basic" autocomplete="off" @finish="handleSubmit">
      <a-form-item name="userAccount" :rules="[{ required: true, message: 'Please enter username' }]">
        <a-input v-model:value="formState.userAccount" placeholder="Username" />
      </a-form-item>
      <a-form-item
        name="userPassword"
        :rules="[
          { required: true, message: 'Please enter password' },
          { min: 8, message: 'Password must be at least 8 characters' },
        ]"
      >
        <a-input-password v-model:value="formState.userPassword" placeholder="Password" />
      </a-form-item>
      <div class="tips">
        No account?
        <RouterLink to="/user/register">Register</RouterLink>
      </div>
      <a-form-item>
        <a-button type="primary" html-type="submit" style="width: 100%">Login</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script lang="ts" setup>
import { reactive } from 'vue'
import { userLogin } from '@/api/user'
import { useLoginUserStore } from '@/stores/loginUser.ts'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'

const formState = reactive<API.UserLoginRequest>({
  userAccount: '',
  userPassword: '',
})

const router = useRouter()
const loginUserStore = useLoginUserStore()

/**
 * Submit Form
 * @param values
 */
const handleSubmit = async (values: any) => {
  const res = await userLogin(values)
  // Login successful
  if (res.data.code === 0 && res.data.data) {
    localStorage.setItem('token', res.data.data.token)
    await loginUserStore.fetchLoginUser()
    message.success('Login successful')
    router.push({
      path: '/',
      replace: true,
    })
  } else {
    message.error('Login failed: ' + res.data.msg)
  }
}
</script>
<style>
#userLoginPage {
  max-width: 480px;
  margin: 0 auto;
}

.title {
  text-align: center;
  margin-bottom: 16px;
}

.desc {
  text-align: center;
  color: #bbb;
  margin-bottom: 16px;
}

.tips {
  text-align: right;
  color: #bbb;
  font-size: 13px;
  margin-bottom: 16px;
}
</style>
