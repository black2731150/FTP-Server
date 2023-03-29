<template>
  <div class="wrapper">
    <div class="body">
      <el-table
        :data="fileList"
        style="height: 100%"
        @cell-dblclick="handledbClick"
      >
        <el-table-column
          prop="Name"
          label="文件名"
          min-width="55%"
        ></el-table-column>

        <el-table-column
          prop="Branch"
          label="分支"
          min-width="40%"
        ></el-table-column>

        <el-table-column
          prop="Size"
          label="大小"
          min-width="30%"
        ></el-table-column>

        <el-table-column
          prop="CreateTime"
          label="上传时间"
          min-width="50%"
        ></el-table-column>

        <el-table-column label="描述" min-width="50%">
          <template #default="scope">
            <span v-if="!scope.row.isEdit">{{ scope.row.Text }}</span>
            <el-input
              v-else
              v-model="desc"
              :ref="'editInput' + scope.$index"
              @blur="handleInputBlur(scope.$index)"
              @keyup.enter="updateDesc(scope.row)"
            >
            </el-input>
          </template>
        </el-table-column>

        <el-table-column
          prop="Md5"
          label="Md5"
          min-width="60%"
        ></el-table-column>

        <el-table-column label="下载" min-width="40%">
          <template #default="scope">
            <a href=":;javascript" @click.prevent="handleDLUrl(scope.row)"
              >下载</a
            >
          </template>
        </el-table-column>
      </el-table>
    </div>

    <div class="footer">
      <div class="pagination">
        <el-pagination
          v-model:current-page="currPage"
          v-model:page-size="pageSize"
          :page-sizes="[15, 20, 50, 100]"
          layout="total, prev, pager, next, jumper,sizes"
          :total="packagesNum"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
      <div class="search">
        <el-input
          v-model="searchKW"
          placeholder="请输入搜索内容"
          class="inputBox"
        ></el-input>
        <el-button type="primary" @click="handleSearch">
          <el-icon>
            <Search />
          </el-icon>
          <span style="vertical-align: middle"> 搜 索 </span>
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { Search } from "@element-plus/icons-vue";
import zhLocale from "element-plus/lib/locale/lang/zh-cn";

export default {
  name: "HomeView",
  components: {
    Search,
  },
  data() {
    return {
      fileList: [],
      pageNum: 1,
      currPage: 1,
      pageSize: 15,
      packagesNum: 0,
      searchKW: "",
      editMode: false,
      desc: "",
    };
  },
  mounted() {
    zhLocale.el.pagination = {
      pagesize: "条/页",
      total: `共{total}条`,
      goto: "前往第",
      pageClassifier: "页",
    };
    this.initData();
  },
  methods: {
    initData() {
      this.fileList = [];
      axios({
        method: "get",
        url: "/getFileInfo",
        params: {
          page: this.currPage,
          size: this.pageSize,
          search: this.searchKW,
        },
      }).then((res) => {
        if (res) {
          const { data, total, size } = res.data;
          for (let key in data) {
            data[key]["isEdit"] = false;
            let packageSize = data[key]["Size"];
            packageSize = (packageSize / (1024 * 1024)).toFixed(2);
            data[key]["Size"] = String(packageSize) + "MB";
            data[key]["index"] = this.fileList.length;
            this.fileList.push(data[key]);
          }
          if (this.searchKW.length === 0) {
            this.packagesNum = total;
          } else {
            this.packagesNum = size;
          }
          this.pageNum = Math.ceil(this.packagesNum / this.pageSize);
        }
      });
    },
    handleDLUrl(row) {
      window.location.href = row.DownLandURL;
    },
    handleSizeChange(val) {
      this.pageSize = val;
      this.currPage = 1;
      this.initData();
    },
    handleCurrentChange(val) {
      this.currPage = val;
      this.initData();
    },
    handleSearch() {
      this.initData();
    },
    handleInputBlur(rowIdx) {
      if (this.editMode) {
        this.fileList[rowIdx].isEdit = false;
        this.editMode = false;
      }
    },
    handledbClick(row, column) {
      if (!this.editMode && column.label === "描述") {
        this.fileList[row.index].isEdit = true;
        this.$forceUpdate();
        this.desc = row.Text;
        let id = "editInput" + row.index;
        this.$nextTick(() => {
          this.$refs[id].focus();
        });
        this.editMode = true;
      }
    },
    updateDesc(row) {
      if (this.desc.length !== 0) {
        axios
          .post("/update", `name=${row.Name}&text=${this.desc}`)
          .then((res) => {
            if (res) {
              this.editMode = false;
              this.initData();
            }
          });
      } else {
        this.$message.warning("输入内容不可为空！");
      }
    },
  },
};
</script>

<style>
.wrapper {
  position: absolute;
  padding: 20px 20px;
  width: 100%;
  height: 100%;
  background-color: #fff;
}
.body {
  height: 90%;
  width: 100%;
}
.footer {
  margin-top: 20px;
  height: 10%;
  min-height: 10px;
  position: relative;
  bottom: 0;
}
.pagination {
  margin: 0;
  position: absolute;
  left: 50%;
  transform: translate(-50%, 0);
}
.search {
  float: left;
}
.inputBox {
  width: 150px !important;
  margin-right: 5px;
}
</style>
