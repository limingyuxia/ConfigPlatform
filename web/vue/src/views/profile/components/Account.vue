<template>
  <div>
    <el-card v-for="(item,key) in basicInfoList" :key="item.path" class="box-card" shadow="hover">

      <span slot="header" class="clearfix">{{ item.title }}</span>

      <el-descriptions v-if="item.editFlag" class="information-title" :column="1" @click.native="item.editFlag = !item.editFlag">
        <template v-for="(itemDe,keyDe) in item.dataList">

          <el-descriptions-item v-if="itemDe.type == 'text'" :key="keyDe" :label="itemDe.title">{{ itemDe.model }}</el-descriptions-item>
          <el-descriptions-item v-if="itemDe.type == 'img'" :key="keyDe" :label="itemDe.title">
            <el-avatar shape="square" :size="50" :src="imageUrl" />
          </el-descriptions-item>
          <el-descriptions-item v-if="itemDe.type == 'select'" :key="keyDe" v-model="itemDe.model" :label="itemDe.title">{{ itemDe.data[itemDe.model].label }}</el-descriptions-item>

        </template>

      </el-descriptions>

      <el-form v-else style="vertical-align:middle; " :inline="true" label-width="70px" class="demo-form-inline">
        <template v-for="(itemDe,key) in item.dataList">

          <el-form-item v-if="itemDe.type == 'text'" :key="key" style="display:flex" :label="itemDe.title">
            <el-input v-model.trim="itemDe.model" />
          </el-form-item>

          <el-form-item v-if="itemDe.type == 'select'" :key="key" style="display:flex" :label="itemDe.title">

            <el-select v-model="itemDe.model" placeholder="请选择">

              <el-option
                v-for="itemSe in itemDe.data"
                :key="itemSe.value"
                :label="itemSe.label"
                :value="itemSe.value"
              />
            </el-select>

          </el-form-item>

          <el-form-item v-if="itemDe.type == 'img'" :key="itemDe.index" :label="itemDe.title">
            <!--:headers="headerObj"-->
            <el-upload
              ref="my-upload"
              class="avatar-uploader"
              action="tmp"
              method="POST"
              :http-request="uploadImage"
              list-type="picture"
              :limit="2"
              :on-change="handleChange"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
              :auto-upload="false"
            >
              <img v-if="imageUrl" :src="imageUrl" class="avatar" @error="loadIMG($event)">

              <i v-else slot="trigger" class="el-icon-plus avatar-uploader-icon" />

            </el-upload>
            <el-button style="margin-left:20px" :disabled="upAvatarButton" size="small" type="primary" @click="upAvatarfuc">点击上传</el-button>

          </el-form-item>

        </template>

        <div class="buttonUp">
          <el-button round @click="cancelFun(key)">取消</el-button>
          <el-button type="danger" round @click="submitFun(item.dataList,key)">保存</el-button>
        </div>
      </el-form>

    </el-card>

  </div>

</template>

<script>

import { uploadAvatar } from '@/api/user'
import { mapGetters } from 'vuex'

export default {
  computed: {

    ...mapGetters([
      'basicInfo'
    ])
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
      basicInfoList: {
        basicInfo: {
          editFlag: true,
          title: '基本信息',
          dataList: {
            nickname: { index: 0, type: 'text', title: '用户昵称', model: '123' },
            // username: { index: 1, type: 'text', title: '用户名', model: '123' },
            arv: { index: 2, type: 'img', title: '头像', model: '123' },
            region: { index: 3, type: 'text', title: '地区', model: '123' },
            gender: { index: 4, type: 'select', title: '性别', model: '0', data: [{ value: 0, label: '男' }, { value: 1, label: '女' }, { value: 2, label: '未知' }] }
          }
        }

      },
      imageUrl: '',
      my_loadingFalg: false,
      upAvatarButton: true
    }
  },
  mounted() {
    this.iniUserInfo()
    // this.upAvatarButton = true
  },

  methods: {
    iniUserInfo() {
      console.log('created_1', this.user)
      console.log('created_2', this.basicInfo)

      this.imageUrl = this.user.avatar

      // 循环赋值
      for (var key in this.basicInfoList.basicInfo.dataList) {
        if (this.basicInfoList.basicInfo.dataList[key].type === 'select') {
          this.basicInfoList.basicInfo.dataList[key].model = this.basicInfo[key]
        } else {
          this.basicInfoList.basicInfo.dataList[key].model = this.basicInfo[key]
        }
      }
      console.log('data_1', this.basicInfoList)

      if (this.$refs['my-upload']) {
        this.$refs['my-upload'][0].clearFiles()
      }
    },
    submitFun(data, key) {
      console.log('submitFun:', data)
      const upData = {
        'chType': 'basicInfo',
        'data': data
      }
      this.$store.dispatch('user/updata', upData).then(() => { // 上传图片
        this.$message({
          message: '用户信息更新成功',
          type: 'success',
          duration: 5 * 1000
        })
        this.basicInfoList[key].editFlag = true
        this.$emit('changeUser', { 'loadingAll': false })
      }).catch((e) => {
        this.$message({
          message: e,
          type: 'error',
          duration: 5 * 1000
        })
        this.$emit('changeUser', { 'loadingAll': false })
      })
    },
    cancelFun(data) {
      console.log('cancelFun:', data)
      this.iniUserInfo()
      this.basicInfoList[data].editFlag = true
    },
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
</style>
