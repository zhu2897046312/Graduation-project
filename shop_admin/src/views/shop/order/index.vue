<template>
  <PageLayout>
    <Card  title="订单列表">
      <TableList ref="tableRef" 
        @on-info="(e: number) => { infoRef.useOpen({ id: e }); }"
        @on-delivery="(e: number) => { deliveryRef.useOpen({ id: e }); }"
        @on-update="(e: number) => { stateRef.useOpen({ id: e }); }"
        @on-refund="(e: number) => { refundRef.useOpen({ id: e }); }"
      />
    </Card>
  </PageLayout>
  <Delivery ref="deliveryRef" @on-change="() => { tableRef.useReload() }" />
  <UpdateState ref="stateRef" @on-change="() => { tableRef.useReload() }" />
  <RefundReq ref="refundRef" @on-change="() => { tableRef.useReload() }" />
  <Info ref="infoRef" />
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TableList from './components/TableList.vue';
  import Info from './components/Info.vue';
  import Delivery from './components/Delivery.vue';
  import UpdateState from './components/UpdateState.vue';
  import RefundReq from './components/RefundReq.vue';
  
  import { ref } from 'vue';

  const refundRef = ref<any>(null);
  const infoRef = ref<any>(null);
  const deliveryRef = ref<any>(null);
  const stateRef = ref<any>(null);
  const tableRef = ref<any>(null);


</script>