/*验证码api
1、图片验证码
*/
import request from '@/utils/request'
export function emailConfirm(params) {
  //console.log("post:",params)
  return request({
    url: '/email/confirm',
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


export function confirmCaptchaCode(params) {
    return request({
      url: '/captcha/confirm',
      method: 'POST',
      data: params
    })
  }
export function getCaptchaId() {
    return request({
      url: '/captcha/get',
      method: 'GET'
  
    })
  }