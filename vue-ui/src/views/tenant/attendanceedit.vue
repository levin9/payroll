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
          <el-form ref="form" :model="formModel" label-width="260px" >
            <el-row>
                <el-col span="8">
                      <el-form-item label="姓名" prop="real_name"  size="mini" label-width="100px">
                        <el-input   v-model="formModel.real_name"   readonly ></el-input></el-form-item>               
                    </el-form-item>

                </el-col>
                <el-col span="8">
                                  <el-form-item label="姓名" prop="dept_name"  size="mini" label-width="100px">
                        <el-input  v-model="formModel.dept_name"  readonly ></el-input></el-form-item>
                </el-col>
                <el-col span="8">
                <el-form-item label="月份" prop="month_id" size="mini" label-width="100px">
                        <el-input v-model="formModel.month_id"  readonly ></el-input></el-form-item>
                </el-col>                        
            </el-row>
              <el-row>
                <el-col span="8">
                    <el-form-item label="迟到次数" prop="chi_dao_count" required="true" size="mini" label-width="100px">
                                                 <inputNumber  v-model="formModel.chi_dao_count"   append="次" ></inputNumber>
                    </el-form-item>
                </el-col>  
                <el-col span="8">
                    <el-form-item label="迟到时长" prop="chi_dao_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.chi_dao_num"   append="分钟" ></inputNumber>
                    </el-form-item>
                </el-col>    
                     <el-col span="8">
                    <el-form-item label="早退时长" prop="zao_tui_count" required="true" size="mini" label-width="100px">
                                                 <inputNumber  v-model="formModel.zao_tui_count"   append="分钟" ></inputNumber>
                    </el-form-item>
                </el-col> 
                                  
            </el-row>
              
                          <el-row>
                <el-col span="8">
                    <el-form-item label="早退次数" prop="zao_tui_num" required="true" size="mini" label-width="100px">
                                                 <inputNumber  v-model="formModel.zao_tui_num"   append="次" ></inputNumber>
                    </el-form-item>
                </el-col>  
                <el-col span="8">
                    <el-form-item label="调休时长" prop="tiao_xiu_num" required="true" size="mini" label-width="100px">
                        <inputNumber  v-model="formModel.tiao_xiu_num"   append="小时" ></inputNumber>
                    </el-form-item>
                </el-col>    
                     <el-col span="8">
                    <el-form-item label="事假时长" prop="shi_jia_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.shi_jia_num"   append="小时" ></inputNumber>
                    </el-form-item>
                </el-col> 
                                  
            </el-row>
                      <el-row>
                <el-col span="8">
                    <el-form-item label="病假时长" prop="sick_num" required="true" size="mini" label-width="100px">
                        <inputNumber  v-model="formModel.sick_num"   append="小时" ></inputNumber>
                    </el-form-item>
                </el-col>  
                <el-col span="8">
                    <el-form-item label="年假时长" prop="annual_leave_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.annual_leave_num"   append="小时" ></inputNumber>
                    </el-form-item>
                </el-col>    
                     <el-col span="8">
                    <el-form-item label="带薪假" prop="dai_xin_shi_jia_num" required="true" size="mini" label-width="100px">
                        <inputNumber  v-model="formModel.dai_xin_shi_jia_num"   append="小时" ></inputNumber>
                    </el-form-item>
                </el-col> 
                                  
            </el-row>
                      <el-row>
                <el-col span="8">
                    <el-form-item label="陪产假" prop="pei_chan_jia_num" required="true" size="mini" label-width="100px">
                                                 <inputNumber  v-model="formModel.pei_chan_jia_num"   append="天" ></inputNumber>
                    </el-form-item>
                </el-col>  
                <el-col span="8">
                    <el-form-item label="产假" prop="chan_jia_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.chan_jia_num"   append="天" ></inputNumber>
                    </el-form-item>
                </el-col>    
                     <el-col span="8">
                    <el-form-item label="婚假" prop="hun_jia_num" required="true" size="mini" label-width="100px">
                                                 <inputNumber  v-model="formModel.hun_jia_num"   append="天" ></inputNumber>
                    </el-form-item>
                </el-col> 
                                  
            </el-row>
                      <el-row>
                <el-col span="8">
                    <el-form-item label="丧假" prop="sang_jia_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.sang_jia_num"   append="天" ></inputNumber>
                    </el-form-item>
                </el-col>  
                <el-col span="8">
                    <el-form-item label="旷工" prop="kuang_gong_num" required="true" size="mini" label-width="100px">
                                                <inputNumber  v-model="formModel.kuang_gong_num"   append="天" ></inputNumber>
                    </el-form-item>
                </el-col>    
                     <el-col span="8">
                    <el-form-item label="考勤空白" prop="kong_bai_num" required="true" size="mini" label-width="100px">
                         <inputNumber  v-model="formModel.kong_bai_num"   append="次数" ></inputNumber>
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
                formModel:{                    
                    id: "",
                    tenant_id: "207fa1a9-160c-4943-a89b-8fa4db0547ce",
                    person_id: "00d7094d-7d2b-48db-94aa-3c02fc1c0eff",
                    month_id: "201910",
                    real_name: "李志刚",
                    dept_name: "李志刚团队",
                    chi_dao_count: 1,
                    chi_dao_num: 2,
                    zao_tui_count: 3,
                    zao_tui_num: 4,
                    tiao_xiu_num: 3,
                    shi_jia_num: 1,
                    sick_num: 0,
                    bu_rv_jia_num: 0,
                    pei_chan_jia_num: 0,
                    chan_jia_num: 0,
                    hun_jia_num: 0,
                    sang_jia_num: 0,
                    kuang_gong_num: 0,
                    kong_bai_num: 0,
                    dai_xin_shi_jia_num: 0,
                    plan_work_day: 23,
                    actual_work_day: 22,
                    annual_leave_num: 0,
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
                   
                }
            }
        },
        mounted(){  
            this.$fetch('attendance/' + this.$route.query.id)
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
