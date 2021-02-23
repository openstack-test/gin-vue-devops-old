<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">创建</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">删除</el-button>
          </el-popover>
        </el-form-item>
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
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="ID" prop="id" width="100"></el-table-column>    
    <el-table-column label="应用" prop="deployment" width="300"></el-table-column> 
    
    <el-table-column label="实例数" prop="replicas" width="200"></el-table-column>
    
    <el-table-column label="创建时间" prop="time" width="200"></el-table-column> 
    <el-table-column label="容器列表" prop="pod" width="200">
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
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateK8sDeployment(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteK8sDeployment(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="创建应用">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="命名空间:">
           <el-input v-model="formData.deployment" clearable placeholder="请选择" ></el-input>
         </el-form-item>
       
         <el-form-item label="应用名称:">
            <el-input v-model="formData.deployment" clearable placeholder="请输入" ></el-input>
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
    createK8sDeployment,
    deleteK8sDeployment,
    deleteK8sDeploymentByIds,
    updateK8sDeployment,
    findK8sDeployment,
    getK8sDeploymentList
} from "@/api/k8sDeployments";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "K8sDeployment",
  mixins: [infoList],
  data() {
    return {
      listApi: getK8sDeploymentList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            deployment:"",
            
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
    handleView(row) {
      this.$router.push({ name: 'k8sPods', query: { deployment: row.deployment, processId: row.process }})
    },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10        
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteK8sDeploymentByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateK8sDeployment(row) {
      const res = await findK8sDeployment({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rek8sDeployments;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          deployment:"",
          
      };
    },
    async deleteK8sDeployment(row) {
      this.visible = false;
      const res = await deleteK8sDeployment({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createK8sDeployment(this.formData);
          break;
        case "update":
          res = await updateK8sDeployment(this.formData);
          break;
        default:
          res = await createK8sDeployment(this.formData);
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

<style scoped>
</style>