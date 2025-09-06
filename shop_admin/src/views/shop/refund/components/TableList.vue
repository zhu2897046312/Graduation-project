<template>
  <div class="mb-2">
    <Form layout="inline">
      <Form.Item label="订单编号">
        <Input v-model:value="search.order_code" placeholder="订单编号检索" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="退款单号">
        <Input v-model:value="search.refund_no" placeholder="退款单号检索" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="用户姓名">
        <Input v-model:value="search.name" placeholder="用户姓名检索" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="用户Email">
        <Input v-model:value="search.email" placeholder="用户Email检索" :allowClear="true" @pressEnter="() => tableRef.useReload()"/>
      </Form.Item>
      <Form.Item label="退款状态">
        <Select
          style="width: 160px;"
          v-model:value="search.status"
          :options="[
            { label: '全部', value: null },
            { label: '处理中', value: 2 },
            { label: '已拒绝', value: 4 },
            { label: '已退款', value: 3 },
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
       <Button type="link" @click="emit('on-info', record.id)">退款详情</Button>
       <Button type="link" @click="emit('on-order', record.order_code)">订单详情</Button>
     </Button.Group>
     <template v-if="column.dataIndex === 'status'">  
      <Tag :color="record.status === 2 ? 'success' : ''">
            {{ record.status === 2 ? '处理中' : record.status === 4 ? '已拒绝' : '已退款' }}
          </Tag>
     </template>
   </template>
 </ApiBasicTable>
 
</template>

<script lang="ts" setup>
 import { ref } from 'vue';
 import { Button, Input, Form, Select, Tag } from 'ant-design-vue';
 import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
 import api from '/@/api/shop/refund'
 import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'

 const emit = defineEmits(['on-info', 'on-order'])
 const tableRef = ref<any>(null);

const search = ref<any>({
  order_code: '',
  refund_no: '',
  status: null,
  name: '',
  email: '',
  page_no: 1,
  page_size: 30
})
const apiLoadData = async (params: any): Promise<any> => {
  return await api.list(params)
}

const columns: any[] = [
  { dataIndex: 'refund_no', title: '退款单号', key: 'refund_no', width: 150 },
  { dataIndex: 'order_code', title: '订单编号', key: 'order_code', width: 130 },
  { dataIndex: 'status', title: '状态', key: 'status', width: 80 },
  { dataIndex: 'refund_amount', title: '退款金额', key: 'refund_amount', width: 100 },
  { dataIndex: 'refund_time', title: '退款时间', key: 'refund_time', width: 130 },
  { dataIndex: 'name', title: '用户姓名', key: 'name', width: 120 },
  { dataIndex: 'email', title: '用户Email', key: 'email', width: 200 },
];

defineExpose({
  useReload: () => { tableRef.value && tableRef.value.useReload() },
})
</script>