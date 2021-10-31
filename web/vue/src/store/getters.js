const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  
  email:state => state.user.email,
  roles:state => state.user.roles,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name
}
export default getters
