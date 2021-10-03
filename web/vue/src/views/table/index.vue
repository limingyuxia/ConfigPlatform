<template>
  <div class="app-container">

    <el-form :inline="true" :model="formInline" class="demo-form-inline">

      <el-form-item label="区县">
        <el-input v-model="formInline.user" placeholder="审批人" />
      </el-form-item>

      <el-form-item label="查看类型">
        <el-select v-model="formInline.region" placeholder="明细">
          <el-option label="明细" value="shanghai" />
          <el-option label="统计" value="beijing" />
        </el-select>
      </el-form-item>

      <el-form-item label="开始时间">
        <el-date-picker
          v-model="formInline.dataStart"
          type="datetime"
          placeholder="选择日期时间"
          default-time="12:00:00"
        />
      </el-form-item>

      <el-form-item label="结束时间">
        <el-date-picker
          v-model="formInline.dataEnd"
          type="datetime"
          placeholder="选择日期时间"
          default-time="12:00:00"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
        <el-button type="primary" @click="onSubmit1">导出</el-button>
      </el-form-item>

    </el-form>

    <!-- 表格 -->
    <el-table
      :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
      style="width: 100%"
    >
      <el-table-column
        prop="date"
        label="日期"
        width="180"
      />
      <el-table-column
        prop="name"
        label="姓名"
        width="180"
      />
      <el-table-column
        prop="address"
        label="地址"
      />
    </el-table>
    <!-- 分页器 -->
    <div
      class="block"
      style="margin-top:15px;"
    >
      <el-pagination
        :current-page="currentPage"
        :page-sizes="[20,30]"
        :page-size="pageSize"
        :total="tableData.length"
        align="center"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { getList } from '@/api/table'

export default {

  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      input1: '',
      input2: '',
      input3: '',
      select: '',
      tableData: [
        {
          date: '2016-05-02',
          name: '第一页',
          address: ' 上海市普陀区金沙江路 1518 弄'
        },
        {
          date: '2016-05-04',
          name: '第二页',
          address: ' 上海市普陀区金沙江路 1517 弄'
        },
        {
          date: '2016-05-01',
          name: '第三页',
          address: ' 上海市普陀区金沙江路 1519 弄'
        },
        {
          date: '2016-05-03',
          name: '第四页',
          address: ' 上海市普陀区金沙江路 1516 弄'
        },
        {
          date: '2016-05-01',
          name: '第五页',
          address: ' 上海市普陀区金沙江路 1519 弄'
        },
        {
          date: '2016-05-03',
          name: '第六页',
          address: ' 上海市普陀区金沙江路 1516 弄'
        }
      ],
      currentPage: 1, // 当前页码
      total: 20, // 总条数
      pageSize: 2, // 每页的数据条数

      formInline: {
        user: '',
        region: '',
        dataStart: '',
        dataEnd: ''
      },
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    onSubmit() {
      console.log('submit!', this.formInline)
    },
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.list = response.data.items
        this.listLoading = false
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
    width: 480px;
  }
</style>
