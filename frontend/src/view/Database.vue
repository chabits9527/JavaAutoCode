<script setup>
import {onMounted, reactive} from 'vue'
import {DeleteDatabaseConfig, Page} from "../../wailsjs/go/openAPI/OpenAPI.js";
import DatabaseConfigForm from "@/components/DatabaseConfigForm.vue";
import {ElMessage, ElNotification} from "element-plus";

const data = reactive({
  tableData: [],
  pageNum: 1,
  pageSize: 10,
  total: 1,
  dialogVisible: false,
  editable: false,
  DatabaseConfig: {
    id: null,
    dbType: 'Mysql',
    name: '',
    host: 'localhost',
    port: '3306',
    username: '',
    password: '',
    database: '',
    params: '?',
  }
})

onMounted(() => {
  flushTableData()
})
const dialogInit = () => {
  data.DatabaseConfig = {
    id: null,
    dbType: 'Mysql',
    name: '',
    host: 'localhost',
    port: '3306',
    username: '',
    password: '',
    database: '',
    params: '?'
  }
  data.editable = false
  data.dialogVisible = false
}

const flushTableData = () => {
  Page(data.pageNum, data.pageSize).then(res => {
    data.tableData = res.list
    data.pageNum = res.pageNum
    data.pageSize = res.pageSize
    data.total = res.total
  }).catch(err => {
    ElNotification({
      title: '获取数据库列表失败',
      message: err,
      type: 'error',
      duration: 5000
    })
  })
}
const handleEdit = (row) => {
  data.DatabaseConfig = {...row}
  data.editable = true
  data.dialogVisible = true
}
const handleSelect = (row) => {
  data.DatabaseConfig = {...row}
  data.dialogVisible = true
}
const handleClose = () => {
  dialogInit()
  flushTableData()
}
const handleDelete = (row) => {
  DeleteDatabaseConfig(row.id).then(()=> {
    ElMessage({
      message: '删除成功',
      type: 'success'
    })
    flushTableData()
  }).catch(err=>{
    ElNotification({
      title: '删除失败',
      message: err,
      type: 'error',
      duration: 5000
    })
  })
}
</script>

<template>
  <el-row justify="end">
    <el-button type="primary" @click="data.editable = true;data.dialogVisible = true">新增</el-button>
    <DatabaseConfigForm :visible="data.dialogVisible" :DatabaseConfig="data.DatabaseConfig" :editable="data.editable"
                        @closeDialog="handleClose"/>
  </el-row>
  <br/>
  <el-row justify="end" style="height: 80%;">
    <el-card style="width: 100%;">
      <el-table
          :data="data.tableData" border width="100%" height="100%"
          :header-cell-style="{background:'#95a0eb',color:'#ffffff',fontSize:'1.5em'}"
          :cell-style="{fontSize:'1.5em'}"
          stripe fit>
        <el-table-column
            prop="name"
            label="数据库名称"
            width="250">
          <template #default="scope">
            <el-button type="primary" text @click="handleSelect(scope.row)" style="font-size: 1em">{{ scope.row.name }}</el-button>
          </template>
        </el-table-column>
        <el-table-column
            prop="dbType"
            label="类型"
            width="120">
        </el-table-column>
        <el-table-column
            prop="host"
            label="主机"
            width="200">
        </el-table-column>
        <el-table-column
            prop="port"
            label="端口"
            width="80">
        </el-table-column>
        <el-table-column
            prop="username"
            label="用户名"
            width="200">
        </el-table-column>
        <el-table-column
            prop="password"
            label="密码"
            width="200">
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="150">
          <template #default="scope">
            <el-button type="primary" size="default" icon="Edit"
                       @click="handleEdit(scope.row)">编辑
            </el-button>
            <el-button type="danger" size="default" icon="Delete" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </el-row>
  <br/>
  <el-row justify="end">
    <el-pagination background layout="prev, pager, next" :total="data.total"/>
  </el-row>
</template>

<style>
</style>