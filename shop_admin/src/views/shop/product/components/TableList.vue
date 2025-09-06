<template>
  <div class="mb-2">
    <Form layout="inline">
      <Form.Item label="商品名称">
        <Input v-model:value="search.title" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="类目">
        <TreeSelect v-model:value="search.category_id" show-search style="width: 200px"
                  :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }" tree-default-expand-all :tree-data="category_list" />
      </Form.Item>
      <Form.Item label="状态">
        <Select
          style="width: 120px"
          v-model:value="search.state"
          :options="[
            {value: 0, label: '全部'}, 
            {value: 1, label: '上线'}, 
            {value: 2, label: '下线'}, 
            {value: 3, label: '待审核'}, 
            {value: 4, label: '审核不通过'}
           ]"
          :allowClear="true" />
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
   </template>
 </ApiBasicTable>
</template>

<script lang="ts" setup>
 import { ref, h, onMounted } from 'vue';
 import { Button, Popconfirm, Input, Form, Select, Image, Tag, TreeSelect, message } from 'ant-design-vue';
 import { loadingTask } from '/@/utils/helper';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import api from '/@/api/index'

 const emit = defineEmits(['on-add', 'on-edit', 'on-value'])
 const tableRef = ref<any>(null);

 const search = ref<any>({
  title: '',
  state: 0,
  category_id: '0',
  page_no: 1,
  page_size: 30
})
 const apiLoadData = async (values: any): Promise<any> => {
   return await api.shop.product.list(values)
 }

 const category_list = ref<any[]>([])

 const stateMap = {
  '1': '上线',
  '2': '下线',
  '3': '待审核',
  '4': '审核不通过'
 }

 onMounted(async () => {
  let _treeData = (await api.shop.category.tree()) as any
  category_list.value = [
    {
      label: '全部',
      value: '0',
      children: _treeData
    }
  ]
 })

 const columns: any[] = [
  { dataIndex: 'id', title: '编号', key: 'id', width: 88 },
   {
      dataIndex: 'picture',
      title: '封面图',
      width: 80,
      customRender: (e: any) => {
         return h(Image, {
          width: 60,
          height: 30,
          src: e.record.picture,
          fallback: "data:image/png;base64,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"
         })
       }
    },
   { dataIndex: 'title', title: '商品名称', key: 'title' },
   { dataIndex: ['category', 'title'], title: '类目', key: 'category_title', width: 138 },
   { dataIndex: 'price', title: '当前价格', key: 'price', width: 108 },
   { dataIndex: 'original_price', title: '原价', key: 'price', width: 108 },
   { dataIndex: 'cost_price', title: '成本价', key: 'price', width: 108 },
   { dataIndex: 'state', title: '状态', key: 'state', width: 98, customRender: (e: any) => {
    return stateMap[e.record.state] ? h(Tag, {}, () => stateMap[e.record.state]) : '-'
   } },
   { dataIndex: 'putaway_time', title: '上架时间', key: 'pataway_time', width: 168 },
   { dataIndex: 'sort_num', title: '排序', width: 68 },
 ];

 const handleDel = (e: number) => {
   console.debug(e);
   loadingTask(async () => {
    await api.shop.product.del(e)
    tableRef.value && tableRef.value.useReload()
    message.success('删除成功')
   }, {
     msg: '删除中...'
   })
 }

 defineExpose({
   useReload: () => { tableRef.value && tableRef.value.useReload() },
 })

</script>
