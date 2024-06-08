<template>
  <el-container>
    <el-header class="header">
      货物管理系统
    </el-header>
    <el-container>
      <el-aside width="200px" class="aside">
        <el-menu
            default-active="1"
            class="el-menu-vertical-demo"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
        >
          <el-menu-item index="1" @click="navigateTo('/dashboard/add')">
            <i class="el-icon-menu"></i>
            <span slot="title">增加</span>
          </el-menu-item>
          <el-menu-item index="2" @click="navigateTo('/dashboard/restock')">
            <i class="el-icon-menu"></i>
            <span slot="title">进货</span>
          </el-menu-item>
          <el-menu-item index="3" @click="navigateTo('/dashboard/products')">
            <i class="el-icon-menu"></i>
            <span slot="title">商店库存</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main class="main">
        <!-- 主内容区域 -->
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  methods: {
    navigateTo(route) {
      this.$router.push(route);
    },
  }
};
</script>

<style>
.header {
  background-color: #545c64;
  color: #fff;
  text-align: center;
  line-height: 60px;
  font-size: 24px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.aside {
  background-color: #f2f2f2;
  padding: 10px;
  height: 100vh;
  box-shadow: 2px 0 12px 0 rgba(0, 0, 0, 0.1);
}

.main {
  padding: 20px;
  background-color: #fff;
  min-height: 100vh;
}

.box-card {
  width: 100%;
}

.item {
  margin-bottom: 18px;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 100%;
  min-height: 100%;
}
</style>
