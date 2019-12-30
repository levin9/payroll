<template>
  <div>
    <editForm @download="download" 
              :apiService="userService" 
              :grid-config="configData.gridConfig" 
              :grid-btn-config="configData.gridBtnConfig"
              :grid-data="testData"
              :form-config="configData.formConfig" 
              :form-data="configData.formModel" 
              :grid-edit-width="220"
              @onEdit="onEdit"
              form-title="用户" 
              :is-async="true">
    </editForm>
  </div>
</template>

<script>
  import editForm from  '../../components/common/TableEdit.vue'
  import {USER_CONFIG} from './configs/user_config'
  import {userService} from '@/api/user.js'

  export default {
    components: {
      editForm
    },
    data () {
      return {
        testData: [],
        formModel: {
                      id: '',
                      tel: '',
                      name: '',
                      email: '',
                      status: '',
                      create_time: '',
                      expand: '',
                      roleIdList: []
                    },
        configData: USER_CONFIG,
        userService: userService
      }
    },
    name: 'user-list',
    mounted () {
     
    },
    methods: {
      
    }
  }
</script>