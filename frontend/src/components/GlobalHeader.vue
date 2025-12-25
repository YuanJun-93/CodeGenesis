<template>
  <a-layout-header class="header">
    <a-row :wrap="false">
      <!-- Left: Logo and Title -->
      <a-col flex="200px">
        <RouterLink to="/">
          <div class="header-left">
            <img class="logo" src="@/assets/logo.png" alt="Logo" />
            <h1 class="site-title">CodeGenesis</h1>
          </div>
        </RouterLink>
      </a-col>
      <!-- Center: Navigation Menu -->
      <a-col flex="auto">
        <a-menu
          v-model:selectedKeys="selectedKeys"
          mode="horizontal"
          :items="menuItems"
          @click="handleMenuClick"
        />
      </a-col>
      <!-- Right: User Operations -->
      <a-col>
        <div class="user-login-status">
          <div v-if="loginUserStore.loginUser.id">
            <a-dropdown>
              <a-space>
                <a-avatar :src="loginUserStore.loginUser.userAvatar" />
                {{ loginUserStore.loginUser.userName ?? 'Guest' }}
              </a-space>
              <template #overlay>
                <a-menu>
                  <a-menu-item @click="doLogout">
                    <LogoutOutlined />
                    Logout
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </div>
          <div v-else>
            <RouterLink to="/user/login">
              <a-button type="primary">Login</a-button>
            </RouterLink>
          </div>
        </div>
      </a-col>
    </a-row>
  </a-layout-header>
</template>

<script setup lang="ts">
import {computed, h, ref} from 'vue'
import { useRouter } from 'vue-router'
import type { MenuProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { useLoginUserStore } from '@/stores/loginUser.ts'
import { LogoutOutlined } from '@ant-design/icons-vue'
import { userLogout } from '@/api/user'

// Get login user state
const loginUserStore = useLoginUserStore()

const router = useRouter()
// Current selected menu
const selectedKeys = ref<string[]>(['/'])
// Listen to route changes, update selected menu
router.afterEach((to, from, next) => {
  selectedKeys.value = [to.path]
})

// Menu items configuration
const originItems = [
  {
    key: '/',
    label: 'Home',
    title: 'Home',
  },
  {
    key: '/admin/userManage',
    label: 'User Management',
    title: 'User Management',
  },
  /*
  {
    key: 'others',
    label: h('a', { href: 'https://github.com/YuanJun-93/CodeGenesis', target: '_blank' }, 'GitHub'),
    title: 'GitHub',
  },
  */
]

// Filter menu items
const filterMenus = (menus = [] as MenuProps['items']) => {
  return menus?.filter((menu) => {
    const menuKey = menu?.key as string
    if (menuKey?.startsWith('/admin')) {
      const loginUser = loginUserStore.loginUser
      if (!loginUser || loginUser.userRole !== 'admin') {
        return false
      }
    }
    return true
  })
}

// Routes to display in menu
const menuItems = computed<MenuProps['items']>(() => filterMenus(originItems))

// Handle menu click
const handleMenuClick: MenuProps['onClick'] = (e) => {
  const key = e.key as string
  selectedKeys.value = [key]
  // Push to page
  if (key.startsWith('/')) {
    router.push(key)
  }
}

// Redirect to login page
const toLogin = () => {
  router.push('/user/login')
}

// Logout
const doLogout = async () => {
  const res = await userLogout()
  if (res.data.code === 0) {
    loginUserStore.setLoginUser({
      userName: 'Not logged in',
    })
    message.success('Logout successful')
    await router.push('/user/login')
  } else {
    message.error('Logout failed: ' + res.data.msg)
  }
}
</script>

<style scoped>
.header {
  background: #fff;
  padding: 0 24px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo {
  height: 48px;
  width: 48px;
}

.site-title {
  margin: 0;
  font-size: 18px;
  color: #1890ff;
}

.ant-menu-horizontal {
  border-bottom: none !important;
}
</style>
