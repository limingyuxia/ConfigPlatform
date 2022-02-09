<template>

  <el-form

    :ref="formRef"
    :disabled="formDisabled"
    :model="formData"
    :rules="formRules"
    size="medium"
    label-width="auto"
    label-position="left"
  >
    <el-row  :gutter="2" :key="index"   v-for="item,index in ConfHeader">

      <el-col :key="colindex" :span="colitem.span || 60" v-for="colitem,colindex in item" >
        <!--tag-->
        <el-form-item :rules="colitem.rules || []" class="tmp1"  :label="colitem.label" v-if="colitem.type == 'tag'" :prop="colitem.model">
          <el-row style="border:1px solid" :gutter="10" >
            
              <el-col class="tag-class" v-show="formData[colitem.model].data.length >0" v-for="tagItem,tagIndex in formData[colitem.model].data" :key="tagIndex" :span="25">
                
                <el-tag
                  style="width: 100%; "
                  
                  v-if="tagItem.type == 'tag'"
                  :ref="'saveTag'+tagIndex" 
                  closable
                  :disable-transitions="true"
                  @click="showInput(tagItem,tagIndex,colitem.model)"
                  @close="handleClose(tagItem,tagIndex,colitem.model)">
                  {{style}}
                  {{tagItem.label.length>10?tagItem.label.substring(0,9)+'...':tagItem.label}}
                </el-tag>
                <el-input
                  :key="tagIndex"
                  
                  v-if="tagItem.type == 'input'"
                  v-model="tagItem.inputTmp"  
                  :ref="'saveTagInput'+tagIndex" 
                  size="small"
                  @keyup.enter.native="handleInputConfirm(tagItem,tagIndex,colitem.model)"
                  @blur="handleInputConfirm(tagItem,tagIndex,colitem.model)"
                >
                </el-input>
              </el-col>

            
            <el-col class="tag-class" :span="25">
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
              <el-button v-if="formDisabled == false && formData[colitem.model].inputVisible == false" class="button-new-tag" size="small" @click="showInputNew(colitem.model)">{{colitem.addButton || "+ New Tag" }} </el-button>
            </el-col>
          </el-row>
        </el-form-item>
        <!--input-->
        <el-form-item :rules="colitem.rules || []" class="tmp1" :label="colitem.label" v-if="colitem.type == 'input'" :prop="colitem.model">
          <el-input
            :type="colitem.type_old || ''"
            v-model="formData[colitem.model]"
            :placeholder="colitem.placeholder"
            :readonly="colitem.readonly || false"
            :disabled="colitem.disabled"
          />
        </el-form-item>
        <!--upData-->
        <el-form-item :rules="colitem.rules || []" class="tmp1" :label="colitem.label" v-if="colitem.type == 'upData'" :prop="colitem.model">
          <el-popover
            placement="top-start"
            width="500"
            title="上传附件"
            trigger="click">
              <el-upload
                drag
                action="123"
                list-type="picture"
                :on-preview="handlePreview"
                :on-change="handleChange"

                :auto-upload="false">

                <i class="el-icon-plus"></i>

              </el-upload>
                             <el-divider></el-divider>
        <div  class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div> 
            <el-button slot="reference"><i class="el-icon-upload el-icon-left"></i>上传附件</el-button>
        </el-popover>
        
        </el-form-item>
       
        <!--select-->
        <el-form-item :rules="colitem.rules || []" class="tmp1" :label="colitem.label" v-if="colitem.type == 'select'" :prop="colitem.model">
          <el-select  v-model="formData[colitem.model]" 
                      :placeholder="colitem.placeholder"
                      :readonly="colitem.readonly || false"
                      :disabled="colitem.disabled"
                      >
            <el-option v-for="tagItem,tagIndex in colitem.data" 
                        :key="tagIndex" 
                        :label="tagItem.label" 
                        :value="tagItem.value">
            </el-option>
          </el-select>
        </el-form-item>

        <!--date_picker-->
        <el-form-item :rules="colitem.rules || []" class="tmp1" :label="colitem.label" v-if="colitem.type == 'date_picker'" :prop="colitem.model">
          <el-col>
          <el-date-picker
                  v-model="formData[colitem.model]"
                  :type="colitem.type_old || ''"
                  :placeholder="colitem.placeholder"
                  :style="{width: '85%'}"
                  :disabled="colitem.disabled"
          />
           </el-col>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>

</template>

<script>
export default {
  name: 'MyUpForm',
  props: {
    formRef:{
      type: String,
      default: ""
    },
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
          {"label":"项目id0","type":"input","span":12,"placeholder":"描述","disabled":true,"model":"id"},
           {"label":"项目id1","type":"input","span":12,"placeholder":"描述","disabled":false,"model":"id1"},


        ],
        [
          {"label":"项目id_1","type":"input","span":20,"placeholder":"描述","disabled":true,"model":"id"},


        ],
        [
          {"label":"部门-0","type":"tag","span":24,"placeholder":"描述","disabled":false,"model":"tag"},

        ],
        [
          {"label":"上传文件","type":"upData","span":24,"placeholder":"描述","disabled":false,"model":"tag"},

        ]

      ]
    }, 



  },
    watch: {
    'formData': function(newVal) {
      console.log('value:', newVal)
    }
  },
  mounted() {
    console.log('mounted:', this.formRef)
  },

  methods: {
    handleChange(file, fileList){
      var url = "https://src.pcsoft.com.cn/d/file/soft/wlgj/wlqt/2020-05-21/3f3678a651765913995ccb458c126320.jpg"
      file.url = url
      console.log(1,file);
      console.log(2,fileList);
    },
    handlePreview(file) {
        console.log(1,file);
      },
    handleClose(tagItem,tagIndex,model) {
       console.log("删除",tagItem,tagIndex,model)

       this.formData[model].data.splice(tagIndex,1)
       console.log("123",this.formData[model].data)
        //this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
      },

    showInputNew(model){
      console.log("添加",this.formData,model,this.formData[model])
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


<style lang="scss" scoped>

  .button-new-tag {
    margin-left: 0px;
    height: 32px;

  }
  .input-new-tag {
    width:auto;
    // height: 10px;
    margin-left: 0px;
    margin-top: 0px;

  }
  .tag-class{
    

      // min-width: 100px;
      margin-top: 5px;
      
      
  }


</style>