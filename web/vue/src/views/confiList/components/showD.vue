<template>
<el-form ref="form" :model="form" label-width="80px">
  <el-form-item label="分组">
    <el-input v-model="form.group"></el-input>
  </el-form-item>
  <el-form-item label="Key">
    <el-input v-model="form.key"></el-input>
  </el-form-item>


  <el-form-item label="类型">
    <el-radio-group @change="changeLanguage" v-model="form.type">

      <el-radio label="text/STRING">STRING</el-radio>
      <el-radio label="text/html">XML</el-radio>
      <el-radio label="text/x-yaml">YAML</el-radio>
      <el-radio label="application/json">JSON</el-radio>

    
    </el-radio-group>
  </el-form-item>

  <el-form-item label="Value">
     <edit :languageMy="language" v-model="value"  />
    
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
      watch: {
        'value': function(newVal) {
          console.log('value:', newVal)

        }
  },
      components: { edit },
    data() {
      return {
        value:"",
        language:"",
        form: {
          group:'',
          key: '',
          type:'',
          value:''

        }
      }
    },
    methods: {
      changeLanguage(value){
        console.log("changeLanguage",value)
        this.language =value
        
      },
      onSubmit() {
        console.log('submit!');
      }
    }
  }
</script>