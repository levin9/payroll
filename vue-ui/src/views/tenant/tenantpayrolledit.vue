

<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 公司设置</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪资项维护</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" >
             </div>                
        </div>
        <div class="container" >
          <el-form ref="form" :rules="rules" :model="formModel" label-width="100px">
            <div style="margin-top:-20px;">
               <el-row>
             
                
                    <el-col :span="8">
                        <el-form-item label="薪资项名称" prop="tenant_payroll_name" required=“true” ><el-input v-model="formModel.tenant_payroll_name" ></el-input>
                    </el-form-item>
                    </el-col>
                    <el-col :span="1"> &nbsp; </el-col>  
                    <el-col :span="8">
                         <el-form-item label="项目类别" prop="item_type" required=“true” >
                             <el-select  placeholder="请选择" v-model="formModel.item_type"  width="100%" >
                                <el-option
                                    v-for="item in allPayrollCategory"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>  
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <el-form-item label="启用" prop="is_used" required=“true” >
                            <el-switch               
                                    :active-value=1   :inactive-value=2
                                    on-color="#5B7BFA"  off-color="#dadde5"
                                    v-model="formModel.is_used"  >
                            </el-switch>
                        </el-form-item>
                    </el-col>
                    <el-col :span="1"> &nbsp; </el-col>    
                    <el-col :span="8">
                        <el-form-item label="序号" prop="f_sort_code" required=“true” ><el-input-number type="number" v-model="formModel.f_sort_code" :min="1" :max="100" ></el-input-number></el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="17" >
                        <el-form-item label="公式说明" prop="tenant_formula_remark" required=“true” >
                            <el-input type="textarea" v-model="formModel.tenant_formula_remark"  ></el-input>
                        </el-form-item>
                    </el-col>                 
                </el-row>
                <el-row>
                    <el-col :span="17">
                        <el-form-item label="公式" prop="tenant_formula" required=“true” >
                            <el-input type="textarea" v-model="formModel.tenant_formula"   :autosize="{ minRows: 2, maxRows: 4}" ></el-input>
                        </el-form-item>
                    </el-col>                 
                </el-row>
                <el-row  style="height:20px; margin:0 0 0 0; " height="20px">
                    <el-col :offset="13" :span="2"  >
                        <el-form-item label="" >
                            <el-button type="info" @click="handerShowVar" >公式变量</el-button>
                        </el-form-item>
                    </el-col>                 
                </el-row>
               
             
                <el-row>
                    <el-col :span="17">
                          <bottomBar @onSave="handleSave"></bottomBar>
                    </el-col>                 
                </el-row>           
            </div>
           </el-form>          
        </div>
        <el-dialog title="变量"  :visible.sync="dialogVarVisible" destroy-on-close="true" :width="width?width:'55%'"> 
            <div style="bgcolor:red; border:2px;margin-top:-40px;">   
                <el-tag closable  v-for="(item,index) in allVarList" :key="index" style="margin-left:10px;margin-bottom:10px;" >  {{item}}  </el-tag>          
                
            </div>
            <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVarVisible = false">取 消</el-button>
                    <el-button type="primary" @click="dialogVarVisible = false">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import bottomBar from    '../../components/common/BottomBar.vue'

  export default {
    components: {        
        bottomBar
    },
    data() {
      return {
        active: 1,
        dialogVarVisible:false,
        allPayrollCategory: 
            [
                { value: '1', label: '黄金糕' }, 
                { value: '2', label: '黄金糕' }
            ],
        allVarList:['基本工资','绩效工资','基本工资金糕','绩工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资','基本工资','绩效工资'],
        formModel:{
            payroll_id:"44444",
            tenant_payroll_name:"",
            item_type:"",
            f_sort_code:0,
            tenant_formula:"",
            tenant_formula_remark:"",
            IsMegerCalc:false
        },
        rules:{
            tenant_payroll_name:[
                {required: true, message: '必填', trigger: 'blur'}    
            ],
             item_type:[
                {required: true, message: '必填', trigger: 'blur'}    
            ],
             f_sort_code:[
                {required: true, message: '必填', trigger: 'blur'}    
            ],
             tenant_formula:[
                {required: true, message: '必填', trigger: 'blur'}    
            ],
            tenant_formula_remark:[
                {required: true, message: '必填', trigger: 'blur'}    
            ]
        }
      };
    },
    mounted(){  
        this.$fetch('tenantpayroll/' + this.$route.query.id )
        .then((response) => {
            this.formModel=response.data;                           
        }).catch(function(error){           
            alert('error..');
        }); 

         this.$fetch('loadvariable/' + this.$route.query.id )
        .then((response) => {
            console.log('dddddddddddddd');
            console.log(response);
            this.allVarList=response.data;                           
        }).catch(function(error){           
            alert('error..');
        }); 
    },
    methods: {
      next() {
        if (this.active++ > 5) this.active = 1;
      },
       previous() {
        if (this.active-- < 1) this.active =1;
      },
       handleSave(formData){
            this.$refs['form'].validate((valid) => {
                if (!valid) {
                    alert('falure');
                    return false
                } else {
                    this.$patch('tenantpayroll',this.formModel)
                    .then((response) => {
                        console.log(response);
                        this.$message.success(`修改成功`);
                    }).catch(function(error){
                        console.log(error);
                        alert('error..');
                    }); 

                    this.$router.push({path:'/companyedit',query: {active_tab_name:"payrollset"}});
                    this.$destroy(true);           
                }
            });
       },
       handerShowVar(){
           this.dialogVarVisible=true;
       }
    }
  }
</script>