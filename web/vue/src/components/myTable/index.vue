<template>
  <div>
    <el-table
      :data="projectConfList"
      empty-text="暂无数据"
      style="width: 100% height: 100%;"
      height="450"
      :header-cell-style="{'border': '1px solid #EEEEEE',background:'#F7F7F7',color:'#606266','text-align':'center','height': '1px'}"
    >

    <el-table-column
      label="序号"
      type="index"
      width="70px"
      align="center"
    />
    <template v-for="item,index in projectConfHeader">
      <el-table-column
        :key="index"
        v-if="item.type == 'button'"
        fixed="right"
        label="操作"
        width="150">
            
        <template slot-scope="scope">
          <template v-for="buttonItem,buttonIndex in item.data">
            <el-button :key="buttonIndex" type="text" size="small" @click="getDetail(scope.row,buttonItem.buttonValue)">{{buttonItem.buttonText}}</el-button>
          </template>
        </template>
              
    </el-table-column>
      <el-table-column
        v-if="item.type == 'text'"
        :key="index"
        :prop="item.valueStr"
        :label="item.label"
        :width="item.width"
        align="center"
      />

      <el-table-column
        v-if="item.type == 'desc'"
        :key="index"
        :prop="item.valueStr"
        :label="item.label"
        :width="item.width"
        align="center"
      >

        <template
          slot-scope="scope"
          style="width:400px;
          height:40px;
          border:1px solid red;"
        >
          <p
            style="overflow: hidden;
            text-overflow: ellipsis;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 1;
            "
          >
            {{ scope.row.Value }}
          </p>
          
        </template>
      </el-table-column>
      <!--:filters="[{ text: '开发中', value: '开发中' },{ text: '运行中', value: '运行中' }]"-->
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
            >{{ scope.row.department[index1] }} </el-tag>
          </template>
        </template>

      </el-table-column>

    </template>

          </el-table>

     <div
        class="block"
        style="margin-top:10px;"
      >
        <el-pagination

          :current-page="pagination.currentPage"
          :page-sizes="[5,10,20]"
          :page-size="pagination.pageSize"
          :total="pagination.total"
          align="center"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
        
      </div>
  </div>
</template>

<script>
export default {
  name: 'myTable',
  props: {
    pagination: {
      type: Object,
      default: () => {
        return {
          currentPage:0, //当前页数
          pageSize:0,      //每页个数
          total:0          //总数
        }
      }
    },
    projectConfHeader:{// 头配置
      type: Array,
      default: []
    }, 

    projectConfList:{// 项目配置列表
      type: Array,
      default: []
    }, 

  },
  methods: {
    handleSizeChange(row){
      this.$emit('pagination-size-change',row)
    },
    handleCurrentChange(row){

      this.$emit('pagination-current-change',row)

    },
    getDetail(row,buttonValue) {
      this.$emit('getDetail',row,buttonValue)
    }
  }
}
</script>

<style scoped>

</style>
