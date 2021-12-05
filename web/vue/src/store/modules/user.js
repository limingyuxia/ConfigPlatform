import { login, logout, getInfo, uploadAvatar, thirdLogin, getAuthInfo, resetToken, updata } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
// import { Base64, encode, decode } from 'js-base64'

// import cypher from '@/utils/cypher'

const getDefaultState = () => {
  return {
    token: getToken(),
    email: '',
    name: '',
    avatar: '../../icons/svg/中国移动.png',
    roles: [],
    phone: '',
    authInfo: {
      'github_avatar': '',
      'github_username': '',
      'qq_avatar': '',
      'qq_username': '',
      'wechat_avatar': '',
      'wechat_username': '',
      'weibo_avatar': '',
      'weibo_username': ''
    },
    basicInfo: {
      gender: 0,
      nickname: 'Null',
      region: 'Null'
    }

  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_AUTHINFO: (state, authInfo) => {
    state.authInfo = authInfo
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },

  SET_PHONE: (state, phone) => {
    state.phone = phone
  },
  SET_EMAIL: (state, email) => {
    state.email = email
  },
  SET_NAME: (state, name) => {
    state.name = name
  },

  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_BASICINFO: (state, basicInfo) => {
    state.basicInfo = basicInfo
  }
}

const actions = {
  async updata({ commit }, userInfo) {
    console.log('更新数据', userInfo)
    var that = this
    // 获取原始数据
    if (userInfo.chType !== 'basicInfo' && userInfo.chType !== 'email') {
      return Promise.reject('修改类型错误[' + userInfo.chType + ']')
    }
    await this.dispatch('user/getInfo')
    console.log('获取数据：', this.getters)
    if (userInfo.chType === 'basicInfo') {
      // 生成数据
      var changeData = {
        'email': this.getters.email,
        'gender': userInfo.data.gender.model,
        'nickname': userInfo.data.nickname.model,
        'phone': this.getters.phone,
        'region': userInfo.data.region.model,
        'username': this.getters.name
      }
    } else if (userInfo.chType === 'email') {
      var changeData = {
        'email': userInfo.data.account,
        'gender': this.getters.basicInfo.gender,
        'nickname': this.getters.basicInfo.nickname,
        'phone': this.getters.phone,
        'region': this.getters.basicInfo.region,
        'username': this.getters.name
      }
    }

    return new Promise((resolve, reject) => {
      updata(changeData).then(async response => {
        await that.dispatch('user/getInfo')

        resolve()
      }).catch(error => {
        console.log('更新失败', error)
        reject(error)
      })
    })

    // await that.dispatch('user/getInfo')
  },
  // 第三方登录
  thirdLogin({ commit }, loginInfo) {
    const that = this
    console.log('第三方登录：', loginInfo.Code)
    console.log('登录类型', loginInfo.typy)
    let baseurl

    if (loginInfo.typy === 'qq') {
      baseurl = '/qqLogin' + loginInfo.Code
    } else if (loginInfo.typy === 'github') {
      baseurl = '/githubLogin' + loginInfo.Code
    }
    return new Promise((resolve, reject) => {
      thirdLogin(loginInfo.Code, baseurl).then(response => {
        // const { data } = response
        console.log('登录:', response)

        commit('SET_TOKEN', response.token)
        setToken(response.token)
        resolve()
      }).catch(error => {
        console.log('第三方登录失败', error)
        reject(error)
      })
    })
  },

  // 上传头像
  upAvatar({ commit }, userInfo) {
    const that = this
    /*
    var  avatarFile = ""
    userInfo.forEach((value, key) => {
      if(key == "file"){
        avatarFile = value
      }
    })
console.log("file",avatarFile)
*/

    // store.dispatch('user/getInfo')
    // this.getInfo()
    // commit('SET_AVATAR', userInfo)

    console.log('上传头像：', userInfo)

    return new Promise((resolve, reject) => {
      uploadAvatar(userInfo).then(async response => {
        // const { data } = response
        console.log('WebData_SET_AVATAR:', response)

        await that.dispatch('user/getInfo')
        resolve()
      }).catch(error => {
        console.log('上传头像失败：', error)
        reject(error)
      })
    })
  },

  // user login
  login({ commit }, userInfo) {
    // const { username, password, type } = userInfo

    return new Promise((resolve, reject) => {
      login(userInfo).then(response => {
        // const { data } = response
        console.log('WebData:', response)

        commit('SET_TOKEN', response.token)
        setToken(response.token)
        resolve()
      }).catch(error => {
        // console.log('token保存失败', error)
        reject(error)
      })
    })
  },
  getAuthInfo({ commit, state }) {
    console.log('第三方信息_1：')
    return new Promise((resolve, reject) => {
      getAuthInfo(state.token).then(response => {
        console.log('第三方信息：', response)

        commit('SET_AUTHINFO', response)
        resolve(response)
      }).catch(error => {
        console.log('第三方信息_error：', error)
        reject(error)
      })
    })
  },
  // get user info
  getInfo({ commit, state }) {
    console.log('获取：', commit, state)

    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        // const { data } = response
        const defaultAvatar = 'https://tse1-mm.cn.bing.net/th/id/R-C.0e74e383596a31c27ea66aaa4933f3c2?rik=w0ORAUQtlQCGAw&riu=http%3a%2f%2fpic.22520.cn%2fup%2f200710%2f1594389918178495.jpeg&ehk=gSSynot63OE6DuCPuDZMOk%2flGpPCXJGzH02vx4DVeUg%3d&risl=&pid=ImgRaw&r=0'

        if (!response) {
          return reject('Verification failed, please Login again.')
        }
        console.log('userInfo:', response)
        const { username, photo, email, phone } = response
        let avatar = photo || defaultAvatar
        const userEmail = email
        const userPhone = phone
        if (avatar !== defaultAvatar) {
          avatar = 'http://' + process.env.VUE_APP_BASE_HTTP_API + avatar
        }
        console.log('avatar', avatar)
        commit('SET_ROLES', [])

        commit('SET_PHONE', userPhone)
        commit('SET_EMAIL', userEmail)
        commit('SET_NAME', username)
        commit('SET_AVATAR', avatar)
        // 基本信息
        var basicInfo = {
          gender: response.gender || 0,
          nickname: response.nickname,
          phone: response.phone,
          region: response.region
        }

        commit('SET_BASICINFO', basicInfo)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    removeToken() // must remove  token  first
    resetRouter()
    commit('RESET_STATE')

    return
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise((resolve, reject) => {
      // 重置token
      resetToken().then((response) => {
        removeToken() // must remove  token  first
        console.log('刷新token：', response)
        commit('SET_TOKEN', response.token)
        setToken(response.token)
        resolve()
      }).catch(error => {
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

