<template>
  <div class="login_container">
    <div class="login_box">
      <el-text class="centered-text" size="large">货物管理系统管理员登陆</el-text>
      <!-- 登录表单区域 -->
      <el-form label-width="0px" class="login_form">
        <!-- 用户名 -->
        <el-form-item>
          <el-text>用户名</el-text>
          <el-input v-model="username"></el-input>
        </el-form-item>
        <!-- 密码 -->
        <el-form-item>
          <el-text>密码</el-text>
          <el-input v-model="password" show-password/>
        </el-form-item>
        <!-- 按钮区域 -->
        <el-row justify="end">
          <el-form-item class="login_btn">
            <el-button type="primary" @click=login>登录</el-button>
          </el-form-item>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
  import axios from "axios";
  import {ElMessage,ElMessageBox} from "element-plus";
  import {useRouter} from "vue-router";

  export default {
    setup() {
      const router = useRouter()

      router.push('/login')
    },

    data() {
      return {
          username: '',
          password: ''
      }
    },

    methods:{
      login(){
        let formData=new FormData();
        formData.append("username",this.username)
        formData.append("password",this.password)

        // 登录逻辑
        axios({
          method:'post',
          url: 'user/login',
          data: formData,
          headers:{
            'Content-Type': 'multipart/form-data'
          }
        }).then(res=>{
          if(res.data.state){
              ElMessage.success('登陆成功');
              this.$router.replace({
                path:'/dashboard'
              }).catch(err => {
                console.error(err);
                ElMessage.error('路由跳转失败');
              })
          }else{
              ElMessageBox.alert(res.data.msg,'登陆失败')
          }
        })
      },
    }
  }
</script>

<style lang="less" scoped>
.login_container {
  height: 100%;
  width: 100%;
  overflow: hidden;
  position: fixed; /* 使元素固定在视口，即使页面滚动也不会移动 */
  top: 0;
  left: 0;

  background-color: #2b4b6b;
}
.login_box {
  // 宽450像素
  width: 450px;
  // 高300像素
  height: 300px;
  // 背景颜色
  background-color: #fff;
  // 圆角边框3像素
  border-radius: 3px;
  // 绝对定位
  position: absolute;
  // 距离左则50%
  left: 50%;
  // 上面距离50%
  top: 50%;
  // 进行位移，并且在横轴上位移-50%，纵轴也位移-50%
  transform: translate(-50%, -50%);
  .avatar_box {
    // 盒子高度130像素
    height: 130px;
    // 宽度130像素
    width: 130px;
    // 边框线1像素 实线
    border: 1px solid #eee;
    // 圆角
    border-radius: 50%;
    // 内padding
    padding: 10px;
    // 添加阴影 向外扩散10像素 颜色ddd
    box-shadow: 0 0 10px #ddd;
    // 绝对定位
    position: absolute;
    // 距离左则
    left: 50%;
    // 位移
    transform: translate(-50%, -50%);
    // 背景
    background-color: #fff;
    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eee;
    }
  }
}
.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}
.login_btn {
  // 设置弹性布局
  display: flex;
  // 横轴上尾部对齐
  justify-content: flex-end;
}

.centered-text {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 8% 0;
}
</style>