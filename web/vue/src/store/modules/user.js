import { reCode,login, logout, getInfo, uploadAvatar} from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
import { Base64,encode, decode } from 'js-base64';

// import cypher from '@/utils/cypher'

const getDefaultState = () => {
  return {
    token: getToken(),
    email:"",
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
  }
}

const actions = {
  // user login
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



    

    //store.dispatch('user/getInfo')
    //this.getInfo()
    //commit('SET_AVATAR', userInfo)


    return new Promise((resolve, reject) => {
      uploadAvatar(userInfo).then(async response => {

        // const { data } = response
        console.log('WebData_SET_AVATAR:', response)
        
        await that.dispatch('user/getInfo')
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
        /*
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
*/
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
    

    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        //const { data } = response
        let defaultAvatar = "https://tse1-mm.cn.bing.net/th/id/R-C.0e74e383596a31c27ea66aaa4933f3c2?rik=w0ORAUQtlQCGAw&riu=http%3a%2f%2fpic.22520.cn%2fup%2f200710%2f1594389918178495.jpeg&ehk=gSSynot63OE6DuCPuDZMOk%2flGpPCXJGzH02vx4DVeUg%3d&risl=&pid=ImgRaw&r=0"

        if (!response) {
          return reject('Verification failed, please Login again.')
        }
        console.log("userInfo:",response)
        const { username, photo ,email} = response
        let avatar = photo || defaultAvatar
        let userEmail = email || "Null"
        if (avatar !== defaultAvatar){
          avatar = 'http://' + process.env.VUE_APP_BASE_HTTP_API + avatar
        }
        console.log("avatar",avatar)
        commit('SET_ROLES', [])
        commit('SET_EMAIL', userEmail)
        commit('SET_NAME', username)
        commit('SET_AVATAR', avatar)
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

