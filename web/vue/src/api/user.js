import request from '@/utils/request'

export function login(params) {
  return request({
    url: '/login',
    method: 'post',
    data: params
    // headers: params.header
  })
}
export function updata(params) {
  return request({
    url: '/user/update',
    method: 'post',
    data: params

  })
}
export function getAuthInfo(params) {
  return request({
    url: '/user/auth2Info',
    method: 'get'

  })
}

export function resetToken() {
  return request({
    url: '/auth/refreshToken',
    method: 'post'

    // headers: params.header
  })
}
export function thirdLogin(params, url) {
  return request({
    url: url,
    method: 'get'
    // params: { params }
    // data: params
    // headers: params.header
  })
}

export function reCode() {
  return request({
    url: 'https://api.btstu.cn/sjbz/api.php?lx=dongman&format=json',
    method: 'GET'

  })
}

export function confirmCode(params) {
  return request({
    url: '/captcha/confirm',
    method: 'POST',
    data: params
  })
}

export function emailSend(params) {
  return request({
    url: '/email/send',
    method: 'POST',
    data: params
  })
}

export function registerUser(params) {
  return request({
    url: '/register',
    method: 'POST',
    data: params
  })
}

// 上传头像
export function uploadAvatar(params) {
  return request({
    url: '/user/avatar/update',
    method: 'POST',
    data: params
    // headers:{"enctype":"multipart/form-data" }
  })
}

export function emailConfirm(params) {
  return request({
    url: '/email/confirm',
    method: 'POST',
    data: params
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get'
    // headers: { token }
  })
}

export function logout() {
  return request({
    url: '/vue-admin-template/user/logout',
    method: 'post'
  })
}
