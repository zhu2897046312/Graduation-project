<template>
    <div class="mb-2">
    <Form layout="inline">
      <Form.Item label="标题">
        <Input v-model:value="search.title" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="状态">
        <Select
          style="width: 160px"
          v-model:value="search.state"
          :allowClear="true"
          :options="[
            {value: 0, label: '全部'}, 
            {value: 1, label: '开启'}, 
            {value: 2, label: '关闭'}
            ]" 
          />
      </Form.Item>
      <Form.Item label="平台">
        <Select
          style="width: 160px"
          v-model:value="search.platform"
          :allowClear="true"
          :options="[
            {value: 0, label: '全部'}, 
            {value: 1, label: '淘宝'}, 
            ]" 
          />
      </Form.Item>
      <Form.Item label="任务类型">
        <Select
          style="width: 160px"
          v-model:value="search.type"
          :allowClear="true"
          :options="[
            {value: 0, label: '全部'}, 
            {value: 1, label: '关键词'}, 
            {value: 2, label: '详情'}, 
            ]" 
          />
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
       <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record.id) }">
         <Button type="link" danger>删除</Button>
       </Popconfirm>
     </Button.Group>
     
     <EnumLabel v-if="column.dataIndex === 'state'" :dict-value="record.state" dict-code="CtTaskConstant$State"/>
     <EnumLabel v-if="column.dataIndex === 'platform'" :dict-value="record.platform" dict-code="CtTaskConstant$Platform"/>
     <EnumLabel v-if="column.dataIndex === 'type'" :dict-value="record.type" dict-code="CtTaskConstant$Type"/>
   </template>
 </ApiBasicTable>
</template>

<script lang="ts" setup>
 import { ref } from 'vue';
 import { Button, Popconfirm, Form, Input, Select } from 'ant-design-vue';
 import { loadingTask } from '/@/utils/helper';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import api from '/@/api/index'
import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'

 const emit = defineEmits(['on-add', 'on-edit'])
 const tableRef = ref<any>(null);

 const search = ref<any>({
  title: '',
  state: 0,
  platform: 0,
  type: 0,
  page_no: 1,
  page_size: 30
})
 const apiLoadData = async (e: any): Promise<any> => {
   return await api.shop.task.list(e)
 }

 const columns: any[] = [
   { dataIndex: 'title', title: '任务名称', key: 'title', width: 200 },
   { dataIndex: 'platform', title: '目标平台', key: 'platform', width: 90 },
   { dataIndex: 'type', title: '类型', key: 'type', width: 90 },
   { dataIndex: 'state', title: '状态', key: 'state', width: 90 },
   { dataIndex: 'interval', title: '采集间隔（小时）', key: 'interval', width: 90 },
   { dataIndex: 'created_time', title: '创建时间', key: 'created_time', width: 90 },
   { dataIndex: 'last_start_time', title: '最后一次执行时间', key: 'last_start_time', width: 90 },
 ];

 const handleDel = (e: number) => {
   console.debug(e);
   loadingTask(async () => {
    await api.shop.task.del(e)
    tableRef.value && tableRef.value.useReload()
   }, {
     msg: '删除中...'
   })
 }

 defineExpose({
   useReload: () => { tableRef.value && tableRef.value.useReload() },
 })

</script>
