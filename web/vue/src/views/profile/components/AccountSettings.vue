<template>
  <div>
    <el-card v-for="(item) in basicInfoList" :key="item.path" class="box-card" shadow="hover">

      <span slot="header" class="clearfix">{{ item.title }}</span>


      <div v-for="(itemDe) in item.dataList" :key="itemDe.index">
        
        <el-row style="margin: auto;" type="flex" class="row-bg">
         
          <el-avatar v-if="itemDe.avatarUrl" style="margin: auto;" :size="30" :src="itemDe.avatarUrl"></el-avatar>
          <el-col style="margin: auto; margin-left:5px;" :span="4"><p>{{ itemDe.title }}：</p></el-col>
          <el-col style="margin: auto;word-break:break-all" :span="24 - 4 - 2">
            <el-avatar v-if="false" style="margin: auto;" :size="30" :src="itemDe.avatarUrl"></el-avatar>
            <div style="margin: auto;">{{itemDe.describe}}</div>

          </el-col>
          <el-col style="margin: auto" :span="2">

            <el-button  @click="showDialog(itemDe)" type="text" size="small">{{ itemDe.button }}</el-button>
          </el-col>
        </el-row>

      </div>

 

    </el-card>
<el-dialog
  title="提示"
  :visible.sync="dialogVisible"
  width="50%"
  :before-close="handleClose">
  <span>这是一段信息</span>
  <span slot="footer" class="dialog-footer">

 </span>
</el-dialog>
  </div>

</template>

<script>

// import { uploadAvatar } from '@/api/user'
import { mapGetters } from 'vuex'
export default {
    computed: {

    ...mapGetters([
      'authInfo',
      'email',
      'phone'
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
      dialogVisible:false,
      basicInfoList: [
        {
          editFlag: true,
          title: '账号设置',
          dataList: {
            nickname: { index: 0, describe: '存在风险，请设置密码 ', button: '设置密码', type: 'text', title: '密码', model: '123' },
            phone:{ index: 1, describe: ' ', button: '修改手机', type: 'text', title: '手机', model: '123' },
            email:{ index: 2, describe: ' ', button: '修改邮箱', type: 'text', title: '邮箱', model: '123' },
          }
        }, 
        {
          editFlag: true,
          title: '帐号关联',
          dataList: {
            wechat: { index: 0,avatarUrl:"https://open.weixin.qq.com/zh_CN/htmledition/res/assets/res-design-download/icon48_appwx_logo.png",desPicture:"", describe: '', button: '绑定', type: 'text', title: '微信', model: '123' },
            qq: { index: 1,avatarUrl:"https://img.ixintu.com/upload/jpg/20210522/abb7b6902f2f6fe5a399809e97aeef6c_61793_800_800.jpg!con",desPicture:"", describe: '', button: '绑定', type: 'text', title: 'QQ', model: '123' },
            github: { index: 2,avatarUrl:"http://bpic.588ku.com/element_pic/00/25/58/9656d059a2d139b.jpg",desPicture:"", describe: '', button: '绑定', type: 'text', title: 'GitHub', model: '123' },
            weibo: { index: 3,avatarUrl:"https://img.ixintu.com/download/jpg/201912/f2b9e19eae370b85e0cf0c23fc18ec30.jpg!ys",desPicture:"", describe: '', button: '绑定', type: 'text', title: '新浪微博', model: '123' },
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
    async iniUserInfo(){
     await this.$store.dispatch('user/getAuthInfo')

      console.log('created_信息')
      console.log('created_原始',this.basicInfoList[0].dataList)
      
      this.basicInfoList[0].dataList.email.describe = this.email
      this.basicInfoList[0].dataList.phone.describe = this.phone
      for (var key in this.authInfo) {//第三方信息
        let myKey = key.split("_")[0]
        let myType = key.split("_")[1]
        if (myType === "username"){
          this.basicInfoList[1].dataList[myKey].describe = this.authInfo[key]

        }
      }
    },
    cardClickFuc(index) {
      console.log('点击卡片', index)
    },
    showDialog(){
      this.dialogVisible = true

    },
    handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
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
