<template>
  <div class="login-container">
    <!--    <span>QQ第三方登录 - 回调页面</span>-->
  </div>
</template>

<script>
import { Message } from 'element-ui'
export default {

  name: 'Proxy',
  data() {
    return {
      redirect: undefined
    }
  },

  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  created() {
    console.log('初始化表格')

    // this.qqLogin()
  },
  methods: {

  }
}
</script>
<style rel="stylesheet/scss" lang="scss"></style>
