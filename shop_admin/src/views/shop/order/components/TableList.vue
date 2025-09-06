<template>
  <div class="mb-2">
    <Form layout="inline">
      <Form.Item label="订单号">
        <Input v-model:value="search.code" placeholder="订单号检索" :allowClear="true"
          @pressEnter="() => tableRef.useReload()" />
      </Form.Item>
      <Form.Item label="用户昵称">
        <Input v-model:value="search.nickname" placeholder="用户昵称检索" :allowClear="true"
          @pressEnter="() => tableRef.useReload()" />
      </Form.Item>
      <Form.Item label="用户Email">
        <Input v-model:value="search.email" placeholder="用户Email检索" :allowClear="true"
          @pressEnter="() => tableRef.useReload()" />
      </Form.Item>
      <Form.Item label="订单状态">
        <Select style="width: 160px;" v-model:value="search.state" :options="[
          { label: '全部', value: 0 },
          { label: '待付款', value: 1 },
          { label: '待发货', value: 2 },
          { label: '已发货', value: 3 },
          { label: '已完成', value: 4 },
          { label: '已关闭', value: 5 },
          { label: '无效订单', value: 6 },
          { label: '部分退款', value: 7 },
          { label: '全部退款', value: 8 },
        ]" />
      </Form.Item>
      <Form.Item>
        <Button type="primary" @click="() => tableRef.useReload()">查询</Button>
      </Form.Item>
    </Form>
  </div>
  <ApiBasicTable ref="tableRef" :api="apiLoadData" :columns="columns" :action-width="200" :search-param="search"
    :show-id="false" :show-page="true">
    <template #bodyCell="{ column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-info', record.id)">详情</Button>
        <Button type="link" @click="emit('on-delivery', record.id)" :disabled="record.state != 2">发货</Button>
        <Button type="link" @click="emit('on-update', record.id)">修改状态</Button>
        <Button type="link" @click="emit('on-refund', record.id)"
          :disabled="record.state == 1 || record.state == 8">退款</Button>
      </Button.Group>
      <div v-if="column.key === 'items'">
        <div class="product-info" v-for="it in record.items">
          <div class="product-item">
            <span class="label" style="width: 40px;">商品: </span>
            <span class="value">{{ it.title }};</span>
          </div>
          <div v-if="it.sku_title && it.sku_title.length > 0" class="product-item"><span class="label">规格: </span><span
              class="value">{{ it.sku_title }};</span></div>
          <div class="product-item"><span class="label">数量: </span><span class="value">{{ it.quantity }};</span></div>
        </div>
      </div>

      <EnumLabel v-if="column.dataIndex === 'state'" :dict-value="record.state" dict-code="SpOrderConstant$State" />
      <!-- <span v-if="column.dataIndex === 'state'">{{ record.state }}</span> -->
    </template>
  </ApiBasicTable>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { Button, Input, Form, Select } from 'ant-design-vue';
import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
import api from '/@/api/index'
import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'

const emit = defineEmits(['on-info', 'on-delivery', 'on-update'])
const tableRef = ref<any>(null);

const search = ref<any>({
  nickname: '',
  code: '',
  email: '',
  state: 0,
  page_no: 1,
  page_size: 30
})
const apiLoadData = async (params: any): Promise<any> => {
  return await api.shop.order.list(params)
}

const columns: any[] = [
  { dataIndex: 'code', title: '订单编号', key: 'code', width: 130 },
  { dataIndex: 'items', title: '商品', key: 'items', width: 360 },
  { dataIndex: 'state', title: '状态', key: 'state', width: 80 },
  { dataIndex: 'created_time', title: '下单时间', key: 'created_time', width: 130 },
  { dataIndex: 'pay_amount', title: '支付金额', key: 'pay_amount', width: 110 },
  { dataIndex: 'nickname', title: '用户昵称', key: 'nickname', width: 120 },
  { dataIndex: 'email', title: 'email', key: 'email', width: 200 },
];


defineExpose({
  useReload: () => { tableRef.value && tableRef.value.useReload() },
})

</script>

<style lang="css" scoped>
.product-info {
  border-bottom: 1px dashed #c4c4c4;
  padding-bottom: 5px;
  padding-top: 5px;
}

.product-info:first-child {
  padding-top: 0;
}

.product-info:last-child {
  border-bottom: none;
}

.product-item {
  display: flex;
  align-items: center;
  gap: 3px;
}

.product-item .label {
  font-size: 14px;
  color: #666;
}

.product-item .value {
  font-size: 14px;
  color: #333;
}

.search-bar {
  background: #fff;
  padding: 16px;
  border-radius: 2px;
}

.search-bar :deep(.ant-form-item) {
  margin-bottom: 16px;
}

.search-bar :deep(.import-button) {
  float: right;
  margin-right: 0;
}

.table-header {
  margin-bottom: 16px;
}

:deep(.ant-btn-link) {
  padding: 4px 8px;
  height: auto;
}

:deep(.ant-space) {
  display: flex;
  justify-content: center;
}
</style>
