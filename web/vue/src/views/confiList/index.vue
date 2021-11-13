<template>

  <el-container>
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
          max-height="500"
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
          <el-table

            :data="projectConfList.list"
            empty-text="暂无数据"
            style="width: 100% height: 100%;"
            height="497"
            width="700"
            :header-cell-style="{'border': '1px solid #EEEEEE',background:'#F7F7F7',color:'#606266','text-align':'center'}"
          >
            <el-table-column
              label="序号"
              type="index"
            />
            <template v-for="item,index in projectConfHeader">

              <el-table-column
                v-if="item.type == 'text'"
                :key="index"
                :prop="item.valueStr"
                :label="item.label"
                :width="item.width"
                align="center"
              />
              <el-table-column
                v-if="item.type == 'tagArray'"
                :key="index"
                align="center"
                :label="item.label"
                :width="item.width"
              >

                <template slot-scope="scope">
                  <template v-for="item1,index1 in scope.row.department">
                    <el-tag
                      v-if="item1 !== ''"
                      :key="index1"
                      style="margin-left: 5px;margin-right: 5px;"
                      size="medium"
                    >{{ item1 }}</el-tag>
                  </template>
                </template>

              </el-table-column>

            </template>

            <el-table-column
              fixed="right"
              label="操作"
              width="150"
            >
              <template slot-scope="scope">
                <el-button type="text" size="small" @click="getDetail(scope.row,'view')">编辑</el-button>
                <el-button type="text" size="small" @click="getDetail(scope.row,'edit')">删除</el-button>
                <el-button type="text" size="small" @click="deleteClick(scope.row)">历史版本</el-button>
              </template>
            </el-table-column>

          </el-table>
        </el-main>

      </el-container>
    </el-container>

  </el-container>

</template>

<script>

export default {
  data() {
    return {
      projectConfHeader: [// 表头配置
        { 'label': '状态', 'valueStr': 'department', 'type': 'tagArray', 'width': '140' },
        { 'label': 'Key', 'valueStr': 'name', 'type': 'text' }
      ],
      projectConfList: { // 项目配置列表
        title: '临时',
        list: [{
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
        }, {
          name: '王小虎'
        }]
      },
      confGroupList: { // 配置分组列表
        value: '',
        search: '',
        list: [{
          name: '王小虎'
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
  },
  methods: {
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
</style>
