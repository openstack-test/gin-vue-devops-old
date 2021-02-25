<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      </el-form>
    </div>
    <el-table
      :data="tableData.slice((page-1)*pageSize,page*pageSize)"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="ID" prop="id" width="100"></el-table-column>
    <el-table-column label="命名空间" prop="namespace" width="200"></el-table-column> 
    
    <el-table-column label="运行状态" prop="status" width="200"></el-table-column> 
    
    <el-table-column label="创建时间" prop="createTime" width="200"></el-table-column> 
	<el-table-column label="应用列表" prop="application">
    <template slot-scope="scope">
    <el-button
              v-permisaction="['process:list:upcoming:select']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleView(scope.row)"
            >查看</el-button>
  </template>
  </el-table-column>   
  </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import {
    getK8sNamespacesList
} from "@/api/k8sNamespaces";

import infoList from "@/mixins/infoList";
export default {
  name: "K8sNamespaces",
  mixins: [infoList],
  data() {
    return {
      listApi: getK8sNamespacesList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            namespace:"",
            status:"",
            time:"",
      }
    };
  },

  async created() {
    await this.getTableData();
},

  methods: {
    handleView(row) {
      this.$router.push({ name: 'k8sDeployments', query: { namespace: row.namespace, processId: row.process }})
    },
  }
};
</script>

<style scoped>
</style>