<template>

  <div class="login-register">
    <div class="contain">
      <div
        class="big-box"
        :class="{ active: isLogin }"
      >
        <el-form
          v-if="isLogin"
          ref="inputLogin"
          class="big-contain"
          :model="inputLogin"
          :rules="inputLoginRules"
        >

          <div class="btitle">账户登录</div>

          <div style="height: 40%;" class="bform">
            <el-form-item style="width: 50%;" prop="username">
              <el-input v-model="inputLogin.username" type="text" title="请输入用户名" placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item style="width: 50%;" prop="password">
              <el-input v-model="inputLogin.password" type="password" title="请输入密码" placeholder="请输入密码" show-password />
            </el-form-item>
          </div>

          <el-button
            round
            type="primary"
            style="width: 30%;"
            @click="login"
          >
            登录
          </el-button>

          <span id="qqLoginBtn" />

          <div class="social-signup-container">

            <div class="sign-btn" @click="wechatHandleClick('wechat')">
              <span class="wx-svg-container"><svg-icon icon-class="wechat" class="icon" /></span> 微信
            </div>
            <div class="sign-btn" @click="tencentHandleClick('tencent')">
              <span class="qq-svg-container"><svg-icon id="qqLoginBtn" icon-class="qq" class="icon" /></span> QQ
            </div>
            <div class="sign-btn" @click="getGithubAuthorizeCode()">
              <span class="qq-svg-container"><svg-icon  icon-class="github-fill" class="icon" /></span> github
            </div>
            <div class="sign-btn" @click="weiboHandleClick('tencent')">
              <span class="qq-svg-container"><svg-icon id="wbLoginBtn" icon-class="weibo" class="icon" /></span> 微博
            </div>
          </div>

        </el-form>
        <el-form
          v-else
          ref="inputRegister"
          class="big-contain"
          :model="inputRegister"
          :rules="inputRegisterRules"
        >
          <div class="btitle">创建账户</div>
          <div v-loading="inputRegister.registerLoading" style="height: 70%;" class="bform">

            <el-form-item style="width: 50%;" prop="username">
              <el-input v-model="inputRegister.username" type="text" title="请输入用户名" placeholder="请输入用户名" />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="email">
              <el-input v-model="inputRegister.email" type="email" title="请输入邮箱" placeholder="请输入邮箱" />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="codeNum">
              <el-input v-model="inputRegister.codeNum" type="text" title="请输入图形验证码" placeholder="请输入图形验证码" @input="codeNumIn"><el-button slot="append" v-loading="reCodeLoading" @click="reCodefuc()"><img onerror="default" style="width: 130px;height: 40px;" :src="src1"></el-button></el-input>
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="codeEmail">
              <el-input v-model="inputRegister.codeEmail" type="text" title="请输入邮箱验证码" placeholder="请输入邮箱验证码" @input="codeEmailIn"><el-button slot="append">{{ emailCodeStr }}</el-button> </el-input>
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="password1">
              <el-input v-model="inputRegister.password1" type="password" title="请输入密码" placeholder="请输入密码" show-password />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="password2">
              <el-input v-model="inputRegister.password2" type="password" title="请输入确认密码" placeholder="请输入确认密码" show-password />
            </el-form-item>
          </div>
          <el-button style="width: 30%;" round type="primary" :disabled="inputRegister.buttonDisabled" @click="register">
            注册
          </el-button>

        </el-form>
      </div>
      <div
        class="small-box"
        :class="{ active: isLogin }"
      >
        <div
          v-if="isLogin"
          class="small-contain"
        >
          <div class="stitle">你好，朋友!</div>
          <p class="scontent">开始注册，和我们一起旅行</p>
          <button
            class="sbutton"
            @click="changeType"
          >注册</button>
        </div>
        <div
          v-else
          class="small-contain"
        >
          <div class="stitle">欢迎回来!</div>
          <p class="scontent">与我们保持联系，请登录你的账户</p>
          <button
            class="sbutton"
            @click="changeType"
          >登录</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import { getCodeId, confirmCode, emailSend, emailConfirm, registerUser } from '@/api/user'
import { Message } from 'element-ui'

export default {
  mounted(){
        const oIframe = document.getElementById('show-iframe');
        const deviceWidth = document.documentElement.clientWidth;
        const deviceHeight = document.documentElement.clientHeight;
        oIframe.style.width = deviceWidth + 'px';
        oIframe.style.height = deviceHeight + 'px';
    },
  name: 'LoginRegister',
  data() {
    const validateEmail = (rule, value, callback) => {
      console.log('validateEmail', value)
      if (value === '') {
        callback(new Error('请正确填写邮箱'))
      } else {
        if (value !== '') {
          var reg = /^\w+([-+.])*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/
          // var reg=/^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
          if (!reg.test(value)) {
            callback(new Error('请输入有效的邮箱'))
          }
        }
        callback()
      }
    }

    const validatepassword1 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.inputRegister.password2 !== '') {
          this.$refs.inputRegister.validateField('password2')
        }
        callback()
      }
    }
    const validatepassword2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.inputRegister.password1) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      url_all:"1",
      src1: 'https://tva2.sinaimg.cn/large/9bd9b167gy1g2qkt9k952j21hc0u01kx.jpg',
      inputRegisterRules: { // 校验

        'username': [// 用户名
          { required: true, message: '用户名不可为空', trigger: 'blur' }
        ],
        'email': [// 邮件

          { required: true, validator: validateEmail, trigger: 'blur' }

        ],

        'codeNum': [
          { required: true, message: '图片验证码不可为空', trigger: 'blur' }
        ],
        'codeEmail': [
          { required: true, message: '邮件验证码不可为空', trigger: 'blur' }
        ],
        'password1': [
          { required: true, validator: validatepassword1, trigger: 'blur' }

        ],
        'password2': [
          { required: true, validator: validatepassword2, trigger: 'blur' }
        ]

      },
      inputLoginRules: {
        'password': [
          { required: true, message: '密码不可为空', trigger: 'blur' }
        ],
        'username': [
          { required: true, message: '用户名不可为空', trigger: 'blur' }
        ]
      },
      inputLogin: { // 登录数据
        'password': '',
        'username': ''
      },

      inputRegister: { // 注册数据
        'password1': '',
        'password2': '',
        'username': '',
        'codeNum': '',
        'codeEmail': '',
        'email': '',
        'codeNumToken': '',
        'buttonDisabled': true,
        'registerLoading': false
      },
      emailCodeStr: '请输入图形验证码',
      captcha_id: '',
      reCodeLoading: false,
      isLogin: true,
      loginLoading: false,
      emailError: false,
      passwordError: false,
      existed: false,
      form: {
        username: '',
        useremail: '',
        userpwd: ''
      }
    }
  },

  methods: {
     goBack(){
            this.goBackState = false;
            this.iframeState = false;
        },
        showIframe(){
            this.goBackState = true;
            this.iframeState = true;
        },
    codeEmailIn(str) {
      const codeStr = str
      console.log('codeEmailIn', codeStr)
      if (codeStr.length === 6) {
        console.log('checkEmail', this.inputRegister.codeNumToken)
        var updata = {
          'captcha_token': this.inputRegister.codeNumToken,
          'email_code': codeStr
        }
        // emailConfirm
        this.inputRegister.registerLoading = true

        emailConfirm(updata).then(response => {
          this.emailCodeStr = '验证成功'
          // 允许注册
          this.inputRegister.registerLoading = false
          this.inputRegister.buttonDisabled = false
          Message({
            message: response,
            type: 'success',
            duration: 5 * 1000
          })
        }, reason => {
          console.error(reason) // 出错了！
        })
      }
    },
    codeNumIn(str) {
      const codeStr = str
      console.log('codeNumIn', codeStr)
      if (codeStr.length === 6) { // 启动验证
        const validateList = []

        this.$refs['inputRegister'].validateField(['email', 'username'], (message) => {
          validateList.push(message)
        })
        console.log('validateList:', validateList)
        if (validateList.every((item) => item === '')) {
          // 咱们的操作
          this.confirmCodeNum(codeStr)
        }
      }
    },
    confirmCodeNum(codeNum) {
      console.log('confirmCodeNum', codeNum)
      const that = this
      const updata = {
        'captcha_solution': codeNum,
        'captcha_id': this.captcha_id
      }

      confirmCode(updata).then(response => {
        // {captcha_token: 'f61daa354e778d0d5282892514f88bfd'}
        console.log('confirmCode:', response)
        // 发送邮箱验证码 emailSend
        that.inputRegister.codeNumToken = response.captcha_token
        var updata = {
          'captcha_token': response.captcha_token,
          'email_address': this.inputRegister.email
        }
        console.log('sendEmail:', updata)
        this.inputRegister.registerLoading = true
        // 发送邮件
        emailSend(updata).then(response => {
          console.log('sendEmail:', response)
          that.inputRegister.emailCodeStr = '请输入邮箱验证码'
          this.inputRegister.registerLoading = false
          Message({
            message: '验证码发送成功，请及时输入验证码',
            type: 'success',
            duration: 5 * 1000
          })
        })
      }, reason => {
        this.reCodefuc()
        console.error(reason) // 出错了！
      })
    },
    changeType() {
      console.log('changeType', this.isLogin)
      if (this.isLogin) {
        this.reCodefuc()
      }
      this.isLogin = !this.isLogin

      this.inputLogin = { // 登录数据
        'password': '',
        'username': ''
      }
      this.inputRegister = { // 注册数据
        'password1': '',
        'password2': '',
        'username': '',
        'codeNum': '',
        'codeEmail': '',
        'email': '',
        'codeNumToken': '',
        'buttonDisabled': true,
        'registerLoading': false
      }
    },

    reCodefuc() {
      console.log('reCodefuc:')
      if (this.reCodeLoading === true) {
        return 0
      }
      this.reCodeLoading = true
      getCodeId().then(response => {
        console.log('response:', response)
        this.captcha_id = response.captcha_id
        var protocolStr = document.location.protocol

        if (protocolStr === 'http:') {
          this.src1 = 'http://' + process.env.VUE_APP_BASE_HTTP_API + "/captcha/" + response.captcha_id + '.png'
          // console.log('protocol = ' + protocolStr)
        } else if (protocolStr === 'https:') {
          this.src1 = 'https://' + process.env.VUE_APP_BASE_HTTPS_API+ "/captcha/" + response.captcha_id + '.png'
          // console.log('protocol = ' + protocolStr)
        } else {
          this.src1 = 'http://' + process.env.VUE_APP_BASE_HTTP_API+ "/captcha/" + response.captcha_id + '.png'
          // console.log('other protocol')
        }
      console.log("this.src1 :",this.src1 )
        this.reCodeLoading = false
        // this.formData = response
      }, reason => {
        this.reCodeLoading = false
        console.error(reason) // 出错了！
      })
    },
    login() { // 登录
      this.$refs.inputLogin.validate(valid => {
        if (valid) {
          this.loading = true

          const data = this.inputLogin
          // data['type'] = 'phone'
          console.log('data', data)
          this.$store.dispatch('user/login', data).then(() => {
            console.log('data,', this.redirect)
            this.$router.push({ path: this.redirect || '/' })

            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    wechatHandleClick(thirdpart) {
  
                let baseUrl = "https://open.weixin.qq.com/connect/qrconnect"
                let appid = "wxbdc5610cc59c1631"
                let redictUrl = "https%3A%2F%2Fpassport.yhd.com%2Fwechat%2Fcallback.do"
                let state = "3d6be0a4035d839573b04816624a415e"

                let url = baseUrl + "?appid=" + appid + "&redirect_uri=" + redictUrl
                    + "&response_type=code" + "&scope=snsapi_login" + "&state=" + state

                window.location = url
            

      //alert('ok')
      // this.$store.commit('SET_AUTH_TYPE', thirdpart)
      // const appid = 'xxxxx'
      // const redirect_uri = encodeURIComponent('xxx/redirect?redirect=' + window.location.origin + '/auth-redirect')
      // const url = 'https://open.weixin.qq.com/connect/qrconnect?appid=' + appid + '&redirect_uri=' + redirect_uri + '&response_type=code&scope=snsapi_login#wechat_redirect'
      // openWindow(url, thirdpart, 540, 540)
    },
    // QQ 第三方登录
    tencentHandleClick(thirdpart) {
      let baseUrl = "https://graph.qq.com/oauth2.0/authorize"
      let appid = 101981088
      let redictUrl = "http://config-platform.top/qqLogin"
      let state = "3d6be0a4035d839573b04816624a415e"
      //let scope = "get_user_info,add_share,add_one_blog,list_album,upload_pic,add_album,list_photo,check_page_fans,get_info,add_t,del_t,add_pic_t,get_repost_list,get_other_info,get_fanslist,get_idollist,add_idol,del_idol,get_tenpay_addr"

      let url = baseUrl + "?response_type=code&client_id=" + appid
          + "&redirect_uri=" + redictUrl + "&state=" + state 
          

      window.location = url


      /*
      // 直接弹出授权页面，授权过后跳转到回调页面进行登录处理
      QC.Login.showPopup({
        //state : "3d6be0a4035d839573b04816624a415e",
        appId: '101490224',
        redirectURI: 'http://localhost:3000/proxy' // 登录成功后会自动跳往该地址
      })
      */
      // 法二
      // var _self = this;// 先将vue这个对象保存在_self对象中
      // _self.$store.commit('SET_AUTH_TYPE', thirdpart)
      // const client_id = 'xxx';
      // const redirect_uri = "xxx";
      // const url = 'https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=' + client_id + '&redirect_uri=' + redirect_uri;
      // 打开QQ授权登录界面，授权成功后会重定向
      // openWindow(url, thirdpart, 540, 540);
    },
    getGithubAuthorizeCode() {
                let baseUrl = "https://github.com/login/oauth/authorize"
                let clientId = "b0f4b22bfa884640f030"
                let redictUrl = "http://config-platform.top/githubLogin"
                //redictUrl ="http://localhost:9745/#/githubLogin"
                let state = "3d6be0a4035d839573b04816624a415e"
             
                let scope = "read:user"

                let url = baseUrl + "?client_id=" + clientId + "&redirect_uri=" + redictUrl
                    + "&scope=" + scope + "&state=" + state

                //window.location = url
                //this.url_all = url
                console.log("url_all",url)
                //windows.parent.location.href=url
                //var mypage = window.open(url,'mypage','address=0,resizable=0,toolbar=0,location=0,status=0,menubar=0,fullscreen=0');
                
               let a = window.open(url,"_top")
              console.log("a:",a)
            },
    register() { // 注册
      this.$refs.inputRegister.validate(valid => {
        console.log(valid)
        if (valid) {
          console.log(this.inputRegister)
          const updata = {
            'email_address': this.inputRegister.email,
            'password': this.inputRegister.password1,
            'username': this.inputRegister.username
          }
          registerUser(updata).then(response => {
            const data = {
              'password': updata.password,
              'username': updata.username
            }
            this.$store.dispatch('user/login', data).then(() => {
              console.log('data,', this.redirect)
              this.$router.push({ path: this.redirect || '/' })
            }).catch(() => {

            })

            // 注册成功
            Message({
              message: '注册成功',
              type: 'success',
              duration: 5 * 1000
            })
          }, reason => {
            console.error(reason) // 出错了！
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped="scoped">
.login-register {
  width: 100vw;
  height: 100vh;
  box-sizing: border-box;
}
.contain {
  width: 60%;
  height: 70%;
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  border-radius: 20px;
  box-shadow: 0 0 3px #f0f0f0, 0 0 6px #f0f0f0;
}
.big-box {
  width: 70%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 30%;
  transform: translateX(0%);
  transition: all 1s;
}
.big-contain {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.btitle {
  font-size: 1.5em;
  font-weight: bold;
  color: rgb(57, 167, 176);
}

.bform {
  width: 120%;
  padding: 2em 0;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}
.bform .errTips {
  display: block;
  width: 50%;
  text-align: left;
  color: red;
  font-size: 0.7em;
  margin-left: 1em;
}
.bform input {
  width: 50%;
  height: 30px;
  border: none;
  outline: none;
  border-radius: 10px;
  padding-left: 2em;
  background-color: #f0f0f0;
}
.bbutton {
  width: 20%;
  height: 40px;
  border-radius: 24px;
  border: none;
  outline: none;

  font-size: 0.9em;
  cursor: pointer;
}
.small-box {
  width: 30%;
  height: 100%;
  background: linear-gradient(135deg, rgb(57, 167, 176), rgb(56, 183, 145));
  position: absolute;
  top: 0;
  left: 0;
  transform: translateX(0%);
  transition: all 1s;
  border-top-left-radius: inherit;
  border-bottom-left-radius: inherit;
}
.small-contain {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.stitle {
  font-size: 1.5em;
  font-weight: bold;
  color: #fff;
}
.scontent {
  font-size: 0.8em;
  color: #fff;
  text-align: center;
  padding: 2em 4em;
  line-height: 1.7em;
}
.sbutton {
  width: 60%;
  height: 40px;
  border-radius: 24px;
  border: 1px solid #fff;
  outline: none;
  background-color: transparent;
  color: #fff;
  font-size: 0.9em;
  cursor: pointer;
}
.big-box.active {
  left: 0;
  transition: all 0.5s;
}
.small-box.active {
  left: 100%;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  border-top-right-radius: inherit;
  border-bottom-right-radius: inherit;
  transform: translateX(-100%);
  transition: all 1s;
}
</style>

<!-- 三方登录样式  -->
<style rel="stylesheet/scss" lang="scss" scoped>
  .social-signup-container {
    margin: 20px 0;

  .sign-btn {
    display: inline-block;
    cursor: pointer;
  }

  .icon {
    color: #fff;
    font-size: 24px;
    margin-top: 8px;
  }

  .wx-svg-container,
  .qq-svg-container {
    display: inline-block;
    width: 40px;
    height: 40px;
    line-height: 40px;
    text-align: center;
    padding-top: 1px;
    border-radius: 4px;
    margin-bottom: 20px;
    margin-right: 5px;
  }

  .wx-svg-container {
    background-color: #24da70;
  }

  .qq-svg-container {
    background-color: #6BA2D6;
    margin-left: 50px;
  }

  }
</style>
