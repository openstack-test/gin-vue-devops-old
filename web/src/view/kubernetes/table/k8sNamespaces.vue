<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="ID" width="180">
         <template slot-scope="scope">{{scope.row.ID}}</template>
    </el-table-column>
    
    <el-table-column label="namespace" prop="namespace" width="200"></el-table-column> 
    
    <el-table-column label="status" prop="status" width="200"></el-table-column> 
    
    <el-table-column label="time" prop="time" width="200"></el-table-column> 
    
      <el-table-column label="更多操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateK8sNamespaces(scope.row)" size="small" type="primary" icon="el-icon-edit">修改</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增界面">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="namespace名称:">
            <el-input v-model="formData.namespace" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createK8sNamespaces,
    updateK8sNamespaces,
    findK8sNamespaces,
    getK8sNamespacesList
} from "@/api/k8sNamespaces";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
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
            age:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10       
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async updateK8sNamespaces(row) {
      const res = await findK8sNamespaces({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rek8sNamespaces;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          namespace:"",
          status:"",
          age:"",
          
      };
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createK8sNamespaces(this.formData);
          break;
        case "update":
          res = await updateK8sNamespaces(this.formData);
          break;
        default:
          res = await createK8sNamespaces(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>