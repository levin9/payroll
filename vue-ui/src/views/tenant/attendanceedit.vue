<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                            <el-breadcrumb-item> 编辑考勤</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" >
             </div>                
        </div>
        <div class="container">
          <el-form ref="form" :model="formModel" label-width="100px" >
            <el-tab-pane label="基本信息" name="baseinfo">
                     <el-row>
                        <el-col span="8">
                            <el-form-item label="公司名称" prop="full_name" required="true">
                                <el-input v-model="formModel.full_name" ></el-input>
                            </el-form-item>
                            <input type="hidden" id="tenant_id" v-bind:value="formModel.tenant_id" > </input>

                        </el-col>
                        <el-col span="1"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="公司简称" prop="simple_name" required="true" ><el-input v-model="formModel.simple_name" ></el-input></el-form-item>
                        </el-col>
                     </el-row>
                     <el-row>
                        <el-col span="8">
                            <el-form-item label="联系人"  prop="contact_name" required="true"  ><el-input v-model="formModel.contact_name" ></el-input></el-form-item>
                        </el-col>
                        <el-col span="1"> &nbsp; </el-col>    
                        <el-col span="12">
                            <el-form-item label="联系人称呼" prop="contact_title" required="true"  >
                                <el-radio-group v-model="formModel.contact_title">                                       
                                        <el-radio v-for="title in allcontact_titles"  :label="title" :key="title">{{title}}</el-radio>
                                    </el-radio-group>
                            </el-form-item>
                        </el-col>
                     </el-row>
                    <el-row>
                        <el-col span="8">
                            <el-form-item label="联系电话" prop="contact_phone" ><el-input v-model="formModel.contact_phone" ></el-input></el-form-item>
                        </el-col>
                        <el-col span="1"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="手机" prop="contact_mobile"><el-input v-model="formModel.contact_mobile" ></el-input></el-form-item>
                        </el-col>
                     </el-row>
                     <el-row>
                        <el-col span="17">
                            <el-form-item label="联系人邮件" prop="contact_email"   >
                                <el-input v-model="formModel.contact_email" ></el-input>
                            </el-form-item>
                        </el-col>
                     
                     </el-row>
                     <el-row>
                        <el-col span="8">
                            <el-form-item label="开始日期" prop="tenant_start_time">
                                <el-date-picker v-model="formModel.tenant_start_time"
                                    dateType="date"
                                    placeholder="选择日期"
                                    format="yyyy 年 MM 月 dd 日"
                                    value-format="2019-12-26T09:46:31 08:00"
                                ></el-date-picker></el-form-item>
                        </el-col>
                        <el-col span="1"> &nbsp; </el-col>    
                        <el-col span="8">
                            <el-form-item label="结束日期" prop="tenant_end_time">
                                <el-date-picker v-model="formModel.tenant_end_time"
                                    dateType="date"
                                    placeholder="选择日期"
                                    format="yyyy 年 MM 月 dd 日"
                                    value-format="2019-12-26T09:46:31 08:00" >
                                </el-date-picker></el-form-item>
                        </el-col>
                     </el-row>                   
                </el-tab-pane>
                <el-tab-pane label="系统设置" name="systemset">               
                     <el-row>
                        <el-col span="18">
                            <el-form-item label="考勤来源" prop="hr_source"  label-width="100px">
                                <el-radio-group v-model="formModel.hr_source">
                                        <el-radio :label="Excel">Excel导入</el-radio>
                                        <el-radio :label="DingTalk">从钉钉抓取</el-radio>
                                    </el-radio-group>
                            </el-form-item>
                        </el-col>                       
                     </el-row>
                     <el-row>
                        <el-col span="18">
                            <el-form-item label="人事来源" prop="att_source"  label-width="100px">
                                  <el-radio-group v-model="formModel.att_source">
                                        <el-radio :label="Excel">Excel导入</el-radio>
                                        <el-radio :label="DingTalk">从钉钉抓取</el-radio>                                        
                                    </el-radio-group>
                            </el-form-item>
                        </el-col>                       
                     </el-row>
                   <el-row>
                        <el-col span="18">
                            <el-form-item label="社保设置" prop="buy_she_bao"  >
                                    <el-radio-group v-model="formModel.buy_she_bao">
                                        <el-radio :label="normal-buy">15日前入职即购买</el-radio>
                                        <el-radio :label="regular-buy">转正月才购买</el-radio>   
                                        <el-radio :label="none-buy">不购买社保</el-radio>                                          
                                    </el-radio-group>
                            </el-form-item>
                        </el-col>                       
                     </el-row>
                            <el-row>
                        <el-col span="10">
                            <el-form-item label="薪资周期" prop="payroll_day" required="true" size="mini" label-width="100px">
                             <el-input-number v-model="formModel.payroll_day"   :min="1" :max="31" label="请选择每月最后一天"></el-input-number>
                            </el-form-item>
                        </el-col>      
                     </el-row>
                     <el-row>
                        <el-col span="10">
                            <el-form-item label="公积金比例" prop="house_rate" required="true" >                            
                                <el-input-number v-model="formModel.house_rate"  :precision="2" :step="1"  :min="1" :max="20" ></el-input-number>
                            </el-form-item>
                        </el-col>                       
                     </el-row>
                     <el-row>
                        <el-col span="10">
                            <el-form-item label="社保比例" prop="she_bao_rate" required="true" >
                                <el-input-number v-model="formModel.she_bao_rate" :precision="2" :step="1"  :min="1" :max="20" ></el-input-number>
                            </el-form-item>
                        </el-col>                       
                     </el-row>  
       
       
       <bottomBar @onSave="handleSave"></bottomBar>
          
        </el-form>
          
        </div>
    </div>
</template>

<script>
import inputNumber from  '../../components/controls/input-number.vue'
import bottomBar from    '../../components/common/BottomBar.vue'
//import ruleUtil from '../../utils/rules';

    export default {
        components: {
            inputNumber,
            bottomBar
        },
        data() {
            var phoneReg = /^[1][3,4,5,7,8][0-9]{9}$/
            var validatePhone = (rule, value, callback) => {
            if (!value) {
                return callback(new Error('号码不能为空!!'))
            }
            setTimeout(() => {
                if (!phoneReg.test(value)) {
                callback(new Error('格式有误'))
                } else {
                callback()
                }
            }, 1000)
            }          
            return {
                message: 'first',
                showHeader: false,
                active_tab_name:"baseinfo",
                allcontact_titles: ['先生','女士','老板'],
                formModel:{                    
                    tenant_id:"112",
                    full_name:"",
                    simple_name:"",
                    contact_name:"",
                    contact_title:"",
                    contact_phone:"",
                    contact_mobile:"",
                    contact_email:"",
                    tenant_start_time:new Date(),
                    tenant_end_time:new Date("9999-12-31"),
                    payroll_day:"31",                    
                    hr_source:"",
                    att_source:"",
                    buy_she_bao:"normal",
                    she_bao_rate:8,
                    house_rate:5
                },
                query:{
                    tenant_id:localStorage.getItem('user_tenant_id'),
                    is_valid:1,
                    order:" IsUsed desc,f_sortcode asc ",
                    limit:30
                },
                payrollData:[ ],
                attData:[],              
                rules:{
                    full_name:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                     simple_name:[
                        {required: true, message: '必填', trigger: 'blur'}   
                    ],  
                    contact_name:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    contact_title:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    contact_mobile:[
                        {required: true, message: '必填', trigger: 'blur'}
                    ],           
                    tenant_start_time:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    tenant_end_time:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    DayIndex:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    hr_source:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    att_source:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    buy_she_bao:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    she_bao_rate:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    house_rate:[
                        {required: true, message: '必填', trigger: 'blur'}    
                    ],
                    contact_email:[
                         { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
                    ]
                }
            }
        },
        mounted(){  
            this.$fetch('tenant/' + localStorage.getItem('user_tenant_id'))
                    .then((response) => {
                        console.log(response);
                        this.formModel=response.data;                        
                    }).catch(function(error){
                        console.log(error);
                        //alert('error..');
                    }); 
                  
        },
        methods: {
            handleRead(index) {
                const item = this.unread.splice(index, 1);
                console.log(item);
                this.read = item.concat(this.read);
            },
            backHistory(){
                this.$router.go(-1);
            },
            handleSave(formModel){
                this.$refs['form'].validate((valid) => {
                if (!valid) {
                    return false
                } else {
                    this.$patch('tenant',this.formModel)
                    .then((response) => {
                        console.log(response);
                        this.$message.success(`修改成功`);
                    }).catch(function(error){
                        console.log(error);
                        alert('error..');
                    }); 
                    console.log(this.formData);

                    //this.dialogUpdateVisible = false;
                    //this.$message.success(`修改成功`);
                    //this.$set(this.tableData, this.idx, this.form);
                }
                })


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
