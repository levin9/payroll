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
            <div style="width:50%;bgcolor:blue;" > </div>                
        </div>
        <div class="container">
            <div class="handle-box" style="text-align:center; margin-top:-20px; margin-bottom:40px;" >
                 <el-col  :span="6">
                    <el-input placeholder="请输入姓名" v-model="query.keyword"  class="input-with-select">                 
                    <el-button slot="append" icon="el-icon-search" @click="handleSearch" ></el-button>
                    </el-input>  
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
                <el-table-column prop="dept_name" label="部门"></el-table-column>
                <el-table-column prop="person_name" label="姓名"></el-table-column>
                <el-table-column prop="month_id" label="月份"></el-table-column>
                <el-table-column prop="amount" label="调整金额"></el-table-column>
                <el-table-column prop="adjest_remark" label="调整原因"></el-table-column>
                          
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
                <el-form-item label="姓名" prop="person_name"><el-input v-model="entity.person_name" :disabled=dialogUpdateVisible ></el-input></el-form-item> 
                <el-form-item label="月份" prop="month_id" ><el-input v-model="entity.month_id" :disabled=dialogUpdateVisible ></el-input></el-form-item> 
                <el-form-item label="调整金额" prop="amount">
                    <inputNumber v-model="entity.amount"  append="元" ></inputNumber>
                    </el-form-item> 
                <el-form-item label="调整原因" prop="adjest_remark" ><el-input v-model="entity.adjest_remark" ></el-input></el-form-item>               
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogUpdateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitEdit">确 定</el-button>
        </span>
        </el-dialog>
        <el-dialog title="新增-手工调整"  :visible.sync="dialogCreateVisible" :width="width?width:'40%'">
            <el-form :model="entity" ref="createForm" :rules="rules" label-width="120px" >          
                <el-form-item label="姓名" width="250" prop="person_name" ><el-input v-model="entity.person_name"   ></el-input></el-form-item> 
                <el-form-item label="月份" prop="month_id" ><el-input v-model="entity.month_id" ></el-input></el-form-item> 
                
                <el-form-item label="调整金额" prop="amount" >
                    <inputNumber v-model="entity.amount"    append="元" ></inputNumber>
                    </el-form-item> 
                <el-form-item label="调整原因" prop="adjest_remark"><el-input v-model="entity.adjest_remark" ></el-input></el-form-item>               
            </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialogCreateVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitCreate">确 定</el-button>
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
                month_id:'201911',
                tenant_id:localStorage.getItem('user_tenant_id'),
                keyword: '',
                order:'',          
                pageSize:6,
                pageIndex:1              
            },
            entity:{
                month_id:'201911',
                tenant_id:'',
                person_name:'',
                person_id:'00d7094d-7d2b-48db-94aa-3c02fc1c0eff',
                amount:0,
                adjest_remark:''
            },
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
            this.$fetch('tenantadjustitem',this.query )
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
            this.entity.person_name="";
            this.entity.amount=0;
            this.entity.adjest_remark="";
           
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
