<script setup>
import icon from '@/assets/images/logo-universal.png'
</script>

<template>
  <el-container class="body-container">
    <el-aside width="250px" class="aside-container">
      <div class="aside-div-container">
      <el-menu class="menu-container" router default-active="/Database">
        <el-menu-item index="/AutoCode">
          <el-image :src="icon" class="el-image"/>
        </el-menu-item>
        <el-menu-item index="/AutoCode" class="el-menu-item">代码生成</el-menu-item>
        <el-menu-item index="/Project" class="el-menu-item">项目管理</el-menu-item>
        <el-menu-item index="/Template" class="el-menu-item">模板管理</el-menu-item>
        <el-menu-item index="/Database" class="el-menu-item">数据库配置</el-menu-item>
      </el-menu>
      </div>
    </el-aside>
    <el-main class="main-container">
      <router-view/>
    </el-main>
  </el-container>

</template>

<style scoped>
.body-container {
  background-color: rgb(235, 235, 235);
  height: 100vh;
  width: 100%;
}

.el-menu--vertical:not(.el-menu--collapse):not(.el-menu--popup-container) .el-menu-item, .el-menu--vertical:not(.el-menu--collapse):not(.el-menu--popup-container) .el-menu-item-group__title, .el-menu--vertical:not(.el-menu--collapse):not(.el-menu--popup-container) .el-sub-menu__title {
  padding-left: 0;
}

.aside-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-color: rgb(255, 255, 255);
}
.aside-div-container{
  height: 80%;
  width: 100%;
}

.main-container {
  width: 100%;
  height: 100%;
}

.menu-container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
  width: 100%;
}

.el-menu-item {
  font-size: 20px;
  text-align: center; /* 文本居中 */
  width: 100%; /* 使每个菜单项占据整个宽度 */
  display: flex;
  justify-content: center; /* 内容居中 */
  align-items: center; /* 内容居中 */
  padding: 50px 0;
}


.el-image {
  width: 100px;
  height: 100px;
  user-select: none;
  -webkit-user-drag: none;
  -webkit-user-select: none;
  pointer-events: none;
}
</style>
