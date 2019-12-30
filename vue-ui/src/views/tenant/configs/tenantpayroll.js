export const USER_CONFIG = {
    gridTitlePath:['测试木块','列表页'],
    gridConfig: [
      {label: '用户ID', prop: 'id', width: '100'},
      //{label: '手机号(登录账号)', prop: 'tel'},
      {label: '邮箱', prop: 'email', },
      {label: '中文名', prop: 'name', width: '100'},
      {
        label: '状态', prop: 'status', width: '100', render: (h, params) => {
          if(params.row.status === '0'){
            return h('el-tag', {
              props:{
                size:'mini',
                type:'warning'
              }
            },'正常');
          }else {
            return h('el-tag', {
              props:{
                size:'mini',
                type:'success'
              }
            },'禁用');
          }
        }
      },
      {label: '创建时间', prop: 'create_time', width: '100'},
      {label: '扩展信息', prop: 'expand'}
    ],
    // crud的模态框表单配置，可配置表单类型，验证规则，是否必填,col-span布局可通过span参数配置
    formConfig: [
      {span: 12, label: '手机号', prop: 'tel', type: 'text'},
      {span: 12, label: '中文名', prop: 'name', type: 'text'},
      {span: 12, label: '邮箱', prop: 'email', type: 'text'},
      {
        span: 12, label: '角色',
        prop: 'roleIdList',
        type: 'checkbox',
        data: [{name: '角色一', id: '1'}, {name: '角色二', id: '2'}],
        rules: { type: 'array', required: true, message: '请选择角色', trigger: 'change' }
      },
      {
        span: 12, label: '状态',
        prop: 'status',
        type: 'radio',
        data: [{name: '正常', id: 1}, {name: '禁用', id: 0}],
        rules: {required: true, message: '请选择角色状态', trigger: 'change'}
      },
      {span: 24, label: '备注', prop: 'expand', type: 'textarea'}
    ],
    // crud的操作按钮配置，基础按钮有添加、修改、删除、查看，还可以配置扩展按钮
    gridBtnConfig: {
      create: true, update: true, delete: true, view: false,search: true,
      expands: [
        { name: '下载', emitName: 'download', type: 'text' }
      ]
    }, 
    gridEditWidth:400,   
    searchModel:{
      keyword:''
    }
  };