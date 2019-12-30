<el-table :data="tableData"  border size="mini" fit highlight-current-row height="500" 
      :row-style="rowStyle" 
      @row-dblclick="rowDblclick"
      v-loading="loading"
      element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading"
      element-loading-background="rgba(0, 0, 0, 0.3)">
      <el-table-column type="index" align="center" fixed></el-table-column>
      <!-- prop: 字段名name, label: 展示的名称, fixed: 是否需要固定(left, right), minWidth: 设置列的最小宽度(不传默认值), oper: 是否有操作列
           oper.name: 操作列字段名称, oper.clickFun: 操作列点击事件, formatData: 格式化内容 -->
      <el-table-column v-for="(th, key) in tableHeader"
        :key="key"
        :prop="th.prop"
        :label="th.label"
        :fixed="th.fixed"
        :min-width="th.minWidth" align="center">  
      </el-table-column>
</el-table>