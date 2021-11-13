const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,

  authInfo:state => state.user.authInfo,
  basicInfo: state => state.user.basicInfo,
  email: state => state.user.email,
  phone:state => state.user.phone,
  roles: state => state.user.roles,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name
}
export default getters
