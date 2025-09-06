<template>
  <PageLayout>
    <Card title="退款列表">
      <TableList ref="tableRef" 
        @on-info="(e: number) => { infoRef.useOpen({ id: e }); }"
        @on-order="(e: string) => { orderRef.useOpen({ orderCode: e }); }"
      />
    </Card>
  </PageLayout>
  <Info ref="infoRef" />
  <OrderInfo ref="orderRef" />
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TableList from './components/TableList.vue';
  import Info from './components/Info.vue';
  import OrderInfo from './components/OrderInfo.vue';
  import { ref } from 'vue';

  const infoRef = ref<any>(null);
  const orderRef = ref<any>(null);
  const tableRef = ref<any>(null);
</script>