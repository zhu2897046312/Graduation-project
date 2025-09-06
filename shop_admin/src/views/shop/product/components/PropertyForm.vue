<template>
  <div>
    <div class="mb-2">
      <Button.Group>
        <Button type="link" @click="handleAddLine">新增一行</Button>
        <Popconfirm :disabled="selectedRowKeys.length == 0" title="是否删除选定行?" @confirm="handleDelLine">
          <Button :disabled="selectedRowKeys.length == 0" type="link">删除选定行</Button>
        </Popconfirm>
        <Popconfirm :disabled="list.length == 0" title="是否清空?" @confirm="handleClearLine">
          <Button :disabled="list.length == 0" type="link">清空</Button>
        </Popconfirm>
      </Button.Group>
    </div>
    <Table size="small" :pagination="false" :columns="columns" bordered row-key="_key"
      :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: handleSelectChange }" :dataSource="list">
      <template #bodyCell="{ column, index }: { column: any, index: number, text: any }">
        <Input v-if="column.key === 'title'" v-model:value="list[index]['title']" />
        <Input v-if="column.key === 'value'" v-model:value="list[index]['value']" />
        <InputNumber :min="0" :max="999999" :precision="0" v-if="column.key === 'sort_num'" v-model:value="list[index]['sort_num']" />
      </template>
    </Table>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Table, Button, InputNumber, Popconfirm, message, Input, Form } from 'ant-design-vue';

const porps = defineProps<{ value: any[] }>()

const emit = defineEmits(['update:value', 'change'])

const columns: any[] = [
  { dataIndex: 'title', key: 'title', title: '属性名', width: 160,  },
  { dataIndex: 'value', key: 'value', title: '属性值', width: 220, },
  { dataIndex: 'sort_num', key: 'sort_num', title: '排序', width: 80, },
]


const list = computed<any[]>({
  get: () => porps.value,
  set: (e: any[]) => {
    emit('update:value', e)
    emit('change', e)
    formItemContext && formItemContext.onFieldChange()
  }
})


let formItemContext: any = null
onMounted(async () => {
  formItemContext = Form.useInjectFormItemContext();
})

const handleAddLine = () => {
  list.value.push({
    title: '',
    value: '',
    sort_num: 100,
    _key: Math.random().toString(36).substring(3, 8)
  })
}

const selectedRowKeys = ref<string[]>([])
const handleSelectChange = (e: any[]) => {
  selectedRowKeys.value = e
}

const handleDelLine = () => {
  if (selectedRowKeys.value.length == 0) {
    message.warning('请选择要删除的行')
    return
  }
  list.value = list.value.filter(e => !selectedRowKeys.value.includes(e._key))
  selectedRowKeys.value = []
}

const handleClearLine = () => {
  list.value = []
}



</script>