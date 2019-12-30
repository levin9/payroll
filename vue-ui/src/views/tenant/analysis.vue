

<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                            <el-breadcrumb-item> 手工调整</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" >
             </div>                
        </div>
        <div class="container">
          <el-form ref="form" :model="form" label-width="70px">
              <el-steps :active="active" finish-status="success" simple style="margin-top: -15px">
                <el-step title="初始化" ></el-step>
                <el-step title="人事信息" ></el-step>
                <el-step title="手工调整" ></el-step>
                <el-step title="考勤信息" ></el-step>
                <el-step title="薪酬分析" ></el-step>
                <el-step title="查看报表" ></el-step>
                </el-steps>
                <el-row>
                        <el-button type="button" icon="el-icon-plus" class="handle-del mr10" @click="previous"  >上一步</el-button>
                        <el-button type="primary" icon="el-icon-plus" class="handle-del mr10" @click="next"  >下一步</el-button>          
                </el-row>
           </el-form>          
        </div>
    </div>
</template>

<script>
  export default {
    data() {
      return {
        active: 1
      };
    },

    methods: {
      next() {
        if (this.active++ > 5) this.active = 1;
      },
       previous() {
        if (this.active-- < 1) this.active =1;
      }
    }
  }
</script>