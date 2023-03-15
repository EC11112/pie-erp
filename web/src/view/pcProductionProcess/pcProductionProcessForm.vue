<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="流程名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属项目:" prop="project_id">
          <el-input v-model.number="formData.project_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车间主任:" prop="user_id">
          <el-input v-model.number="formData.user_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="生产周期:" prop="production_cycle">
          <el-date-picker v-model="formData.production_cycle" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="说明:" prop="explanation">
          <el-input v-model="formData.explanation" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PcProductionProcess'
}
</script>

<script setup>
import {
  createPcProductionProcess,
  updatePcProductionProcess,
  findPcProductionProcess
} from '@/api/pcProductionProcess'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            project_id: 0,
            user_id: 0,
            production_cycle: new Date(),
            explanation: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               project_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               production_cycle : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPcProductionProcess({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repcProductionProcess
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPcProductionProcess(formData.value)
               break
             case 'update':
               res = await updatePcProductionProcess(formData.value)
               break
             default:
               res = await createPcProductionProcess(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
