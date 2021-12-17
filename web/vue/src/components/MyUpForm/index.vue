<template>
  <div>
      <el-form
      
        ref="elForm"
        :disabled="formDisabled"
        :model="formData"
        :rules="formRules"
        size="medium"
        label-width="auto"
        label-position="left"
      >

      <el-row class="tmp2" :gutter="20" :key="index"   v-for="item,index in ConfHeader">

          <el-col :key="colindex" :span="colitem.span || 60" v-for="colitem,colindex in item" >

            <el-form-item class="tmp1"  :label="colitem.label" v-if="colitem.type == 'tag'" :prop="colitem.model">
              <template  v-for="tagItem,tagIndex in formData[colitem.model].data">
                  <el-col  :key="tagIndex" :span=60 >
                <el-tag

                  v-if="tagItem.type == 'tag'"
                  closable
                  :disable-transitions="true"
                  @click="showInput(tagItem,tagIndex,colitem.model)"
                  @close="handleClose(tagItem,tagIndex,colitem.model)">
                  {{tagItem.label}}
                </el-tag>
                
                
                <el-input
                  :key="tagIndex"
                  class="input-new-tag"
                  v-if="tagItem.type == 'input'"
                  v-model="tagItem.inputTmp"
          
                  :ref="'saveTagInput'+tagIndex" 
                  size="small"
                  @keyup.enter.native="handleInputConfirm(tagItem,tagIndex,colitem.model)"
                  @blur="handleInputConfirm(tagItem,tagIndex,colitem.model)"
                >
                </el-input>


                </el-col>
              </template>

                <el-input
                  class="input-new-tag"
                  v-if="formDisabled == false && formData[colitem.model].inputVisible == true"
                  v-model="formData[colitem.model].inputValue"
                  
                  :ref="'saveTagInput'+colitem.model"
                  size="small"
                  @keyup.enter.native="handleInputConfirmNew(colitem.model)"
                  @blur="handleInputConfirmNew(colitem.model)"
                >
                </el-input>

                <el-button v-if="formDisabled == false && formData[colitem.model].inputVisible == false" class="button-new-tag" size="small" @click="showInputNew(colitem.model)">+ New Tag</el-button>

            </el-form-item>

              <el-form-item class="tmp1" :label="colitem.label" v-if="colitem.type == 'input'" :prop="colitem.model">
                <el-input
                  v-model="formData[colitem.model]"
                  :placeholder="'formData.' + colitem.model"
                  :readonly="colitem.readonly || false"
                  :disabled="colitem.disabled"
                />
              </el-form-item>

            </el-col>

      </el-row>

      </el-form>
  </div>
</template>

<script>
export default {
  name: 'MyUpForm',
  props: {
    inputValue:"",
    formRules:{
      type: Object,
      default: () => {}

    },
    formDisabled:{
      type: Boolean,
      default: false
    },
        formData: {
      type: Object,
      default: () => {
        return {
          id: undefined,
          tag:{
            "data":[{"label":"1","type":"tag","inputTmp":""},{"label":"2","type":"input","inputTmp":""}],
            inputVisible: false,
            inputValue: ''
          },
        }
      }
    },

    ConfHeader:{// 头配置
      type: [Array,Object],
      default:()=> [
        [
          {"label":"项目id","type":"input","span":50,"placeholder":"描述","disabled":true,"model":"id"},
           {"label":"项目id","type":"input","span":12,"placeholder":"描述","disabled":false,"model":"id1"},


        ],
               [
          {"label":"项目id","type":"input","span":50,"placeholder":"描述","disabled":true,"model":"id"},


        ],
        [
          {"label":"部门","type":"tag","span":60,"placeholder":"描述","disabled":false,"model":"tag"},

        ]

      ]
    }, 



  },
    watch: {
    'formData': function(newVal) {
      console.log('value:', newVal)
    }
  },

  methods: {
    handleClose(tagItem,tagIndex,model) {
       console.log("删除",tagItem,tagIndex,model)

       this.formData[model].data.splice(tagIndex,1)
       console.log("123",this.formData[model].data)
        //this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
      },

    showInputNew(model){
      console.log("添加",model)
      this.formData[model].inputVisible = true
      //自动对焦
      this.$nextTick(_ => {
          this.$refs["saveTagInput" + model][0].$refs.input.focus();
        });

    },
    
    handleInputConfirmNew(model){
      console.log("保存",model)
      let inputValue = this.formData[model].inputValue
      if (inputValue) {
        let tmp_inputValue = {
          "label":inputValue,"type":"tag","inputTmp":""
        }
        console.log("保存",tmp_inputValue)
          this.formData[model].data.push(tmp_inputValue);
        }
        this.formData[model].inputValue = ""
      this.formData[model].inputVisible = false
      
    },
    
      handleInputConfirm(tagItem,tagIndex,model) {
        

        let inputValue = this.formData[model].data[tagIndex].inputTmp 
        if (inputValue) {
          this.formData[model].data[tagIndex].label= inputValue;
        }
        this.formData[model].data[tagIndex].type = "tag"
        this.formData[model].data[tagIndex].inputTmp  = '';
        //this.formData[model][tagIndex].disableTransitions = true
      },

    showInput(tagItem,tagIndex,model){//开始编辑
        console.log("showInput",tagItem,tagIndex,model)
        
        
        this.formData[model].data[tagIndex].inputTmp = this.formData[model].data[tagIndex].label
        //this.formData[model][tagIndex].disableTransitions = false
        this.formData[model].data[tagIndex].type = "input"

        //自动聚焦 
      
       this.$nextTick(_ => {
          this.$refs["saveTagInput" + tagIndex][0].$refs.input.focus();
        });

    },
    getDetail(row,buttonValue) {
      this.$emit('getDetail',row,buttonValue)
    }
  }
}
</script>

<style scoped>

</style>
<style>
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
    margin-top: 5px;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 0px;
    margin-top: 4px;
   
  }
</style>
<style> 
.tmp1{
  background-color: aquamarine;
}
.tmp2{
  background-color: rgb(23, 77, 59);
}
        .el-form-item{ 

            float: left; 

        } 

    </style> 