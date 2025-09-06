<template>
  <ApiBasicTable
   ref="tableRef"
   :api="apiLoadData"
   :columns="columns"
   :action-width="200"
   :show-id="false"
   :show-page="false">
   <template #bodyCell="{column, record }">
     <div v-if="column.key == 'title'">
      <p :style="{ textIndent: `${record._level * 2}em` }">{{ record.title }}</p>
     </div>
     <Button.Group v-if="column.key == '_action'">
       <Button type="link" @click="emit('on-edit', record.id)">编辑</Button>
       <Button type="link" @click="emit('on-add', record.id)">添加子类目</Button>
       <!-- <Button type="link" @click="emit('on-cont', record.id)">编辑内容</Button> -->
       <!-- <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record.id) }"><Button type="link" danger>删除</Button></Popconfirm> -->
     </Button.Group>
     <p v-if="column.key == 'label'" :style="{ textIndent: `${record.level * 2}em` }">
       {{ record.label }}
     </p>
     <EnumLabel v-if="column.dataIndex === 'state'" :dict-value="record.state" dict-code="SpCategoryConstant$Status"/>
   </template>
 </ApiBasicTable>
</template>

<script lang="ts" setup>
 import { h, ref } from 'vue';
 import {  Button, Image } from 'ant-design-vue';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import { treeToLine } from '/@/utils/tree_util'
 import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'
 import api from '/@/api/index'

 const emit = defineEmits(['on-add', 'on-edit', 'on-cont'])
 const tableRef = ref<any>(null);

 const apiLoadData = async (): Promise<any> => {
   let res = await api.shop.category.tree() as any
   let result = treeToLine(res)
   return {
     list: result,
     total: result.length
   }
 }

 const columns: any[] = [
  {
    dataIndex: 'picture',
    title: '封面图',
    width: 90,
    customRender: (e: any) => {
      
       return h(Image, {
        width: 80,
        height: 45,
         src: e.record.picture,
        fallback:'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI4MCIgaGVpZ2h0PSI0NSIgdmlld0JveD0iMCAwIDgwIDQ1Ij48cmVjdCB3aWR0aD0iODAiIGhlaWdodD0iNDUiIGZpbGw9IiNlZWVlZWUiLz48dGV4dCB4PSI1MCUiIHk9IjUwJSIgZm9udC1mYW1pbHk9IkFyaWFsIiBmb250LXNpemU9IjEwIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSIgZmlsbD0iIzk5OSI+Tk9ORTwvdGV4dD48L3N2Zz4='
       })
     }
  },
   { dataIndex: 'title', title: '类目', key: 'title', width: 260 },
   { dataIndex: 'code', title: '编码', key: 'code', width: 260 },
   { dataIndex: 'state', title: '状态', key: 'state', width: 160 },
   { dataIndex: 'sort_num', title: '排序', width: 120 },
 ];

//  const handleDel = (e: number) => {
//    console.debug(e);
//    loadingTask(async () => {
//      message.info('暂未开放删除接口');
//    }, {
//      msg: '删除中...'
//    })
//  }

 defineExpose({
   useReload: () => { tableRef.value && tableRef.value.useReload() },
 })

</script>
