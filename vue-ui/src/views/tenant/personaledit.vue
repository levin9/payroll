<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                            <el-breadcrumb-item> 员工信息</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" >
             </div>                
        </div>
        <div class="container">
              <div class="handle-box" style="margin-top:-20px; margin-bottom:10px;" >
          <el-form ref="form" :model="formModel"  label-width="100px" width=60%>
            <el-tabs type="border-card">
                <el-tab-pane label="用户信息">
                     <el-row>
                        <el-col span="8">
                            <el-form-item label="姓名" required=“true” prop="formModel.real_name" ><el-input v-model="formModel.real_name" ></el-input></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                             <el-form-item label="工号" prop="formModel.job_no"  ><el-input v-model="formModel.job_no" ></el-input></el-form-item>
                        </el-col>
                     </el-row>
                     <el-row>
                        <el-col span="8">
                             <el-form-item label="部门" prop="formModel.dept_name" required=“true”   ><el-input v-model="formModel.dept_name" ></el-input></el-form-item>
                          
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="在职" prop="formModel.status"> 
                                <el-switch
                                    v-model="formModel.status"
                                    active-color="#13ce66"
                                    inactive-color="#ff4949">
                                </el-switch>                            
                            </el-form-item>
                        </el-col>
                     </el-row>         
                     <el-row>
                        <el-col span="8">
                             <el-form-item label="职务" prop="formModel.job_name"><el-input v-model="formModel.job_name" ></el-input></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="职级" prop="formModel.job_level"><el-input v-model="formModel.job_level" ></el-input></el-form-item>
                        </el-col>
                     </el-row>
                     <el-row>
                        <el-col span="8">
                            <el-form-item label="入职日期" prop="formModel.join_date" required=“true”  ><el-date-picker v-model="formModel.join_date" ></el-date-picker></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="转正日期" prop="formModel.regular_date"><el-date-picker v-model="formModel.regular_date" ></el-date-picker></el-form-item>
                        </el-col>
                     </el-row>
                           <el-row>
                        <el-col span="8">
                            <el-form-item label="联系电话" required=“true” prop="formModel.mobile" ><el-input v-model="formModel.mobile" ></el-input></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="电子邮件" prop="formModel.email"><el-input v-model="formModel.email" ></el-input></el-form-item>
                        </el-col>
                     </el-row>
                </el-tab-pane>
                <el-tab-pane label="个税信息">
                      <el-row>
                        <el-col span="8">
                            <el-form-item label="子女教育" prop="formModel.znjy"><inputNumber v-model="formModel.znjy"  append="元" ></inputNumber></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="继续教育" prop="formModel.jxjy"><inputNumber v-model="formModel.jxjy"  append="元" ></inputNumber></el-form-item>
                        </el-col>
                     </el-row>
                         <el-row>
                        <el-col span="8">
                            <el-form-item label="住房贷款" prop="formModel.zfdk"><inputNumber v-model="formModel.zfdk"  append="元" ></inputNumber></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="住房租金" prop="formModel.zfzj"><inputNumber v-model="formModel.zfzj"  append="元" ></inputNumber></el-form-item>
                        </el-col>
                     </el-row>
                               <el-row>
                        <el-col span="8">
                            <el-form-item label="赡养老人" prop="formModel.sylr"><inputNumber v-model="formModel.sylr"  append="元" ></inputNumber></el-form-item>
                        </el-col>
                        <el-col span="2"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="申报日期" prop="formModel.f_create_date"><el-date-picker v-model="formModel.f_create_date" ></el-date-picker></el-form-item>
                        </el-col>
                     </el-row>                 
                </el-tab-pane>
                <el-tab-pane label="银行卡">
                    <el-row>
                        <el-col span="8">
                            <el-form-item label="银 行" prop="formModel.bank_name"><el-input v-model="formModel.bank_name" ></el-input></el-form-item>
                        </el-col>                      
                     </el-row>
                       <el-row>
                        <el-col span="8">
                            <el-form-item label="开户银行" prop="formModel.bank_branch"><el-input v-model="formModel.bank_branch" ></el-input></el-form-item>
                        </el-col>                      
                     </el-row>
                       <el-row>
                        <el-col span="8">
                            <el-form-item label="卡 号" prop="formModel.account_no"><el-input v-model="formModel.account_no" ></el-input></el-form-item>
                        </el-col>                      
                     </el-row>
                       <el-row>
                        <el-col span="8">
                            <el-form-item label="收款人" prop="formModel.account_name"><el-input v-model="formModel.account_name" ></el-input></el-form-item>
                        </el-col>                      
                     </el-row>
                </el-tab-pane>
                <el-tab-pane label="薪资信息">
                        <el-row>
                        <el-col span="9">
                            <el-form-item label="月薪" prop="formModel.monthly_pay_amount" required=“true”  ><inputNumber v-model="formModel.monthly_pay_amount" append="元" ></inputNumber></el-form-item>
                        </el-col>
                 
                        <el-col span="9">
                            <el-form-item label="生效日期" prop="formModel.valid_date"><el-date-picker v-model="formModel.valid_date" ></el-date-picker></el-form-item>
                        </el-col>
                     </el-row>
                    <el-row>
                        <el-col span="9">
                            <el-form-item label="社保基数" prop="formModel.she_bao_fee"><inputNumber v-model="formModel.she_bao_fee" append="元" ></inputNumber></el-form-item>
                        </el-col>                 
                        <el-col span="9">
                            <el-form-item label="公积金" prop="formModel.house_fee" >                                
                                <inputNumber  v-model="formModel.house_fee" append="元" ></inputNumber>
                            </el-form-item>
                        </el-col>
                     </el-row>
                     <el-col :span="item.span?item.span:9" v-for="(item,index) in payrollcfg" :key="index">         
                            <el-form-item  :label="item.label" :prop="item.prop" >                                
                                <inputNumber v-model="formModel[item.prop]" append="元" ></inputNumber>
                            </el-form-item>
                    </el-col>   
                </el-tab-pane>
            </el-tabs>
            <bottomBar @onSave="handleSave"></bottomBar>
           </el-form>
          
        
        </div></div>
    </div>
</template>

<script>
import inputNumber from  '../../components/controls/input-number.vue'
import bottomBar from    '../../components/common/BottomBar.vue'

    export default {
        components: {
            inputNumber,
            bottomBar
        },
        data() {
            return {
                message: 'first',
                formModel:{
                    person_id:'',
                    tenant_id:localStorage.getItem('user_tenant_id'),
                    real_name:'',                  
                    job_no:'',
                    dept_name:'',
                    job_level:'',
                    job_title:'',
                    join_date:'',
                    regular_date:'',
                    status:'',
                    mobile:'',
                    email:'',                                   
            
                    bank_name:'',
                    bank_branch:'',
                    account_no:'',
                    account_name:'' ,
                    znjy:0,
                    jxjy:0,
                    zfdk:0,
                    zfzj:0,
                    sylr:0,  
                    applicateDate:new Date() ,
                    monthly_pay_amount:0,
                    basic_sallary_amount:0,
                    kpi_amout:0,
                    traffic_fee:0,
                    food_fee:0,
                    att_award_fee:0,
                    att_fee:0,
                    professiona_fee:0,
                    job_allowance:0,
                    assign_fee:0,
                    mobile_fee:0,
                    she_bao_fee:0,
                    house_fee:0,
                    cust_add_fee2:0,
                    cust_add_fee1:0,
                    cust_sub_fee2:0,
                    cust_sub_fee1:0,
                    spec_job_allowance:0,
                    she_bao_account_no:"",
                    house_account_no:"",
                    valid_date:new Date()
                },                
                payrollcfg:[
                    { label:'基本工资',prop:'basic_sallary_amount'},
                    { label:'岗位工资',prop:'job_allowance'},
                    { label:'绩效工资',prop:'kpi_amout'},
                    { label:'全勤奖',prop:'att_award_fee'},
                    { label:'交通补贴',prop:'traffic_fee'},
                    { label:'岗位津贴',prop:'cust_add_fee2'},
                    { label:'发票抵扣',prop:'cust_add_fee1'},
                    { label:'租金抵扣',prop:'cust_sub_fee2'},
                ],                
                rules:{
                    real_name:[
                         {required: true, message: '必填', trigger: 'blur'}
                    ],
                     job_no:[
                        {required: true, message: '必填', trigger: 'blur'},
                        {min: 2,max: 20,message: '长度在 2 到 6 个字符' }
                    ]                    
                }
            }
        },
        mounted(){  
            this.$fetch('personal/' + this.$route.query.id )
                .then((response) => {        
                    this.formModel.person_id=response.data.person_id;
                    this.formModel.real_name=response.data.real_name;                 
                    this.formModel.job_no=response.data.job_no;       
                    this.formModel.dept_name=response.data.dept_name;       
                    this.formModel.job_level=response.data.job_level;       
                    this.formModel.job_title=response.data.job_title;       
                    this.formModel.join_date=response.data.join_date;       
                    this.formModel.regular_date=response.data.regular_date;       
                    this.formModel.status=response.data.status;       
                    this.formModel.mobile=response.data.mobile;       
                    this.formModel.email=response.data.email;                                         
                }).catch(function(error){           
                    alert('error..');
                }); 
            this.$fetch('personpayroll/' + this.$route.query.id )
                .then((response) => {
                    if (response.code ==1){
                        this.formModel=response.data;     
                        this.formModel.monthly_pay_amount=response.data.monthly_pay_amount;   
                        this.formModel.basic_sallary_amount=response.data.basic_sallary_amount;   
                        //this.formModel.kpi_amout=response.data.kpi_amout;   
                        this.formModel.traffic_fee=response.data.traffic_fee;   
                        this.formModel.food_fee=response.data.food_fee;   
                        this.formModel.att_award_fee=response.data.att_award_fee;   
                        this.formModel.att_fee=response.data.att_fee;   
                        this.formModel.professiona_fee=response.professiona_fee.email;   
                        this.formModel.job_allowance=response.data.job_allowance;   
                        this.formModel.assign_fee=response.data.assign_fee;   
                        this.formModel.mobile_fee=response.data.mobile_fee;   
                        this.formModel.she_bao_fee=response.data.she_bao_fee;   
                        this.formModel.house_fee=response.data.house_fee;   
                        this.formModel.cust_add_fee2=response.data.cust_add_fee2;   
                        this.formModel.cust_add_fee1=response.data.cust_add_fee1;   
                        this.formModel.cust_sub_fee2=response.data.cust_sub_fee2;   
                        this.formModel.cust_sub_fee1=response.data.cust_sub_fee1;   
                        this.formModel.spec_job_allowance=response.data.spec_job_allowance;   
                        this.formModel.she_bao_account_no=response.she_bao_account_no.email;   
                        this.formModel.house_account_no=response.data.house_account_no;   
                        //this.formModel.valid_date=new Date()
                    }
                }).catch(function(error){           
                    alert('error..');
                }); 
            this.$fetch('singletax/207fa1a9-160c-4943-a89b-8fa4db0547ce/' + this.$route.query.id )
                .then((response) => {   
                    if (response.code==1){
                        this.formModel.znjy=response.data.znjy;  
                        this.formModel.jxjy=response.data.jxjy;  
                        this.formModel.zfdk=response.data.zfdk;  
                        this.formModel.zfzj=response.data.zfzj;  
                        this.formModel.sylr=response.data.sylr;    
                    }                                           
                }).catch(function(error){           
                    alert('error..');
                }); 
            this.$fetch('bankaccount/' + this.$route.query.id )
                .then((response) => {  
                    if (response.code==1){
                        this.formModel.bank_name=response.data.bank_name;  
                        this.formModel.bank_branch=response.data.bank_branch;  
                        this.formModel.account_no=response.data.account_no;  
                        this.formModel.account_name=response.data.account_name;  
                    } 
                }).catch(function(error){           
                    alert('error..');
                }); 
        },
        methods: {            
            handleSave(formModel){
                alert('dddd');
                alert(this.$refs['form']);
                // this.$refs['form'].validate((valid) => {        
                //     alert(valid);
                // });
                console.log(this.formModel);
                console.log('开始post对象');

                this.$post('modifypersonalinfo',this.formModel)
                    .then((response) => {
                        console.log(response);
                        this.$message.success(`修改成功`);
                    }).catch(function(error){
                        console.log('error对象是：')
                        console.log(error);
                        alert('error..');
                    }); 
                    console.log(this.formModel);

                alert('sdfs');
            },
            handleDel(index) {
                const item = this.read.splice(index, 1);
                this.recycle = item.concat(this.recycle);
            }
        }
    }

</script>

<style>
.message-title{
    cursor: pointer;
}
.handle-row{
    margin-top: 30px;
}
</style>

