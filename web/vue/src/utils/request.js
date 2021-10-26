import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_HTTP_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    // 配置http

    var protocolStr = document.location.protocol

    if (protocolStr === 'http:') {
      config.baseURL = 'http://' + process.env.VUE_APP_BASE_HTTP_API
      console.log('protocol = ' + protocolStr)
    } else if (protocolStr === 'https:') {
      config.baseURL = 'https://' + process.env.VUE_APP_BASE_HTTPS_API
      console.log('protocol = ' + protocolStr)
    } else {
      config.baseURL = 'http://' + process.env.VUE_APP_BASE_HTTP_API
      console.log('other protocol')
    }

    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['Authorization'] = 'Bearer ' + getToken()
      // config.headers['Authorization'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    console.log('webDataAll:', response)

    const res = response.data

    const statusCode = response.status

    // if the custom code is not 20000, it is judged as an error.
    if (statusCode !== 200) {
      Message({
        message: '[' + statusCode + ']' + res.msg || 'Error',
        type: 'error',
        duration: 5 * 1000
      })

      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (statusCode === 40004 || statusCode === 50012 || statusCode === 50014) {
        // to re-login
        // 再次获取token

        // 成功 再次请求

        MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
            location.reload()
          })
        })
      }
      return Promise.reject(new Error(res.msg || 'Error'))
    } else {
      return res
    }
  },
  error => {
    // console.log("errorData",error)

    if (typeof (error.response) === 'undefined') {
      Message({
        message: error,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(error)
    }

    const statusCode = error.response.status
    const res = error.response.data

    if (statusCode === 400) {
      Message({
        message: '[' + res.errcode + ']' + res.errmsg || 'Error',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(error)
    } else { // 其他错误
      Message({
        message: error,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(error)
    }
  }
)

export default service
