<script>
import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";

export default {
  name: "restock",

  data() {
    return {
      restock: {
        name: '',
        store: '',
        vendor: '',
        quantity: '',
        storeID: '',
        vendorID: '',
        upc: '',
        totalPrice: '',
      },
      autoRestock: {
        name: '',
        store: '',
        quantity: '',
        storeID: '',
        upc: '',
        limit: '',
      },
      deleteAutoRestock: {
        name: '',
        store: '',
        upc: '',
        storeID: '',
      },

      autoRestockList: [],
    }
  },

  methods: {
    restockGoods() {
      let formData = new FormData();
      formData.append("UPCCode", this.restock.upc);
      formData.append("vendorID", this.restock.vendorID);
      formData.append("storeID", this.restock.storeID);
      formData.append("purchaseAmount", this.restock.quantity);
      formData.append("totalPrice", this.restock.totalPrice);

      axios({
        method: 'post',
        url: '/cargo/restock',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          this.restock.name = ''
          this.restock.store = ''
          this.restock.quantity = ''
          this.restock.storeID = ''
          this.restock.upc = ''
          this.restock.vendor = ''
          this.restock.vendorID = ''
          this.restock.totalPrice = ''
          ElMessage.success('添加成功');
        } else {
          ElMessageBox.alert(res.data.msg, '添加失败')
        }
      })
    },

    autoRestockGoods() {
      let formData = new FormData();
      formData.append("UPCCode", this.autoRestock.upc);
      formData.append("storeID", this.autoRestock.storeID);
      formData.append("restockAmount", this.autoRestock.quantity);
      formData.append("limit", this.autoRestock.limit);

      axios({
        method: 'post',
        url: '/cargo/ordering/add',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          this.autoRestock.name = ''
          this.autoRestock.store = ''
          this.autoRestock.quantity = ''
          this.autoRestock.storeID = ''
          this.autoRestock.upc = ''
          this.autoRestock.limit = ''
          ElMessage.success('添加成功');
        } else {
          ElMessageBox.alert(res.data.msg, '添加失败')
        }
      })
    },

    deleteAutoRestockGoods() {
      let formData = new FormData();
      formData.append("UPCCode", this.deleteAutoRestock.upc);
      formData.append("storeID", this.deleteAutoRestock.storeID);

      axios({
        method: 'post',
        url: '/cargo/ordering/delete',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          this.deleteAutoRestock.name = ''
          this.deleteAutoRestock.store = ''
          this.deleteAutoRestock.upc = ''
          this.deleteAutoRestock.storeID = ''
          ElMessage.success('删除成功');
        } else {
          ElMessageBox.alert(res.data.msg, '删除失败')
        }
      })
    },

    queryStore(storeName, inputModel) {
      if (storeName === '') {
        return
      }

      let formData = new FormData();
      formData.append("storeName", storeName);

      axios({
        method: 'post',
        url: '/query/store',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          inputModel.storeID = res.data.id
        } else {
          inputModel.store = ''
          ElMessageBox.alert(res.data.msg, '查询失败')
        }
      })
    },

    queryUPC(productName, inputModel) {
      if (productName === '') {
        return
      }

      let formData = new FormData();
      formData.append("productName", productName);

      axios({
        method: 'post',
        url: '/query/upc',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          inputModel.upc = res.data.upc
        } else {
          inputModel.name = ''
          ElMessageBox.alert(res.data.msg, '查询失败')
        }
      })
    },

    queryVendor(vendorName, inputModel) {
      if (vendorName === '') {
        return
      }

      let formData = new FormData();
      formData.append("vendorName", vendorName);

      axios({
        method: 'post',
        url: '/query/vendor',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if (res.data.state) {
          inputModel.vendorID = res.data.id
        } else {
          inputModel.vendor = ''
          ElMessageBox.alert(res.data.msg, '查询失败')
        }
      })
    },
  },
}
</script>

<template>
  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">进货</el-text>
    </div>
    <div>
      <el-form :model="restock">
        <el-form-item label="货物名称">
          <el-input v-model="restock.name" @blur="queryUPC(restock.name,restock)"></el-input>
        </el-form-item>
        <el-form-item label="商店名称">
          <el-input v-model="restock.store" @blur="queryStore(restock.store,restock)"></el-input>
        </el-form-item>
        <el-form-item label="供应商">
          <el-input v-model="restock.vendor" @blur="queryVendor(restock.vendor,restock)"></el-input>
        </el-form-item>
        <el-form-item label="数量">
          <el-input-number v-model="restock.quantity"></el-input-number>
        </el-form-item>
        <el-form-item label="总价">
          <el-input-number v-model="restock.totalPrice" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="restockGoods">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>

  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">自动进货计划</el-text>
    </div>
    <div>
      <el-form :model="autoRestock">
        <el-form-item label="货物名称">
          <el-input v-model="autoRestock.name" @blur="queryUPC(autoRestock.name,autoRestock)"></el-input>
        </el-form-item>
        <el-form-item label="商店名称">
          <el-input v-model="autoRestock.store" @blur="queryStore(autoRestock.store,autoRestock)"></el-input>
        </el-form-item>
        <el-form-item label="进货数量">
          <el-input-number v-model="autoRestock.quantity"></el-input-number>
        </el-form-item>
        <el-form-item label="库存下限">
          <el-input-number v-model="autoRestock.limit"></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="autoRestockGoods">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>

  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">删除自动进货计划</el-text>
    </div>
    <div>
      <el-form :model="deleteAutoRestock">
        <el-form-item label="货物名称">
          <el-input v-model="deleteAutoRestock.name" @blur="queryUPC(deleteAutoRestock.name,deleteAutoRestock)"></el-input>
        </el-form-item>
        <el-form-item label="商店名称">
          <el-input v-model="deleteAutoRestock.store" @blur="queryStore(deleteAutoRestock.store,deleteAutoRestock)"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="deleteAutoRestockGoods">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<style scoped>
.centered-text {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2% 0;
}

.card-margin {
  margin: 20px; /* 设置上下左右的间距为20px */
}

</style>