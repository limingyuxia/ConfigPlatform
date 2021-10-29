import { login, logout, getInfo, uploadAvatar} from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
import { Base64,encode, decode } from 'js-base64';

// import cypher from '@/utils/cypher'

const getDefaultState = () => {
  return {
    token: getToken(),
    name: '',
    avatar:  '../../icons/svg/中国移动.png',
    roles:[]
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  }
}

const actions = {
  // user login
  upAvatar({ commit }, userInfo) {
    // const { username, password, type } = userInfo

    return new Promise((resolve, reject) => {
      uploadAvatar(userInfo).then(response => {

        // const { data } = response
        console.log('WebData:', response)
        
        commit('SET_AVATAR', response)
 
        resolve()
      }).catch(error => {
        // console.log('token保存失败', error)
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
        let defaultAvatar = "https://tse1-mm.cn.bing.net/th/id/R-C.0e74e383596a31c27ea66aaa4933f3c2?rik=w0ORAUQtlQCGAw&riu=http%3a%2f%2fpic.22520.cn%2fup%2f200710%2f1594389918178495.jpeg&ehk=gSSynot63OE6DuCPuDZMOk%2flGpPCXJGzH02vx4DVeUg%3d&risl=&pid=ImgRaw&r=0"
        
        let tokenArry = response.token.split(".") 

        let playload = Base64.decode(tokenArry[1]); // 解码
        console.log("playload_1",playload)

        playload = JSON.parse(playload); 
        let userAvatar = playload.userAvatar || defaultAvatar
        let userRoles = playload.userRoles  || ["tmp"]
        let userName = playload.username  || "Null"
        
        console.log("playload_3",userAvatar,userRoles,userName)

        commit('SET_ROLES', userRoles)
        commit('SET_AVATAR', userAvatar)
        commit('SET_NAME', userName)
        commit('SET_TOKEN', response.token)
        setToken(response.token)
        resolve()
      }).catch(error => {
        // console.log('token保存失败', error)
        reject(error)
      })
    })
  },
  // get user info
  getInfo({ commit, state }) {
    console.log('获取：', commit, state)
    return 0

    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        const { data } = response

        if (!data) {
          return reject('Verification failed, please Login again.')
        }

        const { name, avatar } = data

        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        resolve(data)
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
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

