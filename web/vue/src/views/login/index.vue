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
        </el-form>
        <el-form
          v-else
          class="big-contain"
        >
          <div class="btitle">创建账户</div>
          <div style="height: 70%;" class="bform">

            <el-input v-model="inputRegister.username" type="text" title="请输入用户名" style=" width: 50%;" placeholder="请输入用户名" />
            <el-input v-model="inputRegister.email" type="email" title="请输入邮箱" style=" width: 50%;" placeholder="请输入邮箱" />

            <el-input v-model="inputRegister.codeNum" type="text" title="请输入图形验证码" style=" width: 50%;" placeholder="请输入图形验证码"><el-button slot="append" v-loading="reCodeLoading" @click="reCodefuc()"><img onerror="default" style="width: 125px;height: 30px;" :src="src1"></el-button></el-input>
            
            <el-input v-model="inputRegister.codeEmail" type="text" title="请输入邮箱验证码" style=" width: 50%;" placeholder="请输入邮箱验证码"  ><el-button slot="append">点击获取邮箱验证码</el-button> </el-input>

            <el-input v-model="inputRegister.password1" type="password" title="请输入密码" style=" width: 50%;" placeholder="请输入密码" show-password />
            <el-input v-model="inputRegister.password2" type="password" title="请输入确认密码" style=" width: 50%;" placeholder="请输入确认密码" show-password />

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

import { reCode } from '@/api/user'

export default {

  name: 'LoginRegister',
  data() {
    return {
      src1: 'https://tva2.sinaimg.cn/large/9bd9b167gy1g2qkt9k952j21hc0u01kx.jpg',
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
    changeType() {
      this.isLogin = !this.isLogin
      this.form.username = ''
      this.form.useremail = ''
      this.form.userpwd = ''
    },
    reCodefuc() {
      console.log('reCodefuc:')
      if (this.reCodeLoading === true) {
        return 0
      }
      this.reCodeLoading = true
      reCode().then(response => {
        console.log('response:', response)
        this.src1 = response.imgurl
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
    register() {
      const self = this
      if (
        self.form.username !== '' &&
        self.form.useremail !== '' &&
        self.form.userpwd !== ''
      ) {
        self
          .$axios({
            method: 'post',
            url: 'http://127.0.0.1:10520/api/user/add',
            data: {
              username: self.form.username,
              email: self.form.useremail,
              password: self.form.userpwd
            }
          })
          .then((res) => {
            switch (res.data) {
              case 0:
                alert('注册成功！')
                this.login()
                break
              case -1:
                this.existed = true
                break
            }
          })
          .catch((err) => {
            console.log(err)
          })
      } else {
        alert('填写不能为空！')
      }
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
