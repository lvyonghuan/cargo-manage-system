<script>
import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";

export default {
  data(){
    return{
      'UPCCode':'',
      'ProductName':'',
      'Price':'',
      'ProductBrand':'',
      'ProductType':'',
      'ProductBrandID':'',
      'TypeID':'',

      'BrandName':'',

      'TypeName':'',
      'parentType':'',

      'VendorName':'',
    }
  },

  methods:{
    addProduct(){
      let formData=new FormData();
      formData.append("UPCCode",this.UPCCode);
      formData.append("name",this.ProductName);
      formData.append("price",this.Price);
      formData.append("brand",this.ProductBrandID);
      formData.append("type",this.TypeID);

      axios({
        method:'post',
        url:'/cargo/add/product',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.UPCCode=''
          this.ProductName=''
          this.Price=''
          this.ProductBrand=''
          this.ProductType=''
          this.ProductBrandID=''
          this.TypeID=''
          ElMessage.success('添加成功');
        }else{
          ElMessageBox.alert(res.data.msg,'添加失败')
        }
      })
    },

    addBrand(){
      let formData=new FormData();
      formData.append("name",this.BrandName);

      axios({
        method:'post',
        url:'/cargo/add/brand',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.BrandName=''
          ElMessage.success('添加成功');
        }else{
          ElMessageBox.alert(res.data.msg,'添加失败')
        }
      })
    },

    addType(){
      let formData=new FormData();
      formData.append("name",this.TypeName);
      formData.append("parentTypeID",this.TypeID);

      axios({
        method:'post',
        url:'/cargo/add/type',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.TypeName=''
          this.parentType=''
          ElMessage.success('添加成功');
        }else{
          ElMessageBox.alert(res.data.msg,'添加失败')
        }
      })
    },

    addVendor(){
      let formData=new FormData();
      formData.append("name",this.VendorName);

      axios({
        method:'post',
        url:'/cargo/add/vendor',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.VendorName=''
          ElMessage.success('添加成功');
        }else{
          ElMessageBox.alert(res.data.msg,'添加失败')
        }
      })
    },

    queryBrandID(brandName,inputModel){
      if(brandName===''){
        return
      }
      let formData=new FormData();
      formData.append("brandName",brandName);

      axios({
        method:'post',
        url:'query/brand',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.ProductBrandID=res.data.id
        }else{
          this[inputModel] = ''
          ElMessageBox.alert(res.data.msg,'品牌ID获取失败')
        }
      })
    },

    queryTypeID(typeName,inputModel){
      if(typeName===''){
        return
      }
      let formData=new FormData();
      formData.append("typeName",typeName);

      axios({
        method:'post',
        url:'query/type',
        data:formData,
        headers:{
          'Content-Type': 'multipart/form-data'
        }
      }).then(res=>{
        if(res.data.state){
          this.TypeID=res.data.id
        }else{
          this[inputModel] = ''
          ElMessageBox.alert(res.data.msg,'产品类型ID获取失败')
        }
      })
    }
  }
}
</script>

<template>
  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">新增产品</el-text>
    </div>
    <div>

      <el-form>
        <el-form-item label="UPC码">
          <el-input v-model="UPCCode"></el-input>
        </el-form-item>
        <el-form-item label="产品名称">
          <el-input v-model="ProductName"></el-input>
        </el-form-item>
        <el-form-item label="价格">
          <el-input v-model="Price"></el-input>
        </el-form-item>
        <el-form-item label="品牌">
          <el-input v-model="ProductBrand" @blur="queryBrandID(ProductBrand,'ProductBrand')"></el-input>
        </el-form-item>
        <el-form-item label="类型">
          <el-input v-model="ProductType" @blur="queryTypeID(ProductType,'ProductType')"></el-input>
        </el-form-item>
      </el-form>

      <el-button @click="addProduct">添加</el-button>
    </div>
  </el-card>

  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">新增品牌</el-text>
      <div>
        <el-form>
          <el-form-item label="品牌名称">
            <el-input v-model="BrandName"></el-input>
          </el-form-item>
        </el-form>

        <el-button @click="addBrand">添加</el-button>
      </div>
    </div>
  </el-card>

  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">新增类型</el-text>
      </div>

    <div>
      <el-form>
        <el-form-item label="类型名称">
          <el-input v-model="TypeName"></el-input>
        </el-form-item>
        <el-form-item label="父类型">
          <el-input v-model="parentType" @blur="queryTypeID(parentType,'parentType')"></el-input>
        </el-form-item>
      </el-form>
      <el-button @click="addType">添加</el-button>
    </div>
  </el-card>

  <el-card class="box-card card-margin">
    <div slot="header" class="clearfix">
      <el-text size="large" class="centered-text">新增供应商</el-text>
    </div>

    <div>
      <el-form>
        <el-form-item label="供应商名称">
          <el-input v-model="VendorName"></el-input>
        </el-form-item>
      </el-form>
      <el-button @click="addVendor">添加</el-button>
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