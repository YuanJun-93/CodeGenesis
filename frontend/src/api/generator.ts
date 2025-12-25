// @ts-ignore
/* eslint-disable */
import request from '@/request'

/** Stream Generate Code POST /api/generator/stream */
export async function streamGenerator(
  body: API.GeneratorRequest,
  options?: { [key: string]: any }
) {
  return request<any>('/api/generator/stream', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}

/** Generate Code with AI POST /api/generator/use */
export async function useGenerator(body: API.GeneratorRequest, options?: { [key: string]: any }) {
  return request<API.GeneratorResponse>('/api/generator/use', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  })
}
