<template>

  <el-container>
    <el-dialog
      :visible.sync="dialogData.Visible"
      width="50%"
      :title="dialogData.title"
      center
    >
      <showD :form="showData" />
    </el-dialog>

    <el-header :inline="true" class="demo-form-inline">

      <span style="margin-right: 10px;" type="info">项目</span>

      <el-select v-model="projectList.value" filterable placeholder="请选择">
        <el-option
          v-for="item in projectList.list"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

      <span style="margin-left: 10px;margin-right: 10px;" type="info">环境</span>
      <el-select v-model="envList.value" filterable placeholder="请选择">
        <el-option
          v-for="item in envList.list"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-header>

    <el-container id="box" style="height: 600px; border: 1px solid #eee">   <!--all-->

      <el-aside id="left" width="200px" style="background-color: rgb(238, 241, 246)">   <!--lift-->

        <el-input
          v-model="confGroupList.search"
          size="medium"
          placeholder="输入关键字搜索"
        >
          <el-button slot="append" icon="el-icon-plus" circle />
        </el-input>

        <el-table
          :show-header="false"
          max-height="300px"
          :border="true"
          :data="confGroupList.list.filter(data => !confGroupList.search || data.name.toLowerCase().includes(confGroupList.search.toLowerCase()))"
          @current-change="handleConfGroupChange"
        >
          <el-table-column

            prop="name"
          />

        </el-table>

      </el-aside>

      <div id="resize" />

      <el-container id="right" style="border-top: 1px solid #b5b9a9; ">  <!--right-->

        <el-header style=" font-size: 12px">
          <span style="font-size:18px ;font-family: Microsoft Yahei;"> {{ projectConfList.title }}</span>
        </el-header>

        <el-main>
          <myTable :projectConfList="projectConfList.list" :projectConfHeader="projectConfHeader" @getDetail="getDetail" />
        </el-main>

      </el-container>
    </el-container>

  </el-container>

</template>

<script>

import showD from './components/showD'
import myTable from '@/components/myTable'

import { getList } from '@/api/project'
import { mapGetters } from 'vuex'
export default {

  components: { showD,myTable },
  computed: {

    ...mapGetters([
      'name'
    ])
  },
  data() {
    return {
      input3: '',
      select: '',
      showData: {
        group: 'sdasd',
        key: '',
        type: 'application/json',
        value: ''
      },
      dialogData: {
        'Visible': false,
        'title': '123'

      },
      projectConfHeader: [// 表头配置
        { 'label': 'Key', 'valueStr': 'name', 'type': 'text' },
        { 'label': 'Value', 'valueStr': 'Value', 'type': 'desc' },
        { 'label': '备注', 'valueStr': 'name', 'type': 'text' },
        { 'label': '状态', 'valueStr': 'department', 'type': 'tagArray', 'width': '140' },
        { 'label': '类型', 'valueStr': 'name', 'type': 'text' },
        { 'label': '类型', 'valueStr': 'name', 'type': 'button',"data":[
        {"buttonType":"text","buttonText":"编辑","buttonValue":"edit1"},{"buttonType":"text","buttonText":"删除","buttonValue":"delete"},{"buttonType":"text","buttonText":"历史版本","buttonValue":"history"}] 
          }
      ],
      projectConfList: { // 项目配置列表
        title: '临时',
        list: [{
          Value: { 'adasd': '4561654984ada4561654984adasda4561654984adasdasda4561654984ada4561654984adasda4561654984adasdasda' },
          name: '王小虎3',
          department: ['开发中']
        }, {
          name: '王小虎',
          department: ['开发中']
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }, {
          name: '王小虎'
        }]
      },
      confGroupList: { // 配置分组列表
        value: '',
        search: '',
        list: [{
          name: '王小虎'
        }, {
          name: '密钥'
        }]
      },
      envList: { // 环境列表
        value: '',
        list: [{
          value: '选项1',
          label: 'dev'
        }]
      },
      projectList: { // 项目信息
        value: '',
        list: [{
          value: '选项1',
          label: '北方民族大学'
        }]
      },
      name1: 'Dashboard'
    }
  },

  mounted() {
    this.dragControllerDiv()
    this.iniData()
  },
  methods: {
    historyClick(row){
      console.log('历史版本', row)
      this.projectConfList.title="历史"

      this.projectConfList.list = [{
          Value: { 'adasd': '4561654984ada4561654984adasda4561654984adasdasda4561654984ada4561654984adasda4561654984adasdasda' },
          name: '王小虎3',
          department: ['开发中']
        }]

      this.projectConfHeader =  [// 表头配置
        { 'label': '版本名', 'valueStr': 'name', 'type': 'text' },
        { 'label': 'Key', 'valueStr': 'name', 'type': 'text' },
        { 'label': 'Value', 'valueStr': 'Value', 'type': 'value' },
        { 'label': '修改时间', 'valueStr': 'name', 'type': 'text' },
        { 'label': '修改人', 'valueStr': 'name', 'type': 'text' },
        { 'label': '备注', 'valueStr': 'name', 'type': 'text' },

      ]

    },
    getDetail(value, row, column) {
      console.log('编辑变量', value)
      console.log('编辑变量row', row)
      
      this.showData = {
        group: '分组',
        key: 'key',
        type: 'application/json',
        value: '变量33'
      },

      this.dialogData.Visible = true
    },
    filterTag(value, row, column) { // 删选标签
      if (row.department === undefined) {
        return false
      }

      const index = row.department.indexOf(value)
      if (index === -1) {
        return false
      }
      return true
    },
    iniData() {
      console.log('初始化数据')
      this.projectList.list = []
      const data = {
        page_index: 1,
        page_size: 5,
        project_user: this.name
      }
      getList(data).then(response => {
        const tableDataList = response.project_list
        const total = response.Total

        console.log('tableData', response, total === 0)
        if (total === 0) {
          return 0
        }

        if (tableDataList === null) {
          return 0
        }

        for (let index = 0; index < tableDataList.length; index++) {
          const element = tableDataList[index]
          const dataJson = {
            value: element.id,
            label: element.name
          }

          this.projectList.list.push(dataJson)
        }
      }, reason => {
        console.error(reason) // 出错了！
      })
    },
    handleConfGroupChange(val) {
      console.log('handleConfGroupChange:', val)
      this.projectConfList.title = val.name
    },
    dragControllerDiv: function() {
      const resize = document.getElementById('resize')
      const left = document.getElementById('left')
      const right = document.getElementById('right')
      const box = document.getElementById('box')

      resize.onmousedown = function(e) {
        const startX = e.clientX
        resize.left = resize.offsetLeft
        document.onmousemove = function(e) {
          const endX = e.clientX
          let moveLen = resize.left + (endX - startX)
          const maxT = box.clientWidth - resize.offsetWidth
          if (moveLen < 150) moveLen = 150
          if (moveLen > maxT - 800) moveLen = maxT - 800
          resize.style.left = moveLen
          left.style.width = moveLen + 'px'
          right.style.width = (box.clientWidth - moveLen - 5) + 'px'
        }
        document.onmouseup = function() {
          document.onmousemove = null
          document.onmouseup = null
          resize.releaseCapture && resize.releaseCapture()
        }
        resize.setCapture && resize.setCapture()
        return false
      }
    }
  }

}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>

<style>
  .el-select .el-input {
    width: 130px;
  }
  .input-with-select .el-input-group__prepend {
    background-color: #fff;
  }
  .el-header{
    background-color: #d1d1d1;

    color: #333;
   /* text-align: center;   // 居中
   background-color: #d1d1d1;
   */
    line-height: 60px;

  }

  .el-aside {
    background-color: #ffffff;
    color: rgb(255, 255, 255);

    /* text-align: center;*/
    /*line-height: 50px;*/
  }

  .el-main {
    background-color: #ffffff;
    color: rgb(255, 255, 255);
    text-align: center;
    /*line-height: 160px;*/
  }

</style>
<style scoped>
  #box{
    width:100%;
    height:100%;
    position: relative;
    overflow:hidden;
  }
  .top {
    width: 100%;
    height: 100%;
    background: #ffe0c6;
  }
  #left{
    width:calc(70% - 5px);
    height: 100%;
    float:left;
    overflow: auto;
    background: pink;
  }

  #resize{
    padding: 2px;
    z-index: 10;
    background-color: #bfe2ff;
    color: #fff;
    position: relative;
    width:1px;
    height:20rpx;
    cursor: w-resize;
    float:left;
  }

  #right{
    width:70%;
    height:100%;
    float:left;
    overflow: hidden;
    background: #ffc5c1;
  }
  .value_class{

    display: -webkit-box;/*作为弹性伸缩盒子模型显示*/
    -webkit-line-clamp: 1; /*显示的行数；如果要设置2行加...则设置为2*/
    overflow: hidden; /*超出的文本隐藏*/
    text-overflow: ellipsis; /* 溢出用省略号*/
    -webkit-box-orient: vertical;/*伸缩盒子的子元素排列：从上到下*/

  }

</style>
