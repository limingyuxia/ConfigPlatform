/* 验证码相关函数
1、图片验证码
*/
import { getCaptchaId, confirmCaptchaCode, emailSend, emailConfirm } from '@/api/verify'

const getDefaultState = () => {
  return {

  }
}

const state = getDefaultState()

const mutations = {

}

const actions = {
  // 更新用户数据
  updataUserInfo({ commit }, data) {
    console.log('提交总数据', data)
    // 验证验证码

    // 提交数据
  },
  // 验证验证码
  confirmCode({ commit }, data) {
    console.log('验证总数据', data)
    console.log('验证类型', data.codeType)

    if (data.codeType === 'email') {
      console.log('验证类型', data.codeType)
      var updata = {
        captcha_token: data.captcha_token,
        email_code: data.email_code
      }
      console.log('上传数据', updata)
      return new Promise((resolve, reject) => {
        console.log('调试', updata)
        // 验证邮箱
        emailConfirm(updata).then(response => {
          console.log('验证结果', response)

          resolve(updata)
        }).catch(error => {
          console.log('更新失败2', error)
          reject(error)
        })
      })
    }
  },
  sendCode({ commit }, data) {
    console.log('发送类型：', data.codeType)
    console.log('验证码id', data.captcha_id)
    console.log('captcha_num', data.captcha_num)
    console.log('账号地址', data.account)
    const updata = {
      'captcha_solution': data.captcha_num,
      'captcha_id': data.captcha_id
    }
    var p1 = new Promise(async(resolve, reject) => {
      // 验证图片验证码
      let updata = {
        'captcha_solution': data.captcha_num,
        'captcha_id': data.captcha_id
      }
      var captcha_token = ''
      confirmCaptchaCode(updata).then(response => {
        console.log('验证图片验证码结果', response)
        captcha_token = response.captcha_token
        // 发送邮箱
        updata = {
          'captcha_token': captcha_token,
          'email_address': data.account
        }

        emailSend(updata).then(response => {
          console.log('发送验证码结果', response)

          resolve(updata)
        }).catch(error => {
          console.log('更新失败2', error)
          reject(error)
        })
      }).catch(error => {
        console.log('更新失败1', error)
        reject(error)
      })
    })
    return p1
  },

  getPictureCode({ commit }) {
    // getpictureCode
    console.log('获取图片验证码')

    return new Promise((resolve, reject) => {
      getCaptchaId().then(response => {
        console.log('验证码id', response)

        const captcha_id = response.captcha_id
        let captcha_url = ''
        var protocolStr = document.location.protocol

        if (protocolStr === 'http:') {
          captcha_url = 'http://' + process.env.VUE_APP_BASE_HTTP_API + '/captcha/' + response.captcha_id + '.png'
          // console.log('protocol = ' + protocolStr)
        } else if (protocolStr === 'https:') {
          captcha_url = 'https://' + process.env.VUE_APP_BASE_HTTPS_API + '/captcha/' + response.captcha_id + '.png'
          // console.log('protocol = ' + protocolStr)
        } else {
          captcha_url = 'http://' + process.env.VUE_APP_BASE_HTTP_API + '/captcha/' + response.captcha_id + '.png'
          // console.log('other protocol')
        }
        var returnJson = {
          captcha_url: captcha_url,
          captcha_id: captcha_id
        }

        resolve(returnJson)
      }).catch(error => {
        console.log('更新失败', error)
        reject(error)
      })
    })
  }

}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

