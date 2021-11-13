<template>
  <div>
    <el-card v-for="(item) in basicInfoList" :key="item.path" class="box-card" shadow="hover">

      <span slot="header" class="clearfix">{{ item.title }}</span>

      <div v-for="(itemDe) in item.dataList" :key="itemDe">

        <el-row style="margin: auto;" type="flex" class="row-bg">
          <el-col style="margin: auto;" :span="4"><p>{{ itemDe.title }}：</p></el-col>
          <el-col style="margin: auto;text-align: right;word-break:break-all" :span="24 - 4 - 2">
            <p>{{ itemDe.describe }}</p>

          </el-col>
          <el-col style="margin: auto;text-align: right;" :span="2">

            <el-button type="text" size="small">{{ itemDe.button }}</el-button>
          </el-col>
        </el-row>

      </div>

 

    </el-card>

  </div>

</template>

<script>

// import { uploadAvatar } from '@/api/user'

export default {
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
      basicInfoList: [
        {
          editFlag: true,
          title: '账号设置',
          dataList: {
            nickname: { index: 0, describe: '', button: '设置密码', type: 'text', title: '用户昵称', model: '123' }

          }
        }

      ],
      imageUrl: '',
      my_loadingFalg: false,
      upAvatarButton: true
    }
  },
  mounted() {
    console.log('created_1', this.user)
    this.imageUrl = this.user.avatar
    this.$refs['my-upload'][0].clearFiles()
    // this.upAvatarButton = true
  },

  methods: {
    cardClickFuc(index) {
      console.log('点击卡片', index)
    },
    loadIMG(e) { // 图片加载出错
      console.log(e)
      this.$message.error('图片加载错误')
      this.imageUrl = ''
      this.$refs['my-upload'][0].clearFiles()
      this.upAvatarButton = true
    },

    upAvatarfuc() {
      console.log('upAvatarfuc')
      this.$refs['my-upload'][0].submit()
    },
    /* handleExceed(param,fileList){
console.log("param:",param)
console.log("param0:",param[0])
console.log("Filelist:",fileList)
      this.upAvatarButton=false
      this.imageUrl = URL.createObjectURL(param[0]);
      fileList.splice(0, 1);
    },*/
    handleChange(param, fileList) {
      if (fileList.length > 1) {
        fileList.splice(0, 1)
      }

      console.log('handleChange:', param.raw)
      const file = param.raw
      const isJPG = /^image\/(jpeg|png|jpg|x-icon)$/.test(file.type)

      // const isJPG = file.type === 'image/jpeg'; //x-icon
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
        this.$refs['my-upload'][0].clearFiles()
        return
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
        this.$refs['my-upload'][0].clearFiles()
        return
      }
      this.upAvatarButton = false
      this.imageUrl = URL.createObjectURL(file)
    },
    uploadImage(param) {
      const that = this
      this.$emit('changeUser', { 'loadingAll': true })
      console.log('uploadImage:', param)
      const formData = new FormData()
      console.log('param', param.file)
      formData.append('file', param.file)
      console.log('formData', formData)

      this.$store.dispatch('user/upAvatar', formData, param).then(() => { // 上传图片
        console.log('11')
        that.$refs['my-upload'][0].clearFiles()

        this.$message({
          message: '头像更新成功',
          type: 'success',
          duration: 5 * 1000
        })

        this.$emit('changeUser', { 'loadingAll': false })

        that.upAvatarButton = true
      }).catch(() => {
        that.$refs['my-upload'][0].clearFiles()
        this.$emit('changeUser', { 'loadingAll': false })
        that.upAvatarButton = true
      })
      /*
      uploadAvatar(formData).then(response => {
        console.log('上传图片成功')
        this.form.picUrl = process.env.VUE_APP_BASE_API + response.imgUrl

      }).catch(response => {
        console.log('图片上传失败')
      })
*/
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw)
    },
    beforeAvatarUpload(file) {
      // const isJPG = file.type === 'image/jpeg';
      const isJPG = /^image\/(jpeg|png|jpg|x-icon)$/.test(file.type)
      const isLt2M = file.size / 1024 / 1024 < 2
      console.log('123', isJPG)
      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
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
