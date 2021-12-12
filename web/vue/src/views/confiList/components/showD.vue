<template>
  <el-form  v-loading="loading" ref="form" :model="form" label-width="80px">
    <el-form-item label="分组">
      <el-input v-model="form.group" />
    </el-form-item>
    <el-form-item label="Key">
      <el-input v-model="form.key" />
    </el-form-item>

    <el-form-item label="类型">
      <el-radio-group v-model="form.type" @change="changeLanguage">

        <el-radio label="text/STRING">STRING</el-radio>
        <el-radio label="text/html">XML</el-radio>
        <el-radio label="text/x-yaml">YAML</el-radio>
        <el-radio label="application/json">JSON</el-radio>

      </el-radio-group>
    </el-form-item>

    <el-form-item label="Value">
      <edit v-model="form.value" :language-my="language" />

    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">确定</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import edit from './edit'

export default {
  components: { edit },
  props: {

    form: {
      type: Object,

      default: () => {
        return {
          group: '',
          key: '',
          type: 'application/json',
          value: ''
        }
      }
    }
  },
  data() {
    return {
      // value:"",
      language: '',
loading:false
    }
  },
  watch: {
    'form.value': function(newVal) {
      console.log('value:', newVal)
    }
  },
  onShow() { // 监听页面显示。页面每次出现在屏幕上都触发，包括从下级页面点返回露出当前页面
    console.log('onShow')
  },
  mounted() {
    console.log('123456:', this.form)
  },
  methods: {
    changeLanguage(value) {
      console.log('changeLanguage', value)
      this.language = value
    },
    onSubmit() {
      console.log('submit!',this.form)
      this.loading = true
          setTimeout(()=>{   //设置延迟执行
        console.log('ok');
        this.loading = false
    },1000);

      
    }
  }
}
</script>
