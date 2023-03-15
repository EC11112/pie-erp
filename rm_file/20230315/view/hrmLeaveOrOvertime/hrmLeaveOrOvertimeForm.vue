<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="left" :rules="rule" label-width="120px">
        <el-form-item label="用户ID:" prop="user_id">
          <el-input v-model.number="formData.user_id" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否为请假:" prop="is_leave">
          <el-switch v-model="formData.is_leave" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="起始日期:" prop="start_date">
          <el-date-picker v-model="formData.start_date" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="终止日期:" prop="end_date">
          <el-date-picker v-model="formData.end_date" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
        </el-form-item>
        <el-form-item label="是否批准:" prop="is_approved">
          <el-switch v-model="formData.is_approved" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="说明:" prop="explanation">
          <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 4}" v-model="formData.explanation" :clearable="true" placeholder="请输入" />
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
  name: 'HrmLeaveOrOvertime'
}
</script>

<script setup>
import {
  createHrmLeaveOrOvertime,
  updateHrmLeaveOrOvertime,
  findHrmLeaveOrOvertime
} from '@/api/hrmLeaveOrOvertime'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            user_id: 0,
            is_leave: false,
            start_date: new Date(),
            end_date: new Date(),
            is_approved: false,
            explanation: '',
        })
// 验证规则
const rule = reactive({
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               is_leave : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               start_date : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               end_date : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               is_approved : [{
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
      const res = await findHrmLeaveOrOvertime({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehrmLeaveOrOvertime
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
               res = await createHrmLeaveOrOvertime(formData.value)
               break
             case 'update':
               res = await updateHrmLeaveOrOvertime(formData.value)
               break
             default:
               res = await createHrmLeaveOrOvertime(formData.value)
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
