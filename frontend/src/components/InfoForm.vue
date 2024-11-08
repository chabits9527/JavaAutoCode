<script setup>
import {reactive} from "vue";

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
    }
  }
});

let data = reactive({
  DatabaseConfig: {}
});
// let DatabaseConfig = reactive({});

const emits = defineEmits(['closeDialog']);
const handleClose = () => {
  emits('closeDialog');
}


</script>

<template>
  <el-dialog
      :model-value="visible"
      :title="props.editable ? (data.DatabaseConfig.id == null ? '新增数据库配置' : '编辑数据库配置') : '数据库配置详情'"
      width="500"
      :before-close="handleClose"
      @opened="data.DatabaseConfig = props.DatabaseConfig"
  >
    <!-- form表单 -->
    <p>{{ data.DatabaseConfig }}</p>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="()=>{
          console.log('edit', props.editable)
          handleClose()
        }">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>