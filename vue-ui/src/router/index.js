import Vue from 'vue';
import Router from 'vue-router';
import DynamicForm from 'vue-dynamic-form-component';

Vue.use(Router);
Vue.use(DynamicForm);

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            redirect: '/dashboard'
        },
        {
            path: '/',
            component: () => import(/* webpackChunkName: "home" */ '../components/common/Home.vue'),
            meta: { title: '自述文件' },
            children: [
                {
                    path: '/dashboard',
                    component: () => import(/* webpackChunkName: "dashboard" */ '../components/page/Dashboard.vue'),
                    meta: { title: '系统首页' }
                },
                {
                    path: '/icon',
                    component: () => import(/* webpackChunkName: "icon" */ '../components/page/Icon.vue'),
                    meta: { title: '自定义图标' }
                },
                {
                    path: '/companyedit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/companyedit.vue'),
                    meta: { title: '公司设置' }
                },
                {
                    path: '/personal',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/personal.vue'),
                    meta: { title: '员工管理' }
                },
                {
                    path: '/basemonth',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/basemonth.vue'),
                    meta: { title: '月份设置' }
                },
                {
                    path: '/adjuest',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/adjuest.vue'),
                    meta: { title: '手工调整工资' }
                },
                {
                    path: '/attendance',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/attendance.vue'),
                    meta: { title: '考勤管理' }
                },
                {
                    path: '/attendanceedit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/attendanceedit.vue'),
                    meta: { title: '考勤编辑' }
                },
                // {
                //     path: '/analysis',
                //     component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/analysis.vue'),
                //     meta: { title: '薪酬计算' }
                // },
                {
                    path: '/calculate',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/calculate.vue'),
                    meta: { title: '薪酬分析2' }
                },
                {
                    path: '/personaledit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/personaledit.vue'),
                    meta: { title: '修改员工' }
                },
                {
                    path: '/personpayrolledit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/personpayrolledit.vue'),
                    meta: { title: '修改员工' }
                },
                {
                    path: '/payrolledit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/payrolledit.vue'),
                    meta: { title: '修改员工2' }
                },
                {
                    path: '/tenantpayrolledit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/tenantpayrolledit.vue'),
                    meta: { title: '薪酬项维护' }
                },
                {
                    path: '/tenantattedit',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/tenantattedit.vue'),
                    meta: { title: '考勤项维护' }
                },
                {
                    path: '/curdtest',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tenant/curdtest.vue'),
                    meta: { title: 'test' }
                },
                {
                    path: '/formindex',
                    component: () => import(/* webpackChunkName: "icon" */ '../views/tests/formindex.vue'),
                    meta: { title: 'formindex' }
                },
                {
                    path: '/table',
                    component: () => import(/* webpackChunkName: "table" */ '../components/page/BaseTable.vue'),
                    meta: { title: '基础表格' }
                },
                {
                    path: '/tabs',
                    component: () => import(/* webpackChunkName: "tabs" */ '../components/page/Tabs.vue'),
                    meta: { title: 'tab选项卡' }
                },
                {
                    path: '/form',
                    component: () => import(/* webpackChunkName: "form" */ '../components/page/BaseForm.vue'),
                    meta: { title: '基本表单' }
                },
                {
                    // 富文本编辑器组件
                    path: '/editor',
                    component: () => import(/* webpackChunkName: "editor" */ '../components/page/VueEditor.vue'),
                    meta: { title: '富文本编辑器' }
                },
                {
                    // markdown组件
                    path: '/markdown',
                    component: () => import(/* webpackChunkName: "markdown" */ '../components/page/Markdown.vue'),
                    meta: { title: 'markdown编辑器' }
                },
                {
                    // 图片上传组件
                    path: '/upload',
                    component: () => import(/* webpackChunkName: "upload" */ '../components/page/Upload.vue'),
                    meta: { title: '文件上传' }
                },
                {
                    // vue-schart组件
                    path: '/charts',
                    component: () => import(/* webpackChunkName: "chart" */ '../components/page/BaseCharts.vue'),
                    meta: { title: 'schart图表' }
                },
                {
                    // 拖拽列表组件
                    path: '/drag',
                    component: () => import(/* webpackChunkName: "drag" */ '../components/page/DragList.vue'),
                    meta: { title: '拖拽列表' }
                },
                {
                    // 拖拽Dialog组件
                    path: '/dialog',
                    component: () => import(/* webpackChunkName: "dragdialog" */ '../components/page/DragDialog.vue'),
                    meta: { title: '拖拽弹框' }
                },
                {
                    // 国际化组件
                    path: '/i18n',
                    component: () => import(/* webpackChunkName: "i18n" */ '../components/page/I18n.vue'),
                    meta: { title: '国际化' }
                },
                {
                    // 权限页面
                    path: '/permission',
                    component: () => import(/* webpackChunkName: "permission" */ '../components/page/Permission.vue'),
                    meta: { title: '权限测试', permission: true }
                },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../components/page/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../components/page/403.vue'),
                    meta: { title: '403' }
                },
                {
                    path: '/donate',
                    component: () => import(/* webpackChunkName: "donate" */ '../components/page/Donate.vue'),
                    meta: { title: '支持作者' }
                }
            ]
        },
        {
            path: '/login',
            component: () => import(/* webpackChunkName: "login" */ '../components/page/Login.vue'),
            meta: { title: '登录' }
        },
        {
            path: '*',
            redirect: '/404'
        }
    ]
});
