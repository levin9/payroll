<template>
       <el-row type="flex" class="row-bg" justify="end" style="margin-top:15px;" >
           <el-col>
                <el-button type="button" icon="el-icon-back" @click="backHistory" >返回</el-button>
                
           </el-col>
           <el-col :span="4">
                <el-button type="primary" icon="el-icon-plus" @click="handleSave()" >保存</el-button>            
           </el-col>
       </el-row> 
</template>


<script>
  
  export default {
    name: "botton-bar",   
    props: [       
      'formModel',
      'onSave'     
    ],
    data() {
      return {        
        // 表单数据
       
      }
    },
    mounted() {
      
    },
    methods: {
      backHistory(){
        this.$destroy(true);
        this.$router.go(-1);
         
      },
      handleSave(){
          this.$emit("onSave");
      }
    },
    watch: {
      
    }
  }
</script>



  