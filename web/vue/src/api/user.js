import request from '@/utils/request'

export function login(params) {
  /* return {
    code: 200,
    data: token
  }
  */

  return request({
    url: '/login',
    method: 'post',
    data: params
    // headers: params.header
  })
}

export function getInfo(token) {
  return request({
    url: '/vue-admin-template/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/vue-admin-template/user/logout',
    method: 'post'
  })
}
