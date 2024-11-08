<script setup>
import {onMounted, reactive} from 'vue'
import {Page} from "../../wailsjs/go/openAPI/OpenAPI.js";
import DatabaseConfigForm from "@/components/InfoForm.vue";

const data = reactive({
  tableData: [],
  pageNum: 1,
  pageSize: 10,
  dialogVisible: false,
  editable: false,
  DatabaseConfig: {}
})

onMounted(() => {
  Page(data.pageNum, data.pageSize).then(res => {
    console.log(res)
  })
  data.tableData.push({
    id: 1,
    name: '数据库1',
    host: 'localhost',
    dbType: 'mysql',
    port: '3306',
    username: 'user',
    password: '123',
  })
})


const handleEdit = (row) => {
  data.DatabaseConfig = {...row}
  data.editable = true
  data.dialogVisible = true
}
const handleClose = () => {
  data.DatabaseConfig = {}
  data.editable = false
  data.dialogVisible = false
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
          :header-cell-style="{background:'#95a0eb',color:'#ffffff'}" stripe fit>
        <el-table-column
            prop="name"
            label="数据库名称"
            width="200">
        </el-table-column>
        <el-table-column
            prop="dbType"
            label="数据库类型"
            width="100">
        </el-table-column>
        <el-table-column
            prop="host"
            label="主机"
            width="200">
        </el-table-column>
        <el-table-column
            prop="port"
            label="端口"
            width="75">
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
            <el-button type="primary" size="small"
                       @click="handleEdit(scope.row)">编辑
            </el-button>
            <el-button type="danger" size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </el-row>
  <br/>
  <el-row justify="end">
    <el-pagination background layout="prev, pager, next" :total="1"/>
  </el-row>
</template>

<style>
</style>