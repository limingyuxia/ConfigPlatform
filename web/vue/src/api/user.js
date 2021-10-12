import request from '@/utils/request'

export function login(params) {
  return request({
    url: '/login',
    method: 'post',
    data: params
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
    data:params
  })
}


export function getCodeId() {

  return request({
    url: '/captcha/get',
    method: 'GET'

  })
}

export function emailSend(params) {

  return request({
    url: '/email/send',
    method: 'POST',
    data:params
  })
}

export function emailConfirm(params) {

  return request({
    url: '/email/confirm',
    method: 'POST',
    data:params
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
