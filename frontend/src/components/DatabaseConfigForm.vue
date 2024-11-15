<script setup>
import {reactive, ref} from "vue";
import {SaveDatabaseConfig, TestConnect} from "../../wailsjs/go/openAPI/OpenAPI.js";
import {ElMessage, ElNotification} from "element-plus";

const props = defineProps({
  editable: {
    type: Boolean,
    default: false
  },
  visible: {
    type: Boolean,
    default: false
  },
  DatabaseConfig: {
    type: Object,
    default: {
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
  }
});

let data = reactive({
  DatabaseConfig: {},
})

let testButtonLoading = ref(false);
let testButtonSuccess = ref(false);
const handleTest = () => {
  testButtonLoading.value = true;
  TestConnect(data.DatabaseConfig).then(() => {
    ElMessage({
      message: '连接成功',
      type: 'success',
      duration: 2000
    })
    testButtonSuccess.value = true
    testButtonLoading.value = false
  }).catch(err => {
    ElNotification({
      title: '数据库连接失败',
      message: err,
      type: 'error',
      duration: 5000
    })
    testButtonLoading.value = false
    testButtonSuccess.value = false
  })
}


const init = () => {
  data.DatabaseConfig = props.DatabaseConfig
  testButtonSuccess.value = false
  testButtonLoading.value = false
}


const emits = defineEmits(['closeDialog']);
const handleClose = () => {
  emits('closeDialog');
}
const handleSubmit = () => {
  SaveDatabaseConfig(data.DatabaseConfig).then(() => {
    ElMessage({
      message: '添加成功',
      type: 'success'
    })
    handleClose();
  }).catch(err => {
    ElNotification({
      title: '添加失败',
      message: err,
      type: 'error',
      duration: 5000
    })
  })
}


</script>

<template>
  <el-dialog
      :model-value="visible"
      :title="props.editable ? (data.DatabaseConfig.id == null ? '新增数据库配置' : '编辑数据库配置') : '数据库配置详情'"
      width="500"
      :before-close="handleClose"
      @opened="init"
  >
    <el-form
        style="max-width: 600px"
        :model="data.DatabaseConfig"
        label-width="auto"
        status-icon
    >
      <el-form-item label="数据库类型">
        <span v-if="!editable">{{ data.DatabaseConfig.dbType }}</span>
        <el-select v-if="editable" v-model="data.DatabaseConfig.dbType" placeholder="请选择数据库类型">
          <el-option label="MySQL" value="Mysql"/>
        </el-select>
      </el-form-item>

      <el-form-item label="数据库连接名称">
        <span v-if="!editable">{{ data.DatabaseConfig.name }}</span>
        <el-input v-if="editable" v-model="data.DatabaseConfig.name" placeholder="请输入数据库名称"/>
      </el-form-item>
      <el-form-item label="数据库地址">
        <span v-if="!editable">{{ data.DatabaseConfig.host }}</span>
        <el-input v-if="editable" v-model="data.DatabaseConfig.host" placeholder="请输入数据库地址"/>
      </el-form-item>
      <el-form-item label="数据库端口">
        <span v-if="!editable">{{ data.DatabaseConfig.port }}</span>
        <el-input v-if="editable" type="number" v-model="data.DatabaseConfig.port" placeholder="请输入数据库端口"/>
      </el-form-item>
      <el-form-item label="数据库用户名">
        <span v-if="!editable">{{ data.DatabaseConfig.username }}</span>
        <el-input v-if="editable" v-model="data.DatabaseConfig.username" placeholder="请输入数据库用户名"/>
      </el-form-item>
      <el-form-item label="数据库密码">
        <span v-if="!editable">{{ data.DatabaseConfig.password }}</span>
        <el-input v-if="editable" v-model="data.DatabaseConfig.password" placeholder="请输入数据库密码"/>
      </el-form-item>
      <el-form-item label="数据库名称">
        <span v-if="!editable">{{ data.DatabaseConfig.database }}</span>
        <el-input v-if="editable" v-model="data.DatabaseConfig.database" placeholder="请输入数据库名称"/>
      </el-form-item>
<!--      <el-form-item label="数据库连接参数">-->
<!--        <span v-if="!editable">{{ data.DatabaseConfig.params }}</span>-->
<!--        <el-input v-if="editable" v-model="data.DatabaseConfig.params" placeholder="请输入数据库连接参数"/>-->
<!--      </el-form-item>-->
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button size="large" @click="handleClose">取消</el-button>
        <el-button size="large" type="success"
                   @click="handleTest"
                   :loading="testButtonLoading"
                   :plain="!testButtonSuccess"
                   :icon="testButtonSuccess ? 'Check':'Van'">
          {{ testButtonSuccess ? '连接成功' : '测试连接' }}
        </el-button>
        <el-button type="primary" @click="handleSubmit" v-if="editable" size="large">
          提交
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>