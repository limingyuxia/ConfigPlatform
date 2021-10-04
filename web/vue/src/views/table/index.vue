<template>
  <div class="app-container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="临时1">
        <el-input v-model="formInline.user" placeholder="临时1" />
      </el-form-item>
      <el-form-item label="临时2">
        <el-input v-model="formInline.sn" placeholder="临时2" />
      </el-form-item>
    </el-form>

    <!-- 表格 -->
    <el-card v-loading="listLoading" class="box-card">

      <div slot="header" class="clearfix">

        <el-button style="float: right; padding: 3px 0" type="text" @click="queryData">查询</el-button>
        <div style="width: 30px;text-align: center; float: right; padding: 3px 0">  |  </div>

        <el-button style="float: right; padding: 3px 0" type="text" @click="addData">添加</el-button>

      </div>

      <el-scrollbar>

        <el-table
          border
          :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
          empty-text="暂无数据"
          style="width: 100% height: 100%;"
          height="500"
          width="700"
        >
          <el-table-column
            type="index"
          />

          <el-table-column
            v-for="item,index in Header"
            :key="index"
            :prop="item.valueStr"
            :label="item.label"
          />
          <el-table-column
            fixed="right"
            label="操作"
            width="150"
          >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleClick(scope.row)">查看</el-button>
              <el-button type="text" size="small">编辑</el-button>
              <el-button type="text" size="small">删除</el-button>
            </template>
          </el-table-column>

        </el-table>
      </el-scrollbar>
      <div
        class="block"
        style="margin-top:15px;"
      >
        <el-pagination

          :current-page="currentPage"
          :page-sizes="[20]"
          :page-size="pageSize"
          :total="total"
          align="center"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <!-- 分页器 -->

  </div>
</template>

<script>
import { getList } from '@/api/table'

export default {
  data() {
    return {

      Header: [

        { 'label': '项目id', 'valueStr': 'id' },
        { 'label': '项目名称', 'valueStr': 'name' },
        { 'label': '项目描述', 'valueStr': 'description' },
        { 'label': '项目所属的部门', 'valueStr': 'department' },
        { 'label': '管理员', 'valueStr': 'admin' },
        { 'label': '创建时间', 'valueStr': 'create_time' },
        { 'label': '更新时间', 'valueStr': 'update_time' }

      ],

      tableData: [], // 表格数据
      currentPage: 1, // 当前页码
      total: 0, // 总条数
      pageSize: 20, // 每页的数据条数

      formInline: { // 表单信息
        tmp1: '',
        tmp2: ''

      },

      listLoading: false
    }
  },
  created() {
    console.log('初始化表格')
    // this.fetchData()
  },

  methods: {
    handleSizeChange(pageSize) {
      console.log('pageSize!', pageSize)
      this.pageSize = pageSize
    },
    handleCurrentChange(currentPage) {
      console.log('currentPage!', currentPage)
      this.currentPage = currentPage
    },

    queryData() {
      this.listLoading = true

      var data = {
        'dataEnd': this.formInline['dataEnd'],
        'dataStart': this.formInline['dataStart'],
        'sn': this.formInline['sn'],
        'user': this.formInline['user']
      }
      console.log('传入前:', data)
      getList(data).then(response => {
        console.log('网页返回:', response)

        this.tableData = response.list
        this.listLoading = false
        this.total = 100
      }, reason => {
        this.total = 0
        this.tableData = []
        this.listLoading = false
        console.error(reason) // 出错了！
      })
    }

  }
}
</script>
<style>
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
