<template>
  <div class="app-container">
    <el-form :inline="true" class="demo-form-inline">

      <el-form-item v-for="item,index in formInline" :key="index" :label="item.label">
        <!--formInline.user-->
        <el-input v-model="item.value" :placeholder="item.placeholder" />
      </el-form-item>

    </el-form>
    <!-- 表格 -->
    <el-card v-loading="listLoading" class="box-card">

      <div slot="header" class="clearfix">
        <el-button style="float: right; padding: 3px 0" type="text" @click="queryDataFuc()">查询</el-button>
        <div style="width: 30px;text-align: center; float: right; padding: 3px 0">  |  </div>
        <el-button style="float: right; padding: 3px 0" type="text" @click="dialogShow([],'add')">添加</el-button>

      </div>

          <myTable  @pagination-size-change="handleSizeChange" 
                  @pagination-current-change="handleCurrentChange" 
                  :projectConfList="tableData" 
                  :projectConfHeader="Header" 
                  :pagination="paginationData" 
                  @getDetail="getDetail" />

    </el-card>

    <el-dialog v-loading="dialogFormLoading" :visible.sync="dialogFormVisible" v-bind="$attrs" title="项目详情">
      <!--
      <MyUpForm />
       -->
      <el-form
        ref="elForm"
        :disabled="dialogformDisabled"
        :model="formData"
        :rules="formRules"
        size="medium"
        label-width="100px"
        label-position="left"
      >
        <el-row type="flex" justify="start" align="bottom">
          <el-col :span="11">
            <el-form-item label="项目id" prop="id">
              <el-input
                v-model="formData.id"
                placeholder="由系统自动生成项目id项目id"
                :maxlength="11"
                show-word-limit
                readonly
                :disabled="true"
                clearable
                prefix-icon="el-icon-s-unfold"
                :style="{width: '100%'}"
              />
            </el-form-item>
          </el-col>
          <el-col :span="13">
            <el-form-item label="项目创建者" prop="project_user">
              <el-input
                v-model="formData.project_user"
                placeholder="请输入项目创建者"
                readonly
                :disabled="true"
                clearable
                :style="{width: '100%'}"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="项目名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入项目名称" clearable :style="{width: '100%'}" />
        </el-form-item>
        <el-form-item label="项目描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            placeholder="请输入项目描述"
            :autosize="{minRows: 4, maxRows: 4}"
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-row>
          <el-col :span="13">
            <el-form-item label-width="130px" label="项目所属的部门">

              <el-form-item
                v-for="item,index in formData.department"
                :key="index"
                :label="index +'级部门'"
                style="margin-bottom:10px"
              >
                <el-input v-model="formData.department[index]" style=" float:left;min-width:85%;" placeholder="请输入部门名称">
                  <el-button slot="append" icon="el-icon-close" @click.prevent="removeDepartment(index)" />
                </el-input>

              </el-form-item>
              <div v-show="!dialogformDisabled" align="center">
                <el-button @click="addDepartment">增加部门</el-button>
              </div>

            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="11">

            <el-form-item label="管理员" prop="admin">
              <div class="divTagAll">

                <el-tag
                  v-for="item,index in formData.admin"
                  :key="item.index"
                  color="white"
                  size="medium"
                  :closable="!dialogformDisabled"
                  :disable-transitions="false"
                  @close="handleClose(index,'admin')"
                >
                  {{ item }}
                </el-tag>
                <el-input
                  v-if="adminTagSys.inputVisible"
                  ref="saveTagInput"
                  v-model="adminTagSys.inputValue"
                  class="input-new-tag"
                  size="medium"
                  @keyup.enter.native="handleInputConfirm('admin')"
                  @blur="handleInputConfirm('admin')"
                />
                <el-button v-else v-show="!dialogformDisabled" class="button-new-tag" size="small" @click="showInput('admin')">添加</el-button>

              </div>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="项目研发人员" prop="develop_user">
              <div class="divTagAll">

                <el-tag
                  v-for="item,index in formData.develop_user"
                  :key="item.index"
                  color="white"
                  class="el-tag"
                  :disable-transitions="false"

                  :closable="!dialogformDisabled"
                  size="medium"
                  @close="handleClose(index,'developUser')"
                >
                  {{ item }}
                </el-tag>
                <el-input
                  v-if="developUserTagSys.inputVisible"
                  ref="saveTagInput"
                  v-model="developUserTagSys.inputValue"
                  class="input-new-tag"
                  size="medium"
                  @keyup.enter.native="handleInputConfirm('developUser')"
                  @blur="handleInputConfirm('developUser')"
                />
                <el-button v-else v-show="!dialogformDisabled" class="button-new-tag" size="medium" @click="showInput('developUser')">添加</el-button>

              </div>

            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="创建时间" prop="create_time">
              <!--
               <el-date-picker
                v-model="formData.create_time"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
-->
              <el-date-picker
                v-model="formData.create_time"
                type="datetime"
                :style="{width: '100%'}"
                placeholder="由系统自动生成"
                clearable
                :disabled="true"
              />

            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新时间" prop="update_time">
              <el-date-picker
                v-model="formData.update_time"
                type="datetime"
                :style="{width: '100%'}"
                placeholder="由系统自动生成"
                clearable
                :disabled="true"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer">
        <!--<el-button @click="close">取消</el-button>-->
        <el-button :disabled="dialogformDisabled" type="primary" @click="handelConfirm">保存更改</el-button>
      </div>
    </el-dialog>

  </div>

</template>

<script>
import { getList, getDetailList, addProject, deleteProject, editProject } from '@/api/project'
import { Message } from 'element-ui'
import { mapGetters } from 'vuex'
import myTable from '@/components/myTable'
import MyUpForm from '@/components/MyUpForm'

export default {
  components: { myTable,MyUpForm },
  computed: {

    ...mapGetters([
      'name'
    ])
  },
  data() {
    return {

      formInline: [
        { 'label': '项目名称', 'value': '', 'model': 'name', 'placeholder': '请输入' },
        { 'label': '项目id', 'value': '', 'model': 'id', 'placeholder': '请输入' },
        { 'label': '项目的开发人员', 'value': '', 'model': 'develop_user', 'placeholder': '请输入' },
        { 'label': '项目所属部门', 'value': '', 'model': 'department', 'placeholder': '请输入' },
        { 'label': '管理员', 'value': '', 'model': 'admin', 'placeholder': '请输入' }

      ],
      adminTagSys: {
        'inputVisible': false,
        'inputValue': ''
      },
      developUserTagSys: {
        'inputVisible': false,
        'inputValue': ''
      },
      formRules: {
        id: [],
        project_user: [],
        name: [{

          message: '请输入项目名称',
          trigger: 'blur'
        }],
        description: [{

          message: '请输入项目描述',
          trigger: 'blur'
        }],
        department: [{

          message: '请输入项目所属的部门',
          trigger: 'blur'
        }],
        admin: [],
        develop_user: [],
        create_time: [],
        update_time: []
      },

      formData: {
        systype: '',
        id: undefined,
        project_user: undefined,
        name: undefined,
        description: undefined,
        department: undefined,
        admin: undefined,
        develop_user: undefined,
        create_time: '23:30:09',
        update_time: null
      },

      addform: {
        'name': 'aa',
        'region': ''
      },

      Header: [

        // { 'label': '项目id', 'valueStr': 'id' ,'type':'text'},
        { 'label': '项目名称', 'valueStr': 'name', 'type': 'text' },
        // { 'label': '项目描述', 'valueStr': 'description' ,'type':'text'},
        { 'label': '项目所属的部门', 'valueStr': 'department', 'type': 'tagArray' },
        // { 'label': '管理员', 'valueStr': 'admin' ,'type':'text'},
        { 'label': '创建时间', 'valueStr': 'create_time', 'type': 'text', 'width': 180 },
        { 'label': '更新时间', 'valueStr': 'update_time', 'type': 'text', 'width': 180 },
      { 'label': '类型', 'valueStr': 'name', 'type': 'button',"data":[
        {"buttonType":"text","buttonText":"查看","buttonValue":"view"},{"buttonType":"text","buttonText":"编辑","buttonValue":"edit"},{"buttonType":"text","buttonText":"删除","buttonValue":"delete"}] 
          }
      ],
      paginationData:{
        "currentPage":1,// 当前页码
        "total":20,// 总条数
        "pageSize":20
      
      },
      tableData: [], // 表格数据
      /*
      currentPage: 1, // 当前页码
      total: 20, // 总条数
      pageSize: 20, // 每页的数据条数
*/
      dialogFormVisible: false,
      dialogVisible: false,
      dialogformDisabled: false,
      dialogFormLoading: false,
      listLoading: false
    }
  },
  created() {
    console.log('初始化表格', this)
    var data = {
      'page_index': 1,
      'page_size': 20,
      'project_user': this.name// 用户名 必须
    }
    this.queryData(data)
  },

  methods: {
    queryDataFuc() { // 查询
      console.log('queryDataFuc:')
      const formInline = this.formInline
      var data = {
        'page_index': this.paginationData.currentPage,
        'page_size': this.paginationData.pageSize,
        'project_user': this.name// 用户名 必须
      }

      for (let index = 0; index < formInline.length; index++) {
        const element = formInline[index]
        if (element.value !== '') {
          data[element.model] = element.value
        }
      }

      console.log('upData:', data)
      this.queryData(data)
    },
    handleClose(index, type) {
      if (type === 'admin') {
        this.formData.admin.splice(index, 1)
      } else if (type === 'developUser') {
        this.formData.develop_user.splice(index, 1)
      }
    },

    showInput(type) {
      if (type === 'admin') {
        this.adminTagSys.inputVisible = true
      } else if (type === 'developUser') {
        this.developUserTagSys.inputVisible = true
      } else {
        console.log('onthon')
      }
    },

    handleInputConfirm(type) {
      if (type === 'admin') {
        const inputValue = this.adminTagSys.inputValue
        if (inputValue) {
          this.formData.admin.push(inputValue)
        }
        this.adminTagSys.inputVisible = false
        this.adminTagSys.inputValue = ''
      } else if (type === 'developUser') {
        const inputValue = this.developUserTagSys.inputValue
        if (inputValue) {
          this.formData.develop_user.push(inputValue)
        }
        this.developUserTagSys.inputVisible = false
        this.developUserTagSys.inputValue = ''
      } else {
        console.log('onthon')
      }
    },
    removeDepartment(index) {
      console.log('remove:', index)
      // var index = this.formData.department.indexOf(item)
      if (index !== -1) {
        this.formData.department.splice(index, 1)
      }
    },

    addDepartment() {
      this.formData.department.push('')
    },
    onClose() {
      this.$refs['elForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },

    handelConfirm() {
      console.log('123:', this.formData)

      if (this.formData.systype === 'add') {
        const upData = {
          'admin': this.formData.admin,
          'department': this.formData.department,
          'description': this.formData.description,
          'develop_user': this.formData.develop_user,
          'name': this.formData.name,
          'project_user': this.name
        }
        console.log('upData:', upData)
        this.dialogFormLoading = true
        addProject(upData).then(response => {
          console.log('addProject', response)
          Message({
            message: response,
            type: 'success',
            duration: 5 * 1000
          })
          this.dialogFormLoading = false
          this.dialogFormVisible = false
          var data = {
            'page_index': this.paginationData.currentPage,
            'page_size': this.paginationData.pageSize,
            'project_user': this.name// 用户名 必须
          }
          this.queryData(data)
          // this.formData = response
        }, reason => {
          this.dialogFormLoading = false
          console.error(reason) // 出错了！
        })
      } else if (this.formData.systype === 'edit') {
        const upData = {
          'admin': this.formData.admin,
          'department': this.formData.department,
          'description': this.formData.description,
          'develop_user': this.formData.develop_user,
          'id': this.formData.id

        }
        console.log('upData:', upData)
        this.dialogFormLoading = true

        editProject(upData).then(response => {
          console.log('editProject', response)
          Message({
            message: response,
            type: 'success',
            duration: 5 * 1000
          })
          this.dialogFormLoading = false
          this.dialogFormVisible = false

          var data = {
            'page_index': this.paginationData.currentPage,
            'page_size': this.paginationData.pageSize,
            'project_user': this.name// 用户名 必须
          }
          this.queryData(data)
          // this.formData = response
        }, reason => {
          this.dialogFormLoading = false
          console.error(reason) // 出错了！
        })
      }
    },

    getDetail(data, type) { // 弹窗的
      console.log('dialogShow:', data, type)
      var updata = {}

      if (type === 'view') { // 查看
      // 获取详情
        updata = {
          'id': data['id'],
          'name': data['name']
        }

        this.dialogformDisabled = true
        // this.formData = data
        this.dialogFormVisible = true
        this.dialogFormLoading = true

        getDetailList(updata).then(response => {
          console.log('DatailList', response)
          this.dialogFormLoading = false
          this.formData = response
        }, reason => {
          this.dialogFormLoading = false
          console.error(reason) // 出错了！
        })
      } else if (type === 'edit') { // 编辑
        updata = {
          'id': data['id'],
          'name': data['name']
        }

        this.dialogformDisabled = false
        // this.formData = data
        this.dialogFormVisible = true
        this.dialogFormLoading = true

        getDetailList(updata).then(response => {
          console.log('DatailList', response)
          this.dialogFormLoading = false
          this.formData = response
          this.formData.systype = 'edit'
        }, reason => {
          this.dialogFormLoading = false
          console.error(reason) // 出错了！
        })
      } else if (type === 'add') { // 异常
        console.log('onthon')
      }
      else if(type === 'delete'){//删除项目
        this._deleteClick(data)
      }
      // dialogformDisabled
    },

    dialogShow(data, type) { // 添加弹窗
      console.log('dialogShow:', data, type)

      data = {
        'systype': 'add',
        'admin': [],
        'department': [],
        'description': '',
        'develop_user': [],
        'id': '',
        'name': '',
        'project_user': this.name
      }
      this.dialogformDisabled = false
      this.formData = data
      this.dialogFormVisible = true

      // 获取详情数据
    },

    _deleteClick(data) { // 删除项目
      console.log('delete:', data)
      this.$confirm('此操作将永久删除该项目, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        console.log('delete_1:', data)
        var upData = {
          'id': data.id,
          'project_user': data.project_user
        }

        deleteProject(upData).then(response => {
          console.log('deleteProject', response)

          Message({
            message: response,
            type: 'success',
            duration: 5 * 1000
          })
          // 取余数
          const remainder = this.paginationData.total % this.paginationData.pageSize
          if (remainder === 1 && this.paginationData.currentPage > 1) {
            this.paginationData.currentPage = this.paginationData.currentPage - 1
          }
          var data = {
            'page_index': this.paginationData.currentPage,
            'page_size': this.paginationData.pageSize,
            'project_user': this.name// 用户名 必须
          }
          this.queryData(data)
        }, reason => {
          this.dialogFormLoading = false
          console.log('deleteProjectError', reason)
          console.error(reason) // 出错了！
        })
      }).catch((e) => {
        console.log('deleteProjectError_12', e)
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleSizeChange(pageSize) {
      console.log('pageSize!', pageSize)
      this.paginationData.pageSize = pageSize

      var data = {
        'page_index': this.paginationData.currentPage,
        'page_size': pageSize,
        'project_user': this.name// 用户名 必须
      }
      this.queryData(data)
    },
    handleCurrentChange(currentPage) {
      console.log('currentPage!', currentPage)
      var data = {
        'page_index': currentPage,
        'page_size': this.paginationData.pageSize,
        'project_user': this.name// 用户名 必须
      }
      this.queryData(data)

      this.paginationData.currentPage = currentPage
    },

    queryData(data) {
      this.listLoading = true

      /*
      var data = {
        'page_index': currentPage,
        'page_size': this.paginationData.pageSize,
        'project_user': this.name// 用户名 必须
      }
      */
      console.log('upData:', data)
      this.tableData = []
      /*
       currentPage: 1, // 当前页码
      total: 0, // 总条数
      pageSize: 20, // 每页的数据条数
*/
/*
      for (let index = 0; index < (data.page_index - 1) * data.page_size; index++) {
        this.tableData.push(index)
      }
*/
      console.log('this.tableData_tmp:', this.tableData, this.paginationData.currentPage, this.paginationData.pageSize,this.paginationData)

      getList(data).then(response => {
        const tableDataList = response.project_list
        const total = response.total

        console.log('tableData', response, total)
        if (total === 0) {
          this.paginationData.total = 0
          this.tableData = []
          this.listLoading = false
          return 0
        }

        if (tableDataList === null) {
          this.paginationData.total = total
          this.tableData = []
          this.listLoading = false
          return 0
        }
        

        this.tableData = tableDataList

        console.log('this.tableData:', this.tableData, this.paginationData.currentPage, this.paginationData.pageSize,total)

        this.paginationData.total = total
        this.listLoading = false
      }, reason => {
        this.paginationData.total = 0
        this.tableData = []
        this.listLoading = false
        console.error(reason) // 出错了！
      })
    }

  }
}
</script>
<style>
.divTagAll{
  border:1px solid
}
.el-tag{
    white-space: normal;
    height:auto;
    display:block;
    margin-bottom: 5px;
    margin-top: 5px;

}
.el-textarea .el-textarea__inner{
  resize: none;
}
  .text {
    font-size: 14px;
  }

  .item {
    padding: 18px 0;
  }

  .box-card {
    width: 100%;
  }
</style>
