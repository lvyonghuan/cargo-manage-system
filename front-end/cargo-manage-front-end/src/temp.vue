<script>
export default {
  name: "temp"
}
</script>

<template>
  <el-card>
    <el-card-header>
      <span>新增货物</span>
    </el-card-header>
    <el-card-body>
      <el-form :model="newGoods">
        <el-form-item label="名称">
          <el-input v-model="newGoods.name"></el-input>
        </el-form-item>
        <el-form-item label="价格">
          <el-input-number v-model="newGoods.price"></el-input-number>
        </el-form-item>
        <el-form-item label="数量">
          <el-input-number v-model="newGoods.quantity"></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addGoods">提交</el-button>
        </el-form-item>
      </el-form>
    </el-card-body>
  </el-card>
</template>

<style scoped>

</style>