<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="仓库名称:" prop="warehouse_id">
          <el-input v-model.number="formData.warehouse_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="变更类型:" prop="change_type">
        <el-select v-model="formData.change_type" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['出库','入库']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="物品名称:" prop="item_id">
          <el-input v-model.number="formData.item_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="物品数量:" prop="item_number">
          <el-input v-model.number="formData.item_number" :clearable="true" placeholder="请输入" />
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
  name: 'IcInventoryChanges'
}
</script>

<script setup>
import {
  createIcInventoryChanges,
  updateIcInventoryChanges,
  findIcInventoryChanges
} from '@/api/icInventoryChanges'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            warehouse_id: 0,
            item_id: 0,
            item_number: 0,
        })
// 验证规则
const rule = reactive({
               warehouse_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               change_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               item_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               item_number : [{
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
      const res = await findIcInventoryChanges({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reicInventoryChanges
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
               res = await createIcInventoryChanges(formData.value)
               break
             case 'update':
               res = await updateIcInventoryChanges(formData.value)
               break
             default:
               res = await createIcInventoryChanges(formData.value)
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
