// @ts-ignore
/* eslint-disable */
import request from '@/request'

/** 此处后端没有提供注释 GET /api/healthz */
export async function healthz(options?: { [key: string]: any }) {
  return request<API.BaseResponseString>('/api/healthz', {
    method: 'GET',
    ...(options || {}),
  })
}
