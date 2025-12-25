<script lang="ts" setup>
import { reactive } from 'vue'
import { userRegister } from '@/api/user'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import type { RuleObject } from 'ant-design-vue/es/form'

const formState = reactive<API.UserRegisterRequest>({
  userAccount: '',
  userPassword: '',
  checkPassword: '',
})

const router = useRouter()

// Custom validation function
async function validateCheckPassword(rule: RuleObject, value: string) {
  if (value !== formState.userPassword) {
    return Promise.reject('Passwords do not match')
  }
  return Promise.resolve()
}

/**
 * Submit Form
 * @param values
 */
const handleSubmit = async (values: any) => {
  const res = await userRegister(formState)
  // Register successful
  if (res.data.code === 0 && res.data.data) {
    message.success('Registration successful')
    router.push({
      path: '/user/login',
      replace: true,
    })
  } else {
    message.error('Registration failed: ' + res.data.msg)
  }
}
</script>
<template>
  <div id="userRegisterPage">
    <h2 class="title">CodeGenesis - User Register</h2>
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
      <a-form-item
        name="checkPassword"
        :rules="[
          { required: true, message: 'Please confirm password' },
          { min: 8, message: 'Password must be at least 8 characters' },
          { validator: validateCheckPassword, trigger: 'blur' },
        ]"
      >
        <a-input-password v-model:value="formState.checkPassword" placeholder="Confirm Password" />
      </a-form-item>
      <div class="tips">
        Already have an account?
        <RouterLink to="/user/login">Login</RouterLink>
      </div>
      <a-form-item>
        <a-button type="primary" html-type="submit" style="width: 100%">Register</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<style>
#userRegisterPage {
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
