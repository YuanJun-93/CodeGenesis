import axios from 'axios'
import { message } from 'ant-design-vue'

// Create Axios instance
const myAxios = axios.create({
  baseURL: 'http://localhost:8888',
  timeout: 60000,
  withCredentials: true,
})

// Global request interceptor
myAxios.interceptors.request.use(
  function (config) {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = token
    }
    return config
  },
  function (error) {
    // Handle request error
    return Promise.reject(error)
  },
)

// Global response interceptor
myAxios.interceptors.response.use(
  function (response) {
    const { data } = response
    // Handle 401 Unauthorized
    if (data.code === 40100) {
      // If not fetching user info and not currently on login page, redirect to login
      if (
        !response.request.responseURL.includes('/users/me') &&
        !window.location.pathname.includes('/user/login')
      ) {
        message.warning('Please login first')
        window.location.href = `/user/login?redirect=${window.location.href}`
      }
    }
    return response
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Handle response error
    return Promise.reject(error)
  },
)

export default myAxios
