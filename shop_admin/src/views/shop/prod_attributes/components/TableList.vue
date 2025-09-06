<template>
  <div class="mb-2">
    <Form layout="inline">
      <Form.Item label="名称">
        <Input v-model:value="search.title" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item>
        <Button type="primary" @click="() => tableRef.useReload()">查询</Button>
      </Form.Item>
    </Form>
  </div>
  <ApiBasicTable
   ref="tableRef"
   :api="apiLoadData"
   :columns="columns"
   :action-width="200"
   :search-param="search"
   :show-id="false"
   :show-page="true">
   <template #bodyCell="{column, record }">
     <Button.Group v-if="column.key == '_action'">
       <Button type="link" @click="emit('on-edit', record.id)">编辑</Button>
       <Button type="link" @click="emit('on-value', record.id)">属性值</Button>
       <!-- <Button type="link" @click="emit('on-cont', record.id)">编辑内容</Button> -->
       <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record.id) }">
         <Button type="link" danger>删除</Button>
       </Popconfirm>
     </Button.Group>
   </template>
 </ApiBasicTable>
</template>

<script lang="ts" setup>
 import { ref } from 'vue';
 import { Button, Popconfirm, Input, Form } from 'ant-design-vue';
 import { loadingTask } from '/@/utils/helper';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import api from '/@/api/index'

 const emit = defineEmits(['on-add', 'on-edit', 'on-value'])
 const tableRef = ref<any>(null);

 const search = ref<any>({
  title: '',
  page_no: 1,
  page_size: 30
})
const apiLoadData = async (params: any): Promise<any> => {
   return await api.shop.prodAttributes.list(params)
 }

 const columns: any[] = [
   { dataIndex: 'title', title: '属性名称', key: 'title', width: 260 },
   { dataIndex: 'sort_num', title: '排序', width: 120 },
 ];

 const handleDel = (e: number) => {
   console.debug(e);
   loadingTask(async () => {
    await api.shop.prodAttributes.del(e)
    tableRef.value && tableRef.value.useReload()
   }, {
     msg: '删除中...'
   })
 }

 defineExpose({
   useReload: () => { tableRef.value && tableRef.value.useReload() },
 })

</script>
