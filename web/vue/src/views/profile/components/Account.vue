<template >
  <el-form  >
     
    <el-form-item label="Name">

      <el-input v-model.trim="user.name" />
    </el-form-item>
    <el-form-item label="Email">
      <el-input v-model.trim="user.email" />
    </el-form-item>
    <el-form-item  label="头像">
      <!--:headers="headerObj"-->
      <el-upload
      ref="my-upload"
        class="avatar-uploader"
        action="tmp"
        enctype="multipart/form-data" 
        method="POST"
        
        :http-request="uploadImage"
        list-type="picture"

        :on-change="handleChange"
        :show-file-list="true"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload"
        :auto-upload="false">
        <img v-if="imageUrl" :src="imageUrl" class="avatar" @error="loadIMG($event)"  >

         
        <i v-else slot="trigger" class="el-icon-plus avatar-uploader-icon"></i>

      </el-upload>
      <el-button :disabled="upAvatarButton" @click="upAvatarfuc" size="small" type="primary">点击上传</el-button>

    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">Update</el-button>
    </el-form-item>
  </el-form>
</template>



<script>

import { uploadAvatar} from '@/api/user'

export default {
  data() {
      return {
        imageUrl: "",
        my_loadingFalg:false,
        upAvatarButton:true
      };
    },
  props: {
    user: {
      type: Object,

      default: () => {
        return {
          name: '',
          email: '',
          avatar:''
        }
      }
    }
  },  
  created() {
    console.log("created_1",this.user)
    this.imageUrl = this.user.avatar
    that.$refs['my-upload'].clearFiles();
    //this.upAvatarButton = true
  },


  methods: {
    loadIMG(e) {  // 图片加载出错
        console.log(e)
        this.$message.error('图片加载错误');
        this.imageUrl = ""
        this.$refs['my-upload'].clearFiles();
        this.upAvatarButton=true
    },
      

    upAvatarfuc(){
      console.log("upAvatarfuc")
      this.$refs['my-upload'].submit();

    },
    handleChange(param){
      console.log("handleChange:",param.raw)
      let file = param.raw
      const isJPG = /^image\/(jpeg|png|jpg|x-icon)$/.test(file.type)

      //const isJPG = file.type === 'image/jpeg'; //x-icon
      const isLt2M = file.size / 1024 / 1024 < 2;
      
      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
        this.$refs['my-upload'].clearFiles();
        return
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
        this.$refs['my-upload'].clearFiles();
        return
      }
      this.upAvatarButton=false
      this.imageUrl = URL.createObjectURL(file);

    },
    uploadImage(param){
      const that = this
      this.$emit("changeUser", {"loadingAll":true});
    console.log("uploadImage:",param)
      const formData = new FormData()
      console.log("param",param)
      formData.append('file', param.file)

      this.$store.dispatch('user/upAvatar', formData,param).then(() => {//上传图片

           console.log("11")
            that.$refs['my-upload'].clearFiles();
            
             this.$message({
                message: '头像更新成功',
                type: 'success',
                duration: 5 * 1000
              })

            this.$emit("changeUser", {"loadingAll":false});

           that.upAvatarButton=true
           
          }).catch(() => {
            that.$refs['my-upload'].clearFiles();
            this.$emit("changeUser", {"loadingAll":false});
            that.upAvatarButton=true
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
        this.imageUrl = URL.createObjectURL(file.raw);
      },
      beforeAvatarUpload(file) {
        //const isJPG = file.type === 'image/jpeg';
        const isJPG = /^image\/(jpeg|png|jpg|x-icon)$/.test(file.type)
        const isLt2M = file.size / 1024 / 1024 < 2;
        console.log("123",isJPG)
        if (!isJPG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 2MB!');
        }
        return isJPG && isLt2M;
      },
    submit() {
      this.name = "123456"
      this.user.name = this.user.name + 1

      this.$emit("changeUser", {"loadingAll":false});

      //this.$emit("loadingFalg", !this.loadingFalg);
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
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>