<template>
    <div class="">
        <div class="crumbs">
            <div style="width:50%; bgcolor:red;" >
                <el-breadcrumb separator="/">
                            <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                            <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                            <el-breadcrumb-item> 月份设置</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div style="width:50%;bgcolor:blue;" > </div>                
        </div>
        <div class="container">
            <div class="handle-box" style="text-align:center; margin-top:-20px; margin-bottom:40px;" >
                 <el-col  :span="6">
                 &nbsp;
                 </el-col>
                  <el-col  :span="16">
                      &nbsp;
                 </el-col>
                  <el-col  :span="2">
                        <el-button
                            type="primary"
                            icon="el-icon-plus"
                            class="handle-del mr10"
                            @click="handleInsert"
                        >新增</el-button>
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
                <el-table-column prop="month_id" label="月份"></el-table-column>
                 <el-table-column prop="month_name" label="名称"></el-table-column>
                <el-table-column prop="start_date" label="开始时间" :formatter="dateFormat"></el-table-column>
                <el-table-column prop="end_date" label="结束时间" :formatter="formatDatetime"></el-table-column>
                <el-table-column prop="workday_qty" label="工作天数"></el-table-column>
                <el-table-column prop="state_name" label="状态"></el-table-column>
                          
                <el-table-column label="操作"  align="center">
                    <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(scope.$index, scope.row)"
                        >编辑</el-button>
                        <el-button
                            type="text"
                            icon="el-icon-delete"
                            class="red"
                            @click="handleDelete(scope.$index, scope.row)"
                        >删除</el-button>
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
          <el-dialog title="编辑-手工调整"  :visible.sync="dialogUpdateVisible" :width="width?width:'40%'">
            <el-form :model="entity" ref="editForm" :rules="rules" label-width="100px" >
                <el-form-item label="月份" width="250" prop="month_id" ><el-input v-model="entity.month_id"   ></el-input></el-form-item> 
                <el-form-item label="名称" prop="month_name" ><el-input v-model="entity.month_name" ></el-input></el-form-item> 
                <el-form-item label="开始日期" prop="start_date" > <el-date-picker
                        v-model="entity.start_date" 
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker></el-form-item> 
                <el-form-item label="结束日期" prop="end_date" >
                   
                     <el-date-picker
                        v-model="entity.end_date" 
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker>
                </el-form-item> 
                <el-form-item label="调整金额" prop="workday_qty" >
                    <inputNumber v-model="entity.workday_qty"    append="天" ></inputNumber>
                </el-form-item> 
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogUpdateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitEdit">确 定</el-button>
        </span>
        </el-dialog>
        <el-dialog title="新增-手工调整"  :visible.sync="dialogCreateVisible" :width="width?width:'40%'">
            <el-form :model="entity" ref="createForm" :rules="rules" label-width="120px" >          
                <el-form-item label="月份" width="250" prop="month_id" ><el-input v-model="entity.month_id"   ></el-input></el-form-item> 
                <el-form-item label="名称" prop="month_name" ><el-input v-model="entity.month_name" ></el-input></el-form-item> 
                <el-form-item label="开始日期" prop="start_date" > <el-date-picker
                        v-model="entity.start_date" 
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker></el-form-item> 
                <el-form-item label="结束日期" prop="end_date" >
                   
                     <el-date-picker
                        v-model="entity.end_date" 
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker>
                </el-form-item> 
                
                <el-form-item label="工作天数" prop="workday_qty" >
                    <inputNumber v-model="entity.workday_qty"    append="天" ></inputNumber>
                </el-form-item> 
              
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogCreateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitCreate">确 定</el-button>
        </span>
        </el-dialog>
    </div>    
</template>

<script>

import inputNumber from  '../../components/controls/input-number.vue'
import { uuid } from 'vue-uuid' 

export default {
    name: 'basemonth',
      components: {
            inputNumber
        },
    data() {
        return {            
            query: {
                month_id:'201911',
                tenant_id:localStorage.getItem('user_tenant_id'),
                keyword: '',
                order:'',          
                pageSize:6,
                pageIndex:1              
            },
            entity:{
                id:'',
                month_id:'201911',
                tenant_id:'',
                first_code:'',
                second_code:'',
                start_date:'',
                end_date:'',
                she_bao_due_date:'',
                state:'',
                state_name:''
            },
            width:0,
            tableData: [],
            multipleSelection: [],
            delList: [],
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1,
            dialogCreateVisible:false,
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
            this.$fetch('basemonth',this.query )
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
            this.entity.month_name='';
            this.entity.start_date=null;           
            this.entity.end_date=null;
            this.entity.workday_qty=23;
           
            this.dialogCreateVisible=true;
            this.dialogUpdateVisible=false;

        },
         // 修改行操作
        handleEdit(index, row) {        
            this.dialogUpdateVisible=true;
            this.entity=row;                
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
        submitCreate(){
            this.$refs['createForm'].validate((valid) => {
            if (!valid) {
                return false
            } 
            this.$post('tenantadjustitem',this.entity)
                .then((response) => {                        
                    this.dialogCreateVisible = false;
                    this.$message.success(`新增成功`);  
                }).catch(function(error){
                    this.$message.error(error);                         
                }); 
        })

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
        },
        dateFormat:function(row,column){
        var t=new Date(row.start_date);//row 表示一行数据, updateTime 表示要格式化的字段名称   
        return t.getFullYear()+"-"+(t.getMonth()+1)+"-"+t.getDate();
        }
        ,
        formatDatetime:function( row,column){    
              //alert(column);  
             // alert(row[column])
            
          
        var t=new Date(row.end_date);//row 表示一行数据, updateTime 表示要格式化的字段名称     
        return t.getFullYear()+"-"+(t.getMonth()+1)+"-"+t.getDate();
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
