// @ts-ignore
/* eslint-disable */
import request from '@/request'

/** List Users By Page (Admin) GET /api/users/ */
export async function listUserByPage(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.ListUserByPageParams,
  options?: { [key: string]: any }
) {
  return request<API.BaseResponseUserPage>('/api/users/', {
    method: 'GET',
    params: {
      // current has a default value: 1
      current: '1',
      // pageSize has a default value: 10
      pageSize: '10',

      ...params,
    },
    ...(options || {}),
  })
}

/** Add User (Admin) POST /api/users/ */
export async function addUser(body: API.UserAddRequest, options?: { [key: string]: any }) {
  return request<API.BaseResponseLong>('/api/users/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** Get User By ID (Admin) GET /api/users/${param0} */
export async function getUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.GetUserParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.BaseResponseUserResponse>(`/api/users/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  })
}

/** Update User (Admin) PUT /api/users/${param0} */
export async function updateUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.UpdateUserParams,
  body: API.UserUpdateRequest,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.BaseResponseBoolean>(`/api/users/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  })
}

/** Delete User (Admin) DELETE /api/users/${param0} */
export async function deleteUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.DeleteUserParams,
  body: API.IdPathRequest,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.BaseResponseBoolean>(`/api/users/${param0}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  })
}

/** User Login POST /api/users/login */
export async function userLogin(body: API.UserLoginRequest, options?: { [key: string]: any }) {
  return request<API.BaseResponseLoginUserResponse>('/api/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** User Logout POST /api/users/logout */
export async function userLogout(options?: { [key: string]: any }) {
  return request<API.BaseResponseBoolean>('/api/users/logout', {
    method: 'POST',
    ...(options || {}),
  })
}

/** Get Current Logged In User GET /api/users/me */
export async function getLoginUser(options?: { [key: string]: any }) {
  return request<API.BaseResponseLoginUserResponse>('/api/users/me', {
    method: 'GET',
    ...(options || {}),
  })
}

/** User Register POST /api/users/register */
export async function userRegister(
  body: API.UserRegisterRequest,
  options?: { [key: string]: any }
) {
  return request<API.BaseResponseLong>('/api/users/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}
