<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item><i class="el-icon-lx-home"></i> 首页</el-breadcrumb-item>
                <el-breadcrumb-item> 薪酬管理</el-breadcrumb-item>
                <el-breadcrumb-item> 考勤管理</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="handle-box" style="text-align:center; margin-top:-20px; margin-bottom:40px;" >
                 <el-col  :span="6">
                    <el-input placeholder="请输入内容" v-model="query.name"  class="input-with-select">                 
                    <el-button slot="append" icon="el-icon-search" ></el-button>
                    </el-input>  
                 </el-col>
                  <el-col  :span="12">
                      &nbsp;
                 </el-col>
                  <el-col  :span="6">
                       <!-- <el-button type="primary" icon="el-icon-plus" class="handle-del mr10" @click="handleInsert" >新增</el-button>
                            -->
                        <el-button-group>
                            <!-- <el-button icon="el-icon-plus" @click="handleInsert" >新增</el-button>                 -->
                            <el-button  icon="el-icon-upload" @click="handleImport" >导入</el-button>
                            <el-button  icon="el-icon-plus" @click="handleInsert" >导出</el-button>
                            

                        </el-button-group>

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
               <el-table-column prop="real_name" label="姓名"></el-table-column> 
               <!-- <el-table-column prop="month_id" label="月份"></el-table-column>       -->
               <el-table-column label="迟到"  >     
                    <el-table-column prop="chi_dao_count" label="次数" width="50" ></el-table-column>  
                    <el-table-column prop="chi_dao_num" label="时长" width="50"></el-table-column>  
                </el-table-column>   
                <el-table-column label="早退">   
                    <el-table-column prop="zao_tui_num" label="次数" width="50"></el-table-column>  
                    <el-table-column prop="zao_tui_num" label="时长" width="50"></el-table-column> 
                 </el-table-column>  
                <el-table-column label="常见假期">   
                    <el-table-column prop="tiao_xiu_num" label="调休" width="50"></el-table-column>  
                    <el-table-column prop="shi_jia_num" label="事假" width="50"></el-table-column>  
                    <el-table-column prop="sick_num" label="病假" width="50"></el-table-column>  
                    <el-table-column prop="dai_xin_shi_jia_num" label="带薪假" width="65"></el-table-column>
                </el-table-column>  
                <el-table-column prop="kuang_gong_num" label="考勤空白" width="80" v-if="show" ></el-table-column>
                <el-table-column prop="kong_bai_num" label="旷工" width="50"></el-table-column>
                <el-table-column prop="bu_rv_jia_num" label="哺乳假" width="65" v-if="show" ></el-table-column> 
                <el-table-column prop="pei_chan_jia_num" label="陪产假" width="65"></el-table-column>
                <el-table-column prop="chan_jia_num" label="产假" width="50"></el-table-column>
                <el-table-column prop="hun_jia_num" label="婚假" width="50"></el-table-column>
                <el-table-column prop="sang_jia_num" label="丧假" width="50"></el-table-column>                      
                <el-table-column label="操作" width="70" align="center">
                    <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(scope.$index, scope.row)"
                        >编辑</el-button>                     
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

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="30%">
            <el-form ref="form" :model="form" label-width="70px">
                <el-form-item label="用户名">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model="form.address"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>
        <!-- 导入 -->
        <el-dialog title="导入" :visible.sync="uploadVisible" width="40%">
        <div style="margin-top:-20px;" >
            <el-upload
                class="upload-demo"
                drag
                action="http://127.0.0.1:5555/upload"
                :http-request="uploadFile"
                accept=".xls,.xlsx" 
                :limit="1" 
                :auto-upload="true"
                multiple>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
                </el-upload>
        </div>
            <span slot="footer" class="dialog-footer">               
                <el-button @click="uploadVisible = false">取 消</el-button>
                <el-button type="primary" @click="importData">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: 'basetable',
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
            tableData: [],
            multipleSelection: [],
            delList: [],
            editVisible: false,
            uploadVisible:false,
            uploadFilePath:'',
            pageTotal: 0,
            form: {},
            idx: -1,
            id: -1
        };
    },
     created() {
        //this.getData();
    },
    mounted(){  
        this.getData();         
    },
    methods: {        
         getData() { 
            this.$fetch('attendance',this.query )
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
            this.$router.push({path:'/personaledit'})
        },
         // 修改行操作
        handleEdit(index, row) {
            this.$router.push({path:'/attendanceedit',query: {id: row.id}})             
        },
        // 删除操作
        handleDelete(index, row) {
            // 二次确认删除
            this.$confirm('确定要删除吗？', '提示', {
                type: 'warning'
            }).then(() => {
                this.$del('personal/'+row.person_id)
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
            }).catch(() => {});
        },
        uploadFile(param) {       
            var that = this;
          var fileObj = param.file;
          console.log(fileObj);
          this.file = fileObj;
          var FileController = 'http://127.0.0.1:5555/api/v1/upload'   ; // 接收上传文件的后台地址
          var form = new FormData();    // FormData 对象
          form.append('file', fileObj);    // 文件对象
          form.append('biztype', 'temp_import');    // 其他参数
          var xhr = new XMLHttpRequest();    // XMLHttpRequest 对象
          xhr.open('post', FileController, true);
          xhr.send(form);          
          //若响应完成且请求成功
          xhr.onreadystatechange = function(){
            if(xhr.readyState === 4 && xhr.status === 200){
                //do something, e.g. request.responseText
                var result = JSON.parse(xhr.responseText);                
                that.uploadFilePath=result.path;
            }
          }
         
        },
        handleImport(){
            this.uploadVisible=true;
            this.editVisible=false;            
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
        // 编辑操作
        //handleEdit(index, row) {
        //    this.idx = index;
        //    this.form = row;
        //    this.editVisible = true;
        //},
        // 保存编辑
        importData() {
            var that = this;        
            var filePara = {
                                file_path:this.uploadFilePath,
                                biz_type:'temp_import',
                                tenant_id:localStorage.getItem('user_tenant_id')
                            };            
            this.$post('attendanceimport',filePara )
            .then((response) => {
                console.log(response);
                this.uploadVisible = false;
                this.$message.success(`上传成功`);
                this.$set(this.tableData, this.idx, this.form);
               
            }).catch(function(error){
                console.log(error);
                this.$message.error(error);                
            }); 

            
        },
        beforeUpload(file) {
            //判断文件格式
            const isXlsx = file.type ==='application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
            const isXls = file.type === 'application/vnd.ms-excel'
            const isLt2M = file.size / 1024 / 1024 < 2
            console.log(isXls, isXlsx)
            if (!isXls && !isXlsx) {
                this.$message.warning('上传文件必须是xls/xlsx格式!')
            }
            if (!isLt2M) {
                this.$message.warning('上传图片大小不能超过 2MB!')               
            }
            console.log('error', (isXlsx || isXls) && isLt2M)
            return (isXlsx || isXls)
            
        },
        // 上传文件个数超过定义的数量
        handleExceed(files, fileList) {
            this.$message.warning(`当前限制选择 1 个文件，请删除后继续上传`)
        },
        // 分页导航
        handlePageChange(val) {
            this.$set(this.query, 'pageIndex', val);
            this.getData();
        },
         dateFormat:function(row,column){
        var t=new Date(row.join_date);//row 表示一行数据, updateTime 表示要格式化的字段名称   
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
