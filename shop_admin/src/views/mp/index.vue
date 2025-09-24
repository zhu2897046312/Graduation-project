<template>
    <PageLayout>
        <Card  title="用户列表">
            <TableList ref="tableRef" 
            @on-edit="(e: number) => { router.push({ name: 'ShopCoreProductModify', query: { id: e } }); }"
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
import { useRouter } from 'vue-router';

import { ref } from 'vue';
const router = useRouter();
const refundRef = ref<any>(null);
const infoRef = ref<any>(null);
const deliveryRef = ref<any>(null);
const stateRef = ref<any>(null);
const tableRef = ref<any>(null);
</script>