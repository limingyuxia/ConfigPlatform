<template>
  <div>

    <div>
      登录方式：
    </div>
    <div>
      登录数据：{{ search }}
    </div>
    <div>
      验证结果：
    </div>
    <el-button type="">12</el-button>

  </div>

</template>

<script>

export default {

  name: 'Proxy',
  data() {
    return {
      search: 'ss'

    }
  },

  created() {
    // let query = this.$route.query
    console.log('初始话github')
  },
  mounted() {
    console.log('mounted', window.location)
    var myURL = new URL(window.location.href)
    console.log('mounted_myURL', myURL)
    this.search = window.location.search
    // console.log(window.location.href.indexOf("code=") !==-1);判断是否否找到code=
    if (window.location.href.indexOf('code=') == -1) {

    } else {
      var str = window.location.href.substr(window.location.href.indexOf('code=') + 5)
      // console.log(str);
      var Code = str.substr()
      console.log(Code)
      var upData = {
        'Code': window.location.search,
        'typy': 'qq'
      }
      // 后端接口
      this.$store.dispatch('user/thirdLogin', upData).then((webdata) => {
        console.log('登录成功,', webdata)
        this.$router.push({ path: '/dashboard' })
      }).catch((e) => {
        console.log('登录失败,', e)
        // this.$router.push({ path:  '/dashboard' })
        this.$router.push({ path: '/login' })
      })
    }
    // return Code;
  },

  methods: {

  }
}
</script>

