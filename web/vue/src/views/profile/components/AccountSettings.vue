<template>
  <div>
    <el-card v-for="(item) in basicInfoList" :key="item.path" class="box-card" shadow="hover">

      <span slot="header" class="clearfix">{{ item.title }}</span>

      <div v-for="(itemDe) in item.dataList" :key="itemDe.index">

        <el-row style="margin: auto;" type="flex" class="row-bg">

          <el-avatar v-if="itemDe.avatarUrl" style="margin: auto;" :size="30" :src="itemDe.avatarUrl" />
          <el-col style="margin: auto; margin-left:5px;" :span="4"><p>{{ itemDe.title }}：</p></el-col>
          <el-col style="margin: auto;word-break:break-all" :span="24 - 4 - 2">
            <el-avatar v-if="false" style="margin: auto;" :size="30" :src="itemDe.avatarUrl" />
            <div style="margin: auto;">{{ itemDe.describe }}</div>

          </el-col>
          <el-col style="margin: auto" :span="2">

            <el-button type="text" size="small" @click="showDialog(itemDe)">{{ itemDe.button }}</el-button>
          </el-col>
        </el-row>

      </div>

    </el-card>
    <el-dialog

      :visible.sync="dialogVisible"
      width="50%"
      :before-close="handleClose"
      :title="dialogData.title"
      center
    >
      <el-steps align-center :active="dialogData.active" finish-status="success">
        <el-step v-for="(itemSteps,key) in dialogData.steps" :key="key" :title="itemSteps.title" />

      </el-steps>

      <el-result v-if="dialogData.active == 2" icon="success" title="修改成功" sub-title="请点击下方按钮">
        <template slot="extra">
          <el-button dialog-visible="false" type="primary" size="medium" @click="dialogVisible = false">返回</el-button>
        </template>
      </el-result>

      <el-form
        v-loading="dialogData.loading"
        :model="dialogData.formData"
        label-position="right"
        label-width="20%"
      >

        <template v-for="(itemDe,key) in dialogData.formInline[dialogData.active]">
          <el-form-item v-if="itemDe.type == 'sendCode'" :key="itemDe.key" style="text-align:'center';width: 80%;" label="">
            <el-input v-model="itemDe.model" :placeholder="itemDe.placeholder"> <el-button slot="suffix" :disabled="itemDe.disabled" type="text" @click="getCode(key)">{{ itemDe.label }}</el-button></el-input>
          </el-form-item>

          <el-form-item v-if="itemDe.type == 'text'" :key="itemDe.key" style="text-align:'center';margin:auto;width: 60%;" :label="itemDe.label">
            <span style="margin:auto" class="demonstration">{{ itemDe.placeholder }}</span>
          </el-form-item>

          <el-form-item v-if="itemDe.type == 'code'" :key="itemDe.key" style=" width: 80%;" prop="codeNum">
            <el-input v-model="itemDe.model" type="text" title="请输入图形验证码" placeholder="请输入图形验证码"><el-button slot="append" v-loading="itemDe.Loading" @click="reCodefuc(key)"><img onerror="default" style="width: 130px;height: 40px;" :src="itemDe.value"></el-button></el-input>
          </el-form-item>

          <el-form-item v-if="itemDe.type == 'input'" :key="itemDe.key" style=" width: 80%;" prop="codeNum">
            <el-input v-model="itemDe.model" type="text" :title="itemDe.title" :placeholder="itemDe.placeholder" />
          </el-form-item>

        </template>

      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button v-if="dialogData.steps[dialogData.active].nextButtonLabel" :disabled="dialogData.steps[dialogData.active].disabled" type="danger" round @click="nextStep(dialogData.active)">{{ dialogData.steps[dialogData.active].nextButtonLabel }}</el-button>
      </span>

    </el-dialog>
  </div>

</template>

<script>

import { getCodeId } from '@/api/user'
import { mapGetters } from 'vuex'
import { Message } from 'element-ui'
export default {
  computed: {

    ...mapGetters([
      'authInfo',
      'email',
      'phone'
    ])
  },
  watch: {
    'email': function(newVal) {
      console.log('email_数据:', newVal)
      this.basicInfoList[0].dataList.email.describe = newVal
    }

  },
  props: {
    user: {
      type: Object,

      default: () => {
        return {
          name: '',
          email: '',
          avatar: ''
        }
      }
    }
  },
  data() {
    return {
      timer: '',

      dialogVisible: false,
      dialogData: {
        active: 1,
        title: '修改邮箱',
        formData: {},
        steps: [
          { 'nextButtonLabel': '点击验证' },
          { 'nextButtonLabel': '点击绑定' },
          { 'nextButtonLabel': '' }

        ],
        // isLogin
        loading: false,
        formInline: [
          [
            { 'type': 'text', 'value': '', 'model': '1160201372@qq.com', 'placeholder': '验证码将发送到邮箱\n(1160201372@qq.com)' },
            { 'type': 'code', 'label': '', 'value': '', 'captcha_id': '', 'model': '', 'placeholder': '请输入' },
            { 'type': 'sendCode', 'disabled': false, 'captcha_token': '', 'codeType': 'email', 'label': '获取验证码', 'value': '', 'model': '', 'placeholder': '请输入', 'account': '1160201372@qq.com' }
          ],
          [
            { 'type': 'input', 'value': '', 'model': '', 'placeholder': '请输入新邮箱' },
            { 'type': 'code', 'label': '', 'value': '', 'captcha_id': '', 'model': '', 'placeholder': '请输入' },
            { 'type': 'sendCode', 'disabled': false, 'codeType': 'email', 'label': '获取验证码', 'value': '', 'model': '', 'placeholder': '请输入', 'account': '1160201372@qq.com' }
          ]
        ]
      },
      basicInfoList: [
        {
          editFlag: true,
          title: '账号设置',
          dataList: {
            nickname: { index: 0, describe: '存在风险，请设置密码 ', button: '设置密码', type: 'text', title: '密码', model: '123' },
            phone: { index: 1, describe: ' ', button: '修改手机', type: 'text', title: '手机', model: '123' },
            email: { index: 2, describe: ' ', button: '修改邮箱', type: 'text', title: '邮箱', model: '123', 'showType': 'email' }
          }
        },
        {
          editFlag: true,
          title: '帐号关联',
          dataList: {
            wechat: { index: 0, avatarUrl: 'https://open.weixin.qq.com/zh_CN/htmledition/res/assets/res-design-download/icon48_appwx_logo.png', desPicture: '', describe: '', button: '绑定', type: 'text', title: '微信', model: '123' },
            qq: { index: 1, avatarUrl: 'https://img.ixintu.com/upload/jpg/20210522/abb7b6902f2f6fe5a399809e97aeef6c_61793_800_800.jpg!con', desPicture: '', describe: '', button: '绑定', type: 'text', title: 'QQ', model: '123' },
            github: { index: 2, avatarUrl: 'http://bpic.588ku.com/element_pic/00/25/58/9656d059a2d139b.jpg', desPicture: '', describe: '', button: '绑定', type: 'text', title: 'GitHub', model: '123' },
            weibo: { index: 3, avatarUrl: 'https://img.ixintu.com/download/jpg/201912/f2b9e19eae370b85e0cf0c23fc18ec30.jpg!ys', desPicture: '', describe: '', button: '绑定', type: 'text', title: '新浪微博', model: '123' }
          }
        }
      ],
      imageUrl: '',
      my_loadingFalg: false,
      upAvatarButton: true
    }
  },
  mounted() {
    console.log('created_17687', this.user)
    this.iniUserInfo()
  },

  methods: {
    nextStep(key) {
      console.log('下一步', key)
      var formInline = this.dialogData.formInline[this.dialogData.active]
      var active = this.dialogData.active
      // this.dialogData.active = this.dialogData.active + 1

      // 验证数据
      var storeData = {
        codeType: formInline[2].codeType,
        captcha_token: formInline[2].captcha_token,
        email_code: formInline[2].model,
        account: formInline[0].model
      }
      // confirmCode
      if (active === 0) { // 验证
        this.dialogData.loading = true
        this.$store.dispatch('verify/confirmCode', storeData).then((redirect) => {
          console.log('confirmCode')
          this.dialogData.loading = false
          this.dialogData.active = this.dialogData.active + 1
          this.reCodefuc(1)
        }).catch((e) => {
          console.log('error,', e)
        })
      } else if (active === 1) { // 更新用户数据
        console.log('第二步骤', storeData)
        this.dialogData.loading = true
        this.$store.dispatch('verify/confirmCode', storeData).then(async(redirect) => {
          console.log('confirmCode')
          // 提交数据
          const upData = {
            'chType': 'email',
            'data': storeData
          }
          await this.$store.dispatch('user/updata', upData)
          this.dialogData.loading = false
          this.dialogData.active = this.dialogData.active + 1
        }).catch((e) => {
          console.log('error,', e)
        })
      }
    },
    getCode(key) {
      this.dialogData.loading = true
      var formInline = this.dialogData.formInline[this.dialogData.active][key]
      var formInline_last = this.dialogData.formInline[this.dialogData.active][key - 1]
      var formInline_farst = this.dialogData.formInline[this.dialogData.active][key - 2]

      const storeData = {
        codeType: formInline.codeType,
        account: formInline_farst.model,
        captcha_id: formInline_last.captcha_id,
        captcha_num: formInline_last.model

      }
      this.$store.dispatch('verify/sendCode', storeData).then((redirect) => {
        console.log('sendCode', redirect)
        console.log('formInline', formInline)
        formInline.captcha_token = redirect.captcha_token
        this.dialogData.loading = false
        Message({
          message: '验证码发送成功，请及时输入验证码',
          type: 'success',
          duration: 5 * 1000
        })
        this.dialogData.steps[this.dialogData.active].disabled = false
        // 倒计时
        formInline.disabled = true
        let i = 60 // 倒计时时间
        formInline.label = i + 's'
        this.timer = setInterval(() => {
          formInline.label = i + 's'
          i--
          if (i < 0) {
            formInline.disabled = false
            formInline.label = '获取验证码'
            clearInterval(this.timer)
          }
        }, 1000)
      }).catch((e) => {
        console.log('error,', e)
        this.reCodefuc(1)
        this.dialogData.loading = false
      })
    },
    reCodefuc(key) {
      var formInline = this.dialogData.formInline[this.dialogData.active][key]
      var formInline_last = this.dialogData.formInline[this.dialogData.active][key - 1]

      this.$store.dispatch('verify/getPictureCode').then((redirect) => {
        console.log('data:', redirect)

        formInline.captcha_id = redirect.captcha_id
        formInline.value = redirect.captcha_url
      }).catch((e) => {
        console.log('error,', e)
      })
    },

    async iniUserInfo() {
      await this.$store.dispatch('user/getAuthInfo')

      console.log('created_信息')
      console.log('created_原始', this.basicInfoList[0].dataList)

      this.basicInfoList[0].dataList.email.describe = this.email
      this.basicInfoList[0].dataList.phone.describe = this.phone
      for (var key in this.authInfo) { // 第三方信息
        const myKey = key.split('_')[0]
        const myType = key.split('_')[1]
        if (myType === 'username') {
          this.basicInfoList[1].dataList[myKey].describe = this.authInfo[key]
        }
      }
    },
    cardClickFuc(index) {
      console.log('点击卡片', index)
    },
    showDialog(itemDe) {
      console.log('显示类型', itemDe)
      if (itemDe.showType === 'email') {
        // 弹窗赋值
        this.dialogData = {
          active: 0,
          title: '修改邮箱',
          steps: [
            { 'title': '验证原有邮箱', 'disabled': true, 'nextButtonLabel': '点击验证' },
            { 'title': '修改邮箱', 'disabled': true, 'nextButtonLabel': '点击绑定' },
            { 'title': '完成', 'disabled': true, 'nextButtonLabel': '' }

          ],
          // isLogin
          loading: false,
          formInline: [
            [
              { 'type': 'text', 'value': '', 'model': this.email, 'placeholder': '验证码将发送到邮箱\n(' + this.email + ')' },
              { 'type': 'code', 'label': '', 'value': '', 'captcha_id': '', 'model': '', 'placeholder': '请输入' },
              { 'type': 'sendCode', 'disabled': false, 'captcha_token': '', 'codeType': 'email', 'label': '获取验证码', 'value': '', 'model': '', 'placeholder': '请输入', 'account': '1160201372@qq.com' }
            ],
            [
              { 'type': 'input', 'value': '', 'model': '', 'placeholder': '请输入新邮箱' },
              { 'type': 'code', 'label': '', 'value': '', 'captcha_id': '', 'model': '', 'placeholder': '请输入' },
              { 'type': 'sendCode', 'disabled': false, 'codeType': 'email', 'label': '获取验证码', 'value': '', 'model': '', 'placeholder': '请输入', 'account': '1160201372@qq.com' }
            ]
          ]

        }
        if (this.email.length == 0) {
          this.dialogData.active = 1
        }

        this.reCodefuc(1)
        this.dialogVisible = true
      }
      // this.dialogVisible = true
      // this.reCodefuc(1)
    },
    handleClose(done) {
      this.$confirm('确认关闭？')
        .then(_ => {
          clearInterval(this.timer)
          done()
        })
        .catch(_ => {})
    },

    submit() {
      this.name = '123456'
      this.user.name = this.user.name + 1

      this.$emit('changeUser', { 'loadingAll': false })

      // this.$emit("loadingFalg", !this.loadingFalg);
      this.$message({
        message: 'User information has been updated successfully',
        type: 'success',
        duration: 5 * 1000
      })
    }
  }
}
</script>
<style>
  .avatar-uploader .el-upload {
    display:flex;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
     border: 1px dashed #d9d9d9;
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    border: 1px dashed #d9d9d9;
    width: 178px;
    height: 178px;

  }
   .box-card {
    width: 98%;
    margin-top:15px;
    margin-left:1%;
  }
  .information-title {
  color: #19d3ea;
  font-size: 18px;
  cursor: pointer;  /*鼠标悬停变小手*/
  width: 100%;
}
.buttonUp{

position:absolute;
left: 50%;
transform: translate(-50%,-50%);

}
 .row-bg {
   margin:auto;

    border-bottom: 0.5px solid rgb(240, 240, 245);

  }
</style>
