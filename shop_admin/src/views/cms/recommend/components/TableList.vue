<template>
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
      <Button type="link" @click="emit('on-value', record.id)">内容管理</Button>
       <Button type="link" @click="emit('on-edit', record.id)">编辑</Button>
       <!-- <Button type="link" @click="emit('on-cont', record.id)">编辑内容</Button> -->
       <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record.id) }">
         <Button type="link" danger>删除</Button>
       </Popconfirm>
     </Button.Group>
     
     <EnumLabel v-if="column.dataIndex === 'state'" :dict-value="record.state" dict-code="CmsRecommendConstant$Status"/>
   </template>
 </ApiBasicTable>
</template>

<script lang="ts" setup>
 import { ref } from 'vue';
 import { Button, Popconfirm } from 'ant-design-vue';
 import { loadingTask } from '/@/utils/helper';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import api from '/@/api/index'
import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'

 const emit = defineEmits(['on-add', 'on-edit', 'on-value'])
 const tableRef = ref<any>(null);

 const search = ref<any>({
  page_no: 1,
  page_size: 30
})
 const apiLoadData = async (e: any): Promise<any> => {
   return await api.cms.recommend.curdApi("LIST", e)
 }

 const columns: any[] = [
   { dataIndex: 'title', title: '推荐位名称', key: 'title', width: 200 },
   { dataIndex: 'code', title: '编码', key: 'title', width: 160 },
   { dataIndex: 'state', title: '状态', key: 'title', width: 90 },
 ];

 const handleDel = (e: number) => {
   console.debug(e);
   loadingTask(async () => {
    await api.cms.recommend.curdApi('DELETE', e)
    tableRef.value && tableRef.value.useReload()
   }, {
     msg: '删除中...'
   })
 }

 defineExpose({
   useReload: () => { tableRef.value && tableRef.value.useReload() },
 })

</script>
