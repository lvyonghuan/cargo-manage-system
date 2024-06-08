<script>
import axios from "axios";
import {ElMessageBox} from "element-plus";

export default {
  name: "products",

  data(){
    return{
     "stores":[],
      "selectId":"",
      "products":[],
    }
  },

  methods:{
    getStoresList(){
      axios({
        method:'get',
        url:'/query/stores',
      }).then(res=>{
        if(res.data.state){
          this.stores=res.data.stores
          if (this.stores.length > 0) {
            this.selectId = this.stores[0].storeID;
          }
        }else{
          ElMessageBox.alert(res.data.msg,'获取商店列表失败')
        }
      })
    },

    getProductsList(){
      if(this.selectId===''){
        return
      }

      let formData=new FormData();
      formData.append("storeID",this.selectId)

      axios({
        method:'post',
        url:'/query/products',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.products=res.data.products
        }else{
          ElMessageBox.alert(res.data.msg,'获取商品列表失败')
        }
      })
    },
  },

  mounted() {
    this.getStoresList();
  },

  watch:{
    selectId:{
      immediate:true,
      handler:function(){
        this.getProductsList()
      }
    }
  },
}
</script>

<template>
  <el-select v-model="selectId" placeholder="Select" style="width: 240px">
    <el-option v-for="store in stores" :key="store.storeID" :label="store.storeName" :value="store.storeID"/>
  </el-select>

  <el-table :data="products" stripe style="width: 100%">
    <el-table-column prop="ProductName" label="商品" width="180" />
    <el-table-column prop="ProductType" label="类型" width="180" />
    <el-table-column prop="BrandName" label="品牌" width="180" />
    <el-table-column prop="Price" label="售价" width="180" />
    <el-table-column prop="Inventory" label="库存" />
  </el-table>

</template>

<style scoped>

</style>