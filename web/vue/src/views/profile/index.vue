<template>
  <div v-loading="user.loadingAll" class="app-container">
    <el-button type="" @click="ceshi" />
    <div v-if="user">
      <el-row :gutter="20">

        <el-col :span="6" :xs="24">
          <user-card :user="user" />
        </el-col>

        <el-col :span="18" :xs="24">
          <el-card>
            <el-tabs v-model="activeTab">
              <!--
              <el-tab-pane label="activity" name="activity">
                <activity />
              </el-tab-pane>
              -->

              <el-tab-pane label="个人资料" name="account">
                <account :user="user" @changeUser="changeUser" />
              </el-tab-pane>

              <el-tab-pane label="账号设置" name="AccountSettings">
                <AccountSettings :user="user" @changeUser="changeUser" />

              </el-tab-pane>
            </el-tabs>
          </el-card>
        </el-col>

      </el-row>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import UserCard from './components/UserCard'
// import Activity from './components/Activity'
// import Timeline from './components/Timeline'
import Account from './components/Account'
import AccountSettings from './components/AccountSettings'

export default {
  name: 'Profile',
  // components: { UserCard, Activity, Timeline, Account, AccountSettings },
  components: { UserCard, Account, AccountSettings },
  data() {
    return {

      user: {
        loadingAll: false
      },
      activeTab: 'account'
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'avatar',
      'roles',
      'email'
    ])
  },
  watch: {
    'user': function(newVal) {
      console.log('user:', newVal)
    },
    'avatar': function(newVal) {
      this.getUser()
      // console.log("avatar_ch:",newVal)
      // this.user.avatar = newVal
    }

  },
  created() {
    this.getUser()
  },

  methods: {
    ceshi() {
      this.user.avatar = this.user.avatar + 1
      console.log('userInfo:', this.user)
    },

    changeUser(user) { // 触发加载
      console.log('changeUser', user)

      for (var key in user) {
        console.log('key', key)
        console.log('user', user[key])
        this.user[key] = user[key]
      }
    },
    getUser() {
      console.log('created', this.email)

      this.user.role = []
      this.user.email = this.email
      this.user.avatar = this.avatar
      this.user.name = this.name
      console.log('user_info', this.user)
      /*
      this.user = {
        name: this.name,
        //role: this.roles.join(' | '),
        role: this.roles.join(' | '),
        email: this.email,
        avatar: this.avatar
      }*/
    }
  }
}
</script>
