<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">          
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
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
    <el-table-column label="容器" prop="podName" width="300"></el-table-column> 
    <el-table-column label="容器IP" prop="podIP" width="200"></el-table-column>
    <el-table-column label="主机IP" prop="hostIP" width="200"></el-table-column>
    <el-table-column label="状态" prop="status" width="150"></el-table-column> 
    <el-table-column label="启动时间" prop="startTime" width="200"></el-table-column>   
    <el-table-column label="重启次数" prop="restartCount" width="100"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateK8sPods(scope.row)" size="small" type="primary" icon="el-icon-edit">查看日志</el-button>
          <el-button class="table-button" @click="updateK8sPods(scope.row)" size="small" type="primary" icon="el-icon-s-promotion">进入容器</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要重启吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteK8sPods(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-warning" size="mini" slot="reference">重启</el-button>
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
  </div>
</template>

<script>
import {
    createK8sPods,
    deleteK8sPods,
    deleteK8sPodsByIds,
    updateK8sPods,
    findK8sPods,
    getK8sPodsList
} from "@/api/k8sPods";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "K8sPods",
  mixins: [infoList],
  data() {
    return {
      listApi: getK8sPodsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            id:0,
            pod:"",
            
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
        const res = await deleteK8sPodsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateK8sPods(row) {
      const res = await findK8sPods({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rek8sPods;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          id:0,
          pod:"",
          
      };
    },
    async deleteK8sPods(row) {
      this.visible = false;
      const res = await deleteK8sPods({ ID: row.ID });
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
          res = await createK8sPods(this.formData);
          break;
        case "update":
          res = await updateK8sPods(this.formData);
          break;
        default:
          res = await createK8sPods(this.formData);
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