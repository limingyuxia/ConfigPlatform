<template>
  <div class="login-register">
    <div class="contain">
      <div
        class="big-box"
        :class="{ active: isLogin }"
      >
        <el-form
          v-if="isLogin"
          class="big-contain"
        >

          <div class="btitle">账户登录</div>

          <div style="height: 40%;" class="bform">

            <el-input v-model="inputLogin.username" type="text" title="请输入用户名" style=" width: 50%;" placeholder="请输入用户名" />
            <el-input v-model="inputLogin.password" type="password" title="请输入密码" style=" width: 50%;" placeholder="请输入密码" show-password />

          </div>

          <button
            class="bbutton"
            @click="login"
          >登录</button>
          <span id="qqLoginBtn"></span>
          
          <div class="social-signup-container">

    <div class="sign-btn" @click="wechatHandleClick('wechat')">
      <span class="wx-svg-container"><svg-icon icon-class="wechat" class="icon"/></span> 微信
    </div>
    <div class="sign-btn" @click="tencentHandleClick('tencent')">
      <span class="qq-svg-container"><svg-icon icon-class="qq" class="icon" id="qqLoginBtn"/></span> QQ
    </div>
  </div>

        </el-form>
        <el-form
          v-else
          class="big-contain"
          ref="inputRegister"
          :model="inputRegister"
          :rules="inputRegisterRules"
        >
          <div class="btitle">创建账户</div>
          <div style="height: 70%;" class="bform">

            <el-form-item style="width: 50%;" prop="username">
              <el-input v-model="inputRegister.username" type="text" title="请输入用户名" placeholder="请输入用户名" />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="email">
              <el-input  v-model="inputRegister.email" type="email" title="请输入邮箱"  placeholder="请输入邮箱" />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="codeNum">
              <el-input v-model="inputRegister.codeNum" type="text"  @input="codeNumIn" title="请输入图形验证码" placeholder="请输入图形验证码"><el-button slot="append" v-loading="reCodeLoading" @click="reCodefuc()"><img onerror="default" style="width: 130px;height: 40px;" :src="src1"></el-button></el-input>
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="codeEmail">
              <el-input v-model="inputRegister.codeEmail" type="text" title="请输入邮箱验证码" placeholder="请输入邮箱验证码"  ><el-button slot="append">{{emailCodeStr}}</el-button> </el-input>
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="password1">
              <el-input v-model="inputRegister.password1" type="password" title="请输入密码"  placeholder="请输入密码" show-password />
            </el-form-item>

            <el-form-item style=" width: 50%;" prop="password2">
              <el-input v-model="inputRegister.password2" type="password" title="请输入确认密码"  placeholder="请输入确认密码" show-password />
            </el-form-item>
          </div>
          <button
            class="bbutton"
            @click="register"
          >注册</button>

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

import { getCodeId,confirmCode } from '@/api/user'
import { MessageBox } from 'element-ui'

export default {

  name: 'LoginRegister',
  data() {
    return {
      src1: 'https://tva2.sinaimg.cn/large/9bd9b167gy1g2qkt9k952j21hc0u01kx.jpg',
      inputRegisterRules:{//校验
        
        "username":[//用户名
          { required: true, message: '用户名不可为空', trigger: 'blur' },
        ],
        "email": [//邮件
          { required: true, message: '邮件不可为空', trigger: 'blur' },
        ],
        
        'codeNum': [
          { required: true, message: '图片验证码不可为空', trigger: 'blur' },
        ],
        "codeEmail":[
          { required: true, message: '邮件验证码不可为空', trigger: 'blur' },
        ],
        'password1': [
          { required: true, message: '密码不可为空', trigger: 'blur' },
        ],
        'password2': [
          { required: true, message: '密码不可为空', trigger: 'blur' },
        ],


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
        "codeEmail":'',
        'email': ''
      },
      emailCodeStr:"点击获取邮箱验证码",
      captcha_id:"",
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
    codeNumIn(str){
      const codeStr = str
      if(codeStr.length === 6){//启动验证
        this.confirmCodeNum(codeStr)

      }
    },
    confirmCodeNum(codeNum){
      const that = this
      const updata = {
        "captcha_solution":codeNum,
        "captcha_id":this.captcha_id
      }
      confirmCode(updata).then(response => {
        console.log('confirmCode:', response)
        if (response === "验证成功") {
          that.emailCodeStr ="已经发送验证码"
        }else{
            Message({
            message: response,
            type: 'error',
            duration: 5 * 1000
          })
        }
        
        
      }, reason => {
        this.reCodefuc()
        console.error(reason) // 出错了！
      })
    },
    changeType() {
      console.log("changeType",this.isLogin)
      if (this.isLogin) {
        this.reCodefuc()
      }
      this.isLogin = !this.isLogin

      this.inputLogin= { // 登录数据
        'password': '',
        'username': ''
      }
      this.inputRegister= { // 注册数据
        'password1': '',
        'password2': '',
        'username': '',
        'codeNum': '',
        "codeEmail":'',
        'email': ''
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
          this.src1 = 'http://' + process.env.VUE_APP_BASE_HTTP_API + response.captcha_id + ".png"
          //console.log('protocol = ' + protocolStr)
        } else if (protocolStr === 'https:') {
          this.src1 = 'https://' + process.env.VUE_APP_BASE_HTTPS_API + response.captcha_id + ".png"
          //console.log('protocol = ' + protocolStr)
        } else {
          this.src1 = 'http://' + process.env.VUE_APP_BASE_HTTP_API + response.captcha_id + ".png"
          //console.log('other protocol')
        }

        
        this.reCodeLoading = false
        // this.formData = response
      }, reason => {
        this.reCodeLoading = false
        console.error(reason) // 出错了！
      })
    },
    login() {
      var valid = true
      console.log(valid)

      if (valid) {
        this.loading = true
        const data = this.inputLogin
        // data['type'] = 'phone'

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
    },
    wechatHandleClick(thirdpart) {
                alert('ok')
                // this.$store.commit('SET_AUTH_TYPE', thirdpart)
                // const appid = 'xxxxx'
                // const redirect_uri = encodeURIComponent('xxx/redirect?redirect=' + window.location.origin + '/auth-redirect')
                // const url = 'https://open.weixin.qq.com/connect/qrconnect?appid=' + appid + '&redirect_uri=' + redirect_uri + '&response_type=code&scope=snsapi_login#wechat_redirect'
                // openWindow(url, thirdpart, 540, 540)
            },
            // QQ 第三方登录
            tencentHandleClick(thirdpart) {
                // 直接弹出授权页面，授权过后跳转到回调页面进行登录处理
                QC.Login.showPopup({
                    appId:"123132",
                    redirectURI:"xxx"  //登录成功后会自动跳往该地址
                });

                // 法二
                // var _self = this;// 先将vue这个对象保存在_self对象中
                // _self.$store.commit('SET_AUTH_TYPE', thirdpart)
                // const client_id = 'xxx';
                // const redirect_uri = "xxx";
                // const url = 'https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=' + client_id + '&redirect_uri=' + redirect_uri;
                // 打开QQ授权登录界面，授权成功后会重定向
                // openWindow(url, thirdpart, 540, 540);
            },
    register() {//注册
      this.$refs.inputRegister.validate(valid => {
        console.log(valid)
        if (valid) {
          this.loading = true
          const data = this.loginForm
          // data['type'] = 'phone'

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
  background-color: rgb(57, 167, 176);
  color: #fff;
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
