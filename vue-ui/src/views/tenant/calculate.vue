<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬分析</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" > </div>                
        </div>
        <div class="container">
            <div class="handle-box" style="text-align:center; margin-top:-20px; margin-bottom:40px;" >
                 <el-col  :span="5">
                        <el-date-picker
                            v-model="query.month_id"
                            type="month"
                            placeholder="选择月"
                            @change="getData" >
                            
                            </el-date-picker>
                            <el-button slot="append" icon="el-icon-search" @click="handleSearch" ></el-button>
                    
                 </el-col>
                <el-col  :span="5">
                    <el-input placeholder="请输入姓名" v-model="query.keyword"  class="input-with-select">                 
                    <el-button slot="append" icon="el-icon-search" @click="handleSearch" ></el-button>
                    </el-input>  
                 </el-col>
                  <el-col  :span="8">
                      &nbsp;
                 </el-col>
                  <!-- <el-col  :span="2">
                        <el-button
                            type="primary"
                            icon="el-icon-search"
                            class="handle-del mr10"
                            @click="handleInsert"
                        >查询</el-button>
                 </el-col>    -->
                      <el-col  :span="3">
                        <el-button
                            type="warning"
                            icon="el-icon-caret-right"
                            class="handle-del mr10"
                            @click="handleInsert"
                        >薪酬分析</el-button>
                 </el-col> 
                      <el-col  :span="3">
                        <el-button
                            type="primary"
                            icon="el-icon-download"
                            class="handle-del mr10"
                            @click="handleInsert"
                        >导出银行格式</el-button>
                 </el-col> 
                     
            </div>
               <el-table
                :data="tableData"
                border
                class="table"
                ref="multipleTable"
                header-cell-class-name="table-header"
                @selection-change="handleSelectionChange"
            >
                <el-table-column prop="dept_name" label="部门" fixed width="120" ></el-table-column>
                <el-table-column prop="real_name" label="姓名" fixed ></el-table-column>
                <el-table-column prop="chng_remark" label="异动情况" width="120" ></el-table-column>
                <el-table-column prop="actual_work_day" label="工作天数" v-show:="false"></el-table-column>    
                 <template v-for='(col) in tableHeaders'>
                    <el-table-column                    
                        :show-overflow-tooltip="true"
                        :prop="col.dataItem"
                        :label="col.dataName"
                        :key="col.dataItem"
                    >
                    </el-table-column>
                    </template>
                <el-table-column prop="wages_payable_amount" label="应发金额"></el-table-column>
                <el-table-column label="代扣" align='center' >
                    <el-table-column prop="she_bao_fee" label="社保"></el-table-column>
                    <el-table-column prop="house_fee" label="公积金"></el-table-column>
                    <el-table-column prop="income_tax_amount" label="个税"></el-table-column>
                </el-table-column>
                <el-table-column prop="actual_payable_amount" label="实发金额"></el-table-column>
                
               
                <el-table-column label="详情"  align="center">
                    <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                             @click="handleEdit(scope.$index, scope.row)"
                        >详情</el-button>                       
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination
                    background
                    layout="total, prev, pager, next"
                    :current-page="query.pageIndex"
                    :page-size="query.pageSize"
                    :total="pageTotal"
                    @current-change="handlePageChange"
                ></el-pagination>
            </div>               
        </div>
        <div v-if="dialogUpdateVisible">
          <el-dialog title="详情-薪酬详细"  :visible.sync="dialogUpdateVisible" :append-to-body='true'
        :lock-scroll="false" :width="width?width:'60%'">
            <el-form :model="entity" ref="editForm"  label-width="100px" >
               <div>
               <el-table :data="detailDataList"> 
                    <el-table-column prop="fName" label="公式名称" fixed width="120" ></el-table-column>
                    <el-table-column prop="fValue" label="备注"  ></el-table-column>
               
               </el-table>

               </div>
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogUpdateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitEdit">确 定</el-button>
        </span>
        </el-dialog>
        </div>
        <el-dialog title="薪酬计算"  :visible.sync="dialogCalculateVisible" :width="width?width:'40%'">
            <el-form :model="entity" ref="createForm" :rules="rules" label-width="120px" >          
                <el-form-item label="月份" prop="month_id" >  
                       <el-date-picker
                            v-model="query.month_id"
                            type="month"
                            placeholder="选择月"
                            >                            
                            </el-date-picker>
                </el-form-item>              
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogCalculateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitCalculate">确 定</el-button>
        </span>
        </el-dialog>
    </div>    
</template>

<script>
import { fetchData } from '../../api/adjuest';
import inputNumber from  '../../components/controls/input-number.vue'
import { uuid } from 'vue-uuid' 

export default {
    name: 'adjuestlist',
      components: {
            inputNumber
        },
    data() {
        return {            
            query: {
                month_id:new Date(),
                tenant_id:localStorage.getItem('user_tenant_id'),
                keyword: '',
                order:'',          
                pageSize:6,
                pageIndex:1              
            },
            tableHeaders:[  
                            {dataItem: 'monthly_pay_amount',dataName: '月薪'},
                            {dataItem: 'att_fee',dataName: '考勤扣款'}, 
                            {dataItem: 'att_award_fee',dataName: '全勤奖'},                            
                            {dataItem: 'traffic_fee',dataName: '交通'}, 
                            {dataItem: 'food_fee',dataName: '餐费'}, 
                            {dataItem: 'mobile_fee',dataName: '通讯'} ],
            entity:{
                month_id:new Date(),
                tenant_id:'',
                person_name:'',
                person_id:'00d7094d-7d2b-48db-94aa-3c02fc1c0eff',
                amount:0,
                adjest_remark:''
            },
            calcParameter:{
                month_id:'',
                tenant_id:''
            },
            tableData: [],
            detailDataList:[],
            multipleSelection: [],
            delList: [],
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
            dialogCalculateVisible:false,
            dialogUpdateVisible:false,
        rules: 
            {
            amount: [
                        {required: true, message: '必填', trigger: 'blur'}
                    ],
            adjest_remark:   
                    [
                        {required: true, message: '必填', trigger: 'blur'},
                        {min: 2,max: 20,message: '长度在 2 到 6 个字符' }
                    ]
           }
        };
    },
    created() {
        //this.getData();
    },
    mounted(){  
        this.getData();         
    },
    methods: {
        // 获取 easy-mock 的模拟数据
        getData() {           
            if (this.query.month_id ==''){
                this.$message.error("请输入月份年"); 
                return ;
            }
            this.$fetch('monthpayroll',this.query )
            .then((response) => {
                console.log(response);
                this.tableData = response.data;
                this.pageTotal = response.total ;                    
            }).catch(function(error){
                console.log(error);
                this.$message.error(error);                
            }); 
        },
        // 触发搜索按钮
        handleSearch() {
            this.$set(this.query, 'pageIndex', 1);           
            this.getData();
        },
        handleInsert() {     
            this.entity.Id=uuid.v1();
            this.entity.tenant_id=localStorage.getItem('user_tenant_id');  
            this.entity.month_id="201911";
            this.entity.person_id='00d7094d-7d2b-48db-94aa-3c02fc1c0eff';
            this.entity.person_name="无名";
            this.entity.amount=0;
            this.entity.adjest_remark="";
           
            this.dialogCalculateVisible=true;
            this.dialogUpdateVisible=false;

        },
         // 修改行操作
        handleEdit(index, row) {        
            //this.dialogUpdateVisible=true;
            this.entity=row;      
            this.detailDataList=null;
            if (row.ext_info !=''|| row.ext_info!='[]'){                
                var remark =row.ext_info.substring(1,row.ext_info.length-2);
                alert(remark);
                // var ss = remark.split(",");
                // alert(ss.length);
                // for(var i=0;i<ss.length;i++){
                //     var s =ss[i].split(":");
                //    this.detailDataList.push({fName:"a",fValue:i});
                //     if (s.length==2){
                        
                //     }
                // }
            }
        
            //this.detailDataList=row.ext_info;
            
            //alert(detailDataList);    
        },
           // 保存编辑
        submitEdit(){  
            this.$refs['editForm'].validate((valid) => {        
                if (!valid) {
                    return false
                } 
                this.$patch('tenantadjustitem',this.entity)
                    .then((response) => {
                        this.dialogUpdateVisible = false;
                        this.$message.success(`修改成功`);
                        this.$set(this.tableData, this.idx, this.entity);
                        //todo 没有及时刷新
                    }).catch(function(error){
                        console.log(error);
                        alert('error..');
                    }); 
            })
        },
           // 新增
        submitCalculate(){ 
            this.calcParameter.tenant_id=localStorage.getItem('user_tenant_id');  
            this.calcParameter.month_id=this.query.month_id;
            if (this.calcParameter.month_id==""){
                this.$message.success(`请选择需要计算的月份.`);  
                return false;
            }            
            this.$post('calculatepayroll',this.calcParameter)
                .then((response) => {                        
                    this.dialogCalculateVisible = false;
                    this.$message.success(`计算成功`);  
                }).catch(function(error){
                    this.$message.error(error);                         
                });        

        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            })
                .then(() => {
                    this.$del('tenantadjustitem/'+row.Id)
                    .then((response) => {
                        if (response.code==1){
                            this.$message.success('删除成功');
                            this.tableData.splice(index, 1);
                        }else{
                            this.$message.error('删除失败');
                        }                        
                        //todo 没有及时刷新
                    }).catch(function(error){
                        console.log(error);
                        alert('error..');
                    }); 
                    
                })
                .catch(() => {});
        },
        // 多选操作
        handleSelectionChange(val) {
            this.multipleSelection = val;
        },
        delAllSelection() {
            const length = this.multipleSelection.length;
            let str = '';
            this.delList = this.delList.concat(this.multipleSelection);
            for (let i = 0; i < length; i++) {
                str += this.multipleSelection[i].name + ' ';
            }
            this.$message.error(`删除了${str}`);
            this.multipleSelection = [];
        },       
        // 分页导航
        handlePageChange(val) {
            //alert(val);
            this.query.pageIndex=val;           
            this.getData();
        }
    }
};
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}
.table {
    width: 100%;
    font-size: 14px;
}
.red {
    color: #ff0000;
}
.mr10 {
    margin-right: 10px;
}
.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
</style>
