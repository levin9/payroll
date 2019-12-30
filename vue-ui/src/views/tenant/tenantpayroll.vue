<template>
  <div>
    <TableList @download="download" 
              :apiService="userService" 
              :grid-config="configData.gridConfig" 
              :grid-btn-config="configData.gridBtnConfig"
              :grid-data="testData"
              :form-config="configData.formConfig" 
              :form-data="configData.formModel" 
              :grid-edit-width="220"
              :titlePath="configData.gridTitlePath"
              @onEdit="onEdit"
              @onSearch="onSearch"
              :showTopLayout="true"
              form-title="用户"
              
              :is-async="true">
    </TableList>
  </div>
</template>

<script>
  import TableList from  '../../components/common/TableList.vue'
  import {uiCfg} from './configs/tenantpayroll.js'
  import {userService} from '@/api/user.js'

  export default {
    components: {
      TableList
    },
    data () {
      return {
        testData: [],
        configData: uiCfg,
        userService: userService
      }
    },
    name: 'user-list',
    mounted () {
      this.testData = [
        {
          id: '1',
          tel: '15184318420',
          name: '小白',
          email: '412412@qq.com',
          status: '1',
          create_time: '2018-04-20',
          expand: '扩展信息一',
          role: ['2']
        },
             {
          id: '1',
          tel: '15184318420',
          name: '小白',
          email: '412412@qq.com',
          status: '1',
          create_time: '2018-04-20',
          expand: '扩展信息一',
          role: ['2']
        }
      ];
    },
    methods: {
      download(row){
        console.log('点击了下载按钮 parent',row);
        alert('点击了下载按钮.formindex');
      },   
      onEdit(item){
        console.log(item);
        alert('son alert .');
      },
      onSearch(keyword){
        alert(keyword);
      },
      view(data){
        alert('son page data');
      }
    }
  }
</script>