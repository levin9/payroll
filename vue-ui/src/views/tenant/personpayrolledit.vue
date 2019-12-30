<template>
  <dynamic-form  v-model="data"  :descriptors="descriptors" aria-colspan="8" aria-rowspan="8" >  
      <template slot="operations">
      <el-button @click="clearValidate">clearValidate</el-button>
      <el-button type="primary" @click="resetFields" plain>resetFields</el-button>
      <el-button type="primary" @click="validate">validate</el-button>
    </template>
  </dynamic-form>
</template>

<script>
export default {
  data () {
    return {
      descriptors: {
        date: { type: 'date', label: 'date \'s label', required: false },
        number: { type: 'number', label: '基本工资', required: true, placeholder: 'please input the number' },
        string: { type: 'string', label: '住房公积金', required: true, pattern: /^test$/g },
        url: { type: 'url', label: '社保', required: true, placeholder: 'please input the url' },
        email: { type: 'email', label: 'email \'s label', required: false },
        enum: {
          type: 'enum',
          label: 'enum\'s label',
          enum: ['value-1', 'value-2'],
          options: [
            { label: 'option-1', value: 'value-1', disabled: true },
            { label: 'option-2', value: 'value-2' }
          ]
        }
      },
      data: {}
    }
  },
  methods: {
    clearValidate () {
      this.$refs['dynamic-form'].clearValidate()
    },
    resetFields () {
      this.$refs['dynamic-form'].resetFields()
    },
    async validate () {
      const valid = await this.$refs['dynamic-form'].validate()
      window.alert(`valid result ====> ${valid}`)
    }
  }
}
</script>